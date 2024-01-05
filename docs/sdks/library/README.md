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
* [GetLatestLibraryItems](#getlatestlibraryitems) - Get Latest Library Items
* [GetCommonLibraryItems](#getcommonlibraryitems) - Get Common Library Items
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
	"net/http"
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

    if res.StatusCode == http.StatusOK {
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

Returns details for the library. This can be thought of as an interstitial endpoint because it contains information about the library, rather than content itself. These details are:

- A list of `Directory` objects: These used to be used by clients to build a menuing system. There are four flavors of directory found here:
  - Primary: (e.g. all, On Deck) These are still used in some clients to provide "shortcuts" to subsets of media. However, with the exception of On Deck, all of them can be created by media queries, and the desire is to allow these to be customized by users.
  - Secondary: These are marked with `secondary="1"` and were used by old clients to provide nested menus allowing for primative (but structured) navigation.
  - Special: There is a By Folder entry which allows browsing the media by the underlying filesystem structure, and there's a completely obsolete entry marked `search="1"` which used to be used to allow clients to build search dialogs on the fly.
- A list of `Type` objects: These represent the types of things found in this library, and for each one, a list of `Filter` and `Sort` objects. These can be used to build rich controls around a grid of media to allow filtering and organizing. Note that these filters and sorts are optional, and without them, the client won't render any filtering controls. The `Type` object contains:
  - `key`: This provides the root endpoint returning the actual media list for the type.
  - `type`: This is the metadata type for the type (if a standard Plex type).
  - `title`: The title for for the content of this type (e.g. "Movies").
- Each `Filter` object contains a description of the filter. Note that it is not an exhaustive list of the full media query language, but an inportant subset useful for top-level API.
  - `filter`: This represents the filter name used for the filter, which can be used to construct complex media queries with.
  - `filterType`: This is either `string`, `integer`, or `boolean`, and describes the type of values used for the filter.
  - `key`: This provides the endpoint where the possible range of values for the filter can be retrieved (e.g. for a "Genre" filter, it returns a list of all the genres in the library). This will include a `type` argument that matches the metadata type of the Type element.
  - `title`: The title for the filter.
- Each `Sort` object contains a description of the sort field.
  - `defaultDirection`: Can be either `asc` or `desc`, and specifies the default direction for the sort field (e.g. titles default to alphabetically ascending).
  - `descKey` and `key`: Contains the parameters passed to the `sort=...` media query for each direction of the sort.
  - `title`: The title of the field.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"context"
	"log"
	"net/http"
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

    if res.StatusCode == http.StatusOK {
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

This endpoint will return a list of library items filtered by the filter and type provided


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


    var sectionID float64 = 4510.92

    var type_ *float64 = 760.66

    var filter *string = "string"

    ctx := context.Background()
    res, err := s.Library.GetLibraryItems(ctx, sectionID, type_, filter)
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
| `sectionID`                                           | *float64*                                             | :heavy_check_mark:                                    | the Id of the library to query                        |
| `type_`                                               | **float64*                                            | :heavy_minus_sign:                                    | item type                                             |
| `filter`                                              | **string*                                             | :heavy_minus_sign:                                    | the filter parameter                                  |


### Response

**[*operations.GetLibraryItemsResponse](../../models/operations/getlibraryitemsresponse.md), error**
| Error Object                          | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.GetLibraryItemsResponseBody | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4xx-5xx                               | */*                                   |

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

## GetLatestLibraryItems

This endpoint will return a list of the latest library items filtered by the filter and type provided


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


    var sectionID float64 = 7171.54

    var type_ float64 = 8015.12

    var filter *string = "string"

    ctx := context.Background()
    res, err := s.Library.GetLatestLibraryItems(ctx, sectionID, type_, filter)
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
| `sectionID`                                           | *float64*                                             | :heavy_check_mark:                                    | the Id of the library to query                        |
| `type_`                                               | *float64*                                             | :heavy_check_mark:                                    | item type                                             |
| `filter`                                              | **string*                                             | :heavy_minus_sign:                                    | the filter parameter                                  |


### Response

**[*operations.GetLatestLibraryItemsResponse](../../models/operations/getlatestlibraryitemsresponse.md), error**
| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.GetLatestLibraryItemsResponseBody | 401                                         | application/json                            |
| sdkerrors.SDKError                          | 4xx-5xx                                     | */*                                         |

## GetCommonLibraryItems

Represents a "Common" item. It contains only the common attributes of the items selected by the provided filter


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


    var sectionID float64 = 2710.37

    var type_ float64 = 2760.31

    var filter *string = "string"

    ctx := context.Background()
    res, err := s.Library.GetCommonLibraryItems(ctx, sectionID, type_, filter)
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
| `sectionID`                                           | *float64*                                             | :heavy_check_mark:                                    | the Id of the library to query                        |
| `type_`                                               | *float64*                                             | :heavy_check_mark:                                    | item type                                             |
| `filter`                                              | **string*                                             | :heavy_minus_sign:                                    | the filter parameter                                  |


### Response

**[*operations.GetCommonLibraryItemsResponse](../../models/operations/getcommonlibraryitemsresponse.md), error**
| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.GetCommonLibraryItemsResponseBody | 401                                         | application/json                            |
| sdkerrors.SDKError                          | 4xx-5xx                                     | */*                                         |

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
	"net/http"
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

    if res.StatusCode == http.StatusOK {
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
	"net/http"
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

    if res.StatusCode == http.StatusOK {
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
