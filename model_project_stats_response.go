/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ProjectStatsResponse struct for ProjectStatsResponse
type ProjectStatsResponse struct {
	Id string `json:"id"`
	ServiceTotalNumber *float32 `json:"service_total_number,omitempty"`
	EnvironmentTotalNumber *float32 `json:"environment_total_number,omitempty"`
}

// NewProjectStatsResponse instantiates a new ProjectStatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectStatsResponse(id string) *ProjectStatsResponse {
	this := ProjectStatsResponse{}
	this.Id = id
	return &this
}

// NewProjectStatsResponseWithDefaults instantiates a new ProjectStatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectStatsResponseWithDefaults() *ProjectStatsResponse {
	this := ProjectStatsResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectStatsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectStatsResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectStatsResponse) SetId(v string) {
	o.Id = v
}

// GetServiceTotalNumber returns the ServiceTotalNumber field value if set, zero value otherwise.
func (o *ProjectStatsResponse) GetServiceTotalNumber() float32 {
	if o == nil || o.ServiceTotalNumber == nil {
		var ret float32
		return ret
	}
	return *o.ServiceTotalNumber
}

// GetServiceTotalNumberOk returns a tuple with the ServiceTotalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectStatsResponse) GetServiceTotalNumberOk() (*float32, bool) {
	if o == nil || o.ServiceTotalNumber == nil {
		return nil, false
	}
	return o.ServiceTotalNumber, true
}

// HasServiceTotalNumber returns a boolean if a field has been set.
func (o *ProjectStatsResponse) HasServiceTotalNumber() bool {
	if o != nil && o.ServiceTotalNumber != nil {
		return true
	}

	return false
}

// SetServiceTotalNumber gets a reference to the given float32 and assigns it to the ServiceTotalNumber field.
func (o *ProjectStatsResponse) SetServiceTotalNumber(v float32) {
	o.ServiceTotalNumber = &v
}

// GetEnvironmentTotalNumber returns the EnvironmentTotalNumber field value if set, zero value otherwise.
func (o *ProjectStatsResponse) GetEnvironmentTotalNumber() float32 {
	if o == nil || o.EnvironmentTotalNumber == nil {
		var ret float32
		return ret
	}
	return *o.EnvironmentTotalNumber
}

// GetEnvironmentTotalNumberOk returns a tuple with the EnvironmentTotalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectStatsResponse) GetEnvironmentTotalNumberOk() (*float32, bool) {
	if o == nil || o.EnvironmentTotalNumber == nil {
		return nil, false
	}
	return o.EnvironmentTotalNumber, true
}

// HasEnvironmentTotalNumber returns a boolean if a field has been set.
func (o *ProjectStatsResponse) HasEnvironmentTotalNumber() bool {
	if o != nil && o.EnvironmentTotalNumber != nil {
		return true
	}

	return false
}

// SetEnvironmentTotalNumber gets a reference to the given float32 and assigns it to the EnvironmentTotalNumber field.
func (o *ProjectStatsResponse) SetEnvironmentTotalNumber(v float32) {
	o.EnvironmentTotalNumber = &v
}

func (o ProjectStatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ServiceTotalNumber != nil {
		toSerialize["service_total_number"] = o.ServiceTotalNumber
	}
	if o.EnvironmentTotalNumber != nil {
		toSerialize["environment_total_number"] = o.EnvironmentTotalNumber
	}
	return json.Marshal(toSerialize)
}

type NullableProjectStatsResponse struct {
	value *ProjectStatsResponse
	isSet bool
}

func (v NullableProjectStatsResponse) Get() *ProjectStatsResponse {
	return v.value
}

func (v *NullableProjectStatsResponse) Set(val *ProjectStatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectStatsResponse(val *ProjectStatsResponse) *NullableProjectStatsResponse {
	return &NullableProjectStatsResponse{value: val, isSet: true}
}

func (v NullableProjectStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


