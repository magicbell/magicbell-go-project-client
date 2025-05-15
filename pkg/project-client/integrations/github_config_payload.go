package integrations

import "encoding/json"

type GithubConfigPayload struct {
	// The signing secret to verify incoming requests from Github
	WebhookSigningSecret *string `json:"webhook_signing_secret,omitempty" required:"true" maxLength:"100" minLength:"1"`
}

func (g *GithubConfigPayload) GetWebhookSigningSecret() *string {
	if g == nil {
		return nil
	}
	return g.WebhookSigningSecret
}

func (g *GithubConfigPayload) SetWebhookSigningSecret(webhookSigningSecret string) {
	g.WebhookSigningSecret = &webhookSigningSecret
}

func (g GithubConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GithubConfigPayload to string"
	}
	return string(jsonData)
}
