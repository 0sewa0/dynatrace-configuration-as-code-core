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

// checks if the Model3rdPartySyntheticTests type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model3rdPartySyntheticTests{}

// Model3rdPartySyntheticTests struct for Model3rdPartySyntheticTests
type Model3rdPartySyntheticTests struct {
	// The list of third-party synthetic locations.
	Locations []Model3rdPartySyntheticLocation `json:"locations"`
	// The timestamp of the message creation, in UTC milliseconds.
	MessageTimestamp int64 `json:"messageTimestamp"`
	// The URL of the third-party synthetic monitor icon.
	SyntheticEngineIconUrl *string `json:"syntheticEngineIconUrl,omitempty"`
	// The type of the third-party synthetic monitor.
	SyntheticEngineName string `json:"syntheticEngineName"`
	// The list of results of third-party synthetic monitor execution.
	TestResults []Model3rdPartySyntheticTestResult `json:"testResults,omitempty"`
	// The list of third-party synthetic monitors.
	Tests []Model3rdPartySyntheticMonitor `json:"tests"`
}

type _Model3rdPartySyntheticTests Model3rdPartySyntheticTests

// NewModel3rdPartySyntheticTests instantiates a new Model3rdPartySyntheticTests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel3rdPartySyntheticTests(locations []Model3rdPartySyntheticLocation, messageTimestamp int64, syntheticEngineName string, tests []Model3rdPartySyntheticMonitor) *Model3rdPartySyntheticTests {
	this := Model3rdPartySyntheticTests{}
	this.Locations = locations
	this.MessageTimestamp = messageTimestamp
	this.SyntheticEngineName = syntheticEngineName
	this.Tests = tests
	return &this
}

// NewModel3rdPartySyntheticTestsWithDefaults instantiates a new Model3rdPartySyntheticTests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel3rdPartySyntheticTestsWithDefaults() *Model3rdPartySyntheticTests {
	this := Model3rdPartySyntheticTests{}
	return &this
}

// GetLocations returns the Locations field value
func (o *Model3rdPartySyntheticTests) GetLocations() []Model3rdPartySyntheticLocation {
	if o == nil {
		var ret []Model3rdPartySyntheticLocation
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTests) GetLocationsOk() ([]Model3rdPartySyntheticLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locations, true
}

// SetLocations sets field value
func (o *Model3rdPartySyntheticTests) SetLocations(v []Model3rdPartySyntheticLocation) {
	o.Locations = v
}

// GetMessageTimestamp returns the MessageTimestamp field value
func (o *Model3rdPartySyntheticTests) GetMessageTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MessageTimestamp
}

// GetMessageTimestampOk returns a tuple with the MessageTimestamp field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTests) GetMessageTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageTimestamp, true
}

// SetMessageTimestamp sets field value
func (o *Model3rdPartySyntheticTests) SetMessageTimestamp(v int64) {
	o.MessageTimestamp = v
}

// GetSyntheticEngineIconUrl returns the SyntheticEngineIconUrl field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticTests) GetSyntheticEngineIconUrl() string {
	if o == nil || IsNil(o.SyntheticEngineIconUrl) {
		var ret string
		return ret
	}
	return *o.SyntheticEngineIconUrl
}

// GetSyntheticEngineIconUrlOk returns a tuple with the SyntheticEngineIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTests) GetSyntheticEngineIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SyntheticEngineIconUrl) {
		return nil, false
	}
	return o.SyntheticEngineIconUrl, true
}

// HasSyntheticEngineIconUrl returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticTests) HasSyntheticEngineIconUrl() bool {
	if o != nil && !IsNil(o.SyntheticEngineIconUrl) {
		return true
	}

	return false
}

// SetSyntheticEngineIconUrl gets a reference to the given string and assigns it to the SyntheticEngineIconUrl field.
func (o *Model3rdPartySyntheticTests) SetSyntheticEngineIconUrl(v string) {
	o.SyntheticEngineIconUrl = &v
}

// GetSyntheticEngineName returns the SyntheticEngineName field value
func (o *Model3rdPartySyntheticTests) GetSyntheticEngineName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SyntheticEngineName
}

// GetSyntheticEngineNameOk returns a tuple with the SyntheticEngineName field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTests) GetSyntheticEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyntheticEngineName, true
}

// SetSyntheticEngineName sets field value
func (o *Model3rdPartySyntheticTests) SetSyntheticEngineName(v string) {
	o.SyntheticEngineName = v
}

// GetTestResults returns the TestResults field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticTests) GetTestResults() []Model3rdPartySyntheticTestResult {
	if o == nil || IsNil(o.TestResults) {
		var ret []Model3rdPartySyntheticTestResult
		return ret
	}
	return o.TestResults
}

// GetTestResultsOk returns a tuple with the TestResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTests) GetTestResultsOk() ([]Model3rdPartySyntheticTestResult, bool) {
	if o == nil || IsNil(o.TestResults) {
		return nil, false
	}
	return o.TestResults, true
}

// HasTestResults returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticTests) HasTestResults() bool {
	if o != nil && !IsNil(o.TestResults) {
		return true
	}

	return false
}

// SetTestResults gets a reference to the given []Model3rdPartySyntheticTestResult and assigns it to the TestResults field.
func (o *Model3rdPartySyntheticTests) SetTestResults(v []Model3rdPartySyntheticTestResult) {
	o.TestResults = v
}

// GetTests returns the Tests field value
func (o *Model3rdPartySyntheticTests) GetTests() []Model3rdPartySyntheticMonitor {
	if o == nil {
		var ret []Model3rdPartySyntheticMonitor
		return ret
	}

	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticTests) GetTestsOk() ([]Model3rdPartySyntheticMonitor, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tests, true
}

// SetTests sets field value
func (o *Model3rdPartySyntheticTests) SetTests(v []Model3rdPartySyntheticMonitor) {
	o.Tests = v
}

func (o Model3rdPartySyntheticTests) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model3rdPartySyntheticTests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locations"] = o.Locations
	toSerialize["messageTimestamp"] = o.MessageTimestamp
	if !IsNil(o.SyntheticEngineIconUrl) {
		toSerialize["syntheticEngineIconUrl"] = o.SyntheticEngineIconUrl
	}
	toSerialize["syntheticEngineName"] = o.SyntheticEngineName
	if !IsNil(o.TestResults) {
		toSerialize["testResults"] = o.TestResults
	}
	toSerialize["tests"] = o.Tests
	return toSerialize, nil
}

func (o *Model3rdPartySyntheticTests) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"locations",
		"messageTimestamp",
		"syntheticEngineName",
		"tests",
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

	varModel3rdPartySyntheticTests := _Model3rdPartySyntheticTests{}

	err = json.Unmarshal(bytes, &varModel3rdPartySyntheticTests)

	if err != nil {
		return err
	}

	*o = Model3rdPartySyntheticTests(varModel3rdPartySyntheticTests)

	return err
}

type NullableModel3rdPartySyntheticTests struct {
	value *Model3rdPartySyntheticTests
	isSet bool
}

func (v NullableModel3rdPartySyntheticTests) Get() *Model3rdPartySyntheticTests {
	return v.value
}

func (v *NullableModel3rdPartySyntheticTests) Set(val *Model3rdPartySyntheticTests) {
	v.value = val
	v.isSet = true
}

func (v NullableModel3rdPartySyntheticTests) IsSet() bool {
	return v.isSet
}

func (v *NullableModel3rdPartySyntheticTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel3rdPartySyntheticTests(val *Model3rdPartySyntheticTests) *NullableModel3rdPartySyntheticTests {
	return &NullableModel3rdPartySyntheticTests{value: val, isSet: true}
}

func (v NullableModel3rdPartySyntheticTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel3rdPartySyntheticTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
