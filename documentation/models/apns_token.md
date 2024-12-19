# ApnsToken

**Properties**

| Name           | Type                             | Required | Description                                                                                                                                                                       |
| :------------- | :------------------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| DeviceToken    | string                           | ✅       |                                                                                                                                                                                   |
| AppId          | string                           | ❌       | (Optional) The bundle identifier of the application that is registering this token. Use this field to override the default identifier specified in the projects APNs integration. |
| InstallationId | channels.ApnsTokenInstallationId | ❌       | (Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.                                                     |

# ApnsTokenInstallationId

(Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.

**Properties**

| Name        | Type   | Required | Description   |
| :---------- | :----- | :------- | :------------ |
| development | string | ✅       | "development" |
| production  | string | ✅       | "production"  |
