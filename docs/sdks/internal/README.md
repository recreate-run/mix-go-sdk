# Internal
(*Internal*)

## Overview

### Available Operations

* [RefreshOAuthTokens](#refreshoauthtokens) - Manually refresh OAuth tokens

## RefreshOAuthTokens

Manually trigger OAuth token refresh for all expired tokens. Normally tokens are refreshed automatically by the background service every 30 minutes.

### Example Usage

<!-- UsageSnippet language="go" operationID="refreshOAuthTokens" method="post" path="/internal/auth/refresh-tokens" -->
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

    res, err := s.Internal.RefreshOAuthTokens(ctx)
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

**[*operations.RefreshOAuthTokensResponse](../../models/operations/refreshoauthtokensresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |