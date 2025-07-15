/**
 * (C) Copyright IBM Corp. 2022.
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

package casemanagementv1_test

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
	"github.com/IBM/platform-services-go-sdk/casemanagementv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`CaseManagementV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(caseManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(caseManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
				URL: "https://casemanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(caseManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CASE_MANAGEMENT_URL":       "https://casemanagementv1/api",
				"CASE_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{})
				Expect(caseManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := caseManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != caseManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(caseManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(caseManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(caseManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := caseManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != caseManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(caseManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(caseManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{})
				err := caseManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := caseManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != caseManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(caseManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(caseManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CASE_MANAGEMENT_URL":       "https://casemanagementv1/api",
				"CASE_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(caseManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CASE_MANAGEMENT_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1UsingExternalConfig(&casemanagementv1.CaseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(caseManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = casemanagementv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetCases(getCasesOptions *GetCasesOptions) - Operation response error`, func() {
		getCasesPath := "/cases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCasesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"number"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCases with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := new(casemanagementv1.GetCasesOptions)
				getCasesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCasesOptionsModel.Search = core.StringPtr("testString")
				getCasesOptionsModel.Sort = core.StringPtr("number")
				getCasesOptionsModel.Status = []string{"new"}
				getCasesOptionsModel.Fields = []string{"number"}
				getCasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.GetCases(getCasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.GetCases(getCasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCases(getCasesOptions *GetCasesOptions)`, func() {
		getCasesPath := "/cases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"number"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "first": {"href": "Href"}, "next": {"href": "Href"}, "previous": {"href": "Href"}, "last": {"href": "Href"}, "cases": [{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}]}`)
				}))
			})
			It(`Invoke GetCases successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := new(casemanagementv1.GetCasesOptions)
				getCasesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCasesOptionsModel.Search = core.StringPtr("testString")
				getCasesOptionsModel.Sort = core.StringPtr("number")
				getCasesOptionsModel.Status = []string{"new"}
				getCasesOptionsModel.Fields = []string{"number"}
				getCasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.GetCasesWithContext(ctx, getCasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.GetCases(getCasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.GetCasesWithContext(ctx, getCasesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"number"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "first": {"href": "Href"}, "next": {"href": "Href"}, "previous": {"href": "Href"}, "last": {"href": "Href"}, "cases": [{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}]}`)
				}))
			})
			It(`Invoke GetCases successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.GetCases(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := new(casemanagementv1.GetCasesOptions)
				getCasesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCasesOptionsModel.Search = core.StringPtr("testString")
				getCasesOptionsModel.Sort = core.StringPtr("number")
				getCasesOptionsModel.Status = []string{"new"}
				getCasesOptionsModel.Fields = []string{"number"}
				getCasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.GetCases(getCasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCases with error: Operation request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := new(casemanagementv1.GetCasesOptions)
				getCasesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCasesOptionsModel.Search = core.StringPtr("testString")
				getCasesOptionsModel.Sort = core.StringPtr("number")
				getCasesOptionsModel.Status = []string{"new"}
				getCasesOptionsModel.Fields = []string{"number"}
				getCasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.GetCases(getCasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke GetCases successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := new(casemanagementv1.GetCasesOptions)
				getCasesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getCasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCasesOptionsModel.Search = core.StringPtr("testString")
				getCasesOptionsModel.Sort = core.StringPtr("number")
				getCasesOptionsModel.Status = []string{"new"}
				getCasesOptionsModel.Fields = []string{"number"}
				getCasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.GetCases(getCasesOptionsModel)
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
				responseObject := new(casemanagementv1.CaseList)
				nextObject := new(casemanagementv1.PaginationLink)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(casemanagementv1.CaseList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(casemanagementv1.CaseList)
				nextObject := new(casemanagementv1.PaginationLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(casemanagementv1.CaseList)
				nextObject := new(casemanagementv1.PaginationLink)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCasesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"cases":[{"number":"Number","short_description":"ShortDescription","description":"Description","created_at":"CreatedAt","created_by":{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"},"updated_at":"UpdatedAt","updated_by":{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"},"contact_type":"Cloud Support Center","contact":{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"},"status":"Status","severity":8,"support_tier":"Free","resolution":"Resolution","close_notes":"CloseNotes","eu":{"support":false,"data_center":"DataCenter"},"watchlist":[{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"}],"attachments":[{"id":"ID","filename":"Filename","size_in_bytes":11,"created_at":"CreatedAt","url":"URL"}],"offering":{"name":"Name","type":{"group":"crn_service_name","key":"Key","kind":"Kind","id":"ID"}},"resources":[{"crn":"CRN","name":"Name","type":"Type","url":"URL","note":"Note"}],"comments":[{"value":"Value","added_at":"AddedAt","added_by":{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"}}]}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"cases":[{"number":"Number","short_description":"ShortDescription","description":"Description","created_at":"CreatedAt","created_by":{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"},"updated_at":"UpdatedAt","updated_by":{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"},"contact_type":"Cloud Support Center","contact":{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"},"status":"Status","severity":8,"support_tier":"Free","resolution":"Resolution","close_notes":"CloseNotes","eu":{"support":false,"data_center":"DataCenter"},"watchlist":[{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"}],"attachments":[{"id":"ID","filename":"Filename","size_in_bytes":11,"created_at":"CreatedAt","url":"URL"}],"offering":{"name":"Name","type":{"group":"crn_service_name","key":"Key","kind":"Kind","id":"ID"}},"resources":[{"crn":"CRN","name":"Name","type":"Type","url":"URL","note":"Note"}],"comments":[{"value":"Value","added_at":"AddedAt","added_by":{"name":"Name","realm":"IBMid","user_id":"abc@ibm.com"}}]}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use GetCasesPager.GetNext successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				getCasesOptionsModel := &casemanagementv1.GetCasesOptions{
					Limit:  core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
					Sort:   core.StringPtr("number"),
					Status: []string{"new"},
					Fields: []string{"number"},
				}

				pager, err := caseManagementService.NewGetCasesPager(getCasesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []casemanagementv1.Case
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use GetCasesPager.GetAll successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				getCasesOptionsModel := &casemanagementv1.GetCasesOptions{
					Limit:  core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
					Sort:   core.StringPtr("number"),
					Status: []string{"new"},
					Fields: []string{"number"},
				}

				pager, err := caseManagementService.NewGetCasesPager(getCasesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateCase(createCaseOptions *CreateCaseOptions) - Operation response error`, func() {
		createCasePath := "/cases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCasePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCase with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))

				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				resourcePayloadModel.CRN = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsModel := new(casemanagementv1.CreateCaseOptions)
				createCaseOptionsModel.Type = core.StringPtr("technical")
				createCaseOptionsModel.Subject = core.StringPtr("testString")
				createCaseOptionsModel.Description = core.StringPtr("testString")
				createCaseOptionsModel.Severity = core.Int64Ptr(int64(1))
				createCaseOptionsModel.Eu = casePayloadEuModel
				createCaseOptionsModel.Offering = offeringModel
				createCaseOptionsModel.Resources = []casemanagementv1.ResourcePayload{*resourcePayloadModel}
				createCaseOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				createCaseOptionsModel.InvoiceNumber = core.StringPtr("testString")
				createCaseOptionsModel.SLACreditRequest = core.BoolPtr(false)
				createCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.CreateCase(createCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.CreateCase(createCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCase(createCaseOptions *CreateCaseOptions)`, func() {
		createCasePath := "/cases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCasePath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
				}))
			})
			It(`Invoke CreateCase successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))

				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				resourcePayloadModel.CRN = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsModel := new(casemanagementv1.CreateCaseOptions)
				createCaseOptionsModel.Type = core.StringPtr("technical")
				createCaseOptionsModel.Subject = core.StringPtr("testString")
				createCaseOptionsModel.Description = core.StringPtr("testString")
				createCaseOptionsModel.Severity = core.Int64Ptr(int64(1))
				createCaseOptionsModel.Eu = casePayloadEuModel
				createCaseOptionsModel.Offering = offeringModel
				createCaseOptionsModel.Resources = []casemanagementv1.ResourcePayload{*resourcePayloadModel}
				createCaseOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				createCaseOptionsModel.InvoiceNumber = core.StringPtr("testString")
				createCaseOptionsModel.SLACreditRequest = core.BoolPtr(false)
				createCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.CreateCaseWithContext(ctx, createCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.CreateCase(createCaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.CreateCaseWithContext(ctx, createCaseOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createCasePath))
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
				}))
			})
			It(`Invoke CreateCase successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.CreateCase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))

				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				resourcePayloadModel.CRN = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsModel := new(casemanagementv1.CreateCaseOptions)
				createCaseOptionsModel.Type = core.StringPtr("technical")
				createCaseOptionsModel.Subject = core.StringPtr("testString")
				createCaseOptionsModel.Description = core.StringPtr("testString")
				createCaseOptionsModel.Severity = core.Int64Ptr(int64(1))
				createCaseOptionsModel.Eu = casePayloadEuModel
				createCaseOptionsModel.Offering = offeringModel
				createCaseOptionsModel.Resources = []casemanagementv1.ResourcePayload{*resourcePayloadModel}
				createCaseOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				createCaseOptionsModel.InvoiceNumber = core.StringPtr("testString")
				createCaseOptionsModel.SLACreditRequest = core.BoolPtr(false)
				createCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.CreateCase(createCaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCase with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))

				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				resourcePayloadModel.CRN = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsModel := new(casemanagementv1.CreateCaseOptions)
				createCaseOptionsModel.Type = core.StringPtr("technical")
				createCaseOptionsModel.Subject = core.StringPtr("testString")
				createCaseOptionsModel.Description = core.StringPtr("testString")
				createCaseOptionsModel.Severity = core.Int64Ptr(int64(1))
				createCaseOptionsModel.Eu = casePayloadEuModel
				createCaseOptionsModel.Offering = offeringModel
				createCaseOptionsModel.Resources = []casemanagementv1.ResourcePayload{*resourcePayloadModel}
				createCaseOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				createCaseOptionsModel.InvoiceNumber = core.StringPtr("testString")
				createCaseOptionsModel.SLACreditRequest = core.BoolPtr(false)
				createCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.CreateCase(createCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCaseOptions model with no property values
				createCaseOptionsModelNew := new(casemanagementv1.CreateCaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.CreateCase(createCaseOptionsModelNew)
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
			It(`Invoke CreateCase successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))

				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				resourcePayloadModel.CRN = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsModel := new(casemanagementv1.CreateCaseOptions)
				createCaseOptionsModel.Type = core.StringPtr("technical")
				createCaseOptionsModel.Subject = core.StringPtr("testString")
				createCaseOptionsModel.Description = core.StringPtr("testString")
				createCaseOptionsModel.Severity = core.Int64Ptr(int64(1))
				createCaseOptionsModel.Eu = casePayloadEuModel
				createCaseOptionsModel.Offering = offeringModel
				createCaseOptionsModel.Resources = []casemanagementv1.ResourcePayload{*resourcePayloadModel}
				createCaseOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				createCaseOptionsModel.InvoiceNumber = core.StringPtr("testString")
				createCaseOptionsModel.SLACreditRequest = core.BoolPtr(false)
				createCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.CreateCase(createCaseOptionsModel)
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
	Describe(`GetCase(getCaseOptions *GetCaseOptions) - Operation response error`, func() {
		getCasePath := "/cases/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCasePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCase with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the GetCaseOptions model
				getCaseOptionsModel := new(casemanagementv1.GetCaseOptions)
				getCaseOptionsModel.CaseNumber = core.StringPtr("testString")
				getCaseOptionsModel.Fields = []string{"number"}
				getCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.GetCase(getCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.GetCase(getCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCase(getCaseOptions *GetCaseOptions)`, func() {
		getCasePath := "/cases/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCasePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
				}))
			})
			It(`Invoke GetCase successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetCaseOptions model
				getCaseOptionsModel := new(casemanagementv1.GetCaseOptions)
				getCaseOptionsModel.CaseNumber = core.StringPtr("testString")
				getCaseOptionsModel.Fields = []string{"number"}
				getCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.GetCaseWithContext(ctx, getCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.GetCase(getCaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.GetCaseWithContext(ctx, getCaseOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCasePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
				}))
			})
			It(`Invoke GetCase successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.GetCase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCaseOptions model
				getCaseOptionsModel := new(casemanagementv1.GetCaseOptions)
				getCaseOptionsModel.CaseNumber = core.StringPtr("testString")
				getCaseOptionsModel.Fields = []string{"number"}
				getCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.GetCase(getCaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCase with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the GetCaseOptions model
				getCaseOptionsModel := new(casemanagementv1.GetCaseOptions)
				getCaseOptionsModel.CaseNumber = core.StringPtr("testString")
				getCaseOptionsModel.Fields = []string{"number"}
				getCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.GetCase(getCaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCaseOptions model with no property values
				getCaseOptionsModelNew := new(casemanagementv1.GetCaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.GetCase(getCaseOptionsModelNew)
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
			It(`Invoke GetCase successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the GetCaseOptions model
				getCaseOptionsModel := new(casemanagementv1.GetCaseOptions)
				getCaseOptionsModel.CaseNumber = core.StringPtr("testString")
				getCaseOptionsModel.Fields = []string{"number"}
				getCaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.GetCase(getCaseOptionsModel)
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
	Describe(`UpdateCaseStatus(updateCaseStatusOptions *UpdateCaseStatusOptions) - Operation response error`, func() {
		updateCaseStatusPath := "/cases/testString/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCaseStatusPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCaseStatus with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("It was actually a mistake")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdateCaseStatusOptions model
				updateCaseStatusOptionsModel := new(casemanagementv1.UpdateCaseStatusOptions)
				updateCaseStatusOptionsModel.CaseNumber = core.StringPtr("testString")
				updateCaseStatusOptionsModel.StatusPayload = statusPayloadModel
				updateCaseStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.UpdateCaseStatus(updateCaseStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.UpdateCaseStatus(updateCaseStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCaseStatus(updateCaseStatusOptions *UpdateCaseStatusOptions)`, func() {
		updateCaseStatusPath := "/cases/testString/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCaseStatusPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
				}))
			})
			It(`Invoke UpdateCaseStatus successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("It was actually a mistake")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdateCaseStatusOptions model
				updateCaseStatusOptionsModel := new(casemanagementv1.UpdateCaseStatusOptions)
				updateCaseStatusOptionsModel.CaseNumber = core.StringPtr("testString")
				updateCaseStatusOptionsModel.StatusPayload = statusPayloadModel
				updateCaseStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.UpdateCaseStatusWithContext(ctx, updateCaseStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.UpdateCaseStatus(updateCaseStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.UpdateCaseStatusWithContext(ctx, updateCaseStatusOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateCaseStatusPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"number": "Number", "short_description": "ShortDescription", "description": "Description", "created_at": "CreatedAt", "created_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "updated_at": "UpdatedAt", "updated_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "contact_type": "Cloud Support Center", "contact": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}, "status": "Status", "severity": 8, "support_tier": "Free", "resolution": "Resolution", "close_notes": "CloseNotes", "eu": {"support": false, "data_center": "DataCenter"}, "watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}], "offering": {"name": "Name", "type": {"group": "crn_service_name", "key": "Key", "kind": "Kind", "id": "ID"}}, "resources": [{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}], "comments": [{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}]}`)
				}))
			})
			It(`Invoke UpdateCaseStatus successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.UpdateCaseStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("It was actually a mistake")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdateCaseStatusOptions model
				updateCaseStatusOptionsModel := new(casemanagementv1.UpdateCaseStatusOptions)
				updateCaseStatusOptionsModel.CaseNumber = core.StringPtr("testString")
				updateCaseStatusOptionsModel.StatusPayload = statusPayloadModel
				updateCaseStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.UpdateCaseStatus(updateCaseStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCaseStatus with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("It was actually a mistake")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdateCaseStatusOptions model
				updateCaseStatusOptionsModel := new(casemanagementv1.UpdateCaseStatusOptions)
				updateCaseStatusOptionsModel.CaseNumber = core.StringPtr("testString")
				updateCaseStatusOptionsModel.StatusPayload = statusPayloadModel
				updateCaseStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.UpdateCaseStatus(updateCaseStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCaseStatusOptions model with no property values
				updateCaseStatusOptionsModelNew := new(casemanagementv1.UpdateCaseStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.UpdateCaseStatus(updateCaseStatusOptionsModelNew)
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
			It(`Invoke UpdateCaseStatus successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("It was actually a mistake")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdateCaseStatusOptions model
				updateCaseStatusOptionsModel := new(casemanagementv1.UpdateCaseStatusOptions)
				updateCaseStatusOptionsModel.CaseNumber = core.StringPtr("testString")
				updateCaseStatusOptionsModel.StatusPayload = statusPayloadModel
				updateCaseStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.UpdateCaseStatus(updateCaseStatusOptionsModel)
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
	Describe(`AddComment(addCommentOptions *AddCommentOptions) - Operation response error`, func() {
		addCommentPath := "/cases/testString/comments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addCommentPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddComment with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the AddCommentOptions model
				addCommentOptionsModel := new(casemanagementv1.AddCommentOptions)
				addCommentOptionsModel.CaseNumber = core.StringPtr("testString")
				addCommentOptionsModel.Comment = core.StringPtr("This is a test comment")
				addCommentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.AddComment(addCommentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.AddComment(addCommentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddComment(addCommentOptions *AddCommentOptions)`, func() {
		addCommentPath := "/cases/testString/comments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addCommentPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}`)
				}))
			})
			It(`Invoke AddComment successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the AddCommentOptions model
				addCommentOptionsModel := new(casemanagementv1.AddCommentOptions)
				addCommentOptionsModel.CaseNumber = core.StringPtr("testString")
				addCommentOptionsModel.Comment = core.StringPtr("This is a test comment")
				addCommentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.AddCommentWithContext(ctx, addCommentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.AddComment(addCommentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.AddCommentWithContext(ctx, addCommentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addCommentPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"value": "Value", "added_at": "AddedAt", "added_by": {"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}}`)
				}))
			})
			It(`Invoke AddComment successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.AddComment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddCommentOptions model
				addCommentOptionsModel := new(casemanagementv1.AddCommentOptions)
				addCommentOptionsModel.CaseNumber = core.StringPtr("testString")
				addCommentOptionsModel.Comment = core.StringPtr("This is a test comment")
				addCommentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.AddComment(addCommentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddComment with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the AddCommentOptions model
				addCommentOptionsModel := new(casemanagementv1.AddCommentOptions)
				addCommentOptionsModel.CaseNumber = core.StringPtr("testString")
				addCommentOptionsModel.Comment = core.StringPtr("This is a test comment")
				addCommentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.AddComment(addCommentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddCommentOptions model with no property values
				addCommentOptionsModelNew := new(casemanagementv1.AddCommentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.AddComment(addCommentOptionsModelNew)
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
			It(`Invoke AddComment successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the AddCommentOptions model
				addCommentOptionsModel := new(casemanagementv1.AddCommentOptions)
				addCommentOptionsModel.CaseNumber = core.StringPtr("testString")
				addCommentOptionsModel.Comment = core.StringPtr("This is a test comment")
				addCommentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.AddComment(addCommentOptionsModel)
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
	Describe(`AddWatchlist(addWatchlistOptions *AddWatchlistOptions) - Operation response error`, func() {
		addWatchlistPath := "/cases/testString/watchlist"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addWatchlistPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddWatchlist with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the AddWatchlistOptions model
				addWatchlistOptionsModel := new(casemanagementv1.AddWatchlistOptions)
				addWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				addWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				addWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.AddWatchlist(addWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.AddWatchlist(addWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddWatchlist(addWatchlistOptions *AddWatchlistOptions)`, func() {
		addWatchlistPath := "/cases/testString/watchlist"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addWatchlistPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"added": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "failed": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}]}`)
				}))
			})
			It(`Invoke AddWatchlist successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the AddWatchlistOptions model
				addWatchlistOptionsModel := new(casemanagementv1.AddWatchlistOptions)
				addWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				addWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				addWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.AddWatchlistWithContext(ctx, addWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.AddWatchlist(addWatchlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.AddWatchlistWithContext(ctx, addWatchlistOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addWatchlistPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"added": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}], "failed": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}]}`)
				}))
			})
			It(`Invoke AddWatchlist successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.AddWatchlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the AddWatchlistOptions model
				addWatchlistOptionsModel := new(casemanagementv1.AddWatchlistOptions)
				addWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				addWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				addWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.AddWatchlist(addWatchlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddWatchlist with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the AddWatchlistOptions model
				addWatchlistOptionsModel := new(casemanagementv1.AddWatchlistOptions)
				addWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				addWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				addWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.AddWatchlist(addWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddWatchlistOptions model with no property values
				addWatchlistOptionsModelNew := new(casemanagementv1.AddWatchlistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.AddWatchlist(addWatchlistOptionsModelNew)
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
			It(`Invoke AddWatchlist successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the AddWatchlistOptions model
				addWatchlistOptionsModel := new(casemanagementv1.AddWatchlistOptions)
				addWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				addWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				addWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.AddWatchlist(addWatchlistOptionsModel)
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
	Describe(`RemoveWatchlist(removeWatchlistOptions *RemoveWatchlistOptions) - Operation response error`, func() {
		removeWatchlistPath := "/cases/testString/watchlist"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeWatchlistPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveWatchlist with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the RemoveWatchlistOptions model
				removeWatchlistOptionsModel := new(casemanagementv1.RemoveWatchlistOptions)
				removeWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				removeWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				removeWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.RemoveWatchlist(removeWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.RemoveWatchlist(removeWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveWatchlist(removeWatchlistOptions *RemoveWatchlistOptions)`, func() {
		removeWatchlistPath := "/cases/testString/watchlist"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeWatchlistPath))
					Expect(req.Method).To(Equal("DELETE"))

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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}]}`)
				}))
			})
			It(`Invoke RemoveWatchlist successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the RemoveWatchlistOptions model
				removeWatchlistOptionsModel := new(casemanagementv1.RemoveWatchlistOptions)
				removeWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				removeWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				removeWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.RemoveWatchlistWithContext(ctx, removeWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.RemoveWatchlist(removeWatchlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.RemoveWatchlistWithContext(ctx, removeWatchlistOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(removeWatchlistPath))
					Expect(req.Method).To(Equal("DELETE"))

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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"watchlist": [{"name": "Name", "realm": "IBMid", "user_id": "abc@ibm.com"}]}`)
				}))
			})
			It(`Invoke RemoveWatchlist successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.RemoveWatchlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the RemoveWatchlistOptions model
				removeWatchlistOptionsModel := new(casemanagementv1.RemoveWatchlistOptions)
				removeWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				removeWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				removeWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.RemoveWatchlist(removeWatchlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RemoveWatchlist with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the RemoveWatchlistOptions model
				removeWatchlistOptionsModel := new(casemanagementv1.RemoveWatchlistOptions)
				removeWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				removeWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				removeWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.RemoveWatchlist(removeWatchlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveWatchlistOptions model with no property values
				removeWatchlistOptionsModelNew := new(casemanagementv1.RemoveWatchlistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.RemoveWatchlist(removeWatchlistOptionsModelNew)
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
			It(`Invoke RemoveWatchlist successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")

				// Construct an instance of the RemoveWatchlistOptions model
				removeWatchlistOptionsModel := new(casemanagementv1.RemoveWatchlistOptions)
				removeWatchlistOptionsModel.CaseNumber = core.StringPtr("testString")
				removeWatchlistOptionsModel.Watchlist = []casemanagementv1.User{*userModel}
				removeWatchlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.RemoveWatchlist(removeWatchlistOptionsModel)
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
	Describe(`AddResource(addResourceOptions *AddResourceOptions) - Operation response error`, func() {
		addResourcePath := "/cases/testString/resources"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addResourcePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddResource with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the AddResourceOptions model
				addResourceOptionsModel := new(casemanagementv1.AddResourceOptions)
				addResourceOptionsModel.CaseNumber = core.StringPtr("testString")
				addResourceOptionsModel.CRN = core.StringPtr("testString")
				addResourceOptionsModel.Type = core.StringPtr("testString")
				addResourceOptionsModel.ID = core.Float64Ptr(float64(72.5))
				addResourceOptionsModel.Note = core.StringPtr("testString")
				addResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.AddResource(addResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.AddResource(addResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddResource(addResourceOptions *AddResourceOptions)`, func() {
		addResourcePath := "/cases/testString/resources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addResourcePath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}`)
				}))
			})
			It(`Invoke AddResource successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the AddResourceOptions model
				addResourceOptionsModel := new(casemanagementv1.AddResourceOptions)
				addResourceOptionsModel.CaseNumber = core.StringPtr("testString")
				addResourceOptionsModel.CRN = core.StringPtr("testString")
				addResourceOptionsModel.Type = core.StringPtr("testString")
				addResourceOptionsModel.ID = core.Float64Ptr(float64(72.5))
				addResourceOptionsModel.Note = core.StringPtr("testString")
				addResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.AddResourceWithContext(ctx, addResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.AddResource(addResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.AddResourceWithContext(ctx, addResourceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addResourcePath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "CRN", "name": "Name", "type": "Type", "url": "URL", "note": "Note"}`)
				}))
			})
			It(`Invoke AddResource successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.AddResource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddResourceOptions model
				addResourceOptionsModel := new(casemanagementv1.AddResourceOptions)
				addResourceOptionsModel.CaseNumber = core.StringPtr("testString")
				addResourceOptionsModel.CRN = core.StringPtr("testString")
				addResourceOptionsModel.Type = core.StringPtr("testString")
				addResourceOptionsModel.ID = core.Float64Ptr(float64(72.5))
				addResourceOptionsModel.Note = core.StringPtr("testString")
				addResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.AddResource(addResourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddResource with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the AddResourceOptions model
				addResourceOptionsModel := new(casemanagementv1.AddResourceOptions)
				addResourceOptionsModel.CaseNumber = core.StringPtr("testString")
				addResourceOptionsModel.CRN = core.StringPtr("testString")
				addResourceOptionsModel.Type = core.StringPtr("testString")
				addResourceOptionsModel.ID = core.Float64Ptr(float64(72.5))
				addResourceOptionsModel.Note = core.StringPtr("testString")
				addResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.AddResource(addResourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddResourceOptions model with no property values
				addResourceOptionsModelNew := new(casemanagementv1.AddResourceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.AddResource(addResourceOptionsModelNew)
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
			It(`Invoke AddResource successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the AddResourceOptions model
				addResourceOptionsModel := new(casemanagementv1.AddResourceOptions)
				addResourceOptionsModel.CaseNumber = core.StringPtr("testString")
				addResourceOptionsModel.CRN = core.StringPtr("testString")
				addResourceOptionsModel.Type = core.StringPtr("testString")
				addResourceOptionsModel.ID = core.Float64Ptr(float64(72.5))
				addResourceOptionsModel.Note = core.StringPtr("testString")
				addResourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.AddResource(addResourceOptionsModel)
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
	Describe(`UploadFile(uploadFileOptions *UploadFileOptions) - Operation response error`, func() {
		uploadFilePath := "/cases/testString/attachments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadFilePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UploadFile with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the UploadFileOptions model
				uploadFileOptionsModel := new(casemanagementv1.UploadFileOptions)
				uploadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				uploadFileOptionsModel.File = []casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt")}}
				uploadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.UploadFile(uploadFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.UploadFile(uploadFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UploadFile(uploadFileOptions *UploadFileOptions)`, func() {
		uploadFilePath := "/cases/testString/attachments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadFilePath))
					Expect(req.Method).To(Equal("PUT"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}`)
				}))
			})
			It(`Invoke UploadFile successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the UploadFileOptions model
				uploadFileOptionsModel := new(casemanagementv1.UploadFileOptions)
				uploadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				uploadFileOptionsModel.File = []casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt")}}
				uploadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.UploadFileWithContext(ctx, uploadFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.UploadFile(uploadFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.UploadFileWithContext(ctx, uploadFileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(uploadFilePath))
					Expect(req.Method).To(Equal("PUT"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}`)
				}))
			})
			It(`Invoke UploadFile successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.UploadFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UploadFileOptions model
				uploadFileOptionsModel := new(casemanagementv1.UploadFileOptions)
				uploadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				uploadFileOptionsModel.File = []casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt")}}
				uploadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.UploadFile(uploadFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UploadFile with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the UploadFileOptions model
				uploadFileOptionsModel := new(casemanagementv1.UploadFileOptions)
				uploadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				uploadFileOptionsModel.File = []casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt")}}
				uploadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.UploadFile(uploadFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UploadFileOptions model with no property values
				uploadFileOptionsModelNew := new(casemanagementv1.UploadFileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.UploadFile(uploadFileOptionsModelNew)
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
			It(`Invoke UploadFile successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the UploadFileOptions model
				uploadFileOptionsModel := new(casemanagementv1.UploadFileOptions)
				uploadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				uploadFileOptionsModel.File = []casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt")}}
				uploadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.UploadFile(uploadFileOptionsModel)
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
	Describe(`DownloadFile(downloadFileOptions *DownloadFileOptions)`, func() {
		downloadFilePath := "/cases/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(downloadFilePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/octet-stream")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DownloadFile successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the DownloadFileOptions model
				downloadFileOptionsModel := new(casemanagementv1.DownloadFileOptions)
				downloadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				downloadFileOptionsModel.FileID = core.StringPtr("testString")
				downloadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.DownloadFileWithContext(ctx, downloadFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.DownloadFile(downloadFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.DownloadFileWithContext(ctx, downloadFileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(downloadFilePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/octet-stream")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke DownloadFile successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.DownloadFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DownloadFileOptions model
				downloadFileOptionsModel := new(casemanagementv1.DownloadFileOptions)
				downloadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				downloadFileOptionsModel.FileID = core.StringPtr("testString")
				downloadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.DownloadFile(downloadFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DownloadFile with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the DownloadFileOptions model
				downloadFileOptionsModel := new(casemanagementv1.DownloadFileOptions)
				downloadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				downloadFileOptionsModel.FileID = core.StringPtr("testString")
				downloadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.DownloadFile(downloadFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DownloadFileOptions model with no property values
				downloadFileOptionsModelNew := new(casemanagementv1.DownloadFileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.DownloadFile(downloadFileOptionsModelNew)
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
			It(`Invoke DownloadFile successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the DownloadFileOptions model
				downloadFileOptionsModel := new(casemanagementv1.DownloadFileOptions)
				downloadFileOptionsModel.CaseNumber = core.StringPtr("testString")
				downloadFileOptionsModel.FileID = core.StringPtr("testString")
				downloadFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.DownloadFile(downloadFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteFile(deleteFileOptions *DeleteFileOptions) - Operation response error`, func() {
		deleteFilePath := "/cases/testString/attachments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFilePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteFile with error: Operation response processing error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteFileOptions model
				deleteFileOptionsModel := new(casemanagementv1.DeleteFileOptions)
				deleteFileOptionsModel.CaseNumber = core.StringPtr("testString")
				deleteFileOptionsModel.FileID = core.StringPtr("testString")
				deleteFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := caseManagementService.DeleteFile(deleteFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				caseManagementService.EnableRetries(0, 0)
				result, response, operationErr = caseManagementService.DeleteFile(deleteFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteFile(deleteFileOptions *DeleteFileOptions)`, func() {
		deleteFilePath := "/cases/testString/attachments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFilePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}]}`)
				}))
			})
			It(`Invoke DeleteFile successfully with retries`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())
				caseManagementService.EnableRetries(0, 0)

				// Construct an instance of the DeleteFileOptions model
				deleteFileOptionsModel := new(casemanagementv1.DeleteFileOptions)
				deleteFileOptionsModel.CaseNumber = core.StringPtr("testString")
				deleteFileOptionsModel.FileID = core.StringPtr("testString")
				deleteFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := caseManagementService.DeleteFileWithContext(ctx, deleteFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				caseManagementService.DisableRetries()
				result, response, operationErr := caseManagementService.DeleteFile(deleteFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = caseManagementService.DeleteFileWithContext(ctx, deleteFileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteFilePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"attachments": [{"id": "ID", "filename": "Filename", "size_in_bytes": 11, "created_at": "CreatedAt", "url": "URL"}]}`)
				}))
			})
			It(`Invoke DeleteFile successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := caseManagementService.DeleteFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteFileOptions model
				deleteFileOptionsModel := new(casemanagementv1.DeleteFileOptions)
				deleteFileOptionsModel.CaseNumber = core.StringPtr("testString")
				deleteFileOptionsModel.FileID = core.StringPtr("testString")
				deleteFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = caseManagementService.DeleteFile(deleteFileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteFile with error: Operation validation and request error`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteFileOptions model
				deleteFileOptionsModel := new(casemanagementv1.DeleteFileOptions)
				deleteFileOptionsModel.CaseNumber = core.StringPtr("testString")
				deleteFileOptionsModel.FileID = core.StringPtr("testString")
				deleteFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := caseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := caseManagementService.DeleteFile(deleteFileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteFileOptions model with no property values
				deleteFileOptionsModelNew := new(casemanagementv1.DeleteFileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = caseManagementService.DeleteFile(deleteFileOptionsModelNew)
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
			It(`Invoke DeleteFile successfully`, func() {
				caseManagementService, serviceErr := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(caseManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteFileOptions model
				deleteFileOptionsModel := new(casemanagementv1.DeleteFileOptions)
				deleteFileOptionsModel.CaseNumber = core.StringPtr("testString")
				deleteFileOptionsModel.FileID = core.StringPtr("testString")
				deleteFileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := caseManagementService.DeleteFile(deleteFileOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			caseManagementService, _ := casemanagementv1.NewCaseManagementV1(&casemanagementv1.CaseManagementV1Options{
				URL:           "http://casemanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddCommentOptions successfully`, func() {
				// Construct an instance of the AddCommentOptions model
				caseNumber := "testString"
				addCommentOptionsComment := "This is a test comment"
				addCommentOptionsModel := caseManagementService.NewAddCommentOptions(caseNumber, addCommentOptionsComment)
				addCommentOptionsModel.SetCaseNumber("testString")
				addCommentOptionsModel.SetComment("This is a test comment")
				addCommentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addCommentOptionsModel).ToNot(BeNil())
				Expect(addCommentOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(addCommentOptionsModel.Comment).To(Equal(core.StringPtr("This is a test comment")))
				Expect(addCommentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddResourceOptions successfully`, func() {
				// Construct an instance of the AddResourceOptions model
				caseNumber := "testString"
				addResourceOptionsModel := caseManagementService.NewAddResourceOptions(caseNumber)
				addResourceOptionsModel.SetCaseNumber("testString")
				addResourceOptionsModel.SetCRN("testString")
				addResourceOptionsModel.SetType("testString")
				addResourceOptionsModel.SetID(float64(72.5))
				addResourceOptionsModel.SetNote("testString")
				addResourceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addResourceOptionsModel).ToNot(BeNil())
				Expect(addResourceOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(addResourceOptionsModel.CRN).To(Equal(core.StringPtr("testString")))
				Expect(addResourceOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(addResourceOptionsModel.ID).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(addResourceOptionsModel.Note).To(Equal(core.StringPtr("testString")))
				Expect(addResourceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddWatchlistOptions successfully`, func() {
				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				Expect(userModel).ToNot(BeNil())
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")
				Expect(userModel.Realm).To(Equal(core.StringPtr("IBMid")))
				Expect(userModel.UserID).To(Equal(core.StringPtr("abc@ibm.com")))

				// Construct an instance of the AddWatchlistOptions model
				caseNumber := "testString"
				addWatchlistOptionsModel := caseManagementService.NewAddWatchlistOptions(caseNumber)
				addWatchlistOptionsModel.SetCaseNumber("testString")
				addWatchlistOptionsModel.SetWatchlist([]casemanagementv1.User{*userModel})
				addWatchlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addWatchlistOptionsModel).ToNot(BeNil())
				Expect(addWatchlistOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(addWatchlistOptionsModel.Watchlist).To(Equal([]casemanagementv1.User{*userModel}))
				Expect(addWatchlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCaseOptions successfully`, func() {
				// Construct an instance of the CasePayloadEu model
				casePayloadEuModel := new(casemanagementv1.CasePayloadEu)
				Expect(casePayloadEuModel).ToNot(BeNil())
				casePayloadEuModel.Supported = core.BoolPtr(true)
				casePayloadEuModel.DataCenter = core.Int64Ptr(int64(38))
				Expect(casePayloadEuModel.Supported).To(Equal(core.BoolPtr(true)))
				Expect(casePayloadEuModel.DataCenter).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the OfferingType model
				offeringTypeModel := new(casemanagementv1.OfferingType)
				Expect(offeringTypeModel).ToNot(BeNil())
				offeringTypeModel.Group = core.StringPtr("crn_service_name")
				offeringTypeModel.Key = core.StringPtr("testString")
				offeringTypeModel.Kind = core.StringPtr("testString")
				offeringTypeModel.ID = core.StringPtr("testString")
				Expect(offeringTypeModel.Group).To(Equal(core.StringPtr("crn_service_name")))
				Expect(offeringTypeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(offeringTypeModel.Kind).To(Equal(core.StringPtr("testString")))
				Expect(offeringTypeModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Offering model
				offeringModel := new(casemanagementv1.Offering)
				Expect(offeringModel).ToNot(BeNil())
				offeringModel.Name = core.StringPtr("testString")
				offeringModel.Type = offeringTypeModel
				Expect(offeringModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(offeringModel.Type).To(Equal(offeringTypeModel))

				// Construct an instance of the ResourcePayload model
				resourcePayloadModel := new(casemanagementv1.ResourcePayload)
				Expect(resourcePayloadModel).ToNot(BeNil())
				resourcePayloadModel.CRN = core.StringPtr("testString")
				resourcePayloadModel.Type = core.StringPtr("testString")
				resourcePayloadModel.ID = core.Float64Ptr(float64(72.5))
				resourcePayloadModel.Note = core.StringPtr("testString")
				Expect(resourcePayloadModel.CRN).To(Equal(core.StringPtr("testString")))
				Expect(resourcePayloadModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(resourcePayloadModel.ID).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(resourcePayloadModel.Note).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				Expect(userModel).ToNot(BeNil())
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")
				Expect(userModel.Realm).To(Equal(core.StringPtr("IBMid")))
				Expect(userModel.UserID).To(Equal(core.StringPtr("abc@ibm.com")))

				// Construct an instance of the CreateCaseOptions model
				createCaseOptionsType := "technical"
				createCaseOptionsSubject := "testString"
				createCaseOptionsDescription := "testString"
				createCaseOptionsModel := caseManagementService.NewCreateCaseOptions(createCaseOptionsType, createCaseOptionsSubject, createCaseOptionsDescription)
				createCaseOptionsModel.SetType("technical")
				createCaseOptionsModel.SetSubject("testString")
				createCaseOptionsModel.SetDescription("testString")
				createCaseOptionsModel.SetSeverity(int64(1))
				createCaseOptionsModel.SetEu(casePayloadEuModel)
				createCaseOptionsModel.SetOffering(offeringModel)
				createCaseOptionsModel.SetResources([]casemanagementv1.ResourcePayload{*resourcePayloadModel})
				createCaseOptionsModel.SetWatchlist([]casemanagementv1.User{*userModel})
				createCaseOptionsModel.SetInvoiceNumber("testString")
				createCaseOptionsModel.SetSLACreditRequest(false)
				createCaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCaseOptionsModel).ToNot(BeNil())
				Expect(createCaseOptionsModel.Type).To(Equal(core.StringPtr("technical")))
				Expect(createCaseOptionsModel.Subject).To(Equal(core.StringPtr("testString")))
				Expect(createCaseOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createCaseOptionsModel.Severity).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createCaseOptionsModel.Eu).To(Equal(casePayloadEuModel))
				Expect(createCaseOptionsModel.Offering).To(Equal(offeringModel))
				Expect(createCaseOptionsModel.Resources).To(Equal([]casemanagementv1.ResourcePayload{*resourcePayloadModel}))
				Expect(createCaseOptionsModel.Watchlist).To(Equal([]casemanagementv1.User{*userModel}))
				Expect(createCaseOptionsModel.InvoiceNumber).To(Equal(core.StringPtr("testString")))
				Expect(createCaseOptionsModel.SLACreditRequest).To(Equal(core.BoolPtr(false)))
				Expect(createCaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFileOptions successfully`, func() {
				// Construct an instance of the DeleteFileOptions model
				caseNumber := "testString"
				fileID := "testString"
				deleteFileOptionsModel := caseManagementService.NewDeleteFileOptions(caseNumber, fileID)
				deleteFileOptionsModel.SetCaseNumber("testString")
				deleteFileOptionsModel.SetFileID("testString")
				deleteFileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFileOptionsModel).ToNot(BeNil())
				Expect(deleteFileOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(deleteFileOptionsModel.FileID).To(Equal(core.StringPtr("testString")))
				Expect(deleteFileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDownloadFileOptions successfully`, func() {
				// Construct an instance of the DownloadFileOptions model
				caseNumber := "testString"
				fileID := "testString"
				downloadFileOptionsModel := caseManagementService.NewDownloadFileOptions(caseNumber, fileID)
				downloadFileOptionsModel.SetCaseNumber("testString")
				downloadFileOptionsModel.SetFileID("testString")
				downloadFileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(downloadFileOptionsModel).ToNot(BeNil())
				Expect(downloadFileOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(downloadFileOptionsModel.FileID).To(Equal(core.StringPtr("testString")))
				Expect(downloadFileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFileWithMetadata successfully`, func() {
				data := CreateMockReader("This is a mock file.")
				_model, err := caseManagementService.NewFileWithMetadata(data)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetCaseOptions successfully`, func() {
				// Construct an instance of the GetCaseOptions model
				caseNumber := "testString"
				getCaseOptionsModel := caseManagementService.NewGetCaseOptions(caseNumber)
				getCaseOptionsModel.SetCaseNumber("testString")
				getCaseOptionsModel.SetFields([]string{"number"})
				getCaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCaseOptionsModel).ToNot(BeNil())
				Expect(getCaseOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(getCaseOptionsModel.Fields).To(Equal([]string{"number"}))
				Expect(getCaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCasesOptions successfully`, func() {
				// Construct an instance of the GetCasesOptions model
				getCasesOptionsModel := caseManagementService.NewGetCasesOptions()
				getCasesOptionsModel.SetOffset(int64(38))
				getCasesOptionsModel.SetLimit(int64(10))
				getCasesOptionsModel.SetSearch("testString")
				getCasesOptionsModel.SetSort("number")
				getCasesOptionsModel.SetStatus([]string{"new"})
				getCasesOptionsModel.SetFields([]string{"number"})
				getCasesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCasesOptionsModel).ToNot(BeNil())
				Expect(getCasesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getCasesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(getCasesOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(getCasesOptionsModel.Sort).To(Equal(core.StringPtr("number")))
				Expect(getCasesOptionsModel.Status).To(Equal([]string{"new"}))
				Expect(getCasesOptionsModel.Fields).To(Equal([]string{"number"}))
				Expect(getCasesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewOffering successfully`, func() {
				name := "testString"
				var typeVar *casemanagementv1.OfferingType = nil
				_, err := caseManagementService.NewOffering(name, typeVar)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewOfferingType successfully`, func() {
				group := "crn_service_name"
				key := "testString"
				_model, err := caseManagementService.NewOfferingType(group, key)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRemoveWatchlistOptions successfully`, func() {
				// Construct an instance of the User model
				userModel := new(casemanagementv1.User)
				Expect(userModel).ToNot(BeNil())
				userModel.Realm = core.StringPtr("IBMid")
				userModel.UserID = core.StringPtr("abc@ibm.com")
				Expect(userModel.Realm).To(Equal(core.StringPtr("IBMid")))
				Expect(userModel.UserID).To(Equal(core.StringPtr("abc@ibm.com")))

				// Construct an instance of the RemoveWatchlistOptions model
				caseNumber := "testString"
				removeWatchlistOptionsModel := caseManagementService.NewRemoveWatchlistOptions(caseNumber)
				removeWatchlistOptionsModel.SetCaseNumber("testString")
				removeWatchlistOptionsModel.SetWatchlist([]casemanagementv1.User{*userModel})
				removeWatchlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeWatchlistOptionsModel).ToNot(BeNil())
				Expect(removeWatchlistOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(removeWatchlistOptionsModel.Watchlist).To(Equal([]casemanagementv1.User{*userModel}))
				Expect(removeWatchlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCaseStatusOptions successfully`, func() {
				// Construct an instance of the ResolvePayload model
				statusPayloadModel := new(casemanagementv1.ResolvePayload)
				Expect(statusPayloadModel).ToNot(BeNil())
				statusPayloadModel.Action = core.StringPtr("resolve")
				statusPayloadModel.Comment = core.StringPtr("testString")
				statusPayloadModel.ResolutionCode = core.Int64Ptr(int64(1))
				Expect(statusPayloadModel.Action).To(Equal(core.StringPtr("resolve")))
				Expect(statusPayloadModel.Comment).To(Equal(core.StringPtr("testString")))
				Expect(statusPayloadModel.ResolutionCode).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the UpdateCaseStatusOptions model
				caseNumber := "testString"
				var statusPayload casemanagementv1.StatusPayloadIntf = nil
				updateCaseStatusOptionsModel := caseManagementService.NewUpdateCaseStatusOptions(caseNumber, statusPayload)
				updateCaseStatusOptionsModel.SetCaseNumber("testString")
				updateCaseStatusOptionsModel.SetStatusPayload(statusPayloadModel)
				updateCaseStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCaseStatusOptionsModel).ToNot(BeNil())
				Expect(updateCaseStatusOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(updateCaseStatusOptionsModel.StatusPayload).To(Equal(statusPayloadModel))
				Expect(updateCaseStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadFileOptions successfully`, func() {
				// Construct an instance of the FileWithMetadata model
				fileWithMetadataModel := new(casemanagementv1.FileWithMetadata)
				Expect(fileWithMetadataModel).ToNot(BeNil())
				fileWithMetadataModel.Data = CreateMockReader("This is a mock file.")
				fileWithMetadataModel.Filename = core.StringPtr("testString")
				fileWithMetadataModel.ContentType = core.StringPtr("testString")
				Expect(fileWithMetadataModel.Data).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(fileWithMetadataModel.Filename).To(Equal(core.StringPtr("testString")))
				Expect(fileWithMetadataModel.ContentType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UploadFileOptions model
				caseNumber := "testString"
				file := []casemanagementv1.FileWithMetadata{}
				uploadFileOptionsModel := caseManagementService.NewUploadFileOptions(caseNumber, file)
				uploadFileOptionsModel.SetCaseNumber("testString")
				uploadFileOptionsModel.SetFile([]casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt")}})
				uploadFileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadFileOptionsModel).ToNot(BeNil())
				Expect(uploadFileOptionsModel.CaseNumber).To(Equal(core.StringPtr("testString")))
				Expect(uploadFileOptionsModel.File).To(Equal([]casemanagementv1.FileWithMetadata{casemanagementv1.FileWithMetadata{Data: CreateMockReader("This is a mock file."), Filename: core.StringPtr("mockfilename.txt")}}))
				Expect(uploadFileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUser successfully`, func() {
				realm := "IBMid"
				userID := "abc@ibm.com"
				_model, err := caseManagementService.NewUser(realm, userID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAcceptPayload successfully`, func() {
				action := "accept"
				_model, err := caseManagementService.NewAcceptPayload(action)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResolvePayload successfully`, func() {
				action := "resolve"
				resolutionCode := int64(1)
				_model, err := caseManagementService.NewResolvePayload(action, resolutionCode)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUnresolvePayload successfully`, func() {
				action := "unresolve"
				comment := "testString"
				_model, err := caseManagementService.NewUnresolvePayload(action, comment)
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
