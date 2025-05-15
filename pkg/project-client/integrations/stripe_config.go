package integrations

import "encoding/json"

type StripeConfig struct {
	Config *StripeConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string              `json:"id,omitempty" required:"true"`
	Name   *string              `json:"name,omitempty" required:"true"`
}

func (s *StripeConfig) GetConfig() *StripeConfigPayload {
	if s == nil {
		return nil
	}
	return s.Config
}

func (s *StripeConfig) SetConfig(config StripeConfigPayload) {
	s.Config = &config
}

func (s *StripeConfig) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *StripeConfig) SetId(id string) {
	s.Id = &id
}

func (s *StripeConfig) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *StripeConfig) SetName(name string) {
	s.Name = &name
}

func (s StripeConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StripeConfig to string"
	}
	return string(jsonData)
}
