# Server
(*Server*)

## Overview

Operations against the Plex Media Server System.


### Available Operations

* [GetServerCapabilities](#getservercapabilities) - Server Capabilities
* [GetServerPreferences](#getserverpreferences) - Get Server Preferences
* [GetAvailableClients](#getavailableclients) - Get Available Clients
* [GetDevices](#getdevices) - Get Devices
* [GetServerIdentity](#getserveridentity) - Get Server Identity
* [GetMyPlexAccount](#getmyplexaccount) - Get MyPlex Account
* [GetResizedPhoto](#getresizedphoto) - Get a Resized Photo
* [GetServerList](#getserverlist) - Get Server List

## GetServerCapabilities

Server Capabilities

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetServerCapabilitiesResponse](../../models/operations/getservercapabilitiesresponse.md), error**
| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.GetServerCapabilitiesResponseBody | 401                                         | application/json                            |
| sdkerrors.SDKError                          | 4xx-5xx                                     | */*                                         |

## GetServerPreferences

Get Server Preferences

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
	"net/http"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Server.GetServerPreferences(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetServerPreferencesResponse](../../models/operations/getserverpreferencesresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.GetServerPreferencesResponseBody | 401                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

## GetAvailableClients

Get Available Clients

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Server.GetAvailableClients(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAvailableClientsResponse](../../models/operations/getavailableclientsresponse.md), error**
| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetAvailableClientsResponseBody | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |

## GetDevices

Get Devices

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDevicesResponse](../../models/operations/getdevicesresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.GetDevicesResponseBody | 401                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## GetServerIdentity

Get Server Identity

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetServerIdentityResponse](../../models/operations/getserveridentityresponse.md), error**
| Error Object                            | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.GetServerIdentityResponseBody | 401                                     | application/json                        |
| sdkerrors.SDKError                      | 4xx-5xx                                 | */*                                     |

## GetMyPlexAccount

Returns MyPlex Account Information

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetMyPlexAccountResponse](../../models/operations/getmyplexaccountresponse.md), error**
| Error Object                           | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.GetMyPlexAccountResponseBody | 401                                    | application/json                       |
| sdkerrors.SDKError                     | 4xx-5xx                                | */*                                    |

## GetResizedPhoto

Plex's Photo transcoder is used throughout the service to serve images at specified sizes.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
	"net/http"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Server.GetResizedPhoto(ctx, operations.GetResizedPhotoRequest{
        Width: 110,
        Height: 165,
        Opacity: 643869,
        Blur: 4000,
        MinSize: operations.MinSizeZero,
        Upscale: operations.UpscaleZero,
        URL: "/library/metadata/49564/thumb/1654258204",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetResizedPhotoRequest](../../models/operations/getresizedphotorequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetResizedPhotoResponse](../../models/operations/getresizedphotoresponse.md), error**
| Error Object                          | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.GetResizedPhotoResponseBody | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4xx-5xx                               | */*                                   |

## GetServerList

Get Server List

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetServerListResponse](../../models/operations/getserverlistresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.GetServerListResponseBody | 401                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |
