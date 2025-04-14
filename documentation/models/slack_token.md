# SlackToken

**Properties**

| Name        | Type                       | Required | Description                                                                                         |
| :---------- | :------------------------- | :------- | :-------------------------------------------------------------------------------------------------- |
| CreatedAt   | string                     | ✅       |                                                                                                     |
| Id          | string                     | ✅       |                                                                                                     |
| DiscardedAt | string                     | ❌       |                                                                                                     |
| Oauth       | channels.Oauth             | ❌       |                                                                                                     |
| UpdatedAt   | string                     | ❌       |                                                                                                     |
| Webhook     | channels.SlackTokenWebhook | ❌       | Obtained directly from the incoming_webhook object in the installation response from the Slack API. |

# Oauth

**Properties**

| Name           | Type   | Required | Description |
| :------------- | :----- | :------- | :---------- |
| ChannelId      | string | ✅       |             |
| InstallationId | string | ✅       |             |
| Scope          | string | ❌       |             |

# SlackTokenWebhook

Obtained directly from the incoming_webhook object in the installation response from the Slack API.

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| Url  | string | ✅       |             |
