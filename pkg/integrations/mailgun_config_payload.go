package integrations

import "encoding/json"

type MailgunConfigPayload struct {
	ApiKey *string `json:"api_key,omitempty" required:"true" minLength:"1"`
	Domain *string `json:"domain,omitempty" required:"true" minLength:"1"`
	Region *Region `json:"region,omitempty" required:"true"`
}

func (m *MailgunConfigPayload) GetApiKey() *string {
	if m == nil {
		return nil
	}
	return m.ApiKey
}

func (m *MailgunConfigPayload) SetApiKey(apiKey string) {
	m.ApiKey = &apiKey
}

func (m *MailgunConfigPayload) GetDomain() *string {
	if m == nil {
		return nil
	}
	return m.Domain
}

func (m *MailgunConfigPayload) SetDomain(domain string) {
	m.Domain = &domain
}

func (m *MailgunConfigPayload) GetRegion() *Region {
	if m == nil {
		return nil
	}
	return m.Region
}

func (m *MailgunConfigPayload) SetRegion(region Region) {
	m.Region = &region
}

func (m MailgunConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MailgunConfigPayload to string"
	}
	return string(jsonData)
}

type Region string

const (
	REGION_US Region = "us"
	REGION_EU Region = "eu"
)
