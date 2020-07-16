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

// Package globalloadbalancereventsv1 : Operations and models for the GlobalLoadBalancerEventsV1 service
package globalloadbalancereventsv1

import (
	"encoding/json"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	common "github.ibm.com/ibmcloud/networking-go-sdk/common"
	"reflect"
)

// GlobalLoadBalancerEventsV1 : Global Load Balancer Healthcheck Events
//
// Version: 1.0.1
type GlobalLoadBalancerEventsV1 struct {
	Service *core.BaseService

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "global_load_balancer_events"

// GlobalLoadBalancerEventsV1Options : Service options
type GlobalLoadBalancerEventsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string `validate:"required"`
}

// NewGlobalLoadBalancerEventsV1UsingExternalConfig : constructs an instance of GlobalLoadBalancerEventsV1 with passed in options and external configuration.
func NewGlobalLoadBalancerEventsV1UsingExternalConfig(options *GlobalLoadBalancerEventsV1Options) (globalLoadBalancerEvents *GlobalLoadBalancerEventsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	globalLoadBalancerEvents, err = NewGlobalLoadBalancerEventsV1(options)
	if err != nil {
		return
	}

	err = globalLoadBalancerEvents.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = globalLoadBalancerEvents.Service.SetServiceURL(options.URL)
	}
	return
}

// NewGlobalLoadBalancerEventsV1 : constructs an instance of GlobalLoadBalancerEventsV1 with passed in options.
func NewGlobalLoadBalancerEventsV1(options *GlobalLoadBalancerEventsV1Options) (service *GlobalLoadBalancerEventsV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &GlobalLoadBalancerEventsV1{
		Service: baseService,
		Crn:     options.Crn,
	}

	return
}

// SetServiceURL sets the service URL
func (globalLoadBalancerEvents *GlobalLoadBalancerEventsV1) SetServiceURL(url string) error {
	return globalLoadBalancerEvents.Service.SetServiceURL(url)
}

// GetLoadBalancerEvents : List all load balancer events
// Get load balancer events for all origins.
func (globalLoadBalancerEvents *GlobalLoadBalancerEventsV1) GetLoadBalancerEvents(getLoadBalancerEventsOptions *GetLoadBalancerEventsOptions) (result *ListEventsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getLoadBalancerEventsOptions, "getLoadBalancerEventsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "load_balancers/events"}
	pathParameters := []string{*globalLoadBalancerEvents.Crn}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalLoadBalancerEvents.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLoadBalancerEventsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_load_balancer_events", "V1", "GetLoadBalancerEvents")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalLoadBalancerEvents.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListEventsResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetLoadBalancerEventsOptions : The GetLoadBalancerEvents options.
type GetLoadBalancerEventsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLoadBalancerEventsOptions : Instantiate GetLoadBalancerEventsOptions
func (*GlobalLoadBalancerEventsV1) NewGetLoadBalancerEventsOptions() *GetLoadBalancerEventsOptions {
	return &GetLoadBalancerEventsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetLoadBalancerEventsOptions) SetHeaders(param map[string]string) *GetLoadBalancerEventsOptions {
	options.Headers = param
	return options
}

// ListEventsRespResultInfo : result information.
type ListEventsRespResultInfo struct {
	// Page number.
	Page *int64 `json:"page" validate:"required"`

	// Number of results per page.
	PerPage *int64 `json:"per_page" validate:"required"`

	// Number of results.
	Count *int64 `json:"count" validate:"required"`

	// Total number of results.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalListEventsRespResultInfo unmarshals an instance of ListEventsRespResultInfo from the specified map of raw messages.
func UnmarshalListEventsRespResultInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListEventsRespResultInfo)
	err = core.UnmarshalPrimitive(m, "page", &obj.Page)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "per_page", &obj.PerPage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListEventsRespResultItem : ListEventsRespResultItem struct
type ListEventsRespResultItem struct {
	// ID of the event.
	ID *string `json:"id,omitempty"`

	// Time of the event.
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// Pool information.
	Pool []ListEventsRespResultItemPoolItem `json:"pool,omitempty"`

	// Load balancer origins.
	Origins []ListEventsRespResultItemOriginsItem `json:"origins,omitempty"`
}

// UnmarshalListEventsRespResultItem unmarshals an instance of ListEventsRespResultItem from the specified map of raw messages.
func UnmarshalListEventsRespResultItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListEventsRespResultItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pool", &obj.Pool, UnmarshalListEventsRespResultItemPoolItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "origins", &obj.Origins, UnmarshalListEventsRespResultItemOriginsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListEventsRespResultItemOriginsItem : ListEventsRespResultItemOriginsItem struct
type ListEventsRespResultItemOriginsItem struct {
	// Origin name.
	Name *string `json:"name,omitempty"`

	// Origin address.
	Address *string `json:"address,omitempty"`

	// Origin id.
	Ip *string `json:"ip,omitempty"`

	// Origin enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Origin healthy.
	Healthy *bool `json:"healthy,omitempty"`

	// Origin failure reason.
	FailureReason *string `json:"failure_reason,omitempty"`

	// Origin changed.
	Changed *bool `json:"changed,omitempty"`
}

// UnmarshalListEventsRespResultItemOriginsItem unmarshals an instance of ListEventsRespResultItemOriginsItem from the specified map of raw messages.
func UnmarshalListEventsRespResultItemOriginsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListEventsRespResultItemOriginsItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "healthy", &obj.Healthy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failure_reason", &obj.FailureReason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "changed", &obj.Changed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListEventsRespResultItemPoolItem : ListEventsRespResultItemPoolItem struct
type ListEventsRespResultItemPoolItem struct {
	// Pool id.
	ID *string `json:"id,omitempty"`

	// Pool name.
	Name *string `json:"name,omitempty"`

	// Pool is healthy.
	Healthy *bool `json:"healthy,omitempty"`

	// Pool changed.
	Changed *bool `json:"changed,omitempty"`

	// Minimum origins.
	MinimumOrigins *int64 `json:"minimum_origins,omitempty"`
}

// UnmarshalListEventsRespResultItemPoolItem unmarshals an instance of ListEventsRespResultItemPoolItem from the specified map of raw messages.
func UnmarshalListEventsRespResultItemPoolItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListEventsRespResultItemPoolItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "healthy", &obj.Healthy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "changed", &obj.Changed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_origins", &obj.MinimumOrigins)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListEventsResp : events list response object.
type ListEventsResp struct {
	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Result of the operation.
	Result []ListEventsRespResultItem `json:"result" validate:"required"`

	// result information.
	ResultInfo *ListEventsRespResultInfo `json:"result_info" validate:"required"`

	// Array of errors returned.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalListEventsResp unmarshals an instance of ListEventsResp from the specified map of raw messages.
func UnmarshalListEventsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListEventsResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalListEventsRespResultItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result_info", &obj.ResultInfo, UnmarshalListEventsRespResultInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
