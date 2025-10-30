# Mix Go SDK - Examples

This directory contains comprehensive examples demonstrating all features of the Mix Go SDK.

## Quick Start

1. **Set up your environment:**
   ```bash
   export MIX_SERVER_URL="http://localhost:8088"  # Optional, defaults to localhost:8088
   export ANTHROPIC_API_KEY="your-api-key"        # Optional, for AI provider access
   ```

2. **Run an example:**
   ```bash
   go run basic_client.go
   ```

## Examples Overview

### 🚀 Getting Started

#### `example_template.go`
**Perfect starting point for new users**
- Minimal, well-commented template
- Shows basic session creation and streaming
- Copy this file to start building your own application
- Demonstrates essential patterns and cleanup

#### `basic_client.go`
**Simple end-to-end workflow**
- Health checks
- Session CRUD operations
- Basic messaging
- Synchronous pattern demonstration

### 📝 Core Functionality

#### `sessions_example.go`
**Comprehensive session lifecycle management**
- Session creation with custom configuration
- Session listing and retrieval
- Message sending and activity tracking
- Session forking at specific message indices
- Session callbacks and export
- Session rewinding (deleting messages)
- Processing cancellation
- Metadata and usage statistics

**Key Demonstrations:**
- Session types (main, forked, subagent)
- Session metadata analysis
- Token counting and cost tracking
- Complete cleanup patterns

#### `messages_example.go`
**Message operations and conversation management**
- Global message history with pagination
- Session-specific message listing
- Interactive conversation building
- Tool call analysis
- Message metadata (tokens, cost, reasoning duration)
- Conversation continuity testing
- Message statistics tracking

**Key Demonstrations:**
- Pagination strategies
- Tool integration analysis
- Context management
- Message filtering

### 🔐 Authentication & Configuration

#### `authentication_example.go`
**Multi-provider authentication**
- API key storage for multiple providers:
  - OpenRouter
  - Anthropic
  - OpenAI
  - Gemini
  - Brave
- OAuth flow initiation and callback handling
- Authentication status checking
- Preferred provider validation
- Credential deletion and cleanup
- OAuth health monitoring

**Key Demonstrations:**
- Multi-provider setup
- OAuth workflows
- Credential management
- Health status monitoring

#### `preferences_example.go`
**Model and provider configuration**
- Current preference retrieval
- Available providers discovery
- Model listing per provider
- Dual-agent configuration (main agent vs sub agent)
- Token limit settings
- Reasoning effort configuration
- Preference reset functionality

**Key Demonstrations:**
- Main agent configuration (primary tasks)
- Sub agent configuration (delegated tasks)
- Provider switching
- Model selection strategies

### 🌊 Streaming & Real-time

#### `simple_streaming.go`
**Minimal streaming example**
- Basic SSE connection setup
- Simple event handling
- Content streaming
- Error handling
- Perfect for learning streaming basics

#### `streaming_example.go`
**Advanced SSE event handling**
- Comprehensive event type handling:
  - Thinking events
  - Content events
  - Tool events (call, start, complete, error)
  - Permission events (request, granted, denied)
  - Session events (created, messages)
  - Error and completion events
- Concurrent streaming and messaging
- Event parsing and display
- Timeout management

**Key Demonstrations:**
- SSE event stream parsing
- Goroutine-based concurrent processing
- Event type differentiation
- Real-time UI updates

### 📁 File Operations

#### `files_example.go`
**Complete file management**
- File upload (text, image, binary)
- File listing with metadata
- File download with content preview
- Thumbnail generation (box, width, height constraints)
- File deletion
- Session-based file isolation verification

**Key Demonstrations:**
- Multi-format file handling
- Thumbnail generation strategies
- File isolation between sessions
- Memory-efficient file operations

### 🔧 Tools & System

#### `tools_example.go`
**Tool discovery and status**
- Available LLM tools listing
- Tools status by category (core, search, code, files, external)
- Tool credentials status checking
- Provider analysis
- Authentication requirements analysis
- Capability assessment

**Key Demonstrations:**
- Tool categorization
- Provider-based filtering
- Authentication verification
- Capability discovery

#### `system_example.go`
**System introspection and monitoring**
- System health checks
- Command discovery and listing
- Detailed command inspection
- MCP (Model Context Protocol) server listing
- Command categorization
- Integration verification

**Key Demonstrations:**
- Health monitoring patterns
- Command introspection
- MCP integration
- System capability discovery

### 🔐 Permissions

#### `permissions_example.go`
**Permission management**
- Permission granting workflow
- Permission denial workflow
- Asynchronous permission operations
- Integration with streaming
- Advanced parameter configuration

**Key Demonstrations:**
- Permission request handling
- Streaming integration
- User-driven permission decisions
- Timeout and custom headers

## Running Examples

### Basic Usage
```bash
# Run any example
go run <example_name>.go

# Example:
go run sessions_example.go
```

### With Environment Variables
```bash
# Set server URL
export MIX_SERVER_URL="http://localhost:8088"

# Set API keys for authentication
export ANTHROPIC_API_KEY="sk-ant-..."
export OPENAI_API_KEY="sk-..."
export GEMINI_API_KEY="..."
export OPENROUTER_API_KEY="..."
export BRAVE_API_KEY="..."

# Run example
go run authentication_example.go
```

### Building and Running
```bash
# Build an example
go build -o basic_client basic_client.go

# Run the compiled binary
./basic_client
```

## Example Categories

| Category | Examples | Description |
|----------|----------|-------------|
| **Getting Started** | `example_template.go`, `basic_client.go` | Quick start and basic patterns |
| **Core Features** | `sessions_example.go`, `messages_example.go` | Session and message operations |
| **Authentication** | `authentication_example.go`, `preferences_example.go` | Auth and configuration |
| **Streaming** | `simple_streaming.go`, `streaming_example.go` | Real-time SSE events |
| **Files** | `files_example.go` | File upload, download, thumbnails |
| **Tools** | `tools_example.go`, `system_example.go` | Tool discovery and system info |
| **Permissions** | `permissions_example.go` | Permission management |

## Learning Path

### Beginner
1. Start with `example_template.go` - understand basic structure
2. Run `basic_client.go` - see complete CRUD workflow
3. Try `simple_streaming.go` - learn streaming basics

### Intermediate
4. Explore `sessions_example.go` - advanced session management
5. Study `messages_example.go` - conversation handling
6. Review `authentication_example.go` - multi-provider setup

### Advanced
7. Master `streaming_example.go` - comprehensive event handling
8. Learn `files_example.go` - file operations and isolation
9. Understand `tools_example.go` and `system_example.go` - introspection
10. Implement `permissions_example.go` - permission workflows

## Common Patterns

### Session Management
```go
// Create session
createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
    Title: "My Session",
})
sessionID := createResp.SessionData.ID

// Always cleanup
defer client.Sessions.DeleteSession(ctx, sessionID)
```

### Streaming
```go
// Start stream in goroutine
go func() {
    streamResp, _ := client.Streaming.StreamEvents(ctx, sessionID, nil)
    // Handle events...
}()

// Send message
client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
    Text: "Hello!",
})
```

### Error Handling
```go
resp, err := client.Sessions.CreateSession(ctx, request)
if err != nil {
    log.Fatalf("Operation failed: %v", err)
}
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `MIX_SERVER_URL` | Mix API server URL | `http://localhost:8088` |
| `ANTHROPIC_API_KEY` | Anthropic API key | - |
| `OPENAI_API_KEY` | OpenAI API key | - |
| `GEMINI_API_KEY` | Google Gemini API key | - |
| `OPENROUTER_API_KEY` | OpenRouter API key | - |
| `BRAVE_API_KEY` | Brave Search API key | - |
| `PREFERRED_PROVIDER` | Default AI provider | `anthropic` |

## Best Practices

### 1. Always Use Context
```go
ctx := context.Background()
// Or with timeout:
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
```

### 2. Cleanup Resources
```go
defer func() {
    for _, sessionID := range createdSessions {
        client.Sessions.DeleteSession(ctx, sessionID)
    }
}()
```

### 3. Handle Streaming Properly
```go
// Use goroutines for concurrent streaming
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Stream handling...
}()
// Send message...
wg.Wait()
```

### 4. Error Handling
```go
if err != nil {
    log.Printf("Non-fatal error: %v", err)
    // Handle gracefully
}
```

## Troubleshooting

### Connection Errors
- Verify `MIX_SERVER_URL` is correct
- Ensure the Mix server is running
- Check network connectivity

### Authentication Errors
- Verify API keys are set correctly
- Check provider authentication status with `GetAuthStatus()`
- Ensure credentials are stored with `StoreAPIKey()`

### Streaming Issues
- Use proper timeout contexts
- Ensure goroutine synchronization
- Check for stream close/cleanup

### File Operation Errors
- Verify file paths exist
- Check file permissions
- Ensure files are within size limits

## Additional Resources

- **SDK Documentation:** See `../docs/` directory
- **API Reference:** See `../docs/sdks/` for detailed API docs
- **Model Documentation:** See `../docs/models/` for data structures

## Contributing

When adding new examples:
1. Follow existing code style
2. Include comprehensive comments
3. Demonstrate error handling
4. Add cleanup patterns
5. Update this README

## Support

For issues or questions:
- Check example code comments
- Review SDK documentation
- Open an issue on GitHub
- Consult API reference docs

---

**Happy Coding!** 🚀
