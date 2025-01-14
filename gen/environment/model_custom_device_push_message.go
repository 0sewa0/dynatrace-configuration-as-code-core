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

// checks if the CustomDevicePushMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDevicePushMessage{}

// CustomDevicePushMessage Configuration of a custom device.
type CustomDevicePushMessage struct {
	// The URL of a configuration web page for the custom device, such as a login page for a firewall or router.
	ConfigUrl *string `json:"configUrl,omitempty"`
	// The name of the custom device that will appear in the user interface.
	DisplayName *string `json:"displayName,omitempty"`
	// The icon to be displayed for your custom component within Smartscape. Provide the full URL of the icon file.
	Favicon *string `json:"favicon,omitempty"`
	// User defined group ID of entity.   The group ID helps to keep a consistent picture of device-group relations. One of many cases where a proper group is important is service detection: you can define which custom devices should lead to the same service by defining the same group ID for them.   If you set a group ID, it will be hashed into the Dynatrace entity ID of the custom device. In that case the custom device can only be part of one custom device group.   If you don't set the group ID, Dynatrace will create it based on the ID or type of the custom device. Also, the group will not be hashed into the device ID which means the device may switch groups.
	Group *string `json:"group,omitempty"`
	// The list of host names related to the custom device.   These names are used to automatically discover the horizontal communication relationship between this component and all other observed components within Smartscape. Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If you send a value, the existing values will be overwritten.   If you send `null` or an empty value; or omit this field, the existing values will be kept.
	HostNames []string `json:"hostNames,omitempty"`
	// The list of IP addresses that belong to the custom device.   These addresses are used to automatically discover the horizontal communication relationship between this component and all other observed components within Smartscape. Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If you send a value (including an empty value), the existing values will be overwritten.   If you send `null` or omit this field, the existing values will be kept.
	IpAddresses []string `json:"ipAddresses,omitempty"`
	// The list of ports the custom devices listens to.   These ports are used to discover the horizontal communication relationship between this component and all other observed components within Smartscape.   Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If ports are specified, you should also add at least one IP address or a host name for the custom device.   If you send a value, the existing values will be overwritten.   If you send `null`, or an empty value, or omit this field, the existing values will be kept.
	ListenPorts []int32 `json:"listenPorts,omitempty"`
	// The list of key-value pair properties that will be shown beneath the infographics of your custom device.
	Properties *map[string]string `json:"properties,omitempty"`
	// The list of metric values that are reported for the custom device.   The metric you're reporting must already exist in Dynatrace. To learn how to create a custom metric, see [Timeseries API v1 - PUT a custom metric](https://dt-url.net/5k143rzb).   Dynatrace aggregates all the values you report for a custom device.   If you send a value (including an empty value), it will be added to the set of existing values.   If you send `null` or omit this field, the set of existing values won't change.
	Series []EntityTimeseriesData `json:"series,omitempty"`
	// List of custom tags that you want to attach to your custom device.
	Tags []string `json:"tags,omitempty"`
	// The technology type definition of the custom device.   It must be the same technology type of the metric you're reporting.   If you send a value, the existing value will be overwritten.   If you send `null`, empty or omit this field, the existing value will be kept.
	Type *string `json:"type,omitempty"`
}

// NewCustomDevicePushMessage instantiates a new CustomDevicePushMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDevicePushMessage() *CustomDevicePushMessage {
	this := CustomDevicePushMessage{}
	return &this
}

// NewCustomDevicePushMessageWithDefaults instantiates a new CustomDevicePushMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDevicePushMessageWithDefaults() *CustomDevicePushMessage {
	this := CustomDevicePushMessage{}
	return &this
}

// GetConfigUrl returns the ConfigUrl field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetConfigUrl() string {
	if o == nil || IsNil(o.ConfigUrl) {
		var ret string
		return ret
	}
	return *o.ConfigUrl
}

// GetConfigUrlOk returns a tuple with the ConfigUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetConfigUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigUrl) {
		return nil, false
	}
	return o.ConfigUrl, true
}

// HasConfigUrl returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasConfigUrl() bool {
	if o != nil && !IsNil(o.ConfigUrl) {
		return true
	}

	return false
}

// SetConfigUrl gets a reference to the given string and assigns it to the ConfigUrl field.
func (o *CustomDevicePushMessage) SetConfigUrl(v string) {
	o.ConfigUrl = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CustomDevicePushMessage) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFavicon returns the Favicon field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetFavicon() string {
	if o == nil || IsNil(o.Favicon) {
		var ret string
		return ret
	}
	return *o.Favicon
}

// GetFaviconOk returns a tuple with the Favicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetFaviconOk() (*string, bool) {
	if o == nil || IsNil(o.Favicon) {
		return nil, false
	}
	return o.Favicon, true
}

// HasFavicon returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasFavicon() bool {
	if o != nil && !IsNil(o.Favicon) {
		return true
	}

	return false
}

// SetFavicon gets a reference to the given string and assigns it to the Favicon field.
func (o *CustomDevicePushMessage) SetFavicon(v string) {
	o.Favicon = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *CustomDevicePushMessage) SetGroup(v string) {
	o.Group = &v
}

// GetHostNames returns the HostNames field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetHostNames() []string {
	if o == nil || IsNil(o.HostNames) {
		var ret []string
		return ret
	}
	return o.HostNames
}

// GetHostNamesOk returns a tuple with the HostNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetHostNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.HostNames) {
		return nil, false
	}
	return o.HostNames, true
}

// HasHostNames returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasHostNames() bool {
	if o != nil && !IsNil(o.HostNames) {
		return true
	}

	return false
}

// SetHostNames gets a reference to the given []string and assigns it to the HostNames field.
func (o *CustomDevicePushMessage) SetHostNames(v []string) {
	o.HostNames = v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetIpAddresses() []string {
	if o == nil || IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetIpAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasIpAddresses() bool {
	if o != nil && !IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *CustomDevicePushMessage) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetListenPorts returns the ListenPorts field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetListenPorts() []int32 {
	if o == nil || IsNil(o.ListenPorts) {
		var ret []int32
		return ret
	}
	return o.ListenPorts
}

// GetListenPortsOk returns a tuple with the ListenPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetListenPortsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ListenPorts) {
		return nil, false
	}
	return o.ListenPorts, true
}

// HasListenPorts returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasListenPorts() bool {
	if o != nil && !IsNil(o.ListenPorts) {
		return true
	}

	return false
}

// SetListenPorts gets a reference to the given []int32 and assigns it to the ListenPorts field.
func (o *CustomDevicePushMessage) SetListenPorts(v []int32) {
	o.ListenPorts = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *CustomDevicePushMessage) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetSeries() []EntityTimeseriesData {
	if o == nil || IsNil(o.Series) {
		var ret []EntityTimeseriesData
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetSeriesOk() ([]EntityTimeseriesData, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []EntityTimeseriesData and assigns it to the Series field.
func (o *CustomDevicePushMessage) SetSeries(v []EntityTimeseriesData) {
	o.Series = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CustomDevicePushMessage) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomDevicePushMessage) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushMessage) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomDevicePushMessage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomDevicePushMessage) SetType(v string) {
	o.Type = &v
}

func (o CustomDevicePushMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDevicePushMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigUrl) {
		toSerialize["configUrl"] = o.ConfigUrl
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Favicon) {
		toSerialize["favicon"] = o.Favicon
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.HostNames) {
		toSerialize["hostNames"] = o.HostNames
	}
	if !IsNil(o.IpAddresses) {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if !IsNil(o.ListenPorts) {
		toSerialize["listenPorts"] = o.ListenPorts
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCustomDevicePushMessage struct {
	value *CustomDevicePushMessage
	isSet bool
}

func (v NullableCustomDevicePushMessage) Get() *CustomDevicePushMessage {
	return v.value
}

func (v *NullableCustomDevicePushMessage) Set(val *CustomDevicePushMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDevicePushMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDevicePushMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDevicePushMessage(val *CustomDevicePushMessage) *NullableCustomDevicePushMessage {
	return &NullableCustomDevicePushMessage{value: val, isSet: true}
}

func (v NullableCustomDevicePushMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDevicePushMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
