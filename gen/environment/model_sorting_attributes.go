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

// checks if the SortingAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SortingAttributes{}

// SortingAttributes How to sort the results.
type SortingAttributes struct {
	// Sort ascending (`true`) or descending (`false`).
	SortAscending *bool `json:"sortAscending,omitempty"`
	// The field to sort by.
	SortFieldName *string `json:"sortFieldName,omitempty"`
}

// NewSortingAttributes instantiates a new SortingAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortingAttributes() *SortingAttributes {
	this := SortingAttributes{}
	return &this
}

// NewSortingAttributesWithDefaults instantiates a new SortingAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortingAttributesWithDefaults() *SortingAttributes {
	this := SortingAttributes{}
	return &this
}

// GetSortAscending returns the SortAscending field value if set, zero value otherwise.
func (o *SortingAttributes) GetSortAscending() bool {
	if o == nil || IsNil(o.SortAscending) {
		var ret bool
		return ret
	}
	return *o.SortAscending
}

// GetSortAscendingOk returns a tuple with the SortAscending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortingAttributes) GetSortAscendingOk() (*bool, bool) {
	if o == nil || IsNil(o.SortAscending) {
		return nil, false
	}
	return o.SortAscending, true
}

// HasSortAscending returns a boolean if a field has been set.
func (o *SortingAttributes) HasSortAscending() bool {
	if o != nil && !IsNil(o.SortAscending) {
		return true
	}

	return false
}

// SetSortAscending gets a reference to the given bool and assigns it to the SortAscending field.
func (o *SortingAttributes) SetSortAscending(v bool) {
	o.SortAscending = &v
}

// GetSortFieldName returns the SortFieldName field value if set, zero value otherwise.
func (o *SortingAttributes) GetSortFieldName() string {
	if o == nil || IsNil(o.SortFieldName) {
		var ret string
		return ret
	}
	return *o.SortFieldName
}

// GetSortFieldNameOk returns a tuple with the SortFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortingAttributes) GetSortFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.SortFieldName) {
		return nil, false
	}
	return o.SortFieldName, true
}

// HasSortFieldName returns a boolean if a field has been set.
func (o *SortingAttributes) HasSortFieldName() bool {
	if o != nil && !IsNil(o.SortFieldName) {
		return true
	}

	return false
}

// SetSortFieldName gets a reference to the given string and assigns it to the SortFieldName field.
func (o *SortingAttributes) SetSortFieldName(v string) {
	o.SortFieldName = &v
}

func (o SortingAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SortingAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SortAscending) {
		toSerialize["sortAscending"] = o.SortAscending
	}
	if !IsNil(o.SortFieldName) {
		toSerialize["sortFieldName"] = o.SortFieldName
	}
	return toSerialize, nil
}

type NullableSortingAttributes struct {
	value *SortingAttributes
	isSet bool
}

func (v NullableSortingAttributes) Get() *SortingAttributes {
	return v.value
}

func (v *NullableSortingAttributes) Set(val *SortingAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSortingAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSortingAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortingAttributes(val *SortingAttributes) *NullableSortingAttributes {
	return &NullableSortingAttributes{value: val, isSet: true}
}

func (v NullableSortingAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortingAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
