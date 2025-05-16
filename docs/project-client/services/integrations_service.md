# IntegrationsService

A list of all methods in the `IntegrationsService` service. Click on the method name to view detailed information about that method.

| Methods                                                               | Description                                                                                                                                                             |
| :-------------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [ListIntegrations](#listintegrations)                                 | Lists all available and configured integrations for the project. Returns a summary of each integration including its type, status, and basic configuration information. |
| [GetApnsIntegration](#getapnsintegration)                             | Retrieves the current apns integration configurations for a specific integration type in the project. Returns configuration details and status information.             |
| [SaveApnsIntegration](#saveapnsintegration)                           | Creates or updates a apns integration for the project. Only administrators can configure integrations.                                                                  |
| [DeleteApnsIntegration](#deleteapnsintegration)                       | Removes a apns integration configuration from the project. This will disable the integration's functionality within the project.                                        |
| [DeleteApnsIntegrationById](#deleteapnsintegrationbyid)               | Removes a specific apns integration instance by ID from the project.                                                                                                    |
| [GetAwssnsIntegration](#getawssnsintegration)                         | Retrieves the current awssns integration configurations for a specific integration type in the project. Returns configuration details and status information.           |
| [SaveAwssnsIntegration](#saveawssnsintegration)                       | Creates or updates a awssns integration for the project. Only administrators can configure integrations.                                                                |
| [DeleteAwssnsIntegration](#deleteawssnsintegration)                   | Removes a awssns integration configuration from the project. This will disable the integration's functionality within the project.                                      |
| [DeleteAwssnsIntegrationById](#deleteawssnsintegrationbyid)           | Removes a specific awssns integration instance by ID from the project.                                                                                                  |
| [GetEventsourceIntegration](#geteventsourceintegration)               | Retrieves the current eventsource integration configurations for a specific integration type in the project. Returns configuration details and status information.      |
| [SaveEventsourceIntegration](#saveeventsourceintegration)             | Creates or updates a eventsource integration for the project. Only administrators can configure integrations.                                                           |
| [DeleteEventsourceIntegration](#deleteeventsourceintegration)         | Removes a eventsource integration configuration from the project. This will disable the integration's functionality within the project.                                 |
| [DeleteEventsourceIntegrationById](#deleteeventsourceintegrationbyid) | Removes a specific eventsource integration instance by ID from the project.                                                                                             |
| [GetExpoIntegration](#getexpointegration)                             | Retrieves the current expo integration configurations for a specific integration type in the project. Returns configuration details and status information.             |
| [SaveExpoIntegration](#saveexpointegration)                           | Creates or updates a expo integration for the project. Only administrators can configure integrations.                                                                  |
| [DeleteExpoIntegration](#deleteexpointegration)                       | Removes a expo integration configuration from the project. This will disable the integration's functionality within the project.                                        |
| [DeleteExpoIntegrationById](#deleteexpointegrationbyid)               | Removes a specific expo integration instance by ID from the project.                                                                                                    |
| [GetFcmIntegration](#getfcmintegration)                               | Retrieves the current fcm integration configurations for a specific integration type in the project. Returns configuration details and status information.              |
| [SaveFcmIntegration](#savefcmintegration)                             | Creates or updates a fcm integration for the project. Only administrators can configure integrations.                                                                   |
| [DeleteFcmIntegration](#deletefcmintegration)                         | Removes a fcm integration configuration from the project. This will disable the integration's functionality within the project.                                         |
| [DeleteFcmIntegrationById](#deletefcmintegrationbyid)                 | Removes a specific fcm integration instance by ID from the project.                                                                                                     |
| [GetGithubIntegration](#getgithubintegration)                         | Retrieves the current github integration configurations for a specific integration type in the project. Returns configuration details and status information.           |
| [SaveGithubIntegration](#savegithubintegration)                       | Creates or updates a github integration for the project. Only administrators can configure integrations.                                                                |
| [DeleteGithubIntegration](#deletegithubintegration)                   | Removes a github integration configuration from the project. This will disable the integration's functionality within the project.                                      |
| [DeleteGithubIntegrationById](#deletegithubintegrationbyid)           | Removes a specific github integration instance by ID from the project.                                                                                                  |
| [GetInboxIntegration](#getinboxintegration)                           | Retrieves the current inbox integration configurations for a specific integration type in the project. Returns configuration details and status information.            |
| [SaveInboxIntegration](#saveinboxintegration)                         | Creates or updates a inbox integration for the project. Only administrators can configure integrations.                                                                 |
| [DeleteInboxIntegration](#deleteinboxintegration)                     | Removes a inbox integration configuration from the project. This will disable the integration's functionality within the project.                                       |
| [DeleteInboxIntegrationById](#deleteinboxintegrationbyid)             | Removes a specific inbox integration instance by ID from the project.                                                                                                   |
| [GetMailgunIntegration](#getmailgunintegration)                       | Retrieves the current mailgun integration configurations for a specific integration type in the project. Returns configuration details and status information.          |
| [SaveMailgunIntegration](#savemailgunintegration)                     | Creates or updates a mailgun integration for the project. Only administrators can configure integrations.                                                               |
| [DeleteMailgunIntegration](#deletemailgunintegration)                 | Removes a mailgun integration configuration from the project. This will disable the integration's functionality within the project.                                     |
| [DeleteMailgunIntegrationById](#deletemailgunintegrationbyid)         | Removes a specific mailgun integration instance by ID from the project.                                                                                                 |
| [GetPingEmailIntegration](#getpingemailintegration)                   | Retrieves the current ping_email integration configurations for a specific integration type in the project. Returns configuration details and status information.       |
| [SavePingEmailIntegration](#savepingemailintegration)                 | Creates or updates a ping_email integration for the project. Only administrators can configure integrations.                                                            |
| [DeletePingEmailIntegration](#deletepingemailintegration)             | Removes a ping_email integration configuration from the project. This will disable the integration's functionality within the project.                                  |
| [DeletePingEmailIntegrationById](#deletepingemailintegrationbyid)     | Removes a specific ping_email integration instance by ID from the project.                                                                                              |
| [GetSendgridIntegration](#getsendgridintegration)                     | Retrieves the current sendgrid integration configurations for a specific integration type in the project. Returns configuration details and status information.         |
| [SaveSendgridIntegration](#savesendgridintegration)                   | Creates or updates a sendgrid integration for the project. Only administrators can configure integrations.                                                              |
| [DeleteSendgridIntegration](#deletesendgridintegration)               | Removes a sendgrid integration configuration from the project. This will disable the integration's functionality within the project.                                    |
| [DeleteSendgridIntegrationById](#deletesendgridintegrationbyid)       | Removes a specific sendgrid integration instance by ID from the project.                                                                                                |
| [GetSesIntegration](#getsesintegration)                               | Retrieves the current ses integration configurations for a specific integration type in the project. Returns configuration details and status information.              |
| [SaveSesIntegration](#savesesintegration)                             | Creates or updates a ses integration for the project. Only administrators can configure integrations.                                                                   |
| [DeleteSesIntegration](#deletesesintegration)                         | Removes a ses integration configuration from the project. This will disable the integration's functionality within the project.                                         |
| [DeleteSesIntegrationById](#deletesesintegrationbyid)                 | Removes a specific ses integration instance by ID from the project.                                                                                                     |
| [GetSlackIntegration](#getslackintegration)                           | Retrieves the current slack integration configurations for a specific integration type in the project. Returns configuration details and status information.            |
| [SaveSlackIntegration](#saveslackintegration)                         | Creates or updates a slack integration for the project. Only administrators can configure integrations.                                                                 |
| [DeleteSlackIntegration](#deleteslackintegration)                     | Removes a slack integration configuration from the project. This will disable the integration's functionality within the project.                                       |
| [DeleteSlackIntegrationById](#deleteslackintegrationbyid)             | Removes a specific slack integration instance by ID from the project.                                                                                                   |
| [GetStripeIntegration](#getstripeintegration)                         | Retrieves the current stripe integration configurations for a specific integration type in the project. Returns configuration details and status information.           |
| [SaveStripeIntegration](#savestripeintegration)                       | Creates or updates a stripe integration for the project. Only administrators can configure integrations.                                                                |
| [DeleteStripeIntegration](#deletestripeintegration)                   | Removes a stripe integration configuration from the project. This will disable the integration's functionality within the project.                                      |
| [DeleteStripeIntegrationById](#deletestripeintegrationbyid)           | Removes a specific stripe integration instance by ID from the project.                                                                                                  |
| [GetTemplatesIntegration](#gettemplatesintegration)                   | Retrieves the current templates integration configurations for a specific integration type in the project. Returns configuration details and status information.        |
| [SaveTemplatesIntegration](#savetemplatesintegration)                 | Creates or updates a templates integration for the project. Only administrators can configure integrations.                                                             |
| [DeleteTemplatesIntegration](#deletetemplatesintegration)             | Removes a templates integration configuration from the project. This will disable the integration's functionality within the project.                                   |
| [DeleteTemplatesIntegrationById](#deletetemplatesintegrationbyid)     | Removes a specific templates integration instance by ID from the project.                                                                                               |
| [GetTwilioIntegration](#gettwiliointegration)                         | Retrieves the current twilio integration configurations for a specific integration type in the project. Returns configuration details and status information.           |
| [SaveTwilioIntegration](#savetwiliointegration)                       | Creates or updates a twilio integration for the project. Only administrators can configure integrations.                                                                |
| [DeleteTwilioIntegration](#deletetwiliointegration)                   | Removes a twilio integration configuration from the project. This will disable the integration's functionality within the project.                                      |
| [DeleteTwilioIntegrationById](#deletetwiliointegrationbyid)           | Removes a specific twilio integration instance by ID from the project.                                                                                                  |
| [GetWebPushIntegration](#getwebpushintegration)                       | Retrieves the current web_push integration configurations for a specific integration type in the project. Returns configuration details and status information.         |
| [SaveWebPushIntegration](#savewebpushintegration)                     | Creates or updates a web_push integration for the project. Only administrators can configure integrations.                                                              |
| [DeleteWebPushIntegration](#deletewebpushintegration)                 | Removes a web_push integration configuration from the project. This will disable the integration's functionality within the project.                                    |
| [DeleteWebPushIntegrationById](#deletewebpushintegrationbyid)         | Removes a specific web_push integration instance by ID from the project.                                                                                                |

## ListIntegrations

Lists all available and configured integrations for the project. Returns a summary of each integration including its type, status, and basic configuration information.

- HTTP Method: `GET`
- Endpoint: `/integrations`

**Parameters**

| Name   | Type                          | Required | Description                   |
| :----- | :---------------------------- | :------- | :---------------------------- |
| ctx    | Context                       | ✅       | Default go language context   |
| params | ListIntegrationsRequestParams | ✅       | Additional request parameters |

**Return Type**

`IntegrationConfigCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.ListIntegrationsRequestParams{

}

response, err := client.Integrations.ListIntegrations(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetApnsIntegration

Retrieves the current apns integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/apns`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`ApnsConfigCollection`

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

response, err := client.Integrations.GetApnsIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveApnsIntegration

Creates or updates a apns integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/apns`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| apnsConfigPayload | ApnsConfigPayload | ✅       |                             |

**Return Type**

`ApnsConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)

badge := integrations.BADGE_UNREAD

payloadVersion := integrations.PAYLOAD_VERSION_1

request := integrations.ApnsConfigPayload{
  AppId: util.ToPointer("AppId"),
  Badge: &badge,
  Certificate: util.ToPointer("Certificate"),
  KeyId: util.ToPointer("KeyId"),
  PayloadVersion: &payloadVersion,
  TeamId: util.ToPointer("TeamId"),
}

response, err := client.Integrations.SaveApnsIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteApnsIntegration

Removes a apns integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/apns`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteApnsIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteApnsIntegrationById

Removes a specific apns integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/apns/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteApnsIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetAwssnsIntegration

Retrieves the current awssns integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/awssns`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`AwssnsConfigCollection`

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

response, err := client.Integrations.GetAwssnsIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveAwssnsIntegration

Creates or updates a awssns integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/awssns`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| awssnsConfigPayload | AwssnsConfigPayload | ✅       |                             |

**Return Type**

`AwssnsConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.AwssnsConfigPayload{
  WebhookSigningSecret: util.ToPointer("WebhookSigningSecret"),
}

response, err := client.Integrations.SaveAwssnsIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteAwssnsIntegration

Removes a awssns integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/awssns`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteAwssnsIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteAwssnsIntegrationById

Removes a specific awssns integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/awssns/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteAwssnsIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEventsourceIntegration

Retrieves the current eventsource integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/eventsource`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`EventSourceConfigCollection`

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

response, err := client.Integrations.GetEventsourceIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveEventsourceIntegration

Creates or updates a eventsource integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/eventsource`

**Parameters**

| Name                     | Type                     | Required | Description                 |
| :----------------------- | :----------------------- | :------- | :-------------------------- |
| ctx                      | Context                  | ✅       | Default go language context |
| eventSourceConfigPayload | EventSourceConfigPayload | ✅       |                             |

**Return Type**

`EventSourceConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.EventSourceConfigPayload{
  Source: util.ToPointer("Source"),
}

response, err := client.Integrations.SaveEventsourceIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteEventsourceIntegration

Removes a eventsource integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/eventsource`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteEventsourceIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteEventsourceIntegrationById

Removes a specific eventsource integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/eventsource/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteEventsourceIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetExpoIntegration

Retrieves the current expo integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/expo`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`ExpoConfigCollection`

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

response, err := client.Integrations.GetExpoIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveExpoIntegration

Creates or updates a expo integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/expo`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| expoConfigPayload | ExpoConfigPayload | ✅       |                             |

**Return Type**

`ExpoConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.ExpoConfigPayload{
  AccessToken: util.ToPointer("AccessToken"),
}

response, err := client.Integrations.SaveExpoIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteExpoIntegration

Removes a expo integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/expo`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteExpoIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteExpoIntegrationById

Removes a specific expo integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/expo/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteExpoIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetFcmIntegration

Retrieves the current fcm integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/fcm`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`FcmConfigCollection`

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

response, err := client.Integrations.GetFcmIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveFcmIntegration

Creates or updates a fcm integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/fcm`

**Parameters**

| Name             | Type             | Required | Description                 |
| :--------------- | :--------------- | :------- | :-------------------------- |
| ctx              | Context          | ✅       | Default go language context |
| fcmConfigPayload | FcmConfigPayload | ✅       |                             |

**Return Type**

`FcmConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)

type_ := integrations.TYPE_SERVICE_ACCOUNT

request := integrations.FcmConfigPayload{
  AuthProviderX509CertUrl: util.ToPointer("AuthProviderX509CertUrl"),
  AuthUri: util.ToPointer("AuthUri"),
  ClientEmail: util.ToPointer("ClientEmail"),
  ClientId: util.ToPointer("ClientId"),
  ClientX509CertUrl: util.ToPointer("ClientX509CertUrl"),
  PrivateKey: util.ToPointer("PrivateKey"),
  PrivateKeyId: util.ToPointer("PrivateKeyId"),
  ProjectId: util.ToPointer("ProjectId"),
  TokenUri: util.ToPointer("TokenUri"),
  Type_: &type_,
  UniverseDomain: util.ToPointer("UniverseDomain"),
}

response, err := client.Integrations.SaveFcmIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteFcmIntegration

Removes a fcm integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/fcm`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteFcmIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteFcmIntegrationById

Removes a specific fcm integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/fcm/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteFcmIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetGithubIntegration

Retrieves the current github integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/github`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`GithubConfigCollection`

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

response, err := client.Integrations.GetGithubIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveGithubIntegration

Creates or updates a github integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/github`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| githubConfigPayload | GithubConfigPayload | ✅       |                             |

**Return Type**

`GithubConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.GithubConfigPayload{
  WebhookSigningSecret: util.ToPointer("WebhookSigningSecret"),
}

response, err := client.Integrations.SaveGithubIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteGithubIntegration

Removes a github integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/github`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteGithubIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteGithubIntegrationById

Removes a specific github integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/github/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteGithubIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetInboxIntegration

Retrieves the current inbox integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/inbox`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`InboxConfigCollection`

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

response, err := client.Integrations.GetInboxIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveInboxIntegration

Creates or updates a inbox integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/inbox`

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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
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

notification := integrations.Notification{
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
  Notification: &notification,
  UnseenBadge: &unseenBadge,
}

request := integrations.InboxConfigPayload{
  Images: &images,
  Locale: util.ToPointer(util.Nullable[string]{ Value: "Locale" }),
  Theme: &theme,
}

response, err := client.Integrations.SaveInboxIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteInboxIntegration

Removes a inbox integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/inbox`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteInboxIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteInboxIntegrationById

Removes a specific inbox integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/inbox/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteInboxIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetMailgunIntegration

Retrieves the current mailgun integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/mailgun`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`MailgunConfigCollection`

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

response, err := client.Integrations.GetMailgunIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveMailgunIntegration

Creates or updates a mailgun integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/mailgun`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| mailgunConfigPayload | MailgunConfigPayload | ✅       |                             |

**Return Type**

`MailgunConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


mailgunConfigPayloadFrom := integrations.MailgunConfigPayloadFrom{
  Email: util.ToPointer("Email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "Name" }),
}

region := integrations.REGION_US

request := integrations.MailgunConfigPayload{
  ApiKey: util.ToPointer("ApiKey"),
  Domain: util.ToPointer("Domain"),
  From: &mailgunConfigPayloadFrom,
  Region: &region,
}

response, err := client.Integrations.SaveMailgunIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteMailgunIntegration

Removes a mailgun integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/mailgun`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteMailgunIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteMailgunIntegrationById

Removes a specific mailgun integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/mailgun/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteMailgunIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetPingEmailIntegration

Retrieves the current ping_email integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/ping_email`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`PingConfigCollection`

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

response, err := client.Integrations.GetPingEmailIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SavePingEmailIntegration

Creates or updates a ping_email integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/ping_email`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| pingConfigPayload | PingConfigPayload | ✅       |                             |

**Return Type**

`PingConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.PingConfigPayload{
  Url: util.ToPointer("Url"),
}

response, err := client.Integrations.SavePingEmailIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeletePingEmailIntegration

Removes a ping_email integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/ping_email`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeletePingEmailIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeletePingEmailIntegrationById

Removes a specific ping_email integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/ping_email/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeletePingEmailIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSendgridIntegration

Retrieves the current sendgrid integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/sendgrid`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`SendgridConfigCollection`

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

response, err := client.Integrations.GetSendgridIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSendgridIntegration

Creates or updates a sendgrid integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/sendgrid`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| sendgridConfigPayload | SendgridConfigPayload | ✅       |                             |

**Return Type**

`SendgridConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


sendgridConfigPayloadFrom := integrations.SendgridConfigPayloadFrom{
  Email: util.ToPointer("Email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "Name" }),
}


replyTo := integrations.ReplyTo{
  Email: util.ToPointer("Email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "Name" }),
}

request := integrations.SendgridConfigPayload{
  ApiKey: util.ToPointer("ApiKey"),
  From: &sendgridConfigPayloadFrom,
  ReplyTo: &replyTo,
}

response, err := client.Integrations.SaveSendgridIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSendgridIntegration

Removes a sendgrid integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/sendgrid`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteSendgridIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSendgridIntegrationById

Removes a specific sendgrid integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/sendgrid/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteSendgridIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSesIntegration

Retrieves the current ses integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/ses`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`SesConfigCollection`

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

response, err := client.Integrations.GetSesIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSesIntegration

Creates or updates a ses integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/ses`

**Parameters**

| Name             | Type             | Required | Description                 |
| :--------------- | :--------------- | :------- | :-------------------------- |
| ctx              | Context          | ✅       | Default go language context |
| sesConfigPayload | SesConfigPayload | ✅       |                             |

**Return Type**

`SesConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


sesConfigPayloadFrom := integrations.SesConfigPayloadFrom{
  Email: util.ToPointer("Email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "Name" }),
}

request := integrations.SesConfigPayload{
  From: &sesConfigPayloadFrom,
  KeyId: util.ToPointer("KeyId"),
  Region: util.ToPointer("Region"),
  SecretKey: util.ToPointer("SecretKey"),
}

response, err := client.Integrations.SaveSesIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSesIntegration

Removes a ses integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/ses`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteSesIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSesIntegrationById

Removes a specific ses integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/ses/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteSesIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetSlackIntegration

Retrieves the current slack integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/slack`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`SlackConfigCollection`

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

response, err := client.Integrations.GetSlackIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSlackIntegration

Creates or updates a slack integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/slack`

**Parameters**

| Name               | Type               | Required | Description                 |
| :----------------- | :----------------- | :------- | :-------------------------- |
| ctx                | Context            | ✅       | Default go language context |
| slackConfigPayload | SlackConfigPayload | ✅       |                             |

**Return Type**

`SlackConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.SlackConfigPayload{
  AppId: util.ToPointer("AppId"),
  ClientId: util.ToPointer("ClientId"),
  ClientSecret: util.ToPointer("ClientSecret"),
  SigningSecret: util.ToPointer("SigningSecret"),
}

response, err := client.Integrations.SaveSlackIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSlackIntegration

Removes a slack integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/slack`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteSlackIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSlackIntegrationById

Removes a specific slack integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/slack/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteSlackIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetStripeIntegration

Retrieves the current stripe integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/stripe`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`StripeConfigCollection`

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

response, err := client.Integrations.GetStripeIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveStripeIntegration

Creates or updates a stripe integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/stripe`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| stripeConfigPayload | StripeConfigPayload | ✅       |                             |

**Return Type**

`StripeConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.StripeConfigPayload{
  WebhookSigningSecret: util.ToPointer("WebhookSigningSecret"),
}

response, err := client.Integrations.SaveStripeIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteStripeIntegration

Removes a stripe integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/stripe`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteStripeIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteStripeIntegrationById

Removes a specific stripe integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/stripe/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteStripeIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTemplatesIntegration

Retrieves the current templates integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/templates`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`TemplatesConfigCollection`

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

response, err := client.Integrations.GetTemplatesIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveTemplatesIntegration

Creates or updates a templates integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/templates`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| body | []byte  | ✅       |                             |

**Return Type**

`[]byte`

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

response, err := client.Integrations.SaveTemplatesIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteTemplatesIntegration

Removes a templates integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/templates`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteTemplatesIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteTemplatesIntegrationById

Removes a specific templates integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/templates/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteTemplatesIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTwilioIntegration

Retrieves the current twilio integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/twilio`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`TwilioConfigCollection`

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

response, err := client.Integrations.GetTwilioIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveTwilioIntegration

Creates or updates a twilio integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/twilio`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| twilioConfigPayload | TwilioConfigPayload | ✅       |                             |

**Return Type**

`TwilioConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.TwilioConfigPayload{
  AccountSid: util.ToPointer("AccountSid"),
  ApiKey: util.ToPointer("ApiKey"),
  ApiSecret: util.ToPointer("ApiSecret"),
  From: util.ToPointer("From"),
}

response, err := client.Integrations.SaveTwilioIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteTwilioIntegration

Removes a twilio integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/twilio`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteTwilioIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteTwilioIntegrationById

Removes a specific twilio integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/twilio/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteTwilioIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetWebPushIntegration

Retrieves the current web_push integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/web_push`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`WebpushConfigCollection`

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

response, err := client.Integrations.GetWebPushIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveWebPushIntegration

Creates or updates a web_push integration for the project. Only administrators can configure integrations.

- HTTP Method: `PUT`
- Endpoint: `/integrations/web_push`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| webpushConfigPayload | WebpushConfigPayload | ✅       |                             |

**Return Type**

`WebpushConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.WebpushConfigPayload{
  PrivateKey: util.ToPointer("PrivateKey"),
  PublicKey: util.ToPointer("PublicKey"),
}

response, err := client.Integrations.SaveWebPushIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteWebPushIntegration

Removes a web_push integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/web_push`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteWebPushIntegration(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteWebPushIntegrationById

Removes a specific web_push integration instance by ID from the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/web_push/{id}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| id   | string  | ✅       |                             |

**Return Type**

`any`

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

response, err := client.Integrations.DeleteWebPushIntegrationById(context.Background(), "id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
