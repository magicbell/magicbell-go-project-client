---
title: "project-client"
---

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

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
| [BroadcastsService](services/broadcasts_service.md)       |
| [ChannelsService](services/channels_service.md)           |
| [EventsService](services/events_service.md)               |
| [IntegrationsService](services/integrations_service.md)   |
| [JwtService](services/jwt_service.md)                     |
| [NotificationsService](services/notifications_service.md) |
| [UsersService](services/users_service.md)                 |

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
| [BroadcastCollection](models/broadcast_collection.md)                     |             |
| [Broadcast](models/broadcast.md)                                          |             |
| [CategoryDeliveryConfig](models/category_delivery_config.md)              |             |
| [InboxTokenResponseCollection](models/inbox_token_response_collection.md) |             |
| [InboxTokenResponse](models/inbox_token_response.md)                      |             |
| [DiscardResult](models/discard_result.md)                                 |             |
| [ApnsTokenCollection](models/apns_token_collection.md)                    |             |
| [ApnsToken](models/apns_token.md)                                         |             |
| [ExpoTokenCollection](models/expo_token_collection.md)                    |             |
| [ExpoToken](models/expo_token.md)                                         |             |
| [FcmTokenCollection](models/fcm_token_collection.md)                      |             |
| [FcmToken](models/fcm_token.md)                                           |             |
| [SlackTokenCollection](models/slack_token_collection.md)                  |             |
| [SlackToken](models/slack_token.md)                                       |             |
| [TeamsTokenCollection](models/teams_token_collection.md)                  |             |
| [TeamsToken](models/teams_token.md)                                       |             |
| [WebPushTokenCollection](models/web_push_token_collection.md)             |             |
| [WebPushToken](models/web_push_token.md)                                  |             |
| [EventCollection](models/event_collection.md)                             |             |
| [Event](models/event.md)                                                  |             |
| [IntegrationConfigCollection](models/integration_config_collection.md)    |             |
| [ApnsConfigCollection](models/apns_config_collection.md)                  |             |
| [ApnsConfigPayload](models/apns_config_payload.md)                        |             |
| [AwssnsConfigCollection](models/awssns_config_collection.md)              |             |
| [AwssnsConfigPayload](models/awssns_config_payload.md)                    |             |
| [EventSourceConfigCollection](models/event_source_config_collection.md)   |             |
| [EventSourceConfigPayload](models/event_source_config_payload.md)         |             |
| [ExpoConfigCollection](models/expo_config_collection.md)                  |             |
| [ExpoConfigPayload](models/expo_config_payload.md)                        |             |
| [FcmConfigCollection](models/fcm_config_collection.md)                    |             |
| [FcmConfigPayload](models/fcm_config_payload.md)                          |             |
| [GithubConfigCollection](models/github_config_collection.md)              |             |
| [GithubConfigPayload](models/github_config_payload.md)                    |             |
| [InboxConfigCollection](models/inbox_config_collection.md)                |             |
| [InboxConfigPayload](models/inbox_config_payload.md)                      |             |
| [MailgunConfigCollection](models/mailgun_config_collection.md)            |             |
| [MailgunConfigPayload](models/mailgun_config_payload.md)                  |             |
| [PingConfigCollection](models/ping_config_collection.md)                  |             |
| [PingConfigPayload](models/ping_config_payload.md)                        |             |
| [SendgridConfigCollection](models/sendgrid_config_collection.md)          |             |
| [SendgridConfigPayload](models/sendgrid_config_payload.md)                |             |
| [SesConfigCollection](models/ses_config_collection.md)                    |             |
| [SesConfigPayload](models/ses_config_payload.md)                          |             |
| [SlackConfigCollection](models/slack_config_collection.md)                |             |
| [SlackConfigPayload](models/slack_config_payload.md)                      |             |
| [StripeConfigCollection](models/stripe_config_collection.md)              |             |
| [StripeConfigPayload](models/stripe_config_payload.md)                    |             |
| [TemplatesConfigCollection](models/templates_config_collection.md)        |             |
| [TwilioConfigCollection](models/twilio_config_collection.md)              |             |
| [TwilioConfigPayload](models/twilio_config_payload.md)                    |             |
| [WebpushConfigCollection](models/webpush_config_collection.md)            |             |
| [WebpushConfigPayload](models/webpush_config_payload.md)                  |             |
| [AccessTokenCollection](models/access_token_collection.md)                |             |
| [CreateProjectTokenRequest](models/create_project_token_request.md)       |             |
| [CreateTokenResponse](models/create_token_response.md)                    |             |
| [DiscardTokenResponse](models/discard_token_response.md)                  |             |
| [CreateUserTokenRequest](models/create_user_token_request.md)             |             |
| [DeliveryPlanCollection](models/delivery_plan_collection.md)              |             |
| [UserCollection](models/user_collection.md)                               |             |
| [User](models/user.md)                                                    |             |
| [Links](models/links.md)                                                  |             |
| [IntegrationConfig](models/integration_config.md)                         |             |
| [ApnsConfig](models/apns_config.md)                                       |             |
| [AwssnsConfig](models/awssns_config.md)                                   |             |
| [EventSourceConfig](models/event_source_config.md)                        |             |
| [ExpoConfig](models/expo_config.md)                                       |             |
| [FcmConfig](models/fcm_config.md)                                         |             |
| [GithubConfig](models/github_config.md)                                   |             |
| [InboxConfig](models/inbox_config.md)                                     |             |
| [MailgunConfig](models/mailgun_config.md)                                 |             |
| [PingConfig](models/ping_config.md)                                       |             |
| [SendgridConfig](models/sendgrid_config.md)                               |             |
| [SesConfig](models/ses_config.md)                                         |             |
| [SlackConfig](models/slack_config.md)                                     |             |
| [StripeConfig](models/stripe_config.md)                                   |             |
| [TemplatesConfig](models/templates_config.md)                             |             |
| [TwilioConfig](models/twilio_config.md)                                   |             |
| [WebpushConfig](models/webpush_config.md)                                 |             |
| [AccessToken](models/access_token.md)                                     |             |
| [DeliveryPlan](models/delivery_plan.md)                                   |             |

</details>
