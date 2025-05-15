package integrations

import "encoding/json"

type WebpushConfig struct {
	Config *WebpushConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string               `json:"id,omitempty" required:"true"`
	Name   *string               `json:"name,omitempty" required:"true"`
}

func (w *WebpushConfig) GetConfig() *WebpushConfigPayload {
	if w == nil {
		return nil
	}
	return w.Config
}

func (w *WebpushConfig) SetConfig(config WebpushConfigPayload) {
	w.Config = &config
}

func (w *WebpushConfig) GetId() *string {
	if w == nil {
		return nil
	}
	return w.Id
}

func (w *WebpushConfig) SetId(id string) {
	w.Id = &id
}

func (w *WebpushConfig) GetName() *string {
	if w == nil {
		return nil
	}
	return w.Name
}

func (w *WebpushConfig) SetName(name string) {
	w.Name = &name
}

func (w WebpushConfig) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebpushConfig to string"
	}
	return string(jsonData)
}
