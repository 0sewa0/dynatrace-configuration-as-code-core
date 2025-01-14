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

// checks if the LongProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LongProperty{}

// LongProperty A custom property of the user action with a Long value.
type LongProperty struct {
	// The custom key of the property.
	Key *string `json:"key,omitempty"`
	// The Long value of the property.
	Value *int64 `json:"value,omitempty"`
}

// NewLongProperty instantiates a new LongProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLongProperty() *LongProperty {
	this := LongProperty{}
	return &this
}

// NewLongPropertyWithDefaults instantiates a new LongProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLongPropertyWithDefaults() *LongProperty {
	this := LongProperty{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LongProperty) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongProperty) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LongProperty) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LongProperty) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LongProperty) GetValue() int64 {
	if o == nil || IsNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LongProperty) GetValueOk() (*int64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LongProperty) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *LongProperty) SetValue(v int64) {
	o.Value = &v
}

func (o LongProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LongProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableLongProperty struct {
	value *LongProperty
	isSet bool
}

func (v NullableLongProperty) Get() *LongProperty {
	return v.value
}

func (v *NullableLongProperty) Set(val *LongProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableLongProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableLongProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLongProperty(val *LongProperty) *NullableLongProperty {
	return &NullableLongProperty{value: val, isSet: true}
}

func (v NullableLongProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLongProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
