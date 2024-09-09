# Server
(*Server*)

## Overview

Operations against the Plex Media Server System.


### Available Operations

* [GetServerCapabilities](#getservercapabilities) - Get Server Capabilities
* [GetServerPreferences](#getserverpreferences) - Get Server Preferences
* [GetAvailableClients](#getavailableclients) - Get Available Clients
* [GetDevices](#getdevices) - Get Devices
* [GetServerIdentity](#getserveridentity) - Get Server Identity
* [GetMyPlexAccount](#getmyplexaccount) - Get MyPlex Account
* [GetResizedPhoto](#getresizedphoto) - Get a Resized Photo
* [GetMediaProviders](#getmediaproviders) - Get Media Providers
* [GetServerList](#getserverlist) - Get Server List

## GetServerCapabilities

Get Server Capabilities

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Server.GetServerCapabilities(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetServerCapabilitiesResponse](../../models/operations/getservercapabilitiesresponse.md), error**

### Errors

| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.GetServerCapabilitiesResponseBody       | 400                                               | application/json                                  |
| sdkerrors.GetServerCapabilitiesServerResponseBody | 401                                               | application/json                                  |
| sdkerrors.SDKError                                | 4xx-5xx                                           | */*                                               |


## GetServerPreferences

Get Server Preferences

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Server.GetServerPreferences(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetServerPreferencesResponse](../../models/operations/getserverpreferencesresponse.md), error**

### Errors

| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.GetServerPreferencesResponseBody       | 400                                              | application/json                                 |
| sdkerrors.GetServerPreferencesServerResponseBody | 401                                              | application/json                                 |
| sdkerrors.SDKError                               | 4xx-5xx                                          | */*                                              |


## GetAvailableClients

Get Available Clients

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Server.GetAvailableClients(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAvailableClientsResponse](../../models/operations/getavailableclientsresponse.md), error**

### Errors

| Error Object                                    | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| sdkerrors.GetAvailableClientsResponseBody       | 400                                             | application/json                                |
| sdkerrors.GetAvailableClientsServerResponseBody | 401                                             | application/json                                |
| sdkerrors.SDKError                              | 4xx-5xx                                         | */*                                             |


## GetDevices

Get Devices

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Server.GetDevices(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDevicesResponse](../../models/operations/getdevicesresponse.md), error**

### Errors

| Error Object                           | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.GetDevicesResponseBody       | 400                                    | application/json                       |
| sdkerrors.GetDevicesServerResponseBody | 401                                    | application/json                       |
| sdkerrors.SDKError                     | 4xx-5xx                                | */*                                    |


## GetServerIdentity

This request is useful to determine if the server is online or offline

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Server.GetServerIdentity(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetServerIdentityResponse](../../models/operations/getserveridentityresponse.md), error**

### Errors

| Error Object                            | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.GetServerIdentityResponseBody | 408                                     | application/json                        |
| sdkerrors.SDKError                      | 4xx-5xx                                 | */*                                     |


## GetMyPlexAccount

Returns MyPlex Account Information

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Server.GetMyPlexAccount(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMyPlexAccountResponse](../../models/operations/getmyplexaccountresponse.md), error**

### Errors

| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.GetMyPlexAccountResponseBody       | 400                                          | application/json                             |
| sdkerrors.GetMyPlexAccountServerResponseBody | 401                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |


## GetResizedPhoto

Plex's Photo transcoder is used throughout the service to serve images at specified sizes.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Server.GetResizedPhoto(ctx, operations.GetResizedPhotoRequest{
        Width: 110,
        Height: 165,
        Opacity: 100,
        Blur: 20,
        MinSize: operations.MinSizeOne,
        Upscale: operations.UpscaleOne,
        URL: "/library/metadata/49564/thumb/1654258204",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetResizedPhotoRequest](../../models/operations/getresizedphotorequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetResizedPhotoResponse](../../models/operations/getresizedphotoresponse.md), error**

### Errors

| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.GetResizedPhotoResponseBody       | 400                                         | application/json                            |
| sdkerrors.GetResizedPhotoServerResponseBody | 401                                         | application/json                            |
| sdkerrors.SDKError                          | 4xx-5xx                                     | */*                                         |


## GetMediaProviders

Retrieves media providers and their features from the Plex server.

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Server.GetMediaProviders(ctx, "CV5xoxjTpFKUzBTShsaf")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `xPlexToken`                                             | *string*                                                 | :heavy_check_mark:                                       | Plex Authentication Token                                | CV5xoxjTpFKUzBTShsaf                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMediaProvidersResponse](../../models/operations/getmediaprovidersresponse.md), error**

### Errors

| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.GetMediaProvidersResponseBody       | 400                                           | application/json                              |
| sdkerrors.GetMediaProvidersServerResponseBody | 401                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |


## GetServerList

Get Server List

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Server.GetServerList(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetServerListResponse](../../models/operations/getserverlistresponse.md), error**

### Errors

| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetServerListResponseBody       | 400                                       | application/json                          |
| sdkerrors.GetServerListServerResponseBody | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |
