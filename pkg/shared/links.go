package shared

import (
	"encoding/json"
)

type Links struct {
	First   *string `json:"first,omitempty"`
	Next    *string `json:"next,omitempty"`
	Prev    *string `json:"prev,omitempty"`
	touched map[string]bool
}

func (l *Links) GetFirst() *string {
	if l == nil {
		return nil
	}
	return l.First
}

func (l *Links) SetFirst(first string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["First"] = true
	l.First = &first
}

func (l *Links) SetFirstNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["First"] = true
	l.First = nil
}

func (l *Links) GetNext() *string {
	if l == nil {
		return nil
	}
	return l.Next
}

func (l *Links) SetNext(next string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Next"] = true
	l.Next = &next
}

func (l *Links) SetNextNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Next"] = true
	l.Next = nil
}

func (l *Links) GetPrev() *string {
	if l == nil {
		return nil
	}
	return l.Prev
}

func (l *Links) SetPrev(prev string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Prev"] = true
	l.Prev = &prev
}

func (l *Links) SetPrevNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Prev"] = true
	l.Prev = nil
}
func (l Links) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["First"] && l.First == nil {
		data["first"] = nil
	} else if l.First != nil {
		data["first"] = l.First
	}

	if l.touched["Next"] && l.Next == nil {
		data["next"] = nil
	} else if l.Next != nil {
		data["next"] = l.Next
	}

	if l.touched["Prev"] && l.Prev == nil {
		data["prev"] = nil
	} else if l.Prev != nil {
		data["prev"] = l.Prev
	}

	return json.Marshal(data)
}
