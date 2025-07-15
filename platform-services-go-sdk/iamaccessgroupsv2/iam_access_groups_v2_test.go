/**
 * (C) Copyright IBM Corp. 2025.
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

package iamaccessgroupsv2_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IamAccessGroupsV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(iamAccessGroupsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "https://iamaccessgroupsv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL":       "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := iamAccessGroupsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamAccessGroupsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamAccessGroupsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamAccessGroupsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL: "https://testService/api",
				})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := iamAccessGroupsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamAccessGroupsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamAccessGroupsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamAccessGroupsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{})
				err := iamAccessGroupsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := iamAccessGroupsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamAccessGroupsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamAccessGroupsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamAccessGroupsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL":       "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = iamaccessgroupsv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateAccessGroup(createAccessGroupOptions *CreateAccessGroupOptions) - Operation response error`, func() {
		createAccessGroupPath := "/v2/groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessGroupPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccessGroup(createAccessGroupOptions *CreateAccessGroupOptions)`, func() {
		createAccessGroupPath := "/v2/groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessGroupPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false, "crn": "CRN"}`)
				}))
			})
			It(`Invoke CreateAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.CreateAccessGroupWithContext(ctx, createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.CreateAccessGroupWithContext(ctx, createAccessGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccessGroupPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false, "crn": "CRN"}`)
				}))
			})
			It(`Invoke CreateAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccessGroupOptions model with no property values
				createAccessGroupOptionsModelNew := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
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
	Describe(`ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions) - Operation response error`, func() {
		listAccessGroupsPath := "/v2/groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// TODO: Add check for show_federated query parameter
					// TODO: Add check for hide_public_access query parameter
					// TODO: Add check for show_crn query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccessGroups with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupsOptions model
				listAccessGroupsOptionsModel := new(iamaccessgroupsv2.ListAccessGroupsOptions)
				listAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.Search = core.StringPtr("testString")
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
				listAccessGroupsOptionsModel.ShowCRN = core.BoolPtr(false)
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions)`, func() {
		listAccessGroupsPath := "/v2/groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// TODO: Add check for show_federated query parameter
					// TODO: Add check for hide_public_access query parameter
					// TODO: Add check for show_crn query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "groups": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false, "crn": "CRN"}]}`)
				}))
			})
			It(`Invoke ListAccessGroups successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the ListAccessGroupsOptions model
				listAccessGroupsOptionsModel := new(iamaccessgroupsv2.ListAccessGroupsOptions)
				listAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.Search = core.StringPtr("testString")
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
				listAccessGroupsOptionsModel.ShowCRN = core.BoolPtr(false)
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ListAccessGroupsWithContext(ctx, listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupsWithContext(ctx, listAccessGroupsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// TODO: Add check for show_federated query parameter
					// TODO: Add check for hide_public_access query parameter
					// TODO: Add check for show_crn query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "groups": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false, "crn": "CRN"}]}`)
				}))
			})
			It(`Invoke ListAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessGroupsOptions model
				listAccessGroupsOptionsModel := new(iamaccessgroupsv2.ListAccessGroupsOptions)
				listAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.Search = core.StringPtr("testString")
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
				listAccessGroupsOptionsModel.ShowCRN = core.BoolPtr(false)
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccessGroups with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupsOptions model
				listAccessGroupsOptionsModel := new(iamaccessgroupsv2.ListAccessGroupsOptions)
				listAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.Search = core.StringPtr("testString")
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
				listAccessGroupsOptionsModel.ShowCRN = core.BoolPtr(false)
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAccessGroupsOptions model with no property values
				listAccessGroupsOptionsModelNew := new(iamaccessgroupsv2.ListAccessGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModelNew)
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
			It(`Invoke ListAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupsOptions model
				listAccessGroupsOptionsModel := new(iamaccessgroupsv2.ListAccessGroupsOptions)
				listAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.Search = core.StringPtr("testString")
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
				listAccessGroupsOptionsModel.ShowCRN = core.BoolPtr(false)
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
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
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(iamaccessgroupsv2.GroupsList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(iamaccessgroupsv2.GroupsList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.GroupsList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.GroupsList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"groups":[{"id":"ID","name":"Name","description":"Description","account_id":"AccountID","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID","last_modified_at":"2019-01-01T12:00:00.000Z","last_modified_by_id":"LastModifiedByID","href":"Href","is_federated":false,"crn":"CRN"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"groups":[{"id":"ID","name":"Name","description":"Description","account_id":"AccountID","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID","last_modified_at":"2019-01-01T12:00:00.000Z","last_modified_by_id":"LastModifiedByID","href":"Href","is_federated":false,"crn":"CRN"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AccessGroupsPager.GetNext successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listAccessGroupsOptionsModel := &iamaccessgroupsv2.ListAccessGroupsOptions{
					AccountID:        core.StringPtr("testString"),
					TransactionID:    core.StringPtr("testString"),
					IamID:            core.StringPtr("testString"),
					Search:           core.StringPtr("testString"),
					MembershipType:   core.StringPtr("static"),
					Limit:            core.Int64Ptr(int64(10)),
					Sort:             core.StringPtr("name"),
					ShowFederated:    core.BoolPtr(false),
					HidePublicAccess: core.BoolPtr(false),
					ShowCRN:          core.BoolPtr(false),
				}

				pager, err := iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []iamaccessgroupsv2.Group
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AccessGroupsPager.GetAll successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listAccessGroupsOptionsModel := &iamaccessgroupsv2.ListAccessGroupsOptions{
					AccountID:        core.StringPtr("testString"),
					TransactionID:    core.StringPtr("testString"),
					IamID:            core.StringPtr("testString"),
					Search:           core.StringPtr("testString"),
					MembershipType:   core.StringPtr("static"),
					Limit:            core.Int64Ptr(int64(10)),
					Sort:             core.StringPtr("name"),
					ShowFederated:    core.BoolPtr(false),
					HidePublicAccess: core.BoolPtr(false),
					ShowCRN:          core.BoolPtr(false),
				}

				pager, err := iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetAccessGroup(getAccessGroupOptions *GetAccessGroupOptions) - Operation response error`, func() {
		getAccessGroupPath := "/v2/groups/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for show_federated query parameter
					// TODO: Add check for show_crn query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
				getAccessGroupOptionsModel.ShowCRN = core.BoolPtr(false)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccessGroup(getAccessGroupOptions *GetAccessGroupOptions)`, func() {
		getAccessGroupPath := "/v2/groups/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for show_federated query parameter
					// TODO: Add check for show_crn query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false, "crn": "CRN"}`)
				}))
			})
			It(`Invoke GetAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
				getAccessGroupOptionsModel.ShowCRN = core.BoolPtr(false)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.GetAccessGroupWithContext(ctx, getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.GetAccessGroupWithContext(ctx, getAccessGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for show_federated query parameter
					// TODO: Add check for show_crn query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false, "crn": "CRN"}`)
				}))
			})
			It(`Invoke GetAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
				getAccessGroupOptionsModel.ShowCRN = core.BoolPtr(false)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
				getAccessGroupOptionsModel.ShowCRN = core.BoolPtr(false)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccessGroupOptions model with no property values
				getAccessGroupOptionsModelNew := new(iamaccessgroupsv2.GetAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModelNew)
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
			It(`Invoke GetAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
				getAccessGroupOptionsModel.ShowCRN = core.BoolPtr(false)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
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
	Describe(`UpdateAccessGroup(updateAccessGroupOptions *UpdateAccessGroupOptions) - Operation response error`, func() {
		updateAccessGroupPath := "/v2/groups/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccessGroup(updateAccessGroupOptions *UpdateAccessGroupOptions)`, func() {
		updateAccessGroupPath := "/v2/groups/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessGroupPath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false, "crn": "CRN"}`)
				}))
			})
			It(`Invoke UpdateAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.UpdateAccessGroupWithContext(ctx, updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.UpdateAccessGroupWithContext(ctx, updateAccessGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessGroupPath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false, "crn": "CRN"}`)
				}))
			})
			It(`Invoke UpdateAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccessGroupOptions model with no property values
				updateAccessGroupOptionsModelNew := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModelNew)
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
			It(`Invoke UpdateAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
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
	Describe(`DeleteAccessGroup(deleteAccessGroupOptions *DeleteAccessGroupOptions)`, func() {
		deleteAccessGroupPath := "/v2/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessGroupPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for force query parameter
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.DeleteAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAccessGroupOptions model
				deleteAccessGroupOptionsModel := new(iamaccessgroupsv2.DeleteAccessGroupOptions)
				deleteAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.Force = core.BoolPtr(false)
				deleteAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessGroupOptions model
				deleteAccessGroupOptionsModel := new(iamaccessgroupsv2.DeleteAccessGroupOptions)
				deleteAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.Force = core.BoolPtr(false)
				deleteAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAccessGroupOptions model with no property values
				deleteAccessGroupOptionsModelNew := new(iamaccessgroupsv2.DeleteAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`IsMemberOfAccessGroup(isMemberOfAccessGroupOptions *IsMemberOfAccessGroupOptions)`, func() {
		isMemberOfAccessGroupPath := "/v2/groups/testString/members/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(isMemberOfAccessGroupPath))
					Expect(req.Method).To(Equal("HEAD"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke IsMemberOfAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.IsMemberOfAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the IsMemberOfAccessGroupOptions model
				isMemberOfAccessGroupOptionsModel := new(iamaccessgroupsv2.IsMemberOfAccessGroupOptions)
				isMemberOfAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke IsMemberOfAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the IsMemberOfAccessGroupOptions model
				isMemberOfAccessGroupOptionsModel := new(iamaccessgroupsv2.IsMemberOfAccessGroupOptions)
				isMemberOfAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the IsMemberOfAccessGroupOptions model with no property values
				isMemberOfAccessGroupOptionsModelNew := new(iamaccessgroupsv2.IsMemberOfAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddMembersToAccessGroup(addMembersToAccessGroupOptions *AddMembersToAccessGroupOptions) - Operation response error`, func() {
		addMembersToAccessGroupPath := "/v2/groups/testString/members"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMembersToAccessGroupPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddMembersToAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddMembersToAccessGroup(addMembersToAccessGroupOptions *AddMembersToAccessGroupOptions)`, func() {
		addMembersToAccessGroupPath := "/v2/groups/testString/members"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMembersToAccessGroupPath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"members": [{"iam_id": "IamID", "type": "Type", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke AddMembersToAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.AddMembersToAccessGroupWithContext(ctx, addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.AddMembersToAccessGroupWithContext(ctx, addMembersToAccessGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addMembersToAccessGroupPath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"members": [{"iam_id": "IamID", "type": "Type", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke AddMembersToAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddMembersToAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddMembersToAccessGroupOptions model with no property values
				addMembersToAccessGroupOptionsModelNew := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModelNew)
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
					res.WriteHeader(207)
				}))
			})
			It(`Invoke AddMembersToAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
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
	Describe(`ListAccessGroupMembers(listAccessGroupMembersOptions *ListAccessGroupMembersOptions) - Operation response error`, func() {
		listAccessGroupMembersPath := "/v2/groups/testString/members"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// TODO: Add check for verbose query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccessGroupMembers with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessGroupMembers(listAccessGroupMembersOptions *ListAccessGroupMembersOptions)`, func() {
		listAccessGroupMembersPath := "/v2/groups/testString/members"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// TODO: Add check for verbose query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "members": [{"iam_id": "IamID", "type": "Type", "membership_type": "MembershipType", "name": "Name", "email": "Email", "description": "Description", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupMembers successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ListAccessGroupMembersWithContext(ctx, listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupMembersWithContext(ctx, listAccessGroupMembersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// TODO: Add check for verbose query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "members": [{"iam_id": "IamID", "type": "Type", "membership_type": "MembershipType", "name": "Name", "email": "Email", "description": "Description", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupMembers successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccessGroupMembers with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAccessGroupMembersOptions model with no property values
				listAccessGroupMembersOptionsModelNew := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModelNew)
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
			It(`Invoke ListAccessGroupMembers successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
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
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(iamaccessgroupsv2.GroupMembersList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(iamaccessgroupsv2.GroupMembersList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.GroupMembersList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.GroupMembersList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"members":[{"iam_id":"IamID","type":"Type","membership_type":"MembershipType","name":"Name","email":"Email","description":"Description","href":"Href","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"members":[{"iam_id":"IamID","type":"Type","membership_type":"MembershipType","name":"Name","email":"Email","description":"Description","href":"Href","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID"}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AccessGroupMembersPager.GetNext successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listAccessGroupMembersOptionsModel := &iamaccessgroupsv2.ListAccessGroupMembersOptions{
					AccessGroupID:  core.StringPtr("testString"),
					TransactionID:  core.StringPtr("testString"),
					MembershipType: core.StringPtr("static"),
					Limit:          core.Int64Ptr(int64(10)),
					Type:           core.StringPtr("testString"),
					Verbose:        core.BoolPtr(false),
					Sort:           core.StringPtr("testString"),
				}

				pager, err := iamAccessGroupsService.NewAccessGroupMembersPager(listAccessGroupMembersOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []iamaccessgroupsv2.ListGroupMembersResponseMember
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AccessGroupMembersPager.GetAll successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listAccessGroupMembersOptionsModel := &iamaccessgroupsv2.ListAccessGroupMembersOptions{
					AccessGroupID:  core.StringPtr("testString"),
					TransactionID:  core.StringPtr("testString"),
					MembershipType: core.StringPtr("static"),
					Limit:          core.Int64Ptr(int64(10)),
					Type:           core.StringPtr("testString"),
					Verbose:        core.BoolPtr(false),
					Sort:           core.StringPtr("testString"),
				}

				pager, err := iamAccessGroupsService.NewAccessGroupMembersPager(listAccessGroupMembersOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions *RemoveMemberFromAccessGroupOptions)`, func() {
		removeMemberFromAccessGroupPath := "/v2/groups/testString/members/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMemberFromAccessGroupPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke RemoveMemberFromAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.RemoveMemberFromAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveMemberFromAccessGroupOptions model
				removeMemberFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAccessGroupOptions)
				removeMemberFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RemoveMemberFromAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMemberFromAccessGroupOptions model
				removeMemberFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAccessGroupOptions)
				removeMemberFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RemoveMemberFromAccessGroupOptions model with no property values
				removeMemberFromAccessGroupOptionsModelNew := new(iamaccessgroupsv2.RemoveMemberFromAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptions *RemoveMembersFromAccessGroupOptions) - Operation response error`, func() {
		removeMembersFromAccessGroupPath := "/v2/groups/testString/members/delete"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMembersFromAccessGroupPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveMembersFromAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptions *RemoveMembersFromAccessGroupOptions)`, func() {
		removeMembersFromAccessGroupPath := "/v2/groups/testString/members/delete"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMembersFromAccessGroupPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"access_group_id": "AccessGroupID", "members": [{"iam_id": "IamID", "trace": "Trace", "status_code": 10, "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke RemoveMembersFromAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroupWithContext(ctx, removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroupWithContext(ctx, removeMembersFromAccessGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(removeMembersFromAccessGroupPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"access_group_id": "AccessGroupID", "members": [{"iam_id": "IamID", "trace": "Trace", "status_code": 10, "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke RemoveMembersFromAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RemoveMembersFromAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveMembersFromAccessGroupOptions model with no property values
				removeMembersFromAccessGroupOptionsModelNew := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModelNew)
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
					res.WriteHeader(207)
				}))
			})
			It(`Invoke RemoveMembersFromAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
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
	Describe(`RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions *RemoveMemberFromAllAccessGroupsOptions) - Operation response error`, func() {
		removeMemberFromAllAccessGroupsPath := "/v2/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMemberFromAllAccessGroupsPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveMemberFromAllAccessGroups with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions *RemoveMemberFromAllAccessGroupsOptions)`, func() {
		removeMemberFromAllAccessGroupsPath := "/v2/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMemberFromAllAccessGroupsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke RemoveMemberFromAllAccessGroups successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroupsWithContext(ctx, removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroupsWithContext(ctx, removeMemberFromAllAccessGroupsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(removeMemberFromAllAccessGroupsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke RemoveMemberFromAllAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RemoveMemberFromAllAccessGroups with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveMemberFromAllAccessGroupsOptions model with no property values
				removeMemberFromAllAccessGroupsOptionsModelNew := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModelNew)
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
					res.WriteHeader(207)
				}))
			})
			It(`Invoke RemoveMemberFromAllAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
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
	Describe(`AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions *AddMemberToMultipleAccessGroupsOptions) - Operation response error`, func() {
		addMemberToMultipleAccessGroupsPath := "/v2/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMemberToMultipleAccessGroupsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddMemberToMultipleAccessGroups with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"AccessGroupId-b0d32f56-f85c-4bf1-af37-7bbd92b1b2b3"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions *AddMemberToMultipleAccessGroupsOptions)`, func() {
		addMemberToMultipleAccessGroupsPath := "/v2/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMemberToMultipleAccessGroupsPath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke AddMemberToMultipleAccessGroups successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"AccessGroupId-b0d32f56-f85c-4bf1-af37-7bbd92b1b2b3"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroupsWithContext(ctx, addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroupsWithContext(ctx, addMemberToMultipleAccessGroupsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addMemberToMultipleAccessGroupsPath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke AddMemberToMultipleAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"AccessGroupId-b0d32f56-f85c-4bf1-af37-7bbd92b1b2b3"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddMemberToMultipleAccessGroups with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"AccessGroupId-b0d32f56-f85c-4bf1-af37-7bbd92b1b2b3"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddMemberToMultipleAccessGroupsOptions model with no property values
				addMemberToMultipleAccessGroupsOptionsModelNew := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModelNew)
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
					res.WriteHeader(207)
				}))
			})
			It(`Invoke AddMemberToMultipleAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"AccessGroupId-b0d32f56-f85c-4bf1-af37-7bbd92b1b2b3"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
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
	Describe(`AddAccessGroupRule(addAccessGroupRuleOptions *AddAccessGroupRuleOptions) - Operation response error`, func() {
		addAccessGroupRulePath := "/v2/groups/testString/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addAccessGroupRulePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddAccessGroupRule with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				addAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddAccessGroupRule(addAccessGroupRuleOptions *AddAccessGroupRuleOptions)`, func() {
		addAccessGroupRulePath := "/v2/groups/testString/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addAccessGroupRulePath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke AddAccessGroupRule successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				addAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.AddAccessGroupRuleWithContext(ctx, addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.AddAccessGroupRuleWithContext(ctx, addAccessGroupRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addAccessGroupRulePath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke AddAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				addAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddAccessGroupRule with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				addAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddAccessGroupRuleOptions model with no property values
				addAccessGroupRuleOptionsModelNew := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				addAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
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
	Describe(`ListAccessGroupRules(listAccessGroupRulesOptions *ListAccessGroupRulesOptions) - Operation response error`, func() {
		listAccessGroupRulesPath := "/v2/groups/testString/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccessGroupRules with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessGroupRules(listAccessGroupRulesOptions *ListAccessGroupRulesOptions)`, func() {
		listAccessGroupRulesPath := "/v2/groups/testString/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rules": [{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupRules successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ListAccessGroupRulesWithContext(ctx, listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupRulesWithContext(ctx, listAccessGroupRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rules": [{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupRules successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccessGroupRules with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAccessGroupRulesOptions model with no property values
				listAccessGroupRulesOptionsModelNew := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModelNew)
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
			It(`Invoke ListAccessGroupRules successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
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
	Describe(`GetAccessGroupRule(getAccessGroupRuleOptions *GetAccessGroupRuleOptions) - Operation response error`, func() {
		getAccessGroupRulePath := "/v2/groups/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccessGroupRule with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccessGroupRule(getAccessGroupRuleOptions *GetAccessGroupRuleOptions)`, func() {
		getAccessGroupRulePath := "/v2/groups/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAccessGroupRule successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.GetAccessGroupRuleWithContext(ctx, getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.GetAccessGroupRuleWithContext(ctx, getAccessGroupRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccessGroupRule with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccessGroupRuleOptions model with no property values
				getAccessGroupRuleOptionsModelNew := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModelNew)
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
			It(`Invoke GetAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
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
	Describe(`ReplaceAccessGroupRule(replaceAccessGroupRuleOptions *ReplaceAccessGroupRuleOptions) - Operation response error`, func() {
		replaceAccessGroupRulePath := "/v2/groups/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAccessGroupRulePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceAccessGroupRule with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				replaceAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceAccessGroupRule(replaceAccessGroupRuleOptions *ReplaceAccessGroupRuleOptions)`, func() {
		replaceAccessGroupRulePath := "/v2/groups/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAccessGroupRulePath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplaceAccessGroupRule successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				replaceAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ReplaceAccessGroupRuleWithContext(ctx, replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ReplaceAccessGroupRuleWithContext(ctx, replaceAccessGroupRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceAccessGroupRulePath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplaceAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				replaceAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceAccessGroupRule with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				replaceAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceAccessGroupRuleOptions model with no property values
				replaceAccessGroupRuleOptionsModelNew := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModelNew)
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
			It(`Invoke ReplaceAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				replaceAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
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
	Describe(`RemoveAccessGroupRule(removeAccessGroupRuleOptions *RemoveAccessGroupRuleOptions)`, func() {
		removeAccessGroupRulePath := "/v2/groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeAccessGroupRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke RemoveAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.RemoveAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveAccessGroupRuleOptions model
				removeAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.RemoveAccessGroupRuleOptions)
				removeAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RemoveAccessGroupRule with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveAccessGroupRuleOptions model
				removeAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.RemoveAccessGroupRuleOptions)
				removeAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RemoveAccessGroupRuleOptions model with no property values
				removeAccessGroupRuleOptionsModelNew := new(iamaccessgroupsv2.RemoveAccessGroupRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) - Operation response error`, func() {
		getAccountSettingsPath := "/v2/groups/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountSettings with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions)`, func() {
		getAccountSettingsPath := "/v2/groups/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountSettings with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountSettingsOptions model with no property values
				getAccountSettingsOptionsModelNew := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModelNew)
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
			It(`Invoke GetAccountSettings successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
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
	Describe(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions) - Operation response error`, func() {
		updateAccountSettingsPath := "/v2/groups/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountSettings with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions)`, func() {
		updateAccountSettingsPath := "/v2/groups/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccountSettings with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountSettingsOptions model with no property values
				updateAccountSettingsOptionsModelNew := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModelNew)
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
			It(`Invoke UpdateAccountSettings successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
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
	Describe(`CreateTemplate(createTemplateOptions *CreateTemplateOptions) - Operation response error`, func() {
		createTemplatePath := "/v1/group_templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTemplate with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345", "iam-ServiceId-456"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(iamaccessgroupsv2.CreateTemplateOptions)
				createTemplateOptionsModel.Name = core.StringPtr("IAM Admin Group template")
				createTemplateOptionsModel.AccountID = core.StringPtr("accountID-123")
				createTemplateOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateOptionsModel.Group = accessGroupRequestModel
				createTemplateOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTemplate(createTemplateOptions *CreateTemplateOptions)`, func() {
		createTemplatePath := "/v1/group_templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTemplatePath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateTemplate successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345", "iam-ServiceId-456"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(iamaccessgroupsv2.CreateTemplateOptions)
				createTemplateOptionsModel.Name = core.StringPtr("IAM Admin Group template")
				createTemplateOptionsModel.AccountID = core.StringPtr("accountID-123")
				createTemplateOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateOptionsModel.Group = accessGroupRequestModel
				createTemplateOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.CreateTemplateWithContext(ctx, createTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.CreateTemplateWithContext(ctx, createTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTemplatePath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateTemplate successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.CreateTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345", "iam-ServiceId-456"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(iamaccessgroupsv2.CreateTemplateOptions)
				createTemplateOptionsModel.Name = core.StringPtr("IAM Admin Group template")
				createTemplateOptionsModel.AccountID = core.StringPtr("accountID-123")
				createTemplateOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateOptionsModel.Group = accessGroupRequestModel
				createTemplateOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTemplate with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345", "iam-ServiceId-456"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(iamaccessgroupsv2.CreateTemplateOptions)
				createTemplateOptionsModel.Name = core.StringPtr("IAM Admin Group template")
				createTemplateOptionsModel.AccountID = core.StringPtr("accountID-123")
				createTemplateOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateOptionsModel.Group = accessGroupRequestModel
				createTemplateOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTemplateOptions model with no property values
				createTemplateOptionsModelNew := new(iamaccessgroupsv2.CreateTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.CreateTemplate(createTemplateOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateTemplate successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345", "iam-ServiceId-456"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(iamaccessgroupsv2.CreateTemplateOptions)
				createTemplateOptionsModel.Name = core.StringPtr("IAM Admin Group template")
				createTemplateOptionsModel.AccountID = core.StringPtr("accountID-123")
				createTemplateOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateOptionsModel.Group = accessGroupRequestModel
				createTemplateOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.CreateTemplate(createTemplateOptionsModel)
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
	Describe(`ListTemplates(listTemplatesOptions *ListTemplatesOptions) - Operation response error`, func() {
		listTemplatesPath := "/v1/group_templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTemplatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"accountID-123"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// TODO: Add check for verbose query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTemplates with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(iamaccessgroupsv2.ListTemplatesOptions)
				listTemplatesOptionsModel.AccountID = core.StringPtr("accountID-123")
				listTemplatesOptionsModel.TransactionID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Verbose = core.BoolPtr(true)
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTemplates(listTemplatesOptions *ListTemplatesOptions)`, func() {
		listTemplatesPath := "/v1/group_templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"accountID-123"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// TODO: Add check for verbose query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "group_templates": [{"id": "ID", "name": "Name", "description": "Description", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListTemplates successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(iamaccessgroupsv2.ListTemplatesOptions)
				listTemplatesOptionsModel.AccountID = core.StringPtr("accountID-123")
				listTemplatesOptionsModel.TransactionID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Verbose = core.BoolPtr(true)
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ListTemplatesWithContext(ctx, listTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ListTemplatesWithContext(ctx, listTemplatesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"accountID-123"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// TODO: Add check for verbose query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "group_templates": [{"id": "ID", "name": "Name", "description": "Description", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListTemplates successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListTemplates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(iamaccessgroupsv2.ListTemplatesOptions)
				listTemplatesOptionsModel.AccountID = core.StringPtr("accountID-123")
				listTemplatesOptionsModel.TransactionID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Verbose = core.BoolPtr(true)
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTemplates with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(iamaccessgroupsv2.ListTemplatesOptions)
				listTemplatesOptionsModel.AccountID = core.StringPtr("accountID-123")
				listTemplatesOptionsModel.TransactionID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Verbose = core.BoolPtr(true)
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTemplatesOptions model with no property values
				listTemplatesOptionsModelNew := new(iamaccessgroupsv2.ListTemplatesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ListTemplates(listTemplatesOptionsModelNew)
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
			It(`Invoke ListTemplates successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(iamaccessgroupsv2.ListTemplatesOptions)
				listTemplatesOptionsModel.AccountID = core.StringPtr("accountID-123")
				listTemplatesOptionsModel.TransactionID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Verbose = core.BoolPtr(true)
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ListTemplates(listTemplatesOptionsModel)
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
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(iamaccessgroupsv2.ListTemplatesResponse)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(iamaccessgroupsv2.ListTemplatesResponse)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.ListTemplatesResponse)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.ListTemplatesResponse)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"group_templates":[{"id":"ID","name":"Name","description":"Description","version":"Version","committed":false,"group":{"name":"Name","description":"Description","members":{"users":["Users"],"services":["Services"],"action_controls":{"add":false,"remove":true}},"assertions":{"rules":[{"name":"Name","expiration":10,"realm_name":"RealmName","conditions":[{"claim":"Claim","operator":"Operator","value":"Value"}],"action_controls":{"remove":true}}],"action_controls":{"add":false,"remove":true}},"action_controls":{"access":{"add":false}}},"policy_template_references":[{"id":"ID","version":"Version"}],"href":"Href","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID","last_modified_at":"2019-01-01T12:00:00.000Z","last_modified_by_id":"LastModifiedByID"}],"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"group_templates":[{"id":"ID","name":"Name","description":"Description","version":"Version","committed":false,"group":{"name":"Name","description":"Description","members":{"users":["Users"],"services":["Services"],"action_controls":{"add":false,"remove":true}},"assertions":{"rules":[{"name":"Name","expiration":10,"realm_name":"RealmName","conditions":[{"claim":"Claim","operator":"Operator","value":"Value"}],"action_controls":{"remove":true}}],"action_controls":{"add":false,"remove":true}},"action_controls":{"access":{"add":false}}},"policy_template_references":[{"id":"ID","version":"Version"}],"href":"Href","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID","last_modified_at":"2019-01-01T12:00:00.000Z","last_modified_by_id":"LastModifiedByID"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use TemplatesPager.GetNext successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listTemplatesOptionsModel := &iamaccessgroupsv2.ListTemplatesOptions{
					AccountID:     core.StringPtr("accountID-123"),
					TransactionID: core.StringPtr("testString"),
					Limit:         core.Int64Ptr(int64(50)),
					Verbose:       core.BoolPtr(true),
				}

				pager, err := iamAccessGroupsService.NewTemplatesPager(listTemplatesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []iamaccessgroupsv2.GroupTemplate
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use TemplatesPager.GetAll successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listTemplatesOptionsModel := &iamaccessgroupsv2.ListTemplatesOptions{
					AccountID:     core.StringPtr("accountID-123"),
					TransactionID: core.StringPtr("testString"),
					Limit:         core.Int64Ptr(int64(50)),
					Verbose:       core.BoolPtr(true),
				}

				pager, err := iamAccessGroupsService.NewTemplatesPager(listTemplatesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateTemplateVersion(createTemplateVersionOptions *CreateTemplateVersionOptions) - Operation response error`, func() {
		createTemplateVersionPath := "/v1/group_templates/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTemplateVersionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTemplateVersion with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateVersionOptions model
				createTemplateVersionOptionsModel := new(iamaccessgroupsv2.CreateTemplateVersionOptions)
				createTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				createTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateVersionOptionsModel.Group = accessGroupRequestModel
				createTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTemplateVersion(createTemplateVersionOptions *CreateTemplateVersionOptions)`, func() {
		createTemplateVersionPath := "/v1/group_templates/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTemplateVersionPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateTemplateVersion successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateVersionOptions model
				createTemplateVersionOptionsModel := new(iamaccessgroupsv2.CreateTemplateVersionOptions)
				createTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				createTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateVersionOptionsModel.Group = accessGroupRequestModel
				createTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.CreateTemplateVersionWithContext(ctx, createTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.CreateTemplateVersionWithContext(ctx, createTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTemplateVersionPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateTemplateVersion successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.CreateTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateVersionOptions model
				createTemplateVersionOptionsModel := new(iamaccessgroupsv2.CreateTemplateVersionOptions)
				createTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				createTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateVersionOptionsModel.Group = accessGroupRequestModel
				createTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTemplateVersion with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateVersionOptions model
				createTemplateVersionOptionsModel := new(iamaccessgroupsv2.CreateTemplateVersionOptions)
				createTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				createTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateVersionOptionsModel.Group = accessGroupRequestModel
				createTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTemplateVersionOptions model with no property values
				createTemplateVersionOptionsModelNew := new(iamaccessgroupsv2.CreateTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateTemplateVersion successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the CreateTemplateVersionOptions model
				createTemplateVersionOptionsModel := new(iamaccessgroupsv2.CreateTemplateVersionOptions)
				createTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				createTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateVersionOptionsModel.Group = accessGroupRequestModel
				createTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				createTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				createTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptionsModel)
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
	Describe(`ListTemplateVersions(listTemplateVersionsOptions *ListTemplateVersionsOptions) - Operation response error`, func() {
		listTemplateVersionsPath := "/v1/group_templates/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTemplateVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTemplateVersions with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListTemplateVersionsOptions model
				listTemplateVersionsOptionsModel := new(iamaccessgroupsv2.ListTemplateVersionsOptions)
				listTemplateVersionsOptionsModel.TemplateID = core.StringPtr("testString")
				listTemplateVersionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listTemplateVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ListTemplateVersions(listTemplateVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ListTemplateVersions(listTemplateVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTemplateVersions(listTemplateVersionsOptions *ListTemplateVersionsOptions)`, func() {
		listTemplateVersionsPath := "/v1/group_templates/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTemplateVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "group_template_versions": [{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListTemplateVersions successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the ListTemplateVersionsOptions model
				listTemplateVersionsOptionsModel := new(iamaccessgroupsv2.ListTemplateVersionsOptions)
				listTemplateVersionsOptionsModel.TemplateID = core.StringPtr("testString")
				listTemplateVersionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listTemplateVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ListTemplateVersionsWithContext(ctx, listTemplateVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ListTemplateVersions(listTemplateVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ListTemplateVersionsWithContext(ctx, listTemplateVersionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTemplateVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(100))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "group_template_versions": [{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListTemplateVersions successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListTemplateVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTemplateVersionsOptions model
				listTemplateVersionsOptionsModel := new(iamaccessgroupsv2.ListTemplateVersionsOptions)
				listTemplateVersionsOptionsModel.TemplateID = core.StringPtr("testString")
				listTemplateVersionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listTemplateVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListTemplateVersions(listTemplateVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTemplateVersions with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListTemplateVersionsOptions model
				listTemplateVersionsOptionsModel := new(iamaccessgroupsv2.ListTemplateVersionsOptions)
				listTemplateVersionsOptionsModel.TemplateID = core.StringPtr("testString")
				listTemplateVersionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listTemplateVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ListTemplateVersions(listTemplateVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTemplateVersionsOptions model with no property values
				listTemplateVersionsOptionsModelNew := new(iamaccessgroupsv2.ListTemplateVersionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ListTemplateVersions(listTemplateVersionsOptionsModelNew)
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
			It(`Invoke ListTemplateVersions successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListTemplateVersionsOptions model
				listTemplateVersionsOptionsModel := new(iamaccessgroupsv2.ListTemplateVersionsOptions)
				listTemplateVersionsOptionsModel.TemplateID = core.StringPtr("testString")
				listTemplateVersionsOptionsModel.Limit = core.Int64Ptr(int64(100))
				listTemplateVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ListTemplateVersions(listTemplateVersionsOptionsModel)
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
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(iamaccessgroupsv2.ListTemplateVersionsResponse)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(iamaccessgroupsv2.ListTemplateVersionsResponse)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.ListTemplateVersionsResponse)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.ListTemplateVersionsResponse)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTemplateVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"group_template_versions":[{"name":"Name","description":"Description","account_id":"AccountID","version":"Version","committed":false,"group":{"name":"Name","description":"Description","members":{"users":["Users"],"services":["Services"],"action_controls":{"add":false,"remove":true}},"assertions":{"rules":[{"name":"Name","expiration":10,"realm_name":"RealmName","conditions":[{"claim":"Claim","operator":"Operator","value":"Value"}],"action_controls":{"remove":true}}],"action_controls":{"add":false,"remove":true}},"action_controls":{"access":{"add":false}}},"policy_template_references":[{"id":"ID","version":"Version"}],"href":"Href","created_at":"CreatedAt","created_by_id":"CreatedByID","last_modified_at":"LastModifiedAt","last_modified_by_id":"LastModifiedByID"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"group_template_versions":[{"name":"Name","description":"Description","account_id":"AccountID","version":"Version","committed":false,"group":{"name":"Name","description":"Description","members":{"users":["Users"],"services":["Services"],"action_controls":{"add":false,"remove":true}},"assertions":{"rules":[{"name":"Name","expiration":10,"realm_name":"RealmName","conditions":[{"claim":"Claim","operator":"Operator","value":"Value"}],"action_controls":{"remove":true}}],"action_controls":{"add":false,"remove":true}},"action_controls":{"access":{"add":false}}},"policy_template_references":[{"id":"ID","version":"Version"}],"href":"Href","created_at":"CreatedAt","created_by_id":"CreatedByID","last_modified_at":"LastModifiedAt","last_modified_by_id":"LastModifiedByID"}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use TemplateVersionsPager.GetNext successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listTemplateVersionsOptionsModel := &iamaccessgroupsv2.ListTemplateVersionsOptions{
					TemplateID: core.StringPtr("testString"),
					Limit:      core.Int64Ptr(int64(100)),
				}

				pager, err := iamAccessGroupsService.NewTemplateVersionsPager(listTemplateVersionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []iamaccessgroupsv2.ListTemplateVersionResponse
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use TemplateVersionsPager.GetAll successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listTemplateVersionsOptionsModel := &iamaccessgroupsv2.ListTemplateVersionsOptions{
					TemplateID: core.StringPtr("testString"),
					Limit:      core.Int64Ptr(int64(100)),
				}

				pager, err := iamAccessGroupsService.NewTemplateVersionsPager(listTemplateVersionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetTemplateVersion(getTemplateVersionOptions *GetTemplateVersionOptions) - Operation response error`, func() {
		getTemplateVersionPath := "/v1/group_templates/testString/versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for verbose query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTemplateVersion with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetTemplateVersionOptions model
				getTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetTemplateVersionOptions)
				getTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTemplateVersion(getTemplateVersionOptions *GetTemplateVersionOptions)`, func() {
		getTemplateVersionPath := "/v1/group_templates/testString/versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for verbose query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetTemplateVersion successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the GetTemplateVersionOptions model
				getTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetTemplateVersionOptions)
				getTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.GetTemplateVersionWithContext(ctx, getTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.GetTemplateVersionWithContext(ctx, getTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for verbose query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetTemplateVersion successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTemplateVersionOptions model
				getTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetTemplateVersionOptions)
				getTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTemplateVersion with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetTemplateVersionOptions model
				getTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetTemplateVersionOptions)
				getTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTemplateVersionOptions model with no property values
				getTemplateVersionOptionsModelNew := new(iamaccessgroupsv2.GetTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptionsModelNew)
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
			It(`Invoke GetTemplateVersion successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetTemplateVersionOptions model
				getTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetTemplateVersionOptions)
				getTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptionsModel)
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
	Describe(`UpdateTemplateVersion(updateTemplateVersionOptions *UpdateTemplateVersionOptions) - Operation response error`, func() {
		updateTemplateVersionPath := "/v1/group_templates/testString/versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTemplateVersionPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "83adf5bd-de790caa3")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTemplateVersion with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-e371b0e5-1c80-48e3-bf12-c6a8ef2b1a11"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the UpdateTemplateVersionOptions model
				updateTemplateVersionOptionsModel := new(iamaccessgroupsv2.UpdateTemplateVersionOptions)
				updateTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				updateTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				updateTemplateVersionOptionsModel.Group = accessGroupRequestModel
				updateTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				updateTemplateVersionOptionsModel.TransactionID = core.StringPtr("83adf5bd-de790caa3")
				updateTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTemplateVersion(updateTemplateVersionOptions *UpdateTemplateVersionOptions)`, func() {
		updateTemplateVersionPath := "/v1/group_templates/testString/versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTemplateVersionPath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "83adf5bd-de790caa3")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke UpdateTemplateVersion successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-e371b0e5-1c80-48e3-bf12-c6a8ef2b1a11"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the UpdateTemplateVersionOptions model
				updateTemplateVersionOptionsModel := new(iamaccessgroupsv2.UpdateTemplateVersionOptions)
				updateTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				updateTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				updateTemplateVersionOptionsModel.Group = accessGroupRequestModel
				updateTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				updateTemplateVersionOptionsModel.TransactionID = core.StringPtr("83adf5bd-de790caa3")
				updateTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.UpdateTemplateVersionWithContext(ctx, updateTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.UpdateTemplateVersionWithContext(ctx, updateTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTemplateVersionPath))
					Expect(req.Method).To(Equal("PUT"))

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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "83adf5bd-de790caa3")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke UpdateTemplateVersion successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.UpdateTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-e371b0e5-1c80-48e3-bf12-c6a8ef2b1a11"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the UpdateTemplateVersionOptions model
				updateTemplateVersionOptionsModel := new(iamaccessgroupsv2.UpdateTemplateVersionOptions)
				updateTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				updateTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				updateTemplateVersionOptionsModel.Group = accessGroupRequestModel
				updateTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				updateTemplateVersionOptionsModel.TransactionID = core.StringPtr("83adf5bd-de790caa3")
				updateTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTemplateVersion with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-e371b0e5-1c80-48e3-bf12-c6a8ef2b1a11"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the UpdateTemplateVersionOptions model
				updateTemplateVersionOptionsModel := new(iamaccessgroupsv2.UpdateTemplateVersionOptions)
				updateTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				updateTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				updateTemplateVersionOptionsModel.Group = accessGroupRequestModel
				updateTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				updateTemplateVersionOptionsModel.TransactionID = core.StringPtr("83adf5bd-de790caa3")
				updateTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTemplateVersionOptions model with no property values
				updateTemplateVersionOptionsModelNew := new(iamaccessgroupsv2.UpdateTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke UpdateTemplateVersion successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				membersModel.Users = []string{"IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-e371b0e5-1c80-48e3-bf12-c6a8ef2b1a11"}
				membersModel.ActionControls = membersActionControlsModel

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				ruleActionControlsModel.Remove = core.BoolPtr(false)

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				accessActionControlsModel.Add = core.BoolPtr(false)

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				groupActionControlsModel.Access = accessActionControlsModel

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")

				// Construct an instance of the UpdateTemplateVersionOptions model
				updateTemplateVersionOptionsModel := new(iamaccessgroupsv2.UpdateTemplateVersionOptions)
				updateTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.IfMatch = core.StringPtr("testString")
				updateTemplateVersionOptionsModel.Name = core.StringPtr("IAM Admin Group template 2")
				updateTemplateVersionOptionsModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				updateTemplateVersionOptionsModel.Group = accessGroupRequestModel
				updateTemplateVersionOptionsModel.PolicyTemplateReferences = []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}
				updateTemplateVersionOptionsModel.TransactionID = core.StringPtr("83adf5bd-de790caa3")
				updateTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptionsModel)
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
	Describe(`DeleteTemplateVersion(deleteTemplateVersionOptions *DeleteTemplateVersionOptions)`, func() {
		deleteTemplateVersionPath := "/v1/group_templates/testString/versions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTemplateVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTemplateVersion successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.DeleteTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTemplateVersionOptions model
				deleteTemplateVersionOptionsModel := new(iamaccessgroupsv2.DeleteTemplateVersionOptions)
				deleteTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				deleteTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				deleteTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				deleteTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.DeleteTemplateVersion(deleteTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTemplateVersion with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the DeleteTemplateVersionOptions model
				deleteTemplateVersionOptionsModel := new(iamaccessgroupsv2.DeleteTemplateVersionOptions)
				deleteTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				deleteTemplateVersionOptionsModel.VersionNum = core.StringPtr("testString")
				deleteTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				deleteTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.DeleteTemplateVersion(deleteTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTemplateVersionOptions model with no property values
				deleteTemplateVersionOptionsModelNew := new(iamaccessgroupsv2.DeleteTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.DeleteTemplateVersion(deleteTemplateVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CommitTemplate(commitTemplateOptions *CommitTemplateOptions)`, func() {
		commitTemplatePath := "/v1/group_templates/testString/versions/testString/commit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(commitTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke CommitTemplate successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.CommitTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CommitTemplateOptions model
				commitTemplateOptionsModel := new(iamaccessgroupsv2.CommitTemplateOptions)
				commitTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				commitTemplateOptionsModel.VersionNum = core.StringPtr("testString")
				commitTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				commitTemplateOptionsModel.TransactionID = core.StringPtr("testString")
				commitTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.CommitTemplate(commitTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CommitTemplate with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CommitTemplateOptions model
				commitTemplateOptionsModel := new(iamaccessgroupsv2.CommitTemplateOptions)
				commitTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				commitTemplateOptionsModel.VersionNum = core.StringPtr("testString")
				commitTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				commitTemplateOptionsModel.TransactionID = core.StringPtr("testString")
				commitTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.CommitTemplate(commitTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CommitTemplateOptions model with no property values
				commitTemplateOptionsModelNew := new(iamaccessgroupsv2.CommitTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.CommitTemplate(commitTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLatestTemplateVersion(getLatestTemplateVersionOptions *GetLatestTemplateVersionOptions) - Operation response error`, func() {
		getLatestTemplateVersionPath := "/v1/group_templates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for verbose query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLatestTemplateVersion with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetLatestTemplateVersionOptions model
				getLatestTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetLatestTemplateVersionOptions)
				getLatestTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getLatestTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLatestTemplateVersion(getLatestTemplateVersionOptions *GetLatestTemplateVersionOptions)`, func() {
		getLatestTemplateVersionPath := "/v1/group_templates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for verbose query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetLatestTemplateVersion successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the GetLatestTemplateVersionOptions model
				getLatestTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetLatestTemplateVersionOptions)
				getLatestTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getLatestTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.GetLatestTemplateVersionWithContext(ctx, getLatestTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.GetLatestTemplateVersionWithContext(ctx, getLatestTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLatestTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for verbose query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "group": {"name": "Name", "description": "Description", "members": {"users": ["Users"], "services": ["Services"], "action_controls": {"add": false, "remove": true}}, "assertions": {"rules": [{"name": "Name", "expiration": 10, "realm_name": "RealmName", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "action_controls": {"remove": true}}], "action_controls": {"add": false, "remove": true}}, "action_controls": {"access": {"add": false}}}, "policy_template_references": [{"id": "ID", "version": "Version"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetLatestTemplateVersion successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetLatestTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLatestTemplateVersionOptions model
				getLatestTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetLatestTemplateVersionOptions)
				getLatestTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getLatestTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLatestTemplateVersion with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetLatestTemplateVersionOptions model
				getLatestTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetLatestTemplateVersionOptions)
				getLatestTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getLatestTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLatestTemplateVersionOptions model with no property values
				getLatestTemplateVersionOptionsModelNew := new(iamaccessgroupsv2.GetLatestTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptionsModelNew)
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
			It(`Invoke GetLatestTemplateVersion successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetLatestTemplateVersionOptions model
				getLatestTemplateVersionOptionsModel := new(iamaccessgroupsv2.GetLatestTemplateVersionOptions)
				getLatestTemplateVersionOptionsModel.TemplateID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Verbose = core.BoolPtr(true)
				getLatestTemplateVersionOptionsModel.TransactionID = core.StringPtr("testString")
				getLatestTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptionsModel)
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
	Describe(`DeleteTemplate(deleteTemplateOptions *DeleteTemplateOptions)`, func() {
		deleteTemplatePath := "/v1/group_templates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTemplatePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTemplate successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.DeleteTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTemplateOptions model
				deleteTemplateOptionsModel := new(iamaccessgroupsv2.DeleteTemplateOptions)
				deleteTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				deleteTemplateOptionsModel.TransactionID = core.StringPtr("testString")
				deleteTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.DeleteTemplate(deleteTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTemplate with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the DeleteTemplateOptions model
				deleteTemplateOptionsModel := new(iamaccessgroupsv2.DeleteTemplateOptions)
				deleteTemplateOptionsModel.TemplateID = core.StringPtr("testString")
				deleteTemplateOptionsModel.TransactionID = core.StringPtr("testString")
				deleteTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.DeleteTemplate(deleteTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTemplateOptions model with no property values
				deleteTemplateOptionsModelNew := new(iamaccessgroupsv2.DeleteTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.DeleteTemplate(deleteTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAssignment(createAssignmentOptions *CreateAssignmentOptions) - Operation response error`, func() {
		createAssignmentPath := "/v1/group_assignments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAssignmentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAssignment with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CreateAssignmentOptions model
				createAssignmentOptionsModel := new(iamaccessgroupsv2.CreateAssignmentOptions)
				createAssignmentOptionsModel.TemplateID = core.StringPtr("AccessGroupTemplateId-4be4")
				createAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				createAssignmentOptionsModel.TargetType = core.StringPtr("AccountGroup")
				createAssignmentOptionsModel.Target = core.StringPtr("0a45594d0f-123")
				createAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.CreateAssignment(createAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.CreateAssignment(createAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAssignment(createAssignmentOptions *CreateAssignmentOptions)`, func() {
		createAssignmentPath := "/v1/group_assignments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAssignmentPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": "TemplateVersion", "target_type": "Account", "target": "Target", "operation": "assign", "status": "accepted", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateAssignment successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the CreateAssignmentOptions model
				createAssignmentOptionsModel := new(iamaccessgroupsv2.CreateAssignmentOptions)
				createAssignmentOptionsModel.TemplateID = core.StringPtr("AccessGroupTemplateId-4be4")
				createAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				createAssignmentOptionsModel.TargetType = core.StringPtr("AccountGroup")
				createAssignmentOptionsModel.Target = core.StringPtr("0a45594d0f-123")
				createAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.CreateAssignmentWithContext(ctx, createAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.CreateAssignment(createAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.CreateAssignmentWithContext(ctx, createAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAssignmentPath))
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

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": "TemplateVersion", "target_type": "Account", "target": "Target", "operation": "assign", "status": "accepted", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateAssignment successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.CreateAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAssignmentOptions model
				createAssignmentOptionsModel := new(iamaccessgroupsv2.CreateAssignmentOptions)
				createAssignmentOptionsModel.TemplateID = core.StringPtr("AccessGroupTemplateId-4be4")
				createAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				createAssignmentOptionsModel.TargetType = core.StringPtr("AccountGroup")
				createAssignmentOptionsModel.Target = core.StringPtr("0a45594d0f-123")
				createAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.CreateAssignment(createAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAssignment with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CreateAssignmentOptions model
				createAssignmentOptionsModel := new(iamaccessgroupsv2.CreateAssignmentOptions)
				createAssignmentOptionsModel.TemplateID = core.StringPtr("AccessGroupTemplateId-4be4")
				createAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				createAssignmentOptionsModel.TargetType = core.StringPtr("AccountGroup")
				createAssignmentOptionsModel.Target = core.StringPtr("0a45594d0f-123")
				createAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.CreateAssignment(createAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAssignmentOptions model with no property values
				createAssignmentOptionsModelNew := new(iamaccessgroupsv2.CreateAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.CreateAssignment(createAssignmentOptionsModelNew)
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
			It(`Invoke CreateAssignment successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CreateAssignmentOptions model
				createAssignmentOptionsModel := new(iamaccessgroupsv2.CreateAssignmentOptions)
				createAssignmentOptionsModel.TemplateID = core.StringPtr("AccessGroupTemplateId-4be4")
				createAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				createAssignmentOptionsModel.TargetType = core.StringPtr("AccountGroup")
				createAssignmentOptionsModel.Target = core.StringPtr("0a45594d0f-123")
				createAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				createAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.CreateAssignment(createAssignmentOptionsModel)
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
	Describe(`ListAssignments(listAssignmentsOptions *ListAssignmentsOptions) - Operation response error`, func() {
		listAssignmentsPath := "/v1/group_assignments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"accountID-123"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"accepted"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAssignments with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAssignmentsOptions model
				listAssignmentsOptionsModel := new(iamaccessgroupsv2.ListAssignmentsOptions)
				listAssignmentsOptionsModel.AccountID = core.StringPtr("accountID-123")
				listAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAssignmentsOptionsModel.Status = core.StringPtr("accepted")
				listAssignmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAssignmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ListAssignments(listAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ListAssignments(listAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAssignments(listAssignmentsOptions *ListAssignmentsOptions)`, func() {
		listAssignmentsPath := "/v1/group_assignments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"accountID-123"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"accepted"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "assignments": [{"id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": "TemplateVersion", "target_type": "Account", "target": "Target", "operation": "assign", "status": "accepted", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListAssignments successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the ListAssignmentsOptions model
				listAssignmentsOptionsModel := new(iamaccessgroupsv2.ListAssignmentsOptions)
				listAssignmentsOptionsModel.AccountID = core.StringPtr("accountID-123")
				listAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAssignmentsOptionsModel.Status = core.StringPtr("accepted")
				listAssignmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAssignmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ListAssignmentsWithContext(ctx, listAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ListAssignments(listAssignmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ListAssignmentsWithContext(ctx, listAssignmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"accountID-123"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["target"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"accepted"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "last": {"href": "Href"}, "assignments": [{"id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": "TemplateVersion", "target_type": "Account", "target": "Target", "operation": "assign", "status": "accepted", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListAssignments successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListAssignments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAssignmentsOptions model
				listAssignmentsOptionsModel := new(iamaccessgroupsv2.ListAssignmentsOptions)
				listAssignmentsOptionsModel.AccountID = core.StringPtr("accountID-123")
				listAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAssignmentsOptionsModel.Status = core.StringPtr("accepted")
				listAssignmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAssignmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListAssignments(listAssignmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAssignments with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAssignmentsOptions model
				listAssignmentsOptionsModel := new(iamaccessgroupsv2.ListAssignmentsOptions)
				listAssignmentsOptionsModel.AccountID = core.StringPtr("accountID-123")
				listAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAssignmentsOptionsModel.Status = core.StringPtr("accepted")
				listAssignmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAssignmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ListAssignments(listAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAssignmentsOptions model with no property values
				listAssignmentsOptionsModelNew := new(iamaccessgroupsv2.ListAssignmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ListAssignments(listAssignmentsOptionsModelNew)
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
			It(`Invoke ListAssignments successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAssignmentsOptions model
				listAssignmentsOptionsModel := new(iamaccessgroupsv2.ListAssignmentsOptions)
				listAssignmentsOptionsModel.AccountID = core.StringPtr("accountID-123")
				listAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listAssignmentsOptionsModel.Target = core.StringPtr("testString")
				listAssignmentsOptionsModel.Status = core.StringPtr("accepted")
				listAssignmentsOptionsModel.TransactionID = core.StringPtr("testString")
				listAssignmentsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listAssignmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ListAssignments(listAssignmentsOptionsModel)
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
	Describe(`GetAssignment(getAssignmentOptions *GetAssignmentOptions) - Operation response error`, func() {
		getAssignmentPath := "/v1/group_assignments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAssignmentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for verbose query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAssignment with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAssignmentOptions model
				getAssignmentOptionsModel := new(iamaccessgroupsv2.GetAssignmentOptions)
				getAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				getAssignmentOptionsModel.Verbose = core.BoolPtr(false)
				getAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.GetAssignment(getAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.GetAssignment(getAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAssignment(getAssignmentOptions *GetAssignmentOptions)`, func() {
		getAssignmentPath := "/v1/group_assignments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAssignmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for verbose query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": "TemplateVersion", "target_type": "TargetType", "target": "Target", "operation": "Operation", "status": "Status", "resources": [{"target": "Target", "group": {"group": {"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}, "members": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}], "rules": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}]}, "policy_template_references": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAssignment successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the GetAssignmentOptions model
				getAssignmentOptionsModel := new(iamaccessgroupsv2.GetAssignmentOptions)
				getAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				getAssignmentOptionsModel.Verbose = core.BoolPtr(false)
				getAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.GetAssignmentWithContext(ctx, getAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.GetAssignment(getAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.GetAssignmentWithContext(ctx, getAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAssignmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for verbose query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": "TemplateVersion", "target_type": "TargetType", "target": "Target", "operation": "Operation", "status": "Status", "resources": [{"target": "Target", "group": {"group": {"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}, "members": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}], "rules": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}]}, "policy_template_references": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAssignment successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAssignmentOptions model
				getAssignmentOptionsModel := new(iamaccessgroupsv2.GetAssignmentOptions)
				getAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				getAssignmentOptionsModel.Verbose = core.BoolPtr(false)
				getAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetAssignment(getAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAssignment with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAssignmentOptions model
				getAssignmentOptionsModel := new(iamaccessgroupsv2.GetAssignmentOptions)
				getAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				getAssignmentOptionsModel.Verbose = core.BoolPtr(false)
				getAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.GetAssignment(getAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAssignmentOptions model with no property values
				getAssignmentOptionsModelNew := new(iamaccessgroupsv2.GetAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.GetAssignment(getAssignmentOptionsModelNew)
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
			It(`Invoke GetAssignment successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAssignmentOptions model
				getAssignmentOptionsModel := new(iamaccessgroupsv2.GetAssignmentOptions)
				getAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				getAssignmentOptionsModel.Verbose = core.BoolPtr(false)
				getAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.GetAssignment(getAssignmentOptionsModel)
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
	Describe(`UpdateAssignment(updateAssignmentOptions *UpdateAssignmentOptions) - Operation response error`, func() {
		updateAssignmentPath := "/v1/group_assignments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAssignmentPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAssignment with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAssignmentOptions model
				updateAssignmentOptionsModel := new(iamaccessgroupsv2.UpdateAssignmentOptions)
				updateAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				updateAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.UpdateAssignment(updateAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.UpdateAssignment(updateAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAssignment(updateAssignmentOptions *UpdateAssignmentOptions)`, func() {
		updateAssignmentPath := "/v1/group_assignments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAssignmentPath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": "TemplateVersion", "target_type": "TargetType", "target": "Target", "operation": "Operation", "status": "Status", "resources": [{"target": "Target", "group": {"group": {"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}, "members": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}], "rules": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}]}, "policy_template_references": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke UpdateAssignment successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAssignmentOptions model
				updateAssignmentOptionsModel := new(iamaccessgroupsv2.UpdateAssignmentOptions)
				updateAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				updateAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.UpdateAssignmentWithContext(ctx, updateAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.UpdateAssignment(updateAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.UpdateAssignmentWithContext(ctx, updateAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAssignmentPath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_id": "AccountID", "template_id": "TemplateID", "template_version": "TemplateVersion", "target_type": "TargetType", "target": "Target", "operation": "Operation", "status": "Status", "resources": [{"target": "Target", "group": {"group": {"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}, "members": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}], "rules": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}]}, "policy_template_references": [{"id": "ID", "name": "Name", "version": "Version", "resource": "Resource", "error": "Error", "operation": "Operation", "status": "Status"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke UpdateAssignment successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.UpdateAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAssignmentOptions model
				updateAssignmentOptionsModel := new(iamaccessgroupsv2.UpdateAssignmentOptions)
				updateAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				updateAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.UpdateAssignment(updateAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAssignment with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAssignmentOptions model
				updateAssignmentOptionsModel := new(iamaccessgroupsv2.UpdateAssignmentOptions)
				updateAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				updateAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.UpdateAssignment(updateAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAssignmentOptions model with no property values
				updateAssignmentOptionsModelNew := new(iamaccessgroupsv2.UpdateAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.UpdateAssignment(updateAssignmentOptionsModelNew)
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
			It(`Invoke UpdateAssignment successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAssignmentOptions model
				updateAssignmentOptionsModel := new(iamaccessgroupsv2.UpdateAssignmentOptions)
				updateAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				updateAssignmentOptionsModel.IfMatch = core.StringPtr("testString")
				updateAssignmentOptionsModel.TemplateVersion = core.StringPtr("1")
				updateAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.UpdateAssignment(updateAssignmentOptionsModel)
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
	Describe(`DeleteAssignment(deleteAssignmentOptions *DeleteAssignmentOptions)`, func() {
		deleteAssignmentPath := "/v1/group_assignments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAssignmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteAssignment successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.DeleteAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAssignmentOptions model
				deleteAssignmentOptionsModel := new(iamaccessgroupsv2.DeleteAssignmentOptions)
				deleteAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.DeleteAssignment(deleteAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAssignment with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the DeleteAssignmentOptions model
				deleteAssignmentOptionsModel := new(iamaccessgroupsv2.DeleteAssignmentOptions)
				deleteAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				deleteAssignmentOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.DeleteAssignment(deleteAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAssignmentOptions model with no property values
				deleteAssignmentOptionsModelNew := new(iamaccessgroupsv2.DeleteAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.DeleteAssignment(deleteAssignmentOptionsModelNew)
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
			iamAccessGroupsService, _ := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL:           "http://iamaccessgroupsv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAccessGroupRequest successfully`, func() {
				name := "testString"
				_model, err := iamAccessGroupsService.NewAccessGroupRequest(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAddAccessGroupRuleOptions successfully`, func() {
				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				Expect(ruleConditionsModel).ToNot(BeNil())
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")
				Expect(ruleConditionsModel.Claim).To(Equal(core.StringPtr("isManager")))
				Expect(ruleConditionsModel.Operator).To(Equal(core.StringPtr("EQUALS")))
				Expect(ruleConditionsModel.Value).To(Equal(core.StringPtr("true")))

				// Construct an instance of the AddAccessGroupRuleOptions model
				accessGroupID := "testString"
				addAccessGroupRuleOptionsExpiration := int64(12)
				addAccessGroupRuleOptionsRealmName := "https://idp.example.org/SAML2"
				addAccessGroupRuleOptionsConditions := []iamaccessgroupsv2.RuleConditions{}
				addAccessGroupRuleOptionsModel := iamAccessGroupsService.NewAddAccessGroupRuleOptions(accessGroupID, addAccessGroupRuleOptionsExpiration, addAccessGroupRuleOptionsRealmName, addAccessGroupRuleOptionsConditions)
				addAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				addAccessGroupRuleOptionsModel.SetExpiration(int64(12))
				addAccessGroupRuleOptionsModel.SetRealmName("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.SetConditions([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel})
				addAccessGroupRuleOptionsModel.SetName("Manager group rule")
				addAccessGroupRuleOptionsModel.SetTransactionID("testString")
				addAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(addAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(addAccessGroupRuleOptionsModel.Expiration).To(Equal(core.Int64Ptr(int64(12))))
				Expect(addAccessGroupRuleOptionsModel.RealmName).To(Equal(core.StringPtr("https://idp.example.org/SAML2")))
				Expect(addAccessGroupRuleOptionsModel.Conditions).To(Equal([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}))
				Expect(addAccessGroupRuleOptionsModel.Name).To(Equal(core.StringPtr("Manager group rule")))
				Expect(addAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddGroupMembersRequestMembersItem successfully`, func() {
				iamID := "testString"
				typeVar := "testString"
				_model, err := iamAccessGroupsService.NewAddGroupMembersRequestMembersItem(iamID, typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAddMemberToMultipleAccessGroupsOptions successfully`, func() {
				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				accountID := "testString"
				iamID := "testString"
				addMemberToMultipleAccessGroupsOptionsModel := iamAccessGroupsService.NewAddMemberToMultipleAccessGroupsOptions(accountID, iamID)
				addMemberToMultipleAccessGroupsOptionsModel.SetAccountID("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetIamID("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetType("user")
				addMemberToMultipleAccessGroupsOptionsModel.SetGroups([]string{"AccessGroupId-b0d32f56-f85c-4bf1-af37-7bbd92b1b2b3"})
				addMemberToMultipleAccessGroupsOptionsModel.SetTransactionID("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addMemberToMultipleAccessGroupsOptionsModel).ToNot(BeNil())
				Expect(addMemberToMultipleAccessGroupsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.Type).To(Equal(core.StringPtr("user")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.Groups).To(Equal([]string{"AccessGroupId-b0d32f56-f85c-4bf1-af37-7bbd92b1b2b3"}))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddMembersToAccessGroupOptions successfully`, func() {
				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				Expect(addGroupMembersRequestMembersItemModel).ToNot(BeNil())
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")
				Expect(addGroupMembersRequestMembersItemModel.IamID).To(Equal(core.StringPtr("IBMid-user1")))
				Expect(addGroupMembersRequestMembersItemModel.Type).To(Equal(core.StringPtr("user")))

				// Construct an instance of the AddMembersToAccessGroupOptions model
				accessGroupID := "testString"
				addMembersToAccessGroupOptionsModel := iamAccessGroupsService.NewAddMembersToAccessGroupOptions(accessGroupID)
				addMembersToAccessGroupOptionsModel.SetAccessGroupID("testString")
				addMembersToAccessGroupOptionsModel.SetMembers([]iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel})
				addMembersToAccessGroupOptionsModel.SetTransactionID("testString")
				addMembersToAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addMembersToAccessGroupOptionsModel).ToNot(BeNil())
				Expect(addMembersToAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(addMembersToAccessGroupOptionsModel.Members).To(Equal([]iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}))
				Expect(addMembersToAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addMembersToAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCommitTemplateOptions successfully`, func() {
				// Construct an instance of the CommitTemplateOptions model
				templateID := "testString"
				versionNum := "testString"
				ifMatch := "testString"
				commitTemplateOptionsModel := iamAccessGroupsService.NewCommitTemplateOptions(templateID, versionNum, ifMatch)
				commitTemplateOptionsModel.SetTemplateID("testString")
				commitTemplateOptionsModel.SetVersionNum("testString")
				commitTemplateOptionsModel.SetIfMatch("testString")
				commitTemplateOptionsModel.SetTransactionID("testString")
				commitTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(commitTemplateOptionsModel).ToNot(BeNil())
				Expect(commitTemplateOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(commitTemplateOptionsModel.VersionNum).To(Equal(core.StringPtr("testString")))
				Expect(commitTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(commitTemplateOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(commitTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccessGroupOptions successfully`, func() {
				// Construct an instance of the CreateAccessGroupOptions model
				accountID := "testString"
				createAccessGroupOptionsName := "Managers"
				createAccessGroupOptionsModel := iamAccessGroupsService.NewCreateAccessGroupOptions(accountID, createAccessGroupOptionsName)
				createAccessGroupOptionsModel.SetAccountID("testString")
				createAccessGroupOptionsModel.SetName("Managers")
				createAccessGroupOptionsModel.SetDescription("Group for managers")
				createAccessGroupOptionsModel.SetTransactionID("testString")
				createAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccessGroupOptionsModel).ToNot(BeNil())
				Expect(createAccessGroupOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessGroupOptionsModel.Name).To(Equal(core.StringPtr("Managers")))
				Expect(createAccessGroupOptionsModel.Description).To(Equal(core.StringPtr("Group for managers")))
				Expect(createAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAssignmentOptions successfully`, func() {
				// Construct an instance of the CreateAssignmentOptions model
				createAssignmentOptionsTemplateID := "AccessGroupTemplateId-4be4"
				createAssignmentOptionsTemplateVersion := "1"
				createAssignmentOptionsTargetType := "AccountGroup"
				createAssignmentOptionsTarget := "0a45594d0f-123"
				createAssignmentOptionsModel := iamAccessGroupsService.NewCreateAssignmentOptions(createAssignmentOptionsTemplateID, createAssignmentOptionsTemplateVersion, createAssignmentOptionsTargetType, createAssignmentOptionsTarget)
				createAssignmentOptionsModel.SetTemplateID("AccessGroupTemplateId-4be4")
				createAssignmentOptionsModel.SetTemplateVersion("1")
				createAssignmentOptionsModel.SetTargetType("AccountGroup")
				createAssignmentOptionsModel.SetTarget("0a45594d0f-123")
				createAssignmentOptionsModel.SetTransactionID("testString")
				createAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAssignmentOptionsModel).ToNot(BeNil())
				Expect(createAssignmentOptionsModel.TemplateID).To(Equal(core.StringPtr("AccessGroupTemplateId-4be4")))
				Expect(createAssignmentOptionsModel.TemplateVersion).To(Equal(core.StringPtr("1")))
				Expect(createAssignmentOptionsModel.TargetType).To(Equal(core.StringPtr("AccountGroup")))
				Expect(createAssignmentOptionsModel.Target).To(Equal(core.StringPtr("0a45594d0f-123")))
				Expect(createAssignmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTemplateOptions successfully`, func() {
				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				Expect(membersActionControlsModel).ToNot(BeNil())
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)
				Expect(membersActionControlsModel.Add).To(Equal(core.BoolPtr(true)))
				Expect(membersActionControlsModel.Remove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				Expect(membersModel).ToNot(BeNil())
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345", "iam-ServiceId-456"}
				membersModel.ActionControls = membersActionControlsModel
				Expect(membersModel.Users).To(Equal([]string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}))
				Expect(membersModel.Services).To(Equal([]string{"iam-ServiceId-345", "iam-ServiceId-456"}))
				Expect(membersModel.ActionControls).To(Equal(membersActionControlsModel))

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				Expect(conditionsModel).ToNot(BeNil())
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")
				Expect(conditionsModel.Claim).To(Equal(core.StringPtr("blueGroup")))
				Expect(conditionsModel.Operator).To(Equal(core.StringPtr("CONTAINS")))
				Expect(conditionsModel.Value).To(Equal(core.StringPtr("test-bluegroup-saml")))

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				Expect(ruleActionControlsModel).ToNot(BeNil())
				ruleActionControlsModel.Remove = core.BoolPtr(false)
				Expect(ruleActionControlsModel.Remove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				Expect(assertionsRuleModel).ToNot(BeNil())
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel
				Expect(assertionsRuleModel.Name).To(Equal(core.StringPtr("Manager group rule")))
				Expect(assertionsRuleModel.Expiration).To(Equal(core.Int64Ptr(int64(12))))
				Expect(assertionsRuleModel.RealmName).To(Equal(core.StringPtr("https://idp.example.org/SAML2")))
				Expect(assertionsRuleModel.Conditions).To(Equal([]iamaccessgroupsv2.Conditions{*conditionsModel}))
				Expect(assertionsRuleModel.ActionControls).To(Equal(ruleActionControlsModel))

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				Expect(assertionsActionControlsModel).ToNot(BeNil())
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)
				Expect(assertionsActionControlsModel.Add).To(Equal(core.BoolPtr(false)))
				Expect(assertionsActionControlsModel.Remove).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				Expect(assertionsModel).ToNot(BeNil())
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel
				Expect(assertionsModel.Rules).To(Equal([]iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}))
				Expect(assertionsModel.ActionControls).To(Equal(assertionsActionControlsModel))

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				Expect(accessActionControlsModel).ToNot(BeNil())
				accessActionControlsModel.Add = core.BoolPtr(false)
				Expect(accessActionControlsModel.Add).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				Expect(groupActionControlsModel).ToNot(BeNil())
				groupActionControlsModel.Access = accessActionControlsModel
				Expect(groupActionControlsModel.Access).To(Equal(accessActionControlsModel))

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				Expect(accessGroupRequestModel).ToNot(BeNil())
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel
				Expect(accessGroupRequestModel.Name).To(Equal(core.StringPtr("IAM Admin Group")))
				Expect(accessGroupRequestModel.Description).To(Equal(core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")))
				Expect(accessGroupRequestModel.Members).To(Equal(membersModel))
				Expect(accessGroupRequestModel.Assertions).To(Equal(assertionsModel))
				Expect(accessGroupRequestModel.ActionControls).To(Equal(groupActionControlsModel))

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				Expect(policyTemplatesModel).ToNot(BeNil())
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")
				Expect(policyTemplatesModel.ID).To(Equal(core.StringPtr("policyTemplateId-123")))
				Expect(policyTemplatesModel.Version).To(Equal(core.StringPtr("1")))

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsName := "IAM Admin Group template"
				createTemplateOptionsAccountID := "accountID-123"
				createTemplateOptionsModel := iamAccessGroupsService.NewCreateTemplateOptions(createTemplateOptionsName, createTemplateOptionsAccountID)
				createTemplateOptionsModel.SetName("IAM Admin Group template")
				createTemplateOptionsModel.SetAccountID("accountID-123")
				createTemplateOptionsModel.SetDescription("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateOptionsModel.SetGroup(accessGroupRequestModel)
				createTemplateOptionsModel.SetPolicyTemplateReferences([]iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel})
				createTemplateOptionsModel.SetTransactionID("testString")
				createTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTemplateOptionsModel).ToNot(BeNil())
				Expect(createTemplateOptionsModel.Name).To(Equal(core.StringPtr("IAM Admin Group template")))
				Expect(createTemplateOptionsModel.AccountID).To(Equal(core.StringPtr("accountID-123")))
				Expect(createTemplateOptionsModel.Description).To(Equal(core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")))
				Expect(createTemplateOptionsModel.Group).To(Equal(accessGroupRequestModel))
				Expect(createTemplateOptionsModel.PolicyTemplateReferences).To(Equal([]iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}))
				Expect(createTemplateOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTemplateVersionOptions successfully`, func() {
				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				Expect(membersActionControlsModel).ToNot(BeNil())
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)
				Expect(membersActionControlsModel.Add).To(Equal(core.BoolPtr(true)))
				Expect(membersActionControlsModel.Remove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				Expect(membersModel).ToNot(BeNil())
				membersModel.Users = []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-345"}
				membersModel.ActionControls = membersActionControlsModel
				Expect(membersModel.Users).To(Equal([]string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"}))
				Expect(membersModel.Services).To(Equal([]string{"iam-ServiceId-345"}))
				Expect(membersModel.ActionControls).To(Equal(membersActionControlsModel))

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				Expect(conditionsModel).ToNot(BeNil())
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")
				Expect(conditionsModel.Claim).To(Equal(core.StringPtr("blueGroup")))
				Expect(conditionsModel.Operator).To(Equal(core.StringPtr("CONTAINS")))
				Expect(conditionsModel.Value).To(Equal(core.StringPtr("test-bluegroup-saml")))

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				Expect(ruleActionControlsModel).ToNot(BeNil())
				ruleActionControlsModel.Remove = core.BoolPtr(false)
				Expect(ruleActionControlsModel.Remove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				Expect(assertionsRuleModel).ToNot(BeNil())
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel
				Expect(assertionsRuleModel.Name).To(Equal(core.StringPtr("Manager group rule")))
				Expect(assertionsRuleModel.Expiration).To(Equal(core.Int64Ptr(int64(12))))
				Expect(assertionsRuleModel.RealmName).To(Equal(core.StringPtr("https://idp.example.org/SAML2")))
				Expect(assertionsRuleModel.Conditions).To(Equal([]iamaccessgroupsv2.Conditions{*conditionsModel}))
				Expect(assertionsRuleModel.ActionControls).To(Equal(ruleActionControlsModel))

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				Expect(assertionsActionControlsModel).ToNot(BeNil())
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)
				Expect(assertionsActionControlsModel.Add).To(Equal(core.BoolPtr(false)))
				Expect(assertionsActionControlsModel.Remove).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				Expect(assertionsModel).ToNot(BeNil())
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel
				Expect(assertionsModel.Rules).To(Equal([]iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}))
				Expect(assertionsModel.ActionControls).To(Equal(assertionsActionControlsModel))

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				Expect(accessActionControlsModel).ToNot(BeNil())
				accessActionControlsModel.Add = core.BoolPtr(false)
				Expect(accessActionControlsModel.Add).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				Expect(groupActionControlsModel).ToNot(BeNil())
				groupActionControlsModel.Access = accessActionControlsModel
				Expect(groupActionControlsModel.Access).To(Equal(accessActionControlsModel))

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				Expect(accessGroupRequestModel).ToNot(BeNil())
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel
				Expect(accessGroupRequestModel.Name).To(Equal(core.StringPtr("IAM Admin Group 8")))
				Expect(accessGroupRequestModel.Description).To(Equal(core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")))
				Expect(accessGroupRequestModel.Members).To(Equal(membersModel))
				Expect(accessGroupRequestModel.Assertions).To(Equal(assertionsModel))
				Expect(accessGroupRequestModel.ActionControls).To(Equal(groupActionControlsModel))

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				Expect(policyTemplatesModel).ToNot(BeNil())
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")
				Expect(policyTemplatesModel.ID).To(Equal(core.StringPtr("policyTemplateId-123")))
				Expect(policyTemplatesModel.Version).To(Equal(core.StringPtr("1")))

				// Construct an instance of the CreateTemplateVersionOptions model
				templateID := "testString"
				createTemplateVersionOptionsModel := iamAccessGroupsService.NewCreateTemplateVersionOptions(templateID)
				createTemplateVersionOptionsModel.SetTemplateID("testString")
				createTemplateVersionOptionsModel.SetName("IAM Admin Group template 2")
				createTemplateVersionOptionsModel.SetDescription("This access group template allows admin access to all IAM platform services in the account.")
				createTemplateVersionOptionsModel.SetGroup(accessGroupRequestModel)
				createTemplateVersionOptionsModel.SetPolicyTemplateReferences([]iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel})
				createTemplateVersionOptionsModel.SetTransactionID("testString")
				createTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(createTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(createTemplateVersionOptionsModel.Name).To(Equal(core.StringPtr("IAM Admin Group template 2")))
				Expect(createTemplateVersionOptionsModel.Description).To(Equal(core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")))
				Expect(createTemplateVersionOptionsModel.Group).To(Equal(accessGroupRequestModel))
				Expect(createTemplateVersionOptionsModel.PolicyTemplateReferences).To(Equal([]iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}))
				Expect(createTemplateVersionOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccessGroupOptions successfully`, func() {
				// Construct an instance of the DeleteAccessGroupOptions model
				accessGroupID := "testString"
				deleteAccessGroupOptionsModel := iamAccessGroupsService.NewDeleteAccessGroupOptions(accessGroupID)
				deleteAccessGroupOptionsModel.SetAccessGroupID("testString")
				deleteAccessGroupOptionsModel.SetTransactionID("testString")
				deleteAccessGroupOptionsModel.SetForce(false)
				deleteAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccessGroupOptionsModel).ToNot(BeNil())
				Expect(deleteAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessGroupOptionsModel.Force).To(Equal(core.BoolPtr(false)))
				Expect(deleteAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAssignmentOptions successfully`, func() {
				// Construct an instance of the DeleteAssignmentOptions model
				assignmentID := "testString"
				deleteAssignmentOptionsModel := iamAccessGroupsService.NewDeleteAssignmentOptions(assignmentID)
				deleteAssignmentOptionsModel.SetAssignmentID("testString")
				deleteAssignmentOptionsModel.SetTransactionID("testString")
				deleteAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAssignmentOptionsModel).ToNot(BeNil())
				Expect(deleteAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAssignmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTemplateOptions successfully`, func() {
				// Construct an instance of the DeleteTemplateOptions model
				templateID := "testString"
				deleteTemplateOptionsModel := iamAccessGroupsService.NewDeleteTemplateOptions(templateID)
				deleteTemplateOptionsModel.SetTemplateID("testString")
				deleteTemplateOptionsModel.SetTransactionID("testString")
				deleteTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTemplateOptionsModel).ToNot(BeNil())
				Expect(deleteTemplateOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTemplateOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTemplateVersionOptions successfully`, func() {
				// Construct an instance of the DeleteTemplateVersionOptions model
				templateID := "testString"
				versionNum := "testString"
				deleteTemplateVersionOptionsModel := iamAccessGroupsService.NewDeleteTemplateVersionOptions(templateID, versionNum)
				deleteTemplateVersionOptionsModel.SetTemplateID("testString")
				deleteTemplateVersionOptionsModel.SetVersionNum("testString")
				deleteTemplateVersionOptionsModel.SetTransactionID("testString")
				deleteTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(deleteTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTemplateVersionOptionsModel.VersionNum).To(Equal(core.StringPtr("testString")))
				Expect(deleteTemplateVersionOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccessGroupOptions successfully`, func() {
				// Construct an instance of the GetAccessGroupOptions model
				accessGroupID := "testString"
				getAccessGroupOptionsModel := iamAccessGroupsService.NewGetAccessGroupOptions(accessGroupID)
				getAccessGroupOptionsModel.SetAccessGroupID("testString")
				getAccessGroupOptionsModel.SetTransactionID("testString")
				getAccessGroupOptionsModel.SetShowFederated(false)
				getAccessGroupOptionsModel.SetShowCRN(false)
				getAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccessGroupOptionsModel).ToNot(BeNil())
				Expect(getAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupOptionsModel.ShowFederated).To(Equal(core.BoolPtr(false)))
				Expect(getAccessGroupOptionsModel.ShowCRN).To(Equal(core.BoolPtr(false)))
				Expect(getAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccessGroupRuleOptions successfully`, func() {
				// Construct an instance of the GetAccessGroupRuleOptions model
				accessGroupID := "testString"
				ruleID := "testString"
				getAccessGroupRuleOptionsModel := iamAccessGroupsService.NewGetAccessGroupRuleOptions(accessGroupID, ruleID)
				getAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				getAccessGroupRuleOptionsModel.SetRuleID("testString")
				getAccessGroupRuleOptionsModel.SetTransactionID("testString")
				getAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(getAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountSettingsOptions successfully`, func() {
				// Construct an instance of the GetAccountSettingsOptions model
				accountID := "testString"
				getAccountSettingsOptionsModel := iamAccessGroupsService.NewGetAccountSettingsOptions(accountID)
				getAccountSettingsOptionsModel.SetAccountID("testString")
				getAccountSettingsOptionsModel.SetTransactionID("testString")
				getAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(getAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAssignmentOptions successfully`, func() {
				// Construct an instance of the GetAssignmentOptions model
				assignmentID := "testString"
				getAssignmentOptionsModel := iamAccessGroupsService.NewGetAssignmentOptions(assignmentID)
				getAssignmentOptionsModel.SetAssignmentID("testString")
				getAssignmentOptionsModel.SetTransactionID("testString")
				getAssignmentOptionsModel.SetVerbose(false)
				getAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAssignmentOptionsModel).ToNot(BeNil())
				Expect(getAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(getAssignmentOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAssignmentOptionsModel.Verbose).To(Equal(core.BoolPtr(false)))
				Expect(getAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLatestTemplateVersionOptions successfully`, func() {
				// Construct an instance of the GetLatestTemplateVersionOptions model
				templateID := "testString"
				getLatestTemplateVersionOptionsModel := iamAccessGroupsService.NewGetLatestTemplateVersionOptions(templateID)
				getLatestTemplateVersionOptionsModel.SetTemplateID("testString")
				getLatestTemplateVersionOptionsModel.SetVerbose(true)
				getLatestTemplateVersionOptionsModel.SetTransactionID("testString")
				getLatestTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLatestTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(getLatestTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(getLatestTemplateVersionOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(getLatestTemplateVersionOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getLatestTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTemplateVersionOptions successfully`, func() {
				// Construct an instance of the GetTemplateVersionOptions model
				templateID := "testString"
				versionNum := "testString"
				getTemplateVersionOptionsModel := iamAccessGroupsService.NewGetTemplateVersionOptions(templateID, versionNum)
				getTemplateVersionOptionsModel.SetTemplateID("testString")
				getTemplateVersionOptionsModel.SetVersionNum("testString")
				getTemplateVersionOptionsModel.SetVerbose(true)
				getTemplateVersionOptionsModel.SetTransactionID("testString")
				getTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(getTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateVersionOptionsModel.VersionNum).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateVersionOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(getTemplateVersionOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewIsMemberOfAccessGroupOptions successfully`, func() {
				// Construct an instance of the IsMemberOfAccessGroupOptions model
				accessGroupID := "testString"
				iamID := "testString"
				isMemberOfAccessGroupOptionsModel := iamAccessGroupsService.NewIsMemberOfAccessGroupOptions(accessGroupID, iamID)
				isMemberOfAccessGroupOptionsModel.SetAccessGroupID("testString")
				isMemberOfAccessGroupOptionsModel.SetIamID("testString")
				isMemberOfAccessGroupOptionsModel.SetTransactionID("testString")
				isMemberOfAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(isMemberOfAccessGroupOptionsModel).ToNot(BeNil())
				Expect(isMemberOfAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(isMemberOfAccessGroupOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(isMemberOfAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(isMemberOfAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccessGroupMembersOptions successfully`, func() {
				// Construct an instance of the ListAccessGroupMembersOptions model
				accessGroupID := "testString"
				listAccessGroupMembersOptionsModel := iamAccessGroupsService.NewListAccessGroupMembersOptions(accessGroupID)
				listAccessGroupMembersOptionsModel.SetAccessGroupID("testString")
				listAccessGroupMembersOptionsModel.SetTransactionID("testString")
				listAccessGroupMembersOptionsModel.SetMembershipType("static")
				listAccessGroupMembersOptionsModel.SetLimit(int64(10))
				listAccessGroupMembersOptionsModel.SetOffset(int64(0))
				listAccessGroupMembersOptionsModel.SetType("testString")
				listAccessGroupMembersOptionsModel.SetVerbose(false)
				listAccessGroupMembersOptionsModel.SetSort("testString")
				listAccessGroupMembersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessGroupMembersOptionsModel).ToNot(BeNil())
				Expect(listAccessGroupMembersOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.MembershipType).To(Equal(core.StringPtr("static")))
				Expect(listAccessGroupMembersOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAccessGroupMembersOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listAccessGroupMembersOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.Verbose).To(Equal(core.BoolPtr(false)))
				Expect(listAccessGroupMembersOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccessGroupRulesOptions successfully`, func() {
				// Construct an instance of the ListAccessGroupRulesOptions model
				accessGroupID := "testString"
				listAccessGroupRulesOptionsModel := iamAccessGroupsService.NewListAccessGroupRulesOptions(accessGroupID)
				listAccessGroupRulesOptionsModel.SetAccessGroupID("testString")
				listAccessGroupRulesOptionsModel.SetTransactionID("testString")
				listAccessGroupRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessGroupRulesOptionsModel).ToNot(BeNil())
				Expect(listAccessGroupRulesOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupRulesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccessGroupsOptions successfully`, func() {
				// Construct an instance of the ListAccessGroupsOptions model
				accountID := "testString"
				listAccessGroupsOptionsModel := iamAccessGroupsService.NewListAccessGroupsOptions(accountID)
				listAccessGroupsOptionsModel.SetAccountID("testString")
				listAccessGroupsOptionsModel.SetTransactionID("testString")
				listAccessGroupsOptionsModel.SetIamID("testString")
				listAccessGroupsOptionsModel.SetSearch("testString")
				listAccessGroupsOptionsModel.SetMembershipType("static")
				listAccessGroupsOptionsModel.SetLimit(int64(10))
				listAccessGroupsOptionsModel.SetOffset(int64(0))
				listAccessGroupsOptionsModel.SetSort("name")
				listAccessGroupsOptionsModel.SetShowFederated(false)
				listAccessGroupsOptionsModel.SetHidePublicAccess(false)
				listAccessGroupsOptionsModel.SetShowCRN(false)
				listAccessGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessGroupsOptionsModel).ToNot(BeNil())
				Expect(listAccessGroupsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.MembershipType).To(Equal(core.StringPtr("static")))
				Expect(listAccessGroupsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAccessGroupsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listAccessGroupsOptionsModel.Sort).To(Equal(core.StringPtr("name")))
				Expect(listAccessGroupsOptionsModel.ShowFederated).To(Equal(core.BoolPtr(false)))
				Expect(listAccessGroupsOptionsModel.HidePublicAccess).To(Equal(core.BoolPtr(false)))
				Expect(listAccessGroupsOptionsModel.ShowCRN).To(Equal(core.BoolPtr(false)))
				Expect(listAccessGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAssignmentsOptions successfully`, func() {
				// Construct an instance of the ListAssignmentsOptions model
				accountID := "accountID-123"
				listAssignmentsOptionsModel := iamAccessGroupsService.NewListAssignmentsOptions(accountID)
				listAssignmentsOptionsModel.SetAccountID("accountID-123")
				listAssignmentsOptionsModel.SetTemplateID("testString")
				listAssignmentsOptionsModel.SetTemplateVersion("testString")
				listAssignmentsOptionsModel.SetTarget("testString")
				listAssignmentsOptionsModel.SetStatus("accepted")
				listAssignmentsOptionsModel.SetTransactionID("testString")
				listAssignmentsOptionsModel.SetLimit(int64(50))
				listAssignmentsOptionsModel.SetOffset(int64(0))
				listAssignmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAssignmentsOptionsModel).ToNot(BeNil())
				Expect(listAssignmentsOptionsModel.AccountID).To(Equal(core.StringPtr("accountID-123")))
				Expect(listAssignmentsOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(listAssignmentsOptionsModel.TemplateVersion).To(Equal(core.StringPtr("testString")))
				Expect(listAssignmentsOptionsModel.Target).To(Equal(core.StringPtr("testString")))
				Expect(listAssignmentsOptionsModel.Status).To(Equal(core.StringPtr("accepted")))
				Expect(listAssignmentsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAssignmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listAssignmentsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listAssignmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTemplateVersionsOptions successfully`, func() {
				// Construct an instance of the ListTemplateVersionsOptions model
				templateID := "testString"
				listTemplateVersionsOptionsModel := iamAccessGroupsService.NewListTemplateVersionsOptions(templateID)
				listTemplateVersionsOptionsModel.SetTemplateID("testString")
				listTemplateVersionsOptionsModel.SetLimit(int64(100))
				listTemplateVersionsOptionsModel.SetOffset(int64(0))
				listTemplateVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTemplateVersionsOptionsModel).ToNot(BeNil())
				Expect(listTemplateVersionsOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(listTemplateVersionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(100))))
				Expect(listTemplateVersionsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTemplateVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTemplatesOptions successfully`, func() {
				// Construct an instance of the ListTemplatesOptions model
				accountID := "accountID-123"
				listTemplatesOptionsModel := iamAccessGroupsService.NewListTemplatesOptions(accountID)
				listTemplatesOptionsModel.SetAccountID("accountID-123")
				listTemplatesOptionsModel.SetTransactionID("testString")
				listTemplatesOptionsModel.SetLimit(int64(50))
				listTemplatesOptionsModel.SetOffset(int64(0))
				listTemplatesOptionsModel.SetVerbose(true)
				listTemplatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTemplatesOptionsModel).ToNot(BeNil())
				Expect(listTemplatesOptionsModel.AccountID).To(Equal(core.StringPtr("accountID-123")))
				Expect(listTemplatesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listTemplatesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listTemplatesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTemplatesOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(listTemplatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveAccessGroupRuleOptions successfully`, func() {
				// Construct an instance of the RemoveAccessGroupRuleOptions model
				accessGroupID := "testString"
				ruleID := "testString"
				removeAccessGroupRuleOptionsModel := iamAccessGroupsService.NewRemoveAccessGroupRuleOptions(accessGroupID, ruleID)
				removeAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				removeAccessGroupRuleOptionsModel.SetRuleID("testString")
				removeAccessGroupRuleOptionsModel.SetTransactionID("testString")
				removeAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(removeAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(removeAccessGroupRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(removeAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(removeAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveMemberFromAccessGroupOptions successfully`, func() {
				// Construct an instance of the RemoveMemberFromAccessGroupOptions model
				accessGroupID := "testString"
				iamID := "testString"
				removeMemberFromAccessGroupOptionsModel := iamAccessGroupsService.NewRemoveMemberFromAccessGroupOptions(accessGroupID, iamID)
				removeMemberFromAccessGroupOptionsModel.SetAccessGroupID("testString")
				removeMemberFromAccessGroupOptionsModel.SetIamID("testString")
				removeMemberFromAccessGroupOptionsModel.SetTransactionID("testString")
				removeMemberFromAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeMemberFromAccessGroupOptionsModel).ToNot(BeNil())
				Expect(removeMemberFromAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAccessGroupOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveMemberFromAllAccessGroupsOptions successfully`, func() {
				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				accountID := "testString"
				iamID := "testString"
				removeMemberFromAllAccessGroupsOptionsModel := iamAccessGroupsService.NewRemoveMemberFromAllAccessGroupsOptions(accountID, iamID)
				removeMemberFromAllAccessGroupsOptionsModel.SetAccountID("testString")
				removeMemberFromAllAccessGroupsOptionsModel.SetIamID("testString")
				removeMemberFromAllAccessGroupsOptionsModel.SetTransactionID("testString")
				removeMemberFromAllAccessGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeMemberFromAllAccessGroupsOptionsModel).ToNot(BeNil())
				Expect(removeMemberFromAllAccessGroupsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAllAccessGroupsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAllAccessGroupsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAllAccessGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveMembersFromAccessGroupOptions successfully`, func() {
				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				accessGroupID := "testString"
				removeMembersFromAccessGroupOptionsModel := iamAccessGroupsService.NewRemoveMembersFromAccessGroupOptions(accessGroupID)
				removeMembersFromAccessGroupOptionsModel.SetAccessGroupID("testString")
				removeMembersFromAccessGroupOptionsModel.SetMembers([]string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"})
				removeMembersFromAccessGroupOptionsModel.SetTransactionID("testString")
				removeMembersFromAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeMembersFromAccessGroupOptionsModel).ToNot(BeNil())
				Expect(removeMembersFromAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(removeMembersFromAccessGroupOptionsModel.Members).To(Equal([]string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}))
				Expect(removeMembersFromAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(removeMembersFromAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceAccessGroupRuleOptions successfully`, func() {
				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				Expect(ruleConditionsModel).ToNot(BeNil())
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")
				Expect(ruleConditionsModel.Claim).To(Equal(core.StringPtr("isManager")))
				Expect(ruleConditionsModel.Operator).To(Equal(core.StringPtr("EQUALS")))
				Expect(ruleConditionsModel.Value).To(Equal(core.StringPtr("true")))

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				accessGroupID := "testString"
				ruleID := "testString"
				ifMatch := "testString"
				replaceAccessGroupRuleOptionsExpiration := int64(12)
				replaceAccessGroupRuleOptionsRealmName := "https://idp.example.org/SAML2"
				replaceAccessGroupRuleOptionsConditions := []iamaccessgroupsv2.RuleConditions{}
				replaceAccessGroupRuleOptionsModel := iamAccessGroupsService.NewReplaceAccessGroupRuleOptions(accessGroupID, ruleID, ifMatch, replaceAccessGroupRuleOptionsExpiration, replaceAccessGroupRuleOptionsRealmName, replaceAccessGroupRuleOptionsConditions)
				replaceAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				replaceAccessGroupRuleOptionsModel.SetRuleID("testString")
				replaceAccessGroupRuleOptionsModel.SetIfMatch("testString")
				replaceAccessGroupRuleOptionsModel.SetExpiration(int64(12))
				replaceAccessGroupRuleOptionsModel.SetRealmName("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.SetConditions([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel})
				replaceAccessGroupRuleOptionsModel.SetName("Manager group rule")
				replaceAccessGroupRuleOptionsModel.SetTransactionID("testString")
				replaceAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(replaceAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.Expiration).To(Equal(core.Int64Ptr(int64(12))))
				Expect(replaceAccessGroupRuleOptionsModel.RealmName).To(Equal(core.StringPtr("https://idp.example.org/SAML2")))
				Expect(replaceAccessGroupRuleOptionsModel.Conditions).To(Equal([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}))
				Expect(replaceAccessGroupRuleOptionsModel.Name).To(Equal(core.StringPtr("Manager group rule")))
				Expect(replaceAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRuleConditions successfully`, func() {
				claim := "testString"
				operator := "EQUALS"
				value := "testString"
				_model, err := iamAccessGroupsService.NewRuleConditions(claim, operator, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateAccessGroupOptions successfully`, func() {
				// Construct an instance of the UpdateAccessGroupOptions model
				accessGroupID := "testString"
				ifMatch := "testString"
				updateAccessGroupOptionsModel := iamAccessGroupsService.NewUpdateAccessGroupOptions(accessGroupID, ifMatch)
				updateAccessGroupOptionsModel.SetAccessGroupID("testString")
				updateAccessGroupOptionsModel.SetIfMatch("testString")
				updateAccessGroupOptionsModel.SetName("Awesome Managers")
				updateAccessGroupOptionsModel.SetDescription("Group for awesome managers.")
				updateAccessGroupOptionsModel.SetTransactionID("testString")
				updateAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccessGroupOptionsModel).ToNot(BeNil())
				Expect(updateAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.Name).To(Equal(core.StringPtr("Awesome Managers")))
				Expect(updateAccessGroupOptionsModel.Description).To(Equal(core.StringPtr("Group for awesome managers.")))
				Expect(updateAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountSettingsOptions successfully`, func() {
				// Construct an instance of the UpdateAccountSettingsOptions model
				accountID := "testString"
				updateAccountSettingsOptionsModel := iamAccessGroupsService.NewUpdateAccountSettingsOptions(accountID)
				updateAccountSettingsOptionsModel.SetAccountID("testString")
				updateAccountSettingsOptionsModel.SetPublicAccessEnabled(true)
				updateAccountSettingsOptionsModel.SetTransactionID("testString")
				updateAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(updateAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.PublicAccessEnabled).To(Equal(core.BoolPtr(true)))
				Expect(updateAccountSettingsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAssignmentOptions successfully`, func() {
				// Construct an instance of the UpdateAssignmentOptions model
				assignmentID := "testString"
				ifMatch := "testString"
				updateAssignmentOptionsTemplateVersion := "1"
				updateAssignmentOptionsModel := iamAccessGroupsService.NewUpdateAssignmentOptions(assignmentID, ifMatch, updateAssignmentOptionsTemplateVersion)
				updateAssignmentOptionsModel.SetAssignmentID("testString")
				updateAssignmentOptionsModel.SetIfMatch("testString")
				updateAssignmentOptionsModel.SetTemplateVersion("1")
				updateAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAssignmentOptionsModel).ToNot(BeNil())
				Expect(updateAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateAssignmentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAssignmentOptionsModel.TemplateVersion).To(Equal(core.StringPtr("1")))
				Expect(updateAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTemplateVersionOptions successfully`, func() {
				// Construct an instance of the MembersActionControls model
				membersActionControlsModel := new(iamaccessgroupsv2.MembersActionControls)
				Expect(membersActionControlsModel).ToNot(BeNil())
				membersActionControlsModel.Add = core.BoolPtr(true)
				membersActionControlsModel.Remove = core.BoolPtr(false)
				Expect(membersActionControlsModel.Add).To(Equal(core.BoolPtr(true)))
				Expect(membersActionControlsModel.Remove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the Members model
				membersModel := new(iamaccessgroupsv2.Members)
				Expect(membersModel).ToNot(BeNil())
				membersModel.Users = []string{"IBMid-665000T8WY"}
				membersModel.Services = []string{"iam-ServiceId-e371b0e5-1c80-48e3-bf12-c6a8ef2b1a11"}
				membersModel.ActionControls = membersActionControlsModel
				Expect(membersModel.Users).To(Equal([]string{"IBMid-665000T8WY"}))
				Expect(membersModel.Services).To(Equal([]string{"iam-ServiceId-e371b0e5-1c80-48e3-bf12-c6a8ef2b1a11"}))
				Expect(membersModel.ActionControls).To(Equal(membersActionControlsModel))

				// Construct an instance of the Conditions model
				conditionsModel := new(iamaccessgroupsv2.Conditions)
				Expect(conditionsModel).ToNot(BeNil())
				conditionsModel.Claim = core.StringPtr("blueGroup")
				conditionsModel.Operator = core.StringPtr("CONTAINS")
				conditionsModel.Value = core.StringPtr("test-bluegroup-saml")
				Expect(conditionsModel.Claim).To(Equal(core.StringPtr("blueGroup")))
				Expect(conditionsModel.Operator).To(Equal(core.StringPtr("CONTAINS")))
				Expect(conditionsModel.Value).To(Equal(core.StringPtr("test-bluegroup-saml")))

				// Construct an instance of the RuleActionControls model
				ruleActionControlsModel := new(iamaccessgroupsv2.RuleActionControls)
				Expect(ruleActionControlsModel).ToNot(BeNil())
				ruleActionControlsModel.Remove = core.BoolPtr(false)
				Expect(ruleActionControlsModel.Remove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the AssertionsRule model
				assertionsRuleModel := new(iamaccessgroupsv2.AssertionsRule)
				Expect(assertionsRuleModel).ToNot(BeNil())
				assertionsRuleModel.Name = core.StringPtr("Manager group rule")
				assertionsRuleModel.Expiration = core.Int64Ptr(int64(12))
				assertionsRuleModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				assertionsRuleModel.Conditions = []iamaccessgroupsv2.Conditions{*conditionsModel}
				assertionsRuleModel.ActionControls = ruleActionControlsModel
				Expect(assertionsRuleModel.Name).To(Equal(core.StringPtr("Manager group rule")))
				Expect(assertionsRuleModel.Expiration).To(Equal(core.Int64Ptr(int64(12))))
				Expect(assertionsRuleModel.RealmName).To(Equal(core.StringPtr("https://idp.example.org/SAML2")))
				Expect(assertionsRuleModel.Conditions).To(Equal([]iamaccessgroupsv2.Conditions{*conditionsModel}))
				Expect(assertionsRuleModel.ActionControls).To(Equal(ruleActionControlsModel))

				// Construct an instance of the AssertionsActionControls model
				assertionsActionControlsModel := new(iamaccessgroupsv2.AssertionsActionControls)
				Expect(assertionsActionControlsModel).ToNot(BeNil())
				assertionsActionControlsModel.Add = core.BoolPtr(false)
				assertionsActionControlsModel.Remove = core.BoolPtr(true)
				Expect(assertionsActionControlsModel.Add).To(Equal(core.BoolPtr(false)))
				Expect(assertionsActionControlsModel.Remove).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Assertions model
				assertionsModel := new(iamaccessgroupsv2.Assertions)
				Expect(assertionsModel).ToNot(BeNil())
				assertionsModel.Rules = []iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}
				assertionsModel.ActionControls = assertionsActionControlsModel
				Expect(assertionsModel.Rules).To(Equal([]iamaccessgroupsv2.AssertionsRule{*assertionsRuleModel}))
				Expect(assertionsModel.ActionControls).To(Equal(assertionsActionControlsModel))

				// Construct an instance of the AccessActionControls model
				accessActionControlsModel := new(iamaccessgroupsv2.AccessActionControls)
				Expect(accessActionControlsModel).ToNot(BeNil())
				accessActionControlsModel.Add = core.BoolPtr(false)
				Expect(accessActionControlsModel.Add).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the GroupActionControls model
				groupActionControlsModel := new(iamaccessgroupsv2.GroupActionControls)
				Expect(groupActionControlsModel).ToNot(BeNil())
				groupActionControlsModel.Access = accessActionControlsModel
				Expect(groupActionControlsModel.Access).To(Equal(accessActionControlsModel))

				// Construct an instance of the AccessGroupRequest model
				accessGroupRequestModel := new(iamaccessgroupsv2.AccessGroupRequest)
				Expect(accessGroupRequestModel).ToNot(BeNil())
				accessGroupRequestModel.Name = core.StringPtr("IAM Admin Group 8")
				accessGroupRequestModel.Description = core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")
				accessGroupRequestModel.Members = membersModel
				accessGroupRequestModel.Assertions = assertionsModel
				accessGroupRequestModel.ActionControls = groupActionControlsModel
				Expect(accessGroupRequestModel.Name).To(Equal(core.StringPtr("IAM Admin Group 8")))
				Expect(accessGroupRequestModel.Description).To(Equal(core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")))
				Expect(accessGroupRequestModel.Members).To(Equal(membersModel))
				Expect(accessGroupRequestModel.Assertions).To(Equal(assertionsModel))
				Expect(accessGroupRequestModel.ActionControls).To(Equal(groupActionControlsModel))

				// Construct an instance of the PolicyTemplates model
				policyTemplatesModel := new(iamaccessgroupsv2.PolicyTemplates)
				Expect(policyTemplatesModel).ToNot(BeNil())
				policyTemplatesModel.ID = core.StringPtr("policyTemplateId-123")
				policyTemplatesModel.Version = core.StringPtr("1")
				Expect(policyTemplatesModel.ID).To(Equal(core.StringPtr("policyTemplateId-123")))
				Expect(policyTemplatesModel.Version).To(Equal(core.StringPtr("1")))

				// Construct an instance of the UpdateTemplateVersionOptions model
				templateID := "testString"
				versionNum := "testString"
				ifMatch := "testString"
				updateTemplateVersionOptionsModel := iamAccessGroupsService.NewUpdateTemplateVersionOptions(templateID, versionNum, ifMatch)
				updateTemplateVersionOptionsModel.SetTemplateID("testString")
				updateTemplateVersionOptionsModel.SetVersionNum("testString")
				updateTemplateVersionOptionsModel.SetIfMatch("testString")
				updateTemplateVersionOptionsModel.SetName("IAM Admin Group template 2")
				updateTemplateVersionOptionsModel.SetDescription("This access group template allows admin access to all IAM platform services in the account.")
				updateTemplateVersionOptionsModel.SetGroup(accessGroupRequestModel)
				updateTemplateVersionOptionsModel.SetPolicyTemplateReferences([]iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel})
				updateTemplateVersionOptionsModel.SetTransactionID("83adf5bd-de790caa3")
				updateTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(updateTemplateVersionOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateVersionOptionsModel.VersionNum).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateVersionOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateVersionOptionsModel.Name).To(Equal(core.StringPtr("IAM Admin Group template 2")))
				Expect(updateTemplateVersionOptionsModel.Description).To(Equal(core.StringPtr("This access group template allows admin access to all IAM platform services in the account.")))
				Expect(updateTemplateVersionOptionsModel.Group).To(Equal(accessGroupRequestModel))
				Expect(updateTemplateVersionOptionsModel.PolicyTemplateReferences).To(Equal([]iamaccessgroupsv2.PolicyTemplates{*policyTemplatesModel}))
				Expect(updateTemplateVersionOptionsModel.TransactionID).To(Equal(core.StringPtr("83adf5bd-de790caa3")))
				Expect(updateTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAccessActionControls successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.AccessActionControls)
			model.Add = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.AccessActionControls
			err = iamaccessgroupsv2.UnmarshalAccessActionControls(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAccessGroupRequest successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.AccessGroupRequest)
			model.Name = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Members = nil
			model.Assertions = nil
			model.ActionControls = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.AccessGroupRequest
			err = iamaccessgroupsv2.UnmarshalAccessGroupRequest(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAddGroupMembersRequestMembersItem successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
			model.IamID = core.StringPtr("testString")
			model.Type = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.AddGroupMembersRequestMembersItem
			err = iamaccessgroupsv2.UnmarshalAddGroupMembersRequestMembersItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAssertions successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.Assertions)
			model.Rules = nil
			model.ActionControls = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.Assertions
			err = iamaccessgroupsv2.UnmarshalAssertions(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAssertionsActionControls successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.AssertionsActionControls)
			model.Add = core.BoolPtr(true)
			model.Remove = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.AssertionsActionControls
			err = iamaccessgroupsv2.UnmarshalAssertionsActionControls(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAssertionsRule successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.AssertionsRule)
			model.Name = core.StringPtr("testString")
			model.Expiration = core.Int64Ptr(int64(38))
			model.RealmName = core.StringPtr("testString")
			model.Conditions = nil
			model.ActionControls = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.AssertionsRule
			err = iamaccessgroupsv2.UnmarshalAssertionsRule(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalConditions successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.Conditions)
			model.Claim = core.StringPtr("testString")
			model.Operator = core.StringPtr("testString")
			model.Value = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.Conditions
			err = iamaccessgroupsv2.UnmarshalConditions(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGroupActionControls successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.GroupActionControls)
			model.Access = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.GroupActionControls
			err = iamaccessgroupsv2.UnmarshalGroupActionControls(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalMembers successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.Members)
			model.Users = []string{"testString"}
			model.Services = []string{"testString"}
			model.ActionControls = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.Members
			err = iamaccessgroupsv2.UnmarshalMembers(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalMembersActionControls successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.MembersActionControls)
			model.Add = core.BoolPtr(true)
			model.Remove = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.MembersActionControls
			err = iamaccessgroupsv2.UnmarshalMembersActionControls(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPolicyTemplates successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.PolicyTemplates)
			model.ID = core.StringPtr("testString")
			model.Version = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.PolicyTemplates
			err = iamaccessgroupsv2.UnmarshalPolicyTemplates(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRuleActionControls successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.RuleActionControls)
			model.Remove = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.RuleActionControls
			err = iamaccessgroupsv2.UnmarshalRuleActionControls(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRuleConditions successfully`, func() {
			// Construct an instance of the model.
			model := new(iamaccessgroupsv2.RuleConditions)
			model.Claim = core.StringPtr("testString")
			model.Operator = core.StringPtr("EQUALS")
			model.Value = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *iamaccessgroupsv2.RuleConditions
			err = iamaccessgroupsv2.UnmarshalRuleConditions(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
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

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
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
