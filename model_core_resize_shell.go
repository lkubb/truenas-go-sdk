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

// CoreResizeShell struct for CoreResizeShell
type CoreResizeShell struct {
	Id *string `json:"id,omitempty"`
	Cols *int32 `json:"cols,omitempty"`
	Rows *int32 `json:"rows,omitempty"`
}

// NewCoreResizeShell instantiates a new CoreResizeShell object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreResizeShell() *CoreResizeShell {
	this := CoreResizeShell{}
	return &this
}

// NewCoreResizeShellWithDefaults instantiates a new CoreResizeShell object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreResizeShellWithDefaults() *CoreResizeShell {
	this := CoreResizeShell{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreResizeShell) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreResizeShell) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreResizeShell) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoreResizeShell) SetId(v string) {
	o.Id = &v
}

// GetCols returns the Cols field value if set, zero value otherwise.
func (o *CoreResizeShell) GetCols() int32 {
	if o == nil || o.Cols == nil {
		var ret int32
		return ret
	}
	return *o.Cols
}

// GetColsOk returns a tuple with the Cols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreResizeShell) GetColsOk() (*int32, bool) {
	if o == nil || o.Cols == nil {
		return nil, false
	}
	return o.Cols, true
}

// HasCols returns a boolean if a field has been set.
func (o *CoreResizeShell) HasCols() bool {
	if o != nil && o.Cols != nil {
		return true
	}

	return false
}

// SetCols gets a reference to the given int32 and assigns it to the Cols field.
func (o *CoreResizeShell) SetCols(v int32) {
	o.Cols = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *CoreResizeShell) GetRows() int32 {
	if o == nil || o.Rows == nil {
		var ret int32
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreResizeShell) GetRowsOk() (*int32, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *CoreResizeShell) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given int32 and assigns it to the Rows field.
func (o *CoreResizeShell) SetRows(v int32) {
	o.Rows = &v
}

func (o CoreResizeShell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Cols != nil {
		toSerialize["cols"] = o.Cols
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	return json.Marshal(toSerialize)
}

type NullableCoreResizeShell struct {
	value *CoreResizeShell
	isSet bool
}

func (v NullableCoreResizeShell) Get() *CoreResizeShell {
	return v.value
}

func (v *NullableCoreResizeShell) Set(val *CoreResizeShell) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreResizeShell) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreResizeShell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreResizeShell(val *CoreResizeShell) *NullableCoreResizeShell {
	return &NullableCoreResizeShell{value: val, isSet: true}
}

func (v NullableCoreResizeShell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreResizeShell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

