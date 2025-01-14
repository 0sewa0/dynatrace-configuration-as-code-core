/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environment

import (
	"encoding/json"
)

// checks if the Occurrence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Occurrence{}

// Occurrence struct for Occurrence
type Occurrence struct {
	// Count of top parsing field occurrences
	Count *int64 `json:"count,omitempty"`
	// Value of top parsing field occurrence
	Value *string `json:"value,omitempty"`
}

// NewOccurrence instantiates a new Occurrence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOccurrence() *Occurrence {
	this := Occurrence{}
	return &this
}

// NewOccurrenceWithDefaults instantiates a new Occurrence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOccurrenceWithDefaults() *Occurrence {
	this := Occurrence{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Occurrence) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Occurrence) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Occurrence) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *Occurrence) SetCount(v int64) {
	o.Count = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Occurrence) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Occurrence) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Occurrence) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Occurrence) SetValue(v string) {
	o.Value = &v
}

func (o Occurrence) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Occurrence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableOccurrence struct {
	value *Occurrence
	isSet bool
}

func (v NullableOccurrence) Get() *Occurrence {
	return v.value
}

func (v *NullableOccurrence) Set(val *Occurrence) {
	v.value = val
	v.isSet = true
}

func (v NullableOccurrence) IsSet() bool {
	return v.isSet
}

func (v *NullableOccurrence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOccurrence(val *Occurrence) *NullableOccurrence {
	return &NullableOccurrence{value: val, isSet: true}
}

func (v NullableOccurrence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOccurrence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
