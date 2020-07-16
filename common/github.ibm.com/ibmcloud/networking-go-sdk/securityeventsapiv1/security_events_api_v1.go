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

// Package securityeventsapiv1 : Operations and models for the SecurityEventsApiV1 service
package securityeventsapiv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	common "github.ibm.com/ibmcloud/networking-go-sdk/common"
	"reflect"
)

// SecurityEventsApiV1 : Security Events API
//
// Version: 1.0.0
type SecurityEventsApiV1 struct {
	Service *core.BaseService

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string

	// zone identifier.
	ZoneID *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "security_events_api"

// SecurityEventsApiV1Options : Service options
type SecurityEventsApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string `validate:"required"`

	// zone identifier.
	ZoneID *string `validate:"required"`
}

// NewSecurityEventsApiV1UsingExternalConfig : constructs an instance of SecurityEventsApiV1 with passed in options and external configuration.
func NewSecurityEventsApiV1UsingExternalConfig(options *SecurityEventsApiV1Options) (securityEventsApi *SecurityEventsApiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	securityEventsApi, err = NewSecurityEventsApiV1(options)
	if err != nil {
		return
	}

	err = securityEventsApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = securityEventsApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewSecurityEventsApiV1 : constructs an instance of SecurityEventsApiV1 with passed in options.
func NewSecurityEventsApiV1(options *SecurityEventsApiV1Options) (service *SecurityEventsApiV1, err error) {
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

	service = &SecurityEventsApiV1{
		Service: baseService,
		Crn:     options.Crn,
		ZoneID:  options.ZoneID,
	}

	return
}

// SetServiceURL sets the service URL
func (securityEventsApi *SecurityEventsApiV1) SetServiceURL(url string) error {
	return securityEventsApi.Service.SetServiceURL(url)
}

// SecurityEvents : Logs of the mitigations performed by Firewall features
// Provides a full log of the mitigations performed by the CIS Firewall features including; Firewall Rules, Rate
// Limiting, Security Level, Access Rules (IP, IP Range, ASN, and Country), WAF (Web Application Firewall), User Agent
// Blocking, Zone Lockdown, and Advanced DDoS Protection.
func (securityEventsApi *SecurityEventsApiV1) SecurityEvents(securityEventsOptions *SecurityEventsOptions) (result *SecurityEvents, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(securityEventsOptions, "securityEventsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1", "zones", "security/events"}
	pathParameters := []string{*securityEventsApi.Crn, *securityEventsApi.ZoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(securityEventsApi.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range securityEventsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("security_events_api", "V1", "SecurityEvents")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if securityEventsOptions.IpClass != nil {
		builder.AddQuery("ip_class", fmt.Sprint(*securityEventsOptions.IpClass))
	}
	if securityEventsOptions.Method != nil {
		builder.AddQuery("method", fmt.Sprint(*securityEventsOptions.Method))
	}
	if securityEventsOptions.Scheme != nil {
		builder.AddQuery("scheme", fmt.Sprint(*securityEventsOptions.Scheme))
	}
	if securityEventsOptions.Ip != nil {
		builder.AddQuery("ip", fmt.Sprint(*securityEventsOptions.Ip))
	}
	if securityEventsOptions.Host != nil {
		builder.AddQuery("host", fmt.Sprint(*securityEventsOptions.Host))
	}
	if securityEventsOptions.Proto != nil {
		builder.AddQuery("proto", fmt.Sprint(*securityEventsOptions.Proto))
	}
	if securityEventsOptions.URI != nil {
		builder.AddQuery("uri", fmt.Sprint(*securityEventsOptions.URI))
	}
	if securityEventsOptions.Ua != nil {
		builder.AddQuery("ua", fmt.Sprint(*securityEventsOptions.Ua))
	}
	if securityEventsOptions.Colo != nil {
		builder.AddQuery("colo", fmt.Sprint(*securityEventsOptions.Colo))
	}
	if securityEventsOptions.RayID != nil {
		builder.AddQuery("ray_id", fmt.Sprint(*securityEventsOptions.RayID))
	}
	if securityEventsOptions.Kind != nil {
		builder.AddQuery("kind", fmt.Sprint(*securityEventsOptions.Kind))
	}
	if securityEventsOptions.Action != nil {
		builder.AddQuery("action", fmt.Sprint(*securityEventsOptions.Action))
	}
	if securityEventsOptions.Cursor != nil {
		builder.AddQuery("cursor", fmt.Sprint(*securityEventsOptions.Cursor))
	}
	if securityEventsOptions.Country != nil {
		builder.AddQuery("country", fmt.Sprint(*securityEventsOptions.Country))
	}
	if securityEventsOptions.Since != nil {
		builder.AddQuery("since", fmt.Sprint(*securityEventsOptions.Since))
	}
	if securityEventsOptions.Source != nil {
		builder.AddQuery("source", fmt.Sprint(*securityEventsOptions.Source))
	}
	if securityEventsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*securityEventsOptions.Limit))
	}
	if securityEventsOptions.RuleID != nil {
		builder.AddQuery("rule_id", fmt.Sprint(*securityEventsOptions.RuleID))
	}
	if securityEventsOptions.Until != nil {
		builder.AddQuery("until", fmt.Sprint(*securityEventsOptions.Until))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = securityEventsApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSecurityEvents)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ResultInfoCursors : Cursor positions of the security events.
type ResultInfoCursors struct {
	// The events in the response is after this cursor position.
	After *string `json:"after" validate:"required"`

	// The events in the response is before this cursor position.
	Before *string `json:"before" validate:"required"`
}

// UnmarshalResultInfoCursors unmarshals an instance of ResultInfoCursors from the specified map of raw messages.
func UnmarshalResultInfoCursors(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultInfoCursors)
	err = core.UnmarshalPrimitive(m, "after", &obj.After)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "before", &obj.Before)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResultInfoScannedRange : The time window of the events.
type ResultInfoScannedRange struct {
	// Start date and time of the events.
	Since *string `json:"since" validate:"required"`

	// End date and time of the events.
	Until *string `json:"until" validate:"required"`
}

// UnmarshalResultInfoScannedRange unmarshals an instance of ResultInfoScannedRange from the specified map of raw messages.
func UnmarshalResultInfoScannedRange(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultInfoScannedRange)
	err = core.UnmarshalPrimitive(m, "since", &obj.Since)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "until", &obj.Until)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityEventObjectMatchesItem : SecurityEventObjectMatchesItem struct
type SecurityEventObjectMatchesItem struct {
	// The ID of the rule that triggered the event, should be considered in the context of source.
	RuleID *string `json:"rule_id" validate:"required"`

	// Source of the event.
	Source *string `json:"source" validate:"required"`

	// What type of action was taken.
	Action *string `json:"action" validate:"required"`

	// metadata.
	Metadata interface{} `json:"metadata" validate:"required"`
}

// UnmarshalSecurityEventObjectMatchesItem unmarshals an instance of SecurityEventObjectMatchesItem from the specified map of raw messages.
func UnmarshalSecurityEventObjectMatchesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityEventObjectMatchesItem)
	err = core.UnmarshalPrimitive(m, "rule_id", &obj.RuleID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityEventsOptions : The SecurityEvents options.
type SecurityEventsOptions struct {
	// IP class is a map of client IP to visitor classification.
	IpClass *string `json:"ip_class,omitempty"`

	// The HTTP method of the request.
	Method *string `json:"method,omitempty"`

	// The scheme of the uri.
	Scheme *string `json:"scheme,omitempty"`

	// The IPv4 or IPv6 address from which the request originated.
	Ip *string `json:"ip,omitempty"`

	// The hostname the request attempted to access.
	Host *string `json:"host,omitempty"`

	// The protocol of the request.
	Proto *string `json:"proto,omitempty"`

	// The URI requested from the hostname.
	URI *string `json:"uri,omitempty"`

	// The client user agent that initiated the request.
	Ua *string `json:"ua,omitempty"`

	// The 3-letter CF PoP code.
	Colo *string `json:"colo,omitempty"`

	// Ray ID of the request.
	RayID *string `json:"ray_id,omitempty"`

	// Kind of events. Now it is only firewall.
	Kind *string `json:"kind,omitempty"`

	// What type of action was taken.
	Action *string `json:"action,omitempty"`

	// Cursor position and direction for requesting next set of records when amount of results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the cursors object in the result_info structure.
	Cursor *string `json:"cursor,omitempty"`

	// The 2-digit country code in which the request originated.
	Country *string `json:"country,omitempty"`

	// Start date and time of requesting data period in the ISO8601 format. Can't go back more than a year.
	Since *strfmt.DateTime `json:"since,omitempty"`

	// Source of the event.
	Source *string `json:"source,omitempty"`

	// The number of events to return.
	Limit *int64 `json:"limit,omitempty"`

	// The ID of the rule that triggered the event, should be considered in the context of source.
	RuleID *string `json:"rule_id,omitempty"`

	// End date and time of requesting data period in the ISO8601 format.
	Until *strfmt.DateTime `json:"until,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the SecurityEventsOptions.IpClass property.
// IP class is a map of client IP to visitor classification.
const (
	SecurityEventsOptions_IpClass_Backupservice     = "backupService"
	SecurityEventsOptions_IpClass_Badhost           = "badHost"
	SecurityEventsOptions_IpClass_Clean             = "clean"
	SecurityEventsOptions_IpClass_Greylist          = "greylist"
	SecurityEventsOptions_IpClass_Mobileplatform    = "mobilePlatform"
	SecurityEventsOptions_IpClass_Monitoringservice = "monitoringService"
	SecurityEventsOptions_IpClass_Norecord          = "noRecord"
	SecurityEventsOptions_IpClass_Scan              = "scan"
	SecurityEventsOptions_IpClass_Searchengine      = "searchEngine"
	SecurityEventsOptions_IpClass_Securityscanner   = "securityScanner"
	SecurityEventsOptions_IpClass_Tor               = "tor"
	SecurityEventsOptions_IpClass_Unknown           = "unknown"
	SecurityEventsOptions_IpClass_Whitelist         = "whitelist"
)

// Constants associated with the SecurityEventsOptions.Method property.
// The HTTP method of the request.
const (
	SecurityEventsOptions_Method_Acl             = "ACL"
	SecurityEventsOptions_Method_BaselineControl = "BASELINE-CONTROL"
	SecurityEventsOptions_Method_Bcopy           = "BCOPY"
	SecurityEventsOptions_Method_Bdelete         = "BDELETE"
	SecurityEventsOptions_Method_Bmove           = "BMOVE"
	SecurityEventsOptions_Method_Bpropfind       = "BPROPFIND"
	SecurityEventsOptions_Method_Bproppatch      = "BPROPPATCH"
	SecurityEventsOptions_Method_Checkin         = "CHECKIN"
	SecurityEventsOptions_Method_Checkout        = "CHECKOUT"
	SecurityEventsOptions_Method_Connect         = "CONNECT"
	SecurityEventsOptions_Method_Cook            = "COOK"
	SecurityEventsOptions_Method_Copy            = "COPY"
	SecurityEventsOptions_Method_Delete          = "DELETE"
	SecurityEventsOptions_Method_Get             = "GET"
	SecurityEventsOptions_Method_Head            = "HEAD"
	SecurityEventsOptions_Method_JSON            = "JSON"
	SecurityEventsOptions_Method_Label           = "LABEL"
	SecurityEventsOptions_Method_Lock            = "LOCK"
	SecurityEventsOptions_Method_Merge           = "MERGE"
	SecurityEventsOptions_Method_Mkactivity      = "MKACTIVITY"
	SecurityEventsOptions_Method_Mkcol           = "MKCOL"
	SecurityEventsOptions_Method_Mkworkspace     = "MKWORKSPACE"
	SecurityEventsOptions_Method_Move            = "MOVE"
	SecurityEventsOptions_Method_Notify          = "NOTIFY"
	SecurityEventsOptions_Method_Options         = "OPTIONS"
	SecurityEventsOptions_Method_Orderpatch      = "ORDERPATCH"
	SecurityEventsOptions_Method_Patch           = "PATCH"
	SecurityEventsOptions_Method_Poll            = "POLL"
	SecurityEventsOptions_Method_Post            = "POST"
	SecurityEventsOptions_Method_Propfind        = "PROPFIND"
	SecurityEventsOptions_Method_Proppatch       = "PROPPATCH"
	SecurityEventsOptions_Method_Purge           = "PURGE"
	SecurityEventsOptions_Method_Put             = "PUT"
	SecurityEventsOptions_Method_Report          = "REPORT"
	SecurityEventsOptions_Method_RpcInData       = "RPC_IN_DATA"
	SecurityEventsOptions_Method_RpcOutData      = "RPC_OUT_DATA"
	SecurityEventsOptions_Method_Search          = "SEARCH"
	SecurityEventsOptions_Method_Subscribe       = "SUBSCRIBE"
	SecurityEventsOptions_Method_Trace           = "TRACE"
	SecurityEventsOptions_Method_Track           = "TRACK"
	SecurityEventsOptions_Method_Uncheckout      = "UNCHECKOUT"
	SecurityEventsOptions_Method_Unlock          = "UNLOCK"
	SecurityEventsOptions_Method_Unsubscribe     = "UNSUBSCRIBE"
	SecurityEventsOptions_Method_Update          = "UPDATE"
	SecurityEventsOptions_Method_VersionControl  = "VERSION-CONTROL"
	SecurityEventsOptions_Method_XMsEnumatts     = "X-MS-ENUMATTS"
)

// Constants associated with the SecurityEventsOptions.Scheme property.
// The scheme of the uri.
const (
	SecurityEventsOptions_Scheme_Http    = "http"
	SecurityEventsOptions_Scheme_Https   = "https"
	SecurityEventsOptions_Scheme_Unknown = "unknown"
)

// Constants associated with the SecurityEventsOptions.Proto property.
// The protocol of the request.
const (
	SecurityEventsOptions_Proto_Http10 = "HTTP/1.0"
	SecurityEventsOptions_Proto_Http11 = "HTTP/1.1"
	SecurityEventsOptions_Proto_Http12 = "HTTP/1.2"
	SecurityEventsOptions_Proto_Http2  = "HTTP/2"
	SecurityEventsOptions_Proto_Spdy31 = "SPDY/3.1"
	SecurityEventsOptions_Proto_Unk    = "UNK"
)

// Constants associated with the SecurityEventsOptions.Kind property.
// Kind of events. Now it is only firewall.
const (
	SecurityEventsOptions_Kind_Firewall = "firewall"
)

// Constants associated with the SecurityEventsOptions.Action property.
// What type of action was taken.
const (
	SecurityEventsOptions_Action_Allow           = "allow"
	SecurityEventsOptions_Action_Challenge       = "challenge"
	SecurityEventsOptions_Action_Connectionclose = "connectionClose"
	SecurityEventsOptions_Action_Drop            = "drop"
	SecurityEventsOptions_Action_Jschallenge     = "jschallenge"
	SecurityEventsOptions_Action_Log             = "log"
	SecurityEventsOptions_Action_Simulate        = "simulate"
	SecurityEventsOptions_Action_Unknown         = "unknown"
)

// Constants associated with the SecurityEventsOptions.Source property.
// Source of the event.
const (
	SecurityEventsOptions_Source_Asn           = "asn"
	SecurityEventsOptions_Source_Bic           = "bic"
	SecurityEventsOptions_Source_Country       = "country"
	SecurityEventsOptions_Source_Firewallrules = "firewallRules"
	SecurityEventsOptions_Source_Hot           = "hot"
	SecurityEventsOptions_Source_Ip            = "ip"
	SecurityEventsOptions_Source_Iprange       = "ipRange"
	SecurityEventsOptions_Source_L7ddos        = "l7ddos"
	SecurityEventsOptions_Source_Ratelimit     = "rateLimit"
	SecurityEventsOptions_Source_Securitylevel = "securityLevel"
	SecurityEventsOptions_Source_Uablock       = "uaBlock"
	SecurityEventsOptions_Source_Unknown       = "unknown"
	SecurityEventsOptions_Source_Waf           = "waf"
	SecurityEventsOptions_Source_Zonelockdown  = "zoneLockdown"
)

// NewSecurityEventsOptions : Instantiate SecurityEventsOptions
func (*SecurityEventsApiV1) NewSecurityEventsOptions() *SecurityEventsOptions {
	return &SecurityEventsOptions{}
}

// SetIpClass : Allow user to set IpClass
func (options *SecurityEventsOptions) SetIpClass(ipClass string) *SecurityEventsOptions {
	options.IpClass = core.StringPtr(ipClass)
	return options
}

// SetMethod : Allow user to set Method
func (options *SecurityEventsOptions) SetMethod(method string) *SecurityEventsOptions {
	options.Method = core.StringPtr(method)
	return options
}

// SetScheme : Allow user to set Scheme
func (options *SecurityEventsOptions) SetScheme(scheme string) *SecurityEventsOptions {
	options.Scheme = core.StringPtr(scheme)
	return options
}

// SetIp : Allow user to set Ip
func (options *SecurityEventsOptions) SetIp(ip string) *SecurityEventsOptions {
	options.Ip = core.StringPtr(ip)
	return options
}

// SetHost : Allow user to set Host
func (options *SecurityEventsOptions) SetHost(host string) *SecurityEventsOptions {
	options.Host = core.StringPtr(host)
	return options
}

// SetProto : Allow user to set Proto
func (options *SecurityEventsOptions) SetProto(proto string) *SecurityEventsOptions {
	options.Proto = core.StringPtr(proto)
	return options
}

// SetURI : Allow user to set URI
func (options *SecurityEventsOptions) SetURI(uri string) *SecurityEventsOptions {
	options.URI = core.StringPtr(uri)
	return options
}

// SetUa : Allow user to set Ua
func (options *SecurityEventsOptions) SetUa(ua string) *SecurityEventsOptions {
	options.Ua = core.StringPtr(ua)
	return options
}

// SetColo : Allow user to set Colo
func (options *SecurityEventsOptions) SetColo(colo string) *SecurityEventsOptions {
	options.Colo = core.StringPtr(colo)
	return options
}

// SetRayID : Allow user to set RayID
func (options *SecurityEventsOptions) SetRayID(rayID string) *SecurityEventsOptions {
	options.RayID = core.StringPtr(rayID)
	return options
}

// SetKind : Allow user to set Kind
func (options *SecurityEventsOptions) SetKind(kind string) *SecurityEventsOptions {
	options.Kind = core.StringPtr(kind)
	return options
}

// SetAction : Allow user to set Action
func (options *SecurityEventsOptions) SetAction(action string) *SecurityEventsOptions {
	options.Action = core.StringPtr(action)
	return options
}

// SetCursor : Allow user to set Cursor
func (options *SecurityEventsOptions) SetCursor(cursor string) *SecurityEventsOptions {
	options.Cursor = core.StringPtr(cursor)
	return options
}

// SetCountry : Allow user to set Country
func (options *SecurityEventsOptions) SetCountry(country string) *SecurityEventsOptions {
	options.Country = core.StringPtr(country)
	return options
}

// SetSince : Allow user to set Since
func (options *SecurityEventsOptions) SetSince(since *strfmt.DateTime) *SecurityEventsOptions {
	options.Since = since
	return options
}

// SetSource : Allow user to set Source
func (options *SecurityEventsOptions) SetSource(source string) *SecurityEventsOptions {
	options.Source = core.StringPtr(source)
	return options
}

// SetLimit : Allow user to set Limit
func (options *SecurityEventsOptions) SetLimit(limit int64) *SecurityEventsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetRuleID : Allow user to set RuleID
func (options *SecurityEventsOptions) SetRuleID(ruleID string) *SecurityEventsOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetUntil : Allow user to set Until
func (options *SecurityEventsOptions) SetUntil(until *strfmt.DateTime) *SecurityEventsOptions {
	options.Until = until
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SecurityEventsOptions) SetHeaders(param map[string]string) *SecurityEventsOptions {
	options.Headers = param
	return options
}

// ResultInfo : Statistics of results.
type ResultInfo struct {
	// Cursor positions of the security events.
	Cursors *ResultInfoCursors `json:"cursors" validate:"required"`

	// The time window of the events.
	ScannedRange *ResultInfoScannedRange `json:"scanned_range" validate:"required"`
}

// UnmarshalResultInfo unmarshals an instance of ResultInfo from the specified map of raw messages.
func UnmarshalResultInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultInfo)
	err = core.UnmarshalModel(m, "cursors", &obj.Cursors, UnmarshalResultInfoCursors)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scanned_range", &obj.ScannedRange, UnmarshalResultInfoScannedRange)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityEventObject : Security event object.
type SecurityEventObject struct {
	// Ray ID of the request.
	RayID *string `json:"ray_id" validate:"required"`

	// Kind of events. Now it is only firewall.
	Kind *string `json:"kind" validate:"required"`

	// Source of the event.
	Source *string `json:"source" validate:"required"`

	// What type of action was taken.
	Action *string `json:"action" validate:"required"`

	// The ID of the rule that triggered the event, should be considered in the context of source.
	RuleID *string `json:"rule_id" validate:"required"`

	// The IPv4 or IPv6 address from which the request originated.
	Ip *string `json:"ip" validate:"required"`

	// IP class is a map of client IP to visitor classification.
	IpClass *string `json:"ip_class" validate:"required"`

	// The 2-digit country code in which the request originated.
	Country *string `json:"country" validate:"required"`

	// The 3-letter CF PoP code.
	Colo *string `json:"colo" validate:"required"`

	// The hostname the request attempted to access.
	Host *string `json:"host" validate:"required"`

	// The HTTP method of the request.
	Method *string `json:"method" validate:"required"`

	// The protocol of the request.
	Proto *string `json:"proto" validate:"required"`

	// The scheme of the uri.
	Scheme *string `json:"scheme" validate:"required"`

	// The client user agent that initiated the request.
	Ua *string `json:"ua" validate:"required"`

	// The URI requested from the hostname.
	URI *string `json:"uri" validate:"required"`

	// The time that the event occurred.
	OccurredAt *strfmt.DateTime `json:"occurred_at" validate:"required"`

	// The firewall rules those the event matches.
	Matches []SecurityEventObjectMatchesItem `json:"matches" validate:"required"`
}

// Constants associated with the SecurityEventObject.Kind property.
// Kind of events. Now it is only firewall.
const (
	SecurityEventObject_Kind_Firewall = "firewall"
)

// Constants associated with the SecurityEventObject.Source property.
// Source of the event.
const (
	SecurityEventObject_Source_Asn           = "asn"
	SecurityEventObject_Source_Bic           = "bic"
	SecurityEventObject_Source_Country       = "country"
	SecurityEventObject_Source_Firewallrules = "firewallRules"
	SecurityEventObject_Source_Hot           = "hot"
	SecurityEventObject_Source_Ip            = "ip"
	SecurityEventObject_Source_Iprange       = "ipRange"
	SecurityEventObject_Source_L7ddos        = "l7ddos"
	SecurityEventObject_Source_Ratelimit     = "rateLimit"
	SecurityEventObject_Source_Securitylevel = "securityLevel"
	SecurityEventObject_Source_Uablock       = "uaBlock"
	SecurityEventObject_Source_Unknown       = "unknown"
	SecurityEventObject_Source_Waf           = "waf"
	SecurityEventObject_Source_Zonelockdown  = "zoneLockdown"
)

// Constants associated with the SecurityEventObject.Action property.
// What type of action was taken.
const (
	SecurityEventObject_Action_Allow           = "allow"
	SecurityEventObject_Action_Challenge       = "challenge"
	SecurityEventObject_Action_Connectionclose = "connectionClose"
	SecurityEventObject_Action_Drop            = "drop"
	SecurityEventObject_Action_Jschallenge     = "jschallenge"
	SecurityEventObject_Action_Log             = "log"
	SecurityEventObject_Action_Simulate        = "simulate"
	SecurityEventObject_Action_Unknown         = "unknown"
)

// Constants associated with the SecurityEventObject.IpClass property.
// IP class is a map of client IP to visitor classification.
const (
	SecurityEventObject_IpClass_Backupservice     = "backupService"
	SecurityEventObject_IpClass_Badhost           = "badHost"
	SecurityEventObject_IpClass_Clean             = "clean"
	SecurityEventObject_IpClass_Greylist          = "greylist"
	SecurityEventObject_IpClass_Mobileplatform    = "mobilePlatform"
	SecurityEventObject_IpClass_Monitoringservice = "monitoringService"
	SecurityEventObject_IpClass_Norecord          = "noRecord"
	SecurityEventObject_IpClass_Scan              = "scan"
	SecurityEventObject_IpClass_Searchengine      = "searchEngine"
	SecurityEventObject_IpClass_Securityscanner   = "securityScanner"
	SecurityEventObject_IpClass_Tor               = "tor"
	SecurityEventObject_IpClass_Unknown           = "unknown"
	SecurityEventObject_IpClass_Whitelist         = "whitelist"
)

// Constants associated with the SecurityEventObject.Method property.
// The HTTP method of the request.
const (
	SecurityEventObject_Method_Acl             = "ACL"
	SecurityEventObject_Method_BaselineControl = "BASELINE-CONTROL"
	SecurityEventObject_Method_Bcopy           = "BCOPY"
	SecurityEventObject_Method_Bdelete         = "BDELETE"
	SecurityEventObject_Method_Bmove           = "BMOVE"
	SecurityEventObject_Method_Bpropfind       = "BPROPFIND"
	SecurityEventObject_Method_Bproppatch      = "BPROPPATCH"
	SecurityEventObject_Method_Checkin         = "CHECKIN"
	SecurityEventObject_Method_Checkout        = "CHECKOUT"
	SecurityEventObject_Method_Connect         = "CONNECT"
	SecurityEventObject_Method_Cook            = "COOK"
	SecurityEventObject_Method_Copy            = "COPY"
	SecurityEventObject_Method_Delete          = "DELETE"
	SecurityEventObject_Method_Get             = "GET"
	SecurityEventObject_Method_Head            = "HEAD"
	SecurityEventObject_Method_JSON            = "JSON"
	SecurityEventObject_Method_Label           = "LABEL"
	SecurityEventObject_Method_Lock            = "LOCK"
	SecurityEventObject_Method_Merge           = "MERGE"
	SecurityEventObject_Method_Mkactivity      = "MKACTIVITY"
	SecurityEventObject_Method_Mkcol           = "MKCOL"
	SecurityEventObject_Method_Mkworkspace     = "MKWORKSPACE"
	SecurityEventObject_Method_Move            = "MOVE"
	SecurityEventObject_Method_Notify          = "NOTIFY"
	SecurityEventObject_Method_Options         = "OPTIONS"
	SecurityEventObject_Method_Orderpatch      = "ORDERPATCH"
	SecurityEventObject_Method_Patch           = "PATCH"
	SecurityEventObject_Method_Poll            = "POLL"
	SecurityEventObject_Method_Post            = "POST"
	SecurityEventObject_Method_Propfind        = "PROPFIND"
	SecurityEventObject_Method_Proppatch       = "PROPPATCH"
	SecurityEventObject_Method_Purge           = "PURGE"
	SecurityEventObject_Method_Put             = "PUT"
	SecurityEventObject_Method_Report          = "REPORT"
	SecurityEventObject_Method_RpcInData       = "RPC_IN_DATA"
	SecurityEventObject_Method_RpcOutData      = "RPC_OUT_DATA"
	SecurityEventObject_Method_Search          = "SEARCH"
	SecurityEventObject_Method_Subscribe       = "SUBSCRIBE"
	SecurityEventObject_Method_Trace           = "TRACE"
	SecurityEventObject_Method_Track           = "TRACK"
	SecurityEventObject_Method_Uncheckout      = "UNCHECKOUT"
	SecurityEventObject_Method_Unlock          = "UNLOCK"
	SecurityEventObject_Method_Unsubscribe     = "UNSUBSCRIBE"
	SecurityEventObject_Method_Update          = "UPDATE"
	SecurityEventObject_Method_VersionControl  = "VERSION-CONTROL"
	SecurityEventObject_Method_XMsEnumatts     = "X-MS-ENUMATTS"
)

// Constants associated with the SecurityEventObject.Proto property.
// The protocol of the request.
const (
	SecurityEventObject_Proto_Http10 = "HTTP/1.0"
	SecurityEventObject_Proto_Http11 = "HTTP/1.1"
	SecurityEventObject_Proto_Http12 = "HTTP/1.2"
	SecurityEventObject_Proto_Http2  = "HTTP/2"
	SecurityEventObject_Proto_Spdy31 = "SPDY/3.1"
	SecurityEventObject_Proto_Unk    = "UNK"
)

// Constants associated with the SecurityEventObject.Scheme property.
// The scheme of the uri.
const (
	SecurityEventObject_Scheme_Http    = "http"
	SecurityEventObject_Scheme_Https   = "https"
	SecurityEventObject_Scheme_Unknown = "unknown"
)

// UnmarshalSecurityEventObject unmarshals an instance of SecurityEventObject from the specified map of raw messages.
func UnmarshalSecurityEventObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityEventObject)
	err = core.UnmarshalPrimitive(m, "ray_id", &obj.RayID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rule_id", &obj.RuleID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ip_class", &obj.IpClass)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "country", &obj.Country)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "colo", &obj.Colo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "proto", &obj.Proto)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ua", &obj.Ua)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uri", &obj.URI)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "occurred_at", &obj.OccurredAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "matches", &obj.Matches, UnmarshalSecurityEventObjectMatchesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecurityEvents : security events objects.
type SecurityEvents struct {
	// Container for response information.
	Result []SecurityEventObject `json:"result" validate:"required"`

	// Statistics of results.
	ResultInfo *ResultInfo `json:"result_info" validate:"required"`

	// Was the get successful.
	Success *bool `json:"success" validate:"required"`

	// Array of errors encountered.
	Errors [][]string `json:"errors" validate:"required"`

	// Array of messages returned.
	Messages [][]string `json:"messages" validate:"required"`
}

// UnmarshalSecurityEvents unmarshals an instance of SecurityEvents from the specified map of raw messages.
func UnmarshalSecurityEvents(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecurityEvents)
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalSecurityEventObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result_info", &obj.ResultInfo, UnmarshalResultInfo)
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
