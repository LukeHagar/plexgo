# Playlists
(*Playlists*)

## Overview

Playlists are ordered collections of media. They can be dumb (just a list of media) or smart (based on a media query, such as "all albums from 2017"). 
They can be organized in (optionally nesting) folders.
Retrieving a playlist, or its items, will trigger a refresh of its metadata. 
This may cause the duration and number of items to change.


### Available Operations

* [CreatePlaylist](#createplaylist) - Create a Playlist
* [GetPlaylists](#getplaylists) - Get All Playlists
* [GetPlaylist](#getplaylist) - Retrieve Playlist
* [DeletePlaylist](#deleteplaylist) - Deletes a Playlist
* [UpdatePlaylist](#updateplaylist) - Update a Playlist
* [GetPlaylistContents](#getplaylistcontents) - Retrieve Playlist Contents
* [ClearPlaylistContents](#clearplaylistcontents) - Delete Playlist Contents
* [AddPlaylistContents](#addplaylistcontents) - Adding to a Playlist
* [UploadPlaylist](#uploadplaylist) - Upload Playlist

## CreatePlaylist

Create a new playlist. By default the playlist is blank. To create a playlist along with a first item, pass:
- `uri` - The content URI for what we're playing (e.g. `server://1234/com.plexapp.plugins.library/library/metadata/1`).
- `playQueueID` - To create a playlist from an existing play queue.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("<value>"),
    )

    ctx := context.Background()
    res, err := s.Playlists.CreatePlaylist(ctx, operations.CreatePlaylistRequest{
        Title: "<value>",
        Type: operations.QueryParamTypePhoto,
        Smart: operations.SmartOne,
        URI: "https://inborn-brochure.biz",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreatePlaylistRequest](../../models/operations/createplaylistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreatePlaylistResponse](../../models/operations/createplaylistresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.CreatePlaylistResponseBody | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

## GetPlaylists

Get All Playlists given the specified filters.

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("<value>"),
    )


    var playlistType *operations.PlaylistType = operations.PlaylistTypeAudio.ToPointer()

    var smart *operations.QueryParamSmart = operations.QueryParamSmartZero.ToPointer()

    ctx := context.Background()
    res, err := s.Playlists.GetPlaylists(ctx, playlistType, smart)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `playlistType`                                                            | [*operations.PlaylistType](../../models/operations/playlisttype.md)       | :heavy_minus_sign:                                                        | limit to a type of playlist.                                              |
| `smart`                                                                   | [*operations.QueryParamSmart](../../models/operations/queryparamsmart.md) | :heavy_minus_sign:                                                        | type of playlists to return (default is all).                             |


### Response

**[*operations.GetPlaylistsResponse](../../models/operations/getplaylistsresponse.md), error**
| Error Object                       | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GetPlaylistsResponseBody | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4xx-5xx                            | */*                                |

## GetPlaylist

Gets detailed metadata for a playlist. A playlist for many purposes (rating, editing metadata, tagging), can be treated like a regular metadata item:
Smart playlist details contain the `content` attribute. This is the content URI for the generator. This can then be parsed by a client to provide smart playlist editing.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("<value>"),
    )


    var playlistID float64 = 4109.48

    ctx := context.Background()
    res, err := s.Playlists.GetPlaylist(ctx, playlistID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `playlistID`                                          | *float64*                                             | :heavy_check_mark:                                    | the ID of the playlist                                |


### Response

**[*operations.GetPlaylistResponse](../../models/operations/getplaylistresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetPlaylistResponseBody | 401                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## DeletePlaylist

This endpoint will delete a playlist


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("<value>"),
    )


    var playlistID float64 = 216.22

    ctx := context.Background()
    res, err := s.Playlists.DeletePlaylist(ctx, playlistID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `playlistID`                                          | *float64*                                             | :heavy_check_mark:                                    | the ID of the playlist                                |


### Response

**[*operations.DeletePlaylistResponse](../../models/operations/deleteplaylistresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.DeletePlaylistResponseBody | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

## UpdatePlaylist

From PMS version 1.9.1 clients can also edit playlist metadata using this endpoint as they would via `PUT /library/metadata/{playlistID}`


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("<value>"),
    )


    var playlistID float64 = 3915

    var title *string = plexgo.String("<value>")

    var summary *string = plexgo.String("<value>")

    ctx := context.Background()
    res, err := s.Playlists.UpdatePlaylist(ctx, playlistID, title, summary)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `playlistID`                                          | *float64*                                             | :heavy_check_mark:                                    | the ID of the playlist                                |
| `title`                                               | **string*                                             | :heavy_minus_sign:                                    | name of the playlist                                  |
| `summary`                                             | **string*                                             | :heavy_minus_sign:                                    | summary description of the playlist                   |


### Response

**[*operations.UpdatePlaylistResponse](../../models/operations/updateplaylistresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.UpdatePlaylistResponseBody | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

## GetPlaylistContents

Gets the contents of a playlist. Should be paged by clients via standard mechanisms. 
By default leaves are returned (e.g. episodes, movies). In order to return other types you can use the `type` parameter. 
For example, you could use this to display a list of recently added albums vis a smart playlist. 
Note that for dumb playlists, items have a `playlistItemID` attribute which is used for deleting or moving items.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("<value>"),
    )


    var playlistID float64 = 5004.46

    var type_ float64 = 9403.59

    ctx := context.Background()
    res, err := s.Playlists.GetPlaylistContents(ctx, playlistID, type_)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `playlistID`                                          | *float64*                                             | :heavy_check_mark:                                    | the ID of the playlist                                |
| `type_`                                               | *float64*                                             | :heavy_check_mark:                                    | the metadata type of the item to return               |


### Response

**[*operations.GetPlaylistContentsResponse](../../models/operations/getplaylistcontentsresponse.md), error**
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetPlaylistContentsResponseBody | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |

## ClearPlaylistContents

Clears a playlist, only works with dumb playlists. Returns the playlist.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("<value>"),
    )


    var playlistID float64 = 1893.18

    ctx := context.Background()
    res, err := s.Playlists.ClearPlaylistContents(ctx, playlistID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `playlistID`                                          | *float64*                                             | :heavy_check_mark:                                    | the ID of the playlist                                |


### Response

**[*operations.ClearPlaylistContentsResponse](../../models/operations/clearplaylistcontentsresponse.md), error**
| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.ClearPlaylistContentsResponseBody | 401                                         | application/json                            |
| sdkerrors.SDKError                          | 4xx-5xx                                     | */*                                         |

## AddPlaylistContents

Adds a generator to a playlist, same parameters as the POST to create. With a dumb playlist, this adds the specified items to the playlist.
With a smart playlist, passing a new `uri` parameter replaces the rules for the playlist. Returns the playlist.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("<value>"),
    )


    var playlistID float64 = 8502.01

    var uri string = "server://12345/com.plexapp.plugins.library/library/metadata/1"

    var playQueueID *float64 = plexgo.Float64(123)

    ctx := context.Background()
    res, err := s.Playlists.AddPlaylistContents(ctx, playlistID, uri, playQueueID)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                     | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ctx`                                                         | [context.Context](https://pkg.go.dev/context#Context)         | :heavy_check_mark:                                            | The context to use for the request.                           |                                                               |
| `playlistID`                                                  | *float64*                                                     | :heavy_check_mark:                                            | the ID of the playlist                                        |                                                               |
| `uri`                                                         | *string*                                                      | :heavy_check_mark:                                            | the content URI for the playlist                              | server://12345/com.plexapp.plugins.library/library/metadata/1 |
| `playQueueID`                                                 | **float64*                                                    | :heavy_minus_sign:                                            | the play queue to add to a playlist                           | 123                                                           |


### Response

**[*operations.AddPlaylistContentsResponse](../../models/operations/addplaylistcontentsresponse.md), error**
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.AddPlaylistContentsResponseBody | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |

## UploadPlaylist

Imports m3u playlists by passing a path on the server to scan for m3u-formatted playlist files, or a path to a single playlist file.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("<value>"),
    )


    var path string = "/home/barkley/playlist.m3u"

    var force operations.Force = operations.ForceZero

    ctx := context.Background()
    res, err := s.Playlists.UploadPlaylist(ctx, path, force)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| `path`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | absolute path to a directory on the server where m3u files are stored, or the absolute path to a playlist file on the server. <br/>If the `path` argument is a directory, that path will be scanned for playlist files to be processed. <br/>Each file in that directory creates a separate playlist, with a name based on the filename of the file that created it. <br/>The GUID of each playlist is based on the filename. <br/>If the `path` argument is a file, that file will be used to create a new playlist, with the name based on the filename of the file that created it. <br/>The GUID of each playlist is based on the filename.<br/> | /home/barkley/playlist.m3u                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `force`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [operations.Force](../../models/operations/force.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Force overwriting of duplicate playlists.  <br/>By default, a playlist file uploaded with the same path will overwrite the existing playlist. <br/>The `force` argument is used to disable overwriting.  <br/>If the `force` argument is set to 0, a new playlist will be created suffixed with the date and time that the duplicate was uploaded.<br/>                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |


### Response

**[*operations.UploadPlaylistResponse](../../models/operations/uploadplaylistresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.UploadPlaylistResponseBody | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |
