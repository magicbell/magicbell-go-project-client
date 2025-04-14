package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type SesConfigPayload struct {
	From *SesConfigPayloadFrom `json:"from,omitempty"`
	// AWS Access Key ID
	KeyId *string `json:"key_id,omitempty" required:"true" minLength:"1"`
	// AWS Region
	Region *string `json:"region,omitempty" required:"true" minLength:"1"`
	// AWS Secret Key
	SecretKey *string `json:"secret_key,omitempty" required:"true" minLength:"1"`
}

func (s *SesConfigPayload) GetFrom() *SesConfigPayloadFrom {
	if s == nil {
		return nil
	}
	return s.From
}

func (s *SesConfigPayload) SetFrom(from SesConfigPayloadFrom) {
	s.From = &from
}

func (s *SesConfigPayload) GetKeyId() *string {
	if s == nil {
		return nil
	}
	return s.KeyId
}

func (s *SesConfigPayload) SetKeyId(keyId string) {
	s.KeyId = &keyId
}

func (s *SesConfigPayload) GetRegion() *string {
	if s == nil {
		return nil
	}
	return s.Region
}

func (s *SesConfigPayload) SetRegion(region string) {
	s.Region = &region
}

func (s *SesConfigPayload) GetSecretKey() *string {
	if s == nil {
		return nil
	}
	return s.SecretKey
}

func (s *SesConfigPayload) SetSecretKey(secretKey string) {
	s.SecretKey = &secretKey
}

func (s SesConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SesConfigPayload to string"
	}
	return string(jsonData)
}

type SesConfigPayloadFrom struct {
	// The email address to send from
	Email *string `json:"email,omitempty" required:"true"`
	// The name to send from
	Name *util.Nullable[string] `json:"name,omitempty"`
}

func (s *SesConfigPayloadFrom) GetEmail() *string {
	if s == nil {
		return nil
	}
	return s.Email
}

func (s *SesConfigPayloadFrom) SetEmail(email string) {
	s.Email = &email
}

func (s *SesConfigPayloadFrom) GetName() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SesConfigPayloadFrom) SetName(name util.Nullable[string]) {
	s.Name = &name
}

func (s *SesConfigPayloadFrom) SetNameNull() {
	s.Name = &util.Nullable[string]{IsNull: true}
}

func (s SesConfigPayloadFrom) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SesConfigPayloadFrom to string"
	}
	return string(jsonData)
}
