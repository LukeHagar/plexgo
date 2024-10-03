# Plex
(*Plex*)

## Overview

API Calls that perform operations directly against https://Plex.tv


### Available Operations

* [GetCompanionsData](#getcompanionsdata) - Get Companions Data
* [GetUserFriends](#getuserfriends) - Get list of friends of the user logged in
* [GetGeoData](#getgeodata) - Get Geo Data
* [GetHomeData](#gethomedata) - Get Plex Home Data
* [GetServerResources](#getserverresources) - Get Server Resources
* [GetPin](#getpin) - Get a Pin
* [GetTokenByPinID](#gettokenbypinid) - Get Access Token by PinId

## GetCompanionsData

Get Companions Data

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
    res, err := s.Plex.GetCompanionsData(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
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

**[*operations.GetCompanionsDataResponse](../../models/operations/getcompanionsdataresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.GetCompanionsDataBadRequest   | 400                                     | application/json                        |
| sdkerrors.GetCompanionsDataUnauthorized | 401                                     | application/json                        |
| sdkerrors.SDKError                      | 4XX, 5XX                                | \*/\*                                   |

## GetUserFriends

Get friends of provided auth token.

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
    res, err := s.Plex.GetUserFriends(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Friends != nil {
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

**[*operations.GetUserFriendsResponse](../../models/operations/getuserfriendsresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.GetUserFriendsBadRequest   | 400                                  | application/json                     |
| sdkerrors.GetUserFriendsUnauthorized | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4XX, 5XX                             | \*/\*                                |

## GetGeoData

Returns the geolocation and locale data of the caller

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
    res, err := s.Plex.GetGeoData(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GeoData != nil {
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

**[*operations.GetGeoDataResponse](../../models/operations/getgeodataresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.GetGeoDataBadRequest   | 400                              | application/json                 |
| sdkerrors.GetGeoDataUnauthorized | 401                              | application/json                 |
| sdkerrors.SDKError               | 4XX, 5XX                         | \*/\*                            |

## GetHomeData

Retrieves the home data for the authenticated user, including details like home ID, name, guest access information, and subscription status.

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
    res, err := s.Plex.GetHomeData(ctx)
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

**[*operations.GetHomeDataResponse](../../models/operations/gethomedataresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetHomeDataBadRequest   | 400                               | application/json                  |
| sdkerrors.GetHomeDataUnauthorized | 401                               | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## GetServerResources

Get Plex server access tokens and server connections

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
    res, err := s.Plex.GetServerResources(ctx, operations.IncludeHTTPSEnable.ToPointer(), operations.IncludeRelayEnable.ToPointer(), operations.IncludeIPv6Enable.ToPointer(), plexgo.String("3381b62b-9ab7-4e37-827b-203e9809eb58"))
    if err != nil {
        log.Fatal(err)
    }
    if res.PlexDevices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |                                                                                                                    |
| `includeHTTPS`                                                                                                     | [*operations.IncludeHTTPS](../../models/operations/includehttps.md)                                                | :heavy_minus_sign:                                                                                                 | Include Https entries in the results                                                                               | 1                                                                                                                  |
| `includeRelay`                                                                                                     | [*operations.IncludeRelay](../../models/operations/includerelay.md)                                                | :heavy_minus_sign:                                                                                                 | Include Relay addresses in the results <br/>E.g: https://10-0-0-25.bbf8e10c7fa20447cacee74cd9914cde.plex.direct:32400<br/> | 1                                                                                                                  |
| `includeIPv6`                                                                                                      | [*operations.IncludeIPv6](../../models/operations/includeipv6.md)                                                  | :heavy_minus_sign:                                                                                                 | Include IPv6 entries in the results                                                                                | 1                                                                                                                  |
| `clientID`                                                                                                         | **string*                                                                                                          | :heavy_minus_sign:                                                                                                 | An opaque identifier unique to the client (UUID, serial number, or other unique device ID)                         | 3381b62b-9ab7-4e37-827b-203e9809eb58                                                                               |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |                                                                                                                    |

### Response

**[*operations.GetServerResourcesResponse](../../models/operations/getserverresourcesresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.GetServerResourcesBadRequest   | 400                                      | application/json                         |
| sdkerrors.GetServerResourcesUnauthorized | 401                                      | application/json                         |
| sdkerrors.SDKError                       | 4XX, 5XX                                 | \*/\*                                    |

## GetPin

Retrieve a Pin ID from Plex.tv to use for authentication flows

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
    )

    ctx := context.Background()
    res, err := s.Plex.GetPin(ctx, operations.GetPinRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthPinContainer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [operations.GetPinRequest](../../models/operations/getpinrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.GetPinResponse](../../models/operations/getpinresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.GetPinBadRequest | 400                        | application/json           |
| sdkerrors.SDKError         | 4XX, 5XX                   | \*/\*                      |

## GetTokenByPinID

Retrieve an Access Token from Plex.tv after the Pin has been authenticated

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
        plexgo.WithClientID("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithClientName("Plex for Roku"),
        plexgo.WithClientVersion("2.4.1"),
        plexgo.WithPlatform("Roku"),
        plexgo.WithDeviceNickname("Roku 3"),
    )

    ctx := context.Background()
    res, err := s.Plex.GetTokenByPinID(ctx, operations.GetTokenByPinIDRequest{
        PinID: 408895,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthPinContainer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetTokenByPinIDRequest](../../models/operations/gettokenbypinidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetTokenByPinIDResponse](../../models/operations/gettokenbypinidresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.GetTokenByPinIDBadRequest   | 400                                   | application/json                      |
| sdkerrors.GetTokenByPinIDResponseBody | 404                                   | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |