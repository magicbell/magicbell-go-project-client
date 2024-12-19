# ProjectDeliveryConfig

**Properties**

| Name     | Type                                     | Required | Description |
| :------- | :--------------------------------------- | :------- | :---------- |
| Channels | []channels.ProjectDeliveryConfigChannels | ✅       |             |

# ProjectDeliveryConfigChannels

**Properties**

| Name     | Type                      | Required | Description                                                                       |
| :------- | :------------------------ | :------- | :-------------------------------------------------------------------------------- |
| Channel  | channels.ChannelsChannel1 | ✅       |                                                                                   |
| Delay    | int64                     | ❌       | Delay (in seconds) since the last step, before the message is sent to the channel |
| Disabled | bool                      | ❌       |                                                                                   |
| If\_     | string                    | ❌       |                                                                                   |
| Priority | int64                     | ❌       |                                                                                   |

# ChannelsChannel1

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
