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

	fmt.Println("=== Mix Go SDK - System Example ===\n")

	// 1. Health check
	fmt.Println("1. Checking system health...")
	healthResp, err := client.System.HealthCheck(ctx)
	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}
	fmt.Printf("   Status: %s\n", *healthResp.Object.Status)
	if healthResp.Object.Version != nil {
		fmt.Printf("   Version: %s\n", *healthResp.Object.Version)
	}
	if healthResp.Object.Timestamp != nil {
		fmt.Printf("   Timestamp: %s\n", *healthResp.Object.Timestamp)
	}
	fmt.Println()

	// 2. List available commands
	fmt.Println("2. Listing available system commands...")
	commandsResp, err := client.System.ListCommands(ctx)
	if err != nil {
		log.Fatalf("Failed to list commands: %v", err)
	}
	fmt.Printf("   Found %d command(s):\n", len(commandsResp.ResponseBodies))
	for i, cmd := range commandsResp.ResponseBodies {
		if cmd.Name != nil {
			fmt.Printf("   [%d] %s\n", i+1, *cmd.Name)
		}
		if cmd.Description != nil {
			fmt.Printf("       Description: %s\n", *cmd.Description)
		}
		if i >= 9 {
			fmt.Printf("   ... and %d more\n", len(commandsResp.ResponseBodies)-10)
			break
		}
	}
	fmt.Println()

	// 3. Get detailed command information
	if len(commandsResp.ResponseBodies) > 0 && commandsResp.ResponseBodies[0].Name != nil {
		firstCommand := *commandsResp.ResponseBodies[0].Name
		fmt.Printf("3. Getting details for command: %s...\n", firstCommand)
		cmdResp, err := client.System.GetCommand(ctx, firstCommand)
		if err != nil {
			log.Printf("Failed to get command details: %v", err)
		} else if cmdResp.Object != nil {
			if cmdResp.Object.Name != nil {
				fmt.Printf("   Name: %s\n", *cmdResp.Object.Name)
			}
			if cmdResp.Object.Description != nil {
				fmt.Printf("   Description: %s\n", *cmdResp.Object.Description)
			}
			if cmdResp.Object.Usage != nil {
				fmt.Printf("   Usage: %s\n", *cmdResp.Object.Usage)
			}
		}
		fmt.Println()
	}

	// 4. List MCP (Model Context Protocol) servers
	fmt.Println("4. Listing MCP servers...")
	mcpResp, err := client.System.ListMcpServers(ctx)
	if err != nil {
		log.Printf("Failed to list MCP servers: %v", err)
	} else {
		fmt.Printf("   Found %d MCP server(s):\n", len(mcpResp.ResponseBodies))
		for i, server := range mcpResp.ResponseBodies {
			fmt.Printf("   [%d] %s\n", i+1, server.Name)
			fmt.Printf("       Status: %s\n", server.Status)
			fmt.Printf("       Connected: %v\n", server.Connected)
			if server.Tools.IsSet() {
				tools, _ := server.Tools.Get()
				if tools != nil {
					fmt.Printf("       Tools: %d\n", len(*tools))
				}
			}
		}
	}
	fmt.Println()

	// 5. Analyze commands by name
	fmt.Println("5. Command summary...")
	var namedCommands int
	for _, cmd := range commandsResp.ResponseBodies {
		if cmd.Name != nil {
			namedCommands++
		}
	}
	fmt.Printf("   Commands with names: %d\n", namedCommands)
	fmt.Printf("   Total command entries: %d\n", len(commandsResp.ResponseBodies))
	fmt.Println()

	// 6. System capabilities summary
	fmt.Println("\n6. System capabilities summary:")
	fmt.Printf("   Total commands: %d\n", len(commandsResp.ResponseBodies))
	if mcpResp != nil && mcpResp.ResponseBodies != nil {
		fmt.Printf("   MCP servers: %d\n", len(mcpResp.ResponseBodies))
	}
	fmt.Printf("   Health status: %s\n", *healthResp.Object.Status)
	fmt.Println()

	// 7. System information
	fmt.Println("7. System information:")
	fmt.Println("   API Endpoints:")
	fmt.Println("   - Health: GET /health")
	fmt.Println("   - Commands: GET /commands")
	fmt.Println("   - Command Detail: GET /commands/{name}")
	fmt.Println("   - MCP Servers: GET /mcp/servers")
	fmt.Println()
	fmt.Println("   Integration points:")
	fmt.Println("   - System health monitoring")
	fmt.Println("   - Command discovery and introspection")
	fmt.Println("   - MCP server management")
	fmt.Println("   - Tool ecosystem status")
	fmt.Println()

	// 8. Health monitoring recommendations
	fmt.Println("8. Health monitoring best practices:")
	fmt.Println("   - Call HealthCheck() before critical operations")
	fmt.Println("   - Implement retry logic for failed health checks")
	fmt.Println("   - Monitor health status for service degradation")
	fmt.Println("   - Use version information for compatibility checks")
	fmt.Println("   - Check MCP server status for tool availability")
	fmt.Println()

	// 9. Command usage patterns
	fmt.Println("9. Command usage patterns:")
	fmt.Println("   - Use ListCommands() to discover available commands")
	fmt.Println("   - Use GetCommand(name) for detailed command information")
	fmt.Println("   - Check command parameters before invocation")
	fmt.Println("   - Implement command validation based on category")
	fmt.Println("   - Use command descriptions for user-facing documentation")
	fmt.Println()

	// 10. MCP integration
	fmt.Println("10. MCP (Model Context Protocol) integration:")
	fmt.Println("   - MCP servers extend AI capabilities")
	fmt.Println("   - Check server status before using MCP features")
	fmt.Println("   - Monitor server URLs for connectivity")
	fmt.Println("   - Use ListMCPServers() to discover available servers")
	fmt.Println("   - Implement fallback logic for unavailable servers")
	fmt.Println()

	fmt.Println("=== System Example Completed Successfully! ===")
	fmt.Println("\nKey Points:")
	fmt.Println("  - HealthCheck() verifies system availability")
	fmt.Println("  - ListCommands() discovers available system commands")
	fmt.Println("  - GetCommand(name) provides detailed command information")
	fmt.Println("  - ListMCPServers() shows available MCP servers")
	fmt.Println("  - Use system endpoints for monitoring and introspection")
	fmt.Println("  - Implement health checks in production applications")
}
