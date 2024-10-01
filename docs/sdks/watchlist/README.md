# Watchlist
(*Watchlist*)

## Overview

API Calls that perform operations with Plex Media Server Watchlists


### Available Operations

* [GetWatchList](#getwatchlist) - Get User Watchlist

## GetWatchList

Get User Watchlist

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
    res, err := s.Watchlist.GetWatchList(ctx, operations.GetWatchListRequest{
        Filter: operations.FilterAvailable,
        XPlexContainerStart: plexgo.Int(0),
        XPlexContainerSize: plexgo.Int(50),
        XPlexToken: "CV5xoxjTpFKUzBTShsaf",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetWatchListRequest](../../models/operations/getwatchlistrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetWatchListResponse](../../models/operations/getwatchlistresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GetWatchListBadRequest   | 400                                | application/json                   |
| sdkerrors.GetWatchListUnauthorized | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |