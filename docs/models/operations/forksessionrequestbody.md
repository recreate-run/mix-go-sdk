# ForkSessionRequestBody


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `MessageIndex`                                                       | *int64*                                                              | :heavy_check_mark:                                                   | Index of the last message to include in the fork (0-based)           |
| `Title`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | Optional title for the forked session (defaults to 'Forked Session') |