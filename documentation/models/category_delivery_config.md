# CategoryDeliveryConfig

**Properties**

| Name     | Type                                      | Required | Description |
| :------- | :---------------------------------------- | :------- | :---------- |
| Category | string                                    | ✅       |             |
| Channels | []channels.CategoryDeliveryConfigChannels | ✅       |             |
| Disabled | bool                                      | ❌       |             |

# CategoryDeliveryConfigChannels

**Properties**

| Name     | Type                      | Required | Description |
| :------- | :------------------------ | :------- | :---------- |
| Channel  | channels.ChannelsChannel2 | ✅       |             |
| Delay    | int64                     | ❌       |             |
| Disabled | bool                      | ❌       |             |
| If\_     | string                    | ❌       |             |
| Priority | int64                     | ❌       |             |

# ChannelsChannel2

**Properties**

| Name        | Type   | Required | Description   |
| :---------- | :----- | :------- | :------------ |
| in_app      | string | ✅       | "in_app"      |
| slack       | string | ✅       | "slack"       |
| web_push    | string | ✅       | "web_push"    |
| mobile_push | string | ✅       | "mobile_push" |
| teams       | string | ✅       | "teams"       |
| email       | string | ✅       | "email"       |
