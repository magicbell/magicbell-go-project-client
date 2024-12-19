package integrations

import (
	"encoding/json"
)

type SendgridConfig struct {
	// The API key for Sendgrid
	ApiKey  *string             `json:"api_key,omitempty" required:"true"`
	From    *SendgridConfigFrom `json:"from,omitempty"`
	ReplyTo *ReplyTo            `json:"reply_to,omitempty"`
	touched map[string]bool
}

func (s *SendgridConfig) GetApiKey() *string {
	if s == nil {
		return nil
	}
	return s.ApiKey
}

func (s *SendgridConfig) SetApiKey(apiKey string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ApiKey"] = true
	s.ApiKey = &apiKey
}

func (s *SendgridConfig) SetApiKeyNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ApiKey"] = true
	s.ApiKey = nil
}

func (s *SendgridConfig) GetFrom() *SendgridConfigFrom {
	if s == nil {
		return nil
	}
	return s.From
}

func (s *SendgridConfig) SetFrom(from SendgridConfigFrom) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["From"] = true
	s.From = &from
}

func (s *SendgridConfig) SetFromNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["From"] = true
	s.From = nil
}

func (s *SendgridConfig) GetReplyTo() *ReplyTo {
	if s == nil {
		return nil
	}
	return s.ReplyTo
}

func (s *SendgridConfig) SetReplyTo(replyTo ReplyTo) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ReplyTo"] = true
	s.ReplyTo = &replyTo
}

func (s *SendgridConfig) SetReplyToNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ReplyTo"] = true
	s.ReplyTo = nil
}
func (s SendgridConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["ApiKey"] && s.ApiKey == nil {
		data["api_key"] = nil
	} else if s.ApiKey != nil {
		data["api_key"] = s.ApiKey
	}

	if s.touched["From"] && s.From == nil {
		data["from"] = nil
	} else if s.From != nil {
		data["from"] = s.From
	}

	if s.touched["ReplyTo"] && s.ReplyTo == nil {
		data["reply_to"] = nil
	} else if s.ReplyTo != nil {
		data["reply_to"] = s.ReplyTo
	}

	return json.Marshal(data)
}

type SendgridConfigFrom struct {
	// The email address to send from
	Email *string `json:"email,omitempty" required:"true"`
	// The name to send from
	Name    *string `json:"name,omitempty"`
	touched map[string]bool
}

func (s *SendgridConfigFrom) GetEmail() *string {
	if s == nil {
		return nil
	}
	return s.Email
}

func (s *SendgridConfigFrom) SetEmail(email string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Email"] = true
	s.Email = &email
}

func (s *SendgridConfigFrom) SetEmailNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Email"] = true
	s.Email = nil
}

func (s *SendgridConfigFrom) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SendgridConfigFrom) SetName(name string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = &name
}

func (s *SendgridConfigFrom) SetNameNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Name"] = true
	s.Name = nil
}
func (s SendgridConfigFrom) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Email"] && s.Email == nil {
		data["email"] = nil
	} else if s.Email != nil {
		data["email"] = s.Email
	}

	if s.touched["Name"] && s.Name == nil {
		data["name"] = nil
	} else if s.Name != nil {
		data["name"] = s.Name
	}

	return json.Marshal(data)
}

type ReplyTo struct {
	// The email address to reply to
	Email *string `json:"email,omitempty" required:"true"`
	// The name to reply to
	Name    *string `json:"name,omitempty"`
	touched map[string]bool
}

func (r *ReplyTo) GetEmail() *string {
	if r == nil {
		return nil
	}
	return r.Email
}

func (r *ReplyTo) SetEmail(email string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Email"] = true
	r.Email = &email
}

func (r *ReplyTo) SetEmailNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Email"] = true
	r.Email = nil
}

func (r *ReplyTo) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *ReplyTo) SetName(name string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Name"] = true
	r.Name = &name
}

func (r *ReplyTo) SetNameNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Name"] = true
	r.Name = nil
}
func (r ReplyTo) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["Email"] && r.Email == nil {
		data["email"] = nil
	} else if r.Email != nil {
		data["email"] = r.Email
	}

	if r.touched["Name"] && r.Name == nil {
		data["name"] = nil
	} else if r.Name != nil {
		data["name"] = r.Name
	}

	return json.Marshal(data)
}
