package integrations

import "encoding/json"

type PingConfigPayload struct {
	// URL to ping
	Url *string `json:"url,omitempty" required:"true" maxLength:"100" minLength:"1"`
}

func (p *PingConfigPayload) GetUrl() *string {
	if p == nil {
		return nil
	}
	return p.Url
}

func (p *PingConfigPayload) SetUrl(url string) {
	p.Url = &url
}

func (p PingConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PingConfigPayload to string"
	}
	return string(jsonData)
}
