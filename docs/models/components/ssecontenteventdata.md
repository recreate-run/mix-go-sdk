# SSEContentEventData


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AssistantMessageID`                                                      | **string*                                                                 | :heavy_minus_sign:                                                        | ID of the assistant message this content belongs to                       |
| `Content`                                                                 | *string*                                                                  | :heavy_check_mark:                                                        | Streaming content delta                                                   |
| `ParentToolCallID`                                                        | **string*                                                                 | :heavy_minus_sign:                                                        | ID of the parent tool call that spawned this subagent (for nested events) |
| `Type`                                                                    | *string*                                                                  | :heavy_check_mark:                                                        | Content event type                                                        |