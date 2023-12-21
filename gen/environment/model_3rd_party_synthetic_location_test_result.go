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

// checks if the Model3rdPartySyntheticLocationTestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model3rdPartySyntheticLocationTestResult{}

// Model3rdPartySyntheticLocationTestResult Results of third-party monitor executions per location.
type Model3rdPartySyntheticLocationTestResult struct {
	// The ID of the location.
	Id string `json:"id"`
	// The overall response time of the monitor from this location, in milliseconds.    If absent, it is calculated as the sum of response times of all steps.
	ResponseTimeMillis *int64 `json:"responseTimeMillis,omitempty"`
	// The timestamp of text execution start, in UTC milliseconds.
	StartTimestamp int64 `json:"startTimestamp"`
	// Results of individual monitor steps.
	StepResults []SyntheticMonitorStepResult `json:"stepResults"`
	// If the test was successful (`true`) or failed (`false`) - will influence availability timeseries.
	Success bool `json:"success"`
	// The overall availability of the monitor from this location, percent.    If absent, calculated as the number of successful steps compared to the overall number of steps.
	SuccessRate *float64 `json:"successRate,omitempty"`
}

type _Model3rdPartySyntheticLocationTestResult Model3rdPartySyntheticLocationTestResult

// NewModel3rdPartySyntheticLocationTestResult instantiates a new Model3rdPartySyntheticLocationTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel3rdPartySyntheticLocationTestResult(id string, startTimestamp int64, stepResults []SyntheticMonitorStepResult, success bool) *Model3rdPartySyntheticLocationTestResult {
	this := Model3rdPartySyntheticLocationTestResult{}
	this.Id = id
	this.StartTimestamp = startTimestamp
	this.StepResults = stepResults
	this.Success = success
	return &this
}

// NewModel3rdPartySyntheticLocationTestResultWithDefaults instantiates a new Model3rdPartySyntheticLocationTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel3rdPartySyntheticLocationTestResultWithDefaults() *Model3rdPartySyntheticLocationTestResult {
	this := Model3rdPartySyntheticLocationTestResult{}
	return &this
}

// GetId returns the Id field value
func (o *Model3rdPartySyntheticLocationTestResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticLocationTestResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Model3rdPartySyntheticLocationTestResult) SetId(v string) {
	o.Id = v
}

// GetResponseTimeMillis returns the ResponseTimeMillis field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticLocationTestResult) GetResponseTimeMillis() int64 {
	if o == nil || IsNil(o.ResponseTimeMillis) {
		var ret int64
		return ret
	}
	return *o.ResponseTimeMillis
}

// GetResponseTimeMillisOk returns a tuple with the ResponseTimeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticLocationTestResult) GetResponseTimeMillisOk() (*int64, bool) {
	if o == nil || IsNil(o.ResponseTimeMillis) {
		return nil, false
	}
	return o.ResponseTimeMillis, true
}

// HasResponseTimeMillis returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticLocationTestResult) HasResponseTimeMillis() bool {
	if o != nil && !IsNil(o.ResponseTimeMillis) {
		return true
	}

	return false
}

// SetResponseTimeMillis gets a reference to the given int64 and assigns it to the ResponseTimeMillis field.
func (o *Model3rdPartySyntheticLocationTestResult) SetResponseTimeMillis(v int64) {
	o.ResponseTimeMillis = &v
}

// GetStartTimestamp returns the StartTimestamp field value
func (o *Model3rdPartySyntheticLocationTestResult) GetStartTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticLocationTestResult) GetStartTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTimestamp, true
}

// SetStartTimestamp sets field value
func (o *Model3rdPartySyntheticLocationTestResult) SetStartTimestamp(v int64) {
	o.StartTimestamp = v
}

// GetStepResults returns the StepResults field value
func (o *Model3rdPartySyntheticLocationTestResult) GetStepResults() []SyntheticMonitorStepResult {
	if o == nil {
		var ret []SyntheticMonitorStepResult
		return ret
	}

	return o.StepResults
}

// GetStepResultsOk returns a tuple with the StepResults field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticLocationTestResult) GetStepResultsOk() ([]SyntheticMonitorStepResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.StepResults, true
}

// SetStepResults sets field value
func (o *Model3rdPartySyntheticLocationTestResult) SetStepResults(v []SyntheticMonitorStepResult) {
	o.StepResults = v
}

// GetSuccess returns the Success field value
func (o *Model3rdPartySyntheticLocationTestResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticLocationTestResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *Model3rdPartySyntheticLocationTestResult) SetSuccess(v bool) {
	o.Success = v
}

// GetSuccessRate returns the SuccessRate field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticLocationTestResult) GetSuccessRate() float64 {
	if o == nil || IsNil(o.SuccessRate) {
		var ret float64
		return ret
	}
	return *o.SuccessRate
}

// GetSuccessRateOk returns a tuple with the SuccessRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticLocationTestResult) GetSuccessRateOk() (*float64, bool) {
	if o == nil || IsNil(o.SuccessRate) {
		return nil, false
	}
	return o.SuccessRate, true
}

// HasSuccessRate returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticLocationTestResult) HasSuccessRate() bool {
	if o != nil && !IsNil(o.SuccessRate) {
		return true
	}

	return false
}

// SetSuccessRate gets a reference to the given float64 and assigns it to the SuccessRate field.
func (o *Model3rdPartySyntheticLocationTestResult) SetSuccessRate(v float64) {
	o.SuccessRate = &v
}

func (o Model3rdPartySyntheticLocationTestResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model3rdPartySyntheticLocationTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ResponseTimeMillis) {
		toSerialize["responseTimeMillis"] = o.ResponseTimeMillis
	}
	toSerialize["startTimestamp"] = o.StartTimestamp
	toSerialize["stepResults"] = o.StepResults
	toSerialize["success"] = o.Success
	if !IsNil(o.SuccessRate) {
		toSerialize["successRate"] = o.SuccessRate
	}
	return toSerialize, nil
}

func (o *Model3rdPartySyntheticLocationTestResult) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"startTimestamp",
		"stepResults",
		"success",
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

	varModel3rdPartySyntheticLocationTestResult := _Model3rdPartySyntheticLocationTestResult{}

	err = json.Unmarshal(bytes, &varModel3rdPartySyntheticLocationTestResult)

	if err != nil {
		return err
	}

	*o = Model3rdPartySyntheticLocationTestResult(varModel3rdPartySyntheticLocationTestResult)

	return err
}

type NullableModel3rdPartySyntheticLocationTestResult struct {
	value *Model3rdPartySyntheticLocationTestResult
	isSet bool
}

func (v NullableModel3rdPartySyntheticLocationTestResult) Get() *Model3rdPartySyntheticLocationTestResult {
	return v.value
}

func (v *NullableModel3rdPartySyntheticLocationTestResult) Set(val *Model3rdPartySyntheticLocationTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModel3rdPartySyntheticLocationTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModel3rdPartySyntheticLocationTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel3rdPartySyntheticLocationTestResult(val *Model3rdPartySyntheticLocationTestResult) *NullableModel3rdPartySyntheticLocationTestResult {
	return &NullableModel3rdPartySyntheticLocationTestResult{value: val, isSet: true}
}

func (v NullableModel3rdPartySyntheticLocationTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel3rdPartySyntheticLocationTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
