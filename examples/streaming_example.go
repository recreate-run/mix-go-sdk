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

	fmt.Println("=== Mix Go SDK - Streaming Example ===\n")

	var createdSessions []string

	// Cleanup function
	defer func() {
		if len(createdSessions) > 0 {
			fmt.Println("\n=== Cleanup ===")
			for _, sessionID := range createdSessions {
				fmt.Printf("Deleting session: %s...\n", sessionID)
				_, err := client.Sessions.DeleteSession(ctx, sessionID)
				if err != nil {
					log.Printf("Failed to delete session %s: %v", sessionID, err)
				}
			}
		}
	}()

	// 1. Create a session for streaming
	fmt.Println("1. Creating a new session...")
	createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
		Title: "Streaming Example Session",
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	sessionID := createResp.SessionData.ID
	createdSessions = append(createdSessions, sessionID)
	fmt.Printf("   Created session: %s\n\n", sessionID)

	// 2. Start streaming in a goroutine and send message
	fmt.Println("2. Starting SSE stream and sending message...")
	message := "Explain the concept of recursion with a simple example. Think through your explanation step by step."

	var wg sync.WaitGroup
	wg.Add(1)

	// Start streaming
	go func() {
		defer wg.Done()
		streamAndProcess(ctx, client, sessionID)
	}()

	// Brief delay to ensure stream is connected
	time.Sleep(500 * time.Millisecond)

	// Send message
	fmt.Printf("   User: %s\n\n", message)
	_, err = client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
		Text: message,
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Wait for streaming to complete
	wg.Wait()

	fmt.Println("\n=== Streaming Example Completed Successfully! ===")
}

func streamAndProcess(ctx context.Context, client *mix.Mix, sessionID string) {
	fmt.Println("   📡 Opening SSE stream...\n")

	// Create timeout context for stream
	streamCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Start streaming
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
		processEvent(event)
	}

	if err := streamResp.SSEEventStream.Err(); err != nil {
		log.Printf("Stream error: %v", err)
	}

	fmt.Println("\n   📡 Stream closed")
}

func processEvent(event *components.SSEEventStream) {
	switch event.Type {
	case components.SSEEventStreamTypeThinking:
		if event.SSEThinkingEvent != nil {
			fmt.Printf("   💭 Thinking...\n")
		}

	case components.SSEEventStreamTypeContent:
		if event.SSEContentEvent != nil && event.SSEContentEvent.Data.Content != "" {
			fmt.Printf("   💬 Content: %s\n", truncate(event.SSEContentEvent.Data.Content, 100))
		}

	case components.SSEEventStreamTypeTool:
		if event.SSEToolEvent != nil {
			fmt.Printf("   🔧 Tool Call\n")
		}

	case components.SSEEventStreamTypeToolExecutionStart:
		if event.SSEToolExecutionStartEvent != nil {
			fmt.Printf("   ▶️  Tool Started\n")
		}

	case components.SSEEventStreamTypeToolExecutionComplete:
		if event.SSEToolExecutionCompleteEvent != nil {
			fmt.Printf("   ✅ Tool Completed\n")
		}

	case components.SSEEventStreamTypePermission:
		if event.SSEPermissionEvent != nil {
			fmt.Printf("   🔐 Permission Request\n")
		}

	case components.SSEEventStreamTypeUserMessageCreated:
		if event.SSEUserMessageCreatedEvent != nil {
			fmt.Printf("   📝 User Message Created\n")
		}

	case components.SSEEventStreamTypeSessionCreated:
		if event.SSESessionCreatedEvent != nil {
			fmt.Printf("   🆕 Session Created\n")
		}

	case components.SSEEventStreamTypeSessionDeleted:
		if event.SSESessionDeletedEvent != nil {
			fmt.Printf("   🗑️  Session Deleted\n")
		}

	case components.SSEEventStreamTypeComplete:
		fmt.Println("   ✨ Processing Complete")

	case components.SSEEventStreamTypeError:
		if event.SSEErrorEvent != nil {
			fmt.Printf("   ⚠️  Error: %s\n", event.SSEErrorEvent.Data.Error)
		}

	case components.SSEEventStreamTypeHeartbeat:
		// Heartbeat event - can be ignored or logged
		// fmt.Println("   💓 Heartbeat")

	case components.SSEEventStreamTypeConnected:
		fmt.Println("   🔌 Connected")

	default:
		fmt.Printf("   ℹ️  Event type: %s\n", event.Type)
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
