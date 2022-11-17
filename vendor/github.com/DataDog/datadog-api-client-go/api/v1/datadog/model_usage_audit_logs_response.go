// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"encoding/json"
)

// UsageAuditLogsResponse Response containing the audit logs usage for each hour for a given organization.
type UsageAuditLogsResponse struct {
	// Get hourly usage for audit logs.
	Usage []UsageAuditLogsHour `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewUsageAuditLogsResponse instantiates a new UsageAuditLogsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageAuditLogsResponse() *UsageAuditLogsResponse {
	this := UsageAuditLogsResponse{}
	return &this
}

// NewUsageAuditLogsResponseWithDefaults instantiates a new UsageAuditLogsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageAuditLogsResponseWithDefaults() *UsageAuditLogsResponse {
	this := UsageAuditLogsResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageAuditLogsResponse) GetUsage() []UsageAuditLogsHour {
	if o == nil || o.Usage == nil {
		var ret []UsageAuditLogsHour
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAuditLogsResponse) GetUsageOk() (*[]UsageAuditLogsHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageAuditLogsResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageAuditLogsHour and assigns it to the Usage field.
func (o *UsageAuditLogsResponse) SetUsage(v []UsageAuditLogsHour) {
	o.Usage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageAuditLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageAuditLogsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Usage []UsageAuditLogsHour `json:"usage,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Usage = all.Usage
	return nil
}