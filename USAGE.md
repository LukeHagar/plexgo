<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
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
		TranscodeType:             components.TranscodeTypeMusic,
		Extension:                 operations.ExtensionMpd,
		AdvancedSubtitles:         components.AdvancedSubtitlesBurn.ToPointer(),
		AudioBoost:                plexgo.Pointer[int64](50),
		AudioChannelCount:         plexgo.Pointer[int64](5),
		AutoAdjustQuality:         components.BoolIntOne.ToPointer(),
		AutoAdjustSubtitle:        components.BoolIntOne.ToPointer(),
		DirectPlay:                components.BoolIntOne.ToPointer(),
		DirectStream:              components.BoolIntOne.ToPointer(),
		DirectStreamAudio:         components.BoolIntOne.ToPointer(),
		DisableResolutionRotation: components.BoolIntOne.ToPointer(),
		HasMDE:                    components.BoolIntOne.ToPointer(),
		Location:                  operations.StartTranscodeSessionQueryParamLocationWan.ToPointer(),
		MediaBufferSize:           plexgo.Pointer[int64](102400),
		MediaIndex:                plexgo.Pointer[int64](0),
		MusicBitrate:              plexgo.Pointer[int64](5000),
		Offset:                    plexgo.Pointer[float64](90.5),
		PartIndex:                 plexgo.Pointer[int64](0),
		Path:                      plexgo.Pointer("/library/metadata/151671"),
		PeakBitrate:               plexgo.Pointer[int64](12000),
		PhotoResolution:           plexgo.Pointer("1080x1080"),
		Protocol:                  operations.StartTranscodeSessionQueryParamProtocolDash.ToPointer(),
		SecondsPerSegment:         plexgo.Pointer[int64](5),
		SubtitleSize:              plexgo.Pointer[int64](50),
		VideoBitrate:              plexgo.Pointer[int64](12000),
		VideoQuality:              plexgo.Pointer[int64](50),
		VideoResolution:           plexgo.Pointer("1080x1080"),
		XPlexClientProfileExtra:   plexgo.Pointer("add-limitation(scope=videoCodec&scopeName=*&type=upperBound&name=video.frameRate&value=60&replace=true)+append-transcode-target-codec(type=videoProfile&context=streaming&videoCodec=h264%2Chevc&audioCodec=aac&protocol=dash)"),
		XPlexClientProfileName:    plexgo.Pointer("generic"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseStream != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->