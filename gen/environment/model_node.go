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

// checks if the Node type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Node{}

// Node Configuration of a synthetic node.    A *synthetic node* is an ActiveGate that is able to execute synthetic monitors.
type Node struct {
	// The version of the Active Gate.
	ActiveGateVersion string `json:"activeGateVersion"`
	// The Active Gate has the Auto update option enabled ('true') or not ('false')
	AutoUpdateEnabled bool `json:"autoUpdateEnabled"`
	// The synthetic node is able to execute browser monitors (`true`) or not (`false`).
	BrowserMonitorsEnabled bool `json:"browserMonitorsEnabled"`
	// The browser type.
	BrowserType string `json:"browserType"`
	// The browser version.
	BrowserVersion string `json:"browserVersion"`
	// The ID of the synthetic node.
	EntityId string `json:"entityId"`
	// The health check status of the synthetic node.
	HealthCheckStatus string `json:"healthCheckStatus"`
	// The hostname of the synthetic node.
	Hostname string `json:"hostname"`
	// The IP of the synthetic node.
	Ips []string `json:"ips"`
	// The Active Gate has the One Agent routing enabled ('true') or not ('false').
	OneAgentRoutingEnabled bool `json:"oneAgentRoutingEnabled"`
	// The Active Gate's host operating system.
	OperatingSystem string `json:"operatingSystem"`
	// The version of the synthetic player.
	PlayerVersion string `json:"playerVersion"`
	// The status of the synthetic node.
	Status string `json:"status"`
	// The version of the synthetic node.
	Version string `json:"version"`
}

type _Node Node

// NewNode instantiates a new Node object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNode(activeGateVersion string, autoUpdateEnabled bool, browserMonitorsEnabled bool, browserType string, browserVersion string, entityId string, healthCheckStatus string, hostname string, ips []string, oneAgentRoutingEnabled bool, operatingSystem string, playerVersion string, status string, version string) *Node {
	this := Node{}
	this.ActiveGateVersion = activeGateVersion
	this.AutoUpdateEnabled = autoUpdateEnabled
	this.BrowserMonitorsEnabled = browserMonitorsEnabled
	this.BrowserType = browserType
	this.BrowserVersion = browserVersion
	this.EntityId = entityId
	this.HealthCheckStatus = healthCheckStatus
	this.Hostname = hostname
	this.Ips = ips
	this.OneAgentRoutingEnabled = oneAgentRoutingEnabled
	this.OperatingSystem = operatingSystem
	this.PlayerVersion = playerVersion
	this.Status = status
	this.Version = version
	return &this
}

// NewNodeWithDefaults instantiates a new Node object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeWithDefaults() *Node {
	this := Node{}
	return &this
}

// GetActiveGateVersion returns the ActiveGateVersion field value
func (o *Node) GetActiveGateVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveGateVersion
}

// GetActiveGateVersionOk returns a tuple with the ActiveGateVersion field value
// and a boolean to check if the value has been set.
func (o *Node) GetActiveGateVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveGateVersion, true
}

// SetActiveGateVersion sets field value
func (o *Node) SetActiveGateVersion(v string) {
	o.ActiveGateVersion = v
}

// GetAutoUpdateEnabled returns the AutoUpdateEnabled field value
func (o *Node) GetAutoUpdateEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoUpdateEnabled
}

// GetAutoUpdateEnabledOk returns a tuple with the AutoUpdateEnabled field value
// and a boolean to check if the value has been set.
func (o *Node) GetAutoUpdateEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoUpdateEnabled, true
}

// SetAutoUpdateEnabled sets field value
func (o *Node) SetAutoUpdateEnabled(v bool) {
	o.AutoUpdateEnabled = v
}

// GetBrowserMonitorsEnabled returns the BrowserMonitorsEnabled field value
func (o *Node) GetBrowserMonitorsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BrowserMonitorsEnabled
}

// GetBrowserMonitorsEnabledOk returns a tuple with the BrowserMonitorsEnabled field value
// and a boolean to check if the value has been set.
func (o *Node) GetBrowserMonitorsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserMonitorsEnabled, true
}

// SetBrowserMonitorsEnabled sets field value
func (o *Node) SetBrowserMonitorsEnabled(v bool) {
	o.BrowserMonitorsEnabled = v
}

// GetBrowserType returns the BrowserType field value
func (o *Node) GetBrowserType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrowserType
}

// GetBrowserTypeOk returns a tuple with the BrowserType field value
// and a boolean to check if the value has been set.
func (o *Node) GetBrowserTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserType, true
}

// SetBrowserType sets field value
func (o *Node) SetBrowserType(v string) {
	o.BrowserType = v
}

// GetBrowserVersion returns the BrowserVersion field value
func (o *Node) GetBrowserVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrowserVersion
}

// GetBrowserVersionOk returns a tuple with the BrowserVersion field value
// and a boolean to check if the value has been set.
func (o *Node) GetBrowserVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserVersion, true
}

// SetBrowserVersion sets field value
func (o *Node) SetBrowserVersion(v string) {
	o.BrowserVersion = v
}

// GetEntityId returns the EntityId field value
func (o *Node) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *Node) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *Node) SetEntityId(v string) {
	o.EntityId = v
}

// GetHealthCheckStatus returns the HealthCheckStatus field value
func (o *Node) GetHealthCheckStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HealthCheckStatus
}

// GetHealthCheckStatusOk returns a tuple with the HealthCheckStatus field value
// and a boolean to check if the value has been set.
func (o *Node) GetHealthCheckStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthCheckStatus, true
}

// SetHealthCheckStatus sets field value
func (o *Node) SetHealthCheckStatus(v string) {
	o.HealthCheckStatus = v
}

// GetHostname returns the Hostname field value
func (o *Node) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *Node) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *Node) SetHostname(v string) {
	o.Hostname = v
}

// GetIps returns the Ips field value
func (o *Node) GetIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *Node) GetIpsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *Node) SetIps(v []string) {
	o.Ips = v
}

// GetOneAgentRoutingEnabled returns the OneAgentRoutingEnabled field value
func (o *Node) GetOneAgentRoutingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OneAgentRoutingEnabled
}

// GetOneAgentRoutingEnabledOk returns a tuple with the OneAgentRoutingEnabled field value
// and a boolean to check if the value has been set.
func (o *Node) GetOneAgentRoutingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OneAgentRoutingEnabled, true
}

// SetOneAgentRoutingEnabled sets field value
func (o *Node) SetOneAgentRoutingEnabled(v bool) {
	o.OneAgentRoutingEnabled = v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *Node) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *Node) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *Node) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetPlayerVersion returns the PlayerVersion field value
func (o *Node) GetPlayerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerVersion
}

// GetPlayerVersionOk returns a tuple with the PlayerVersion field value
// and a boolean to check if the value has been set.
func (o *Node) GetPlayerVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerVersion, true
}

// SetPlayerVersion sets field value
func (o *Node) SetPlayerVersion(v string) {
	o.PlayerVersion = v
}

// GetStatus returns the Status field value
func (o *Node) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Node) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Node) SetStatus(v string) {
	o.Status = v
}

// GetVersion returns the Version field value
func (o *Node) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Node) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Node) SetVersion(v string) {
	o.Version = v
}

func (o Node) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Node) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activeGateVersion"] = o.ActiveGateVersion
	toSerialize["autoUpdateEnabled"] = o.AutoUpdateEnabled
	toSerialize["browserMonitorsEnabled"] = o.BrowserMonitorsEnabled
	toSerialize["browserType"] = o.BrowserType
	toSerialize["browserVersion"] = o.BrowserVersion
	toSerialize["entityId"] = o.EntityId
	toSerialize["healthCheckStatus"] = o.HealthCheckStatus
	toSerialize["hostname"] = o.Hostname
	toSerialize["ips"] = o.Ips
	toSerialize["oneAgentRoutingEnabled"] = o.OneAgentRoutingEnabled
	toSerialize["operatingSystem"] = o.OperatingSystem
	toSerialize["playerVersion"] = o.PlayerVersion
	toSerialize["status"] = o.Status
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *Node) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activeGateVersion",
		"autoUpdateEnabled",
		"browserMonitorsEnabled",
		"browserType",
		"browserVersion",
		"entityId",
		"healthCheckStatus",
		"hostname",
		"ips",
		"oneAgentRoutingEnabled",
		"operatingSystem",
		"playerVersion",
		"status",
		"version",
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

	varNode := _Node{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNode)

	if err != nil {
		return err
	}

	*o = Node(varNode)

	return err
}

type NullableNode struct {
	value *Node
	isSet bool
}

func (v NullableNode) Get() *Node {
	return v.value
}

func (v *NullableNode) Set(val *Node) {
	v.value = val
	v.isSet = true
}

func (v NullableNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode(val *Node) *NullableNode {
	return &NullableNode{value: val, isSet: true}
}

func (v NullableNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
