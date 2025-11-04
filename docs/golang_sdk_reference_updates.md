# Mix Go SDK Documentation - Required Updates

> This file contains critical updates that need to be added to `golang_sdk_reference.md`

## 1. SDK Metadata Section (Add after "Initialization")

### SDK Information

The SDK includes the following metadata:

| Property     | Value                                                                          |
| :----------- | :----------------------------------------------------------------------------- |
| Version      | `0.0.1`                                                                        |
| User Agent   | `speakeasy-sdk/go 0.0.1 2.730.0 1.0.0 github.com/recreate-run/mix-go-sdk`     |
| Server List  | `["http://localhost:8088"]`                                                    |

Access SDK version:
```go
client := mix.New()
fmt.Println(client.SDKVersion) // "0.0.1"
```

---

## 2. SDK Structure Table - Add Internal Resource

Add this row to the SDK Structure table:

| Resource       | Description                              |
| :------------- | :--------------------------------------- |
| `Internal`     | Internal SDK operations and utilities    |

---

## 3. Default Retry Configuration (Add to "Configuration Options" section)

### Default Retry Behavior

All SDK methods automatically retry failed requests with the following default configuration:

```go
retry.Config{
    Strategy: "backoff",
    Backoff: &retry.BackoffStrategy{
        InitialInterval: 500,        // 500ms initial delay
        MaxInterval:     60000,      // 60 seconds maximum delay
        Exponent:        1.5,        // Exponential backoff multiplier
        MaxElapsedTime:  600000,     // 10 minutes total retry time
    },
    RetryConnectionErrors: true,
}
```

**Retry Status Codes:** `5XX`, `408` (Request Timeout), `429` (Too Many Requests)

**Behavior:**
- Connection errors are automatically retried
- Exponential backoff between retry attempts
- Maximum 10 minutes of total retry time
- Can be overridden per-request using `operations.WithRetries()`

---

## 4. Hooks System (NEW SECTION - Add after "Configuration Options")

## Hooks and Middleware

The SDK provides a hooks system for intercepting and modifying HTTP requests and responses.

### Hook Types

The SDK executes hooks at three points in the request lifecycle:

1. **`BeforeRequest`** - Called before sending the HTTP request
2. **`AfterSuccess`** - Called after receiving a successful response
3. **`AfterError`** - Called after receiving an error response

### Hook Context

Each hook receives a `hooks.HookContext` containing:

```go
type HookContext struct {
    SDK              *Mix           // SDK instance
    SDKConfiguration config.SDKConfiguration
    BaseURL          string
    Context          context.Context
    OperationID      string         // e.g., "createSession"
    OAuth2Scopes     []string
    SecuritySource   interface{}
}
```

### Hook Execution Order

For every API call:
```
1. BeforeRequest hook → 2. HTTP Request → 3. AfterSuccess/AfterError hook
```

**Note:** Hooks are primarily used internally by the SDK for request/response processing. Custom hook implementation requires access to internal SDK structures.

---

## 5. Authentication API - MISSING METHODS

Add these methods to the Authentication API section:

### `StartOAuthFlow()`

Initiate OAuth authentication flow for a specific provider.

```go
func (s *Authentication) StartOAuthFlow(
    ctx context.Context,
    provider string,
    opts ...operations.Option,
) (*operations.StartOAuthFlowResponse, error)
```

#### Parameters

| Parameter  | Type                   | Description                              |
| :--------- | :--------------------- | :--------------------------------------- |
| `ctx`      | `context.Context`      | Context for the request                  |
| `provider` | `string`               | Provider name (e.g., "google", "github") |
| `opts`     | `...operations.Option` | Additional options                       |

#### Returns

Returns `*operations.StartOAuthFlowResponse` containing the OAuth authorization URL or an error.

#### Example

```go
ctx := context.Background()
resp, err := client.Authentication.StartOAuthFlow(ctx, "google")
if err != nil {
    log.Fatalf("Failed to start OAuth flow: %v", err)
}

fmt.Printf("Visit this URL to authorize: %s\n", resp.AuthURL)
```

---

### `HandleOAuthCallback()`

Process OAuth callback and exchange authorization code for access token.

```go
func (s *Authentication) HandleOAuthCallback(
    ctx context.Context,
    request operations.HandleOAuthCallbackRequest,
    opts ...operations.Option,
) (*operations.HandleOAuthCallbackResponse, error)
```

#### Parameters

| Parameter | Type                                       | Description             |
| :-------- | :----------------------------------------- | :---------------------- |
| `ctx`     | `context.Context`                          | Context for the request |
| `request` | `operations.HandleOAuthCallbackRequest`    | Callback data           |
| `opts`    | `...operations.Option`                     | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.Authentication.HandleOAuthCallback(ctx,
    operations.HandleOAuthCallbackRequest{
        Code:     "auth_code_from_callback",
        State:    "state_token",
        Provider: "google",
    },
)
if err != nil {
    log.Fatalf("Failed to handle OAuth callback: %v", err)
}
```

---

### `ValidatePreferredProvider()`

Check if the user's preferred provider is authenticated.

```go
func (s *Authentication) ValidatePreferredProvider(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.ValidatePreferredProviderResponse, error)
```

#### Returns

Returns validation status indicating if the preferred provider has valid credentials.

#### Example

```go
ctx := context.Background()
resp, err := client.Authentication.ValidatePreferredProvider(ctx)
if err != nil {
    log.Fatalf("Failed to validate provider: %v", err)
}

if resp.IsValid {
    fmt.Println("Preferred provider is authenticated")
}
```

---

### `GetOAuthHealth()`

Get health status of all OAuth credentials.

```go
func (s *Authentication) GetOAuthHealth(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.GetOAuthHealthResponse, error)
```

#### Returns

Returns health status for each OAuth provider (healthy/degraded/unhealthy).

#### Example

```go
ctx := context.Background()
resp, err := client.Authentication.GetOAuthHealth(ctx)
if err != nil {
    log.Fatalf("Failed to get OAuth health: %v", err)
}

for provider, status := range resp.ProviderHealth {
    fmt.Printf("%s: %s\n", provider, status)
}
```

**Note:** This method is also available in the `Health` API resource.

---

### `RefreshOAuthTokens()`

Manually trigger OAuth token refresh for all expired tokens.

```go
func (s *Authentication) RefreshOAuthTokens(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.RefreshOAuthTokensResponse, error)
```

#### Returns

Returns status of token refresh operations.

#### Example

```go
ctx := context.Background()
resp, err := client.Authentication.RefreshOAuthTokens(ctx)
if err != nil {
    log.Fatalf("Failed to refresh tokens: %v", err)
}

fmt.Printf("Refreshed %d tokens\n", resp.RefreshedCount)
```

---

### `DeleteCredentials()` (CORRECT METHOD NAME)

**Note:** The correct method name is `DeleteCredentials()`, not `DeleteAPIKey()`.

Delete stored API key and/or OAuth credentials for a provider.

```go
func (s *Authentication) DeleteCredentials(
    ctx context.Context,
    provider string,
    opts ...operations.Option,
) (*operations.DeleteCredentialsResponse, error)
```

#### Parameters

| Parameter  | Type                   | Description             |
| :--------- | :--------------------- | :---------------------- |
| `ctx`      | `context.Context`      | Context for the request |
| `provider` | `string`               | Provider name to delete |
| `opts`     | `...operations.Option` | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.Authentication.DeleteCredentials(ctx, "anthropic")
if err != nil {
    log.Fatalf("Failed to delete credentials: %v", err)
}

fmt.Println("Credentials deleted successfully")
```

---

## 6. Files API - METHOD NAME CORRECTIONS

**IMPORTANT:** All Files API methods have different names than initially documented.

### Correct Method Names

Replace all Files API method signatures with these:

### `ListSessionFiles()`

**Correct method name** (not `ListFiles()`).

```go
func (s *Files) ListSessionFiles(
    ctx context.Context,
    sessionID string,
    opts ...operations.Option,
) (*operations.ListSessionFilesResponse, error)
```

---

### `UploadSessionFile()`

**Correct method name** (not `UploadFile()`).

```go
func (s *Files) UploadSessionFile(
    ctx context.Context,
    sessionID string,
    request operations.UploadSessionFileRequest,
    opts ...operations.Option,
) (*operations.UploadSessionFileResponse, error)
```

---

### `DeleteSessionFile()`

**Correct method name** (not `DeleteFile()`).

```go
func (s *Files) DeleteSessionFile(
    ctx context.Context,
    sessionID string,
    filename string,
    opts ...operations.Option,
) (*operations.DeleteSessionFileResponse, error)
```

---

### `GetSessionFile()`

**This single method handles both file download AND thumbnail generation.**

```go
func (s *Files) GetSessionFile(
    ctx context.Context,
    id string,
    filename string,
    thumb *string,
    time *float64,
    opts ...operations.Option,
) (*operations.GetSessionFileResponse, error)
```

#### Parameters

| Parameter  | Type                   | Description                                    |
| :--------- | :--------------------- | :--------------------------------------------- |
| `ctx`      | `context.Context`      | Context for the request                        |
| `id`       | `string`               | Session ID                                     |
| `filename` | `string`               | File name to retrieve                          |
| `thumb`    | `*string`              | Thumbnail type: `"box"`, `"width"`, `"height"` |
| `time`     | `*float64`             | Timestamp for video thumbnail (seconds)        |
| `opts`     | `...operations.Option` | Additional options                             |

#### Usage

**Download file:**
```go
resp, err := client.Files.GetSessionFile(ctx, sessionID, "document.pdf", nil, nil)
```

**Generate image thumbnail (box constraint):**
```go
thumb := "box"
resp, err := client.Files.GetSessionFile(ctx, sessionID, "image.jpg", &thumb, nil)
```

**Generate video thumbnail at 30 seconds:**
```go
thumb := "box"
timestamp := 30.0
resp, err := client.Files.GetSessionFile(ctx, sessionID, "video.mp4", &thumb, &timestamp)
```

---

## 7. Preferences API - MISSING METHOD

Add this method to the Preferences API section:

### `ResetPreferences()`

Reset user preferences to default values.

```go
func (s *Preferences) ResetPreferences(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.ResetPreferencesResponse, error)
```

#### Example

```go
ctx := context.Background()
resp, err := client.Preferences.ResetPreferences(ctx)
if err != nil {
    log.Fatalf("Failed to reset preferences: %v", err)
}

fmt.Println("Preferences reset to defaults")
```

---

### `UpdatePreferences()` (CORRECT METHOD NAME)

**Note:** The correct method name is `UpdatePreferences()`, not `SetPreferences()`.

Update preferences for model and provider selection.

```go
func (s *Preferences) UpdatePreferences(
    ctx context.Context,
    request operations.UpdatePreferencesRequest,
    opts ...operations.Option,
) (*operations.UpdatePreferencesResponse, error)
```

---

## 8. Tools API - METHOD NAME CORRECTIONS AND ADDITIONS

### `ListLLMTools()` (CORRECT METHOD NAME)

**Correct method name** (not `ListTools()`).

List all available LLM tools that Claude can invoke.

```go
func (s *Tools) ListLLMTools(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.ListLLMToolsResponse, error)
```

---

### `GetToolCredentialsStatus()` (NEW METHOD)

Get authentication/credential status for external tool integrations.

```go
func (s *Tools) GetToolCredentialsStatus(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.GetToolCredentialsStatusResponse, error)
```

#### Returns

Returns credential status for tools that require external authentication (e.g., Brave Search, web tools).

#### Example

```go
ctx := context.Background()
resp, err := client.Tools.GetToolCredentialsStatus(ctx)
if err != nil {
    log.Fatalf("Failed to get tool credentials: %v", err)
}

for tool, status := range resp.ToolCredentials {
    fmt.Printf("%s: %v\n", tool, status.IsAuthenticated)
}
```

---

## 9. System API - MISSING METHODS

Add these methods to the System API section:

### `GetCommand()`

Retrieve detailed information about a specific command.

```go
func (s *System) GetCommand(
    ctx context.Context,
    name string,
    opts ...operations.Option,
) (*operations.GetCommandResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `name`    | `string`               | Command name            |
| `opts`    | `...operations.Option` | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.System.GetCommand(ctx, "analyze-code")
if err != nil {
    log.Fatalf("Failed to get command: %v", err)
}

fmt.Printf("Command: %s\n", resp.Command.Name)
fmt.Printf("Description: %s\n", resp.Command.Description)
```

---

### `GetSystemInfo()`

Retrieve system information including storage configuration and paths.

```go
func (s *System) GetSystemInfo(
    ctx context.Context,
    opts ...operations.Option,
) (*operations.GetSystemInfoResponse, error)
```

#### Returns

Returns system configuration including storage directories, version information, and environment details.

#### Example

```go
ctx := context.Background()
resp, err := client.System.GetSystemInfo(ctx)
if err != nil {
    log.Fatalf("Failed to get system info: %v", err)
}

fmt.Printf("Storage path: %s\n", resp.SystemInfo.StoragePath)
fmt.Printf("Version: %s\n", resp.SystemInfo.Version)
```

---

### `ListMcpServers()` (CORRECT METHOD NAME)

**Note:** The correct casing is `ListMcpServers()`, not `ListMCPServers()`.

---

## 10. Permissions API - SIGNATURE CORRECTIONS

### Correct Method Signatures

The Permissions methods take the permission `id` directly as a parameter, not as part of a request object:

### `GrantPermission()`

```go
func (s *Permissions) GrantPermission(
    ctx context.Context,
    id string,
    opts ...operations.Option,
) (*operations.GrantPermissionResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `id`      | `string`               | Permission request ID   |
| `opts`    | `...operations.Option` | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.Permissions.GrantPermission(ctx, permissionID)
if err != nil {
    log.Fatalf("Failed to grant permission: %v", err)
}
```

---

### `DenyPermission()`

```go
func (s *Permissions) DenyPermission(
    ctx context.Context,
    id string,
    opts ...operations.Option,
) (*operations.DenyPermissionResponse, error)
```

#### Parameters

| Parameter | Type                   | Description             |
| :-------- | :--------------------- | :---------------------- |
| `ctx`     | `context.Context`      | Context for the request |
| `id`      | `string`               | Permission request ID   |
| `opts`    | `...operations.Option` | Additional options      |

#### Example

```go
ctx := context.Background()
resp, err := client.Permissions.DenyPermission(ctx, permissionID)
if err != nil {
    log.Fatalf("Failed to deny permission: %v", err)
}
```

---

## 11. Internal API Resource (NEW SECTION)

## Internal API

The `Internal` resource provides internal SDK operations and utilities.

**Note:** This is an internal API resource primarily used by the SDK itself. Direct usage is generally not required for typical applications.

Access via:
```go
client.Internal
```

The Internal API is initialized automatically when creating a new Mix client instance.

---

## 12. Request Options - ADD SUPPORTED OPTIONS CONSTANTS

Add this subsection to "Request Options":

### Supported Options Constants

The SDK defines constants for supported options on each operation:

```go
operations.SupportedOptionRetries
operations.SupportedOptionTimeout
```

These constants are used internally to validate which options are supported for each API method. All methods support at minimum:
- `Retries` configuration
- `Timeout` configuration

---

## 13. Complete OAuth Workflow Example (NEW EXAMPLE)

Add this to the "Advanced Examples" section:

### Complete OAuth Authentication Flow

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/recreate-run/mix-go-sdk"
    "github.com/recreate-run/mix-go-sdk/models/operations"
)

func main() {
    client := mix.New(
        mix.WithServerURL("http://localhost:8088"),
    )

    ctx := context.Background()

    // 1. Start OAuth flow
    flowResp, err := client.Authentication.StartOAuthFlow(ctx, "google")
    if err != nil {
        log.Fatalf("Failed to start OAuth flow: %v", err)
    }

    fmt.Printf("Visit this URL to authorize:\n%s\n\n", flowResp.AuthURL)

    // 2. User visits URL and authorizes, then is redirected to callback URL
    // Your callback handler receives the code and state parameters

    // 3. Handle the OAuth callback
    callbackResp, err := client.Authentication.HandleOAuthCallback(ctx,
        operations.HandleOAuthCallbackRequest{
            Code:     "received_auth_code",
            State:    "received_state",
            Provider: "google",
        },
    )
    if err != nil {
        log.Fatalf("Failed to handle callback: %v", err)
    }

    fmt.Println("OAuth authentication successful!")

    // 4. Validate the authentication
    validateResp, err := client.Authentication.ValidatePreferredProvider(ctx)
    if err != nil {
        log.Fatalf("Failed to validate: %v", err)
    }

    if validateResp.IsValid {
        fmt.Println("Provider is authenticated and ready to use")
    }

    // 5. Check OAuth health periodically
    healthResp, err := client.Authentication.GetOAuthHealth(ctx)
    if err != nil {
        log.Fatalf("Failed to get health: %v", err)
    }

    for provider, status := range healthResp.ProviderHealth {
        fmt.Printf("%s: %s\n", provider, status)
    }

    // 6. Refresh tokens if needed
    refreshResp, err := client.Authentication.RefreshOAuthTokens(ctx)
    if err != nil {
        log.Fatalf("Failed to refresh tokens: %v", err)
    }

    fmt.Printf("Refreshed %d tokens\n", refreshResp.RefreshedCount)
}
```

---

## 14. Health API Clarification

Add this note to the Health API section:

**Note:** The `GetOAuthHealth()` method is available in both the `Authentication` API and the `Health` API resources. Both methods provide identical functionality:

```go
// Both are equivalent
authHealth, _ := client.Authentication.GetOAuthHealth(ctx)
healthCheck, _ := client.Health.GetOAuthHealth(ctx)
```

---

## Summary of Critical Corrections

### Method Name Mismatches (CRITICAL)

1. **Files API:** All methods are `*SessionFile()` not just `*File()`
2. **Tools API:** `ListLLMTools()` not `ListTools()`
3. **Preferences API:** `UpdatePreferences()` not `SetPreferences()`
4. **Authentication API:** `DeleteCredentials()` not `DeleteAPIKey()`
5. **System API:** `ListMcpServers()` not `ListMCPServers()`

### Missing Methods (CRITICAL)

**Authentication API:**
- `StartOAuthFlow()`
- `HandleOAuthCallback()`
- `ValidatePreferredProvider()`
- `GetOAuthHealth()`
- `RefreshOAuthTokens()`

**Files API:**
- `GetSessionFile()` handles both download and thumbnails with `thumb` and `time` parameters

**Preferences API:**
- `ResetPreferences()`

**Tools API:**
- `GetToolCredentialsStatus()`

**System API:**
- `GetCommand()`
- `GetSystemInfo()`

### Missing Documentation Sections

1. Hooks system
2. Default retry configuration details
3. SDK metadata (version, user agent, server list)
4. Internal API resource
5. Supported options constants
