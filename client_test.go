package oncall

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("requires API key", func(t *testing.T) {
		_, err := NewClient(Config{})
		if err == nil {
			t.Fatal("expected error when API key is missing")
		}
		if err.Error() != "apiKey is required" {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("creates client with API key", func(t *testing.T) {
		client, err := NewClient(Config{APIKey: "test-key"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if client == nil {
			t.Fatal("expected client to be created")
		}
		if client.Relay == nil {
			t.Fatal("expected Relay resource to be initialized")
		}
		if client.Schedule == nil {
			t.Fatal("expected Schedule resource to be initialized")
		}
		if client.ContactMethod == nil {
			t.Fatal("expected ContactMethod resource to be initialized")
		}
		if client.Alert == nil {
			t.Fatal("expected Alert resource to be initialized")
		}
		if client.Integration == nil {
			t.Fatal("expected Integration resource to be initialized")
		}
	})
}

func TestErrorTypes(t *testing.T) {
	t.Run("AuthError", func(t *testing.T) {
		err := &AuthError{OnCallError: OnCallError{Message: "unauthorized", RequestID: "req123"}}
		if err.Error() != "unauthorized (request_id: req123)" {
			t.Fatalf("unexpected error message: %v", err.Error())
		}
	})

	t.Run("ValidationError", func(t *testing.T) {
		err := &ValidationError{OnCallError: OnCallError{Message: "invalid input"}}
		if err.Error() != "invalid input" {
			t.Fatalf("unexpected error message: %v", err.Error())
		}
	})

	t.Run("NotFoundError", func(t *testing.T) {
		err := &NotFoundError{OnCallError: OnCallError{Message: "not found"}}
		if err.Error() != "not found" {
			t.Fatalf("unexpected error message: %v", err.Error())
		}
	})
}

func TestMapHTTPError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		message    string
		requestID  string
		wantType   string
	}{
		{"401 -> AuthError", 401, "unauthorized", "req123", "*oncall.AuthError"},
		{"403 -> AuthError", 403, "forbidden", "req124", "*oncall.AuthError"},
		{"400 -> ValidationError", 400, "bad request", "req125", "*oncall.ValidationError"},
		{"422 -> ValidationError", 422, "unprocessable", "req126", "*oncall.ValidationError"},
		{"404 -> NotFoundError", 404, "not found", "req127", "*oncall.NotFoundError"},
		{"429 -> RateLimitError", 429, "rate limited", "req128", "*oncall.RateLimitError"},
		{"500 -> ServerError", 500, "server error", "req129", "*oncall.ServerError"},
		{"502 -> ServerError", 502, "bad gateway", "req130", "*oncall.ServerError"},
		{"418 -> HTTPError", 418, "teapot", "req131", "*oncall.HTTPError"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mapHTTPError(tt.statusCode, tt.message, tt.requestID)
			errType := ""
			switch err.(type) {
			case *AuthError:
				errType = "*oncall.AuthError"
			case *ValidationError:
				errType = "*oncall.ValidationError"
			case *NotFoundError:
				errType = "*oncall.NotFoundError"
			case *RateLimitError:
				errType = "*oncall.RateLimitError"
			case *ServerError:
				errType = "*oncall.ServerError"
			case *HTTPError:
				errType = "*oncall.HTTPError"
			}
			if errType != tt.wantType {
				t.Fatalf("expected error type %s, got %s", tt.wantType, errType)
			}
		})
	}
}

func TestResultType(t *testing.T) {
	t.Run("Result with data", func(t *testing.T) {
		relay := &Relay{ID: "relay123"}
		result := Result[Relay]{Data: relay}
		if result.Data == nil {
			t.Fatal("expected data to be set")
		}
		if result.Data.ID != "relay123" {
			t.Fatalf("unexpected relay ID: %v", result.Data.ID)
		}
		if result.Error != nil {
			t.Fatal("expected error to be nil")
		}
	})

	t.Run("Result with error", func(t *testing.T) {
		err := &AuthError{OnCallError: OnCallError{Message: "unauthorized"}}
		result := Result[Relay]{Error: err}
		if result.Data != nil {
			t.Fatal("expected data to be nil")
		}
		if result.Error == nil {
			t.Fatal("expected error to be set")
		}
	})
}

func ExampleClient() {
	client, err := NewClient(Config{
		APIKey: "your-api-key",
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	relays, err := client.Relay.List(ctx)
	if err != nil {
		panic(err)
	}

	for _, relay := range relays {
		_ = relay.Name
	}
}
