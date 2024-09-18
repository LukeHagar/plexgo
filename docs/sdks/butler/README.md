# Butler
(*Butler*)

## Overview

Butler is the task manager of the Plex Media Server Ecosystem.


### Available Operations

* [GetButlerTasks](#getbutlertasks) - Get Butler tasks
* [StartAllTasks](#startalltasks) - Start all Butler tasks
* [StopAllTasks](#stopalltasks) - Stop all Butler tasks
* [StartTask](#starttask) - Start a single Butler task
* [StopTask](#stoptask) - Stop a single Butler task

## GetButlerTasks

Returns a list of butler tasks

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
    res, err := s.Butler.GetButlerTasks(ctx)
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

**[*operations.GetButlerTasksResponse](../../models/operations/getbutlertasksresponse.md), error**

### Errors

| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.GetButlerTasksBadRequest   | 400                                  | application/json                     |
| sdkerrors.GetButlerTasksUnauthorized | 401                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |


## StartAllTasks

This endpoint will attempt to start all Butler tasks that are enabled in the settings. Butler tasks normally run automatically during a time window configured on the server's Settings page but can be manually started using this endpoint. Tasks will run with the following criteria:
1. Any tasks not scheduled to run on the current day will be skipped.
2. If a task is configured to run at a random time during the configured window and we are outside that window, the task will start immediately.
3. If a task is configured to run at a random time during the configured window and we are within that window, the task will be scheduled at a random time within the window.
4. If we are outside the configured window, the task will start immediately.


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
    res, err := s.Butler.StartAllTasks(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.StartAllTasksResponse](../../models/operations/startalltasksresponse.md), error**

### Errors

| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.StartAllTasksBadRequest   | 400                                 | application/json                    |
| sdkerrors.StartAllTasksUnauthorized | 401                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |


## StopAllTasks

This endpoint will stop all currently running tasks and remove any scheduled tasks from the queue.


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
    res, err := s.Butler.StopAllTasks(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.StopAllTasksResponse](../../models/operations/stopalltasksresponse.md), error**

### Errors

| Error Object                       | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.StopAllTasksBadRequest   | 400                                | application/json                   |
| sdkerrors.StopAllTasksUnauthorized | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4xx-5xx                            | */*                                |


## StartTask

This endpoint will attempt to start a single Butler task that is enabled in the settings. Butler tasks normally run automatically during a time window configured on the server's Settings page but can be manually started using this endpoint. Tasks will run with the following criteria:
1. Any tasks not scheduled to run on the current day will be skipped.
2. If a task is configured to run at a random time during the configured window and we are outside that window, the task will start immediately.
3. If a task is configured to run at a random time during the configured window and we are within that window, the task will be scheduled at a random time within the window.
4. If we are outside the configured window, the task will start immediately.


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
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Butler.StartTask(ctx, operations.TaskNameCleanOldBundles)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `taskName`                                                 | [operations.TaskName](../../models/operations/taskname.md) | :heavy_check_mark:                                         | the name of the task to be started.                        |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.StartTaskResponse](../../models/operations/starttaskresponse.md), error**

### Errors

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.StartTaskBadRequest   | 400                             | application/json                |
| sdkerrors.StartTaskUnauthorized | 401                             | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |


## StopTask

This endpoint will stop a currently running task by name, or remove it from the list of scheduled tasks if it exists. See the section above for a list of task names for this endpoint.


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
        plexgo.WithClientID("gcgzw5rz2xovp84b4vha3a40"),
        plexgo.WithClientName("Plex Web"),
        plexgo.WithClientVersion("4.133.0"),
        plexgo.WithClientPlatform("Chrome"),
        plexgo.WithDeviceName("Linux"),
    )

    ctx := context.Background()
    res, err := s.Butler.StopTask(ctx, operations.PathParamTaskNameBackupDatabase)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `taskName`                                                                   | [operations.PathParamTaskName](../../models/operations/pathparamtaskname.md) | :heavy_check_mark:                                                           | The name of the task to be started.                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.StopTaskResponse](../../models/operations/stoptaskresponse.md), error**

### Errors

| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.StopTaskBadRequest   | 400                            | application/json               |
| sdkerrors.StopTaskUnauthorized | 401                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
