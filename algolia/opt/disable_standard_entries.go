// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

// DisableStandardEntriesOption is a wrapper for an DisableStandardEntries option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type DisableStandardEntriesOption struct {
	value map[string]map[string]bool
}

// DisableStandardEntries wraps the given value into a DisableStandardEntriesOption.
func DisableStandardEntries(v map[string]map[string]bool) *DisableStandardEntriesOption {
	return &DisableStandardEntriesOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *DisableStandardEntriesOption) Get() map[string]map[string]bool {
	if o == nil {
		return map[string]map[string]bool{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// DisableStandardEntriesOption.
func (o DisableStandardEntriesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// DisableStandardEntriesOption.
func (o *DisableStandardEntriesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = map[string]map[string]bool{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *DisableStandardEntriesOption) Equal(o2 *DisableStandardEntriesOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.value, map[string]map[string]bool{})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.value, map[string]map[string]bool{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// DisableStandardEntriesEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func DisableStandardEntriesEqual(o1, o2 *DisableStandardEntriesOption) bool {
	return o1.Equal(o2)
}