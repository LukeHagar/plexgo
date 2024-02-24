# Plex
(*Plex*)

## Overview

API Calls that perform operations directly against https://Plex.tv


### Available Operations

* [GetPin](#getpin) - Get a Pin
* [GetToken](#gettoken) - Get Access Token

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
    s := plexgo.New()


    var xPlexClientIdentifier string = "<value>"

    var strong *bool = plexgo.Bool(false)

    ctx := context.Background()
    res, err := s.Plex.GetPin(ctx, xPlexClientIdentifier, strong)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                             | Type                                                                                                                                                                  | Required                                                                                                                                                              | Description                                                                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                 | :heavy_check_mark:                                                                                                                                                    | The context to use for the request.                                                                                                                                   |
| `xPlexClientIdentifier`                                                                                                                                               | *string*                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                    | The unique identifier for the client application<br/>This is used to track the client application and its usage<br/>(UUID, serial number, or other number unique per device)<br/> |
| `strong`                                                                                                                                                              | **bool*                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                    | Determines the kind of code returned by the API call<br/>Strong codes are used for Pin authentication flows<br/>Non-Strong codes are used for `Plex.tv/link`<br/>     |
| `opts`                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                              | :heavy_minus_sign:                                                                                                                                                    | The options for this request.                                                                                                                                         |


### Response

**[*operations.GetPinResponse](../../models/operations/getpinresponse.md), error**
| Error Object                 | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.GetPinResponseBody | 400                          | application/json             |
| sdkerrors.SDKError           | 4xx-5xx                      | */*                          |

## GetToken

Retrieve an Access Token from Plex.tv after the Pin has already been authenticated

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"context"
	"log"
	"net/http"
)

func main() {
    s := plexgo.New()


    var pinID string = "<value>"

    var xPlexClientIdentifier string = "<value>"

    ctx := context.Background()
    res, err := s.Plex.GetToken(ctx, pinID, xPlexClientIdentifier)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                             | Type                                                                                                                                                                  | Required                                                                                                                                                              | Description                                                                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                 | :heavy_check_mark:                                                                                                                                                    | The context to use for the request.                                                                                                                                   |
| `pinID`                                                                                                                                                               | *string*                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                    | The PinID to retrieve an access token for                                                                                                                             |
| `xPlexClientIdentifier`                                                                                                                                               | *string*                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                    | The unique identifier for the client application<br/>This is used to track the client application and its usage<br/>(UUID, serial number, or other number unique per device)<br/> |
| `opts`                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                              | :heavy_minus_sign:                                                                                                                                                    | The options for this request.                                                                                                                                         |


### Response

**[*operations.GetTokenResponse](../../models/operations/gettokenresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.GetTokenResponseBody | 400                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
