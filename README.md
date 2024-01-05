# github.com/LukeHagar/plexgo

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/LukeHagar/plexgo
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Server](docs/sdks/server/README.md)

* [GetServerCapabilities](docs/sdks/server/README.md#getservercapabilities) - Server Capabilities
* [GetServerPreferences](docs/sdks/server/README.md#getserverpreferences) - Get Server Preferences
* [GetAvailableClients](docs/sdks/server/README.md#getavailableclients) - Get Available Clients
* [GetDevices](docs/sdks/server/README.md#getdevices) - Get Devices
* [GetServerIdentity](docs/sdks/server/README.md#getserveridentity) - Get Server Identity
* [GetMyPlexAccount](docs/sdks/server/README.md#getmyplexaccount) - Get MyPlex Account
* [GetResizedPhoto](docs/sdks/server/README.md#getresizedphoto) - Get a Resized Photo
* [GetServerList](docs/sdks/server/README.md#getserverlist) - Get Server List

### [Media](docs/sdks/media/README.md)

* [MarkPlayed](docs/sdks/media/README.md#markplayed) - Mark Media Played
* [MarkUnplayed](docs/sdks/media/README.md#markunplayed) - Mark Media Unplayed
* [UpdatePlayProgress](docs/sdks/media/README.md#updateplayprogress) - Update Media Play Progress

### [Activities](docs/sdks/activities/README.md)

* [GetServerActivities](docs/sdks/activities/README.md#getserveractivities) - Get Server Activities
* [CancelServerActivities](docs/sdks/activities/README.md#cancelserveractivities) - Cancel Server Activities

### [Butler](docs/sdks/butler/README.md)

* [GetButlerTasks](docs/sdks/butler/README.md#getbutlertasks) - Get Butler tasks
* [StartAllTasks](docs/sdks/butler/README.md#startalltasks) - Start all Butler tasks
* [StopAllTasks](docs/sdks/butler/README.md#stopalltasks) - Stop all Butler tasks
* [StartTask](docs/sdks/butler/README.md#starttask) - Start a single Butler task
* [StopTask](docs/sdks/butler/README.md#stoptask) - Stop a single Butler task

### [Hubs](docs/sdks/hubs/README.md)

* [GetGlobalHubs](docs/sdks/hubs/README.md#getglobalhubs) - Get Global Hubs
* [GetLibraryHubs](docs/sdks/hubs/README.md#getlibraryhubs) - Get library specific hubs

### [Search](docs/sdks/search/README.md)

* [PerformSearch](docs/sdks/search/README.md#performsearch) - Perform a search
* [PerformVoiceSearch](docs/sdks/search/README.md#performvoicesearch) - Perform a voice search
* [GetSearchResults](docs/sdks/search/README.md#getsearchresults) - Get Search Results

### [Library](docs/sdks/library/README.md)

* [GetFileHash](docs/sdks/library/README.md#getfilehash) - Get Hash Value
* [GetRecentlyAdded](docs/sdks/library/README.md#getrecentlyadded) - Get Recently Added
* [GetLibraries](docs/sdks/library/README.md#getlibraries) - Get All Libraries
* [GetLibrary](docs/sdks/library/README.md#getlibrary) - Get Library Details
* [DeleteLibrary](docs/sdks/library/README.md#deletelibrary) - Delete Library Section
* [GetLibraryItems](docs/sdks/library/README.md#getlibraryitems) - Get Library Items
* [RefreshLibrary](docs/sdks/library/README.md#refreshlibrary) - Refresh Library
* [GetLatestLibraryItems](docs/sdks/library/README.md#getlatestlibraryitems) - Get Latest Library Items
* [GetCommonLibraryItems](docs/sdks/library/README.md#getcommonlibraryitems) - Get Common Library Items
* [GetMetadata](docs/sdks/library/README.md#getmetadata) - Get Items Metadata
* [GetMetadataChildren](docs/sdks/library/README.md#getmetadatachildren) - Get Items Children
* [GetOnDeck](docs/sdks/library/README.md#getondeck) - Get On Deck

### [Log](docs/sdks/log/README.md)

* [LogLine](docs/sdks/log/README.md#logline) - Logging a single line message.
* [LogMultiLine](docs/sdks/log/README.md#logmultiline) - Logging a multi-line message
* [EnablePaperTrail](docs/sdks/log/README.md#enablepapertrail) - Enabling Papertrail

### [Playlists](docs/sdks/playlists/README.md)

* [CreatePlaylist](docs/sdks/playlists/README.md#createplaylist) - Create a Playlist
* [GetPlaylists](docs/sdks/playlists/README.md#getplaylists) - Get All Playlists
* [GetPlaylist](docs/sdks/playlists/README.md#getplaylist) - Retrieve Playlist
* [DeletePlaylist](docs/sdks/playlists/README.md#deleteplaylist) - Deletes a Playlist
* [UpdatePlaylist](docs/sdks/playlists/README.md#updateplaylist) - Update a Playlist
* [GetPlaylistContents](docs/sdks/playlists/README.md#getplaylistcontents) - Retrieve Playlist Contents
* [ClearPlaylistContents](docs/sdks/playlists/README.md#clearplaylistcontents) - Delete Playlist Contents
* [AddPlaylistContents](docs/sdks/playlists/README.md#addplaylistcontents) - Adding to a Playlist
* [UploadPlaylist](docs/sdks/playlists/README.md#uploadplaylist) - Upload Playlist

### [Security](docs/sdks/security/README.md)

* [GetTransientToken](docs/sdks/security/README.md#gettransienttoken) - Get a Transient Token.
* [GetSourceConnectionInformation](docs/sdks/security/README.md#getsourceconnectioninformation) - Get Source Connection Information

### [Sessions](docs/sdks/sessions/README.md)

* [GetSessions](docs/sdks/sessions/README.md#getsessions) - Get Active Sessions
* [GetSessionHistory](docs/sdks/sessions/README.md#getsessionhistory) - Get Session History
* [GetTranscodeSessions](docs/sdks/sessions/README.md#gettranscodesessions) - Get Transcode Sessions
* [StopTranscodeSession](docs/sdks/sessions/README.md#stoptranscodesession) - Stop a Transcode Session

### [Updater](docs/sdks/updater/README.md)

* [GetUpdateStatus](docs/sdks/updater/README.md#getupdatestatus) - Querying status of updates
* [CheckForUpdates](docs/sdks/updater/README.md#checkforupdates) - Checking for updates
* [ApplyUpdates](docs/sdks/updater/README.md#applyupdates) - Apply Updates

### [Video](docs/sdks/video/README.md)

* [StartUniversalTranscode](docs/sdks/video/README.md#startuniversaltranscode) - Start Universal Transcode
* [GetTimeline](docs/sdks/video/README.md#gettimeline) - Get the timeline for a media item
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.GetServerCapabilitiesResponseBody | 401                                         | application/json                            |
| sdkerrors.SDKError                          | 4xx-5xx                                     | */*                                         |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/sdkerrors"
	"log"
)

func main() {
	s := plexgo.New(
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Server.GetServerCapabilities(ctx)
	if err != nil {

		var e *sdkerrors.GetServerCapabilitiesResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
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

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `http://10.10.10.47:32400` | None |
| 1 | `{protocol}://{ip}:{port}` | `protocol` (default is `http`), `ip` (default is `10.10.10.47`), `port` (default is `32400`) |

#### Example

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
		plexgo.WithServerIndex(1),
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
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

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithProtocol plexgo.ServerProtocol`
 * `WithIP string`
 * `WithPort string`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		plexgo.WithServerURL("http://10.10.10.47:32400"),
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
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
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name          | Type          | Scheme        |
| ------------- | ------------- | ------------- |
| `AccessToken` | apiKey        | API key       |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
