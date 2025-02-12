# UsersService

A list of all methods in the `UsersService` service. Click on the method name to view detailed information about that method.

| Methods                 | Description |
| :---------------------- | :---------- |
| [ListUsers](#listusers) |             |

## ListUsers

- HTTP Method: `GET`
- Endpoint: `/users`

**Parameters**

| Name   | Type                   | Required | Description                   |
| :----- | :--------------------- | :------- | :---------------------------- |
| ctx    | Context                | ✅       | Default go language context   |
| params | ListUsersRequestParams | ✅       | Additional request parameters |

**Return Type**

`ArrayOfUsers`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/users"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := users.ListUsersRequestParams{}


response, err := client.Users.ListUsers(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
