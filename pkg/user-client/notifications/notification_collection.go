package notifications

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/shared"
)

type NotificationCollection struct {
	Data  []Notification `json:"data,omitempty"`
	Links *shared.Links  `json:"links,omitempty"`
}

func (n *NotificationCollection) GetData() []Notification {
	if n == nil {
		return nil
	}
	return n.Data
}

func (n *NotificationCollection) SetData(data []Notification) {
	n.Data = data
}

func (n *NotificationCollection) GetLinks() *shared.Links {
	if n == nil {
		return nil
	}
	return n.Links
}

func (n *NotificationCollection) SetLinks(links shared.Links) {
	n.Links = &links
}

func (n NotificationCollection) String() string {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return "error converting struct: NotificationCollection to string"
	}
	return string(jsonData)
}
