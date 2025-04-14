# CategoryDeliveryConfig

**Properties**

| Name     | Type                                      | Required | Description |
| :------- | :---------------------------------------- | :------- | :---------- |
| Channels | []channels.CategoryDeliveryConfigChannels | ✅       |             |
| Key      | string                                    | ✅       |             |
| Disabled | bool                                      | ❌       |             |

# CategoryDeliveryConfigChannels

**Properties**

| Name    | Type             | Required | Description |
| :------ | :--------------- | :------- | :---------- |
| Channel | channels.Channel | ✅       |             |
| Delay   | int64            | ❌       |             |
| If\_    | string           | ❌       |             |

# Channel

**Properties**

| Name        | Type   | Required | Description   |
| :---------- | :----- | :------- | :------------ |
| in_app      | string | ✅       | "in_app"      |
| slack       | string | ✅       | "slack"       |
| web_push    | string | ✅       | "web_push"    |
| mobile_push | string | ✅       | "mobile_push" |
| teams       | string | ✅       | "teams"       |
| email       | string | ✅       | "email"       |
| sms         | string | ✅       | "sms"         |
