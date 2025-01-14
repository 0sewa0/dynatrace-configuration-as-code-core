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

// checks if the ServiceToRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceToRelationships{}

// ServiceToRelationships struct for ServiceToRelationships
type ServiceToRelationships struct {
	Calls []string `json:"calls,omitempty"`
}

// NewServiceToRelationships instantiates a new ServiceToRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceToRelationships() *ServiceToRelationships {
	this := ServiceToRelationships{}
	return &this
}

// NewServiceToRelationshipsWithDefaults instantiates a new ServiceToRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceToRelationshipsWithDefaults() *ServiceToRelationships {
	this := ServiceToRelationships{}
	return &this
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *ServiceToRelationships) GetCalls() []string {
	if o == nil || IsNil(o.Calls) {
		var ret []string
		return ret
	}
	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceToRelationships) GetCallsOk() ([]string, bool) {
	if o == nil || IsNil(o.Calls) {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *ServiceToRelationships) HasCalls() bool {
	if o != nil && !IsNil(o.Calls) {
		return true
	}

	return false
}

// SetCalls gets a reference to the given []string and assigns it to the Calls field.
func (o *ServiceToRelationships) SetCalls(v []string) {
	o.Calls = v
}

func (o ServiceToRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceToRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Calls) {
		toSerialize["calls"] = o.Calls
	}
	return toSerialize, nil
}

type NullableServiceToRelationships struct {
	value *ServiceToRelationships
	isSet bool
}

func (v NullableServiceToRelationships) Get() *ServiceToRelationships {
	return v.value
}

func (v *NullableServiceToRelationships) Set(val *ServiceToRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceToRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceToRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceToRelationships(val *ServiceToRelationships) *NullableServiceToRelationships {
	return &NullableServiceToRelationships{value: val, isSet: true}
}

func (v NullableServiceToRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceToRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
