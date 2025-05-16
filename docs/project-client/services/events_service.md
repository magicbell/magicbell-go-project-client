# EventsService

A list of all methods in the `EventsService` service. Click on the method name to view detailed information about that method.

| Methods                   | Description                                           |
| :------------------------ | :---------------------------------------------------- |
| [ListEvents](#listevents) | Retrieves a paginated list of events for the project. |
| [GetEvent](#getevent)     | Retrieves a project event by its ID.                  |

## ListEvents

Retrieves a paginated list of events for the project.

- HTTP Method: `GET`
- Endpoint: `/events`

**Parameters**

| Name   | Type                    | Required | Description                   |
| :----- | :---------------------- | :------- | :---------------------------- |
| ctx    | Context                 | ✅       | Default go language context   |
| params | ListEventsRequestParams | ✅       | Additional request parameters |

**Return Type**

`EventCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/events"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := events.ListEventsRequestParams{

}

response, err := client.Events.ListEvents(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEvent

Retrieves a project event by its ID.

- HTTP Method: `GET`
- Endpoint: `/events/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`Event`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Events.GetEvent(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
