# Tools

## Overview

### Available Operations

* [ListLLMTools](#listllmtools) - List LLM tools
* [GetToolCredentialsStatus](#gettoolcredentialsstatus) - Get tool credentials status
* [GetToolsStatus](#gettoolsstatus) - Get tools status

## ListLLMTools

Returns the list of all LLM tools that Claude can invoke. The list is dynamically extracted from the actual tools registered in CoderAgentTools() (agent/tools.go), ensuring it always reflects the current tool availability. Typical tools include: Bash, Edit, Read, Write, Grep, Glob, WebFetch, WebSearch, ReadMedia, TodoWrite, ExitPlanMode, and Task. This endpoint is useful for creating tool callbacks or understanding available agent capabilities.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLLMTools" method="get" path="/api/tools" -->
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

    res, err := s.Tools.ListLLMTools(ctx)
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

**[*operations.ListLLMToolsResponse](../../models/operations/listllmtoolsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetToolCredentialsStatus

Returns authentication/credential status for external tool integrations (Brave Search, Gemini Vision, etc.). This endpoint checks if API keys are configured for tools that require external service credentials.

### Example Usage

<!-- UsageSnippet language="go" operationID="getToolCredentialsStatus" method="get" path="/api/tools/credentials-status" -->
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

    res, err := s.Tools.GetToolCredentialsStatus(ctx)
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

**[*operations.GetToolCredentialsStatusResponse](../../models/operations/gettoolcredentialsstatusresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetToolsStatus

Get status and authentication information for all available tools and categories

### Example Usage

<!-- UsageSnippet language="go" operationID="getToolsStatus" method="get" path="/api/tools/status" -->
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

    res, err := s.Tools.GetToolsStatus(ctx)
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

**[*operations.GetToolsStatusResponse](../../models/operations/gettoolsstatusresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |