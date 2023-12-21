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

// checks if the EnvironmentListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentListDto{}

// EnvironmentListDto struct for EnvironmentListDto
type EnvironmentListDto struct {
	// Lists all environments in an account
	Data []EnvironmentDto `json:"data"`
}

type _EnvironmentListDto EnvironmentListDto

// NewEnvironmentListDto instantiates a new EnvironmentListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentListDto(data []EnvironmentDto) *EnvironmentListDto {
	this := EnvironmentListDto{}
	this.Data = data
	return &this
}

// NewEnvironmentListDtoWithDefaults instantiates a new EnvironmentListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentListDtoWithDefaults() *EnvironmentListDto {
	this := EnvironmentListDto{}
	return &this
}

// GetData returns the Data field value
func (o *EnvironmentListDto) GetData() []EnvironmentDto {
	if o == nil {
		var ret []EnvironmentDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EnvironmentListDto) GetDataOk() ([]EnvironmentDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *EnvironmentListDto) SetData(v []EnvironmentDto) {
	o.Data = v
}

func (o EnvironmentListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EnvironmentListDto) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varEnvironmentListDto := _EnvironmentListDto{}

	err = json.Unmarshal(bytes, &varEnvironmentListDto)

	if err != nil {
		return err
	}

	*o = EnvironmentListDto(varEnvironmentListDto)

	return err
}

type NullableEnvironmentListDto struct {
	value *EnvironmentListDto
	isSet bool
}

func (v NullableEnvironmentListDto) Get() *EnvironmentListDto {
	return v.value
}

func (v *NullableEnvironmentListDto) Set(val *EnvironmentListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentListDto(val *EnvironmentListDto) *NullableEnvironmentListDto {
	return &NullableEnvironmentListDto{value: val, isSet: true}
}

func (v NullableEnvironmentListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
