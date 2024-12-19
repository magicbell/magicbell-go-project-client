# SendgridConfig

**Properties**

| Name    | Type                            | Required | Description              |
| :------ | :------------------------------ | :------- | :----------------------- |
| ApiKey  | string                          | ✅       | The API key for Sendgrid |
| From    | integrations.SendgridConfigFrom | ❌       |                          |
| ReplyTo | integrations.ReplyTo            | ❌       |                          |

# SendgridConfigFrom

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
