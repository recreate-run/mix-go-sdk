# GetPreferencesResponseBody

User preferences and available providers


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `AvailableProviders`                                                                      | map[string][operations.AvailableProviders](../../models/operations/availableproviders.md) | :heavy_check_mark:                                                                        | Map of available AI providers and their models                                            |
| `Preferences`                                                                             | [*components.UserPreferencesResponse](../../models/components/userpreferencesresponse.md) | :heavy_minus_sign:                                                                        | User preferences configuration                                                            |