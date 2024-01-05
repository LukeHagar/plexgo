# StartUniversalTranscodeRequest


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `HasMDE`                                      | *float64*                                     | :heavy_check_mark:                            | Whether the media item has MDE                |
| `Path`                                        | *string*                                      | :heavy_check_mark:                            | The path to the media item to transcode       |
| `MediaIndex`                                  | *float64*                                     | :heavy_check_mark:                            | The index of the media item to transcode      |
| `PartIndex`                                   | *float64*                                     | :heavy_check_mark:                            | The index of the part to transcode            |
| `Protocol`                                    | *string*                                      | :heavy_check_mark:                            | The protocol to use for the transcode session |
| `FastSeek`                                    | **float64*                                    | :heavy_minus_sign:                            | Whether to use fast seek or not               |
| `DirectPlay`                                  | **float64*                                    | :heavy_minus_sign:                            | Whether to use direct play or not             |
| `DirectStream`                                | **float64*                                    | :heavy_minus_sign:                            | Whether to use direct stream or not           |
| `SubtitleSize`                                | **float64*                                    | :heavy_minus_sign:                            | The size of the subtitles                     |
| `Subtites`                                    | **string*                                     | :heavy_minus_sign:                            | The subtitles                                 |
| `AudioBoost`                                  | **float64*                                    | :heavy_minus_sign:                            | The audio boost                               |
| `Location`                                    | **string*                                     | :heavy_minus_sign:                            | The location of the transcode session         |
| `MediaBufferSize`                             | **float64*                                    | :heavy_minus_sign:                            | The size of the media buffer                  |
| `Session`                                     | **string*                                     | :heavy_minus_sign:                            | The session ID                                |
| `AddDebugOverlay`                             | **float64*                                    | :heavy_minus_sign:                            | Whether to add a debug overlay or not         |
| `AutoAdjustQuality`                           | **float64*                                    | :heavy_minus_sign:                            | Whether to auto adjust quality or not         |