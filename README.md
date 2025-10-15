# oncall-go

Go SDK for the oncall.sh API by beeps.

## Sign up at oncall.sh

- [Documentation](https://oncall.sh): Full documentation site

## Installation

```bash
go get github.com/boopshq/oncall-go
```

## Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/boopshq/oncall-go"
)

func main() {
    client, err := oncall.NewClient(oncall.Config{
        APIKey: "your-api-key-here",
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Create a relay
    description := "Primary on-call relay"
    relay, err := client.Relay.Create(ctx, oncall.CreateRelayInput{
        Name:        "Engineering Primary",
        Description: &description,
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Created relay: %s (ID: %s)\n", relay.Name, relay.ID)

    // List all relays
    relays, err := client.Relay.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    for _, r := range relays {
        fmt.Printf("  - %s\n", r.Name)
    }
}
```

## Configuration

The SDK supports several configuration options:

```go
client, err := oncall.NewClient(oncall.Config{
    APIKey:     "your-api-key",
    BaseURL:    "https://api.oncall.sh/v0", // Optional, defaults to production API
    Timeout:    30 * time.Second,            // Optional, defaults to 10 seconds
    MaxRetries: 3,                           // Optional, defaults to 2
    BackoffMs:  500,                         // Optional, defaults to 300ms
})
```

## Error Handling

The SDK provides two patterns for error handling:

### Standard (throws errors)

```go
relay, err := client.Relay.Create(ctx, oncall.CreateRelayInput{
    Name: "Test Relay",
})
if err != nil {
    switch e := err.(type) {
    case *oncall.AuthError:
        fmt.Println("Authentication failed")
    case *oncall.ValidationError:
        fmt.Println("Invalid input")
    case *oncall.NotFoundError:
        fmt.Println("Resource not found")
    case *oncall.RateLimitError:
        fmt.Println("Rate limit exceeded")
    case *oncall.ServerError:
        fmt.Println("Server error")
    case *oncall.NetworkError:
        fmt.Println("Network error")
    default:
        fmt.Printf("Unexpected error: %v\n", e)
    }
    return
}
```

### Safe variants (returns Result type)

All resource methods have a "Safe" variant that returns a `Result[T]` instead of throwing:

```go
result := client.Relay.CreateSafe(ctx, oncall.CreateRelayInput{
    Name: "Test Relay",
})
if result.Error != nil {
    fmt.Printf("Error: %v\n", result.Error)
    return
}

relay := result.Data
fmt.Printf("Created: %s\n", relay.Name)
```

## Available Resources

### Relay

```go
// Create a relay
relay, err := client.Relay.Create(ctx, oncall.CreateRelayInput{...})

// List relays
relays, err := client.Relay.List(ctx)

// Access relay rules
rules, err := client.Relay.Rules.List(ctx, relayID, nil)
rule, err := client.Relay.Rules.Create(ctx, relayID, oncall.CreateRelayRuleInput{...})
rule, err := client.Relay.Rules.Get(ctx, relayID, ruleID)
rule, err := client.Relay.Rules.Update(ctx, relayID, ruleID, oncall.UpdateRelayRuleInput{...})
err := client.Relay.Rules.Delete(ctx, relayID, ruleID)
rules, err := client.Relay.Rules.Reorder(ctx, relayID, oncall.ReorderRelayRulesInput{...})
```

### Schedule

```go
// Create a schedule
schedule, err := client.Schedule.Create(ctx, oncall.CreateScheduleInput{
    Name:      "Weekly Rotation",
    RelayID:   relayID,
    Type:      oncall.ScheduleTypeWeekly,
    StartDay:  oncall.Monday,
    StartTime: "09:00",
})

// List schedules
schedules, err := client.Schedule.List(ctx)

// Add member to schedule
member, err := client.Schedule.AddMember(ctx, scheduleID, oncall.AddScheduleMemberInput{
    UserID: "user-123",
})

// Get assignments
count := 5
assignments, err := client.Schedule.GetAssignments(ctx, scheduleID, &oncall.GetAssignmentsParams{
    Count: &count,
})

// Get current on-call user
onCall, err := client.Schedule.GetOnCall(ctx, scheduleID)
```

### Alert

```go
// List alerts
alerts, err := client.Alert.List(ctx)
activeAlerts, err := client.Alert.ListActive(ctx)
resolvedAlerts, err := client.Alert.ListResolved(ctx)

// Get alert
alert, err := client.Alert.Get(ctx, alertID)

// Acknowledge alert
alert, err := client.Alert.Acknowledge(ctx, alertID, oncall.AcknowledgeAlertInput{
    UserID: "user-123",
})

// Resolve alert
alert, err := client.Alert.Resolve(ctx, alertID)

// Assign alert
alert, err := client.Alert.Assign(ctx, alertID, "user-123")
```

### Contact Method

```go
// List contact methods
methods, err := client.ContactMethod.List(ctx, oncall.ListContactMethodsParams{
    UserID: "user-123",
})

// Create contact method
method, err := client.ContactMethod.Create(ctx, oncall.CreateContactMethodInput{
    UserID:    "user-123",
    Transport: oncall.TransportEmail,
    Value:     "user@example.com",
})

// Delete contact method
err := client.ContactMethod.Delete(ctx, methodID, oncall.DeleteContactMethodParams{
    UserID: "user-123",
})
```

### Integration

```go
// List integrations
integrations, err := client.Integration.List(ctx)

// Create integration
integration, err := client.Integration.Create(ctx, oncall.CreateIntegrationInput{
    Name:     "Devin Integration",
    Provider: oncall.ProviderDevin,
    APIKey:   "devin-api-key",
})

// Get integration
integration, err := client.Integration.Get(ctx, integrationID)

// Update integration
integration, err := client.Integration.Update(ctx, integrationID, oncall.UpdateIntegrationInput{
    Name: func() *string { s := "Updated Name"; return &s }(),
})

// Delete integration
err := client.Integration.Delete(ctx, integrationID)
```

## Context Support

All methods accept a `context.Context` as the first parameter, allowing you to:

- Set timeouts
- Cancel requests
- Pass request-scoped values

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

relays, err := client.Relay.List(ctx)
```

## Examples

See the [examples](./examples) directory for more complete examples:

- [basic.go](./examples/basic.go) - Basic relay creation and listing
- [schedule.go](./examples/schedule.go) - Schedule management and assignments
- [alerts.go](./examples/alerts.go) - Alert handling and acknowledgment
- [safe_variants.go](./examples/safe_variants.go) - Using safe variants for error handling

## Development

Run tests:

```bash
go test ./...
```

## License

See LICENSE file for details.
