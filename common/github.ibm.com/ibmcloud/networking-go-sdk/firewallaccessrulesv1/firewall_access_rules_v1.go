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

// Package firewallaccessrulesv1 : Operations and models for the FirewallAccessRulesV1 service
package firewallaccessrulesv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	common "github.ibm.com/ibmcloud/networking-go-sdk/common"
	"reflect"
)

// FirewallAccessRulesV1 : Instance Level Firewall Access Rules
//
// Version: 1.0.1
type FirewallAccessRulesV1 struct {
	Service *core.BaseService

	// Full crn of the service instance.
	Crn *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "firewall_access_rules"

// FirewallAccessRulesV1Options : Service options
type FirewallAccessRulesV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full crn of the service instance.
	Crn *string `validate:"required"`
}

// NewFirewallAccessRulesV1UsingExternalConfig : constructs an instance of FirewallAccessRulesV1 with passed in options and external configuration.
func NewFirewallAccessRulesV1UsingExternalConfig(options *FirewallAccessRulesV1Options) (firewallAccessRules *FirewallAccessRulesV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	firewallAccessRules, err = NewFirewallAccessRulesV1(options)
	if err != nil {
		return
	}

	err = firewallAccessRules.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = firewallAccessRules.Service.SetServiceURL(options.URL)
	}
	return
}

// NewFirewallAccessRulesV1 : constructs an instance of FirewallAccessRulesV1 with passed in options.
func NewFirewallAccessRulesV1(options *FirewallAccessRulesV1Options) (service *FirewallAccessRulesV1, err error) {
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

	service = &FirewallAccessRulesV1{
		Service: baseService,
		Crn:     options.Crn,
	}

	return
}

// SetServiceURL sets the service URL
func (firewallAccessRules *FirewallAccessRulesV1) SetServiceURL(url string) error {
	return firewallAccessRules.Service.SetServiceURL(url)
}

// ListAllAccountAccessRules : List all instance level firewall access rules
// List all instance level firewall access rules.
func (firewallAccessRules *FirewallAccessRulesV1) ListAllAccountAccessRules(listAllAccountAccessRulesOptions *ListAllAccountAccessRulesOptions) (result *ListAccountAccessRulesResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listAllAccountAccessRulesOptions, "listAllAccountAccessRulesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "firewall/access_rules/rules"}
	pathParameters := []string{*firewallAccessRules.Crn}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(firewallAccessRules.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAllAccountAccessRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("firewall_access_rules", "V1", "ListAllAccountAccessRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAllAccountAccessRulesOptions.Notes != nil {
		builder.AddQuery("notes", fmt.Sprint(*listAllAccountAccessRulesOptions.Notes))
	}
	if listAllAccountAccessRulesOptions.Mode != nil {
		builder.AddQuery("mode", fmt.Sprint(*listAllAccountAccessRulesOptions.Mode))
	}
	if listAllAccountAccessRulesOptions.ConfigurationTarget != nil {
		builder.AddQuery("configuration.target", fmt.Sprint(*listAllAccountAccessRulesOptions.ConfigurationTarget))
	}
	if listAllAccountAccessRulesOptions.ConfigurationValue != nil {
		builder.AddQuery("configuration.value", fmt.Sprint(*listAllAccountAccessRulesOptions.ConfigurationValue))
	}
	if listAllAccountAccessRulesOptions.Page != nil {
		builder.AddQuery("page", fmt.Sprint(*listAllAccountAccessRulesOptions.Page))
	}
	if listAllAccountAccessRulesOptions.PerPage != nil {
		builder.AddQuery("per_page", fmt.Sprint(*listAllAccountAccessRulesOptions.PerPage))
	}
	if listAllAccountAccessRulesOptions.Order != nil {
		builder.AddQuery("order", fmt.Sprint(*listAllAccountAccessRulesOptions.Order))
	}
	if listAllAccountAccessRulesOptions.Direction != nil {
		builder.AddQuery("direction", fmt.Sprint(*listAllAccountAccessRulesOptions.Direction))
	}
	if listAllAccountAccessRulesOptions.Match != nil {
		builder.AddQuery("match", fmt.Sprint(*listAllAccountAccessRulesOptions.Match))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = firewallAccessRules.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListAccountAccessRulesResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateAccountAccessRule : Create an instance level firewall access rule
// Create a new instance level firewall access rule for a given service instance.
func (firewallAccessRules *FirewallAccessRulesV1) CreateAccountAccessRule(createAccountAccessRuleOptions *CreateAccountAccessRuleOptions) (result *AccountAccessRuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createAccountAccessRuleOptions, "createAccountAccessRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "firewall/access_rules/rules"}
	pathParameters := []string{*firewallAccessRules.Crn}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(firewallAccessRules.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAccountAccessRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("firewall_access_rules", "V1", "CreateAccountAccessRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createAccountAccessRuleOptions.Mode != nil {
		body["mode"] = createAccountAccessRuleOptions.Mode
	}
	if createAccountAccessRuleOptions.Notes != nil {
		body["notes"] = createAccountAccessRuleOptions.Notes
	}
	if createAccountAccessRuleOptions.Configuration != nil {
		body["configuration"] = createAccountAccessRuleOptions.Configuration
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
	response, err = firewallAccessRules.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountAccessRuleResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteAccountAccessRule : Delete an instance level access rule
// Delete an instance level access rule given its id.
func (firewallAccessRules *FirewallAccessRulesV1) DeleteAccountAccessRule(deleteAccountAccessRuleOptions *DeleteAccountAccessRuleOptions) (result *DeleteAccountAccessRuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAccountAccessRuleOptions, "deleteAccountAccessRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAccountAccessRuleOptions, "deleteAccountAccessRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "firewall/access_rules/rules"}
	pathParameters := []string{*firewallAccessRules.Crn, *deleteAccountAccessRuleOptions.AccessruleIdentifier}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(firewallAccessRules.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAccountAccessRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("firewall_access_rules", "V1", "DeleteAccountAccessRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = firewallAccessRules.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteAccountAccessRuleResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetAccountAccessRule : Get the details of an instance level
// Get the details of an instance level firewall access rule for a given  service instance.
func (firewallAccessRules *FirewallAccessRulesV1) GetAccountAccessRule(getAccountAccessRuleOptions *GetAccountAccessRuleOptions) (result *AccountAccessRuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountAccessRuleOptions, "getAccountAccessRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountAccessRuleOptions, "getAccountAccessRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "firewall/access_rules/rules"}
	pathParameters := []string{*firewallAccessRules.Crn, *getAccountAccessRuleOptions.AccessruleIdentifier}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(firewallAccessRules.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountAccessRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("firewall_access_rules", "V1", "GetAccountAccessRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = firewallAccessRules.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountAccessRuleResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateAccountAccessRule : Update an instance level firewall access rule
// Update an existing instance level firewall access rule for a given service instance.
func (firewallAccessRules *FirewallAccessRulesV1) UpdateAccountAccessRule(updateAccountAccessRuleOptions *UpdateAccountAccessRuleOptions) (result *AccountAccessRuleResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccountAccessRuleOptions, "updateAccountAccessRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccountAccessRuleOptions, "updateAccountAccessRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "firewall/access_rules/rules"}
	pathParameters := []string{*firewallAccessRules.Crn, *updateAccountAccessRuleOptions.AccessruleIdentifier}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(firewallAccessRules.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAccountAccessRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("firewall_access_rules", "V1", "UpdateAccountAccessRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAccountAccessRuleOptions.Mode != nil {
		body["mode"] = updateAccountAccessRuleOptions.Mode
	}
	if updateAccountAccessRuleOptions.Notes != nil {
		body["notes"] = updateAccountAccessRuleOptions.Notes
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
	response, err = firewallAccessRules.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountAccessRuleResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AccountAccessRuleInputConfiguration : Configuration object specifying access rule.
type AccountAccessRuleInputConfiguration struct {
	// The request property to target.
	Target *string `json:"target" validate:"required"`

	// The value for the selected target.For ip the value is a valid ip address.For ip_range the value specifies ip range
	// limited to /16 and /24. For asn the value is an AS number. For country the value is a country code for the country.
	Value *string `json:"value" validate:"required"`
}

// Constants associated with the AccountAccessRuleInputConfiguration.Target property.
// The request property to target.
const (
	AccountAccessRuleInputConfiguration_Target_Asn     = "asn"
	AccountAccessRuleInputConfiguration_Target_Country = "country"
	AccountAccessRuleInputConfiguration_Target_Ip      = "ip"
	AccountAccessRuleInputConfiguration_Target_IpRange = "ip_range"
)

// NewAccountAccessRuleInputConfiguration : Instantiate AccountAccessRuleInputConfiguration (Generic Model Constructor)
func (*FirewallAccessRulesV1) NewAccountAccessRuleInputConfiguration(target string, value string) (model *AccountAccessRuleInputConfiguration, err error) {
	model = &AccountAccessRuleInputConfiguration{
		Target: core.StringPtr(target),
		Value:  core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalAccountAccessRuleInputConfiguration unmarshals an instance of AccountAccessRuleInputConfiguration from the specified map of raw messages.
func UnmarshalAccountAccessRuleInputConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountAccessRuleInputConfiguration)
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccountAccessRuleObjectConfiguration : configuration.
type AccountAccessRuleObjectConfiguration struct {
	// target ip address.
	Target *string `json:"target" validate:"required"`

	// Value for the given target. For ip the value is a valid ip address.For ip_range the value specifies ip range limited
	// to /16 and /24. For asn the value is an AS number. For country the value is a country code for the country.
	Value *string `json:"value" validate:"required"`
}

// Constants associated with the AccountAccessRuleObjectConfiguration.Target property.
// target ip address.
const (
	AccountAccessRuleObjectConfiguration_Target_Asn     = "asn"
	AccountAccessRuleObjectConfiguration_Target_Country = "country"
	AccountAccessRuleObjectConfiguration_Target_Ip      = "ip"
	AccountAccessRuleObjectConfiguration_Target_IpRange = "ip_range"
)

// UnmarshalAccountAccessRuleObjectConfiguration unmarshals an instance of AccountAccessRuleObjectConfiguration from the specified map of raw messages.
func UnmarshalAccountAccessRuleObjectConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountAccessRuleObjectConfiguration)
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccountAccessRuleObjectScope : The scope definition of the access rule.
type AccountAccessRuleObjectScope struct {
	// The scope of the access rule.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the AccountAccessRuleObjectScope.Type property.
// The scope of the access rule.
const (
	AccountAccessRuleObjectScope_Type_Account      = "account"
	AccountAccessRuleObjectScope_Type_Organization = "organization"
)

// UnmarshalAccountAccessRuleObjectScope unmarshals an instance of AccountAccessRuleObjectScope from the specified map of raw messages.
func UnmarshalAccountAccessRuleObjectScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountAccessRuleObjectScope)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateAccountAccessRuleOptions : The CreateAccountAccessRule options.
type CreateAccountAccessRuleOptions struct {
	// The action to apply to a matched request.
	Mode *string `json:"mode,omitempty"`

	// A personal note about the rule. Typically used as a reminder or explanation for the rule.
	Notes *string `json:"notes,omitempty"`

	// Configuration object specifying access rule.
	Configuration *AccountAccessRuleInputConfiguration `json:"configuration,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateAccountAccessRuleOptions.Mode property.
// The action to apply to a matched request.
const (
	CreateAccountAccessRuleOptions_Mode_Block       = "block"
	CreateAccountAccessRuleOptions_Mode_Challenge   = "challenge"
	CreateAccountAccessRuleOptions_Mode_JsChallenge = "js_challenge"
	CreateAccountAccessRuleOptions_Mode_Whitelist   = "whitelist"
)

// NewCreateAccountAccessRuleOptions : Instantiate CreateAccountAccessRuleOptions
func (*FirewallAccessRulesV1) NewCreateAccountAccessRuleOptions() *CreateAccountAccessRuleOptions {
	return &CreateAccountAccessRuleOptions{}
}

// SetMode : Allow user to set Mode
func (options *CreateAccountAccessRuleOptions) SetMode(mode string) *CreateAccountAccessRuleOptions {
	options.Mode = core.StringPtr(mode)
	return options
}

// SetNotes : Allow user to set Notes
func (options *CreateAccountAccessRuleOptions) SetNotes(notes string) *CreateAccountAccessRuleOptions {
	options.Notes = core.StringPtr(notes)
	return options
}

// SetConfiguration : Allow user to set Configuration
func (options *CreateAccountAccessRuleOptions) SetConfiguration(configuration *AccountAccessRuleInputConfiguration) *CreateAccountAccessRuleOptions {
	options.Configuration = configuration
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAccountAccessRuleOptions) SetHeaders(param map[string]string) *CreateAccountAccessRuleOptions {
	options.Headers = param
	return options
}

// DeleteAccountAccessRuleOptions : The DeleteAccountAccessRule options.
type DeleteAccountAccessRuleOptions struct {
	// Identifier of the access rule to be deleted.
	AccessruleIdentifier *string `json:"accessrule_identifier" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAccountAccessRuleOptions : Instantiate DeleteAccountAccessRuleOptions
func (*FirewallAccessRulesV1) NewDeleteAccountAccessRuleOptions(accessruleIdentifier string) *DeleteAccountAccessRuleOptions {
	return &DeleteAccountAccessRuleOptions{
		AccessruleIdentifier: core.StringPtr(accessruleIdentifier),
	}
}

// SetAccessruleIdentifier : Allow user to set AccessruleIdentifier
func (options *DeleteAccountAccessRuleOptions) SetAccessruleIdentifier(accessruleIdentifier string) *DeleteAccountAccessRuleOptions {
	options.AccessruleIdentifier = core.StringPtr(accessruleIdentifier)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAccountAccessRuleOptions) SetHeaders(param map[string]string) *DeleteAccountAccessRuleOptions {
	options.Headers = param
	return options
}

// DeleteAccountAccessRuleRespResult : Container for response information.
type DeleteAccountAccessRuleRespResult struct {
	// ID.
	ID *string `json:"id" validate:"required"`
}

// UnmarshalDeleteAccountAccessRuleRespResult unmarshals an instance of DeleteAccountAccessRuleRespResult from the specified map of raw messages.
func UnmarshalDeleteAccountAccessRuleRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteAccountAccessRuleRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAccountAccessRuleOptions : The GetAccountAccessRule options.
type GetAccountAccessRuleOptions struct {
	// Identifier of firewall access rule for the given zone.
	AccessruleIdentifier *string `json:"accessrule_identifier" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountAccessRuleOptions : Instantiate GetAccountAccessRuleOptions
func (*FirewallAccessRulesV1) NewGetAccountAccessRuleOptions(accessruleIdentifier string) *GetAccountAccessRuleOptions {
	return &GetAccountAccessRuleOptions{
		AccessruleIdentifier: core.StringPtr(accessruleIdentifier),
	}
}

// SetAccessruleIdentifier : Allow user to set AccessruleIdentifier
func (options *GetAccountAccessRuleOptions) SetAccessruleIdentifier(accessruleIdentifier string) *GetAccountAccessRuleOptions {
	options.AccessruleIdentifier = core.StringPtr(accessruleIdentifier)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountAccessRuleOptions) SetHeaders(param map[string]string) *GetAccountAccessRuleOptions {
	options.Headers = param
	return options
}

// ListAccountAccessRulesRespResultInfo : Statistics of results.
type ListAccountAccessRulesRespResultInfo struct {
	// Page number.
	Page *int64 `json:"page" validate:"required"`

	// Number of results per page.
	PerPage *int64 `json:"per_page" validate:"required"`

	// Number of results.
	Count *int64 `json:"count" validate:"required"`

	// Total number of results.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalListAccountAccessRulesRespResultInfo unmarshals an instance of ListAccountAccessRulesRespResultInfo from the specified map of raw messages.
func UnmarshalListAccountAccessRulesRespResultInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListAccountAccessRulesRespResultInfo)
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

// ListAllAccountAccessRulesOptions : The ListAllAccountAccessRules options.
type ListAllAccountAccessRulesOptions struct {
	// Search access rules by note.(Not case sensitive).
	Notes *string `json:"notes,omitempty"`

	// Search access rules by mode.
	Mode *string `json:"mode,omitempty"`

	// Search access rules by configuration target.
	ConfigurationTarget *string `json:"configuration.target,omitempty"`

	// Search access rules by configuration value which can be IP, IPrange, or country code.
	ConfigurationValue *string `json:"configuration.value,omitempty"`

	// Page number of paginated results.
	Page *int64 `json:"page,omitempty"`

	// Maximum number of access rules per page.
	PerPage *int64 `json:"per_page,omitempty"`

	// Field by which to order list of access rules.
	Order *string `json:"order,omitempty"`

	// Direction in which to order results [ascending/descending order].
	Direction *string `json:"direction,omitempty"`

	// Whether to match all (all) or atleast one search parameter (any).
	Match *string `json:"match,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListAllAccountAccessRulesOptions.Mode property.
// Search access rules by mode.
const (
	ListAllAccountAccessRulesOptions_Mode_Block       = "block"
	ListAllAccountAccessRulesOptions_Mode_Challenge   = "challenge"
	ListAllAccountAccessRulesOptions_Mode_JsChallenge = "js_challenge"
	ListAllAccountAccessRulesOptions_Mode_Whitelist   = "whitelist"
)

// Constants associated with the ListAllAccountAccessRulesOptions.ConfigurationTarget property.
// Search access rules by configuration target.
const (
	ListAllAccountAccessRulesOptions_ConfigurationTarget_Asn     = "asn"
	ListAllAccountAccessRulesOptions_ConfigurationTarget_Country = "country"
	ListAllAccountAccessRulesOptions_ConfigurationTarget_Ip      = "ip"
	ListAllAccountAccessRulesOptions_ConfigurationTarget_IpRange = "ip_range"
)

// Constants associated with the ListAllAccountAccessRulesOptions.Order property.
// Field by which to order list of access rules.
const (
	ListAllAccountAccessRulesOptions_Order_Mode   = "mode"
	ListAllAccountAccessRulesOptions_Order_Target = "target"
	ListAllAccountAccessRulesOptions_Order_Value  = "value"
)

// Constants associated with the ListAllAccountAccessRulesOptions.Direction property.
// Direction in which to order results [ascending/descending order].
const (
	ListAllAccountAccessRulesOptions_Direction_Asc  = "asc"
	ListAllAccountAccessRulesOptions_Direction_Desc = "desc"
)

// Constants associated with the ListAllAccountAccessRulesOptions.Match property.
// Whether to match all (all) or atleast one search parameter (any).
const (
	ListAllAccountAccessRulesOptions_Match_All = "all"
	ListAllAccountAccessRulesOptions_Match_Any = "any"
)

// NewListAllAccountAccessRulesOptions : Instantiate ListAllAccountAccessRulesOptions
func (*FirewallAccessRulesV1) NewListAllAccountAccessRulesOptions() *ListAllAccountAccessRulesOptions {
	return &ListAllAccountAccessRulesOptions{}
}

// SetNotes : Allow user to set Notes
func (options *ListAllAccountAccessRulesOptions) SetNotes(notes string) *ListAllAccountAccessRulesOptions {
	options.Notes = core.StringPtr(notes)
	return options
}

// SetMode : Allow user to set Mode
func (options *ListAllAccountAccessRulesOptions) SetMode(mode string) *ListAllAccountAccessRulesOptions {
	options.Mode = core.StringPtr(mode)
	return options
}

// SetConfigurationTarget : Allow user to set ConfigurationTarget
func (options *ListAllAccountAccessRulesOptions) SetConfigurationTarget(configurationTarget string) *ListAllAccountAccessRulesOptions {
	options.ConfigurationTarget = core.StringPtr(configurationTarget)
	return options
}

// SetConfigurationValue : Allow user to set ConfigurationValue
func (options *ListAllAccountAccessRulesOptions) SetConfigurationValue(configurationValue string) *ListAllAccountAccessRulesOptions {
	options.ConfigurationValue = core.StringPtr(configurationValue)
	return options
}

// SetPage : Allow user to set Page
func (options *ListAllAccountAccessRulesOptions) SetPage(page int64) *ListAllAccountAccessRulesOptions {
	options.Page = core.Int64Ptr(page)
	return options
}

// SetPerPage : Allow user to set PerPage
func (options *ListAllAccountAccessRulesOptions) SetPerPage(perPage int64) *ListAllAccountAccessRulesOptions {
	options.PerPage = core.Int64Ptr(perPage)
	return options
}

// SetOrder : Allow user to set Order
func (options *ListAllAccountAccessRulesOptions) SetOrder(order string) *ListAllAccountAccessRulesOptions {
	options.Order = core.StringPtr(order)
	return options
}

// SetDirection : Allow user to set Direction
func (options *ListAllAccountAccessRulesOptions) SetDirection(direction string) *ListAllAccountAccessRulesOptions {
	options.Direction = core.StringPtr(direction)
	return options
}

// SetMatch : Allow user to set Match
func (options *ListAllAccountAccessRulesOptions) SetMatch(match string) *ListAllAccountAccessRulesOptions {
	options.Match = core.StringPtr(match)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAllAccountAccessRulesOptions) SetHeaders(param map[string]string) *ListAllAccountAccessRulesOptions {
	options.Headers = param
	return options
}

// UpdateAccountAccessRuleOptions : The UpdateAccountAccessRule options.
type UpdateAccountAccessRuleOptions struct {
	// Identifier of firewall access rule.
	AccessruleIdentifier *string `json:"accessrule_identifier" validate:"required"`

	// The action to apply to a matched request.
	Mode *string `json:"mode,omitempty"`

	// A personal note about the rule. Typically used as a reminder or explanation for the rule.
	Notes *string `json:"notes,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateAccountAccessRuleOptions.Mode property.
// The action to apply to a matched request.
const (
	UpdateAccountAccessRuleOptions_Mode_Block       = "block"
	UpdateAccountAccessRuleOptions_Mode_Challenge   = "challenge"
	UpdateAccountAccessRuleOptions_Mode_JsChallenge = "js_challenge"
	UpdateAccountAccessRuleOptions_Mode_Whitelist   = "whitelist"
)

// NewUpdateAccountAccessRuleOptions : Instantiate UpdateAccountAccessRuleOptions
func (*FirewallAccessRulesV1) NewUpdateAccountAccessRuleOptions(accessruleIdentifier string) *UpdateAccountAccessRuleOptions {
	return &UpdateAccountAccessRuleOptions{
		AccessruleIdentifier: core.StringPtr(accessruleIdentifier),
	}
}

// SetAccessruleIdentifier : Allow user to set AccessruleIdentifier
func (options *UpdateAccountAccessRuleOptions) SetAccessruleIdentifier(accessruleIdentifier string) *UpdateAccountAccessRuleOptions {
	options.AccessruleIdentifier = core.StringPtr(accessruleIdentifier)
	return options
}

// SetMode : Allow user to set Mode
func (options *UpdateAccountAccessRuleOptions) SetMode(mode string) *UpdateAccountAccessRuleOptions {
	options.Mode = core.StringPtr(mode)
	return options
}

// SetNotes : Allow user to set Notes
func (options *UpdateAccountAccessRuleOptions) SetNotes(notes string) *UpdateAccountAccessRuleOptions {
	options.Notes = core.StringPtr(notes)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccountAccessRuleOptions) SetHeaders(param map[string]string) *UpdateAccountAccessRuleOptions {
	options.Headers = param
	return options
}

// AccountAccessRuleObject : access rule objects.
type AccountAccessRuleObject struct {
	// Identifier of the instance level firewall access rule.
	ID *string `json:"id" validate:"required"`

	// A personal note about the rule. Typically used as a reminder or explanation for the rule.
	Notes *string `json:"notes" validate:"required"`

	// List of modes that are allowed.
	AllowedModes []string `json:"allowed_modes" validate:"required"`

	// The action to be applied to a request matching the instance level access rule.
	Mode *string `json:"mode" validate:"required"`

	// The scope definition of the access rule.
	Scope *AccountAccessRuleObjectScope `json:"scope,omitempty"`

	// The creation date-time of the instance level firewall access rule.
	CreatedOn *strfmt.DateTime `json:"created_on" validate:"required"`

	// The modification date-time of the instance level firewall access rule.
	ModifiedOn *strfmt.DateTime `json:"modified_on" validate:"required"`

	// configuration.
	Configuration *AccountAccessRuleObjectConfiguration `json:"configuration" validate:"required"`
}

// Constants associated with the AccountAccessRuleObject.AllowedModes property.
const (
	AccountAccessRuleObject_AllowedModes_Block       = "block"
	AccountAccessRuleObject_AllowedModes_Challenge   = "challenge"
	AccountAccessRuleObject_AllowedModes_JsChallenge = "js_challenge"
	AccountAccessRuleObject_AllowedModes_Whitelist   = "whitelist"
)

// Constants associated with the AccountAccessRuleObject.Mode property.
// The action to be applied to a request matching the instance level access rule.
const (
	AccountAccessRuleObject_Mode_Block       = "block"
	AccountAccessRuleObject_Mode_Challenge   = "challenge"
	AccountAccessRuleObject_Mode_JsChallenge = "js_challenge"
	AccountAccessRuleObject_Mode_Whitelist   = "whitelist"
)

// UnmarshalAccountAccessRuleObject unmarshals an instance of AccountAccessRuleObject from the specified map of raw messages.
func UnmarshalAccountAccessRuleObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountAccessRuleObject)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "notes", &obj.Notes)
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
	err = core.UnmarshalModel(m, "scope", &obj.Scope, UnmarshalAccountAccessRuleObjectScope)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configuration", &obj.Configuration, UnmarshalAccountAccessRuleObjectConfiguration)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccountAccessRuleResp : access rule response output.
type AccountAccessRuleResp struct {
	// Operation success flag.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages encountered.
	Messages [][]string `json:"messages" validate:"required"`

	// access rule objects.
	Result *AccountAccessRuleObject `json:"result" validate:"required"`
}

// UnmarshalAccountAccessRuleResp unmarshals an instance of AccountAccessRuleResp from the specified map of raw messages.
func UnmarshalAccountAccessRuleResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountAccessRuleResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalAccountAccessRuleObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteAccountAccessRuleResp : delete access rule response.
type DeleteAccountAccessRuleResp struct {
	// Operation success flag.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages encountered.
	Messages [][]string `json:"messages" validate:"required"`

	// Container for response information.
	Result *DeleteAccountAccessRuleRespResult `json:"result" validate:"required"`
}

// UnmarshalDeleteAccountAccessRuleResp unmarshals an instance of DeleteAccountAccessRuleResp from the specified map of raw messages.
func UnmarshalDeleteAccountAccessRuleResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteAccountAccessRuleResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalDeleteAccountAccessRuleRespResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListAccountAccessRulesResp : access rule list response.
type ListAccountAccessRulesResp struct {
	// Operation success flag.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages encountered.
	Messages [][]string `json:"messages" validate:"required"`

	// Container for response information.
	Result []AccountAccessRuleObject `json:"result" validate:"required"`

	// Statistics of results.
	ResultInfo *ListAccountAccessRulesRespResultInfo `json:"result_info" validate:"required"`
}

// UnmarshalListAccountAccessRulesResp unmarshals an instance of ListAccountAccessRulesResp from the specified map of raw messages.
func UnmarshalListAccountAccessRulesResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListAccountAccessRulesResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalAccountAccessRuleObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result_info", &obj.ResultInfo, UnmarshalListAccountAccessRulesRespResultInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
