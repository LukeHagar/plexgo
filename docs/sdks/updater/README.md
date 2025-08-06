# Updater
(*Updater*)

## Overview

This describes the API for searching and applying updates to the Plex Media Server.
Updates to the status can be observed via the Event API.


### Available Operations

* [GetUpdateStatus](#getupdatestatus) - Querying status of updates
* [CheckForUpdates](#checkforupdates) - Checking for updates
* [ApplyUpdates](#applyupdates) - Apply Updates

## GetUpdateStatus

Querying status of updates

### Example Usage

<!-- UsageSnippet language="go" operationID="getUpdateStatus" method="get" path="/updater/status" -->
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

    res, err := s.Updater.GetUpdateStatus(ctx)
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

**[*operations.GetUpdateStatusResponse](../../models/operations/getupdatestatusresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.GetUpdateStatusBadRequest   | 400                                   | application/json                      |
| sdkerrors.GetUpdateStatusUnauthorized | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## CheckForUpdates

Checking for updates

### Example Usage

<!-- UsageSnippet language="go" operationID="checkForUpdates" method="put" path="/updater/check" -->
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

    res, err := s.Updater.CheckForUpdates(ctx, operations.DownloadOne.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |                                                             |
| `download`                                                  | [*operations.Download](../../models/operations/download.md) | :heavy_minus_sign:                                          | Indicate that you want to start download any updates found. | 1                                                           |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |                                                             |

### Response

**[*operations.CheckForUpdatesResponse](../../models/operations/checkforupdatesresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.CheckForUpdatesBadRequest   | 400                                   | application/json                      |
| sdkerrors.CheckForUpdatesUnauthorized | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## ApplyUpdates

Note that these two parameters are effectively mutually exclusive. The `tonight` parameter takes precedence and `skip` will be ignored if `tonight` is also passed


### Example Usage

<!-- UsageSnippet language="go" operationID="applyUpdates" method="put" path="/updater/apply" -->
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

    res, err := s.Updater.ApplyUpdates(ctx, operations.TonightOne.ToPointer(), operations.SkipOne.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              | Example                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |                                                                                                                                                          |
| `tonight`                                                                                                                                                | [*operations.Tonight](../../models/operations/tonight.md)                                                                                                | :heavy_minus_sign:                                                                                                                                       | Indicate that you want the update to run during the next Butler execution. Omitting this or setting it to false indicates that the update should install | 1                                                                                                                                                        |
| `skip`                                                                                                                                                   | [*operations.Skip](../../models/operations/skip.md)                                                                                                      | :heavy_minus_sign:                                                                                                                                       | Indicate that the latest version should be marked as skipped. The [Release] entry for this version will have the `state` set to `skipped`.               | 1                                                                                                                                                        |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |                                                                                                                                                          |

### Response

**[*operations.ApplyUpdatesResponse](../../models/operations/applyupdatesresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.ApplyUpdatesBadRequest   | 400                                | application/json                   |
| sdkerrors.ApplyUpdatesUnauthorized | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |