# \ActivitiesApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelServerActivities**](ActivitiesApi.md#CancelServerActivities) | **Delete** /activities/{activityUUID} | Cancel Server Activities
[**GetServerActivities**](ActivitiesApi.md#GetServerActivities) | **Get** /activities | Get Server Activities



## CancelServerActivities

> CancelServerActivities(ctx, activityUUID).Execute()

Cancel Server Activities



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
    activityUUID := TODO // interface{} | The UUID of the activity to cancel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivitiesApi.CancelServerActivities(context.Background(), activityUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesApi.CancelServerActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityUUID** | [**interface{}**](.md) | The UUID of the activity to cancel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelServerActivitiesRequest struct via the builder pattern


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


## GetServerActivities

> GetServerActivities200Response GetServerActivities(ctx).Execute()

Get Server Activities



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
    resp, r, err := apiClient.ActivitiesApi.GetServerActivities(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesApi.GetServerActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerActivities`: GetServerActivities200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivitiesApi.GetServerActivities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerActivitiesRequest struct via the builder pattern


### Return type

[**GetServerActivities200Response**](GetServerActivities200Response.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

