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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
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

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.GetServerCapabilitiesBadRequest   | 400                                         | application/json                            |
| sdkerrors.GetServerCapabilitiesUnauthorized | 401                                         | application/json                            |
| sdkerrors.SDKError                          | 4XX, 5XX                                    | \*/\*                                       |

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
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

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.GetServerPreferencesBadRequest   | 400                                        | application/json                           |
| sdkerrors.GetServerPreferencesUnauthorized | 401                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
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

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetAvailableClientsBadRequest   | 400                                       | application/json                          |
| sdkerrors.GetAvailableClientsUnauthorized | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4XX, 5XX                                  | \*/\*                                     |

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
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

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.GetDevicesBadRequest   | 400                              | application/json                 |
| sdkerrors.GetDevicesUnauthorized | 401                              | application/json                 |
| sdkerrors.SDKError               | 4XX, 5XX                         | \*/\*                            |

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
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

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetServerIdentityRequestTimeout | 408                                       | application/json                          |
| sdkerrors.SDKError                        | 4XX, 5XX                                  | \*/\*                                     |

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
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

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.GetMyPlexAccountBadRequest   | 400                                    | application/json                       |
| sdkerrors.GetMyPlexAccountUnauthorized | 401                                    | application/json                       |
| sdkerrors.SDKError                     | 4XX, 5XX                               | \*/\*                                  |

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
    )

    ctx := context.Background()
    res, err := s.Server.GetResizedPhoto(ctx, operations.GetResizedPhotoRequest{
        Width: 110,
        Height: 165,
        Opacity: 100,
        Blur: 0,
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

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.GetResizedPhotoBadRequest   | 400                                   | application/json                      |
| sdkerrors.GetResizedPhotoUnauthorized | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
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
| `xPlexToken`                                             | *string*                                                 | :heavy_check_mark:                                       | An authentication token, obtained from plex.tv           | CV5xoxjTpFKUzBTShsaf                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMediaProvidersResponse](../../models/operations/getmediaprovidersresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.GetMediaProvidersBadRequest   | 400                                     | application/json                        |
| sdkerrors.GetMediaProvidersUnauthorized | 401                                     | application/json                        |
| sdkerrors.SDKError                      | 4XX, 5XX                                | \*/\*                                   |

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
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

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.GetServerListBadRequest   | 400                                 | application/json                    |
| sdkerrors.GetServerListUnauthorized | 401                                 | application/json                    |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |