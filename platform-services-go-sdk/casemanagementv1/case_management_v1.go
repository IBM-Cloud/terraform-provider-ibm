/**
 * (C) Copyright IBM Corp. 2020, 2022.
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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.60.0-13f6e1ba-20221019-164457
 */

// Package casemanagementv1 : Operations and models for the CaseManagementV1 service
package casemanagementv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
)

// CaseManagementV1 : Case management API for creating cases, getting case statuses, adding comments to a case, adding
// and removing users from a case watchlist, downloading and adding attachments, and more.
//
// API Version: 1.0.0
type CaseManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://support-center.cloud.ibm.com/case-management/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "case_management"

// CaseManagementV1Options : Service options
type CaseManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewCaseManagementV1UsingExternalConfig : constructs an instance of CaseManagementV1 with passed in options and external configuration.
func NewCaseManagementV1UsingExternalConfig(options *CaseManagementV1Options) (caseManagement *CaseManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	caseManagement, err = NewCaseManagementV1(options)
	if err != nil {
		return
	}

	err = caseManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = caseManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewCaseManagementV1 : constructs an instance of CaseManagementV1 with passed in options.
func NewCaseManagementV1(options *CaseManagementV1Options) (service *CaseManagementV1, err error) {
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

	service = &CaseManagementV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "caseManagement" suitable for processing requests.
func (caseManagement *CaseManagementV1) Clone() *CaseManagementV1 {
	if core.IsNil(caseManagement) {
		return nil
	}
	clone := *caseManagement
	clone.Service = caseManagement.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (caseManagement *CaseManagementV1) SetServiceURL(url string) error {
	return caseManagement.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (caseManagement *CaseManagementV1) GetServiceURL() string {
	return caseManagement.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (caseManagement *CaseManagementV1) SetDefaultHeaders(headers http.Header) {
	caseManagement.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (caseManagement *CaseManagementV1) SetEnableGzipCompression(enableGzip bool) {
	caseManagement.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (caseManagement *CaseManagementV1) GetEnableGzipCompression() bool {
	return caseManagement.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (caseManagement *CaseManagementV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	caseManagement.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (caseManagement *CaseManagementV1) DisableRetries() {
	caseManagement.Service.DisableRetries()
}

// GetCases : Get cases in account
// Get cases in the account that are specified by the content of the IAM token.
func (caseManagement *CaseManagementV1) GetCases(getCasesOptions *GetCasesOptions) (result *CaseList, response *core.DetailedResponse, err error) {
	return caseManagement.GetCasesWithContext(context.Background(), getCasesOptions)
}

// GetCasesWithContext is an alternate form of the GetCases method which supports a Context parameter
func (caseManagement *CaseManagementV1) GetCasesWithContext(ctx context.Context, getCasesOptions *GetCasesOptions) (result *CaseList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCasesOptions, "getCasesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCasesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "GetCases")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getCasesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getCasesOptions.Offset))
	}
	if getCasesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getCasesOptions.Limit))
	}
	if getCasesOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*getCasesOptions.Search))
	}
	if getCasesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*getCasesOptions.Sort))
	}
	if getCasesOptions.Status != nil {
		builder.AddQuery("status", strings.Join(getCasesOptions.Status, ","))
	}
	if getCasesOptions.Fields != nil {
		builder.AddQuery("fields", strings.Join(getCasesOptions.Fields, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCaseList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCase : Create a case
// Create a support case to resolve issues in your account.
func (caseManagement *CaseManagementV1) CreateCase(createCaseOptions *CreateCaseOptions) (result *Case, response *core.DetailedResponse, err error) {
	return caseManagement.CreateCaseWithContext(context.Background(), createCaseOptions)
}

// CreateCaseWithContext is an alternate form of the CreateCase method which supports a Context parameter
func (caseManagement *CaseManagementV1) CreateCaseWithContext(ctx context.Context, createCaseOptions *CreateCaseOptions) (result *Case, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCaseOptions, "createCaseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCaseOptions, "createCaseOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "CreateCase")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCaseOptions.Type != nil {
		body["type"] = createCaseOptions.Type
	}
	if createCaseOptions.Subject != nil {
		body["subject"] = createCaseOptions.Subject
	}
	if createCaseOptions.Description != nil {
		body["description"] = createCaseOptions.Description
	}
	if createCaseOptions.Severity != nil {
		body["severity"] = createCaseOptions.Severity
	}
	if createCaseOptions.Eu != nil {
		body["eu"] = createCaseOptions.Eu
	}
	if createCaseOptions.Offering != nil {
		body["offering"] = createCaseOptions.Offering
	}
	if createCaseOptions.Resources != nil {
		body["resources"] = createCaseOptions.Resources
	}
	if createCaseOptions.Watchlist != nil {
		body["watchlist"] = createCaseOptions.Watchlist
	}
	if createCaseOptions.InvoiceNumber != nil {
		body["invoice_number"] = createCaseOptions.InvoiceNumber
	}
	if createCaseOptions.SLACreditRequest != nil {
		body["sla_credit_request"] = createCaseOptions.SLACreditRequest
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
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCase)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCase : Get a case in account
// View a case in the account that is specified by the case number.
func (caseManagement *CaseManagementV1) GetCase(getCaseOptions *GetCaseOptions) (result *Case, response *core.DetailedResponse, err error) {
	return caseManagement.GetCaseWithContext(context.Background(), getCaseOptions)
}

// GetCaseWithContext is an alternate form of the GetCase method which supports a Context parameter
func (caseManagement *CaseManagementV1) GetCaseWithContext(ctx context.Context, getCaseOptions *GetCaseOptions) (result *Case, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCaseOptions, "getCaseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCaseOptions, "getCaseOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"case_number": *getCaseOptions.CaseNumber,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases/{case_number}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "GetCase")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getCaseOptions.Fields != nil {
		builder.AddQuery("fields", strings.Join(getCaseOptions.Fields, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCase)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateCaseStatus : Update case status
// Mark the case as resolved or unresolved, or accept the provided resolution.
func (caseManagement *CaseManagementV1) UpdateCaseStatus(updateCaseStatusOptions *UpdateCaseStatusOptions) (result *Case, response *core.DetailedResponse, err error) {
	return caseManagement.UpdateCaseStatusWithContext(context.Background(), updateCaseStatusOptions)
}

// UpdateCaseStatusWithContext is an alternate form of the UpdateCaseStatus method which supports a Context parameter
func (caseManagement *CaseManagementV1) UpdateCaseStatusWithContext(ctx context.Context, updateCaseStatusOptions *UpdateCaseStatusOptions) (result *Case, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCaseStatusOptions, "updateCaseStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCaseStatusOptions, "updateCaseStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"case_number": *updateCaseStatusOptions.CaseNumber,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases/{case_number}/status`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCaseStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "UpdateCaseStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(updateCaseStatusOptions.StatusPayload)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCase)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddComment : Add comment to case
// Add a comment to a case to be viewed by a support engineer.
func (caseManagement *CaseManagementV1) AddComment(addCommentOptions *AddCommentOptions) (result *Comment, response *core.DetailedResponse, err error) {
	return caseManagement.AddCommentWithContext(context.Background(), addCommentOptions)
}

// AddCommentWithContext is an alternate form of the AddComment method which supports a Context parameter
func (caseManagement *CaseManagementV1) AddCommentWithContext(ctx context.Context, addCommentOptions *AddCommentOptions) (result *Comment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addCommentOptions, "addCommentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addCommentOptions, "addCommentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"case_number": *addCommentOptions.CaseNumber,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases/{case_number}/comments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addCommentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "AddComment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addCommentOptions.Comment != nil {
		body["comment"] = addCommentOptions.Comment
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
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalComment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddWatchlist : Add users to watchlist of case
// Add users to the watchlist of case. By adding a user to the watchlist of the case, you are granting them read and
// write permissions, so the user can view the case, receive updates, and make updates to the case. Note that the user
// must be in the account to be added to the watchlist.
func (caseManagement *CaseManagementV1) AddWatchlist(addWatchlistOptions *AddWatchlistOptions) (result *WatchlistAddResponse, response *core.DetailedResponse, err error) {
	return caseManagement.AddWatchlistWithContext(context.Background(), addWatchlistOptions)
}

// AddWatchlistWithContext is an alternate form of the AddWatchlist method which supports a Context parameter
func (caseManagement *CaseManagementV1) AddWatchlistWithContext(ctx context.Context, addWatchlistOptions *AddWatchlistOptions) (result *WatchlistAddResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addWatchlistOptions, "addWatchlistOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addWatchlistOptions, "addWatchlistOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"case_number": *addWatchlistOptions.CaseNumber,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases/{case_number}/watchlist`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addWatchlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "AddWatchlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addWatchlistOptions.Watchlist != nil {
		body["watchlist"] = addWatchlistOptions.Watchlist
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
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWatchlistAddResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoveWatchlist : Remove users from watchlist of case
// Remove users from the watchlist of a case if you don't want them to view the case, receive updates, or make updates
// to the case.
func (caseManagement *CaseManagementV1) RemoveWatchlist(removeWatchlistOptions *RemoveWatchlistOptions) (result *Watchlist, response *core.DetailedResponse, err error) {
	return caseManagement.RemoveWatchlistWithContext(context.Background(), removeWatchlistOptions)
}

// RemoveWatchlistWithContext is an alternate form of the RemoveWatchlist method which supports a Context parameter
func (caseManagement *CaseManagementV1) RemoveWatchlistWithContext(ctx context.Context, removeWatchlistOptions *RemoveWatchlistOptions) (result *Watchlist, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeWatchlistOptions, "removeWatchlistOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeWatchlistOptions, "removeWatchlistOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"case_number": *removeWatchlistOptions.CaseNumber,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases/{case_number}/watchlist`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeWatchlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "RemoveWatchlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if removeWatchlistOptions.Watchlist != nil {
		body["watchlist"] = removeWatchlistOptions.Watchlist
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
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWatchlist)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddResource : Add a resource to case
// Add a resource to case by specifying the Cloud Resource Name (CRN), or id and type if attaching a class iaaS
// resource.
func (caseManagement *CaseManagementV1) AddResource(addResourceOptions *AddResourceOptions) (result *Resource, response *core.DetailedResponse, err error) {
	return caseManagement.AddResourceWithContext(context.Background(), addResourceOptions)
}

// AddResourceWithContext is an alternate form of the AddResource method which supports a Context parameter
func (caseManagement *CaseManagementV1) AddResourceWithContext(ctx context.Context, addResourceOptions *AddResourceOptions) (result *Resource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addResourceOptions, "addResourceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addResourceOptions, "addResourceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"case_number": *addResourceOptions.CaseNumber,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases/{case_number}/resources`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addResourceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "AddResource")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addResourceOptions.CRN != nil {
		body["crn"] = addResourceOptions.CRN
	}
	if addResourceOptions.Type != nil {
		body["type"] = addResourceOptions.Type
	}
	if addResourceOptions.ID != nil {
		body["id"] = addResourceOptions.ID
	}
	if addResourceOptions.Note != nil {
		body["note"] = addResourceOptions.Note
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
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UploadFile : Add attachments to a support case
// You can add attachments to a case to provide more information for the support team about the issue that you're
// experiencing.
func (caseManagement *CaseManagementV1) UploadFile(uploadFileOptions *UploadFileOptions) (result *Attachment, response *core.DetailedResponse, err error) {
	return caseManagement.UploadFileWithContext(context.Background(), uploadFileOptions)
}

// UploadFileWithContext is an alternate form of the UploadFile method which supports a Context parameter
func (caseManagement *CaseManagementV1) UploadFileWithContext(ctx context.Context, uploadFileOptions *UploadFileOptions) (result *Attachment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(uploadFileOptions, "uploadFileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(uploadFileOptions, "uploadFileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"case_number": *uploadFileOptions.CaseNumber,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases/{case_number}/attachments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range uploadFileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "UploadFile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	for _, item := range uploadFileOptions.File {
		builder.AddFormData("file", core.StringNilMapper(item.Filename), core.StringNilMapper(item.ContentType), item.Data)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachment)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DownloadFile : Download an attachment
// Download an attachment from a case.
func (caseManagement *CaseManagementV1) DownloadFile(downloadFileOptions *DownloadFileOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return caseManagement.DownloadFileWithContext(context.Background(), downloadFileOptions)
}

// DownloadFileWithContext is an alternate form of the DownloadFile method which supports a Context parameter
func (caseManagement *CaseManagementV1) DownloadFileWithContext(ctx context.Context, downloadFileOptions *DownloadFileOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(downloadFileOptions, "downloadFileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(downloadFileOptions, "downloadFileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"case_number": *downloadFileOptions.CaseNumber,
		"file_id":     *downloadFileOptions.FileID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases/{case_number}/attachments/{file_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range downloadFileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "DownloadFile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/octet-stream")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = caseManagement.Service.Request(request, &result)

	return
}

// DeleteFile : Remove attachment from case
// Remove an attachment from a case.
func (caseManagement *CaseManagementV1) DeleteFile(deleteFileOptions *DeleteFileOptions) (result *AttachmentList, response *core.DetailedResponse, err error) {
	return caseManagement.DeleteFileWithContext(context.Background(), deleteFileOptions)
}

// DeleteFileWithContext is an alternate form of the DeleteFile method which supports a Context parameter
func (caseManagement *CaseManagementV1) DeleteFileWithContext(ctx context.Context, deleteFileOptions *DeleteFileOptions) (result *AttachmentList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteFileOptions, "deleteFileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteFileOptions, "deleteFileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"case_number": *deleteFileOptions.CaseNumber,
		"file_id":     *deleteFileOptions.FileID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = caseManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(caseManagement.Service.Options.URL, `/cases/{case_number}/attachments/{file_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteFileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("case_management", "V1", "DeleteFile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = caseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddCommentOptions : The AddComment options.
type AddCommentOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required,ne="`

	// Comment to add to the case.
	Comment *string `json:"comment" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddCommentOptions : Instantiate AddCommentOptions
func (*CaseManagementV1) NewAddCommentOptions(caseNumber string, comment string) *AddCommentOptions {
	return &AddCommentOptions{
		CaseNumber: core.StringPtr(caseNumber),
		Comment:    core.StringPtr(comment),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (_options *AddCommentOptions) SetCaseNumber(caseNumber string) *AddCommentOptions {
	_options.CaseNumber = core.StringPtr(caseNumber)
	return _options
}

// SetComment : Allow user to set Comment
func (_options *AddCommentOptions) SetComment(comment string) *AddCommentOptions {
	_options.Comment = core.StringPtr(comment)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddCommentOptions) SetHeaders(param map[string]string) *AddCommentOptions {
	options.Headers = param
	return options
}

// AddResourceOptions : The AddResource options.
type AddResourceOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required,ne="`

	// Cloud Resource Name of the resource.
	CRN *string `json:"crn,omitempty"`

	// Only used to attach Classic IaaS devices that have no CRN.
	Type *string `json:"type,omitempty"`

	// Only used to attach Classic IaaS devices that have no CRN. Id of Classic IaaS device. This is deprecated in favor of
	// the crn field.
	// Deprecated: this field is deprecated and may be removed in a future release.
	ID *float64 `json:"id,omitempty"`

	// A note about this resource.
	Note *string `json:"note,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddResourceOptions : Instantiate AddResourceOptions
func (*CaseManagementV1) NewAddResourceOptions(caseNumber string) *AddResourceOptions {
	return &AddResourceOptions{
		CaseNumber: core.StringPtr(caseNumber),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (_options *AddResourceOptions) SetCaseNumber(caseNumber string) *AddResourceOptions {
	_options.CaseNumber = core.StringPtr(caseNumber)
	return _options
}

// SetCRN : Allow user to set CRN
func (_options *AddResourceOptions) SetCRN(crn string) *AddResourceOptions {
	_options.CRN = core.StringPtr(crn)
	return _options
}

// SetType : Allow user to set Type
func (_options *AddResourceOptions) SetType(typeVar string) *AddResourceOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetID : Allow user to set ID
// Deprecated: this method is deprecated and may be removed in a future release.
func (_options *AddResourceOptions) SetID(id float64) *AddResourceOptions {
	_options.ID = core.Float64Ptr(id)
	return _options
}

// SetNote : Allow user to set Note
func (_options *AddResourceOptions) SetNote(note string) *AddResourceOptions {
	_options.Note = core.StringPtr(note)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddResourceOptions) SetHeaders(param map[string]string) *AddResourceOptions {
	options.Headers = param
	return options
}

// AddWatchlistOptions : The AddWatchlist options.
type AddWatchlistOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required,ne="`

	// Array of user ID objects.
	Watchlist []User `json:"watchlist,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddWatchlistOptions : Instantiate AddWatchlistOptions
func (*CaseManagementV1) NewAddWatchlistOptions(caseNumber string) *AddWatchlistOptions {
	return &AddWatchlistOptions{
		CaseNumber: core.StringPtr(caseNumber),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (_options *AddWatchlistOptions) SetCaseNumber(caseNumber string) *AddWatchlistOptions {
	_options.CaseNumber = core.StringPtr(caseNumber)
	return _options
}

// SetWatchlist : Allow user to set Watchlist
func (_options *AddWatchlistOptions) SetWatchlist(watchlist []User) *AddWatchlistOptions {
	_options.Watchlist = watchlist
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddWatchlistOptions) SetHeaders(param map[string]string) *AddWatchlistOptions {
	options.Headers = param
	return options
}

// Attachment : Details of an attachment.
type Attachment struct {
	// Unique identifier of the attachment in database.
	ID *string `json:"id,omitempty"`

	// Name of the attachment.
	Filename *string `json:"filename,omitempty"`

	// Size of the attachment in bytes.
	SizeInBytes *int64 `json:"size_in_bytes,omitempty"`

	// Date time of uploading in UTC.
	CreatedAt *string `json:"created_at,omitempty"`

	// URL of the attachment used to download.
	URL *string `json:"url,omitempty"`
}

// UnmarshalAttachment unmarshals an instance of Attachment from the specified map of raw messages.
func UnmarshalAttachment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Attachment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "filename", &obj.Filename)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size_in_bytes", &obj.SizeInBytes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentList : List of attachments in the case.
type AttachmentList struct {
	// New attachments array.
	Attachments []Attachment `json:"attachments,omitempty"`
}

// UnmarshalAttachmentList unmarshals an instance of AttachmentList from the specified map of raw messages.
func UnmarshalAttachmentList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentList)
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Case : The support case.
type Case struct {
	// Identifying number of a created case.
	Number *string `json:"number,omitempty"`

	// Short description of what the case is about.
	ShortDescription *string `json:"short_description,omitempty"`

	// Full description of what the case is about.
	Description *string `json:"description,omitempty"`

	// Date and time of case creation in UTC.
	CreatedAt *string `json:"created_at,omitempty"`

	// User info in a case.
	CreatedBy *User `json:"created_by,omitempty"`

	// Date and time of the last update on the case in UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// User info in a case.
	UpdatedBy *User `json:"updated_by,omitempty"`

	// Name of the console to interact with the contact.
	ContactType *string `json:"contact_type,omitempty"`

	// User info in a case.
	Contact *User `json:"contact,omitempty"`

	// Status type of the case.
	Status *string `json:"status,omitempty"`

	// Severity level of the case.
	Severity *float64 `json:"severity,omitempty"`

	// Support tier of the account.
	SupportTier *string `json:"support_tier,omitempty"`

	// Standard reasons of resolving case.
	Resolution *string `json:"resolution,omitempty"`

	// Notes of case closing.
	CloseNotes *string `json:"close_notes,omitempty"`

	// EU support.
	Eu *CaseEu `json:"eu,omitempty"`

	// List of users in the case watchlist.
	Watchlist []User `json:"watchlist,omitempty"`

	// List of files that are attached to the case.
	Attachments []Attachment `json:"attachments,omitempty"`

	// Offering details.
	Offering *Offering `json:"offering,omitempty"`

	// List of attached resources.
	Resources []Resource `json:"resources,omitempty"`

	// List of comments and updates that are sorted in chronological order.
	Comments []Comment `json:"comments,omitempty"`
}

// Constants associated with the Case.ContactType property.
// Name of the console to interact with the contact.
const (
	CaseContactTypeCloudSupportCenterConst = "Cloud Support Center"
	CaseContactTypeImsConsoleConst         = "IMS Console"
)

// Constants associated with the Case.SupportTier property.
// Support tier of the account.
const (
	CaseSupportTierBasicConst    = "Basic"
	CaseSupportTierFreeConst     = "Free"
	CaseSupportTierPremiumConst  = "Premium"
	CaseSupportTierStandardConst = "Standard"
)

// UnmarshalCase unmarshals an instance of Case from the specified map of raw messages.
func UnmarshalCase(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Case)
	err = core.UnmarshalPrimitive(m, "number", &obj.Number)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "short_description", &obj.ShortDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "created_by", &obj.CreatedBy, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "updated_by", &obj.UpdatedBy, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "contact_type", &obj.ContactType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contact", &obj.Contact, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "severity", &obj.Severity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "support_tier", &obj.SupportTier)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resolution", &obj.Resolution)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "close_notes", &obj.CloseNotes)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "eu", &obj.Eu, UnmarshalCaseEu)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "watchlist", &obj.Watchlist, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "offering", &obj.Offering, UnmarshalOffering)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "comments", &obj.Comments, UnmarshalComment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CaseEu : EU support.
type CaseEu struct {
	// Identifying whether the case has EU Support.
	Support *bool `json:"support,omitempty"`

	// Information about the data center.
	DataCenter *string `json:"data_center,omitempty"`
}

// UnmarshalCaseEu unmarshals an instance of CaseEu from the specified map of raw messages.
func UnmarshalCaseEu(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CaseEu)
	err = core.UnmarshalPrimitive(m, "support", &obj.Support)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_center", &obj.DataCenter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CaseList : Response of a GET /cases request.
type CaseList struct {
	// Total number of cases that satisfy the query.
	TotalCount *int64 `json:"total_count,omitempty"`

	// Container for URL pointer to related pages of cases.
	First *PaginationLink `json:"first,omitempty"`

	// Container for URL pointer to related pages of cases.
	Next *PaginationLink `json:"next,omitempty"`

	// Container for URL pointer to related pages of cases.
	Previous *PaginationLink `json:"previous,omitempty"`

	// Container for URL pointer to related pages of cases.
	Last *PaginationLink `json:"last,omitempty"`

	// List of cases.
	Cases []Case `json:"cases,omitempty"`
}

// UnmarshalCaseList unmarshals an instance of CaseList from the specified map of raw messages.
func UnmarshalCaseList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CaseList)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cases", &obj.Cases, UnmarshalCase)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *CaseList) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// CasePayloadEu : Specify if the case should be treated as EU regulated. Only one of the following properties is required. Call EU
// support utility endpoint to determine which property must be specified for your account.
type CasePayloadEu struct {
	// indicating whether the case is EU supported.
	Supported *bool `json:"supported,omitempty"`

	// If EU supported utility endpoint specifies data center, then pass the data center id to mark a case as EU supported.
	DataCenter *int64 `json:"data_center,omitempty"`
}

// UnmarshalCasePayloadEu unmarshals an instance of CasePayloadEu from the specified map of raw messages.
func UnmarshalCasePayloadEu(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CasePayloadEu)
	err = core.UnmarshalPrimitive(m, "supported", &obj.Supported)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_center", &obj.DataCenter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Comment : A comment in a case.
type Comment struct {
	// The comment.
	Value *string `json:"value,omitempty"`

	// Date time when comment was added in UTC.
	AddedAt *string `json:"added_at,omitempty"`

	// User info in a case.
	AddedBy *User `json:"added_by,omitempty"`
}

// UnmarshalComment unmarshals an instance of Comment from the specified map of raw messages.
func UnmarshalComment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Comment)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "added_at", &obj.AddedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "added_by", &obj.AddedBy, UnmarshalUser)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCaseOptions : The CreateCase options.
type CreateCaseOptions struct {
	// Case type.
	Type *string `json:"type" validate:"required"`

	// Short description used to identify the case.
	Subject *string `json:"subject" validate:"required"`

	// Detailed description of the issue.
	Description *string `json:"description" validate:"required"`

	// Severity of the case. Smaller values mean higher severity.
	Severity *int64 `json:"severity,omitempty"`

	// Specify if the case should be treated as EU regulated. Only one of the following properties is required. Call EU
	// support utility endpoint to determine which property must be specified for your account.
	Eu *CasePayloadEu `json:"eu,omitempty"`

	// Offering details.
	Offering *Offering `json:"offering,omitempty"`

	// List of resources to attach to case. If you attach Classic IaaS devices, use the type and id fields if the Cloud
	// Resource Name (CRN) is unavailable. Otherwise, pass the resource CRN. The resource list must be consistent with the
	// value that is selected for the resource offering.
	Resources []ResourcePayload `json:"resources,omitempty"`

	// Array of user IDs to add to the watchlist.
	Watchlist []User `json:"watchlist,omitempty"`

	// Invoice number of "Billing and Invoice" case type.
	InvoiceNumber *string `json:"invoice_number,omitempty"`

	// Flag to indicate if case is for an Service Level Agreement (SLA) credit request.
	SLACreditRequest *bool `json:"sla_credit_request,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCaseOptions.Type property.
// Case type.
const (
	CreateCaseOptionsTypeAccountAndAccessConst  = "account_and_access"
	CreateCaseOptionsTypeBillingAndInvoiceConst = "billing_and_invoice"
	CreateCaseOptionsTypeSalesConst             = "sales"
	CreateCaseOptionsTypeTechnicalConst         = "technical"
)

// NewCreateCaseOptions : Instantiate CreateCaseOptions
func (*CaseManagementV1) NewCreateCaseOptions(typeVar string, subject string, description string) *CreateCaseOptions {
	return &CreateCaseOptions{
		Type:        core.StringPtr(typeVar),
		Subject:     core.StringPtr(subject),
		Description: core.StringPtr(description),
	}
}

// SetType : Allow user to set Type
func (_options *CreateCaseOptions) SetType(typeVar string) *CreateCaseOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetSubject : Allow user to set Subject
func (_options *CreateCaseOptions) SetSubject(subject string) *CreateCaseOptions {
	_options.Subject = core.StringPtr(subject)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateCaseOptions) SetDescription(description string) *CreateCaseOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetSeverity : Allow user to set Severity
func (_options *CreateCaseOptions) SetSeverity(severity int64) *CreateCaseOptions {
	_options.Severity = core.Int64Ptr(severity)
	return _options
}

// SetEu : Allow user to set Eu
func (_options *CreateCaseOptions) SetEu(eu *CasePayloadEu) *CreateCaseOptions {
	_options.Eu = eu
	return _options
}

// SetOffering : Allow user to set Offering
func (_options *CreateCaseOptions) SetOffering(offering *Offering) *CreateCaseOptions {
	_options.Offering = offering
	return _options
}

// SetResources : Allow user to set Resources
func (_options *CreateCaseOptions) SetResources(resources []ResourcePayload) *CreateCaseOptions {
	_options.Resources = resources
	return _options
}

// SetWatchlist : Allow user to set Watchlist
func (_options *CreateCaseOptions) SetWatchlist(watchlist []User) *CreateCaseOptions {
	_options.Watchlist = watchlist
	return _options
}

// SetInvoiceNumber : Allow user to set InvoiceNumber
func (_options *CreateCaseOptions) SetInvoiceNumber(invoiceNumber string) *CreateCaseOptions {
	_options.InvoiceNumber = core.StringPtr(invoiceNumber)
	return _options
}

// SetSLACreditRequest : Allow user to set SLACreditRequest
func (_options *CreateCaseOptions) SetSLACreditRequest(slaCreditRequest bool) *CreateCaseOptions {
	_options.SLACreditRequest = core.BoolPtr(slaCreditRequest)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCaseOptions) SetHeaders(param map[string]string) *CreateCaseOptions {
	options.Headers = param
	return options
}

// DeleteFileOptions : The DeleteFile options.
type DeleteFileOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required,ne="`

	// Unique identifier of a file.
	FileID *string `json:"file_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteFileOptions : Instantiate DeleteFileOptions
func (*CaseManagementV1) NewDeleteFileOptions(caseNumber string, fileID string) *DeleteFileOptions {
	return &DeleteFileOptions{
		CaseNumber: core.StringPtr(caseNumber),
		FileID:     core.StringPtr(fileID),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (_options *DeleteFileOptions) SetCaseNumber(caseNumber string) *DeleteFileOptions {
	_options.CaseNumber = core.StringPtr(caseNumber)
	return _options
}

// SetFileID : Allow user to set FileID
func (_options *DeleteFileOptions) SetFileID(fileID string) *DeleteFileOptions {
	_options.FileID = core.StringPtr(fileID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteFileOptions) SetHeaders(param map[string]string) *DeleteFileOptions {
	options.Headers = param
	return options
}

// DownloadFileOptions : The DownloadFile options.
type DownloadFileOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required,ne="`

	// Unique identifier of a file.
	FileID *string `json:"file_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDownloadFileOptions : Instantiate DownloadFileOptions
func (*CaseManagementV1) NewDownloadFileOptions(caseNumber string, fileID string) *DownloadFileOptions {
	return &DownloadFileOptions{
		CaseNumber: core.StringPtr(caseNumber),
		FileID:     core.StringPtr(fileID),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (_options *DownloadFileOptions) SetCaseNumber(caseNumber string) *DownloadFileOptions {
	_options.CaseNumber = core.StringPtr(caseNumber)
	return _options
}

// SetFileID : Allow user to set FileID
func (_options *DownloadFileOptions) SetFileID(fileID string) *DownloadFileOptions {
	_options.FileID = core.StringPtr(fileID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DownloadFileOptions) SetHeaders(param map[string]string) *DownloadFileOptions {
	options.Headers = param
	return options
}

// FileWithMetadata : A file with its associated metadata.
type FileWithMetadata struct {
	// The data / content for the file.
	Data io.ReadCloser `json:"data" validate:"required"`

	// The filename of the file.
	Filename *string `json:"filename,omitempty"`

	// The content type of the file.
	ContentType *string `json:"content_type,omitempty"`
}

// NewFileWithMetadata : Instantiate FileWithMetadata (Generic Model Constructor)
func (*CaseManagementV1) NewFileWithMetadata(data io.ReadCloser) (_model *FileWithMetadata, err error) {
	_model = &FileWithMetadata{
		Data: data,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalFileWithMetadata unmarshals an instance of FileWithMetadata from the specified map of raw messages.
func UnmarshalFileWithMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(core.FileWithMetadata)
	err = core.UnmarshalFileWithMetadata(m, &obj)
	if err != nil {
		return
	}

	// do a simple conversion from the core type to the service type
	// they have identical fields
	convertedModel := FileWithMetadata(*obj)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(&convertedModel))

	return
}

// GetCaseOptions : The GetCase options.
type GetCaseOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required,ne="`

	// Selected fields of interest instead of all of the case information.
	Fields []string `json:"fields,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetCaseOptions.Fields property.
const (
	GetCaseOptionsFieldsAgentCloseOnlyConst   = "agent_close_only"
	GetCaseOptionsFieldsAttachmentsConst      = "attachments"
	GetCaseOptionsFieldsCloseNotesConst       = "close_notes"
	GetCaseOptionsFieldsCommentsConst         = "comments"
	GetCaseOptionsFieldsContactConst          = "contact"
	GetCaseOptionsFieldsContactTypeConst      = "contact_type"
	GetCaseOptionsFieldsCreatedAtConst        = "created_at"
	GetCaseOptionsFieldsCreatedByConst        = "created_by"
	GetCaseOptionsFieldsDescriptionConst      = "description"
	GetCaseOptionsFieldsEuConst               = "eu"
	GetCaseOptionsFieldsInvoiceNumberConst    = "invoice_number"
	GetCaseOptionsFieldsNumberConst           = "number"
	GetCaseOptionsFieldsOfferingConst         = "offering"
	GetCaseOptionsFieldsResolutionConst       = "resolution"
	GetCaseOptionsFieldsResourcesConst        = "resources"
	GetCaseOptionsFieldsSeverityConst         = "severity"
	GetCaseOptionsFieldsShortDescriptionConst = "short_description"
	GetCaseOptionsFieldsStatusConst           = "status"
	GetCaseOptionsFieldsSupportTierConst      = "support_tier"
	GetCaseOptionsFieldsUpdatedAtConst        = "updated_at"
	GetCaseOptionsFieldsUpdatedByConst        = "updated_by"
	GetCaseOptionsFieldsWatchlistConst        = "watchlist"
)

// NewGetCaseOptions : Instantiate GetCaseOptions
func (*CaseManagementV1) NewGetCaseOptions(caseNumber string) *GetCaseOptions {
	return &GetCaseOptions{
		CaseNumber: core.StringPtr(caseNumber),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (_options *GetCaseOptions) SetCaseNumber(caseNumber string) *GetCaseOptions {
	_options.CaseNumber = core.StringPtr(caseNumber)
	return _options
}

// SetFields : Allow user to set Fields
func (_options *GetCaseOptions) SetFields(fields []string) *GetCaseOptions {
	_options.Fields = fields
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCaseOptions) SetHeaders(param map[string]string) *GetCaseOptions {
	options.Headers = param
	return options
}

// GetCasesOptions : The GetCases options.
type GetCasesOptions struct {
	// Number of cases that are skipped.
	Offset *int64 `json:"offset,omitempty"`

	// Number of cases that are returned.
	Limit *int64 `json:"limit,omitempty"`

	// String that a case might contain.
	Search *string `json:"search,omitempty"`

	// Sort field and direction. If omitted, default to descending of updated date. Prefix "~" signifies sort in
	// descending.
	Sort *string `json:"sort,omitempty"`

	// Case status filter.
	Status []string `json:"status,omitempty"`

	// Selected fields of interest instead of all of the case information.
	Fields []string `json:"fields,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetCasesOptions.Status property.
const (
	GetCasesOptionsStatusClosedConst             = "closed"
	GetCasesOptionsStatusInProgressConst         = "in_progress"
	GetCasesOptionsStatusNewConst                = "new"
	GetCasesOptionsStatusResolutionProvidedConst = "resolution_provided"
	GetCasesOptionsStatusResolvedConst           = "resolved"
	GetCasesOptionsStatusWaitingOnClientConst    = "waiting_on_client"
)

// Constants associated with the GetCasesOptions.Fields property.
const (
	GetCasesOptionsFieldsAgentCloseOnlyConst   = "agent_close_only"
	GetCasesOptionsFieldsAttachmentsConst      = "attachments"
	GetCasesOptionsFieldsCloseNotesConst       = "close_notes"
	GetCasesOptionsFieldsCommentsConst         = "comments"
	GetCasesOptionsFieldsContactConst          = "contact"
	GetCasesOptionsFieldsContactTypeConst      = "contact_type"
	GetCasesOptionsFieldsCreatedAtConst        = "created_at"
	GetCasesOptionsFieldsCreatedByConst        = "created_by"
	GetCasesOptionsFieldsDescriptionConst      = "description"
	GetCasesOptionsFieldsEuConst               = "eu"
	GetCasesOptionsFieldsInvoiceNumberConst    = "invoice_number"
	GetCasesOptionsFieldsNumberConst           = "number"
	GetCasesOptionsFieldsOfferingConst         = "offering"
	GetCasesOptionsFieldsResolutionConst       = "resolution"
	GetCasesOptionsFieldsResourcesConst        = "resources"
	GetCasesOptionsFieldsSeverityConst         = "severity"
	GetCasesOptionsFieldsShortDescriptionConst = "short_description"
	GetCasesOptionsFieldsStatusConst           = "status"
	GetCasesOptionsFieldsSupportTierConst      = "support_tier"
	GetCasesOptionsFieldsUpdatedAtConst        = "updated_at"
	GetCasesOptionsFieldsUpdatedByConst        = "updated_by"
	GetCasesOptionsFieldsWatchlistConst        = "watchlist"
)

// NewGetCasesOptions : Instantiate GetCasesOptions
func (*CaseManagementV1) NewGetCasesOptions() *GetCasesOptions {
	return &GetCasesOptions{}
}

// SetOffset : Allow user to set Offset
func (_options *GetCasesOptions) SetOffset(offset int64) *GetCasesOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *GetCasesOptions) SetLimit(limit int64) *GetCasesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *GetCasesOptions) SetSearch(search string) *GetCasesOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *GetCasesOptions) SetSort(sort string) *GetCasesOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *GetCasesOptions) SetStatus(status []string) *GetCasesOptions {
	_options.Status = status
	return _options
}

// SetFields : Allow user to set Fields
func (_options *GetCasesOptions) SetFields(fields []string) *GetCasesOptions {
	_options.Fields = fields
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCasesOptions) SetHeaders(param map[string]string) *GetCasesOptions {
	options.Headers = param
	return options
}

// Offering : Offering details.
type Offering struct {
	// Name of the offering.
	Name *string `json:"name" validate:"required"`

	// Offering type.
	Type *OfferingType `json:"type" validate:"required"`
}

// NewOffering : Instantiate Offering (Generic Model Constructor)
func (*CaseManagementV1) NewOffering(name string, typeVar *OfferingType) (_model *Offering, err error) {
	_model = &Offering{
		Name: core.StringPtr(name),
		Type: typeVar,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalOffering unmarshals an instance of Offering from the specified map of raw messages.
func UnmarshalOffering(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Offering)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "type", &obj.Type, UnmarshalOfferingType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OfferingType : Offering type.
type OfferingType struct {
	// Offering type group. "crn_service_name" is preferred over "category" as the latter is legacy and will be deprecated
	// in the future.
	Group *string `json:"group" validate:"required"`

	// CRN service name of the offering.
	Key *string `json:"key" validate:"required"`

	// Optional. Platform kind of the offering.
	Kind *string `json:"kind,omitempty"`

	// Offering id in the catalog. This alone is enough to identify the offering.
	ID *string `json:"id,omitempty"`
}

// Constants associated with the OfferingType.Group property.
// Offering type group. "crn_service_name" is preferred over "category" as the latter is legacy and will be deprecated
// in the future.
const (
	OfferingTypeGroupCRNServiceNameConst = "crn_service_name"
	OfferingTypeGroupCategoryConst       = "category"
)

// NewOfferingType : Instantiate OfferingType (Generic Model Constructor)
func (*CaseManagementV1) NewOfferingType(group string, key string) (_model *OfferingType, err error) {
	_model = &OfferingType{
		Group: core.StringPtr(group),
		Key:   core.StringPtr(key),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalOfferingType unmarshals an instance of OfferingType from the specified map of raw messages.
func UnmarshalOfferingType(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OfferingType)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginationLink : Container for URL pointer to related pages of cases.
type PaginationLink struct {
	// URL to related pages of cases.
	Href *string `json:"href,omitempty"`
}

// UnmarshalPaginationLink unmarshals an instance of PaginationLink from the specified map of raw messages.
func UnmarshalPaginationLink(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationLink)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RemoveWatchlistOptions : The RemoveWatchlist options.
type RemoveWatchlistOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required,ne="`

	// Array of user ID objects.
	Watchlist []User `json:"watchlist,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveWatchlistOptions : Instantiate RemoveWatchlistOptions
func (*CaseManagementV1) NewRemoveWatchlistOptions(caseNumber string) *RemoveWatchlistOptions {
	return &RemoveWatchlistOptions{
		CaseNumber: core.StringPtr(caseNumber),
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (_options *RemoveWatchlistOptions) SetCaseNumber(caseNumber string) *RemoveWatchlistOptions {
	_options.CaseNumber = core.StringPtr(caseNumber)
	return _options
}

// SetWatchlist : Allow user to set Watchlist
func (_options *RemoveWatchlistOptions) SetWatchlist(watchlist []User) *RemoveWatchlistOptions {
	_options.Watchlist = watchlist
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveWatchlistOptions) SetHeaders(param map[string]string) *RemoveWatchlistOptions {
	options.Headers = param
	return options
}

// Resource : A resource record of a case.
type Resource struct {
	// ID of the resource.
	CRN *string `json:"crn,omitempty"`

	// Name of the resource.
	Name *string `json:"name,omitempty"`

	// Type of resource.
	Type *string `json:"type,omitempty"`

	// URL of resource.
	URL *string `json:"url,omitempty"`

	// Note about resource.
	Note *string `json:"note,omitempty"`
}

// UnmarshalResource unmarshals an instance of Resource from the specified map of raw messages.
func UnmarshalResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resource)
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "note", &obj.Note)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourcePayload : Payload to add a resource to a case.
type ResourcePayload struct {
	// Cloud Resource Name of the resource.
	CRN *string `json:"crn,omitempty"`

	// Only used to attach Classic IaaS devices that have no CRN.
	Type *string `json:"type,omitempty"`

	// Only used to attach Classic IaaS devices that have no CRN. Id of Classic IaaS device. This is deprecated in favor of
	// the crn field.
	// Deprecated: this field is deprecated and may be removed in a future release.
	ID *float64 `json:"id,omitempty"`

	// A note about this resource.
	Note *string `json:"note,omitempty"`
}

// UnmarshalResourcePayload unmarshals an instance of ResourcePayload from the specified map of raw messages.
func UnmarshalResourcePayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourcePayload)
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "note", &obj.Note)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusPayload : Payload to update status of the case.
// Models which "extend" this model:
// - ResolvePayload
// - UnresolvePayload
// - AcceptPayload
type StatusPayload struct {
	// action to perform on the case.
	Action *string `json:"action" validate:"required"`

	// comment of resolution.
	Comment *string `json:"comment,omitempty"`

	// * 1: Client error
	// * 2: Defect found with Component/Service
	// * 3: Documentation Error
	// * 4: Solution found in forums
	// * 5: Solution found in public Documentation
	// * 6: Solution no longer required
	// * 7: Solution provided by IBM outside of support case
	// * 8: Solution provided by IBM support engineer.
	ResolutionCode *int64 `json:"resolution_code,omitempty"`
}

// Constants associated with the StatusPayload.Action property.
// action to perform on the case.
const (
	StatusPayloadActionAcceptConst    = "accept"
	StatusPayloadActionResolveConst   = "resolve"
	StatusPayloadActionUnresolveConst = "unresolve"
)

func (*StatusPayload) isaStatusPayload() bool {
	return true
}

type StatusPayloadIntf interface {
	isaStatusPayload() bool
}

// UnmarshalStatusPayload unmarshals an instance of StatusPayload from the specified map of raw messages.
func UnmarshalStatusPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "action", &discValue)
	if err != nil {
		err = fmt.Errorf("error unmarshalling discriminator property 'action': %s", err.Error())
		return
	}
	if discValue == "" {
		err = fmt.Errorf("required discriminator property 'action' not found in JSON object")
		return
	}
	if discValue == "resolve" {
		err = core.UnmarshalModel(m, "", result, UnmarshalResolvePayload)
	} else if discValue == "unresolve" {
		err = core.UnmarshalModel(m, "", result, UnmarshalUnresolvePayload)
	} else if discValue == "accept" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAcceptPayload)
	} else {
		err = fmt.Errorf("unrecognized value for discriminator property 'action': %s", discValue)
	}
	return
}

// UpdateCaseStatusOptions : The UpdateCaseStatus options.
type UpdateCaseStatusOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required,ne="`

	// Payload to update status of the case.
	StatusPayload StatusPayloadIntf `json:"StatusPayload" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCaseStatusOptions : Instantiate UpdateCaseStatusOptions
func (*CaseManagementV1) NewUpdateCaseStatusOptions(caseNumber string, statusPayload StatusPayloadIntf) *UpdateCaseStatusOptions {
	return &UpdateCaseStatusOptions{
		CaseNumber:    core.StringPtr(caseNumber),
		StatusPayload: statusPayload,
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (_options *UpdateCaseStatusOptions) SetCaseNumber(caseNumber string) *UpdateCaseStatusOptions {
	_options.CaseNumber = core.StringPtr(caseNumber)
	return _options
}

// SetStatusPayload : Allow user to set StatusPayload
func (_options *UpdateCaseStatusOptions) SetStatusPayload(statusPayload StatusPayloadIntf) *UpdateCaseStatusOptions {
	_options.StatusPayload = statusPayload
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCaseStatusOptions) SetHeaders(param map[string]string) *UpdateCaseStatusOptions {
	options.Headers = param
	return options
}

// UploadFileOptions : The UploadFile options.
type UploadFileOptions struct {
	// Unique identifier of a case.
	CaseNumber *string `json:"case_number" validate:"required,ne="`

	// file of supported types, 8MB in size limit.
	File []FileWithMetadata `json:"file" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUploadFileOptions : Instantiate UploadFileOptions
func (*CaseManagementV1) NewUploadFileOptions(caseNumber string, file []FileWithMetadata) *UploadFileOptions {
	return &UploadFileOptions{
		CaseNumber: core.StringPtr(caseNumber),
		File:       file,
	}
}

// SetCaseNumber : Allow user to set CaseNumber
func (_options *UploadFileOptions) SetCaseNumber(caseNumber string) *UploadFileOptions {
	_options.CaseNumber = core.StringPtr(caseNumber)
	return _options
}

// SetFile : Allow user to set File
func (_options *UploadFileOptions) SetFile(file []FileWithMetadata) *UploadFileOptions {
	_options.File = file
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UploadFileOptions) SetHeaders(param map[string]string) *UploadFileOptions {
	options.Headers = param
	return options
}

// User : User info in a case.
type User struct {
	// Full name of the user.
	Name *string `json:"name,omitempty"`

	// the ID realm.
	Realm *string `json:"realm" validate:"required"`

	// unique user ID in the realm specified by the type.
	UserID *string `json:"user_id" validate:"required"`
}

// Constants associated with the User.Realm property.
// the ID realm.
const (
	UserRealmBssConst   = "BSS"
	UserRealmIbmidConst = "IBMid"
	UserRealmSlConst    = "SL"
)

// NewUser : Instantiate User (Generic Model Constructor)
func (*CaseManagementV1) NewUser(realm string, userID string) (_model *User, err error) {
	_model = &User{
		Realm:  core.StringPtr(realm),
		UserID: core.StringPtr(userID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalUser unmarshals an instance of User from the specified map of raw messages.
func UnmarshalUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(User)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "realm", &obj.Realm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Watchlist : Payload to add/remove users to/from the case watchlist.
type Watchlist struct {
	// Array of user ID objects.
	Watchlist []User `json:"watchlist,omitempty"`
}

// UnmarshalWatchlist unmarshals an instance of Watchlist from the specified map of raw messages.
func UnmarshalWatchlist(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Watchlist)
	err = core.UnmarshalModel(m, "watchlist", &obj.Watchlist, UnmarshalUser)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WatchlistAddResponse : Response of a request when adding to watchlist.
type WatchlistAddResponse struct {
	// List of added user.
	Added []User `json:"added,omitempty"`

	// List of failed to add user.
	Failed []User `json:"failed,omitempty"`
}

// UnmarshalWatchlistAddResponse unmarshals an instance of WatchlistAddResponse from the specified map of raw messages.
func UnmarshalWatchlistAddResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WatchlistAddResponse)
	err = core.UnmarshalModel(m, "added", &obj.Added, UnmarshalUser)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failed", &obj.Failed, UnmarshalUser)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AcceptPayload : Payload to accept the proposed resolution of the case.
// This model "extends" StatusPayload
type AcceptPayload struct {
	// action to perform on the case.
	Action *string `json:"action" validate:"required"`

	// Comment about accepting the proposed resolution.
	Comment *string `json:"comment,omitempty"`
}

// Constants associated with the AcceptPayload.Action property.
// action to perform on the case.
const (
	AcceptPayloadActionAcceptConst    = "accept"
	AcceptPayloadActionResolveConst   = "resolve"
	AcceptPayloadActionUnresolveConst = "unresolve"
)

// NewAcceptPayload : Instantiate AcceptPayload (Generic Model Constructor)
func (*CaseManagementV1) NewAcceptPayload(action string) (_model *AcceptPayload, err error) {
	_model = &AcceptPayload{
		Action: core.StringPtr(action),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*AcceptPayload) isaStatusPayload() bool {
	return true
}

// UnmarshalAcceptPayload unmarshals an instance of AcceptPayload from the specified map of raw messages.
func UnmarshalAcceptPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AcceptPayload)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResolvePayload : Payload to resolve the case.
// This model "extends" StatusPayload
type ResolvePayload struct {
	// action to perform on the case.
	Action *string `json:"action" validate:"required"`

	// comment of resolution.
	Comment *string `json:"comment,omitempty"`

	// * 1: Client error
	// * 2: Defect found with Component/Service
	// * 3: Documentation Error
	// * 4: Solution found in forums
	// * 5: Solution found in public Documentation
	// * 6: Solution no longer required
	// * 7: Solution provided by IBM outside of support case
	// * 8: Solution provided by IBM support engineer.
	ResolutionCode *int64 `json:"resolution_code" validate:"required"`
}

// Constants associated with the ResolvePayload.Action property.
// action to perform on the case.
const (
	ResolvePayloadActionAcceptConst    = "accept"
	ResolvePayloadActionResolveConst   = "resolve"
	ResolvePayloadActionUnresolveConst = "unresolve"
)

// NewResolvePayload : Instantiate ResolvePayload (Generic Model Constructor)
func (*CaseManagementV1) NewResolvePayload(action string, resolutionCode int64) (_model *ResolvePayload, err error) {
	_model = &ResolvePayload{
		Action:         core.StringPtr(action),
		ResolutionCode: core.Int64Ptr(resolutionCode),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResolvePayload) isaStatusPayload() bool {
	return true
}

// UnmarshalResolvePayload unmarshals an instance of ResolvePayload from the specified map of raw messages.
func UnmarshalResolvePayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResolvePayload)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resolution_code", &obj.ResolutionCode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UnresolvePayload : Payload to unresolve the case.
// This model "extends" StatusPayload
type UnresolvePayload struct {
	// action to perform on the case.
	Action *string `json:"action" validate:"required"`

	// Comment why the case should be unresolved.
	Comment *string `json:"comment" validate:"required"`
}

// Constants associated with the UnresolvePayload.Action property.
// action to perform on the case.
const (
	UnresolvePayloadActionAcceptConst    = "accept"
	UnresolvePayloadActionResolveConst   = "resolve"
	UnresolvePayloadActionUnresolveConst = "unresolve"
)

// NewUnresolvePayload : Instantiate UnresolvePayload (Generic Model Constructor)
func (*CaseManagementV1) NewUnresolvePayload(action string, comment string) (_model *UnresolvePayload, err error) {
	_model = &UnresolvePayload{
		Action:  core.StringPtr(action),
		Comment: core.StringPtr(comment),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*UnresolvePayload) isaStatusPayload() bool {
	return true
}

// UnmarshalUnresolvePayload unmarshals an instance of UnresolvePayload from the specified map of raw messages.
func UnmarshalUnresolvePayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UnresolvePayload)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetCasesPager can be used to simplify the use of the "GetCases" method.
type GetCasesPager struct {
	hasNext     bool
	options     *GetCasesOptions
	client      *CaseManagementV1
	pageContext struct {
		next *int64
	}
}

// NewGetCasesPager returns a new GetCasesPager instance.
func (caseManagement *CaseManagementV1) NewGetCasesPager(options *GetCasesOptions) (pager *GetCasesPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy GetCasesOptions = *options
	pager = &GetCasesPager{
		hasNext: true,
		options: &optionsCopy,
		client:  caseManagement,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *GetCasesPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *GetCasesPager) GetNextWithContext(ctx context.Context) (page []Case, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.GetCasesWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Cases

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *GetCasesPager) GetAllWithContext(ctx context.Context) (allItems []Case, err error) {
	for pager.HasNext() {
		var nextPage []Case
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *GetCasesPager) GetNext() (page []Case, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *GetCasesPager) GetAll() (allItems []Case, err error) {
	return pager.GetAllWithContext(context.Background())
}
