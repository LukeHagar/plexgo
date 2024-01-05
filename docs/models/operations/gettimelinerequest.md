# GetTimelineRequest


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `RatingKey`                                          | *float64*                                            | :heavy_check_mark:                                   | The rating key of the media item                     |
| `Key`                                                | *string*                                             | :heavy_check_mark:                                   | The key of the media item to get the timeline for    |
| `State`                                              | [operations.State](../../models/operations/state.md) | :heavy_check_mark:                                   | The state of the media item                          |
| `HasMDE`                                             | *float64*                                            | :heavy_check_mark:                                   | Whether the media item has MDE                       |
| `Time`                                               | *float64*                                            | :heavy_check_mark:                                   | The time of the media item                           |
| `Duration`                                           | *float64*                                            | :heavy_check_mark:                                   | The duration of the media item                       |
| `Context`                                            | *string*                                             | :heavy_check_mark:                                   | The context of the media item                        |
| `PlayQueueItemID`                                    | *float64*                                            | :heavy_check_mark:                                   | The play queue item ID of the media item             |
| `PlayBackTime`                                       | *float64*                                            | :heavy_check_mark:                                   | The playback time of the media item                  |
| `Row`                                                | *float64*                                            | :heavy_check_mark:                                   | The row of the media item                            |