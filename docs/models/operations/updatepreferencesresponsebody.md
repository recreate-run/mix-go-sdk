# UpdatePreferencesResponseBody

Updated preferences


## Fields

| Field                           | Type                            | Required                        | Description                     |
| ------------------------------- | ------------------------------- | ------------------------------- | ------------------------------- |
| `CreatedAt`                     | **int64*                        | :heavy_minus_sign:              | Creation timestamp              |
| `MainAgentMaxTokens`            | **int64*                        | :heavy_minus_sign:              | Maximum tokens for main agent   |
| `MainAgentModel`                | **string*                       | :heavy_minus_sign:              | Main agent model ID             |
| `MainAgentReasoningEffort`      | **string*                       | :heavy_minus_sign:              | Reasoning effort for main agent |
| `PreferredProvider`             | **string*                       | :heavy_minus_sign:              | Preferred AI provider           |
| `SubAgentMaxTokens`             | **int64*                        | :heavy_minus_sign:              | Maximum tokens for sub agent    |
| `SubAgentModel`                 | **string*                       | :heavy_minus_sign:              | Sub agent model ID              |
| `SubAgentReasoningEffort`       | **string*                       | :heavy_minus_sign:              | Reasoning effort for sub agent  |
| `UpdatedAt`                     | **int64*                        | :heavy_minus_sign:              | Last update timestamp           |