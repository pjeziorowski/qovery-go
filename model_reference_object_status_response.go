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
)

// ReferenceObjectStatusResponse struct for ReferenceObjectStatusResponse
type ReferenceObjectStatusResponse struct {
	Id string `json:"id"`
	// Status is a state machine. It starts with `BUILDING` or `DEPLOYING` state (or `INITIALIZED`if auto-deploy is deactivated). Then finish with `*_ERROR` or any termination state.
	State string `json:"state"`
	// message related to the state
	Message NullableString `json:"message,omitempty"`
}

// NewReferenceObjectStatusResponse instantiates a new ReferenceObjectStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceObjectStatusResponse(id string, state string) *ReferenceObjectStatusResponse {
	this := ReferenceObjectStatusResponse{}
	this.Id = id
	this.State = state
	return &this
}

// NewReferenceObjectStatusResponseWithDefaults instantiates a new ReferenceObjectStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceObjectStatusResponseWithDefaults() *ReferenceObjectStatusResponse {
	this := ReferenceObjectStatusResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ReferenceObjectStatusResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatusResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReferenceObjectStatusResponse) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *ReferenceObjectStatusResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ReferenceObjectStatusResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ReferenceObjectStatusResponse) SetState(v string) {
	o.State = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferenceObjectStatusResponse) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferenceObjectStatusResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ReferenceObjectStatusResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ReferenceObjectStatusResponse) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *ReferenceObjectStatusResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ReferenceObjectStatusResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o ReferenceObjectStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableReferenceObjectStatusResponse struct {
	value *ReferenceObjectStatusResponse
	isSet bool
}

func (v NullableReferenceObjectStatusResponse) Get() *ReferenceObjectStatusResponse {
	return v.value
}

func (v *NullableReferenceObjectStatusResponse) Set(val *ReferenceObjectStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceObjectStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceObjectStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceObjectStatusResponse(val *ReferenceObjectStatusResponse) *NullableReferenceObjectStatusResponse {
	return &NullableReferenceObjectStatusResponse{value: val, isSet: true}
}

func (v NullableReferenceObjectStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceObjectStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
