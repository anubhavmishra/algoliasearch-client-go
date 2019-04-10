// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type SortFacetValuesByOption struct {
	value string
}

func SortFacetValuesBy(v string) *SortFacetValuesByOption {
	return &SortFacetValuesByOption{v}
}

func (o SortFacetValuesByOption) Get() string {
	return o.value
}

func (o SortFacetValuesByOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *SortFacetValuesByOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "count"
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *SortFacetValuesByOption) Equal(o2 *SortFacetValuesByOption) bool {
	if o2 == nil {
		return o.value == "count"
	}
	return o.value == o2.value
}

func SortFacetValuesByEqual(o1, o2 *SortFacetValuesByOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}