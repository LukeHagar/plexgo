# \UpdaterApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyUpdates**](UpdaterApi.md#ApplyUpdates) | **Put** /updater/apply | Apply Updates
[**CheckForUpdates**](UpdaterApi.md#CheckForUpdates) | **Put** /updater/check | Checking for updates
[**GetUpdateStatus**](UpdaterApi.md#GetUpdateStatus) | **Get** /updater/status | Querying status of updates



## ApplyUpdates

> ApplyUpdates(ctx).Tonight(tonight).Skip(skip).Execute()

Apply Updates



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
    tonight := TODO // interface{} | Indicate that you want the update to run during the next Butler execution. Omitting this or setting it to false indicates that the update should install (optional)
    skip := TODO // interface{} | Indicate that the latest version should be marked as skipped. The <Release> entry for this version will have the `state` set to `skipped`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdaterApi.ApplyUpdates(context.Background()).Tonight(tonight).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdaterApi.ApplyUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tonight** | [**interface{}**](interface{}.md) | Indicate that you want the update to run during the next Butler execution. Omitting this or setting it to false indicates that the update should install | 
 **skip** | [**interface{}**](interface{}.md) | Indicate that the latest version should be marked as skipped. The &lt;Release&gt; entry for this version will have the &#x60;state&#x60; set to &#x60;skipped&#x60;. | 

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


## CheckForUpdates

> CheckForUpdates(ctx).Download(download).Execute()

Checking for updates



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
    download := TODO // interface{} | Indicate that you want to start download any updates found. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdaterApi.CheckForUpdates(context.Background()).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdaterApi.CheckForUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckForUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **download** | [**interface{}**](interface{}.md) | Indicate that you want to start download any updates found. | 

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


## GetUpdateStatus

> GetUpdateStatus(ctx).Execute()

Querying status of updates



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
    resp, r, err := apiClient.UpdaterApi.GetUpdateStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdaterApi.GetUpdateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateStatusRequest struct via the builder pattern


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

