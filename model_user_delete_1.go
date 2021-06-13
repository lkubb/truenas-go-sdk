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

// UserDelete1 struct for UserDelete1
type UserDelete1 struct {
	DeleteGroup *bool `json:"delete_group,omitempty"`
}

// NewUserDelete1 instantiates a new UserDelete1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDelete1() *UserDelete1 {
	this := UserDelete1{}
	return &this
}

// NewUserDelete1WithDefaults instantiates a new UserDelete1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDelete1WithDefaults() *UserDelete1 {
	this := UserDelete1{}
	return &this
}

// GetDeleteGroup returns the DeleteGroup field value if set, zero value otherwise.
func (o *UserDelete1) GetDeleteGroup() bool {
	if o == nil || o.DeleteGroup == nil {
		var ret bool
		return ret
	}
	return *o.DeleteGroup
}

// GetDeleteGroupOk returns a tuple with the DeleteGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDelete1) GetDeleteGroupOk() (*bool, bool) {
	if o == nil || o.DeleteGroup == nil {
		return nil, false
	}
	return o.DeleteGroup, true
}

// HasDeleteGroup returns a boolean if a field has been set.
func (o *UserDelete1) HasDeleteGroup() bool {
	if o != nil && o.DeleteGroup != nil {
		return true
	}

	return false
}

// SetDeleteGroup gets a reference to the given bool and assigns it to the DeleteGroup field.
func (o *UserDelete1) SetDeleteGroup(v bool) {
	o.DeleteGroup = &v
}

func (o UserDelete1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteGroup != nil {
		toSerialize["delete_group"] = o.DeleteGroup
	}
	return json.Marshal(toSerialize)
}

type NullableUserDelete1 struct {
	value *UserDelete1
	isSet bool
}

func (v NullableUserDelete1) Get() *UserDelete1 {
	return v.value
}

func (v *NullableUserDelete1) Set(val *UserDelete1) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDelete1) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDelete1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDelete1(val *UserDelete1) *NullableUserDelete1 {
	return &NullableUserDelete1{value: val, isSet: true}
}

func (v NullableUserDelete1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDelete1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

