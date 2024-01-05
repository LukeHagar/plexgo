# Security
(*Security*)

## Overview

API Calls against Security for Plex Media Server


### Available Operations

* [GetTransientToken](#gettransienttoken) - Get a Transient Token.
* [GetSourceConnectionInformation](#getsourceconnectioninformation) - Get Source Connection Information

## GetTransientToken

This endpoint provides the caller with a temporary token with the same access level as the caller's token. These tokens are valid for up to 48 hours and are destroyed if the server instance is restarted.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"context"
	"log"
	"net/http"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var type_ operations.QueryParamType = operations.QueryParamTypeDelegation

    var scope operations.Scope = operations.ScopeAll

    ctx := context.Background()
    res, err := s.Security.GetTransientToken(ctx, type_, scope)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `type_`                                                                | [operations.QueryParamType](../../models/operations/queryparamtype.md) | :heavy_check_mark:                                                     | `delegation` - This is the only supported `type` parameter.            |
| `scope`                                                                | [operations.Scope](../../models/operations/scope.md)                   | :heavy_check_mark:                                                     | `all` - This is the only supported `scope` parameter.                  |


### Response

**[*operations.GetTransientTokenResponse](../../models/operations/gettransienttokenresponse.md), error**
| Error Object                            | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.GetTransientTokenResponseBody | 401                                     | application/json                        |
| sdkerrors.SDKError                      | 4xx-5xx                                 | */*                                     |

## GetSourceConnectionInformation

If a caller requires connection details and a transient token for a source that is known to the server, for example a cloud media provider or shared PMS, then this endpoint can be called. This endpoint is only accessible with either an admin token or a valid transient token generated from an admin token.
Note: requires Plex Media Server >= 1.15.4.


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


    var source string = "server://client-identifier"

    ctx := context.Background()
    res, err := s.Security.GetSourceConnectionInformation(ctx, source)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `source`                                              | *string*                                              | :heavy_check_mark:                                    | The source identifier with an included prefix.        | server://client-identifier                            |


### Response

**[*operations.GetSourceConnectionInformationResponse](../../models/operations/getsourceconnectioninformationresponse.md), error**
| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.GetSourceConnectionInformationResponseBody | 401                                                  | application/json                                     |
| sdkerrors.SDKError                                   | 4xx-5xx                                              | */*                                                  |