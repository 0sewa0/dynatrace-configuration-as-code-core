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

// checks if the MaintenanceWindowRecurrence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaintenanceWindowRecurrence{}

// MaintenanceWindowRecurrence The recurrence of the maintenance window.
type MaintenanceWindowRecurrence struct {
	// The day of the week for weekly maintenance.    The format is the full name of the day in upper case, for example `WEDNESDAY`.
	Day *string `json:"day,omitempty"`
	// The day of the month for monthly maintenance.
	DayOfMonth *int32 `json:"dayOfMonth,omitempty"`
	// The duration of the maintenance window in minutes.
	Duration int32 `json:"duration"`
	// The start time of the maintenance window. The format is `HH:mm`.
	Start string `json:"start"`
}

type _MaintenanceWindowRecurrence MaintenanceWindowRecurrence

// NewMaintenanceWindowRecurrence instantiates a new MaintenanceWindowRecurrence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindowRecurrence(duration int32, start string) *MaintenanceWindowRecurrence {
	this := MaintenanceWindowRecurrence{}
	this.Duration = duration
	this.Start = start
	return &this
}

// NewMaintenanceWindowRecurrenceWithDefaults instantiates a new MaintenanceWindowRecurrence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowRecurrenceWithDefaults() *MaintenanceWindowRecurrence {
	this := MaintenanceWindowRecurrence{}
	return &this
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *MaintenanceWindowRecurrence) GetDay() string {
	if o == nil || IsNil(o.Day) {
		var ret string
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRecurrence) GetDayOk() (*string, bool) {
	if o == nil || IsNil(o.Day) {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *MaintenanceWindowRecurrence) HasDay() bool {
	if o != nil && !IsNil(o.Day) {
		return true
	}

	return false
}

// SetDay gets a reference to the given string and assigns it to the Day field.
func (o *MaintenanceWindowRecurrence) SetDay(v string) {
	o.Day = &v
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise.
func (o *MaintenanceWindowRecurrence) GetDayOfMonth() int32 {
	if o == nil || IsNil(o.DayOfMonth) {
		var ret int32
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRecurrence) GetDayOfMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.DayOfMonth) {
		return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *MaintenanceWindowRecurrence) HasDayOfMonth() bool {
	if o != nil && !IsNil(o.DayOfMonth) {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given int32 and assigns it to the DayOfMonth field.
func (o *MaintenanceWindowRecurrence) SetDayOfMonth(v int32) {
	o.DayOfMonth = &v
}

// GetDuration returns the Duration field value
func (o *MaintenanceWindowRecurrence) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRecurrence) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *MaintenanceWindowRecurrence) SetDuration(v int32) {
	o.Duration = v
}

// GetStart returns the Start field value
func (o *MaintenanceWindowRecurrence) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRecurrence) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *MaintenanceWindowRecurrence) SetStart(v string) {
	o.Start = v
}

func (o MaintenanceWindowRecurrence) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaintenanceWindowRecurrence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Day) {
		toSerialize["day"] = o.Day
	}
	if !IsNil(o.DayOfMonth) {
		toSerialize["dayOfMonth"] = o.DayOfMonth
	}
	toSerialize["duration"] = o.Duration
	toSerialize["start"] = o.Start
	return toSerialize, nil
}

func (o *MaintenanceWindowRecurrence) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"duration",
		"start",
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

	varMaintenanceWindowRecurrence := _MaintenanceWindowRecurrence{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMaintenanceWindowRecurrence)

	if err != nil {
		return err
	}

	*o = MaintenanceWindowRecurrence(varMaintenanceWindowRecurrence)

	return err
}

type NullableMaintenanceWindowRecurrence struct {
	value *MaintenanceWindowRecurrence
	isSet bool
}

func (v NullableMaintenanceWindowRecurrence) Get() *MaintenanceWindowRecurrence {
	return v.value
}

func (v *NullableMaintenanceWindowRecurrence) Set(val *MaintenanceWindowRecurrence) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindowRecurrence) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindowRecurrence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindowRecurrence(val *MaintenanceWindowRecurrence) *NullableMaintenanceWindowRecurrence {
	return &NullableMaintenanceWindowRecurrence{value: val, isSet: true}
}

func (v NullableMaintenanceWindowRecurrence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindowRecurrence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
