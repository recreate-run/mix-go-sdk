package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/components"
	"github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
	serverURL := os.Getenv("MIX_SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8088"
	}

	client := mix.New(mix.WithServerURL(serverURL))
	ctx := context.Background()

	fmt.Println("=== Mix Go SDK - Simple Streaming Example ===\n")

	// Create a session
	createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
		Title: "Simple Streaming Example",
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	sessionID := createResp.SessionData.ID
	fmt.Printf("Created session: %s\n\n", sessionID)

	// Cleanup
	defer func() {
		fmt.Println("\nCleaning up...")
		client.Sessions.DeleteSession(ctx, sessionID)
	}()

	// Start streaming and send message concurrently
	var wg sync.WaitGroup
	wg.Add(1)

	// Start streaming
	go func() {
		defer wg.Done()
		streamMessages(ctx, client, sessionID)
	}()

	// Brief delay to ensure stream is connected
	time.Sleep(500 * time.Millisecond)

	// Send message
	message := "Hello! Can you explain what recursion is in one sentence?"
	fmt.Printf("User: %s\n\n", message)
	_, err = client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
		Text: message,
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Wait for streaming to complete
	wg.Wait()

	fmt.Println("\n=== Simple Streaming Completed! ===")
}

func streamMessages(ctx context.Context, client *mix.Mix, sessionID string) {
	fmt.Println("Opening stream...\n")

	streamCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	lastEventID := ""
	streamResp, err := client.Streaming.StreamEvents(streamCtx, sessionID, &lastEventID)
	if err != nil {
		log.Printf("Failed to start stream: %v", err)
		return
	}
	defer streamResp.SSEEventStream.Close()

	// Use the EventStream API
	for streamResp.SSEEventStream.Next() {
		event := streamResp.SSEEventStream.Value()
		if event == nil {
			continue
		}
		handleEvent(event)
	}

	if err := streamResp.SSEEventStream.Err(); err != nil {
		log.Printf("Stream error: %v", err)
	}
}

func handleEvent(event *components.SSEEventStream) {
	switch event.Type {
	case components.SSEEventStreamTypeContent:
		if event.SSEContentEvent != nil && event.SSEContentEvent.Data.Content != "" {
			fmt.Printf("Assistant: %s\n", event.SSEContentEvent.Data.Content)
		}
	case components.SSEEventStreamTypeComplete:
		fmt.Println("\n✓ Complete")
	case components.SSEEventStreamTypeError:
		if event.SSEErrorEvent != nil {
			fmt.Printf("Error: %s\n", event.SSEErrorEvent.Data.Error)
		}
	}
	// Ignore other event types for simplicity
}
