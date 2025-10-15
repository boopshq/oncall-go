package oncall

import "context"

type RelayResource struct {
	http  *httpClient
	Rules *RelayRulesResource
}

func newRelayResource(http *httpClient) *RelayResource {
	return &RelayResource{
		http:  http,
		Rules: newRelayRulesResource(http),
	}
}

func (r *RelayResource) Create(ctx context.Context, input CreateRelayInput) (*Relay, error) {
	var result struct {
		Relay Relay `json:"relay"`
	}
	if err := r.http.post(ctx, "/relay", input, &result); err != nil {
		return nil, err
	}
	return &result.Relay, nil
}

func (r *RelayResource) List(ctx context.Context) ([]Relay, error) {
	var result struct {
		Relays []Relay `json:"relays"`
	}
	if err := r.http.get(ctx, "/relay", &result); err != nil {
		return nil, err
	}
	return result.Relays, nil
}

func (r *RelayResource) CreateSafe(ctx context.Context, input CreateRelayInput) Result[Relay] {
	relay, err := r.Create(ctx, input)
	if err != nil {
		return Result[Relay]{Error: err}
	}
	return Result[Relay]{Data: relay}
}

func (r *RelayResource) ListSafe(ctx context.Context) Result[[]Relay] {
	relays, err := r.List(ctx)
	if err != nil {
		return Result[[]Relay]{Error: err}
	}
	return Result[[]Relay]{Data: &relays}
}
