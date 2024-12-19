package integrations

import (
	"encoding/json"
)

type GithubConfig struct {
	// The signing secret to verify incoming requests from Github
	WebhookSigningSecret *string `json:"webhook_signing_secret,omitempty" required:"true" maxLength:"100" minLength:"1"`
	touched              map[string]bool
}

func (g *GithubConfig) GetWebhookSigningSecret() *string {
	if g == nil {
		return nil
	}
	return g.WebhookSigningSecret
}

func (g *GithubConfig) SetWebhookSigningSecret(webhookSigningSecret string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["WebhookSigningSecret"] = true
	g.WebhookSigningSecret = &webhookSigningSecret
}

func (g *GithubConfig) SetWebhookSigningSecretNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["WebhookSigningSecret"] = true
	g.WebhookSigningSecret = nil
}
func (g GithubConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["WebhookSigningSecret"] && g.WebhookSigningSecret == nil {
		data["webhook_signing_secret"] = nil
	} else if g.WebhookSigningSecret != nil {
		data["webhook_signing_secret"] = g.WebhookSigningSecret
	}

	return json.Marshal(data)
}
