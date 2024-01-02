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

// checks if the Binding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Binding{}

// Binding struct for Binding
type Binding struct {
	// The ID of the policy.
	PolicyUuid string `json:"policyUuid"`
	// A list of user groups to which the policy applies.
	Groups []string `json:"groups"`
}

type _Binding Binding

// NewBinding instantiates a new Binding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinding(policyUuid string, groups []string) *Binding {
	this := Binding{}
	this.PolicyUuid = policyUuid
	this.Groups = groups
	return &this
}

// NewBindingWithDefaults instantiates a new Binding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindingWithDefaults() *Binding {
	this := Binding{}
	return &this
}

// GetPolicyUuid returns the PolicyUuid field value
func (o *Binding) GetPolicyUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyUuid
}

// GetPolicyUuidOk returns a tuple with the PolicyUuid field value
// and a boolean to check if the value has been set.
func (o *Binding) GetPolicyUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyUuid, true
}

// SetPolicyUuid sets field value
func (o *Binding) SetPolicyUuid(v string) {
	o.PolicyUuid = v
}

// GetGroups returns the Groups field value
func (o *Binding) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *Binding) GetGroupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *Binding) SetGroups(v []string) {
	o.Groups = v
}

func (o Binding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Binding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policyUuid"] = o.PolicyUuid
	toSerialize["groups"] = o.Groups
	return toSerialize, nil
}

func (o *Binding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"policyUuid",
		"groups",
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

	varBinding := _Binding{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBinding)

	if err != nil {
		return err
	}

	*o = Binding(varBinding)

	return err
}

type NullableBinding struct {
	value *Binding
	isSet bool
}

func (v NullableBinding) Get() *Binding {
	return v.value
}

func (v *NullableBinding) Set(val *Binding) {
	v.value = val
	v.isSet = true
}

func (v NullableBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinding(val *Binding) *NullableBinding {
	return &NullableBinding{value: val, isSet: true}
}

func (v NullableBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
