package integrations

import "encoding/json"

type AwssnsConfigPayload struct {
	// The signing certificate from AWS SNS
	WebhookSigningSecret *string `json:"webhook_signing_secret,omitempty" required:"true" minLength:"1"`
}

func (a *AwssnsConfigPayload) GetWebhookSigningSecret() *string {
	if a == nil {
		return nil
	}
	return a.WebhookSigningSecret
}

func (a *AwssnsConfigPayload) SetWebhookSigningSecret(webhookSigningSecret string) {
	a.WebhookSigningSecret = &webhookSigningSecret
}

func (a AwssnsConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AwssnsConfigPayload to string"
	}
	return string(jsonData)
}
