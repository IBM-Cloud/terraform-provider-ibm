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

// Package routingv1 : Operations and models for the RoutingV1 service
package routingv1

import (
	"encoding/json"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	common "github.ibm.com/ibmcloud/networking-go-sdk/common"
	"reflect"
)

// RoutingV1 : Routing
//
// Version: 1.0.1
type RoutingV1 struct {
	Service *core.BaseService

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string

	// Zone identifier.
	ZoneIdentifier *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "routing"

// RoutingV1Options : Service options
type RoutingV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string `validate:"required"`

	// Zone identifier.
	ZoneIdentifier *string `validate:"required"`
}

// NewRoutingV1UsingExternalConfig : constructs an instance of RoutingV1 with passed in options and external configuration.
func NewRoutingV1UsingExternalConfig(options *RoutingV1Options) (routing *RoutingV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	routing, err = NewRoutingV1(options)
	if err != nil {
		return
	}

	err = routing.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = routing.Service.SetServiceURL(options.URL)
	}
	return
}

// NewRoutingV1 : constructs an instance of RoutingV1 with passed in options.
func NewRoutingV1(options *RoutingV1Options) (service *RoutingV1, err error) {
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

	service = &RoutingV1{
		Service:        baseService,
		Crn:            options.Crn,
		ZoneIdentifier: options.ZoneIdentifier,
	}

	return
}

// SetServiceURL sets the service URL
func (routing *RoutingV1) SetServiceURL(url string) error {
	return routing.Service.SetServiceURL(url)
}

// GetSmartRouting : Get Routing feature smart routing setting
// Get Routing feature smart routing setting for a zone.
func (routing *RoutingV1) GetSmartRouting(getSmartRoutingOptions *GetSmartRoutingOptions) (result *SmartRoutingResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSmartRoutingOptions, "getSmartRoutingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "routing/smart_routing"}
	pathParameters := []string{*routing.Crn, *routing.ZoneIdentifier}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(routing.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSmartRoutingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("routing", "V1", "GetSmartRouting")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = routing.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSmartRoutingResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateSmartRouting : Update Routing feature smart route setting
// Update Routing feature smart route setting for a zone.
func (routing *RoutingV1) UpdateSmartRouting(updateSmartRoutingOptions *UpdateSmartRoutingOptions) (result *SmartRoutingResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateSmartRoutingOptions, "updateSmartRoutingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "routing/smart_routing"}
	pathParameters := []string{*routing.Crn, *routing.ZoneIdentifier}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(routing.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSmartRoutingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("routing", "V1", "UpdateSmartRouting")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateSmartRoutingOptions.Value != nil {
		body["value"] = updateSmartRoutingOptions.Value
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = routing.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSmartRoutingResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSmartRoutingOptions : The GetSmartRouting options.
type GetSmartRoutingOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSmartRoutingOptions : Instantiate GetSmartRoutingOptions
func (*RoutingV1) NewGetSmartRoutingOptions() *GetSmartRoutingOptions {
	return &GetSmartRoutingOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSmartRoutingOptions) SetHeaders(param map[string]string) *GetSmartRoutingOptions {
	options.Headers = param
	return options
}

// SmartRoutingRespResult : Container for response information.
type SmartRoutingRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`

	// Value.
	Value *string `json:"value" validate:"required"`

	// Editable.
	Editable *bool `json:"editable" validate:"required"`

	// Modified date.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`
}

// UnmarshalSmartRoutingRespResult unmarshals an instance of SmartRoutingRespResult from the specified map of raw messages.
func UnmarshalSmartRoutingRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SmartRoutingRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "editable", &obj.Editable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateSmartRoutingOptions : The UpdateSmartRouting options.
type UpdateSmartRoutingOptions struct {
	// Value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateSmartRoutingOptions.Value property.
// Value.
const (
	UpdateSmartRoutingOptions_Value_Off = "off"
	UpdateSmartRoutingOptions_Value_On  = "on"
)

// NewUpdateSmartRoutingOptions : Instantiate UpdateSmartRoutingOptions
func (*RoutingV1) NewUpdateSmartRoutingOptions() *UpdateSmartRoutingOptions {
	return &UpdateSmartRoutingOptions{}
}

// SetValue : Allow user to set Value
func (options *UpdateSmartRoutingOptions) SetValue(value string) *UpdateSmartRoutingOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSmartRoutingOptions) SetHeaders(param map[string]string) *UpdateSmartRoutingOptions {
	options.Headers = param
	return options
}

// SmartRoutingResp : smart routing response.
type SmartRoutingResp struct {
	// Container for response information.
	Result *SmartRoutingRespResult `json:"result" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalSmartRoutingResp unmarshals an instance of SmartRoutingResp from the specified map of raw messages.
func UnmarshalSmartRoutingResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SmartRoutingResp)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalSmartRoutingRespResult)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
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
