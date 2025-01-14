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

// checks if the EventCreation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventCreation{}

// EventCreation Configuration of the custom event.
type EventCreation struct {
	// Allow Davis AI to merge this event into existing problems (true) or force creating a new problem (false).  This only applies to problem-opening event types.
	AllowDavisMerge *bool `json:"allowDavisMerge,omitempty"`
	// A detailed description of the custom annotation, for example `DNS route has been changed to x.mydomain.com`.
	AnnotationDescription *string `json:"annotationDescription,omitempty"`
	// The type of the custom annotation, for example `DNS route has been changed`.
	AnnotationType *string              `json:"annotationType,omitempty"`
	AttachRules    PushEventAttachRules `json:"attachRules"`
	// The new value of the configuration that has been set by the event.
	Changed *string `json:"changed,omitempty"`
	// The link to the deployed artifact within the 3rd party system.
	CiBackLink *string `json:"ciBackLink,omitempty"`
	// The ID or the name of the configuration that has been changed by the event.
	Configuration *string `json:"configuration,omitempty"`
	// The set of any properties related to the event, in the *\"key\" : \"value\"* format.
	CustomProperties map[string]map[string]interface{} `json:"customProperties,omitempty"`
	// The ID of the triggered deployment.
	DeploymentName *string `json:"deploymentName,omitempty"`
	// The project name of the deployed package.
	DeploymentProject *string `json:"deploymentProject,omitempty"`
	// The version of the triggered deployment.
	DeploymentVersion *string `json:"deploymentVersion,omitempty"`
	// The textual description of the configuration change.
	Description *string `json:"description,omitempty"`
	// The end timestamp of the event, in UTC milliseconds.   If not set, the current time is used for information-only events.   Not applicable to problem-opening events. Such an event stays open until it times out depending on the **timeoutMinutes** parameter.
	End *int64 `json:"end,omitempty"`
	// The type of the event.
	EventType string `json:"eventType"`
	// The previous value of the configuration.
	Original *string `json:"original,omitempty"`
	// The link to the deployment related remediation action within the external tool.
	RemediationAction *string `json:"remediationAction,omitempty"`
	// The name or ID of the external source of the event.
	Source string `json:"source"`
	// The start timestamp of the event, in UTC milliseconds.   If not set, the current timestamp is used.   You can report information-only events up to **30 days** into the past.   You can report problem-opening events up to **60 minutes** into the past.
	Start *int64 `json:"start,omitempty"`
	// The timeout for problem-opening events in minutes. Not applicable to information-only events.   If not set, 15 minutes is used. The maximum allowed value is 120 minutes.   You can refresh the event by sending the same payload again.
	TimeoutMinutes *int32 `json:"timeoutMinutes,omitempty"`
	// A list of timeseries IDs, related to the event.
	TimeseriesIds []string `json:"timeseriesIds,omitempty"`
	// The title of the configuration that has been set by the event.
	Title *string `json:"title,omitempty"`
}

type _EventCreation EventCreation

// NewEventCreation instantiates a new EventCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCreation(attachRules PushEventAttachRules, eventType string, source string) *EventCreation {
	this := EventCreation{}
	this.AttachRules = attachRules
	this.EventType = eventType
	this.Source = source
	return &this
}

// NewEventCreationWithDefaults instantiates a new EventCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCreationWithDefaults() *EventCreation {
	this := EventCreation{}
	return &this
}

// GetAllowDavisMerge returns the AllowDavisMerge field value if set, zero value otherwise.
func (o *EventCreation) GetAllowDavisMerge() bool {
	if o == nil || IsNil(o.AllowDavisMerge) {
		var ret bool
		return ret
	}
	return *o.AllowDavisMerge
}

// GetAllowDavisMergeOk returns a tuple with the AllowDavisMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetAllowDavisMergeOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowDavisMerge) {
		return nil, false
	}
	return o.AllowDavisMerge, true
}

// HasAllowDavisMerge returns a boolean if a field has been set.
func (o *EventCreation) HasAllowDavisMerge() bool {
	if o != nil && !IsNil(o.AllowDavisMerge) {
		return true
	}

	return false
}

// SetAllowDavisMerge gets a reference to the given bool and assigns it to the AllowDavisMerge field.
func (o *EventCreation) SetAllowDavisMerge(v bool) {
	o.AllowDavisMerge = &v
}

// GetAnnotationDescription returns the AnnotationDescription field value if set, zero value otherwise.
func (o *EventCreation) GetAnnotationDescription() string {
	if o == nil || IsNil(o.AnnotationDescription) {
		var ret string
		return ret
	}
	return *o.AnnotationDescription
}

// GetAnnotationDescriptionOk returns a tuple with the AnnotationDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetAnnotationDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.AnnotationDescription) {
		return nil, false
	}
	return o.AnnotationDescription, true
}

// HasAnnotationDescription returns a boolean if a field has been set.
func (o *EventCreation) HasAnnotationDescription() bool {
	if o != nil && !IsNil(o.AnnotationDescription) {
		return true
	}

	return false
}

// SetAnnotationDescription gets a reference to the given string and assigns it to the AnnotationDescription field.
func (o *EventCreation) SetAnnotationDescription(v string) {
	o.AnnotationDescription = &v
}

// GetAnnotationType returns the AnnotationType field value if set, zero value otherwise.
func (o *EventCreation) GetAnnotationType() string {
	if o == nil || IsNil(o.AnnotationType) {
		var ret string
		return ret
	}
	return *o.AnnotationType
}

// GetAnnotationTypeOk returns a tuple with the AnnotationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetAnnotationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AnnotationType) {
		return nil, false
	}
	return o.AnnotationType, true
}

// HasAnnotationType returns a boolean if a field has been set.
func (o *EventCreation) HasAnnotationType() bool {
	if o != nil && !IsNil(o.AnnotationType) {
		return true
	}

	return false
}

// SetAnnotationType gets a reference to the given string and assigns it to the AnnotationType field.
func (o *EventCreation) SetAnnotationType(v string) {
	o.AnnotationType = &v
}

// GetAttachRules returns the AttachRules field value
func (o *EventCreation) GetAttachRules() PushEventAttachRules {
	if o == nil {
		var ret PushEventAttachRules
		return ret
	}

	return o.AttachRules
}

// GetAttachRulesOk returns a tuple with the AttachRules field value
// and a boolean to check if the value has been set.
func (o *EventCreation) GetAttachRulesOk() (*PushEventAttachRules, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttachRules, true
}

// SetAttachRules sets field value
func (o *EventCreation) SetAttachRules(v PushEventAttachRules) {
	o.AttachRules = v
}

// GetChanged returns the Changed field value if set, zero value otherwise.
func (o *EventCreation) GetChanged() string {
	if o == nil || IsNil(o.Changed) {
		var ret string
		return ret
	}
	return *o.Changed
}

// GetChangedOk returns a tuple with the Changed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetChangedOk() (*string, bool) {
	if o == nil || IsNil(o.Changed) {
		return nil, false
	}
	return o.Changed, true
}

// HasChanged returns a boolean if a field has been set.
func (o *EventCreation) HasChanged() bool {
	if o != nil && !IsNil(o.Changed) {
		return true
	}

	return false
}

// SetChanged gets a reference to the given string and assigns it to the Changed field.
func (o *EventCreation) SetChanged(v string) {
	o.Changed = &v
}

// GetCiBackLink returns the CiBackLink field value if set, zero value otherwise.
func (o *EventCreation) GetCiBackLink() string {
	if o == nil || IsNil(o.CiBackLink) {
		var ret string
		return ret
	}
	return *o.CiBackLink
}

// GetCiBackLinkOk returns a tuple with the CiBackLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetCiBackLinkOk() (*string, bool) {
	if o == nil || IsNil(o.CiBackLink) {
		return nil, false
	}
	return o.CiBackLink, true
}

// HasCiBackLink returns a boolean if a field has been set.
func (o *EventCreation) HasCiBackLink() bool {
	if o != nil && !IsNil(o.CiBackLink) {
		return true
	}

	return false
}

// SetCiBackLink gets a reference to the given string and assigns it to the CiBackLink field.
func (o *EventCreation) SetCiBackLink(v string) {
	o.CiBackLink = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *EventCreation) GetConfiguration() string {
	if o == nil || IsNil(o.Configuration) {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetConfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *EventCreation) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *EventCreation) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *EventCreation) GetCustomProperties() map[string]map[string]interface{} {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetCustomPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return map[string]map[string]interface{}{}, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *EventCreation) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the CustomProperties field.
func (o *EventCreation) SetCustomProperties(v map[string]map[string]interface{}) {
	o.CustomProperties = v
}

// GetDeploymentName returns the DeploymentName field value if set, zero value otherwise.
func (o *EventCreation) GetDeploymentName() string {
	if o == nil || IsNil(o.DeploymentName) {
		var ret string
		return ret
	}
	return *o.DeploymentName
}

// GetDeploymentNameOk returns a tuple with the DeploymentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetDeploymentNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentName) {
		return nil, false
	}
	return o.DeploymentName, true
}

// HasDeploymentName returns a boolean if a field has been set.
func (o *EventCreation) HasDeploymentName() bool {
	if o != nil && !IsNil(o.DeploymentName) {
		return true
	}

	return false
}

// SetDeploymentName gets a reference to the given string and assigns it to the DeploymentName field.
func (o *EventCreation) SetDeploymentName(v string) {
	o.DeploymentName = &v
}

// GetDeploymentProject returns the DeploymentProject field value if set, zero value otherwise.
func (o *EventCreation) GetDeploymentProject() string {
	if o == nil || IsNil(o.DeploymentProject) {
		var ret string
		return ret
	}
	return *o.DeploymentProject
}

// GetDeploymentProjectOk returns a tuple with the DeploymentProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetDeploymentProjectOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentProject) {
		return nil, false
	}
	return o.DeploymentProject, true
}

// HasDeploymentProject returns a boolean if a field has been set.
func (o *EventCreation) HasDeploymentProject() bool {
	if o != nil && !IsNil(o.DeploymentProject) {
		return true
	}

	return false
}

// SetDeploymentProject gets a reference to the given string and assigns it to the DeploymentProject field.
func (o *EventCreation) SetDeploymentProject(v string) {
	o.DeploymentProject = &v
}

// GetDeploymentVersion returns the DeploymentVersion field value if set, zero value otherwise.
func (o *EventCreation) GetDeploymentVersion() string {
	if o == nil || IsNil(o.DeploymentVersion) {
		var ret string
		return ret
	}
	return *o.DeploymentVersion
}

// GetDeploymentVersionOk returns a tuple with the DeploymentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetDeploymentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentVersion) {
		return nil, false
	}
	return o.DeploymentVersion, true
}

// HasDeploymentVersion returns a boolean if a field has been set.
func (o *EventCreation) HasDeploymentVersion() bool {
	if o != nil && !IsNil(o.DeploymentVersion) {
		return true
	}

	return false
}

// SetDeploymentVersion gets a reference to the given string and assigns it to the DeploymentVersion field.
func (o *EventCreation) SetDeploymentVersion(v string) {
	o.DeploymentVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventCreation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventCreation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventCreation) SetDescription(v string) {
	o.Description = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *EventCreation) GetEnd() int64 {
	if o == nil || IsNil(o.End) {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetEndOk() (*int64, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *EventCreation) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *EventCreation) SetEnd(v int64) {
	o.End = &v
}

// GetEventType returns the EventType field value
func (o *EventCreation) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EventCreation) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EventCreation) SetEventType(v string) {
	o.EventType = v
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *EventCreation) GetOriginal() string {
	if o == nil || IsNil(o.Original) {
		var ret string
		return ret
	}
	return *o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetOriginalOk() (*string, bool) {
	if o == nil || IsNil(o.Original) {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *EventCreation) HasOriginal() bool {
	if o != nil && !IsNil(o.Original) {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given string and assigns it to the Original field.
func (o *EventCreation) SetOriginal(v string) {
	o.Original = &v
}

// GetRemediationAction returns the RemediationAction field value if set, zero value otherwise.
func (o *EventCreation) GetRemediationAction() string {
	if o == nil || IsNil(o.RemediationAction) {
		var ret string
		return ret
	}
	return *o.RemediationAction
}

// GetRemediationActionOk returns a tuple with the RemediationAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetRemediationActionOk() (*string, bool) {
	if o == nil || IsNil(o.RemediationAction) {
		return nil, false
	}
	return o.RemediationAction, true
}

// HasRemediationAction returns a boolean if a field has been set.
func (o *EventCreation) HasRemediationAction() bool {
	if o != nil && !IsNil(o.RemediationAction) {
		return true
	}

	return false
}

// SetRemediationAction gets a reference to the given string and assigns it to the RemediationAction field.
func (o *EventCreation) SetRemediationAction(v string) {
	o.RemediationAction = &v
}

// GetSource returns the Source field value
func (o *EventCreation) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EventCreation) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *EventCreation) SetSource(v string) {
	o.Source = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *EventCreation) GetStart() int64 {
	if o == nil || IsNil(o.Start) {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetStartOk() (*int64, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *EventCreation) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *EventCreation) SetStart(v int64) {
	o.Start = &v
}

// GetTimeoutMinutes returns the TimeoutMinutes field value if set, zero value otherwise.
func (o *EventCreation) GetTimeoutMinutes() int32 {
	if o == nil || IsNil(o.TimeoutMinutes) {
		var ret int32
		return ret
	}
	return *o.TimeoutMinutes
}

// GetTimeoutMinutesOk returns a tuple with the TimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetTimeoutMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutMinutes) {
		return nil, false
	}
	return o.TimeoutMinutes, true
}

// HasTimeoutMinutes returns a boolean if a field has been set.
func (o *EventCreation) HasTimeoutMinutes() bool {
	if o != nil && !IsNil(o.TimeoutMinutes) {
		return true
	}

	return false
}

// SetTimeoutMinutes gets a reference to the given int32 and assigns it to the TimeoutMinutes field.
func (o *EventCreation) SetTimeoutMinutes(v int32) {
	o.TimeoutMinutes = &v
}

// GetTimeseriesIds returns the TimeseriesIds field value if set, zero value otherwise.
func (o *EventCreation) GetTimeseriesIds() []string {
	if o == nil || IsNil(o.TimeseriesIds) {
		var ret []string
		return ret
	}
	return o.TimeseriesIds
}

// GetTimeseriesIdsOk returns a tuple with the TimeseriesIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetTimeseriesIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TimeseriesIds) {
		return nil, false
	}
	return o.TimeseriesIds, true
}

// HasTimeseriesIds returns a boolean if a field has been set.
func (o *EventCreation) HasTimeseriesIds() bool {
	if o != nil && !IsNil(o.TimeseriesIds) {
		return true
	}

	return false
}

// SetTimeseriesIds gets a reference to the given []string and assigns it to the TimeseriesIds field.
func (o *EventCreation) SetTimeseriesIds(v []string) {
	o.TimeseriesIds = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EventCreation) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreation) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EventCreation) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *EventCreation) SetTitle(v string) {
	o.Title = &v
}

func (o EventCreation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventCreation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowDavisMerge) {
		toSerialize["allowDavisMerge"] = o.AllowDavisMerge
	}
	if !IsNil(o.AnnotationDescription) {
		toSerialize["annotationDescription"] = o.AnnotationDescription
	}
	if !IsNil(o.AnnotationType) {
		toSerialize["annotationType"] = o.AnnotationType
	}
	toSerialize["attachRules"] = o.AttachRules
	if !IsNil(o.Changed) {
		toSerialize["changed"] = o.Changed
	}
	if !IsNil(o.CiBackLink) {
		toSerialize["ciBackLink"] = o.CiBackLink
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.DeploymentName) {
		toSerialize["deploymentName"] = o.DeploymentName
	}
	if !IsNil(o.DeploymentProject) {
		toSerialize["deploymentProject"] = o.DeploymentProject
	}
	if !IsNil(o.DeploymentVersion) {
		toSerialize["deploymentVersion"] = o.DeploymentVersion
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	toSerialize["eventType"] = o.EventType
	if !IsNil(o.Original) {
		toSerialize["original"] = o.Original
	}
	if !IsNil(o.RemediationAction) {
		toSerialize["remediationAction"] = o.RemediationAction
	}
	toSerialize["source"] = o.Source
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.TimeoutMinutes) {
		toSerialize["timeoutMinutes"] = o.TimeoutMinutes
	}
	if !IsNil(o.TimeseriesIds) {
		toSerialize["timeseriesIds"] = o.TimeseriesIds
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

func (o *EventCreation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attachRules",
		"eventType",
		"source",
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

	varEventCreation := _EventCreation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventCreation)

	if err != nil {
		return err
	}

	*o = EventCreation(varEventCreation)

	return err
}

type NullableEventCreation struct {
	value *EventCreation
	isSet bool
}

func (v NullableEventCreation) Get() *EventCreation {
	return v.value
}

func (v *NullableEventCreation) Set(val *EventCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCreation(val *EventCreation) *NullableEventCreation {
	return &NullableEventCreation{value: val, isSet: true}
}

func (v NullableEventCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
