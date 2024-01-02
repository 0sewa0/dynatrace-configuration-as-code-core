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

// checks if the SubscriptionEnvironmentCostListV2Dto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionEnvironmentCostListV2Dto{}

// SubscriptionEnvironmentCostListV2Dto struct for SubscriptionEnvironmentCostListV2Dto
type SubscriptionEnvironmentCostListV2Dto struct {
	// Cost data of the subscription.
	Data []SubscriptionEnvironmentCostV2Dto `json:"data"`
	// The time the subscription data was last modified in `2021-05-01T15:11:00Z` format.
	LastModifiedTime string `json:"lastModifiedTime"`
}

type _SubscriptionEnvironmentCostListV2Dto SubscriptionEnvironmentCostListV2Dto

// NewSubscriptionEnvironmentCostListV2Dto instantiates a new SubscriptionEnvironmentCostListV2Dto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionEnvironmentCostListV2Dto(data []SubscriptionEnvironmentCostV2Dto, lastModifiedTime string) *SubscriptionEnvironmentCostListV2Dto {
	this := SubscriptionEnvironmentCostListV2Dto{}
	this.Data = data
	this.LastModifiedTime = lastModifiedTime
	return &this
}

// NewSubscriptionEnvironmentCostListV2DtoWithDefaults instantiates a new SubscriptionEnvironmentCostListV2Dto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionEnvironmentCostListV2DtoWithDefaults() *SubscriptionEnvironmentCostListV2Dto {
	this := SubscriptionEnvironmentCostListV2Dto{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionEnvironmentCostListV2Dto) GetData() []SubscriptionEnvironmentCostV2Dto {
	if o == nil {
		var ret []SubscriptionEnvironmentCostV2Dto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionEnvironmentCostListV2Dto) GetDataOk() ([]SubscriptionEnvironmentCostV2Dto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionEnvironmentCostListV2Dto) SetData(v []SubscriptionEnvironmentCostV2Dto) {
	o.Data = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *SubscriptionEnvironmentCostListV2Dto) GetLastModifiedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *SubscriptionEnvironmentCostListV2Dto) GetLastModifiedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *SubscriptionEnvironmentCostListV2Dto) SetLastModifiedTime(v string) {
	o.LastModifiedTime = v
}

func (o SubscriptionEnvironmentCostListV2Dto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionEnvironmentCostListV2Dto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["lastModifiedTime"] = o.LastModifiedTime
	return toSerialize, nil
}

func (o *SubscriptionEnvironmentCostListV2Dto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"lastModifiedTime",
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

	varSubscriptionEnvironmentCostListV2Dto := _SubscriptionEnvironmentCostListV2Dto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionEnvironmentCostListV2Dto)

	if err != nil {
		return err
	}

	*o = SubscriptionEnvironmentCostListV2Dto(varSubscriptionEnvironmentCostListV2Dto)

	return err
}

type NullableSubscriptionEnvironmentCostListV2Dto struct {
	value *SubscriptionEnvironmentCostListV2Dto
	isSet bool
}

func (v NullableSubscriptionEnvironmentCostListV2Dto) Get() *SubscriptionEnvironmentCostListV2Dto {
	return v.value
}

func (v *NullableSubscriptionEnvironmentCostListV2Dto) Set(val *SubscriptionEnvironmentCostListV2Dto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionEnvironmentCostListV2Dto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionEnvironmentCostListV2Dto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionEnvironmentCostListV2Dto(val *SubscriptionEnvironmentCostListV2Dto) *NullableSubscriptionEnvironmentCostListV2Dto {
	return &NullableSubscriptionEnvironmentCostListV2Dto{value: val, isSet: true}
}

func (v NullableSubscriptionEnvironmentCostListV2Dto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionEnvironmentCostListV2Dto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
