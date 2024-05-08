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
	"github.com/LukeHagar/plexgo/models/components"
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetButlerTasksResponse](../../models/operations/getbutlertasksresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.GetButlerTasksResponseBody | 401                                  | application/json                     |
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
	"github.com/LukeHagar/plexgo/models/components"
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.StartAllTasksResponse](../../models/operations/startalltasksresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.StartAllTasksResponseBody | 401                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## StopAllTasks

This endpoint will stop all currently running tasks and remove any scheduled tasks from the queue.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.StopAllTasksResponse](../../models/operations/stopalltasksresponse.md), error**
| Error Object                       | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.StopAllTasksResponseBody | 401                                | application/json                   |
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
	"github.com/LukeHagar/plexgo/models/components"
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

    var taskName operations.TaskName = operations.TaskNameCleanOldBundles
    
    ctx := context.Background()
    res, err := s.Butler.StartTask(ctx, taskName)
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


### Response

**[*operations.StartTaskResponse](../../models/operations/starttaskresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.StartTaskResponseBody | 401                             | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## StopTask

This endpoint will stop a currently running task by name, or remove it from the list of scheduled tasks if it exists. See the section above for a list of task names for this endpoint.


### Example Usage

```go
package main

import(
	"github.com/LukeHagar/plexgo/models/components"
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

    var taskName operations.PathParamTaskName = operations.PathParamTaskNameBackupDatabase
    
    ctx := context.Background()
    res, err := s.Butler.StopTask(ctx, taskName)
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


### Response

**[*operations.StopTaskResponse](../../models/operations/stoptaskresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.StopTaskResponseBody | 401                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
