package notifications

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/util"
)

type Notification struct {
	ActionUrl        *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	ArchivedAt       *util.Nullable[string] `json:"archived_at,omitempty"`
	Category         *util.Nullable[string] `json:"category,omitempty" maxLength:"100"`
	Content          *util.Nullable[string] `json:"content,omitempty" maxLength:"10485760"`
	CreatedAt        *string                `json:"created_at,omitempty" required:"true"`
	CustomAttributes *util.Nullable[any]    `json:"custom_attributes,omitempty"`
	DiscardedAt      *util.Nullable[string] `json:"discarded_at,omitempty"`
	Id               *string                `json:"id,omitempty" required:"true"`
	ReadAt           *util.Nullable[string] `json:"read_at,omitempty"`
	SeenAt           *util.Nullable[string] `json:"seen_at,omitempty"`
	SentAt           *util.Nullable[string] `json:"sent_at,omitempty"`
	Title            *string                `json:"title,omitempty" required:"true" maxLength:"255" minLength:"1"`
	Topic            *util.Nullable[string] `json:"topic,omitempty" maxLength:"100"`
	UpdatedAt        *string                `json:"updated_at,omitempty" required:"true"`
	UserId           *string                `json:"user_id,omitempty" required:"true"`
}

func (n *Notification) GetActionUrl() *util.Nullable[string] {
	if n == nil {
		return nil
	}
	return n.ActionUrl
}

func (n *Notification) SetActionUrl(actionUrl util.Nullable[string]) {
	n.ActionUrl = &actionUrl
}

func (n *Notification) SetActionUrlNull() {
	n.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (n *Notification) GetArchivedAt() *util.Nullable[string] {
	if n == nil {
		return nil
	}
	return n.ArchivedAt
}

func (n *Notification) SetArchivedAt(archivedAt util.Nullable[string]) {
	n.ArchivedAt = &archivedAt
}

func (n *Notification) SetArchivedAtNull() {
	n.ArchivedAt = &util.Nullable[string]{IsNull: true}
}

func (n *Notification) GetCategory() *util.Nullable[string] {
	if n == nil {
		return nil
	}
	return n.Category
}

func (n *Notification) SetCategory(category util.Nullable[string]) {
	n.Category = &category
}

func (n *Notification) SetCategoryNull() {
	n.Category = &util.Nullable[string]{IsNull: true}
}

func (n *Notification) GetContent() *util.Nullable[string] {
	if n == nil {
		return nil
	}
	return n.Content
}

func (n *Notification) SetContent(content util.Nullable[string]) {
	n.Content = &content
}

func (n *Notification) SetContentNull() {
	n.Content = &util.Nullable[string]{IsNull: true}
}

func (n *Notification) GetCreatedAt() *string {
	if n == nil {
		return nil
	}
	return n.CreatedAt
}

func (n *Notification) SetCreatedAt(createdAt string) {
	n.CreatedAt = &createdAt
}

func (n *Notification) GetCustomAttributes() *util.Nullable[any] {
	if n == nil {
		return nil
	}
	return n.CustomAttributes
}

func (n *Notification) SetCustomAttributes(customAttributes util.Nullable[any]) {
	n.CustomAttributes = &customAttributes
}

func (n *Notification) SetCustomAttributesNull() {
	n.CustomAttributes = &util.Nullable[any]{IsNull: true}
}

func (n *Notification) GetDiscardedAt() *util.Nullable[string] {
	if n == nil {
		return nil
	}
	return n.DiscardedAt
}

func (n *Notification) SetDiscardedAt(discardedAt util.Nullable[string]) {
	n.DiscardedAt = &discardedAt
}

func (n *Notification) SetDiscardedAtNull() {
	n.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (n *Notification) GetId() *string {
	if n == nil {
		return nil
	}
	return n.Id
}

func (n *Notification) SetId(id string) {
	n.Id = &id
}

func (n *Notification) GetReadAt() *util.Nullable[string] {
	if n == nil {
		return nil
	}
	return n.ReadAt
}

func (n *Notification) SetReadAt(readAt util.Nullable[string]) {
	n.ReadAt = &readAt
}

func (n *Notification) SetReadAtNull() {
	n.ReadAt = &util.Nullable[string]{IsNull: true}
}

func (n *Notification) GetSeenAt() *util.Nullable[string] {
	if n == nil {
		return nil
	}
	return n.SeenAt
}

func (n *Notification) SetSeenAt(seenAt util.Nullable[string]) {
	n.SeenAt = &seenAt
}

func (n *Notification) SetSeenAtNull() {
	n.SeenAt = &util.Nullable[string]{IsNull: true}
}

func (n *Notification) GetSentAt() *util.Nullable[string] {
	if n == nil {
		return nil
	}
	return n.SentAt
}

func (n *Notification) SetSentAt(sentAt util.Nullable[string]) {
	n.SentAt = &sentAt
}

func (n *Notification) SetSentAtNull() {
	n.SentAt = &util.Nullable[string]{IsNull: true}
}

func (n *Notification) GetTitle() *string {
	if n == nil {
		return nil
	}
	return n.Title
}

func (n *Notification) SetTitle(title string) {
	n.Title = &title
}

func (n *Notification) GetTopic() *util.Nullable[string] {
	if n == nil {
		return nil
	}
	return n.Topic
}

func (n *Notification) SetTopic(topic util.Nullable[string]) {
	n.Topic = &topic
}

func (n *Notification) SetTopicNull() {
	n.Topic = &util.Nullable[string]{IsNull: true}
}

func (n *Notification) GetUpdatedAt() *string {
	if n == nil {
		return nil
	}
	return n.UpdatedAt
}

func (n *Notification) SetUpdatedAt(updatedAt string) {
	n.UpdatedAt = &updatedAt
}

func (n *Notification) GetUserId() *string {
	if n == nil {
		return nil
	}
	return n.UserId
}

func (n *Notification) SetUserId(userId string) {
	n.UserId = &userId
}

func (n Notification) String() string {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return "error converting struct: Notification to string"
	}
	return string(jsonData)
}
