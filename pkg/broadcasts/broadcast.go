package broadcasts

import (
	"encoding/json"
)

type Broadcast struct {
	ActionUrl *string `json:"action_url,omitempty" maxLength:"2048"`
	Category  *string `json:"category,omitempty" maxLength:"255" pattern:"^[A-Za-z0-9_\.\-/:]+$"`
	Content   *string `json:"content,omitempty" maxLength:"10485760"`
	// The timestamp when the broadcast was created.
	CreatedAt        *string `json:"created_at,omitempty"`
	CustomAttributes any     `json:"custom_attributes,omitempty"`
	// The unique id for this broadcast.
	Id         *string          `json:"id,omitempty"`
	Overrides  *Overrides       `json:"overrides,omitempty"`
	Recipients []any            `json:"recipients,omitempty" required:"true" minItems:"1" maxItems:"1000"`
	Status     *BroadcastStatus `json:"status,omitempty"`
	Title      *string          `json:"title,omitempty" required:"true" maxLength:"255" minLength:"1"`
	Topic      *string          `json:"topic,omitempty" maxLength:"255" pattern:"^[A-Za-z0-9_\.\-/:]+$"`
	touched    map[string]bool
}

func (b *Broadcast) GetActionUrl() *string {
	if b == nil {
		return nil
	}
	return b.ActionUrl
}

func (b *Broadcast) SetActionUrl(actionUrl string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["ActionUrl"] = true
	b.ActionUrl = &actionUrl
}

func (b *Broadcast) SetActionUrlNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["ActionUrl"] = true
	b.ActionUrl = nil
}

func (b *Broadcast) GetCategory() *string {
	if b == nil {
		return nil
	}
	return b.Category
}

func (b *Broadcast) SetCategory(category string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Category"] = true
	b.Category = &category
}

func (b *Broadcast) SetCategoryNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Category"] = true
	b.Category = nil
}

func (b *Broadcast) GetContent() *string {
	if b == nil {
		return nil
	}
	return b.Content
}

func (b *Broadcast) SetContent(content string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Content"] = true
	b.Content = &content
}

func (b *Broadcast) SetContentNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Content"] = true
	b.Content = nil
}

func (b *Broadcast) GetCreatedAt() *string {
	if b == nil {
		return nil
	}
	return b.CreatedAt
}

func (b *Broadcast) SetCreatedAt(createdAt string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["CreatedAt"] = true
	b.CreatedAt = &createdAt
}

func (b *Broadcast) SetCreatedAtNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["CreatedAt"] = true
	b.CreatedAt = nil
}

func (b *Broadcast) GetCustomAttributes() any {
	if b == nil {
		return nil
	}
	return b.CustomAttributes
}

func (b *Broadcast) SetCustomAttributes(customAttributes any) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["CustomAttributes"] = true
	b.CustomAttributes = customAttributes
}

func (b *Broadcast) SetCustomAttributesNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["CustomAttributes"] = true
	b.CustomAttributes = nil
}

func (b *Broadcast) GetId() *string {
	if b == nil {
		return nil
	}
	return b.Id
}

func (b *Broadcast) SetId(id string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Id"] = true
	b.Id = &id
}

func (b *Broadcast) SetIdNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Id"] = true
	b.Id = nil
}

func (b *Broadcast) GetOverrides() *Overrides {
	if b == nil {
		return nil
	}
	return b.Overrides
}

func (b *Broadcast) SetOverrides(overrides Overrides) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Overrides"] = true
	b.Overrides = &overrides
}

func (b *Broadcast) SetOverridesNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Overrides"] = true
	b.Overrides = nil
}

func (b *Broadcast) GetRecipients() []any {
	if b == nil {
		return nil
	}
	return b.Recipients
}

func (b *Broadcast) SetRecipients(recipients []any) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Recipients"] = true
	b.Recipients = recipients
}

func (b *Broadcast) SetRecipientsNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Recipients"] = true
	b.Recipients = nil
}

func (b *Broadcast) GetStatus() *BroadcastStatus {
	if b == nil {
		return nil
	}
	return b.Status
}

func (b *Broadcast) SetStatus(status BroadcastStatus) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Status"] = true
	b.Status = &status
}

func (b *Broadcast) SetStatusNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Status"] = true
	b.Status = nil
}

func (b *Broadcast) GetTitle() *string {
	if b == nil {
		return nil
	}
	return b.Title
}

func (b *Broadcast) SetTitle(title string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Title"] = true
	b.Title = &title
}

func (b *Broadcast) SetTitleNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Title"] = true
	b.Title = nil
}

func (b *Broadcast) GetTopic() *string {
	if b == nil {
		return nil
	}
	return b.Topic
}

func (b *Broadcast) SetTopic(topic string) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Topic"] = true
	b.Topic = &topic
}

func (b *Broadcast) SetTopicNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Topic"] = true
	b.Topic = nil
}
func (b Broadcast) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if b.touched["ActionUrl"] && b.ActionUrl == nil {
		data["action_url"] = nil
	} else if b.ActionUrl != nil {
		data["action_url"] = b.ActionUrl
	}

	if b.touched["Category"] && b.Category == nil {
		data["category"] = nil
	} else if b.Category != nil {
		data["category"] = b.Category
	}

	if b.touched["Content"] && b.Content == nil {
		data["content"] = nil
	} else if b.Content != nil {
		data["content"] = b.Content
	}

	if b.touched["CreatedAt"] && b.CreatedAt == nil {
		data["created_at"] = nil
	} else if b.CreatedAt != nil {
		data["created_at"] = b.CreatedAt
	}

	if b.touched["CustomAttributes"] && b.CustomAttributes == nil {
		data["custom_attributes"] = nil
	} else if b.CustomAttributes != nil {
		data["custom_attributes"] = b.CustomAttributes
	}

	if b.touched["Id"] && b.Id == nil {
		data["id"] = nil
	} else if b.Id != nil {
		data["id"] = b.Id
	}

	if b.touched["Overrides"] && b.Overrides == nil {
		data["overrides"] = nil
	} else if b.Overrides != nil {
		data["overrides"] = b.Overrides
	}

	if b.touched["Recipients"] && b.Recipients == nil {
		data["recipients"] = nil
	} else if b.Recipients != nil {
		data["recipients"] = b.Recipients
	}

	if b.touched["Status"] && b.Status == nil {
		data["status"] = nil
	} else if b.Status != nil {
		data["status"] = b.Status
	}

	if b.touched["Title"] && b.Title == nil {
		data["title"] = nil
	} else if b.Title != nil {
		data["title"] = b.Title
	}

	if b.touched["Topic"] && b.Topic == nil {
		data["topic"] = nil
	} else if b.Topic != nil {
		data["topic"] = b.Topic
	}

	return json.Marshal(data)
}

type Overrides struct {
	Channels  *OverridesChannels `json:"channels,omitempty"`
	Providers *Providers         `json:"providers,omitempty"`
	touched   map[string]bool
}

func (o *Overrides) GetChannels() *OverridesChannels {
	if o == nil {
		return nil
	}
	return o.Channels
}

func (o *Overrides) SetChannels(channels OverridesChannels) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Channels"] = true
	o.Channels = &channels
}

func (o *Overrides) SetChannelsNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Channels"] = true
	o.Channels = nil
}

func (o *Overrides) GetProviders() *Providers {
	if o == nil {
		return nil
	}
	return o.Providers
}

func (o *Overrides) SetProviders(providers Providers) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Providers"] = true
	o.Providers = &providers
}

func (o *Overrides) SetProvidersNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Providers"] = true
	o.Providers = nil
}
func (o Overrides) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if o.touched["Channels"] && o.Channels == nil {
		data["channels"] = nil
	} else if o.Channels != nil {
		data["channels"] = o.Channels
	}

	if o.touched["Providers"] && o.Providers == nil {
		data["providers"] = nil
	} else if o.Providers != nil {
		data["providers"] = o.Providers
	}

	return json.Marshal(data)
}

type OverridesChannels struct {
	Email      *Email      `json:"email,omitempty"`
	InApp      *InApp      `json:"in_app,omitempty"`
	MobilePush *MobilePush `json:"mobile_push,omitempty"`
	Slack      *Slack      `json:"slack,omitempty"`
	Sms        *Sms        `json:"sms,omitempty"`
	WebPush    *WebPush    `json:"web_push,omitempty"`
	touched    map[string]bool
}

func (o *OverridesChannels) GetEmail() *Email {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *OverridesChannels) SetEmail(email Email) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Email"] = true
	o.Email = &email
}

func (o *OverridesChannels) SetEmailNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Email"] = true
	o.Email = nil
}

func (o *OverridesChannels) GetInApp() *InApp {
	if o == nil {
		return nil
	}
	return o.InApp
}

func (o *OverridesChannels) SetInApp(inApp InApp) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["InApp"] = true
	o.InApp = &inApp
}

func (o *OverridesChannels) SetInAppNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["InApp"] = true
	o.InApp = nil
}

func (o *OverridesChannels) GetMobilePush() *MobilePush {
	if o == nil {
		return nil
	}
	return o.MobilePush
}

func (o *OverridesChannels) SetMobilePush(mobilePush MobilePush) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["MobilePush"] = true
	o.MobilePush = &mobilePush
}

func (o *OverridesChannels) SetMobilePushNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["MobilePush"] = true
	o.MobilePush = nil
}

func (o *OverridesChannels) GetSlack() *Slack {
	if o == nil {
		return nil
	}
	return o.Slack
}

func (o *OverridesChannels) SetSlack(slack Slack) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Slack"] = true
	o.Slack = &slack
}

func (o *OverridesChannels) SetSlackNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Slack"] = true
	o.Slack = nil
}

func (o *OverridesChannels) GetSms() *Sms {
	if o == nil {
		return nil
	}
	return o.Sms
}

func (o *OverridesChannels) SetSms(sms Sms) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Sms"] = true
	o.Sms = &sms
}

func (o *OverridesChannels) SetSmsNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Sms"] = true
	o.Sms = nil
}

func (o *OverridesChannels) GetWebPush() *WebPush {
	if o == nil {
		return nil
	}
	return o.WebPush
}

func (o *OverridesChannels) SetWebPush(webPush WebPush) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["WebPush"] = true
	o.WebPush = &webPush
}

func (o *OverridesChannels) SetWebPushNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["WebPush"] = true
	o.WebPush = nil
}
func (o OverridesChannels) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if o.touched["Email"] && o.Email == nil {
		data["email"] = nil
	} else if o.Email != nil {
		data["email"] = o.Email
	}

	if o.touched["InApp"] && o.InApp == nil {
		data["in_app"] = nil
	} else if o.InApp != nil {
		data["in_app"] = o.InApp
	}

	if o.touched["MobilePush"] && o.MobilePush == nil {
		data["mobile_push"] = nil
	} else if o.MobilePush != nil {
		data["mobile_push"] = o.MobilePush
	}

	if o.touched["Slack"] && o.Slack == nil {
		data["slack"] = nil
	} else if o.Slack != nil {
		data["slack"] = o.Slack
	}

	if o.touched["Sms"] && o.Sms == nil {
		data["sms"] = nil
	} else if o.Sms != nil {
		data["sms"] = o.Sms
	}

	if o.touched["WebPush"] && o.WebPush == nil {
		data["web_push"] = nil
	} else if o.WebPush != nil {
		data["web_push"] = o.WebPush
	}

	return json.Marshal(data)
}

type Email struct {
	ActionUrl *string `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string `json:"content,omitempty" maxLength:"1048576"`
	Title     *string `json:"title,omitempty" maxLength:"255" minLength:"1"`
	touched   map[string]bool
}

func (e *Email) GetActionUrl() *string {
	if e == nil {
		return nil
	}
	return e.ActionUrl
}

func (e *Email) SetActionUrl(actionUrl string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["ActionUrl"] = true
	e.ActionUrl = &actionUrl
}

func (e *Email) SetActionUrlNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["ActionUrl"] = true
	e.ActionUrl = nil
}

func (e *Email) GetContent() *string {
	if e == nil {
		return nil
	}
	return e.Content
}

func (e *Email) SetContent(content string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Content"] = true
	e.Content = &content
}

func (e *Email) SetContentNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Content"] = true
	e.Content = nil
}

func (e *Email) GetTitle() *string {
	if e == nil {
		return nil
	}
	return e.Title
}

func (e *Email) SetTitle(title string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Title"] = true
	e.Title = &title
}

func (e *Email) SetTitleNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Title"] = true
	e.Title = nil
}
func (e Email) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["ActionUrl"] && e.ActionUrl == nil {
		data["action_url"] = nil
	} else if e.ActionUrl != nil {
		data["action_url"] = e.ActionUrl
	}

	if e.touched["Content"] && e.Content == nil {
		data["content"] = nil
	} else if e.Content != nil {
		data["content"] = e.Content
	}

	if e.touched["Title"] && e.Title == nil {
		data["title"] = nil
	} else if e.Title != nil {
		data["title"] = e.Title
	}

	return json.Marshal(data)
}

type InApp struct {
	ActionUrl *string `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string `json:"content,omitempty" maxLength:"1048576"`
	Title     *string `json:"title,omitempty" maxLength:"255" minLength:"1"`
	touched   map[string]bool
}

func (i *InApp) GetActionUrl() *string {
	if i == nil {
		return nil
	}
	return i.ActionUrl
}

func (i *InApp) SetActionUrl(actionUrl string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["ActionUrl"] = true
	i.ActionUrl = &actionUrl
}

func (i *InApp) SetActionUrlNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["ActionUrl"] = true
	i.ActionUrl = nil
}

func (i *InApp) GetContent() *string {
	if i == nil {
		return nil
	}
	return i.Content
}

func (i *InApp) SetContent(content string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Content"] = true
	i.Content = &content
}

func (i *InApp) SetContentNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Content"] = true
	i.Content = nil
}

func (i *InApp) GetTitle() *string {
	if i == nil {
		return nil
	}
	return i.Title
}

func (i *InApp) SetTitle(title string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Title"] = true
	i.Title = &title
}

func (i *InApp) SetTitleNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Title"] = true
	i.Title = nil
}
func (i InApp) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["ActionUrl"] && i.ActionUrl == nil {
		data["action_url"] = nil
	} else if i.ActionUrl != nil {
		data["action_url"] = i.ActionUrl
	}

	if i.touched["Content"] && i.Content == nil {
		data["content"] = nil
	} else if i.Content != nil {
		data["content"] = i.Content
	}

	if i.touched["Title"] && i.Title == nil {
		data["title"] = nil
	} else if i.Title != nil {
		data["title"] = i.Title
	}

	return json.Marshal(data)
}

type MobilePush struct {
	ActionUrl *string `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string `json:"content,omitempty" maxLength:"1048576"`
	Title     *string `json:"title,omitempty" maxLength:"255" minLength:"1"`
	touched   map[string]bool
}

func (m *MobilePush) GetActionUrl() *string {
	if m == nil {
		return nil
	}
	return m.ActionUrl
}

func (m *MobilePush) SetActionUrl(actionUrl string) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["ActionUrl"] = true
	m.ActionUrl = &actionUrl
}

func (m *MobilePush) SetActionUrlNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["ActionUrl"] = true
	m.ActionUrl = nil
}

func (m *MobilePush) GetContent() *string {
	if m == nil {
		return nil
	}
	return m.Content
}

func (m *MobilePush) SetContent(content string) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Content"] = true
	m.Content = &content
}

func (m *MobilePush) SetContentNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Content"] = true
	m.Content = nil
}

func (m *MobilePush) GetTitle() *string {
	if m == nil {
		return nil
	}
	return m.Title
}

func (m *MobilePush) SetTitle(title string) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Title"] = true
	m.Title = &title
}

func (m *MobilePush) SetTitleNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Title"] = true
	m.Title = nil
}
func (m MobilePush) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if m.touched["ActionUrl"] && m.ActionUrl == nil {
		data["action_url"] = nil
	} else if m.ActionUrl != nil {
		data["action_url"] = m.ActionUrl
	}

	if m.touched["Content"] && m.Content == nil {
		data["content"] = nil
	} else if m.Content != nil {
		data["content"] = m.Content
	}

	if m.touched["Title"] && m.Title == nil {
		data["title"] = nil
	} else if m.Title != nil {
		data["title"] = m.Title
	}

	return json.Marshal(data)
}

type Slack struct {
	ActionUrl *string `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string `json:"content,omitempty" maxLength:"1048576"`
	Title     *string `json:"title,omitempty" maxLength:"255" minLength:"1"`
	touched   map[string]bool
}

func (s *Slack) GetActionUrl() *string {
	if s == nil {
		return nil
	}
	return s.ActionUrl
}

func (s *Slack) SetActionUrl(actionUrl string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ActionUrl"] = true
	s.ActionUrl = &actionUrl
}

func (s *Slack) SetActionUrlNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ActionUrl"] = true
	s.ActionUrl = nil
}

func (s *Slack) GetContent() *string {
	if s == nil {
		return nil
	}
	return s.Content
}

func (s *Slack) SetContent(content string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Content"] = true
	s.Content = &content
}

func (s *Slack) SetContentNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Content"] = true
	s.Content = nil
}

func (s *Slack) GetTitle() *string {
	if s == nil {
		return nil
	}
	return s.Title
}

func (s *Slack) SetTitle(title string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Title"] = true
	s.Title = &title
}

func (s *Slack) SetTitleNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Title"] = true
	s.Title = nil
}
func (s Slack) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["ActionUrl"] && s.ActionUrl == nil {
		data["action_url"] = nil
	} else if s.ActionUrl != nil {
		data["action_url"] = s.ActionUrl
	}

	if s.touched["Content"] && s.Content == nil {
		data["content"] = nil
	} else if s.Content != nil {
		data["content"] = s.Content
	}

	if s.touched["Title"] && s.Title == nil {
		data["title"] = nil
	} else if s.Title != nil {
		data["title"] = s.Title
	}

	return json.Marshal(data)
}

type Sms struct {
	ActionUrl *string `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string `json:"content,omitempty" maxLength:"1048576"`
	Title     *string `json:"title,omitempty" maxLength:"255" minLength:"1"`
	touched   map[string]bool
}

func (s *Sms) GetActionUrl() *string {
	if s == nil {
		return nil
	}
	return s.ActionUrl
}

func (s *Sms) SetActionUrl(actionUrl string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ActionUrl"] = true
	s.ActionUrl = &actionUrl
}

func (s *Sms) SetActionUrlNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ActionUrl"] = true
	s.ActionUrl = nil
}

func (s *Sms) GetContent() *string {
	if s == nil {
		return nil
	}
	return s.Content
}

func (s *Sms) SetContent(content string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Content"] = true
	s.Content = &content
}

func (s *Sms) SetContentNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Content"] = true
	s.Content = nil
}

func (s *Sms) GetTitle() *string {
	if s == nil {
		return nil
	}
	return s.Title
}

func (s *Sms) SetTitle(title string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Title"] = true
	s.Title = &title
}

func (s *Sms) SetTitleNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Title"] = true
	s.Title = nil
}
func (s Sms) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["ActionUrl"] && s.ActionUrl == nil {
		data["action_url"] = nil
	} else if s.ActionUrl != nil {
		data["action_url"] = s.ActionUrl
	}

	if s.touched["Content"] && s.Content == nil {
		data["content"] = nil
	} else if s.Content != nil {
		data["content"] = s.Content
	}

	if s.touched["Title"] && s.Title == nil {
		data["title"] = nil
	} else if s.Title != nil {
		data["title"] = s.Title
	}

	return json.Marshal(data)
}

type WebPush struct {
	ActionUrl *string `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string `json:"content,omitempty" maxLength:"1048576"`
	Title     *string `json:"title,omitempty" maxLength:"255" minLength:"1"`
	touched   map[string]bool
}

func (w *WebPush) GetActionUrl() *string {
	if w == nil {
		return nil
	}
	return w.ActionUrl
}

func (w *WebPush) SetActionUrl(actionUrl string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["ActionUrl"] = true
	w.ActionUrl = &actionUrl
}

func (w *WebPush) SetActionUrlNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["ActionUrl"] = true
	w.ActionUrl = nil
}

func (w *WebPush) GetContent() *string {
	if w == nil {
		return nil
	}
	return w.Content
}

func (w *WebPush) SetContent(content string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Content"] = true
	w.Content = &content
}

func (w *WebPush) SetContentNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Content"] = true
	w.Content = nil
}

func (w *WebPush) GetTitle() *string {
	if w == nil {
		return nil
	}
	return w.Title
}

func (w *WebPush) SetTitle(title string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Title"] = true
	w.Title = &title
}

func (w *WebPush) SetTitleNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Title"] = true
	w.Title = nil
}
func (w WebPush) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if w.touched["ActionUrl"] && w.ActionUrl == nil {
		data["action_url"] = nil
	} else if w.ActionUrl != nil {
		data["action_url"] = w.ActionUrl
	}

	if w.touched["Content"] && w.Content == nil {
		data["content"] = nil
	} else if w.Content != nil {
		data["content"] = w.Content
	}

	if w.touched["Title"] && w.Title == nil {
		data["title"] = nil
	} else if w.Title != nil {
		data["title"] = w.Title
	}

	return json.Marshal(data)
}

type Providers struct {
	AmazonSes any `json:"amazon_ses,omitempty"`
	Android   any `json:"android,omitempty"`
	Ios       any `json:"ios,omitempty"`
	Mailgun   any `json:"mailgun,omitempty"`
	Postmark  any `json:"postmark,omitempty"`
	Sendgrid  any `json:"sendgrid,omitempty"`
	Slack     any `json:"slack,omitempty"`
	touched   map[string]bool
}

func (p *Providers) GetAmazonSes() any {
	if p == nil {
		return nil
	}
	return p.AmazonSes
}

func (p *Providers) SetAmazonSes(amazonSes any) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["AmazonSes"] = true
	p.AmazonSes = amazonSes
}

func (p *Providers) SetAmazonSesNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["AmazonSes"] = true
	p.AmazonSes = nil
}

func (p *Providers) GetAndroid() any {
	if p == nil {
		return nil
	}
	return p.Android
}

func (p *Providers) SetAndroid(android any) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Android"] = true
	p.Android = android
}

func (p *Providers) SetAndroidNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Android"] = true
	p.Android = nil
}

func (p *Providers) GetIos() any {
	if p == nil {
		return nil
	}
	return p.Ios
}

func (p *Providers) SetIos(ios any) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Ios"] = true
	p.Ios = ios
}

func (p *Providers) SetIosNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Ios"] = true
	p.Ios = nil
}

func (p *Providers) GetMailgun() any {
	if p == nil {
		return nil
	}
	return p.Mailgun
}

func (p *Providers) SetMailgun(mailgun any) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Mailgun"] = true
	p.Mailgun = mailgun
}

func (p *Providers) SetMailgunNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Mailgun"] = true
	p.Mailgun = nil
}

func (p *Providers) GetPostmark() any {
	if p == nil {
		return nil
	}
	return p.Postmark
}

func (p *Providers) SetPostmark(postmark any) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Postmark"] = true
	p.Postmark = postmark
}

func (p *Providers) SetPostmarkNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Postmark"] = true
	p.Postmark = nil
}

func (p *Providers) GetSendgrid() any {
	if p == nil {
		return nil
	}
	return p.Sendgrid
}

func (p *Providers) SetSendgrid(sendgrid any) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Sendgrid"] = true
	p.Sendgrid = sendgrid
}

func (p *Providers) SetSendgridNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Sendgrid"] = true
	p.Sendgrid = nil
}

func (p *Providers) GetSlack() any {
	if p == nil {
		return nil
	}
	return p.Slack
}

func (p *Providers) SetSlack(slack any) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Slack"] = true
	p.Slack = slack
}

func (p *Providers) SetSlackNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Slack"] = true
	p.Slack = nil
}
func (p Providers) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["AmazonSes"] && p.AmazonSes == nil {
		data["amazon_ses"] = nil
	} else if p.AmazonSes != nil {
		data["amazon_ses"] = p.AmazonSes
	}

	if p.touched["Android"] && p.Android == nil {
		data["android"] = nil
	} else if p.Android != nil {
		data["android"] = p.Android
	}

	if p.touched["Ios"] && p.Ios == nil {
		data["ios"] = nil
	} else if p.Ios != nil {
		data["ios"] = p.Ios
	}

	if p.touched["Mailgun"] && p.Mailgun == nil {
		data["mailgun"] = nil
	} else if p.Mailgun != nil {
		data["mailgun"] = p.Mailgun
	}

	if p.touched["Postmark"] && p.Postmark == nil {
		data["postmark"] = nil
	} else if p.Postmark != nil {
		data["postmark"] = p.Postmark
	}

	if p.touched["Sendgrid"] && p.Sendgrid == nil {
		data["sendgrid"] = nil
	} else if p.Sendgrid != nil {
		data["sendgrid"] = p.Sendgrid
	}

	if p.touched["Slack"] && p.Slack == nil {
		data["slack"] = nil
	} else if p.Slack != nil {
		data["slack"] = p.Slack
	}

	return json.Marshal(data)
}

type BroadcastStatus struct {
	Errors  []Errors      `json:"errors,omitempty" required:"true"`
	Status  *StatusStatus `json:"status,omitempty" required:"true"`
	Summary *Summary      `json:"summary,omitempty" required:"true"`
	touched map[string]bool
}

func (b *BroadcastStatus) GetErrors() []Errors {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BroadcastStatus) SetErrors(errors []Errors) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Errors"] = true
	b.Errors = errors
}

func (b *BroadcastStatus) SetErrorsNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Errors"] = true
	b.Errors = nil
}

func (b *BroadcastStatus) GetStatus() *StatusStatus {
	if b == nil {
		return nil
	}
	return b.Status
}

func (b *BroadcastStatus) SetStatus(status StatusStatus) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Status"] = true
	b.Status = &status
}

func (b *BroadcastStatus) SetStatusNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Status"] = true
	b.Status = nil
}

func (b *BroadcastStatus) GetSummary() *Summary {
	if b == nil {
		return nil
	}
	return b.Summary
}

func (b *BroadcastStatus) SetSummary(summary Summary) {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Summary"] = true
	b.Summary = &summary
}

func (b *BroadcastStatus) SetSummaryNil() {
	if b.touched == nil {
		b.touched = map[string]bool{}
	}
	b.touched["Summary"] = true
	b.Summary = nil
}
func (b BroadcastStatus) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if b.touched["Errors"] && b.Errors == nil {
		data["errors"] = nil
	} else if b.Errors != nil {
		data["errors"] = b.Errors
	}

	if b.touched["Status"] && b.Status == nil {
		data["status"] = nil
	} else if b.Status != nil {
		data["status"] = b.Status
	}

	if b.touched["Summary"] && b.Summary == nil {
		data["summary"] = nil
	} else if b.Summary != nil {
		data["summary"] = b.Summary
	}

	return json.Marshal(data)
}

type Errors struct {
	Message *string `json:"message,omitempty"`
	touched map[string]bool
}

func (e *Errors) GetMessage() *string {
	if e == nil {
		return nil
	}
	return e.Message
}

func (e *Errors) SetMessage(message string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Message"] = true
	e.Message = &message
}

func (e *Errors) SetMessageNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Message"] = true
	e.Message = nil
}
func (e Errors) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["Message"] && e.Message == nil {
		data["message"] = nil
	} else if e.Message != nil {
		data["message"] = e.Message
	}

	return json.Marshal(data)
}

type StatusStatus string

const (
	STATUS_STATUS_ENQUEUED   StatusStatus = "enqueued"
	STATUS_STATUS_PROCESSING StatusStatus = "processing"
	STATUS_STATUS_PROCESSED  StatusStatus = "processed"
)

type Summary struct {
	// The number of failures while processing the broadcast.
	Failures *int64 `json:"failures,omitempty" required:"true"`
	// The number of recipients that the broadcast was sent to.
	Total   *int64 `json:"total,omitempty" required:"true"`
	touched map[string]bool
}

func (s *Summary) GetFailures() *int64 {
	if s == nil {
		return nil
	}
	return s.Failures
}

func (s *Summary) SetFailures(failures int64) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Failures"] = true
	s.Failures = &failures
}

func (s *Summary) SetFailuresNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Failures"] = true
	s.Failures = nil
}

func (s *Summary) GetTotal() *int64 {
	if s == nil {
		return nil
	}
	return s.Total
}

func (s *Summary) SetTotal(total int64) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Total"] = true
	s.Total = &total
}

func (s *Summary) SetTotalNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Total"] = true
	s.Total = nil
}
func (s Summary) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Failures"] && s.Failures == nil {
		data["failures"] = nil
	} else if s.Failures != nil {
		data["failures"] = s.Failures
	}

	if s.touched["Total"] && s.Total == nil {
		data["total"] = nil
	} else if s.Total != nil {
		data["total"] = s.Total
	}

	return json.Marshal(data)
}
