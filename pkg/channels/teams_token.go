package channels

import (
	"encoding/json"
)

type TeamsToken struct {
	Webhook *TeamsTokenWebhook `json:"webhook,omitempty"`
	touched map[string]bool
}

func (t *TeamsToken) GetWebhook() *TeamsTokenWebhook {
	if t == nil {
		return nil
	}
	return t.Webhook
}

func (t *TeamsToken) SetWebhook(webhook TeamsTokenWebhook) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Webhook"] = true
	t.Webhook = &webhook
}

func (t *TeamsToken) SetWebhookNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Webhook"] = true
	t.Webhook = nil
}
func (t TeamsToken) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Webhook"] && t.Webhook == nil {
		data["webhook"] = nil
	} else if t.Webhook != nil {
		data["webhook"] = t.Webhook
	}

	return json.Marshal(data)
}

type TeamsTokenWebhook struct {
	Url     *string `json:"url,omitempty"`
	touched map[string]bool
}

func (t *TeamsTokenWebhook) GetUrl() *string {
	if t == nil {
		return nil
	}
	return t.Url
}

func (t *TeamsTokenWebhook) SetUrl(url string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Url"] = true
	t.Url = &url
}

func (t *TeamsTokenWebhook) SetUrlNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Url"] = true
	t.Url = nil
}
func (t TeamsTokenWebhook) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Url"] && t.Url == nil {
		data["url"] = nil
	} else if t.Url != nil {
		data["url"] = t.Url
	}

	return json.Marshal(data)
}
