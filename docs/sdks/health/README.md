# Health
(*Health*)

## Overview

### Available Operations

* [GetOAuthHealth](#getoauthhealth) - Get OAuth authentication health

## GetOAuthHealth

Get health status of all OAuth credentials. Background service refreshes tokens 35 minutes before expiry. API calls mark tokens expired 5 minutes before expiry. Health statuses: 'healthy' (tokens valid, >5min remaining), 'degraded' (some tokens within 5min of expiry but refreshable), 'unhealthy' (tokens expired without refresh capability)

### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthHealth" method="get" path="/health/auth" -->
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

    res, err := s.Health.GetOAuthHealth(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOAuthHealthResponse](../../models/operations/getoauthhealthresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |