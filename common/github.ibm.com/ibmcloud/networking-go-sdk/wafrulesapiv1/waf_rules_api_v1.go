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

// Package wafrulesapiv1 : Operations and models for the WafRulesApiV1 service
package wafrulesapiv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.ibm.com/ibmcloud/networking-go-sdk/common"
	"reflect"
)

// WafRulesApiV1 : This document describes CIS WAF Rules API.
//
// Version: 1.0.0
type WafRulesApiV1 struct {
	Service *core.BaseService

	// cloud resource name.
	Crn *string

	// zone id.
	ZoneID *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "waf_rules_api"

// WafRulesApiV1Options : Service options
type WafRulesApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// cloud resource name.
	Crn *string `validate:"required"`

	// zone id.
	ZoneID *string `validate:"required"`
}

// NewWafRulesApiV1UsingExternalConfig : constructs an instance of WafRulesApiV1 with passed in options and external configuration.
func NewWafRulesApiV1UsingExternalConfig(options *WafRulesApiV1Options) (wafRulesApi *WafRulesApiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	wafRulesApi, err = NewWafRulesApiV1(options)
	if err != nil {
		return
	}

	err = wafRulesApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = wafRulesApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewWafRulesApiV1 : constructs an instance of WafRulesApiV1 with passed in options.
func NewWafRulesApiV1(options *WafRulesApiV1Options) (service *WafRulesApiV1, err error) {
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

	service = &WafRulesApiV1{
		Service: baseService,
		Crn:     options.Crn,
		ZoneID:  options.ZoneID,
	}

	return
}

// SetServiceURL sets the service URL
func (wafRulesApi *WafRulesApiV1) SetServiceURL(url string) error {
	return wafRulesApi.Service.SetServiceURL(url)
}

// ListWafRules : List all WAF rules
// List all Web Application Firewall (WAF) rules.
func (wafRulesApi *WafRulesApiV1) ListWafRules(listWafRulesOptions *ListWafRulesOptions) (result *WafRulesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listWafRulesOptions, "listWafRulesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listWafRulesOptions, "listWafRulesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "firewall/waf/packages", "rules"}
	pathParameters := []string{*wafRulesApi.Crn, *wafRulesApi.ZoneID, *listWafRulesOptions.PackageID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(wafRulesApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listWafRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("waf_rules_api", "V1", "ListWafRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listWafRulesOptions.Mode != nil {
		builder.AddQuery("mode", fmt.Sprint(*listWafRulesOptions.Mode))
	}
	if listWafRulesOptions.Priority != nil {
		builder.AddQuery("priority", fmt.Sprint(*listWafRulesOptions.Priority))
	}
	if listWafRulesOptions.Match != nil {
		builder.AddQuery("match", fmt.Sprint(*listWafRulesOptions.Match))
	}
	if listWafRulesOptions.Order != nil {
		builder.AddQuery("order", fmt.Sprint(*listWafRulesOptions.Order))
	}
	if listWafRulesOptions.GroupID != nil {
		builder.AddQuery("group_id", fmt.Sprint(*listWafRulesOptions.GroupID))
	}
	if listWafRulesOptions.Description != nil {
		builder.AddQuery("description", fmt.Sprint(*listWafRulesOptions.Description))
	}
	if listWafRulesOptions.Direction != nil {
		builder.AddQuery("direction", fmt.Sprint(*listWafRulesOptions.Direction))
	}
	if listWafRulesOptions.Page != nil {
		builder.AddQuery("page", fmt.Sprint(*listWafRulesOptions.Page))
	}
	if listWafRulesOptions.PerPage != nil {
		builder.AddQuery("per_page", fmt.Sprint(*listWafRulesOptions.PerPage))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = wafRulesApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWafRulesResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetWafRule : Get WAF Rule info
// Get individual information about a rule.
func (wafRulesApi *WafRulesApiV1) GetWafRule(getWafRuleOptions *GetWafRuleOptions) (result *WafRuleResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getWafRuleOptions, "getWafRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getWafRuleOptions, "getWafRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "firewall/waf/packages", "rules"}
	pathParameters := []string{*wafRulesApi.Crn, *wafRulesApi.ZoneID, *getWafRuleOptions.PackageID, *getWafRuleOptions.Identifier}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(wafRulesApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getWafRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("waf_rules_api", "V1", "GetWafRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = wafRulesApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWafRuleResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateWafRule : Update WAF rule
// Update the action the rule will perform if triggered on the zone.
func (wafRulesApi *WafRulesApiV1) UpdateWafRule(updateWafRuleOptions *UpdateWafRuleOptions) (result *WafRuleResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateWafRuleOptions, "updateWafRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateWafRuleOptions, "updateWafRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "firewall/waf/packages", "rules"}
	pathParameters := []string{*wafRulesApi.Crn, *wafRulesApi.ZoneID, *updateWafRuleOptions.PackageID, *updateWafRuleOptions.Identifier}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(wafRulesApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateWafRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("waf_rules_api", "V1", "UpdateWafRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateWafRuleOptions.Cis != nil {
		body["cis"] = updateWafRuleOptions.Cis
	}
	if updateWafRuleOptions.Owasp != nil {
		body["owasp"] = updateWafRuleOptions.Owasp
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
	response, err = wafRulesApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWafRuleResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetWafRuleOptions : The GetWafRule options.
type GetWafRuleOptions struct {
	// package id.
	PackageID *string `json:"package_id" validate:"required"`

	// rule identifier.
	Identifier *string `json:"identifier" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetWafRuleOptions : Instantiate GetWafRuleOptions
func (*WafRulesApiV1) NewGetWafRuleOptions(packageID string, identifier string) *GetWafRuleOptions {
	return &GetWafRuleOptions{
		PackageID:  core.StringPtr(packageID),
		Identifier: core.StringPtr(identifier),
	}
}

// SetPackageID : Allow user to set PackageID
func (options *GetWafRuleOptions) SetPackageID(packageID string) *GetWafRuleOptions {
	options.PackageID = core.StringPtr(packageID)
	return options
}

// SetIdentifier : Allow user to set Identifier
func (options *GetWafRuleOptions) SetIdentifier(identifier string) *GetWafRuleOptions {
	options.Identifier = core.StringPtr(identifier)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetWafRuleOptions) SetHeaders(param map[string]string) *GetWafRuleOptions {
	options.Headers = param
	return options
}

// ListWafRulesOptions : The ListWafRules options.
type ListWafRulesOptions struct {
	// package id.
	PackageID *string `json:"package_id" validate:"required"`

	// The Rule Mode.
	Mode *string `json:"mode,omitempty"`

	// The order in which the individual rule is executed within the related group.
	Priority *string `json:"priority,omitempty"`

	// Whether to match all search requirements or at least one. default value: all. valid values: any, all.
	Match *string `json:"match,omitempty"`

	// Field to order rules by. valid values: priority, group_id, description.
	Order *string `json:"order,omitempty"`

	// WAF group identifier tag. max length: 32; Read-only.
	GroupID *string `json:"group_id,omitempty"`

	// Public description of the rule.
	Description *string `json:"description,omitempty"`

	// Direction to order rules. valid values: asc, desc.
	Direction *string `json:"direction,omitempty"`

	// Page number of paginated results. default value: 1; min value:1.
	Page *int64 `json:"page,omitempty"`

	// Number of rules per page. default value: 50; min value:5; max value:100.
	PerPage *int64 `json:"per_page,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListWafRulesOptions.Mode property.
// The Rule Mode.
const (
	ListWafRulesOptions_Mode_Off = "off"
	ListWafRulesOptions_Mode_On  = "on"
)

// NewListWafRulesOptions : Instantiate ListWafRulesOptions
func (*WafRulesApiV1) NewListWafRulesOptions(packageID string) *ListWafRulesOptions {
	return &ListWafRulesOptions{
		PackageID: core.StringPtr(packageID),
	}
}

// SetPackageID : Allow user to set PackageID
func (options *ListWafRulesOptions) SetPackageID(packageID string) *ListWafRulesOptions {
	options.PackageID = core.StringPtr(packageID)
	return options
}

// SetMode : Allow user to set Mode
func (options *ListWafRulesOptions) SetMode(mode string) *ListWafRulesOptions {
	options.Mode = core.StringPtr(mode)
	return options
}

// SetPriority : Allow user to set Priority
func (options *ListWafRulesOptions) SetPriority(priority string) *ListWafRulesOptions {
	options.Priority = core.StringPtr(priority)
	return options
}

// SetMatch : Allow user to set Match
func (options *ListWafRulesOptions) SetMatch(match string) *ListWafRulesOptions {
	options.Match = core.StringPtr(match)
	return options
}

// SetOrder : Allow user to set Order
func (options *ListWafRulesOptions) SetOrder(order string) *ListWafRulesOptions {
	options.Order = core.StringPtr(order)
	return options
}

// SetGroupID : Allow user to set GroupID
func (options *ListWafRulesOptions) SetGroupID(groupID string) *ListWafRulesOptions {
	options.GroupID = core.StringPtr(groupID)
	return options
}

// SetDescription : Allow user to set Description
func (options *ListWafRulesOptions) SetDescription(description string) *ListWafRulesOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetDirection : Allow user to set Direction
func (options *ListWafRulesOptions) SetDirection(direction string) *ListWafRulesOptions {
	options.Direction = core.StringPtr(direction)
	return options
}

// SetPage : Allow user to set Page
func (options *ListWafRulesOptions) SetPage(page int64) *ListWafRulesOptions {
	options.Page = core.Int64Ptr(page)
	return options
}

// SetPerPage : Allow user to set PerPage
func (options *ListWafRulesOptions) SetPerPage(perPage int64) *ListWafRulesOptions {
	options.PerPage = core.Int64Ptr(perPage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListWafRulesOptions) SetHeaders(param map[string]string) *ListWafRulesOptions {
	options.Headers = param
	return options
}

// UpdateWafRuleOptions : The UpdateWafRule options.
type UpdateWafRuleOptions struct {
	// package id.
	PackageID *string `json:"package_id" validate:"required"`

	// rule identifier.
	Identifier *string `json:"identifier" validate:"required"`

	// cis package.
	Cis *WafRuleBodyCis `json:"cis,omitempty"`

	// owasp package.
	Owasp *WafRuleBodyOwasp `json:"owasp,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateWafRuleOptions : Instantiate UpdateWafRuleOptions
func (*WafRulesApiV1) NewUpdateWafRuleOptions(packageID string, identifier string) *UpdateWafRuleOptions {
	return &UpdateWafRuleOptions{
		PackageID:  core.StringPtr(packageID),
		Identifier: core.StringPtr(identifier),
	}
}

// SetPackageID : Allow user to set PackageID
func (options *UpdateWafRuleOptions) SetPackageID(packageID string) *UpdateWafRuleOptions {
	options.PackageID = core.StringPtr(packageID)
	return options
}

// SetIdentifier : Allow user to set Identifier
func (options *UpdateWafRuleOptions) SetIdentifier(identifier string) *UpdateWafRuleOptions {
	options.Identifier = core.StringPtr(identifier)
	return options
}

// SetCis : Allow user to set Cis
func (options *UpdateWafRuleOptions) SetCis(cis *WafRuleBodyCis) *UpdateWafRuleOptions {
	options.Cis = cis
	return options
}

// SetOwasp : Allow user to set Owasp
func (options *UpdateWafRuleOptions) SetOwasp(owasp *WafRuleBodyOwasp) *UpdateWafRuleOptions {
	options.Owasp = owasp
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateWafRuleOptions) SetHeaders(param map[string]string) *UpdateWafRuleOptions {
	options.Headers = param
	return options
}

// WafRuleBodyCis : cis package.
type WafRuleBodyCis struct {
	// mode to choose from.
	Mode *string `json:"mode" validate:"required"`
}

// Constants associated with the WafRuleBodyCis.Mode property.
// mode to choose from.
const (
	WafRuleBodyCis_Mode_Block     = "block"
	WafRuleBodyCis_Mode_Challenge = "challenge"
	WafRuleBodyCis_Mode_Default   = "default"
	WafRuleBodyCis_Mode_Disable   = "disable"
	WafRuleBodyCis_Mode_Simulate  = "simulate"
)

// NewWafRuleBodyCis : Instantiate WafRuleBodyCis (Generic Model Constructor)
func (*WafRulesApiV1) NewWafRuleBodyCis(mode string) (model *WafRuleBodyCis, err error) {
	model = &WafRuleBodyCis{
		Mode: core.StringPtr(mode),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalWafRuleBodyCis unmarshals an instance of WafRuleBodyCis from the specified map of raw messages.
func UnmarshalWafRuleBodyCis(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRuleBodyCis)
	err = core.UnmarshalPrimitive(m, "mode", &obj.Mode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WafRuleBodyOwasp : owasp package.
type WafRuleBodyOwasp struct {
	// mode to choose from. 'owasp' limited modes - on and off.
	Mode *string `json:"mode" validate:"required"`
}

// Constants associated with the WafRuleBodyOwasp.Mode property.
// mode to choose from. 'owasp' limited modes - on and off.
const (
	WafRuleBodyOwasp_Mode_Off = "off"
	WafRuleBodyOwasp_Mode_On  = "on"
)

// NewWafRuleBodyOwasp : Instantiate WafRuleBodyOwasp (Generic Model Constructor)
func (*WafRulesApiV1) NewWafRuleBodyOwasp(mode string) (model *WafRuleBodyOwasp, err error) {
	model = &WafRuleBodyOwasp{
		Mode: core.StringPtr(mode),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalWafRuleBodyOwasp unmarshals an instance of WafRuleBodyOwasp from the specified map of raw messages.
func UnmarshalWafRuleBodyOwasp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRuleBodyOwasp)
	err = core.UnmarshalPrimitive(m, "mode", &obj.Mode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WafRuleResponseResult : Information about a Rule.
type WafRuleResponseResult struct {
	// ID.
	ID *string `json:"id,omitempty"`

	// description.
	Description *string `json:"description,omitempty"`

	// priority.
	Priority *string `json:"priority,omitempty"`

	// group definition.
	Group *WafRuleResponseResultGroup `json:"group,omitempty"`

	// package id.
	PackageID *string `json:"package_id,omitempty"`

	// allowed modes.
	AllowedModes []string `json:"allowed_modes,omitempty"`

	// mode.
	Mode *string `json:"mode,omitempty"`
}

// Constants associated with the WafRuleResponseResult.Mode property.
// mode.
const (
	WafRuleResponseResult_Mode_Off = "off"
	WafRuleResponseResult_Mode_On  = "on"
)

// UnmarshalWafRuleResponseResult unmarshals an instance of WafRuleResponseResult from the specified map of raw messages.
func UnmarshalWafRuleResponseResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRuleResponseResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "priority", &obj.Priority)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "group", &obj.Group, UnmarshalWafRuleResponseResultGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "package_id", &obj.PackageID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allowed_modes", &obj.AllowedModes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mode", &obj.Mode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WafRuleResponseResultGroup : group definition.
type WafRuleResponseResultGroup struct {
	// group id.
	ID *string `json:"id,omitempty"`

	// group name.
	Name *string `json:"name,omitempty"`
}

// UnmarshalWafRuleResponseResultGroup unmarshals an instance of WafRuleResponseResultGroup from the specified map of raw messages.
func UnmarshalWafRuleResponseResultGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRuleResponseResultGroup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WafRulesResponseResultInfo : result information.
type WafRulesResponseResultInfo struct {
	// current page.
	Page *int64 `json:"page,omitempty"`

	// number of data per page.
	PerPage *int64 `json:"per_page,omitempty"`

	// count.
	Count *int64 `json:"count,omitempty"`

	// total count of data.
	TotalCount *int64 `json:"total_count,omitempty"`
}

// UnmarshalWafRulesResponseResultInfo unmarshals an instance of WafRulesResponseResultInfo from the specified map of raw messages.
func UnmarshalWafRulesResponseResultInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRulesResponseResultInfo)
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

// WafRulesResponseResultItem : WafRulesResponseResultItem struct
type WafRulesResponseResultItem struct {
	// ID.
	ID *string `json:"id,omitempty"`

	// description.
	Description *string `json:"description,omitempty"`

	// priority.
	Priority *string `json:"priority,omitempty"`

	// group definition.
	Group *WafRulesResponseResultItemGroup `json:"group,omitempty"`

	// package id.
	PackageID *string `json:"package_id,omitempty"`

	// allowed modes.
	AllowedModes []string `json:"allowed_modes,omitempty"`

	// mode.
	Mode *string `json:"mode,omitempty"`
}

// Constants associated with the WafRulesResponseResultItem.Mode property.
// mode.
const (
	WafRulesResponseResultItem_Mode_Off = "off"
	WafRulesResponseResultItem_Mode_On  = "on"
)

// UnmarshalWafRulesResponseResultItem unmarshals an instance of WafRulesResponseResultItem from the specified map of raw messages.
func UnmarshalWafRulesResponseResultItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRulesResponseResultItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "priority", &obj.Priority)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "group", &obj.Group, UnmarshalWafRulesResponseResultItemGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "package_id", &obj.PackageID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allowed_modes", &obj.AllowedModes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mode", &obj.Mode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WafRulesResponseResultItemGroup : group definition.
type WafRulesResponseResultItemGroup struct {
	// group id.
	ID *string `json:"id,omitempty"`

	// group name.
	Name *string `json:"name,omitempty"`
}

// UnmarshalWafRulesResponseResultItemGroup unmarshals an instance of WafRulesResponseResultItemGroup from the specified map of raw messages.
func UnmarshalWafRulesResponseResultItemGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRulesResponseResultItemGroup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WafRuleResponse : waf rule response.
type WafRuleResponse struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	// Information about a Rule.
	Result *WafRuleResponseResult `json:"result" validate:"required"`
}

// UnmarshalWafRuleResponse unmarshals an instance of WafRuleResponse from the specified map of raw messages.
func UnmarshalWafRuleResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRuleResponse)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalWafRuleResponseResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WafRulesResponse : waf rule response.
type WafRulesResponse struct {
	// Was operation successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`

	// Array of Rules.
	Result []WafRulesResponseResultItem `json:"result" validate:"required"`

	// result information.
	ResultInfo *WafRulesResponseResultInfo `json:"result_info,omitempty"`
}

// UnmarshalWafRulesResponse unmarshals an instance of WafRulesResponse from the specified map of raw messages.
func UnmarshalWafRulesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WafRulesResponse)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalWafRulesResponseResultItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result_info", &obj.ResultInfo, UnmarshalWafRulesResponseResultInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
