# ChannelsService

A list of all methods in the `ChannelsService` service. Click on the method name to view detailed information about that method.

| Methods                                                           | Description                                                                                                                                                                                                                                                        |
| :---------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [GetProjectDeliveryconfig](#getprojectdeliveryconfig)             |                                                                                                                                                                                                                                                                    |
| [SaveProjectDeliveryconfig](#saveprojectdeliveryconfig)           |                                                                                                                                                                                                                                                                    |
| [SaveCategoriesDeliveryconfig](#savecategoriesdeliveryconfig)     |                                                                                                                                                                                                                                                                    |
| [GetMobilePushApnsUserTokens](#getmobilepushapnsusertokens)       | Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                               |
| [GetMobilePushApnsUserToken](#getmobilepushapnsusertoken)         | Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata. |
| [DiscardMobilePushApnsUserToken](#discardmobilepushapnsusertoken) | Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.         |
| [GetMobilePushExpoUserTokens](#getmobilepushexpousertokens)       | Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                               |
| [GetMobilePushExpoUserToken](#getmobilepushexpousertoken)         | Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata. |
| [DiscardMobilePushExpoUserToken](#discardmobilepushexpousertoken) | Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.         |
| [GetMobilePushFcmUserTokens](#getmobilepushfcmusertokens)         | Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                               |
| [GetMobilePushFcmUserToken](#getmobilepushfcmusertoken)           | Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata. |
| [DiscardMobilePushFcmUserToken](#discardmobilepushfcmusertoken)   | Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.         |
| [GetSlackUserTokens](#getslackusertokens)                         | Lists all slack tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                     |
| [GetSlackUserToken](#getslackusertoken)                           | Retrieves a specific slack token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.       |
| [DiscardSlackUserToken](#discardslackusertoken)                   | Revokes a specific user's slack token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.               |
| [GetTeamsUserTokens](#getteamsusertokens)                         | Lists all teams tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                     |
| [GetTeamsUserToken](#getteamsusertoken)                           | Retrieves a specific teams token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.       |
| [DiscardTeamsUserToken](#discardteamsusertoken)                   | Revokes a specific user's teams token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.               |
| [GetWebPushUserTokens](#getwebpushusertokens)                     | Lists all web_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                  |
| [GetWebPushUserToken](#getwebpushusertoken)                       | Retrieves a specific web_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.    |
| [DiscardWebPushUserToken](#discardwebpushusertoken)               | Revokes a specific user's web_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.            |

## GetProjectDeliveryconfig

- HTTP Method: `GET`
- Endpoint: `/channels/deliveryconfig`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`ProjectDeliveryConfig`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.GetProjectDeliveryconfig(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveProjectDeliveryconfig

- HTTP Method: `PUT`
- Endpoint: `/channels/deliveryconfig`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| projectDeliveryConfig | ProjectDeliveryConfig | ✅       |                             |

**Return Type**

`ProjectDeliveryConfig`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/channels"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

channelsChannel1 := channels.CHANNELS_CHANNEL1_IN_APP

projectDeliveryConfigChannels := channels.ProjectDeliveryConfigChannels{}
projectDeliveryConfigChannels.SetChannel(channelsChannel1)
projectDeliveryConfigChannels.SetDelay(int64(123))
projectDeliveryConfigChannels.SetDisabled(true)
projectDeliveryConfigChannels.SetIf_("If_")
projectDeliveryConfigChannels.SetPriority(int64(123))

request := channels.ProjectDeliveryConfig{}
request.SetChannels([]channels.ProjectDeliveryConfigChannels{projectDeliveryConfigChannels})

response, err := client.Channels.SaveProjectDeliveryconfig(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveCategoriesDeliveryconfig

- HTTP Method: `POST`
- Endpoint: `/channels/deliveryconfig/categories`

**Parameters**

| Name                   | Type                   | Required | Description                 |
| :--------------------- | :--------------------- | :------- | :-------------------------- |
| ctx                    | Context                | ✅       | Default go language context |
| categoryDeliveryConfig | CategoryDeliveryConfig | ✅       |                             |

**Return Type**

`CategoryDeliveryConfig`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/channels"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

channelsChannel2 := channels.CHANNELS_CHANNEL2_IN_APP

categoryDeliveryConfigChannels := channels.CategoryDeliveryConfigChannels{}
categoryDeliveryConfigChannels.SetChannel(channelsChannel2)
categoryDeliveryConfigChannels.SetDelay(int64(123))
categoryDeliveryConfigChannels.SetDisabled(true)
categoryDeliveryConfigChannels.SetIf_("If_")
categoryDeliveryConfigChannels.SetPriority(int64(123))

request := channels.CategoryDeliveryConfig{}
request.SetCategory("Category")
request.SetChannels([]channels.CategoryDeliveryConfigChannels{categoryDeliveryConfigChannels})
request.SetDisabled(true)

response, err := client.Channels.SaveCategoriesDeliveryconfig(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushApnsUserTokens

Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/apns/tokens`

**Parameters**

| Name   | Type                                     | Required | Description                   |
| :----- | :--------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                  | ✅       | Default go language context   |
| userId | string                                   | ✅       |                               |
| params | GetMobilePushApnsUserTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ArrayOfMetadataApnsTokens`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/channels"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := channels.GetMobilePushApnsUserTokensRequestParams{}


response, err := client.Channels.GetMobilePushApnsUserTokens(context.Background(), "userId", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushApnsUserToken

Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/apns/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`MetadataApnsToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.GetMobilePushApnsUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardMobilePushApnsUserToken

Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/mobile_push/apns/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.DiscardMobilePushApnsUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushExpoUserTokens

Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/expo/tokens`

**Parameters**

| Name   | Type                                     | Required | Description                   |
| :----- | :--------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                  | ✅       | Default go language context   |
| userId | string                                   | ✅       |                               |
| params | GetMobilePushExpoUserTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ArrayOfMetadataExpoTokens`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/channels"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := channels.GetMobilePushExpoUserTokensRequestParams{}


response, err := client.Channels.GetMobilePushExpoUserTokens(context.Background(), "userId", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushExpoUserToken

Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/expo/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`MetadataExpoToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.GetMobilePushExpoUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardMobilePushExpoUserToken

Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/mobile_push/expo/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.DiscardMobilePushExpoUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushFcmUserTokens

Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/fcm/tokens`

**Parameters**

| Name   | Type                                    | Required | Description                   |
| :----- | :-------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                 | ✅       | Default go language context   |
| userId | string                                  | ✅       |                               |
| params | GetMobilePushFcmUserTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ArrayOfMetadataFcmTokens`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/channels"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := channels.GetMobilePushFcmUserTokensRequestParams{}


response, err := client.Channels.GetMobilePushFcmUserTokens(context.Background(), "userId", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushFcmUserToken

Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/fcm/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`MetadataFcmToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.GetMobilePushFcmUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardMobilePushFcmUserToken

Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/mobile_push/fcm/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.DiscardMobilePushFcmUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSlackUserTokens

Lists all slack tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/slack/tokens`

**Parameters**

| Name   | Type                            | Required | Description                   |
| :----- | :------------------------------ | :------- | :---------------------------- |
| ctx    | Context                         | ✅       | Default go language context   |
| userId | string                          | ✅       |                               |
| params | GetSlackUserTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ArrayOfMetadataSlackTokens`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/channels"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := channels.GetSlackUserTokensRequestParams{}


response, err := client.Channels.GetSlackUserTokens(context.Background(), "userId", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSlackUserToken

Retrieves a specific slack token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/slack/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`MetadataSlackToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.GetSlackUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardSlackUserToken

Revokes a specific user's slack token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/slack/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.DiscardSlackUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTeamsUserTokens

Lists all teams tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/teams/tokens`

**Parameters**

| Name   | Type                            | Required | Description                   |
| :----- | :------------------------------ | :------- | :---------------------------- |
| ctx    | Context                         | ✅       | Default go language context   |
| userId | string                          | ✅       |                               |
| params | GetTeamsUserTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ArrayOfMetadataTeamsTokens`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/channels"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := channels.GetTeamsUserTokensRequestParams{}


response, err := client.Channels.GetTeamsUserTokens(context.Background(), "userId", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTeamsUserToken

Retrieves a specific teams token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/teams/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`MetadataTeamsToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.GetTeamsUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardTeamsUserToken

Revokes a specific user's teams token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/teams/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.DiscardTeamsUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetWebPushUserTokens

Lists all web_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/web_push/tokens`

**Parameters**

| Name   | Type                              | Required | Description                   |
| :----- | :-------------------------------- | :------- | :---------------------------- |
| ctx    | Context                           | ✅       | Default go language context   |
| userId | string                            | ✅       |                               |
| params | GetWebPushUserTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ArrayOfMetadataWebPushTokens`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/channels"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := channels.GetWebPushUserTokensRequestParams{}


response, err := client.Channels.GetWebPushUserTokens(context.Background(), "userId", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetWebPushUserToken

Retrieves a specific web_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/web_push/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`MetadataWebPushToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.GetWebPushUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardWebPushUserToken

Revokes a specific user's web_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/web_push/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

response, err := client.Channels.DiscardWebPushUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
