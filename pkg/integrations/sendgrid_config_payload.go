package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type SendgridConfigPayload struct {
	// The API key for Sendgrid
	ApiKey  *string                    `json:"api_key,omitempty" required:"true"`
	From    *SendgridConfigPayloadFrom `json:"from,omitempty"`
	ReplyTo *ReplyTo                   `json:"reply_to,omitempty"`
}

func (s *SendgridConfigPayload) GetApiKey() *string {
	if s == nil {
		return nil
	}
	return s.ApiKey
}

func (s *SendgridConfigPayload) SetApiKey(apiKey string) {
	s.ApiKey = &apiKey
}

func (s *SendgridConfigPayload) GetFrom() *SendgridConfigPayloadFrom {
	if s == nil {
		return nil
	}
	return s.From
}

func (s *SendgridConfigPayload) SetFrom(from SendgridConfigPayloadFrom) {
	s.From = &from
}

func (s *SendgridConfigPayload) GetReplyTo() *ReplyTo {
	if s == nil {
		return nil
	}
	return s.ReplyTo
}

func (s *SendgridConfigPayload) SetReplyTo(replyTo ReplyTo) {
	s.ReplyTo = &replyTo
}

func (s SendgridConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendgridConfigPayload to string"
	}
	return string(jsonData)
}

type SendgridConfigPayloadFrom struct {
	// The email address to send from
	Email *string `json:"email,omitempty" required:"true"`
	// The name to send from
	Name *util.Nullable[string] `json:"name,omitempty"`
}

func (s *SendgridConfigPayloadFrom) GetEmail() *string {
	if s == nil {
		return nil
	}
	return s.Email
}

func (s *SendgridConfigPayloadFrom) SetEmail(email string) {
	s.Email = &email
}

func (s *SendgridConfigPayloadFrom) GetName() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SendgridConfigPayloadFrom) SetName(name util.Nullable[string]) {
	s.Name = &name
}

func (s *SendgridConfigPayloadFrom) SetNameNull() {
	s.Name = &util.Nullable[string]{IsNull: true}
}

func (s SendgridConfigPayloadFrom) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendgridConfigPayloadFrom to string"
	}
	return string(jsonData)
}

type ReplyTo struct {
	// The email address to reply to
	Email *string `json:"email,omitempty" required:"true"`
	// The name to reply to
	Name *util.Nullable[string] `json:"name,omitempty"`
}

func (r *ReplyTo) GetEmail() *string {
	if r == nil {
		return nil
	}
	return r.Email
}

func (r *ReplyTo) SetEmail(email string) {
	r.Email = &email
}

func (r *ReplyTo) GetName() *util.Nullable[string] {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *ReplyTo) SetName(name util.Nullable[string]) {
	r.Name = &name
}

func (r *ReplyTo) SetNameNull() {
	r.Name = &util.Nullable[string]{IsNull: true}
}

func (r ReplyTo) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReplyTo to string"
	}
	return string(jsonData)
}
