package integrations

import (
	"encoding/json"
)

type MailgunConfig struct {
	ApiKey  *string              `json:"api_key,omitempty" required:"true" minLength:"1"`
	Domain  *string              `json:"domain,omitempty" required:"true" minLength:"1"`
	Region  *MailgunConfigRegion `json:"region,omitempty" required:"true"`
	touched map[string]bool
}

func (m *MailgunConfig) GetApiKey() *string {
	if m == nil {
		return nil
	}
	return m.ApiKey
}

func (m *MailgunConfig) SetApiKey(apiKey string) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["ApiKey"] = true
	m.ApiKey = &apiKey
}

func (m *MailgunConfig) SetApiKeyNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["ApiKey"] = true
	m.ApiKey = nil
}

func (m *MailgunConfig) GetDomain() *string {
	if m == nil {
		return nil
	}
	return m.Domain
}

func (m *MailgunConfig) SetDomain(domain string) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Domain"] = true
	m.Domain = &domain
}

func (m *MailgunConfig) SetDomainNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Domain"] = true
	m.Domain = nil
}

func (m *MailgunConfig) GetRegion() *MailgunConfigRegion {
	if m == nil {
		return nil
	}
	return m.Region
}

func (m *MailgunConfig) SetRegion(region MailgunConfigRegion) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Region"] = true
	m.Region = &region
}

func (m *MailgunConfig) SetRegionNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Region"] = true
	m.Region = nil
}
func (m MailgunConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if m.touched["ApiKey"] && m.ApiKey == nil {
		data["api_key"] = nil
	} else if m.ApiKey != nil {
		data["api_key"] = m.ApiKey
	}

	if m.touched["Domain"] && m.Domain == nil {
		data["domain"] = nil
	} else if m.Domain != nil {
		data["domain"] = m.Domain
	}

	if m.touched["Region"] && m.Region == nil {
		data["region"] = nil
	} else if m.Region != nil {
		data["region"] = m.Region
	}

	return json.Marshal(data)
}

type MailgunConfigRegion string

const (
	MAILGUN_CONFIG_REGION_US MailgunConfigRegion = "us"
	MAILGUN_CONFIG_REGION_EU MailgunConfigRegion = "eu"
)
