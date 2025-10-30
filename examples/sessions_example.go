package main

import (
	"context"
	"fmt"
	"log"
	"os"
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

	fmt.Println("=== Mix Go SDK - Sessions Example ===\n")

	// Track created sessions for cleanup
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

	// 1. List existing sessions with metadata
	fmt.Println("1. Listing existing sessions...")
	listResp, err := client.Sessions.ListSessions(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list sessions: %v", err)
	}
	fmt.Printf("   Found %d session(s)\n", len(listResp.SessionData))
	if len(listResp.SessionData) > 0 {
		fmt.Println("   Recent sessions:")
		for i, session := range listResp.SessionData {
			if i >= 3 {
				break // Show only first 3
			}
			fmt.Printf("   - %s: %s (Created: %s)\n", session.ID, session.Title, session.CreatedAt.Format(time.RFC3339))
		}
	}
	fmt.Println()

	// 2. Create a new session with custom configuration
	fmt.Println("2. Creating a new session...")
	createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
		Title:       "Session Lifecycle Example",
		SessionType: operations.SessionTypeMain.ToPointer(),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	sessionID := createResp.SessionData.ID
	createdSessions = append(createdSessions, sessionID)
	fmt.Printf("   Created session: %s\n", sessionID)
	fmt.Printf("   Title: %s\n", createResp.SessionData.Title)
	fmt.Printf("   Type: %s\n\n", createResp.SessionData.SessionType)

	// 3. Get session details
	fmt.Println("3. Retrieving session details...")
	sessionResp, err := client.Sessions.GetSession(ctx, sessionID)
	if err != nil {
		log.Fatalf("Failed to get session: %v", err)
	}
	fmt.Printf("   Session ID: %s\n", sessionResp.SessionData.ID)
	fmt.Printf("   Title: %s\n", sessionResp.SessionData.Title)
	fmt.Printf("   Created: %s\n", sessionResp.SessionData.CreatedAt.Format(time.RFC3339))
	// Note: UpdatedAt field is not available in the current API
	fmt.Printf("   Type: %s\n\n", sessionResp.SessionData.SessionType)

	// 4. Send messages to create activity
	fmt.Println("4. Sending messages to create activity...")
	messages := []string{
		"What is 2 + 2?",
		"What is the capital of France?",
		"Tell me a short joke",
	}
	for i, text := range messages {
		fmt.Printf("   Sending message %d: %s\n", i+1, text)
		_, err := client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
			Text: text,
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
		time.Sleep(500 * time.Millisecond) // Brief delay between messages
	}
	fmt.Println("   Messages sent (processing asynchronously)\n")

	// 5. Get session messages
	fmt.Println("5. Retrieving session messages...")
	time.Sleep(2 * time.Second) // Wait for some processing
	// Note: ListSessionMessages is not available, using GetMessageHistory
	messagesResp, err := client.Messages.GetMessageHistory(ctx, nil, nil)
	if err != nil {
		log.Fatalf("Failed to get messages: %v", err)
	}
	fmt.Printf("   Found %d message(s) in history\n", len(messagesResp.BackendMessages))
	for i, msg := range messagesResp.BackendMessages {
		if i >= 5 {
			break // Show only first 5
		}
		role := msg.Role
		content := msg.UserInput
		if role == "assistant" && msg.AssistantResponse != nil {
			content = *msg.AssistantResponse
		}
		fmt.Printf("   [%d] %s: %s\n", i+1, role, truncate(content, 50))
	}
	fmt.Println()

	// 6. Create a forked session at message index 2
	if len(messagesResp.BackendMessages) >= 2 {
		fmt.Println("6. Forking session at message index 2...")
		messageIndex := int64(1) // Fork after second message
		forkResp, err := client.Sessions.ForkSession(ctx, sessionID, operations.ForkSessionRequestBody{
			MessageIndex: messageIndex,
			Title:        mix.String("Forked Session - Alternative Path"),
		})
		if err != nil {
			log.Fatalf("Failed to fork session: %v", err)
		}
		forkedSessionID := forkResp.SessionData.ID
		createdSessions = append(createdSessions, forkedSessionID)
		fmt.Printf("   Forked session created: %s\n", forkedSessionID)
		fmt.Printf("   Title: %s\n", forkResp.SessionData.Title)
		fmt.Printf("   Type: %s\n\n", forkResp.SessionData.SessionType)
	}

	// 7. Update session callbacks
	fmt.Println("7. Updating session callbacks...")
	updateResp, err := client.Sessions.UpdateSessionCallbacks(ctx, sessionID, operations.UpdateSessionCallbacksRequestBody{
		Callbacks: []components.Callback{
			{
				Type: components.CallbackTypeBashScript,
				// Note: Config field is not available in the current API
			},
		},
	})
	if err != nil {
		log.Printf("Failed to update callbacks: %v", err)
	} else {
		fmt.Printf("   Callbacks updated (Status: %d)\n\n", updateResp.HTTPMeta.Response.StatusCode)
	}

	// 8. Export session
	fmt.Println("8. Exporting session...")
	exportResp, err := client.Sessions.ExportSession(ctx, sessionID)
	if err != nil {
		log.Fatalf("Failed to export session: %v", err)
	}
	fmt.Printf("   Session ID: %s\n", exportResp.ExportSession.ID)
	fmt.Printf("   Title: %s\n", exportResp.ExportSession.Title)
	fmt.Printf("   Messages: %d\n", len(exportResp.ExportSession.Messages))
	fmt.Printf("   Exported at: %s\n\n", time.Now().Format(time.RFC3339))

	// 9. Session statistics
	fmt.Println("9. Session statistics:")
	sessionResp2, err := client.Sessions.GetSession(ctx, sessionID)
	if err != nil {
		log.Fatalf("Failed to get session: %v", err)
	}
	// Note: TotalTokens and TotalCost fields are not available in SessionData
	// They are available in ExportSession from the export endpoint
	fmt.Printf("   Session ID: %s\n", sessionResp2.SessionData.ID)
	fmt.Printf("   Title: %s\n", sessionResp2.SessionData.Title)
	fmt.Printf("   Type: %s\n", sessionResp2.SessionData.SessionType)
	fmt.Println()

	// 10. Cancel processing (if any)
	fmt.Println("10. Testing cancel processing...")
	// Note: CancelProcessing method is not available in the current API
	fmt.Println("   Skipping cancel processing test (method not available)")
	fmt.Println()

	// 11. Rewind session (delete messages after specific message ID)
	fmt.Println("11. Rewinding session...")
	// Note: RewindSession requires a MessageID, not MessageIndex
	// We need to get a message ID from the session to use
	if len(messagesResp.BackendMessages) > 1 {
		// Rewind to the first message
		firstMessageID := messagesResp.BackendMessages[0].ID
		rewindResp, err := client.Sessions.RewindSession(ctx, sessionID, operations.RewindSessionRequestBody{
			MessageID: firstMessageID,
		})
		if err != nil {
			log.Printf("Failed to rewind session: %v", err)
		} else {
			fmt.Printf("   Session rewound (Status: %d)\n", rewindResp.HTTPMeta.Response.StatusCode)
			fmt.Printf("   Messages after ID %s have been deleted\n", firstMessageID)
		}
	} else {
		fmt.Println("   Not enough messages to demonstrate rewind")
	}
	fmt.Println()

	fmt.Println("=== Sessions Example Completed Successfully! ===")
}

// truncate truncates a string to maxLen characters
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
