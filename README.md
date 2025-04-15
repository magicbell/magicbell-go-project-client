# MagicbellProjectClient Go SDK 0.2.0

Welcome to the MagicbellProjectClient SDK documentation. This guide will help you get started with integrating and using the MagicbellProjectClient SDK in your project.

[![This SDK was generated by liblab](https://public-liblab-readme-assets.s3.us-east-1.amazonaws.com/built-by-liblab-icon.svg)](https://liblab.com/?utm_source=readme)

## Versions

- API version: `2.0.0`
- SDK version: `0.2.0`

## About the API

OpenAPI 3.1.0 Specification for MagicBell API.

## Table of Contents

- [Setup & Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
  - [Installation](#installation)
- [Authentication](#authentication)
  - [Access Token Authentication](#access-token-authentication)
- [Services](#services)
  - [Response Wrappers](#response-wrappers)
- [Models](#models)
- [License](#license)

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Authentication

### Access Token Authentication

The magicbell-project-client API uses an Access Token for authentication.

This token must be provided to authenticate your requests to the API.

#### Setting the Access Token

When you initialize the SDK, you can set the access token as follows:

```go
import (
    "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
    "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  )

config := magicbellprojectclientconfig.NewConfig()
config.SetAccessToken("YOUR-TOKEN")

sdk := magicbellprojectclient.NewMagicbellProjectClient(config)
```

If you need to set or update the access token after initializing the SDK, you can use:

```go
import (
    "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
    "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  )

config := magicbellprojectclientconfig.NewConfig()

sdk := magicbellprojectclient.NewMagicbellProjectClient(config)
sdk.SetAccessToken("YOUR-TOKEN")
```

## Services

The SDK provides various services to interact with the API.

<details> 
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                                    |
| :---------------------------------------------------------------------- |
| [BroadcastsService](documentation/services/broadcasts_service.md)       |
| [ChannelsService](documentation/services/channels_service.md)           |
| [EventsService](documentation/services/events_service.md)               |
| [IntegrationsService](documentation/services/integrations_service.md)   |
| [JwtService](documentation/services/jwt_service.md)                     |
| [NotificationsService](documentation/services/notifications_service.md) |
| [UsersService](documentation/services/users_service.md)                 |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `MagicbellProjectClientResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                                     | Description                                 |
| :------- | :--------------------------------------- | :------------------------------------------ |
| Data     | `T`                                      | The body of the API response                |
| Metadata | `MagicbellProjectClientResponseMetadata` | Status code and headers returned by the API |

#### `MagicbellProjectClientError`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                                     | Description                                 |
| :------- | :--------------------------------------- | :------------------------------------------ |
| Err      | `error`                                  | The error that occurred                     |
| Body     | `T`                                      | The body of the API response                |
| Metadata | `MagicbellProjectClientResponseMetadata` | Status code and headers returned by the API |

#### `MagicbellProjectClientResponseMetadata`

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
| [BroadcastCollection](documentation/models/broadcast_collection.md)                     |             |
| [Broadcast](documentation/models/broadcast.md)                                          |             |
| [CategoryDeliveryConfig](documentation/models/category_delivery_config.md)              |             |
| [InboxTokenResponseCollection](documentation/models/inbox_token_response_collection.md) |             |
| [InboxTokenResponse](documentation/models/inbox_token_response.md)                      |             |
| [DiscardResult](documentation/models/discard_result.md)                                 |             |
| [ApnsTokenCollection](documentation/models/apns_token_collection.md)                    |             |
| [ApnsToken](documentation/models/apns_token.md)                                         |             |
| [ExpoTokenCollection](documentation/models/expo_token_collection.md)                    |             |
| [ExpoToken](documentation/models/expo_token.md)                                         |             |
| [FcmTokenCollection](documentation/models/fcm_token_collection.md)                      |             |
| [FcmToken](documentation/models/fcm_token.md)                                           |             |
| [SlackTokenCollection](documentation/models/slack_token_collection.md)                  |             |
| [SlackToken](documentation/models/slack_token.md)                                       |             |
| [TeamsTokenCollection](documentation/models/teams_token_collection.md)                  |             |
| [TeamsToken](documentation/models/teams_token.md)                                       |             |
| [WebPushTokenCollection](documentation/models/web_push_token_collection.md)             |             |
| [WebPushToken](documentation/models/web_push_token.md)                                  |             |
| [EventCollection](documentation/models/event_collection.md)                             |             |
| [Event](documentation/models/event.md)                                                  |             |
| [IntegrationConfigCollection](documentation/models/integration_config_collection.md)    |             |
| [ApnsConfigCollection](documentation/models/apns_config_collection.md)                  |             |
| [ApnsConfigPayload](documentation/models/apns_config_payload.md)                        |             |
| [AwssnsConfigCollection](documentation/models/awssns_config_collection.md)              |             |
| [AwssnsConfigPayload](documentation/models/awssns_config_payload.md)                    |             |
| [EventSourceConfigCollection](documentation/models/event_source_config_collection.md)   |             |
| [EventSourceConfigPayload](documentation/models/event_source_config_payload.md)         |             |
| [ExpoConfigCollection](documentation/models/expo_config_collection.md)                  |             |
| [ExpoConfigPayload](documentation/models/expo_config_payload.md)                        |             |
| [FcmConfigCollection](documentation/models/fcm_config_collection.md)                    |             |
| [FcmConfigPayload](documentation/models/fcm_config_payload.md)                          |             |
| [GithubConfigCollection](documentation/models/github_config_collection.md)              |             |
| [GithubConfigPayload](documentation/models/github_config_payload.md)                    |             |
| [InboxConfigCollection](documentation/models/inbox_config_collection.md)                |             |
| [InboxConfigPayload](documentation/models/inbox_config_payload.md)                      |             |
| [MailgunConfigCollection](documentation/models/mailgun_config_collection.md)            |             |
| [MailgunConfigPayload](documentation/models/mailgun_config_payload.md)                  |             |
| [PingConfigCollection](documentation/models/ping_config_collection.md)                  |             |
| [PingConfigPayload](documentation/models/ping_config_payload.md)                        |             |
| [SendgridConfigCollection](documentation/models/sendgrid_config_collection.md)          |             |
| [SendgridConfigPayload](documentation/models/sendgrid_config_payload.md)                |             |
| [SesConfigCollection](documentation/models/ses_config_collection.md)                    |             |
| [SesConfigPayload](documentation/models/ses_config_payload.md)                          |             |
| [SlackConfigCollection](documentation/models/slack_config_collection.md)                |             |
| [SlackConfigPayload](documentation/models/slack_config_payload.md)                      |             |
| [StripeConfigCollection](documentation/models/stripe_config_collection.md)              |             |
| [StripeConfigPayload](documentation/models/stripe_config_payload.md)                    |             |
| [TemplatesConfigCollection](documentation/models/templates_config_collection.md)        |             |
| [TwilioConfigCollection](documentation/models/twilio_config_collection.md)              |             |
| [TwilioConfigPayload](documentation/models/twilio_config_payload.md)                    |             |
| [WebpushConfigCollection](documentation/models/webpush_config_collection.md)            |             |
| [WebpushConfigPayload](documentation/models/webpush_config_payload.md)                  |             |
| [AccessTokenCollection](documentation/models/access_token_collection.md)                |             |
| [CreateProjectTokenRequest](documentation/models/create_project_token_request.md)       |             |
| [CreateTokenResponse](documentation/models/create_token_response.md)                    |             |
| [DiscardTokenResponse](documentation/models/discard_token_response.md)                  |             |
| [CreateUserTokenRequest](documentation/models/create_user_token_request.md)             |             |
| [DeliveryPlanCollection](documentation/models/delivery_plan_collection.md)              |             |
| [UserCollection](documentation/models/user_collection.md)                               |             |
| [User](documentation/models/user.md)                                                    |             |
| [Links](documentation/models/links.md)                                                  |             |
| [IntegrationConfig](documentation/models/integration_config.md)                         |             |
| [ApnsConfig](documentation/models/apns_config.md)                                       |             |
| [AwssnsConfig](documentation/models/awssns_config.md)                                   |             |
| [EventSourceConfig](documentation/models/event_source_config.md)                        |             |
| [ExpoConfig](documentation/models/expo_config.md)                                       |             |
| [FcmConfig](documentation/models/fcm_config.md)                                         |             |
| [GithubConfig](documentation/models/github_config.md)                                   |             |
| [InboxConfig](documentation/models/inbox_config.md)                                     |             |
| [MailgunConfig](documentation/models/mailgun_config.md)                                 |             |
| [PingConfig](documentation/models/ping_config.md)                                       |             |
| [SendgridConfig](documentation/models/sendgrid_config.md)                               |             |
| [SesConfig](documentation/models/ses_config.md)                                         |             |
| [SlackConfig](documentation/models/slack_config.md)                                     |             |
| [StripeConfig](documentation/models/stripe_config.md)                                   |             |
| [TemplatesConfig](documentation/models/templates_config.md)                             |             |
| [TwilioConfig](documentation/models/twilio_config.md)                                   |             |
| [WebpushConfig](documentation/models/webpush_config.md)                                 |             |
| [AccessToken](documentation/models/access_token.md)                                     |             |
| [DeliveryPlan](documentation/models/delivery_plan.md)                                   |             |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.
