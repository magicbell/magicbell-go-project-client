package notifications

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type DeliveryPlan struct {
	Category       *util.Nullable[string] `json:"category,omitempty"`
	CompletedAt    *util.Nullable[string] `json:"completed_at,omitempty"`
	Id             *string                `json:"id,omitempty"`
	NextStep       *int64                 `json:"next_step,omitempty"`
	NotificationId *string                `json:"notification_id,omitempty"`
	ScheduledAt    *string                `json:"scheduled_at,omitempty"`
	StartedAt      *util.Nullable[string] `json:"started_at,omitempty"`
	Status         *string                `json:"status,omitempty"`
	Topic          *util.Nullable[string] `json:"topic,omitempty"`
	UserId         *string                `json:"user_id,omitempty"`
}

func (d *DeliveryPlan) GetCategory() *util.Nullable[string] {
	if d == nil {
		return nil
	}
	return d.Category
}

func (d *DeliveryPlan) SetCategory(category util.Nullable[string]) {
	d.Category = &category
}

func (d *DeliveryPlan) SetCategoryNull() {
	d.Category = &util.Nullable[string]{IsNull: true}
}

func (d *DeliveryPlan) GetCompletedAt() *util.Nullable[string] {
	if d == nil {
		return nil
	}
	return d.CompletedAt
}

func (d *DeliveryPlan) SetCompletedAt(completedAt util.Nullable[string]) {
	d.CompletedAt = &completedAt
}

func (d *DeliveryPlan) SetCompletedAtNull() {
	d.CompletedAt = &util.Nullable[string]{IsNull: true}
}

func (d *DeliveryPlan) GetId() *string {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *DeliveryPlan) SetId(id string) {
	d.Id = &id
}

func (d *DeliveryPlan) GetNextStep() *int64 {
	if d == nil {
		return nil
	}
	return d.NextStep
}

func (d *DeliveryPlan) SetNextStep(nextStep int64) {
	d.NextStep = &nextStep
}

func (d *DeliveryPlan) GetNotificationId() *string {
	if d == nil {
		return nil
	}
	return d.NotificationId
}

func (d *DeliveryPlan) SetNotificationId(notificationId string) {
	d.NotificationId = &notificationId
}

func (d *DeliveryPlan) GetScheduledAt() *string {
	if d == nil {
		return nil
	}
	return d.ScheduledAt
}

func (d *DeliveryPlan) SetScheduledAt(scheduledAt string) {
	d.ScheduledAt = &scheduledAt
}

func (d *DeliveryPlan) GetStartedAt() *util.Nullable[string] {
	if d == nil {
		return nil
	}
	return d.StartedAt
}

func (d *DeliveryPlan) SetStartedAt(startedAt util.Nullable[string]) {
	d.StartedAt = &startedAt
}

func (d *DeliveryPlan) SetStartedAtNull() {
	d.StartedAt = &util.Nullable[string]{IsNull: true}
}

func (d *DeliveryPlan) GetStatus() *string {
	if d == nil {
		return nil
	}
	return d.Status
}

func (d *DeliveryPlan) SetStatus(status string) {
	d.Status = &status
}

func (d *DeliveryPlan) GetTopic() *util.Nullable[string] {
	if d == nil {
		return nil
	}
	return d.Topic
}

func (d *DeliveryPlan) SetTopic(topic util.Nullable[string]) {
	d.Topic = &topic
}

func (d *DeliveryPlan) SetTopicNull() {
	d.Topic = &util.Nullable[string]{IsNull: true}
}

func (d *DeliveryPlan) GetUserId() *string {
	if d == nil {
		return nil
	}
	return d.UserId
}

func (d *DeliveryPlan) SetUserId(userId string) {
	d.UserId = &userId
}

func (d DeliveryPlan) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeliveryPlan to string"
	}
	return string(jsonData)
}
