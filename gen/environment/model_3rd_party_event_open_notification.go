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

// checks if the Model3rdPartyEventOpenNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model3rdPartyEventOpenNotification{}

// Model3rdPartyEventOpenNotification The open third-party synthetic event.
type Model3rdPartyEventOpenNotification struct {
	// The unique ID of the event.
	EventId string `json:"eventId"`
	// The type of the event.
	EventType string `json:"eventType"`
	// The list of IDs of third-party synthetic locations where the event happens.
	LocationIds []string `json:"locationIds"`
	// The name of the event.
	Name string `json:"name"`
	// The cause of the event.
	Reason string `json:"reason"`
	// The start timestamp of the event, in UTC milliseconds.
	StartTimestamp int64 `json:"startTimestamp"`
	// The ID of the third-party synthetic monitor.
	TestId string `json:"testId"`
}

type _Model3rdPartyEventOpenNotification Model3rdPartyEventOpenNotification

// NewModel3rdPartyEventOpenNotification instantiates a new Model3rdPartyEventOpenNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel3rdPartyEventOpenNotification(eventId string, eventType string, locationIds []string, name string, reason string, startTimestamp int64, testId string) *Model3rdPartyEventOpenNotification {
	this := Model3rdPartyEventOpenNotification{}
	this.EventId = eventId
	this.EventType = eventType
	this.LocationIds = locationIds
	this.Name = name
	this.Reason = reason
	this.StartTimestamp = startTimestamp
	this.TestId = testId
	return &this
}

// NewModel3rdPartyEventOpenNotificationWithDefaults instantiates a new Model3rdPartyEventOpenNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel3rdPartyEventOpenNotificationWithDefaults() *Model3rdPartyEventOpenNotification {
	this := Model3rdPartyEventOpenNotification{}
	return &this
}

// GetEventId returns the EventId field value
func (o *Model3rdPartyEventOpenNotification) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartyEventOpenNotification) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *Model3rdPartyEventOpenNotification) SetEventId(v string) {
	o.EventId = v
}

// GetEventType returns the EventType field value
func (o *Model3rdPartyEventOpenNotification) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartyEventOpenNotification) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *Model3rdPartyEventOpenNotification) SetEventType(v string) {
	o.EventType = v
}

// GetLocationIds returns the LocationIds field value
func (o *Model3rdPartyEventOpenNotification) GetLocationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LocationIds
}

// GetLocationIdsOk returns a tuple with the LocationIds field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartyEventOpenNotification) GetLocationIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocationIds, true
}

// SetLocationIds sets field value
func (o *Model3rdPartyEventOpenNotification) SetLocationIds(v []string) {
	o.LocationIds = v
}

// GetName returns the Name field value
func (o *Model3rdPartyEventOpenNotification) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartyEventOpenNotification) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Model3rdPartyEventOpenNotification) SetName(v string) {
	o.Name = v
}

// GetReason returns the Reason field value
func (o *Model3rdPartyEventOpenNotification) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartyEventOpenNotification) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *Model3rdPartyEventOpenNotification) SetReason(v string) {
	o.Reason = v
}

// GetStartTimestamp returns the StartTimestamp field value
func (o *Model3rdPartyEventOpenNotification) GetStartTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartyEventOpenNotification) GetStartTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTimestamp, true
}

// SetStartTimestamp sets field value
func (o *Model3rdPartyEventOpenNotification) SetStartTimestamp(v int64) {
	o.StartTimestamp = v
}

// GetTestId returns the TestId field value
func (o *Model3rdPartyEventOpenNotification) GetTestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartyEventOpenNotification) GetTestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestId, true
}

// SetTestId sets field value
func (o *Model3rdPartyEventOpenNotification) SetTestId(v string) {
	o.TestId = v
}

func (o Model3rdPartyEventOpenNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model3rdPartyEventOpenNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	toSerialize["eventType"] = o.EventType
	toSerialize["locationIds"] = o.LocationIds
	toSerialize["name"] = o.Name
	toSerialize["reason"] = o.Reason
	toSerialize["startTimestamp"] = o.StartTimestamp
	toSerialize["testId"] = o.TestId
	return toSerialize, nil
}

func (o *Model3rdPartyEventOpenNotification) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventId",
		"eventType",
		"locationIds",
		"name",
		"reason",
		"startTimestamp",
		"testId",
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

	varModel3rdPartyEventOpenNotification := _Model3rdPartyEventOpenNotification{}

	err = json.Unmarshal(bytes, &varModel3rdPartyEventOpenNotification)

	if err != nil {
		return err
	}

	*o = Model3rdPartyEventOpenNotification(varModel3rdPartyEventOpenNotification)

	return err
}

type NullableModel3rdPartyEventOpenNotification struct {
	value *Model3rdPartyEventOpenNotification
	isSet bool
}

func (v NullableModel3rdPartyEventOpenNotification) Get() *Model3rdPartyEventOpenNotification {
	return v.value
}

func (v *NullableModel3rdPartyEventOpenNotification) Set(val *Model3rdPartyEventOpenNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableModel3rdPartyEventOpenNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableModel3rdPartyEventOpenNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel3rdPartyEventOpenNotification(val *Model3rdPartyEventOpenNotification) *NullableModel3rdPartyEventOpenNotification {
	return &NullableModel3rdPartyEventOpenNotification{value: val, isSet: true}
}

func (v NullableModel3rdPartyEventOpenNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel3rdPartyEventOpenNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
