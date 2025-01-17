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

// CronJobSchedule struct for CronJobSchedule
type CronJobSchedule struct {
	Minute               *string `json:"minute,omitempty"`
	Hour                 *string `json:"hour,omitempty"`
	Dom                  *string `json:"dom,omitempty"`
	Month                *string `json:"month,omitempty"`
	Dow                  *string `json:"dow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CronJobSchedule CronJobSchedule

// NewCronJobSchedule instantiates a new CronJobSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronJobSchedule() *CronJobSchedule {
	this := CronJobSchedule{}
	return &this
}

// NewCronJobScheduleWithDefaults instantiates a new CronJobSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronJobScheduleWithDefaults() *CronJobSchedule {
	this := CronJobSchedule{}
	return &this
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *CronJobSchedule) GetMinute() string {
	if o == nil || o.Minute == nil {
		var ret string
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobSchedule) GetMinuteOk() (*string, bool) {
	if o == nil || o.Minute == nil {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *CronJobSchedule) HasMinute() bool {
	if o != nil && o.Minute != nil {
		return true
	}

	return false
}

// SetMinute gets a reference to the given string and assigns it to the Minute field.
func (o *CronJobSchedule) SetMinute(v string) {
	o.Minute = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *CronJobSchedule) GetHour() string {
	if o == nil || o.Hour == nil {
		var ret string
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobSchedule) GetHourOk() (*string, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *CronJobSchedule) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given string and assigns it to the Hour field.
func (o *CronJobSchedule) SetHour(v string) {
	o.Hour = &v
}

// GetDom returns the Dom field value if set, zero value otherwise.
func (o *CronJobSchedule) GetDom() string {
	if o == nil || o.Dom == nil {
		var ret string
		return ret
	}
	return *o.Dom
}

// GetDomOk returns a tuple with the Dom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobSchedule) GetDomOk() (*string, bool) {
	if o == nil || o.Dom == nil {
		return nil, false
	}
	return o.Dom, true
}

// HasDom returns a boolean if a field has been set.
func (o *CronJobSchedule) HasDom() bool {
	if o != nil && o.Dom != nil {
		return true
	}

	return false
}

// SetDom gets a reference to the given string and assigns it to the Dom field.
func (o *CronJobSchedule) SetDom(v string) {
	o.Dom = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *CronJobSchedule) GetMonth() string {
	if o == nil || o.Month == nil {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobSchedule) GetMonthOk() (*string, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *CronJobSchedule) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *CronJobSchedule) SetMonth(v string) {
	o.Month = &v
}

// GetDow returns the Dow field value if set, zero value otherwise.
func (o *CronJobSchedule) GetDow() string {
	if o == nil || o.Dow == nil {
		var ret string
		return ret
	}
	return *o.Dow
}

// GetDowOk returns a tuple with the Dow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobSchedule) GetDowOk() (*string, bool) {
	if o == nil || o.Dow == nil {
		return nil, false
	}
	return o.Dow, true
}

// HasDow returns a boolean if a field has been set.
func (o *CronJobSchedule) HasDow() bool {
	if o != nil && o.Dow != nil {
		return true
	}

	return false
}

// SetDow gets a reference to the given string and assigns it to the Dow field.
func (o *CronJobSchedule) SetDow(v string) {
	o.Dow = &v
}

func (o CronJobSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Minute != nil {
		toSerialize["minute"] = o.Minute
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.Dom != nil {
		toSerialize["dom"] = o.Dom
	}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.Dow != nil {
		toSerialize["dow"] = o.Dow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CronJobSchedule) UnmarshalJSON(bytes []byte) (err error) {
	varCronJobSchedule := _CronJobSchedule{}

	if err = json.Unmarshal(bytes, &varCronJobSchedule); err == nil {
		*o = CronJobSchedule(varCronJobSchedule)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "minute")
		delete(additionalProperties, "hour")
		delete(additionalProperties, "dom")
		delete(additionalProperties, "month")
		delete(additionalProperties, "dow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCronJobSchedule struct {
	value *CronJobSchedule
	isSet bool
}

func (v NullableCronJobSchedule) Get() *CronJobSchedule {
	return v.value
}

func (v *NullableCronJobSchedule) Set(val *CronJobSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableCronJobSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableCronJobSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronJobSchedule(val *CronJobSchedule) *NullableCronJobSchedule {
	return &NullableCronJobSchedule{value: val, isSet: true}
}

func (v NullableCronJobSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronJobSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
