# \PlexTvApi

All URIs are relative to *https://plex.tv/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanionsData**](PlexTvApi.md#GetCompanionsData) | **Get** /companions | Get Companions Data
[**GetDevices**](PlexTvApi.md#GetDevices) | **Get** /resources | Get Devices
[**GetGeoData**](PlexTvApi.md#GetGeoData) | **Get** /geoip | Get Geo Data
[**GetHomeData**](PlexTvApi.md#GetHomeData) | **Get** /home | Get Home Data
[**GetPin**](PlexTvApi.md#GetPin) | **Post** /pins | Get a Pin
[**GetToken**](PlexTvApi.md#GetToken) | **Get** /pins/{pinID} | Get Access Token
[**GetUserDetails**](PlexTvApi.md#GetUserDetails) | **Get** /user | Get Logged in User
[**GetUserOptOutSettings**](PlexTvApi.md#GetUserOptOutSettings) | **Get** /user/settings/opt_outs | Get User Opt Out Settings
[**GetUserSettings**](PlexTvApi.md#GetUserSettings) | **Get** /user/settings | Get User Settings



## GetCompanionsData

> interface{} GetCompanionsData(ctx).Execute()

Get Companions Data



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
    resp, r, err := apiClient.PlexTvApi.GetCompanionsData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlexTvApi.GetCompanionsData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanionsData`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlexTvApi.GetCompanionsData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanionsDataRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[PlatformVersion](../README.md#PlatformVersion), [ClientIdentifier](../README.md#ClientIdentifier), [Platform](../README.md#Platform), [Version](../README.md#Version), [Device](../README.md#Device), [Product](../README.md#Product), [Token](../README.md#Token), [DeviceName](../README.md#DeviceName)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> interface{} GetDevices(ctx).IncludeHttps(includeHttps).IncludeRelay(includeRelay).IncludeIPv6(includeIPv6).Execute()

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
    includeHttps := TODO // interface{} | Include Https entries in the results (optional)
    includeRelay := TODO // interface{} | Include Relay addresses in the results (optional)
    includeIPv6 := TODO // interface{} | Include IPv6 entries in the results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlexTvApi.GetDevices(context.Background()).IncludeHttps(includeHttps).IncludeRelay(includeRelay).IncludeIPv6(includeIPv6).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlexTvApi.GetDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevices`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlexTvApi.GetDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeHttps** | [**interface{}**](interface{}.md) | Include Https entries in the results | 
 **includeRelay** | [**interface{}**](interface{}.md) | Include Relay addresses in the results | 
 **includeIPv6** | [**interface{}**](interface{}.md) | Include IPv6 entries in the results | 

### Return type

**interface{}**

### Authorization

[PlatformVersion](../README.md#PlatformVersion), [ClientIdentifier](../README.md#ClientIdentifier), [Platform](../README.md#Platform), [Version](../README.md#Version), [Device](../README.md#Device), [Product](../README.md#Product), [Token](../README.md#Token), [DeviceName](../README.md#DeviceName)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeoData

> GetGeoData200Response GetGeoData(ctx).Execute()

Get Geo Data



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
    resp, r, err := apiClient.PlexTvApi.GetGeoData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlexTvApi.GetGeoData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGeoData`: GetGeoData200Response
    fmt.Fprintf(os.Stdout, "Response from `PlexTvApi.GetGeoData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGeoDataRequest struct via the builder pattern


### Return type

[**GetGeoData200Response**](GetGeoData200Response.md)

### Authorization

[PlatformVersion](../README.md#PlatformVersion), [ClientIdentifier](../README.md#ClientIdentifier), [Platform](../README.md#Platform), [Version](../README.md#Version), [Device](../README.md#Device), [Product](../README.md#Product), [Token](../README.md#Token), [DeviceName](../README.md#DeviceName)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHomeData

> GetGeoData200Response GetHomeData(ctx).Execute()

Get Home Data



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
    resp, r, err := apiClient.PlexTvApi.GetHomeData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlexTvApi.GetHomeData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHomeData`: GetGeoData200Response
    fmt.Fprintf(os.Stdout, "Response from `PlexTvApi.GetHomeData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHomeDataRequest struct via the builder pattern


### Return type

[**GetGeoData200Response**](GetGeoData200Response.md)

### Authorization

[PlatformVersion](../README.md#PlatformVersion), [ClientIdentifier](../README.md#ClientIdentifier), [Platform](../README.md#Platform), [Version](../README.md#Version), [Device](../README.md#Device), [Product](../README.md#Product), [Token](../README.md#Token), [DeviceName](../README.md#DeviceName)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPin

> GetPin200Response GetPin(ctx).Strong(strong).Execute()

Get a Pin



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
    strong := TODO // interface{} | Determines the kind of code returned by the API call Strong codes are used for Pin authentication flows Non-Strong codes are used for `Plex.tv/link`  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlexTvApi.GetPin(context.Background()).Strong(strong).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlexTvApi.GetPin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPin`: GetPin200Response
    fmt.Fprintf(os.Stdout, "Response from `PlexTvApi.GetPin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **strong** | [**interface{}**](interface{}.md) | Determines the kind of code returned by the API call Strong codes are used for Pin authentication flows Non-Strong codes are used for &#x60;Plex.tv/link&#x60;  | [default to false]

### Return type

[**GetPin200Response**](GetPin200Response.md)

### Authorization

[ClientIdentifier](../README.md#ClientIdentifier)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> GetToken(ctx, pinID).Execute()

Get Access Token



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
    pinID := TODO // interface{} | The PinID to retrieve an access token for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlexTvApi.GetToken(context.Background(), pinID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlexTvApi.GetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pinID** | [**interface{}**](.md) | The PinID to retrieve an access token for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ClientIdentifier](../README.md#ClientIdentifier)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDetails

> GetUserDetails(ctx).Execute()

Get Logged in User



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
    resp, r, err := apiClient.PlexTvApi.GetUserDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlexTvApi.GetUserDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDetailsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[PlatformVersion](../README.md#PlatformVersion), [ClientIdentifier](../README.md#ClientIdentifier), [Platform](../README.md#Platform), [Version](../README.md#Version), [Device](../README.md#Device), [Product](../README.md#Product), [Token](../README.md#Token), [DeviceName](../README.md#DeviceName)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserOptOutSettings

> GetUserOptOutSettings200Response GetUserOptOutSettings(ctx).Execute()

Get User Opt Out Settings



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
    resp, r, err := apiClient.PlexTvApi.GetUserOptOutSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlexTvApi.GetUserOptOutSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserOptOutSettings`: GetUserOptOutSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `PlexTvApi.GetUserOptOutSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserOptOutSettingsRequest struct via the builder pattern


### Return type

[**GetUserOptOutSettings200Response**](GetUserOptOutSettings200Response.md)

### Authorization

[PlatformVersion](../README.md#PlatformVersion), [ClientIdentifier](../README.md#ClientIdentifier), [Platform](../README.md#Platform), [Version](../README.md#Version), [Device](../README.md#Device), [Product](../README.md#Product), [Token](../README.md#Token), [DeviceName](../README.md#DeviceName)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettings

> interface{} GetUserSettings(ctx).Execute()

Get User Settings



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
    resp, r, err := apiClient.PlexTvApi.GetUserSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlexTvApi.GetUserSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSettings`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlexTvApi.GetUserSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[PlatformVersion](../README.md#PlatformVersion), [ClientIdentifier](../README.md#ClientIdentifier), [Platform](../README.md#Platform), [Version](../README.md#Version), [Device](../README.md#Device), [Product](../README.md#Product), [Token](../README.md#Token), [DeviceName](../README.md#DeviceName)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

