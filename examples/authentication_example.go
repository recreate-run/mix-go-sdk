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

	fmt.Println("=== Mix Go SDK - Authentication Example ===\n")

	// 1. Check initial authentication status
	fmt.Println("1. Checking current authentication status...")
	statusResp, err := client.Authentication.GetAuthStatus(ctx)
	if err != nil {
		log.Fatalf("Failed to get auth status: %v", err)
	}
	fmt.Println("   Authentication status by provider:")
	if statusResp.Object != nil && statusResp.Object.Providers != nil {
		providers := []string{"anthropic", "openai", "gemini", "brave", "openrouter"}
		for _, providerName := range providers {
			if provider, ok := statusResp.Object.Providers[providerName]; ok {
				authenticated := false
				if provider.Authenticated != nil {
					authenticated = *provider.Authenticated
				}
				fmt.Printf("   - %s: %t", providerName, authenticated)
				if provider.AuthMethod != nil {
					fmt.Printf(" (method: %s)", *provider.AuthMethod)
				}
				fmt.Println()
			}
		}
	}
	fmt.Println()

	// 2. Store API key for OpenRouter (example provider)
	openrouterKey := os.Getenv("OPENROUTER_API_KEY")
	if openrouterKey != "" {
		fmt.Println("2. Storing OpenRouter API key...")
		storeResp, err := client.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
			Provider: "openrouter",
			APIKey:   openrouterKey,
		})
		if err != nil {
			log.Printf("Failed to store OpenRouter API key: %v", err)
		} else {
			fmt.Printf("   API key stored successfully (Status: %d)\n", storeResp.HTTPMeta.Response.StatusCode)
		}
	} else {
		fmt.Println("2. Skipping OpenRouter API key storage (OPENROUTER_API_KEY not set)")
	}
	fmt.Println()

	// 3. Store API key for Anthropic
	anthropicKey := os.Getenv("ANTHROPIC_API_KEY")
	if anthropicKey != "" {
		fmt.Println("3. Storing Anthropic API key...")
		storeResp, err := client.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
			Provider: "anthropic",
			APIKey:   anthropicKey,
		})
		if err != nil {
			log.Printf("Failed to store Anthropic API key: %v", err)
		} else {
			fmt.Printf("   API key stored successfully (Status: %d)\n", storeResp.HTTPMeta.Response.StatusCode)
		}
	} else {
		fmt.Println("3. Skipping Anthropic API key storage (ANTHROPIC_API_KEY not set)")
	}
	fmt.Println()

	// 4. Store API key for OpenAI
	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey != "" {
		fmt.Println("4. Storing OpenAI API key...")
		storeResp, err := client.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
			Provider: "openai",
			APIKey:   openaiKey,
		})
		if err != nil {
			log.Printf("Failed to store OpenAI API key: %v", err)
		} else {
			fmt.Printf("   API key stored successfully (Status: %d)\n", storeResp.HTTPMeta.Response.StatusCode)
		}
	} else {
		fmt.Println("4. Skipping OpenAI API key storage (OPENAI_API_KEY not set)")
	}
	fmt.Println()

	// 5. Store API key for Gemini
	geminiKey := os.Getenv("GEMINI_API_KEY")
	if geminiKey != "" {
		fmt.Println("5. Storing Gemini API key...")
		storeResp, err := client.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
			Provider: "gemini",
			APIKey:   geminiKey,
		})
		if err != nil {
			log.Printf("Failed to store Gemini API key: %v", err)
		} else {
			fmt.Printf("   API key stored successfully (Status: %d)\n", storeResp.HTTPMeta.Response.StatusCode)
		}
	} else {
		fmt.Println("5. Skipping Gemini API key storage (GEMINI_API_KEY not set)")
	}
	fmt.Println()

	// 6. Store API key for Brave
	braveKey := os.Getenv("BRAVE_API_KEY")
	if braveKey != "" {
		fmt.Println("6. Storing Brave API key...")
		storeResp, err := client.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
			Provider: "brave",
			APIKey:   braveKey,
		})
		if err != nil {
			log.Printf("Failed to store Brave API key: %v", err)
		} else {
			fmt.Printf("   API key stored successfully (Status: %d)\n", storeResp.HTTPMeta.Response.StatusCode)
		}
	} else {
		fmt.Println("6. Skipping Brave API key storage (BRAVE_API_KEY not set)")
	}
	fmt.Println()

	// 7. Check authentication status after storing keys
	fmt.Println("7. Checking authentication status after storing keys...")
	statusResp2, err := client.Authentication.GetAuthStatus(ctx)
	if err != nil {
		log.Fatalf("Failed to get auth status: %v", err)
	}
	fmt.Println("   Updated authentication status:")
	if statusResp2.Object != nil && statusResp2.Object.Providers != nil {
		providers := []string{"anthropic", "openai", "gemini", "brave", "openrouter"}
		for _, providerName := range providers {
			if provider, ok := statusResp2.Object.Providers[providerName]; ok {
				authenticated := false
				if provider.Authenticated != nil {
					authenticated = *provider.Authenticated
				}
				fmt.Printf("   - %s: %t", providerName, authenticated)
				if provider.AuthMethod != nil {
					fmt.Printf(" (method: %s)", *provider.AuthMethod)
				}
				fmt.Println()
			}
		}
	}
	fmt.Println()

	// 8. Validate preferred provider
	fmt.Println("8. Validating preferred provider...")
	validateResp, err := client.Authentication.ValidatePreferredProvider(ctx)
	if err != nil {
		log.Printf("Failed to validate preferred provider: %v", err)
	} else if validateResp.Object != nil {
		isValid := false
		if validateResp.Object.Valid != nil {
			isValid = *validateResp.Object.Valid
		}
		provider := "unknown"
		if validateResp.Object.Provider != nil {
			provider = *validateResp.Object.Provider
		}
		fmt.Printf("   Provider '%s' is valid: %t\n", provider, isValid)
		if validateResp.Object.AuthMethod != nil {
			fmt.Printf("   Auth method: %s\n", *validateResp.Object.AuthMethod)
		}
		if validateResp.Object.Message != nil {
			fmt.Printf("   Message: %s\n", *validateResp.Object.Message)
		}
	}
	fmt.Println()

	// 9. OAuth flow initiation (example - won't complete without browser)
	fmt.Println("9. Demonstrating OAuth flow initiation...")
	fmt.Println("   Note: OAuth requires browser interaction, showing URL only")
	oauthProvider := "anthropic" // Example OAuth provider (currently only anthropic supported)
	oauthResp, err := client.Authentication.StartOAuthFlow(ctx, oauthProvider)
	if err != nil {
		log.Printf("Failed to start OAuth flow: %v", err)
	} else if oauthResp.Object != nil {
		if oauthResp.Object.AuthURL != nil {
			fmt.Printf("   OAuth URL: %s\n", *oauthResp.Object.AuthURL)
			fmt.Println("   (In a real application, redirect user to this URL)")
		}
		if oauthResp.Object.State != nil {
			fmt.Printf("   State: %s\n", *oauthResp.Object.State)
		}
		if oauthResp.Object.Message != nil {
			fmt.Printf("   Message: %s\n", *oauthResp.Object.Message)
		}
	}
	fmt.Println()

	// 10. Get OAuth health status
	fmt.Println("10. Checking OAuth health status...")
	healthResp, err := client.Authentication.GetOAuthHealth(ctx)
	if err != nil {
		log.Printf("Failed to get OAuth health: %v", err)
	} else {
		fmt.Printf("   OAuth health status (Status: %d)\n", healthResp.HTTPMeta.Response.StatusCode)
		if healthResp.Object != nil {
			fmt.Printf("   Details: %+v\n", healthResp.Object)
		}
	}
	fmt.Println()

	// 11. Refresh OAuth tokens
	fmt.Println("11. Testing OAuth token refresh...")
	refreshResp, err := client.Internal.RefreshOAuthTokens(ctx)
	if err != nil {
		log.Printf("No OAuth tokens to refresh or refresh failed: %v", err)
	} else {
		fmt.Printf("   OAuth tokens refreshed (Status: %d)\n", refreshResp.HTTPMeta.Response.StatusCode)
	}
	fmt.Println()

	// 12. Delete credentials (cleanup example)
	fmt.Println("12. Credential deletion example...")
	fmt.Println("   Note: Uncommenting the code below will delete stored credentials")
	fmt.Println("   // To delete OpenRouter credentials:")
	fmt.Println("   // client.Authentication.DeleteCredentials(ctx, \"openrouter\")")
	fmt.Println()

	// Uncomment below to actually delete credentials
	// deleteProvider := "openrouter"
	// fmt.Printf("   Deleting %s credentials...\n", deleteProvider)
	// deleteResp, err := client.Authentication.DeleteCredentials(ctx, deleteProvider)
	// if err != nil {
	// 	log.Printf("Failed to delete credentials: %v", err)
	// } else {
	// 	fmt.Printf("   Credentials deleted (Status: %d)\n", deleteResp.HTTPMeta.Response.StatusCode)
	// }

	fmt.Println("=== Authentication Example Completed Successfully! ===")
	fmt.Println("\nTips:")
	fmt.Println("  - Set environment variables (ANTHROPIC_API_KEY, OPENAI_API_KEY, etc.) to test API key storage")
	fmt.Println("  - OAuth flows require browser interaction for completion")
	fmt.Println("  - Use GetAuthStatus() to verify which providers are configured")
	fmt.Println("  - Credentials are stored securely and persist across sessions")
}
