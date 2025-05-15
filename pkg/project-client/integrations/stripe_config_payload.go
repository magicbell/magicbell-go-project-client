package integrations

import "encoding/json"

type StripeConfigPayload struct {
	// The signing secret to verify incoming requests from Stripe
	WebhookSigningSecret *string `json:"webhook_signing_secret,omitempty" required:"true" maxLength:"100" minLength:"1"`
}

func (s *StripeConfigPayload) GetWebhookSigningSecret() *string {
	if s == nil {
		return nil
	}
	return s.WebhookSigningSecret
}

func (s *StripeConfigPayload) SetWebhookSigningSecret(webhookSigningSecret string) {
	s.WebhookSigningSecret = &webhookSigningSecret
}

func (s StripeConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StripeConfigPayload to string"
	}
	return string(jsonData)
}
