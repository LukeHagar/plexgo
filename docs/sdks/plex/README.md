# Plex
(*Plex*)

## Overview

### Available Operations

* [GetServerResources](#getserverresources) - Get Server Resources

## GetServerResources

Get Plex server access tokens and server connections

### Example Usage

<!-- UsageSnippet language="go" operationID="get-server-resources" method="get" path="/resources" -->
```go
package main

import(
	"context"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := plexgo.New(
        plexgo.WithAccepts(components.AcceptsApplicationXML),
        plexgo.WithClientIdentifier("3381b62b-9ab7-4e37-827b-203e9809eb58"),
        plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Plex.GetServerResources(ctx, operations.GetServerResourcesRequest{
        IncludeHTTPS: operations.IncludeHTTPSTrue.ToPointer(),
        IncludeRelay: operations.IncludeRelayTrue.ToPointer(),
        IncludeIPv6: operations.IncludeIPv6True.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlexDevices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetServerResourcesRequest](../../models/operations/getserverresourcesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetServerResourcesResponse](../../models/operations/getserverresourcesresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.GetServerResourcesUnauthorized | 401                                      | application/json                         |
| sdkerrors.SDKError                       | 4XX, 5XX                                 | \*/\*                                    |