// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAnalyticsTags(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AnalyticsTagsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AnalyticsTags([]string{}...),
		},
		{
			opts:     []interface{}{opt.AnalyticsTags("value1")},
			expected: opt.AnalyticsTags("value1"),
		},
		{
			opts:     []interface{}{opt.AnalyticsTags("value1", "value2", "value3")},
			expected: opt.AnalyticsTags("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractAnalyticsTags(c.opts...)
			out opt.AnalyticsTagsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestAnalyticsTags_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.AnalyticsTagsOption
	}{
		{
			payload:  `""`,
			expected: opt.AnalyticsTags([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.AnalyticsTags("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.AnalyticsTags("value1", "value2", "value3"),
		},
	} {
		var got opt.AnalyticsTagsOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}
