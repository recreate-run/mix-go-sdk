# SSEThinkingEventData


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AssistantMessageID`                                                      | **string*                                                                 | :heavy_minus_sign:                                                        | ID of the assistant message this thinking belongs to                      |
| `Content`                                                                 | *string*                                                                  | :heavy_check_mark:                                                        | Thinking or reasoning content                                             |
| `ParentToolCallID`                                                        | **string*                                                                 | :heavy_minus_sign:                                                        | ID of the parent tool call that spawned this subagent (for nested events) |
| `Type`                                                                    | *string*                                                                  | :heavy_check_mark:                                                        | Thinking event type                                                       |