# FileInfo


## Fields

| Field                               | Type                                | Required                            | Description                         |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `IsDir`                             | *bool*                              | :heavy_check_mark:                  | Whether this is a directory         |
| `Modified`                          | *int64*                             | :heavy_check_mark:                  | Last modified timestamp (Unix time) |
| `Name`                              | *string*                            | :heavy_check_mark:                  | File name                           |
| `Size`                              | *int64*                             | :heavy_check_mark:                  | File size in bytes                  |
| `URL`                               | *string*                            | :heavy_check_mark:                  | Static URL to access the file       |