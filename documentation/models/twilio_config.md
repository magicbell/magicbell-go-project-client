# TwilioConfig

**Properties**

| Name       | Type                            | Required | Description                                     |
| :--------- | :------------------------------ | :------- | :---------------------------------------------- |
| AccountSid | string                          | ✅       | The SID for your Twilio account                 |
| ApiKey     | string                          | ✅       | The API key for Twilio                          |
| ApiSecret  | string                          | ✅       | The API Secret for Twilio                       |
| From       | string                          | ✅       | The phone number to send from, in E.164 format  |
| Region     | integrations.TwilioConfigRegion | ❌       | The region to use for Twilio, defaults to 'us1' |

# TwilioConfigRegion

The region to use for Twilio, defaults to 'us1'

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| us1  | string | ✅       | "us1"       |
| ie1  | string | ✅       | "ie1"       |
| au1  | string | ✅       | "au1"       |
