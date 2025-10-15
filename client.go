package oncall

import (
	"errors"
	"time"
)

type Config struct {
	APIKey     string
	BaseURL    string
	Timeout    time.Duration
	MaxRetries int
	BackoffMs  int
}

type Client struct {
	Relay         *RelayResource
	Schedule      *ScheduleResource
	ContactMethod *ContactMethodResource
	Alert         *AlertResource
	Integration   *IntegrationResource
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, errors.New("apiKey is required")
	}

	http := newHTTPClient(&cfg)

	return &Client{
		Relay:         newRelayResource(http),
		Schedule:      newScheduleResource(http),
		ContactMethod: newContactMethodResource(http),
		Alert:         newAlertResource(http),
		Integration:   newIntegrationResource(http),
	}, nil
}
