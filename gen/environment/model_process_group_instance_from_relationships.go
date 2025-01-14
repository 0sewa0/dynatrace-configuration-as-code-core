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

// checks if the ProcessGroupInstanceFromRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessGroupInstanceFromRelationships{}

// ProcessGroupInstanceFromRelationships struct for ProcessGroupInstanceFromRelationships
type ProcessGroupInstanceFromRelationships struct {
	IsInstanceOf      []string `json:"isInstanceOf,omitempty"`
	IsNetworkClientOf []string `json:"isNetworkClientOf,omitempty"`
	IsProcessOf       []string `json:"isProcessOf,omitempty"`
}

// NewProcessGroupInstanceFromRelationships instantiates a new ProcessGroupInstanceFromRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupInstanceFromRelationships() *ProcessGroupInstanceFromRelationships {
	this := ProcessGroupInstanceFromRelationships{}
	return &this
}

// NewProcessGroupInstanceFromRelationshipsWithDefaults instantiates a new ProcessGroupInstanceFromRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupInstanceFromRelationshipsWithDefaults() *ProcessGroupInstanceFromRelationships {
	this := ProcessGroupInstanceFromRelationships{}
	return &this
}

// GetIsInstanceOf returns the IsInstanceOf field value if set, zero value otherwise.
func (o *ProcessGroupInstanceFromRelationships) GetIsInstanceOf() []string {
	if o == nil || IsNil(o.IsInstanceOf) {
		var ret []string
		return ret
	}
	return o.IsInstanceOf
}

// GetIsInstanceOfOk returns a tuple with the IsInstanceOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupInstanceFromRelationships) GetIsInstanceOfOk() ([]string, bool) {
	if o == nil || IsNil(o.IsInstanceOf) {
		return nil, false
	}
	return o.IsInstanceOf, true
}

// HasIsInstanceOf returns a boolean if a field has been set.
func (o *ProcessGroupInstanceFromRelationships) HasIsInstanceOf() bool {
	if o != nil && !IsNil(o.IsInstanceOf) {
		return true
	}

	return false
}

// SetIsInstanceOf gets a reference to the given []string and assigns it to the IsInstanceOf field.
func (o *ProcessGroupInstanceFromRelationships) SetIsInstanceOf(v []string) {
	o.IsInstanceOf = v
}

// GetIsNetworkClientOf returns the IsNetworkClientOf field value if set, zero value otherwise.
func (o *ProcessGroupInstanceFromRelationships) GetIsNetworkClientOf() []string {
	if o == nil || IsNil(o.IsNetworkClientOf) {
		var ret []string
		return ret
	}
	return o.IsNetworkClientOf
}

// GetIsNetworkClientOfOk returns a tuple with the IsNetworkClientOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupInstanceFromRelationships) GetIsNetworkClientOfOk() ([]string, bool) {
	if o == nil || IsNil(o.IsNetworkClientOf) {
		return nil, false
	}
	return o.IsNetworkClientOf, true
}

// HasIsNetworkClientOf returns a boolean if a field has been set.
func (o *ProcessGroupInstanceFromRelationships) HasIsNetworkClientOf() bool {
	if o != nil && !IsNil(o.IsNetworkClientOf) {
		return true
	}

	return false
}

// SetIsNetworkClientOf gets a reference to the given []string and assigns it to the IsNetworkClientOf field.
func (o *ProcessGroupInstanceFromRelationships) SetIsNetworkClientOf(v []string) {
	o.IsNetworkClientOf = v
}

// GetIsProcessOf returns the IsProcessOf field value if set, zero value otherwise.
func (o *ProcessGroupInstanceFromRelationships) GetIsProcessOf() []string {
	if o == nil || IsNil(o.IsProcessOf) {
		var ret []string
		return ret
	}
	return o.IsProcessOf
}

// GetIsProcessOfOk returns a tuple with the IsProcessOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupInstanceFromRelationships) GetIsProcessOfOk() ([]string, bool) {
	if o == nil || IsNil(o.IsProcessOf) {
		return nil, false
	}
	return o.IsProcessOf, true
}

// HasIsProcessOf returns a boolean if a field has been set.
func (o *ProcessGroupInstanceFromRelationships) HasIsProcessOf() bool {
	if o != nil && !IsNil(o.IsProcessOf) {
		return true
	}

	return false
}

// SetIsProcessOf gets a reference to the given []string and assigns it to the IsProcessOf field.
func (o *ProcessGroupInstanceFromRelationships) SetIsProcessOf(v []string) {
	o.IsProcessOf = v
}

func (o ProcessGroupInstanceFromRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessGroupInstanceFromRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsInstanceOf) {
		toSerialize["isInstanceOf"] = o.IsInstanceOf
	}
	if !IsNil(o.IsNetworkClientOf) {
		toSerialize["isNetworkClientOf"] = o.IsNetworkClientOf
	}
	if !IsNil(o.IsProcessOf) {
		toSerialize["isProcessOf"] = o.IsProcessOf
	}
	return toSerialize, nil
}

type NullableProcessGroupInstanceFromRelationships struct {
	value *ProcessGroupInstanceFromRelationships
	isSet bool
}

func (v NullableProcessGroupInstanceFromRelationships) Get() *ProcessGroupInstanceFromRelationships {
	return v.value
}

func (v *NullableProcessGroupInstanceFromRelationships) Set(val *ProcessGroupInstanceFromRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupInstanceFromRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupInstanceFromRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupInstanceFromRelationships(val *ProcessGroupInstanceFromRelationships) *NullableProcessGroupInstanceFromRelationships {
	return &NullableProcessGroupInstanceFromRelationships{value: val, isSet: true}
}

func (v NullableProcessGroupInstanceFromRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupInstanceFromRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
