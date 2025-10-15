package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/oncallsh/oncall-go"
)

func main() {
	client, err := oncall.NewClient(oncall.Config{
		APIKey: os.Getenv("ONCALL_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	result := client.Relay.ListSafe(ctx)
	if result.Error != nil {
		fmt.Printf("Error listing relays: %v\n", result.Error)
		switch e := result.Error.(type) {
		case *oncall.AuthError:
			fmt.Println("Authentication failed. Check your API key.")
		case *oncall.RateLimitError:
			fmt.Println("Rate limit exceeded. Try again later.")
		case *oncall.NetworkError:
			fmt.Println("Network error. Check your connection.")
		default:
			fmt.Printf("Unexpected error: %v\n", e)
		}
		os.Exit(1)
	}

	fmt.Printf("Found %d relays:\n", len(*result.Data))
	for _, relay := range *result.Data {
		fmt.Printf("  - %s (ID: %s)\n", relay.Name, relay.ID)
	}

	description := "Test relay with safe variant"
	createResult := client.Relay.CreateSafe(ctx, oncall.CreateRelayInput{
		Name:        "Test Relay",
		Description: &description,
	})
	if createResult.Error != nil {
		fmt.Printf("Error creating relay: %v\n", createResult.Error)
		os.Exit(1)
	}

	fmt.Printf("\nCreated relay: %s (ID: %s)\n", createResult.Data.Name, createResult.Data.ID)
}
