# NotificationsService

A list of all methods in the `NotificationsService` service. Click on the method name to view detailed information about that method.

| Methods                                           | Description                         |
| :------------------------------------------------ | :---------------------------------- |
| [ListNotifications](#listnotifications)           | Lists all notifications for a user. |
| [ArchiveNotification](#archivenotification)       | Archives a notification.            |
| [MarkNotificationRead](#marknotificationread)     | Marks a notification as read.       |
| [UnarchiveNotification](#unarchivenotification)   | Unarchives a notification.          |
| [MarkNotificationUnread](#marknotificationunread) | Marks a notification as unread.     |

## ListNotifications

Lists all notifications for a user.

- HTTP Method: `GET`
- Endpoint: `/notifications`

**Parameters**

| Name   | Type                           | Required | Description                   |
| :----- | :----------------------------- | :------- | :---------------------------- |
| ctx    | Context                        | ✅       | Default go language context   |
| params | ListNotificationsRequestParams | ✅       | Additional request parameters |

**Return Type**

`NotificationCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/userclient"

  "github.com/magicbell/magicbell-go/pkg/user-client/notifications"
)

config := userclientconfig.NewConfig()
client := userclient.NewUserClient(config)


params := notifications.ListNotificationsRequestParams{

}

response, err := client.Notifications.ListNotifications(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ArchiveNotification

Archives a notification.

- HTTP Method: `POST`
- Endpoint: `/notifications/{notification_id}/archive`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Notifications.ArchiveNotification(context.Background(), "notificationId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## MarkNotificationRead

Marks a notification as read.

- HTTP Method: `POST`
- Endpoint: `/notifications/{notification_id}/read`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Notifications.MarkNotificationRead(context.Background(), "notificationId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## UnarchiveNotification

Unarchives a notification.

- HTTP Method: `POST`
- Endpoint: `/notifications/{notification_id}/unarchive`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Notifications.UnarchiveNotification(context.Background(), "notificationId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## MarkNotificationUnread

Marks a notification as unread.

- HTTP Method: `POST`
- Endpoint: `/notifications/{notification_id}/unread`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Notifications.MarkNotificationUnread(context.Background(), "notificationId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
