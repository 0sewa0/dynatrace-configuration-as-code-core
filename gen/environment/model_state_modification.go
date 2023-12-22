/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environment

import (
	"encoding/json"
	"fmt"
)

// checks if the StateModification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateModification{}

// StateModification Operation state to be set for all third-party Synthetic monitors
type StateModification struct {
	// The new operation state for all third-party Synthetic monitors.
	State string `json:"state"`
}

type _StateModification StateModification

// NewStateModification instantiates a new StateModification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateModification(state string) *StateModification {
	this := StateModification{}
	this.State = state
	return &this
}

// NewStateModificationWithDefaults instantiates a new StateModification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateModificationWithDefaults() *StateModification {
	this := StateModification{}
	return &this
}

// GetState returns the State field value
func (o *StateModification) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *StateModification) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *StateModification) SetState(v string) {
	o.State = v
}

func (o StateModification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateModification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *StateModification) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
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

	varStateModification := _StateModification{}

	err = json.Unmarshal(bytes, &varStateModification)

	if err != nil {
		return err
	}

	*o = StateModification(varStateModification)

	return err
}

type NullableStateModification struct {
	value *StateModification
	isSet bool
}

func (v NullableStateModification) Get() *StateModification {
	return v.value
}

func (v *NullableStateModification) Set(val *StateModification) {
	v.value = val
	v.isSet = true
}

func (v NullableStateModification) IsSet() bool {
	return v.isSet
}

func (v *NullableStateModification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateModification(val *StateModification) *NullableStateModification {
	return &NullableStateModification{value: val, isSet: true}
}

func (v NullableStateModification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateModification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}