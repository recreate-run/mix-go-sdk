# StreamEventsRequest


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `SessionID`                                              | *string*                                                 | :heavy_check_mark:                                       | Session ID to stream events for                          |
| `LastEventID`                                            | **string*                                                | :heavy_minus_sign:                                       | Last received event ID for reconnection and event replay |