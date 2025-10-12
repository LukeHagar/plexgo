# ListPlaybackHistoryMetadata


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `AccountID`                                       | **int64*                                          | :heavy_minus_sign:                                | The account id of this playback                   |
| `DeviceID`                                        | **int64*                                          | :heavy_minus_sign:                                | The device id which played the item               |
| `HistoryKey`                                      | **string*                                         | :heavy_minus_sign:                                | The key for this individual history item          |
| `Key`                                             | **string*                                         | :heavy_minus_sign:                                | The metadata key for the item played              |
| `LibrarySectionID`                                | **string*                                         | :heavy_minus_sign:                                | The library section id containing the item played |
| `OriginallyAvailableAt`                           | **string*                                         | :heavy_minus_sign:                                | The originally available at of the item played    |
| `RatingKey`                                       | **string*                                         | :heavy_minus_sign:                                | The rating key for the item played                |
| `Thumb`                                           | **string*                                         | :heavy_minus_sign:                                | The thumb of the item played                      |
| `Title`                                           | **string*                                         | :heavy_minus_sign:                                | The title of the item played                      |
| `Type`                                            | **string*                                         | :heavy_minus_sign:                                | The metadata type of the item played              |
| `ViewedAt`                                        | **int64*                                          | :heavy_minus_sign:                                | The time when the item was played                 |