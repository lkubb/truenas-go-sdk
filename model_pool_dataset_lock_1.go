/*
 * TrueNAS RESTful API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// PoolDatasetLock1 struct for PoolDatasetLock1
type PoolDatasetLock1 struct {
	ForceUmount *bool `json:"force_umount,omitempty"`
}

// NewPoolDatasetLock1 instantiates a new PoolDatasetLock1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDatasetLock1() *PoolDatasetLock1 {
	this := PoolDatasetLock1{}
	return &this
}

// NewPoolDatasetLock1WithDefaults instantiates a new PoolDatasetLock1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDatasetLock1WithDefaults() *PoolDatasetLock1 {
	this := PoolDatasetLock1{}
	return &this
}

// GetForceUmount returns the ForceUmount field value if set, zero value otherwise.
func (o *PoolDatasetLock1) GetForceUmount() bool {
	if o == nil || o.ForceUmount == nil {
		var ret bool
		return ret
	}
	return *o.ForceUmount
}

// GetForceUmountOk returns a tuple with the ForceUmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetLock1) GetForceUmountOk() (*bool, bool) {
	if o == nil || o.ForceUmount == nil {
		return nil, false
	}
	return o.ForceUmount, true
}

// HasForceUmount returns a boolean if a field has been set.
func (o *PoolDatasetLock1) HasForceUmount() bool {
	if o != nil && o.ForceUmount != nil {
		return true
	}

	return false
}

// SetForceUmount gets a reference to the given bool and assigns it to the ForceUmount field.
func (o *PoolDatasetLock1) SetForceUmount(v bool) {
	o.ForceUmount = &v
}

func (o PoolDatasetLock1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceUmount != nil {
		toSerialize["force_umount"] = o.ForceUmount
	}
	return json.Marshal(toSerialize)
}

type NullablePoolDatasetLock1 struct {
	value *PoolDatasetLock1
	isSet bool
}

func (v NullablePoolDatasetLock1) Get() *PoolDatasetLock1 {
	return v.value
}

func (v *NullablePoolDatasetLock1) Set(val *PoolDatasetLock1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDatasetLock1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDatasetLock1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDatasetLock1(val *PoolDatasetLock1) *NullablePoolDatasetLock1 {
	return &NullablePoolDatasetLock1{value: val, isSet: true}
}

func (v NullablePoolDatasetLock1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDatasetLock1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

