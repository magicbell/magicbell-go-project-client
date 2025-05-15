# CreateUserTokenRequest

**Properties**

| Name       | Type   | Required | Description                                                                                                                                                                                                       |
| :--------- | :----- | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Email      | string | ❌       | The user's email.                                                                                                                                                                                                 |
| Expiry     | int64  | ❌       | The duration for which the token is valid (in seconds)                                                                                                                                                            |
| ExternalId | string | ❌       | A unique string that MagicBell can utilize to identify the user uniquely. We recommend setting this attribute to the ID of the user in your database. Provide the external id if the user's email is unavailable. |
| Name       | string | ❌       | The name of the token.                                                                                                                                                                                            |
