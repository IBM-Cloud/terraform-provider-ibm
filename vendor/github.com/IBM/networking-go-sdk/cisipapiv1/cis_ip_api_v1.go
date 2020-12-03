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

// Package cisipapiv1 : Operations and models for the CisIpApiV1 service
package cisipapiv1

import (
	"encoding/json"
	"reflect"

	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/networking-go-sdk/common"
)

// CisIpApiV1 : This document describes CIS IP API.
//
// Version: 1.0.0
type CisIpApiV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "cis_ip_api"

// CisIpApiV1Options : Service options
type CisIpApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewCisIpApiV1UsingExternalConfig : constructs an instance of CisIpApiV1 with passed in options and external configuration.
func NewCisIpApiV1UsingExternalConfig(options *CisIpApiV1Options) (cisIpApi *CisIpApiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	cisIpApi, err = NewCisIpApiV1(options)
	if err != nil {
		return
	}

	err = cisIpApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = cisIpApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewCisIpApiV1 : constructs an instance of CisIpApiV1 with passed in options.
func NewCisIpApiV1(options *CisIpApiV1Options) (service *CisIpApiV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
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

	service = &CisIpApiV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (cisIpApi *CisIpApiV1) SetServiceURL(url string) error {
	return cisIpApi.Service.SetServiceURL(url)
}

// ListIps : List of all IP addresses used by the CIS proxy
// List of all IP addresses used by the CIS proxy.
func (cisIpApi *CisIpApiV1) ListIps(listIpsOptions *ListIpsOptions) (result *IpResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listIpsOptions, "listIpsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/ips"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(cisIpApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listIpsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("cis_ip_api", "V1", "ListIps")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = cisIpApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIpResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// IpResponseResult : Container for response information.
type IpResponseResult struct {
	// List of IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs,omitempty"`

	// List of IPv6 CIDR addresses.
	Ipv6Cidrs []string `json:"ipv6_cidrs,omitempty"`
}


// UnmarshalIpResponseResult unmarshals an instance of IpResponseResult from the specified map of raw messages.
func UnmarshalIpResponseResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IpResponseResult)
	err = core.UnmarshalPrimitive(m, "ipv4_cidrs", &obj.Ipv4Cidrs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ipv6_cidrs", &obj.Ipv6Cidrs)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListIpsOptions : The ListIps options.
type ListIpsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListIpsOptions : Instantiate ListIpsOptions
func (*CisIpApiV1) NewListIpsOptions() *ListIpsOptions {
	return &ListIpsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListIpsOptions) SetHeaders(param map[string]string) *ListIpsOptions {
	options.Headers = param
	return options
}

// IpResponse : ip response.
type IpResponse struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// Container for response information.
	Result *IpResponseResult `json:"result" validate:"required"`
}


// UnmarshalIpResponse unmarshals an instance of IpResponse from the specified map of raw messages.
func UnmarshalIpResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IpResponse)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalIpResponseResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
