package oncall

import (
	"context"
	"fmt"
)

type IntegrationResource struct {
	http *httpClient
}

func newIntegrationResource(http *httpClient) *IntegrationResource {
	return &IntegrationResource{http: http}
}

func (i *IntegrationResource) List(ctx context.Context) ([]Integration, error) {
	var result struct {
		Integrations []Integration `json:"integrations"`
	}
	if err := i.http.get(ctx, "/integrations", &result); err != nil {
		return nil, err
	}
	return result.Integrations, nil
}

func (i *IntegrationResource) Create(ctx context.Context, input CreateIntegrationInput) (*Integration, error) {
	var result struct {
		Integration Integration `json:"integration"`
	}
	if err := i.http.post(ctx, "/integrations", input, &result); err != nil {
		return nil, err
	}
	return &result.Integration, nil
}

func (i *IntegrationResource) Get(ctx context.Context, integrationID string) (*Integration, error) {
	var result struct {
		Integration Integration `json:"integration"`
	}
	path := fmt.Sprintf("/integrations/%s", integrationID)
	if err := i.http.get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result.Integration, nil
}

func (i *IntegrationResource) Update(ctx context.Context, integrationID string, input UpdateIntegrationInput) (*Integration, error) {
	var result struct {
		Integration Integration `json:"integration"`
	}
	path := fmt.Sprintf("/integrations/%s", integrationID)
	if err := i.http.put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result.Integration, nil
}

func (i *IntegrationResource) Delete(ctx context.Context, integrationID string) error {
	var result struct {
		Success bool `json:"success"`
	}
	path := fmt.Sprintf("/integrations/%s", integrationID)
	if err := i.http.delete(ctx, path, &result); err != nil {
		return err
	}
	return nil
}

func (i *IntegrationResource) ListSafe(ctx context.Context) Result[[]Integration] {
	integrations, err := i.List(ctx)
	if err != nil {
		return Result[[]Integration]{Error: err}
	}
	return Result[[]Integration]{Data: &integrations}
}

func (i *IntegrationResource) CreateSafe(ctx context.Context, input CreateIntegrationInput) Result[Integration] {
	integration, err := i.Create(ctx, input)
	if err != nil {
		return Result[Integration]{Error: err}
	}
	return Result[Integration]{Data: integration}
}

func (i *IntegrationResource) GetSafe(ctx context.Context, integrationID string) Result[Integration] {
	integration, err := i.Get(ctx, integrationID)
	if err != nil {
		return Result[Integration]{Error: err}
	}
	return Result[Integration]{Data: integration}
}

func (i *IntegrationResource) UpdateSafe(ctx context.Context, integrationID string, input UpdateIntegrationInput) Result[Integration] {
	integration, err := i.Update(ctx, integrationID, input)
	if err != nil {
		return Result[Integration]{Error: err}
	}
	return Result[Integration]{Data: integration}
}

func (i *IntegrationResource) DeleteSafe(ctx context.Context, integrationID string) Result[bool] {
	err := i.Delete(ctx, integrationID)
	if err != nil {
		return Result[bool]{Error: err}
	}
	success := true
	return Result[bool]{Data: &success}
}
