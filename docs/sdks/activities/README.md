# Activities
(*Activities*)

## Overview

Activities are awesome. They provide a way to monitor and control asynchronous operations on the server. In order to receive real-time updates for activities, a client would normally subscribe via either EventSource or Websocket endpoints.
Activities are associated with HTTP replies via a special `X-Plex-Activity` header which contains the UUID of the activity.
Activities are optional cancellable. If cancellable, they may be cancelled via the `DELETE` endpoint. Other details:
- They can contain a `progress` (from 0 to 100) marking the percent completion of the activity.
- They must contain an `type` which is used by clients to distinguish the specific activity.
- They may contain a `Context` object with attributes which associate the activity with various specific entities (items, libraries, etc.)
- The may contain a `Response` object which attributes which represent the result of the asynchronous operation.


### Available Operations

* [GetServerActivities](#getserveractivities) - Get Server Activities
* [CancelServerActivities](#cancelserveractivities) - Cancel Server Activities

## GetServerActivities

Get Server Activities

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
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Activities.GetServerActivities(ctx)
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

**[*operations.GetServerActivitiesResponse](../../models/operations/getserveractivitiesresponse.md), error**

### Errors

| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.GetServerActivitiesBadRequest   | 400                                       | application/json                          |
| sdkerrors.GetServerActivitiesUnauthorized | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |


## CancelServerActivities

Cancel Server Activities

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
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Activities.CancelServerActivities(ctx, "25b71ed5-0f9d-461c-baa7-d404e9e10d3e")
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
| `activityUUID`                                           | *string*                                                 | :heavy_check_mark:                                       | The UUID of the activity to cancel.                      | 25b71ed5-0f9d-461c-baa7-d404e9e10d3e                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.CancelServerActivitiesResponse](../../models/operations/cancelserveractivitiesresponse.md), error**

### Errors

| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.CancelServerActivitiesBadRequest   | 400                                          | application/json                             |
| sdkerrors.CancelServerActivitiesUnauthorized | 401                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |
