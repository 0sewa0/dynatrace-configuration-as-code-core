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

// checks if the LoadingTimeThresholdsPolicyDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadingTimeThresholdsPolicyDto{}

// LoadingTimeThresholdsPolicyDto Performance thresholds configuration.
type LoadingTimeThresholdsPolicyDto struct {
	// Performance threshold is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// The list of performance threshold rules.
	Thresholds []LoadingTimeThreshold `json:"thresholds"`
}

type _LoadingTimeThresholdsPolicyDto LoadingTimeThresholdsPolicyDto

// NewLoadingTimeThresholdsPolicyDto instantiates a new LoadingTimeThresholdsPolicyDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadingTimeThresholdsPolicyDto(enabled bool, thresholds []LoadingTimeThreshold) *LoadingTimeThresholdsPolicyDto {
	this := LoadingTimeThresholdsPolicyDto{}
	this.Enabled = enabled
	this.Thresholds = thresholds
	return &this
}

// NewLoadingTimeThresholdsPolicyDtoWithDefaults instantiates a new LoadingTimeThresholdsPolicyDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadingTimeThresholdsPolicyDtoWithDefaults() *LoadingTimeThresholdsPolicyDto {
	this := LoadingTimeThresholdsPolicyDto{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *LoadingTimeThresholdsPolicyDto) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LoadingTimeThresholdsPolicyDto) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LoadingTimeThresholdsPolicyDto) SetEnabled(v bool) {
	o.Enabled = v
}

// GetThresholds returns the Thresholds field value
func (o *LoadingTimeThresholdsPolicyDto) GetThresholds() []LoadingTimeThreshold {
	if o == nil {
		var ret []LoadingTimeThreshold
		return ret
	}

	return o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value
// and a boolean to check if the value has been set.
func (o *LoadingTimeThresholdsPolicyDto) GetThresholdsOk() ([]LoadingTimeThreshold, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thresholds, true
}

// SetThresholds sets field value
func (o *LoadingTimeThresholdsPolicyDto) SetThresholds(v []LoadingTimeThreshold) {
	o.Thresholds = v
}

func (o LoadingTimeThresholdsPolicyDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadingTimeThresholdsPolicyDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["thresholds"] = o.Thresholds
	return toSerialize, nil
}

func (o *LoadingTimeThresholdsPolicyDto) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"thresholds",
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

	varLoadingTimeThresholdsPolicyDto := _LoadingTimeThresholdsPolicyDto{}

	err = json.Unmarshal(bytes, &varLoadingTimeThresholdsPolicyDto)

	if err != nil {
		return err
	}

	*o = LoadingTimeThresholdsPolicyDto(varLoadingTimeThresholdsPolicyDto)

	return err
}

type NullableLoadingTimeThresholdsPolicyDto struct {
	value *LoadingTimeThresholdsPolicyDto
	isSet bool
}

func (v NullableLoadingTimeThresholdsPolicyDto) Get() *LoadingTimeThresholdsPolicyDto {
	return v.value
}

func (v *NullableLoadingTimeThresholdsPolicyDto) Set(val *LoadingTimeThresholdsPolicyDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadingTimeThresholdsPolicyDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadingTimeThresholdsPolicyDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadingTimeThresholdsPolicyDto(val *LoadingTimeThresholdsPolicyDto) *NullableLoadingTimeThresholdsPolicyDto {
	return &NullableLoadingTimeThresholdsPolicyDto{value: val, isSet: true}
}

func (v NullableLoadingTimeThresholdsPolicyDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadingTimeThresholdsPolicyDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
