# UsersService

A list of all methods in the `UsersService` service. Click on the method name to view detailed information about that method.

| Methods                   | Description                                                                                                              |
| :------------------------ | :----------------------------------------------------------------------------------------------------------------------- |
| [ListUsers](#listusers)   |                                                                                                                          |
| [CreateUser](#createuser) | Creates a user with the provided details. The user will be associated with the project specified in the request context. |
| [DeleteUser](#deleteuser) |                                                                                                                          |

## ListUsers

- HTTP Method: `GET`
- Endpoint: `/users`

**Parameters**

| Name   | Type                   | Required | Description                   |
| :----- | :--------------------- | :------- | :---------------------------- |
| ctx    | Context                | ✅       | Default go language context   |
| params | ListUsersRequestParams | ✅       | Additional request parameters |

**Return Type**

`UserCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/projectclientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/projectclient"

  "github.com/magicbell/magicbell-go/pkg/project-client/users"
)

config := projectclientconfig.NewConfig()
client := projectclient.NewProjectClient(config)


params := users.ListUsersRequestParams{

}

response, err := client.Users.ListUsers(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateUser

Creates a user with the provided details. The user will be associated with the project specified in the request context.

- HTTP Method: `POST`
- Endpoint: `/users`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| user | User    | ✅       |                             |

**Return Type**

`User`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/projectclientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/projectclient"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/users"
)

config := projectclientconfig.NewConfig()
client := projectclient.NewProjectClient(config)


request := users.User{
  CreatedAt: util.ToPointer(util.Nullable[string]{ Value: "CreatedAt" }),
  CustomAttributes: []byte{},
  Email: util.ToPointer(util.Nullable[string]{ Value: "Email" }),
  ExternalId: util.ToPointer(util.Nullable[string]{ Value: "ExternalId" }),
  FirstName: util.ToPointer(util.Nullable[string]{ Value: "FirstName" }),
  Id: util.ToPointer("Id"),
  LastName: util.ToPointer(util.Nullable[string]{ Value: "LastName" }),
  LastNotifiedAt: util.ToPointer(util.Nullable[string]{ Value: "LastNotifiedAt" }),
  LastSeenAt: util.ToPointer(util.Nullable[string]{ Value: "LastSeenAt" }),
  UpdatedAt: util.ToPointer(util.Nullable[string]{ Value: "UpdatedAt" }),
}

response, err := client.Users.CreateUser(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUser

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}`

**Parameters**

| Name   | Type    | Required | Description                 |
| :----- | :------ | :------- | :-------------------------- |
| ctx    | Context | ✅       | Default go language context |
| userId | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Users.DeleteUser(context.Background(), "userId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
