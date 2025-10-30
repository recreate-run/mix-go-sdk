# UpdatePreferencesRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `MainAgentMaxTokens`                                  | **int64*                                              | :heavy_minus_sign:                                    | Maximum tokens for main agent responses               |
| `MainAgentModel`                                      | **string*                                             | :heavy_minus_sign:                                    | Main agent model ID                                   |
| `MainAgentReasoningEffort`                            | **string*                                             | :heavy_minus_sign:                                    | Reasoning effort setting for main agent               |
| `PreferredProvider`                                   | **string*                                             | :heavy_minus_sign:                                    | Preferred AI provider (anthropic, openai, openrouter) |
| `SubAgentMaxTokens`                                   | **int64*                                              | :heavy_minus_sign:                                    | Maximum tokens for sub agent responses                |
| `SubAgentModel`                                       | **string*                                             | :heavy_minus_sign:                                    | Sub agent model ID                                    |
| `SubAgentReasoningEffort`                             | **string*                                             | :heavy_minus_sign:                                    | Reasoning effort setting for sub agent                |