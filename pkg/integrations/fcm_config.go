package integrations

import "encoding/json"

type FcmConfig struct {
	Config *FcmConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string           `json:"id,omitempty" required:"true"`
	Name   *string           `json:"name,omitempty" required:"true"`
}

func (f *FcmConfig) GetConfig() *FcmConfigPayload {
	if f == nil {
		return nil
	}
	return f.Config
}

func (f *FcmConfig) SetConfig(config FcmConfigPayload) {
	f.Config = &config
}

func (f *FcmConfig) GetId() *string {
	if f == nil {
		return nil
	}
	return f.Id
}

func (f *FcmConfig) SetId(id string) {
	f.Id = &id
}

func (f *FcmConfig) GetName() *string {
	if f == nil {
		return nil
	}
	return f.Name
}

func (f *FcmConfig) SetName(name string) {
	f.Name = &name
}

func (f FcmConfig) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FcmConfig to string"
	}
	return string(jsonData)
}
