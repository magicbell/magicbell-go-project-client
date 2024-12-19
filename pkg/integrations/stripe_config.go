package integrations

import (
	"encoding/json"
)

type StripeConfig struct {
	// The signing secret to verify incoming requests from Stripe
	WebhookSigningSecret *string `json:"webhook_signing_secret,omitempty" required:"true" maxLength:"100" minLength:"1"`
	touched              map[string]bool
}

func (s *StripeConfig) GetWebhookSigningSecret() *string {
	if s == nil {
		return nil
	}
	return s.WebhookSigningSecret
}

func (s *StripeConfig) SetWebhookSigningSecret(webhookSigningSecret string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["WebhookSigningSecret"] = true
	s.WebhookSigningSecret = &webhookSigningSecret
}

func (s *StripeConfig) SetWebhookSigningSecretNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["WebhookSigningSecret"] = true
	s.WebhookSigningSecret = nil
}
func (s StripeConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["WebhookSigningSecret"] && s.WebhookSigningSecret == nil {
		data["webhook_signing_secret"] = nil
	} else if s.WebhookSigningSecret != nil {
		data["webhook_signing_secret"] = s.WebhookSigningSecret
	}

	return json.Marshal(data)
}
