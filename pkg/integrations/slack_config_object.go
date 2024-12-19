package integrations

import (
	"encoding/json"
)

type SlackConfigObject struct {
	Config  *SlackConfig `json:"config,omitempty" required:"true"`
	Id      *string      `json:"id,omitempty" required:"true"`
	Name    *string      `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (s *SlackConfigObject) GetConfig() *SlackConfig {
	if s == nil {
		return nil
	}
	return s.Config
}

func (s *SlackConfigObject) SetConfig(config SlackConfig) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Config"] = true
	s.Config = &config
}

func (s *SlackConfigObject) SetConfigNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Config"] = true
	s.Config = nil
}

func (s *SlackConfigObject) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SlackConfigObject) SetId(id string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Id"] = true
	s.Id = &id
}

func (s *SlackConfigObject) SetIdNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Id"] = true
	s.Id = nil
}

func (s *SlackConfigObject) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SlackConfigObject) SetName(name string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = &name
}

func (s *SlackConfigObject) SetNameNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = nil
}
func (s SlackConfigObject) MarshalJSON() ([]byte, error) {
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
