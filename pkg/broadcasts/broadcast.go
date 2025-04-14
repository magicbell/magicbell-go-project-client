package broadcasts

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type Broadcast struct {
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	Category  *util.Nullable[string] `json:"category,omitempty" maxLength:"255" pattern:"^[A-Za-z0-9_\.\-/:]+$"`
	Content   *util.Nullable[string] `json:"content,omitempty" maxLength:"10485760"`
	// The timestamp when the broadcast was created.
	CreatedAt        *string             `json:"created_at,omitempty"`
	CustomAttributes *util.Nullable[any] `json:"custom_attributes,omitempty"`
	// The unique id for this broadcast.
	Id         *string                         `json:"id,omitempty"`
	Overrides  *util.Nullable[Overrides]       `json:"overrides,omitempty"`
	Recipients *util.Nullable[[]any]           `json:"recipients,omitempty" required:"true" minItems:"1" maxItems:"1000"`
	Status     *util.Nullable[BroadcastStatus] `json:"status,omitempty"`
	Title      *string                         `json:"title,omitempty" required:"true" maxLength:"255" minLength:"1"`
	Topic      *util.Nullable[string]          `json:"topic,omitempty" maxLength:"255" pattern:"^[A-Za-z0-9_\.\-/:]+$"`
}

func (b *Broadcast) GetActionUrl() *util.Nullable[string] {
	if b == nil {
		return nil
	}
	return b.ActionUrl
}

func (b *Broadcast) SetActionUrl(actionUrl util.Nullable[string]) {
	b.ActionUrl = &actionUrl
}

func (b *Broadcast) SetActionUrlNull() {
	b.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (b *Broadcast) GetCategory() *util.Nullable[string] {
	if b == nil {
		return nil
	}
	return b.Category
}

func (b *Broadcast) SetCategory(category util.Nullable[string]) {
	b.Category = &category
}

func (b *Broadcast) SetCategoryNull() {
	b.Category = &util.Nullable[string]{IsNull: true}
}

func (b *Broadcast) GetContent() *util.Nullable[string] {
	if b == nil {
		return nil
	}
	return b.Content
}

func (b *Broadcast) SetContent(content util.Nullable[string]) {
	b.Content = &content
}

func (b *Broadcast) SetContentNull() {
	b.Content = &util.Nullable[string]{IsNull: true}
}

func (b *Broadcast) GetCreatedAt() *string {
	if b == nil {
		return nil
	}
	return b.CreatedAt
}

func (b *Broadcast) SetCreatedAt(createdAt string) {
	b.CreatedAt = &createdAt
}

func (b *Broadcast) GetCustomAttributes() *util.Nullable[any] {
	if b == nil {
		return nil
	}
	return b.CustomAttributes
}

func (b *Broadcast) SetCustomAttributes(customAttributes util.Nullable[any]) {
	b.CustomAttributes = &customAttributes
}

func (b *Broadcast) SetCustomAttributesNull() {
	b.CustomAttributes = &util.Nullable[any]{IsNull: true}
}

func (b *Broadcast) GetId() *string {
	if b == nil {
		return nil
	}
	return b.Id
}

func (b *Broadcast) SetId(id string) {
	b.Id = &id
}

func (b *Broadcast) GetOverrides() *util.Nullable[Overrides] {
	if b == nil {
		return nil
	}
	return b.Overrides
}

func (b *Broadcast) SetOverrides(overrides util.Nullable[Overrides]) {
	b.Overrides = &overrides
}

func (b *Broadcast) SetOverridesNull() {
	b.Overrides = &util.Nullable[Overrides]{IsNull: true}
}

func (b *Broadcast) GetRecipients() *util.Nullable[[]any] {
	if b == nil {
		return nil
	}
	return b.Recipients
}

func (b *Broadcast) SetRecipients(recipients util.Nullable[[]any]) {
	b.Recipients = &recipients
}

func (b *Broadcast) SetRecipientsNull() {
	b.Recipients = &util.Nullable[[]any]{IsNull: true}
}

func (b *Broadcast) GetStatus() *util.Nullable[BroadcastStatus] {
	if b == nil {
		return nil
	}
	return b.Status
}

func (b *Broadcast) SetStatus(status util.Nullable[BroadcastStatus]) {
	b.Status = &status
}

func (b *Broadcast) SetStatusNull() {
	b.Status = &util.Nullable[BroadcastStatus]{IsNull: true}
}

func (b *Broadcast) GetTitle() *string {
	if b == nil {
		return nil
	}
	return b.Title
}

func (b *Broadcast) SetTitle(title string) {
	b.Title = &title
}

func (b *Broadcast) GetTopic() *util.Nullable[string] {
	if b == nil {
		return nil
	}
	return b.Topic
}

func (b *Broadcast) SetTopic(topic util.Nullable[string]) {
	b.Topic = &topic
}

func (b *Broadcast) SetTopicNull() {
	b.Topic = &util.Nullable[string]{IsNull: true}
}

func (b Broadcast) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: Broadcast to string"
	}
	return string(jsonData)
}

type Overrides struct {
	Channels  *OverridesChannels `json:"channels,omitempty"`
	Providers *Providers         `json:"providers,omitempty"`
}

func (o *Overrides) GetChannels() *OverridesChannels {
	if o == nil {
		return nil
	}
	return o.Channels
}

func (o *Overrides) SetChannels(channels OverridesChannels) {
	o.Channels = &channels
}

func (o *Overrides) GetProviders() *Providers {
	if o == nil {
		return nil
	}
	return o.Providers
}

func (o *Overrides) SetProviders(providers Providers) {
	o.Providers = &providers
}

func (o Overrides) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: Overrides to string"
	}
	return string(jsonData)
}

type OverridesChannels struct {
	Email      *Email      `json:"email,omitempty"`
	InApp      *InApp      `json:"in_app,omitempty"`
	MobilePush *MobilePush `json:"mobile_push,omitempty"`
	Slack      *Slack      `json:"slack,omitempty"`
	Sms        *Sms        `json:"sms,omitempty"`
	WebPush    *WebPush    `json:"web_push,omitempty"`
}

func (o *OverridesChannels) GetEmail() *Email {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *OverridesChannels) SetEmail(email Email) {
	o.Email = &email
}

func (o *OverridesChannels) GetInApp() *InApp {
	if o == nil {
		return nil
	}
	return o.InApp
}

func (o *OverridesChannels) SetInApp(inApp InApp) {
	o.InApp = &inApp
}

func (o *OverridesChannels) GetMobilePush() *MobilePush {
	if o == nil {
		return nil
	}
	return o.MobilePush
}

func (o *OverridesChannels) SetMobilePush(mobilePush MobilePush) {
	o.MobilePush = &mobilePush
}

func (o *OverridesChannels) GetSlack() *Slack {
	if o == nil {
		return nil
	}
	return o.Slack
}

func (o *OverridesChannels) SetSlack(slack Slack) {
	o.Slack = &slack
}

func (o *OverridesChannels) GetSms() *Sms {
	if o == nil {
		return nil
	}
	return o.Sms
}

func (o *OverridesChannels) SetSms(sms Sms) {
	o.Sms = &sms
}

func (o *OverridesChannels) GetWebPush() *WebPush {
	if o == nil {
		return nil
	}
	return o.WebPush
}

func (o *OverridesChannels) SetWebPush(webPush WebPush) {
	o.WebPush = &webPush
}

func (o OverridesChannels) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OverridesChannels to string"
	}
	return string(jsonData)
}

type Email struct {
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string                `json:"content,omitempty" maxLength:"1048576"`
	Title     *string                `json:"title,omitempty" maxLength:"255" minLength:"1"`
}

func (e *Email) GetActionUrl() *util.Nullable[string] {
	if e == nil {
		return nil
	}
	return e.ActionUrl
}

func (e *Email) SetActionUrl(actionUrl util.Nullable[string]) {
	e.ActionUrl = &actionUrl
}

func (e *Email) SetActionUrlNull() {
	e.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (e *Email) GetContent() *string {
	if e == nil {
		return nil
	}
	return e.Content
}

func (e *Email) SetContent(content string) {
	e.Content = &content
}

func (e *Email) GetTitle() *string {
	if e == nil {
		return nil
	}
	return e.Title
}

func (e *Email) SetTitle(title string) {
	e.Title = &title
}

func (e Email) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Email to string"
	}
	return string(jsonData)
}

type InApp struct {
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string                `json:"content,omitempty" maxLength:"1048576"`
	Title     *string                `json:"title,omitempty" maxLength:"255" minLength:"1"`
}

func (i *InApp) GetActionUrl() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.ActionUrl
}

func (i *InApp) SetActionUrl(actionUrl util.Nullable[string]) {
	i.ActionUrl = &actionUrl
}

func (i *InApp) SetActionUrlNull() {
	i.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (i *InApp) GetContent() *string {
	if i == nil {
		return nil
	}
	return i.Content
}

func (i *InApp) SetContent(content string) {
	i.Content = &content
}

func (i *InApp) GetTitle() *string {
	if i == nil {
		return nil
	}
	return i.Title
}

func (i *InApp) SetTitle(title string) {
	i.Title = &title
}

func (i InApp) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InApp to string"
	}
	return string(jsonData)
}

type MobilePush struct {
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string                `json:"content,omitempty" maxLength:"1048576"`
	Title     *string                `json:"title,omitempty" maxLength:"255" minLength:"1"`
}

func (m *MobilePush) GetActionUrl() *util.Nullable[string] {
	if m == nil {
		return nil
	}
	return m.ActionUrl
}

func (m *MobilePush) SetActionUrl(actionUrl util.Nullable[string]) {
	m.ActionUrl = &actionUrl
}

func (m *MobilePush) SetActionUrlNull() {
	m.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (m *MobilePush) GetContent() *string {
	if m == nil {
		return nil
	}
	return m.Content
}

func (m *MobilePush) SetContent(content string) {
	m.Content = &content
}

func (m *MobilePush) GetTitle() *string {
	if m == nil {
		return nil
	}
	return m.Title
}

func (m *MobilePush) SetTitle(title string) {
	m.Title = &title
}

func (m MobilePush) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MobilePush to string"
	}
	return string(jsonData)
}

type Slack struct {
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string                `json:"content,omitempty" maxLength:"1048576"`
	Title     *string                `json:"title,omitempty" maxLength:"255" minLength:"1"`
}

func (s *Slack) GetActionUrl() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.ActionUrl
}

func (s *Slack) SetActionUrl(actionUrl util.Nullable[string]) {
	s.ActionUrl = &actionUrl
}

func (s *Slack) SetActionUrlNull() {
	s.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (s *Slack) GetContent() *string {
	if s == nil {
		return nil
	}
	return s.Content
}

func (s *Slack) SetContent(content string) {
	s.Content = &content
}

func (s *Slack) GetTitle() *string {
	if s == nil {
		return nil
	}
	return s.Title
}

func (s *Slack) SetTitle(title string) {
	s.Title = &title
}

func (s Slack) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: Slack to string"
	}
	return string(jsonData)
}

type Sms struct {
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string                `json:"content,omitempty" maxLength:"1048576"`
	Title     *string                `json:"title,omitempty" maxLength:"255" minLength:"1"`
}

func (s *Sms) GetActionUrl() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.ActionUrl
}

func (s *Sms) SetActionUrl(actionUrl util.Nullable[string]) {
	s.ActionUrl = &actionUrl
}

func (s *Sms) SetActionUrlNull() {
	s.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (s *Sms) GetContent() *string {
	if s == nil {
		return nil
	}
	return s.Content
}

func (s *Sms) SetContent(content string) {
	s.Content = &content
}

func (s *Sms) GetTitle() *string {
	if s == nil {
		return nil
	}
	return s.Title
}

func (s *Sms) SetTitle(title string) {
	s.Title = &title
}

func (s Sms) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: Sms to string"
	}
	return string(jsonData)
}

type WebPush struct {
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	Content   *string                `json:"content,omitempty" maxLength:"1048576"`
	Title     *string                `json:"title,omitempty" maxLength:"255" minLength:"1"`
}

func (w *WebPush) GetActionUrl() *util.Nullable[string] {
	if w == nil {
		return nil
	}
	return w.ActionUrl
}

func (w *WebPush) SetActionUrl(actionUrl util.Nullable[string]) {
	w.ActionUrl = &actionUrl
}

func (w *WebPush) SetActionUrlNull() {
	w.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (w *WebPush) GetContent() *string {
	if w == nil {
		return nil
	}
	return w.Content
}

func (w *WebPush) SetContent(content string) {
	w.Content = &content
}

func (w *WebPush) GetTitle() *string {
	if w == nil {
		return nil
	}
	return w.Title
}

func (w *WebPush) SetTitle(title string) {
	w.Title = &title
}

func (w WebPush) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebPush to string"
	}
	return string(jsonData)
}

type Providers struct {
	AmazonSes any `json:"amazon_ses,omitempty"`
	Android   any `json:"android,omitempty"`
	Ios       any `json:"ios,omitempty"`
	Mailgun   any `json:"mailgun,omitempty"`
	Postmark  any `json:"postmark,omitempty"`
	Sendgrid  any `json:"sendgrid,omitempty"`
	Slack     any `json:"slack,omitempty"`
}

func (p *Providers) GetAmazonSes() any {
	if p == nil {
		return nil
	}
	return p.AmazonSes
}

func (p *Providers) SetAmazonSes(amazonSes any) {
	p.AmazonSes = &amazonSes
}

func (p *Providers) GetAndroid() any {
	if p == nil {
		return nil
	}
	return p.Android
}

func (p *Providers) SetAndroid(android any) {
	p.Android = &android
}

func (p *Providers) GetIos() any {
	if p == nil {
		return nil
	}
	return p.Ios
}

func (p *Providers) SetIos(ios any) {
	p.Ios = &ios
}

func (p *Providers) GetMailgun() any {
	if p == nil {
		return nil
	}
	return p.Mailgun
}

func (p *Providers) SetMailgun(mailgun any) {
	p.Mailgun = &mailgun
}

func (p *Providers) GetPostmark() any {
	if p == nil {
		return nil
	}
	return p.Postmark
}

func (p *Providers) SetPostmark(postmark any) {
	p.Postmark = &postmark
}

func (p *Providers) GetSendgrid() any {
	if p == nil {
		return nil
	}
	return p.Sendgrid
}

func (p *Providers) SetSendgrid(sendgrid any) {
	p.Sendgrid = &sendgrid
}

func (p *Providers) GetSlack() any {
	if p == nil {
		return nil
	}
	return p.Slack
}

func (p *Providers) SetSlack(slack any) {
	p.Slack = &slack
}

func (p Providers) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Providers to string"
	}
	return string(jsonData)
}

type BroadcastStatus struct {
	Errors  *util.Nullable[[]Errors] `json:"errors,omitempty" required:"true"`
	Status  *StatusStatus            `json:"status,omitempty" required:"true"`
	Summary *Summary                 `json:"summary,omitempty" required:"true"`
}

func (b *BroadcastStatus) GetErrors() *util.Nullable[[]Errors] {
	if b == nil {
		return nil
	}
	return b.Errors
}

func (b *BroadcastStatus) SetErrors(errors util.Nullable[[]Errors]) {
	b.Errors = &errors
}

func (b *BroadcastStatus) SetErrorsNull() {
	b.Errors = &util.Nullable[[]Errors]{IsNull: true}
}

func (b *BroadcastStatus) GetStatus() *StatusStatus {
	if b == nil {
		return nil
	}
	return b.Status
}

func (b *BroadcastStatus) SetStatus(status StatusStatus) {
	b.Status = &status
}

func (b *BroadcastStatus) GetSummary() *Summary {
	if b == nil {
		return nil
	}
	return b.Summary
}

func (b *BroadcastStatus) SetSummary(summary Summary) {
	b.Summary = &summary
}

func (b BroadcastStatus) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BroadcastStatus to string"
	}
	return string(jsonData)
}

type Errors struct {
	Message *string `json:"message,omitempty"`
}

func (e *Errors) GetMessage() *string {
	if e == nil {
		return nil
	}
	return e.Message
}

func (e *Errors) SetMessage(message string) {
	e.Message = &message
}

func (e Errors) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Errors to string"
	}
	return string(jsonData)
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
	Total *int64 `json:"total,omitempty" required:"true"`
}

func (s *Summary) GetFailures() *int64 {
	if s == nil {
		return nil
	}
	return s.Failures
}

func (s *Summary) SetFailures(failures int64) {
	s.Failures = &failures
}

func (s *Summary) GetTotal() *int64 {
	if s == nil {
		return nil
	}
	return s.Total
}

func (s *Summary) SetTotal(total int64) {
	s.Total = &total
}

func (s Summary) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: Summary to string"
	}
	return string(jsonData)
}
