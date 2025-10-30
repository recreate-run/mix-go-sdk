# RewindSessionRequestBody


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `CleanupMedia`                                                                           | **bool*                                                                                  | :heavy_minus_sign:                                                                       | Whether to clean up media files created after the rewind point (based on file timestamp) |
| `MessageID`                                                                              | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the last message to keep. All messages after this message will be deleted.         |