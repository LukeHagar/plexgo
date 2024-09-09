# GetServerResourcesResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | HTTP response content type for this operation                       |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | HTTP response status code for this operation                        |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_check_mark:                                                  | Raw HTTP response; suitable for custom response parsing             |
| `PlexDevices`                                                       | [][operations.PlexDevice](../../models/operations/plexdevice.md)    | :heavy_minus_sign:                                                  | List of Plex Devices. This includes Plex hosted servers and clients |