package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/boopshq/oncall-go"
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
		acknowledged, err := client.Alert.Acknowledge(ctx, alert.ID, oncall.AcknowledgeAlertInput{
			UserID: "user-123",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nAcknowledged alert: %s\n", acknowledged.Title)

		resolved, err := client.Alert.Resolve(ctx, alert.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Resolved alert: %s\n", resolved.Title)
	}
}
