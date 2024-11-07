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

// checks if the OrganizationEventResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationEventResponseList{}

// OrganizationEventResponseList struct for OrganizationEventResponseList
type OrganizationEventResponseList struct {
	Links *OrganizationEventResponseListLinks `json:"links,omitempty"`
	// Indicates if you cannot see previous logs according to your organization max limit
	OrganizationMaxLimitReached *bool                       `json:"organization_max_limit_reached,omitempty"`
	Events                      []OrganizationEventResponse `json:"events,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _OrganizationEventResponseList OrganizationEventResponseList

// NewOrganizationEventResponseList instantiates a new OrganizationEventResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationEventResponseList() *OrganizationEventResponseList {
	this := OrganizationEventResponseList{}
	return &this
}

// NewOrganizationEventResponseListWithDefaults instantiates a new OrganizationEventResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationEventResponseListWithDefaults() *OrganizationEventResponseList {
	this := OrganizationEventResponseList{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrganizationEventResponseList) GetLinks() OrganizationEventResponseListLinks {
	if o == nil || IsNil(o.Links) {
		var ret OrganizationEventResponseListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponseList) GetLinksOk() (*OrganizationEventResponseListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrganizationEventResponseList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrganizationEventResponseListLinks and assigns it to the Links field.
func (o *OrganizationEventResponseList) SetLinks(v OrganizationEventResponseListLinks) {
	o.Links = &v
}

// GetOrganizationMaxLimitReached returns the OrganizationMaxLimitReached field value if set, zero value otherwise.
func (o *OrganizationEventResponseList) GetOrganizationMaxLimitReached() bool {
	if o == nil || IsNil(o.OrganizationMaxLimitReached) {
		var ret bool
		return ret
	}
	return *o.OrganizationMaxLimitReached
}

// GetOrganizationMaxLimitReachedOk returns a tuple with the OrganizationMaxLimitReached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponseList) GetOrganizationMaxLimitReachedOk() (*bool, bool) {
	if o == nil || IsNil(o.OrganizationMaxLimitReached) {
		return nil, false
	}
	return o.OrganizationMaxLimitReached, true
}

// HasOrganizationMaxLimitReached returns a boolean if a field has been set.
func (o *OrganizationEventResponseList) HasOrganizationMaxLimitReached() bool {
	if o != nil && !IsNil(o.OrganizationMaxLimitReached) {
		return true
	}

	return false
}

// SetOrganizationMaxLimitReached gets a reference to the given bool and assigns it to the OrganizationMaxLimitReached field.
func (o *OrganizationEventResponseList) SetOrganizationMaxLimitReached(v bool) {
	o.OrganizationMaxLimitReached = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrganizationEventResponseList) GetEvents() []OrganizationEventResponse {
	if o == nil || IsNil(o.Events) {
		var ret []OrganizationEventResponse
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponseList) GetEventsOk() ([]OrganizationEventResponse, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrganizationEventResponseList) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []OrganizationEventResponse and assigns it to the Events field.
func (o *OrganizationEventResponseList) SetEvents(v []OrganizationEventResponse) {
	o.Events = v
}

func (o OrganizationEventResponseList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationEventResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.OrganizationMaxLimitReached) {
		toSerialize["organization_max_limit_reached"] = o.OrganizationMaxLimitReached
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationEventResponseList) UnmarshalJSON(data []byte) (err error) {
	varOrganizationEventResponseList := _OrganizationEventResponseList{}

	err = json.Unmarshal(data, &varOrganizationEventResponseList)

	if err != nil {
		return err
	}

	*o = OrganizationEventResponseList(varOrganizationEventResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "organization_max_limit_reached")
		delete(additionalProperties, "events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationEventResponseList struct {
	value *OrganizationEventResponseList
	isSet bool
}

func (v NullableOrganizationEventResponseList) Get() *OrganizationEventResponseList {
	return v.value
}

func (v *NullableOrganizationEventResponseList) Set(val *OrganizationEventResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEventResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEventResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEventResponseList(val *OrganizationEventResponseList) *NullableOrganizationEventResponseList {
	return &NullableOrganizationEventResponseList{value: val, isSet: true}
}

func (v NullableOrganizationEventResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEventResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
