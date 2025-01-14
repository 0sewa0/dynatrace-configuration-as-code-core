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

// checks if the ThresholdRegistrationMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdRegistrationMessage{}

// ThresholdRegistrationMessage Parameters of a single plugin or custom event threshold.
type ThresholdRegistrationMessage struct {
	// The condition for the threshold value check: above or below.
	AlertCondition *string `json:"alertCondition,omitempty"`
	// How many one-minute samples within the evaluation window should go back to normal to close the event.
	DealertingSamples *int32 `json:"dealertingSamples,omitempty"`
	// A description of the event, triggered by a threshold violation.    You can use the following placeholders:  {severity}: the actual metric value,  {alert_condition}: above or below threshold condition,  {threshold}: the threshold value,{violating_samples}: the number of samples, violating the threshold,  {dimensions}: metric's dimension that violated the threshold.
	Description *string `json:"description,omitempty"`
	// The threshold is enabled/disabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the event, displayed in the UI.
	EventName *string `json:"eventName,omitempty"`
	// The type of the event to trigger on the threshold violation.
	EventType *string `json:"eventType,omitempty"`
	// The number of one-minute samples to form the sliding evaluation window.
	Samples *int32 `json:"samples,omitempty"`
	// The value of the threshold.
	Threshold *float64 `json:"threshold,omitempty"`
	// The ID of the threshold.
	ThresholdId *string `json:"thresholdId,omitempty"`
	// The case-sensitive ID of the metric evaluated by threshold.    You can find metric IDs, by performing the GET request to the `timeseries` endpoint of the API.
	TimeseriesId *string `json:"timeseriesId,omitempty"`
	// How many one-minute samples within the evaluation window should violate the threshold to trigger an event.
	ViolatingSamples *int32 `json:"violatingSamples,omitempty"`
}

// NewThresholdRegistrationMessage instantiates a new ThresholdRegistrationMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdRegistrationMessage() *ThresholdRegistrationMessage {
	this := ThresholdRegistrationMessage{}
	return &this
}

// NewThresholdRegistrationMessageWithDefaults instantiates a new ThresholdRegistrationMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdRegistrationMessageWithDefaults() *ThresholdRegistrationMessage {
	this := ThresholdRegistrationMessage{}
	return &this
}

// GetAlertCondition returns the AlertCondition field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetAlertCondition() string {
	if o == nil || IsNil(o.AlertCondition) {
		var ret string
		return ret
	}
	return *o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetAlertConditionOk() (*string, bool) {
	if o == nil || IsNil(o.AlertCondition) {
		return nil, false
	}
	return o.AlertCondition, true
}

// HasAlertCondition returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasAlertCondition() bool {
	if o != nil && !IsNil(o.AlertCondition) {
		return true
	}

	return false
}

// SetAlertCondition gets a reference to the given string and assigns it to the AlertCondition field.
func (o *ThresholdRegistrationMessage) SetAlertCondition(v string) {
	o.AlertCondition = &v
}

// GetDealertingSamples returns the DealertingSamples field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetDealertingSamples() int32 {
	if o == nil || IsNil(o.DealertingSamples) {
		var ret int32
		return ret
	}
	return *o.DealertingSamples
}

// GetDealertingSamplesOk returns a tuple with the DealertingSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetDealertingSamplesOk() (*int32, bool) {
	if o == nil || IsNil(o.DealertingSamples) {
		return nil, false
	}
	return o.DealertingSamples, true
}

// HasDealertingSamples returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasDealertingSamples() bool {
	if o != nil && !IsNil(o.DealertingSamples) {
		return true
	}

	return false
}

// SetDealertingSamples gets a reference to the given int32 and assigns it to the DealertingSamples field.
func (o *ThresholdRegistrationMessage) SetDealertingSamples(v int32) {
	o.DealertingSamples = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThresholdRegistrationMessage) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ThresholdRegistrationMessage) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *ThresholdRegistrationMessage) SetEventName(v string) {
	o.EventName = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *ThresholdRegistrationMessage) SetEventType(v string) {
	o.EventType = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetSamples() int32 {
	if o == nil || IsNil(o.Samples) {
		var ret int32
		return ret
	}
	return *o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetSamplesOk() (*int32, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given int32 and assigns it to the Samples field.
func (o *ThresholdRegistrationMessage) SetSamples(v int32) {
	o.Samples = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetThreshold() float64 {
	if o == nil || IsNil(o.Threshold) {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *ThresholdRegistrationMessage) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetThresholdId returns the ThresholdId field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetThresholdId() string {
	if o == nil || IsNil(o.ThresholdId) {
		var ret string
		return ret
	}
	return *o.ThresholdId
}

// GetThresholdIdOk returns a tuple with the ThresholdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetThresholdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThresholdId) {
		return nil, false
	}
	return o.ThresholdId, true
}

// HasThresholdId returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasThresholdId() bool {
	if o != nil && !IsNil(o.ThresholdId) {
		return true
	}

	return false
}

// SetThresholdId gets a reference to the given string and assigns it to the ThresholdId field.
func (o *ThresholdRegistrationMessage) SetThresholdId(v string) {
	o.ThresholdId = &v
}

// GetTimeseriesId returns the TimeseriesId field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetTimeseriesId() string {
	if o == nil || IsNil(o.TimeseriesId) {
		var ret string
		return ret
	}
	return *o.TimeseriesId
}

// GetTimeseriesIdOk returns a tuple with the TimeseriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetTimeseriesIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimeseriesId) {
		return nil, false
	}
	return o.TimeseriesId, true
}

// HasTimeseriesId returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasTimeseriesId() bool {
	if o != nil && !IsNil(o.TimeseriesId) {
		return true
	}

	return false
}

// SetTimeseriesId gets a reference to the given string and assigns it to the TimeseriesId field.
func (o *ThresholdRegistrationMessage) SetTimeseriesId(v string) {
	o.TimeseriesId = &v
}

// GetViolatingSamples returns the ViolatingSamples field value if set, zero value otherwise.
func (o *ThresholdRegistrationMessage) GetViolatingSamples() int32 {
	if o == nil || IsNil(o.ViolatingSamples) {
		var ret int32
		return ret
	}
	return *o.ViolatingSamples
}

// GetViolatingSamplesOk returns a tuple with the ViolatingSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdRegistrationMessage) GetViolatingSamplesOk() (*int32, bool) {
	if o == nil || IsNil(o.ViolatingSamples) {
		return nil, false
	}
	return o.ViolatingSamples, true
}

// HasViolatingSamples returns a boolean if a field has been set.
func (o *ThresholdRegistrationMessage) HasViolatingSamples() bool {
	if o != nil && !IsNil(o.ViolatingSamples) {
		return true
	}

	return false
}

// SetViolatingSamples gets a reference to the given int32 and assigns it to the ViolatingSamples field.
func (o *ThresholdRegistrationMessage) SetViolatingSamples(v int32) {
	o.ViolatingSamples = &v
}

func (o ThresholdRegistrationMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdRegistrationMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertCondition) {
		toSerialize["alertCondition"] = o.AlertCondition
	}
	if !IsNil(o.DealertingSamples) {
		toSerialize["dealertingSamples"] = o.DealertingSamples
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.ThresholdId) {
		toSerialize["thresholdId"] = o.ThresholdId
	}
	if !IsNil(o.TimeseriesId) {
		toSerialize["timeseriesId"] = o.TimeseriesId
	}
	if !IsNil(o.ViolatingSamples) {
		toSerialize["violatingSamples"] = o.ViolatingSamples
	}
	return toSerialize, nil
}

type NullableThresholdRegistrationMessage struct {
	value *ThresholdRegistrationMessage
	isSet bool
}

func (v NullableThresholdRegistrationMessage) Get() *ThresholdRegistrationMessage {
	return v.value
}

func (v *NullableThresholdRegistrationMessage) Set(val *ThresholdRegistrationMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdRegistrationMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdRegistrationMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdRegistrationMessage(val *ThresholdRegistrationMessage) *NullableThresholdRegistrationMessage {
	return &NullableThresholdRegistrationMessage{value: val, isSet: true}
}

func (v NullableThresholdRegistrationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdRegistrationMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
