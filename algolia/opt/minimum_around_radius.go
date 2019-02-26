// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type MinimumAroundRadiusOption struct {
	value int
}

func MinimumAroundRadius(v int) MinimumAroundRadiusOption {
	return MinimumAroundRadiusOption{v}
}

func (o MinimumAroundRadiusOption) Get() int {
	return o.value
}

func (o MinimumAroundRadiusOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *MinimumAroundRadiusOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 0
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
