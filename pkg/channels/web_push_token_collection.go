package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type WebPushTokenCollection struct {
	Data  []WebPushToken `json:"data,omitempty"`
	Links *shared.Links  `json:"links,omitempty"`
}

func (w *WebPushTokenCollection) GetData() []WebPushToken {
	if w == nil {
		return nil
	}
	return w.Data
}

func (w *WebPushTokenCollection) SetData(data []WebPushToken) {
	w.Data = data
}

func (w *WebPushTokenCollection) GetLinks() *shared.Links {
	if w == nil {
		return nil
	}
	return w.Links
}

func (w *WebPushTokenCollection) SetLinks(links shared.Links) {
	w.Links = &links
}

func (w WebPushTokenCollection) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebPushTokenCollection to string"
	}
	return string(jsonData)
}
