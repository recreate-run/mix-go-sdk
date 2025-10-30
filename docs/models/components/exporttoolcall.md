# ExportToolCall

Complete tool call information for export


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Finished`                                                    | *bool*                                                        | :heavy_check_mark:                                            | Whether tool execution finished                               |
| `ID`                                                          | *string*                                                      | :heavy_check_mark:                                            | Tool call identifier                                          |
| `Input`                                                       | *string*                                                      | :heavy_check_mark:                                            | Tool input as JSON string                                     |
| `InputJSON`                                                   | [*components.InputJSON](../../models/components/inputjson.md) | :heavy_minus_sign:                                            | Parsed tool input (optional)                                  |
| `IsError`                                                     | **bool*                                                       | :heavy_minus_sign:                                            | Whether execution resulted in error (optional)                |
| `Metadata`                                                    | **string*                                                     | :heavy_minus_sign:                                            | Additional tool metadata (optional)                           |
| `Name`                                                        | *string*                                                      | :heavy_check_mark:                                            | Tool name                                                     |
| `Result`                                                      | **string*                                                     | :heavy_minus_sign:                                            | Tool execution result (optional)                              |
| `Type`                                                        | *string*                                                      | :heavy_check_mark:                                            | Tool type                                                     |