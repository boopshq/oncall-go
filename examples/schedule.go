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

	relays, err := client.Relay.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(relays) == 0 {
		log.Fatal("No relays found. Create a relay first.")
	}
	relay := relays[0]

	schedule, err := client.Schedule.Create(ctx, oncall.CreateScheduleInput{
		Name:      "Weekly rotation",
		RelayID:   relay.ID,
		Type:      oncall.ScheduleTypeWeekly,
		StartDay:  oncall.Monday,
		StartTime: "09:00",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created schedule: %s (ID: %s)\n", schedule.Name, schedule.ID)

	member, err := client.Schedule.AddMember(ctx, schedule.ID, oncall.AddScheduleMemberInput{
		UserID: "user-123",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added member: %s to schedule\n", member.UserID)

	assignments, err := client.Schedule.GetAssignments(ctx, schedule.ID, &oncall.GetAssignmentsParams{
		Count: func() *int { c := 5; return &c }(),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nNext %d assignments:\n", len(assignments))
	for _, a := range assignments {
		fmt.Printf("  - User %s: %s to %s (assignment #%d)\n",
			a.UserID, a.StartDate, a.EndDate, a.AssignmentNumber)
	}

	onCall, err := client.Schedule.GetOnCall(ctx, schedule.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nCurrently on-call: %s (%s)\n", onCall.Email, onCall.UserID)
}
