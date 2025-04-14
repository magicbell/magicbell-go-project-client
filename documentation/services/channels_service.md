# ChannelsService

A list of all methods in the `ChannelsService` service. Click on the method name to view detailed information about that method.

| Methods                                                           | Description                                                                                                                                                                                                                                                        |
| :---------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [GetDeliveryconfig](#getdeliveryconfig)                           |                                                                                                                                                                                                                                                                    |
| [SaveDeliveryconfig](#savedeliveryconfig)                         |                                                                                                                                                                                                                                                                    |
| [GetInAppInboxUserTokens](#getinappinboxusertokens)               | Lists all in_app tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                    |
| [GetInAppInboxUserToken](#getinappinboxusertoken)                 | Retrieves a specific in_app token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.      |
| [DiscardInAppInboxUserToken](#discardinappinboxusertoken)         | Revokes a specific user's in_app token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.              |
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

## GetDeliveryconfig

- HTTP Method: `GET`
- Endpoint: `/channels/deliveryconfig`

**Parameters**

| Name   | Type                           | Required | Description                   |
| :----- | :----------------------------- | :------- | :---------------------------- |
| ctx    | Context                        | ✅       | Default go language context   |
| params | GetDeliveryconfigRequestParams | ✅       | Additional request parameters |

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


params := channels.GetDeliveryconfigRequestParams{

}

response, err := client.Channels.GetDeliveryconfig(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveDeliveryconfig

- HTTP Method: `PUT`
- Endpoint: `/channels/deliveryconfig`

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
  "github.com/magicbell/magicbell-go-project-client/pkg/util"
  "github.com/magicbell/magicbell-go-project-client/pkg/channels"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)

channel := channels.CHANNEL_IN_APP

categoryDeliveryConfigChannels := channels.CategoryDeliveryConfigChannels{
  Channel: &channel,
  Delay: util.ToPointer(int64(123)),
  If_: util.ToPointer(util.Nullable[string]{ Value: "If_" }),
}

request := channels.CategoryDeliveryConfig{
  Channels: []channels.CategoryDeliveryConfigChannels{categoryDeliveryConfigChannels},
  Disabled: util.ToPointer(true),
  Key: util.ToPointer("Key"),
}

response, err := client.Channels.SaveDeliveryconfig(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetInAppInboxUserTokens

Lists all in_app tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/in_app/inbox/tokens`

**Parameters**

| Name   | Type                                 | Required | Description                   |
| :----- | :----------------------------------- | :------- | :---------------------------- |
| ctx    | Context                              | ✅       | Default go language context   |
| userId | string                               | ✅       |                               |
| params | GetInAppInboxUserTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`InboxTokenResponseCollection`

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


params := channels.GetInAppInboxUserTokensRequestParams{

}

response, err := client.Channels.GetInAppInboxUserTokens(context.Background(), "userId", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetInAppInboxUserToken

Retrieves a specific in_app token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/in_app/inbox/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`InboxTokenResponse`

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

response, err := client.Channels.GetInAppInboxUserToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardInAppInboxUserToken

Revokes a specific user's in_app token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/in_app/inbox/tokens/{token_id}`

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

response, err := client.Channels.DiscardInAppInboxUserToken(context.Background(), "userId", "tokenId")
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

`ApnsTokenCollection`

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


params := channels.GetMobilePushApnsUserTokensRequestParams{

}

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

`ApnsToken`

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

`ExpoTokenCollection`

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


params := channels.GetMobilePushExpoUserTokensRequestParams{

}

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

`ExpoToken`

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

`FcmTokenCollection`

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


params := channels.GetMobilePushFcmUserTokensRequestParams{

}

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

`FcmToken`

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

`SlackTokenCollection`

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


params := channels.GetSlackUserTokensRequestParams{

}

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

`SlackToken`

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

`TeamsTokenCollection`

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


params := channels.GetTeamsUserTokensRequestParams{

}

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

`TeamsToken`

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

`WebPushTokenCollection`

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


params := channels.GetWebPushUserTokensRequestParams{

}

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

`WebPushToken`

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
