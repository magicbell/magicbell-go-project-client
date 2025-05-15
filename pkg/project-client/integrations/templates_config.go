package integrations

import "encoding/json"

type TemplatesConfig struct {
	Config any     `json:"config,omitempty" required:"true"`
	Id     *string `json:"id,omitempty" required:"true"`
	Name   *string `json:"name,omitempty" required:"true"`
}

func (t *TemplatesConfig) GetConfig() any {
	if t == nil {
		return nil
	}
	return t.Config
}

func (t *TemplatesConfig) SetConfig(config any) {
	t.Config = &config
}

func (t *TemplatesConfig) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TemplatesConfig) SetId(id string) {
	t.Id = &id
}

func (t *TemplatesConfig) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *TemplatesConfig) SetName(name string) {
	t.Name = &name
}

func (t TemplatesConfig) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplatesConfig to string"
	}
	return string(jsonData)
}
