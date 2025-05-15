---
title: "user-client"
---

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Authentication

### Access Token Authentication

The UserClient API uses an Access Token for authentication.

This token must be provided to authenticate your requests to the API.

#### Setting the Access Token

When you initialize the SDK, you can set the access token as follows:

```go
import (
    "github.com/magicbell/magicbell-go/pkg/user-client/userclient"
    "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  )

config := userclientconfig.NewConfig()
config.SetAccessToken("YOUR-TOKEN")

sdk := userclient.NewUserClient(config)
```

If you need to set or update the access token after initializing the SDK, you can use:

```go
import (
    "github.com/magicbell/magicbell-go/pkg/user-client/userclient"
    "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  )

config := userclientconfig.NewConfig()

sdk := userclient.NewUserClient(config)
sdk.SetAccessToken("YOUR-TOKEN")
```

## Services

The SDK provides various services to interact with the API.

<details> 
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                                    |
| :---------------------------------------------------------------------- |
| [ChannelsService](services/channels_service.md)           |
| [IntegrationsService](services/integrations_service.md)   |
| [NotificationsService](services/notifications_service.md) |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `UserClientResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                         | Description                                 |
| :------- | :--------------------------- | :------------------------------------------ |
| Data     | `T`                          | The body of the API response                |
| Metadata | `UserClientResponseMetadata` | Status code and headers returned by the API |

#### `UserClientError`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                         | Description                                 |
| :------- | :--------------------------- | :------------------------------------------ |
| Err      | `error`                      | The error that occurred                     |
| Body     | `T`                          | The body of the API response                |
| Metadata | `UserClientResponseMetadata` | Status code and headers returned by the API |

#### `UserClientResponseMetadata`

This struct is shared by both response wrappers and contains the following fields:

| Name       | Type                | Description                                      |
| :--------- | :------------------ | :----------------------------------------------- |
| Headers    | `map[string]string` | A map containing the headers returned by the API |
| StatusCode | `int`               | The status code returned by the API              |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details> 
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                                             | Description |
| :----------------------------------------------------------------------------------------------- | :---------- |
| [InboxTokenResponseCollection](models/inbox_token_response_collection.md)          |             |
| [InboxToken](models/inbox_token.md)                                                |             |
| [InboxTokenResponse](models/inbox_token_response.md)                               |             |
| [DiscardResult](models/discard_result.md)                                          |             |
| [ApnsTokenCollection](models/apns_token_collection.md)                             |             |
| [ApnsTokenPayload](models/apns_token_payload.md)                                   |             |
| [ApnsToken](models/apns_token.md)                                                  |             |
| [ExpoTokenCollection](models/expo_token_collection.md)                             |             |
| [ExpoTokenPayload](models/expo_token_payload.md)                                   |             |
| [ExpoToken](models/expo_token.md)                                                  |             |
| [FcmTokenCollection](models/fcm_token_collection.md)                               |             |
| [FcmTokenPayload](models/fcm_token_payload.md)                                     |             |
| [FcmToken](models/fcm_token.md)                                                    |             |
| [SlackTokenCollection](models/slack_token_collection.md)                           |             |
| [SlackTokenPayload](models/slack_token_payload.md)                                 |             |
| [SlackToken](models/slack_token.md)                                                |             |
| [TeamsTokenCollection](models/teams_token_collection.md)                           |             |
| [TeamsTokenPayload](models/teams_token_payload.md)                                 |             |
| [TeamsToken](models/teams_token.md)                                                |             |
| [WebPushTokenCollection](models/web_push_token_collection.md)                      |             |
| [WebPushTokenPayload](models/web_push_token_payload.md)                            |             |
| [WebPushToken](models/web_push_token.md)                                           |             |
| [InboxConfigPayload](models/inbox_config_payload.md)                               |             |
| [SlackInstallation](models/slack_installation.md)                                  |             |
| [SlackFinishInstallResponse](models/slack_finish_install_response.md)              |             |
| [SlackStartInstall](models/slack_start_install.md)                                 |             |
| [SlackStartInstallResponseContent](models/slack_start_install_response_content.md) |             |
| [TemplatesInstallation](models/templates_installation.md)                          |             |
| [WebPushStartInstallationResponse](models/web_push_start_installation_response.md) |             |
| [NotificationCollection](models/notification_collection.md)                        |             |
| [Links](models/links.md)                                                           |             |
| [Notification](models/notification.md)                                             |             |

</details>
