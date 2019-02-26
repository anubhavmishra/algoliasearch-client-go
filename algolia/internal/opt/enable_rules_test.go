// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestEnableRules(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.EnableRulesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.EnableRules(false),
		},
		{
			opts:     []interface{}{opt.EnableRules(true)},
			expected: opt.EnableRules(true),
		},
		{
			opts:     []interface{}{opt.EnableRules(false)},
			expected: opt.EnableRules(false),
		},
	} {
		var (
			in  = ExtractEnableRules(c.opts...)
			out opt.EnableRulesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
