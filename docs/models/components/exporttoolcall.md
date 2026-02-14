# ExportToolCall

Complete tool call information for export


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Finished`                                                    | *bool*                                                        | :heavy_check_mark:                                            | Whether tool execution finished                               |
| `ID`                                                          | *string*                                                      | :heavy_check_mark:                                            | Tool call identifier                                          |
| `Input`                                                       | *string*                                                      | :heavy_check_mark:                                            | Tool input as JSON string                                     |
| `InputJSON`                                                   | [*components.InputJSON](../../models/components/inputjson.md) | :heavy_minus_sign:                                            | Parsed tool input (optional)                                  |
| `Name`                                                        | *string*                                                      | :heavy_check_mark:                                            | Tool name                                                     |
| `Result`                                                      | **string*                                                     | :heavy_minus_sign:                                            | Tool execution result (optional)                              |
| `ScreenshotUrls`                                              | []*string*                                                    | :heavy_minus_sign:                                            | Screenshot URLs captured during tool execution (optional)     |
| `Type`                                                        | *string*                                                      | :heavy_check_mark:                                            | Tool type                                                     |