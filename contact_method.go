package oncall

import (
	"context"
	"fmt"
	"net/url"
)

type ContactMethodResource struct {
	http *httpClient
}

func newContactMethodResource(http *httpClient) *ContactMethodResource {
	return &ContactMethodResource{http: http}
}

func (c *ContactMethodResource) List(ctx context.Context, params ListContactMethodsParams) ([]ContactMethod, error) {
	query := url.Values{}
	query.Set("userId", params.UserID)
	path := fmt.Sprintf("/contact-methods?%s", query.Encode())

	var result struct {
		ContactMethods []ContactMethod `json:"contactMethods"`
	}
	if err := c.http.get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result.ContactMethods, nil
}

func (c *ContactMethodResource) Create(ctx context.Context, input CreateContactMethodInput) (*ContactMethod, error) {
	var result struct {
		ContactMethod ContactMethod `json:"contactMethod"`
	}
	if err := c.http.post(ctx, "/contact-methods", input, &result); err != nil {
		return nil, err
	}
	return &result.ContactMethod, nil
}

func (c *ContactMethodResource) Delete(ctx context.Context, id string, params DeleteContactMethodParams) error {
	query := url.Values{}
	query.Set("userId", params.UserID)
	path := fmt.Sprintf("/contact-methods/%s?%s", id, query.Encode())

	var result struct {
		Success bool `json:"success"`
	}
	if err := c.http.delete(ctx, path, &result); err != nil {
		return err
	}
	return nil
}

func (c *ContactMethodResource) ListSafe(ctx context.Context, params ListContactMethodsParams) Result[[]ContactMethod] {
	methods, err := c.List(ctx, params)
	if err != nil {
		return Result[[]ContactMethod]{Error: err}
	}
	return Result[[]ContactMethod]{Data: &methods}
}

func (c *ContactMethodResource) CreateSafe(ctx context.Context, input CreateContactMethodInput) Result[ContactMethod] {
	method, err := c.Create(ctx, input)
	if err != nil {
		return Result[ContactMethod]{Error: err}
	}
	return Result[ContactMethod]{Data: method}
}

func (c *ContactMethodResource) DeleteSafe(ctx context.Context, id string, params DeleteContactMethodParams) Result[bool] {
	err := c.Delete(ctx, id, params)
	if err != nil {
		return Result[bool]{Error: err}
	}
	success := true
	return Result[bool]{Data: &success}
}
