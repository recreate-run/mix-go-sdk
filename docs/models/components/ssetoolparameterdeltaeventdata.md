# SSEToolParameterDeltaEventData


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AssistantMessageID`                                                      | **string*                                                                 | :heavy_minus_sign:                                                        | ID of the assistant message this tool parameter delta belongs to          |
| `Input`                                                                   | *string*                                                                  | :heavy_check_mark:                                                        | Partial JSON parameter delta - may not be parseable until complete        |
| `ParentToolCallID`                                                        | **string*                                                                 | :heavy_minus_sign:                                                        | ID of the parent tool call that spawned this subagent (for nested events) |
| `ToolCallID`                                                              | *string*                                                                  | :heavy_check_mark:                                                        | Tool call identifier for correlation                                      |
| `Type`                                                                    | *string*                                                                  | :heavy_check_mark:                                                        | Tool parameter delta event type                                           |