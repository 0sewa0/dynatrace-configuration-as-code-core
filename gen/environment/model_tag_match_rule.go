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

// checks if the TagMatchRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagMatchRule{}

// TagMatchRule The list of tags to be used for matching Dynatrace entities.
type TagMatchRule struct {
	// The list of types of the Dynatrace entities (for example hosts or services) you want to pick up by matching.
	MeTypes []string `json:"meTypes"`
	// The list of tags you want to use for matching. At least one tag is required.    You can use custom tags from the UI, imported tags, and tags based on environment variables.
	Tags []TagInfo `json:"tags"`
}

type _TagMatchRule TagMatchRule

// NewTagMatchRule instantiates a new TagMatchRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagMatchRule(meTypes []string, tags []TagInfo) *TagMatchRule {
	this := TagMatchRule{}
	this.MeTypes = meTypes
	this.Tags = tags
	return &this
}

// NewTagMatchRuleWithDefaults instantiates a new TagMatchRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagMatchRuleWithDefaults() *TagMatchRule {
	this := TagMatchRule{}
	return &this
}

// GetMeTypes returns the MeTypes field value
func (o *TagMatchRule) GetMeTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MeTypes
}

// GetMeTypesOk returns a tuple with the MeTypes field value
// and a boolean to check if the value has been set.
func (o *TagMatchRule) GetMeTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeTypes, true
}

// SetMeTypes sets field value
func (o *TagMatchRule) SetMeTypes(v []string) {
	o.MeTypes = v
}

// GetTags returns the Tags field value
func (o *TagMatchRule) GetTags() []TagInfo {
	if o == nil {
		var ret []TagInfo
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagMatchRule) GetTagsOk() ([]TagInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *TagMatchRule) SetTags(v []TagInfo) {
	o.Tags = v
}

func (o TagMatchRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagMatchRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meTypes"] = o.MeTypes
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

func (o *TagMatchRule) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meTypes",
		"tags",
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

	varTagMatchRule := _TagMatchRule{}

	err = json.Unmarshal(bytes, &varTagMatchRule)

	if err != nil {
		return err
	}

	*o = TagMatchRule(varTagMatchRule)

	return err
}

type NullableTagMatchRule struct {
	value *TagMatchRule
	isSet bool
}

func (v NullableTagMatchRule) Get() *TagMatchRule {
	return v.value
}

func (v *NullableTagMatchRule) Set(val *TagMatchRule) {
	v.value = val
	v.isSet = true
}

func (v NullableTagMatchRule) IsSet() bool {
	return v.isSet
}

func (v *NullableTagMatchRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagMatchRule(val *TagMatchRule) *NullableTagMatchRule {
	return &NullableTagMatchRule{value: val, isSet: true}
}

func (v NullableTagMatchRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagMatchRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}