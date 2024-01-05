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


    var level operations.Level = operations.LevelThree

    var message string = "string"

    var source string = "string"

    ctx := context.Background()
    res, err := s.Log.LogLine(ctx, level, message, source)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `level`                                                                                                      | [operations.Level](../../models/operations/level.md)                                                         | :heavy_check_mark:                                                                                           | An integer log level to write to the PMS log with.  <br/>0: Error  <br/>1: Warning  <br/>2: Info <br/>3: Debug  <br/>4: Verbose<br/> |                                                                                                              |
| `message`                                                                                                    | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The text of the message to write to the log.                                                                 |                                                                                                              |
| `source`                                                                                                     | *string*                                                                                                     | :heavy_check_mark:                                                                                           | a string indicating the source of the message.                                                               |                                                                                                              |


### Response

**[*operations.LogLineResponse](../../models/operations/loglineresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.LogLineResponseBody | 401                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## LogMultiLine

This endpoint will write multiple lines to the main Plex Media Server log in a single request. It takes a set of query strings as would normally sent to the above GET endpoint as a linefeed-separated block of POST data. The parameters for each query string match as above.


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

    ctx := context.Background()
    res, err := s.Log.LogMultiLine(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.LogMultiLineResponse](../../models/operations/logmultilineresponse.md), error**
| Error Object                       | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.LogMultiLineResponseBody | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4xx-5xx                            | */*                                |

## EnablePaperTrail

This endpoint will enable all Plex Media Serverlogs to be sent to the Papertrail networked logging site for a period of time.


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

    ctx := context.Background()
    res, err := s.Log.EnablePaperTrail(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.EnablePaperTrailResponse](../../models/operations/enablepapertrailresponse.md), error**
| Error Object                           | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.EnablePaperTrailResponseBody | 401                                    | application/json                       |
| sdkerrors.SDKError                     | 4xx-5xx                                | */*                                    |
