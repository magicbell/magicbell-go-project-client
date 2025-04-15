package integrations

import "encoding/json"

type SendgridConfig struct {
	Config *SendgridConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string                `json:"id,omitempty" required:"true"`
	Name   *string                `json:"name,omitempty" required:"true"`
}

func (s *SendgridConfig) GetConfig() *SendgridConfigPayload {
	if s == nil {
		return nil
	}
	return s.Config
}

func (s *SendgridConfig) SetConfig(config SendgridConfigPayload) {
	s.Config = &config
}

func (s *SendgridConfig) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SendgridConfig) SetId(id string) {
	s.Id = &id
}

func (s *SendgridConfig) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SendgridConfig) SetName(name string) {
	s.Name = &name
}

func (s SendgridConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendgridConfig to string"
	}
	return string(jsonData)
}
