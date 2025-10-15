package oncall

import (
	"context"
	"fmt"
)

type AlertResource struct {
	http *httpClient
}

func newAlertResource(http *httpClient) *AlertResource {
	return &AlertResource{http: http}
}

func (a *AlertResource) List(ctx context.Context) ([]Alert, error) {
	var result struct {
		Alerts []Alert `json:"alerts"`
	}
	if err := a.http.get(ctx, "/alerts", &result); err != nil {
		return nil, err
	}
	return result.Alerts, nil
}

func (a *AlertResource) ListActive(ctx context.Context) ([]Alert, error) {
	var result struct {
		Alerts []Alert `json:"alerts"`
	}
	if err := a.http.get(ctx, "/alerts/active", &result); err != nil {
		return nil, err
	}
	return result.Alerts, nil
}

func (a *AlertResource) ListResolved(ctx context.Context) ([]Alert, error) {
	var result struct {
		Alerts []Alert `json:"alerts"`
	}
	if err := a.http.get(ctx, "/alerts/resolved", &result); err != nil {
		return nil, err
	}
	return result.Alerts, nil
}

func (a *AlertResource) Get(ctx context.Context, alertID string) (*Alert, error) {
	var result struct {
		Alert Alert `json:"alert"`
	}
	path := fmt.Sprintf("/alerts/%s", alertID)
	if err := a.http.get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result.Alert, nil
}

func (a *AlertResource) Acknowledge(ctx context.Context, alertID string, input AcknowledgeAlertInput) (*Alert, error) {
	var result struct {
		Alert Alert `json:"alert"`
	}
	path := fmt.Sprintf("/alerts/%s/acknowledge", alertID)
	if err := a.http.post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result.Alert, nil
}

func (a *AlertResource) Resolve(ctx context.Context, alertID string) (*Alert, error) {
	var result struct {
		Alert Alert `json:"alert"`
	}
	path := fmt.Sprintf("/alerts/%s/resolve", alertID)
	if err := a.http.post(ctx, path, struct{}{}, &result); err != nil {
		return nil, err
	}
	return &result.Alert, nil
}

func (a *AlertResource) Assign(ctx context.Context, alertID string, userID string) (*Alert, error) {
	var result struct {
		Alert Alert `json:"alert"`
	}
	path := fmt.Sprintf("/alerts/%s/assign", alertID)
	body := map[string]string{"userId": userID}
	if err := a.http.post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result.Alert, nil
}

func (a *AlertResource) ListSafe(ctx context.Context) Result[[]Alert] {
	alerts, err := a.List(ctx)
	if err != nil {
		return Result[[]Alert]{Error: err}
	}
	return Result[[]Alert]{Data: &alerts}
}

func (a *AlertResource) ListActiveSafe(ctx context.Context) Result[[]Alert] {
	alerts, err := a.ListActive(ctx)
	if err != nil {
		return Result[[]Alert]{Error: err}
	}
	return Result[[]Alert]{Data: &alerts}
}

func (a *AlertResource) ListResolvedSafe(ctx context.Context) Result[[]Alert] {
	alerts, err := a.ListResolved(ctx)
	if err != nil {
		return Result[[]Alert]{Error: err}
	}
	return Result[[]Alert]{Data: &alerts}
}

func (a *AlertResource) GetSafe(ctx context.Context, alertID string) Result[Alert] {
	alert, err := a.Get(ctx, alertID)
	if err != nil {
		return Result[Alert]{Error: err}
	}
	return Result[Alert]{Data: alert}
}

func (a *AlertResource) AcknowledgeSafe(ctx context.Context, alertID string, input AcknowledgeAlertInput) Result[Alert] {
	alert, err := a.Acknowledge(ctx, alertID, input)
	if err != nil {
		return Result[Alert]{Error: err}
	}
	return Result[Alert]{Data: alert}
}

func (a *AlertResource) ResolveSafe(ctx context.Context, alertID string) Result[Alert] {
	alert, err := a.Resolve(ctx, alertID)
	if err != nil {
		return Result[Alert]{Error: err}
	}
	return Result[Alert]{Data: alert}
}

func (a *AlertResource) AssignSafe(ctx context.Context, alertID string, userID string) Result[Alert] {
	alert, err := a.Assign(ctx, alertID, userID)
	if err != nil {
		return Result[Alert]{Error: err}
	}
	return Result[Alert]{Data: alert}
}
