# \ServerApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableClients**](ServerApi.md#GetAvailableClients) | **Get** /clients | Get Available Clients
[**GetDevices**](ServerApi.md#GetDevices) | **Get** /devices | Get Devices
[**GetMyPlexAccount**](ServerApi.md#GetMyPlexAccount) | **Get** /myplex/account | Get MyPlex Account
[**GetResizedPhoto**](ServerApi.md#GetResizedPhoto) | **Get** /photo/:/transcode | Get a Resized Photo
[**GetServerCapabilities**](ServerApi.md#GetServerCapabilities) | **Get** / | Server Capabilities
[**GetServerIdentity**](ServerApi.md#GetServerIdentity) | **Get** /identity | Get Server Identity
[**GetServerList**](ServerApi.md#GetServerList) | **Get** /servers | Get Server List
[**GetServerPreferences**](ServerApi.md#GetServerPreferences) | **Get** /:/prefs | Get Server Preferences



## GetAvailableClients

> interface{} GetAvailableClients(ctx).Execute()

Get Available Clients



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.GetAvailableClients(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetAvailableClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableClients`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.GetAvailableClients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableClientsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> GetDevices200Response GetDevices(ctx).Execute()

Get Devices



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.GetDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevices`: GetDevices200Response
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.GetDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


### Return type

[**GetDevices200Response**](GetDevices200Response.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyPlexAccount

> GetMyPlexAccount200Response GetMyPlexAccount(ctx).Execute()

Get MyPlex Account



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.GetMyPlexAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetMyPlexAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMyPlexAccount`: GetMyPlexAccount200Response
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.GetMyPlexAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPlexAccountRequest struct via the builder pattern


### Return type

[**GetMyPlexAccount200Response**](GetMyPlexAccount200Response.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResizedPhoto

> GetResizedPhoto(ctx).Width(width).Height(height).Opacity(opacity).Blur(blur).MinSize(minSize).Upscale(upscale).Url(url).Execute()

Get a Resized Photo



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
    width := TODO // interface{} | The width for the resized photo
    height := TODO // interface{} | The height for the resized photo
    opacity := TODO // interface{} | The opacity for the resized photo (default to 100)
    blur := TODO // interface{} | The width for the resized photo
    minSize := TODO // interface{} | images are always scaled proportionally. A value of '1' in minSize will make the smaller native dimension the dimension resized against.
    upscale := TODO // interface{} | allow images to be resized beyond native dimensions.
    url := TODO // interface{} | path to image within Plex

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.GetResizedPhoto(context.Background()).Width(width).Height(height).Opacity(opacity).Blur(blur).MinSize(minSize).Upscale(upscale).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetResizedPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResizedPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **width** | [**interface{}**](interface{}.md) | The width for the resized photo | 
 **height** | [**interface{}**](interface{}.md) | The height for the resized photo | 
 **opacity** | [**interface{}**](interface{}.md) | The opacity for the resized photo | [default to 100]
 **blur** | [**interface{}**](interface{}.md) | The width for the resized photo | 
 **minSize** | [**interface{}**](interface{}.md) | images are always scaled proportionally. A value of &#39;1&#39; in minSize will make the smaller native dimension the dimension resized against. | 
 **upscale** | [**interface{}**](interface{}.md) | allow images to be resized beyond native dimensions. | 
 **url** | [**interface{}**](interface{}.md) | path to image within Plex | 

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


## GetServerCapabilities

> GetServerCapabilities200Response GetServerCapabilities(ctx).Execute()

Server Capabilities



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.GetServerCapabilities(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetServerCapabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerCapabilities`: GetServerCapabilities200Response
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.GetServerCapabilities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerCapabilitiesRequest struct via the builder pattern


### Return type

[**GetServerCapabilities200Response**](GetServerCapabilities200Response.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerIdentity

> GetServerIdentity200Response GetServerIdentity(ctx).Execute()

Get Server Identity



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.GetServerIdentity(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetServerIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerIdentity`: GetServerIdentity200Response
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.GetServerIdentity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerIdentityRequest struct via the builder pattern


### Return type

[**GetServerIdentity200Response**](GetServerIdentity200Response.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerList

> GetServerList200Response GetServerList(ctx).Execute()

Get Server List



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.GetServerList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetServerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerList`: GetServerList200Response
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.GetServerList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerListRequest struct via the builder pattern


### Return type

[**GetServerList200Response**](GetServerList200Response.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerPreferences

> GetServerPreferences(ctx).Execute()

Get Server Preferences



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.GetServerPreferences(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.GetServerPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerPreferencesRequest struct via the builder pattern


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

