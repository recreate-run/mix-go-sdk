package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
	serverURL := os.Getenv("MIX_SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8088"
	}

	client := mix.New(mix.WithServerURL(serverURL))
	ctx := context.Background()

	fmt.Println("=== Mix Go SDK - Messages Example ===\n")

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

	// 1. Get global message history
	fmt.Println("1. Retrieving global message history...")
	limit := int64(10)
	historyResp, err := client.Messages.GetMessageHistory(ctx, &limit, nil)
	if err != nil {
		log.Fatalf("Failed to get message history: %v", err)
	}
	fmt.Printf("   Found %d message(s) in global history\n", len(historyResp.BackendMessages))
	if len(historyResp.BackendMessages) > 0 {
		fmt.Println("   Recent messages:")
		for i, msg := range historyResp.BackendMessages {
			if i >= 3 {
				break
			}
			role := msg.Role
			content := msg.UserInput
			if role == "assistant" && msg.AssistantResponse != nil {
				content = *msg.AssistantResponse
			}
			fmt.Printf("   [%d] %s: %s\n", i+1, role, truncate(content, 60))
		}
	}
	fmt.Println()

	// 2. Create a session for messaging
	fmt.Println("2. Creating a new session for messaging...")
	createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
		Title: "Messages Example - Interactive Conversation",
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	sessionID := createResp.SessionData.ID
	createdSessions = append(createdSessions, sessionID)
	fmt.Printf("   Created session: %s\n", sessionID)
	fmt.Printf("   Title: %s\n\n", createResp.SessionData.Title)

	// 3. Send multiple messages to build conversation
	fmt.Println("3. Building a conversation with multiple messages...")
	conversationMessages := []string{
		"Hello! Can you introduce yourself?",
		"What can you help me with?",
		"Can you explain what the Fibonacci sequence is?",
	}

	for i, text := range conversationMessages {
		fmt.Printf("   [%d] User: %s\n", i+1, text)
		sendResp, err := client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
			Text: text,
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			continue
		}
		fmt.Printf("       Sent (Status: %d)\n", sendResp.HTTPMeta.Response.StatusCode)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()

	// 4. Wait for processing and retrieve session messages
	fmt.Println("4. Waiting for message processing...")
	time.Sleep(3 * time.Second)
	fmt.Println("   Retrieving session messages...\n")

	// Note: ListSessionMessages is not available in the current API
	// Using GetMessageHistory as an alternative approach
	messagesResp, err := client.Messages.GetMessageHistory(ctx, nil, nil)
	if err != nil {
		log.Fatalf("Failed to get message history: %v", err)
	}

	// 5. Display conversation with metadata
	fmt.Println("5. Conversation history with metadata:")
	fmt.Printf("   Total messages: %d\n\n", len(messagesResp.BackendMessages))

	for i, msg := range messagesResp.BackendMessages {
		role := msg.Role
		content := msg.UserInput
		if role == "assistant" && msg.AssistantResponse != nil {
			content = *msg.AssistantResponse
		}

		fmt.Printf("   Message %d [%s]:\n", i+1, role)
		fmt.Printf("   ID: %s\n", msg.ID)
		fmt.Printf("   Content: %s\n", truncate(content, 100))

		// Display metadata if available
		if msg.ReasoningDuration != nil {
			fmt.Printf("   Reasoning Duration: %dms\n", *msg.ReasoningDuration)
		}
		if msg.Reasoning != nil {
			fmt.Printf("   Reasoning: %s\n", truncate(*msg.Reasoning, 100))
		}

		// Display tool calls if present
		if msg.ToolCalls != nil && len(msg.ToolCalls) > 0 {
			fmt.Printf("   Tool Calls: %d\n", len(msg.ToolCalls))
			for j, toolCall := range msg.ToolCalls {
				fmt.Printf("     [%d] %s\n", j+1, toolCall.Name)
			}
		}

		fmt.Println()
	}

	// 6. Pagination example - get messages with limit and offset
	fmt.Println("6. Demonstrating pagination...")
	limit = int64(2)
	// Note: Using GetMessageHistory with limit parameter
	paginatedResp, err := client.Messages.GetMessageHistory(ctx, &limit, nil)
	if err != nil {
		log.Fatalf("Failed to get paginated messages: %v", err)
	}
	fmt.Printf("   Retrieved %d message(s) with limit=%d\n", len(paginatedResp.BackendMessages), limit)
	for i, msg := range paginatedResp.BackendMessages {
		role := msg.Role
		content := msg.UserInput
		if role == "assistant" && msg.AssistantResponse != nil {
			content = *msg.AssistantResponse
		}
		fmt.Printf("   [%d] %s: %s\n", i+1, role, truncate(content, 50))
	}
	fmt.Println()

	// 7. Send a message with tool use
	fmt.Println("7. Sending a message that might trigger tool use...")
	toolMessage := "What's the current weather in San Francisco? Use a tool if available."
	fmt.Printf("   User: %s\n", toolMessage)
	_, err = client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
		Text: toolMessage,
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	} else {
		fmt.Println("   Message sent (processing asynchronously)")
	}
	fmt.Println()

	// 8. Message statistics
	fmt.Println("8. Analyzing message statistics...")
	time.Sleep(2 * time.Second)

	// Note: Using GetMessageHistory for statistics
	statsResp, err := client.Messages.GetMessageHistory(ctx, nil, nil)
	if err != nil {
		log.Fatalf("Failed to get messages for stats: %v", err)
	}

	var messagesWithTools int
	var userMessages int
	var assistantMessages int

	for _, msg := range statsResp.BackendMessages {
		if msg.ToolCalls != nil && len(msg.ToolCalls) > 0 {
			messagesWithTools++
		}
		if msg.Role == "user" {
			userMessages++
		} else if msg.Role == "assistant" {
			assistantMessages++
		}
	}

	fmt.Printf("   Total messages: %d\n", len(statsResp.BackendMessages))
	fmt.Printf("   User messages: %d\n", userMessages)
	fmt.Printf("   Assistant messages: %d\n", assistantMessages)
	fmt.Printf("   Messages with tools: %d\n", messagesWithTools)
	fmt.Println()

	// 9. Test conversation continuity
	fmt.Println("9. Testing conversation continuity...")
	continuityMessage := "Based on what we discussed earlier, can you remind me what the Fibonacci sequence is?"
	fmt.Printf("   User: %s\n", continuityMessage)
	_, err = client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
		Text: continuityMessage,
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	} else {
		fmt.Println("   Message sent - AI will use conversation context")
	}
	fmt.Println()

	fmt.Println("=== Messages Example Completed Successfully! ===")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
