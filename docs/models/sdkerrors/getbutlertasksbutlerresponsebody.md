# GetButlerTasksButlerResponseBody

Unauthorized - Returned if the X-Plex-Token is missing from the header or query.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `Errors`                                                                                       | [][sdkerrors.GetButlerTasksButlerErrors](../../models/sdkerrors/getbutlertasksbutlererrors.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |
| `RawResponse`                                                                                  | [*http.Response](https://pkg.go.dev/net/http#Response)                                         | :heavy_minus_sign:                                                                             | Raw HTTP response; suitable for custom response parsing                                        |