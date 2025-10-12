# Player

Information about the player being used for playback


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Title`                                                     | **string*                                                   | :heavy_minus_sign:                                          | The title of the client                                     |
| `Address`                                                   | **string*                                                   | :heavy_minus_sign:                                          | The remote address                                          |
| `Local`                                                     | **bool*                                                     | :heavy_minus_sign:                                          | Indicating if the client is playing from the local LAN      |
| `MachineIdentifier`                                         | **string*                                                   | :heavy_minus_sign:                                          | The identifier of the client                                |
| `Model`                                                     | **string*                                                   | :heavy_minus_sign:                                          | The model of the client                                     |
| `Platform`                                                  | **string*                                                   | :heavy_minus_sign:                                          | The platform of the client                                  |
| `PlatformVersion`                                           | **string*                                                   | :heavy_minus_sign:                                          | The platformVersion of the client                           |
| `Product`                                                   | **string*                                                   | :heavy_minus_sign:                                          | The product name of the client                              |
| `Relayed`                                                   | **bool*                                                     | :heavy_minus_sign:                                          | Indicating if the client is playing over a relay connection |
| `RemotePublicAddress`                                       | **string*                                                   | :heavy_minus_sign:                                          | The client's public address                                 |
| `Secure`                                                    | **bool*                                                     | :heavy_minus_sign:                                          | Indicating if the client is playing over HTTPS              |
| `State`                                                     | **string*                                                   | :heavy_minus_sign:                                          | The client's last reported state                            |
| `UserID`                                                    | **int64*                                                    | :heavy_minus_sign:                                          | The id of the user                                          |
| `Vendor`                                                    | **string*                                                   | :heavy_minus_sign:                                          | The vendor of the client                                    |
| `Version`                                                   | **string*                                                   | :heavy_minus_sign:                                          | The version of the client                                   |