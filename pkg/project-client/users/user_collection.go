package users

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type UserCollection struct {
	Data  []User        `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (u *UserCollection) GetData() []User {
	if u == nil {
		return nil
	}
	return u.Data
}

func (u *UserCollection) SetData(data []User) {
	u.Data = data
}

func (u *UserCollection) GetLinks() *shared.Links {
	if u == nil {
		return nil
	}
	return u.Links
}

func (u *UserCollection) SetLinks(links shared.Links) {
	u.Links = &links
}

func (u UserCollection) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserCollection to string"
	}
	return string(jsonData)
}
