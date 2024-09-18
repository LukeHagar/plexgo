# Library
(*Library*)

## Overview

API Calls interacting with Plex Media Server Libraries


### Available Operations

* [GetFileHash](#getfilehash) - Get Hash Value
* [GetRecentlyAdded](#getrecentlyadded) - Get Recently Added
* [GetAllLibraries](#getalllibraries) - Get All Libraries
* [GetLibraryDetails](#getlibrarydetails) - Get Library Details
* [DeleteLibrary](#deletelibrary) - Delete Library Section
* [GetLibraryItems](#getlibraryitems) - Get Library Items
* [GetRefreshLibraryMetadata](#getrefreshlibrarymetadata) - Refresh Metadata Of The Library
* [GetSearchLibrary](#getsearchlibrary) - Search Library
* [GetMetaDataByRatingKey](#getmetadatabyratingkey) - Get Metadata by RatingKey
* [GetMetadataChildren](#getmetadatachildren) - Get Items Children
* [GetTopWatchedContent](#gettopwatchedcontent) - Get Top Watched Content
* [GetOnDeck](#getondeck) - Get On Deck

## GetFileHash

This resource returns hash values for local files

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetFileHash(ctx, "file://C:\Image.png&type=13", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |                                                                   |
| `url_`                                                            | *string*                                                          | :heavy_check_mark:                                                | This is the path to the local file, must be prefixed by `file://` | file://C:\Image.png&type=13                                       |
| `type_`                                                           | **float64*                                                        | :heavy_minus_sign:                                                | Item type                                                         |                                                                   |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |                                                                   |

### Response

**[*operations.GetFileHashResponse](../../models/operations/getfilehashresponse.md), error**

### Errors

| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetFileHashBadRequest   | 400                               | application/json                  |
| sdkerrors.GetFileHashUnauthorized | 401                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |


## GetRecentlyAdded

This endpoint will return the recently added content.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetRecentlyAdded(ctx, plexgo.Int(0), plexgo.Int(50))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                 | Type                                                                                                                                                                                      | Required                                                                                                                                                                                  | Description                                                                                                                                                                               | Example                                                                                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                        | The context to use for the request.                                                                                                                                                       |                                                                                                                                                                                           |
| `xPlexContainerStart`                                                                                                                                                                     | **int*                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                        | The index of the first item to return. If not specified, the first item will be returned.<br/>If the number of items exceeds the limit, the response will be paginated.<br/>By default this is 0<br/> | 0                                                                                                                                                                                         |
| `xPlexContainerSize`                                                                                                                                                                      | **int*                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                        | The number of items to return. If not specified, all items will be returned.<br/>If the number of items exceeds the limit, the response will be paginated.<br/>By default this is 50<br/> | 50                                                                                                                                                                                        |
| `opts`                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                        | The options for this request.                                                                                                                                                             |                                                                                                                                                                                           |

### Response

**[*operations.GetRecentlyAddedResponse](../../models/operations/getrecentlyaddedresponse.md), error**

### Errors

| Error Object                           | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.GetRecentlyAddedBadRequest   | 400                                    | application/json                       |
| sdkerrors.GetRecentlyAddedUnauthorized | 401                                    | application/json                       |
| sdkerrors.SDKError                     | 4xx-5xx                                | */*                                    |


## GetAllLibraries

A library section (commonly referred to as just a library) is a collection of media. 
Libraries are typed, and depending on their type provide either a flat or a hierarchical view of the media. 
For example, a music library has an artist > albums > tracks structure, whereas a movie library is flat.

Libraries have features beyond just being a collection of media; for starters, they include information about supported types, filters and sorts. 
This allows a client to provide a rich interface around the media (e.g. allow sorting movies by release year).


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetAllLibraries(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAllLibrariesResponse](../../models/operations/getalllibrariesresponse.md), error**

### Errors

| Error Object                          | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.GetAllLibrariesBadRequest   | 400                                   | application/json                      |
| sdkerrors.GetAllLibrariesUnauthorized | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4xx-5xx                               | */*                                   |


## GetLibraryDetails

## Library Details Endpoint

This endpoint provides comprehensive details about the library, focusing on organizational aspects rather than the content itself.   

The details include:

### Directories
Organized into three categories:

- **Primary Directories**: 
  - Used in some clients for quick access to media subsets (e.g., "All", "On Deck").
  - Most can be replicated via media queries.
  - Customizable by users.

- **Secondary Directories**:
  - Marked with `secondary="1"`.
  - Used in older clients for structured navigation.

- **Special Directories**:
  - Includes a "By Folder" entry for filesystem-based browsing.
  - Contains an obsolete `search="1"` entry for on-the-fly search dialog creation.

### Types
Each type in the library comes with a set of filters and sorts, aiding in building dynamic media controls:

- **Type Object Attributes**:
  - `key`: Endpoint for the media list of this type.
  - `type`: Metadata type (if standard Plex type).
  - `title`: Title for this content type (e.g., "Movies").

- **Filter Objects**:
  - Subset of the media query language.
  - Attributes include `filter` (name), `filterType` (data type), `key` (endpoint for value range), and `title`.

- **Sort Objects**:
  - Description of sort fields.
  - Attributes include `defaultDirection` (asc/desc), `descKey` and `key` (sort parameters), and `title`.

> **Note**: Filters and sorts are optional; without them, no filtering controls are rendered.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetLibraryDetails(ctx, 9518, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                  | Type                                                                                                                                                                                       | Required                                                                                                                                                                                   | Description                                                                                                                                                                                | Example                                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                         | The context to use for the request.                                                                                                                                                        |                                                                                                                                                                                            |
| `sectionKey`                                                                                                                                                                               | *int*                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                         | The unique key of the Plex library. <br/>Note: This is unique in the context of the Plex server.<br/>                                                                                      | 9518                                                                                                                                                                                       |
| `includeDetails`                                                                                                                                                                           | [*operations.IncludeDetails](../../models/operations/includedetails.md)                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                         | Whether or not to include details for a section (types, filters, and sorts). <br/>Only exists for backwards compatibility, media providers other than the server libraries have it on always.<br/> |                                                                                                                                                                                            |
| `opts`                                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                         | The options for this request.                                                                                                                                                              |                                                                                                                                                                                            |

### Response

**[*operations.GetLibraryDetailsResponse](../../models/operations/getlibrarydetailsresponse.md), error**

### Errors

| Error Object                            | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.GetLibraryDetailsBadRequest   | 400                                     | application/json                        |
| sdkerrors.GetLibraryDetailsUnauthorized | 401                                     | application/json                        |
| sdkerrors.SDKError                      | 4xx-5xx                                 | */*                                     |


## DeleteLibrary

Delete a library using a specific section id

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.DeleteLibrary(ctx, 9518)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |                                                                                               |
| `sectionKey`                                                                                  | *int*                                                                                         | :heavy_check_mark:                                                                            | The unique key of the Plex library. <br/>Note: This is unique in the context of the Plex server.<br/> | 9518                                                                                          |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |                                                                                               |

### Response

**[*operations.DeleteLibraryResponse](../../models/operations/deletelibraryresponse.md), error**

### Errors

| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.DeleteLibraryBadRequest   | 400                                 | application/json                    |
| sdkerrors.DeleteLibraryUnauthorized | 401                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |


## GetLibraryItems

Fetches details from a specific section of the library identified by a section key and a tag. The tag parameter accepts the following values:
- `all`: All items in the section.
- `unwatched`: Items that have not been played.
- `newest`: Items that are recently released.
- `recentlyAdded`: Items that are recently added to the library.
- `recentlyViewed`: Items that were recently viewed.
- `onDeck`: Items to continue watching.
- `collection`: Items categorized by collection.
- `edition`: Items categorized by edition.
- `genre`: Items categorized by genre.
- `year`: Items categorized by year of release.
- `decade`: Items categorized by decade.
- `director`: Items categorized by director.
- `actor`: Items categorized by starring actor.
- `country`: Items categorized by country of origin.
- `contentRating`: Items categorized by content rating.
- `rating`: Items categorized by rating.
- `resolution`: Items categorized by resolution.
- `firstCharacter`: Items categorized by the first letter.
- `folder`: Items categorized by folder.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetLibraryItems(ctx, operations.GetLibraryItemsRequest{
        SectionKey: 9518,
        Tag: operations.TagEdition,
        IncludeGuids: operations.IncludeGuidsOne.ToPointer(),
        IncludeMeta: operations.IncludeMetaOne.ToPointer(),
        Type: operations.TypeTwo,
        XPlexContainerStart: plexgo.Int(0),
        XPlexContainerSize: plexgo.Int(50),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLibraryItemsRequest](../../models/operations/getlibraryitemsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetLibraryItemsResponse](../../models/operations/getlibraryitemsresponse.md), error**

### Errors

| Error Object                          | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.GetLibraryItemsBadRequest   | 400                                   | application/json                      |
| sdkerrors.GetLibraryItemsUnauthorized | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4xx-5xx                               | */*                                   |


## GetRefreshLibraryMetadata

This endpoint Refreshes all the Metadata of the library.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetRefreshLibraryMetadata(ctx, 9518, operations.ForceOne.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |                                                                                               |
| `sectionKey`                                                                                  | *int*                                                                                         | :heavy_check_mark:                                                                            | The unique key of the Plex library. <br/>Note: This is unique in the context of the Plex server.<br/> | 9518                                                                                          |
| `force`                                                                                       | [*operations.Force](../../models/operations/force.md)                                         | :heavy_minus_sign:                                                                            | Force the refresh even if the library is already being refreshed.                             | 0                                                                                             |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |                                                                                               |

### Response

**[*operations.GetRefreshLibraryMetadataResponse](../../models/operations/getrefreshlibrarymetadataresponse.md), error**

### Errors

| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.GetRefreshLibraryMetadataBadRequest   | 400                                             | application/json                                |
| sdkerrors.GetRefreshLibraryMetadataUnauthorized | 401                                             | application/json                                |
| sdkerrors.SDKError                              | 4xx-5xx                                         | */*                                             |


## GetSearchLibrary

Search for content within a specific section of the library.

### Types
Each type in the library comes with a set of filters and sorts, aiding in building dynamic media controls:

- **Type Object Attributes**:
  - `type`: Metadata type (if standard Plex type).  
  - `title`: Title for this content type (e.g., "Movies").

- **Filter Objects**:
  - Subset of the media query language.
  - Attributes include `filter` (name), `filterType` (data type), `key` (endpoint for value range), and `title`.

- **Sort Objects**:
  - Description of sort fields.
  - Attributes include `defaultDirection` (asc/desc), `descKey` and `key` (sort parameters), and `title`.

> **Note**: Filters and sorts are optional; without them, no filtering controls are rendered.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetSearchLibrary(ctx, 9518, operations.QueryParamTypeTwo)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                       | Type                                                                                                                                                                            | Required                                                                                                                                                                        | Description                                                                                                                                                                     | Example                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                           | :heavy_check_mark:                                                                                                                                                              | The context to use for the request.                                                                                                                                             |                                                                                                                                                                                 |
| `sectionKey`                                                                                                                                                                    | *int*                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                              | The unique key of the Plex library. <br/>Note: This is unique in the context of the Plex server.<br/>                                                                           | 9518                                                                                                                                                                            |
| `type_`                                                                                                                                                                         | [operations.QueryParamType](../../models/operations/queryparamtype.md)                                                                                                          | :heavy_check_mark:                                                                                                                                                              | The type of media to retrieve.<br/>1 = movie<br/>2 = show<br/>3 = season<br/>4 = episode<br/>E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries<br/> | 2                                                                                                                                                                               |
| `opts`                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                        | :heavy_minus_sign:                                                                                                                                                              | The options for this request.                                                                                                                                                   |                                                                                                                                                                                 |

### Response

**[*operations.GetSearchLibraryResponse](../../models/operations/getsearchlibraryresponse.md), error**

### Errors

| Error Object                           | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.GetSearchLibraryBadRequest   | 400                                    | application/json                       |
| sdkerrors.GetSearchLibraryUnauthorized | 401                                    | application/json                       |
| sdkerrors.SDKError                     | 4xx-5xx                                | */*                                    |


## GetMetaDataByRatingKey

This endpoint will return the metadata of a library item specified with the ratingKey.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetMetaDataByRatingKey(ctx, 9518)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `ratingKey`                                              | *int64*                                                  | :heavy_check_mark:                                       | the id of the library item to return the children of.    | 9518                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMetaDataByRatingKeyResponse](../../models/operations/getmetadatabyratingkeyresponse.md), error**

### Errors

| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.GetMetaDataByRatingKeyBadRequest   | 400                                          | application/json                             |
| sdkerrors.GetMetaDataByRatingKeyUnauthorized | 401                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |


## GetMetadataChildren

This endpoint will return the children of of a library item specified with the ratingKey.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetMetadataChildren(ctx, 1539.14, plexgo.String("Stream"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `ratingKey`                                                             | *float64*                                                               | :heavy_check_mark:                                                      | the id of the library item to return the children of.                   |
| `includeElements`                                                       | **string*                                                               | :heavy_minus_sign:                                                      | Adds additional elements to the response. Supported types are (Stream)<br/> |
| `opts`                                                                  | [][operations.Option](../../models/operations/option.md)                | :heavy_minus_sign:                                                      | The options for this request.                                           |

### Response

**[*operations.GetMetadataChildrenResponse](../../models/operations/getmetadatachildrenresponse.md), error**

### Errors

| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetMetadataChildrenBadRequest   | 400                                       | application/json                          |
| sdkerrors.GetMetadataChildrenUnauthorized | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |


## GetTopWatchedContent

This endpoint will return the top watched content from libraries of a certain type


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetTopWatchedContent(ctx, operations.GetTopWatchedContentQueryParamTypeTwo, plexgo.Int64(1))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                       | Type                                                                                                                                                                            | Required                                                                                                                                                                        | Description                                                                                                                                                                     | Example                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                           | :heavy_check_mark:                                                                                                                                                              | The context to use for the request.                                                                                                                                             |                                                                                                                                                                                 |
| `type_`                                                                                                                                                                         | [operations.GetTopWatchedContentQueryParamType](../../models/operations/gettopwatchedcontentqueryparamtype.md)                                                                  | :heavy_check_mark:                                                                                                                                                              | The type of media to retrieve.<br/>1 = movie<br/>2 = show<br/>3 = season<br/>4 = episode<br/>E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries<br/> | 2                                                                                                                                                                               |
| `includeGuids`                                                                                                                                                                  | **int64*                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                              | Adds the Guids object to the response<br/>                                                                                                                                      | 1                                                                                                                                                                               |
| `opts`                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                        | :heavy_minus_sign:                                                                                                                                                              | The options for this request.                                                                                                                                                   |                                                                                                                                                                                 |

### Response

**[*operations.GetTopWatchedContentResponse](../../models/operations/gettopwatchedcontentresponse.md), error**

### Errors

| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.GetTopWatchedContentBadRequest   | 400                                        | application/json                           |
| sdkerrors.GetTopWatchedContentUnauthorized | 401                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |


## GetOnDeck

This endpoint will return the on deck content.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Library.GetOnDeck(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOnDeckResponse](../../models/operations/getondeckresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.GetOnDeckBadRequest   | 400                             | application/json                |
| sdkerrors.GetOnDeckUnauthorized | 401                             | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
