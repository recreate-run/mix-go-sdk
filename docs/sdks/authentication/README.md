# Authentication

## Overview

### Available Operations

* [StoreAPIKey](#storeapikey) - Store API key
* [HandleOAuthCallback](#handleoauthcallback) - Handle OAuth callback
* [StartOAuthFlow](#startoauthflow) - Start OAuth authentication
* [GetAuthStatus](#getauthstatus) - Get authentication status
* [ValidatePreferredProvider](#validatepreferredprovider) - Validate preferred provider
* [DeleteCredentials](#deletecredentials) - Delete provider credentials
* [GetOAuthHealth](#getoauthhealth) - Get OAuth authentication health
* [RefreshOAuthTokens](#refreshoauthtokens) - Manually refresh OAuth tokens

## StoreAPIKey

Store API key for direct authentication with a specific provider

### Example Usage

<!-- UsageSnippet language="go" operationID="storeApiKey" method="post" path="/api/auth/api-key" -->
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

    res, err := s.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
        APIKey: "<value>",
        Provider: operations.ProviderBrave,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.StoreAPIKeyRequest](../../models/operations/storeapikeyrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.StoreAPIKeyResponse](../../models/operations/storeapikeyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## HandleOAuthCallback

Process OAuth callback and exchange code for access token

### Example Usage

<!-- UsageSnippet language="go" operationID="handleOAuthCallback" method="post" path="/api/auth/oauth-callback" -->
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

    res, err := s.Authentication.HandleOAuthCallback(ctx, operations.HandleOAuthCallbackRequest{
        Code: "<value>",
        Provider: "<value>",
        State: "Arizona",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.HandleOAuthCallbackRequest](../../models/operations/handleoauthcallbackrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.HandleOAuthCallbackResponse](../../models/operations/handleoauthcallbackresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## StartOAuthFlow

Initiate OAuth authentication flow for a specific provider

### Example Usage

<!-- UsageSnippet language="go" operationID="startOAuthFlow" method="post" path="/api/auth/oauth/{provider}" -->
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

    res, err := s.Authentication.StartOAuthFlow(ctx, "<value>")
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
| `provider`                                               | *string*                                                 | :heavy_check_mark:                                       | Provider name (currently only 'anthropic')               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.StartOAuthFlowResponse](../../models/operations/startoauthflowresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetAuthStatus

Get authentication status for all supported providers

### Example Usage

<!-- UsageSnippet language="go" operationID="getAuthStatus" method="get" path="/api/auth/status" -->
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

    res, err := s.Authentication.GetAuthStatus(ctx)
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

**[*operations.GetAuthStatusResponse](../../models/operations/getauthstatusresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ValidatePreferredProvider

Check if the user's preferred provider is authenticated

### Example Usage

<!-- UsageSnippet language="go" operationID="validatePreferredProvider" method="get" path="/api/auth/validate" -->
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

    res, err := s.Authentication.ValidatePreferredProvider(ctx)
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

**[*operations.ValidatePreferredProviderResponse](../../models/operations/validatepreferredproviderresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteCredentials

Delete stored API key and/or OAuth credentials for a provider

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCredentials" method="delete" path="/api/auth/{provider}" -->
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

    res, err := s.Authentication.DeleteCredentials(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `provider`                                                   | *string*                                                     | :heavy_check_mark:                                           | Provider name (anthropic, openai, openrouter, gemini, brave) |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.DeleteCredentialsResponse](../../models/operations/deletecredentialsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

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

    res, err := s.Authentication.GetOAuthHealth(ctx)
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

    res, err := s.Authentication.RefreshOAuthTokens(ctx)
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