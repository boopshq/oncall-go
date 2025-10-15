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

	description := "Primary on-call relay for the engineering team"
	relay, err := client.Relay.Create(ctx, oncall.CreateRelayInput{
		Name:        "Engineering Primary",
		Description: &description,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created relay: %s (ID: %s)\n", relay.Name, relay.ID)

	relays, err := client.Relay.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nAll relays (%d):\n", len(relays))
	for _, r := range relays {
		fmt.Printf("  - %s (ID: %s)\n", r.Name, r.ID)
	}
}
