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
	"time"
)

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event struct for Event
type Event struct {
	// The UUID of the environment that raised the event.
	EnvironmentUuid string `json:"environmentUuid"`
	// The subscription capability that raised the event.
	Capability string `json:"capability"`
	// The time when the event was raised, in `2021-05-01T15:11:00Z` format.
	Date time.Time `json:"date"`
	// The time when the notification was created, in `2021-05-01T15:11:00Z` format.
	CreatedAt time.Time `json:"createdAt"`
	// The severity of the event.
	Severity string `json:"severity"`
	// The message from the event.
	Message string `json:"message"`
	// The type of event: forecast or usage.
	EventType string `json:"eventType"`
	// The notification level of the event.
	NotificationLevel string `json:"notificationLevel"`
}

type _Event Event

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(environmentUuid string, capability string, date time.Time, createdAt time.Time, severity string, message string, eventType string, notificationLevel string) *Event {
	this := Event{}
	this.EnvironmentUuid = environmentUuid
	this.Capability = capability
	this.Date = date
	this.CreatedAt = createdAt
	this.Severity = severity
	this.Message = message
	this.EventType = eventType
	this.NotificationLevel = notificationLevel
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetEnvironmentUuid returns the EnvironmentUuid field value
func (o *Event) GetEnvironmentUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentUuid
}

// GetEnvironmentUuidOk returns a tuple with the EnvironmentUuid field value
// and a boolean to check if the value has been set.
func (o *Event) GetEnvironmentUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentUuid, true
}

// SetEnvironmentUuid sets field value
func (o *Event) SetEnvironmentUuid(v string) {
	o.EnvironmentUuid = v
}

// GetCapability returns the Capability field value
func (o *Event) GetCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *Event) GetCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *Event) SetCapability(v string) {
	o.Capability = v
}

// GetDate returns the Date field value
func (o *Event) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Event) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Event) SetDate(v time.Time) {
	o.Date = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Event) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Event) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Event) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetSeverity returns the Severity field value
func (o *Event) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *Event) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *Event) SetSeverity(v string) {
	o.Severity = v
}

// GetMessage returns the Message field value
func (o *Event) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Event) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Event) SetMessage(v string) {
	o.Message = v
}

// GetEventType returns the EventType field value
func (o *Event) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *Event) SetEventType(v string) {
	o.EventType = v
}

// GetNotificationLevel returns the NotificationLevel field value
func (o *Event) GetNotificationLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationLevel
}

// GetNotificationLevelOk returns a tuple with the NotificationLevel field value
// and a boolean to check if the value has been set.
func (o *Event) GetNotificationLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationLevel, true
}

// SetNotificationLevel sets field value
func (o *Event) SetNotificationLevel(v string) {
	o.NotificationLevel = v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentUuid"] = o.EnvironmentUuid
	toSerialize["capability"] = o.Capability
	toSerialize["date"] = o.Date
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["severity"] = o.Severity
	toSerialize["message"] = o.Message
	toSerialize["eventType"] = o.EventType
	toSerialize["notificationLevel"] = o.NotificationLevel
	return toSerialize, nil
}

func (o *Event) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environmentUuid",
		"capability",
		"date",
		"createdAt",
		"severity",
		"message",
		"eventType",
		"notificationLevel",
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

	varEvent := _Event{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEvent)

	if err != nil {
		return err
	}

	*o = Event(varEvent)

	return err
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
