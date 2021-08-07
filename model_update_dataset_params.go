/*
 * TrueNAS RESTful API
 *
 * Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// UpdateDatasetParams struct for UpdateDatasetParams
type UpdateDatasetParams struct {
	Atime                *string `json:"atime,omitempty"`
	Aclmode              *string `json:"aclmode,omitempty"`
	Comments             *string `json:"comments,omitempty"`
	Compression          *string `json:"compression,omitempty"`
	Copies               *int32  `json:"copies,omitempty"`
	Deduplication        *string `json:"deduplication,omitempty"`
	Exec                 *string `json:"exec,omitempty"`
	Quota                *int32  `json:"quota,omitempty"`
	Readonly             *string `json:"readonly,omitempty"`
	Recordsize           *string `json:"recordsize,omitempty"`
	Refquota             *int32  `json:"refquota,omitempty"`
	Refreservation       *int32  `json:"refreservation,omitempty"`
	Snapdir              *string `json:"snapdir,omitempty"`
	Sync                 *string `json:"sync,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateDatasetParams UpdateDatasetParams

// NewUpdateDatasetParams instantiates a new UpdateDatasetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDatasetParams() *UpdateDatasetParams {
	this := UpdateDatasetParams{}
	return &this
}

// NewUpdateDatasetParamsWithDefaults instantiates a new UpdateDatasetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDatasetParamsWithDefaults() *UpdateDatasetParams {
	this := UpdateDatasetParams{}
	return &this
}

// GetAtime returns the Atime field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetAtime() string {
	if o == nil || o.Atime == nil {
		var ret string
		return ret
	}
	return *o.Atime
}

// GetAtimeOk returns a tuple with the Atime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetAtimeOk() (*string, bool) {
	if o == nil || o.Atime == nil {
		return nil, false
	}
	return o.Atime, true
}

// HasAtime returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasAtime() bool {
	if o != nil && o.Atime != nil {
		return true
	}

	return false
}

// SetAtime gets a reference to the given string and assigns it to the Atime field.
func (o *UpdateDatasetParams) SetAtime(v string) {
	o.Atime = &v
}

// GetAclmode returns the Aclmode field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetAclmode() string {
	if o == nil || o.Aclmode == nil {
		var ret string
		return ret
	}
	return *o.Aclmode
}

// GetAclmodeOk returns a tuple with the Aclmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetAclmodeOk() (*string, bool) {
	if o == nil || o.Aclmode == nil {
		return nil, false
	}
	return o.Aclmode, true
}

// HasAclmode returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasAclmode() bool {
	if o != nil && o.Aclmode != nil {
		return true
	}

	return false
}

// SetAclmode gets a reference to the given string and assigns it to the Aclmode field.
func (o *UpdateDatasetParams) SetAclmode(v string) {
	o.Aclmode = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *UpdateDatasetParams) SetComments(v string) {
	o.Comments = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetCompression() string {
	if o == nil || o.Compression == nil {
		var ret string
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetCompressionOk() (*string, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasCompression() bool {
	if o != nil && o.Compression != nil {
		return true
	}

	return false
}

// SetCompression gets a reference to the given string and assigns it to the Compression field.
func (o *UpdateDatasetParams) SetCompression(v string) {
	o.Compression = &v
}

// GetCopies returns the Copies field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetCopies() int32 {
	if o == nil || o.Copies == nil {
		var ret int32
		return ret
	}
	return *o.Copies
}

// GetCopiesOk returns a tuple with the Copies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetCopiesOk() (*int32, bool) {
	if o == nil || o.Copies == nil {
		return nil, false
	}
	return o.Copies, true
}

// HasCopies returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasCopies() bool {
	if o != nil && o.Copies != nil {
		return true
	}

	return false
}

// SetCopies gets a reference to the given int32 and assigns it to the Copies field.
func (o *UpdateDatasetParams) SetCopies(v int32) {
	o.Copies = &v
}

// GetDeduplication returns the Deduplication field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetDeduplication() string {
	if o == nil || o.Deduplication == nil {
		var ret string
		return ret
	}
	return *o.Deduplication
}

// GetDeduplicationOk returns a tuple with the Deduplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetDeduplicationOk() (*string, bool) {
	if o == nil || o.Deduplication == nil {
		return nil, false
	}
	return o.Deduplication, true
}

// HasDeduplication returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasDeduplication() bool {
	if o != nil && o.Deduplication != nil {
		return true
	}

	return false
}

// SetDeduplication gets a reference to the given string and assigns it to the Deduplication field.
func (o *UpdateDatasetParams) SetDeduplication(v string) {
	o.Deduplication = &v
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetExec() string {
	if o == nil || o.Exec == nil {
		var ret string
		return ret
	}
	return *o.Exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetExecOk() (*string, bool) {
	if o == nil || o.Exec == nil {
		return nil, false
	}
	return o.Exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasExec() bool {
	if o != nil && o.Exec != nil {
		return true
	}

	return false
}

// SetExec gets a reference to the given string and assigns it to the Exec field.
func (o *UpdateDatasetParams) SetExec(v string) {
	o.Exec = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetQuota() int32 {
	if o == nil || o.Quota == nil {
		var ret int32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetQuotaOk() (*int32, bool) {
	if o == nil || o.Quota == nil {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasQuota() bool {
	if o != nil && o.Quota != nil {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int32 and assigns it to the Quota field.
func (o *UpdateDatasetParams) SetQuota(v int32) {
	o.Quota = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetReadonly() string {
	if o == nil || o.Readonly == nil {
		var ret string
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetReadonlyOk() (*string, bool) {
	if o == nil || o.Readonly == nil {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasReadonly() bool {
	if o != nil && o.Readonly != nil {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given string and assigns it to the Readonly field.
func (o *UpdateDatasetParams) SetReadonly(v string) {
	o.Readonly = &v
}

// GetRecordsize returns the Recordsize field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetRecordsize() string {
	if o == nil || o.Recordsize == nil {
		var ret string
		return ret
	}
	return *o.Recordsize
}

// GetRecordsizeOk returns a tuple with the Recordsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetRecordsizeOk() (*string, bool) {
	if o == nil || o.Recordsize == nil {
		return nil, false
	}
	return o.Recordsize, true
}

// HasRecordsize returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasRecordsize() bool {
	if o != nil && o.Recordsize != nil {
		return true
	}

	return false
}

// SetRecordsize gets a reference to the given string and assigns it to the Recordsize field.
func (o *UpdateDatasetParams) SetRecordsize(v string) {
	o.Recordsize = &v
}

// GetRefquota returns the Refquota field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetRefquota() int32 {
	if o == nil || o.Refquota == nil {
		var ret int32
		return ret
	}
	return *o.Refquota
}

// GetRefquotaOk returns a tuple with the Refquota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetRefquotaOk() (*int32, bool) {
	if o == nil || o.Refquota == nil {
		return nil, false
	}
	return o.Refquota, true
}

// HasRefquota returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasRefquota() bool {
	if o != nil && o.Refquota != nil {
		return true
	}

	return false
}

// SetRefquota gets a reference to the given int32 and assigns it to the Refquota field.
func (o *UpdateDatasetParams) SetRefquota(v int32) {
	o.Refquota = &v
}

// GetRefreservation returns the Refreservation field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetRefreservation() int32 {
	if o == nil || o.Refreservation == nil {
		var ret int32
		return ret
	}
	return *o.Refreservation
}

// GetRefreservationOk returns a tuple with the Refreservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetRefreservationOk() (*int32, bool) {
	if o == nil || o.Refreservation == nil {
		return nil, false
	}
	return o.Refreservation, true
}

// HasRefreservation returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasRefreservation() bool {
	if o != nil && o.Refreservation != nil {
		return true
	}

	return false
}

// SetRefreservation gets a reference to the given int32 and assigns it to the Refreservation field.
func (o *UpdateDatasetParams) SetRefreservation(v int32) {
	o.Refreservation = &v
}

// GetSnapdir returns the Snapdir field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetSnapdir() string {
	if o == nil || o.Snapdir == nil {
		var ret string
		return ret
	}
	return *o.Snapdir
}

// GetSnapdirOk returns a tuple with the Snapdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetSnapdirOk() (*string, bool) {
	if o == nil || o.Snapdir == nil {
		return nil, false
	}
	return o.Snapdir, true
}

// HasSnapdir returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasSnapdir() bool {
	if o != nil && o.Snapdir != nil {
		return true
	}

	return false
}

// SetSnapdir gets a reference to the given string and assigns it to the Snapdir field.
func (o *UpdateDatasetParams) SetSnapdir(v string) {
	o.Snapdir = &v
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *UpdateDatasetParams) GetSync() string {
	if o == nil || o.Sync == nil {
		var ret string
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatasetParams) GetSyncOk() (*string, bool) {
	if o == nil || o.Sync == nil {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *UpdateDatasetParams) HasSync() bool {
	if o != nil && o.Sync != nil {
		return true
	}

	return false
}

// SetSync gets a reference to the given string and assigns it to the Sync field.
func (o *UpdateDatasetParams) SetSync(v string) {
	o.Sync = &v
}

func (o UpdateDatasetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Atime != nil {
		toSerialize["atime"] = o.Atime
	}
	if o.Aclmode != nil {
		toSerialize["aclmode"] = o.Aclmode
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	if o.Copies != nil {
		toSerialize["copies"] = o.Copies
	}
	if o.Deduplication != nil {
		toSerialize["deduplication"] = o.Deduplication
	}
	if o.Exec != nil {
		toSerialize["exec"] = o.Exec
	}
	if o.Quota != nil {
		toSerialize["quota"] = o.Quota
	}
	if o.Readonly != nil {
		toSerialize["readonly"] = o.Readonly
	}
	if o.Recordsize != nil {
		toSerialize["recordsize"] = o.Recordsize
	}
	if o.Refquota != nil {
		toSerialize["refquota"] = o.Refquota
	}
	if o.Refreservation != nil {
		toSerialize["refreservation"] = o.Refreservation
	}
	if o.Snapdir != nil {
		toSerialize["snapdir"] = o.Snapdir
	}
	if o.Sync != nil {
		toSerialize["sync"] = o.Sync
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateDatasetParams) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateDatasetParams := _UpdateDatasetParams{}

	if err = json.Unmarshal(bytes, &varUpdateDatasetParams); err == nil {
		*o = UpdateDatasetParams(varUpdateDatasetParams)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "atime")
		delete(additionalProperties, "aclmode")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "compression")
		delete(additionalProperties, "copies")
		delete(additionalProperties, "deduplication")
		delete(additionalProperties, "exec")
		delete(additionalProperties, "quota")
		delete(additionalProperties, "readonly")
		delete(additionalProperties, "recordsize")
		delete(additionalProperties, "refquota")
		delete(additionalProperties, "refreservation")
		delete(additionalProperties, "snapdir")
		delete(additionalProperties, "sync")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateDatasetParams struct {
	value *UpdateDatasetParams
	isSet bool
}

func (v NullableUpdateDatasetParams) Get() *UpdateDatasetParams {
	return v.value
}

func (v *NullableUpdateDatasetParams) Set(val *UpdateDatasetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDatasetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDatasetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDatasetParams(val *UpdateDatasetParams) *NullableUpdateDatasetParams {
	return &NullableUpdateDatasetParams{value: val, isSet: true}
}

func (v NullableUpdateDatasetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDatasetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
