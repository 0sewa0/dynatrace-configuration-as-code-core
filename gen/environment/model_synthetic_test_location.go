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

// checks if the SyntheticTestLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyntheticTestLocation{}

// SyntheticTestLocation Synthetic test location.
type SyntheticTestLocation struct {
	// The location is enabled/disabled. Default is `true`, enabling the location.
	Enabled *bool `json:"enabled,omitempty"`
	// The ID of the location.
	Id string `json:"id"`
}

type _SyntheticTestLocation SyntheticTestLocation

// NewSyntheticTestLocation instantiates a new SyntheticTestLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticTestLocation(id string) *SyntheticTestLocation {
	this := SyntheticTestLocation{}
	this.Id = id
	return &this
}

// NewSyntheticTestLocationWithDefaults instantiates a new SyntheticTestLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticTestLocationWithDefaults() *SyntheticTestLocation {
	this := SyntheticTestLocation{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SyntheticTestLocation) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestLocation) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SyntheticTestLocation) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SyntheticTestLocation) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value
func (o *SyntheticTestLocation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SyntheticTestLocation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SyntheticTestLocation) SetId(v string) {
	o.Id = v
}

func (o SyntheticTestLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyntheticTestLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *SyntheticTestLocation) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varSyntheticTestLocation := _SyntheticTestLocation{}

	err = json.Unmarshal(bytes, &varSyntheticTestLocation)

	if err != nil {
		return err
	}

	*o = SyntheticTestLocation(varSyntheticTestLocation)

	return err
}

type NullableSyntheticTestLocation struct {
	value *SyntheticTestLocation
	isSet bool
}

func (v NullableSyntheticTestLocation) Get() *SyntheticTestLocation {
	return v.value
}

func (v *NullableSyntheticTestLocation) Set(val *SyntheticTestLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticTestLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticTestLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticTestLocation(val *SyntheticTestLocation) *NullableSyntheticTestLocation {
	return &NullableSyntheticTestLocation{value: val, isSet: true}
}

func (v NullableSyntheticTestLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticTestLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}