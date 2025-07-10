/**
 * (C) Copyright IBM Corp. 2023.
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
 * IBM OpenAPI SDK Code Generator Version: 3.70.0-7df966bf-20230419-195904
 */

// Package usermanagementv1 : Operations and models for the UserManagementV1 service
package usermanagementv1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"net/http"
	"reflect"
	"time"
)

// UserManagementV1 : Manage the lifecycle of your users using User Management APIs.
//
// API Version: 1.0
type UserManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://user-management.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "user_management"

// UserManagementV1Options : Service options
type UserManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewUserManagementV1UsingExternalConfig : constructs an instance of UserManagementV1 with passed in options and external configuration.
func NewUserManagementV1UsingExternalConfig(options *UserManagementV1Options) (userManagement *UserManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	userManagement, err = NewUserManagementV1(options)
	if err != nil {
		return
	}

	err = userManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = userManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewUserManagementV1 : constructs an instance of UserManagementV1 with passed in options.
func NewUserManagementV1(options *UserManagementV1Options) (service *UserManagementV1, err error) {
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

	service = &UserManagementV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "userManagement" suitable for processing requests.
func (userManagement *UserManagementV1) Clone() *UserManagementV1 {
	if core.IsNil(userManagement) {
		return nil
	}
	clone := *userManagement
	clone.Service = userManagement.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (userManagement *UserManagementV1) SetServiceURL(url string) error {
	return userManagement.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (userManagement *UserManagementV1) GetServiceURL() string {
	return userManagement.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (userManagement *UserManagementV1) SetDefaultHeaders(headers http.Header) {
	userManagement.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (userManagement *UserManagementV1) SetEnableGzipCompression(enableGzip bool) {
	userManagement.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (userManagement *UserManagementV1) GetEnableGzipCompression() bool {
	return userManagement.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (userManagement *UserManagementV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	userManagement.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (userManagement *UserManagementV1) DisableRetries() {
	userManagement.Service.DisableRetries()
}

// ListUsers : List users
// Retrieve users in the account. You can use the IAM service token or a user token for authorization. To use this
// method, the requesting user or service ID must have at least the viewer, editor, or administrator role on the User
// Management service. If unrestricted view is enabled, the user can see all users in the same account without an IAM
// role. If restricted view is enabled and user has the viewer, editor, or administrator role on the user management
// service, the API returns all users in the account. If unrestricted view is enabled and the user does not have these
// roles, the API returns only the current user. Users are returned in a paginated list with a default limit of 100
// users. You can iterate through all users by following the `next_url` field. Additional substring search fields are
// supported to filter the users.
func (userManagement *UserManagementV1) ListUsers(listUsersOptions *ListUsersOptions) (result *UserList, response *core.DetailedResponse, err error) {
	return userManagement.ListUsersWithContext(context.Background(), listUsersOptions)
}

// ListUsersWithContext is an alternate form of the ListUsers method which supports a Context parameter
func (userManagement *UserManagementV1) ListUsersWithContext(ctx context.Context, listUsersOptions *ListUsersOptions) (result *UserList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listUsersOptions, "listUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listUsersOptions, "listUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *listUsersOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = userManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(userManagement.Service.Options.URL, `/v2/accounts/{account_id}/users`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "ListUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listUsersOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listUsersOptions.Limit))
	}
	if listUsersOptions.IncludeSettings != nil {
		builder.AddQuery("include_settings", fmt.Sprint(*listUsersOptions.IncludeSettings))
	}
	if listUsersOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*listUsersOptions.Search))
	}
	if listUsersOptions.Start != nil {
		builder.AddQuery("_start", fmt.Sprint(*listUsersOptions.Start))
	}
	if listUsersOptions.UserID != nil {
		builder.AddQuery("user_id", fmt.Sprint(*listUsersOptions.UserID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = userManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUserList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// InviteUsers : Invite users to an account
// Invite users to the account. You must use a user token for authorization. Service IDs can't invite users to the
// account. To use this method, the requesting user must have the editor or administrator role on the User Management
// service. For more information, see the [Inviting users](https://cloud.ibm.com/docs/account?topic=account-iamuserinv)
// documentation. You can specify the user account role and the corresponding IAM policy information in the request
// body. <br/><br/>When you invite a user to an account, the user is initially created in the `PROCESSING` state. After
// the user is successfully created, all specified permissions are configured, and the activation email is sent, the
// invited user is transitioned to the `PENDING` state. When the invited user clicks the activation email and creates
// and confirms their IBM Cloud account, the user is transitioned to `ACTIVE` state. If the user email is already
// verified, no email is generated.
func (userManagement *UserManagementV1) InviteUsers(inviteUsersOptions *InviteUsersOptions) (result *InvitedUserList, response *core.DetailedResponse, err error) {
	return userManagement.InviteUsersWithContext(context.Background(), inviteUsersOptions)
}

// InviteUsersWithContext is an alternate form of the InviteUsers method which supports a Context parameter
func (userManagement *UserManagementV1) InviteUsersWithContext(ctx context.Context, inviteUsersOptions *InviteUsersOptions) (result *InvitedUserList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(inviteUsersOptions, "inviteUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(inviteUsersOptions, "inviteUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *inviteUsersOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = userManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(userManagement.Service.Options.URL, `/v2/accounts/{account_id}/users`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range inviteUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "InviteUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if inviteUsersOptions.Users != nil {
		body["users"] = inviteUsersOptions.Users
	}
	if inviteUsersOptions.IamPolicy != nil {
		body["iam_policy"] = inviteUsersOptions.IamPolicy
	}
	if inviteUsersOptions.AccessGroups != nil {
		body["access_groups"] = inviteUsersOptions.AccessGroups
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
	response, err = userManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInvitedUserList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetUserProfile : Get user profile
// Retrieve a user's profile by the user's IAM ID in your account. You can use the IAM service token or a user token for
// authorization. To use this method, the requesting user or service ID must have at least the viewer, editor, or
// administrator role on the User Management service.
func (userManagement *UserManagementV1) GetUserProfile(getUserProfileOptions *GetUserProfileOptions) (result *UserProfile, response *core.DetailedResponse, err error) {
	return userManagement.GetUserProfileWithContext(context.Background(), getUserProfileOptions)
}

// GetUserProfileWithContext is an alternate form of the GetUserProfile method which supports a Context parameter
func (userManagement *UserManagementV1) GetUserProfileWithContext(ctx context.Context, getUserProfileOptions *GetUserProfileOptions) (result *UserProfile, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getUserProfileOptions, "getUserProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getUserProfileOptions, "getUserProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *getUserProfileOptions.AccountID,
		"iam_id":     *getUserProfileOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = userManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(userManagement.Service.Options.URL, `/v2/accounts/{account_id}/users/{iam_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getUserProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "GetUserProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getUserProfileOptions.IncludeActivity != nil {
		builder.AddQuery("include_activity", fmt.Sprint(*getUserProfileOptions.IncludeActivity))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = userManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUserProfile)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateUserProfile : Partially update user profile
// Partially update a user's profile by user's IAM ID. You can use the IAM service token or a user token for
// authorization. To use this method, the requesting user or service ID must have at least the editor or administrator
// role on the User Management service. A user or service ID with these roles can change a user's state between
// `ACTIVE`, `VPN_ONLY`, or `DISABLED_CLASSIC_INFRASTRUCTURE`, but they can't change the state to `PROCESSING` or
// `PENDING` because these are system states. For other request body fields, a user can update their own profile without
// having User Management service permissions.
func (userManagement *UserManagementV1) UpdateUserProfile(updateUserProfileOptions *UpdateUserProfileOptions) (response *core.DetailedResponse, err error) {
	return userManagement.UpdateUserProfileWithContext(context.Background(), updateUserProfileOptions)
}

// UpdateUserProfileWithContext is an alternate form of the UpdateUserProfile method which supports a Context parameter
func (userManagement *UserManagementV1) UpdateUserProfileWithContext(ctx context.Context, updateUserProfileOptions *UpdateUserProfileOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateUserProfileOptions, "updateUserProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateUserProfileOptions, "updateUserProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *updateUserProfileOptions.AccountID,
		"iam_id":     *updateUserProfileOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = userManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(userManagement.Service.Options.URL, `/v2/accounts/{account_id}/users/{iam_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateUserProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "UpdateUserProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	if updateUserProfileOptions.IncludeActivity != nil {
		builder.AddQuery("include_activity", fmt.Sprint(*updateUserProfileOptions.IncludeActivity))
	}

	body := make(map[string]interface{})
	if updateUserProfileOptions.Firstname != nil {
		body["firstname"] = updateUserProfileOptions.Firstname
	}
	if updateUserProfileOptions.Lastname != nil {
		body["lastname"] = updateUserProfileOptions.Lastname
	}
	if updateUserProfileOptions.State != nil {
		body["state"] = updateUserProfileOptions.State
	}
	if updateUserProfileOptions.Email != nil {
		body["email"] = updateUserProfileOptions.Email
	}
	if updateUserProfileOptions.Phonenumber != nil {
		body["phonenumber"] = updateUserProfileOptions.Phonenumber
	}
	if updateUserProfileOptions.Altphonenumber != nil {
		body["altphonenumber"] = updateUserProfileOptions.Altphonenumber
	}
	if updateUserProfileOptions.Photo != nil {
		body["photo"] = updateUserProfileOptions.Photo
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// RemoveUser : Remove user from account
// Remove users from an account by user's IAM ID. You must use a user token for authorization. Service IDs can't remove
// users from an account. To use this method, the requesting user must have the editor or administrator role on the User
// Management service. For more information, see the [Removing
// users](https://cloud.ibm.com/docs/account?topic=account-remove) documentation.
func (userManagement *UserManagementV1) RemoveUser(removeUserOptions *RemoveUserOptions) (response *core.DetailedResponse, err error) {
	return userManagement.RemoveUserWithContext(context.Background(), removeUserOptions)
}

// RemoveUserWithContext is an alternate form of the RemoveUser method which supports a Context parameter
func (userManagement *UserManagementV1) RemoveUserWithContext(ctx context.Context, removeUserOptions *RemoveUserOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeUserOptions, "removeUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeUserOptions, "removeUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *removeUserOptions.AccountID,
		"iam_id":     *removeUserOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = userManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(userManagement.Service.Options.URL, `/v2/accounts/{account_id}/users/{iam_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "RemoveUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if removeUserOptions.IncludeActivity != nil {
		builder.AddQuery("include_activity", fmt.Sprint(*removeUserOptions.IncludeActivity))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// Accept : Accept an invitation
// Accept a user invitation to an account. You can use the user's token for authorization. To use this method, the
// requesting user must provide the account ID for the account that they are accepting an invitation for. If the user
// already accepted the invitation request, it returns 204 with no response body.
func (userManagement *UserManagementV1) Accept(acceptOptions *AcceptOptions) (response *core.DetailedResponse, err error) {
	return userManagement.AcceptWithContext(context.Background(), acceptOptions)
}

// AcceptWithContext is an alternate form of the Accept method which supports a Context parameter
func (userManagement *UserManagementV1) AcceptWithContext(ctx context.Context, acceptOptions *AcceptOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(acceptOptions, "acceptOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = userManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(userManagement.Service.Options.URL, `/v2/users/accept`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range acceptOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "Accept")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if acceptOptions.AccountID != nil {
		body["account_id"] = acceptOptions.AccountID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// V3RemoveUser : Remove user from account (Asynchronous)
// Remove users from an account by using the user's IAM ID. You must use a user token for authorization. Service IDs
// can't remove users from an account. If removing the user fails it will set the user's state to ERROR_WHILE_DELETING.
// To use this method, the requesting user must have the editor or administrator role on the User Management service.
// For more information, see the [Removing users](https://cloud.ibm.com/docs/account?topic=account-remove)
// documentation.
func (userManagement *UserManagementV1) V3RemoveUser(v3RemoveUserOptions *V3RemoveUserOptions) (response *core.DetailedResponse, err error) {
	return userManagement.V3RemoveUserWithContext(context.Background(), v3RemoveUserOptions)
}

// V3RemoveUserWithContext is an alternate form of the V3RemoveUser method which supports a Context parameter
func (userManagement *UserManagementV1) V3RemoveUserWithContext(ctx context.Context, v3RemoveUserOptions *V3RemoveUserOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(v3RemoveUserOptions, "v3RemoveUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(v3RemoveUserOptions, "v3RemoveUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *v3RemoveUserOptions.AccountID,
		"iam_id":     *v3RemoveUserOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = userManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(userManagement.Service.Options.URL, `/v3/accounts/{account_id}/users/{iam_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range v3RemoveUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "V3RemoveUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// GetUserSettings : Get user settings
// Retrieve a user's settings by the user's IAM ID. You can use the IAM service token or a user token for authorization.
// To use this method, the requesting user or service ID must have the viewer, editor, or administrator role on the User
// Management service. <br/><br/>The user settings have several fields. The `language` field is the language setting for
// the user interface display language. The `notification_language` field is the language setting for phone and email
// notifications. The `allowed_ip_addresses` field specifies a list of IP addresses that the user can log in and perform
// operations from as described in [Allowing specific IP addresses for a
// user](https://cloud.ibm.com/docs/account?topic=account-ips). For information about the `self_manage` field, review
// information about the [user-managed login setting](https://cloud.ibm.com/docs/account?topic=account-types).
func (userManagement *UserManagementV1) GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions) (result *UserSettings, response *core.DetailedResponse, err error) {
	return userManagement.GetUserSettingsWithContext(context.Background(), getUserSettingsOptions)
}

// GetUserSettingsWithContext is an alternate form of the GetUserSettings method which supports a Context parameter
func (userManagement *UserManagementV1) GetUserSettingsWithContext(ctx context.Context, getUserSettingsOptions *GetUserSettingsOptions) (result *UserSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getUserSettingsOptions, "getUserSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getUserSettingsOptions, "getUserSettingsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *getUserSettingsOptions.AccountID,
		"iam_id":     *getUserSettingsOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = userManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(userManagement.Service.Options.URL, `/v2/accounts/{account_id}/users/{iam_id}/settings`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getUserSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "GetUserSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = userManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUserSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateUserSettings : Partially update user settings
// Update a user's settings by the user's IAM ID. You can use the IAM service token or a user token for authorization.
// To fully use this method, the user or service ID must have the editor or administrator role on the User Management
// service. Without these roles, a user can update only their own `language` or `notification_language` fields. If
// `self_manage` is `true`, the user can also update the `allowed_ip_addresses` field.
func (userManagement *UserManagementV1) UpdateUserSettings(updateUserSettingsOptions *UpdateUserSettingsOptions) (response *core.DetailedResponse, err error) {
	return userManagement.UpdateUserSettingsWithContext(context.Background(), updateUserSettingsOptions)
}

// UpdateUserSettingsWithContext is an alternate form of the UpdateUserSettings method which supports a Context parameter
func (userManagement *UserManagementV1) UpdateUserSettingsWithContext(ctx context.Context, updateUserSettingsOptions *UpdateUserSettingsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateUserSettingsOptions, "updateUserSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateUserSettingsOptions, "updateUserSettingsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *updateUserSettingsOptions.AccountID,
		"iam_id":     *updateUserSettingsOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = userManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(userManagement.Service.Options.URL, `/v2/accounts/{account_id}/users/{iam_id}/settings`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateUserSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "UpdateUserSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateUserSettingsOptions.Language != nil {
		body["language"] = updateUserSettingsOptions.Language
	}
	if updateUserSettingsOptions.NotificationLanguage != nil {
		body["notification_language"] = updateUserSettingsOptions.NotificationLanguage
	}
	if updateUserSettingsOptions.AllowedIPAddresses != nil {
		body["allowed_ip_addresses"] = updateUserSettingsOptions.AllowedIPAddresses
	}
	if updateUserSettingsOptions.SelfManage != nil {
		body["self_manage"] = updateUserSettingsOptions.SelfManage
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// AcceptOptions : The Accept options.
type AcceptOptions struct {
	// The account ID.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAcceptOptions : Instantiate AcceptOptions
func (*UserManagementV1) NewAcceptOptions() *AcceptOptions {
	return &AcceptOptions{}
}

// SetAccountID : Allow user to set AccountID
func (_options *AcceptOptions) SetAccountID(accountID string) *AcceptOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AcceptOptions) SetHeaders(param map[string]string) *AcceptOptions {
	options.Headers = param
	return options
}

// GetUserProfileOptions : The GetUserProfile options.
type GetUserProfileOptions struct {
	// The account ID of the specified user.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The user's IAM ID.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// Include activity information of the user, such as the last authentication timestamp.
	IncludeActivity *string `json:"include_activity,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetUserProfileOptions : Instantiate GetUserProfileOptions
func (*UserManagementV1) NewGetUserProfileOptions(accountID string, iamID string) *GetUserProfileOptions {
	return &GetUserProfileOptions{
		AccountID: core.StringPtr(accountID),
		IamID:     core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *GetUserProfileOptions) SetAccountID(accountID string) *GetUserProfileOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *GetUserProfileOptions) SetIamID(iamID string) *GetUserProfileOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetIncludeActivity : Allow user to set IncludeActivity
func (_options *GetUserProfileOptions) SetIncludeActivity(includeActivity string) *GetUserProfileOptions {
	_options.IncludeActivity = core.StringPtr(includeActivity)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetUserProfileOptions) SetHeaders(param map[string]string) *GetUserProfileOptions {
	options.Headers = param
	return options
}

// GetUserSettingsOptions : The GetUserSettings options.
type GetUserSettingsOptions struct {
	// The account ID of the specified user.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The user's IAM ID.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetUserSettingsOptions : Instantiate GetUserSettingsOptions
func (*UserManagementV1) NewGetUserSettingsOptions(accountID string, iamID string) *GetUserSettingsOptions {
	return &GetUserSettingsOptions{
		AccountID: core.StringPtr(accountID),
		IamID:     core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *GetUserSettingsOptions) SetAccountID(accountID string) *GetUserSettingsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *GetUserSettingsOptions) SetIamID(iamID string) *GetUserSettingsOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetUserSettingsOptions) SetHeaders(param map[string]string) *GetUserSettingsOptions {
	options.Headers = param
	return options
}

// InviteUsersOptions : The InviteUsers options.
type InviteUsersOptions struct {
	// The account ID of the specified user.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// A list of users to be invited.
	Users []InviteUser `json:"users,omitempty"`

	// A list of IAM policies.
	IamPolicy []InviteUserIamPolicy `json:"iam_policy,omitempty"`

	// A list of access groups.
	AccessGroups []string `json:"access_groups,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInviteUsersOptions : Instantiate InviteUsersOptions
func (*UserManagementV1) NewInviteUsersOptions(accountID string) *InviteUsersOptions {
	return &InviteUsersOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *InviteUsersOptions) SetAccountID(accountID string) *InviteUsersOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetUsers : Allow user to set Users
func (_options *InviteUsersOptions) SetUsers(users []InviteUser) *InviteUsersOptions {
	_options.Users = users
	return _options
}

// SetIamPolicy : Allow user to set IamPolicy
func (_options *InviteUsersOptions) SetIamPolicy(iamPolicy []InviteUserIamPolicy) *InviteUsersOptions {
	_options.IamPolicy = iamPolicy
	return _options
}

// SetAccessGroups : Allow user to set AccessGroups
func (_options *InviteUsersOptions) SetAccessGroups(accessGroups []string) *InviteUsersOptions {
	_options.AccessGroups = accessGroups
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InviteUsersOptions) SetHeaders(param map[string]string) *InviteUsersOptions {
	options.Headers = param
	return options
}

// InvitedUser : Information about a user that has been invited to join an account.
type InvitedUser struct {
	// The email address associated with the invited user.
	Email *string `json:"email,omitempty"`

	// The id associated with the invited user.
	ID *string `json:"id,omitempty"`

	// The state of the invitation for the user.
	State *string `json:"state,omitempty"`
}

// UnmarshalInvitedUser unmarshals an instance of InvitedUser from the specified map of raw messages.
func UnmarshalInvitedUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InvitedUser)
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InvitedUserList : A collection of invited users.  This is the response returned by the invite_users operation.
type InvitedUserList struct {
	// The list of users that have been invited to join the account.
	Resources []InvitedUser `json:"resources,omitempty"`
}

// UnmarshalInvitedUserList unmarshals an instance of InvitedUserList from the specified map of raw messages.
func UnmarshalInvitedUserList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InvitedUserList)
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalInvitedUser)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListUsersOptions : The ListUsers options.
type ListUsersOptions struct {
	// The account ID of the specified user.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The number of results to be returned.
	Limit *int64 `json:"limit,omitempty"`

	// The user settings to be returned. Set to true to view language, allowed IP address, and authentication settings.
	IncludeSettings *bool `json:"include_settings,omitempty"`

	// The desired search results to be returned. To view the list of users with the additional search filter, use the
	// following query options: `firstname`, `lastname`, `email`, `state`, `substate`, `iam_id`, `realm`, and `userId`.
	// HTML URL encoding for the search query and `:` must be used. For example, search=state%3AINVALID returns a list of
	// invalid users. Multiple search queries can be combined to obtain `OR` results using `,` operator (not URL encoded).
	// For example, search=state%3AINVALID,email%3Amail.test.ibm.com.
	Search *string `json:"search,omitempty"`

	// An optional token that indicates the beginning of the page of results to be returned. If omitted, the first page of
	// results is returned. This value is obtained from the 'next_url' field of the operation response.
	Start *string `json:"_start,omitempty"`

	// Filter users based on their user ID.
	UserID *string `json:"user_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListUsersOptions : Instantiate ListUsersOptions
func (*UserManagementV1) NewListUsersOptions(accountID string) *ListUsersOptions {
	return &ListUsersOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *ListUsersOptions) SetAccountID(accountID string) *ListUsersOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListUsersOptions) SetLimit(limit int64) *ListUsersOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetIncludeSettings : Allow user to set IncludeSettings
func (_options *ListUsersOptions) SetIncludeSettings(includeSettings bool) *ListUsersOptions {
	_options.IncludeSettings = core.BoolPtr(includeSettings)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *ListUsersOptions) SetSearch(search string) *ListUsersOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListUsersOptions) SetStart(start string) *ListUsersOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetUserID : Allow user to set UserID
func (_options *ListUsersOptions) SetUserID(userID string) *ListUsersOptions {
	_options.UserID = core.StringPtr(userID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListUsersOptions) SetHeaders(param map[string]string) *ListUsersOptions {
	options.Headers = param
	return options
}

// RemoveUserOptions : The RemoveUser options.
type RemoveUserOptions struct {
	// The account ID of the specified user.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The user's IAM ID.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// Include activity information of the user, such as the last authentication timestamp.
	IncludeActivity *string `json:"include_activity,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveUserOptions : Instantiate RemoveUserOptions
func (*UserManagementV1) NewRemoveUserOptions(accountID string, iamID string) *RemoveUserOptions {
	return &RemoveUserOptions{
		AccountID: core.StringPtr(accountID),
		IamID:     core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *RemoveUserOptions) SetAccountID(accountID string) *RemoveUserOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *RemoveUserOptions) SetIamID(iamID string) *RemoveUserOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetIncludeActivity : Allow user to set IncludeActivity
func (_options *RemoveUserOptions) SetIncludeActivity(includeActivity string) *RemoveUserOptions {
	_options.IncludeActivity = core.StringPtr(includeActivity)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveUserOptions) SetHeaders(param map[string]string) *RemoveUserOptions {
	options.Headers = param
	return options
}

// UpdateUserProfileOptions : The UpdateUserProfile options.
type UpdateUserProfileOptions struct {
	// The account ID of the specified user.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The user's IAM ID.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// The first name of the user.
	Firstname *string `json:"firstname,omitempty"`

	// The last name of the user.
	Lastname *string `json:"lastname,omitempty"`

	// The state of the user. Possible values are `PROCESSING`, `PENDING`, `ACTIVE`, `DISABLED_CLASSIC_INFRASTRUCTURE`, and
	// `VPN_ONLY`.
	State *string `json:"state,omitempty"`

	// The email address of the user.
	Email *string `json:"email,omitempty"`

	// The phone number of the user.
	Phonenumber *string `json:"phonenumber,omitempty"`

	// The alternative phone number of the user.
	Altphonenumber *string `json:"altphonenumber,omitempty"`

	// A link to a photo of the user.
	Photo *string `json:"photo,omitempty"`

	// Include activity information of the user, such as the last authentication timestamp.
	IncludeActivity *string `json:"include_activity,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateUserProfileOptions : Instantiate UpdateUserProfileOptions
func (*UserManagementV1) NewUpdateUserProfileOptions(accountID string, iamID string) *UpdateUserProfileOptions {
	return &UpdateUserProfileOptions{
		AccountID: core.StringPtr(accountID),
		IamID:     core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *UpdateUserProfileOptions) SetAccountID(accountID string) *UpdateUserProfileOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *UpdateUserProfileOptions) SetIamID(iamID string) *UpdateUserProfileOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetFirstname : Allow user to set Firstname
func (_options *UpdateUserProfileOptions) SetFirstname(firstname string) *UpdateUserProfileOptions {
	_options.Firstname = core.StringPtr(firstname)
	return _options
}

// SetLastname : Allow user to set Lastname
func (_options *UpdateUserProfileOptions) SetLastname(lastname string) *UpdateUserProfileOptions {
	_options.Lastname = core.StringPtr(lastname)
	return _options
}

// SetState : Allow user to set State
func (_options *UpdateUserProfileOptions) SetState(state string) *UpdateUserProfileOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetEmail : Allow user to set Email
func (_options *UpdateUserProfileOptions) SetEmail(email string) *UpdateUserProfileOptions {
	_options.Email = core.StringPtr(email)
	return _options
}

// SetPhonenumber : Allow user to set Phonenumber
func (_options *UpdateUserProfileOptions) SetPhonenumber(phonenumber string) *UpdateUserProfileOptions {
	_options.Phonenumber = core.StringPtr(phonenumber)
	return _options
}

// SetAltphonenumber : Allow user to set Altphonenumber
func (_options *UpdateUserProfileOptions) SetAltphonenumber(altphonenumber string) *UpdateUserProfileOptions {
	_options.Altphonenumber = core.StringPtr(altphonenumber)
	return _options
}

// SetPhoto : Allow user to set Photo
func (_options *UpdateUserProfileOptions) SetPhoto(photo string) *UpdateUserProfileOptions {
	_options.Photo = core.StringPtr(photo)
	return _options
}

// SetIncludeActivity : Allow user to set IncludeActivity
func (_options *UpdateUserProfileOptions) SetIncludeActivity(includeActivity string) *UpdateUserProfileOptions {
	_options.IncludeActivity = core.StringPtr(includeActivity)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateUserProfileOptions) SetHeaders(param map[string]string) *UpdateUserProfileOptions {
	options.Headers = param
	return options
}

// UpdateUserSettingsOptions : The UpdateUserSettings options.
type UpdateUserSettingsOptions struct {
	// The account ID of the specified user.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The user's IAM ID.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// The console UI language. By default, this field is empty.
	Language *string `json:"language,omitempty"`

	// The language for email and phone notifications. By default, this field is empty.
	NotificationLanguage *string `json:"notification_language,omitempty"`

	// A comma-separated list of IP addresses.
	AllowedIPAddresses *string `json:"allowed_ip_addresses,omitempty"`

	// Whether user managed login is enabled. The default value is `false`.
	SelfManage *bool `json:"self_manage,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateUserSettingsOptions : Instantiate UpdateUserSettingsOptions
func (*UserManagementV1) NewUpdateUserSettingsOptions(accountID string, iamID string) *UpdateUserSettingsOptions {
	return &UpdateUserSettingsOptions{
		AccountID: core.StringPtr(accountID),
		IamID:     core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *UpdateUserSettingsOptions) SetAccountID(accountID string) *UpdateUserSettingsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *UpdateUserSettingsOptions) SetIamID(iamID string) *UpdateUserSettingsOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetLanguage : Allow user to set Language
func (_options *UpdateUserSettingsOptions) SetLanguage(language string) *UpdateUserSettingsOptions {
	_options.Language = core.StringPtr(language)
	return _options
}

// SetNotificationLanguage : Allow user to set NotificationLanguage
func (_options *UpdateUserSettingsOptions) SetNotificationLanguage(notificationLanguage string) *UpdateUserSettingsOptions {
	_options.NotificationLanguage = core.StringPtr(notificationLanguage)
	return _options
}

// SetAllowedIPAddresses : Allow user to set AllowedIPAddresses
func (_options *UpdateUserSettingsOptions) SetAllowedIPAddresses(allowedIPAddresses string) *UpdateUserSettingsOptions {
	_options.AllowedIPAddresses = core.StringPtr(allowedIPAddresses)
	return _options
}

// SetSelfManage : Allow user to set SelfManage
func (_options *UpdateUserSettingsOptions) SetSelfManage(selfManage bool) *UpdateUserSettingsOptions {
	_options.SelfManage = core.BoolPtr(selfManage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateUserSettingsOptions) SetHeaders(param map[string]string) *UpdateUserSettingsOptions {
	options.Headers = param
	return options
}

// UserList : The users returned.
type UserList struct {
	// The number of users returned.
	TotalResults *int64 `json:"total_results" validate:"required"`

	// A limit to the number of users returned in a page.
	Limit *int64 `json:"limit" validate:"required"`

	// The first URL of the get users API.
	FirstURL *string `json:"first_url,omitempty"`

	// The next URL of the get users API.
	NextURL *string `json:"next_url,omitempty"`

	// A list of users in the account.
	Resources []UserProfile `json:"resources,omitempty"`
}

// UnmarshalUserList unmarshals an instance of UserList from the specified map of raw messages.
func UnmarshalUserList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserList)
	err = core.UnmarshalPrimitive(m, "total_results", &obj.TotalResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first_url", &obj.FirstURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalUserProfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *UserList) GetNextStart() (*string, error) {
	if core.IsNil(resp.NextURL) {
		return nil, nil
	}
	_start, err := core.GetQueryParam(resp.NextURL, "_start")
	if err != nil || _start == nil {
		return nil, err
	}
	return _start, nil
}

// UserProfile : Returned the user profile.
type UserProfile struct {
	// An alphanumeric value identifying the user profile.
	ID *string `json:"id,omitempty"`

	// An alphanumeric value identifying the user's IAM ID.
	IamID *string `json:"iam_id,omitempty"`

	// The realm of the user. The value is either `IBMid` or `SL`.
	Realm *string `json:"realm,omitempty"`

	// The user ID used for login.
	UserID *string `json:"user_id,omitempty"`

	// The first name of the user.
	Firstname *string `json:"firstname,omitempty"`

	// The last name of the user.
	Lastname *string `json:"lastname,omitempty"`

	// The state of the user. Possible values are `PROCESSING`, `PENDING`, `ACTIVE`, `DISABLED_CLASSIC_INFRASTRUCTURE`, and
	// `VPN_ONLY`.
	State *string `json:"state,omitempty"`

	// The email address of the user.
	Email *string `json:"email,omitempty"`

	// The phone number of the user.
	Phonenumber *string `json:"phonenumber,omitempty"`

	// The alternative phone number of the user.
	Altphonenumber *string `json:"altphonenumber,omitempty"`

	// A link to a photo of the user.
	Photo *string `json:"photo,omitempty"`

	// An alphanumeric value identifying the account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The timestamp for when the user was added to the account.
	AddedOn *string `json:"added_on,omitempty"`
}

// UnmarshalUserProfile unmarshals an instance of UserProfile from the specified map of raw messages.
func UnmarshalUserProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserProfile)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
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
	err = core.UnmarshalPrimitive(m, "firstname", &obj.Firstname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "lastname", &obj.Lastname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "phonenumber", &obj.Phonenumber)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "altphonenumber", &obj.Altphonenumber)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "photo", &obj.Photo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "added_on", &obj.AddedOn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserSettings : The user settings returned.
type UserSettings struct {
	// The console UI language. By default, this field is empty.
	Language *string `json:"language,omitempty"`

	// The language for email and phone notifications. By default, this field is empty.
	NotificationLanguage *string `json:"notification_language,omitempty"`

	// A comma-separated list of IP addresses.
	AllowedIPAddresses *string `json:"allowed_ip_addresses,omitempty"`

	// Whether user managed login is enabled. The default value is `false`.
	SelfManage *bool `json:"self_manage,omitempty"`
}

// UnmarshalUserSettings unmarshals an instance of UserSettings from the specified map of raw messages.
func UnmarshalUserSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserSettings)
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "notification_language", &obj.NotificationLanguage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allowed_ip_addresses", &obj.AllowedIPAddresses)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "self_manage", &obj.SelfManage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// V3RemoveUserOptions : The V3RemoveUser options.
type V3RemoveUserOptions struct {
	// The account ID of the specified user.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The user's IAM ID.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewV3RemoveUserOptions : Instantiate V3RemoveUserOptions
func (*UserManagementV1) NewV3RemoveUserOptions(accountID string, iamID string) *V3RemoveUserOptions {
	return &V3RemoveUserOptions{
		AccountID: core.StringPtr(accountID),
		IamID:     core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *V3RemoveUserOptions) SetAccountID(accountID string) *V3RemoveUserOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *V3RemoveUserOptions) SetIamID(iamID string) *V3RemoveUserOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *V3RemoveUserOptions) SetHeaders(param map[string]string) *V3RemoveUserOptions {
	options.Headers = param
	return options
}

// Attribute : An attribute/value pair.
type Attribute struct {
	// The name of the attribute.
	Name *string `json:"name,omitempty"`

	// The value of the attribute.
	Value *string `json:"value,omitempty"`
}

// UnmarshalAttribute unmarshals an instance of Attribute from the specified map of raw messages.
func UnmarshalAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Attribute)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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

// InviteUser : Invite a user.
type InviteUser struct {
	// The email of the user to be invited.
	Email *string `json:"email,omitempty"`

	// The account role of the user to be invited.
	AccountRole *string `json:"account_role,omitempty"`
}

// UnmarshalInviteUser unmarshals an instance of InviteUser from the specified map of raw messages.
func UnmarshalInviteUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InviteUser)
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_role", &obj.AccountRole)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InviteUserIamPolicy : Invite a user to an IAM policy.
type InviteUserIamPolicy struct {
	// The policy type. This can be either "access" or "authorization".
	Type *string `json:"type" validate:"required"`

	// A list of IAM roles.
	Roles []Role `json:"roles,omitempty"`

	// A list of resources.
	Resources []Resource `json:"resources,omitempty"`
}

// NewInviteUserIamPolicy : Instantiate InviteUserIamPolicy (Generic Model Constructor)
func (*UserManagementV1) NewInviteUserIamPolicy(typeVar string) (_model *InviteUserIamPolicy, err error) {
	_model = &InviteUserIamPolicy{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalInviteUserIamPolicy unmarshals an instance of InviteUserIamPolicy from the specified map of raw messages.
func UnmarshalInviteUserIamPolicy(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InviteUserIamPolicy)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "roles", &obj.Roles, UnmarshalRole)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resource : A collection of attribute value pairs.
type Resource struct {
	// A list of IAM attributes.
	Attributes []Attribute `json:"attributes,omitempty"`
}

// UnmarshalResource unmarshals an instance of Resource from the specified map of raw messages.
func UnmarshalResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resource)
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Role : The role of an IAM policy.
type Role struct {
	// An alphanumeric value identifying the origin.
	RoleID *string `json:"role_id,omitempty"`
}

// UnmarshalRole unmarshals an instance of Role from the specified map of raw messages.
func UnmarshalRole(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Role)
	err = core.UnmarshalPrimitive(m, "role_id", &obj.RoleID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UsersPager can be used to simplify the use of the "ListUsers" method.
type UsersPager struct {
	hasNext     bool
	options     *ListUsersOptions
	client      *UserManagementV1
	pageContext struct {
		next *string
	}
}

// NewUsersPager returns a new UsersPager instance.
func (userManagement *UserManagementV1) NewUsersPager(options *ListUsersOptions) (pager *UsersPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListUsersOptions = *options
	pager = &UsersPager{
		hasNext: true,
		options: &optionsCopy,
		client:  userManagement,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *UsersPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *UsersPager) GetNextWithContext(ctx context.Context) (page []UserProfile, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListUsersWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.NextURL != nil {
		var _start *string
		_start, err = core.GetQueryParam(result.NextURL, "_start")
		if err != nil {
			err = fmt.Errorf("error retrieving '_start' query parameter from URL '%s': %s", *result.NextURL, err.Error())
			return
		}
		next = _start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *UsersPager) GetAllWithContext(ctx context.Context) (allItems []UserProfile, err error) {
	for pager.HasNext() {
		var nextPage []UserProfile
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *UsersPager) GetNext() (page []UserProfile, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *UsersPager) GetAll() (allItems []UserProfile, err error) {
	return pager.GetAllWithContext(context.Background())
}
