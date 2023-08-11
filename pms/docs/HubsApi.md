# \HubsApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalHubs**](HubsApi.md#GetGlobalHubs) | **Get** /hubs | Get Global Hubs
[**GetLibraryHubs**](HubsApi.md#GetLibraryHubs) | **Get** /hubs/sections/{sectionId} | Get library specific hubs



## GetGlobalHubs

> GetGlobalHubs(ctx).Count(count).OnlyTransient(onlyTransient).Execute()

Get Global Hubs



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
    count := TODO // interface{} | The number of items to return with each hub. (optional)
    onlyTransient := TODO // interface{} | Only return hubs which are \"transient\", meaning those which are prone to changing after media playback or addition (e.g. On Deck, or Recently Added). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HubsApi.GetGlobalHubs(context.Background()).Count(count).OnlyTransient(onlyTransient).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubsApi.GetGlobalHubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalHubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | [**interface{}**](interface{}.md) | The number of items to return with each hub. | 
 **onlyTransient** | [**interface{}**](interface{}.md) | Only return hubs which are \&quot;transient\&quot;, meaning those which are prone to changing after media playback or addition (e.g. On Deck, or Recently Added). | 

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


## GetLibraryHubs

> GetLibraryHubs(ctx, sectionId).Count(count).OnlyTransient(onlyTransient).Execute()

Get library specific hubs



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
    sectionId := TODO // interface{} | the Id of the library to query
    count := TODO // interface{} | The number of items to return with each hub. (optional)
    onlyTransient := TODO // interface{} | Only return hubs which are \"transient\", meaning those which are prone to changing after media playback or addition (e.g. On Deck, or Recently Added). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HubsApi.GetLibraryHubs(context.Background(), sectionId).Count(count).OnlyTransient(onlyTransient).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubsApi.GetLibraryHubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionId** | [**interface{}**](.md) | the Id of the library to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryHubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | [**interface{}**](interface{}.md) | The number of items to return with each hub. | 
 **onlyTransient** | [**interface{}**](interface{}.md) | Only return hubs which are \&quot;transient\&quot;, meaning those which are prone to changing after media playback or addition (e.g. On Deck, or Recently Added). | 

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

