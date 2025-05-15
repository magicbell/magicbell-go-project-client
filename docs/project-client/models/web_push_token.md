# WebPushToken

**Properties**

| Name        | Type          | Required | Description                                                                                                                                                                                      |
| :---------- | :------------ | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CreatedAt   | string        | ✅       |                                                                                                                                                                                                  |
| Endpoint    | string        | ✅       | The push subscription URL obtained from PushSubscription.endpoint after calling registration.pushManager.subscribe(). This is the unique URL for this device that push messages will be sent to. |
| Id          | string        | ✅       |                                                                                                                                                                                                  |
| Keys        | channels.Keys | ✅       | The encryption keys from the PushSubscription.getKey() method, needed to encrypt push messages for this subscription.                                                                            |
| DiscardedAt | string        | ❌       |                                                                                                                                                                                                  |
| UpdatedAt   | string        | ❌       |                                                                                                                                                                                                  |

# Keys

The encryption keys from the PushSubscription.getKey() method, needed to encrypt push messages for this subscription.

**Properties**

| Name   | Type   | Required | Description                                                                                                                     |
| :----- | :----- | :------- | :------------------------------------------------------------------------------------------------------------------------------ |
| Auth   | string | ✅       | The authentication secret obtained from PushSubscription.getKey('auth'). Used to encrypt push messages for this subscription.   |
| P256dh | string | ✅       | The P-256 ECDH public key obtained from PushSubscription.getKey('p256dh'). Used to encrypt push messages for this subscription. |
