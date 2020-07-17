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

// Package wafapiv1 : Operations and models for the WafApiV1 service
package wafapiv1

import (
	"encoding/json"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.ibm.com/ibmcloud/networking-go-sdk/common"
	"reflect"
)

// WafApiV1 : This document describes CIS WAF API.
//
// Version: 1.0.0
type WafApiV1 struct {
	Service *core.BaseService

	// cloud resource name.
	Crn *string

	// zone id.
	ZoneID *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "waf_api"

// WafApiV1Options : Service options
type WafApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// cloud resource name.
	Crn *string `validate:"required"`

	// zone id.
	ZoneID *string `validate:"required"`
}

// NewWafApiV1UsingExternalConfig : constructs an instance of WafApiV1 with passed in options and external configuration.
func NewWafApiV1UsingExternalConfig(options *WafApiV1Options) (wafApi *WafApiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	wafApi, err = NewWafApiV1(options)
	if err != nil {
		return
	}

	err = wafApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = wafApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewWafApiV1 : constructs an instance of WafApiV1 with passed in options.
func NewWafApiV1(options *WafApiV1Options) (service *WafApiV1, err error) {
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

	service = &WafApiV1{
		Service: baseService,
		Crn:     options.Crn,
		ZoneID:  options.ZoneID,
	}

	return
}

// SetServiceURL sets the service URL
func (wafApi *WafApiV1) SetServiceURL(url string) error {
	return wafApi.Service.SetServiceURL(url)
}

// GetWafSettings : Get WAF setting of a specific zone
// Get WAF of a specific zone.
func (wafApi *WafApiV1) GetWafSettings(getWafSettingsOptions *GetWafSettingsOptions) (result *WafResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getWafSettingsOptions, "getWafSettingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/waf"}
	pathParameters := []string{*wafApi.Crn, *wafApi.ZoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(wafApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getWafSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("waf_api", "V1", "GetWafSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = wafApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWafResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateWafSettings : Set WAF setting for a specific zone
// Set WAF (on | off) for a specific zone.
func (wafApi *WafApiV1) UpdateWafSettings(updateWafSettingsOptions *UpdateWafSettingsOptions) (result *WafResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateWafSettingsOptions, "updateWafSettingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "settings/waf"}
	pathParameters := []string{*wafApi.Crn, *wafApi.ZoneID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(wafApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateWafSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("waf_api", "V1", "UpdateWafSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateWafSettingsOptions.Value != nil {
		body["value"] = updateWafSettingsOptions.Value
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
	response, err = wafApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWafResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetWafSettingsOptions : The GetWafSettings options.
type GetWafSettingsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetWafSettingsOptions : Instantiate GetWafSettingsOptions
func (*WafApiV1) NewGetWafSettingsOptions() *GetWafSettingsOptions {
	return &GetWafSettingsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetWafSettingsOptions) SetHeaders(param map[string]string) *GetWafSettingsOptions {
	options.Headers = param
	return options
}

// UpdateWafSettingsOptions : The UpdateWafSettings options.
type UpdateWafSettingsOptions struct {
	// value.
	Value *string `json:"value,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateWafSettingsOptions.Value property.
// value.
const (
	UpdateWafSettingsOptions_Value_Off = "off"
	UpdateWafSettingsOptions_Value_On  = "on"
)

// NewUpdateWafSettingsOptions : Instantiate UpdateWafSettingsOptions
func (*WafApiV1) NewUpdateWafSettingsOptions() *UpdateWafSettingsOptions {
	return &UpdateWafSettingsOptions{}
}

// SetValue : Allow user to set Value
func (options *UpdateWafSettingsOptions) SetValue(value string) *UpdateWafSettingsOptions {
	options.Value = core.StringPtr(value)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateWafSettingsOptions) SetHeaders(param map[string]string) *UpdateWafSettingsOptions {
	options.Headers = param
	return options
}

// WafResponseResult : result.
type WafResponseResult struct {
	// id.
	ID *string `json:"id,omitempty"`

	// value.
	Value *string `json:"value,omitempty"`

	// editable.
	Editable *bool `json:"editable,omitempty"`

	// modified date.
	ModifiedOn *string `json:"modified_on,omitempty"`
}

// UnmarshalWafResponseResult unmarshals an instance of WafResponseResult from the specified map of raw messages.
func UnmarshalWafResponseResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafResponseResult)
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

// WafResponse : waf response.
type WafResponse struct {
	// success.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// result.
	Result *WafResponseResult `json:"result" validate:"required"`
}

// UnmarshalWafResponse unmarshals an instance of WafResponse from the specified map of raw messages.
func UnmarshalWafResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafResponse)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalWafResponseResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
