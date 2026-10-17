# Content

## Overview

The actual content of the media provider

### Available Operations

* [GetCollectionItems](#getcollectionitems) - Get items in a collection
* [GetMetadataItem](#getmetadataitem) - Get a metadata item
* [GetAlbums](#getalbums) - Set section albums
* [ListContent](#listcontent) - Get items in the section
* [GetAllLeaves](#getallleaves) - Set section leaves
* [GetArts](#getarts) - Set section artwork
* [GetCategories](#getcategories) - Set section categories
* [GetCluster](#getcluster) - Set section clusters
* [GetSonicPath](#getsonicpath) - Similar tracks to transition from one to another
* [GetFolders](#getfolders) - Get all folder locations
* [ListMoments](#listmoments) - Set section moments
* [GetSonicallySimilar](#getsonicallysimilar) - The nearest audio tracks
* [GetCollectionImage](#getcollectionimage) - Get a collection's image

## GetCollectionItems

Get items in a collection.  Note if this collection contains more than 100 items, paging must be used.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCollectionItems" method="get" path="/library/collections/{collectionId}/items" -->
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

    res, err := s.Content.GetCollectionItems(ctx, operations.GetCollectionItemsRequest{
        CollectionID: 314585,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetCollectionItemsRequest](../../models/operations/getcollectionitemsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetCollectionItemsResponse](../../models/operations/getcollectionitemsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMetadataItem

Get one or more metadata items.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMetadataItem" method="get" path="/library/metadata/{ids}" -->
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

    res, err := s.Content.GetMetadataItem(ctx, operations.GetMetadataItemRequest{
        Ids: []string{},
        AsyncCheckFiles: components.BoolIntTrue.ToPointer(),
        AsyncRefreshLocalMediaAgent: components.BoolIntTrue.ToPointer(),
        AsyncRefreshAnalysis: components.BoolIntTrue.ToPointer(),
        CheckFiles: components.BoolIntTrue.ToPointer(),
        SkipRefresh: components.BoolIntTrue.ToPointer(),
        CheckFileAvailability: components.BoolIntTrue.ToPointer(),
        AsyncAugmentMetadata: components.BoolIntTrue.ToPointer(),
        AugmentCount: components.BoolIntTrue.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetMetadataItemRequest](../../models/operations/getmetadataitemrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetMetadataItemResponse](../../models/operations/getmetadataitemresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAlbums

Get all albums in a music section

### Example Usage

<!-- UsageSnippet language="go" operationID="getAlbums" method="get" path="/library/sections/{sectionId}/albums" -->
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

    res, err := s.Content.GetAlbums(ctx, operations.GetAlbumsRequest{
        SectionID: 817133,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetAlbumsRequest](../../models/operations/getalbumsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetAlbumsResponse](../../models/operations/getalbumsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListContent

Get the items in a section, potentially filtering them.
When `includeCollections=1` is passed, the response may also contain `Collection` items.

### Example Usage

<!-- UsageSnippet language="go" operationID="listContent" method="get" path="/library/sections/{sectionId}/all" -->
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

    res, err := s.Content.ListContent(ctx, operations.ListContentRequest{
        PmsAPIVersion: plexgo.Pointer("1.0"),
        MediaQuery: &components.MediaQuery{
            Type: components.MediaTypeEpisode.ToPointer(),
            Sort: plexgo.Pointer("duration:desc,index"),
            SourceType: plexgo.Pointer[int64](2),
        },
        SectionID: 813218,
        UpdatedAtAfter: plexgo.Pointer[int64](1792000000),
        UpdatedAtBefore: plexgo.Pointer[int64](1792054845),
        AddedAtAfter: plexgo.Pointer[int64](1791200000),
        AddedAtBefore: plexgo.Pointer[int64](1791252847),
        IncludeMeta: components.BoolIntTrue.ToPointer(),
        IncludeGuids: components.BoolIntTrue.ToPointer(),
        IncludeCollections: components.BoolIntTrue.ToPointer(),
        IncludeExternalMedia: components.BoolIntTrue.ToPointer(),
        IncludeAdvanced: components.BoolIntTrue.ToPointer(),
        CheckFiles: components.BoolIntTrue.ToPointer(),
        IncludeRelated: components.BoolIntTrue.ToPointer(),
        IncludeExtras: components.BoolIntTrue.ToPointer(),
        IncludePopularLeaves: components.BoolIntTrue.ToPointer(),
        IncludeConcerts: components.BoolIntTrue.ToPointer(),
        IncludeOnDeck: components.BoolIntTrue.ToPointer(),
        IncludeChapters: components.BoolIntTrue.ToPointer(),
        IncludePreferences: components.BoolIntTrue.ToPointer(),
        IncludeBandwidths: components.BoolIntTrue.ToPointer(),
        IncludeLoudnessRamps: components.BoolIntTrue.ToPointer(),
        IncludeStations: components.BoolIntTrue.ToPointer(),
        IncludeExternalIds: components.BoolIntTrue.ToPointer(),
        IncludeReviews: components.BoolIntTrue.ToPointer(),
        IncludeCredits: components.BoolIntTrue.ToPointer(),
        IncludeArt: components.BoolIntTrue.ToPointer(),
        IncludeThumb: components.BoolIntTrue.ToPointer(),
        IncludeBanner: components.BoolIntTrue.ToPointer(),
        IncludeTheme: components.BoolIntTrue.ToPointer(),
        IncludeFields: plexgo.Pointer("ratingKey,updatedAt,size,totalSize,offset"),
        AsyncAugmentMetadata: components.BoolIntTrue.ToPointer(),
        AsyncRefreshLocalMediaAgent: components.BoolIntTrue.ToPointer(),
        Nocache: components.BoolIntTrue.ToPointer(),
        SkipRefresh: components.BoolIntTrue.ToPointer(),
        IncludeOptionalElements: plexgo.Pointer("Stream"),
        Unwatched: components.BoolIntTrue.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListContentRequest](../../models/operations/listcontentrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListContentResponse](../../models/operations/listcontentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAllLeaves

Get all leaves in a section (such as episodes in a show section)

### Example Usage

<!-- UsageSnippet language="go" operationID="getAllLeaves" method="get" path="/library/sections/{sectionId}/allLeaves" -->
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

    res, err := s.Content.GetAllLeaves(ctx, operations.GetAllLeavesRequest{
        SectionID: 633197,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetAllLeavesRequest](../../models/operations/getallleavesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetAllLeavesResponse](../../models/operations/getallleavesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetArts

Get artwork for a library section

### Example Usage

<!-- UsageSnippet language="go" operationID="getArts" method="get" path="/library/sections/{sectionId}/arts" -->
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

    res, err := s.Content.GetArts(ctx, operations.GetArtsRequest{
        SectionID: 859200,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithArtwork != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetArtsRequest](../../models/operations/getartsrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.GetArtsResponse](../../models/operations/getartsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCategories

Get categories in a library section

### Example Usage

<!-- UsageSnippet language="go" operationID="getCategories" method="get" path="/library/sections/{sectionId}/categories" -->
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

    res, err := s.Content.GetCategories(ctx, operations.GetCategoriesRequest{
        SectionID: 21841,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithArtwork != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetCategoriesRequest](../../models/operations/getcategoriesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetCategoriesResponse](../../models/operations/getcategoriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCluster

Get clusters in a library section (typically for photos)

### Example Usage

<!-- UsageSnippet language="go" operationID="getCluster" method="get" path="/library/sections/{sectionId}/cluster" -->
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

    res, err := s.Content.GetCluster(ctx, operations.GetClusterRequest{
        SectionID: 138560,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithArtwork != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetClusterRequest](../../models/operations/getclusterrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetClusterResponse](../../models/operations/getclusterresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSonicPath

Get a list of audio tracks starting at one and ending at another which are similar across the path

### Example Usage

<!-- UsageSnippet language="go" operationID="getSonicPath" method="get" path="/library/sections/{sectionId}/computePath" -->
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

    res, err := s.Content.GetSonicPath(ctx, operations.GetSonicPathRequest{
        SectionID: 914549,
        StartID: 629990,
        EndID: 687740,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetSonicPathRequest](../../models/operations/getsonicpathrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetSonicPathResponse](../../models/operations/getsonicpathresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetFolders

Get all folder locations of the media in a section

### Example Usage

<!-- UsageSnippet language="go" operationID="getFolders" method="get" path="/library/sections/{sectionId}/location" -->
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

    res, err := s.Content.GetFolders(ctx, operations.GetFoldersRequest{
        SectionID: 892532,
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetFoldersRequest](../../models/operations/getfoldersrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetFoldersResponse](../../models/operations/getfoldersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMoments

Get moments in a library section (typically for photos)

### Example Usage

<!-- UsageSnippet language="go" operationID="listMoments" method="get" path="/library/sections/{sectionId}/moment" -->
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

    res, err := s.Content.ListMoments(ctx, operations.ListMomentsRequest{
        SectionID: 403239,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithArtwork != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListMomentsRequest](../../models/operations/listmomentsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListMomentsResponse](../../models/operations/listmomentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSonicallySimilar

Get the nearest audio tracks to a particular analysis

### Example Usage

<!-- UsageSnippet language="go" operationID="getSonicallySimilar" method="get" path="/library/sections/{sectionId}/nearest" -->
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

    res, err := s.Content.GetSonicallySimilar(ctx, operations.GetSonicallySimilarRequest{
        SectionID: 525956,
        Values: []int64{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetSonicallySimilarRequest](../../models/operations/getsonicallysimilarrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetSonicallySimilarResponse](../../models/operations/getsonicallysimilarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 401                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCollectionImage

Get an image for the collection based on the items within

### Example Usage

<!-- UsageSnippet language="go" operationID="getCollectionImage" method="get" path="/library/collections/{collectionId}/composite/{updatedAt}" -->
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

    res, err := s.Content.GetCollectionImage(ctx, operations.GetCollectionImageRequest{
        CollectionID: 474227,
        UpdatedAt: 759379,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetCollectionImageRequest](../../models/operations/getcollectionimagerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetCollectionImageResponse](../../models/operations/getcollectionimageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |