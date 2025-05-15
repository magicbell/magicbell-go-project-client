package integrations

import "encoding/json"

type SlackConfig struct {
	Config *SlackConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string             `json:"id,omitempty" required:"true"`
	Name   *string             `json:"name,omitempty" required:"true"`
}

func (s *SlackConfig) GetConfig() *SlackConfigPayload {
	if s == nil {
		return nil
	}
	return s.Config
}

func (s *SlackConfig) SetConfig(config SlackConfigPayload) {
	s.Config = &config
}

func (s *SlackConfig) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SlackConfig) SetId(id string) {
	s.Id = &id
}

func (s *SlackConfig) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SlackConfig) SetName(name string) {
	s.Name = &name
}

func (s SlackConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackConfig to string"
	}
	return string(jsonData)
}
