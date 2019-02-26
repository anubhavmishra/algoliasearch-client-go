// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type PageOption struct {
	value int
}

func Page(v int) PageOption {
	return PageOption{v}
}

func (o PageOption) Get() int {
	return o.value
}

func (o PageOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *PageOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 0
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
