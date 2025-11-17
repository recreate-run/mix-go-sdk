# Files
(*Files*)

## Overview

### Available Operations

* [ListSessionFiles](#listsessionfiles) - List session files
* [UploadSessionFile](#uploadsessionfile) - Upload file to session
* [DeleteSessionFile](#deletesessionfile) - Delete session file
* [GetSessionFile](#getsessionfile) - Get session file

## ListSessionFiles

List all files in session storage directory

### Example Usage

<!-- UsageSnippet language="go" operationID="listSessionFiles" method="get" path="/api/sessions/{id}/files" -->
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

    res, err := s.Files.ListSessionFiles(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.FileInfos != nil {
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

**[*operations.ListSessionFilesResponse](../../models/operations/listsessionfilesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UploadSessionFile

Upload a file to session-specific storage directory

### Example Usage

<!-- UsageSnippet language="go" operationID="uploadSessionFile" method="post" path="/api/sessions/{id}/files/upload" -->
```go
package main

import(
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"os"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := mix.New(
        "https://api.example.com",
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Files.UploadSessionFile(ctx, "<id>", operations.UploadSessionFileRequestBody{
        File: operations.File{
            FileName: "example.file",
            Content: example,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FileInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | Session ID                                                                                         |
| `requestBody`                                                                                      | [operations.UploadSessionFileRequestBody](../../models/operations/uploadsessionfilerequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UploadSessionFileResponse](../../models/operations/uploadsessionfileresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 413           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteSessionFile

Delete a specific file from session storage. Only files are supported - directories cannot be deleted.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSessionFile" method="delete" path="/api/sessions/{id}/files/{filename}" -->
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

    res, err := s.Files.DeleteSessionFile(ctx, "<id>", "example.file")
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
| `filename`                                               | *string*                                                 | :heavy_check_mark:                                       | Filename to delete                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteSessionFileResponse](../../models/operations/deletesessionfileresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetSessionFile

Download or serve a specific file from session storage. Supports thumbnail generation with ?thumb parameter.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSessionFile" method="get" path="/api/sessions/{id}/files/{filename}" -->
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

    res, err := s.Files.GetSessionFile(ctx, "<id>", "example.file", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `id`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | Session ID                                                            |
| `filename`                                                            | *string*                                                              | :heavy_check_mark:                                                    | Filename to retrieve                                                  |
| `thumb`                                                               | **string*                                                             | :heavy_minus_sign:                                                    | Thumbnail specification: '100' (box), 'w100' (width), 'h100' (height) |
| `time`                                                                | **float64*                                                            | :heavy_minus_sign:                                                    | Time offset in seconds for video thumbnails (default: 1.0)            |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*operations.GetSessionFileResponse](../../models/operations/getsessionfileresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |