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

// checks if the MonitorCollectionElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorCollectionElement{}

// MonitorCollectionElement The short representation of a synthetic monitor.
type MonitorCollectionElement struct {
	// The state of a synthetic monitor.
	Enabled bool `json:"enabled"`
	// The ID of a synthetic object.
	EntityId string `json:"entityId"`
	// The name of a synthetic object.
	Name string `json:"name"`
	// The type of a synthetic monitor.
	Type string `json:"type"`
}

type _MonitorCollectionElement MonitorCollectionElement

// NewMonitorCollectionElement instantiates a new MonitorCollectionElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorCollectionElement(enabled bool, entityId string, name string, type_ string) *MonitorCollectionElement {
	this := MonitorCollectionElement{}
	this.Enabled = enabled
	this.EntityId = entityId
	this.Name = name
	this.Type = type_
	return &this
}

// NewMonitorCollectionElementWithDefaults instantiates a new MonitorCollectionElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorCollectionElementWithDefaults() *MonitorCollectionElement {
	this := MonitorCollectionElement{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *MonitorCollectionElement) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MonitorCollectionElement) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MonitorCollectionElement) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEntityId returns the EntityId field value
func (o *MonitorCollectionElement) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *MonitorCollectionElement) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *MonitorCollectionElement) SetEntityId(v string) {
	o.EntityId = v
}

// GetName returns the Name field value
func (o *MonitorCollectionElement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorCollectionElement) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MonitorCollectionElement) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *MonitorCollectionElement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MonitorCollectionElement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MonitorCollectionElement) SetType(v string) {
	o.Type = v
}

func (o MonitorCollectionElement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorCollectionElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["entityId"] = o.EntityId
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *MonitorCollectionElement) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"entityId",
		"name",
		"type",
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

	varMonitorCollectionElement := _MonitorCollectionElement{}

	err = json.Unmarshal(bytes, &varMonitorCollectionElement)

	if err != nil {
		return err
	}

	*o = MonitorCollectionElement(varMonitorCollectionElement)

	return err
}

type NullableMonitorCollectionElement struct {
	value *MonitorCollectionElement
	isSet bool
}

func (v NullableMonitorCollectionElement) Get() *MonitorCollectionElement {
	return v.value
}

func (v *NullableMonitorCollectionElement) Set(val *MonitorCollectionElement) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorCollectionElement) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorCollectionElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorCollectionElement(val *MonitorCollectionElement) *NullableMonitorCollectionElement {
	return &NullableMonitorCollectionElement{value: val, isSet: true}
}

func (v NullableMonitorCollectionElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorCollectionElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
