<!-- Start SDK Example Usage [usage] -->
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