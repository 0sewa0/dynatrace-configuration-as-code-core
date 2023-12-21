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

// checks if the SubscriptionUsageListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionUsageListDto{}

// SubscriptionUsageListDto struct for SubscriptionUsageListDto
type SubscriptionUsageListDto struct {
	// Usage data of the subscription.
	Data []SubscriptionUsageDto `json:"data"`
	// The time the subscription data was last modified in `2021-05-01T15:11:00Z` format.
	LastModifiedTime string `json:"lastModifiedTime"`
}

type _SubscriptionUsageListDto SubscriptionUsageListDto

// NewSubscriptionUsageListDto instantiates a new SubscriptionUsageListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionUsageListDto(data []SubscriptionUsageDto, lastModifiedTime string) *SubscriptionUsageListDto {
	this := SubscriptionUsageListDto{}
	this.Data = data
	this.LastModifiedTime = lastModifiedTime
	return &this
}

// NewSubscriptionUsageListDtoWithDefaults instantiates a new SubscriptionUsageListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionUsageListDtoWithDefaults() *SubscriptionUsageListDto {
	this := SubscriptionUsageListDto{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionUsageListDto) GetData() []SubscriptionUsageDto {
	if o == nil {
		var ret []SubscriptionUsageDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUsageListDto) GetDataOk() ([]SubscriptionUsageDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionUsageListDto) SetData(v []SubscriptionUsageDto) {
	o.Data = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *SubscriptionUsageListDto) GetLastModifiedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUsageListDto) GetLastModifiedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *SubscriptionUsageListDto) SetLastModifiedTime(v string) {
	o.LastModifiedTime = v
}

func (o SubscriptionUsageListDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionUsageListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["lastModifiedTime"] = o.LastModifiedTime
	return toSerialize, nil
}

func (o *SubscriptionUsageListDto) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"lastModifiedTime",
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

	varSubscriptionUsageListDto := _SubscriptionUsageListDto{}

	err = json.Unmarshal(bytes, &varSubscriptionUsageListDto)

	if err != nil {
		return err
	}

	*o = SubscriptionUsageListDto(varSubscriptionUsageListDto)

	return err
}

type NullableSubscriptionUsageListDto struct {
	value *SubscriptionUsageListDto
	isSet bool
}

func (v NullableSubscriptionUsageListDto) Get() *SubscriptionUsageListDto {
	return v.value
}

func (v *NullableSubscriptionUsageListDto) Set(val *SubscriptionUsageListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionUsageListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionUsageListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionUsageListDto(val *SubscriptionUsageListDto) *NullableSubscriptionUsageListDto {
	return &NullableSubscriptionUsageListDto{value: val, isSet: true}
}

func (v NullableSubscriptionUsageListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionUsageListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
