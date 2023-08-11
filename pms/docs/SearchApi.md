# \SearchApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchResults**](SearchApi.md#GetSearchResults) | **Get** /search | Get Search Results
[**PerformSearch**](SearchApi.md#PerformSearch) | **Get** /hubs/search | Perform a search
[**PerformVoiceSearch**](SearchApi.md#PerformVoiceSearch) | **Get** /hubs/search/voice | Perform a voice search



## GetSearchResults

> GetSearchResults200Response GetSearchResults(ctx).Query(query).Execute()

Get Search Results



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
    query := TODO // interface{} | The search query string to use

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetSearchResults(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetSearchResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchResults`: GetSearchResults200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**interface{}**](interface{}.md) | The search query string to use | 

### Return type

[**GetSearchResults200Response**](GetSearchResults200Response.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformSearch

> PerformSearch(ctx).Query(query).SectionId(sectionId).Limit(limit).Execute()

Perform a search



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
    query := TODO // interface{} | The query term
    sectionId := TODO // interface{} | This gives context to the search, and can result in re-ordering of search result hubs (optional)
    limit := TODO // interface{} | The number of items to return per hub (optional) (default to 3)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.PerformSearch(context.Background()).Query(query).SectionId(sectionId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.PerformSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**interface{}**](interface{}.md) | The query term | 
 **sectionId** | [**interface{}**](interface{}.md) | This gives context to the search, and can result in re-ordering of search result hubs | 
 **limit** | [**interface{}**](interface{}.md) | The number of items to return per hub | [default to 3]

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


## PerformVoiceSearch

> PerformVoiceSearch(ctx).Query(query).SectionId(sectionId).Limit(limit).Execute()

Perform a voice search



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
    query := TODO // interface{} | The query term
    sectionId := TODO // interface{} | This gives context to the search, and can result in re-ordering of search result hubs (optional)
    limit := TODO // interface{} | The number of items to return per hub (optional) (default to 3)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.PerformVoiceSearch(context.Background()).Query(query).SectionId(sectionId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.PerformVoiceSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformVoiceSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**interface{}**](interface{}.md) | The query term | 
 **sectionId** | [**interface{}**](interface{}.md) | This gives context to the search, and can result in re-ordering of search result hubs | 
 **limit** | [**interface{}**](interface{}.md) | The number of items to return per hub | [default to 3]

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

