# IntegrationsService

A list of all methods in the `IntegrationsService` service. Click on the method name to view detailed information about that method.

| Methods                                                 | Description                                                                                                                                                                                         |
| :------------------------------------------------------ | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [SaveInboxInstallation](#saveinboxinstallation)         | Creates a new installation of a inbox integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.                         |
| [StartInboxInstallation](#startinboxinstallation)       | Initiates the installation flow for a inbox integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.    |
| [SaveSlackInstallation](#saveslackinstallation)         | Creates a new installation of a slack integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.                         |
| [FinishSlackInstallation](#finishslackinstallation)     | Completes the installation flow for a slack integration. This endpoint is typically called after the user has completed any required authorization steps with slack.                                |
| [StartSlackInstallation](#startslackinstallation)       | Initiates the installation flow for a slack integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.    |
| [SaveTemplatesInstallation](#savetemplatesinstallation) | Creates a new installation of a templates integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.                     |
| [SaveWebPushInstallation](#savewebpushinstallation)     | Creates a new installation of a web_push integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.                      |
| [StartWebPushInstallation](#startwebpushinstallation)   | Initiates the installation flow for a web_push integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required. |

## SaveInboxInstallation

Creates a new installation of a inbox integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.

- HTTP Method: `POST`
- Endpoint: `/integrations/inbox/installations`

**Parameters**

| Name               | Type               | Required | Description                 |
| :----------------- | :----------------- | :------- | :-------------------------- |
| ctx                | Context            | ✅       | Default go language context |
| inboxConfigPayload | InboxConfigPayload | ✅       |                             |

**Return Type**

`InboxConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


images := integrations.Images{
  EmptyInboxUrl: util.ToPointer("EmptyInboxUrl"),
}


banner := integrations.Banner{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  BackgroundOpacity: util.ToPointer(float64(123)),
  FontSize: util.ToPointer("FontSize"),
  TextColor: util.ToPointer("TextColor"),
}


dialog := integrations.Dialog{
  AccentColor: util.ToPointer("AccentColor"),
  BackgroundColor: util.ToPointer("BackgroundColor"),
  TextColor: util.ToPointer("TextColor"),
}


footer := integrations.Footer{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  BorderRadius: util.ToPointer("BorderRadius"),
  FontSize: util.ToPointer("FontSize"),
  TextColor: util.ToPointer("TextColor"),
}


header := integrations.Header{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  BorderRadius: util.ToPointer("BorderRadius"),
  FontFamily: util.ToPointer("FontFamily"),
  FontSize: util.ToPointer("FontSize"),
  TextColor: util.ToPointer("TextColor"),
}


icon := integrations.Icon{
  BorderColor: util.ToPointer("BorderColor"),
  Width: util.ToPointer("Width"),
}


defaultHover := integrations.DefaultHover{
  BackgroundColor: util.ToPointer("BackgroundColor"),
}


defaultState := integrations.DefaultState{
  Color: util.ToPointer("Color"),
}

default_ := integrations.Default_{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  BorderRadius: util.ToPointer("BorderRadius"),
  FontFamily: util.ToPointer("FontFamily"),
  FontSize: util.ToPointer("FontSize"),
  Hover: &defaultHover,
  Margin: util.ToPointer("Margin"),
  State: &defaultState,
  TextColor: util.ToPointer("TextColor"),
}


unreadHover := integrations.UnreadHover{
  BackgroundColor: util.ToPointer("BackgroundColor"),
}


unreadState := integrations.UnreadState{
  Color: util.ToPointer("Color"),
}

unread := integrations.Unread{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  Hover: &unreadHover,
  State: &unreadState,
  TextColor: util.ToPointer("TextColor"),
}


unseenHover := integrations.UnseenHover{
  BackgroundColor: util.ToPointer("BackgroundColor"),
}


unseenState := integrations.UnseenState{
  Color: util.ToPointer("Color"),
}

unseen := integrations.Unseen{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  Hover: &unseenHover,
  State: &unseenState,
  TextColor: util.ToPointer("TextColor"),
}

themeNotification := integrations.ThemeNotification{
  Default_: &default_,
  Unread: &unread,
  Unseen: &unseen,
}


unseenBadge := integrations.UnseenBadge{
  BackgroundColor: util.ToPointer("BackgroundColor"),
}

theme := integrations.Theme{
  Banner: &banner,
  Dialog: &dialog,
  Footer: &footer,
  Header: &header,
  Icon: &icon,
  Notification: &themeNotification,
  UnseenBadge: &unseenBadge,
}

request := integrations.InboxConfigPayload{
  Images: &images,
  Locale: util.ToPointer(util.Nullable[string]{ Value: "Locale" }),
  Theme: &theme,
}

response, err := client.Integrations.SaveInboxInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## StartInboxInstallation

Initiates the installation flow for a inbox integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.

- HTTP Method: `POST`
- Endpoint: `/integrations/inbox/installations/start`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`InboxConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Integrations.StartInboxInstallation(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSlackInstallation

Creates a new installation of a slack integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.

- HTTP Method: `POST`
- Endpoint: `/integrations/slack/installations`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| slackInstallation | SlackInstallation | ✅       |                             |

**Return Type**

`SlackInstallation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


authedUser := integrations.AuthedUser{
  AccessToken: util.ToPointer("AccessToken"),
  ExpiresIn: util.ToPointer(int64(123)),
  Id: util.ToPointer("Id"),
  RefreshToken: util.ToPointer("RefreshToken"),
  Scope: util.ToPointer("Scope"),
  TokenType: util.ToPointer("TokenType"),
}


enterprise := integrations.Enterprise{
  Id: util.ToPointer("Id"),
  Name: util.ToPointer("Name"),
}


incomingWebhook := integrations.IncomingWebhook{
  Channel: util.ToPointer("Channel"),
  ConfigurationUrl: util.ToPointer("ConfigurationUrl"),
  Url: util.ToPointer("Url"),
}


team := integrations.Team{
  Id: util.ToPointer("Id"),
  Name: util.ToPointer("Name"),
}

request := integrations.SlackInstallation{
  AccessToken: util.ToPointer("AccessToken"),
  AppId: util.ToPointer("AppId"),
  AuthedUser: &authedUser,
  BotUserId: util.ToPointer("BotUserId"),
  Enterprise: &enterprise,
  ExpiresIn: util.ToPointer(int64(123)),
  Id: util.ToPointer("Id"),
  IncomingWebhook: &incomingWebhook,
  IsEnterpriseInstall: util.ToPointer(true),
  RefreshToken: util.ToPointer("RefreshToken"),
  Scope: util.ToPointer("Scope"),
  Team: &team,
  TokenType: util.ToPointer("TokenType"),
}

response, err := client.Integrations.SaveSlackInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FinishSlackInstallation

Completes the installation flow for a slack integration. This endpoint is typically called after the user has completed any required authorization steps with slack.

- HTTP Method: `POST`
- Endpoint: `/integrations/slack/installations/finish`

**Parameters**

| Name                       | Type                       | Required | Description                 |
| :------------------------- | :------------------------- | :------- | :-------------------------- |
| ctx                        | Context                    | ✅       | Default go language context |
| slackFinishInstallResponse | SlackFinishInstallResponse | ✅       |                             |

**Return Type**

`SlackInstallation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.SlackFinishInstallResponse{
  AppId: util.ToPointer("AppId"),
  Code: util.ToPointer("Code"),
  RedirectUrl: util.ToPointer("RedirectUrl"),
}

response, err := client.Integrations.FinishSlackInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## StartSlackInstallation

Initiates the installation flow for a slack integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.

- HTTP Method: `POST`
- Endpoint: `/integrations/slack/installations/start`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| slackStartInstall | SlackStartInstall | ✅       |                             |

**Return Type**

`SlackStartInstallResponseContent`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.SlackStartInstall{
  AppId: util.ToPointer("AppId"),
  AuthUrl: util.ToPointer("AuthUrl"),
  ExtraScopes: []string{},
  RedirectUrl: util.ToPointer("RedirectUrl"),
}

response, err := client.Integrations.StartSlackInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveTemplatesInstallation

Creates a new installation of a templates integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.

- HTTP Method: `POST`
- Endpoint: `/integrations/templates/installations`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| templatesInstallation | TemplatesInstallation | ✅       |                             |

**Return Type**

`TemplatesInstallation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.TemplatesInstallation{
  Category: util.ToPointer(util.Nullable[string]{ Value: "Category" }),
  Channel: util.ToPointer("Channel"),
  Text: util.ToPointer("Text"),
}

response, err := client.Integrations.SaveTemplatesInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveWebPushInstallation

Creates a new installation of a web_push integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.

- HTTP Method: `POST`
- Endpoint: `/integrations/web_push/installations`

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
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/shared"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


webPushTokenPayloadKeys := shared.WebPushTokenPayloadKeys{
  Auth: util.ToPointer("Auth"),
  P256dh: util.ToPointer("P256dh"),
}

request := shared.WebPushTokenPayload{
  Endpoint: util.ToPointer("Endpoint"),
  Keys: &webPushTokenPayloadKeys,
}

response, err := client.Integrations.SaveWebPushInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## StartWebPushInstallation

Initiates the installation flow for a web_push integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.

- HTTP Method: `POST`
- Endpoint: `/integrations/web_push/installations/start`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`WebPushStartInstallationResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Integrations.StartWebPushInstallation(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```
