package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/components"
	"github.com/recreate-run/mix-go-sdk/models/operations"
)

// Example Template - Copy this file to get started with Mix Go SDK
func main() {
	// Setup: Get server URL from environment or use default
	serverURL := os.Getenv("MIX_SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8088"
	}

	// Initialize the SDK client
	client := mix.New(mix.WithServerURL(serverURL))
	ctx := context.Background()

	fmt.Println("=== Mix Go SDK - Your Application ===\n")

	// Optional: Store API key if needed
	// Uncomment and set your API key
	// apiKey := os.Getenv("ANTHROPIC_API_KEY")
	// if apiKey != "" {
	// 	_, err := client.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequestBody{
	// 		Provider: "anthropic",
	// 		APIKey:   apiKey,
	// 	})
	// 	if err != nil {
	// 		log.Fatalf("Failed to store API key: %v", err)
	// 	}
	// }

	// Step 1: Create a session
	createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
		Title: "My Application Session",
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	sessionID := createResp.SessionData.ID
	fmt.Printf("Created session: %s\n\n", sessionID)

	// Cleanup on exit
	defer func() {
		fmt.Println("\nCleaning up...")
		client.Sessions.DeleteSession(ctx, sessionID)
	}()

	// Step 2: Send a message with streaming
	var wg sync.WaitGroup
	wg.Add(1)

	// Start streaming in background
	go func() {
		defer wg.Done()
		streamAndDisplay(ctx, client, sessionID)
	}()

	// Brief delay to ensure stream is connected
	time.Sleep(500 * time.Millisecond)

	// Send your message
	message := "Hello! Please introduce yourself and explain what you can help me with."
	fmt.Printf("User: %s\n\n", message)
	_, err = client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
		Text: message,
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Wait for streaming to complete
	wg.Wait()

	fmt.Println("\n=== Completed Successfully! ===")
}

// streamAndDisplay handles SSE stream and displays events
func streamAndDisplay(ctx context.Context, client *mix.Mix, sessionID string) {
	fmt.Println("📡 Opening stream...\n")

	streamCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	lastEventID := ""
	streamResp, err := client.Streaming.StreamEvents(streamCtx, sessionID, &lastEventID)
	if err != nil {
		log.Printf("Failed to start stream: %v", err)
		return
	}
	defer streamResp.SSEEventStream.Close()

	// Use the EventStream API: Next(), Value(), Type
	for streamResp.SSEEventStream.Next() {
		event := streamResp.SSEEventStream.Value()
		if event == nil {
			continue
		}

		displayEvent(event)
	}

	if err := streamResp.SSEEventStream.Err(); err != nil {
		log.Printf("Stream error: %v", err)
	}
}

// displayEvent shows events to the user based on event type
func displayEvent(event *components.SSEEventStream) {
	switch event.Type {
	case components.SSEEventStreamTypeThinking:
		fmt.Printf("💭 Thinking...\n")
	case components.SSEEventStreamTypeContent:
		if event.SSEContentEvent != nil && event.SSEContentEvent.Data.Content != "" {
			fmt.Printf("💬 %s\n", event.SSEContentEvent.Data.Content)
		}
	case components.SSEEventStreamTypeTool:
		fmt.Printf("🔧 Tool event\n")
	case components.SSEEventStreamTypeComplete:
		fmt.Println("\n✓ Complete")
	case components.SSEEventStreamTypeError:
		if event.SSEErrorEvent != nil {
			fmt.Printf("❌ Error: %s\n", event.SSEErrorEvent.Data.Error)
		}
	}
}

// truncate is a helper to shorten long strings
func truncate(s string, maxLen int) string {
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
