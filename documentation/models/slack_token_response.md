# SlackTokenResponse

**Properties**

| Name        | Type                               | Required | Description |
| :---------- | :--------------------------------- | :------- | :---------- |
| CreatedAt   | string                             | ✅       |             |
| Id          | string                             | ✅       |             |
| DiscardedAt | string                             | ❌       |             |
| Oauth       | channels.Oauth                     | ❌       |             |
| UpdatedAt   | string                             | ❌       |             |
| Webhook     | channels.SlackTokenResponseWebhook | ❌       |             |

# Oauth

**Properties**

| Name           | Type   | Required | Description |
| :------------- | :----- | :------- | :---------- |
| ChannelId      | string | ✅       |             |
| InstallationId | string | ✅       |             |
| Scope          | string | ❌       |             |

# SlackTokenResponseWebhook

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| Url  | string | ✅       |             |
