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

// checks if the Nodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Nodes{}

// Nodes A list of synthetic nodes
type Nodes struct {
	// A list of synthetic nodes
	Nodes []NodeCollectionElement `json:"nodes"`
}

type _Nodes Nodes

// NewNodes instantiates a new Nodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodes(nodes []NodeCollectionElement) *Nodes {
	this := Nodes{}
	this.Nodes = nodes
	return &this
}

// NewNodesWithDefaults instantiates a new Nodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodesWithDefaults() *Nodes {
	this := Nodes{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *Nodes) GetNodes() []NodeCollectionElement {
	if o == nil {
		var ret []NodeCollectionElement
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *Nodes) GetNodesOk() ([]NodeCollectionElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *Nodes) SetNodes(v []NodeCollectionElement) {
	o.Nodes = v
}

func (o Nodes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Nodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodes"] = o.Nodes
	return toSerialize, nil
}

func (o *Nodes) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nodes",
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

	varNodes := _Nodes{}

	err = json.Unmarshal(bytes, &varNodes)

	if err != nil {
		return err
	}

	*o = Nodes(varNodes)

	return err
}

type NullableNodes struct {
	value *Nodes
	isSet bool
}

func (v NullableNodes) Get() *Nodes {
	return v.value
}

func (v *NullableNodes) Set(val *Nodes) {
	v.value = val
	v.isSet = true
}

func (v NullableNodes) IsSet() bool {
	return v.isSet
}

func (v *NullableNodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodes(val *Nodes) *NullableNodes {
	return &NullableNodes{value: val, isSet: true}
}

func (v NullableNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
