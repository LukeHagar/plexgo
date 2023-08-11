# \LibraryApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLibrary**](LibraryApi.md#DeleteLibrary) | **Delete** /library/sections/{sectionId} | Delete Library Section
[**GetCommonLibraryItems**](LibraryApi.md#GetCommonLibraryItems) | **Get** /library/sections/{sectionId}/common | Get Common Library Items
[**GetFileHash**](LibraryApi.md#GetFileHash) | **Get** /library/hashes | Get Hash Value
[**GetLatestLibraryItems**](LibraryApi.md#GetLatestLibraryItems) | **Get** /library/sections/{sectionId}/latest | Get Latest Library Items
[**GetLibraries**](LibraryApi.md#GetLibraries) | **Get** /library/sections | Get All Libraries
[**GetLibrary**](LibraryApi.md#GetLibrary) | **Get** /library/sections/{sectionId} | Get Library Details
[**GetLibraryItems**](LibraryApi.md#GetLibraryItems) | **Get** /library/sections/{sectionId}/all | Get Library Items
[**GetMetadata**](LibraryApi.md#GetMetadata) | **Get** /library/metadata/{ratingKey} | Get Items Metadata
[**GetMetadataChildren**](LibraryApi.md#GetMetadataChildren) | **Get** /library/metadata/{ratingKey}/children | Get Items Children
[**GetOnDeck**](LibraryApi.md#GetOnDeck) | **Get** /library/onDeck | Get On Deck
[**GetRecentlyAdded**](LibraryApi.md#GetRecentlyAdded) | **Get** /library/recentlyAdded | Get Recently Added
[**RefreshLibrary**](LibraryApi.md#RefreshLibrary) | **Get** /library/sections/{sectionId}/refresh | Refresh Library



## DeleteLibrary

> DeleteLibrary(ctx, sectionId).Execute()

Delete Library Section



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LibraryApi.DeleteLibrary(context.Background(), sectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.DeleteLibrary``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLibraryRequest struct via the builder pattern


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


## GetCommonLibraryItems

> GetCommonLibraryItems(ctx, sectionId).Type_(type_).Filter(filter).Execute()

Get Common Library Items



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
    type_ := TODO // interface{} | item type
    filter := TODO // interface{} | the filter parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LibraryApi.GetCommonLibraryItems(context.Background(), sectionId).Type_(type_).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetCommonLibraryItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetCommonLibraryItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**interface{}**](interface{}.md) | item type | 
 **filter** | [**interface{}**](interface{}.md) | the filter parameter | 

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


## GetFileHash

> GetFileHash(ctx).Url(url).Type_(type_).Execute()

Get Hash Value



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
    url := TODO // interface{} | This is the path to the local file, must be prefixed by `file://`
    type_ := TODO // interface{} | Item type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LibraryApi.GetFileHash(context.Background()).Url(url).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetFileHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | [**interface{}**](interface{}.md) | This is the path to the local file, must be prefixed by &#x60;file://&#x60; | 
 **type_** | [**interface{}**](interface{}.md) | Item type | 

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


## GetLatestLibraryItems

> GetLatestLibraryItems(ctx, sectionId).Type_(type_).Filter(filter).Execute()

Get Latest Library Items



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
    type_ := TODO // interface{} | item type
    filter := TODO // interface{} | the filter parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LibraryApi.GetLatestLibraryItems(context.Background(), sectionId).Type_(type_).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetLatestLibraryItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetLatestLibraryItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**interface{}**](interface{}.md) | item type | 
 **filter** | [**interface{}**](interface{}.md) | the filter parameter | 

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


## GetLibraries

> GetLibraries(ctx).Execute()

Get All Libraries



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
    resp, r, err := apiClient.LibraryApi.GetLibraries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetLibraries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibrariesRequest struct via the builder pattern


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


## GetLibrary

> GetLibrary(ctx, sectionId).IncludeDetails(includeDetails).Execute()

Get Library Details



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
    includeDetails := TODO // interface{} | Whether or not to include details for a section (types, filters, and sorts).  Only exists for backwards compatibility, media providers other than the server libraries have it on always.  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LibraryApi.GetLibrary(context.Background(), sectionId).IncludeDetails(includeDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetLibrary``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDetails** | [**interface{}**](interface{}.md) | Whether or not to include details for a section (types, filters, and sorts).  Only exists for backwards compatibility, media providers other than the server libraries have it on always.  | [default to 0]

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


## GetLibraryItems

> GetLibraryItems(ctx, sectionId).Type_(type_).Filter(filter).Execute()

Get Library Items



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
    type_ := TODO // interface{} | item type (optional)
    filter := TODO // interface{} | the filter parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LibraryApi.GetLibraryItems(context.Background(), sectionId).Type_(type_).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetLibraryItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetLibraryItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**interface{}**](interface{}.md) | item type | 
 **filter** | [**interface{}**](interface{}.md) | the filter parameter | 

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


## GetMetadata

> GetMetadata(ctx, ratingKey).Execute()

Get Items Metadata



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
    ratingKey := TODO // interface{} | the id of the library item to return the children of.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LibraryApi.GetMetadata(context.Background(), ratingKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ratingKey** | [**interface{}**](.md) | the id of the library item to return the children of. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataRequest struct via the builder pattern


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


## GetMetadataChildren

> GetMetadataChildren(ctx, ratingKey).Execute()

Get Items Children



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
    ratingKey := TODO // interface{} | the id of the library item to return the children of.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LibraryApi.GetMetadataChildren(context.Background(), ratingKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetMetadataChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ratingKey** | [**interface{}**](.md) | the id of the library item to return the children of. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataChildrenRequest struct via the builder pattern


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


## GetOnDeck

> GetOnDeck200Response GetOnDeck(ctx).Execute()

Get On Deck



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
    resp, r, err := apiClient.LibraryApi.GetOnDeck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetOnDeck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOnDeck`: GetOnDeck200Response
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetOnDeck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnDeckRequest struct via the builder pattern


### Return type

[**GetOnDeck200Response**](GetOnDeck200Response.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecentlyAdded

> GetRecentlyAdded200Response GetRecentlyAdded(ctx).Execute()

Get Recently Added



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
    resp, r, err := apiClient.LibraryApi.GetRecentlyAdded(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.GetRecentlyAdded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecentlyAdded`: GetRecentlyAdded200Response
    fmt.Fprintf(os.Stdout, "Response from `LibraryApi.GetRecentlyAdded`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentlyAddedRequest struct via the builder pattern


### Return type

[**GetRecentlyAdded200Response**](GetRecentlyAdded200Response.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshLibrary

> RefreshLibrary(ctx, sectionId).Execute()

Refresh Library



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
    sectionId := TODO // interface{} | the Id of the library to refresh

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LibraryApi.RefreshLibrary(context.Background(), sectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LibraryApi.RefreshLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionId** | [**interface{}**](.md) | the Id of the library to refresh | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshLibraryRequest struct via the builder pattern


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

