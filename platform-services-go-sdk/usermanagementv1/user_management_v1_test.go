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

package usermanagementv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/usermanagementv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`UserManagementV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(userManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(userManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				URL: "https://usermanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(userManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_MANAGEMENT_URL":       "https://usermanagementv1/api",
				"USER_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{})
				Expect(userManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := userManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != userManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(userManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(userManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(userManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := userManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != userManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(userManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(userManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{})
				err := userManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := userManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != userManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(userManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(userManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_MANAGEMENT_URL":       "https://usermanagementv1/api",
				"USER_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(userManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_MANAGEMENT_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(userManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = usermanagementv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListUsers(listUsersOptions *ListUsersOptions) - Operation response error`, func() {
		listUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// TODO: Add check for include_settings query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["user_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListUsers with error: Operation response processing error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(usermanagementv1.ListUsersOptions)
				listUsersOptionsModel.AccountID = core.StringPtr("testString")
				listUsersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listUsersOptionsModel.IncludeSettings = core.BoolPtr(true)
				listUsersOptionsModel.Search = core.StringPtr("testString")
				listUsersOptionsModel.Start = core.StringPtr("testString")
				listUsersOptionsModel.UserID = core.StringPtr("testString")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userManagementService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				userManagementService.EnableRetries(0, 0)
				result, response, operationErr = userManagementService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListUsers(listUsersOptions *ListUsersOptions)`, func() {
		listUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// TODO: Add check for include_settings query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["user_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_results": 12, "limit": 5, "first_url": "FirstURL", "next_url": "NextURL", "resources": [{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID", "added_on": "AddedOn"}]}`)
				}))
			})
			It(`Invoke ListUsers successfully with retries`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())
				userManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(usermanagementv1.ListUsersOptions)
				listUsersOptionsModel.AccountID = core.StringPtr("testString")
				listUsersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listUsersOptionsModel.IncludeSettings = core.BoolPtr(true)
				listUsersOptionsModel.Search = core.StringPtr("testString")
				listUsersOptionsModel.Start = core.StringPtr("testString")
				listUsersOptionsModel.UserID = core.StringPtr("testString")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := userManagementService.ListUsersWithContext(ctx, listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				userManagementService.DisableRetries()
				result, response, operationErr := userManagementService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = userManagementService.ListUsersWithContext(ctx, listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// TODO: Add check for include_settings query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["user_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_results": 12, "limit": 5, "first_url": "FirstURL", "next_url": "NextURL", "resources": [{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID", "added_on": "AddedOn"}]}`)
				}))
			})
			It(`Invoke ListUsers successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userManagementService.ListUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(usermanagementv1.ListUsersOptions)
				listUsersOptionsModel.AccountID = core.StringPtr("testString")
				listUsersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listUsersOptionsModel.IncludeSettings = core.BoolPtr(true)
				listUsersOptionsModel.Search = core.StringPtr("testString")
				listUsersOptionsModel.Start = core.StringPtr("testString")
				listUsersOptionsModel.UserID = core.StringPtr("testString")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userManagementService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListUsers with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(usermanagementv1.ListUsersOptions)
				listUsersOptionsModel.AccountID = core.StringPtr("testString")
				listUsersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listUsersOptionsModel.IncludeSettings = core.BoolPtr(true)
				listUsersOptionsModel.Search = core.StringPtr("testString")
				listUsersOptionsModel.Start = core.StringPtr("testString")
				listUsersOptionsModel.UserID = core.StringPtr("testString")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userManagementService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListUsersOptions model with no property values
				listUsersOptionsModelNew := new(usermanagementv1.ListUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userManagementService.ListUsers(listUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListUsers successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(usermanagementv1.ListUsersOptions)
				listUsersOptionsModel.AccountID = core.StringPtr("testString")
				listUsersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listUsersOptionsModel.IncludeSettings = core.BoolPtr(true)
				listUsersOptionsModel.Search = core.StringPtr("testString")
				listUsersOptionsModel.Start = core.StringPtr("testString")
				listUsersOptionsModel.UserID = core.StringPtr("testString")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := userManagementService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(usermanagementv1.UserList)
				responseObject.NextURL = core.StringPtr("ibm.com?_start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(usermanagementv1.UserList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(usermanagementv1.UserList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?_start=1","resources":[{"id":"ID","iam_id":"IamID","realm":"Realm","user_id":"UserID","firstname":"Firstname","lastname":"Lastname","state":"State","email":"Email","phonenumber":"Phonenumber","altphonenumber":"Altphonenumber","photo":"Photo","account_id":"AccountID","added_on":"AddedOn"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","iam_id":"IamID","realm":"Realm","user_id":"UserID","firstname":"Firstname","lastname":"Lastname","state":"State","email":"Email","phonenumber":"Phonenumber","altphonenumber":"Altphonenumber","photo":"Photo","account_id":"AccountID","added_on":"AddedOn"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use UsersPager.GetNext successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				listUsersOptionsModel := &usermanagementv1.ListUsersOptions{
					AccountID:       core.StringPtr("testString"),
					Limit:           core.Int64Ptr(int64(10)),
					IncludeSettings: core.BoolPtr(true),
					Search:          core.StringPtr("testString"),
					UserID:          core.StringPtr("testString"),
				}

				pager, err := userManagementService.NewUsersPager(listUsersOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []usermanagementv1.UserProfile
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use UsersPager.GetAll successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				listUsersOptionsModel := &usermanagementv1.ListUsersOptions{
					AccountID:       core.StringPtr("testString"),
					Limit:           core.Int64Ptr(int64(10)),
					IncludeSettings: core.BoolPtr(true),
					Search:          core.StringPtr("testString"),
					UserID:          core.StringPtr("testString"),
				}

				pager, err := userManagementService.NewUsersPager(listUsersOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`InviteUsers(inviteUsersOptions *InviteUsersOptions) - Operation response error`, func() {
		inviteUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(inviteUsersPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke InviteUsers with error: Operation response processing error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				roleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				inviteUserIamPolicyModel.Type = core.StringPtr("testString")
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}

				// Construct an instance of the InviteUsersOptions model
				inviteUsersOptionsModel := new(usermanagementv1.InviteUsersOptions)
				inviteUsersOptionsModel.AccountID = core.StringPtr("testString")
				inviteUsersOptionsModel.Users = []usermanagementv1.InviteUser{*inviteUserModel}
				inviteUsersOptionsModel.IamPolicy = []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}
				inviteUsersOptionsModel.AccessGroups = []string{"testString"}
				inviteUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userManagementService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				userManagementService.EnableRetries(0, 0)
				result, response, operationErr = userManagementService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`InviteUsers(inviteUsersOptions *InviteUsersOptions)`, func() {
		inviteUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(inviteUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"resources": [{"email": "Email", "id": "ID", "state": "State"}]}`)
				}))
			})
			It(`Invoke InviteUsers successfully with retries`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())
				userManagementService.EnableRetries(0, 0)

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				roleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				inviteUserIamPolicyModel.Type = core.StringPtr("testString")
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}

				// Construct an instance of the InviteUsersOptions model
				inviteUsersOptionsModel := new(usermanagementv1.InviteUsersOptions)
				inviteUsersOptionsModel.AccountID = core.StringPtr("testString")
				inviteUsersOptionsModel.Users = []usermanagementv1.InviteUser{*inviteUserModel}
				inviteUsersOptionsModel.IamPolicy = []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}
				inviteUsersOptionsModel.AccessGroups = []string{"testString"}
				inviteUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := userManagementService.InviteUsersWithContext(ctx, inviteUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				userManagementService.DisableRetries()
				result, response, operationErr := userManagementService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = userManagementService.InviteUsersWithContext(ctx, inviteUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(inviteUsersPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"resources": [{"email": "Email", "id": "ID", "state": "State"}]}`)
				}))
			})
			It(`Invoke InviteUsers successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userManagementService.InviteUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				roleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				inviteUserIamPolicyModel.Type = core.StringPtr("testString")
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}

				// Construct an instance of the InviteUsersOptions model
				inviteUsersOptionsModel := new(usermanagementv1.InviteUsersOptions)
				inviteUsersOptionsModel.AccountID = core.StringPtr("testString")
				inviteUsersOptionsModel.Users = []usermanagementv1.InviteUser{*inviteUserModel}
				inviteUsersOptionsModel.IamPolicy = []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}
				inviteUsersOptionsModel.AccessGroups = []string{"testString"}
				inviteUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userManagementService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke InviteUsers with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				roleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				inviteUserIamPolicyModel.Type = core.StringPtr("testString")
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}

				// Construct an instance of the InviteUsersOptions model
				inviteUsersOptionsModel := new(usermanagementv1.InviteUsersOptions)
				inviteUsersOptionsModel.AccountID = core.StringPtr("testString")
				inviteUsersOptionsModel.Users = []usermanagementv1.InviteUser{*inviteUserModel}
				inviteUsersOptionsModel.IamPolicy = []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}
				inviteUsersOptionsModel.AccessGroups = []string{"testString"}
				inviteUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userManagementService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the InviteUsersOptions model with no property values
				inviteUsersOptionsModelNew := new(usermanagementv1.InviteUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userManagementService.InviteUsers(inviteUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke InviteUsers successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				roleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				inviteUserIamPolicyModel.Type = core.StringPtr("testString")
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}

				// Construct an instance of the InviteUsersOptions model
				inviteUsersOptionsModel := new(usermanagementv1.InviteUsersOptions)
				inviteUsersOptionsModel.AccountID = core.StringPtr("testString")
				inviteUsersOptionsModel.Users = []usermanagementv1.InviteUser{*inviteUserModel}
				inviteUsersOptionsModel.IamPolicy = []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}
				inviteUsersOptionsModel.AccessGroups = []string{"testString"}
				inviteUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := userManagementService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUserProfile(getUserProfileOptions *GetUserProfileOptions) - Operation response error`, func() {
		getUserProfilePath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserProfilePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["include_activity"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUserProfile with error: Operation response processing error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserProfileOptions model
				getUserProfileOptionsModel := new(usermanagementv1.GetUserProfileOptions)
				getUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				getUserProfileOptionsModel.IamID = core.StringPtr("testString")
				getUserProfileOptionsModel.IncludeActivity = core.StringPtr("testString")
				getUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userManagementService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				userManagementService.EnableRetries(0, 0)
				result, response, operationErr = userManagementService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUserProfile(getUserProfileOptions *GetUserProfileOptions)`, func() {
		getUserProfilePath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["include_activity"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID", "added_on": "AddedOn"}`)
				}))
			})
			It(`Invoke GetUserProfile successfully with retries`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())
				userManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetUserProfileOptions model
				getUserProfileOptionsModel := new(usermanagementv1.GetUserProfileOptions)
				getUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				getUserProfileOptionsModel.IamID = core.StringPtr("testString")
				getUserProfileOptionsModel.IncludeActivity = core.StringPtr("testString")
				getUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := userManagementService.GetUserProfileWithContext(ctx, getUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				userManagementService.DisableRetries()
				result, response, operationErr := userManagementService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = userManagementService.GetUserProfileWithContext(ctx, getUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["include_activity"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID", "added_on": "AddedOn"}`)
				}))
			})
			It(`Invoke GetUserProfile successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userManagementService.GetUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserProfileOptions model
				getUserProfileOptionsModel := new(usermanagementv1.GetUserProfileOptions)
				getUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				getUserProfileOptionsModel.IamID = core.StringPtr("testString")
				getUserProfileOptionsModel.IncludeActivity = core.StringPtr("testString")
				getUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userManagementService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetUserProfile with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserProfileOptions model
				getUserProfileOptionsModel := new(usermanagementv1.GetUserProfileOptions)
				getUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				getUserProfileOptionsModel.IamID = core.StringPtr("testString")
				getUserProfileOptionsModel.IncludeActivity = core.StringPtr("testString")
				getUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userManagementService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserProfileOptions model with no property values
				getUserProfileOptionsModelNew := new(usermanagementv1.GetUserProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userManagementService.GetUserProfile(getUserProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetUserProfile successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserProfileOptions model
				getUserProfileOptionsModel := new(usermanagementv1.GetUserProfileOptions)
				getUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				getUserProfileOptionsModel.IamID = core.StringPtr("testString")
				getUserProfileOptionsModel.IncludeActivity = core.StringPtr("testString")
				getUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := userManagementService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateUserProfile(updateUserProfileOptions *UpdateUserProfileOptions)`, func() {
		updateUserProfilePath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserProfilePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["include_activity"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateUserProfile successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := userManagementService.UpdateUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateUserProfileOptions model
				updateUserProfileOptionsModel := new(usermanagementv1.UpdateUserProfileOptions)
				updateUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				updateUserProfileOptionsModel.IamID = core.StringPtr("testString")
				updateUserProfileOptionsModel.Firstname = core.StringPtr("testString")
				updateUserProfileOptionsModel.Lastname = core.StringPtr("testString")
				updateUserProfileOptionsModel.State = core.StringPtr("testString")
				updateUserProfileOptionsModel.Email = core.StringPtr("testString")
				updateUserProfileOptionsModel.Phonenumber = core.StringPtr("testString")
				updateUserProfileOptionsModel.Altphonenumber = core.StringPtr("testString")
				updateUserProfileOptionsModel.Photo = core.StringPtr("testString")
				updateUserProfileOptionsModel.IncludeActivity = core.StringPtr("testString")
				updateUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = userManagementService.UpdateUserProfile(updateUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateUserProfile with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserProfileOptions model
				updateUserProfileOptionsModel := new(usermanagementv1.UpdateUserProfileOptions)
				updateUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				updateUserProfileOptionsModel.IamID = core.StringPtr("testString")
				updateUserProfileOptionsModel.Firstname = core.StringPtr("testString")
				updateUserProfileOptionsModel.Lastname = core.StringPtr("testString")
				updateUserProfileOptionsModel.State = core.StringPtr("testString")
				updateUserProfileOptionsModel.Email = core.StringPtr("testString")
				updateUserProfileOptionsModel.Phonenumber = core.StringPtr("testString")
				updateUserProfileOptionsModel.Altphonenumber = core.StringPtr("testString")
				updateUserProfileOptionsModel.Photo = core.StringPtr("testString")
				updateUserProfileOptionsModel.IncludeActivity = core.StringPtr("testString")
				updateUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := userManagementService.UpdateUserProfile(updateUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateUserProfileOptions model with no property values
				updateUserProfileOptionsModelNew := new(usermanagementv1.UpdateUserProfileOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = userManagementService.UpdateUserProfile(updateUserProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveUser(removeUserOptions *RemoveUserOptions)`, func() {
		removeUserPath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["include_activity"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke RemoveUser successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := userManagementService.RemoveUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveUserOptions model
				removeUserOptionsModel := new(usermanagementv1.RemoveUserOptions)
				removeUserOptionsModel.AccountID = core.StringPtr("testString")
				removeUserOptionsModel.IamID = core.StringPtr("testString")
				removeUserOptionsModel.IncludeActivity = core.StringPtr("testString")
				removeUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = userManagementService.RemoveUser(removeUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RemoveUser with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the RemoveUserOptions model
				removeUserOptionsModel := new(usermanagementv1.RemoveUserOptions)
				removeUserOptionsModel.AccountID = core.StringPtr("testString")
				removeUserOptionsModel.IamID = core.StringPtr("testString")
				removeUserOptionsModel.IncludeActivity = core.StringPtr("testString")
				removeUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := userManagementService.RemoveUser(removeUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RemoveUserOptions model with no property values
				removeUserOptionsModelNew := new(usermanagementv1.RemoveUserOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = userManagementService.RemoveUser(removeUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Accept(acceptOptions *AcceptOptions)`, func() {
		acceptPath := "/v2/users/accept"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(acceptPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(202)
				}))
			})
			It(`Invoke Accept successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := userManagementService.Accept(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AcceptOptions model
				acceptOptionsModel := new(usermanagementv1.AcceptOptions)
				acceptOptionsModel.AccountID = core.StringPtr("testString")
				acceptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = userManagementService.Accept(acceptOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke Accept with error: Operation request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the AcceptOptions model
				acceptOptionsModel := new(usermanagementv1.AcceptOptions)
				acceptOptionsModel.AccountID = core.StringPtr("testString")
				acceptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := userManagementService.Accept(acceptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`V3RemoveUser(v3RemoveUserOptions *V3RemoveUserOptions)`, func() {
		v3RemoveUserPath := "/v3/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v3RemoveUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke V3RemoveUser successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := userManagementService.V3RemoveUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the V3RemoveUserOptions model
				v3RemoveUserOptionsModel := new(usermanagementv1.V3RemoveUserOptions)
				v3RemoveUserOptionsModel.AccountID = core.StringPtr("testString")
				v3RemoveUserOptionsModel.IamID = core.StringPtr("testString")
				v3RemoveUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = userManagementService.V3RemoveUser(v3RemoveUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke V3RemoveUser with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the V3RemoveUserOptions model
				v3RemoveUserOptionsModel := new(usermanagementv1.V3RemoveUserOptions)
				v3RemoveUserOptionsModel.AccountID = core.StringPtr("testString")
				v3RemoveUserOptionsModel.IamID = core.StringPtr("testString")
				v3RemoveUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := userManagementService.V3RemoveUser(v3RemoveUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the V3RemoveUserOptions model with no property values
				v3RemoveUserOptionsModelNew := new(usermanagementv1.V3RemoveUserOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = userManagementService.V3RemoveUser(v3RemoveUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions) - Operation response error`, func() {
		getUserSettingsPath := "/v2/accounts/testString/users/testString/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUserSettings with error: Operation response processing error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserSettingsOptions model
				getUserSettingsOptionsModel := new(usermanagementv1.GetUserSettingsOptions)
				getUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				getUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userManagementService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				userManagementService.EnableRetries(0, 0)
				result, response, operationErr = userManagementService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions)`, func() {
		getUserSettingsPath := "/v2/accounts/testString/users/testString/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"language": "Language", "notification_language": "NotificationLanguage", "allowed_ip_addresses": "32.96.110.50,172.16.254.1", "self_manage": true}`)
				}))
			})
			It(`Invoke GetUserSettings successfully with retries`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())
				userManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetUserSettingsOptions model
				getUserSettingsOptionsModel := new(usermanagementv1.GetUserSettingsOptions)
				getUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				getUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := userManagementService.GetUserSettingsWithContext(ctx, getUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				userManagementService.DisableRetries()
				result, response, operationErr := userManagementService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = userManagementService.GetUserSettingsWithContext(ctx, getUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"language": "Language", "notification_language": "NotificationLanguage", "allowed_ip_addresses": "32.96.110.50,172.16.254.1", "self_manage": true}`)
				}))
			})
			It(`Invoke GetUserSettings successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userManagementService.GetUserSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserSettingsOptions model
				getUserSettingsOptionsModel := new(usermanagementv1.GetUserSettingsOptions)
				getUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				getUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userManagementService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetUserSettings with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserSettingsOptions model
				getUserSettingsOptionsModel := new(usermanagementv1.GetUserSettingsOptions)
				getUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				getUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userManagementService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserSettingsOptions model with no property values
				getUserSettingsOptionsModelNew := new(usermanagementv1.GetUserSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userManagementService.GetUserSettings(getUserSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetUserSettings successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserSettingsOptions model
				getUserSettingsOptionsModel := new(usermanagementv1.GetUserSettingsOptions)
				getUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				getUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := userManagementService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateUserSettings(updateUserSettingsOptions *UpdateUserSettingsOptions)`, func() {
		updateUserSettingsPath := "/v2/accounts/testString/users/testString/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateUserSettings successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := userManagementService.UpdateUserSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateUserSettingsOptions model
				updateUserSettingsOptionsModel := new(usermanagementv1.UpdateUserSettingsOptions)
				updateUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.Language = core.StringPtr("testString")
				updateUserSettingsOptionsModel.NotificationLanguage = core.StringPtr("testString")
				updateUserSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("32.96.110.50,172.16.254.1")
				updateUserSettingsOptionsModel.SelfManage = core.BoolPtr(true)
				updateUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = userManagementService.UpdateUserSettings(updateUserSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateUserSettings with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserSettingsOptions model
				updateUserSettingsOptionsModel := new(usermanagementv1.UpdateUserSettingsOptions)
				updateUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.Language = core.StringPtr("testString")
				updateUserSettingsOptionsModel.NotificationLanguage = core.StringPtr("testString")
				updateUserSettingsOptionsModel.AllowedIPAddresses = core.StringPtr("32.96.110.50,172.16.254.1")
				updateUserSettingsOptionsModel.SelfManage = core.BoolPtr(true)
				updateUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := userManagementService.UpdateUserSettings(updateUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateUserSettingsOptions model with no property values
				updateUserSettingsOptionsModelNew := new(usermanagementv1.UpdateUserSettingsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = userManagementService.UpdateUserSettings(updateUserSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			userManagementService, _ := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				URL:           "http://usermanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAcceptOptions successfully`, func() {
				// Construct an instance of the AcceptOptions model
				acceptOptionsModel := userManagementService.NewAcceptOptions()
				acceptOptionsModel.SetAccountID("testString")
				acceptOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(acceptOptionsModel).ToNot(BeNil())
				Expect(acceptOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(acceptOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserProfileOptions successfully`, func() {
				// Construct an instance of the GetUserProfileOptions model
				accountID := "testString"
				iamID := "testString"
				getUserProfileOptionsModel := userManagementService.NewGetUserProfileOptions(accountID, iamID)
				getUserProfileOptionsModel.SetAccountID("testString")
				getUserProfileOptionsModel.SetIamID("testString")
				getUserProfileOptionsModel.SetIncludeActivity("testString")
				getUserProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserProfileOptionsModel).ToNot(BeNil())
				Expect(getUserProfileOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getUserProfileOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(getUserProfileOptionsModel.IncludeActivity).To(Equal(core.StringPtr("testString")))
				Expect(getUserProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserSettingsOptions successfully`, func() {
				// Construct an instance of the GetUserSettingsOptions model
				accountID := "testString"
				iamID := "testString"
				getUserSettingsOptionsModel := userManagementService.NewGetUserSettingsOptions(accountID, iamID)
				getUserSettingsOptionsModel.SetAccountID("testString")
				getUserSettingsOptionsModel.SetIamID("testString")
				getUserSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserSettingsOptionsModel).ToNot(BeNil())
				Expect(getUserSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getUserSettingsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(getUserSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInviteUsersOptions successfully`, func() {
				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				Expect(inviteUserModel).ToNot(BeNil())
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")
				Expect(inviteUserModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(inviteUserModel.AccountRole).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				Expect(roleModel).ToNot(BeNil())
				roleModel.RoleID = core.StringPtr("testString")
				Expect(roleModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				Expect(attributeModel).ToNot(BeNil())
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")
				Expect(attributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(attributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				Expect(resourceModel).ToNot(BeNil())
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}
				Expect(resourceModel.Attributes).To(Equal([]usermanagementv1.Attribute{*attributeModel}))

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				Expect(inviteUserIamPolicyModel).ToNot(BeNil())
				inviteUserIamPolicyModel.Type = core.StringPtr("testString")
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}
				Expect(inviteUserIamPolicyModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(inviteUserIamPolicyModel.Roles).To(Equal([]usermanagementv1.Role{*roleModel}))
				Expect(inviteUserIamPolicyModel.Resources).To(Equal([]usermanagementv1.Resource{*resourceModel}))

				// Construct an instance of the InviteUsersOptions model
				accountID := "testString"
				inviteUsersOptionsModel := userManagementService.NewInviteUsersOptions(accountID)
				inviteUsersOptionsModel.SetAccountID("testString")
				inviteUsersOptionsModel.SetUsers([]usermanagementv1.InviteUser{*inviteUserModel})
				inviteUsersOptionsModel.SetIamPolicy([]usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel})
				inviteUsersOptionsModel.SetAccessGroups([]string{"testString"})
				inviteUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(inviteUsersOptionsModel).ToNot(BeNil())
				Expect(inviteUsersOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(inviteUsersOptionsModel.Users).To(Equal([]usermanagementv1.InviteUser{*inviteUserModel}))
				Expect(inviteUsersOptionsModel.IamPolicy).To(Equal([]usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}))
				Expect(inviteUsersOptionsModel.AccessGroups).To(Equal([]string{"testString"}))
				Expect(inviteUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListUsersOptions successfully`, func() {
				// Construct an instance of the ListUsersOptions model
				accountID := "testString"
				listUsersOptionsModel := userManagementService.NewListUsersOptions(accountID)
				listUsersOptionsModel.SetAccountID("testString")
				listUsersOptionsModel.SetLimit(int64(10))
				listUsersOptionsModel.SetIncludeSettings(true)
				listUsersOptionsModel.SetSearch("testString")
				listUsersOptionsModel.SetStart("testString")
				listUsersOptionsModel.SetUserID("testString")
				listUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listUsersOptionsModel).ToNot(BeNil())
				Expect(listUsersOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listUsersOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listUsersOptionsModel.IncludeSettings).To(Equal(core.BoolPtr(true)))
				Expect(listUsersOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listUsersOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listUsersOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(listUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveUserOptions successfully`, func() {
				// Construct an instance of the RemoveUserOptions model
				accountID := "testString"
				iamID := "testString"
				removeUserOptionsModel := userManagementService.NewRemoveUserOptions(accountID, iamID)
				removeUserOptionsModel.SetAccountID("testString")
				removeUserOptionsModel.SetIamID("testString")
				removeUserOptionsModel.SetIncludeActivity("testString")
				removeUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeUserOptionsModel).ToNot(BeNil())
				Expect(removeUserOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(removeUserOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(removeUserOptionsModel.IncludeActivity).To(Equal(core.StringPtr("testString")))
				Expect(removeUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserProfileOptions successfully`, func() {
				// Construct an instance of the UpdateUserProfileOptions model
				accountID := "testString"
				iamID := "testString"
				updateUserProfileOptionsModel := userManagementService.NewUpdateUserProfileOptions(accountID, iamID)
				updateUserProfileOptionsModel.SetAccountID("testString")
				updateUserProfileOptionsModel.SetIamID("testString")
				updateUserProfileOptionsModel.SetFirstname("testString")
				updateUserProfileOptionsModel.SetLastname("testString")
				updateUserProfileOptionsModel.SetState("testString")
				updateUserProfileOptionsModel.SetEmail("testString")
				updateUserProfileOptionsModel.SetPhonenumber("testString")
				updateUserProfileOptionsModel.SetAltphonenumber("testString")
				updateUserProfileOptionsModel.SetPhoto("testString")
				updateUserProfileOptionsModel.SetIncludeActivity("testString")
				updateUserProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateUserProfileOptionsModel).ToNot(BeNil())
				Expect(updateUserProfileOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.Firstname).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.Lastname).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.State).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.Phonenumber).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.Altphonenumber).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.Photo).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.IncludeActivity).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserSettingsOptions successfully`, func() {
				// Construct an instance of the UpdateUserSettingsOptions model
				accountID := "testString"
				iamID := "testString"
				updateUserSettingsOptionsModel := userManagementService.NewUpdateUserSettingsOptions(accountID, iamID)
				updateUserSettingsOptionsModel.SetAccountID("testString")
				updateUserSettingsOptionsModel.SetIamID("testString")
				updateUserSettingsOptionsModel.SetLanguage("testString")
				updateUserSettingsOptionsModel.SetNotificationLanguage("testString")
				updateUserSettingsOptionsModel.SetAllowedIPAddresses("32.96.110.50,172.16.254.1")
				updateUserSettingsOptionsModel.SetSelfManage(true)
				updateUserSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateUserSettingsOptionsModel).ToNot(BeNil())
				Expect(updateUserSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserSettingsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserSettingsOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(updateUserSettingsOptionsModel.NotificationLanguage).To(Equal(core.StringPtr("testString")))
				Expect(updateUserSettingsOptionsModel.AllowedIPAddresses).To(Equal(core.StringPtr("32.96.110.50,172.16.254.1")))
				Expect(updateUserSettingsOptionsModel.SelfManage).To(Equal(core.BoolPtr(true)))
				Expect(updateUserSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewV3RemoveUserOptions successfully`, func() {
				// Construct an instance of the V3RemoveUserOptions model
				accountID := "testString"
				iamID := "testString"
				v3RemoveUserOptionsModel := userManagementService.NewV3RemoveUserOptions(accountID, iamID)
				v3RemoveUserOptionsModel.SetAccountID("testString")
				v3RemoveUserOptionsModel.SetIamID("testString")
				v3RemoveUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(v3RemoveUserOptionsModel).ToNot(BeNil())
				Expect(v3RemoveUserOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(v3RemoveUserOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(v3RemoveUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInviteUserIamPolicy successfully`, func() {
				typeVar := "testString"
				_model, err := userManagementService.NewInviteUserIamPolicy(typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
