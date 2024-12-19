package integrations

import (
	"encoding/json"
)

type StripeConfigObject struct {
	Config  *StripeConfig `json:"config,omitempty" required:"true"`
	Id      *string       `json:"id,omitempty" required:"true"`
	Name    *string       `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (s *StripeConfigObject) GetConfig() *StripeConfig {
	if s == nil {
		return nil
	}
	return s.Config
}

func (s *StripeConfigObject) SetConfig(config StripeConfig) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Config"] = true
	s.Config = &config
}

func (s *StripeConfigObject) SetConfigNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Config"] = true
	s.Config = nil
}

func (s *StripeConfigObject) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *StripeConfigObject) SetId(id string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Id"] = true
	s.Id = &id
}

func (s *StripeConfigObject) SetIdNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Id"] = true
	s.Id = nil
}

func (s *StripeConfigObject) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *StripeConfigObject) SetName(name string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = &name
}

func (s *StripeConfigObject) SetNameNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = nil
}
func (s StripeConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Config"] && s.Config == nil {
		data["config"] = nil
	} else if s.Config != nil {
		data["config"] = s.Config
	}

	if s.touched["Id"] && s.Id == nil {
		data["id"] = nil
	} else if s.Id != nil {
		data["id"] = s.Id
	}

	if s.touched["Name"] && s.Name == nil {
		data["name"] = nil
	} else if s.Name != nil {
		data["name"] = s.Name
	}

	return json.Marshal(data)
}
