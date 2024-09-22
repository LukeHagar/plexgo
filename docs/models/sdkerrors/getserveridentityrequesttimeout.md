# GetServerIdentityRequestTimeout

Request Timeout


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Code`                                                  | **int*                                                  | :heavy_minus_sign:                                      | N/A                                                     | 408                                                     |
| `Message`                                               | **string*                                               | :heavy_minus_sign:                                      | N/A                                                     | The server timed out waiting for the request.           |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_minus_sign:                                      | Raw HTTP response; suitable for custom response parsing |                                                         |