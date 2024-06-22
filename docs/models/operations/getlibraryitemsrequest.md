# GetLibraryItemsRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `SectionID`                                           | *any*                                                 | :heavy_check_mark:                                    | the Id of the library to query                        |                                                       |
| `Tag`                                                 | [operations.Tag](../../models/operations/tag.md)      | :heavy_check_mark:                                    | A key representing a specific tag within the section. |                                                       |
| `IncludeGuids`                                        | **int64*                                              | :heavy_minus_sign:                                    | Adds the Guids object to the response<br/>            | 1                                                     |