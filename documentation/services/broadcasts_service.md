# BroadcastsService

A list of all methods in the `BroadcastsService` service. Click on the method name to view detailed information about that method.

| Methods                             | Description                                                                                                                                                                           |
| :---------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [ListBroadcasts](#listbroadcasts)   | Retrieves a paginated list of broadcasts for the project. Returns basic information about each broadcast including its creation time and status.                                      |
| [CreateBroadcast](#createbroadcast) | Creates a new broadcast message. When a broadcast is created, it generates individual notifications for relevant users within the project. Only administrators can create broadcasts. |
| [FetchBroadcast](#fetchbroadcast)   | Retrieves detailed information about a specific broadcast by its ID. Includes the broadcast's configuration and current status.                                                       |

## ListBroadcasts

Retrieves a paginated list of broadcasts for the project. Returns basic information about each broadcast including its creation time and status.

- HTTP Method: `GET`
- Endpoint: `/broadcasts`

**Parameters**

| Name   | Type                        | Required | Description                   |
| :----- | :-------------------------- | :------- | :---------------------------- |
| ctx    | Context                     | ✅       | Default go language context   |
| params | ListBroadcastsRequestParams | ✅       | Additional request parameters |

**Return Type**

`ArrayOfBroadcasts`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/broadcasts"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := broadcasts.ListBroadcastsRequestParams{}


response, err := client.Broadcasts.ListBroadcasts(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateBroadcast

Creates a new broadcast message. When a broadcast is created, it generates individual notifications for relevant users within the project. Only administrators can create broadcasts.

- HTTP Method: `POST`
- Endpoint: `/broadcasts`

**Parameters**

| Name      | Type      | Required | Description                 |
| :-------- | :-------- | :------- | :-------------------------- |
| ctx       | Context   | ✅       | Default go language context |
| broadcast | Broadcast | ✅       |                             |

**Return Type**

`Broadcast`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/broadcasts"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


email := broadcasts.Email{}
email.SetActionUrl("ActionUrl")
email.SetContent("Content")
email.SetTitle("Title")


inApp := broadcasts.InApp{}
inApp.SetActionUrl("ActionUrl")
inApp.SetContent("Content")
inApp.SetTitle("Title")


mobilePush := broadcasts.MobilePush{}
mobilePush.SetActionUrl("ActionUrl")
mobilePush.SetContent("Content")
mobilePush.SetTitle("Title")


slack := broadcasts.Slack{}
slack.SetActionUrl("ActionUrl")
slack.SetContent("Content")
slack.SetTitle("Title")


sms := broadcasts.Sms{}
sms.SetActionUrl("ActionUrl")
sms.SetContent("Content")
sms.SetTitle("Title")


webPush := broadcasts.WebPush{}
webPush.SetActionUrl("ActionUrl")
webPush.SetContent("Content")
webPush.SetTitle("Title")

overridesChannels := broadcasts.OverridesChannels{}
overridesChannels.SetEmail(email)
overridesChannels.SetInApp(inApp)
overridesChannels.SetMobilePush(mobilePush)
overridesChannels.SetSlack(slack)
overridesChannels.SetSms(sms)
overridesChannels.SetWebPush(webPush)


providers := broadcasts.Providers{}
providers.SetAmazonSes("string")
providers.SetAndroid("string")
providers.SetIos("string")
providers.SetMailgun("string")
providers.SetPostmark("string")
providers.SetSendgrid("string")
providers.SetSlack("string")

overrides := broadcasts.Overrides{}
overrides.SetChannels(overridesChannels)
overrides.SetProviders(providers)


errors := broadcasts.Errors{}
errors.SetMessage("Message")

statusStatus := broadcasts.STATUS_STATUS_ENQUEUED


summary := broadcasts.Summary{}
summary.SetFailures(int64(123))
summary.SetTotal(int64(123))

broadcastStatus := broadcasts.BroadcastStatus{}
broadcastStatus.SetErrors([]broadcasts.Errors{errors})
broadcastStatus.SetStatus(statusStatus)
broadcastStatus.SetSummary(summary)

request := broadcasts.Broadcast{}
request.SetActionUrl("ActionUrl")
request.SetCategory("Category")
request.SetContent("Content")
request.SetCreatedAt("CreatedAt")
request.SetCustomAttributes("string")
request.SetId("Id")
request.SetOverrides(overrides)
request.SetRecipients([]any{})
request.SetStatus(broadcastStatus)
request.SetTitle("Title")
request.SetTopic("Topic")

response, err := client.Broadcasts.CreateBroadcast(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchBroadcast

Retrieves detailed information about a specific broadcast by its ID. Includes the broadcast's configuration and current status.

- HTTP Method: `GET`
- Endpoint: `/broadcasts/{broadcast_id}`

**Parameters**

| Name        | Type    | Required | Description                 |
| :---------- | :------ | :------- | :-------------------------- |
| ctx         | Context | ✅       | Default go language context |
| broadcastId | string  | ✅       |                             |

**Return Type**

`Broadcast`

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

response, err := client.Broadcasts.FetchBroadcast(context.Background(), "broadcastId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
