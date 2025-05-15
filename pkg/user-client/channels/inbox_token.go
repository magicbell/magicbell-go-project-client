package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/util"
)

type InboxToken struct {
	ConnectionId *util.Nullable[string] `json:"connection_id,omitempty"`
	Token        *string                `json:"token,omitempty" required:"true" minLength:"64"`
}

func (i *InboxToken) GetConnectionId() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.ConnectionId
}

func (i *InboxToken) SetConnectionId(connectionId util.Nullable[string]) {
	i.ConnectionId = &connectionId
}

func (i *InboxToken) SetConnectionIdNull() {
	i.ConnectionId = &util.Nullable[string]{IsNull: true}
}

func (i *InboxToken) GetToken() *string {
	if i == nil {
		return nil
	}
	return i.Token
}

func (i *InboxToken) SetToken(token string) {
	i.Token = &token
}

func (i InboxToken) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InboxToken to string"
	}
	return string(jsonData)
}
