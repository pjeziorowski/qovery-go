/*
 * [BETA] Qovery API
 *
 * - Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.
 *
 * API version: 1.0.0
 * Contact: support+api+documentation@qovery.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// MetricMemoryDatapointResponse struct for MetricMemoryDatapointResponse
type MetricMemoryDatapointResponse struct {
	CreatedAt         time.Time `json:"created_at"`
	RequestedInMb     int32     `json:"requested_in_mb"`
	ConsumedInMb      int32     `json:"consumed_in_mb"`
	ConsumedInPercent float32   `json:"consumed_in_percent"`
}

// NewMetricMemoryDatapointResponse instantiates a new MetricMemoryDatapointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricMemoryDatapointResponse(createdAt time.Time, requestedInMb int32, consumedInMb int32, consumedInPercent float32) *MetricMemoryDatapointResponse {
	this := MetricMemoryDatapointResponse{}
	this.CreatedAt = createdAt
	this.RequestedInMb = requestedInMb
	this.ConsumedInMb = consumedInMb
	this.ConsumedInPercent = consumedInPercent
	return &this
}

// NewMetricMemoryDatapointResponseWithDefaults instantiates a new MetricMemoryDatapointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricMemoryDatapointResponseWithDefaults() *MetricMemoryDatapointResponse {
	this := MetricMemoryDatapointResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *MetricMemoryDatapointResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MetricMemoryDatapointResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MetricMemoryDatapointResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetRequestedInMb returns the RequestedInMb field value
func (o *MetricMemoryDatapointResponse) GetRequestedInMb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequestedInMb
}

// GetRequestedInMbOk returns a tuple with the RequestedInMb field value
// and a boolean to check if the value has been set.
func (o *MetricMemoryDatapointResponse) GetRequestedInMbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedInMb, true
}

// SetRequestedInMb sets field value
func (o *MetricMemoryDatapointResponse) SetRequestedInMb(v int32) {
	o.RequestedInMb = v
}

// GetConsumedInMb returns the ConsumedInMb field value
func (o *MetricMemoryDatapointResponse) GetConsumedInMb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConsumedInMb
}

// GetConsumedInMbOk returns a tuple with the ConsumedInMb field value
// and a boolean to check if the value has been set.
func (o *MetricMemoryDatapointResponse) GetConsumedInMbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumedInMb, true
}

// SetConsumedInMb sets field value
func (o *MetricMemoryDatapointResponse) SetConsumedInMb(v int32) {
	o.ConsumedInMb = v
}

// GetConsumedInPercent returns the ConsumedInPercent field value
func (o *MetricMemoryDatapointResponse) GetConsumedInPercent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ConsumedInPercent
}

// GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field value
// and a boolean to check if the value has been set.
func (o *MetricMemoryDatapointResponse) GetConsumedInPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumedInPercent, true
}

// SetConsumedInPercent sets field value
func (o *MetricMemoryDatapointResponse) SetConsumedInPercent(v float32) {
	o.ConsumedInPercent = v
}

func (o MetricMemoryDatapointResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["requested_in_mb"] = o.RequestedInMb
	}
	if true {
		toSerialize["consumed_in_mb"] = o.ConsumedInMb
	}
	if true {
		toSerialize["consumed_in_percent"] = o.ConsumedInPercent
	}
	return json.Marshal(toSerialize)
}

type NullableMetricMemoryDatapointResponse struct {
	value *MetricMemoryDatapointResponse
	isSet bool
}

func (v NullableMetricMemoryDatapointResponse) Get() *MetricMemoryDatapointResponse {
	return v.value
}

func (v *NullableMetricMemoryDatapointResponse) Set(val *MetricMemoryDatapointResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricMemoryDatapointResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricMemoryDatapointResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricMemoryDatapointResponse(val *MetricMemoryDatapointResponse) *NullableMetricMemoryDatapointResponse {
	return &NullableMetricMemoryDatapointResponse{value: val, isSet: true}
}

func (v NullableMetricMemoryDatapointResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricMemoryDatapointResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
