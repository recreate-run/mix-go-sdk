# GetSessionFileRequest


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ID`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | Session ID                                                            |
| `Filename`                                                            | *string*                                                              | :heavy_check_mark:                                                    | Filename to retrieve                                                  |
| `Thumb`                                                               | **string*                                                             | :heavy_minus_sign:                                                    | Thumbnail specification: '100' (box), 'w100' (width), 'h100' (height) |
| `Time`                                                                | **float64*                                                            | :heavy_minus_sign:                                                    | Time offset in seconds for video thumbnails (default: 1.0)            |