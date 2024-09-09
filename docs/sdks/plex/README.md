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
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
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

| Error Object                                | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| sdkerrors.GetCompanionsDataResponseBody     | 400                                         | application/json                            |
| sdkerrors.GetCompanionsDataPlexResponseBody | 401                                         | application/json                            |
| sdkerrors.SDKError                          | 4xx-5xx                                     | */*                                         |


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
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
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

| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.GetUserFriendsResponseBody     | 400                                      | application/json                         |
| sdkerrors.GetUserFriendsPlexResponseBody | 401                                      | application/json                         |
| sdkerrors.SDKError                       | 4xx-5xx                                  | */*                                      |


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
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
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

| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.GetGeoDataResponseBody     | 400                                  | application/json                     |
| sdkerrors.GetGeoDataPlexResponseBody | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |


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
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
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

| Error Object                          | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.GetHomeDataResponseBody     | 400                                   | application/json                      |
| sdkerrors.GetHomeDataPlexResponseBody | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4xx-5xx                               | */*                                   |


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
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Plex.GetServerResources(ctx, operations.GetServerResourcesRequest{
        XPlexToken: "CV5xoxjTpFKUzBTShsaf",
        IncludeHTTPS: operations.IncludeHTTPSOne.ToPointer(),
        IncludeRelay: operations.IncludeRelayOne.ToPointer(),
        IncludeIPv6: operations.IncludeIPv6One.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlexDevices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetServerResourcesRequest](../../models/operations/getserverresourcesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetServerResourcesResponse](../../models/operations/getserverresourcesresponse.md), error**

### Errors

| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.GetServerResourcesResponseBody     | 400                                          | application/json                             |
| sdkerrors.GetServerResourcesPlexResponseBody | 401                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |


## GetPin

Retrieve a Pin from Plex.tv for authentication flows

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
    res, err := s.Plex.GetPin(ctx, nil, plexgo.String("gcgzw5rz2xovp84b4vha3a40"), plexgo.String("Plex Web"))
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthPinContainer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                             | Type                                                                                                                                                                  | Required                                                                                                                                                              | Description                                                                                                                                                           | Example                                                                                                                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                 | :heavy_check_mark:                                                                                                                                                    | The context to use for the request.                                                                                                                                   |                                                                                                                                                                       |
| `strong`                                                                                                                                                              | **bool*                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                    | Determines the kind of code returned by the API call<br/>Strong codes are used for Pin authentication flows<br/>Non-Strong codes are used for `Plex.tv/link`<br/>     |                                                                                                                                                                       |
| `xPlexClientIdentifier`                                                                                                                                               | **string*                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                    | The unique identifier for the client application<br/>This is used to track the client application and its usage<br/>(UUID, serial number, or other number unique per device)<br/> | gcgzw5rz2xovp84b4vha3a40                                                                                                                                              |
| `xPlexProduct`                                                                                                                                                        | **string*                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                    | N/A                                                                                                                                                                   | Plex Web                                                                                                                                                              |
| `opts`                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                              | :heavy_minus_sign:                                                                                                                                                    | The options for this request.                                                                                                                                         |                                                                                                                                                                       |

### Response

**[*operations.GetPinResponse](../../models/operations/getpinresponse.md), error**

### Errors

| Error Object                 | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.GetPinResponseBody | 400                          | application/json             |
| sdkerrors.SDKError           | 4xx-5xx                      | */*                          |


## GetTokenByPinID

Retrieve an Access Token from Plex.tv after the Pin has been authenticated

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
    res, err := s.Plex.GetTokenByPinID(ctx, 408895, plexgo.String("gcgzw5rz2xovp84b4vha3a40"))
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthPinContainer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                             | Type                                                                                                                                                                  | Required                                                                                                                                                              | Description                                                                                                                                                           | Example                                                                                                                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                 | :heavy_check_mark:                                                                                                                                                    | The context to use for the request.                                                                                                                                   |                                                                                                                                                                       |
| `pinID`                                                                                                                                                               | *int64*                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                    | The PinID to retrieve an access token for                                                                                                                             |                                                                                                                                                                       |
| `xPlexClientIdentifier`                                                                                                                                               | **string*                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                    | The unique identifier for the client application<br/>This is used to track the client application and its usage<br/>(UUID, serial number, or other number unique per device)<br/> | gcgzw5rz2xovp84b4vha3a40                                                                                                                                              |
| `opts`                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                              | :heavy_minus_sign:                                                                                                                                                    | The options for this request.                                                                                                                                         |                                                                                                                                                                       |

### Response

**[*operations.GetTokenByPinIDResponse](../../models/operations/gettokenbypinidresponse.md), error**

### Errors

| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetTokenByPinIDResponseBody     | 400                                       | application/json                          |
| sdkerrors.GetTokenByPinIDPlexResponseBody | 404                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |
