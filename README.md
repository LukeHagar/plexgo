# github.com/LukeHagar/plexgo

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

<!-- Start Summary [summary] -->
## Summary

Plex-API: An Open API Spec for interacting with Plex.tv and Plex Media Server

# Plex Media Server OpenAPI Specification

An Open Source OpenAPI Specification for Plex Media Server

Automation and SDKs provided by [Speakeasy](https://speakeasyapi.dev/)

## Documentation

[API Documentation](https://plexapi.dev)

## SDKs

The following SDKs are generated from the OpenAPI Specification. They are automatically generated and may not be fully tested. If you find any issues, please open an issue on the [main specification Repository](https://github.com/LukeHagar/plex-api-spec).

| Language              | Repository                                        | Releases                                                                                         | Other                                                   |
| --------------------- | ------------------------------------------------- | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------- |
| Python                | [GitHub](https://github.com/LukeHagar/plexpy)     | [PyPI](https://pypi.org/project/plex-api-client/)                                                | -                                                       |
| JavaScript/TypeScript | [GitHub](https://github.com/LukeHagar/plexjs)     | [NPM](https://www.npmjs.com/package/@lukehagar/plexjs) \ [JSR](https://jsr.io/@lukehagar/plexjs) | -                                                       |
| Go                    | [GitHub](https://github.com/LukeHagar/plexgo)     | [Releases](https://github.com/LukeHagar/plexgo/releases)                                         | [GoDoc](https://pkg.go.dev/github.com/LukeHagar/plexgo) |
| Ruby                  | [GitHub](https://github.com/LukeHagar/plexruby)   | [Releases](https://github.com/LukeHagar/plexruby/releases)                                       | -                                                       |
| Swift                 | [GitHub](https://github.com/LukeHagar/plexswift)  | [Releases](https://github.com/LukeHagar/plexswift/releases)                                      | -                                                       |
| PHP                   | [GitHub](https://github.com/LukeHagar/plexphp)    | [Releases](https://github.com/LukeHagar/plexphp/releases)                                        | -                                                       |
| Java                  | [GitHub](https://github.com/LukeHagar/plexjava)   | [Releases](https://github.com/LukeHagar/plexjava/releases)                                       | -                                                       |
| C#                    | [GitHub](https://github.com/LukeHagar/plexcsharp) | [Releases](https://github.com/LukeHagar/plexcsharp/releases)                                     | -
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/LukeHagar/plexgo](#githubcomlukehagarplexgo)
* [Plex Media Server OpenAPI Specification](#plex-media-server-openapi-specification)
  * [Documentation](#documentation)
  * [SDKs](#sdks)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Activities](docs/sdks/activities/README.md)

* [GetServerActivities](docs/sdks/activities/README.md#getserveractivities) - Get Server Activities
* [CancelServerActivities](docs/sdks/activities/README.md#cancelserveractivities) - Cancel Server Activities

### [Authentication](docs/sdks/authentication/README.md)

* [GetTransientToken](docs/sdks/authentication/README.md#gettransienttoken) - Get a Transient Token
* [GetSourceConnectionInformation](docs/sdks/authentication/README.md#getsourceconnectioninformation) - Get Source Connection Information
* [GetTokenDetails](docs/sdks/authentication/README.md#gettokendetails) - Get Token Details
* [PostUsersSignInData](docs/sdks/authentication/README.md#postuserssignindata) - Get User Sign In Data

### [Butler](docs/sdks/butler/README.md)

* [GetButlerTasks](docs/sdks/butler/README.md#getbutlertasks) - Get Butler tasks
* [StartAllTasks](docs/sdks/butler/README.md#startalltasks) - Start all Butler tasks
* [StopAllTasks](docs/sdks/butler/README.md#stopalltasks) - Stop all Butler tasks
* [StartTask](docs/sdks/butler/README.md#starttask) - Start a single Butler task
* [StopTask](docs/sdks/butler/README.md#stoptask) - Stop a single Butler task

### [Hubs](docs/sdks/hubs/README.md)

* [GetGlobalHubs](docs/sdks/hubs/README.md#getglobalhubs) - Get Global Hubs
* [GetRecentlyAdded](docs/sdks/hubs/README.md#getrecentlyadded) - Get Recently Added
* [GetLibraryHubs](docs/sdks/hubs/README.md#getlibraryhubs) - Get library specific hubs

### [Library](docs/sdks/library/README.md)

* [GetFileHash](docs/sdks/library/README.md#getfilehash) - Get Hash Value
* [GetRecentlyAddedLibrary](docs/sdks/library/README.md#getrecentlyaddedlibrary) - Get Recently Added
* [GetAllLibraries](docs/sdks/library/README.md#getalllibraries) - Get All Libraries
* [GetLibraryDetails](docs/sdks/library/README.md#getlibrarydetails) - Get Library Details
* [DeleteLibrary](docs/sdks/library/README.md#deletelibrary) - Delete Library Section
* [GetLibraryItems](docs/sdks/library/README.md#getlibraryitems) - Get Library Items
* [GetLibrarySectionsAll](docs/sdks/library/README.md#getlibrarysectionsall) - Get Library section media by tag ALL
* [GetRefreshLibraryMetadata](docs/sdks/library/README.md#getrefreshlibrarymetadata) - Refresh Metadata Of The Library
* [GetSearchLibrary](docs/sdks/library/README.md#getsearchlibrary) - Search Library
* [GetGenresLibrary](docs/sdks/library/README.md#getgenreslibrary) - Get Genres of library media
* [GetCountriesLibrary](docs/sdks/library/README.md#getcountrieslibrary) - Get Countries of library media
* [GetActorsLibrary](docs/sdks/library/README.md#getactorslibrary) - Get Actors of library media
* [GetSearchAllLibraries](docs/sdks/library/README.md#getsearchalllibraries) - Search All Libraries
* [GetMediaMetaData](docs/sdks/library/README.md#getmediametadata) - Get Media Metadata
* [GetMediaArts](docs/sdks/library/README.md#getmediaarts) - Get Media Background Artwork
* [PostMediaArts](docs/sdks/library/README.md#postmediaarts) - Upload Media Background Artwork
* [GetMediaPosters](docs/sdks/library/README.md#getmediaposters) - Get Media Posters
* [PostMediaPoster](docs/sdks/library/README.md#postmediaposter) - Upload Media Poster
* [GetMetadataChildren](docs/sdks/library/README.md#getmetadatachildren) - Get Items Children
* [GetTopWatchedContent](docs/sdks/library/README.md#gettopwatchedcontent) - Get Top Watched Content

### [Log](docs/sdks/log/README.md)

* [LogLine](docs/sdks/log/README.md#logline) - Logging a single line message.
* [LogMultiLine](docs/sdks/log/README.md#logmultiline) - Logging a multi-line message
* [EnablePaperTrail](docs/sdks/log/README.md#enablepapertrail) - Enabling Papertrail

### [Media](docs/sdks/media/README.md)

* [MarkPlayed](docs/sdks/media/README.md#markplayed) - Mark Media Played
* [MarkUnplayed](docs/sdks/media/README.md#markunplayed) - Mark Media Unplayed
* [UpdatePlayProgress](docs/sdks/media/README.md#updateplayprogress) - Update Media Play Progress
* [GetBannerImage](docs/sdks/media/README.md#getbannerimage) - Get Banner Image
* [GetThumbImage](docs/sdks/media/README.md#getthumbimage) - Get Thumb Image

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

### [Plex](docs/sdks/plex/README.md)

* [GetCompanionsData](docs/sdks/plex/README.md#getcompanionsdata) - Get Companions Data
* [GetUserFriends](docs/sdks/plex/README.md#getuserfriends) - Get list of friends of the user logged in
* [GetGeoData](docs/sdks/plex/README.md#getgeodata) - Get Geo Data
* [GetHomeData](docs/sdks/plex/README.md#gethomedata) - Get Plex Home Data
* [GetServerResources](docs/sdks/plex/README.md#getserverresources) - Get Server Resources
* [GetPin](docs/sdks/plex/README.md#getpin) - Get a Pin
* [GetTokenByPinID](docs/sdks/plex/README.md#gettokenbypinid) - Get Access Token by PinId


### [Search](docs/sdks/search/README.md)

* [PerformSearch](docs/sdks/search/README.md#performsearch) - Perform a search
* [PerformVoiceSearch](docs/sdks/search/README.md#performvoicesearch) - Perform a voice search
* [GetSearchResults](docs/sdks/search/README.md#getsearchresults) - Get Search Results

### [Server](docs/sdks/server/README.md)

* [GetServerCapabilities](docs/sdks/server/README.md#getservercapabilities) - Get Server Capabilities
* [GetServerPreferences](docs/sdks/server/README.md#getserverpreferences) - Get Server Preferences
* [GetAvailableClients](docs/sdks/server/README.md#getavailableclients) - Get Available Clients
* [GetDevices](docs/sdks/server/README.md#getdevices) - Get Devices
* [GetServerIdentity](docs/sdks/server/README.md#getserveridentity) - Get Server Identity
* [GetMyPlexAccount](docs/sdks/server/README.md#getmyplexaccount) - Get MyPlex Account
* [GetResizedPhoto](docs/sdks/server/README.md#getresizedphoto) - Get a Resized Photo
* [GetMediaProviders](docs/sdks/server/README.md#getmediaproviders) - Get Media Providers
* [GetServerList](docs/sdks/server/README.md#getserverlist) - Get Server List

### [Sessions](docs/sdks/sessions/README.md)

* [GetSessions](docs/sdks/sessions/README.md#getsessions) - Get Active Sessions
* [GetSessionHistory](docs/sdks/sessions/README.md#getsessionhistory) - Get Session History
* [GetTranscodeSessions](docs/sdks/sessions/README.md#gettranscodesessions) - Get Transcode Sessions
* [StopTranscodeSession](docs/sdks/sessions/README.md#stoptranscodesession) - Stop a Transcode Session

### [Statistics](docs/sdks/statistics/README.md)

* [GetStatistics](docs/sdks/statistics/README.md#getstatistics) - Get Media Statistics
* [GetResourcesStatistics](docs/sdks/statistics/README.md#getresourcesstatistics) - Get Resources Statistics
* [GetBandwidthStatistics](docs/sdks/statistics/README.md#getbandwidthstatistics) - Get Bandwidth Statistics

### [Updater](docs/sdks/updater/README.md)

* [GetUpdateStatus](docs/sdks/updater/README.md#getupdatestatus) - Querying status of updates
* [CheckForUpdates](docs/sdks/updater/README.md#checkforupdates) - Checking for updates
* [ApplyUpdates](docs/sdks/updater/README.md#applyupdates) - Apply Updates

### [Users](docs/sdks/users/README.md)

* [GetUsers](docs/sdks/users/README.md#getusers) - Get list of all connected users

### [Video](docs/sdks/video/README.md)

* [GetTimeline](docs/sdks/video/README.md#gettimeline) - Get the timeline for a media item
* [StartUniversalTranscode](docs/sdks/video/README.md#startuniversaltranscode) - Start Universal Transcode

### [Watchlist](docs/sdks/watchlist/README.md)

* [GetWatchList](docs/sdks/watchlist/README.md#getwatchlist) - Get User Watchlist

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Server.GetServerCapabilities(ctx, operations.WithRetries(
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
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithRetryConfig(
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetServerCapabilities` function may return the following errors:

| Error Type                                  | Status Code | Content Type     |
| ------------------------------------------- | ----------- | ---------------- |
| sdkerrors.GetServerCapabilitiesBadRequest   | 400         | application/json |
| sdkerrors.GetServerCapabilitiesUnauthorized | 401         | application/json |
| sdkerrors.SDKError                          | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Server.GetServerCapabilities(ctx)
	if err != nil {

		var e *sdkerrors.GetServerCapabilitiesBadRequest
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.GetServerCapabilitiesUnauthorized
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

### Server Variables

The default server `{protocol}://{ip}:{port}` contains variables and is set to `https://10.10.10.47:32400` by default. To override default values, the following options are available when initializing the SDK client instance:

| Variable   | Option                                  | Supported Values           | Default         | Description                                    |
| ---------- | --------------------------------------- | -------------------------- | --------------- | ---------------------------------------------- |
| `protocol` | `WithProtocol(protocol ServerProtocol)` | - `"http"`<br/>- `"https"` | `"https"`       | The protocol to use for the server connection  |
| `ip`       | `WithIP(ip string)`                     | string                     | `"10.10.10.47"` | The IP address or hostname of your Plex Server |
| `port`     | `WithPort(port string)`                 | string                     | `"32400"`       | The port of your Plex Server                   |

#### Example

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
		plexgo.WithProtocol("https"),
		plexgo.WithIP("4982:bc2a:b4f8:efb5:2394:5bc3:ab4f:0e6d"),
		plexgo.WithPort("44765"),
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

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
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
		plexgo.WithServerURL("https://10.10.10.47:32400"),
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

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Plex.GetCompanionsData(ctx, operations.WithServerURL("https://plex.tv/api/v2"))
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
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

| Name          | Type   | Scheme  |
| ------------- | ------ | ------- |
| `AccessToken` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
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
