# GetThumbImageRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `RatingKey`                                           | *int64*                                               | :heavy_check_mark:                                    | the id of the library item to return the children of. | 9518                                                  |
| `Width`                                               | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   | 396                                                   |
| `Height`                                              | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   | 396                                                   |
| `MinSize`                                             | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   | 1                                                     |
| `Upscale`                                             | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   | 1                                                     |
| `XPlexToken`                                          | *string*                                              | :heavy_check_mark:                                    | An authentication token, obtained from plex.tv        | CV5xoxjTpFKUzBTShsaf                                  |