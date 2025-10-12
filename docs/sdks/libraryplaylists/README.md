# LibraryPlaylists
(*LibraryPlaylists*)

## Overview

Endpoints for manipulating playlists.

### Available Operations

* [CreatePlaylist](#createplaylist) - Create a Playlist
* [UploadPlaylist](#uploadplaylist) - Upload
* [DeletePlaylist](#deleteplaylist) - Delete a Playlist
* [UpdatePlaylist](#updateplaylist) - Editing a Playlist
* [GetPlaylistGenerators](#getplaylistgenerators) - Get a playlist's generators
* [ClearPlaylistItems](#clearplaylistitems) - Clearing a playlist
* [AddPlaylistItems](#addplaylistitems) - Adding to  a Playlist
* [DeletePlaylistItem](#deleteplaylistitem) - Delete a Generator
* [GetPlaylistGenerator](#getplaylistgenerator) - Get a playlist generator
* [ModifyPlaylistGenerator](#modifyplaylistgenerator) - Modify a Generator
* [GetPlaylistGeneratorItems](#getplaylistgeneratoritems) - Get a playlist generator's items
* [MovePlaylistItem](#moveplaylistitem) - Moving items in a playlist
* [RefreshPlaylist](#refreshplaylist) - Reprocess a generator

## CreatePlaylist

Create a new playlist. By default the playlist is blank.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPlaylist" method="post" path="/playlists" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.CreatePlaylist(ctx, operations.CreatePlaylistRequest{
        URI: plexgo.Pointer("https://short-term-disconnection.name/"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithPlaylistMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreatePlaylistRequest](../../models/operations/createplaylistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreatePlaylistResponse](../../models/operations/createplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UploadPlaylist

Imports m3u playlists by passing a path on the server to scan for m3u-formatted playlist files, or a path to a single playlist file.

### Example Usage

<!-- UsageSnippet language="go" operationID="uploadPlaylist" method="post" path="/playlists/upload" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.UploadPlaylist(ctx, operations.UploadPlaylistRequest{
        Path: plexgo.Pointer("/home/barkley/playlist.m3u"),
        Force: components.BoolIntOne.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UploadPlaylistRequest](../../models/operations/uploadplaylistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UploadPlaylistResponse](../../models/operations/uploadplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeletePlaylist

Deletes a playlist by provided id

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePlaylist" method="delete" path="/playlists/{playlistId}" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.DeletePlaylist(ctx, operations.DeletePlaylistRequest{
        PlaylistID: 343293,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeletePlaylistRequest](../../models/operations/deleteplaylistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeletePlaylistResponse](../../models/operations/deleteplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePlaylist

Edits a playlist in the same manner as [editing metadata](#tag/Provider/operation/metadataPutItem)

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePlaylist" method="put" path="/playlists/{playlistId}" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.UpdatePlaylist(ctx, operations.UpdatePlaylistRequest{
        PlaylistID: 157966,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdatePlaylistRequest](../../models/operations/updateplaylistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdatePlaylistResponse](../../models/operations/updateplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPlaylistGenerators

Get all the generators in a playlist

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlaylistGenerators" method="get" path="/playlists/{playlistId}/generators" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.GetPlaylistGenerators(ctx, operations.GetPlaylistGeneratorsRequest{
        PlaylistID: 162342,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPlaylistGeneratorsRequest](../../models/operations/getplaylistgeneratorsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetPlaylistGeneratorsResponse](../../models/operations/getplaylistgeneratorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ClearPlaylistItems

Clears a playlist, only works with dumb playlists. Returns the playlist.

### Example Usage

<!-- UsageSnippet language="go" operationID="clearPlaylistItems" method="delete" path="/playlists/{playlistId}/items" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.ClearPlaylistItems(ctx, operations.ClearPlaylistItemsRequest{
        PlaylistID: 552140,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithPlaylistMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ClearPlaylistItemsRequest](../../models/operations/clearplaylistitemsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ClearPlaylistItemsResponse](../../models/operations/clearplaylistitemsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddPlaylistItems

Adds a generator to a playlist, same parameters as the POST above. With a dumb playlist, this adds the specified items to the playlist. With a smart playlist, passing a new `uri` parameter replaces the rules for the playlist. Returns the playlist.

### Example Usage

<!-- UsageSnippet language="go" operationID="addPlaylistItems" method="put" path="/playlists/{playlistId}/items" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.AddPlaylistItems(ctx, operations.AddPlaylistItemsRequest{
        PlaylistID: 533723,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithPlaylistMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.AddPlaylistItemsRequest](../../models/operations/addplaylistitemsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.AddPlaylistItemsResponse](../../models/operations/addplaylistitemsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeletePlaylistItem

Deletes an item from a playlist. Only works with dumb playlists.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePlaylistItem" method="delete" path="/playlists/{playlistId}/items/{generatorId}" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.DeletePlaylistItem(ctx, operations.DeletePlaylistItemRequest{
        PlaylistID: 981646,
        GeneratorID: 194010,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithPlaylistMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeletePlaylistItemRequest](../../models/operations/deleteplaylistitemrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DeletePlaylistItemResponse](../../models/operations/deleteplaylistitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPlaylistGenerator

Get a playlist's generator.  Only used for optimized versions

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlaylistGenerator" method="get" path="/playlists/{playlistId}/items/{generatorId}" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.GetPlaylistGenerator(ctx, operations.GetPlaylistGeneratorRequest{
        PlaylistID: 744880,
        GeneratorID: 322168,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetPlaylistGeneratorRequest](../../models/operations/getplaylistgeneratorrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetPlaylistGeneratorResponse](../../models/operations/getplaylistgeneratorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ModifyPlaylistGenerator

Modify a playlist generator.  Only used for optimizer

### Example Usage

<!-- UsageSnippet language="go" operationID="modifyPlaylistGenerator" method="put" path="/playlists/{playlistId}/items/{generatorId}" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.ModifyPlaylistGenerator(ctx, operations.ModifyPlaylistGeneratorRequest{
        PlaylistID: 972007,
        GeneratorID: 208334,
        Item: &operations.Item{
            Location: &operations.ModifyPlaylistGeneratorQueryParamLocation{
                URI: plexgo.Pointer("library://82503060-0d68-4603-b594-8b071d54819e/item/%2Flibrary%2Fmetadata%2F146"),
            },
            LocationID: plexgo.Pointer[int64](-1),
            Policy: &operations.Policy{
                Value: plexgo.Pointer[int64](),
                Scope: operations.QueryParamScopeAll.ToPointer(),
                Unwatched: components.BoolIntZero.ToPointer(),
            },
            Target: plexgo.Pointer(""),
            TargetTagID: plexgo.Pointer[int64](1),
            Title: plexgo.Pointer("Jack-Jack Attack"),
            Type: plexgo.Pointer[int64](42),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithPlaylistMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ModifyPlaylistGeneratorRequest](../../models/operations/modifyplaylistgeneratorrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ModifyPlaylistGeneratorResponse](../../models/operations/modifyplaylistgeneratorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPlaylistGeneratorItems

Get a playlist generator's items

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlaylistGeneratorItems" method="get" path="/playlists/{playlistId}/items/{generatorId}/items" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.GetPlaylistGeneratorItems(ctx, operations.GetPlaylistGeneratorItemsRequest{
        PlaylistID: 77230,
        GeneratorID: 979714,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetPlaylistGeneratorItemsRequest](../../models/operations/getplaylistgeneratoritemsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetPlaylistGeneratorItemsResponse](../../models/operations/getplaylistgeneratoritemsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## MovePlaylistItem

Moves an item in a playlist. Only works with dumb playlists.

### Example Usage

<!-- UsageSnippet language="go" operationID="movePlaylistItem" method="put" path="/playlists/{playlistId}/items/{playlistItemId}/move" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.MovePlaylistItem(ctx, operations.MovePlaylistItemRequest{
        PlaylistID: 940298,
        PlaylistItemID: 375626,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithPlaylistMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.MovePlaylistItemRequest](../../models/operations/moveplaylistitemrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.MovePlaylistItemResponse](../../models/operations/moveplaylistitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RefreshPlaylist

Make a generator reprocess (refresh)

### Example Usage

<!-- UsageSnippet language="go" operationID="refreshPlaylist" method="put" path="/playlists/{playlistId}/items/{generatorId}/{metadataId}/{action}" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("abc123"),
        plexgo.WithProduct("Plex for Roku"),
        plexgo.WithVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithPlatformVersion("4.3 build 1057"),
        plexgo.WithDevice("Roku 3"),
        plexgo.WithModel("4200X"),
        plexgo.WithDeviceVendor("Roku"),
        plexgo.WithDeviceName("Living Room TV"),
        plexgo.WithMarketplace("googlePlay"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.LibraryPlaylists.RefreshPlaylist(ctx, operations.RefreshPlaylistRequest{
        PlaylistID: 895314,
        GeneratorID: 629742,
        MetadataID: 724422,
        Action: operations.ActionDisable,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RefreshPlaylistRequest](../../models/operations/refreshplaylistrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RefreshPlaylistResponse](../../models/operations/refreshplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |