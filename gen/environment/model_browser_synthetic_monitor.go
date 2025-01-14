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

// checks if the BrowserSyntheticMonitor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrowserSyntheticMonitor{}

// BrowserSyntheticMonitor Browser synthetic monitor. Some fields are inherited from the base `SyntheticMonitor` model.
type BrowserSyntheticMonitor struct {
	SyntheticMonitor
	// A list of events for this monitor
	Events                []EventDto             `json:"events,omitempty"`
	KeyPerformanceMetrics *KeyPerformanceMetrics `json:"keyPerformanceMetrics,omitempty"`
}

type _BrowserSyntheticMonitor BrowserSyntheticMonitor

// NewBrowserSyntheticMonitor instantiates a new BrowserSyntheticMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserSyntheticMonitor(automaticallyAssignedApps []string, createdFrom string, enabled bool, entityId string, frequencyMin int32, locations []string, managementZones []ManagementZone, manuallyAssignedApps []string, name string, script map[string]interface{}, tags []TagWithSourceInfo, type_ string) *BrowserSyntheticMonitor {
	this := BrowserSyntheticMonitor{}
	this.AutomaticallyAssignedApps = automaticallyAssignedApps
	this.CreatedFrom = createdFrom
	this.Enabled = enabled
	this.EntityId = entityId
	this.FrequencyMin = frequencyMin
	this.Locations = locations
	this.ManagementZones = managementZones
	this.ManuallyAssignedApps = manuallyAssignedApps
	this.Name = name
	this.Script = script
	this.Tags = tags
	this.Type = type_
	return &this
}

// NewBrowserSyntheticMonitorWithDefaults instantiates a new BrowserSyntheticMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserSyntheticMonitorWithDefaults() *BrowserSyntheticMonitor {
	this := BrowserSyntheticMonitor{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *BrowserSyntheticMonitor) GetEvents() []EventDto {
	if o == nil || IsNil(o.Events) {
		var ret []EventDto
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserSyntheticMonitor) GetEventsOk() ([]EventDto, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *BrowserSyntheticMonitor) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []EventDto and assigns it to the Events field.
func (o *BrowserSyntheticMonitor) SetEvents(v []EventDto) {
	o.Events = v
}

// GetKeyPerformanceMetrics returns the KeyPerformanceMetrics field value if set, zero value otherwise.
func (o *BrowserSyntheticMonitor) GetKeyPerformanceMetrics() KeyPerformanceMetrics {
	if o == nil || IsNil(o.KeyPerformanceMetrics) {
		var ret KeyPerformanceMetrics
		return ret
	}
	return *o.KeyPerformanceMetrics
}

// GetKeyPerformanceMetricsOk returns a tuple with the KeyPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserSyntheticMonitor) GetKeyPerformanceMetricsOk() (*KeyPerformanceMetrics, bool) {
	if o == nil || IsNil(o.KeyPerformanceMetrics) {
		return nil, false
	}
	return o.KeyPerformanceMetrics, true
}

// HasKeyPerformanceMetrics returns a boolean if a field has been set.
func (o *BrowserSyntheticMonitor) HasKeyPerformanceMetrics() bool {
	if o != nil && !IsNil(o.KeyPerformanceMetrics) {
		return true
	}

	return false
}

// SetKeyPerformanceMetrics gets a reference to the given KeyPerformanceMetrics and assigns it to the KeyPerformanceMetrics field.
func (o *BrowserSyntheticMonitor) SetKeyPerformanceMetrics(v KeyPerformanceMetrics) {
	o.KeyPerformanceMetrics = &v
}

func (o BrowserSyntheticMonitor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrowserSyntheticMonitor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSyntheticMonitor, errSyntheticMonitor := json.Marshal(o.SyntheticMonitor)
	if errSyntheticMonitor != nil {
		return map[string]interface{}{}, errSyntheticMonitor
	}
	errSyntheticMonitor = json.Unmarshal([]byte(serializedSyntheticMonitor), &toSerialize)
	if errSyntheticMonitor != nil {
		return map[string]interface{}{}, errSyntheticMonitor
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.KeyPerformanceMetrics) {
		toSerialize["keyPerformanceMetrics"] = o.KeyPerformanceMetrics
	}
	return toSerialize, nil
}

func (o *BrowserSyntheticMonitor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"automaticallyAssignedApps",
		"createdFrom",
		"enabled",
		"entityId",
		"frequencyMin",
		"locations",
		"managementZones",
		"manuallyAssignedApps",
		"name",
		"script",
		"tags",
		"type",
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

	varBrowserSyntheticMonitor := _BrowserSyntheticMonitor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBrowserSyntheticMonitor)

	if err != nil {
		return err
	}

	*o = BrowserSyntheticMonitor(varBrowserSyntheticMonitor)

	return err
}

type NullableBrowserSyntheticMonitor struct {
	value *BrowserSyntheticMonitor
	isSet bool
}

func (v NullableBrowserSyntheticMonitor) Get() *BrowserSyntheticMonitor {
	return v.value
}

func (v *NullableBrowserSyntheticMonitor) Set(val *BrowserSyntheticMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserSyntheticMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserSyntheticMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserSyntheticMonitor(val *BrowserSyntheticMonitor) *NullableBrowserSyntheticMonitor {
	return &NullableBrowserSyntheticMonitor{value: val, isSet: true}
}

func (v NullableBrowserSyntheticMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserSyntheticMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
