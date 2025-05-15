# FcmTokenPayload

**Properties**

| Name           | Type                                   | Required | Description |
| :------------- | :------------------------------------- | :------- | :---------- |
| DeviceToken    | string                                 | ✅       |             |
| InstallationId | channels.FcmTokenPayloadInstallationId | ❌       |             |

# FcmTokenPayloadInstallationId

**Properties**

| Name        | Type   | Required | Description   |
| :---------- | :----- | :------- | :------------ |
| development | string | ✅       | "development" |
| production  | string | ✅       | "production"  |
