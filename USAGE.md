<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"log"
)

func main() {
	s := plexgo.New(
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
		plexgo.WithXPlexClientIdentifier("<value>"),
	)

	ctx := context.Background()
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