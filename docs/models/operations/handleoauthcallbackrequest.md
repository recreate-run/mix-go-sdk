# HandleOAuthCallbackRequest


## Fields

| Field                                  | Type                                   | Required                               | Description                            |
| -------------------------------------- | -------------------------------------- | -------------------------------------- | -------------------------------------- |
| `Code`                                 | *string*                               | :heavy_check_mark:                     | Authorization code from OAuth provider |
| `Provider`                             | *string*                               | :heavy_check_mark:                     | Provider name (anthropic)              |
| `State`                                | *string*                               | :heavy_check_mark:                     | OAuth state for verification           |