/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environment

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EventDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventDto{}

// EventDto struct for EventDto
type EventDto struct {
	// Event identifier
	EntityId string `json:"entityId"`
	// Event name
	Name string `json:"name"`
	// Event sequence number
	SequenceNumber int32 `json:"sequenceNumber"`
}

type _EventDto EventDto

// NewEventDto instantiates a new EventDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventDto(entityId string, name string, sequenceNumber int32) *EventDto {
	this := EventDto{}
	this.EntityId = entityId
	this.Name = name
	this.SequenceNumber = sequenceNumber
	return &this
}

// NewEventDtoWithDefaults instantiates a new EventDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDtoWithDefaults() *EventDto {
	this := EventDto{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *EventDto) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *EventDto) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *EventDto) SetEntityId(v string) {
	o.EntityId = v
}

// GetName returns the Name field value
func (o *EventDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventDto) SetName(v string) {
	o.Name = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *EventDto) GetSequenceNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *EventDto) GetSequenceNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *EventDto) SetSequenceNumber(v int32) {
	o.SequenceNumber = v
}

func (o EventDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityId"] = o.EntityId
	toSerialize["name"] = o.Name
	toSerialize["sequenceNumber"] = o.SequenceNumber
	return toSerialize, nil
}

func (o *EventDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entityId",
		"name",
		"sequenceNumber",
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

	varEventDto := _EventDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventDto)

	if err != nil {
		return err
	}

	*o = EventDto(varEventDto)

	return err
}

type NullableEventDto struct {
	value *EventDto
	isSet bool
}

func (v NullableEventDto) Get() *EventDto {
	return v.value
}

func (v *NullableEventDto) Set(val *EventDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDto(val *EventDto) *NullableEventDto {
	return &NullableEventDto{value: val, isSet: true}
}

func (v NullableEventDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
