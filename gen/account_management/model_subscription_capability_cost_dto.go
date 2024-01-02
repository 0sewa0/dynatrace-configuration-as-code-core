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

// checks if the SubscriptionCapabilityCostDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionCapabilityCostDto{}

// SubscriptionCapabilityCostDto struct for SubscriptionCapabilityCostDto
type SubscriptionCapabilityCostDto struct {
	// The start time for the capability cost in `2021-05-01T15:11:00Z` format.
	StartTime string `json:"startTime"`
	// The end time for the capability cost in `2021-05-01T15:11:00Z` format.
	EndTime string `json:"endTime"`
	// The total cost for all the capabilities.
	Value float32 `json:"value"`
	// The currency of the cost.
	CurrencyCode string `json:"currencyCode"`
	// The key of the subscription capability.
	CapabilityKey string `json:"capabilityKey"`
	// The display name of the subscription capability.
	CapabilityName string `json:"capabilityName"`
}

type _SubscriptionCapabilityCostDto SubscriptionCapabilityCostDto

// NewSubscriptionCapabilityCostDto instantiates a new SubscriptionCapabilityCostDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionCapabilityCostDto(startTime string, endTime string, value float32, currencyCode string, capabilityKey string, capabilityName string) *SubscriptionCapabilityCostDto {
	this := SubscriptionCapabilityCostDto{}
	this.StartTime = startTime
	this.EndTime = endTime
	this.Value = value
	this.CurrencyCode = currencyCode
	this.CapabilityKey = capabilityKey
	this.CapabilityName = capabilityName
	return &this
}

// NewSubscriptionCapabilityCostDtoWithDefaults instantiates a new SubscriptionCapabilityCostDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionCapabilityCostDtoWithDefaults() *SubscriptionCapabilityCostDto {
	this := SubscriptionCapabilityCostDto{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *SubscriptionCapabilityCostDto) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCapabilityCostDto) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *SubscriptionCapabilityCostDto) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *SubscriptionCapabilityCostDto) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCapabilityCostDto) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *SubscriptionCapabilityCostDto) SetEndTime(v string) {
	o.EndTime = v
}

// GetValue returns the Value field value
func (o *SubscriptionCapabilityCostDto) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCapabilityCostDto) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SubscriptionCapabilityCostDto) SetValue(v float32) {
	o.Value = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *SubscriptionCapabilityCostDto) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCapabilityCostDto) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *SubscriptionCapabilityCostDto) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetCapabilityKey returns the CapabilityKey field value
func (o *SubscriptionCapabilityCostDto) GetCapabilityKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CapabilityKey
}

// GetCapabilityKeyOk returns a tuple with the CapabilityKey field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCapabilityCostDto) GetCapabilityKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapabilityKey, true
}

// SetCapabilityKey sets field value
func (o *SubscriptionCapabilityCostDto) SetCapabilityKey(v string) {
	o.CapabilityKey = v
}

// GetCapabilityName returns the CapabilityName field value
func (o *SubscriptionCapabilityCostDto) GetCapabilityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CapabilityName
}

// GetCapabilityNameOk returns a tuple with the CapabilityName field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCapabilityCostDto) GetCapabilityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapabilityName, true
}

// SetCapabilityName sets field value
func (o *SubscriptionCapabilityCostDto) SetCapabilityName(v string) {
	o.CapabilityName = v
}

func (o SubscriptionCapabilityCostDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionCapabilityCostDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	toSerialize["value"] = o.Value
	toSerialize["currencyCode"] = o.CurrencyCode
	toSerialize["capabilityKey"] = o.CapabilityKey
	toSerialize["capabilityName"] = o.CapabilityName
	return toSerialize, nil
}

func (o *SubscriptionCapabilityCostDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startTime",
		"endTime",
		"value",
		"currencyCode",
		"capabilityKey",
		"capabilityName",
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

	varSubscriptionCapabilityCostDto := _SubscriptionCapabilityCostDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionCapabilityCostDto)

	if err != nil {
		return err
	}

	*o = SubscriptionCapabilityCostDto(varSubscriptionCapabilityCostDto)

	return err
}

type NullableSubscriptionCapabilityCostDto struct {
	value *SubscriptionCapabilityCostDto
	isSet bool
}

func (v NullableSubscriptionCapabilityCostDto) Get() *SubscriptionCapabilityCostDto {
	return v.value
}

func (v *NullableSubscriptionCapabilityCostDto) Set(val *SubscriptionCapabilityCostDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionCapabilityCostDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionCapabilityCostDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionCapabilityCostDto(val *SubscriptionCapabilityCostDto) *NullableSubscriptionCapabilityCostDto {
	return &NullableSubscriptionCapabilityCostDto{value: val, isSet: true}
}

func (v NullableSubscriptionCapabilityCostDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionCapabilityCostDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
