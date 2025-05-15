package integrations

import "encoding/json"

type SesConfig struct {
	Config *SesConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string           `json:"id,omitempty" required:"true"`
	Name   *string           `json:"name,omitempty" required:"true"`
}

func (s *SesConfig) GetConfig() *SesConfigPayload {
	if s == nil {
		return nil
	}
	return s.Config
}

func (s *SesConfig) SetConfig(config SesConfigPayload) {
	s.Config = &config
}

func (s *SesConfig) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SesConfig) SetId(id string) {
	s.Id = &id
}

func (s *SesConfig) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SesConfig) SetName(name string) {
	s.Name = &name
}

func (s SesConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SesConfig to string"
	}
	return string(jsonData)
}
