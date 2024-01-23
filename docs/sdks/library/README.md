# Library
(*Library*)

## Overview

API Calls interacting with Plex Media Server Libraries


### Available Operations

* [GetFileHash](#getfilehash) - Get Hash Value
* [GetRecentlyAdded](#getrecentlyadded) - Get Recently Added
* [GetLibraries](#getlibraries) - Get All Libraries
* [GetLibrary](#getlibrary) - Get Library Details
* [DeleteLibrary](#deletelibrary) - Delete Library Section
* [GetLibraryItems](#getlibraryitems) - Get Library Items
* [RefreshLibrary](#refreshlibrary) - Refresh Library
* [SearchLibrary](#searchlibrary) - Search Library
* [GetMetadata](#getmetadata) - Get Items Metadata
* [GetMetadataChildren](#getmetadatachildren) - Get Items Children
* [GetOnDeck](#getondeck) - Get On Deck

## GetFileHash

This resource returns hash values for local files

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
	"net/http"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var url_ string = "file://C:\Image.png&type=13"

    var type_ *float64 = 4462.17

    ctx := context.Background()
    res, err := s.Library.GetFileHash(ctx, url_, type_)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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


### Response

**[*operations.GetFileHashResponse](../../models/operations/getfilehashresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetFileHashResponseBody | 401                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## GetRecentlyAdded

This endpoint will return the recently added content.


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
    )

    ctx := context.Background()
    res, err := s.Library.GetRecentlyAdded(ctx)
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


### Response

**[*operations.GetRecentlyAddedResponse](../../models/operations/getrecentlyaddedresponse.md), error**
| Error Object                           | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.GetRecentlyAddedResponseBody | 401                                    | application/json                       |
| sdkerrors.SDKError                     | 4xx-5xx                                | */*                                    |

## GetLibraries

A library section (commonly referred to as just a library) is a collection of media. 
Libraries are typed, and depending on their type provide either a flat or a hierarchical view of the media. 
For example, a music library has an artist > albums > tracks structure, whereas a movie library is flat.

Libraries have features beyond just being a collection of media; for starters, they include information about supported types, filters and sorts. 
This allows a client to provide a rich interface around the media (e.g. allow sorting movies by release year).


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
    )

    ctx := context.Background()
    res, err := s.Library.GetLibraries(ctx)
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


### Response

**[*operations.GetLibrariesResponse](../../models/operations/getlibrariesresponse.md), error**
| Error Object                       | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GetLibrariesResponseBody | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4xx-5xx                            | */*                                |

## GetLibrary

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
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var sectionID float64 = 1000

    var includeDetails *operations.IncludeDetails = operations.IncludeDetailsZero

    ctx := context.Background()
    res, err := s.Library.GetLibrary(ctx, sectionID, includeDetails)
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
| `sectionID`                                                                                                                                                                                | *float64*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                         | the Id of the library to query                                                                                                                                                             | 1000                                                                                                                                                                                       |
| `includeDetails`                                                                                                                                                                           | [*operations.IncludeDetails](../../models/operations/includedetails.md)                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                         | Whether or not to include details for a section (types, filters, and sorts). <br/>Only exists for backwards compatibility, media providers other than the server libraries have it on always.<br/> |                                                                                                                                                                                            |


### Response

**[*operations.GetLibraryResponse](../../models/operations/getlibraryresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.GetLibraryResponseBody | 401                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## DeleteLibrary

Delate a library using a specific section

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
	"net/http"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var sectionID float64 = 1000

    ctx := context.Background()
    res, err := s.Library.DeleteLibrary(ctx, sectionID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `sectionID`                                           | *float64*                                             | :heavy_check_mark:                                    | the Id of the library to query                        | 1000                                                  |


### Response

**[*operations.DeleteLibraryResponse](../../models/operations/deletelibraryresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.DeleteLibraryResponseBody | 401                                 | application/json                    |
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
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var sectionID int64 = 451092

    var tag operations.Tag = operations.TagUnwatched

    ctx := context.Background()
    res, err := s.Library.GetLibraryItems(ctx, sectionID, tag)
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
| `sectionID`                                           | *int64*                                               | :heavy_check_mark:                                    | the Id of the library to query                        |
| `tag`                                                 | [operations.Tag](../../models/operations/tag.md)      | :heavy_check_mark:                                    | A key representing a specific tag within the section. |


### Response

**[*operations.GetLibraryItemsResponse](../../models/operations/getlibraryitemsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RefreshLibrary

This endpoint Refreshes the library.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
	"net/http"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var sectionID float64 = 934.16

    ctx := context.Background()
    res, err := s.Library.RefreshLibrary(ctx, sectionID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `sectionID`                                           | *float64*                                             | :heavy_check_mark:                                    | the Id of the library to refresh                      |


### Response

**[*operations.RefreshLibraryResponse](../../models/operations/refreshlibraryresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.RefreshLibraryResponseBody | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

## SearchLibrary

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
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var sectionID int64 = 933505

    var type_ operations.Type = operations.TypeFour

    ctx := context.Background()
    res, err := s.Library.SearchLibrary(ctx, sectionID, type_)
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
| `sectionID`                                           | *int64*                                               | :heavy_check_mark:                                    | the Id of the library to query                        |
| `type_`                                               | [operations.Type](../../models/operations/type.md)    | :heavy_check_mark:                                    | Plex content type to search for                       |


### Response

**[*operations.SearchLibraryResponse](../../models/operations/searchlibraryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetMetadata

This endpoint will return the metadata of a library item specified with the ratingKey.


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
    )


    var ratingKey float64 = 8382.31

    ctx := context.Background()
    res, err := s.Library.GetMetadata(ctx, ratingKey)
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
| `ratingKey`                                           | *float64*                                             | :heavy_check_mark:                                    | the id of the library item to return the children of. |


### Response

**[*operations.GetMetadataResponse](../../models/operations/getmetadataresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetMetadataResponseBody | 401                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## GetMetadataChildren

This endpoint will return the children of of a library item specified with the ratingKey.


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
    )


    var ratingKey float64 = 1539.14

    ctx := context.Background()
    res, err := s.Library.GetMetadataChildren(ctx, ratingKey)
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
| `ratingKey`                                           | *float64*                                             | :heavy_check_mark:                                    | the id of the library item to return the children of. |


### Response

**[*operations.GetMetadataChildrenResponse](../../models/operations/getmetadatachildrenresponse.md), error**
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetMetadataChildrenResponseBody | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |

## GetOnDeck

This endpoint will return the on deck content.


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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetOnDeckResponse](../../models/operations/getondeckresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.GetOnDeckResponseBody | 401                             | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
