// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

type OptionalWordsOption struct {
	value []string
}

func OptionalWords(v ...string) *OptionalWordsOption {
	return &OptionalWordsOption{v}
}

func (o OptionalWordsOption) Get() []string {
	return o.value
}

func (o OptionalWordsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *OptionalWordsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *OptionalWordsOption) Equal(o2 *OptionalWordsOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

func OptionalWordsEqual(o1, o2 *OptionalWordsOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}