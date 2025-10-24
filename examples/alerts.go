package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/oncall-sh/oncall-go"
)

func main() {
	client, err := oncall.NewClient(oncall.Config{
		APIKey: os.Getenv("ONCALL_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	activeAlerts, err := client.Alert.ListActive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Active alerts (%d):\n", len(activeAlerts))
	for _, alert := range activeAlerts {
		fmt.Printf("  [%s] %s - %s\n", alert.Severity, alert.Title, alert.Source)
		if alert.AssignedToUserID != nil {
			fmt.Printf("    Assigned to: %s\n", *alert.AssignedToUserID)
		}
	}

	if len(activeAlerts) > 0 {
		alert := activeAlerts[0]

		// Acknowledge using the API key's associated user (userId is optional)
		acknowledged, err := client.Alert.Acknowledge(ctx, alert.ID, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nAcknowledged alert: %s\n", acknowledged.Title)

		// Or explicitly provide a different userId
		userID := "user-123"
		acknowledged2, err := client.Alert.Acknowledge(ctx, alert.ID, &oncall.AcknowledgeAlertInput{
			UserID: &userID,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Acknowledged alert with specific user: %s\n", acknowledged2.Title)

		resolved, err := client.Alert.Resolve(ctx, alert.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Resolved alert: %s\n", resolved.Title)
	}
}
