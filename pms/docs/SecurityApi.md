# \SecurityApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSourceConnectionInformation**](SecurityApi.md#GetSourceConnectionInformation) | **Get** /security/resources | Get Source Connection Information
[**GetTransientToken**](SecurityApi.md#GetTransientToken) | **Get** /security/token | Get a Transient Token.



## GetSourceConnectionInformation

> GetSourceConnectionInformation(ctx).Source(source).Execute()

Get Source Connection Information



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
    source := TODO // interface{} | The source identifier with an included prefix.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetSourceConnectionInformation(context.Background()).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetSourceConnectionInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceConnectionInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | [**interface{}**](interface{}.md) | The source identifier with an included prefix. | 

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


## GetTransientToken

> GetTransientToken(ctx).Type_(type_).Scope(scope).Execute()

Get a Transient Token.



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
    type_ := TODO // interface{} | `delegation` - This is the only supported `type` parameter.
    scope := TODO // interface{} | `all` - This is the only supported `scope` parameter.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetTransientToken(context.Background()).Type_(type_).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetTransientToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransientTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**interface{}**](interface{}.md) | &#x60;delegation&#x60; - This is the only supported &#x60;type&#x60; parameter. | 
 **scope** | [**interface{}**](interface{}.md) | &#x60;all&#x60; - This is the only supported &#x60;scope&#x60; parameter. | 

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

