package integrations

import (
	"encoding/json"
)

type SesConfigObject struct {
	Config  *SesConfig `json:"config,omitempty" required:"true"`
	Id      *string    `json:"id,omitempty" required:"true"`
	Name    *string    `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (s *SesConfigObject) GetConfig() *SesConfig {
	if s == nil {
		return nil
	}
	return s.Config
}

func (s *SesConfigObject) SetConfig(config SesConfig) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Config"] = true
	s.Config = &config
}

func (s *SesConfigObject) SetConfigNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Config"] = true
	s.Config = nil
}

func (s *SesConfigObject) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SesConfigObject) SetId(id string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Id"] = true
	s.Id = &id
}

func (s *SesConfigObject) SetIdNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Id"] = true
	s.Id = nil
}

func (s *SesConfigObject) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SesConfigObject) SetName(name string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = &name
}

func (s *SesConfigObject) SetNameNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = nil
}
func (s SesConfigObject) MarshalJSON() ([]byte, error) {
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
