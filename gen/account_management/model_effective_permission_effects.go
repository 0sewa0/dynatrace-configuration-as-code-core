/*
Dynatrace Account Management API

The enterprise management API for Dynatrace SaaS enables automation of operational tasks related to user access and environment lifecycle management.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmanagement

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EffectivePermissionEffects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EffectivePermissionEffects{}

// EffectivePermissionEffects struct for EffectivePermissionEffects
type EffectivePermissionEffects struct {
	// Effect of policy
	Effect string `json:"effect"`
	// Policy condition
	Conditions []Condition `json:"conditions"`
	// A list of effective policies.
	EffectivePolicies []EffectivePolicyWithBinding `json:"effectivePolicies"`
}

type _EffectivePermissionEffects EffectivePermissionEffects

// NewEffectivePermissionEffects instantiates a new EffectivePermissionEffects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectivePermissionEffects(effect string, conditions []Condition, effectivePolicies []EffectivePolicyWithBinding) *EffectivePermissionEffects {
	this := EffectivePermissionEffects{}
	this.Effect = effect
	this.Conditions = conditions
	this.EffectivePolicies = effectivePolicies
	return &this
}

// NewEffectivePermissionEffectsWithDefaults instantiates a new EffectivePermissionEffects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectivePermissionEffectsWithDefaults() *EffectivePermissionEffects {
	this := EffectivePermissionEffects{}
	return &this
}

// GetEffect returns the Effect field value
func (o *EffectivePermissionEffects) GetEffect() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Effect
}

// GetEffectOk returns a tuple with the Effect field value
// and a boolean to check if the value has been set.
func (o *EffectivePermissionEffects) GetEffectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Effect, true
}

// SetEffect sets field value
func (o *EffectivePermissionEffects) SetEffect(v string) {
	o.Effect = v
}

// GetConditions returns the Conditions field value
func (o *EffectivePermissionEffects) GetConditions() []Condition {
	if o == nil {
		var ret []Condition
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *EffectivePermissionEffects) GetConditionsOk() ([]Condition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *EffectivePermissionEffects) SetConditions(v []Condition) {
	o.Conditions = v
}

// GetEffectivePolicies returns the EffectivePolicies field value
func (o *EffectivePermissionEffects) GetEffectivePolicies() []EffectivePolicyWithBinding {
	if o == nil {
		var ret []EffectivePolicyWithBinding
		return ret
	}

	return o.EffectivePolicies
}

// GetEffectivePoliciesOk returns a tuple with the EffectivePolicies field value
// and a boolean to check if the value has been set.
func (o *EffectivePermissionEffects) GetEffectivePoliciesOk() ([]EffectivePolicyWithBinding, bool) {
	if o == nil {
		return nil, false
	}
	return o.EffectivePolicies, true
}

// SetEffectivePolicies sets field value
func (o *EffectivePermissionEffects) SetEffectivePolicies(v []EffectivePolicyWithBinding) {
	o.EffectivePolicies = v
}

func (o EffectivePermissionEffects) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EffectivePermissionEffects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["effect"] = o.Effect
	toSerialize["conditions"] = o.Conditions
	toSerialize["effectivePolicies"] = o.EffectivePolicies
	return toSerialize, nil
}

func (o *EffectivePermissionEffects) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"effect",
		"conditions",
		"effectivePolicies",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEffectivePermissionEffects := _EffectivePermissionEffects{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEffectivePermissionEffects)

	if err != nil {
		return err
	}

	*o = EffectivePermissionEffects(varEffectivePermissionEffects)

	return err
}

type NullableEffectivePermissionEffects struct {
	value *EffectivePermissionEffects
	isSet bool
}

func (v NullableEffectivePermissionEffects) Get() *EffectivePermissionEffects {
	return v.value
}

func (v *NullableEffectivePermissionEffects) Set(val *EffectivePermissionEffects) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectivePermissionEffects) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectivePermissionEffects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectivePermissionEffects(val *EffectivePermissionEffects) *NullableEffectivePermissionEffects {
	return &NullableEffectivePermissionEffects{value: val, isSet: true}
}

func (v NullableEffectivePermissionEffects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectivePermissionEffects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
