# Hubs
(*Hubs*)

## Overview

Hubs are a structured two-dimensional container for media, generally represented by multiple horizontal rows.


### Available Operations

* [GetGlobalHubs](#getglobalhubs) - Get Global Hubs
* [GetLibraryHubs](#getlibraryhubs) - Get library specific hubs

## GetGlobalHubs

Get Global Hubs filtered by the parameters provided.

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
        plexgo.WithXPlexClientIdentifier("Postman"),
    )

    var count *float64 = plexgo.Float64(1262.49)

    var onlyTransient *operations.OnlyTransient = operations.OnlyTransientOne.ToPointer()
    
    ctx := context.Background()
    res, err := s.Hubs.GetGlobalHubs(ctx, count, onlyTransient)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                             | Type                                                                                                                                                  | Required                                                                                                                                              | Description                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                 | :heavy_check_mark:                                                                                                                                    | The context to use for the request.                                                                                                                   |
| `count`                                                                                                                                               | **float64*                                                                                                                                            | :heavy_minus_sign:                                                                                                                                    | The number of items to return with each hub.                                                                                                          |
| `onlyTransient`                                                                                                                                       | [*operations.OnlyTransient](../../models/operations/onlytransient.md)                                                                                 | :heavy_minus_sign:                                                                                                                                    | Only return hubs which are "transient", meaning those which are prone to changing after media playback or addition (e.g. On Deck, or Recently Added). |


### Response

**[*operations.GetGlobalHubsResponse](../../models/operations/getglobalhubsresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.GetGlobalHubsResponseBody | 401                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## GetLibraryHubs

This endpoint will return a list of library specific hubs


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
        plexgo.WithXPlexClientIdentifier("Postman"),
    )

    var sectionID float64 = 6728.76

    var count *float64 = plexgo.Float64(9010.22)

    var onlyTransient *operations.QueryParamOnlyTransient = operations.QueryParamOnlyTransientZero.ToPointer()
    
    ctx := context.Background()
    res, err := s.Hubs.GetLibraryHubs(ctx, sectionID, count, onlyTransient)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                             | Type                                                                                                                                                  | Required                                                                                                                                              | Description                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                 | :heavy_check_mark:                                                                                                                                    | The context to use for the request.                                                                                                                   |
| `sectionID`                                                                                                                                           | *float64*                                                                                                                                             | :heavy_check_mark:                                                                                                                                    | the Id of the library to query                                                                                                                        |
| `count`                                                                                                                                               | **float64*                                                                                                                                            | :heavy_minus_sign:                                                                                                                                    | The number of items to return with each hub.                                                                                                          |
| `onlyTransient`                                                                                                                                       | [*operations.QueryParamOnlyTransient](../../models/operations/queryparamonlytransient.md)                                                             | :heavy_minus_sign:                                                                                                                                    | Only return hubs which are "transient", meaning those which are prone to changing after media playback or addition (e.g. On Deck, or Recently Added). |


### Response

**[*operations.GetLibraryHubsResponse](../../models/operations/getlibraryhubsresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.GetLibraryHubsResponseBody | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |
