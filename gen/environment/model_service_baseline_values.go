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

// checks if the ServiceBaselineValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceBaselineValues{}

// ServiceBaselineValues The baseline data for a service and its children for the **Response time** duration metric.
type ServiceBaselineValues struct {
	// The display name of the service.
	DisplayName *string `json:"displayName,omitempty"`
	// The ID of the service.
	EntityId string `json:"entityId"`
	// The baseline data for the **Response time** duration metric.
	ServiceResponseTimeBaselines []EntityBaselineData `json:"serviceResponseTimeBaselines,omitempty"`
}

type _ServiceBaselineValues ServiceBaselineValues

// NewServiceBaselineValues instantiates a new ServiceBaselineValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBaselineValues(entityId string) *ServiceBaselineValues {
	this := ServiceBaselineValues{}
	this.EntityId = entityId
	return &this
}

// NewServiceBaselineValuesWithDefaults instantiates a new ServiceBaselineValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBaselineValuesWithDefaults() *ServiceBaselineValues {
	this := ServiceBaselineValues{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ServiceBaselineValues) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBaselineValues) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ServiceBaselineValues) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ServiceBaselineValues) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEntityId returns the EntityId field value
func (o *ServiceBaselineValues) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *ServiceBaselineValues) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *ServiceBaselineValues) SetEntityId(v string) {
	o.EntityId = v
}

// GetServiceResponseTimeBaselines returns the ServiceResponseTimeBaselines field value if set, zero value otherwise.
func (o *ServiceBaselineValues) GetServiceResponseTimeBaselines() []EntityBaselineData {
	if o == nil || IsNil(o.ServiceResponseTimeBaselines) {
		var ret []EntityBaselineData
		return ret
	}
	return o.ServiceResponseTimeBaselines
}

// GetServiceResponseTimeBaselinesOk returns a tuple with the ServiceResponseTimeBaselines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBaselineValues) GetServiceResponseTimeBaselinesOk() ([]EntityBaselineData, bool) {
	if o == nil || IsNil(o.ServiceResponseTimeBaselines) {
		return nil, false
	}
	return o.ServiceResponseTimeBaselines, true
}

// HasServiceResponseTimeBaselines returns a boolean if a field has been set.
func (o *ServiceBaselineValues) HasServiceResponseTimeBaselines() bool {
	if o != nil && !IsNil(o.ServiceResponseTimeBaselines) {
		return true
	}

	return false
}

// SetServiceResponseTimeBaselines gets a reference to the given []EntityBaselineData and assigns it to the ServiceResponseTimeBaselines field.
func (o *ServiceBaselineValues) SetServiceResponseTimeBaselines(v []EntityBaselineData) {
	o.ServiceResponseTimeBaselines = v
}

func (o ServiceBaselineValues) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceBaselineValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["entityId"] = o.EntityId
	if !IsNil(o.ServiceResponseTimeBaselines) {
		toSerialize["serviceResponseTimeBaselines"] = o.ServiceResponseTimeBaselines
	}
	return toSerialize, nil
}

func (o *ServiceBaselineValues) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entityId",
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

	varServiceBaselineValues := _ServiceBaselineValues{}

	err = json.Unmarshal(bytes, &varServiceBaselineValues)

	if err != nil {
		return err
	}

	*o = ServiceBaselineValues(varServiceBaselineValues)

	return err
}

type NullableServiceBaselineValues struct {
	value *ServiceBaselineValues
	isSet bool
}

func (v NullableServiceBaselineValues) Get() *ServiceBaselineValues {
	return v.value
}

func (v *NullableServiceBaselineValues) Set(val *ServiceBaselineValues) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBaselineValues) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBaselineValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBaselineValues(val *ServiceBaselineValues) *NullableServiceBaselineValues {
	return &NullableServiceBaselineValues{value: val, isSet: true}
}

func (v NullableServiceBaselineValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBaselineValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
