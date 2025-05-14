# MailgunConfigPayload

**Properties**

| Name   | Type                                  | Required | Description |
| :----- | :------------------------------------ | :------- | :---------- |
| ApiKey | string                                | ✅       |             |
| Domain | string                                | ✅       |             |
| Region | integrations.Region                   | ✅       |             |
| From   | integrations.MailgunConfigPayloadFrom | ❌       |             |

# Region

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| us   | string | ✅       | "us"        |
| eu   | string | ✅       | "eu"        |

# MailgunConfigPayloadFrom

**Properties**

| Name  | Type   | Required | Description                    |
| :---- | :----- | :------- | :----------------------------- |
| Email | string | ✅       | The email address to send from |
| Name  | string | ❌       | The name to send from          |
