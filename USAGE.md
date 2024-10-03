<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"log"
)

func main() {
	s := plexgo.New(
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
		plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
		plexgo.WithClientName("Plex for Roku"),
		plexgo.WithClientVersion("2.4.1"),
		plexgo.WithPlatform("Roku"),
		plexgo.WithDeviceNickname("Roku 3"),
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