# Messages
(*Messages*)

## Overview

### Available Operations

* [GetMessageHistory](#getmessagehistory) - Get global message history
* [CancelSessionProcessing](#cancelsessionprocessing) - Cancel agent processing
* [GetSessionMessages](#getsessionmessages) - List session messages
* [SendMessage](#sendmessage) - Send a message to session (async)

## GetMessageHistory

Retrieve message history across all sessions with optional pagination

### Example Usage

<!-- UsageSnippet language="go" operationID="getMessageHistory" method="get" path="/api/messages/history" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New()

    res, err := s.Messages.GetMessageHistory(ctx, mix.Pointer[int64](50), mix.Pointer[int64](0))
    if err != nil {
        log.Fatal(err)
    }
    if res.BackendMessages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | Maximum number of messages to return                     |
| `offset`                                                 | **int64*                                                 | :heavy_minus_sign:                                       | Number of messages to skip                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMessageHistoryResponse](../../models/operations/getmessagehistoryresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CancelSessionProcessing

Cancel any ongoing agent processing in the specified session

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelSessionProcessing" method="post" path="/api/sessions/{id}/cancel" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New()

    res, err := s.Messages.CancelSessionProcessing(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.CancelSessionProcessingResponse](../../models/operations/cancelsessionprocessingresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetSessionMessages

Retrieve all messages from a specific session

### Example Usage

<!-- UsageSnippet language="go" operationID="getSessionMessages" method="get" path="/api/sessions/{id}/messages" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New()

    res, err := s.Messages.GetSessionMessages(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.BackendMessages != nil {
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

**[*operations.GetSessionMessagesResponse](../../models/operations/getsessionmessagesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SendMessage

Send a user message to a specific session for AI processing. Returns immediately with 202 Accepted. All results stream via Server-Sent Events (SSE) connection.

### Example Usage

<!-- UsageSnippet language="go" operationID="sendMessage" method="post" path="/api/sessions/{id}/messages" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"github.com/recreate-run/mix-go-sdk/optionalnullable"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New()

    res, err := s.Messages.SendMessage(ctx, "<id>", operations.SendMessageRequestBody{
        Text: "<value>",
        ThinkingLevel: optionalnullable.From(mix.Pointer(operations.ThinkingLevelMedium.ToPointer())),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | Session ID                                                                             |
| `requestBody`                                                                          | [operations.SendMessageRequestBody](../../models/operations/sendmessagerequestbody.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.SendMessageResponse](../../models/operations/sendmessageresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |