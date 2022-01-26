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
	"time"
)

// ServiceResponse struct for ServiceResponse
type ServiceResponse struct {
	// type of the service (application, database, job, gateway...)
	Type *string `json:"type,omitempty"`
	// name of the service
	Name *string `json:"name,omitempty"`
	// uuid of the associated service (application, database, job, gateway...)
	Id *string `json:"id,omitempty"`
	// Git commit ID corresponding to the deployed version of the application
	DeployedCommitId *string `json:"deployed_commit_id,omitempty"`
	// uuid of the user that made the last update
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// global overview of resources consumption of the service
	ConsumedResourcesInPercent *float32 `json:"consumed_resources_in_percent,omitempty"`
	// describes the typology of service (container, postgresl, redis...)
	ServiceTypology *string `json:"service_typology,omitempty"`
	// for databases this field exposes the database version
	ServiceVersion *string `json:"service_version,omitempty"`
	ToUpdate *bool `json:"to_update,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewServiceResponse instantiates a new ServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceResponse(id string, createdAt time.Time) *ServiceResponse {
	this := ServiceResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewServiceResponseWithDefaults instantiates a new ServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceResponseWithDefaults() *ServiceResponse {
	this := ServiceResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceResponse) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceResponse) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceResponse) SetId(v string) {
	o.Id = &v
}

// GetDeployedCommitId returns the DeployedCommitId field value if set, zero value otherwise.
func (o *ServiceResponse) GetDeployedCommitId() string {
	if o == nil || o.DeployedCommitId == nil {
		var ret string
		return ret
	}
	return *o.DeployedCommitId
}

// GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetDeployedCommitIdOk() (*string, bool) {
	if o == nil || o.DeployedCommitId == nil {
		return nil, false
	}
	return o.DeployedCommitId, true
}

// HasDeployedCommitId returns a boolean if a field has been set.
func (o *ServiceResponse) HasDeployedCommitId() bool {
	if o != nil && o.DeployedCommitId != nil {
		return true
	}

	return false
}

// SetDeployedCommitId gets a reference to the given string and assigns it to the DeployedCommitId field.
func (o *ServiceResponse) SetDeployedCommitId(v string) {
	o.DeployedCommitId = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *ServiceResponse) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *ServiceResponse) HasLastUpdatedBy() bool {
	if o != nil && o.LastUpdatedBy != nil {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *ServiceResponse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetConsumedResourcesInPercent returns the ConsumedResourcesInPercent field value if set, zero value otherwise.
func (o *ServiceResponse) GetConsumedResourcesInPercent() float32 {
	if o == nil || o.ConsumedResourcesInPercent == nil {
		var ret float32
		return ret
	}
	return *o.ConsumedResourcesInPercent
}

// GetConsumedResourcesInPercentOk returns a tuple with the ConsumedResourcesInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetConsumedResourcesInPercentOk() (*float32, bool) {
	if o == nil || o.ConsumedResourcesInPercent == nil {
		return nil, false
	}
	return o.ConsumedResourcesInPercent, true
}

// HasConsumedResourcesInPercent returns a boolean if a field has been set.
func (o *ServiceResponse) HasConsumedResourcesInPercent() bool {
	if o != nil && o.ConsumedResourcesInPercent != nil {
		return true
	}

	return false
}

// SetConsumedResourcesInPercent gets a reference to the given float32 and assigns it to the ConsumedResourcesInPercent field.
func (o *ServiceResponse) SetConsumedResourcesInPercent(v float32) {
	o.ConsumedResourcesInPercent = &v
}

// GetServiceTypology returns the ServiceTypology field value if set, zero value otherwise.
func (o *ServiceResponse) GetServiceTypology() string {
	if o == nil || o.ServiceTypology == nil {
		var ret string
		return ret
	}
	return *o.ServiceTypology
}

// GetServiceTypologyOk returns a tuple with the ServiceTypology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetServiceTypologyOk() (*string, bool) {
	if o == nil || o.ServiceTypology == nil {
		return nil, false
	}
	return o.ServiceTypology, true
}

// HasServiceTypology returns a boolean if a field has been set.
func (o *ServiceResponse) HasServiceTypology() bool {
	if o != nil && o.ServiceTypology != nil {
		return true
	}

	return false
}

// SetServiceTypology gets a reference to the given string and assigns it to the ServiceTypology field.
func (o *ServiceResponse) SetServiceTypology(v string) {
	o.ServiceTypology = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *ServiceResponse) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *ServiceResponse) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *ServiceResponse) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetToUpdate returns the ToUpdate field value if set, zero value otherwise.
func (o *ServiceResponse) GetToUpdate() bool {
	if o == nil || o.ToUpdate == nil {
		var ret bool
		return ret
	}
	return *o.ToUpdate
}

// GetToUpdateOk returns a tuple with the ToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetToUpdateOk() (*bool, bool) {
	if o == nil || o.ToUpdate == nil {
		return nil, false
	}
	return o.ToUpdate, true
}

// HasToUpdate returns a boolean if a field has been set.
func (o *ServiceResponse) HasToUpdate() bool {
	if o != nil && o.ToUpdate != nil {
		return true
	}

	return false
}

// SetToUpdate gets a reference to the given bool and assigns it to the ToUpdate field.
func (o *ServiceResponse) SetToUpdate(v bool) {
	o.ToUpdate = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServiceResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServiceResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ServiceResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ServiceResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeployedCommitId != nil {
		toSerialize["deployed_commit_id"] = o.DeployedCommitId
	}
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if o.ConsumedResourcesInPercent != nil {
		toSerialize["consumed_resources_in_percent"] = o.ConsumedResourcesInPercent
	}
	if o.ServiceTypology != nil {
		toSerialize["service_typology"] = o.ServiceTypology
	}
	if o.ServiceVersion != nil {
		toSerialize["service_version"] = o.ServiceVersion
	}
	if o.ToUpdate != nil {
		toSerialize["to_update"] = o.ToUpdate
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableServiceResponse struct {
	value *ServiceResponse
	isSet bool
}

func (v NullableServiceResponse) Get() *ServiceResponse {
	return v.value
}

func (v *NullableServiceResponse) Set(val *ServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceResponse(val *ServiceResponse) *NullableServiceResponse {
	return &NullableServiceResponse{value: val, isSet: true}
}

func (v NullableServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


