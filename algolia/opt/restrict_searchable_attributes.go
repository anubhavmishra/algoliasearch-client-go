// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

type RestrictSearchableAttributesOption struct {
	value []string
}

func RestrictSearchableAttributes(v ...string) *RestrictSearchableAttributesOption {
	return &RestrictSearchableAttributesOption{v}
}

func (o RestrictSearchableAttributesOption) Get() []string {
	return o.value
}

func (o RestrictSearchableAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *RestrictSearchableAttributesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *RestrictSearchableAttributesOption) Equal(o2 *RestrictSearchableAttributesOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

func RestrictSearchableAttributesEqual(o1, o2 *RestrictSearchableAttributesOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}