# SlackInstallation

**Properties**

| Name                | Type                         | Required | Description |
| :------------------ | :--------------------------- | :------- | :---------- |
| AccessToken         | string                       | ✅       |             |
| AppId               | string                       | ✅       |             |
| AuthedUser          | integrations.AuthedUser      | ✅       |             |
| Team                | integrations.Team            | ✅       |             |
| BotUserId           | string                       | ❌       |             |
| Enterprise          | integrations.Enterprise      | ❌       |             |
| ExpiresIn           | int64                        | ❌       |             |
| Id                  | string                       | ❌       |             |
| IncomingWebhook     | integrations.IncomingWebhook | ❌       |             |
| IsEnterpriseInstall | bool                         | ❌       |             |
| RefreshToken        | string                       | ❌       |             |
| Scope               | string                       | ❌       |             |
| TokenType           | string                       | ❌       |             |

# AuthedUser

**Properties**

| Name         | Type   | Required | Description |
| :----------- | :----- | :------- | :---------- |
| Id           | string | ✅       |             |
| AccessToken  | string | ❌       |             |
| ExpiresIn    | int64  | ❌       |             |
| RefreshToken | string | ❌       |             |
| Scope        | string | ❌       |             |
| TokenType    | string | ❌       |             |

# Team

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| Id   | string | ✅       |             |
| Name | string | ❌       |             |

# Enterprise

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| Id   | string | ✅       |             |
| Name | string | ✅       |             |

# IncomingWebhook

**Properties**

| Name             | Type   | Required | Description |
| :--------------- | :----- | :------- | :---------- |
| Channel          | string | ✅       |             |
| ConfigurationUrl | string | ✅       |             |
| Url              | string | ✅       |             |
