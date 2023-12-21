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

// checks if the EventSeverity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSeverity{}

// EventSeverity Additional data on the event severity.
type EventSeverity struct {
	// The metric used in the event severity calculation.
	Context *string `json:"context,omitempty"`
	// The unit of the severity value.
	Unit *string `json:"unit,omitempty"`
	// The value of the severity.
	Value *float32 `json:"value,omitempty"`
}

// NewEventSeverity instantiates a new EventSeverity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSeverity() *EventSeverity {
	this := EventSeverity{}
	return &this
}

// NewEventSeverityWithDefaults instantiates a new EventSeverity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSeverityWithDefaults() *EventSeverity {
	this := EventSeverity{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *EventSeverity) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSeverity) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *EventSeverity) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *EventSeverity) SetContext(v string) {
	o.Context = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *EventSeverity) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSeverity) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *EventSeverity) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *EventSeverity) SetUnit(v string) {
	o.Unit = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EventSeverity) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSeverity) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EventSeverity) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *EventSeverity) SetValue(v float32) {
	o.Value = &v
}

func (o EventSeverity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSeverity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableEventSeverity struct {
	value *EventSeverity
	isSet bool
}

func (v NullableEventSeverity) Get() *EventSeverity {
	return v.value
}

func (v *NullableEventSeverity) Set(val *EventSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSeverity(val *EventSeverity) *NullableEventSeverity {
	return &NullableEventSeverity{value: val, isSet: true}
}

func (v NullableEventSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
