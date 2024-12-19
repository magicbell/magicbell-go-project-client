# SlackToken

**Properties**

| Name    | Type                       | Required | Description |
| :------ | :------------------------- | :------- | :---------- |
| Oauth   | channels.Oauth             | ❌       |             |
| Webhook | channels.SlackTokenWebhook | ❌       |             |

# Oauth

**Properties**

| Name           | Type   | Required | Description |
| :------------- | :----- | :------- | :---------- |
| ChannelId      | string | ✅       |             |
| InstallationId | string | ✅       |             |
| Scope          | string | ❌       |             |

# SlackTokenWebhook

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| Url  | string | ✅       |             |
