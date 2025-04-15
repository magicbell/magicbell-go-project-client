package events

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type Event struct {
	Code      *int64                 `json:"code,omitempty"`
	Context   *util.Nullable[any]    `json:"context,omitempty"`
	Id        *string                `json:"id,omitempty" required:"true"`
	Level     *string                `json:"level,omitempty"`
	Log       *util.Nullable[string] `json:"log,omitempty"`
	Timestamp *string                `json:"timestamp,omitempty" required:"true"`
	Type_     *string                `json:"type,omitempty" required:"true"`
}

func (e *Event) GetCode() *int64 {
	if e == nil {
		return nil
	}
	return e.Code
}

func (e *Event) SetCode(code int64) {
	e.Code = &code
}

func (e *Event) GetContext() *util.Nullable[any] {
	if e == nil {
		return nil
	}
	return e.Context
}

func (e *Event) SetContext(context util.Nullable[any]) {
	e.Context = &context
}

func (e *Event) SetContextNull() {
	e.Context = &util.Nullable[any]{IsNull: true}
}

func (e *Event) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *Event) SetId(id string) {
	e.Id = &id
}

func (e *Event) GetLevel() *string {
	if e == nil {
		return nil
	}
	return e.Level
}

func (e *Event) SetLevel(level string) {
	e.Level = &level
}

func (e *Event) GetLog() *util.Nullable[string] {
	if e == nil {
		return nil
	}
	return e.Log
}

func (e *Event) SetLog(log util.Nullable[string]) {
	e.Log = &log
}

func (e *Event) SetLogNull() {
	e.Log = &util.Nullable[string]{IsNull: true}
}

func (e *Event) GetTimestamp() *string {
	if e == nil {
		return nil
	}
	return e.Timestamp
}

func (e *Event) SetTimestamp(timestamp string) {
	e.Timestamp = &timestamp
}

func (e *Event) GetType_() *string {
	if e == nil {
		return nil
	}
	return e.Type_
}

func (e *Event) SetType_(type_ string) {
	e.Type_ = &type_
}

func (e Event) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Event to string"
	}
	return string(jsonData)
}
