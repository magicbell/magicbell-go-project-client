# SlackFinishInstallResponse

**Properties**

| Name        | Type   | Required | Description                                                                                        |
| :---------- | :----- | :------- | :------------------------------------------------------------------------------------------------- |
| AppId       | string | ✅       | The app ID of the Slack app that was originally configured at the project-level.                   |
| Code        | string | ✅       | The code that was returned from the OAuth flow, and found in the query string of the redirect URL. |
| RedirectUrl | string | ❌       |                                                                                                    |
