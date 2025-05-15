# MagicBell Go SDK

Welcome to the MagiBell GO SDK documentation. This guide will help you get started with integrating and using the MagicBell GO SDK in your project.

# MagicBell GO ProjectClient

<!-- AUTO-GENERATED-CONTENT:START (project-client) -->

## Authentication

### Access Token Authentication

The ProjectClient API uses an Access Token for authentication.

This token must be provided to authenticate your requests to the API.

#### Setting the Access Token

When you initialize the SDK, you can set the access token as follows:

```go
import (
    "github.com/magicbell/magicbell-go/pkg/project-client/projectclient"
    "github.com/magicbell/magicbell-go/pkg/project-client/projectclientconfig"
  )

config := projectclientconfig.NewConfig()
config.SetAccessToken("YOUR-TOKEN")

sdk := projectclient.NewProjectClient(config)
```

If you need to set or update the access token after initializing the SDK, you can use:

```go
import (
    "github.com/magicbell/magicbell-go/pkg/project-client/projectclient"
    "github.com/magicbell/magicbell-go/pkg/project-client/projectclientconfig"
  )

config := projectclientconfig.NewConfig()

sdk := projectclient.NewProjectClient(config)
sdk.SetAccessToken("YOUR-TOKEN")
```

## Services

The SDK provides various services to interact with the API.

<details> 
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                                    |
| :---------------------------------------------------------------------- |
| [BroadcastsService](docs/project-client/services/broadcasts_service.md)       |
| [ChannelsService](docs/project-client/services/channels_service.md)           |
| [EventsService](docs/project-client/services/events_service.md)               |
| [IntegrationsService](docs/project-client/services/integrations_service.md)   |
| [JwtService](docs/project-client/services/jwt_service.md)                     |
| [NotificationsService](docs/project-client/services/notifications_service.md) |
| [UsersService](docs/project-client/services/users_service.md)                 |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `ProjectClientResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                            | Description                                 |
| :------- | :------------------------------ | :------------------------------------------ |
| Data     | `T`                             | The body of the API response                |
| Metadata | `ProjectClientResponseMetadata` | Status code and headers returned by the API |

#### `ProjectClientError`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                            | Description                                 |
| :------- | :------------------------------ | :------------------------------------------ |
| Err      | `error`                         | The error that occurred                     |
| Body     | `T`                             | The body of the API response                |
| Metadata | `ProjectClientResponseMetadata` | Status code and headers returned by the API |

#### `ProjectClientResponseMetadata`

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

| Name                                                                                    | Description |
| :-------------------------------------------------------------------------------------- | :---------- |
| [BroadcastCollection](docs/project-client/models/broadcast_collection.md)                     |             |
| [Broadcast](docs/project-client/models/broadcast.md)                                          |             |
| [CategoryDeliveryConfig](docs/project-client/models/category_delivery_config.md)              |             |
| [InboxTokenResponseCollection](docs/project-client/models/inbox_token_response_collection.md) |             |
| [InboxTokenResponse](docs/project-client/models/inbox_token_response.md)                      |             |
| [DiscardResult](docs/project-client/models/discard_result.md)                                 |             |
| [ApnsTokenCollection](docs/project-client/models/apns_token_collection.md)                    |             |
| [ApnsToken](docs/project-client/models/apns_token.md)                                         |             |
| [ExpoTokenCollection](docs/project-client/models/expo_token_collection.md)                    |             |
| [ExpoToken](docs/project-client/models/expo_token.md)                                         |             |
| [FcmTokenCollection](docs/project-client/models/fcm_token_collection.md)                      |             |
| [FcmToken](docs/project-client/models/fcm_token.md)                                           |             |
| [SlackTokenCollection](docs/project-client/models/slack_token_collection.md)                  |             |
| [SlackToken](docs/project-client/models/slack_token.md)                                       |             |
| [TeamsTokenCollection](docs/project-client/models/teams_token_collection.md)                  |             |
| [TeamsToken](docs/project-client/models/teams_token.md)                                       |             |
| [WebPushTokenCollection](docs/project-client/models/web_push_token_collection.md)             |             |
| [WebPushToken](docs/project-client/models/web_push_token.md)                                  |             |
| [EventCollection](docs/project-client/models/event_collection.md)                             |             |
| [Event](docs/project-client/models/event.md)                                                  |             |
| [IntegrationConfigCollection](docs/project-client/models/integration_config_collection.md)    |             |
| [ApnsConfigCollection](docs/project-client/models/apns_config_collection.md)                  |             |
| [ApnsConfigPayload](docs/project-client/models/apns_config_payload.md)                        |             |
| [AwssnsConfigCollection](docs/project-client/models/awssns_config_collection.md)              |             |
| [AwssnsConfigPayload](docs/project-client/models/awssns_config_payload.md)                    |             |
| [EventSourceConfigCollection](docs/project-client/models/event_source_config_collection.md)   |             |
| [EventSourceConfigPayload](docs/project-client/models/event_source_config_payload.md)         |             |
| [ExpoConfigCollection](docs/project-client/models/expo_config_collection.md)                  |             |
| [ExpoConfigPayload](docs/project-client/models/expo_config_payload.md)                        |             |
| [FcmConfigCollection](docs/project-client/models/fcm_config_collection.md)                    |             |
| [FcmConfigPayload](docs/project-client/models/fcm_config_payload.md)                          |             |
| [GithubConfigCollection](docs/project-client/models/github_config_collection.md)              |             |
| [GithubConfigPayload](docs/project-client/models/github_config_payload.md)                    |             |
| [InboxConfigCollection](docs/project-client/models/inbox_config_collection.md)                |             |
| [InboxConfigPayload](docs/project-client/models/inbox_config_payload.md)                      |             |
| [MailgunConfigCollection](docs/project-client/models/mailgun_config_collection.md)            |             |
| [MailgunConfigPayload](docs/project-client/models/mailgun_config_payload.md)                  |             |
| [PingConfigCollection](docs/project-client/models/ping_config_collection.md)                  |             |
| [PingConfigPayload](docs/project-client/models/ping_config_payload.md)                        |             |
| [SendgridConfigCollection](docs/project-client/models/sendgrid_config_collection.md)          |             |
| [SendgridConfigPayload](docs/project-client/models/sendgrid_config_payload.md)                |             |
| [SesConfigCollection](docs/project-client/models/ses_config_collection.md)                    |             |
| [SesConfigPayload](docs/project-client/models/ses_config_payload.md)                          |             |
| [SlackConfigCollection](docs/project-client/models/slack_config_collection.md)                |             |
| [SlackConfigPayload](docs/project-client/models/slack_config_payload.md)                      |             |
| [StripeConfigCollection](docs/project-client/models/stripe_config_collection.md)              |             |
| [StripeConfigPayload](docs/project-client/models/stripe_config_payload.md)                    |             |
| [TemplatesConfigCollection](docs/project-client/models/templates_config_collection.md)        |             |
| [TwilioConfigCollection](docs/project-client/models/twilio_config_collection.md)              |             |
| [TwilioConfigPayload](docs/project-client/models/twilio_config_payload.md)                    |             |
| [WebpushConfigCollection](docs/project-client/models/webpush_config_collection.md)            |             |
| [WebpushConfigPayload](docs/project-client/models/webpush_config_payload.md)                  |             |
| [AccessTokenCollection](docs/project-client/models/access_token_collection.md)                |             |
| [CreateProjectTokenRequest](docs/project-client/models/create_project_token_request.md)       |             |
| [CreateTokenResponse](docs/project-client/models/create_token_response.md)                    |             |
| [DiscardTokenResponse](docs/project-client/models/discard_token_response.md)                  |             |
| [CreateUserTokenRequest](docs/project-client/models/create_user_token_request.md)             |             |
| [DeliveryPlanCollection](docs/project-client/models/delivery_plan_collection.md)              |             |
| [UserCollection](docs/project-client/models/user_collection.md)                               |             |
| [User](docs/project-client/models/user.md)                                                    |             |
| [Links](docs/project-client/models/links.md)                                                  |             |
| [IntegrationConfig](docs/project-client/models/integration_config.md)                         |             |
| [ApnsConfig](docs/project-client/models/apns_config.md)                                       |             |
| [AwssnsConfig](docs/project-client/models/awssns_config.md)                                   |             |
| [EventSourceConfig](docs/project-client/models/event_source_config.md)                        |             |
| [ExpoConfig](docs/project-client/models/expo_config.md)                                       |             |
| [FcmConfig](docs/project-client/models/fcm_config.md)                                         |             |
| [GithubConfig](docs/project-client/models/github_config.md)                                   |             |
| [InboxConfig](docs/project-client/models/inbox_config.md)                                     |             |
| [MailgunConfig](docs/project-client/models/mailgun_config.md)                                 |             |
| [PingConfig](docs/project-client/models/ping_config.md)                                       |             |
| [SendgridConfig](docs/project-client/models/sendgrid_config.md)                               |             |
| [SesConfig](docs/project-client/models/ses_config.md)                                         |             |
| [SlackConfig](docs/project-client/models/slack_config.md)                                     |             |
| [StripeConfig](docs/project-client/models/stripe_config.md)                                   |             |
| [TemplatesConfig](docs/project-client/models/templates_config.md)                             |             |
| [TwilioConfig](docs/project-client/models/twilio_config.md)                                   |             |
| [WebpushConfig](docs/project-client/models/webpush_config.md)                                 |             |
| [AccessToken](docs/project-client/models/access_token.md)                                     |             |
| [DeliveryPlan](docs/project-client/models/delivery_plan.md)                                   |             |

</details>

<!-- AUTO-GENERATED-CONTENT:END (project-client) -->

# MagicBell GO UserClient

<!-- AUTO-GENERATED-CONTENT:START (user-client) -->

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
| [ChannelsService](docs/user-client/services/channels_service.md)           |
| [IntegrationsService](docs/user-client/services/integrations_service.md)   |
| [NotificationsService](docs/user-client/services/notifications_service.md) |

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
| [InboxTokenResponseCollection](docs/user-client/models/inbox_token_response_collection.md)          |             |
| [InboxToken](docs/user-client/models/inbox_token.md)                                                |             |
| [InboxTokenResponse](docs/user-client/models/inbox_token_response.md)                               |             |
| [DiscardResult](docs/user-client/models/discard_result.md)                                          |             |
| [ApnsTokenCollection](docs/user-client/models/apns_token_collection.md)                             |             |
| [ApnsTokenPayload](docs/user-client/models/apns_token_payload.md)                                   |             |
| [ApnsToken](docs/user-client/models/apns_token.md)                                                  |             |
| [ExpoTokenCollection](docs/user-client/models/expo_token_collection.md)                             |             |
| [ExpoTokenPayload](docs/user-client/models/expo_token_payload.md)                                   |             |
| [ExpoToken](docs/user-client/models/expo_token.md)                                                  |             |
| [FcmTokenCollection](docs/user-client/models/fcm_token_collection.md)                               |             |
| [FcmTokenPayload](docs/user-client/models/fcm_token_payload.md)                                     |             |
| [FcmToken](docs/user-client/models/fcm_token.md)                                                    |             |
| [SlackTokenCollection](docs/user-client/models/slack_token_collection.md)                           |             |
| [SlackTokenPayload](docs/user-client/models/slack_token_payload.md)                                 |             |
| [SlackToken](docs/user-client/models/slack_token.md)                                                |             |
| [TeamsTokenCollection](docs/user-client/models/teams_token_collection.md)                           |             |
| [TeamsTokenPayload](docs/user-client/models/teams_token_payload.md)                                 |             |
| [TeamsToken](docs/user-client/models/teams_token.md)                                                |             |
| [WebPushTokenCollection](docs/user-client/models/web_push_token_collection.md)                      |             |
| [WebPushTokenPayload](docs/user-client/models/web_push_token_payload.md)                            |             |
| [WebPushToken](docs/user-client/models/web_push_token.md)                                           |             |
| [InboxConfigPayload](docs/user-client/models/inbox_config_payload.md)                               |             |
| [SlackInstallation](docs/user-client/models/slack_installation.md)                                  |             |
| [SlackFinishInstallResponse](docs/user-client/models/slack_finish_install_response.md)              |             |
| [SlackStartInstall](docs/user-client/models/slack_start_install.md)                                 |             |
| [SlackStartInstallResponseContent](docs/user-client/models/slack_start_install_response_content.md) |             |
| [TemplatesInstallation](docs/user-client/models/templates_installation.md)                          |             |
| [WebPushStartInstallationResponse](docs/user-client/models/web_push_start_installation_response.md) |             |
| [NotificationCollection](docs/user-client/models/notification_collection.md)                        |             |
| [Links](docs/user-client/models/links.md)                                                           |             |
| [Notification](docs/user-client/models/notification.md)                                             |             |

</details>

<!-- AUTO-GENERATED-CONTENT:END (user-client) -->
