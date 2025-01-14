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

// checks if the SyntheticLocationUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyntheticLocationUpdate{}

// SyntheticLocationUpdate The update of a synthetic location. The actual object depends on the **type** of the location.
type SyntheticLocationUpdate struct {
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `PUBLIC` -> SyntheticPublicLocationUpdate  * `PRIVATE` -> PrivateSyntheticLocationUpdate
	Type string `json:"type"`
}

type _SyntheticLocationUpdate SyntheticLocationUpdate

// NewSyntheticLocationUpdate instantiates a new SyntheticLocationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticLocationUpdate(type_ string) *SyntheticLocationUpdate {
	this := SyntheticLocationUpdate{}
	this.Type = type_
	return &this
}

// NewSyntheticLocationUpdateWithDefaults instantiates a new SyntheticLocationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticLocationUpdateWithDefaults() *SyntheticLocationUpdate {
	this := SyntheticLocationUpdate{}
	return &this
}

// GetType returns the Type field value
func (o *SyntheticLocationUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticLocationUpdate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntheticLocationUpdate) SetType(v string) {
	o.Type = v
}

func (o SyntheticLocationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyntheticLocationUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *SyntheticLocationUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varSyntheticLocationUpdate := _SyntheticLocationUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSyntheticLocationUpdate)

	if err != nil {
		return err
	}

	*o = SyntheticLocationUpdate(varSyntheticLocationUpdate)

	return err
}

type NullableSyntheticLocationUpdate struct {
	value *SyntheticLocationUpdate
	isSet bool
}

func (v NullableSyntheticLocationUpdate) Get() *SyntheticLocationUpdate {
	return v.value
}

func (v *NullableSyntheticLocationUpdate) Set(val *SyntheticLocationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticLocationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticLocationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticLocationUpdate(val *SyntheticLocationUpdate) *NullableSyntheticLocationUpdate {
	return &NullableSyntheticLocationUpdate{value: val, isSet: true}
}

func (v NullableSyntheticLocationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticLocationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
