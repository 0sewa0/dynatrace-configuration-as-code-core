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

// checks if the EntityBaselineData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityBaselineData{}

// EntityBaselineData The baseline data for a Dynatrace entity for a specific duration metric.
type EntityBaselineData struct {
	// The baseline data for child entities of this entity, for example a `SERVICE_METHOD` of a `SERVICE_METHOD_GROUP`.
	ChildBaselines []EntityBaselineData `json:"childBaselines,omitempty"`
	// The display name of the entity.
	DisplayName *string `json:"displayName,omitempty"`
	// The ID of the Dynatrace entity.
	EntityId string `json:"entityId"`
	// The error rate baseline.
	ErrorRate *float32 `json:"errorRate,omitempty"`
	// The entity has a load baseline (`true`) or doesn't (`false`).
	HasLoadBaseline *bool `json:"hasLoadBaseline,omitempty"`
	// The 90th percentile baseline, in microseconds.
	Micros90thPercentile *float32 `json:"micros90thPercentile,omitempty"`
	// The median baseline, in microseconds.
	MicrosMedian *float32 `json:"microsMedian,omitempty"`
}

type _EntityBaselineData EntityBaselineData

// NewEntityBaselineData instantiates a new EntityBaselineData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityBaselineData(entityId string) *EntityBaselineData {
	this := EntityBaselineData{}
	this.EntityId = entityId
	return &this
}

// NewEntityBaselineDataWithDefaults instantiates a new EntityBaselineData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityBaselineDataWithDefaults() *EntityBaselineData {
	this := EntityBaselineData{}
	return &this
}

// GetChildBaselines returns the ChildBaselines field value if set, zero value otherwise.
func (o *EntityBaselineData) GetChildBaselines() []EntityBaselineData {
	if o == nil || IsNil(o.ChildBaselines) {
		var ret []EntityBaselineData
		return ret
	}
	return o.ChildBaselines
}

// GetChildBaselinesOk returns a tuple with the ChildBaselines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityBaselineData) GetChildBaselinesOk() ([]EntityBaselineData, bool) {
	if o == nil || IsNil(o.ChildBaselines) {
		return nil, false
	}
	return o.ChildBaselines, true
}

// HasChildBaselines returns a boolean if a field has been set.
func (o *EntityBaselineData) HasChildBaselines() bool {
	if o != nil && !IsNil(o.ChildBaselines) {
		return true
	}

	return false
}

// SetChildBaselines gets a reference to the given []EntityBaselineData and assigns it to the ChildBaselines field.
func (o *EntityBaselineData) SetChildBaselines(v []EntityBaselineData) {
	o.ChildBaselines = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EntityBaselineData) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityBaselineData) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EntityBaselineData) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EntityBaselineData) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEntityId returns the EntityId field value
func (o *EntityBaselineData) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *EntityBaselineData) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *EntityBaselineData) SetEntityId(v string) {
	o.EntityId = v
}

// GetErrorRate returns the ErrorRate field value if set, zero value otherwise.
func (o *EntityBaselineData) GetErrorRate() float32 {
	if o == nil || IsNil(o.ErrorRate) {
		var ret float32
		return ret
	}
	return *o.ErrorRate
}

// GetErrorRateOk returns a tuple with the ErrorRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityBaselineData) GetErrorRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ErrorRate) {
		return nil, false
	}
	return o.ErrorRate, true
}

// HasErrorRate returns a boolean if a field has been set.
func (o *EntityBaselineData) HasErrorRate() bool {
	if o != nil && !IsNil(o.ErrorRate) {
		return true
	}

	return false
}

// SetErrorRate gets a reference to the given float32 and assigns it to the ErrorRate field.
func (o *EntityBaselineData) SetErrorRate(v float32) {
	o.ErrorRate = &v
}

// GetHasLoadBaseline returns the HasLoadBaseline field value if set, zero value otherwise.
func (o *EntityBaselineData) GetHasLoadBaseline() bool {
	if o == nil || IsNil(o.HasLoadBaseline) {
		var ret bool
		return ret
	}
	return *o.HasLoadBaseline
}

// GetHasLoadBaselineOk returns a tuple with the HasLoadBaseline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityBaselineData) GetHasLoadBaselineOk() (*bool, bool) {
	if o == nil || IsNil(o.HasLoadBaseline) {
		return nil, false
	}
	return o.HasLoadBaseline, true
}

// HasHasLoadBaseline returns a boolean if a field has been set.
func (o *EntityBaselineData) HasHasLoadBaseline() bool {
	if o != nil && !IsNil(o.HasLoadBaseline) {
		return true
	}

	return false
}

// SetHasLoadBaseline gets a reference to the given bool and assigns it to the HasLoadBaseline field.
func (o *EntityBaselineData) SetHasLoadBaseline(v bool) {
	o.HasLoadBaseline = &v
}

// GetMicros90thPercentile returns the Micros90thPercentile field value if set, zero value otherwise.
func (o *EntityBaselineData) GetMicros90thPercentile() float32 {
	if o == nil || IsNil(o.Micros90thPercentile) {
		var ret float32
		return ret
	}
	return *o.Micros90thPercentile
}

// GetMicros90thPercentileOk returns a tuple with the Micros90thPercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityBaselineData) GetMicros90thPercentileOk() (*float32, bool) {
	if o == nil || IsNil(o.Micros90thPercentile) {
		return nil, false
	}
	return o.Micros90thPercentile, true
}

// HasMicros90thPercentile returns a boolean if a field has been set.
func (o *EntityBaselineData) HasMicros90thPercentile() bool {
	if o != nil && !IsNil(o.Micros90thPercentile) {
		return true
	}

	return false
}

// SetMicros90thPercentile gets a reference to the given float32 and assigns it to the Micros90thPercentile field.
func (o *EntityBaselineData) SetMicros90thPercentile(v float32) {
	o.Micros90thPercentile = &v
}

// GetMicrosMedian returns the MicrosMedian field value if set, zero value otherwise.
func (o *EntityBaselineData) GetMicrosMedian() float32 {
	if o == nil || IsNil(o.MicrosMedian) {
		var ret float32
		return ret
	}
	return *o.MicrosMedian
}

// GetMicrosMedianOk returns a tuple with the MicrosMedian field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityBaselineData) GetMicrosMedianOk() (*float32, bool) {
	if o == nil || IsNil(o.MicrosMedian) {
		return nil, false
	}
	return o.MicrosMedian, true
}

// HasMicrosMedian returns a boolean if a field has been set.
func (o *EntityBaselineData) HasMicrosMedian() bool {
	if o != nil && !IsNil(o.MicrosMedian) {
		return true
	}

	return false
}

// SetMicrosMedian gets a reference to the given float32 and assigns it to the MicrosMedian field.
func (o *EntityBaselineData) SetMicrosMedian(v float32) {
	o.MicrosMedian = &v
}

func (o EntityBaselineData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityBaselineData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChildBaselines) {
		toSerialize["childBaselines"] = o.ChildBaselines
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["entityId"] = o.EntityId
	if !IsNil(o.ErrorRate) {
		toSerialize["errorRate"] = o.ErrorRate
	}
	if !IsNil(o.HasLoadBaseline) {
		toSerialize["hasLoadBaseline"] = o.HasLoadBaseline
	}
	if !IsNil(o.Micros90thPercentile) {
		toSerialize["micros90thPercentile"] = o.Micros90thPercentile
	}
	if !IsNil(o.MicrosMedian) {
		toSerialize["microsMedian"] = o.MicrosMedian
	}
	return toSerialize, nil
}

func (o *EntityBaselineData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entityId",
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

	varEntityBaselineData := _EntityBaselineData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntityBaselineData)

	if err != nil {
		return err
	}

	*o = EntityBaselineData(varEntityBaselineData)

	return err
}

type NullableEntityBaselineData struct {
	value *EntityBaselineData
	isSet bool
}

func (v NullableEntityBaselineData) Get() *EntityBaselineData {
	return v.value
}

func (v *NullableEntityBaselineData) Set(val *EntityBaselineData) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityBaselineData) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityBaselineData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityBaselineData(val *EntityBaselineData) *NullableEntityBaselineData {
	return &NullableEntityBaselineData{value: val, isSet: true}
}

func (v NullableEntityBaselineData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityBaselineData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
