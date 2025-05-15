package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/shared"
)

type ApnsTokenCollection struct {
	Data  []ApnsToken   `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (a *ApnsTokenCollection) GetData() []ApnsToken {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ApnsTokenCollection) SetData(data []ApnsToken) {
	a.Data = data
}

func (a *ApnsTokenCollection) GetLinks() *shared.Links {
	if a == nil {
		return nil
	}
	return a.Links
}

func (a *ApnsTokenCollection) SetLinks(links shared.Links) {
	a.Links = &links
}

func (a ApnsTokenCollection) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApnsTokenCollection to string"
	}
	return string(jsonData)
}
