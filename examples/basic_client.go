package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
	// Get server URL from environment variable or use default
	serverURL := os.Getenv("MIX_SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8088"
	}

	// Initialize the SDK client
	client := mix.New(
		mix.WithServerURL(serverURL),
	)

	ctx := context.Background()

	fmt.Println("=== Mix Go SDK - Basic Client Example ===\n")

	// 1. Health Check
	fmt.Println("1. Checking server health...")
	healthResp, err := client.System.HealthCheck(ctx)
	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}
	fmt.Printf("   Status: %s\n\n", *healthResp.Object.Status)

	// 2. List existing sessions
	fmt.Println("2. Listing existing sessions...")
	listResp, err := client.Sessions.ListSessions(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list sessions: %v", err)
	}
	fmt.Printf("   Found %d session(s)\n\n", len(listResp.SessionData))

	// 3. Create a new session
	fmt.Println("3. Creating a new session...")
	createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
		Title: "Basic Client Example Session",
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	sessionID := createResp.SessionData.ID
	fmt.Printf("   Created session: %s\n", sessionID)
	fmt.Printf("   Title: %s\n\n", createResp.SessionData.Title)

	// 4. Get session details
	fmt.Println("4. Getting session details...")
	sessionResp, err := client.Sessions.GetSession(ctx, sessionID)
	if err != nil {
		log.Fatalf("Failed to get session: %v", err)
	}
	fmt.Printf("   Session ID: %s\n", sessionResp.SessionData.ID)
	fmt.Printf("   Title: %s\n", sessionResp.SessionData.Title)
	fmt.Printf("   Created: %s\n\n", sessionResp.SessionData.CreatedAt.String())

	// 5. Send a message to the session
	fmt.Println("5. Sending a message to the session...")
	messageResp, err := client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
		Text: "Hello! This is a test message from the Go SDK.",
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Printf("   Message sent successfully (Status: %d)\n", messageResp.HTTPMeta.Response.StatusCode)
	fmt.Printf("   Note: Message is being processed asynchronously\n\n")

	// 6. List sessions again to see the updated session
	fmt.Println("6. Listing sessions again...")
	listResp2, err := client.Sessions.ListSessions(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list sessions: %v", err)
	}
	fmt.Printf("   Found %d session(s)\n", len(listResp2.SessionData))
	for i, session := range listResp2.SessionData {
		fmt.Printf("   [%d] ID: %s, Title: %s\n", i+1, session.ID, session.Title)
	}
	fmt.Println()

	// 7. Delete the session
	fmt.Println("7. Deleting the session...")
	deleteResp, err := client.Sessions.DeleteSession(ctx, sessionID)
	if err != nil {
		log.Fatalf("Failed to delete session: %v", err)
	}
	fmt.Printf("   Session deleted successfully (Status: %d)\n\n", deleteResp.HTTPMeta.Response.StatusCode)

	// 8. Verify deletion
	fmt.Println("8. Verifying deletion...")
	listResp3, err := client.Sessions.ListSessions(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list sessions: %v", err)
	}
	fmt.Printf("   Found %d session(s) after deletion\n\n", len(listResp3.SessionData))

	fmt.Println("=== Example completed successfully! ===")
}
