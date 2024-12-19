# SesConfig

**Properties**

| Name      | Type                       | Required | Description                                      |
| :-------- | :------------------------- | :------- | :----------------------------------------------- |
| KeyId     | string                     | ✅       | AWS Access Key ID                                |
| Region    | string                     | ✅       | AWS Region                                       |
| SecretKey | string                     | ✅       | AWS Secret Key                                   |
| Endpoint  | integrations.Endpoint      | ❌       | HTTP endpoint to send requests to (testing only) |
| From      | integrations.SesConfigFrom | ❌       |                                                  |

# Endpoint

HTTP endpoint to send requests to (testing only)

# SesConfigFrom

**Properties**

| Name  | Type   | Required | Description                    |
| :---- | :----- | :------- | :----------------------------- |
| Email | string | ✅       | The email address to send from |
| Name  | string | ❌       | The name to send from          |
