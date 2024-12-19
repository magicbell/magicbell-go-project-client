package events

import (
	"encoding/json"
)

type Event struct {
	Code      *int64  `json:"code,omitempty"`
	Context   any     `json:"context,omitempty"`
	Id        *string `json:"id,omitempty" required:"true"`
	Level     *string `json:"level,omitempty"`
	Log       *string `json:"log,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" required:"true"`
	Type_     *string `json:"type,omitempty" required:"true"`
	touched   map[string]bool
}

func (e *Event) GetCode() *int64 {
	if e == nil {
		return nil
	}
	return e.Code
}

func (e *Event) SetCode(code int64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Code"] = true
	e.Code = &code
}

func (e *Event) SetCodeNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Code"] = true
	e.Code = nil
}

func (e *Event) GetContext() any {
	if e == nil {
		return nil
	}
	return e.Context
}

func (e *Event) SetContext(context any) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Context"] = true
	e.Context = context
}

func (e *Event) SetContextNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Context"] = true
	e.Context = nil
}

func (e *Event) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *Event) SetId(id string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Id"] = true
	e.Id = &id
}

func (e *Event) SetIdNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Id"] = true
	e.Id = nil
}

func (e *Event) GetLevel() *string {
	if e == nil {
		return nil
	}
	return e.Level
}

func (e *Event) SetLevel(level string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Level"] = true
	e.Level = &level
}

func (e *Event) SetLevelNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Level"] = true
	e.Level = nil
}

func (e *Event) GetLog() *string {
	if e == nil {
		return nil
	}
	return e.Log
}

func (e *Event) SetLog(log string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Log"] = true
	e.Log = &log
}

func (e *Event) SetLogNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Log"] = true
	e.Log = nil
}

func (e *Event) GetTimestamp() *string {
	if e == nil {
		return nil
	}
	return e.Timestamp
}

func (e *Event) SetTimestamp(timestamp string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Timestamp"] = true
	e.Timestamp = &timestamp
}

func (e *Event) SetTimestampNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Timestamp"] = true
	e.Timestamp = nil
}

func (e *Event) GetType_() *string {
	if e == nil {
		return nil
	}
	return e.Type_
}

func (e *Event) SetType_(type_ string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Type_"] = true
	e.Type_ = &type_
}

func (e *Event) SetType_Nil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Type_"] = true
	e.Type_ = nil
}
func (e Event) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["Code"] && e.Code == nil {
		data["code"] = nil
	} else if e.Code != nil {
		data["code"] = e.Code
	}

	if e.touched["Context"] && e.Context == nil {
		data["context"] = nil
	} else if e.Context != nil {
		data["context"] = e.Context
	}

	if e.touched["Id"] && e.Id == nil {
		data["id"] = nil
	} else if e.Id != nil {
		data["id"] = e.Id
	}

	if e.touched["Level"] && e.Level == nil {
		data["level"] = nil
	} else if e.Level != nil {
		data["level"] = e.Level
	}

	if e.touched["Log"] && e.Log == nil {
		data["log"] = nil
	} else if e.Log != nil {
		data["log"] = e.Log
	}

	if e.touched["Timestamp"] && e.Timestamp == nil {
		data["timestamp"] = nil
	} else if e.Timestamp != nil {
		data["timestamp"] = e.Timestamp
	}

	if e.touched["Type_"] && e.Type_ == nil {
		data["type"] = nil
	} else if e.Type_ != nil {
		data["type"] = e.Type_
	}

	return json.Marshal(data)
}
