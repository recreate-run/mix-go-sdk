# SendMessageResponseBody

Message accepted for processing. Agent runs asynchronously and streams results via SSE.


## Fields

| Field                              | Type                               | Required                           | Description                        | Example                            |
| ---------------------------------- | ---------------------------------- | ---------------------------------- | ---------------------------------- | ---------------------------------- |
| `SessionID`                        | *string*                           | :heavy_check_mark:                 | Session ID for the processing task |                                    |
| `Status`                           | *string*                           | :heavy_check_mark:                 | Processing status                  | processing                         |