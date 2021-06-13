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

// ServiceStop1 struct for ServiceStop1
type ServiceStop1 struct {
	HaPropagate *bool `json:"ha_propagate,omitempty"`
}

// NewServiceStop1 instantiates a new ServiceStop1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStop1() *ServiceStop1 {
	this := ServiceStop1{}
	return &this
}

// NewServiceStop1WithDefaults instantiates a new ServiceStop1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStop1WithDefaults() *ServiceStop1 {
	this := ServiceStop1{}
	return &this
}

// GetHaPropagate returns the HaPropagate field value if set, zero value otherwise.
func (o *ServiceStop1) GetHaPropagate() bool {
	if o == nil || o.HaPropagate == nil {
		var ret bool
		return ret
	}
	return *o.HaPropagate
}

// GetHaPropagateOk returns a tuple with the HaPropagate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStop1) GetHaPropagateOk() (*bool, bool) {
	if o == nil || o.HaPropagate == nil {
		return nil, false
	}
	return o.HaPropagate, true
}

// HasHaPropagate returns a boolean if a field has been set.
func (o *ServiceStop1) HasHaPropagate() bool {
	if o != nil && o.HaPropagate != nil {
		return true
	}

	return false
}

// SetHaPropagate gets a reference to the given bool and assigns it to the HaPropagate field.
func (o *ServiceStop1) SetHaPropagate(v bool) {
	o.HaPropagate = &v
}

func (o ServiceStop1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HaPropagate != nil {
		toSerialize["ha_propagate"] = o.HaPropagate
	}
	return json.Marshal(toSerialize)
}

type NullableServiceStop1 struct {
	value *ServiceStop1
	isSet bool
}

func (v NullableServiceStop1) Get() *ServiceStop1 {
	return v.value
}

func (v *NullableServiceStop1) Set(val *ServiceStop1) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStop1) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStop1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStop1(val *ServiceStop1) *NullableServiceStop1 {
	return &NullableServiceStop1{value: val, isSet: true}
}

func (v NullableServiceStop1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStop1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

