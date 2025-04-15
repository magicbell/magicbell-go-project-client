# BroadcastsService

A list of all methods in the `BroadcastsService` service. Click on the method name to view detailed information about that method.

| Methods                             | Description                                                                                                                                      |
| :---------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------- |
| [ListBroadcasts](#listbroadcasts)   | Retrieves a paginated list of broadcasts for the project. Returns basic information about each broadcast including its creation time and status. |
| [CreateBroadcast](#createbroadcast) | Creates a new broadcast message. When a broadcast is created, it generates individual notifications for relevant users within the project.       |
| [FetchBroadcast](#fetchbroadcast)   | Retrieves detailed information about a specific broadcast by its ID. Includes the broadcast's configuration and current status.                  |

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

`BroadcastCollection`

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


params := broadcasts.ListBroadcastsRequestParams{

}

response, err := client.Broadcasts.ListBroadcasts(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateBroadcast

Creates a new broadcast message. When a broadcast is created, it generates individual notifications for relevant users within the project.

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
  "github.com/magicbell/magicbell-go-project-client/pkg/util"
  "github.com/magicbell/magicbell-go-project-client/pkg/broadcasts"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


email := broadcasts.Email{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "ActionUrl" }),
  Content: util.ToPointer("Content"),
  Title: util.ToPointer("Title"),
}


inApp := broadcasts.InApp{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "ActionUrl" }),
  Content: util.ToPointer("Content"),
  Title: util.ToPointer("Title"),
}


mobilePush := broadcasts.MobilePush{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "ActionUrl" }),
  Content: util.ToPointer("Content"),
  Title: util.ToPointer("Title"),
}


slack := broadcasts.Slack{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "ActionUrl" }),
  Content: util.ToPointer("Content"),
  Title: util.ToPointer("Title"),
}


sms := broadcasts.Sms{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "ActionUrl" }),
  Content: util.ToPointer("Content"),
  Title: util.ToPointer("Title"),
}


webPush := broadcasts.WebPush{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "ActionUrl" }),
  Content: util.ToPointer("Content"),
  Title: util.ToPointer("Title"),
}

overridesChannels := broadcasts.OverridesChannels{
  Email: &email,
  InApp: &inApp,
  MobilePush: &mobilePush,
  Slack: &slack,
  Sms: &sms,
  WebPush: &webPush,
}


providers := broadcasts.Providers{
  AmazonSes: []byte{},
  Android: []byte{},
  Ios: []byte{},
  Mailgun: []byte{},
  Postmark: []byte{},
  Sendgrid: []byte{},
  Slack: []byte{},
}

overrides := broadcasts.Overrides{
  Channels: &overridesChannels,
  Providers: &providers,
}


errors := broadcasts.Errors{
  Message: util.ToPointer("Message"),
}

statusStatus := broadcasts.STATUS_STATUS_ENQUEUED


summary := broadcasts.Summary{
  Failures: util.ToPointer(int64(123)),
  Total: util.ToPointer(int64(123)),
}

broadcastStatus := broadcasts.BroadcastStatus{
  Errors: []broadcasts.Errors{errors},
  Status: &statusStatus,
  Summary: &summary,
}

request := broadcasts.Broadcast{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "ActionUrl" }),
  Category: util.ToPointer(util.Nullable[string]{ Value: "Category" }),
  Content: util.ToPointer(util.Nullable[string]{ Value: "Content" }),
  CreatedAt: util.ToPointer("CreatedAt"),
  CustomAttributes: []byte{},
  Id: util.ToPointer("Id"),
  Overrides: &overrides,
  Recipients: []any{},
  Status: &broadcastStatus,
  Title: util.ToPointer("Title"),
  Topic: util.ToPointer(util.Nullable[string]{ Value: "Topic" }),
}

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
