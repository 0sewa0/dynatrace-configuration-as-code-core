/*
Dynatrace Account Management API

The enterprise management API for Dynatrace SaaS enables automation of operational tasks related to user access and environment lifecycle management.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmanagement

import (
	"encoding/json"
	"fmt"
)

// checks if the EffectivePolicyWithBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EffectivePolicyWithBinding{}

// EffectivePolicyWithBinding struct for EffectivePolicyWithBinding
type EffectivePolicyWithBinding struct {
	Policy  EffectivePolicy  `json:"policy"`
	Binding EffectiveBinding `json:"binding"`
}

type _EffectivePolicyWithBinding EffectivePolicyWithBinding

// NewEffectivePolicyWithBinding instantiates a new EffectivePolicyWithBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectivePolicyWithBinding(policy EffectivePolicy, binding EffectiveBinding) *EffectivePolicyWithBinding {
	this := EffectivePolicyWithBinding{}
	this.Policy = policy
	this.Binding = binding
	return &this
}

// NewEffectivePolicyWithBindingWithDefaults instantiates a new EffectivePolicyWithBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectivePolicyWithBindingWithDefaults() *EffectivePolicyWithBinding {
	this := EffectivePolicyWithBinding{}
	return &this
}

// GetPolicy returns the Policy field value
func (o *EffectivePolicyWithBinding) GetPolicy() EffectivePolicy {
	if o == nil {
		var ret EffectivePolicy
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *EffectivePolicyWithBinding) GetPolicyOk() (*EffectivePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *EffectivePolicyWithBinding) SetPolicy(v EffectivePolicy) {
	o.Policy = v
}

// GetBinding returns the Binding field value
func (o *EffectivePolicyWithBinding) GetBinding() EffectiveBinding {
	if o == nil {
		var ret EffectiveBinding
		return ret
	}

	return o.Binding
}

// GetBindingOk returns a tuple with the Binding field value
// and a boolean to check if the value has been set.
func (o *EffectivePolicyWithBinding) GetBindingOk() (*EffectiveBinding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Binding, true
}

// SetBinding sets field value
func (o *EffectivePolicyWithBinding) SetBinding(v EffectiveBinding) {
	o.Binding = v
}

func (o EffectivePolicyWithBinding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EffectivePolicyWithBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policy"] = o.Policy
	toSerialize["binding"] = o.Binding
	return toSerialize, nil
}

func (o *EffectivePolicyWithBinding) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"policy",
		"binding",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEffectivePolicyWithBinding := _EffectivePolicyWithBinding{}

	err = json.Unmarshal(bytes, &varEffectivePolicyWithBinding)

	if err != nil {
		return err
	}

	*o = EffectivePolicyWithBinding(varEffectivePolicyWithBinding)

	return err
}

type NullableEffectivePolicyWithBinding struct {
	value *EffectivePolicyWithBinding
	isSet bool
}

func (v NullableEffectivePolicyWithBinding) Get() *EffectivePolicyWithBinding {
	return v.value
}

func (v *NullableEffectivePolicyWithBinding) Set(val *EffectivePolicyWithBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectivePolicyWithBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectivePolicyWithBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectivePolicyWithBinding(val *EffectivePolicyWithBinding) *NullableEffectivePolicyWithBinding {
	return &NullableEffectivePolicyWithBinding{value: val, isSet: true}
}

func (v NullableEffectivePolicyWithBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectivePolicyWithBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
