// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type RemoveWordsIfNoResultsOption struct {
	value string
}

func RemoveWordsIfNoResults(v string) *RemoveWordsIfNoResultsOption {
	return &RemoveWordsIfNoResultsOption{v}
}

func (o RemoveWordsIfNoResultsOption) Get() string {
	return o.value
}

func (o RemoveWordsIfNoResultsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *RemoveWordsIfNoResultsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "none"
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *RemoveWordsIfNoResultsOption) Equal(o2 *RemoveWordsIfNoResultsOption) bool {
	if o2 == nil {
		return o.value == "none"
	}
	return o.value == o2.value
}

func RemoveWordsIfNoResultsEqual(o1, o2 *RemoveWordsIfNoResultsOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}