package oncall

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type httpClient struct {
	apiKey     string
	baseURL    string
	timeout    time.Duration
	maxRetries int
	backoffMs  int
	client     *http.Client
}

func newHTTPClient(cfg *Config) *httpClient {
	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = "https://api.oncall.sh/v0"
	}

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	maxRetries := cfg.MaxRetries
	if maxRetries == 0 {
		maxRetries = 2
	}

	backoffMs := cfg.BackoffMs
	if backoffMs == 0 {
		backoffMs = 300
	}

	return &httpClient{
		apiKey:     cfg.APIKey,
		baseURL:    strings.TrimSuffix(baseURL, "/"),
		timeout:    timeout,
		maxRetries: maxRetries,
		backoffMs:  backoffMs,
		client:     &http.Client{Timeout: timeout},
	}
}

func (c *httpClient) post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.request(ctx, http.MethodPost, path, body, result)
}

func (c *httpClient) get(ctx context.Context, path string, result interface{}) error {
	return c.request(ctx, http.MethodGet, path, nil, result)
}

func (c *httpClient) put(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.request(ctx, http.MethodPut, path, body, result)
}

func (c *httpClient) delete(ctx context.Context, path string, result interface{}) error {
	return c.request(ctx, http.MethodDelete, path, nil, result)
}

func (c *httpClient) request(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	url := c.baseURL + "/" + strings.TrimPrefix(path, "/")

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			backoff := c.jitterBackoff(attempt)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(backoff):
			}

			if body != nil {
				jsonBody, _ := json.Marshal(body)
				bodyReader = bytes.NewReader(jsonBody)
			}
		}

		req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API-Key", c.apiKey)
		req.Header.Set("User-Agent", fmt.Sprintf("oncall-go/%s", Version))

		resp, err := c.client.Do(req)
		if err != nil {
			if attempt < c.maxRetries {
				lastErr = &NetworkError{OnCallError: OnCallError{Message: "network error", Err: err}}
				continue
			}
			return &NetworkError{OnCallError: OnCallError{Message: "network error", Err: err}}
		}

		requestID := resp.Header.Get("x-request-id")
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			if result != nil && len(body) > 0 {
				if err := json.Unmarshal(body, result); err != nil {
					return fmt.Errorf("failed to unmarshal response: %w", err)
				}
			}
			return nil
		}

		var errResp struct {
			Error   string `json:"error"`
			Message string `json:"message"`
		}
		json.Unmarshal(body, &errResp)

		message := errResp.Error
		if message == "" {
			message = errResp.Message
		}
		if message == "" {
			message = "Request failed"
		}

		mappedErr := mapHTTPError(resp.StatusCode, message, requestID)

		if c.shouldRetry(resp.StatusCode) && attempt < c.maxRetries {
			lastErr = mappedErr
			continue
		}

		return mappedErr
	}

	return lastErr
}

func (c *httpClient) shouldRetry(statusCode int) bool {
	if statusCode == 429 {
		return true
	}
	if statusCode >= 500 && statusCode != 501 && statusCode != 505 {
		return true
	}
	return false
}

func (c *httpClient) jitterBackoff(attempt int) time.Duration {
	base := float64(c.backoffMs) * math.Pow(2, float64(attempt-1))
	jitter := rand.Float64() * 100
	return time.Duration(base+jitter) * time.Millisecond
}
