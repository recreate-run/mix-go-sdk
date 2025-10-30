# SendMessageResponse


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                                | [components.HTTPMetadata](../../models/components/httpmetadata.md)                        | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `Object`                                                                                  | [*operations.SendMessageResponseBody](../../models/operations/sendmessageresponsebody.md) | :heavy_minus_sign:                                                                        | Message accepted for processing. Agent runs asynchronously and streams results via SSE.   |