/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environment

import (
	"encoding/json"
)

// checks if the EventStoreResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventStoreResult{}

// EventStoreResult Contains IDs of all custom events, created by an event push call.
type EventStoreResult struct {
	// List of correlation IDs for problem-opening-events.
	StoredCorrelationIds []string `json:"storedCorrelationIds,omitempty"`
	// List of event IDs for information-only-events.    This field is provided for compatibility reasons. You should use the values from the **storedIds** field instead.
	StoredEventIds []int64 `json:"storedEventIds,omitempty"`
	// List of **encoded** event IDs for information-only-events.
	StoredIds []string `json:"storedIds,omitempty"`
}

// NewEventStoreResult instantiates a new EventStoreResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventStoreResult() *EventStoreResult {
	this := EventStoreResult{}
	return &this
}

// NewEventStoreResultWithDefaults instantiates a new EventStoreResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventStoreResultWithDefaults() *EventStoreResult {
	this := EventStoreResult{}
	return &this
}

// GetStoredCorrelationIds returns the StoredCorrelationIds field value if set, zero value otherwise.
func (o *EventStoreResult) GetStoredCorrelationIds() []string {
	if o == nil || IsNil(o.StoredCorrelationIds) {
		var ret []string
		return ret
	}
	return o.StoredCorrelationIds
}

// GetStoredCorrelationIdsOk returns a tuple with the StoredCorrelationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStoreResult) GetStoredCorrelationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StoredCorrelationIds) {
		return nil, false
	}
	return o.StoredCorrelationIds, true
}

// HasStoredCorrelationIds returns a boolean if a field has been set.
func (o *EventStoreResult) HasStoredCorrelationIds() bool {
	if o != nil && !IsNil(o.StoredCorrelationIds) {
		return true
	}

	return false
}

// SetStoredCorrelationIds gets a reference to the given []string and assigns it to the StoredCorrelationIds field.
func (o *EventStoreResult) SetStoredCorrelationIds(v []string) {
	o.StoredCorrelationIds = v
}

// GetStoredEventIds returns the StoredEventIds field value if set, zero value otherwise.
func (o *EventStoreResult) GetStoredEventIds() []int64 {
	if o == nil || IsNil(o.StoredEventIds) {
		var ret []int64
		return ret
	}
	return o.StoredEventIds
}

// GetStoredEventIdsOk returns a tuple with the StoredEventIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStoreResult) GetStoredEventIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.StoredEventIds) {
		return nil, false
	}
	return o.StoredEventIds, true
}

// HasStoredEventIds returns a boolean if a field has been set.
func (o *EventStoreResult) HasStoredEventIds() bool {
	if o != nil && !IsNil(o.StoredEventIds) {
		return true
	}

	return false
}

// SetStoredEventIds gets a reference to the given []int64 and assigns it to the StoredEventIds field.
func (o *EventStoreResult) SetStoredEventIds(v []int64) {
	o.StoredEventIds = v
}

// GetStoredIds returns the StoredIds field value if set, zero value otherwise.
func (o *EventStoreResult) GetStoredIds() []string {
	if o == nil || IsNil(o.StoredIds) {
		var ret []string
		return ret
	}
	return o.StoredIds
}

// GetStoredIdsOk returns a tuple with the StoredIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventStoreResult) GetStoredIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StoredIds) {
		return nil, false
	}
	return o.StoredIds, true
}

// HasStoredIds returns a boolean if a field has been set.
func (o *EventStoreResult) HasStoredIds() bool {
	if o != nil && !IsNil(o.StoredIds) {
		return true
	}

	return false
}

// SetStoredIds gets a reference to the given []string and assigns it to the StoredIds field.
func (o *EventStoreResult) SetStoredIds(v []string) {
	o.StoredIds = v
}

func (o EventStoreResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventStoreResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StoredCorrelationIds) {
		toSerialize["storedCorrelationIds"] = o.StoredCorrelationIds
	}
	if !IsNil(o.StoredEventIds) {
		toSerialize["storedEventIds"] = o.StoredEventIds
	}
	if !IsNil(o.StoredIds) {
		toSerialize["storedIds"] = o.StoredIds
	}
	return toSerialize, nil
}

type NullableEventStoreResult struct {
	value *EventStoreResult
	isSet bool
}

func (v NullableEventStoreResult) Get() *EventStoreResult {
	return v.value
}

func (v *NullableEventStoreResult) Set(val *EventStoreResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEventStoreResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEventStoreResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventStoreResult(val *EventStoreResult) *NullableEventStoreResult {
	return &NullableEventStoreResult{value: val, isSet: true}
}

func (v NullableEventStoreResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventStoreResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
