package integrations

import (
	"encoding/json"
)

type AwssnsConfig struct {
	// The signing certificate from AWS SNS
	WebhookSigningSecret *string `json:"webhook_signing_secret,omitempty" required:"true" minLength:"1"`
	touched              map[string]bool
}

func (a *AwssnsConfig) GetWebhookSigningSecret() *string {
	if a == nil {
		return nil
	}
	return a.WebhookSigningSecret
}

func (a *AwssnsConfig) SetWebhookSigningSecret(webhookSigningSecret string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["WebhookSigningSecret"] = true
	a.WebhookSigningSecret = &webhookSigningSecret
}

func (a *AwssnsConfig) SetWebhookSigningSecretNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["WebhookSigningSecret"] = true
	a.WebhookSigningSecret = nil
}
func (a AwssnsConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["WebhookSigningSecret"] && a.WebhookSigningSecret == nil {
		data["webhook_signing_secret"] = nil
	} else if a.WebhookSigningSecret != nil {
		data["webhook_signing_secret"] = a.WebhookSigningSecret
	}

	return json.Marshal(data)
}
