package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type WebpushConfigCollection struct {
	Data  []WebpushConfig `json:"data,omitempty"`
	Links *shared.Links   `json:"links,omitempty"`
}

func (w *WebpushConfigCollection) GetData() []WebpushConfig {
	if w == nil {
		return nil
	}
	return w.Data
}

func (w *WebpushConfigCollection) SetData(data []WebpushConfig) {
	w.Data = data
}

func (w *WebpushConfigCollection) GetLinks() *shared.Links {
	if w == nil {
		return nil
	}
	return w.Links
}

func (w *WebpushConfigCollection) SetLinks(links shared.Links) {
	w.Links = &links
}

func (w WebpushConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebpushConfigCollection to string"
	}
	return string(jsonData)
}
