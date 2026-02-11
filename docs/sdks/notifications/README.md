# Notifications

## Overview

### Available Operations

* [RespondToNotification](#respondtonotification) - Respond to notification

## RespondToNotification

Send user's response to a notification request

### Example Usage

<!-- UsageSnippet language="go" operationID="respondToNotification" method="post" path="/api/notifications/{id}/respond" -->
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

    res, err := s.Notifications.RespondToNotification(ctx, "<id>", operations.RespondToNotificationRequestBody{
        Type: operations.TypeAcknowledge,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `id`                                                                                                       | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Notification ID                                                                                            |
| `requestBody`                                                                                              | [operations.RespondToNotificationRequestBody](../../models/operations/respondtonotificationrequestbody.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RespondToNotificationResponse](../../models/operations/respondtonotificationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 404                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |