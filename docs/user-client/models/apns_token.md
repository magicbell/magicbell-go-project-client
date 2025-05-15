# ApnsToken

**Properties**

| Name           | Type                             | Required | Description                                                                                                                                                                       |
| :------------- | :------------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CreatedAt      | string                           | ✅       |                                                                                                                                                                                   |
| DeviceToken    | string                           | ✅       |                                                                                                                                                                                   |
| Id             | string                           | ✅       |                                                                                                                                                                                   |
| AppId          | string                           | ❌       | (Optional) The bundle identifier of the application that is registering this token. Use this field to override the default identifier specified in the projects APNs integration. |
| DiscardedAt    | string                           | ❌       |                                                                                                                                                                                   |
| InstallationId | channels.ApnsTokenInstallationId | ❌       | (Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.                                                     |
| UpdatedAt      | string                           | ❌       |                                                                                                                                                                                   |

# ApnsTokenInstallationId

(Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.

**Properties**

| Name        | Type   | Required | Description   |
| :---------- | :----- | :------- | :------------ |
| development | string | ✅       | "development" |
| production  | string | ✅       | "production"  |
