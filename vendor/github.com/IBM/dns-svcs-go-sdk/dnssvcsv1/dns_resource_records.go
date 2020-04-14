/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package dnssvcsv1 : Operations and models for the DnsSvcsV1 service
package dnssvcsv1

import (
	"fmt"

	common "github.com/IBM/dns-svcs-go-sdk/common"
	"github.com/IBM/go-sdk-core/v3/core"
)

// DnsSvcsV1 : Manage DNS Resource Records
//

// ListResourceRecords : List Resource Records
// List the Resource Records for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) ListResourceRecords(listResourceRecordsOptions *ListResourceRecordsOptions) (result *ListResourceRecords, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listResourceRecordsOptions, "listResourceRecordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listResourceRecordsOptions, "listResourceRecordsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "resource_records"}
	pathParameters := []string{*listResourceRecordsOptions.InstanceID, *listResourceRecordsOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listResourceRecordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListResourceRecords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if listResourceRecordsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listResourceRecordsOptions.XCorrelationID))
	}

	if listResourceRecordsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listResourceRecordsOptions.Offset))
	}
	if listResourceRecordsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listResourceRecordsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalListResourceRecords(m)
		response.Result = result
	}

	return
}

// CreateResourceRecord : Create a resource record
// Create a resource record for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) CreateResourceRecord(createResourceRecordOptions *CreateResourceRecordOptions) (result *ResourceRecord, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createResourceRecordOptions, "createResourceRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createResourceRecordOptions, "createResourceRecordOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "resource_records"}
	pathParameters := []string{*createResourceRecordOptions.InstanceID, *createResourceRecordOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createResourceRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateResourceRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createResourceRecordOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createResourceRecordOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createResourceRecordOptions.Name != nil {
		body["name"] = createResourceRecordOptions.Name
	}
	if createResourceRecordOptions.Protocol != nil {
		body["protocol"] = createResourceRecordOptions.Protocol
	}
	if createResourceRecordOptions.Rdata != nil {
		body["rdata"] = createResourceRecordOptions.Rdata
	}
	if createResourceRecordOptions.Service != nil {
		body["service"] = createResourceRecordOptions.Service
	}
	if createResourceRecordOptions.TTL != nil {
		body["ttl"] = createResourceRecordOptions.TTL
	}
	if createResourceRecordOptions.Type != nil {
		body["type"] = createResourceRecordOptions.Type
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceRecord(m)
		response.Result = result
	}

	return
}

// DeleteResourceRecord : Delete a resource record
// Delete a resource record.
func (dnsSvcs *DnsSvcsV1) DeleteResourceRecord(deleteResourceRecordOptions *DeleteResourceRecordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteResourceRecordOptions, "deleteResourceRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteResourceRecordOptions, "deleteResourceRecordOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "resource_records"}
	pathParameters := []string{*deleteResourceRecordOptions.InstanceID, *deleteResourceRecordOptions.DnszoneID, *deleteResourceRecordOptions.RecordID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteResourceRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteResourceRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if deleteResourceRecordOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteResourceRecordOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetResourceRecord : Get a resource record
// Get details of a resource record.
func (dnsSvcs *DnsSvcsV1) GetResourceRecord(getResourceRecordOptions *GetResourceRecordOptions) (result *ResourceRecord, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceRecordOptions, "getResourceRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceRecordOptions, "getResourceRecordOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "resource_records"}
	pathParameters := []string{*getResourceRecordOptions.InstanceID, *getResourceRecordOptions.DnszoneID, *getResourceRecordOptions.RecordID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetResourceRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if getResourceRecordOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getResourceRecordOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceRecord(m)
		response.Result = result
	}

	return
}

// UpdateResourceRecord : Update the properties of a resource record
// Update the properties of a resource record.
func (dnsSvcs *DnsSvcsV1) UpdateResourceRecord(updateResourceRecordOptions *UpdateResourceRecordOptions) (result *ResourceRecord, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateResourceRecordOptions, "updateResourceRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateResourceRecordOptions, "updateResourceRecordOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "resource_records"}
	pathParameters := []string{*updateResourceRecordOptions.InstanceID, *updateResourceRecordOptions.DnszoneID, *updateResourceRecordOptions.RecordID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateResourceRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateResourceRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateResourceRecordOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateResourceRecordOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateResourceRecordOptions.Name != nil {
		body["name"] = updateResourceRecordOptions.Name
	}
	if updateResourceRecordOptions.Rdata != nil {
		body["rdata"] = updateResourceRecordOptions.Rdata
	}
	if updateResourceRecordOptions.TTL != nil {
		body["ttl"] = updateResourceRecordOptions.TTL
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceRecord(m)
		response.Result = result
	}

	return
}

// CreateResourceRecordOptions : The CreateResourceRecord options.
type CreateResourceRecordOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Name of the resource record.
	Name *string `json:"name,omitempty"`

	// Only used for SRV record.
	Protocol *string `json:"protocol,omitempty"`

	// Content of the resource record.
	Rdata ResourceRecordInputRdataIntf `json:"rdata,omitempty"`

	// Only used for SRV record.
	Service *string `json:"service,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// Type of the resource record.
	Type *string `json:"type,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateResourceRecordOptions.Type property.
// Type of the resource record.
const (
	CreateResourceRecordOptions_Type_A     = "A"
	CreateResourceRecordOptions_Type_Aaaa  = "AAAA"
	CreateResourceRecordOptions_Type_Cname = "CNAME"
	CreateResourceRecordOptions_Type_Mx    = "MX"
	CreateResourceRecordOptions_Type_Ptr   = "PTR"
	CreateResourceRecordOptions_Type_Srv   = "SRV"
	CreateResourceRecordOptions_Type_Txt   = "TXT"
)

// NewCreateResourceRecordOptions : Instantiate CreateResourceRecordOptions
func (*DnsSvcsV1) NewCreateResourceRecordOptions(instanceID string, dnszoneID string) *CreateResourceRecordOptions {
	return &CreateResourceRecordOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreateResourceRecordOptions) SetInstanceID(instanceID string) *CreateResourceRecordOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *CreateResourceRecordOptions) SetDnszoneID(dnszoneID string) *CreateResourceRecordOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateResourceRecordOptions) SetName(name string) *CreateResourceRecordOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetProtocol : Allow user to set Protocol
func (options *CreateResourceRecordOptions) SetProtocol(protocol string) *CreateResourceRecordOptions {
	options.Protocol = core.StringPtr(protocol)
	return options
}

// SetRdata : Allow user to set Rdata
func (options *CreateResourceRecordOptions) SetRdata(rdata ResourceRecordInputRdataIntf) *CreateResourceRecordOptions {
	options.Rdata = rdata
	return options
}

// SetService : Allow user to set Service
func (options *CreateResourceRecordOptions) SetService(service string) *CreateResourceRecordOptions {
	options.Service = core.StringPtr(service)
	return options
}

// SetTTL : Allow user to set TTL
func (options *CreateResourceRecordOptions) SetTTL(ttl int64) *CreateResourceRecordOptions {
	options.TTL = core.Int64Ptr(ttl)
	return options
}

// SetType : Allow user to set Type
func (options *CreateResourceRecordOptions) SetType(typeVar string) *CreateResourceRecordOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreateResourceRecordOptions) SetXCorrelationID(xCorrelationID string) *CreateResourceRecordOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateResourceRecordOptions) SetHeaders(param map[string]string) *CreateResourceRecordOptions {
	options.Headers = param
	return options
}

// DeleteResourceRecordOptions : The DeleteResourceRecord options.
type DeleteResourceRecordOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a resource record.
	RecordID *string `json:"record_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteResourceRecordOptions : Instantiate DeleteResourceRecordOptions
func (*DnsSvcsV1) NewDeleteResourceRecordOptions(instanceID string, dnszoneID string, recordID string) *DeleteResourceRecordOptions {
	return &DeleteResourceRecordOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		RecordID:   core.StringPtr(recordID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteResourceRecordOptions) SetInstanceID(instanceID string) *DeleteResourceRecordOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *DeleteResourceRecordOptions) SetDnszoneID(dnszoneID string) *DeleteResourceRecordOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetRecordID : Allow user to set RecordID
func (options *DeleteResourceRecordOptions) SetRecordID(recordID string) *DeleteResourceRecordOptions {
	options.RecordID = core.StringPtr(recordID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeleteResourceRecordOptions) SetXCorrelationID(xCorrelationID string) *DeleteResourceRecordOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteResourceRecordOptions) SetHeaders(param map[string]string) *DeleteResourceRecordOptions {
	options.Headers = param
	return options
}

// GetResourceRecordOptions : The GetResourceRecord options.
type GetResourceRecordOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a resource record.
	RecordID *string `json:"record_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceRecordOptions : Instantiate GetResourceRecordOptions
func (*DnsSvcsV1) NewGetResourceRecordOptions(instanceID string, dnszoneID string, recordID string) *GetResourceRecordOptions {
	return &GetResourceRecordOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		RecordID:   core.StringPtr(recordID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetResourceRecordOptions) SetInstanceID(instanceID string) *GetResourceRecordOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *GetResourceRecordOptions) SetDnszoneID(dnszoneID string) *GetResourceRecordOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetRecordID : Allow user to set RecordID
func (options *GetResourceRecordOptions) SetRecordID(recordID string) *GetResourceRecordOptions {
	options.RecordID = core.StringPtr(recordID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetResourceRecordOptions) SetXCorrelationID(xCorrelationID string) *GetResourceRecordOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceRecordOptions) SetHeaders(param map[string]string) *GetResourceRecordOptions {
	options.Headers = param
	return options
}

// ListResourceRecordsOptions : The ListResourceRecords options.
type ListResourceRecordsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Specify how many resource records to skip over, the default value is 0.
	Offset *string `json:"offset,omitempty"`

	// Specify how many resource records are returned, the default value is 20.
	Limit *string `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListResourceRecordsOptions : Instantiate ListResourceRecordsOptions
func (*DnsSvcsV1) NewListResourceRecordsOptions(instanceID string, dnszoneID string) *ListResourceRecordsOptions {
	return &ListResourceRecordsOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListResourceRecordsOptions) SetInstanceID(instanceID string) *ListResourceRecordsOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *ListResourceRecordsOptions) SetDnszoneID(dnszoneID string) *ListResourceRecordsOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListResourceRecordsOptions) SetXCorrelationID(xCorrelationID string) *ListResourceRecordsOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListResourceRecordsOptions) SetOffset(offset string) *ListResourceRecordsOptions {
	options.Offset = core.StringPtr(offset)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListResourceRecordsOptions) SetLimit(limit string) *ListResourceRecordsOptions {
	options.Limit = core.StringPtr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListResourceRecordsOptions) SetHeaders(param map[string]string) *ListResourceRecordsOptions {
	options.Headers = param
	return options
}

// ResourceRecordInputRdata : Content of the resource record.
// Models which "extend" this model:
// - ResourceRecordInputRdataRdataARecord
// - ResourceRecordInputRdataRdataAaaaRecord
// - ResourceRecordInputRdataRdataCnameRecord
// - ResourceRecordInputRdataRdataMxRecord
// - ResourceRecordInputRdataRdataSrvRecord
// - ResourceRecordInputRdataRdataTxtRecord
// - ResourceRecordInputRdataRdataPtrRecord
type ResourceRecordInputRdata struct {
	// IPv4 address.
	Ip *string `json:"ip,omitempty"`

	// Canonical name.
	Cname *string `json:"cname,omitempty"`

	// Hostname of Exchange server.
	Exchange *string `json:"exchange,omitempty"`

	// Preference of the MX record.
	Preference *int64 `json:"preference,omitempty"`

	// Port number of the target server.
	Port *int64 `json:"port,omitempty"`

	// Priority of the SRV record.
	Priority *int64 `json:"priority,omitempty"`

	// Hostname of the target server.
	Target *string `json:"target,omitempty"`

	// Weight of distributing queries among multiple target servers.
	Weight *int64 `json:"weight,omitempty"`

	// Human readable text.
	Txtdata *string `json:"text,omitempty"`

	// Hostname of the relevant A or AAAA record.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

func (*ResourceRecordInputRdata) isaResourceRecordInputRdata() bool {
	return true
}

type ResourceRecordInputRdataIntf interface {
	isaResourceRecordInputRdata() bool
}

// UnmarshalResourceRecordInputRdata constructs an instance of ResourceRecordInputRdata from the specified map.
func UnmarshalResourceRecordInputRdata(m map[string]interface{}) (result ResourceRecordInputRdataIntf, err error) {
	obj := new(ResourceRecordInputRdata)
	obj.Ip, err = core.UnmarshalString(m, "ip")
	if err != nil {
		return
	}
	obj.Cname, err = core.UnmarshalString(m, "cname")
	if err != nil {
		return
	}
	obj.Exchange, err = core.UnmarshalString(m, "exchange")
	if err != nil {
		return
	}
	obj.Preference, err = core.UnmarshalInt64(m, "preference")
	if err != nil {
		return
	}
	obj.Port, err = core.UnmarshalInt64(m, "port")
	if err != nil {
		return
	}
	obj.Priority, err = core.UnmarshalInt64(m, "priority")
	if err != nil {
		return
	}
	obj.Target, err = core.UnmarshalString(m, "target")
	if err != nil {
		return
	}
	obj.Weight, err = core.UnmarshalInt64(m, "weight")
	if err != nil {
		return
	}
	obj.Txtdata, err = core.UnmarshalString(m, "text")
	if err != nil {
		return
	}
	obj.Ptrdname, err = core.UnmarshalString(m, "ptrdname")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordUpdateInputRdata : Content of the resource record.
// Models which "extend" this model:
// - ResourceRecordUpdateInputRdataRdataARecord
// - ResourceRecordUpdateInputRdataRdataAaaaRecord
// - ResourceRecordUpdateInputRdataRdataCnameRecord
// - ResourceRecordUpdateInputRdataRdataMxRecord
// - ResourceRecordUpdateInputRdataRdataSrvRecord
// - ResourceRecordUpdateInputRdataRdataTxtRecord
// - ResourceRecordUpdateInputRdataRdataPtrRecord
type ResourceRecordUpdateInputRdata struct {
	// IPv4 address.
	Ip *string `json:"ip,omitempty"`

	// Canonical name.
	Cname *string `json:"cname,omitempty"`

	// Hostname of Exchange server.
	Exchange *string `json:"exchange,omitempty"`

	// Preference of the MX record.
	Preference *int64 `json:"preference,omitempty"`

	// Port number of the target server.
	Port *int64 `json:"port,omitempty"`

	// Priority of the SRV record.
	Priority *int64 `json:"priority,omitempty"`

	// Hostname of the target server.
	Target *string `json:"target,omitempty"`

	// Weight of distributing queries among multiple target servers.
	Weight *int64 `json:"weight,omitempty"`

	// Human readable text.
	Txtdata *string `json:"text,omitempty"`

	// Hostname of the relevant A or AAAA record.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

func (*ResourceRecordUpdateInputRdata) isaResourceRecordUpdateInputRdata() bool {
	return true
}

type ResourceRecordUpdateInputRdataIntf interface {
	isaResourceRecordUpdateInputRdata() bool
}

// UnmarshalResourceRecordUpdateInputRdata constructs an instance of ResourceRecordUpdateInputRdata from the specified map.
func UnmarshalResourceRecordUpdateInputRdata(m map[string]interface{}) (result ResourceRecordUpdateInputRdataIntf, err error) {
	obj := new(ResourceRecordUpdateInputRdata)
	obj.Ip, err = core.UnmarshalString(m, "ip")
	if err != nil {
		return
	}
	obj.Cname, err = core.UnmarshalString(m, "cname")
	if err != nil {
		return
	}
	obj.Exchange, err = core.UnmarshalString(m, "exchange")
	if err != nil {
		return
	}
	obj.Preference, err = core.UnmarshalInt64(m, "preference")
	if err != nil {
		return
	}
	obj.Port, err = core.UnmarshalInt64(m, "port")
	if err != nil {
		return
	}
	obj.Priority, err = core.UnmarshalInt64(m, "priority")
	if err != nil {
		return
	}
	obj.Target, err = core.UnmarshalString(m, "target")
	if err != nil {
		return
	}
	obj.Weight, err = core.UnmarshalInt64(m, "weight")
	if err != nil {
		return
	}
	obj.Txtdata, err = core.UnmarshalString(m, "text")
	if err != nil {
		return
	}
	obj.Ptrdname, err = core.UnmarshalString(m, "ptrdname")
	if err != nil {
		return
	}
	result = obj
	return
}

// UpdateResourceRecordOptions : The UpdateResourceRecord options.
type UpdateResourceRecordOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a resource record.
	RecordID *string `json:"record_id" validate:"required"`

	// Name of the resource record.
	Name *string `json:"name,omitempty"`

	// Content of the resource record.
	Rdata ResourceRecordUpdateInputRdataIntf `json:"rdata,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// Only used for SRV record.
	Protocol *string `json:"protocol,omitempty"`

	// Only used for SRV record.
	Service *string `json:"service,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateResourceRecordOptions : Instantiate UpdateResourceRecordOptions
func (*DnsSvcsV1) NewUpdateResourceRecordOptions(instanceID string, dnszoneID string, recordID string) *UpdateResourceRecordOptions {
	return &UpdateResourceRecordOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		RecordID:   core.StringPtr(recordID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdateResourceRecordOptions) SetInstanceID(instanceID string) *UpdateResourceRecordOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *UpdateResourceRecordOptions) SetDnszoneID(dnszoneID string) *UpdateResourceRecordOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetRecordID : Allow user to set RecordID
func (options *UpdateResourceRecordOptions) SetRecordID(recordID string) *UpdateResourceRecordOptions {
	options.RecordID = core.StringPtr(recordID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateResourceRecordOptions) SetName(name string) *UpdateResourceRecordOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetRdata : Allow user to set Rdata
func (options *UpdateResourceRecordOptions) SetRdata(rdata ResourceRecordUpdateInputRdataIntf) *UpdateResourceRecordOptions {
	options.Rdata = rdata
	return options
}

// SetTTL : Allow user to set TTL
func (options *UpdateResourceRecordOptions) SetTTL(ttl int64) *UpdateResourceRecordOptions {
	options.TTL = core.Int64Ptr(ttl)
	return options
}

// SetProtocol : Allow user to set Protocol
func (options *UpdateResourceRecordOptions) SetProtocol(protocol string) *UpdateResourceRecordOptions {
	options.Protocol = core.StringPtr(protocol)
	return options
}

// SetService : Allow user to set Service
func (options *UpdateResourceRecordOptions) SetService(service string) *UpdateResourceRecordOptions {
	options.Service = core.StringPtr(service)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *UpdateResourceRecordOptions) SetXCorrelationID(xCorrelationID string) *UpdateResourceRecordOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateResourceRecordOptions) SetHeaders(param map[string]string) *UpdateResourceRecordOptions {
	options.Headers = param
	return options
}

// ListResourceRecords : List Resource Records response.
type ListResourceRecords struct {
	// An array of resource records.
	ResourceRecords []ResourceRecord `json:"resource_records" validate:"required"`

	// Specify how many resource records to skip over.
	Offset *int64 `json:"offset" validate:"required"`

	// Specify how many resource records are returned, the default value is 20.
	Limit *int64 `json:"limit" validate:"required"`

	// Total number of resource records.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next,omitempty"`
}

// UnmarshalListResourceRecords constructs an instance of ListResourceRecords from the specified map.
func UnmarshalListResourceRecords(m map[string]interface{}) (result *ListResourceRecords, err error) {
	obj := new(ListResourceRecords)
	obj.First, err = UnmarshalFirstHrefAsProperty(m, "first")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.Next, err = UnmarshalNextHrefAsProperty(m, "next")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.ResourceRecords, err = UnmarshalResourceRecordSliceAsProperty(m, "resource_records")
	if err != nil {
		return
	}
	obj.TotalCount, err = core.UnmarshalInt64(m, "total_count")
	if err != nil {
		return
	}
	result = obj
	return
}

// NextHref : href.
type NextHref struct {
	// href.
	Href *string `json:"href,omitempty"`
}

// UnmarshalNextHref constructs an instance of NextHref from the specified map.
func UnmarshalNextHref(m map[string]interface{}) (result *NextHref, err error) {
	obj := new(NextHref)
	obj.Href, err = core.UnmarshalString(m, "href")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalNextHrefAsProperty unmarshals an instance of NextHref that is stored as a property
// within the specified map.
func UnmarshalNextHrefAsProperty(m map[string]interface{}, propertyName string) (result *NextHref, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'NextHref'", propertyName)
			return
		}
		result, err = UnmarshalNextHref(objMap)
	}
	return
}

// ResourceRecord : Resource record details.
type ResourceRecord struct {
	// the time when a resource record is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// Identifier of the resource record.
	ID *string `json:"id,omitempty"`

	// the recent time when a resource record is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Name of the resource record.
	Name *string `json:"name,omitempty"`

	// Only used for SRV record.
	Protocol *string `json:"protocol,omitempty"`

	// Content of the resource record.
	Rdata interface{} `json:"rdata,omitempty"`

	// Only used for SRV record.
	Service *string `json:"service,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// Type of the resource record.
	Type *string `json:"type,omitempty"`

	// Linked PTR record in A/AAAA record
	LinkedPTR interface{} `json:"linked_ptr_record,omitempty"`
}

// Constants associated with the ResourceRecord.Type property.
// Type of the resource record.
const (
	ResourceRecord_Type_A     = "A"
	ResourceRecord_Type_Aaaa  = "AAAA"
	ResourceRecord_Type_Cname = "CNAME"
	ResourceRecord_Type_Mx    = "MX"
	ResourceRecord_Type_Ptr   = "PTR"
	ResourceRecord_Type_Srv   = "SRV"
	ResourceRecord_Type_Txt   = "TXT"
)

// UnmarshalResourceRecord constructs an instance of ResourceRecord from the specified map.
func UnmarshalResourceRecord(m map[string]interface{}) (result *ResourceRecord, err error) {
	obj := new(ResourceRecord)
	obj.CreatedOn, err = core.UnmarshalString(m, "created_on")
	if err != nil {
		return
	}
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.ModifiedOn, err = core.UnmarshalString(m, "modified_on")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Protocol, err = core.UnmarshalString(m, "protocol")
	if err != nil {
		return
	}
	obj.Rdata, err = core.UnmarshalObject(m, "rdata")
	if err != nil {
		return
	}
	obj.Service, err = core.UnmarshalString(m, "service")
	if err != nil {
		return
	}
	obj.TTL, err = core.UnmarshalInt64(m, "ttl")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	if m["linked_ptr_record"] != nil {
		obj.LinkedPTR, err = core.UnmarshalObject(m, "linked_ptr_record")
	}
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceRecordSlice unmarshals a slice of ResourceRecord instances from the specified list of maps.
func UnmarshalResourceRecordSlice(s []interface{}) (slice []ResourceRecord, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceRecord'")
			return
		}
		obj, e := UnmarshalResourceRecord(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceRecordSliceAsProperty unmarshals a slice of ResourceRecord instances that are stored as a property
// within the specified map.
func UnmarshalResourceRecordSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceRecord, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceRecord'", propertyName)
			return
		}
		slice, err = UnmarshalResourceRecordSlice(vSlice)
	}
	return
}

// ResourceRecordInputRdataRdataARecord : The content of type-A resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataARecord struct {
	// IPv4 address.
	Ip *string `json:"ip" validate:"required"`
}

// NewResourceRecordInputRdataRdataARecord : Instantiate ResourceRecordInputRdataRdataARecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataARecord(ip string) (model *ResourceRecordInputRdataRdataARecord, err error) {
	model = &ResourceRecordInputRdataRdataARecord{
		Ip: core.StringPtr(ip),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataARecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataARecord constructs an instance of ResourceRecordInputRdataRdataARecord from the specified map.
func UnmarshalResourceRecordInputRdataRdataARecord(m map[string]interface{}) (result *ResourceRecordInputRdataRdataARecord, err error) {
	obj := new(ResourceRecordInputRdataRdataARecord)
	obj.Ip, err = core.UnmarshalString(m, "ip")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordInputRdataRdataAaaaRecord : The content of type-AAAA resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataAaaaRecord struct {
	// IPv6 address.
	Ip *string `json:"ip" validate:"required"`
}

// NewResourceRecordInputRdataRdataAaaaRecord : Instantiate ResourceRecordInputRdataRdataAaaaRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataAaaaRecord(ip string) (model *ResourceRecordInputRdataRdataAaaaRecord, err error) {
	model = &ResourceRecordInputRdataRdataAaaaRecord{
		Ip: core.StringPtr(ip),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataAaaaRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataAaaaRecord constructs an instance of ResourceRecordInputRdataRdataAaaaRecord from the specified map.
func UnmarshalResourceRecordInputRdataRdataAaaaRecord(m map[string]interface{}) (result *ResourceRecordInputRdataRdataAaaaRecord, err error) {
	obj := new(ResourceRecordInputRdataRdataAaaaRecord)
	obj.Ip, err = core.UnmarshalString(m, "ip")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordInputRdataRdataCnameRecord : The content of type-CNAME resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataCnameRecord struct {
	// Canonical name.
	Cname *string `json:"cname" validate:"required"`
}

// NewResourceRecordInputRdataRdataCnameRecord : Instantiate ResourceRecordInputRdataRdataCnameRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataCnameRecord(cname string) (model *ResourceRecordInputRdataRdataCnameRecord, err error) {
	model = &ResourceRecordInputRdataRdataCnameRecord{
		Cname: core.StringPtr(cname),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataCnameRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataCnameRecord constructs an instance of ResourceRecordInputRdataRdataCnameRecord from the specified map.
func UnmarshalResourceRecordInputRdataRdataCnameRecord(m map[string]interface{}) (result *ResourceRecordInputRdataRdataCnameRecord, err error) {
	obj := new(ResourceRecordInputRdataRdataCnameRecord)
	obj.Cname, err = core.UnmarshalString(m, "cname")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordInputRdataRdataMxRecord : The content of type-MX resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataMxRecord struct {
	// Hostname of Exchange server.
	Exchange *string `json:"exchange" validate:"required"`

	// Preference of the MX record.
	Preference *int64 `json:"preference" validate:"required"`
}

// NewResourceRecordInputRdataRdataMxRecord : Instantiate ResourceRecordInputRdataRdataMxRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataMxRecord(exchange string, preference int64) (model *ResourceRecordInputRdataRdataMxRecord, err error) {
	model = &ResourceRecordInputRdataRdataMxRecord{
		Exchange:   core.StringPtr(exchange),
		Preference: core.Int64Ptr(preference),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataMxRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataMxRecord constructs an instance of ResourceRecordInputRdataRdataMxRecord from the specified map.
func UnmarshalResourceRecordInputRdataRdataMxRecord(m map[string]interface{}) (result *ResourceRecordInputRdataRdataMxRecord, err error) {
	obj := new(ResourceRecordInputRdataRdataMxRecord)
	obj.Exchange, err = core.UnmarshalString(m, "exchange")
	if err != nil {
		return
	}
	obj.Preference, err = core.UnmarshalInt64(m, "preference")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordInputRdataRdataPtrRecord : The content of type-PTR resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataPtrRecord struct {
	// Hostname of the relevant A or AAAA record.
	Ptrdname *string `json:"ptrdname" validate:"required"`
}

// NewResourceRecordInputRdataRdataPtrRecord : Instantiate ResourceRecordInputRdataRdataPtrRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataPtrRecord(ptrdname string) (model *ResourceRecordInputRdataRdataPtrRecord, err error) {
	model = &ResourceRecordInputRdataRdataPtrRecord{
		Ptrdname: core.StringPtr(ptrdname),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataPtrRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataPtrRecord constructs an instance of ResourceRecordInputRdataRdataPtrRecord from the specified map.
func UnmarshalResourceRecordInputRdataRdataPtrRecord(m map[string]interface{}) (result *ResourceRecordInputRdataRdataPtrRecord, err error) {
	obj := new(ResourceRecordInputRdataRdataPtrRecord)
	obj.Ptrdname, err = core.UnmarshalString(m, "ptrdname")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordInputRdataRdataSrvRecord : The content of type-SRV resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataSrvRecord struct {
	// Port number of the target server.
	Port *int64 `json:"port" validate:"required"`

	// Priority of the SRV record.
	Priority *int64 `json:"priority" validate:"required"`

	// Hostname of the target server.
	Target *string `json:"target" validate:"required"`

	// Weight of distributing queries among multiple target servers.
	Weight *int64 `json:"weight" validate:"required"`
}

// NewResourceRecordInputRdataRdataSrvRecord : Instantiate ResourceRecordInputRdataRdataSrvRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataSrvRecord(port int64, priority int64, target string, weight int64) (model *ResourceRecordInputRdataRdataSrvRecord, err error) {
	model = &ResourceRecordInputRdataRdataSrvRecord{
		Port:     core.Int64Ptr(port),
		Priority: core.Int64Ptr(priority),
		Target:   core.StringPtr(target),
		Weight:   core.Int64Ptr(weight),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataSrvRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataSrvRecord constructs an instance of ResourceRecordInputRdataRdataSrvRecord from the specified map.
func UnmarshalResourceRecordInputRdataRdataSrvRecord(m map[string]interface{}) (result *ResourceRecordInputRdataRdataSrvRecord, err error) {
	obj := new(ResourceRecordInputRdataRdataSrvRecord)
	obj.Port, err = core.UnmarshalInt64(m, "port")
	if err != nil {
		return
	}
	obj.Priority, err = core.UnmarshalInt64(m, "priority")
	if err != nil {
		return
	}
	obj.Target, err = core.UnmarshalString(m, "target")
	if err != nil {
		return
	}
	obj.Weight, err = core.UnmarshalInt64(m, "weight")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordInputRdataRdataTxtRecord : The content of type-TXT resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataTxtRecord struct {
	// Human readable text.
	Txtdata *string `json:"text" validate:"required"`
}

// NewResourceRecordInputRdataRdataTxtRecord : Instantiate ResourceRecordInputRdataRdataTxtRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataTxtRecord(txtdata string) (model *ResourceRecordInputRdataRdataTxtRecord, err error) {
	model = &ResourceRecordInputRdataRdataTxtRecord{
		Txtdata: core.StringPtr(txtdata),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataTxtRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataTxtRecord constructs an instance of ResourceRecordInputRdataRdataTxtRecord from the specified map.
func UnmarshalResourceRecordInputRdataRdataTxtRecord(m map[string]interface{}) (result *ResourceRecordInputRdataRdataTxtRecord, err error) {
	obj := new(ResourceRecordInputRdataRdataTxtRecord)
	obj.Txtdata, err = core.UnmarshalString(m, "text")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordUpdateInputRdataRdataARecord : The content of type-A resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataARecord struct {
	// IPv4 address.
	Ip *string `json:"ip" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataARecord : Instantiate ResourceRecordUpdateInputRdataRdataARecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataARecord(ip string) (model *ResourceRecordUpdateInputRdataRdataARecord, err error) {
	model = &ResourceRecordUpdateInputRdataRdataARecord{
		Ip: core.StringPtr(ip),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataARecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataARecord constructs an instance of ResourceRecordUpdateInputRdataRdataARecord from the specified map.
func UnmarshalResourceRecordUpdateInputRdataRdataARecord(m map[string]interface{}) (result *ResourceRecordUpdateInputRdataRdataARecord, err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataARecord)
	obj.Ip, err = core.UnmarshalString(m, "ip")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordUpdateInputRdataRdataAaaaRecord : The content of type-AAAA resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataAaaaRecord struct {
	// IPv6 address.
	Ip *string `json:"ip" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataAaaaRecord : Instantiate ResourceRecordUpdateInputRdataRdataAaaaRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataAaaaRecord(ip string) (model *ResourceRecordUpdateInputRdataRdataAaaaRecord, err error) {
	model = &ResourceRecordUpdateInputRdataRdataAaaaRecord{
		Ip: core.StringPtr(ip),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataAaaaRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataAaaaRecord constructs an instance of ResourceRecordUpdateInputRdataRdataAaaaRecord from the specified map.
func UnmarshalResourceRecordUpdateInputRdataRdataAaaaRecord(m map[string]interface{}) (result *ResourceRecordUpdateInputRdataRdataAaaaRecord, err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataAaaaRecord)
	obj.Ip, err = core.UnmarshalString(m, "ip")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordUpdateInputRdataRdataCnameRecord : The content of type-CNAME resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataCnameRecord struct {
	// Canonical name.
	Cname *string `json:"cname" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataCnameRecord : Instantiate ResourceRecordUpdateInputRdataRdataCnameRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataCnameRecord(cname string) (model *ResourceRecordUpdateInputRdataRdataCnameRecord, err error) {
	model = &ResourceRecordUpdateInputRdataRdataCnameRecord{
		Cname: core.StringPtr(cname),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataCnameRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataCnameRecord constructs an instance of ResourceRecordUpdateInputRdataRdataCnameRecord from the specified map.
func UnmarshalResourceRecordUpdateInputRdataRdataCnameRecord(m map[string]interface{}) (result *ResourceRecordUpdateInputRdataRdataCnameRecord, err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataCnameRecord)
	obj.Cname, err = core.UnmarshalString(m, "cname")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordUpdateInputRdataRdataMxRecord : The content of type-MX resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataMxRecord struct {
	// Hostname of Exchange server.
	Exchange *string `json:"exchange" validate:"required"`

	// Preference of the MX record.
	Preference *int64 `json:"preference" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataMxRecord : Instantiate ResourceRecordUpdateInputRdataRdataMxRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataMxRecord(exchange string, preference int64) (model *ResourceRecordUpdateInputRdataRdataMxRecord, err error) {
	model = &ResourceRecordUpdateInputRdataRdataMxRecord{
		Exchange:   core.StringPtr(exchange),
		Preference: core.Int64Ptr(preference),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataMxRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataMxRecord constructs an instance of ResourceRecordUpdateInputRdataRdataMxRecord from the specified map.
func UnmarshalResourceRecordUpdateInputRdataRdataMxRecord(m map[string]interface{}) (result *ResourceRecordUpdateInputRdataRdataMxRecord, err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataMxRecord)
	obj.Exchange, err = core.UnmarshalString(m, "exchange")
	if err != nil {
		return
	}
	obj.Preference, err = core.UnmarshalInt64(m, "preference")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordUpdateInputRdataRdataPtrRecord : The content of type-PTR resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataPtrRecord struct {
	// Hostname of the relevant A or AAAA record.
	Ptrdname *string `json:"ptrdname" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataPtrRecord : Instantiate ResourceRecordUpdateInputRdataRdataPtrRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataPtrRecord(ptrdname string) (model *ResourceRecordUpdateInputRdataRdataPtrRecord, err error) {
	model = &ResourceRecordUpdateInputRdataRdataPtrRecord{
		Ptrdname: core.StringPtr(ptrdname),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataPtrRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataPtrRecord constructs an instance of ResourceRecordUpdateInputRdataRdataPtrRecord from the specified map.
func UnmarshalResourceRecordUpdateInputRdataRdataPtrRecord(m map[string]interface{}) (result *ResourceRecordUpdateInputRdataRdataPtrRecord, err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataPtrRecord)
	obj.Ptrdname, err = core.UnmarshalString(m, "ptrdname")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordUpdateInputRdataRdataSrvRecord : The content of type-SRV resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataSrvRecord struct {
	// Port number of the target server.
	Port *int64 `json:"port" validate:"required"`

	// Priority of the SRV record.
	Priority *int64 `json:"priority" validate:"required"`

	// Hostname of the target server.
	Target *string `json:"target" validate:"required"`

	// Weight of distributing queries among multiple target servers.
	Weight *int64 `json:"weight" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataSrvRecord : Instantiate ResourceRecordUpdateInputRdataRdataSrvRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataSrvRecord(port int64, priority int64, target string, weight int64) (model *ResourceRecordUpdateInputRdataRdataSrvRecord, err error) {
	model = &ResourceRecordUpdateInputRdataRdataSrvRecord{
		Port:     core.Int64Ptr(port),
		Priority: core.Int64Ptr(priority),
		Target:   core.StringPtr(target),
		Weight:   core.Int64Ptr(weight),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataSrvRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataSrvRecord constructs an instance of ResourceRecordUpdateInputRdataRdataSrvRecord from the specified map.
func UnmarshalResourceRecordUpdateInputRdataRdataSrvRecord(m map[string]interface{}) (result *ResourceRecordUpdateInputRdataRdataSrvRecord, err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataSrvRecord)
	obj.Port, err = core.UnmarshalInt64(m, "port")
	if err != nil {
		return
	}
	obj.Priority, err = core.UnmarshalInt64(m, "priority")
	if err != nil {
		return
	}
	obj.Target, err = core.UnmarshalString(m, "target")
	if err != nil {
		return
	}
	obj.Weight, err = core.UnmarshalInt64(m, "weight")
	if err != nil {
		return
	}
	result = obj
	return
}

// ResourceRecordUpdateInputRdataRdataTxtRecord : The content of type-TXT resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataTxtRecord struct {
	// Human readable text.
	Txtdata *string `json:"text" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataTxtRecord : Instantiate ResourceRecordUpdateInputRdataRdataTxtRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataTxtRecord(txtdata string) (model *ResourceRecordUpdateInputRdataRdataTxtRecord, err error) {
	model = &ResourceRecordUpdateInputRdataRdataTxtRecord{
		Txtdata: core.StringPtr(txtdata),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataTxtRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataTxtRecord constructs an instance of ResourceRecordUpdateInputRdataRdataTxtRecord from the specified map.
func UnmarshalResourceRecordUpdateInputRdataRdataTxtRecord(m map[string]interface{}) (result *ResourceRecordUpdateInputRdataRdataTxtRecord, err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataTxtRecord)
	obj.Txtdata, err = core.UnmarshalString(m, "text")
	if err != nil {
		return
	}
	result = obj
	return
}
