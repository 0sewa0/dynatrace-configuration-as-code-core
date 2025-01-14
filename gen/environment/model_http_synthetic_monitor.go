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

// checks if the HttpSyntheticMonitor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpSyntheticMonitor{}

// HttpSyntheticMonitor HTTP synthetic monitor. Some fields are inherited from base `SyntheticMonitor` model.
type HttpSyntheticMonitor struct {
	SyntheticMonitor
	// A list of events for this monitor
	Requests []RequestDto `json:"requests,omitempty"`
}

type _HttpSyntheticMonitor HttpSyntheticMonitor

// NewHttpSyntheticMonitor instantiates a new HttpSyntheticMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpSyntheticMonitor(automaticallyAssignedApps []string, createdFrom string, enabled bool, entityId string, frequencyMin int32, locations []string, managementZones []ManagementZone, manuallyAssignedApps []string, name string, script map[string]interface{}, tags []TagWithSourceInfo, type_ string) *HttpSyntheticMonitor {
	this := HttpSyntheticMonitor{}
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

// NewHttpSyntheticMonitorWithDefaults instantiates a new HttpSyntheticMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpSyntheticMonitorWithDefaults() *HttpSyntheticMonitor {
	this := HttpSyntheticMonitor{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *HttpSyntheticMonitor) GetRequests() []RequestDto {
	if o == nil || IsNil(o.Requests) {
		var ret []RequestDto
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpSyntheticMonitor) GetRequestsOk() ([]RequestDto, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *HttpSyntheticMonitor) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given []RequestDto and assigns it to the Requests field.
func (o *HttpSyntheticMonitor) SetRequests(v []RequestDto) {
	o.Requests = v
}

func (o HttpSyntheticMonitor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpSyntheticMonitor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSyntheticMonitor, errSyntheticMonitor := json.Marshal(o.SyntheticMonitor)
	if errSyntheticMonitor != nil {
		return map[string]interface{}{}, errSyntheticMonitor
	}
	errSyntheticMonitor = json.Unmarshal([]byte(serializedSyntheticMonitor), &toSerialize)
	if errSyntheticMonitor != nil {
		return map[string]interface{}{}, errSyntheticMonitor
	}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

func (o *HttpSyntheticMonitor) UnmarshalJSON(data []byte) (err error) {
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

	varHttpSyntheticMonitor := _HttpSyntheticMonitor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHttpSyntheticMonitor)

	if err != nil {
		return err
	}

	*o = HttpSyntheticMonitor(varHttpSyntheticMonitor)

	return err
}

type NullableHttpSyntheticMonitor struct {
	value *HttpSyntheticMonitor
	isSet bool
}

func (v NullableHttpSyntheticMonitor) Get() *HttpSyntheticMonitor {
	return v.value
}

func (v *NullableHttpSyntheticMonitor) Set(val *HttpSyntheticMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpSyntheticMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpSyntheticMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpSyntheticMonitor(val *HttpSyntheticMonitor) *NullableHttpSyntheticMonitor {
	return &NullableHttpSyntheticMonitor{value: val, isSet: true}
}

func (v NullableHttpSyntheticMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpSyntheticMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
