# SesConfigPayload

**Properties**

| Name      | Type                              | Required | Description       |
| :-------- | :-------------------------------- | :------- | :---------------- |
| KeyId     | string                            | ✅       | AWS Access Key ID |
| Region    | string                            | ✅       | AWS Region        |
| SecretKey | string                            | ✅       | AWS Secret Key    |
| From      | integrations.SesConfigPayloadFrom | ❌       |                   |

# SesConfigPayloadFrom

**Properties**

| Name  | Type   | Required | Description                    |
| :---- | :----- | :------- | :----------------------------- |
| Email | string | ✅       | The email address to send from |
| Name  | string | ❌       | The name to send from          |
