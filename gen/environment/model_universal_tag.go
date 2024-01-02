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

// checks if the UniversalTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniversalTag{}

// UniversalTag struct for UniversalTag
type UniversalTag struct {
	// The origin of the tag, such as AWS or Cloud Foundry. For custom tags use the `CONTEXTLESS` value.   The context is set for tags that are automatically imported by OneAgent (for example, from the AWS console or environment variables). It’s useful for determining the origin of tags when not manually defined, and it also helps to prevent clashes with other existing tags. If the tag is not automatically imported, `CONTEXTLESS` set.
	Context *string `json:"context,omitempty"`
	// The key of the tag. For custom tags, put the tag value here.  The key allows categorization of multiple tags. It is possible that there are multiple values for a single key which will all be represented as standalone tags. Therefore, the key does not have the semantic of a map key but is more like a key of a key-value tuple. In some cases, for example custom tags, the key represents the actual tag value and the value field is not set – those are called valueless tags.
	Key    string           `json:"key"`
	TagKey *UniversalTagKey `json:"tagKey,omitempty"`
	// The value of the tag. Not applicable to custom tags.   If a tag does have a separate key and value (in the textual representation they are split by the colon ‘:’), this field is set with the actual value. Key-value pairs can occur for automatically imported tags and tags set by rules if extractors are used.
	Value *string `json:"value,omitempty"`
}

type _UniversalTag UniversalTag

// NewUniversalTag instantiates a new UniversalTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniversalTag(key string) *UniversalTag {
	this := UniversalTag{}
	this.Key = key
	return &this
}

// NewUniversalTagWithDefaults instantiates a new UniversalTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniversalTagWithDefaults() *UniversalTag {
	this := UniversalTag{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *UniversalTag) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalTag) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *UniversalTag) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *UniversalTag) SetContext(v string) {
	o.Context = &v
}

// GetKey returns the Key field value
func (o *UniversalTag) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *UniversalTag) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *UniversalTag) SetKey(v string) {
	o.Key = v
}

// GetTagKey returns the TagKey field value if set, zero value otherwise.
func (o *UniversalTag) GetTagKey() UniversalTagKey {
	if o == nil || IsNil(o.TagKey) {
		var ret UniversalTagKey
		return ret
	}
	return *o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalTag) GetTagKeyOk() (*UniversalTagKey, bool) {
	if o == nil || IsNil(o.TagKey) {
		return nil, false
	}
	return o.TagKey, true
}

// HasTagKey returns a boolean if a field has been set.
func (o *UniversalTag) HasTagKey() bool {
	if o != nil && !IsNil(o.TagKey) {
		return true
	}

	return false
}

// SetTagKey gets a reference to the given UniversalTagKey and assigns it to the TagKey field.
func (o *UniversalTag) SetTagKey(v UniversalTagKey) {
	o.TagKey = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UniversalTag) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniversalTag) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UniversalTag) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UniversalTag) SetValue(v string) {
	o.Value = &v
}

func (o UniversalTag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniversalTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.TagKey) {
		toSerialize["tagKey"] = o.TagKey
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *UniversalTag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varUniversalTag := _UniversalTag{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUniversalTag)

	if err != nil {
		return err
	}

	*o = UniversalTag(varUniversalTag)

	return err
}

type NullableUniversalTag struct {
	value *UniversalTag
	isSet bool
}

func (v NullableUniversalTag) Get() *UniversalTag {
	return v.value
}

func (v *NullableUniversalTag) Set(val *UniversalTag) {
	v.value = val
	v.isSet = true
}

func (v NullableUniversalTag) IsSet() bool {
	return v.isSet
}

func (v *NullableUniversalTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniversalTag(val *UniversalTag) *NullableUniversalTag {
	return &NullableUniversalTag{value: val, isSet: true}
}

func (v NullableUniversalTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniversalTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
