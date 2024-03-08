# Media
(*Media*)

## Overview

API Calls interacting with Plex Media Server Media


### Available Operations

* [MarkPlayed](#markplayed) - Mark Media Played
* [MarkUnplayed](#markunplayed) - Mark Media Unplayed
* [UpdatePlayProgress](#updateplayprogress) - Update Media Play Progress

## MarkPlayed

This will mark the provided media key as Played.

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


    var key float64 = 59398

    ctx := context.Background()
    res, err := s.Media.MarkPlayed(ctx, key)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `key`                                                 | *float64*                                             | :heavy_check_mark:                                    | The media key to mark as played                       | 59398                                                 |


### Response

**[*operations.MarkPlayedResponse](../../models/operations/markplayedresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.MarkPlayedResponseBody | 401                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## MarkUnplayed

This will mark the provided media key as Unplayed.

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


    var key float64 = 59398

    ctx := context.Background()
    res, err := s.Media.MarkUnplayed(ctx, key)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `key`                                                 | *float64*                                             | :heavy_check_mark:                                    | The media key to mark as Unplayed                     | 59398                                                 |


### Response

**[*operations.MarkUnplayedResponse](../../models/operations/markunplayedresponse.md), error**
| Error Object                       | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.MarkUnplayedResponseBody | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4xx-5xx                            | */*                                |

## UpdatePlayProgress

This API command can be used to update the play progress of a media item.


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


    var key string = "<value>"

    var time float64 = 6900.91

    var state string = "<value>"

    ctx := context.Background()
    res, err := s.Media.UpdatePlayProgress(ctx, key, time, state)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |
| `key`                                                               | *string*                                                            | :heavy_check_mark:                                                  | the media key                                                       |
| `time`                                                              | *float64*                                                           | :heavy_check_mark:                                                  | The time, in milliseconds, used to set the media playback progress. |
| `state`                                                             | *string*                                                            | :heavy_check_mark:                                                  | The playback state of the media item.                               |


### Response

**[*operations.UpdatePlayProgressResponse](../../models/operations/updateplayprogressresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.UpdatePlayProgressResponseBody | 401                                      | application/json                         |
| sdkerrors.SDKError                       | 4xx-5xx                                  | */*                                      |
