# GetOnDeckResponseBody

Bad Request - A parameter was not specified, or was specified incorrectly.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Errors`                                                                 | [][sdkerrors.GetOnDeckErrors](../../models/sdkerrors/getondeckerrors.md) | :heavy_minus_sign:                                                       | N/A                                                                      |
| `RawResponse`                                                            | [*http.Response](https://pkg.go.dev/net/http#Response)                   | :heavy_minus_sign:                                                       | Raw HTTP response; suitable for custom response parsing                  |