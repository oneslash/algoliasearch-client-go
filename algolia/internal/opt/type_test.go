package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.TypeOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Type([]string{}...),
		},
		{
			opts:     []interface{}{opt.Type("value1")},
			expected: opt.Type("value1"),
		},
		{
			opts:     []interface{}{opt.Type("value1", "value2", "value3")},
			expected: opt.Type("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractType(c.opts...)
			out opt.TypeOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
