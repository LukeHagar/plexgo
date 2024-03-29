# Statistics
(*Statistics*)

## Overview

API Calls that perform operations with Plex Media Server Statistics


### Available Operations

* [GetStatistics](#getstatistics) - Get Media Statistics

## GetStatistics

This will return the media statistics for the server

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
        plexgo.WithXPlexClientIdentifier("<value>"),
    )


    var timespan *int64 = plexgo.Int64(411769)

    ctx := context.Background()
    res, err := s.Statistics.GetStatistics(ctx, timespan)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `timespan`                                                                                | **int64*                                                                                  | :heavy_minus_sign:                                                                        | The timespan to retrieve statistics for<br/>the exact meaning of this parameter is not known<br/> |


### Response

**[*operations.GetStatisticsResponse](../../models/operations/getstatisticsresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.GetStatisticsResponseBody | 401                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |
