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

// PoolCreate0Topology struct for PoolCreate0Topology
type PoolCreate0Topology struct {
	Data *[]map[string]interface{} `json:"data,omitempty"`
	Special *[]map[string]interface{} `json:"special,omitempty"`
	Dedup *[]map[string]interface{} `json:"dedup,omitempty"`
	Cache *[]map[string]interface{} `json:"cache,omitempty"`
	Log *[]map[string]interface{} `json:"log,omitempty"`
	Spares *[]string `json:"spares,omitempty"`
}

// NewPoolCreate0Topology instantiates a new PoolCreate0Topology object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolCreate0Topology() *PoolCreate0Topology {
	this := PoolCreate0Topology{}
	return &this
}

// NewPoolCreate0TopologyWithDefaults instantiates a new PoolCreate0Topology object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolCreate0TopologyWithDefaults() *PoolCreate0Topology {
	this := PoolCreate0Topology{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PoolCreate0Topology) GetData() []map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0Topology) GetDataOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PoolCreate0Topology) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []map[string]interface{} and assigns it to the Data field.
func (o *PoolCreate0Topology) SetData(v []map[string]interface{}) {
	o.Data = &v
}

// GetSpecial returns the Special field value if set, zero value otherwise.
func (o *PoolCreate0Topology) GetSpecial() []map[string]interface{} {
	if o == nil || o.Special == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Special
}

// GetSpecialOk returns a tuple with the Special field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0Topology) GetSpecialOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Special == nil {
		return nil, false
	}
	return o.Special, true
}

// HasSpecial returns a boolean if a field has been set.
func (o *PoolCreate0Topology) HasSpecial() bool {
	if o != nil && o.Special != nil {
		return true
	}

	return false
}

// SetSpecial gets a reference to the given []map[string]interface{} and assigns it to the Special field.
func (o *PoolCreate0Topology) SetSpecial(v []map[string]interface{}) {
	o.Special = &v
}

// GetDedup returns the Dedup field value if set, zero value otherwise.
func (o *PoolCreate0Topology) GetDedup() []map[string]interface{} {
	if o == nil || o.Dedup == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Dedup
}

// GetDedupOk returns a tuple with the Dedup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0Topology) GetDedupOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Dedup == nil {
		return nil, false
	}
	return o.Dedup, true
}

// HasDedup returns a boolean if a field has been set.
func (o *PoolCreate0Topology) HasDedup() bool {
	if o != nil && o.Dedup != nil {
		return true
	}

	return false
}

// SetDedup gets a reference to the given []map[string]interface{} and assigns it to the Dedup field.
func (o *PoolCreate0Topology) SetDedup(v []map[string]interface{}) {
	o.Dedup = &v
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *PoolCreate0Topology) GetCache() []map[string]interface{} {
	if o == nil || o.Cache == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0Topology) GetCacheOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Cache == nil {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *PoolCreate0Topology) HasCache() bool {
	if o != nil && o.Cache != nil {
		return true
	}

	return false
}

// SetCache gets a reference to the given []map[string]interface{} and assigns it to the Cache field.
func (o *PoolCreate0Topology) SetCache(v []map[string]interface{}) {
	o.Cache = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *PoolCreate0Topology) GetLog() []map[string]interface{} {
	if o == nil || o.Log == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0Topology) GetLogOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Log == nil {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *PoolCreate0Topology) HasLog() bool {
	if o != nil && o.Log != nil {
		return true
	}

	return false
}

// SetLog gets a reference to the given []map[string]interface{} and assigns it to the Log field.
func (o *PoolCreate0Topology) SetLog(v []map[string]interface{}) {
	o.Log = &v
}

// GetSpares returns the Spares field value if set, zero value otherwise.
func (o *PoolCreate0Topology) GetSpares() []string {
	if o == nil || o.Spares == nil {
		var ret []string
		return ret
	}
	return *o.Spares
}

// GetSparesOk returns a tuple with the Spares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolCreate0Topology) GetSparesOk() (*[]string, bool) {
	if o == nil || o.Spares == nil {
		return nil, false
	}
	return o.Spares, true
}

// HasSpares returns a boolean if a field has been set.
func (o *PoolCreate0Topology) HasSpares() bool {
	if o != nil && o.Spares != nil {
		return true
	}

	return false
}

// SetSpares gets a reference to the given []string and assigns it to the Spares field.
func (o *PoolCreate0Topology) SetSpares(v []string) {
	o.Spares = &v
}

func (o PoolCreate0Topology) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Special != nil {
		toSerialize["special"] = o.Special
	}
	if o.Dedup != nil {
		toSerialize["dedup"] = o.Dedup
	}
	if o.Cache != nil {
		toSerialize["cache"] = o.Cache
	}
	if o.Log != nil {
		toSerialize["log"] = o.Log
	}
	if o.Spares != nil {
		toSerialize["spares"] = o.Spares
	}
	return json.Marshal(toSerialize)
}

type NullablePoolCreate0Topology struct {
	value *PoolCreate0Topology
	isSet bool
}

func (v NullablePoolCreate0Topology) Get() *PoolCreate0Topology {
	return v.value
}

func (v *NullablePoolCreate0Topology) Set(val *PoolCreate0Topology) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolCreate0Topology) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolCreate0Topology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolCreate0Topology(val *PoolCreate0Topology) *NullablePoolCreate0Topology {
	return &NullablePoolCreate0Topology{value: val, isSet: true}
}

func (v NullablePoolCreate0Topology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolCreate0Topology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

