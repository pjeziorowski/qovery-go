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

// OrganizationApiTokenResponseList struct for OrganizationApiTokenResponseList
type OrganizationApiTokenResponseList struct {
	Results []OrganizationApiToken `json:"results,omitempty"`
}

// NewOrganizationApiTokenResponseList instantiates a new OrganizationApiTokenResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationApiTokenResponseList() *OrganizationApiTokenResponseList {
	this := OrganizationApiTokenResponseList{}
	return &this
}

// NewOrganizationApiTokenResponseListWithDefaults instantiates a new OrganizationApiTokenResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationApiTokenResponseListWithDefaults() *OrganizationApiTokenResponseList {
	this := OrganizationApiTokenResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *OrganizationApiTokenResponseList) GetResults() []OrganizationApiToken {
	if o == nil || o.Results == nil {
		var ret []OrganizationApiToken
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenResponseList) GetResultsOk() ([]OrganizationApiToken, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OrganizationApiTokenResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OrganizationApiToken and assigns it to the Results field.
func (o *OrganizationApiTokenResponseList) SetResults(v []OrganizationApiToken) {
	o.Results = v
}

func (o OrganizationApiTokenResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationApiTokenResponseList struct {
	value *OrganizationApiTokenResponseList
	isSet bool
}

func (v NullableOrganizationApiTokenResponseList) Get() *OrganizationApiTokenResponseList {
	return v.value
}

func (v *NullableOrganizationApiTokenResponseList) Set(val *OrganizationApiTokenResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationApiTokenResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationApiTokenResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationApiTokenResponseList(val *OrganizationApiTokenResponseList) *NullableOrganizationApiTokenResponseList {
	return &NullableOrganizationApiTokenResponseList{value: val, isSet: true}
}

func (v NullableOrganizationApiTokenResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationApiTokenResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
