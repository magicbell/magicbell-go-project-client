package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type MailgunConfigPayload struct {
	ApiKey *string                   `json:"api_key,omitempty" required:"true" minLength:"1"`
	Domain *string                   `json:"domain,omitempty" required:"true" minLength:"1"`
	From   *MailgunConfigPayloadFrom `json:"from,omitempty"`
	Region *Region                   `json:"region,omitempty" required:"true"`
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

func (m *MailgunConfigPayload) GetFrom() *MailgunConfigPayloadFrom {
	if m == nil {
		return nil
	}
	return m.From
}

func (m *MailgunConfigPayload) SetFrom(from MailgunConfigPayloadFrom) {
	m.From = &from
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

type MailgunConfigPayloadFrom struct {
	// The email address to send from
	Email *string `json:"email,omitempty" required:"true"`
	// The name to send from
	Name *util.Nullable[string] `json:"name,omitempty"`
}

func (m *MailgunConfigPayloadFrom) GetEmail() *string {
	if m == nil {
		return nil
	}
	return m.Email
}

func (m *MailgunConfigPayloadFrom) SetEmail(email string) {
	m.Email = &email
}

func (m *MailgunConfigPayloadFrom) GetName() *util.Nullable[string] {
	if m == nil {
		return nil
	}
	return m.Name
}

func (m *MailgunConfigPayloadFrom) SetName(name util.Nullable[string]) {
	m.Name = &name
}

func (m *MailgunConfigPayloadFrom) SetNameNull() {
	m.Name = &util.Nullable[string]{IsNull: true}
}

func (m MailgunConfigPayloadFrom) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MailgunConfigPayloadFrom to string"
	}
	return string(jsonData)
}

type Region string

const (
	REGION_US Region = "us"
	REGION_EU Region = "eu"
)
