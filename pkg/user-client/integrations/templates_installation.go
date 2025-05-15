package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/util"
)

type TemplatesInstallation struct {
	Category *util.Nullable[string] `json:"category,omitempty"`
	Channel  *string                `json:"channel,omitempty" required:"true"`
	Text     *string                `json:"text,omitempty" required:"true"`
}

func (t *TemplatesInstallation) GetCategory() *util.Nullable[string] {
	if t == nil {
		return nil
	}
	return t.Category
}

func (t *TemplatesInstallation) SetCategory(category util.Nullable[string]) {
	t.Category = &category
}

func (t *TemplatesInstallation) SetCategoryNull() {
	t.Category = &util.Nullable[string]{IsNull: true}
}

func (t *TemplatesInstallation) GetChannel() *string {
	if t == nil {
		return nil
	}
	return t.Channel
}

func (t *TemplatesInstallation) SetChannel(channel string) {
	t.Channel = &channel
}

func (t *TemplatesInstallation) GetText() *string {
	if t == nil {
		return nil
	}
	return t.Text
}

func (t *TemplatesInstallation) SetText(text string) {
	t.Text = &text
}

func (t TemplatesInstallation) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplatesInstallation to string"
	}
	return string(jsonData)
}
