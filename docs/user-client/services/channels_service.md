# ChannelsService

A list of all methods in the `ChannelsService` service. Click on the method name to view detailed information about that method.

| Methods                                                   | Description                                                                                                                                                                                                                                                      |
| :-------------------------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [GetInAppInboxTokens](#getinappinboxtokens)               | Lists all in_app tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                |
| [SaveInAppInboxToken](#saveinappinboxtoken)               | Saves a in_app token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.      |
| [GetInAppInboxToken](#getinappinboxtoken)                 | Retrieves details of a specific in_app token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                              |
| [DiscardInAppInboxToken](#discardinappinboxtoken)         | Revokes one of the authenticated user's in_app tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                          |
| [GetMobilePushApnsTokens](#getmobilepushapnstokens)       | Lists all mobile_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                           |
| [SaveMobilePushApnsToken](#savemobilepushapnstoken)       | Saves a mobile_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel. |
| [GetMobilePushApnsToken](#getmobilepushapnstoken)         | Retrieves details of a specific mobile_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                         |
| [DiscardMobilePushApnsToken](#discardmobilepushapnstoken) | Revokes one of the authenticated user's mobile_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                     |
| [GetMobilePushExpoTokens](#getmobilepushexpotokens)       | Lists all mobile_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                           |
| [SaveMobilePushExpoToken](#savemobilepushexpotoken)       | Saves a mobile_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel. |
| [GetMobilePushExpoToken](#getmobilepushexpotoken)         | Retrieves details of a specific mobile_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                         |
| [DiscardMobilePushExpoToken](#discardmobilepushexpotoken) | Revokes one of the authenticated user's mobile_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                     |
| [GetMobilePushFcmTokens](#getmobilepushfcmtokens)         | Lists all mobile_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                           |
| [SaveMobilePushFcmToken](#savemobilepushfcmtoken)         | Saves a mobile_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel. |
| [GetMobilePushFcmToken](#getmobilepushfcmtoken)           | Retrieves details of a specific mobile_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                         |
| [DiscardMobilePushFcmToken](#discardmobilepushfcmtoken)   | Revokes one of the authenticated user's mobile_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                     |
| [GetSlackTokens](#getslacktokens)                         | Lists all slack tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                 |
| [SaveSlackToken](#saveslacktoken)                         | Saves a slack token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.       |
| [GetSlackToken](#getslacktoken)                           | Retrieves details of a specific slack token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                               |
| [DiscardSlackToken](#discardslacktoken)                   | Revokes one of the authenticated user's slack tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                           |
| [GetTeamsTokens](#getteamstokens)                         | Lists all teams tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                 |
| [SaveTeamsToken](#saveteamstoken)                         | Saves a teams token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.       |
| [GetTeamsToken](#getteamstoken)                           | Retrieves details of a specific teams token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                               |
| [DiscardTeamsToken](#discardteamstoken)                   | Revokes one of the authenticated user's teams tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                           |
| [GetWebPushTokens](#getwebpushtokens)                     | Lists all web_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                              |
| [SaveWebPushToken](#savewebpushtoken)                     | Saves a web_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.    |
| [GetWebPushToken](#getwebpushtoken)                       | Retrieves details of a specific web_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                            |
| [DiscardWebPushToken](#discardwebpushtoken)               | Revokes one of the authenticated user's web_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                        |

## GetInAppInboxTokens

Lists all in_app tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/in_app/inbox/tokens`

**Parameters**

| Name   | Type                             | Required | Description                   |
| :----- | :------------------------------- | :------- | :---------------------------- |
| ctx    | Context                          | ✅       | Default go language context   |
| params | GetInAppInboxTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`InboxTokenResponseCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


params := channels.GetInAppInboxTokensRequestParams{

}

response, err := client.Channels.GetInAppInboxTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveInAppInboxToken

Saves a in_app token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `POST`
- Endpoint: `/channels/in_app/inbox/tokens`

**Parameters**

| Name       | Type       | Required | Description                 |
| :--------- | :--------- | :------- | :-------------------------- |
| ctx        | Context    | ✅       | Default go language context |
| inboxToken | InboxToken | ✅       |                             |

**Return Type**

`InboxToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


request := channels.InboxToken{
  ConnectionId: util.ToPointer(util.Nullable[string]{ Value: "ConnectionId" }),
  Token: util.ToPointer("Token"),
}

response, err := client.Channels.SaveInAppInboxToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetInAppInboxToken

Retrieves details of a specific in_app token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/in_app/inbox/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`InboxTokenResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.GetInAppInboxToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardInAppInboxToken

Revokes one of the authenticated user's in_app tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/in_app/inbox/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.DiscardInAppInboxToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushApnsTokens

Lists all mobile_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/apns/tokens`

**Parameters**

| Name   | Type                                 | Required | Description                   |
| :----- | :----------------------------------- | :------- | :---------------------------- |
| ctx    | Context                              | ✅       | Default go language context   |
| params | GetMobilePushApnsTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ApnsTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


params := channels.GetMobilePushApnsTokensRequestParams{

}

response, err := client.Channels.GetMobilePushApnsTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveMobilePushApnsToken

Saves a mobile_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `POST`
- Endpoint: `/channels/mobile_push/apns/tokens`

**Parameters**

| Name             | Type             | Required | Description                 |
| :--------------- | :--------------- | :------- | :-------------------------- |
| ctx              | Context          | ✅       | Default go language context |
| apnsTokenPayload | ApnsTokenPayload | ✅       |                             |

**Return Type**

`ApnsTokenPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

apnsTokenPayloadInstallationId := channels.APNS_TOKEN_PAYLOAD_INSTALLATION_ID_DEVELOPMENT

request := channels.ApnsTokenPayload{
  AppId: util.ToPointer("AppId"),
  DeviceToken: util.ToPointer("DeviceToken"),
  InstallationId: &apnsTokenPayloadInstallationId,
}

response, err := client.Channels.SaveMobilePushApnsToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushApnsToken

Retrieves details of a specific mobile_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/apns/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`ApnsToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.GetMobilePushApnsToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardMobilePushApnsToken

Revokes one of the authenticated user's mobile_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/mobile_push/apns/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.DiscardMobilePushApnsToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushExpoTokens

Lists all mobile_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/expo/tokens`

**Parameters**

| Name   | Type                                 | Required | Description                   |
| :----- | :----------------------------------- | :------- | :---------------------------- |
| ctx    | Context                              | ✅       | Default go language context   |
| params | GetMobilePushExpoTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ExpoTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


params := channels.GetMobilePushExpoTokensRequestParams{

}

response, err := client.Channels.GetMobilePushExpoTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveMobilePushExpoToken

Saves a mobile_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `POST`
- Endpoint: `/channels/mobile_push/expo/tokens`

**Parameters**

| Name             | Type             | Required | Description                 |
| :--------------- | :--------------- | :------- | :-------------------------- |
| ctx              | Context          | ✅       | Default go language context |
| expoTokenPayload | ExpoTokenPayload | ✅       |                             |

**Return Type**

`ExpoTokenPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


request := channels.ExpoTokenPayload{
  DeviceToken: util.ToPointer("DeviceToken"),
}

response, err := client.Channels.SaveMobilePushExpoToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushExpoToken

Retrieves details of a specific mobile_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/expo/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`ExpoToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.GetMobilePushExpoToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardMobilePushExpoToken

Revokes one of the authenticated user's mobile_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/mobile_push/expo/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.DiscardMobilePushExpoToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushFcmTokens

Lists all mobile_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/fcm/tokens`

**Parameters**

| Name   | Type                                | Required | Description                   |
| :----- | :---------------------------------- | :------- | :---------------------------- |
| ctx    | Context                             | ✅       | Default go language context   |
| params | GetMobilePushFcmTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`FcmTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


params := channels.GetMobilePushFcmTokensRequestParams{

}

response, err := client.Channels.GetMobilePushFcmTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveMobilePushFcmToken

Saves a mobile_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `POST`
- Endpoint: `/channels/mobile_push/fcm/tokens`

**Parameters**

| Name            | Type            | Required | Description                 |
| :-------------- | :-------------- | :------- | :-------------------------- |
| ctx             | Context         | ✅       | Default go language context |
| fcmTokenPayload | FcmTokenPayload | ✅       |                             |

**Return Type**

`FcmTokenPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

fcmTokenPayloadInstallationId := channels.FCM_TOKEN_PAYLOAD_INSTALLATION_ID_DEVELOPMENT

request := channels.FcmTokenPayload{
  DeviceToken: util.ToPointer("DeviceToken"),
  InstallationId: &fcmTokenPayloadInstallationId,
}

response, err := client.Channels.SaveMobilePushFcmToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMobilePushFcmToken

Retrieves details of a specific mobile_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/fcm/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`FcmToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.GetMobilePushFcmToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardMobilePushFcmToken

Revokes one of the authenticated user's mobile_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/mobile_push/fcm/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.DiscardMobilePushFcmToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSlackTokens

Lists all slack tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/slack/tokens`

**Parameters**

| Name   | Type                        | Required | Description                   |
| :----- | :-------------------------- | :------- | :---------------------------- |
| ctx    | Context                     | ✅       | Default go language context   |
| params | GetSlackTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`SlackTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


params := channels.GetSlackTokensRequestParams{

}

response, err := client.Channels.GetSlackTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSlackToken

Saves a slack token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `POST`
- Endpoint: `/channels/slack/tokens`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| slackTokenPayload | SlackTokenPayload | ✅       |                             |

**Return Type**

`SlackTokenPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


slackTokenPayloadOauth := channels.SlackTokenPayloadOauth{
  ChannelId: util.ToPointer("ChannelId"),
  InstallationId: util.ToPointer("InstallationId"),
  Scope: util.ToPointer("Scope"),
}


slackTokenPayloadWebhook := channels.SlackTokenPayloadWebhook{
  Url: util.ToPointer("Url"),
}

request := channels.SlackTokenPayload{
  Oauth: &slackTokenPayloadOauth,
  Webhook: &slackTokenPayloadWebhook,
}

response, err := client.Channels.SaveSlackToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSlackToken

Retrieves details of a specific slack token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/slack/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`SlackToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.GetSlackToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardSlackToken

Revokes one of the authenticated user's slack tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/slack/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.DiscardSlackToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTeamsTokens

Lists all teams tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/teams/tokens`

**Parameters**

| Name   | Type                        | Required | Description                   |
| :----- | :-------------------------- | :------- | :---------------------------- |
| ctx    | Context                     | ✅       | Default go language context   |
| params | GetTeamsTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`TeamsTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


params := channels.GetTeamsTokensRequestParams{

}

response, err := client.Channels.GetTeamsTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveTeamsToken

Saves a teams token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `POST`
- Endpoint: `/channels/teams/tokens`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| teamsTokenPayload | TeamsTokenPayload | ✅       |                             |

**Return Type**

`TeamsTokenPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


teamsTokenPayloadWebhook := channels.TeamsTokenPayloadWebhook{
  Url: util.ToPointer("Url"),
}

request := channels.TeamsTokenPayload{
  Webhook: &teamsTokenPayloadWebhook,
}

response, err := client.Channels.SaveTeamsToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTeamsToken

Retrieves details of a specific teams token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/teams/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`TeamsToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.GetTeamsToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardTeamsToken

Revokes one of the authenticated user's teams tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/teams/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.DiscardTeamsToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetWebPushTokens

Lists all web_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/web_push/tokens`

**Parameters**

| Name   | Type                          | Required | Description                   |
| :----- | :---------------------------- | :------- | :---------------------------- |
| ctx    | Context                       | ✅       | Default go language context   |
| params | GetWebPushTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`WebPushTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


params := channels.GetWebPushTokensRequestParams{

}

response, err := client.Channels.GetWebPushTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveWebPushToken

Saves a web_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `POST`
- Endpoint: `/channels/web_push/tokens`

**Parameters**

| Name                | Type                       | Required | Description                 |
| :------------------ | :------------------------- | :------- | :-------------------------- |
| ctx                 | Context                    | ✅       | Default go language context |
| webPushTokenPayload | shared.WebPushTokenPayload | ✅       |                             |

**Return Type**

`shared.WebPushTokenPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/shared"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


webPushTokenPayloadKeys := shared.WebPushTokenPayloadKeys{
  Auth: util.ToPointer("Auth"),
  P256dh: util.ToPointer("P256dh"),
}

request := shared.WebPushTokenPayload{
  Endpoint: util.ToPointer("Endpoint"),
  Keys: &webPushTokenPayloadKeys,
}

response, err := client.Channels.SaveWebPushToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetWebPushToken

Retrieves details of a specific web_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/web_push/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`WebPushToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.GetWebPushToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardWebPushToken

Revokes one of the authenticated user's web_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/web_push/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)

response, err := client.Channels.DiscardWebPushToken(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
