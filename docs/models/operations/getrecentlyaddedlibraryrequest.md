# GetRecentlyAddedLibraryRequest


## Fields

| Field                                                                                                                                                                                     | Type                                                                                                                                                                                      | Required                                                                                                                                                                                  | Description                                                                                                                                                                               | Example                                                                                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ContentDirectoryID`                                                                                                                                                                      | **int64*                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                        | N/A                                                                                                                                                                                       | 2                                                                                                                                                                                         |
| `PinnedContentDirectoryID`                                                                                                                                                                | []*int64*                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                        | N/A                                                                                                                                                                                       | [<br/>3,<br/>5,<br/>7,<br/>13,<br/>12,<br/>1,<br/>6,<br/>14,<br/>2,<br/>10,<br/>16,<br/>17<br/>]                                                                                          |
| `SectionID`                                                                                                                                                                               | **int64*                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                        | The library section ID for filtering content.                                                                                                                                             | 2                                                                                                                                                                                         |
| `Type`                                                                                                                                                                                    | [operations.QueryParamType](../../models/operations/queryparamtype.md)                                                                                                                    | :heavy_check_mark:                                                                                                                                                                        | The type of media to retrieve.<br/>1 = movie<br/>2 = show<br/>3 = season<br/>4 = episode<br/>E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries<br/> | 2                                                                                                                                                                                         |
| `IncludeMeta`                                                                                                                                                                             | [*operations.QueryParamIncludeMeta](../../models/operations/queryparamincludemeta.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                                                        | Adds the Meta object to the response<br/>                                                                                                                                                 | 1                                                                                                                                                                                         |
| `XPlexContainerStart`                                                                                                                                                                     | **int*                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                        | The index of the first item to return. If not specified, the first item will be returned.<br/>If the number of items exceeds the limit, the response will be paginated.<br/>By default this is 0<br/> | 0                                                                                                                                                                                         |
| `XPlexContainerSize`                                                                                                                                                                      | **int*                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                        | The number of items to return. If not specified, all items will be returned.<br/>If the number of items exceeds the limit, the response will be paginated.<br/>By default this is 50<br/> | 50                                                                                                                                                                                        |