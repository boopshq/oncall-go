package oncall

import (
	"context"
	"fmt"
	"net/url"
)

type ScheduleResource struct {
	http *httpClient
}

func newScheduleResource(http *httpClient) *ScheduleResource {
	return &ScheduleResource{http: http}
}

func (s *ScheduleResource) Create(ctx context.Context, input CreateScheduleInput) (*Schedule, error) {
	var result struct {
		Schedule Schedule `json:"schedule"`
	}
	if err := s.http.post(ctx, "/schedule", input, &result); err != nil {
		return nil, err
	}
	return &result.Schedule, nil
}

func (s *ScheduleResource) List(ctx context.Context) ([]Schedule, error) {
	var result struct {
		Schedules []Schedule `json:"schedules"`
	}
	if err := s.http.get(ctx, "/schedule", &result); err != nil {
		return nil, err
	}
	return result.Schedules, nil
}

func (s *ScheduleResource) AddMember(ctx context.Context, scheduleID string, input AddScheduleMemberInput) (*ScheduleMember, error) {
	var result struct {
		Member ScheduleMember `json:"member"`
	}
	path := fmt.Sprintf("/schedule/%s/members", scheduleID)
	if err := s.http.post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result.Member, nil
}

func (s *ScheduleResource) GetAssignments(ctx context.Context, scheduleID string, params *GetAssignmentsParams) ([]ScheduleAssignment, error) {
	path := fmt.Sprintf("/schedule/%s/assignments", scheduleID)

	if params != nil {
		query := url.Values{}
		if params.Type != nil {
			query.Set("type", *params.Type)
		}
		if params.Count != nil {
			query.Set("count", fmt.Sprintf("%d", *params.Count))
		}
		if params.Date != nil {
			query.Set("date", *params.Date)
		}
		if len(query) > 0 {
			path = fmt.Sprintf("%s?%s", path, query.Encode())
		}
	}

	var result struct {
		Assignments []ScheduleAssignment `json:"assignments"`
	}
	if err := s.http.get(ctx, path, &result); err != nil {
		return nil, err
	}
	return result.Assignments, nil
}

func (s *ScheduleResource) GetOnCall(ctx context.Context, scheduleID string) (*OnCallUser, error) {
	var result struct {
		OnCall OnCallUser `json:"onCall"`
	}
	path := fmt.Sprintf("/schedule/%s/on-call", scheduleID)
	if err := s.http.get(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result.OnCall, nil
}

func (s *ScheduleResource) GetOnCallSafe(ctx context.Context, scheduleID string) Result[OnCallUser] {
	onCall, err := s.GetOnCall(ctx, scheduleID)
	if err != nil {
		return Result[OnCallUser]{Error: err}
	}
	return Result[OnCallUser]{Data: onCall}
}
