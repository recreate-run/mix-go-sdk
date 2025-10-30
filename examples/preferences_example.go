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

	fmt.Println("=== Mix Go SDK - Preferences Example ===\n")

	// 1. Get current preferences
	fmt.Println("1. Getting current user preferences...")
	prefsResp, err := client.Preferences.GetPreferences(ctx)
	if err != nil {
		log.Fatalf("Failed to get preferences: %v", err)
	}
	fmt.Println("   Current preferences:")
	if prefs, ok := prefsResp.Object.Preferences.Get(); ok {
		if prefs.PreferredProvider != nil {
			fmt.Printf("   - Preferred Provider: %s\n", *prefs.PreferredProvider)
		}
		if prefs.MainAgentModel != nil {
			fmt.Printf("   - Main Agent Model: %s\n", *prefs.MainAgentModel)
		}
		if prefs.SubAgentModel != nil {
			fmt.Printf("   - Sub Agent Model: %s\n", *prefs.SubAgentModel)
		}
		if prefs.MainAgentMaxTokens != nil {
			fmt.Printf("   - Main Agent Max Tokens: %d\n", *prefs.MainAgentMaxTokens)
		}
		if prefs.SubAgentMaxTokens != nil {
			fmt.Printf("   - Sub Agent Max Tokens: %d\n", *prefs.SubAgentMaxTokens)
		}
		if prefs.MainAgentReasoningEffort != nil {
			fmt.Printf("   - Main Agent Reasoning Effort: %s\n", *prefs.MainAgentReasoningEffort)
		}
		if prefs.SubAgentReasoningEffort != nil {
			fmt.Printf("   - Sub Agent Reasoning Effort: %s\n", *prefs.SubAgentReasoningEffort)
		}
	}
	fmt.Println()

	// 2. Get available providers and models
	fmt.Println("2. Discovering available providers and models...")
	fmt.Println("   Available providers from preferences response:")
	if prefsResp.Object != nil && len(prefsResp.Object.AvailableProviders) > 0 {
		count := 0
		for providerName, providerInfo := range prefsResp.Object.AvailableProviders {
			if count >= 5 {
				fmt.Printf("   ... and %d more\n", len(prefsResp.Object.AvailableProviders)-5)
				break
			}
			count++
			fmt.Printf("   [%d] %s\n", count, providerName)
			if providerInfo.Models != nil && len(providerInfo.Models) > 0 {
				fmt.Printf("       Models: ")
				for j, model := range providerInfo.Models {
					if j >= 3 {
						fmt.Printf("... (%d total)", len(providerInfo.Models))
						break
					}
					if j > 0 {
						fmt.Printf(", ")
					}
					fmt.Printf("%s", model)
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Println("   No providers available")
	}
	fmt.Println()

	// 3. Update preferences - Main agent model
	fmt.Println("3. Updating main agent model preference...")
	updateMainAgent := "claude-3-5-sonnet-20241022" // Example model
	updateResp, err := client.Preferences.UpdatePreferences(ctx, operations.UpdatePreferencesRequest{
		MainAgentModel: &updateMainAgent,
	})
	if err != nil {
		log.Printf("Failed to update main agent preference: %v", err)
	} else if updateResp.Object != nil {
		fmt.Printf("   Updated main agent model to: %s\n", *updateResp.Object.MainAgentModel)
		fmt.Printf("   Status: %d\n", updateResp.HTTPMeta.Response.StatusCode)
	}
	fmt.Println()

	// 4. Update preferences - Sub agent model
	fmt.Println("4. Updating sub agent model preference...")
	updateSubAgent := "claude-3-5-haiku-20241022" // Example model
	updateResp2, err := client.Preferences.UpdatePreferences(ctx, operations.UpdatePreferencesRequest{
		SubAgentModel: &updateSubAgent,
	})
	if err != nil {
		log.Printf("Failed to update sub agent preference: %v", err)
	} else if updateResp2.Object != nil {
		fmt.Printf("   Updated sub agent model to: %s\n", *updateResp2.Object.SubAgentModel)
		fmt.Printf("   Status: %d\n", updateResp2.HTTPMeta.Response.StatusCode)
	}
	fmt.Println()

	// 5. Update preferences - Main agent max tokens
	fmt.Println("5. Updating main agent max tokens preference...")
	mainAgentMaxTokens := int64(4096)
	updateResp3, err := client.Preferences.UpdatePreferences(ctx, operations.UpdatePreferencesRequest{
		MainAgentMaxTokens: &mainAgentMaxTokens,
	})
	if err != nil {
		log.Printf("Failed to update main agent max tokens: %v", err)
	} else if updateResp3.Object != nil {
		fmt.Printf("   Updated main agent max tokens to: %d\n", *updateResp3.Object.MainAgentMaxTokens)
		fmt.Printf("   Status: %d\n", updateResp3.HTTPMeta.Response.StatusCode)
	}
	fmt.Println()

	// 6. Update preferences - Main agent reasoning effort
	fmt.Println("6. Updating main agent reasoning effort preference...")
	mainAgentReasoningEffort := "medium" // Options: low, medium, high
	updateResp4, err := client.Preferences.UpdatePreferences(ctx, operations.UpdatePreferencesRequest{
		MainAgentReasoningEffort: &mainAgentReasoningEffort,
	})
	if err != nil {
		log.Printf("Failed to update main agent reasoning effort: %v", err)
	} else if updateResp4.Object != nil {
		fmt.Printf("   Updated main agent reasoning effort to: %s\n", *updateResp4.Object.MainAgentReasoningEffort)
		fmt.Printf("   Status: %d\n", updateResp4.HTTPMeta.Response.StatusCode)
	}
	fmt.Println()

	// 7. Update multiple preferences at once
	fmt.Println("7. Updating multiple preferences simultaneously...")
	preferredProvider := "anthropic"
	mainAgentModel := "claude-3-5-sonnet-20241022"
	subAgentModel := "claude-3-5-haiku-20241022"
	newMainAgentMaxTokens := int64(8192)
	newSubAgentMaxTokens := int64(4096)
	newMainAgentReasoningEffort := "high"
	newSubAgentReasoningEffort := "medium"

	updateResp5, err := client.Preferences.UpdatePreferences(ctx, operations.UpdatePreferencesRequest{
		PreferredProvider:        &preferredProvider,
		MainAgentModel:           &mainAgentModel,
		SubAgentModel:            &subAgentModel,
		MainAgentMaxTokens:       &newMainAgentMaxTokens,
		SubAgentMaxTokens:        &newSubAgentMaxTokens,
		MainAgentReasoningEffort: &newMainAgentReasoningEffort,
		SubAgentReasoningEffort:  &newSubAgentReasoningEffort,
	})
	if err != nil {
		log.Printf("Failed to update preferences: %v", err)
	} else if updateResp5.Object != nil {
		fmt.Println("   Updated preferences:")
		if updateResp5.Object.PreferredProvider != nil {
			fmt.Printf("   - Preferred Provider: %s\n", *updateResp5.Object.PreferredProvider)
		}
		if updateResp5.Object.MainAgentModel != nil {
			fmt.Printf("   - Main Agent Model: %s\n", *updateResp5.Object.MainAgentModel)
		}
		if updateResp5.Object.SubAgentModel != nil {
			fmt.Printf("   - Sub Agent Model: %s\n", *updateResp5.Object.SubAgentModel)
		}
		if updateResp5.Object.MainAgentMaxTokens != nil {
			fmt.Printf("   - Main Agent Max Tokens: %d\n", *updateResp5.Object.MainAgentMaxTokens)
		}
		if updateResp5.Object.SubAgentMaxTokens != nil {
			fmt.Printf("   - Sub Agent Max Tokens: %d\n", *updateResp5.Object.SubAgentMaxTokens)
		}
		if updateResp5.Object.MainAgentReasoningEffort != nil {
			fmt.Printf("   - Main Agent Reasoning Effort: %s\n", *updateResp5.Object.MainAgentReasoningEffort)
		}
		if updateResp5.Object.SubAgentReasoningEffort != nil {
			fmt.Printf("   - Sub Agent Reasoning Effort: %s\n", *updateResp5.Object.SubAgentReasoningEffort)
		}
	}
	fmt.Println()

	// 8. Provider-specific model selection
	fmt.Println("8. Provider-specific configuration examples:")
	fmt.Println("   Anthropic models:")
	fmt.Println("     - claude-3-5-sonnet-20241022 (most capable)")
	fmt.Println("     - claude-3-5-haiku-20241022 (fast, efficient)")
	fmt.Println("   OpenAI models:")
	fmt.Println("     - gpt-4-turbo")
	fmt.Println("     - gpt-3.5-turbo")
	fmt.Println("   Gemini models:")
	fmt.Println("     - gemini-pro")
	fmt.Println("   Note: Available models depend on your authentication")
	fmt.Println()

	// 9. Get updated preferences
	fmt.Println("9. Verifying updated preferences...")
	prefsResp2, err := client.Preferences.GetPreferences(ctx)
	if err != nil {
		log.Fatalf("Failed to get preferences: %v", err)
	}
	fmt.Println("   Current preferences:")
	if prefs, ok := prefsResp2.Object.Preferences.Get(); ok {
		if prefs.PreferredProvider != nil {
			fmt.Printf("   - Preferred Provider: %s\n", *prefs.PreferredProvider)
		}
		if prefs.MainAgentModel != nil {
			fmt.Printf("   - Main Agent Model: %s\n", *prefs.MainAgentModel)
		}
		if prefs.SubAgentModel != nil {
			fmt.Printf("   - Sub Agent Model: %s\n", *prefs.SubAgentModel)
		}
		if prefs.MainAgentMaxTokens != nil {
			fmt.Printf("   - Main Agent Max Tokens: %d\n", *prefs.MainAgentMaxTokens)
		}
		if prefs.SubAgentMaxTokens != nil {
			fmt.Printf("   - Sub Agent Max Tokens: %d\n", *prefs.SubAgentMaxTokens)
		}
		if prefs.MainAgentReasoningEffort != nil {
			fmt.Printf("   - Main Agent Reasoning Effort: %s\n", *prefs.MainAgentReasoningEffort)
		}
		if prefs.SubAgentReasoningEffort != nil {
			fmt.Printf("   - Sub Agent Reasoning Effort: %s\n", *prefs.SubAgentReasoningEffort)
		}
	}
	fmt.Println()

	// 10. Reset preferences to defaults
	fmt.Println("10. Resetting preferences to defaults...")
	resetResp, err := client.Preferences.ResetPreferences(ctx)
	if err != nil {
		log.Printf("Failed to reset preferences: %v", err)
	} else {
		fmt.Printf("   Preferences reset (Status: %d)\n", resetResp.HTTPMeta.Response.StatusCode)
		if resetResp.Object != nil {
			fmt.Println("   Default preferences:")
			if resetResp.Object.PreferredProvider != nil {
				fmt.Printf("   - Preferred Provider: %s\n", *resetResp.Object.PreferredProvider)
			}
			if resetResp.Object.MainAgentModel != nil {
				fmt.Printf("   - Main Agent Model: %s\n", *resetResp.Object.MainAgentModel)
			}
			if resetResp.Object.SubAgentModel != nil {
				fmt.Printf("   - Sub Agent Model: %s\n", *resetResp.Object.SubAgentModel)
			}
		}
	}
	fmt.Println()

	fmt.Println("=== Preferences Example Completed Successfully! ===")
	fmt.Println("\nKey Points:")
	fmt.Println("  - Preferences control which AI models are used")
	fmt.Println("  - Main agent handles primary tasks, sub agent handles delegated tasks")
	fmt.Println("  - Max tokens controls maximum response length for each agent")
	fmt.Println("  - Reasoning effort affects response quality vs. speed for each agent")
	fmt.Println("  - Available models depend on configured API keys")
	fmt.Println("  - Available providers are included in GetPreferences() response")
}
