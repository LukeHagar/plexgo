# Users
(*Users*)

## Overview

### Available Operations

* [GetUsers](#getusers) - Get list of all connected users

## GetUsers

Get list of all users that are friends and have library access with the provided Plex authentication token

### Example Usage

```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New()

    res, err := s.Users.GetUsers(ctx, operations.GetUsersRequest{
        ClientID: "3381b62b-9ab7-4e37-827b-203e9809eb58",
        ClientName: plexgo.String("Plex for Roku"),
        DeviceNickname: plexgo.String("Roku 3"),
        DeviceName: plexgo.String("Chrome"),
        DeviceScreenResolution: plexgo.String("1487x1165,2560x1440"),
        ClientVersion: plexgo.String("2.4.1"),
        Platform: plexgo.String("Roku"),
        ClientFeatures: plexgo.String("external-media,indirect-media,hub-style-list"),
        Model: plexgo.String("4200X"),
        XPlexSessionID: plexgo.String("97e136ef-4ddd-4ff3-89a7-a5820c96c2ca"),
        XPlexLanguage: plexgo.String("en"),
        PlatformVersion: plexgo.String("4.3 build 1057"),
        XPlexToken: "CV5xoxjTpFKUzBTShsaf",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetUsersRequest](../../models/operations/getusersrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.GetUsersResponse](../../models/operations/getusersresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.GetUsersBadRequest   | 400                            | application/json               |
| sdkerrors.GetUsersUnauthorized | 401                            | application/json               |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |