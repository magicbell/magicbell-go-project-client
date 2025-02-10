# WebPushTokenResponse

**Properties**

| Name        | Type          | Required | Description |
| :---------- | :------------ | :------- | :---------- |
| CreatedAt   | string        | ✅       |             |
| Endpoint    | string        | ✅       |             |
| Id          | string        | ✅       |             |
| Keys        | channels.Keys | ✅       |             |
| DiscardedAt | string        | ❌       |             |
| UpdatedAt   | string        | ❌       |             |

# Keys

**Properties**

| Name   | Type   | Required | Description |
| :----- | :----- | :------- | :---------- |
| Auth   | string | ✅       |             |
| P256dh | string | ✅       |             |
