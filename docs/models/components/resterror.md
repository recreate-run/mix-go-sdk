# RESTError


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Code`                                                               | *int64*                                                              | :heavy_check_mark:                                                   | HTTP status code                                                     |
| `Message`                                                            | *string*                                                             | :heavy_check_mark:                                                   | Error message                                                        |
| `Type`                                                               | [components.RESTErrorType](../../models/components/resterrortype.md) | :heavy_check_mark:                                                   | Error type                                                           |