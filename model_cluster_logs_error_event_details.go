/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ClusterLogsErrorEventDetails struct for ClusterLogsErrorEventDetails
type ClusterLogsErrorEventDetails struct {
	// cloud provider used
	ProviderKind    *string                                      `json:"provider_kind,omitempty"`
	Region          *string                                      `json:"region,omitempty"`
	Transmitter     *ClusterLogsErrorEventDetailsTransmitter     `json:"transmitter,omitempty"`
	UnderlyingError *ClusterLogsErrorEventDetailsUnderlyingError `json:"underlying_error,omitempty"`
}

// NewClusterLogsErrorEventDetails instantiates a new ClusterLogsErrorEventDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLogsErrorEventDetails() *ClusterLogsErrorEventDetails {
	this := ClusterLogsErrorEventDetails{}
	return &this
}

// NewClusterLogsErrorEventDetailsWithDefaults instantiates a new ClusterLogsErrorEventDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLogsErrorEventDetailsWithDefaults() *ClusterLogsErrorEventDetails {
	this := ClusterLogsErrorEventDetails{}
	return &this
}

// GetProviderKind returns the ProviderKind field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetails) GetProviderKind() string {
	if o == nil || o.ProviderKind == nil {
		var ret string
		return ret
	}
	return *o.ProviderKind
}

// GetProviderKindOk returns a tuple with the ProviderKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetails) GetProviderKindOk() (*string, bool) {
	if o == nil || o.ProviderKind == nil {
		return nil, false
	}
	return o.ProviderKind, true
}

// HasProviderKind returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetails) HasProviderKind() bool {
	if o != nil && o.ProviderKind != nil {
		return true
	}

	return false
}

// SetProviderKind gets a reference to the given string and assigns it to the ProviderKind field.
func (o *ClusterLogsErrorEventDetails) SetProviderKind(v string) {
	o.ProviderKind = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetails) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetails) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetails) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ClusterLogsErrorEventDetails) SetRegion(v string) {
	o.Region = &v
}

// GetTransmitter returns the Transmitter field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetails) GetTransmitter() ClusterLogsErrorEventDetailsTransmitter {
	if o == nil || o.Transmitter == nil {
		var ret ClusterLogsErrorEventDetailsTransmitter
		return ret
	}
	return *o.Transmitter
}

// GetTransmitterOk returns a tuple with the Transmitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetails) GetTransmitterOk() (*ClusterLogsErrorEventDetailsTransmitter, bool) {
	if o == nil || o.Transmitter == nil {
		return nil, false
	}
	return o.Transmitter, true
}

// HasTransmitter returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetails) HasTransmitter() bool {
	if o != nil && o.Transmitter != nil {
		return true
	}

	return false
}

// SetTransmitter gets a reference to the given ClusterLogsErrorEventDetailsTransmitter and assigns it to the Transmitter field.
func (o *ClusterLogsErrorEventDetails) SetTransmitter(v ClusterLogsErrorEventDetailsTransmitter) {
	o.Transmitter = &v
}

// GetUnderlyingError returns the UnderlyingError field value if set, zero value otherwise.
func (o *ClusterLogsErrorEventDetails) GetUnderlyingError() ClusterLogsErrorEventDetailsUnderlyingError {
	if o == nil || o.UnderlyingError == nil {
		var ret ClusterLogsErrorEventDetailsUnderlyingError
		return ret
	}
	return *o.UnderlyingError
}

// GetUnderlyingErrorOk returns a tuple with the UnderlyingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogsErrorEventDetails) GetUnderlyingErrorOk() (*ClusterLogsErrorEventDetailsUnderlyingError, bool) {
	if o == nil || o.UnderlyingError == nil {
		return nil, false
	}
	return o.UnderlyingError, true
}

// HasUnderlyingError returns a boolean if a field has been set.
func (o *ClusterLogsErrorEventDetails) HasUnderlyingError() bool {
	if o != nil && o.UnderlyingError != nil {
		return true
	}

	return false
}

// SetUnderlyingError gets a reference to the given ClusterLogsErrorEventDetailsUnderlyingError and assigns it to the UnderlyingError field.
func (o *ClusterLogsErrorEventDetails) SetUnderlyingError(v ClusterLogsErrorEventDetailsUnderlyingError) {
	o.UnderlyingError = &v
}

func (o ClusterLogsErrorEventDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProviderKind != nil {
		toSerialize["provider_kind"] = o.ProviderKind
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Transmitter != nil {
		toSerialize["transmitter"] = o.Transmitter
	}
	if o.UnderlyingError != nil {
		toSerialize["underlying_error"] = o.UnderlyingError
	}
	return json.Marshal(toSerialize)
}

type NullableClusterLogsErrorEventDetails struct {
	value *ClusterLogsErrorEventDetails
	isSet bool
}

func (v NullableClusterLogsErrorEventDetails) Get() *ClusterLogsErrorEventDetails {
	return v.value
}

func (v *NullableClusterLogsErrorEventDetails) Set(val *ClusterLogsErrorEventDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLogsErrorEventDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLogsErrorEventDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLogsErrorEventDetails(val *ClusterLogsErrorEventDetails) *NullableClusterLogsErrorEventDetails {
	return &NullableClusterLogsErrorEventDetails{value: val, isSet: true}
}

func (v NullableClusterLogsErrorEventDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLogsErrorEventDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
