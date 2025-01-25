# Authentication
(*Authentication*)

## Overview

API Calls regarding authentication for Plex Media Server


### Available Operations

* [GetTransientToken](#gettransienttoken) - Get a Transient Token
* [GetSourceConnectionInformation](#getsourceconnectioninformation) - Get Source Connection Information
* [GetTokenDetails](#gettokendetails) - Get Token Details
* [PostUsersSignInData](#postuserssignindata) - Get User Sign In Data

## GetTransientToken

This endpoint provides the caller with a temporary token with the same access level as the caller's token. These tokens are valid for up to 48 hours and are destroyed if the server instance is restarted.


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
    
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Authentication.GetTransientToken(ctx, operations.GetTransientTokenQueryParamTypeDelegation, operations.ScopeAll)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `type_`                                                                                                  | [operations.GetTransientTokenQueryParamType](../../models/operations/gettransienttokenqueryparamtype.md) | :heavy_check_mark:                                                                                       | `delegation` - This is the only supported `type` parameter.                                              |
| `scope`                                                                                                  | [operations.Scope](../../models/operations/scope.md)                                                     | :heavy_check_mark:                                                                                       | `all` - This is the only supported `scope` parameter.                                                    |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetTransientTokenResponse](../../models/operations/gettransienttokenresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.GetTransientTokenBadRequest   | 400                                     | application/json                        |
| sdkerrors.GetTransientTokenUnauthorized | 401                                     | application/json                        |
| sdkerrors.SDKError                      | 4XX, 5XX                                | \*/\*                                   |

## GetSourceConnectionInformation

If a caller requires connection details and a transient token for a source that is known to the server, for example a cloud media provider or shared PMS, then this endpoint can be called. This endpoint is only accessible with either an admin token or a valid transient token generated from an admin token.
Note: requires Plex Media Server >= 1.15.4.


### Example Usage

```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Authentication.GetSourceConnectionInformation(ctx, "provider://provider-identifier")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `source`                                                 | *string*                                                 | :heavy_check_mark:                                       | The source identifier with an included prefix.           | server://client-identifier                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetSourceConnectionInformationResponse](../../models/operations/getsourceconnectioninformationresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.GetSourceConnectionInformationBadRequest   | 400                                                  | application/json                                     |
| sdkerrors.GetSourceConnectionInformationUnauthorized | 401                                                  | application/json                                     |
| sdkerrors.SDKError                                   | 4XX, 5XX                                             | \*/\*                                                |

## GetTokenDetails

Get the User data from the provided X-Plex-Token

### Example Usage

```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Authentication.GetTokenDetails(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.UserPlexAccount != nil {
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

**[*operations.GetTokenDetailsResponse](../../models/operations/gettokendetailsresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.GetTokenDetailsBadRequest   | 400                                   | application/json                      |
| sdkerrors.GetTokenDetailsUnauthorized | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## PostUsersSignInData

Sign in user with username and password and return user data with Plex authentication token

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

    res, err := s.Authentication.PostUsersSignInData(ctx, operations.PostUsersSignInDataRequest{
        ClientID: "3381b62b-9ab7-4e37-827b-203e9809eb58",
        ClientName: plexgo.String("Plex for Roku"),
        DeviceNickname: plexgo.String("Roku 3"),
        ClientVersion: plexgo.String("2.4.1"),
        Platform: plexgo.String("Roku"),
        RequestBody: &operations.PostUsersSignInDataRequestBody{
            Login: "username@email.com",
            Password: "password123",
            RememberMe: plexgo.Bool(false),
            VerificationCode: plexgo.String("123456"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserPlexAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PostUsersSignInDataRequest](../../models/operations/postuserssignindatarequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PostUsersSignInDataResponse](../../models/operations/postuserssignindataresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.PostUsersSignInDataBadRequest   | 400                                       | application/json                          |
| sdkerrors.PostUsersSignInDataUnauthorized | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4XX, 5XX                                  | \*/\*                                     |