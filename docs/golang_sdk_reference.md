# Mix SDK Reference - Go

> Complete API reference for the Mix Go SDK, including all functions, types, and structures.

## Installation

```bash
go get github.com/recreate-run/mix-go-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "github.com/recreate-run/mix-go-sdk"
    "github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
    // Initialize the SDK client
    client := mix.New("http://localhost:8088")

    ctx := context.Background()

    // Create a session
    session, _ := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Title:       "My First Session",
        BrowserMode: operations.BrowserModeElectronEmbeddedBrowser,
    })

    fmt.Printf("Created session: %s\n", session.SessionData.ID)
}
```

## Choosing Between API Patterns

The Mix Go SDK provides different patterns for interacting with the Mix application:

### Quick Comparison

| Feature                  | REST API (Polling)           | Streaming API (SSE)                |
| :----------------------- | :--------------------------- | :--------------------------------- |
| **Response Model**       | Request-response             | Real-time event stream             |
| **Connection**           | One-off HTTP requests        | Persistent SSE connection          |
| **Message Processing**   | Poll for results             | Push-based events                  |
| **Latency**              | Higher (polling interval)    | Lower (immediate events)           |
| **Complexity**           | Simple                       | Moderate (goroutine management)    |
| **Use Case**             | Simple request/response      | Real-time updates, long operations |
| **Resource Usage**       | Lower (no persistent conn)   | Higher (persistent connection)     |
| **Best For**             | CRUD operations              | Interactive AI conversations       |

### When to Use REST API (Request-Response)

**Best for:**

* Simple session management (create, list, delete)
* Retrieving conversation history
* File upload/download operations
* Configuration and preferences management
* One-off queries where you don't need real-time updates
* Stateless operations

**Example:**

```go
// Create session, send message, poll for results later
client.Sessions.CreateSession(ctx, request)
client.Messages.SendMessage(ctx, sessionID, message)
// Poll messages later
time.Sleep(2 * time.Second)
messages, _ := client.Messages.GetSessionMessages(ctx, sessionID)
```

### When to Use Streaming API (Server-Sent Events)

**Best for:**

* **Real-time AI interactions** - See responses as they're generated
* **Long-running operations** - Monitor progress of complex tasks
* **Interactive applications** - Chat interfaces, live dashboards
* **Event-driven workflows** - React to specific events (thinking, tool calls, errors)
* **Low-latency requirements** - Immediate notification of state changes

**Example:**

```go
// Start stream, send message, process events in real-time
go func() {
    stream, _ := client.Streaming.StreamEvents(ctx, sessionID, nil)
    defer stream.SSEEventStream.Close()
    for stream.SSEEventStream.Next() {
        processEvent(stream.SSEEventStream.Value())
    }
}()
client.Messages.SendMessage(ctx, sessionID, message)
```

### Hybrid Approach

Most production applications use both:

* **REST API** for session lifecycle and data retrieval
* **Streaming API** for real-time message processing

```go
client := mix.New("http://localhost:8088")

// Create session (REST)
session, _ := client.Sessions.CreateSession(ctx, request)

// Start streaming (SSE)
go streamHandler(client, session.SessionData.ID)

// Send messages (REST)
client.Messages.SendMessage(ctx, session.SessionData.ID, message)

// Export results (REST)
export, _ := client.Sessions.ExportSession(ctx, session.SessionData.ID)
```

## Core API

### `New()`

Creates a new Mix SDK client instance.

```go
func New(serverURL string, opts ...SDKOption) *Mix
```

#### Parameters

| Parameter   | Type        | Description                              |
| :---------- | :---------- | :--------------------------------------- |
| `serverURL` | `string`    | Required server URL                      |
| `opts`      | `SDKOption` | Optional configuration functions         |

#### Returns

Returns a pointer to a `Mix` client instance.

#### Example

```go
client := mix.New(
    "http://localhost:8088",
    mix.WithTimeout(30 * time.Second),
)
```

## Configuration Options

### `WithServerURL()`

Override the server URL after initialization.

```go
func WithServerURL(serverURL string) SDKOption
```

#### Parameters

| Parameter   | Type     | Description                  |
| :---------- | :------- | :--------------------------- |
| `serverURL` | `string` | Override the initial URL     |

#### Example

```go
client := mix.New(
    "http://localhost:8088",
    mix.WithServerURL("https://production.example.com"),
)
```

### `WithClient()`

Provide a custom HTTP client.

```go
func WithClient(client HTTPClient) SDKOption
```

#### Parameters

| Parameter | Type         | Description              |
| :-------- | :----------- | :----------------------- |
| `client`  | `HTTPClient` | Custom HTTP client       |

#### Example

```go
customClient := &http.Client{
    Timeout: 60 * time.Second,
}

client := mix.New(
    "http://localhost:8088",
    mix.WithClient(customClient),
)
```

### `WithRetryConfig()`

Configure retry behavior for failed requests.

```go
func WithRetryConfig(retryConfig retry.Config) SDKOption
```

#### Parameters

| Parameter     | Type           | Description           |
| :------------ | :------------- | :-------------------- |
| `retryConfig` | `retry.Config` | Retry configuration   |

#### Example

```go
client := mix.New(
    "http://localhost:8088",
    mix.WithRetryConfig(retry.Config{
        Strategy: "backoff",
        Backoff: &retry.BackoffStrategy{
            InitialInterval: 500,
            MaxInterval:     60000,
            Exponent:        1.5,
            MaxElapsedTime:  600000,
        },
        RetryConnectionErrors: true,
    }),
)
```

### `WithTimeout()`

Set a timeout for all requests.

```go
func WithTimeout(timeout time.Duration) SDKOption
```

#### Parameters

| Parameter | Type            | Description              |
| :-------- | :-------------- | :----------------------- |
| `timeout` | `time.Duration` | Request timeout duration |

#### Example

```go
client := mix.New(
    "http://localhost:8088",
    mix.WithTimeout(30 * time.Second),
)
```

## Client Structure

The SDK is organized into resource-based modules accessed through the main `Mix` client:

| Resource         | Description                              |
| :--------------- | :--------------------------------------- |
| `Sessions`       | Session lifecycle management             |
| `Messages`       | Message sending and history retrieval    |
| `Streaming`      | Server-Sent Events for real-time updates |
| `Files`          | File upload, download, and management    |
| `Authentication` | API key and OAuth management             |
| `Preferences`    | Model and provider configuration         |
| `Permissions`    | Permission granting and denial           |
| `Tools`          | Tool discovery and status                |
| `System`         | Health checks and system information     |
| `Notifications`  | Notification response handling           |
| `Health`         | OAuth and service health checks          |
| `Internal`       | Internal operations and token refresh    |

```go
client := mix.New("http://localhost:8088")

// Access resources through the client
client.Sessions.CreateSession(...)
client.Messages.SendMessage(...)
client.Streaming.StreamEvents(...)
client.Files.UploadSessionFile(...)
client.Notifications.RespondToNotification(...)
```

## Types

### Configuration Types

#### `SDKOption`

Function type for configuring SDK client initialization.

```go
type SDKOption func(*Mix)
```

Available options: `WithServerURL()`, `WithClient()`, `WithRetryConfig()`, `WithTimeout()`

#### `retry.Config`

Configuration for automatic retry behavior.

```go
type Config struct {
    Strategy              string
    Backoff               *BackoffStrategy
    RetryConnectionErrors bool
}

type BackoffStrategy struct {
    InitialInterval int64   // milliseconds
    MaxInterval     int64   // milliseconds
    Exponent        float64 // multiplier
    MaxElapsedTime  int64   // milliseconds
}
```

### Data Types

#### `SessionData`

Represents session metadata.

```go
type SessionData struct {
    ID                    string
    Title                 string
    SessionType           string          // "main" or "subagent"
    CreatedAt             time.Time
    AssistantMessageCount int64
    BrowserMode           BrowserMode     // Browser automation mode
    Callbacks             []Callback
    CdpURL                *string         // CDP WebSocket URL (remote-cdp-websocket only)
    CompletionTokens      int64
    Cost                  float64
    FirstUserMessage      *string
    ParentSessionID       *string         // For subagent sessions
    ParentToolCallID      *string         // For subagent sessions
    PromptTokens          int64
    SubagentType          *SubagentType   // e.g., "general-purpose"
    ToolCallCount         int64
    UserMessageCount      int64
}
```

#### `BackendMessage`

Represents a message in the conversation.

```go
type BackendMessage struct {
    ID                string
    SessionID         string
    Role              string
    UserInput         string
    AssistantResponse *string
    Reasoning         *string
    ReasoningDuration *int64
    ToolCalls         []ToolCallData
    CallbackResults   []CallbackResultData
}
```

#### `ExportSession`

Complete session export data.

```go
type ExportSession struct {
    ID          string
    Title       string
    SessionType string
    Messages    []ExportMessage
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

#### `ExportMessage`

Message data in session export.

```go
type ExportMessage struct {
    ID                string
    Role              string
    Content           string
    Reasoning         *string
    ReasoningDuration *int64
    ToolCalls         []ExportToolCall
    Timestamp         time.Time
    CreatedAt         time.Time
    UpdatedAt         time.Time
    FinishReason      *string
    Model             *string
}
```

#### `ToolCallData`

Tool call information.

```go
type ToolCallData struct {
    ID       string
    Name     string
    Type     string
    Input    string    // JSON string
    Result   *string   // Execution result
    Finished bool
    IsError  *bool
}
```

#### `FileInfo`

File metadata.

```go
type FileInfo struct {
    IsDir    bool   // Whether this is a directory
    Modified int64  // Last modified timestamp (Unix time)
    Name     string // File name
    Size     int64  // File size in bytes
    URL      string // Static URL to access the file
}
```

#### `BrowserMode`

Browser automation mode enum.

```go
type BrowserMode string

const (
    BrowserModeElectronEmbeddedBrowser  // Electron with embedded Chromium
    BrowserModeLocalBrowserService      // Local GoRod-based browser
    BrowserModeRemoteCdpWebsocket       // Remote CDP via WebSocket
)
```

#### `Callback`

Session callback configuration for automated actions.

```go
type Callback struct {
    Type                 string   // "bash_script", "sub_agent", "send_message"
    ToolName             *string
    BashCommand          *string
    SubAgentPrompt       *string
    MessageContent       *string
    BashTimeout          *int64
    ExcludeFromContext   *bool
    IncludeFullHistory   *bool
    SubAgentType         *string
    Name                 *string
}
```

#### `CallbackResultData`

Result data from callback execution.

```go
type CallbackResultData struct {
    CallbackName       string
    CallbackType       string
    Success            bool
    Error              *string
    ToolCallID         *string
    ToolName           *string
    ExitCode           *int
    Stdout             *string
    Stderr             *string
    SubagentID         *string
    SubagentResult     *string
    ExcludeFromContext *bool
}
```

#### `ExportToolCall`

Extended tool call data used in session exports.

```go
type ExportToolCall struct {
    ID        string
    Name      string
    Input     string      // JSON string
    InputJSON *InputJSON  // Optional parsed input
    Result    *string
    Finished  bool
    Metadata  *string
}
```

### Request Option Types

#### `operations.Option`

Interface for per-request configuration options.

```go
type Option interface {
    apply(*requestOptions)
}
```

Available options:

* `operations.WithRetries(retry.Config)` - Override retry configuration
* `operations.WithTimeout(time.Duration)` - Set custom timeout
* `operations.WithServerURL(string)` - Override server URL
* `operations.WithHeaders(map[string]string)` - Add custom headers

## Message Types

### `BackendMessage`

The primary message type used for conversation history and message retrieval.

```go
type BackendMessage struct {
    ID                string                 // Unique message identifier
    SessionID         string                 // Session identifier
    Role              string                 // "user" or "assistant"
    UserInput         string                 // User's message text
    AssistantResponse *string                // AI response text
    Reasoning         *string                // AI reasoning/thinking content
    ReasoningDuration *int64                 // Time spent reasoning (milliseconds)
    ToolCalls         []ToolCallData         // Tools invoked during processing
    CallbackResults   []CallbackResultData   // Callback execution results
}
```

**Usage Example:**

```go
messages, _ := client.Messages.GetSessionMessages(ctx, sessionID)
for _, msg := range messages.BackendMessages {
    fmt.Printf("[%s] %s\n", msg.Role, msg.UserInput)
    if msg.AssistantResponse != nil {
        fmt.Printf("Response: %s\n", *msg.AssistantResponse)
    }
    if len(msg.ToolCalls) > 0 {
        fmt.Printf("Tools used: %d\n", len(msg.ToolCalls))
    }
}
```

### `ExportMessage`

Extended message format used in session exports with complete metadata.

```go
type ExportMessage struct {
    ID                string           // Unique message identifier
    Role              string           // "user" or "assistant"
    Content           string           // Message content
    Reasoning         *string          // AI reasoning/thinking content
    ReasoningDuration *int64           // Time spent reasoning (milliseconds)
    ToolCalls         []ExportToolCall // Complete tool call data with results
    Timestamp         time.Time        // Message timestamp
    CreatedAt         time.Time        // Creation time
    UpdatedAt         time.Time        // Last update time
    FinishReason      *string          // Completion finish reason
    Model             *string          // Model used for message
}
```

**Usage Example:**

```go
export, _ := client.Sessions.ExportSession(ctx, sessionID)
for _, msg := range export.ExportSession.Messages {
    fmt.Printf("[%s at %s]\n", msg.Role, msg.Timestamp)
    fmt.Printf("Content: %s\n", msg.Content)
    for _, tool := range msg.ToolCalls {
        fmt.Printf("  Tool: %s -> %s\n", tool.Name, tool.Result)
    }
}
```

## Content Block Types

**PLACEHOLDER**: The Mix SDK primarily uses string-based message content. Currently, all message content is represented as strings in `BackendMessage.UserInput` and `BackendMessage.AssistantResponse`.

## Error Types

### `ErrorResponse`

Standard API error response returned by the Mix server.

```go
type ErrorResponse struct {
    HTTPMeta HTTPMetadata
    Message  string
    Code     *int
}
```

**Fields:**

* `HTTPMeta` - HTTP response metadata (status code, headers)
* `Message` - Human-readable error message
* `Code` - Optional application-specific error code

### `APIError`

Generic API error for HTTP-level failures.

```go
type APIError struct {
    Message    string
    StatusCode int
    Body       string
    Response   *http.Response
}
```

**Fields:**

* `Message` - Error description
* `StatusCode` - HTTP status code (4xx, 5xx)
* `Body` - Raw response body
* `Response` - Complete HTTP response object

### Error Handling Example

```go
import "github.com/recreate-run/mix-go-sdk/models/apierrors"

resp, err := client.Sessions.CreateSession(ctx, request)
if err != nil {
    // Check for specific error types
    var apiErr *apierrors.APIError
    if errors.As(err, &apiErr) {
        log.Printf("API error %d: %s", apiErr.StatusCode, apiErr.Message)
        log.Printf("Response body: %s", apiErr.Body)
    }

    var errResp *apierrors.ErrorResponse
    if errors.As(err, &errResp) {
        log.Printf("Server error: %s (code: %v)", errResp.Message, errResp.Code)
    }

    return err
}
```

### Common Error Status Codes

| Status Code | Meaning                | Typical Cause                           |
| :---------- | :--------------------- | :-------------------------------------- |
| 400         | Bad Request            | Invalid request parameters              |
| 401         | Unauthorized           | Missing or invalid authentication       |
| 403         | Forbidden              | Insufficient permissions                |
| 404         | Not Found              | Session or resource doesn't exist       |
| 409         | Conflict               | Resource state conflict                 |
| 422         | Unprocessable Entity   | Validation error                        |
| 429         | Too Many Requests      | Rate limit exceeded                     |
| 500         | Internal Server Error  | Server-side error                       |
| 502         | Bad Gateway            | Upstream service failure                |
| 503         | Service Unavailable    | Server overloaded or maintenance        |

## Hook Types

### Hook System Overview

The SDK provides a hooks system for intercepting and modifying HTTP requests and responses at three points in the request lifecycle.

### Hook Execution Points

Hooks execute at three stages:

1. **`BeforeRequest`** - Called before sending the HTTP request
2. **`AfterSuccess`** - Called after receiving a successful response (2xx status)
3. **`AfterError`** - Called after receiving an error response (non-2xx status)

### Execution Order

For every API call:

```
BeforeRequest → HTTP Request → AfterSuccess/AfterError
```

### `HookContext`

Each hook receives a `hooks.HookContext` containing request metadata:

```go
type HookContext struct {
    SDK              *Mix                     // SDK instance
    SDKConfiguration config.SDKConfiguration  // SDK configuration
    BaseURL          string                   // Base server URL
    Context          context.Context          // Request context
    OperationID      string                   // Operation identifier (e.g., "createSession")
    OAuth2Scopes     []string                 // OAuth scopes if applicable
    SecuritySource   interface{}              // Security/auth information
}
```

### Hook Function Signatures

```go
// BeforeRequest hook signature
type BeforeRequestHook func(
    ctx hooks.HookContext,
    req *http.Request,
) (*http.Request, error)

// AfterSuccess hook signature
type AfterSuccessHook func(
    ctx hooks.HookContext,
    resp *http.Response,
) (*http.Response, error)

// AfterError hook signature
type AfterErrorHook func(
    ctx hooks.HookContext,
    resp *http.Response,
    err error,
) (*http.Response, error)
```

### Hook Registration

**PLACEHOLDER**: Hook registration API is currently internal to the SDK. Custom hook implementation requires access to internal SDK structures and is intended for advanced use cases only.

Future public API may look like:

```go
// Proposed public hook registration API
client := mix.New(
    "http://localhost:8088",
    mix.WithHooks(hooks.Hooks{
        BeforeRequest: []hooks.BeforeRequestHook{
            logRequestHook,
            authInjectionHook,
        },
        AfterSuccess: []hooks.AfterSuccessHook{
            logResponseHook,
        },
        AfterError: []hooks.AfterErrorHook{
            errorLoggingHook,
            retryDecisionHook,
        },
    }),
)
```

### Example Hook Implementations

```go
// Example: Request logging hook
func logRequestHook(ctx hooks.HookContext, req *http.Request) (*http.Request, error) {
    log.Printf("[%s] %s %s", ctx.OperationID, req.Method, req.URL.Path)
    return req, nil
}

// Example: Response logging hook
func logResponseHook(ctx hooks.HookContext, resp *http.Response) (*http.Response, error) {
    log.Printf("[%s] Response: %d", ctx.OperationID, resp.StatusCode)
    return resp, nil
}

// Example: Custom auth injection hook
func authInjectionHook(ctx hooks.HookContext, req *http.Request) (*http.Request, error) {
    token := os.Getenv("CUSTOM_AUTH_TOKEN")
    if token != "" {
        req.Header.Set("Authorization", "Bearer "+token)
    }
    return req, nil
}

// Example: Error logging with context
func errorLoggingHook(ctx hooks.HookContext, resp *http.Response, err error) (*http.Response, error) {
    log.Printf("[%s] Error: %v (status: %d)",
        ctx.OperationID,
        err,
        resp.StatusCode,
    )
    return resp, err
}
```

### Use Cases for Hooks

1. **Request/Response Logging** - Audit all API calls
2. **Custom Authentication** - Inject auth headers or tokens
3. **Request Modification** - Add custom headers, modify payloads
4. **Response Transformation** - Parse or transform responses
5. **Error Handling** - Custom error logging or recovery
6. **Metrics Collection** - Track API usage and performance
7. **Debugging** - Inspect requests and responses

**Note:** Hooks are primarily used internally by the SDK for request/response processing. Public hook APIs for custom implementations may be added in future versions.

## SSE Event Types

The streaming API emits various Server-Sent Events during message processing. Each event represents a different stage in the AI's response generation.

| Event Type | Description | When Emitted | Common Use Case |
| :--------- | :---------- | :----------- | :-------------- |
| `SSEEventStreamTypeConnected` | Connection established | Stream initialization | Confirm stream is ready |
| `SSEEventStreamTypeUserMessageCreated` | User message created | After message submission | Confirm message received |
| `SSEEventStreamTypeThinking` | AI thinking/reasoning | Before response generation | Show "thinking" indicator |
| `SSEEventStreamTypeContent` | Content generation | During response streaming | Display streaming text response |
| `SSEEventStreamTypeToolUseStart` | Tool use started | AI declares tool to use | Show tool being called |
| `SSEEventStreamTypeToolUseParameterDelta` | Tool parameters streaming | Parameters being streamed | Display partial parameters |
| `SSEEventStreamTypeToolUseParameterStreamingComplete` | Tool parameters complete | All parameters received | Prepare for execution |
| `SSEEventStreamTypeToolExecutionStart` | Tool execution started | Tool begins running | Display execution progress |
| `SSEEventStreamTypeToolExecutionComplete` | Tool execution finished | Tool completes/fails | Show tool results or errors |
| `SSEEventStreamTypePermission` | Permission request | Action requires approval | Prompt user for permission |
| `SSEEventStreamTypeComplete` | Processing complete | Response finished | End of stream, cleanup |
| `SSEEventStreamTypeError` | Error occurred | Processing failure | Display error message |
| `SSEEventStreamTypeHeartbeat` | Keep-alive signal | Periodic during stream | Maintain connection |
| `SSEEventStreamTypeSessionCreated` | Session created | New session initialized | Notify of new session |
| `SSEEventStreamTypeSessionDeleted` | Session deleted | Session removed | Notify of deletion |
| `SSEEventStreamTypeSSENotificationEvent` | Notification request | User action required | Prompt user for response |

### Event Processing Pattern

Events should be processed in a loop using the `SSEEventStream` API:

```go
for stream.SSEEventStream.Next() {
    event := stream.SSEEventStream.Value()

    switch event.Type {
    case components.SSEEventStreamTypeContent:
        // Handle streaming content
    case components.SSEEventStreamTypeComplete:
        // End processing
        return
    case components.SSEEventStreamTypeError:
        // Handle errors
    }
}
```

## Tool Input/Output Types

Documentation of input/output schemas for all built-in Mix tools. These tools are invoked by the AI during message processing and appear in SSE tool event streams (`SSEToolUseStartEvent`, `SSEToolUseParameterDeltaEvent`, `SSEToolUseParameterStreamingCompleteEvent`).

### Bash

**Tool name:** `Bash`

**Input:**

```go
{
    "command": string,              // The command to execute
    "timeout": int | nil,           // Optional timeout in milliseconds (max 600000)
}
```

**Output:**

```go
{
    "output": string,               // Combined stdout and stderr output
    "exitCode": int,                // Exit code of the command
    "start_time": int64,            // Start timestamp in milliseconds
    "end_time": int64,              // End timestamp in milliseconds
}
```

**Constants:**

* Default timeout: 60,000ms (1 minute)
* Max timeout: 600,000ms (10 minutes)
* Max output length: 30,000 characters (truncates if exceeded)

### Edit

**Tool name:** `Edit`

**Input:**

```go
{
    "file_path": string,            // Absolute path to the file
    "old_string": string,           // Text to replace (empty for new file creation)
    "new_string": string,           // Replacement text (empty for deletion)
    "replace_all": bool | nil,      // Replace all occurrences (default: false)
}
```

**Output:**

```go
{
    "message": string,              // Success message
    "diff": string,                 // Diff showing changes
    "additions": int,               // Number of lines added
    "removals": int,                // Number of lines removed
}
```

**Special Behaviors:**

* If `old_string` is empty: Creates new file with `new_string` content
* If `new_string` is empty: Deletes `old_string` from file
* Requires file to be read first before editing

### Write

**Tool name:** `Write`

**Input:**

```go
{
    "file_path": string,            // Path to the file (absolute or relative)
    "content": string,              // Content to write
}
```

**Output:**

```go
{
    "message": string,              // Success message
    "diff": string,                 // Diff showing changes
    "additions": int,               // Number of lines added
    "removals": int,                // Number of lines removed
}
```

### ReadText

**Tool name:** `ReadText`

**Input:**

```go
{
    "file_path": string,            // Absolute path or HTTP/HTTPS URL
    "offset": int | nil,            // Line number to start from (0-based)
    "limit": int | nil,             // Number of lines to read (default: 2000)
}
```

**Output:**

```go
{
    "content": string,              // File content with line numbers (cat -n format)
    "file_path": string,            // Path/URL of the file
}
```

**Constants:**

* Default read limit: 2000 lines
* Max line length: 2000 characters (truncates with "...")

### Grep

**Tool name:** `Grep`

**Input:**

```go
{
    "pattern": string,              // Regex pattern to search for
    "path": string | nil,           // File/directory to search
    "glob": string | nil,           // Glob pattern to filter files (e.g., "*.js")
    "type": string | nil,           // File type filter (e.g., "js", "py", "rust")
    "output_mode": string | nil,    // "content", "files_with_matches", "count" (default: "files_with_matches")
    "-i": bool | nil,               // Case insensitive search
    "-n": bool | nil,               // Show line numbers (content mode only)
    "-A": int | nil,                // Lines after match (content mode only)
    "-B": int | nil,                // Lines before match (content mode only)
    "-C": int | nil,                // Lines around match (content mode only)
    "multiline": bool | nil,        // Enable multiline mode
    "head_limit": int | nil,        // Limit output lines/entries
}
```

**Output (content mode):**

```go
{
    "results": string,              // Search results with context
    "number_of_matches": int,       // Count of matches
    "truncated": bool,              // Whether results were truncated
}
```

**Output (files_with_matches mode):**

```go
{
    "files": []string,              // Files containing matches (sorted by modification time)
    "count": int,                   // Number of files with matches
}
```

**Output (count mode):**

```go
{
    "counts": map[string]int,       // Match counts per file
    "total_matches": int,           // Total number of matches
}
```

### Glob

**Tool name:** `Glob`

**Input:**

```go
{
    "pattern": string,              // Glob pattern (e.g., "**/*.js")
    "path": string | nil,           // Directory to search
}
```

**Output:**

```go
{
    "matches": []string,            // Array of matching file paths (sorted by modification time)
    "count": int,                   // Number of matches found
    "truncated": bool,              // Whether results were truncated
}
```

**Constants:**

* Result limit: 100 files

### ReadMedia

**Tool name:** `ReadMedia`

**Input:**

```go
{
    "file_path": string,            // Absolute path or URL
    "media_type": string,           // "image", "audio", "video", or "pdf"
    "prompt": string,               // Analysis prompt
    "pdf_pages": string | nil,      // PDF page selection (e.g., "1-3,7,10-12")
    "video_interval": string | nil, // Video time interval (e.g., "00:20:50-00:26:10")
}
```

**Output:**

```go
{
    "results": []struct{
        "file_path": string,
        "media_type": string,
        "analysis": string,
        "error": string | nil,
    },
    "summary": string,              // Overall summary
}
```

**Supported Formats:**

* Images: .jpg, .jpeg, .png, .gif, .webp, .bmp
* Audio: .mp3, .wav, .m4a, .aac, .ogg, .flac
* Video: .mp4, .avi, .mov, .wmv, .flv, .webm, .mkv
* PDF: .pdf

**Special Behaviors:**

* Requires Gemini API key configuration
* Auto-truncates videos to first 10 minutes if no interval specified
* Auto-truncates PDFs to first 10 pages if no range specified
* Supports YouTube URLs for video analysis

### Show

**Tool name:** `Show`

**Input:**

```go
{
    "outputs": []struct{
        "type": string,             // "image", "video", "audio", "pdf", "csv", "markdown", "json", or "status"
        "path": string | nil,       // HTTP/HTTPS URL (for image/video/audio/pdf/csv). Not used for markdown, json, or status.
        "data": string | nil,       // Inline content string (for markdown/json/status types). For json, this should be a JSON string. For status, this is the status message.
        "title": string,            // Display title for the content (required for all types)
        "startTime": int | nil,     // Start time in seconds for video/audio segments (min: 0)
        "duration": int | nil,      // Duration in seconds for video/audio segments (min: 1)
    },
}
```

**Output:**

```go
{
    "message": string,              // Success message with count and titles
}
```

**Content Types:**

* **Files** (image, video, audio, pdf, csv): Require HTTP/HTTPS URLs in `path` and `title`
* **Inline content** (markdown, json, status): Use `data` field for content string and `title`

**Special Behaviors:**

* All types require `title` field
* Status type uses `data` field for the status message
* Markdown type passes content directly in `data` field
* JSON type requires valid JSON string in `data` field and validates it
* All file-based types require valid HTTP/HTTPS URLs

### WebFetch

**Tool name:** `WebFetch`

**Input:**

```go
{
    "url": string,                  // URL to fetch (format: uri)
    "prompt": string,               // Analysis prompt for the content
}
```

**Output:**

```go
{
    "content": string,              // Formatted markdown content
    "url": string,                  // Original URL
    "fetched_url": string,          // Final URL after redirects
    "content_type": string,         // Content-Type header
    "start_time": int64,            // Start timestamp
    "end_time": int64,              // End timestamp
}
```

**Constants:**

* Max content size: 5MB
* Cache TTL: 15 minutes
* Cache size: 100 entries
* Timeout: 30 seconds

**Special Behaviors:**

* Auto-upgrades HTTP to HTTPS
* Only supports HTTPS URLs
* Converts HTML to markdown

### Search

**Tool name:** `Search`

**Input:**

```go
{
    "query": string,                // Search query (minLength: 2)
    "search_type": string | nil,    // "web", "images", or "videos" (default: "web")
    "allowed_domains": []string | nil, // Only include results from these domains
    "blocked_domains": []string | nil, // Exclude results from these domains
    "safesearch": string | nil,     // "strict", "moderate", or "off" (default: "strict")
    "spellcheck": bool | nil,       // Enable spell correction (default: true)
}
```

**Output (web search):**

```go
{
    "results": []struct{
        "title": string,
        "url": string,
        "description": string,
    },
    "total_results": int,
}
```

**Output (image search):**

```go
{
    "results": []struct{
        "title": string,
        "image_url": string,
        "thumbnail": string,
        "source": string,
        "confidence": float64,
    },
    "total_results": int,
}
```

**Output (video search):**

```go
{
    "results": []struct{
        "title": string,
        "video_url": string,
        "duration": string,
        "views": int,
        "upload_date": string,
        "thumbnail": string,
        "source": string,
    },
    "total_results": int,
}
```

**Constants:**

* Max search results: 3
* Timeout: 30 seconds

**Special Behaviors:**

* Requires Brave Search API key configuration

### TodoWrite

**Tool name:** `TodoWrite`

**Input:**

```go
{
    "todos": []struct{
        "id": string,               // Unique identifier
        "content": string,          // Task description (minLength: 1)
        "status": string,           // "pending", "in_progress", or "completed"
        "priority": string,         // "high", "medium", or "low"
    },
}
```

**Output:**

```go
{
    "message": string,              // Success message with todo count
}
```

**Storage:**

* Location: `{data_directory}/todos/todos.json`
* Format: JSON array of todo objects

### Task

**Tool name:** `Task`

**Input:**

```go
{
    "description": string,          // Short task description (3-5 words)
    "prompt": string,               // Detailed task instructions
    "subagent_type": string,        // Agent type (e.g., "general-purpose", "Explore", "Plan")
    "model": string | nil,          // Optional model override ("sonnet", "opus", "haiku")
    "max_turns": int | nil,         // Maximum agentic turns before stopping
    "run_in_background": bool | nil, // Run task in background
}
```

**Output:**

```go
{
    "result": string,               // Task execution result or agent response
    "agent_id": string,             // Agent ID for resuming
}
```

**Special Behaviors:**

* Spawns specialized sub-agents for complex multi-step tasks
* Different subagent types have specific capabilities and tools
* Can run autonomously in background with `run_in_background: true`

### ExitPlanMode

**Tool name:** `ExitPlanMode`

**Input:**

```go
{
    "plan": string,                 // The plan (supports markdown, should be concise)
}
```

**Output:**

```go
{
    "message": string,              // Formatted plan with approval prompt
}
```

**Constants:**

* Default timeout: 30,000ms (30 seconds)
* Max timeout: 120,000ms (2 minutes)
* Max output length: 30,000 characters

**Special Behaviors:**

* Rejects code containing: `subprocess`, `os`, `exec()`, `eval()`, `__import__`
* Runs with `--with numpy --with pandas`
* Executes in isolated sandboxed environment using `uv run --isolated`

## Example Usage

### Basic Session Workflow

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/recreate-run/mix-go-sdk"
    "github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
    // Initialize client
    client := mix.New(
        "http://localhost:8088",
        mix.WithTimeout(30 * time.Second),
    )

    ctx := context.Background()

    // Create session
    createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Title:       "Example Session",
        BrowserMode: operations.BrowserModeElectronEmbeddedBrowser,
    })
    if err != nil {
        log.Fatalf("Failed to create session: %v", err)
    }
    sessionID := createResp.SessionData.ID

    // Cleanup on exit
    defer func() {
        client.Sessions.DeleteSession(ctx, sessionID)
    }()

    // Send message
    _, err = client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
        Text: "Hello, Claude!",
    })
    if err != nil {
        log.Fatalf("Failed to send message: %v", err)
    }

    // Wait for processing
    time.Sleep(2 * time.Second)

    // Get messages
    messagesResp, err := client.Messages.GetSessionMessages(ctx, sessionID)
    if err != nil {
        log.Fatalf("Failed to get messages: %v", err)
    }

    for _, msg := range messagesResp.BackendMessages {
        fmt.Printf("[%s]: %s\n", msg.Role, msg.UserInput)
    }
}
```

### Streaming with Event Processing

```go
package main

import (
    "context"
    "fmt"
    "log"
    "sync"
    "time"

    "github.com/recreate-run/mix-go-sdk"
    "github.com/recreate-run/mix-go-sdk/models/components"
    "github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
    client := mix.New("http://localhost:8088")

    ctx := context.Background()

    // Create session
    createResp, _ := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Title:       "Streaming Example",
        BrowserMode: operations.BrowserModeElectronEmbeddedBrowser,
    })
    sessionID := createResp.SessionData.ID
    defer client.Sessions.DeleteSession(ctx, sessionID)

    var wg sync.WaitGroup
    wg.Add(1)

    // Start streaming
    go func() {
        defer wg.Done()

        streamResp, err := client.Streaming.StreamEvents(ctx, sessionID, nil)
        if err != nil {
            log.Printf("Stream failed: %v", err)
            return
        }
        defer streamResp.SSEEventStream.Close()

        for streamResp.SSEEventStream.Next() {
            event := streamResp.SSEEventStream.Value()
            if event == nil {
                continue
            }

            switch event.Type {
            case components.SSEEventStreamTypeThinking:
                fmt.Println("💭 Thinking...")
            case components.SSEEventStreamTypeContent:
                if event.SSEContentEvent != nil {
                    fmt.Printf("💬 %s\n", event.SSEContentEvent.Data.Content)
                }
            case components.SSEEventStreamTypeComplete:
                fmt.Println("✅ Complete")
                return
            case components.SSEEventStreamTypeError:
                if event.SSEErrorEvent != nil {
                    fmt.Printf("❌ Error: %s\n", event.SSEErrorEvent.Data.Error)
                }
            }
        }

        if err := streamResp.SSEEventStream.Err(); err != nil {
            log.Printf("Stream error: %v", err)
        }
    }()

    // Brief delay to ensure stream is connected
    time.Sleep(500 * time.Millisecond)

    // Send message
    client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
        Text: "Explain recursion with an example",
    })

    // Wait for streaming to complete
    wg.Wait()
}
```

### File Upload and Download

```go
package main

import (
    "context"
    "fmt"
    "io"
    "io/ioutil"
    "log"

    "github.com/recreate-run/mix-go-sdk"
    "github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
    client := mix.New("http://localhost:8088")

    ctx := context.Background()

    // Create session
    createResp, _ := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Title:       "File Example",
        BrowserMode: operations.BrowserModeElectronEmbeddedBrowser,
    })
    sessionID := createResp.SessionData.ID
    defer client.Sessions.DeleteSession(ctx, sessionID)

    // Read file
    fileData, err := ioutil.ReadFile("example.txt")
    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }

    // Upload file
    uploadResp, err := client.Files.UploadSessionFile(ctx, sessionID, operations.UploadSessionFileRequestBody{
        File: operations.File{
            FileName: "example.txt",
            Content:  fileData,
        },
    })
    if err != nil {
        log.Fatalf("Failed to upload file: %v", err)
    }
    fmt.Printf("Uploaded: %s (URL: %s)\n", uploadResp.FileInfo.Name, uploadResp.FileInfo.URL)

    // List files
    listResp, err := client.Files.ListSessionFiles(ctx, sessionID)
    if err != nil {
        log.Fatalf("Failed to list files: %v", err)
    }

    for _, file := range listResp.FileInfos {
        fmt.Printf("- %s (%d bytes, %s)\n", file.Name, file.Size, file.URL)
    }

    // Download file
    downloadResp, err := client.Files.GetSessionFile(ctx, sessionID, "example.txt", nil, nil)
    if err != nil {
        log.Fatalf("Failed to download file: %v", err)
    }
    defer downloadResp.ResponseStream.Close()

    data, err := io.ReadAll(downloadResp.ResponseStream)
    if err != nil {
        log.Fatalf("Failed to read stream: %v", err)
    }
    ioutil.WriteFile("downloaded.txt", data, 0644)
}
```

## See Also

* [Go SDK Examples](../examples/) - Comprehensive examples directory
* [API Models](../models/) - Data structures and types
* [Operations](../models/operations/) - Request and response types
* [Components](../models/components/) - Shared component types

---

**Version:** 0.2.1
**Last Updated:** 2025

**Note:** This SDK is a REST API client for the Mix application. Event schemas and hook implementations are subject to change as the API evolves.
