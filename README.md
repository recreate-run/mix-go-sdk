# mix

Developer-friendly & type-safe Go SDK specifically catered to leverage *mix* API.

<div align="left" style="margin-bottom: 0;">
    <a href="https://www.speakeasy.com/?utm_source=mix&utm_campaign=go" class="badge-link">
        <span class="badge-container">
            <span class="badge-icon-section">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 30 30" fill="none" style="vertical-align: middle;"><title>Speakeasy Logo</title><path fill="currentColor" d="m20.639 27.548-19.17-2.724L0 26.1l20.639 2.931 8.456-7.336-1.468-.208-6.988 6.062Z"></path><path fill="currentColor" d="m20.639 23.1 8.456-7.336-1.468-.207-6.988 6.06-6.84-.972-9.394-1.333-2.936-.417L0 20.169l2.937.416L0 23.132l20.639 2.931 8.456-7.334-1.468-.208-6.986 6.062-9.78-1.39 1.468-1.273 8.31 1.18Z"></path><path fill="currentColor" d="m20.639 18.65-19.17-2.724L0 17.201l20.639 2.931 8.456-7.334-1.468-.208-6.988 6.06Z"></path><path fill="currentColor" d="M27.627 6.658 24.69 9.205 20.64 12.72l-7.923-1.126L1.469 9.996 0 11.271l11.246 1.596-1.467 1.275-8.311-1.181L0 14.235l20.639 2.932 8.456-7.334-2.937-.418 2.937-2.549-1.468-.208Z"></path><path fill="currentColor" d="M29.095 3.902 8.456.971 0 8.305l20.639 2.934 8.456-7.337Z"></path></svg>
            </span>
            <span class="badge-text badge-text-section">BUILT BY SPEAKEASY</span>
        </span>
    </a>
    <a href="https://opensource.org/licenses/MIT" class="badge-link">
        <span class="badge-container blue">
            <span class="badge-text badge-text-section">LICENSE // MIT</span>
        </span>
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/recreate/mix). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Mix REST API: REST API for the Mix application - session management, messaging, and system operations
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [mix](#mix)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Server-sent event streaming](#server-sent-event-streaming)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/recreate-run/mix-go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := mix.New()

	res, err := s.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
		APIKey:   "<value>",
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Authentication](docs/sdks/authentication/README.md)

* [StoreAPIKey](docs/sdks/authentication/README.md#storeapikey) - Store API key
* [HandleOAuthCallback](docs/sdks/authentication/README.md#handleoauthcallback) - Handle OAuth callback
* [StartOAuthFlow](docs/sdks/authentication/README.md#startoauthflow) - Start OAuth authentication
* [GetAuthStatus](docs/sdks/authentication/README.md#getauthstatus) - Get authentication status
* [ValidatePreferredProvider](docs/sdks/authentication/README.md#validatepreferredprovider) - Validate preferred provider
* [DeleteCredentials](docs/sdks/authentication/README.md#deletecredentials) - Delete provider credentials
* [GetOAuthHealth](docs/sdks/authentication/README.md#getoauthhealth) - Get OAuth authentication health
* [RefreshOAuthTokens](docs/sdks/authentication/README.md#refreshoauthtokens) - Manually refresh OAuth tokens

### [Files](docs/sdks/files/README.md)

* [ListSessionFiles](docs/sdks/files/README.md#listsessionfiles) - List session files
* [UploadSessionFile](docs/sdks/files/README.md#uploadsessionfile) - Upload file to session
* [DeleteSessionFile](docs/sdks/files/README.md#deletesessionfile) - Delete session file
* [GetSessionFile](docs/sdks/files/README.md#getsessionfile) - Get session file

### [Health](docs/sdks/health/README.md)

* [GetOAuthHealth](docs/sdks/health/README.md#getoauthhealth) - Get OAuth authentication health

### [Internal](docs/sdks/internal/README.md)

* [RefreshOAuthTokens](docs/sdks/internal/README.md#refreshoauthtokens) - Manually refresh OAuth tokens

### [Messages](docs/sdks/messages/README.md)

* [GetMessageHistory](docs/sdks/messages/README.md#getmessagehistory) - Get global message history
* [CancelSessionProcessing](docs/sdks/messages/README.md#cancelsessionprocessing) - Cancel agent processing
* [GetSessionMessages](docs/sdks/messages/README.md#getsessionmessages) - List session messages
* [SendMessage](docs/sdks/messages/README.md#sendmessage) - Send a message to session (async)

### [Permissions](docs/sdks/permissions/README.md)

* [DenyPermission](docs/sdks/permissions/README.md#denypermission) - Deny permission
* [GrantPermission](docs/sdks/permissions/README.md#grantpermission) - Grant permission

### [Preferences](docs/sdks/preferences/README.md)

* [GetPreferences](docs/sdks/preferences/README.md#getpreferences) - Get user preferences
* [UpdatePreferences](docs/sdks/preferences/README.md#updatepreferences) - Update user preferences
* [GetAvailableProviders](docs/sdks/preferences/README.md#getavailableproviders) - Get available providers
* [ResetPreferences](docs/sdks/preferences/README.md#resetpreferences) - Reset preferences

### [Sessions](docs/sdks/sessions/README.md)

* [ListSessions](docs/sdks/sessions/README.md#listsessions) - List all sessions
* [CreateSession](docs/sdks/sessions/README.md#createsession) - Create a new session
* [DeleteSession](docs/sdks/sessions/README.md#deletesession) - Delete a session
* [GetSession](docs/sdks/sessions/README.md#getsession) - Get a specific session
* [UpdateSessionCallbacks](docs/sdks/sessions/README.md#updatesessioncallbacks) - Update session callbacks
* [ExportSession](docs/sdks/sessions/README.md#exportsession) - Export session transcript
* [ForkSession](docs/sdks/sessions/README.md#forksession) - Fork a session
* [RewindSession](docs/sdks/sessions/README.md#rewindsession) - Rewind a session

### [Streaming](docs/sdks/streaming/README.md)

* [StreamEvents](docs/sdks/streaming/README.md#streamevents) - Server-Sent Events stream for real-time updates

### [System](docs/sdks/system/README.md)

* [ListCommands](docs/sdks/system/README.md#listcommands) - List available commands
* [GetCommand](docs/sdks/system/README.md#getcommand) - Get specific command
* [ListMcpServers](docs/sdks/system/README.md#listmcpservers) - List MCP servers
* [GetSystemInfo](docs/sdks/system/README.md#getsysteminfo) - Get system information
* [HealthCheck](docs/sdks/system/README.md#healthcheck) - Health check

### [Tools](docs/sdks/tools/README.md)

* [ListLLMTools](docs/sdks/tools/README.md#listllmtools) - List LLM tools
* [GetToolCredentialsStatus](docs/sdks/tools/README.md#gettoolcredentialsstatus) - Get tool credentials status
* [GetToolsStatus](docs/sdks/tools/README.md#gettoolsstatus) - Get tools status

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Server-sent event streaming [eventstream] -->
## Server-sent event streaming

[Server-sent events][mdn-sse] are used to stream content from certain
operations. These operations will expose the stream as an iterable that
can be consumed using a simple `for` loop. The loop will
terminate when the server no longer has any events to send and closes the
underlying connection.

```go
package main

import (
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"log"
)

func main() {
	ctx := context.Background()

	s := mix.New()

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

[mdn-sse]: https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
<!-- End Server-sent event streaming [eventstream] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"github.com/recreate-run/mix-go-sdk/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := mix.New()

	res, err := s.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
		APIKey:   "<value>",
		Provider: operations.ProviderBrave,
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"github.com/recreate-run/mix-go-sdk/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := mix.New(
		mix.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
	)

	res, err := s.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
		APIKey:   "<value>",
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `StoreAPIKey` function may return the following errors:

| Error Type              | Status Code | Content Type     |
| ----------------------- | ----------- | ---------------- |
| apierrors.ErrorResponse | 400         | application/json |
| apierrors.ErrorResponse | 500         | application/json |
| apierrors.APIError      | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	mix "github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/apierrors"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := mix.New()

	res, err := s.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
		APIKey:   "<value>",
		Provider: operations.ProviderBrave,
	})
	if err != nil {

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	mix "github.com/recreate-run/mix-go-sdk"
	"github.com/recreate-run/mix-go-sdk/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := mix.New(
		mix.WithServerURL("http://localhost:8088"),
	)

	res, err := s.Authentication.StoreAPIKey(ctx, operations.StoreAPIKeyRequest{
		APIKey:   "<value>",
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/recreate-run/mix-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = mix.New(mix.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=mix&utm_campaign=go)

<style>
  :root {
    --badge-gray-bg: #f3f4f6;
    --badge-gray-border: #d1d5db;
    --badge-gray-text: #374151;
    --badge-blue-bg: #eff6ff;
    --badge-blue-border: #3b82f6;
    --badge-blue-text: #3b82f6;
  }

  @media (prefers-color-scheme: dark) {
    :root {
      --badge-gray-bg: #374151;
      --badge-gray-border: #4b5563;
      --badge-gray-text: #f3f4f6;
      --badge-blue-bg: #1e3a8a;
      --badge-blue-border: #3b82f6;
      --badge-blue-text: #93c5fd;
    }
  }
  
  h1 {
    border-bottom: none !important;
    margin-bottom: 4px;
    margin-top: 0;
    letter-spacing: 0.5px;
    font-weight: 600;
  }
  
  .badge-text {
    letter-spacing: 1px;
    font-weight: 300;
  }
  
  .badge-container {
    display: inline-flex;
    align-items: center;
    background: var(--badge-gray-bg);
    border: 1px solid var(--badge-gray-border);
    border-radius: 6px;
    overflow: hidden;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif;
    font-size: 11px;
    text-decoration: none;
    vertical-align: middle;
  }

  .badge-container.blue {
    background: var(--badge-blue-bg);
    border-color: var(--badge-blue-border);
  }

  .badge-icon-section {
    padding: 4px 8px;
    border-right: 1px solid var(--badge-gray-border);
    display: flex;
    align-items: center;
  }

  .badge-text-section {
    padding: 4px 10px;
    color: var(--badge-gray-text);
    font-weight: 400;
  }

  .badge-container.blue .badge-text-section {
    color: var(--badge-blue-text);
  }
  
  .badge-link {
    text-decoration: none;
    margin-left: 8px;
    display: inline-flex;
    vertical-align: middle;
  }

  .badge-link:hover {
    text-decoration: none;
  }
  
  .badge-link:first-child {
    margin-left: 0;
  }
  
  .badge-icon-section svg {
    color: var(--badge-gray-text);
  }

  .badge-container.blue .badge-icon-section svg {
    color: var(--badge-blue-text);
  }
</style> 