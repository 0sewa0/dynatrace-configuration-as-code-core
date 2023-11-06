/*
Dynatrace Account Management API

The enterprise management API for Dynatrace SaaS enables automation of operational tasks related to user access and environment lifecycle management.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmanagement

import (
	"encoding/json"
)

// checks if the SubscriptionAccountDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAccountDto{}

// SubscriptionAccountDto struct for SubscriptionAccountDto
type SubscriptionAccountDto struct {
	// The UUID of the account
	Uuid string `json:"uuid"`
}

// NewSubscriptionAccountDto instantiates a new SubscriptionAccountDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAccountDto(uuid string) *SubscriptionAccountDto {
	this := SubscriptionAccountDto{}
	this.Uuid = uuid
	return &this
}

// NewSubscriptionAccountDtoWithDefaults instantiates a new SubscriptionAccountDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAccountDtoWithDefaults() *SubscriptionAccountDto {
	this := SubscriptionAccountDto{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *SubscriptionAccountDto) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SubscriptionAccountDto) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SubscriptionAccountDto) SetUuid(v string) {
	o.Uuid = v
}

func (o SubscriptionAccountDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAccountDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

type NullableSubscriptionAccountDto struct {
	value *SubscriptionAccountDto
	isSet bool
}

func (v NullableSubscriptionAccountDto) Get() *SubscriptionAccountDto {
	return v.value
}

func (v *NullableSubscriptionAccountDto) Set(val *SubscriptionAccountDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAccountDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAccountDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAccountDto(val *SubscriptionAccountDto) *NullableSubscriptionAccountDto {
	return &NullableSubscriptionAccountDto{value: val, isSet: true}
}

func (v NullableSubscriptionAccountDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAccountDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}