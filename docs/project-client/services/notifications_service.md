# NotificationsService

A list of all methods in the `NotificationsService` service. Click on the method name to view detailed information about that method.

| Methods                             | Description                               |
| :---------------------------------- | :---------------------------------------- |
| [GetDeliveryplan](#getdeliveryplan) | Get the delivery plan for a notification. |

## GetDeliveryplan

Get the delivery plan for a notification.

- HTTP Method: `GET`
- Endpoint: `/notifications/{notification_id}/deliveryplan`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

**Return Type**

`DeliveryPlanCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/projectclientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/projectclient"

)

config := projectclientconfig.NewConfig()
client := projectclient.NewProjectClient(config)

response, err := client.Notifications.GetDeliveryplan(context.Background(), "notificationId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
