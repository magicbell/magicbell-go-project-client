package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type InboxTokenResponseCollection struct {
	Data  []InboxTokenResponse `json:"data,omitempty"`
	Links *shared.Links        `json:"links,omitempty"`
}

func (i *InboxTokenResponseCollection) GetData() []InboxTokenResponse {
	if i == nil {
		return nil
	}
	return i.Data
}

func (i *InboxTokenResponseCollection) SetData(data []InboxTokenResponse) {
	i.Data = data
}

func (i *InboxTokenResponseCollection) GetLinks() *shared.Links {
	if i == nil {
		return nil
	}
	return i.Links
}

func (i *InboxTokenResponseCollection) SetLinks(links shared.Links) {
	i.Links = &links
}

func (i InboxTokenResponseCollection) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InboxTokenResponseCollection to string"
	}
	return string(jsonData)
}
