# \VideoApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTimeline**](VideoApi.md#GetTimeline) | **Get** /:/timeline | Get the timeline for a media item
[**StartUniversalTranscode**](VideoApi.md#StartUniversalTranscode) | **Get** /video/:/transcode/universal/start.mpd | Start Universal Transcode



## GetTimeline

> GetTimeline(ctx).RatingKey(ratingKey).Key(key).State(state).HasMDE(hasMDE).Time(time).Duration(duration).Context(context).PlayQueueItemID(playQueueItemID).PlayBackTime(playBackTime).Row(row).Execute()

Get the timeline for a media item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ratingKey := TODO // interface{} | The rating key of the media item
    key := TODO // interface{} | The key of the media item to get the timeline for
    state := TODO // interface{} | The state of the media item
    hasMDE := TODO // interface{} | Whether the media item has MDE
    time := TODO // interface{} | The time of the media item
    duration := TODO // interface{} | The duration of the media item
    context := TODO // interface{} | The context of the media item
    playQueueItemID := TODO // interface{} | The play queue item ID of the media item
    playBackTime := TODO // interface{} | The playback time of the media item
    row := TODO // interface{} | The row of the media item

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VideoApi.GetTimeline(context.Background()).RatingKey(ratingKey).Key(key).State(state).HasMDE(hasMDE).Time(time).Duration(duration).Context(context).PlayQueueItemID(playQueueItemID).PlayBackTime(playBackTime).Row(row).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoApi.GetTimeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratingKey** | [**interface{}**](interface{}.md) | The rating key of the media item | 
 **key** | [**interface{}**](interface{}.md) | The key of the media item to get the timeline for | 
 **state** | [**interface{}**](interface{}.md) | The state of the media item | 
 **hasMDE** | [**interface{}**](interface{}.md) | Whether the media item has MDE | 
 **time** | [**interface{}**](interface{}.md) | The time of the media item | 
 **duration** | [**interface{}**](interface{}.md) | The duration of the media item | 
 **context** | [**interface{}**](interface{}.md) | The context of the media item | 
 **playQueueItemID** | [**interface{}**](interface{}.md) | The play queue item ID of the media item | 
 **playBackTime** | [**interface{}**](interface{}.md) | The playback time of the media item | 
 **row** | [**interface{}**](interface{}.md) | The row of the media item | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartUniversalTranscode

> StartUniversalTranscode(ctx).HasMDE(hasMDE).Path(path).MediaIndex(mediaIndex).PartIndex(partIndex).Protocol(protocol).FastSeek(fastSeek).DirectPlay(directPlay).DirectStream(directStream).SubtitleSize(subtitleSize).Subtites(subtites).AudioBoost(audioBoost).Location(location).MediaBufferSize(mediaBufferSize).Session(session).AddDebugOverlay(addDebugOverlay).AutoAdjustQuality(autoAdjustQuality).Execute()

Start Universal Transcode



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hasMDE := TODO // interface{} | Whether the media item has MDE
    path := TODO // interface{} | The path to the media item to transcode
    mediaIndex := TODO // interface{} | The index of the media item to transcode
    partIndex := TODO // interface{} | The index of the part to transcode
    protocol := TODO // interface{} | The protocol to use for the transcode session
    fastSeek := TODO // interface{} | Whether to use fast seek or not (optional)
    directPlay := TODO // interface{} | Whether to use direct play or not (optional)
    directStream := TODO // interface{} | Whether to use direct stream or not (optional)
    subtitleSize := TODO // interface{} | The size of the subtitles (optional)
    subtites := TODO // interface{} | The subtitles (optional)
    audioBoost := TODO // interface{} | The audio boost (optional)
    location := TODO // interface{} | The location of the transcode session (optional)
    mediaBufferSize := TODO // interface{} | The size of the media buffer (optional)
    session := TODO // interface{} | The session ID (optional)
    addDebugOverlay := TODO // interface{} | Whether to add a debug overlay or not (optional)
    autoAdjustQuality := TODO // interface{} | Whether to auto adjust quality or not (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VideoApi.StartUniversalTranscode(context.Background()).HasMDE(hasMDE).Path(path).MediaIndex(mediaIndex).PartIndex(partIndex).Protocol(protocol).FastSeek(fastSeek).DirectPlay(directPlay).DirectStream(directStream).SubtitleSize(subtitleSize).Subtites(subtites).AudioBoost(audioBoost).Location(location).MediaBufferSize(mediaBufferSize).Session(session).AddDebugOverlay(addDebugOverlay).AutoAdjustQuality(autoAdjustQuality).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoApi.StartUniversalTranscode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartUniversalTranscodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hasMDE** | [**interface{}**](interface{}.md) | Whether the media item has MDE | 
 **path** | [**interface{}**](interface{}.md) | The path to the media item to transcode | 
 **mediaIndex** | [**interface{}**](interface{}.md) | The index of the media item to transcode | 
 **partIndex** | [**interface{}**](interface{}.md) | The index of the part to transcode | 
 **protocol** | [**interface{}**](interface{}.md) | The protocol to use for the transcode session | 
 **fastSeek** | [**interface{}**](interface{}.md) | Whether to use fast seek or not | 
 **directPlay** | [**interface{}**](interface{}.md) | Whether to use direct play or not | 
 **directStream** | [**interface{}**](interface{}.md) | Whether to use direct stream or not | 
 **subtitleSize** | [**interface{}**](interface{}.md) | The size of the subtitles | 
 **subtites** | [**interface{}**](interface{}.md) | The subtitles | 
 **audioBoost** | [**interface{}**](interface{}.md) | The audio boost | 
 **location** | [**interface{}**](interface{}.md) | The location of the transcode session | 
 **mediaBufferSize** | [**interface{}**](interface{}.md) | The size of the media buffer | 
 **session** | [**interface{}**](interface{}.md) | The session ID | 
 **addDebugOverlay** | [**interface{}**](interface{}.md) | Whether to add a debug overlay or not | 
 **autoAdjustQuality** | [**interface{}**](interface{}.md) | Whether to auto adjust quality or not | 

### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

