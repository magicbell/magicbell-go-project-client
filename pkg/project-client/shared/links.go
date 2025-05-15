package shared

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type Links struct {
	First *string                `json:"first,omitempty"`
	Next  *util.Nullable[string] `json:"next,omitempty"`
	Prev  *util.Nullable[string] `json:"prev,omitempty"`
}

func (l *Links) GetFirst() *string {
	if l == nil {
		return nil
	}
	return l.First
}

func (l *Links) SetFirst(first string) {
	l.First = &first
}

func (l *Links) GetNext() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.Next
}

func (l *Links) SetNext(next util.Nullable[string]) {
	l.Next = &next
}

func (l *Links) SetNextNull() {
	l.Next = &util.Nullable[string]{IsNull: true}
}

func (l *Links) GetPrev() *util.Nullable[string] {
	if l == nil {
		return nil
	}
	return l.Prev
}

func (l *Links) SetPrev(prev util.Nullable[string]) {
	l.Prev = &prev
}

func (l *Links) SetPrevNull() {
	l.Prev = &util.Nullable[string]{IsNull: true}
}

func (l Links) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: Links to string"
	}
	return string(jsonData)
}
