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
)

func main() {
    // Initialize the SDK client
    client := mix.New(
        mix.WithServerURL("http://localhost:8088"),
    )

    ctx := context.Background()

    // Create a session
    session, _ := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Title: "My First Session",
    })

    fmt.Printf("Created session: %s\n", session.SessionData.ID)
}
```

## Core Concepts

The Mix Go SDK is a REST API client for interacting with the Mix application. It provides:

- **Session Management**: Create, list, update, and delete sessions
- **Messaging**: Send messages and retrieve conversation history
- **Streaming**: Real-time Server-Sent Events (SSE) for live updates
- **File Operations**: Upload, download, and manage files
- **Authentication**: Manage API keys for multiple providers
- **System Operations**: Health checks, tool discovery, and configuration

## SDK Structure

The SDK is organized into resource-based modules accessed through the main `Mix` client:

| Resource       | Description                              |
| :------------- | :--------------------------------------- |
| `Sessions`     | Session lifecycle management             |
| `Messages`     | Message sending and history retrieval    |
| `Streaming`    | Server-Sent Events for real-time updates |
| `Files`        | File upload, download, and management    |
| `Authentication` | API key and OAuth management           |
| `Preferences`  | Model and provider configuration         |
| `Permissions`  | Permission granting and denial           |
| `Tools`        | Tool discovery and status                |
| `System`       | Health checks and system information     |
| `Health`       | Health monitoring endpoints              |

## Initialization

### `New()`

Creates a new Mix SDK client instance.

```go
func New(opts ...SDKOption) *Mix
```

#### Parameters

Accepts variadic `SDKOption` functions for configuration.

#### Returns

Returns a pointer to a `Mix` client instance.

#### Example

```go
client := mix.New(
    mix.WithServerURL("http://localhost:8088"),
    mix.WithTimeout(30 * time.Second),
)
```

## Configuration Options

### `WithServerURL()`

Override the default server URL.

```go
func WithServerURL(serverURL string) SDKOption
```

#### Parameters

| Parameter   | Type     | Description            |
| :---------- | :------- | :--------------------- |
| `serverURL` | `string` | The server URL to use  |

#### Example

```go
client := mix.New(
    mix.WithServerURL("http://localhost:8088"),
)
```

### `WithServerIndex()`

Override the default server by index.

```go
func WithServerIndex(serverIndex int) SDKOption
```

#### Parameters

| Parameter     | Type  | Description                 |
| :------------ | :---- | :-------------------------- |
| `serverIndex` | `int` | Index in the ServerList     |

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
    mix.WithTimeout(30 * time.Second),
)
```

### Default Retry Behavior

All SDK methods automatically retry failed requests with the following default configuration:

```go
retry.Config{
    Strategy: "backoff",
    Backoff: &retry.BackoffStrategy{
        InitialInterval: 500,        // 500ms initial delay
        MaxInterval:     60000,      // 60 seconds maximum delay
        Exponent:        1.5,        // Exponential backoff multiplier
        MaxElapsedTime:  600000,     // 10 minutes total retry time
    },
    RetryConnectionErrors: true,
}
```

**Retry Behavior:**
- **Retryable Status Codes:** `5XX`, `408` (Request Timeout), `429` (Too Many Requests)
- **Connection Errors:** Automatically retried
- **Backoff Strategy:** Exponential backoff between retry attempts (starts at 500ms, increases by 1.5x each time, max 60 seconds)
- **Maximum Retry Time:** 10 minutes total across all retry attempts
- **Override:** Use `operations.WithRetries()` to customize retry behavior per request

## Hooks and Middleware

The SDK provides a hooks system for intercepting and modifying HTTP requests and responses.

### Hook Types

Hooks execute at three points in the request lifecycle:

1. **`BeforeRequest`** - Called before sending the HTTP request
2. **`AfterSuccess`** - Called after receiving a successful response (2xx status)
3. **`AfterError`** - Called after receiving an error response (non-2xx status)

### Hook Context

Each hook receives a `hooks.HookContext` containing:

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

### Execution Order

For every API call:
```
BeforeRequest → HTTP Request → AfterSuccess/AfterError
```

**Note:** Hooks are primarily used internally by the SDK for request/response processing. Custom hook implementation requires access to internal SDK structures and is intended for advanced use cases only.

## Helper Functions

The SDK provides helper functions for creating pointers to primitive types:

```go
// String returns a pointer to a string
func String(s string) *string

// Bool returns a pointer to a bool
func Bool(b bool) *bool

// Int returns a pointer to an int
func Int(i int) *int

// Int64 returns a pointer to an int64
func Int64(i int64) *int64

// Float32 returns a pointer to a float32
func Float32(f float32) *float32

// Float64 returns a pointer to a float64
func Float64(f float64) *float64

// Pointer returns a pointer to any type
func Pointer[T any](v T) *T
```

### Example

```go
// Use helpers when creating requests
request := operations.CreateSessionRequest{
    Title:       "My Session",
    SessionType: operations.SessionTypeMain.ToPointer(),
}

// Or use the generic Pointer helper
limit := mix.Int64(10)
```

## Sessions API

The `Sessions` resource provides methods for session lifecycle management.

### `ListSessions()`

Retrieve a list of all available sessions with their metadata.

```go
func (s *Sessions) ListSessions(
    ctx context.Context,
    includeSubagents *bool,
    opts ...operations.Option,
) (*operations.ListSessionsResponse, error)
```

#### Parameters

| Parameter          | Type               | Description                          |
| :----------------- | :----------------- | :----------------------------------- |
| `ctx`              | `context.Context`  | Context for the request              |
| `includeSubagents` | `*bool`            | Include subagent sessions (optional) |
| `opts`             | `...operations.Option` | Additional request options       |

#### Returns

Returns `*operations.ListSessionsResponse` containing an array of `SessionData` or an error.

#### Example

```go
ctx := context.Background()
resp, err := client.Sessions.ListSessions(ctx, nil)
if err != nil {
    log.Fatalf("Failed to list sessions: %v", err)
}

for _, session := range resp.SessionData {
    fmt.Printf("Session: %s - %s\n", session.ID, session.Title)
}
```

### `CreateSession()`

Create a new session with a required title and optional custom configuration.

```go
func (s *Sessions) CreateSession(
    ctx context.Context,
    request operations.CreateSessionRequest,
    opts ...operations.Option,
) (*operations.CreateSessionResponse, error)
```

#### Parameters

| Parameter | Type                               | Description            |
| :-------- | :--------------------------------- | :--------------------- |
| `ctx`     | `context.Context`                  | Context for the request |
| `request` | `operations.CreateSessionRequest`  | Session creation data   |
| `opts`    | `...operations.Option`             | Additional options      |

#### Request Fields

| Field            | Type                          | Required | Description                |
| :--------------- | :---------------------------- | :------- | :------------------------- |
| `Title`          | `string`                      | Yes      | Session title              |
| `SessionType`    | `*operations.SessionType`     | No       | Session type (main/forked) |
| `SystemPrompt`   | `*string`                     | No       | Custom system prompt       |

#### Returns

Returns `*operations.CreateSessionResponse` with the created session data or an error.

#### Example

```go
ctx := context.Background()
resp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
    Title:       "My New Session",
    SessionType: operations.SessionTypeMain.ToPointer(),
})
if err != nil {
    log.Fatalf("Failed to create session: %v", err)
}

sessionID := resp.SessionData.ID
fmt.Printf("Created session: %s\n", sessionID)
```

### `GetSession()`

Retrieve detailed information about a specific session.

```go
func (s *Sessions) GetSession(
    ctx context.Context,
    id string,
    opts ...operations.Option,
) (*operations.GetSessionResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `id`      | `string`               | Session ID              |
| `opts`    | `...operations.Option` | Additional options      |

#### Returns

Returns `*operations.GetSessionResponse` with session details or an error.

#### Example

```go
ctx := context.Background()
resp, err := client.Sessions.GetSession(ctx, sessionID)
if err != nil {
    log.Fatalf("Failed to get session: %v", err)
}

fmt.Printf("Session: %s\n", resp.SessionData.Title)
fmt.Printf("Created: %s\n", resp.SessionData.CreatedAt)
```

### `DeleteSession()`

Permanently delete a session and all its data.

```go
func (s *Sessions) DeleteSession(
    ctx context.Context,
    id string,
    opts ...operations.Option,
) (*operations.DeleteSessionResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `id`      | `string`               | Session ID to delete    |
| `opts`    | `...operations.Option` | Additional options      |

#### Returns

Returns `*operations.DeleteSessionResponse` or an error.

#### Example

```go
ctx := context.Background()
resp, err := client.Sessions.DeleteSession(ctx, sessionID)
if err != nil {
    log.Fatalf("Failed to delete session: %v", err)
}

fmt.Println("Session deleted successfully")
```

### `ForkSession()`

Create a new session based on an existing session, copying messages up to a specified index.

```go
func (s *Sessions) ForkSession(
    ctx context.Context,
    id string,
    requestBody operations.ForkSessionRequestBody,
    opts ...operations.Option,
) (*operations.ForkSessionResponse, error)
```

#### Parameters

| Parameter     | Type                                  | Description             |
| :------------ | :------------------------------------ | :---------------------- |
| `ctx`         | `context.Context`                     | Context for the request |
| `id`          | `string`                              | Source session ID       |
| `requestBody` | `operations.ForkSessionRequestBody`   | Fork configuration      |
| `opts`        | `...operations.Option`                | Additional options      |

#### Request Body Fields

| Field          | Type      | Required | Description                        |
| :------------- | :-------- | :------- | :--------------------------------- |
| `MessageIndex` | `int64`   | Yes      | Fork after this message index      |
| `Title`        | `*string` | No       | Title for the forked session       |

#### Returns

Returns `*operations.ForkSessionResponse` with the new forked session or an error.

#### Example

```go
ctx := context.Background()
resp, err := client.Sessions.ForkSession(ctx, sessionID, operations.ForkSessionRequestBody{
    MessageIndex: 2,
    Title:        mix.String("Forked Session - Alternative Path"),
})
if err != nil {
    log.Fatalf("Failed to fork session: %v", err)
}

forkedSessionID := resp.SessionData.ID
fmt.Printf("Forked session: %s\n", forkedSessionID)
```

### `RewindSession()`

Delete messages after a specified message in the current session.

```go
func (s *Sessions) RewindSession(
    ctx context.Context,
    id string,
    requestBody operations.RewindSessionRequestBody,
    opts ...operations.Option,
) (*operations.RewindSessionResponse, error)
```

#### Parameters

| Parameter     | Type                                   | Description              |
| :------------ | :------------------------------------- | :----------------------- |
| `ctx`         | `context.Context`                      | Context for the request  |
| `id`          | `string`                               | Session ID               |
| `requestBody` | `operations.RewindSessionRequestBody`  | Rewind configuration     |
| `opts`        | `...operations.Option`                 | Additional options       |

#### Request Body Fields

| Field        | Type      | Required | Description                   |
| :----------- | :-------- | :------- | :---------------------------- |
| `MessageID`  | `string`  | Yes      | Delete messages after this ID |
| `CleanMedia` | `*bool`   | No       | Clean up media files          |

#### Example

```go
ctx := context.Background()
resp, err := client.Sessions.RewindSession(ctx, sessionID, operations.RewindSessionRequestBody{
    MessageID:  messageID,
    CleanMedia: mix.Bool(true),
})
if err != nil {
    log.Fatalf("Failed to rewind session: %v", err)
}

fmt.Println("Session rewound successfully")
```

### `ExportSession()`

Export complete session transcript with all messages, tool calls, reasoning, and metadata as JSON.

```go
func (s *Sessions) ExportSession(
    ctx context.Context,
    id string,
    opts ...operations.Option,
) (*operations.ExportSessionResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `id`      | `string`               | Session ID to export    |
| `opts`    | `...operations.Option` | Additional options      |

#### Returns

Returns `*operations.ExportSessionResponse` with complete session data or an error.

#### Example

```go
ctx := context.Background()
resp, err := client.Sessions.ExportSession(ctx, sessionID)
if err != nil {
    log.Fatalf("Failed to export session: %v", err)
}

fmt.Printf("Exported %d messages\n", len(resp.ExportSession.Messages))
fmt.Printf("Session: %s\n", resp.ExportSession.Title)
```

### `UpdateSessionCallbacks()`

Update the callback configurations for a session.

```go
func (s *Sessions) UpdateSessionCallbacks(
    ctx context.Context,
    id string,
    requestBody operations.UpdateSessionCallbacksRequestBody,
    opts ...operations.Option,
) (*operations.UpdateSessionCallbacksResponse, error)
```

#### Parameters

| Parameter     | Type                                            | Description              |
| :------------ | :---------------------------------------------- | :----------------------- |
| `ctx`         | `context.Context`                               | Context for the request  |
| `id`          | `string`                                        | Session ID               |
| `requestBody` | `operations.UpdateSessionCallbacksRequestBody`  | Callback configuration   |
| `opts`        | `...operations.Option`                          | Additional options       |

#### Example

```go
ctx := context.Background()
resp, err := client.Sessions.UpdateSessionCallbacks(ctx, sessionID,
    operations.UpdateSessionCallbacksRequestBody{
        Callbacks: []components.Callback{
            {
                Type: components.CallbackTypeBashScript,
            },
        },
    },
)
if err != nil {
    log.Fatalf("Failed to update callbacks: %v", err)
}
```

## Messages API

The `Messages` resource provides methods for message operations and conversation management.

### `SendMessage()`

Send a user message to a specific session for AI processing.

```go
func (s *Messages) SendMessage(
    ctx context.Context,
    id string,
    requestBody operations.SendMessageRequestBody,
    opts ...operations.Option,
) (*operations.SendMessageResponse, error)
```

#### Parameters

| Parameter     | Type                                 | Description              |
| :------------ | :----------------------------------- | :----------------------- |
| `ctx`         | `context.Context`                    | Context for the request  |
| `id`          | `string`                             | Session ID               |
| `requestBody` | `operations.SendMessageRequestBody`  | Message data             |
| `opts`        | `...operations.Option`               | Additional options       |

#### Request Body Fields

| Field   | Type     | Required | Description        |
| :------ | :------- | :------- | :----------------- |
| `Text`  | `string` | Yes      | Message text       |

#### Returns

Returns `*operations.SendMessageResponse` with 202 Accepted status. The message is processed asynchronously and results stream via SSE.

#### Example

```go
ctx := context.Background()
resp, err := client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
    Text: "What is the capital of France?",
})
if err != nil {
    log.Fatalf("Failed to send message: %v", err)
}

fmt.Println("Message sent (processing asynchronously)")
```

### `GetMessageHistory()`

Retrieve message history across all sessions with optional pagination.

```go
func (s *Messages) GetMessageHistory(
    ctx context.Context,
    limit *int64,
    offset *int64,
    opts ...operations.Option,
) (*operations.GetMessageHistoryResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `limit`   | `*int64`               | Maximum messages to return (optional) |
| `offset`  | `*int64`               | Skip this many messages (optional) |
| `opts`    | `...operations.Option` | Additional options      |

#### Returns

Returns `*operations.GetMessageHistoryResponse` containing an array of `BackendMessage` or an error.

#### Example

```go
ctx := context.Background()
limit := int64(10)
resp, err := client.Messages.GetMessageHistory(ctx, &limit, nil)
if err != nil {
    log.Fatalf("Failed to get message history: %v", err)
}

for _, msg := range resp.BackendMessages {
    fmt.Printf("%s: %s\n", msg.Role, msg.UserInput)
}
```

### `GetSessionMessages()`

Retrieve all messages from a specific session.

```go
func (s *Messages) GetSessionMessages(
    ctx context.Context,
    id string,
    opts ...operations.Option,
) (*operations.GetSessionMessagesResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `id`      | `string`               | Session ID              |
| `opts`    | `...operations.Option` | Additional options      |

#### Returns

Returns `*operations.GetSessionMessagesResponse` with messages or an error.

#### Example

```go
ctx := context.Background()
resp, err := client.Messages.GetSessionMessages(ctx, sessionID)
if err != nil {
    log.Fatalf("Failed to get session messages: %v", err)
}

for _, msg := range resp.BackendMessages {
    fmt.Printf("[%s] %s\n", msg.Role, msg.UserInput)
}
```

### `CancelSessionProcessing()`

Cancel any ongoing agent processing in the specified session.

```go
func (s *Messages) CancelSessionProcessing(
    ctx context.Context,
    id string,
    opts ...operations.Option,
) (*operations.CancelSessionProcessingResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `id`      | `string`               | Session ID              |
| `opts`    | `...operations.Option` | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.Messages.CancelSessionProcessing(ctx, sessionID)
if err != nil {
    log.Fatalf("Failed to cancel processing: %v", err)
}

fmt.Println("Processing cancelled")
```

## Streaming API

The `Streaming` resource provides Server-Sent Events (SSE) for real-time updates.

### `StreamEvents()`

Establishes a persistent SSE connection for receiving real-time updates during message processing.

```go
func (s *Streaming) StreamEvents(
    ctx context.Context,
    sessionID string,
    lastEventID *string,
    opts ...operations.Option,
) (*operations.StreamEventsResponse, error)
```

#### Parameters

| Parameter     | Type                   | Description                  |
| :------------ | :--------------------- | :--------------------------- |
| `ctx`         | `context.Context`      | Context for the request      |
| `sessionID`   | `string`               | Session ID to stream from    |
| `lastEventID` | `*string`              | Last event ID for reconnection |
| `opts`        | `...operations.Option` | Additional options           |

#### Returns

Returns `*operations.StreamEventsResponse` containing an `SSEEventStream` or an error.

#### Example - Basic Streaming

```go
ctx := context.Background()
lastEventID := ""

resp, err := client.Streaming.StreamEvents(ctx, sessionID, &lastEventID)
if err != nil {
    log.Fatalf("Failed to start stream: %v", err)
}
defer resp.SSEEventStream.Close()

// Use the EventStream API
for resp.SSEEventStream.Next() {
    event := resp.SSEEventStream.Value()
    if event == nil {
        continue
    }

    // Process event based on type
    switch event.Type {
    case components.SSEEventStreamTypeContent:
        if event.SSEContentEvent != nil {
            fmt.Printf("Content: %s\n", event.SSEContentEvent.Data.Content)
        }
    case components.SSEEventStreamTypeComplete:
        fmt.Println("Processing complete")
    case components.SSEEventStreamTypeError:
        if event.SSEErrorEvent != nil {
            fmt.Printf("Error: %s\n", event.SSEErrorEvent.Data.Error)
        }
    }
}

if err := resp.SSEEventStream.Err(); err != nil {
    log.Printf("Stream error: %v", err)
}
```

#### Example - Concurrent Streaming and Messaging

```go
var wg sync.WaitGroup
wg.Add(1)

// Start streaming in goroutine
go func() {
    defer wg.Done()

    resp, err := client.Streaming.StreamEvents(ctx, sessionID, nil)
    if err != nil {
        log.Printf("Stream failed: %v", err)
        return
    }
    defer resp.SSEEventStream.Close()

    for resp.SSEEventStream.Next() {
        event := resp.SSEEventStream.Value()
        processEvent(event)
    }
}()

// Brief delay to ensure stream is connected
time.Sleep(500 * time.Millisecond)

// Send message
client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
    Text: "Hello!",
})

// Wait for streaming to complete
wg.Wait()
```

### SSE Event Types

The streaming API emits various event types:

| Event Type                        | Description                      |
| :-------------------------------- | :------------------------------- |
| `SSEEventStreamTypeThinking`      | AI thinking event                |
| `SSEEventStreamTypeContent`       | Content generation event         |
| `SSEEventStreamTypeTool`          | Tool call event                  |
| `SSEEventStreamTypeToolExecutionStart` | Tool execution started      |
| `SSEEventStreamTypeToolExecutionComplete` | Tool execution completed |
| `SSEEventStreamTypePermission`    | Permission request event         |
| `SSEEventStreamTypeUserMessageCreated` | User message created       |
| `SSEEventStreamTypeSessionCreated` | Session created event           |
| `SSEEventStreamTypeSessionDeleted` | Session deleted event           |
| `SSEEventStreamTypeComplete`      | Processing complete              |
| `SSEEventStreamTypeError`         | Error event                      |
| `SSEEventStreamTypeHeartbeat`     | Heartbeat event                  |
| `SSEEventStreamTypeConnected`     | Connection established           |

## Files API

The `Files` resource provides file upload, download, and management operations.

### `UploadSessionFile()`

Upload a file to a session's isolated storage.

```go
func (s *Files) UploadSessionFile(
    ctx context.Context,
    sessionID string,
    request operations.UploadSessionFileRequest,
    opts ...operations.Option,
) (*operations.UploadSessionFileResponse, error)
```

#### Parameters

| Parameter   | Type                            | Description             |
| :---------- | :------------------------------ | :---------------------- |
| `ctx`       | `context.Context`               | Context for the request |
| `sessionID` | `string`                        | Session ID              |
| `request`   | `operations.UploadFileRequest`  | File data               |
| `opts`      | `...operations.Option`          | Additional options      |

### `ListSessionFiles()`

List all files in a session.

```go
func (s *Files) ListSessionFiles(
    ctx context.Context,
    sessionID string,
    opts ...operations.Option,
) (*operations.ListSessionFilesResponse, error)
```

### `GetSessionFile()`

Download a file from a session or generate a thumbnail for an image/video file.

```go
func (s *Files) GetSessionFile(
    ctx context.Context,
    id string,
    filename string,
    thumb *string,
    time *float64,
    opts ...operations.Option,
) (*operations.GetSessionFileResponse, error)
```

#### Parameters

| Parameter  | Type                   | Description                                    |
| :--------- | :--------------------- | :--------------------------------------------- |
| `ctx`      | `context.Context`      | Context for the request                        |
| `id`       | `string`               | Session ID                                     |
| `filename` | `string`               | File name to retrieve                          |
| `thumb`    | `*string`              | Thumbnail type: `"box"`, `"width"`, `"height"` |
| `time`     | `*float64`             | Timestamp for video thumbnail (seconds)        |
| `opts`     | `...operations.Option` | Additional options                             |

#### Usage Examples

**Download file:**
```go
resp, err := client.Files.GetSessionFile(ctx, sessionID, "document.pdf", nil, nil)
```

**Generate image thumbnail (box constraint):**
```go
thumb := "box"
resp, err := client.Files.GetSessionFile(ctx, sessionID, "image.jpg", &thumb, nil)
```

**Generate video thumbnail at 30 seconds:**
```go
thumb := "box"
timestamp := 30.0
resp, err := client.Files.GetSessionFile(ctx, sessionID, "video.mp4", &thumb, &timestamp)
```

### `DeleteSessionFile()`

Delete a file from a session.

```go
func (s *Files) DeleteSessionFile(
    ctx context.Context,
    sessionID string,
    filename string,
    opts ...operations.Option,
) (*operations.DeleteSessionFileResponse, error)
```

## Authentication API

The `Authentication` resource manages API keys and OAuth flows.

### `StoreAPIKey()`

Store an API key for a specific provider.

```go
func (s *Authentication) StoreAPIKey(
    ctx context.Context,
    request operations.StoreAPIKeyRequest,
    opts ...operations.Option,
) (*operations.StoreAPIKeyResponse, error)
```

#### Parameters

| Parameter | Type                              | Description             |
| :-------- | :-------------------------------- | :---------------------- |
| `ctx`     | `context.Context`                 | Context for the request |
| `request` | `operations.StoreAPIKeyRequest`   | API key data            |
| `opts`    | `...operations.Option`            | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
    Provider: "anthropic",
    APIKey:   "sk-ant-...",
})
if err != nil {
    log.Fatalf("Failed to store API key: %v", err)
}
```

### `GetAuthStatus()`

Check authentication status for all providers.

```go
func (s *Authentication) GetAuthStatus(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.GetAuthStatusResponse, error)
```

### `DeleteCredentials()`

Delete stored API key and/or OAuth credentials for a provider.

```go
func (s *Authentication) DeleteCredentials(
    ctx context.Context,
    provider string,
    opts ...operations.Option,
) (*operations.DeleteCredentialsResponse, error)
```

#### Parameters

| Parameter  | Type                   | Description             |
| :--------- | :--------------------- | :---------------------- |
| `ctx`      | `context.Context`      | Context for the request |
| `provider` | `string`               | Provider name to delete |
| `opts`     | `...operations.Option` | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.Authentication.DeleteCredentials(ctx, "anthropic")
if err != nil {
    log.Fatalf("Failed to delete credentials: %v", err)
}

fmt.Println("Credentials deleted successfully")
```

## Preferences API

The `Preferences` resource manages model and provider configuration.

### `GetPreferences()`

Retrieve current preferences.

```go
func (s *Preferences) GetPreferences(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.GetPreferencesResponse, error)
```

### `UpdatePreferences()`

Update preferences for model and provider selection.

```go
func (s *Preferences) UpdatePreferences(
    ctx context.Context,
    request operations.UpdatePreferencesRequest,
    opts ...operations.Option,
) (*operations.UpdatePreferencesResponse, error)
```

### `GetAvailableProviders()`

List all available AI providers.

```go
func (s *Preferences) GetAvailableProviders(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.GetAvailableProvidersResponse, error)
```

### `GetAvailableModels()`

List available models for a specific provider.

```go
func (s *Preferences) GetAvailableModels(
    ctx context.Context,
    provider string,
    opts ...operations.Option,
) (*operations.GetAvailableModelsResponse, error)
```

## System API

The `System` resource provides system introspection and monitoring.

### `HealthCheck()`

Check system health status.

```go
func (s *System) HealthCheck(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.HealthCheckResponse, error)
```

#### Example

```go
ctx := context.Background()
resp, err := client.System.HealthCheck(ctx)
if err != nil {
    log.Fatalf("Health check failed: %v", err)
}

fmt.Printf("Status: %s\n", *resp.Object.Status)
```

### `ListCommands()`

List all available commands.

```go
func (s *System) ListCommands(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.ListCommandsResponse, error)
```

### `ListMcpServers()`

List all Model Context Protocol servers.

```go
func (s *System) ListMcpServers(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.ListMcpServersResponse, error)
```

## Tools API

The `Tools` resource provides tool discovery and status information.

### `ListLLMTools()`

List all available LLM tools that Claude can invoke.

```go
func (s *Tools) ListLLMTools(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.ListLLMToolsResponse, error)
```

### `GetToolsStatus()`

Get status of tools by category.

```go
func (s *Tools) GetToolsStatus(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.GetToolsStatusResponse, error)
```

## Permissions API

The `Permissions` resource manages permission granting and denial.

### `GrantPermission()`

Grant a pending permission request.

```go
func (s *Permissions) GrantPermission(
    ctx context.Context,
    id string,
    opts ...operations.Option,
) (*operations.GrantPermissionResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `id`      | `string`               | Permission request ID   |
| `opts`    | `...operations.Option` | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.Permissions.GrantPermission(ctx, permissionID)
if err != nil {
    log.Fatalf("Failed to grant permission: %v", err)
}
```

### `DenyPermission()`

Deny a pending permission request.

```go
func (s *Permissions) DenyPermission(
    ctx context.Context,
    id string,
    opts ...operations.Option,
) (*operations.DenyPermissionResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `id`      | `string`               | Permission request ID   |
| `opts`    | `...operations.Option` | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.Permissions.DenyPermission(ctx, permissionID)
if err != nil {
    log.Fatalf("Failed to deny permission: %v", err)
}
```

## Type Definitions

### `SessionData`

Represents session metadata.

```go
type SessionData struct {
    ID          string
    Title       string
    SessionType string
    CreatedAt   time.Time
    UpdatedAt   *time.Time
}
```

### `BackendMessage`

Represents a message in the conversation.

```go
type BackendMessage struct {
    ID                string
    Role              string
    UserInput         string
    AssistantResponse *string
    Reasoning         *string
    ReasoningDuration *int64
    ToolCalls         []ToolCallData
    CreatedAt         time.Time
}
```

### `ExportSession`

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

### `ExportMessage`

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
}
```

### `ToolCallData`

Tool call information.

```go
type ToolCallData struct {
    ID     string
    Name   string
    Input  map[string]interface{}
    Output *string
}
```

### `FileInfo`

File metadata.

```go
type FileInfo struct {
    Filename string
    Size     int64
    MimeType string
    UploadedAt time.Time
}
```

### `Callback`

Session callback configuration.

```go
type Callback struct {
    Type   CallbackType
    Config map[string]interface{}
}
```

## Error Handling

All SDK methods return an error as the second return value. Always check for errors:

```go
resp, err := client.Sessions.CreateSession(ctx, request)
if err != nil {
    // Handle error
    log.Fatalf("Operation failed: %v", err)
}
```

### Error Types

The SDK uses the `apierrors` package for API-specific errors:

```go
import "github.com/recreate-run/mix-go-sdk/models/apierrors"
```

#### `ErrorResponse`

Standard API error response.

```go
type ErrorResponse struct {
    HTTPMeta HTTPMetadata
    Message  string
    Code     *int
}
```

#### `APIError`

Generic API error.

```go
type APIError struct {
    Message    string
    StatusCode int
    Body       string
    Response   *http.Response
}
```

## Advanced Examples

### Complete Session Workflow

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
        mix.WithServerURL("http://localhost:8088"),
        mix.WithTimeout(30 * time.Second),
    )

    ctx := context.Background()

    // Create session
    createResp, err := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Title: "Example Session",
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
    client := mix.New(
        mix.WithServerURL("http://localhost:8088"),
    )

    ctx := context.Background()

    // Create session
    createResp, _ := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Title: "Streaming Example",
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
    "io/ioutil"
    "log"
    "os"

    "github.com/recreate-run/mix-go-sdk"
    "github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
    client := mix.New(
        mix.WithServerURL("http://localhost:8088"),
    )

    ctx := context.Background()

    // Create session
    createResp, _ := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Title: "File Example",
    })
    sessionID := createResp.SessionData.ID
    defer client.Sessions.DeleteSession(ctx, sessionID)

    // Read file
    fileData, err := ioutil.ReadFile("example.txt")
    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }

    // Upload file
    uploadResp, err := client.Files.UploadSessionFile(ctx, sessionID, operations.UploadSessionFileRequest{
        Filename: "example.txt",
        Content:  fileData,
    })
    if err != nil {
        log.Fatalf("Failed to upload file: %v", err)
    }
    fmt.Printf("Uploaded: %s\n", uploadResp.FileInfo.Filename)

    // List files
    listResp, err := client.Files.ListSessionFiles(ctx, sessionID)
    if err != nil {
        log.Fatalf("Failed to list files: %v", err)
    }

    for _, file := range listResp.Files {
        fmt.Printf("- %s (%d bytes)\n", file.Filename, file.Size)
    }

    // Download file
    downloadResp, err := client.Files.GetSessionFile(ctx, sessionID, "example.txt", nil, nil)
    if err != nil {
        log.Fatalf("Failed to download file: %v", err)
    }

    // Save downloaded file
    err = ioutil.WriteFile("downloaded.txt", downloadResp.Content, 0644)
    if err != nil {
        log.Fatalf("Failed to save file: %v", err)
    }
    fmt.Println("File downloaded successfully")
}
```

### Session Forking and Rewinding

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
    client := mix.New(
        mix.WithServerURL("http://localhost:8088"),
    )

    ctx := context.Background()

    // Create original session
    createResp, _ := client.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Title: "Original Session",
    })
    sessionID := createResp.SessionData.ID
    defer client.Sessions.DeleteSession(ctx, sessionID)

    // Send multiple messages
    messages := []string{
        "What is 2 + 2?",
        "What is the capital of France?",
        "Tell me a joke",
    }

    for _, text := range messages {
        client.Messages.SendMessage(ctx, sessionID, operations.SendMessageRequestBody{
            Text: text,
        })
        time.Sleep(500 * time.Millisecond)
    }

    // Wait for processing
    time.Sleep(3 * time.Second)

    // Fork at message index 1
    forkResp, err := client.Sessions.ForkSession(ctx, sessionID, operations.ForkSessionRequestBody{
        MessageIndex: 1,
        Title:        mix.String("Forked Session"),
    })
    if err != nil {
        log.Fatalf("Failed to fork: %v", err)
    }
    forkedSessionID := forkResp.SessionData.ID
    defer client.Sessions.DeleteSession(ctx, forkedSessionID)

    fmt.Printf("Forked session: %s\n", forkedSessionID)

    // Get messages from original session
    messagesResp, _ := client.Messages.GetSessionMessages(ctx, sessionID)
    if len(messagesResp.BackendMessages) > 0 {
        // Rewind to first message
        firstMessageID := messagesResp.BackendMessages[0].ID
        rewindResp, err := client.Sessions.RewindSession(ctx, sessionID, operations.RewindSessionRequestBody{
            MessageID: firstMessageID,
        })
        if err != nil {
            log.Fatalf("Failed to rewind: %v", err)
        }
        fmt.Println("Session rewound successfully")
    }
}
```

### Multi-Provider Authentication

```go
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
    client := mix.New(
        mix.WithServerURL("http://localhost:8088"),
    )

    ctx := context.Background()

    // Store API keys for multiple providers
    providers := map[string]string{
        "anthropic":  os.Getenv("ANTHROPIC_API_KEY"),
        "openai":     os.Getenv("OPENAI_API_KEY"),
        "gemini":     os.Getenv("GEMINI_API_KEY"),
    }

    for provider, apiKey := range providers {
        if apiKey != "" {
            _, err := client.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
                Provider: provider,
                APIKey:   apiKey,
            })
            if err != nil {
                log.Printf("Failed to store %s key: %v", provider, err)
            } else {
                fmt.Printf("Stored API key for %s\n", provider)
            }
        }
    }

    // Check authentication status
    statusResp, err := client.Authentication.GetAuthStatus(ctx)
    if err != nil {
        log.Fatalf("Failed to get auth status: %v", err)
    }

    fmt.Println("\nAuthentication Status:")
    for provider, status := range statusResp.Status {
        fmt.Printf("- %s: %v\n", provider, status)
    }

    // Get available models for authenticated providers
    providersResp, _ := client.Preferences.GetAvailableProviders(ctx)
    for _, provider := range providersResp.Providers {
        modelsResp, err := client.Preferences.GetAvailableModels(ctx, provider)
        if err != nil {
            continue
        }
        fmt.Printf("\n%s models: %v\n", provider, modelsResp.Models)
    }
}
```

## Request Options

All API methods accept optional `operations.Option` parameters for fine-grained control:

### `operations.WithRetries()`

Override retry configuration for a specific request.

```go
resp, err := client.Sessions.CreateSession(ctx, request,
    operations.WithRetries(retry.Config{
        Strategy: "backoff",
        Backoff: &retry.BackoffStrategy{
            InitialInterval: 1000,
            MaxInterval:     30000,
            Exponent:        2.0,
        },
    }),
)
```

### `operations.WithTimeout()`

Set a custom timeout for a specific request.

```go
resp, err := client.Messages.SendMessage(ctx, sessionID, requestBody,
    operations.WithTimeout(60 * time.Second),
)
```

### `operations.WithServerURL()`

Override server URL for a specific request.

```go
resp, err := client.Sessions.ListSessions(ctx, nil,
    operations.WithServerURL("http://alternative-server:8088"),
)
```

### `operations.WithHeaders()`

Add custom headers to a request.

```go
resp, err := client.Sessions.CreateSession(ctx, request,
    operations.WithHeaders(map[string]string{
        "X-Custom-Header": "value",
    }),
)
```

## Best Practices

### 1. Always Use Context

```go
// Use background context
ctx := context.Background()

// Or use timeout context for long operations
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
```

### 2. Cleanup Resources

```go
// Always cleanup sessions
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
resp, err := client.Sessions.CreateSession(ctx, request)
if err != nil {
    log.Printf("Non-fatal error: %v", err)
    // Handle gracefully
}
```

### 5. Proper Stream Cleanup

```go
streamResp, err := client.Streaming.StreamEvents(ctx, sessionID, nil)
if err != nil {
    return err
}
defer streamResp.SSEEventStream.Close()
```

## Troubleshooting

### Connection Errors

- Verify server URL is correct
- Ensure the Mix server is running
- Check network connectivity

### Authentication Errors

- Verify API keys are correct
- Check provider authentication status
- Ensure credentials are stored properly

### Streaming Issues

- Use proper timeout contexts
- Ensure goroutine synchronization
- Check for stream close/cleanup

### File Operation Errors

- Verify file paths exist
- Check file permissions
- Ensure files are within size limits

## See Also

- [Go SDK Examples](/examples/) - Comprehensive examples directory
- [API Models](/models/) - Data structures and types
- [Operations](/models/operations/) - Request and response types
- [Components](/models/components/) - Shared component types

## Contributing

When adding new features:
1. Follow existing code style
2. Include comprehensive comments
3. Demonstrate error handling
4. Add cleanup patterns
5. Update documentation

---

**Version:** 0.0.1
**Generated:** 2025

**Note:** Complete OAuth authentication methods and related authentication information are not included in this reference documentation.
