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
		plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
		plexgo.WithClientName("Plex Web"),
		plexgo.WithDeviceName("Linux"),
		plexgo.WithClientVersion("4.133.0"),
		plexgo.WithXPlexPlatform("Chrome"),
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