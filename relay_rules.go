package oncall

import (
	"context"
	"fmt"
	"net/url"
)

type RelayRulesResource struct {
	http *httpClient
}

func newRelayRulesResource(http *httpClient) *RelayRulesResource {
	return &RelayRulesResource{http: http}
}

func (r *RelayRulesResource) List(ctx context.Context, relayID string, params *ListRelayRulesParams) ([]RelayRule, error) {
	path := fmt.Sprintf("/relay/%s/rules", relayID)

	if params != nil {
		query := url.Values{}
		if params.Enabled != nil {
			query.Set("enabled", fmt.Sprintf("%t", *params.Enabled))
		}
		if params.RuleType != nil {
			query.Set("ruleType", string(*params.RuleType))
		}
		if len(query) > 0 {
			path = fmt.Sprintf("%s?%s", path, query.Encode())
		}
	}

	var result struct {
		Rules []RelayRule `json:"rules"`
	}
	if err := r.http.get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result.Rules, nil
}

func (r *RelayRulesResource) Create(ctx context.Context, relayID string, input CreateRelayRuleInput) (*RelayRule, error) {
	var result struct {
		Rule RelayRule `json:"rule"`
	}
	path := fmt.Sprintf("/relay/%s/rules", relayID)
	if err := r.http.post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result.Rule, nil
}

func (r *RelayRulesResource) Get(ctx context.Context, relayID, ruleID string) (*RelayRule, error) {
	var result struct {
		Rule RelayRule `json:"rule"`
	}
	path := fmt.Sprintf("/relay/%s/rules/%s", relayID, ruleID)
	if err := r.http.get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result.Rule, nil
}

func (r *RelayRulesResource) Update(ctx context.Context, relayID, ruleID string, input UpdateRelayRuleInput) (*RelayRule, error) {
	var result struct {
		Rule RelayRule `json:"rule"`
	}
	path := fmt.Sprintf("/relay/%s/rules/%s", relayID, ruleID)
	if err := r.http.put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result.Rule, nil
}

func (r *RelayRulesResource) Delete(ctx context.Context, relayID, ruleID string) error {
	var result struct {
		Success bool `json:"success"`
	}
	path := fmt.Sprintf("/relay/%s/rules/%s", relayID, ruleID)
	if err := r.http.delete(ctx, path, &result); err != nil {
		return err
	}
	return nil
}

func (r *RelayRulesResource) Reorder(ctx context.Context, relayID string, input ReorderRelayRulesInput) ([]RelayRule, error) {
	var result struct {
		Rules []RelayRule `json:"rules"`
	}
	path := fmt.Sprintf("/relay/%s/rules/reorder", relayID)
	if err := r.http.put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return result.Rules, nil
}

func (r *RelayRulesResource) ListSafe(ctx context.Context, relayID string, params *ListRelayRulesParams) Result[[]RelayRule] {
	rules, err := r.List(ctx, relayID, params)
	if err != nil {
		return Result[[]RelayRule]{Error: err}
	}
	return Result[[]RelayRule]{Data: &rules}
}

func (r *RelayRulesResource) CreateSafe(ctx context.Context, relayID string, input CreateRelayRuleInput) Result[RelayRule] {
	rule, err := r.Create(ctx, relayID, input)
	if err != nil {
		return Result[RelayRule]{Error: err}
	}
	return Result[RelayRule]{Data: rule}
}

func (r *RelayRulesResource) GetSafe(ctx context.Context, relayID, ruleID string) Result[RelayRule] {
	rule, err := r.Get(ctx, relayID, ruleID)
	if err != nil {
		return Result[RelayRule]{Error: err}
	}
	return Result[RelayRule]{Data: rule}
}

func (r *RelayRulesResource) UpdateSafe(ctx context.Context, relayID, ruleID string, input UpdateRelayRuleInput) Result[RelayRule] {
	rule, err := r.Update(ctx, relayID, ruleID, input)
	if err != nil {
		return Result[RelayRule]{Error: err}
	}
	return Result[RelayRule]{Data: rule}
}

func (r *RelayRulesResource) DeleteSafe(ctx context.Context, relayID, ruleID string) Result[bool] {
	err := r.Delete(ctx, relayID, ruleID)
	if err != nil {
		return Result[bool]{Error: err}
	}
	success := true
	return Result[bool]{Data: &success}
}

func (r *RelayRulesResource) ReorderSafe(ctx context.Context, relayID string, input ReorderRelayRulesInput) Result[[]RelayRule] {
	rules, err := r.Reorder(ctx, relayID, input)
	if err != nil {
		return Result[[]RelayRule]{Error: err}
	}
	return Result[[]RelayRule]{Data: &rules}
}
