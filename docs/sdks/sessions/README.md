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
        plexgo.WithXPlexClientIdentifier("Postman"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSessionsResponse](../../models/operations/getsessionsresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetSessionsResponseBody | 401                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## GetSessionHistory

This will Retrieve a listing of all history views.

### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"context"
	"log"
)

func main() {
    s := plexgo.New(
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        plexgo.WithXPlexClientIdentifier("Postman"),
    )
    var sort *string = plexgo.String("<value>")

    var accountID *int64 = plexgo.Int64(1)

    var filter *operations.Filter = &operations.Filter{}

    var librarySectionID *int64 = plexgo.Int64(12)
    ctx := context.Background()
    res, err := s.Sessions.GetSessionHistory(ctx, sort, accountID, filter, librarySectionID)
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
| `filter`                                                                                                                                                                                      | [*operations.Filter](../../models/operations/filter.md)                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                            | Filters content by field and direction/equality<br/>(Unknown if viewedAt is the only supported column)<br/>                                                                                   | {<br/>"viewed-at-greater-than": {<br/>"value": "viewedAt\u003e"<br/>},<br/>"viewed-at-greater-than-or-equal-to": {<br/>"value": "viewedAt\u003e=\u003e"<br/>},<br/>"viewed-at-less-than": {<br/>"value": "viewedAt\u003c"<br/>}<br/>} |
| `librarySectionID`                                                                                                                                                                            | **int64*                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                            | Filters the results based on the id of a valid library section<br/>                                                                                                                           | 12                                                                                                                                                                                            |


### Response

**[*operations.GetSessionHistoryResponse](../../models/operations/getsessionhistoryresponse.md), error**
| Error Object                            | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.GetSessionHistoryResponseBody | 401                                     | application/json                        |
| sdkerrors.SDKError                      | 4xx-5xx                                 | */*                                     |

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
        plexgo.WithXPlexClientIdentifier("Postman"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetTranscodeSessionsResponse](../../models/operations/gettranscodesessionsresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.GetTranscodeSessionsResponseBody | 401                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

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
        plexgo.WithXPlexClientIdentifier("Postman"),
    )
    var sessionKey string = "zz7llzqlx8w9vnrsbnwhbmep"
    ctx := context.Background()
    res, err := s.Sessions.StopTranscodeSession(ctx, sessionKey)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `sessionKey`                                          | *string*                                              | :heavy_check_mark:                                    | the Key of the transcode session to stop              | zz7llzqlx8w9vnrsbnwhbmep                              |


### Response

**[*operations.StopTranscodeSessionResponse](../../models/operations/stoptranscodesessionresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.StopTranscodeSessionResponseBody | 401                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |
