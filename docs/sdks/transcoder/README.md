# Transcoder
(*Transcoder*)

## Overview

API Operations against the Transcoder

### Available Operations

* [TranscodeImage](#transcodeimage) - Transcode an image
* [MakeDecision](#makedecision) - Make a decision on media playback
* [TriggerFallback](#triggerfallback) - Manually trigger a transcoder fallback
* [TranscodeSubtitles](#transcodesubtitles) - Transcode subtitles
* [StartTranscodeSession](#starttranscodesession) - Start A Transcoding Session

## TranscodeImage

Transcode an image, possibly changing format or size

### Example Usage

<!-- UsageSnippet language="go" operationID="transcodeImage" method="get" path="/photo/:/transcode" -->
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

    res, err := s.Transcoder.TranscodeImage(ctx, operations.TranscodeImageRequest{
        URL: plexgo.Pointer("/library/metadata/265/thumb/1715112705"),
        Background: plexgo.Pointer("#ff5522"),
        Upscale: components.BoolIntOne.ToPointer(),
        MinSize: components.BoolIntOne.ToPointer(),
        Rotate: components.BoolIntOne.ToPointer(),
        BlendColor: plexgo.Pointer("#ff5522"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredImageJpegResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.TranscodeImageRequest](../../models/operations/transcodeimagerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.TranscodeImageResponse](../../models/operations/transcodeimageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## MakeDecision

Make a decision on media playback based on client profile, and requested settings such as bandwidth and resolution.

### Example Usage

<!-- UsageSnippet language="go" operationID="makeDecision" method="get" path="/{transcodeType}/:/transcode/universal/decision" -->
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

    res, err := s.Transcoder.MakeDecision(ctx, operations.MakeDecisionRequest{
        TranscodeType: components.TranscodeTypeMusic,
        AdvancedSubtitles: components.AdvancedSubtitlesBurn.ToPointer(),
        AudioBoost: plexgo.Pointer[int64](50),
        AudioChannelCount: plexgo.Pointer[int64](5),
        AutoAdjustQuality: components.BoolIntOne.ToPointer(),
        AutoAdjustSubtitle: components.BoolIntOne.ToPointer(),
        DirectPlay: components.BoolIntOne.ToPointer(),
        DirectStream: components.BoolIntOne.ToPointer(),
        DirectStreamAudio: components.BoolIntOne.ToPointer(),
        DisableResolutionRotation: components.BoolIntOne.ToPointer(),
        HasMDE: components.BoolIntOne.ToPointer(),
        Location: operations.LocationWan.ToPointer(),
        MediaBufferSize: plexgo.Pointer[int64](102400),
        MediaIndex: plexgo.Pointer[int64](0),
        MusicBitrate: plexgo.Pointer[int64](5000),
        Offset: plexgo.Pointer[float64](90.5),
        PartIndex: plexgo.Pointer[int64](0),
        Path: plexgo.Pointer("/library/metadata/151671"),
        PeakBitrate: plexgo.Pointer[int64](12000),
        PhotoResolution: plexgo.Pointer("1080x1080"),
        Protocol: operations.ProtocolDash.ToPointer(),
        SecondsPerSegment: plexgo.Pointer[int64](5),
        SubtitleSize: plexgo.Pointer[int64](50),
        VideoBitrate: plexgo.Pointer[int64](12000),
        VideoQuality: plexgo.Pointer[int64](50),
        VideoResolution: plexgo.Pointer("1080x1080"),
        XPlexClientProfileExtra: plexgo.Pointer("add-limitation(scope=videoCodec&scopeName=*&type=upperBound&name=video.frameRate&value=60&replace=true)+append-transcode-target-codec(type=videoProfile&context=streaming&videoCodec=h264%2Chevc&audioCodec=aac&protocol=dash)"),
        XPlexClientProfileName: plexgo.Pointer("generic"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MediaContainerWithDecision != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.MakeDecisionRequest](../../models/operations/makedecisionrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.MakeDecisionResponse](../../models/operations/makedecisionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## TriggerFallback

Manually trigger a transcoder fallback ex: HEVC to h.264 or hw to sw

### Example Usage

<!-- UsageSnippet language="go" operationID="triggerFallback" method="post" path="/{transcodeType}/:/transcode/universal/fallback" -->
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

    res, err := s.Transcoder.TriggerFallback(ctx, operations.TriggerFallbackRequest{
        TranscodeType: components.TranscodeTypeAudio,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.TriggerFallbackRequest](../../models/operations/triggerfallbackrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.TriggerFallbackResponse](../../models/operations/triggerfallbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## TranscodeSubtitles

Only transcode subtitle streams.

### Example Usage

<!-- UsageSnippet language="go" operationID="transcodeSubtitles" method="get" path="/{transcodeType}/:/transcode/universal/subtitles" -->
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

    res, err := s.Transcoder.TranscodeSubtitles(ctx, operations.TranscodeSubtitlesRequest{
        TranscodeType: components.TranscodeTypeAudio,
        AdvancedSubtitles: components.AdvancedSubtitlesBurn.ToPointer(),
        AudioBoost: plexgo.Pointer[int64](50),
        AudioChannelCount: plexgo.Pointer[int64](5),
        AutoAdjustQuality: components.BoolIntOne.ToPointer(),
        AutoAdjustSubtitle: components.BoolIntOne.ToPointer(),
        DirectPlay: components.BoolIntOne.ToPointer(),
        DirectStream: components.BoolIntOne.ToPointer(),
        DirectStreamAudio: components.BoolIntOne.ToPointer(),
        DisableResolutionRotation: components.BoolIntOne.ToPointer(),
        HasMDE: components.BoolIntOne.ToPointer(),
        Location: operations.QueryParamLocationWan.ToPointer(),
        MediaBufferSize: plexgo.Pointer[int64](102400),
        MediaIndex: plexgo.Pointer[int64](0),
        MusicBitrate: plexgo.Pointer[int64](5000),
        Offset: plexgo.Pointer[float64](90.5),
        PartIndex: plexgo.Pointer[int64](0),
        Path: plexgo.Pointer("/library/metadata/151671"),
        PeakBitrate: plexgo.Pointer[int64](12000),
        PhotoResolution: plexgo.Pointer("1080x1080"),
        Protocol: operations.QueryParamProtocolDash.ToPointer(),
        SecondsPerSegment: plexgo.Pointer[int64](5),
        SubtitleSize: plexgo.Pointer[int64](50),
        VideoBitrate: plexgo.Pointer[int64](12000),
        VideoQuality: plexgo.Pointer[int64](50),
        VideoResolution: plexgo.Pointer("1080x1080"),
        XPlexClientProfileExtra: plexgo.Pointer("add-limitation(scope=videoCodec&scopeName=*&type=upperBound&name=video.frameRate&value=60&replace=true)+append-transcode-target-codec(type=videoProfile&context=streaming&videoCodec=h264%2Chevc&audioCodec=aac&protocol=dash)"),
        XPlexClientProfileName: plexgo.Pointer("generic"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.TranscodeSubtitlesRequest](../../models/operations/transcodesubtitlesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.TranscodeSubtitlesResponse](../../models/operations/transcodesubtitlesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## StartTranscodeSession

Starts the transcoder and returns the corresponding streaming resource document.

### Example Usage

<!-- UsageSnippet language="go" operationID="startTranscodeSession" method="get" path="/{transcodeType}/:/transcode/universal/start.{extension}" -->
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

    res, err := s.Transcoder.StartTranscodeSession(ctx, operations.StartTranscodeSessionRequest{
        TranscodeType: components.TranscodeTypeMusic,
        Extension: operations.ExtensionMpd,
        AdvancedSubtitles: components.AdvancedSubtitlesBurn.ToPointer(),
        AudioBoost: plexgo.Pointer[int64](50),
        AudioChannelCount: plexgo.Pointer[int64](5),
        AutoAdjustQuality: components.BoolIntOne.ToPointer(),
        AutoAdjustSubtitle: components.BoolIntOne.ToPointer(),
        DirectPlay: components.BoolIntOne.ToPointer(),
        DirectStream: components.BoolIntOne.ToPointer(),
        DirectStreamAudio: components.BoolIntOne.ToPointer(),
        DisableResolutionRotation: components.BoolIntOne.ToPointer(),
        HasMDE: components.BoolIntOne.ToPointer(),
        Location: operations.StartTranscodeSessionQueryParamLocationWan.ToPointer(),
        MediaBufferSize: plexgo.Pointer[int64](102400),
        MediaIndex: plexgo.Pointer[int64](0),
        MusicBitrate: plexgo.Pointer[int64](5000),
        Offset: plexgo.Pointer[float64](90.5),
        PartIndex: plexgo.Pointer[int64](0),
        Path: plexgo.Pointer("/library/metadata/151671"),
        PeakBitrate: plexgo.Pointer[int64](12000),
        PhotoResolution: plexgo.Pointer("1080x1080"),
        Protocol: operations.StartTranscodeSessionQueryParamProtocolDash.ToPointer(),
        SecondsPerSegment: plexgo.Pointer[int64](5),
        SubtitleSize: plexgo.Pointer[int64](50),
        VideoBitrate: plexgo.Pointer[int64](12000),
        VideoQuality: plexgo.Pointer[int64](50),
        VideoResolution: plexgo.Pointer("1080x1080"),
        XPlexClientProfileExtra: plexgo.Pointer("add-limitation(scope=videoCodec&scopeName=*&type=upperBound&name=video.frameRate&value=60&replace=true)+append-transcode-target-codec(type=videoProfile&context=streaming&videoCodec=h264%2Chevc&audioCodec=aac&protocol=dash)"),
        XPlexClientProfileName: plexgo.Pointer("generic"),
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.StartTranscodeSessionRequest](../../models/operations/starttranscodesessionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.StartTranscodeSessionResponse](../../models/operations/starttranscodesessionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |