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

// checks if the HostToRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostToRelationships{}

// HostToRelationships struct for HostToRelationships
type HostToRelationships struct {
	IsNetworkClientOfHost []string `json:"isNetworkClientOfHost,omitempty"`
	IsProcessOf           []string `json:"isProcessOf,omitempty"`
	IsSiteOf              []string `json:"isSiteOf,omitempty"`
	RunsOn                []string `json:"runsOn,omitempty"`
}

// NewHostToRelationships instantiates a new HostToRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostToRelationships() *HostToRelationships {
	this := HostToRelationships{}
	return &this
}

// NewHostToRelationshipsWithDefaults instantiates a new HostToRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostToRelationshipsWithDefaults() *HostToRelationships {
	this := HostToRelationships{}
	return &this
}

// GetIsNetworkClientOfHost returns the IsNetworkClientOfHost field value if set, zero value otherwise.
func (o *HostToRelationships) GetIsNetworkClientOfHost() []string {
	if o == nil || IsNil(o.IsNetworkClientOfHost) {
		var ret []string
		return ret
	}
	return o.IsNetworkClientOfHost
}

// GetIsNetworkClientOfHostOk returns a tuple with the IsNetworkClientOfHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostToRelationships) GetIsNetworkClientOfHostOk() ([]string, bool) {
	if o == nil || IsNil(o.IsNetworkClientOfHost) {
		return nil, false
	}
	return o.IsNetworkClientOfHost, true
}

// HasIsNetworkClientOfHost returns a boolean if a field has been set.
func (o *HostToRelationships) HasIsNetworkClientOfHost() bool {
	if o != nil && !IsNil(o.IsNetworkClientOfHost) {
		return true
	}

	return false
}

// SetIsNetworkClientOfHost gets a reference to the given []string and assigns it to the IsNetworkClientOfHost field.
func (o *HostToRelationships) SetIsNetworkClientOfHost(v []string) {
	o.IsNetworkClientOfHost = v
}

// GetIsProcessOf returns the IsProcessOf field value if set, zero value otherwise.
func (o *HostToRelationships) GetIsProcessOf() []string {
	if o == nil || IsNil(o.IsProcessOf) {
		var ret []string
		return ret
	}
	return o.IsProcessOf
}

// GetIsProcessOfOk returns a tuple with the IsProcessOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostToRelationships) GetIsProcessOfOk() ([]string, bool) {
	if o == nil || IsNil(o.IsProcessOf) {
		return nil, false
	}
	return o.IsProcessOf, true
}

// HasIsProcessOf returns a boolean if a field has been set.
func (o *HostToRelationships) HasIsProcessOf() bool {
	if o != nil && !IsNil(o.IsProcessOf) {
		return true
	}

	return false
}

// SetIsProcessOf gets a reference to the given []string and assigns it to the IsProcessOf field.
func (o *HostToRelationships) SetIsProcessOf(v []string) {
	o.IsProcessOf = v
}

// GetIsSiteOf returns the IsSiteOf field value if set, zero value otherwise.
func (o *HostToRelationships) GetIsSiteOf() []string {
	if o == nil || IsNil(o.IsSiteOf) {
		var ret []string
		return ret
	}
	return o.IsSiteOf
}

// GetIsSiteOfOk returns a tuple with the IsSiteOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostToRelationships) GetIsSiteOfOk() ([]string, bool) {
	if o == nil || IsNil(o.IsSiteOf) {
		return nil, false
	}
	return o.IsSiteOf, true
}

// HasIsSiteOf returns a boolean if a field has been set.
func (o *HostToRelationships) HasIsSiteOf() bool {
	if o != nil && !IsNil(o.IsSiteOf) {
		return true
	}

	return false
}

// SetIsSiteOf gets a reference to the given []string and assigns it to the IsSiteOf field.
func (o *HostToRelationships) SetIsSiteOf(v []string) {
	o.IsSiteOf = v
}

// GetRunsOn returns the RunsOn field value if set, zero value otherwise.
func (o *HostToRelationships) GetRunsOn() []string {
	if o == nil || IsNil(o.RunsOn) {
		var ret []string
		return ret
	}
	return o.RunsOn
}

// GetRunsOnOk returns a tuple with the RunsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostToRelationships) GetRunsOnOk() ([]string, bool) {
	if o == nil || IsNil(o.RunsOn) {
		return nil, false
	}
	return o.RunsOn, true
}

// HasRunsOn returns a boolean if a field has been set.
func (o *HostToRelationships) HasRunsOn() bool {
	if o != nil && !IsNil(o.RunsOn) {
		return true
	}

	return false
}

// SetRunsOn gets a reference to the given []string and assigns it to the RunsOn field.
func (o *HostToRelationships) SetRunsOn(v []string) {
	o.RunsOn = v
}

func (o HostToRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostToRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsNetworkClientOfHost) {
		toSerialize["isNetworkClientOfHost"] = o.IsNetworkClientOfHost
	}
	if !IsNil(o.IsProcessOf) {
		toSerialize["isProcessOf"] = o.IsProcessOf
	}
	if !IsNil(o.IsSiteOf) {
		toSerialize["isSiteOf"] = o.IsSiteOf
	}
	if !IsNil(o.RunsOn) {
		toSerialize["runsOn"] = o.RunsOn
	}
	return toSerialize, nil
}

type NullableHostToRelationships struct {
	value *HostToRelationships
	isSet bool
}

func (v NullableHostToRelationships) Get() *HostToRelationships {
	return v.value
}

func (v *NullableHostToRelationships) Set(val *HostToRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableHostToRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableHostToRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostToRelationships(val *HostToRelationships) *NullableHostToRelationships {
	return &NullableHostToRelationships{value: val, isSet: true}
}

func (v NullableHostToRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostToRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
