# Statistics
(*Statistics*)

## Overview

API Calls that perform operations with Plex Media Server Statistics


### Available Operations

* [GetStatistics](#getstatistics) - Get Media Statistics
* [GetResourcesStatistics](#getresourcesstatistics) - Get Resources Statistics
* [GetBandwidthStatistics](#getbandwidthstatistics) - Get Bandwidth Statistics

## GetStatistics

This will return the media statistics for the server

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
    var timespan *int64 = plexgo.Int64(4)
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

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `timespan`                                                                                | **int64*                                                                                  | :heavy_minus_sign:                                                                        | The timespan to retrieve statistics for<br/>the exact meaning of this parameter is not known<br/> | 4                                                                                         |


### Response

**[*operations.GetStatisticsResponse](../../models/operations/getstatisticsresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.GetStatisticsResponseBody | 401                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## GetResourcesStatistics

This will return the resources for the server

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
    var timespan *int64 = plexgo.Int64(4)
    ctx := context.Background()
    res, err := s.Statistics.GetResourcesStatistics(ctx, timespan)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `timespan`                                                                                | **int64*                                                                                  | :heavy_minus_sign:                                                                        | The timespan to retrieve statistics for<br/>the exact meaning of this parameter is not known<br/> | 4                                                                                         |


### Response

**[*operations.GetResourcesStatisticsResponse](../../models/operations/getresourcesstatisticsresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.GetResourcesStatisticsResponseBody | 401                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |

## GetBandwidthStatistics

This will return the bandwidth statistics for the server

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
    var timespan *int64 = plexgo.Int64(4)
    ctx := context.Background()
    res, err := s.Statistics.GetBandwidthStatistics(ctx, timespan)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `timespan`                                                                                | **int64*                                                                                  | :heavy_minus_sign:                                                                        | The timespan to retrieve statistics for<br/>the exact meaning of this parameter is not known<br/> | 4                                                                                         |


### Response

**[*operations.GetBandwidthStatisticsResponse](../../models/operations/getbandwidthstatisticsresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.GetBandwidthStatisticsResponseBody | 401                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |
