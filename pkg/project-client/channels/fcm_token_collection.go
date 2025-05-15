package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type FcmTokenCollection struct {
	Data  []FcmToken    `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (f *FcmTokenCollection) GetData() []FcmToken {
	if f == nil {
		return nil
	}
	return f.Data
}

func (f *FcmTokenCollection) SetData(data []FcmToken) {
	f.Data = data
}

func (f *FcmTokenCollection) GetLinks() *shared.Links {
	if f == nil {
		return nil
	}
	return f.Links
}

func (f *FcmTokenCollection) SetLinks(links shared.Links) {
	f.Links = &links
}

func (f FcmTokenCollection) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FcmTokenCollection to string"
	}
	return string(jsonData)
}
