# Streaming

## Overview

### Available Operations

* [StreamEvents](#streamevents) - Server-Sent Events stream for real-time updates

## StreamEvents

Establishes a persistent SSE connection for receiving real-time updates during message processing. Connection remains open for multiple messages and includes proper reconnection support with Last-Event-ID header.

### Example Usage

<!-- UsageSnippet language="go" operationID="streamEvents" method="get" path="/stream" -->
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

    res, err := s.Streaming.StreamEvents(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SSEEventStream != nil {
        defer res.SSEEventStream.Close()

        for res.SSEEventStream.Next() {
            event := res.SSEEventStream.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `sessionID`                                              | *string*                                                 | :heavy_check_mark:                                       | Session ID to stream events for                          |
| `lastEventID`                                            | **string*                                                | :heavy_minus_sign:                                       | Last received event ID for reconnection and event replay |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.StreamEventsResponse](../../models/operations/streameventsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |