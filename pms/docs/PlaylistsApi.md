# \PlaylistsApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPlaylistContents**](PlaylistsApi.md#AddPlaylistContents) | **Put** /playlists/{playlistID}/items | Adding to a Playlist
[**ClearPlaylistContents**](PlaylistsApi.md#ClearPlaylistContents) | **Delete** /playlists/{playlistID}/items | Delete Playlist Contents
[**CreatePlaylist**](PlaylistsApi.md#CreatePlaylist) | **Post** /playlists | Create a Playlist
[**DeletePlaylist**](PlaylistsApi.md#DeletePlaylist) | **Delete** /playlists/{playlistID} | Deletes a Playlist
[**GetPlaylist**](PlaylistsApi.md#GetPlaylist) | **Get** /playlists/{playlistID} | Retrieve Playlist
[**GetPlaylistContents**](PlaylistsApi.md#GetPlaylistContents) | **Get** /playlists/{playlistID}/items | Retrieve Playlist Contents
[**GetPlaylists**](PlaylistsApi.md#GetPlaylists) | **Get** /playlists/all | Get All Playlists
[**UpdatePlaylist**](PlaylistsApi.md#UpdatePlaylist) | **Put** /playlists/{playlistID} | Update a Playlist
[**UploadPlaylist**](PlaylistsApi.md#UploadPlaylist) | **Post** /playlists/upload | Upload Playlist



## AddPlaylistContents

> AddPlaylistContents(ctx, playlistID).Uri(uri).PlayQueueID(playQueueID).Execute()

Adding to a Playlist



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
    playlistID := TODO // interface{} | the ID of the playlist
    uri := TODO // interface{} | the content URI for the playlist
    playQueueID := TODO // interface{} | the play queue to add to a playlist

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.AddPlaylistContents(context.Background(), playlistID).Uri(uri).PlayQueueID(playQueueID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.AddPlaylistContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistID** | [**interface{}**](.md) | the ID of the playlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPlaylistContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uri** | [**interface{}**](interface{}.md) | the content URI for the playlist | 
 **playQueueID** | [**interface{}**](interface{}.md) | the play queue to add to a playlist | 

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


## ClearPlaylistContents

> ClearPlaylistContents(ctx, playlistID).Execute()

Delete Playlist Contents



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
    playlistID := TODO // interface{} | the ID of the playlist

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.ClearPlaylistContents(context.Background(), playlistID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.ClearPlaylistContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistID** | [**interface{}**](.md) | the ID of the playlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearPlaylistContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CreatePlaylist

> CreatePlaylist(ctx).Title(title).Type_(type_).Smart(smart).Uri(uri).PlayQueueID(playQueueID).Execute()

Create a Playlist



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
    title := TODO // interface{} | name of the playlist
    type_ := TODO // interface{} | type of playlist to create
    smart := TODO // interface{} | whether the playlist is smart or not
    uri := TODO // interface{} | the content URI for the playlist (optional)
    playQueueID := TODO // interface{} | the play queue to copy to a playlist (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.CreatePlaylist(context.Background()).Title(title).Type_(type_).Smart(smart).Uri(uri).PlayQueueID(playQueueID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.CreatePlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | [**interface{}**](interface{}.md) | name of the playlist | 
 **type_** | [**interface{}**](interface{}.md) | type of playlist to create | 
 **smart** | [**interface{}**](interface{}.md) | whether the playlist is smart or not | 
 **uri** | [**interface{}**](interface{}.md) | the content URI for the playlist | 
 **playQueueID** | [**interface{}**](interface{}.md) | the play queue to copy to a playlist | 

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


## DeletePlaylist

> DeletePlaylist(ctx, playlistID).Execute()

Deletes a Playlist



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
    playlistID := TODO // interface{} | the ID of the playlist

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.DeletePlaylist(context.Background(), playlistID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.DeletePlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistID** | [**interface{}**](.md) | the ID of the playlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetPlaylist

> GetPlaylist(ctx, playlistID).Execute()

Retrieve Playlist



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
    playlistID := TODO // interface{} | the ID of the playlist

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.GetPlaylist(context.Background(), playlistID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistID** | [**interface{}**](.md) | the ID of the playlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetPlaylistContents

> GetPlaylistContents(ctx, playlistID).Type_(type_).Execute()

Retrieve Playlist Contents



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
    playlistID := TODO // interface{} | the ID of the playlist
    type_ := TODO // interface{} | the metadata type of the item to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.GetPlaylistContents(context.Background(), playlistID).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetPlaylistContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistID** | [**interface{}**](.md) | the ID of the playlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**interface{}**](interface{}.md) | the metadata type of the item to return | 

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


## GetPlaylists

> GetPlaylists(ctx).PlaylistType(playlistType).Smart(smart).Execute()

Get All Playlists



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
    playlistType := TODO // interface{} | limit to a type of playlist. (optional)
    smart := TODO // interface{} | type of playlists to return (default is all). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.GetPlaylists(context.Background()).PlaylistType(playlistType).Smart(smart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetPlaylists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playlistType** | [**interface{}**](interface{}.md) | limit to a type of playlist. | 
 **smart** | [**interface{}**](interface{}.md) | type of playlists to return (default is all). | 

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


## UpdatePlaylist

> UpdatePlaylist(ctx, playlistID).Execute()

Update a Playlist



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
    playlistID := TODO // interface{} | the ID of the playlist

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.UpdatePlaylist(context.Background(), playlistID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.UpdatePlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistID** | [**interface{}**](.md) | the ID of the playlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UploadPlaylist

> UploadPlaylist(ctx).Path(path).Force(force).Execute()

Upload Playlist



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
    path := TODO // interface{} | absolute path to a directory on the server where m3u files are stored, or the absolute path to a playlist file on the server.  If the `path` argument is a directory, that path will be scanned for playlist files to be processed.  Each file in that directory creates a separate playlist, with a name based on the filename of the file that created it.  The GUID of each playlist is based on the filename.  If the `path` argument is a file, that file will be used to create a new playlist, with the name based on the filename of the file that created it.  The GUID of each playlist is based on the filename. 
    force := TODO // interface{} | force overwriting of duplicate playlists. By default, a playlist file uploaded with the same path will overwrite the existing playlist.  The `force` argument is used to disable overwriting. If the `force` argument is set to 0, a new playlist will be created suffixed with the date and time that the duplicate was uploaded. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.UploadPlaylist(context.Background()).Path(path).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.UploadPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | [**interface{}**](interface{}.md) | absolute path to a directory on the server where m3u files are stored, or the absolute path to a playlist file on the server.  If the &#x60;path&#x60; argument is a directory, that path will be scanned for playlist files to be processed.  Each file in that directory creates a separate playlist, with a name based on the filename of the file that created it.  The GUID of each playlist is based on the filename.  If the &#x60;path&#x60; argument is a file, that file will be used to create a new playlist, with the name based on the filename of the file that created it.  The GUID of each playlist is based on the filename.  | 
 **force** | [**interface{}**](interface{}.md) | force overwriting of duplicate playlists. By default, a playlist file uploaded with the same path will overwrite the existing playlist.  The &#x60;force&#x60; argument is used to disable overwriting. If the &#x60;force&#x60; argument is set to 0, a new playlist will be created suffixed with the date and time that the duplicate was uploaded.  | 

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

