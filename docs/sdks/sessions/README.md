# Sessions
(*Sessions*)

## Overview

API Calls that perform search operations with Plex Media Server Sessions


### Available Operations

* [GetSessions](#getsessions) - Get Active Sessions
* [GetSessionHistory](#getsessionhistory) - Get Session History
* [GetTranscodeSessions](#gettranscodesessions) - Get Transcode Sessions
* [StopTranscodeSession](#stoptranscodesession) - Stop a Transcode Session

## GetSessions

This will retrieve the "Now Playing" Information of the PMS.

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
    res, err := s.Sessions.GetSessions(ctx)
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

**[*operations.GetSessionsResponse](../../models/operations/getsessionsresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetSessionsBadRequest   | 400                               | application/json                  |
| sdkerrors.GetSessionsUnauthorized | 401                               | application/json                  |
| sdkerrors.SDKError                | 4XX, 5XX                          | \*/\*                             |

## GetSessionHistory

This will Retrieve a listing of all history views.

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
    res, err := s.Sessions.GetSessionHistory(ctx, plexgo.String("viewedAt:desc"), plexgo.Int64(1), &operations.QueryParamFilter{}, plexgo.Int64(12))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                     | Type                                                                                                                                                                                          | Required                                                                                                                                                                                      | Description                                                                                                                                                                                   | Example                                                                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                            | The context to use for the request.                                                                                                                                                           |                                                                                                                                                                                               |
| `sort`                                                                                                                                                                                        | **string*                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                            | Sorts the results by the specified field followed by the direction (asc, desc)<br/>                                                                                                           |                                                                                                                                                                                               |
| `accountID`                                                                                                                                                                                   | **int64*                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                            | Filter results by those that are related to a specific users id<br/>                                                                                                                          | 1                                                                                                                                                                                             |
| `filter`                                                                                                                                                                                      | [*operations.QueryParamFilter](../../models/operations/queryparamfilter.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                            | Filters content by field and direction/equality<br/>(Unknown if viewedAt is the only supported column)<br/>                                                                                   | {<br/>"viewed-at-greater-than": {<br/>"value": "viewedAt\u003e"<br/>},<br/>"viewed-at-greater-than-or-equal-to": {<br/>"value": "viewedAt\u003e=\u003e"<br/>},<br/>"viewed-at-less-than": {<br/>"value": "viewedAt\u003c"<br/>}<br/>} |
| `librarySectionID`                                                                                                                                                                            | **int64*                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                            | Filters the results based on the id of a valid library section<br/>                                                                                                                           | 12                                                                                                                                                                                            |
| `opts`                                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                            | The options for this request.                                                                                                                                                                 |                                                                                                                                                                                               |

### Response

**[*operations.GetSessionHistoryResponse](../../models/operations/getsessionhistoryresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.GetSessionHistoryBadRequest   | 400                                     | application/json                        |
| sdkerrors.GetSessionHistoryUnauthorized | 401                                     | application/json                        |
| sdkerrors.SDKError                      | 4XX, 5XX                                | \*/\*                                   |

## GetTranscodeSessions

Get Transcode Sessions

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
    res, err := s.Sessions.GetTranscodeSessions(ctx)
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

**[*operations.GetTranscodeSessionsResponse](../../models/operations/gettranscodesessionsresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.GetTranscodeSessionsBadRequest   | 400                                        | application/json                           |
| sdkerrors.GetTranscodeSessionsUnauthorized | 401                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |

## StopTranscodeSession

Stop a Transcode Session

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
    res, err := s.Sessions.StopTranscodeSession(ctx, "zz7llzqlx8w9vnrsbnwhbmep")
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
| `sessionKey`                                             | *string*                                                 | :heavy_check_mark:                                       | the Key of the transcode session to stop                 | zz7llzqlx8w9vnrsbnwhbmep                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.StopTranscodeSessionResponse](../../models/operations/stoptranscodesessionresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.StopTranscodeSessionBadRequest   | 400                                        | application/json                           |
| sdkerrors.StopTranscodeSessionUnauthorized | 401                                        | application/json                           |
| sdkerrors.SDKError                         | 4XX, 5XX                                   | \*/\*                                      |