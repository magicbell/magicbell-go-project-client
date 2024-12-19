package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/internal/marshal"
	"github.com/magicbell/magicbell-go-project-client/internal/unmarshal"
)

type SesConfig struct {
	// HTTP endpoint to send requests to (testing only)
	Endpoint *Endpoint      `json:"endpoint,omitempty"`
	From     *SesConfigFrom `json:"from,omitempty"`
	// AWS Access Key ID
	KeyId *string `json:"key_id,omitempty" required:"true" minLength:"1"`
	// AWS Region
	Region *string `json:"region,omitempty" required:"true" minLength:"1"`
	// AWS Secret Key
	SecretKey *string `json:"secret_key,omitempty" required:"true" minLength:"1"`
	touched   map[string]bool
}

func (s *SesConfig) GetEndpoint() *Endpoint {
	if s == nil {
		return nil
	}
	return s.Endpoint
}

func (s *SesConfig) SetEndpoint(endpoint Endpoint) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Endpoint"] = true
	s.Endpoint = &endpoint
}

func (s *SesConfig) SetEndpointNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Endpoint"] = true
	s.Endpoint = nil
}

func (s *SesConfig) GetFrom() *SesConfigFrom {
	if s == nil {
		return nil
	}
	return s.From
}

func (s *SesConfig) SetFrom(from SesConfigFrom) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["From"] = true
	s.From = &from
}

func (s *SesConfig) SetFromNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["From"] = true
	s.From = nil
}

func (s *SesConfig) GetKeyId() *string {
	if s == nil {
		return nil
	}
	return s.KeyId
}

func (s *SesConfig) SetKeyId(keyId string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["KeyId"] = true
	s.KeyId = &keyId
}

func (s *SesConfig) SetKeyIdNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["KeyId"] = true
	s.KeyId = nil
}

func (s *SesConfig) GetRegion() *string {
	if s == nil {
		return nil
	}
	return s.Region
}

func (s *SesConfig) SetRegion(region string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Region"] = true
	s.Region = &region
}

func (s *SesConfig) SetRegionNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Region"] = true
	s.Region = nil
}

func (s *SesConfig) GetSecretKey() *string {
	if s == nil {
		return nil
	}
	return s.SecretKey
}

func (s *SesConfig) SetSecretKey(secretKey string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["SecretKey"] = true
	s.SecretKey = &secretKey
}

func (s *SesConfig) SetSecretKeyNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["SecretKey"] = true
	s.SecretKey = nil
}
func (s SesConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Endpoint"] && s.Endpoint == nil {
		data["endpoint"] = nil
	} else if s.Endpoint != nil {
		data["endpoint"] = s.Endpoint
	}

	if s.touched["From"] && s.From == nil {
		data["from"] = nil
	} else if s.From != nil {
		data["from"] = s.From
	}

	if s.touched["KeyId"] && s.KeyId == nil {
		data["key_id"] = nil
	} else if s.KeyId != nil {
		data["key_id"] = s.KeyId
	}

	if s.touched["Region"] && s.Region == nil {
		data["region"] = nil
	} else if s.Region != nil {
		data["region"] = s.Region
	}

	if s.touched["SecretKey"] && s.SecretKey == nil {
		data["secret_key"] = nil
	} else if s.SecretKey != nil {
		data["secret_key"] = s.SecretKey
	}

	return json.Marshal(data)
}

// HTTP endpoint to send requests to (testing only)
type Endpoint struct {
	AnyValue  any
	AnyValue1 any
}

func (e *Endpoint) SetAnyValue(anyValue any) {
	e.AnyValue = anyValue
}

func (e *Endpoint) GetAnyValue() any {
	if e == nil {
		return nil
	}
	return e.AnyValue
}

func (e *Endpoint) SetAnyValue1(anyValue1 any) {
	e.AnyValue1 = anyValue1
}

func (e *Endpoint) GetAnyValue1() any {
	if e == nil {
		return nil
	}
	return e.AnyValue1
}

func (e Endpoint) MarshalJSON() ([]byte, error) {
	return marshal.FromComplexObject(e)
}

func (e *Endpoint) UnmarshalJSON(data []byte) error {
	return unmarshal.ToComplexObject[Endpoint](data, e)
}

type SesConfigFrom struct {
	// The email address to send from
	Email *string `json:"email,omitempty" required:"true"`
	// The name to send from
	Name    *string `json:"name,omitempty"`
	touched map[string]bool
}

func (s *SesConfigFrom) GetEmail() *string {
	if s == nil {
		return nil
	}
	return s.Email
}

func (s *SesConfigFrom) SetEmail(email string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Email"] = true
	s.Email = &email
}

func (s *SesConfigFrom) SetEmailNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Email"] = true
	s.Email = nil
}

func (s *SesConfigFrom) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SesConfigFrom) SetName(name string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = &name
}

func (s *SesConfigFrom) SetNameNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = nil
}
func (s SesConfigFrom) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Email"] && s.Email == nil {
		data["email"] = nil
	} else if s.Email != nil {
		data["email"] = s.Email
	}

	if s.touched["Name"] && s.Name == nil {
		data["name"] = nil
	} else if s.Name != nil {
		data["name"] = s.Name
	}

	return json.Marshal(data)
}
