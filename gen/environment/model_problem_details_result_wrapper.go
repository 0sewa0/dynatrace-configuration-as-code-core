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

// checks if the ProblemDetailsResultWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemDetailsResultWrapper{}

// ProblemDetailsResultWrapper struct for ProblemDetailsResultWrapper
type ProblemDetailsResultWrapper struct {
	Result *Problem `json:"result,omitempty"`
}

// NewProblemDetailsResultWrapper instantiates a new ProblemDetailsResultWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetailsResultWrapper() *ProblemDetailsResultWrapper {
	this := ProblemDetailsResultWrapper{}
	return &this
}

// NewProblemDetailsResultWrapperWithDefaults instantiates a new ProblemDetailsResultWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetailsResultWrapperWithDefaults() *ProblemDetailsResultWrapper {
	this := ProblemDetailsResultWrapper{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ProblemDetailsResultWrapper) GetResult() Problem {
	if o == nil || IsNil(o.Result) {
		var ret Problem
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsResultWrapper) GetResultOk() (*Problem, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ProblemDetailsResultWrapper) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Problem and assigns it to the Result field.
func (o *ProblemDetailsResultWrapper) SetResult(v Problem) {
	o.Result = &v
}

func (o ProblemDetailsResultWrapper) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemDetailsResultWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableProblemDetailsResultWrapper struct {
	value *ProblemDetailsResultWrapper
	isSet bool
}

func (v NullableProblemDetailsResultWrapper) Get() *ProblemDetailsResultWrapper {
	return v.value
}

func (v *NullableProblemDetailsResultWrapper) Set(val *ProblemDetailsResultWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetailsResultWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetailsResultWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetailsResultWrapper(val *ProblemDetailsResultWrapper) *NullableProblemDetailsResultWrapper {
	return &NullableProblemDetailsResultWrapper{value: val, isSet: true}
}

func (v NullableProblemDetailsResultWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetailsResultWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
