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

// checks if the Model3rdPartySyntheticEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model3rdPartySyntheticEvents{}

// Model3rdPartySyntheticEvents The list of third-party synthetic events.
type Model3rdPartySyntheticEvents struct {
	// The list of open third-party synthetic events.
	Open []Model3rdPartyEventOpenNotification `json:"open,omitempty"`
	// The list of closed third-party synthetic events.
	Resolved []Model3rdPartyEventResolvedNotification `json:"resolved,omitempty"`
	// The type of the third-party synthetic monitor.
	SyntheticEngineName string `json:"syntheticEngineName"`
}

type _Model3rdPartySyntheticEvents Model3rdPartySyntheticEvents

// NewModel3rdPartySyntheticEvents instantiates a new Model3rdPartySyntheticEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel3rdPartySyntheticEvents(syntheticEngineName string) *Model3rdPartySyntheticEvents {
	this := Model3rdPartySyntheticEvents{}
	this.SyntheticEngineName = syntheticEngineName
	return &this
}

// NewModel3rdPartySyntheticEventsWithDefaults instantiates a new Model3rdPartySyntheticEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel3rdPartySyntheticEventsWithDefaults() *Model3rdPartySyntheticEvents {
	this := Model3rdPartySyntheticEvents{}
	return &this
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticEvents) GetOpen() []Model3rdPartyEventOpenNotification {
	if o == nil || IsNil(o.Open) {
		var ret []Model3rdPartyEventOpenNotification
		return ret
	}
	return o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticEvents) GetOpenOk() ([]Model3rdPartyEventOpenNotification, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticEvents) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given []Model3rdPartyEventOpenNotification and assigns it to the Open field.
func (o *Model3rdPartySyntheticEvents) SetOpen(v []Model3rdPartyEventOpenNotification) {
	o.Open = v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticEvents) GetResolved() []Model3rdPartyEventResolvedNotification {
	if o == nil || IsNil(o.Resolved) {
		var ret []Model3rdPartyEventResolvedNotification
		return ret
	}
	return o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticEvents) GetResolvedOk() ([]Model3rdPartyEventResolvedNotification, bool) {
	if o == nil || IsNil(o.Resolved) {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticEvents) HasResolved() bool {
	if o != nil && !IsNil(o.Resolved) {
		return true
	}

	return false
}

// SetResolved gets a reference to the given []Model3rdPartyEventResolvedNotification and assigns it to the Resolved field.
func (o *Model3rdPartySyntheticEvents) SetResolved(v []Model3rdPartyEventResolvedNotification) {
	o.Resolved = v
}

// GetSyntheticEngineName returns the SyntheticEngineName field value
func (o *Model3rdPartySyntheticEvents) GetSyntheticEngineName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SyntheticEngineName
}

// GetSyntheticEngineNameOk returns a tuple with the SyntheticEngineName field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticEvents) GetSyntheticEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyntheticEngineName, true
}

// SetSyntheticEngineName sets field value
func (o *Model3rdPartySyntheticEvents) SetSyntheticEngineName(v string) {
	o.SyntheticEngineName = v
}

func (o Model3rdPartySyntheticEvents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model3rdPartySyntheticEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.Resolved) {
		toSerialize["resolved"] = o.Resolved
	}
	toSerialize["syntheticEngineName"] = o.SyntheticEngineName
	return toSerialize, nil
}

func (o *Model3rdPartySyntheticEvents) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"syntheticEngineName",
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

	varModel3rdPartySyntheticEvents := _Model3rdPartySyntheticEvents{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModel3rdPartySyntheticEvents)

	if err != nil {
		return err
	}

	*o = Model3rdPartySyntheticEvents(varModel3rdPartySyntheticEvents)

	return err
}

type NullableModel3rdPartySyntheticEvents struct {
	value *Model3rdPartySyntheticEvents
	isSet bool
}

func (v NullableModel3rdPartySyntheticEvents) Get() *Model3rdPartySyntheticEvents {
	return v.value
}

func (v *NullableModel3rdPartySyntheticEvents) Set(val *Model3rdPartySyntheticEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableModel3rdPartySyntheticEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableModel3rdPartySyntheticEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel3rdPartySyntheticEvents(val *Model3rdPartySyntheticEvents) *NullableModel3rdPartySyntheticEvents {
	return &NullableModel3rdPartySyntheticEvents{value: val, isSet: true}
}

func (v NullableModel3rdPartySyntheticEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel3rdPartySyntheticEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
