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

// checks if the MonitoredEntityFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoredEntityFilter{}

// MonitoredEntityFilter Filters monitored entities by their type/tags.
type MonitoredEntityFilter struct {
	// The tag you want to use for matching.   You can use custom tags from the UI, AWS tags, Cloud Foundry tags, OpenShift/Kubernetes, and tags based on environment variables.
	Tags []UniversalTag `json:"tags,omitempty"`
	// The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching.
	Type *string `json:"type,omitempty"`
}

// NewMonitoredEntityFilter instantiates a new MonitoredEntityFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoredEntityFilter() *MonitoredEntityFilter {
	this := MonitoredEntityFilter{}
	return &this
}

// NewMonitoredEntityFilterWithDefaults instantiates a new MonitoredEntityFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoredEntityFilterWithDefaults() *MonitoredEntityFilter {
	this := MonitoredEntityFilter{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MonitoredEntityFilter) GetTags() []UniversalTag {
	if o == nil || IsNil(o.Tags) {
		var ret []UniversalTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoredEntityFilter) GetTagsOk() ([]UniversalTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MonitoredEntityFilter) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []UniversalTag and assigns it to the Tags field.
func (o *MonitoredEntityFilter) SetTags(v []UniversalTag) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitoredEntityFilter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoredEntityFilter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitoredEntityFilter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MonitoredEntityFilter) SetType(v string) {
	o.Type = &v
}

func (o MonitoredEntityFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoredEntityFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMonitoredEntityFilter struct {
	value *MonitoredEntityFilter
	isSet bool
}

func (v NullableMonitoredEntityFilter) Get() *MonitoredEntityFilter {
	return v.value
}

func (v *NullableMonitoredEntityFilter) Set(val *MonitoredEntityFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoredEntityFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoredEntityFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoredEntityFilter(val *MonitoredEntityFilter) *NullableMonitoredEntityFilter {
	return &NullableMonitoredEntityFilter{value: val, isSet: true}
}

func (v NullableMonitoredEntityFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoredEntityFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
