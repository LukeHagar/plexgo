# Video
(*Video*)

## Overview

API Calls that perform operations with Plex Media Server Videos


### Available Operations

* [GetTimeline](#gettimeline) - Get the timeline for a media item
* [StartUniversalTranscode](#startuniversaltranscode) - Start Universal Transcode

## GetTimeline

Get the timeline for a media item

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
        plexgo.WithXPlexClientIdentifier("Postman"),
    )

    ctx := context.Background()
    res, err := s.Video.GetTimeline(ctx, operations.GetTimelineRequest{
        RatingKey: 23409,
        Key: "/library/metadata/23409",
        State: operations.StatePlaying,
        HasMDE: 1,
        Time: 2000,
        Duration: 10000,
        Context: "home:hub.continueWatching",
        PlayQueueItemID: 1,
        PlayBackTime: 2000,
        Row: 1,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetTimelineRequest](../../models/operations/gettimelinerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetTimelineResponse](../../models/operations/gettimelineresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetTimelineResponseBody | 401                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## StartUniversalTranscode

Begin a Universal Transcode Session

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
        plexgo.WithXPlexClientIdentifier("Postman"),
    )

    ctx := context.Background()
    res, err := s.Video.StartUniversalTranscode(ctx, operations.StartUniversalTranscodeRequest{
        HasMDE: 1,
        Path: "/library/metadata/23409",
        MediaIndex: 0,
        PartIndex: 0,
        Protocol: "hls",
        FastSeek: plexgo.Float64(0),
        DirectPlay: plexgo.Float64(0),
        DirectStream: plexgo.Float64(0),
        SubtitleSize: plexgo.Float64(100),
        Subtites: plexgo.String("burn"),
        AudioBoost: plexgo.Float64(100),
        Location: plexgo.String("lan"),
        MediaBufferSize: plexgo.Float64(102400),
        Session: plexgo.String("zvcage8b7rkioqcm8f4uns4c"),
        AddDebugOverlay: plexgo.Float64(0),
        AutoAdjustQuality: plexgo.Float64(0),
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.StartUniversalTranscodeRequest](../../models/operations/startuniversaltranscoderequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.StartUniversalTranscodeResponse](../../models/operations/startuniversaltranscoderesponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.StartUniversalTranscodeResponseBody | 401                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |
