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

// CostRangeResponse struct for CostRangeResponse
type CostRangeResponse struct {
	MinCostInCents *int32 `json:"min_cost_in_cents,omitempty"`
	MinCost *float32 `json:"min_cost,omitempty"`
	MaxCostInCents *int32 `json:"max_cost_in_cents,omitempty"`
	MaxCost *float32 `json:"max_cost,omitempty"`
	CurrencyCode string `json:"currency_code"`
}

// NewCostRangeResponse instantiates a new CostRangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostRangeResponse(currencyCode string) *CostRangeResponse {
	this := CostRangeResponse{}
	this.CurrencyCode = currencyCode
	return &this
}

// NewCostRangeResponseWithDefaults instantiates a new CostRangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostRangeResponseWithDefaults() *CostRangeResponse {
	this := CostRangeResponse{}
	return &this
}

// GetMinCostInCents returns the MinCostInCents field value if set, zero value otherwise.
func (o *CostRangeResponse) GetMinCostInCents() int32 {
	if o == nil || o.MinCostInCents == nil {
		var ret int32
		return ret
	}
	return *o.MinCostInCents
}

// GetMinCostInCentsOk returns a tuple with the MinCostInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRangeResponse) GetMinCostInCentsOk() (*int32, bool) {
	if o == nil || o.MinCostInCents == nil {
		return nil, false
	}
	return o.MinCostInCents, true
}

// HasMinCostInCents returns a boolean if a field has been set.
func (o *CostRangeResponse) HasMinCostInCents() bool {
	if o != nil && o.MinCostInCents != nil {
		return true
	}

	return false
}

// SetMinCostInCents gets a reference to the given int32 and assigns it to the MinCostInCents field.
func (o *CostRangeResponse) SetMinCostInCents(v int32) {
	o.MinCostInCents = &v
}

// GetMinCost returns the MinCost field value if set, zero value otherwise.
func (o *CostRangeResponse) GetMinCost() float32 {
	if o == nil || o.MinCost == nil {
		var ret float32
		return ret
	}
	return *o.MinCost
}

// GetMinCostOk returns a tuple with the MinCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRangeResponse) GetMinCostOk() (*float32, bool) {
	if o == nil || o.MinCost == nil {
		return nil, false
	}
	return o.MinCost, true
}

// HasMinCost returns a boolean if a field has been set.
func (o *CostRangeResponse) HasMinCost() bool {
	if o != nil && o.MinCost != nil {
		return true
	}

	return false
}

// SetMinCost gets a reference to the given float32 and assigns it to the MinCost field.
func (o *CostRangeResponse) SetMinCost(v float32) {
	o.MinCost = &v
}

// GetMaxCostInCents returns the MaxCostInCents field value if set, zero value otherwise.
func (o *CostRangeResponse) GetMaxCostInCents() int32 {
	if o == nil || o.MaxCostInCents == nil {
		var ret int32
		return ret
	}
	return *o.MaxCostInCents
}

// GetMaxCostInCentsOk returns a tuple with the MaxCostInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRangeResponse) GetMaxCostInCentsOk() (*int32, bool) {
	if o == nil || o.MaxCostInCents == nil {
		return nil, false
	}
	return o.MaxCostInCents, true
}

// HasMaxCostInCents returns a boolean if a field has been set.
func (o *CostRangeResponse) HasMaxCostInCents() bool {
	if o != nil && o.MaxCostInCents != nil {
		return true
	}

	return false
}

// SetMaxCostInCents gets a reference to the given int32 and assigns it to the MaxCostInCents field.
func (o *CostRangeResponse) SetMaxCostInCents(v int32) {
	o.MaxCostInCents = &v
}

// GetMaxCost returns the MaxCost field value if set, zero value otherwise.
func (o *CostRangeResponse) GetMaxCost() float32 {
	if o == nil || o.MaxCost == nil {
		var ret float32
		return ret
	}
	return *o.MaxCost
}

// GetMaxCostOk returns a tuple with the MaxCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRangeResponse) GetMaxCostOk() (*float32, bool) {
	if o == nil || o.MaxCost == nil {
		return nil, false
	}
	return o.MaxCost, true
}

// HasMaxCost returns a boolean if a field has been set.
func (o *CostRangeResponse) HasMaxCost() bool {
	if o != nil && o.MaxCost != nil {
		return true
	}

	return false
}

// SetMaxCost gets a reference to the given float32 and assigns it to the MaxCost field.
func (o *CostRangeResponse) SetMaxCost(v float32) {
	o.MaxCost = &v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *CostRangeResponse) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *CostRangeResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *CostRangeResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

func (o CostRangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinCostInCents != nil {
		toSerialize["min_cost_in_cents"] = o.MinCostInCents
	}
	if o.MinCost != nil {
		toSerialize["min_cost"] = o.MinCost
	}
	if o.MaxCostInCents != nil {
		toSerialize["max_cost_in_cents"] = o.MaxCostInCents
	}
	if o.MaxCost != nil {
		toSerialize["max_cost"] = o.MaxCost
	}
	if true {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableCostRangeResponse struct {
	value *CostRangeResponse
	isSet bool
}

func (v NullableCostRangeResponse) Get() *CostRangeResponse {
	return v.value
}

func (v *NullableCostRangeResponse) Set(val *CostRangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCostRangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCostRangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostRangeResponse(val *CostRangeResponse) *NullableCostRangeResponse {
	return &NullableCostRangeResponse{value: val, isSet: true}
}

func (v NullableCostRangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostRangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


