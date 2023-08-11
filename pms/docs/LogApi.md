# \LogApi

All URIs are relative to *http://10.10.10.47:32400*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnablePaperTrail**](LogApi.md#EnablePaperTrail) | **Get** /log/networked | Enabling Papertrail
[**LogLine**](LogApi.md#LogLine) | **Get** /log | Logging a single line message.
[**LogMultiLine**](LogApi.md#LogMultiLine) | **Post** /log | Logging a multi-line message



## EnablePaperTrail

> EnablePaperTrail(ctx).Execute()

Enabling Papertrail



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
    resp, r, err := apiClient.LogApi.EnablePaperTrail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogApi.EnablePaperTrail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePaperTrailRequest struct via the builder pattern


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


## LogLine

> LogLine(ctx).Level(level).Message(message).Source(source).Execute()

Logging a single line message.



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
    level := TODO // interface{} | An integer log level to write to the PMS log with.   0: Error   1: Warning   2: Info  3: Debug   4: Verbose 
    message := TODO // interface{} | The text of the message to write to the log.
    source := TODO // interface{} | a string indicating the source of the message.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogApi.LogLine(context.Background()).Level(level).Message(message).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogApi.LogLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | [**interface{}**](interface{}.md) | An integer log level to write to the PMS log with.   0: Error   1: Warning   2: Info  3: Debug   4: Verbose  | 
 **message** | [**interface{}**](interface{}.md) | The text of the message to write to the log. | 
 **source** | [**interface{}**](interface{}.md) | a string indicating the source of the message. | 

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


## LogMultiLine

> LogMultiLine(ctx).Execute()

Logging a multi-line message



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
    resp, r, err := apiClient.LogApi.LogMultiLine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogApi.LogMultiLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogMultiLineRequest struct via the builder pattern


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

