# JwtService

A list of all methods in the `JwtService` service. Click on the method name to view detailed information about that method.

| Methods                                   | Description                                                                                                                                                                                                                                                                                        |
| :---------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [FetchProjectTokens](#fetchprojecttokens) | Retrieves a list of all active project-level JWT tokens. Returns a paginated list showing token metadata including creation date, last used date, and expiration time. For security reasons, the actual token values are not included in the response.                                             |
| [CreateProjectJwt](#createprojectjwt)     | Creates a new project-level JWT token. These tokens provide project-wide access and should be carefully managed. Only administrators can create project tokens. The returned token should be securely stored as it cannot be retrieved again after creation.                                       |
| [DiscardProjectJwt](#discardprojectjwt)   | Immediately revokes a project-level JWT token. Once revoked, any requests using this token will be rejected. This action is immediate and cannot be undone. Active sessions using this token will be terminated.                                                                                   |
| [CreateUserJwt](#createuserjwt)           | Issues a new user-specific JWT token. These tokens are scoped to individual user permissions and access levels. Only administrators can create user tokens. The token is returned only once at creation time and cannot be retrieved later.                                                        |
| [DiscardUserJwt](#discarduserjwt)         | Revokes a specific user's JWT token. This immediately invalidates the token and terminates any active sessions using it. This action cannot be undone. Administrators should use this to revoke access when needed for security purposes.                                                          |
| [FetchUserTokens](#fetchusertokens)       | Lists all JWT tokens associated with a specific user. Returns token metadata including creation time, last access time, and expiration date. Administrators can use this to audit user token usage and manage active sessions. Token values are not included in the response for security reasons. |

## FetchProjectTokens

Retrieves a list of all active project-level JWT tokens. Returns a paginated list showing token metadata including creation date, last used date, and expiration time. For security reasons, the actual token values are not included in the response.

- HTTP Method: `GET`
- Endpoint: `/jwt/project`

**Parameters**

| Name   | Type                            | Required | Description                   |
| :----- | :------------------------------ | :------- | :---------------------------- |
| ctx    | Context                         | ✅       | Default go language context   |
| params | FetchProjectTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`AccessTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"

  "github.com/magicbell/magicbell-go-project-client/pkg/jwt"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := jwt.FetchProjectTokensRequestParams{

}

response, err := client.Jwt.FetchProjectTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateProjectJwt

Creates a new project-level JWT token. These tokens provide project-wide access and should be carefully managed. Only administrators can create project tokens. The returned token should be securely stored as it cannot be retrieved again after creation.

- HTTP Method: `POST`
- Endpoint: `/jwt/project`

**Parameters**

| Name                      | Type                      | Required | Description                 |
| :------------------------ | :------------------------ | :------- | :-------------------------- |
| ctx                       | Context                   | ✅       | Default go language context |
| createProjectTokenRequest | CreateProjectTokenRequest | ✅       |                             |

**Return Type**

`CreateTokenResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/util"
  "github.com/magicbell/magicbell-go-project-client/pkg/jwt"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


request := jwt.CreateProjectTokenRequest{
  Expiry: util.ToPointer(util.Nullable[int64]{ Value: int64(123) }),
  Name: util.ToPointer("Name"),
}

response, err := client.Jwt.CreateProjectJwt(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardProjectJwt

Immediately revokes a project-level JWT token. Once revoked, any requests using this token will be rejected. This action is immediate and cannot be undone. Active sessions using this token will be terminated.

- HTTP Method: `DELETE`
- Endpoint: `/jwt/project/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardTokenResponse`

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

response, err := client.Jwt.DiscardProjectJwt(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateUserJwt

Issues a new user-specific JWT token. These tokens are scoped to individual user permissions and access levels. Only administrators can create user tokens. The token is returned only once at creation time and cannot be retrieved later.

- HTTP Method: `POST`
- Endpoint: `/jwt/user`

**Parameters**

| Name                   | Type                   | Required | Description                 |
| :--------------------- | :--------------------- | :------- | :-------------------------- |
| ctx                    | Context                | ✅       | Default go language context |
| createUserTokenRequest | CreateUserTokenRequest | ✅       |                             |

**Return Type**

`CreateTokenResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"
  "github.com/magicbell/magicbell-go-project-client/pkg/util"
  "github.com/magicbell/magicbell-go-project-client/pkg/jwt"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


request := jwt.CreateUserTokenRequest{
  Email: util.ToPointer(util.Nullable[string]{ Value: "Email" }),
  Expiry: util.ToPointer(util.Nullable[int64]{ Value: int64(123) }),
  ExternalId: util.ToPointer(util.Nullable[string]{ Value: "ExternalId" }),
  Name: util.ToPointer(util.Nullable[string]{ Value: "Name" }),
}

response, err := client.Jwt.CreateUserJwt(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DiscardUserJwt

Revokes a specific user's JWT token. This immediately invalidates the token and terminates any active sessions using it. This action cannot be undone. Administrators should use this to revoke access when needed for security purposes.

- HTTP Method: `DELETE`
- Endpoint: `/jwt/user/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardTokenResponse`

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

response, err := client.Jwt.DiscardUserJwt(context.Background(), "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserTokens

Lists all JWT tokens associated with a specific user. Returns token metadata including creation time, last access time, and expiration date. Administrators can use this to audit user token usage and manage active sessions. Token values are not included in the response for security reasons.

- HTTP Method: `GET`
- Endpoint: `/jwt/user/{user_id}`

**Parameters**

| Name   | Type                         | Required | Description                   |
| :----- | :--------------------------- | :------- | :---------------------------- |
| ctx    | Context                      | ✅       | Default go language context   |
| userId | string                       | ✅       |                               |
| params | FetchUserTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`AccessTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
  "github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclient"

  "github.com/magicbell/magicbell-go-project-client/pkg/jwt"
)

config := magicbellprojectclientconfig.NewConfig()
client := magicbellprojectclient.NewMagicbellProjectClient(config)


params := jwt.FetchUserTokensRequestParams{

}

response, err := client.Jwt.FetchUserTokens(context.Background(), "userId", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
