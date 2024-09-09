# Log
(*Log*)

## Overview

Submit logs to the Log Handler for Plex Media Server


### Available Operations

* [LogLine](#logline) - Logging a single line message.
* [LogMultiLine](#logmultiline) - Logging a multi-line message
* [EnablePaperTrail](#enablepapertrail) - Enabling Papertrail

## LogLine

This endpoint will write a single-line log message, including a level and source to the main Plex Media Server log.


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
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Log.LogLine(ctx, operations.LevelThree, "Test log message", "Postman")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                     | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   | Example                                                                                                       |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                         | :heavy_check_mark:                                                                                            | The context to use for the request.                                                                           |                                                                                                               |
| `level`                                                                                                       | [operations.Level](../../models/operations/level.md)                                                          | :heavy_check_mark:                                                                                            | An integer log level to write to the PMS log with.  <br/>0: Error  <br/>1: Warning  <br/>2: Info  <br/>3: Debug  <br/>4: Verbose<br/> |                                                                                                               |
| `message`                                                                                                     | *string*                                                                                                      | :heavy_check_mark:                                                                                            | The text of the message to write to the log.                                                                  | Test log message                                                                                              |
| `source`                                                                                                      | *string*                                                                                                      | :heavy_check_mark:                                                                                            | a string indicating the source of the message.                                                                | Postman                                                                                                       |
| `opts`                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                      | :heavy_minus_sign:                                                                                            | The options for this request.                                                                                 |                                                                                                               |

### Response

**[*operations.LogLineResponse](../../models/operations/loglineresponse.md), error**

### Errors

| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.LogLineResponseBody    | 400                              | application/json                 |
| sdkerrors.LogLineLogResponseBody | 401                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |


## LogMultiLine

This endpoint allows for the batch addition of log entries to the main Plex Media Server log.  
It accepts a text/plain request body, where each line represents a distinct log entry.  
Each log entry consists of URL-encoded key-value pairs, specifying log attributes such as 'level', 'message', and 'source'.  

Log entries are separated by a newline character (`\n`).  
Each entry's parameters should be URL-encoded to ensure accurate parsing and handling of special characters.  
This method is efficient for logging multiple entries in a single API call, reducing the overhead of multiple individual requests.  

The 'level' parameter specifies the log entry's severity or importance, with the following integer values:
- `0`: Error - Critical issues that require immediate attention.
- `1`: Warning - Important events that are not critical but may indicate potential issues.
- `2`: Info - General informational messages about system operation.
- `3`: Debug - Detailed information useful for debugging purposes.
- `4`: Verbose - Highly detailed diagnostic information for in-depth analysis.

The 'message' parameter contains the log text, and 'source' identifies the log message's origin (e.g., an application name or module).

Example of a single log entry format:
`level=4&message=Sample%20log%20entry&source=applicationName`

Ensure each parameter is properly URL-encoded to avoid interpretation issues.


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
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Log.LogMultiLine(ctx, "level=4&message=Test%20message%201&source=postman\n" +
"level=3&message=Test%20message%202&source=postman\n" +
"level=1&message=Test%20message%203&source=postman")
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
| `request`                                                | [string](../../.md)                                      | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.LogMultiLineResponse](../../models/operations/logmultilineresponse.md), error**

### Errors

| Error Object                          | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.LogMultiLineResponseBody    | 400                                   | application/json                      |
| sdkerrors.LogMultiLineLogResponseBody | 401                                   | application/json                      |
| sdkerrors.SDKError                    | 4xx-5xx                               | */*                                   |


## EnablePaperTrail

This endpoint will enable all Plex Media Serverlogs to be sent to the Papertrail networked logging site for a period of time.


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
        plexgo.WithXPlexClientIdentifier("gcgzw5rz2xovp84b4vha3a40"),
    )

    ctx := context.Background()
    res, err := s.Log.EnablePaperTrail(ctx)
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

**[*operations.EnablePaperTrailResponse](../../models/operations/enablepapertrailresponse.md), error**

### Errors

| Error Object                              | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| sdkerrors.EnablePaperTrailResponseBody    | 400                                       | application/json                          |
| sdkerrors.EnablePaperTrailLogResponseBody | 401                                       | application/json                          |
| sdkerrors.SDKError                        | 4xx-5xx                                   | */*                                       |
