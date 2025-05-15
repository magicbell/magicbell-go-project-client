package integrations

import "encoding/json"

type TwilioConfig struct {
	Config *TwilioConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string              `json:"id,omitempty" required:"true"`
	Name   *string              `json:"name,omitempty" required:"true"`
}

func (t *TwilioConfig) GetConfig() *TwilioConfigPayload {
	if t == nil {
		return nil
	}
	return t.Config
}

func (t *TwilioConfig) SetConfig(config TwilioConfigPayload) {
	t.Config = &config
}

func (t *TwilioConfig) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TwilioConfig) SetId(id string) {
	t.Id = &id
}

func (t *TwilioConfig) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *TwilioConfig) SetName(name string) {
	t.Name = &name
}

func (t TwilioConfig) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioConfig to string"
	}
	return string(jsonData)
}
