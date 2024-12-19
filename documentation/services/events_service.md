# EventsService

A list of all methods in the `EventsService` service. Click on the method name to view detailed information about that method.

| Methods                 | Description                                           |
| :---------------------- | :---------------------------------------------------- |
| [GetEvents](#getevents) | Retrieves a paginated list of events for the project. |

## GetEvents

Retrieves a paginated list of events for the project.

- HTTP Method: `GET`
- Endpoint: `/events`

**Parameters**

| Name   | Type                   | Required | Description                   |
| :----- | :--------------------- | :------- | :---------------------------- |
| ctx    | Context                | ✅       | Default go language context   |
| params | GetEventsRequestParams | ✅       | Additional request parameters |

**Return Type**

`ArrayOfEvents`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/events"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := events.GetEventsRequestParams{}


response, err := client.Events.GetEvents(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
