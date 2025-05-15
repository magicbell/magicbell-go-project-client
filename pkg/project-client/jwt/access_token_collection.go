package jwt

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type AccessTokenCollection struct {
	Data  []AccessToken `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (a *AccessTokenCollection) GetData() []AccessToken {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *AccessTokenCollection) SetData(data []AccessToken) {
	a.Data = data
}

func (a *AccessTokenCollection) GetLinks() *shared.Links {
	if a == nil {
		return nil
	}
	return a.Links
}

func (a *AccessTokenCollection) SetLinks(links shared.Links) {
	a.Links = &links
}

func (a AccessTokenCollection) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AccessTokenCollection to string"
	}
	return string(jsonData)
}
