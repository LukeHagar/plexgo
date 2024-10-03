# Media
(*Media*)

## Overview

API Calls interacting with Plex Media Server Media


### Available Operations

* [MarkPlayed](#markplayed) - Mark Media Played
* [MarkUnplayed](#markunplayed) - Mark Media Unplayed
* [UpdatePlayProgress](#updateplayprogress) - Update Media Play Progress
* [GetBannerImage](#getbannerimage) - Get Banner Image
* [GetThumbImage](#getthumbimage) - Get Thumb Image

## MarkPlayed

This will mark the provided media key as Played.

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
    )

    ctx := context.Background()
    res, err := s.Media.MarkPlayed(ctx, 59398)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `key`                                                    | *float64*                                                | :heavy_check_mark:                                       | The media key to mark as played                          | 59398                                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.MarkPlayedResponse](../../models/operations/markplayedresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.MarkPlayedBadRequest   | 400                              | application/json                 |
| sdkerrors.MarkPlayedUnauthorized | 401                              | application/json                 |
| sdkerrors.SDKError               | 4XX, 5XX                         | \*/\*                            |

## MarkUnplayed

This will mark the provided media key as Unplayed.

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
    )

    ctx := context.Background()
    res, err := s.Media.MarkUnplayed(ctx, 59398)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `key`                                                    | *float64*                                                | :heavy_check_mark:                                       | The media key to mark as Unplayed                        | 59398                                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.MarkUnplayedResponse](../../models/operations/markunplayedresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.MarkUnplayedBadRequest   | 400                                | application/json                   |
| sdkerrors.MarkUnplayedUnauthorized | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpdatePlayProgress

This API command can be used to update the play progress of a media item.


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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
    )

    ctx := context.Background()
    res, err := s.Media.UpdatePlayProgress(ctx, "<key>", 90000, "played")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |                                                                     |
| `key`                                                               | *string*                                                            | :heavy_check_mark:                                                  | the media key                                                       |                                                                     |
| `time`                                                              | *float64*                                                           | :heavy_check_mark:                                                  | The time, in milliseconds, used to set the media playback progress. | 90000                                                               |
| `state`                                                             | *string*                                                            | :heavy_check_mark:                                                  | The playback state of the media item.                               | played                                                              |
| `opts`                                                              | [][operations.Option](../../models/operations/option.md)            | :heavy_minus_sign:                                                  | The options for this request.                                       |                                                                     |

### Response

**[*operations.UpdatePlayProgressResponse](../../models/operations/updateplayprogressresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.UpdatePlayProgressBadRequest   | 400                                      | application/json                         |
| sdkerrors.UpdatePlayProgressUnauthorized | 401                                      | application/json                         |
| sdkerrors.SDKError                       | 4XX, 5XX                                 | \*/\*                                    |

## GetBannerImage

Gets the banner image of the media item

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
    )

    ctx := context.Background()
    res, err := s.Media.GetBannerImage(ctx, operations.GetBannerImageRequest{
        RatingKey: 9518,
        Width: 396,
        Height: 396,
        MinSize: 1,
        Upscale: 1,
        XPlexToken: "CV5xoxjTpFKUzBTShsaf",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetBannerImageRequest](../../models/operations/getbannerimagerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetBannerImageResponse](../../models/operations/getbannerimageresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.GetBannerImageBadRequest   | 400                                  | application/json                     |
| sdkerrors.GetBannerImageUnauthorized | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4XX, 5XX                             | \*/\*                                |

## GetThumbImage

Gets the thumbnail image of the media item

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
    )

    ctx := context.Background()
    res, err := s.Media.GetThumbImage(ctx, operations.GetThumbImageRequest{
        RatingKey: 9518,
        Width: 396,
        Height: 396,
        MinSize: 1,
        Upscale: 1,
        XPlexToken: "CV5xoxjTpFKUzBTShsaf",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetThumbImageRequest](../../models/operations/getthumbimagerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetThumbImageResponse](../../models/operations/getthumbimageresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.GetThumbImageBadRequest   | 400                                 | application/json                    |
| sdkerrors.GetThumbImageUnauthorized | 401                                 | application/json                    |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |