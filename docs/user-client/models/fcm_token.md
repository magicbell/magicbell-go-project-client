# FcmToken

**Properties**

| Name           | Type                            | Required | Description |
| :------------- | :------------------------------ | :------- | :---------- |
| CreatedAt      | string                          | ✅       |             |
| DeviceToken    | string                          | ✅       |             |
| Id             | string                          | ✅       |             |
| DiscardedAt    | string                          | ❌       |             |
| InstallationId | channels.FcmTokenInstallationId | ❌       |             |
| UpdatedAt      | string                          | ❌       |             |

# FcmTokenInstallationId

**Properties**

| Name        | Type   | Required | Description   |
| :---------- | :----- | :------- | :------------ |
| development | string | ✅       | "development" |
| production  | string | ✅       | "production"  |
