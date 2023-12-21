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

// checks if the KeyPerformanceMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyPerformanceMetrics{}

// KeyPerformanceMetrics The key performance metrics configuration.
type KeyPerformanceMetrics struct {
	// Defines the key performance metric for load actions.
	LoadActionKpm string `json:"loadActionKpm"`
	// Defines the key performance metric for XHR actions.
	XhrActionKpm string `json:"xhrActionKpm"`
}

type _KeyPerformanceMetrics KeyPerformanceMetrics

// NewKeyPerformanceMetrics instantiates a new KeyPerformanceMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyPerformanceMetrics(loadActionKpm string, xhrActionKpm string) *KeyPerformanceMetrics {
	this := KeyPerformanceMetrics{}
	this.LoadActionKpm = loadActionKpm
	this.XhrActionKpm = xhrActionKpm
	return &this
}

// NewKeyPerformanceMetricsWithDefaults instantiates a new KeyPerformanceMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyPerformanceMetricsWithDefaults() *KeyPerformanceMetrics {
	this := KeyPerformanceMetrics{}
	return &this
}

// GetLoadActionKpm returns the LoadActionKpm field value
func (o *KeyPerformanceMetrics) GetLoadActionKpm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadActionKpm
}

// GetLoadActionKpmOk returns a tuple with the LoadActionKpm field value
// and a boolean to check if the value has been set.
func (o *KeyPerformanceMetrics) GetLoadActionKpmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadActionKpm, true
}

// SetLoadActionKpm sets field value
func (o *KeyPerformanceMetrics) SetLoadActionKpm(v string) {
	o.LoadActionKpm = v
}

// GetXhrActionKpm returns the XhrActionKpm field value
func (o *KeyPerformanceMetrics) GetXhrActionKpm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.XhrActionKpm
}

// GetXhrActionKpmOk returns a tuple with the XhrActionKpm field value
// and a boolean to check if the value has been set.
func (o *KeyPerformanceMetrics) GetXhrActionKpmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XhrActionKpm, true
}

// SetXhrActionKpm sets field value
func (o *KeyPerformanceMetrics) SetXhrActionKpm(v string) {
	o.XhrActionKpm = v
}

func (o KeyPerformanceMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyPerformanceMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["loadActionKpm"] = o.LoadActionKpm
	toSerialize["xhrActionKpm"] = o.XhrActionKpm
	return toSerialize, nil
}

func (o *KeyPerformanceMetrics) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"loadActionKpm",
		"xhrActionKpm",
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

	varKeyPerformanceMetrics := _KeyPerformanceMetrics{}

	err = json.Unmarshal(bytes, &varKeyPerformanceMetrics)

	if err != nil {
		return err
	}

	*o = KeyPerformanceMetrics(varKeyPerformanceMetrics)

	return err
}

type NullableKeyPerformanceMetrics struct {
	value *KeyPerformanceMetrics
	isSet bool
}

func (v NullableKeyPerformanceMetrics) Get() *KeyPerformanceMetrics {
	return v.value
}

func (v *NullableKeyPerformanceMetrics) Set(val *KeyPerformanceMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyPerformanceMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyPerformanceMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyPerformanceMetrics(val *KeyPerformanceMetrics) *NullableKeyPerformanceMetrics {
	return &NullableKeyPerformanceMetrics{value: val, isSet: true}
}

func (v NullableKeyPerformanceMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyPerformanceMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
