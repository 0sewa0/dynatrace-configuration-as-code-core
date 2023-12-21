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

// checks if the TimeZoneDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeZoneDto{}

// TimeZoneDto struct for TimeZoneDto
type TimeZoneDto struct {
	// The UTC-based name of the time zone.
	DisplayName string `json:"displayName"`
	// The standard name of the time zone.
	Name string `json:"name"`
}

type _TimeZoneDto TimeZoneDto

// NewTimeZoneDto instantiates a new TimeZoneDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeZoneDto(displayName string, name string) *TimeZoneDto {
	this := TimeZoneDto{}
	this.DisplayName = displayName
	this.Name = name
	return &this
}

// NewTimeZoneDtoWithDefaults instantiates a new TimeZoneDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeZoneDtoWithDefaults() *TimeZoneDto {
	this := TimeZoneDto{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *TimeZoneDto) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *TimeZoneDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *TimeZoneDto) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetName returns the Name field value
func (o *TimeZoneDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TimeZoneDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TimeZoneDto) SetName(v string) {
	o.Name = v
}

func (o TimeZoneDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeZoneDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *TimeZoneDto) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"name",
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

	varTimeZoneDto := _TimeZoneDto{}

	err = json.Unmarshal(bytes, &varTimeZoneDto)

	if err != nil {
		return err
	}

	*o = TimeZoneDto(varTimeZoneDto)

	return err
}

type NullableTimeZoneDto struct {
	value *TimeZoneDto
	isSet bool
}

func (v NullableTimeZoneDto) Get() *TimeZoneDto {
	return v.value
}

func (v *NullableTimeZoneDto) Set(val *TimeZoneDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeZoneDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeZoneDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeZoneDto(val *TimeZoneDto) *NullableTimeZoneDto {
	return &NullableTimeZoneDto{value: val, isSet: true}
}

func (v NullableTimeZoneDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeZoneDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
