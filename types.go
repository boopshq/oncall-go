package oncall

import "time"

type Result[T any] struct {
	Data  *T
	Error error
}

type CreateRelayInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	ExternalKey *string `json:"externalKey,omitempty"`
}

type Relay struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organizationId"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ExternalKey    *string   `json:"externalKey,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type DayOfWeek string

const (
	Monday    DayOfWeek = "monday"
	Tuesday   DayOfWeek = "tuesday"
	Wednesday DayOfWeek = "wednesday"
	Thursday  DayOfWeek = "thursday"
	Friday    DayOfWeek = "friday"
	Saturday  DayOfWeek = "saturday"
	Sunday    DayOfWeek = "sunday"
)

type ScheduleType string

const (
	ScheduleTypeDaily  ScheduleType = "daily"
	ScheduleTypeWeekly ScheduleType = "weekly"
)

type CreateScheduleInput struct {
	Name        string       `json:"name"`
	RelayID     string       `json:"relayId"`
	Type        ScheduleType `json:"type"`
	StartDay    DayOfWeek    `json:"startDay"`
	StartTime   string       `json:"startTime"`
	ExternalKey *string      `json:"externalKey,omitempty"`
}

type Schedule struct {
	ID             string       `json:"id"`
	OrganizationID string       `json:"organizationId"`
	RelayID        string       `json:"relayId"`
	Name           string       `json:"name"`
	Type           ScheduleType `json:"type"`
	StartDay       DayOfWeek    `json:"startDay"`
	StartTime      string       `json:"startTime"`
	ExternalKey    *string      `json:"externalKey,omitempty"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
	DeletedAt      *time.Time   `json:"deletedAt,omitempty"`
}

type AddScheduleMemberInput struct {
	UserID string `json:"userId"`
}

type ScheduleMember struct {
	ScheduleID string     `json:"scheduleId"`
	UserID     string     `json:"userId"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
}

type GetAssignmentsParams struct {
	Type  *string `json:"type,omitempty"`
	Count *int    `json:"count,omitempty"`
	Date  *string `json:"date,omitempty"`
}

type ScheduleAssignment struct {
	UserID           string `json:"userId"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	AssignmentNumber int    `json:"assignmentNumber"`
}

type RelayRuleType string

const (
	RuleTypeScheduleNotify RelayRuleType = "schedule_notify"
	RuleTypeWebhook        RelayRuleType = "webhook"
	RuleTypeAgent          RelayRuleType = "agent"
	RuleTypeExternalAPI    RelayRuleType = "external_api"
	RuleTypeWait           RelayRuleType = "wait"
	RuleTypeConditional    RelayRuleType = "conditional"
	RuleTypeEscalate       RelayRuleType = "escalate"
)

type NotificationMethod string

const (
	NotificationEmail NotificationMethod = "email"
	NotificationSMS   NotificationMethod = "sms"
)

type ScheduleNotifyConfig struct {
	ScheduleID         string              `json:"scheduleId"`
	NotificationMethod *NotificationMethod `json:"notificationMethod,omitempty"`
}

type HTTPMethod string

const (
	MethodGET    HTTPMethod = "GET"
	MethodPOST   HTTPMethod = "POST"
	MethodPUT    HTTPMethod = "PUT"
	MethodPATCH  HTTPMethod = "PATCH"
	MethodDELETE HTTPMethod = "DELETE"
)

type WebhookConfig struct {
	Endpoint string            `json:"endpoint"`
	Method   *HTTPMethod       `json:"method,omitempty"`
	Headers  map[string]string `json:"headers,omitempty"`
	Payload  map[string]any    `json:"payload,omitempty"`
	Timeout  *int              `json:"timeout,omitempty"`
}

type AgentType string

const (
	AgentTypeDevin  AgentType = "devin"
	AgentTypeRhythm AgentType = "rhythm"
)

type AgentConfig struct {
	AgentType       AgentType      `json:"agentType"`
	IntegrationID   *string        `json:"integrationId,omitempty"`
	Endpoint        *string        `json:"endpoint,omitempty"`
	Config          map[string]any `json:"config,omitempty"`
	PollInterval    *int           `json:"pollInterval,omitempty"`
	MaxPollAttempts *int           `json:"maxPollAttempts,omitempty"`
}

type ExternalApiConfig struct {
	APIType       string            `json:"apiType"`
	Endpoint      *string           `json:"endpoint,omitempty"`
	IntegrationID *string           `json:"integrationId,omitempty"`
	Method        *HTTPMethod       `json:"method,omitempty"`
	Headers       map[string]string `json:"headers,omitempty"`
	Payload       map[string]any    `json:"payload,omitempty"`
	Timeout       *int              `json:"timeout,omitempty"`
	UseContextFrom *string          `json:"useContextFrom,omitempty"`
}

type WaitConfig struct {
	Duration int     `json:"duration"`
	WaitType *string `json:"waitType,omitempty"`
}

type ConditionalConfig struct {
	Condition      string  `json:"condition"`
	ConditionValue *string `json:"conditionValue,omitempty"`
	TrueRuleID     *string `json:"trueRuleId,omitempty"`
	FalseRuleID    *string `json:"falseRuleId,omitempty"`
}

type EscalateConfig struct {
	MaxAttempts   int     `json:"maxAttempts"`
	EscalateAfter int     `json:"escalateAfter"`
	ResetToRuleID *string `json:"resetToRuleId,omitempty"`
}

type CreateRelayRuleInput struct {
	ExternalKey *string        `json:"externalKey,omitempty"`
	Name        string         `json:"name"`
	RuleType    RelayRuleType  `json:"ruleType"`
	Group       *string        `json:"group,omitempty"`
	Order       *int           `json:"order,omitempty"`
	Config      map[string]any `json:"config"`
	Enabled     *bool          `json:"enabled,omitempty"`
}

type UpdateRelayRuleInput struct {
	Name     *string        `json:"name,omitempty"`
	RuleType *RelayRuleType `json:"ruleType,omitempty"`
	Group    *string        `json:"group,omitempty"`
	Order    *int           `json:"order,omitempty"`
	Config   map[string]any `json:"config,omitempty"`
	Enabled  *bool          `json:"enabled,omitempty"`
}

type RelayRule struct {
	ID             string         `json:"id"`
	OrganizationID string         `json:"organizationId"`
	RelayID        string         `json:"relayId"`
	Group          string         `json:"group"`
	ExternalKey    *string        `json:"externalKey,omitempty"`
	Name           string         `json:"name"`
	Order          int            `json:"order"`
	RuleType       RelayRuleType  `json:"ruleType"`
	Config         map[string]any `json:"config"`
	Enabled        bool           `json:"enabled"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      *time.Time     `json:"deletedAt,omitempty"`
}

type ListRelayRulesParams struct {
	Enabled  *bool          `json:"enabled,omitempty"`
	RuleType *RelayRuleType `json:"ruleType,omitempty"`
}

type ReorderRelayRulesInput struct {
	Rules []struct {
		ID    string `json:"id"`
		Order int    `json:"order"`
	} `json:"rules"`
}

type ContactMethodTransport string

const (
	TransportEmail ContactMethodTransport = "email"
	TransportSMS   ContactMethodTransport = "sms"
)

type CreateContactMethodInput struct {
	UserID    string                 `json:"userId"`
	Transport ContactMethodTransport `json:"transport"`
	Value     string                 `json:"value"`
}

type ContactMethod struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Transport string    `json:"transport"`
	Value     string    `json:"value"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type ListContactMethodsParams struct {
	UserID string `json:"userId"`
}

type DeleteContactMethodParams struct {
	UserID string `json:"userId"`
}

type OnCallUser struct {
	UserID           string `json:"userId"`
	Email            string `json:"email"`
	AssignmentNumber int    `json:"assignmentNumber"`
	ScheduleID       string `json:"scheduleId"`
}

type AlertSeverity string

const (
	SeverityCritical AlertSeverity = "critical"
	SeverityHigh     AlertSeverity = "high"
	SeverityMedium   AlertSeverity = "medium"
	SeverityLow      AlertSeverity = "low"
	SeverityInfo     AlertSeverity = "info"
)

type Alert struct {
	ID               string         `json:"id"`
	OrganizationID   string         `json:"organizationId"`
	WebhookID        string         `json:"webhookId"`
	Title            string         `json:"title"`
	Message          *string        `json:"message,omitempty"`
	Severity         AlertSeverity  `json:"severity"`
	Source           string         `json:"source"`
	ExternalID       *string        `json:"externalId,omitempty"`
	Metadata         map[string]any `json:"metadata,omitempty"`
	AssignedToUserID *string        `json:"assignedToUserId,omitempty"`
	AcknowledgedAt   *time.Time     `json:"acknowledgedAt,omitempty"`
	AcknowledgedBy   *string        `json:"acknowledgedBy,omitempty"`
	ResolvedAt       *time.Time     `json:"resolvedAt,omitempty"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
}

type AcknowledgeAlertInput struct {
	UserID *string `json:"userId,omitempty"`
}

type IntegrationProvider string

const (
	ProviderDevin  IntegrationProvider = "devin"
	ProviderRhythm IntegrationProvider = "rhythm"
)

type CreateIntegrationInput struct {
	Name     string              `json:"name"`
	Provider IntegrationProvider `json:"provider"`
	APIKey   string              `json:"apiKey"`
	Metadata map[string]any      `json:"metadata,omitempty"`
}

type UpdateIntegrationInput struct {
	Name     *string        `json:"name,omitempty"`
	APIKey   *string        `json:"apiKey,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type Integration struct {
	ID             string              `json:"id"`
	OrganizationID string              `json:"organizationId"`
	Name           string              `json:"name"`
	Provider       IntegrationProvider `json:"provider"`
	APIKey         string              `json:"apiKey"`
	Metadata       map[string]any      `json:"metadata,omitempty"`
	CreatedBy      string              `json:"createdBy"`
	CreatedAt      time.Time           `json:"createdAt"`
	UpdatedAt      time.Time           `json:"updatedAt"`
	DeletedAt      *time.Time          `json:"deletedAt,omitempty"`
}
