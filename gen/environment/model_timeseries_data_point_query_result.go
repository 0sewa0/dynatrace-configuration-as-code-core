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

// checks if the TimeseriesDataPointQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeseriesDataPointQueryResult{}

// TimeseriesDataPointQueryResult List of metric's datapoints, as well as their parameters.
type TimeseriesDataPointQueryResult struct {
	// The type of data points aggregation.
	AggregationType *string `json:"aggregationType,omitempty"`
	// Data points of a metric.   A JSON object that maps the ID of the entity that delivered the data points and an array, which consists of arrays of the data point float values.   May contain more than one entity ID per record (for example, a host and its network interface). In such cases, entity IDs are separated by commas.   A datapoint contains a value and a timestamp, at which the value was recorded.    Dynatrace stores data in time slots. The **dataPoints** object shows the *starting* timestamp of the slot. If the **startTimestamp** or **endTimestamp** of your query lies inside of the data time slot, this time slot is included into response. Due to the fact that the timestamp of the first data point lies outside of the specified timeframe, you will see *earlier* timestamp than the specified **startTimestamp** in the first data point of the response.   There are three versions of data points:   * Numeric datapoint: Contains a numeric value.   * Enum datapoint: Contains an enum value, currently only for availability timeseries.   * Prediction datapoint: Similar to the numeric datapoint, but it contains a confidence interval, within which the future values are expected to be.
	DataPoints *map[string][][]float32 `json:"dataPoints,omitempty"`
	// The list of entities where the data points originate.  A JSON object that maps the entity ID in Dynatrace and the actual name of the entity.
	Entities *map[string]string `json:"entities,omitempty"`
	// The resolution of data points.
	ResolutionInMillisUTC *int64 `json:"resolutionInMillisUTC,omitempty"`
	// The ID of the metric.
	TimeseriesId *string `json:"timeseriesId,omitempty"`
	// The unit of data points.
	Unit *string `json:"unit,omitempty"`
}

// NewTimeseriesDataPointQueryResult instantiates a new TimeseriesDataPointQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesDataPointQueryResult() *TimeseriesDataPointQueryResult {
	this := TimeseriesDataPointQueryResult{}
	return &this
}

// NewTimeseriesDataPointQueryResultWithDefaults instantiates a new TimeseriesDataPointQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesDataPointQueryResultWithDefaults() *TimeseriesDataPointQueryResult {
	this := TimeseriesDataPointQueryResult{}
	return &this
}

// GetAggregationType returns the AggregationType field value if set, zero value otherwise.
func (o *TimeseriesDataPointQueryResult) GetAggregationType() string {
	if o == nil || IsNil(o.AggregationType) {
		var ret string
		return ret
	}
	return *o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDataPointQueryResult) GetAggregationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationType) {
		return nil, false
	}
	return o.AggregationType, true
}

// HasAggregationType returns a boolean if a field has been set.
func (o *TimeseriesDataPointQueryResult) HasAggregationType() bool {
	if o != nil && !IsNil(o.AggregationType) {
		return true
	}

	return false
}

// SetAggregationType gets a reference to the given string and assigns it to the AggregationType field.
func (o *TimeseriesDataPointQueryResult) SetAggregationType(v string) {
	o.AggregationType = &v
}

// GetDataPoints returns the DataPoints field value if set, zero value otherwise.
func (o *TimeseriesDataPointQueryResult) GetDataPoints() map[string][][]float32 {
	if o == nil || IsNil(o.DataPoints) {
		var ret map[string][][]float32
		return ret
	}
	return *o.DataPoints
}

// GetDataPointsOk returns a tuple with the DataPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDataPointQueryResult) GetDataPointsOk() (*map[string][][]float32, bool) {
	if o == nil || IsNil(o.DataPoints) {
		return nil, false
	}
	return o.DataPoints, true
}

// HasDataPoints returns a boolean if a field has been set.
func (o *TimeseriesDataPointQueryResult) HasDataPoints() bool {
	if o != nil && !IsNil(o.DataPoints) {
		return true
	}

	return false
}

// SetDataPoints gets a reference to the given map[string][][]float32 and assigns it to the DataPoints field.
func (o *TimeseriesDataPointQueryResult) SetDataPoints(v map[string][][]float32) {
	o.DataPoints = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *TimeseriesDataPointQueryResult) GetEntities() map[string]string {
	if o == nil || IsNil(o.Entities) {
		var ret map[string]string
		return ret
	}
	return *o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDataPointQueryResult) GetEntitiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *TimeseriesDataPointQueryResult) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given map[string]string and assigns it to the Entities field.
func (o *TimeseriesDataPointQueryResult) SetEntities(v map[string]string) {
	o.Entities = &v
}

// GetResolutionInMillisUTC returns the ResolutionInMillisUTC field value if set, zero value otherwise.
func (o *TimeseriesDataPointQueryResult) GetResolutionInMillisUTC() int64 {
	if o == nil || IsNil(o.ResolutionInMillisUTC) {
		var ret int64
		return ret
	}
	return *o.ResolutionInMillisUTC
}

// GetResolutionInMillisUTCOk returns a tuple with the ResolutionInMillisUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDataPointQueryResult) GetResolutionInMillisUTCOk() (*int64, bool) {
	if o == nil || IsNil(o.ResolutionInMillisUTC) {
		return nil, false
	}
	return o.ResolutionInMillisUTC, true
}

// HasResolutionInMillisUTC returns a boolean if a field has been set.
func (o *TimeseriesDataPointQueryResult) HasResolutionInMillisUTC() bool {
	if o != nil && !IsNil(o.ResolutionInMillisUTC) {
		return true
	}

	return false
}

// SetResolutionInMillisUTC gets a reference to the given int64 and assigns it to the ResolutionInMillisUTC field.
func (o *TimeseriesDataPointQueryResult) SetResolutionInMillisUTC(v int64) {
	o.ResolutionInMillisUTC = &v
}

// GetTimeseriesId returns the TimeseriesId field value if set, zero value otherwise.
func (o *TimeseriesDataPointQueryResult) GetTimeseriesId() string {
	if o == nil || IsNil(o.TimeseriesId) {
		var ret string
		return ret
	}
	return *o.TimeseriesId
}

// GetTimeseriesIdOk returns a tuple with the TimeseriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDataPointQueryResult) GetTimeseriesIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimeseriesId) {
		return nil, false
	}
	return o.TimeseriesId, true
}

// HasTimeseriesId returns a boolean if a field has been set.
func (o *TimeseriesDataPointQueryResult) HasTimeseriesId() bool {
	if o != nil && !IsNil(o.TimeseriesId) {
		return true
	}

	return false
}

// SetTimeseriesId gets a reference to the given string and assigns it to the TimeseriesId field.
func (o *TimeseriesDataPointQueryResult) SetTimeseriesId(v string) {
	o.TimeseriesId = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *TimeseriesDataPointQueryResult) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesDataPointQueryResult) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *TimeseriesDataPointQueryResult) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *TimeseriesDataPointQueryResult) SetUnit(v string) {
	o.Unit = &v
}

func (o TimeseriesDataPointQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeseriesDataPointQueryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationType) {
		toSerialize["aggregationType"] = o.AggregationType
	}
	if !IsNil(o.DataPoints) {
		toSerialize["dataPoints"] = o.DataPoints
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.ResolutionInMillisUTC) {
		toSerialize["resolutionInMillisUTC"] = o.ResolutionInMillisUTC
	}
	if !IsNil(o.TimeseriesId) {
		toSerialize["timeseriesId"] = o.TimeseriesId
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableTimeseriesDataPointQueryResult struct {
	value *TimeseriesDataPointQueryResult
	isSet bool
}

func (v NullableTimeseriesDataPointQueryResult) Get() *TimeseriesDataPointQueryResult {
	return v.value
}

func (v *NullableTimeseriesDataPointQueryResult) Set(val *TimeseriesDataPointQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesDataPointQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesDataPointQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesDataPointQueryResult(val *TimeseriesDataPointQueryResult) *NullableTimeseriesDataPointQueryResult {
	return &NullableTimeseriesDataPointQueryResult{value: val, isSet: true}
}

func (v NullableTimeseriesDataPointQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesDataPointQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}