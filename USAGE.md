<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Server.GetServerCapabilities(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->