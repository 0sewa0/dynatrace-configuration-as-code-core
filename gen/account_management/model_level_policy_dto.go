/*
Dynatrace Account Management API

The enterprise management API for Dynatrace SaaS enables automation of operational tasks related to user access and environment lifecycle management.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmanagement

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the LevelPolicyDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LevelPolicyDto{}

// LevelPolicyDto struct for LevelPolicyDto
type LevelPolicyDto struct {
	// The ID of the policy.
	Uuid string `json:"uuid"`
	// The display name of the policy.
	Name string `json:"name"`
	// A list of tags.
	Tags []string `json:"tags"`
	// A short description of the policy.
	Description string `json:"description"`
	// The [statement](https://dt-url.net/ht03ucb) of the policy.
	StatementQuery string `json:"statementQuery"`
	// The expanded form of the policy statement.
	Statements []Statement `json:"statements"`
}

type _LevelPolicyDto LevelPolicyDto

// NewLevelPolicyDto instantiates a new LevelPolicyDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLevelPolicyDto(uuid string, name string, tags []string, description string, statementQuery string, statements []Statement) *LevelPolicyDto {
	this := LevelPolicyDto{}
	this.Uuid = uuid
	this.Name = name
	this.Tags = tags
	this.Description = description
	this.StatementQuery = statementQuery
	this.Statements = statements
	return &this
}

// NewLevelPolicyDtoWithDefaults instantiates a new LevelPolicyDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLevelPolicyDtoWithDefaults() *LevelPolicyDto {
	this := LevelPolicyDto{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *LevelPolicyDto) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *LevelPolicyDto) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *LevelPolicyDto) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *LevelPolicyDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LevelPolicyDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LevelPolicyDto) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value
func (o *LevelPolicyDto) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *LevelPolicyDto) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *LevelPolicyDto) SetTags(v []string) {
	o.Tags = v
}

// GetDescription returns the Description field value
func (o *LevelPolicyDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LevelPolicyDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LevelPolicyDto) SetDescription(v string) {
	o.Description = v
}

// GetStatementQuery returns the StatementQuery field value
func (o *LevelPolicyDto) GetStatementQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatementQuery
}

// GetStatementQueryOk returns a tuple with the StatementQuery field value
// and a boolean to check if the value has been set.
func (o *LevelPolicyDto) GetStatementQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatementQuery, true
}

// SetStatementQuery sets field value
func (o *LevelPolicyDto) SetStatementQuery(v string) {
	o.StatementQuery = v
}

// GetStatements returns the Statements field value
func (o *LevelPolicyDto) GetStatements() []Statement {
	if o == nil {
		var ret []Statement
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *LevelPolicyDto) GetStatementsOk() ([]Statement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statements, true
}

// SetStatements sets field value
func (o *LevelPolicyDto) SetStatements(v []Statement) {
	o.Statements = v
}

func (o LevelPolicyDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LevelPolicyDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	toSerialize["tags"] = o.Tags
	toSerialize["description"] = o.Description
	toSerialize["statementQuery"] = o.StatementQuery
	toSerialize["statements"] = o.Statements
	return toSerialize, nil
}

func (o *LevelPolicyDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"name",
		"tags",
		"description",
		"statementQuery",
		"statements",
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

	varLevelPolicyDto := _LevelPolicyDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLevelPolicyDto)

	if err != nil {
		return err
	}

	*o = LevelPolicyDto(varLevelPolicyDto)

	return err
}

type NullableLevelPolicyDto struct {
	value *LevelPolicyDto
	isSet bool
}

func (v NullableLevelPolicyDto) Get() *LevelPolicyDto {
	return v.value
}

func (v *NullableLevelPolicyDto) Set(val *LevelPolicyDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLevelPolicyDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLevelPolicyDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLevelPolicyDto(val *LevelPolicyDto) *NullableLevelPolicyDto {
	return &NullableLevelPolicyDto{value: val, isSet: true}
}

func (v NullableLevelPolicyDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLevelPolicyDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
