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

// checks if the TimeseriesDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeseriesDefinition{}

// TimeseriesDefinition The configuration of a metric with all its parameters.
type TimeseriesDefinition struct {
	// The list of allowed aggregations for this metric.
	AggregationTypes []string `json:"aggregationTypes,omitempty"`
	// The feature, where the metric originates.
	DetailedSource *string `json:"detailedSource,omitempty"`
	// The fine metric division, for example process group and process ID for some process-related metric.
	Dimensions []string `json:"dimensions,omitempty"`
	// The name of the metric in the user interface.
	DisplayName *string `json:"displayName,omitempty"`
	// The feature, where the metric originates.
	Filter *string `json:"filter,omitempty"`
	// The ID of the plugin, where the metric originates.
	PluginId *string `json:"pluginId,omitempty"`
	// The ID of the metric.
	TimeseriesId *string `json:"timeseriesId,omitempty"`
	// Technology type definition. Used to group metrics under a logical technology name.
	Types []string `json:"types,omitempty"`
	// The unit of the metric.
	Unit *string `json:"unit,omitempty"`
	// The warnings that occurred while creating the metric.
	Warnings []string `json:"warnings,omitempty"`
}

// NewTimeseriesDefinition instantiates a new TimeseriesDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesDefinition() *TimeseriesDefinition {
	this := TimeseriesDefinition{}
	return &this
}

// NewTimeseriesDefinitionWithDefaults instantiates a new TimeseriesDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesDefinitionWithDefaults() *TimeseriesDefinition {
	this := TimeseriesDefinition{}
	return &this
}

// GetAggregationTypes returns the AggregationTypes field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetAggregationTypes() []string {
	if o == nil || IsNil(o.AggregationTypes) {
		var ret []string
		return ret
	}
	return o.AggregationTypes
}

// GetAggregationTypesOk returns a tuple with the AggregationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetAggregationTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AggregationTypes) {
		return nil, false
	}
	return o.AggregationTypes, true
}

// HasAggregationTypes returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasAggregationTypes() bool {
	if o != nil && !IsNil(o.AggregationTypes) {
		return true
	}

	return false
}

// SetAggregationTypes gets a reference to the given []string and assigns it to the AggregationTypes field.
func (o *TimeseriesDefinition) SetAggregationTypes(v []string) {
	o.AggregationTypes = v
}

// GetDetailedSource returns the DetailedSource field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetDetailedSource() string {
	if o == nil || IsNil(o.DetailedSource) {
		var ret string
		return ret
	}
	return *o.DetailedSource
}

// GetDetailedSourceOk returns a tuple with the DetailedSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetDetailedSourceOk() (*string, bool) {
	if o == nil || IsNil(o.DetailedSource) {
		return nil, false
	}
	return o.DetailedSource, true
}

// HasDetailedSource returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasDetailedSource() bool {
	if o != nil && !IsNil(o.DetailedSource) {
		return true
	}

	return false
}

// SetDetailedSource gets a reference to the given string and assigns it to the DetailedSource field.
func (o *TimeseriesDefinition) SetDetailedSource(v string) {
	o.DetailedSource = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetDimensions() []string {
	if o == nil || IsNil(o.Dimensions) {
		var ret []string
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetDimensionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given []string and assigns it to the Dimensions field.
func (o *TimeseriesDefinition) SetDimensions(v []string) {
	o.Dimensions = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TimeseriesDefinition) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *TimeseriesDefinition) SetFilter(v string) {
	o.Filter = &v
}

// GetPluginId returns the PluginId field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetPluginId() string {
	if o == nil || IsNil(o.PluginId) {
		var ret string
		return ret
	}
	return *o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetPluginIdOk() (*string, bool) {
	if o == nil || IsNil(o.PluginId) {
		return nil, false
	}
	return o.PluginId, true
}

// HasPluginId returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasPluginId() bool {
	if o != nil && !IsNil(o.PluginId) {
		return true
	}

	return false
}

// SetPluginId gets a reference to the given string and assigns it to the PluginId field.
func (o *TimeseriesDefinition) SetPluginId(v string) {
	o.PluginId = &v
}

// GetTimeseriesId returns the TimeseriesId field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetTimeseriesId() string {
	if o == nil || IsNil(o.TimeseriesId) {
		var ret string
		return ret
	}
	return *o.TimeseriesId
}

// GetTimeseriesIdOk returns a tuple with the TimeseriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetTimeseriesIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimeseriesId) {
		return nil, false
	}
	return o.TimeseriesId, true
}

// HasTimeseriesId returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasTimeseriesId() bool {
	if o != nil && !IsNil(o.TimeseriesId) {
		return true
	}

	return false
}

// SetTimeseriesId gets a reference to the given string and assigns it to the TimeseriesId field.
func (o *TimeseriesDefinition) SetTimeseriesId(v string) {
	o.TimeseriesId = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *TimeseriesDefinition) SetTypes(v []string) {
	o.Types = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *TimeseriesDefinition) SetUnit(v string) {
	o.Unit = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *TimeseriesDefinition) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDefinition) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TimeseriesDefinition) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *TimeseriesDefinition) SetWarnings(v []string) {
	o.Warnings = v
}

func (o TimeseriesDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeseriesDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationTypes) {
		toSerialize["aggregationTypes"] = o.AggregationTypes
	}
	if !IsNil(o.DetailedSource) {
		toSerialize["detailedSource"] = o.DetailedSource
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.PluginId) {
		toSerialize["pluginId"] = o.PluginId
	}
	if !IsNil(o.TimeseriesId) {
		toSerialize["timeseriesId"] = o.TimeseriesId
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableTimeseriesDefinition struct {
	value *TimeseriesDefinition
	isSet bool
}

func (v NullableTimeseriesDefinition) Get() *TimeseriesDefinition {
	return v.value
}

func (v *NullableTimeseriesDefinition) Set(val *TimeseriesDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesDefinition(val *TimeseriesDefinition) *NullableTimeseriesDefinition {
	return &NullableTimeseriesDefinition{value: val, isSet: true}
}

func (v NullableTimeseriesDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
