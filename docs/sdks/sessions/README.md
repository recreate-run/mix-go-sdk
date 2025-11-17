# Sessions
(*Sessions*)

## Overview

### Available Operations

* [ListSessions](#listsessions) - List all sessions
* [CreateSession](#createsession) - Create a new session
* [DeleteSession](#deletesession) - Delete a session
* [GetSession](#getsession) - Get a specific session
* [UpdateSessionCallbacks](#updatesessioncallbacks) - Update session callbacks
* [ExportSession](#exportsession) - Export session transcript
* [RewindSession](#rewindsession) - Rewind a session

## ListSessions

Retrieve a list of all available sessions with their metadata

### Example Usage

<!-- UsageSnippet language="go" operationID="listSessions" method="get" path="/api/sessions" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New(
        "https://api.example.com",
    )

    res, err := s.Sessions.ListSessions(ctx, mix.Pointer(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `includeSubagents`                                                                   | **bool*                                                                              | :heavy_minus_sign:                                                                   | Include subagent sessions in response (default: false, subagent sessions are hidden) |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListSessionsResponse](../../models/operations/listsessionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreateSession

Create a new session with required title and optional custom system prompt. Session automatically gets isolated storage directory. Supports session-level callbacks for automated actions after tool execution.

### Example Usage

<!-- UsageSnippet language="go" operationID="createSession" method="post" path="/api/sessions" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/components"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New(
        "https://api.example.com",
    )

    res, err := s.Sessions.CreateSession(ctx, operations.CreateSessionRequest{
        Callbacks: []components.Callback{
            components.Callback{
                MessageContent: mix.Pointer("Please review the changes and run tests"),
                Name: mix.Pointer("Log Output"),
                ToolName: "*",
                Type: components.CallbackTypeSendMessage,
            },
        },
        CustomSystemPrompt: mix.Pointer("You are a helpful assistant specialized in $<domain>. Always be concise and accurate."),
        PromptMode: operations.PromptModeAppend.ToPointer(),
        SubagentType: mix.Pointer(""),
        Title: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateSessionRequest](../../models/operations/createsessionrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateSessionResponse](../../models/operations/createsessionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteSession

Permanently delete a session and all its data

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSession" method="delete" path="/api/sessions/{id}" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New(
        "https://api.example.com",
    )

    res, err := s.Sessions.DeleteSession(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Session ID                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteSessionResponse](../../models/operations/deletesessionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetSession

Retrieve detailed information about a specific session

### Example Usage

<!-- UsageSnippet language="go" operationID="getSession" method="get" path="/api/sessions/{id}" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New(
        "https://api.example.com",
    )

    res, err := s.Sessions.GetSession(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Session ID                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSessionResponse](../../models/operations/getsessionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateSessionCallbacks

Update the callback configurations for a session. Callbacks execute automatically after tool completion. Pass an empty array to clear all callbacks.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSessionCallbacks" method="patch" path="/api/sessions/{id}/callbacks" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/components"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New(
        "https://api.example.com",
    )

    res, err := s.Sessions.UpdateSessionCallbacks(ctx, "<id>", operations.UpdateSessionCallbacksRequestBody{
        Callbacks: []components.Callback{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `id`                                                                                                         | *string*                                                                                                     | :heavy_check_mark:                                                                                           | Session ID to update                                                                                         |
| `requestBody`                                                                                                | [operations.UpdateSessionCallbacksRequestBody](../../models/operations/updatesessioncallbacksrequestbody.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateSessionCallbacksResponse](../../models/operations/updatesessioncallbacksresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ExportSession

Export complete session transcript with all messages, tool calls, reasoning, and metadata as JSON

### Example Usage

<!-- UsageSnippet language="go" operationID="exportSession" method="get" path="/api/sessions/{id}/export" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New(
        "https://api.example.com",
    )

    res, err := s.Sessions.ExportSession(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ExportSession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Session ID to export                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ExportSessionResponse](../../models/operations/exportsessionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## RewindSession

Delete messages after a specified message in the current session, optionally cleaning up media files created after that point

### Example Usage

<!-- UsageSnippet language="go" operationID="rewindSession" method="post" path="/api/sessions/{id}/rewind" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New(
        "https://api.example.com",
    )

    res, err := s.Sessions.RewindSession(ctx, "<id>", operations.RewindSessionRequestBody{
        MessageID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `id`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | Session ID to rewind                                                                       |
| `requestBody`                                                                              | [operations.RewindSessionRequestBody](../../models/operations/rewindsessionrequestbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RewindSessionResponse](../../models/operations/rewindsessionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |