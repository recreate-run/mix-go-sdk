package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/recreate-run/mix-go-sdk"
)

func main() {
	serverURL := os.Getenv("MIX_SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8088"
	}

	client := mix.New(mix.WithServerURL(serverURL))
	ctx := context.Background()

	fmt.Println("=== Mix Go SDK - Tools Example ===\n")

	// 1. List all available LLM tools
	fmt.Println("1. Discovering available LLM tools...")
	toolsResp, err := client.Tools.ListLLMTools(ctx)
	if err != nil {
		log.Fatalf("Failed to list LLM tools: %v", err)
	}
	fmt.Printf("   Found %d tool(s):\n", len(toolsResp.Object.Tools))
	for i, tool := range toolsResp.Object.Tools {
		if tool.Name != nil {
			fmt.Printf("   [%d] %s\n", i+1, *tool.Name)
		}
		if tool.Description != nil {
			fmt.Printf("       Description: %s\n", *tool.Description)
		}
		if tool.Parameters != nil {
			fmt.Printf("       Has parameters: Yes\n")
		}
		if len(tool.Required) > 0 {
			fmt.Printf("       Required fields: %v\n", tool.Required)
		}
		fmt.Println()
	}

	// 2. Get tools status
	fmt.Println("2. Getting tools status...")
	statusResp, err := client.Tools.GetToolsStatus(ctx)
	if err != nil {
		log.Printf("Failed to get tools status: %v", err)
	} else if statusResp.Object != nil {
		fmt.Println("   Tools status response received")
		// Note: The actual structure of GetToolsStatus response may vary
		// Display the raw response for inspection
		fmt.Printf("   Status data: %+v\n", statusResp.Object)
	}
	fmt.Println()

	// 3. Check tool credentials status
	fmt.Println("3. Checking tool credentials status...")
	credsResp, err := client.Tools.GetToolCredentialsStatus(ctx)
	if err != nil {
		log.Printf("Failed to get tool credentials status: %v", err)
	} else if credsResp.Object != nil {
		fmt.Println("   Tool credentials status by category:")
		if credsResp.Object.Categories != nil {
			for categoryName, category := range credsResp.Object.Categories {
				fmt.Printf("   - %s", categoryName)
				if category.DisplayName != nil {
					fmt.Printf(" (%s)", *category.DisplayName)
				}
				fmt.Println()
				if len(category.Tools) > 0 {
					for i, tool := range category.Tools {
						if i >= 3 {
							fmt.Printf("     ... and %d more tools\n", len(category.Tools)-3)
							break
						}
						if tool.DisplayName != nil {
							fmt.Printf("     • %s", *tool.DisplayName)
							if tool.Authenticated != nil {
								if *tool.Authenticated {
									fmt.Printf(" (authenticated)")
								} else {
									fmt.Printf(" (not authenticated)")
								}
							}
							fmt.Println()
						}
					}
				}
			}
		}
	}
	fmt.Println()

	// 4. Analyze tools by name
	fmt.Println("4. Analyzing available tools...")
	fmt.Printf("   Total tools: %d\n", len(toolsResp.Object.Tools))
	fmt.Println("   Sample tools:")
	for i, tool := range toolsResp.Object.Tools {
		if i >= 10 {
			fmt.Printf("   ... and %d more\n", len(toolsResp.Object.Tools)-10)
			break
		}
		if tool.Name != nil {
			fmt.Printf("   - %s", *tool.Name)
			if tool.Description != nil {
				// Truncate description if too long
				desc := *tool.Description
				if len(desc) > 60 {
					desc = desc[:57] + "..."
				}
				fmt.Printf(": %s", desc)
			}
			fmt.Println()
		}
	}
	fmt.Println()

	// 5. Analyze tools with parameters
	fmt.Println("5. Analyzing tools with parameters...")
	var toolsWithParams []string
	var requiredParamsCount int
	for _, tool := range toolsResp.Object.Tools {
		if tool.Parameters != nil && tool.Name != nil {
			toolsWithParams = append(toolsWithParams, *tool.Name)
			if len(tool.Required) > 0 {
				requiredParamsCount++
			}
		}
	}
	fmt.Printf("   Tools with parameters: %d\n", len(toolsWithParams))
	fmt.Printf("   Tools with required parameters: %d\n", requiredParamsCount)
	fmt.Println("   Sample tools with parameters:")
	for i, toolName := range toolsWithParams {
		if i >= 5 {
			fmt.Printf("   ... and %d more\n", len(toolsWithParams)-5)
			break
		}
		fmt.Printf("   - %s\n", toolName)
	}
	fmt.Println()

	// 6. Tools capability summary
	fmt.Println("6. Tools capability summary:")
	fmt.Printf("   Total tools available: %d\n", len(toolsResp.Object.Tools))
	fmt.Printf("   Tools with parameters: %d\n", len(toolsWithParams))
	fmt.Printf("   Tools with required parameters: %d\n", requiredParamsCount)

	// Count tools with descriptions
	toolsWithDesc := 0
	for _, tool := range toolsResp.Object.Tools {
		if tool.Description != nil {
			toolsWithDesc++
		}
	}
	fmt.Printf("   Tools with descriptions: %d\n", toolsWithDesc)
	fmt.Println()

	fmt.Println("=== Tools Example Completed Successfully! ===")
	fmt.Println("\nKey Points:")
	fmt.Println("  - Tools extend AI capabilities for various tasks")
	fmt.Println("  - Each tool has a name, description, and optional parameters")
	fmt.Println("  - Some tools have required parameters that must be provided")
	fmt.Println("  - Use ListLLMTools() to discover all available tools")
	fmt.Println("  - Use GetToolsStatus() to check overall tool availability")
	fmt.Println("  - Use GetToolCredentialsStatus() to verify authentication status")
}
