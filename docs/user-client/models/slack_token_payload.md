# SlackTokenPayload

**Properties**

| Name    | Type                              | Required | Description                                                                                         |
| :------ | :-------------------------------- | :------- | :-------------------------------------------------------------------------------------------------- |
| Oauth   | channels.SlackTokenPayloadOauth   | ❌       |                                                                                                     |
| Webhook | channels.SlackTokenPayloadWebhook | ❌       | Obtained directly from the incoming_webhook object in the installation response from the Slack API. |

# SlackTokenPayloadOauth

**Properties**

| Name           | Type   | Required | Description |
| :------------- | :----- | :------- | :---------- |
| ChannelId      | string | ✅       |             |
| InstallationId | string | ✅       |             |
| Scope          | string | ❌       |             |

# SlackTokenPayloadWebhook

Obtained directly from the incoming_webhook object in the installation response from the Slack API.

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| Url  | string | ✅       |             |
