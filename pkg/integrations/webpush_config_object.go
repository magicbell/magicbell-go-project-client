package integrations

import (
	"encoding/json"
)

type WebpushConfigObject struct {
	Config  *WebpushConfig `json:"config,omitempty" required:"true"`
	Id      *string        `json:"id,omitempty" required:"true"`
	Name    *string        `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (w *WebpushConfigObject) GetConfig() *WebpushConfig {
	if w == nil {
		return nil
	}
	return w.Config
}

func (w *WebpushConfigObject) SetConfig(config WebpushConfig) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Config"] = true
	w.Config = &config
}

func (w *WebpushConfigObject) SetConfigNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Config"] = true
	w.Config = nil
}

func (w *WebpushConfigObject) GetId() *string {
	if w == nil {
		return nil
	}
	return w.Id
}

func (w *WebpushConfigObject) SetId(id string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Id"] = true
	w.Id = &id
}

func (w *WebpushConfigObject) SetIdNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Id"] = true
	w.Id = nil
}

func (w *WebpushConfigObject) GetName() *string {
	if w == nil {
		return nil
	}
	return w.Name
}

func (w *WebpushConfigObject) SetName(name string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Name"] = true
	w.Name = &name
}

func (w *WebpushConfigObject) SetNameNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Name"] = true
	w.Name = nil
}
func (w WebpushConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if w.touched["Config"] && w.Config == nil {
		data["config"] = nil
	} else if w.Config != nil {
		data["config"] = w.Config
	}

	if w.touched["Id"] && w.Id == nil {
		data["id"] = nil
	} else if w.Id != nil {
		data["id"] = w.Id
	}

	if w.touched["Name"] && w.Name == nil {
		data["name"] = nil
	} else if w.Name != nil {
		data["name"] = w.Name
	}

	return json.Marshal(data)
}
