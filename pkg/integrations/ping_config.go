package integrations

import (
	"encoding/json"
)

type PingConfig struct {
	// URL to ping
	Url     *string `json:"url,omitempty" required:"true" maxLength:"100" minLength:"1"`
	touched map[string]bool
}

func (p *PingConfig) GetUrl() *string {
	if p == nil {
		return nil
	}
	return p.Url
}

func (p *PingConfig) SetUrl(url string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Url"] = true
	p.Url = &url
}

func (p *PingConfig) SetUrlNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Url"] = true
	p.Url = nil
}
func (p PingConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["Url"] && p.Url == nil {
		data["url"] = nil
	} else if p.Url != nil {
		data["url"] = p.Url
	}

	return json.Marshal(data)
}
