package oncall

import "fmt"

type OnCallError struct {
	Message   string
	RequestID string
	Err       error
}

func (e *OnCallError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("%s (request_id: %s)", e.Message, e.RequestID)
	}
	return e.Message
}

func (e *OnCallError) Unwrap() error {
	return e.Err
}

type AuthError struct {
	OnCallError
}

type ValidationError struct {
	OnCallError
}

type NotFoundError struct {
	OnCallError
}

type RateLimitError struct {
	OnCallError
}

type ServerError struct {
	OnCallError
}

type NetworkError struct {
	OnCallError
}

type HTTPError struct {
	OnCallError
	StatusCode int
}

func mapHTTPError(statusCode int, message, requestID string) error {
	base := OnCallError{Message: message, RequestID: requestID}

	switch statusCode {
	case 401, 403:
		return &AuthError{OnCallError: base}
	case 400, 422:
		return &ValidationError{OnCallError: base}
	case 404:
		return &NotFoundError{OnCallError: base}
	case 429:
		return &RateLimitError{OnCallError: base}
	case 500, 502, 503, 504:
		return &ServerError{OnCallError: base}
	default:
		return &HTTPError{OnCallError: base, StatusCode: statusCode}
	}
}
