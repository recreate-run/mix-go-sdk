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
	serverURL := os.Getenv("MIX_SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8088"
	}

	client := mix.New(mix.WithServerURL(serverURL))
	ctx := context.Background()

	fmt.Println("=== Mix Go SDK - Permissions Example ===\n")

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

	// 1. Create a session
	fmt.Println("1. Creating a session...")
	createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
		Title: "Permissions Example Session",
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	sessionID := createResp.SessionData.ID
	createdSessions = append(createdSessions, sessionID)
	fmt.Printf("   Created session: %s\n\n", sessionID)

	// 2. Send a message that might trigger permission requests
	fmt.Println("2. Sending a message that might trigger permission requests...")
	message := "Can you help me search for the latest news about AI? Use web search if you need to ask for permission."
	fmt.Printf("   User: %s\n", message)
	_, err = client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
		Text: message,
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Println("   Message sent (processing asynchronously)")
	fmt.Println("   Note: In a real application, permission requests would be handled via SSE stream")
	fmt.Println()

	// 3. Grant permission (example)
	fmt.Println("3. Granting permission (example)...")
	fmt.Println("   Note: Permission ID would come from SSE permission_request event")
	fmt.Println("   Example usage:")
	fmt.Println("   permissionID := \"perm_123abc\"")
	fmt.Println("   grantResp, err := client.Permissions.GrantPermission(ctx, permissionID)")
	fmt.Println()

	// Uncomment below to actually grant a permission (requires valid permission ID)
	// permissionID := "perm_123abc" // Replace with actual permission ID from SSE stream
	// grantResp, err := client.Permissions.GrantPermission(ctx, permissionID)
	// if err != nil {
	// 	log.Printf("Failed to grant permission: %v", err)
	// } else {
	// 	fmt.Printf("   Permission granted (Status: %d)\n", grantResp.HTTPMeta.Response.StatusCode)
	// }

	// 4. Deny permission (example)
	fmt.Println("4. Denying permission (example)...")
	fmt.Println("   Note: Permission ID would come from SSE permission_request event")
	fmt.Println("   Example usage:")
	fmt.Println("   permissionID := \"perm_456def\"")
	fmt.Println("   denyResp, err := client.Permissions.DenyPermission(ctx, permissionID)")
	fmt.Println()

	// Uncomment below to actually deny a permission (requires valid permission ID)
	// permissionID := "perm_456def" // Replace with actual permission ID from SSE stream
	// denyResp, err := client.Permissions.DenyPermission(ctx, permissionID)
	// if err != nil {
	// 	log.Printf("Failed to deny permission: %v", err)
	// } else {
	// 	fmt.Printf("   Permission denied (Status: %d)\n", denyResp.HTTPMeta.Response.StatusCode)
	// }

	// 5. Permission workflow explanation
	fmt.Println("5. Permission workflow:")
	fmt.Println("   a) Send a message that requires permission (e.g., web search, file access)")
	fmt.Println("   b) Listen to SSE stream for 'permission_request' events")
	fmt.Println("   c) Extract permission ID from the event")
	fmt.Println("   d) Call GrantPermission(permissionID) or DenyPermission(permissionID)")
	fmt.Println("   e) Processing continues based on permission decision")
	fmt.Println()

	// 6. Advanced permission handling with custom parameters
	fmt.Println("6. Advanced permission handling:")
	fmt.Println("   Permissions API supports custom parameters:")
	fmt.Println("   - Custom timeout values")
	fmt.Println("   - Custom headers for authentication")
	fmt.Println("   - Server URL override for different environments")
	fmt.Println()

	// Example with custom timeout
	fmt.Println("   Example with custom timeout:")
	fmt.Println("   import \"time\"")
	fmt.Println("   timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)")
	fmt.Println("   defer cancel()")
	fmt.Println("   grantResp, err := client.Permissions.GrantPermission(timeoutCtx, permissionID)")
	fmt.Println()

	// 7. Integration with streaming
	fmt.Println("7. Complete permission handling example with streaming:")
	fmt.Println("   See streaming_example.go for handling permission_request events")
	fmt.Println("   When you receive a permission_request event:")
	fmt.Println("   - Parse the event data to get permission ID and details")
	fmt.Println("   - Show permission request to user (UI/CLI prompt)")
	fmt.Println("   - Based on user response, call GrantPermission or DenyPermission")
	fmt.Println("   - Continue listening to stream for processing updates")
	fmt.Println()

	fmt.Println("=== Permissions Example Completed Successfully! ===")
	fmt.Println("\nKey Points:")
	fmt.Println("  - Permissions are requested during message processing")
	fmt.Println("  - Permission requests come via SSE 'permission_request' events")
	fmt.Println("  - Use GrantPermission(id) or DenyPermission(id) to respond")
	fmt.Println("  - Permission decisions affect how the AI continues processing")
	fmt.Println("  - Combine with streaming for real-time permission handling")
}
