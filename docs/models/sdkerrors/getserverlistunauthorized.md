# GetServerListUnauthorized

Unauthorized - Returned if the X-Plex-Token is missing from the header or query.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Errors`                                                                                     | [][sdkerrors.GetServerListServerErrors](../../models/sdkerrors/getserverlistservererrors.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `RawResponse`                                                                                | [*http.Response](https://pkg.go.dev/net/http#Response)                                       | :heavy_minus_sign:                                                                           | Raw HTTP response; suitable for custom response parsing                                      |