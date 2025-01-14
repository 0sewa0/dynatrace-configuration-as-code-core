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

// checks if the ProcessGroupInstanceModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessGroupInstanceModule{}

// ProcessGroupInstanceModule struct for ProcessGroupInstanceModule
type ProcessGroupInstanceModule struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewProcessGroupInstanceModule instantiates a new ProcessGroupInstanceModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupInstanceModule() *ProcessGroupInstanceModule {
	this := ProcessGroupInstanceModule{}
	return &this
}

// NewProcessGroupInstanceModuleWithDefaults instantiates a new ProcessGroupInstanceModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupInstanceModuleWithDefaults() *ProcessGroupInstanceModule {
	this := ProcessGroupInstanceModule{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProcessGroupInstanceModule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupInstanceModule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProcessGroupInstanceModule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProcessGroupInstanceModule) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ProcessGroupInstanceModule) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupInstanceModule) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ProcessGroupInstanceModule) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ProcessGroupInstanceModule) SetVersion(v string) {
	o.Version = &v
}

func (o ProcessGroupInstanceModule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessGroupInstanceModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableProcessGroupInstanceModule struct {
	value *ProcessGroupInstanceModule
	isSet bool
}

func (v NullableProcessGroupInstanceModule) Get() *ProcessGroupInstanceModule {
	return v.value
}

func (v *NullableProcessGroupInstanceModule) Set(val *ProcessGroupInstanceModule) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupInstanceModule) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupInstanceModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupInstanceModule(val *ProcessGroupInstanceModule) *NullableProcessGroupInstanceModule {
	return &NullableProcessGroupInstanceModule{value: val, isSet: true}
}

func (v NullableProcessGroupInstanceModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupInstanceModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
