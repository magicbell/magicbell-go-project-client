# SendgridConfigPayload

**Properties**

| Name    | Type                                   | Required | Description              |
| :------ | :------------------------------------- | :------- | :----------------------- |
| ApiKey  | string                                 | ✅       | The API key for Sendgrid |
| From    | integrations.SendgridConfigPayloadFrom | ❌       |                          |
| ReplyTo | integrations.ReplyTo                   | ❌       |                          |

# SendgridConfigPayloadFrom

**Properties**

| Name  | Type   | Required | Description                    |
| :---- | :----- | :------- | :----------------------------- |
| Email | string | ✅       | The email address to send from |
| Name  | string | ❌       | The name to send from          |

# ReplyTo

**Properties**

| Name  | Type   | Required | Description                   |
| :---- | :----- | :------- | :---------------------------- |
| Email | string | ✅       | The email address to reply to |
| Name  | string | ❌       | The name to reply to          |
