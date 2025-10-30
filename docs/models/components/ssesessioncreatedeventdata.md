# SSESessionCreatedEventData


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `CreatedAt`                                 | *int64*                                     | :heavy_check_mark:                          | Unix timestamp when the session was created |                                             |
| `SessionID`                                 | *string*                                    | :heavy_check_mark:                          | ID of the newly created session             |                                             |
| `Title`                                     | *string*                                    | :heavy_check_mark:                          | Title of the newly created session          |                                             |
| `Type`                                      | *string*                                    | :heavy_check_mark:                          | Event type                                  | session_created                             |