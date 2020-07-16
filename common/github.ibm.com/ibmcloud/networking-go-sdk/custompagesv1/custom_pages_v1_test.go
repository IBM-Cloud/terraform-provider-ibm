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

package custompagesv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/custompagesv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`CustomPagesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
				URL:            "https://custompagesv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CUSTOM_PAGES_URL":       "https://custompagesv1/api",
				"CUSTOM_PAGES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CUSTOM_PAGES_URL":       "https://custompagesv1/api",
				"CUSTOM_PAGES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CUSTOM_PAGES_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := custompagesv1.NewCustomPagesV1UsingExternalConfig(&custompagesv1.CustomPagesV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListInstanceCustomPages(listInstanceCustomPagesOptions *ListInstanceCustomPagesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listInstanceCustomPagesPath := "/v1/testString/custom_pages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listInstanceCustomPagesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListInstanceCustomPages with error: Operation response processing error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListInstanceCustomPagesOptions model
				listInstanceCustomPagesOptionsModel := new(custompagesv1.ListInstanceCustomPagesOptions)
				listInstanceCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListInstanceCustomPages(listInstanceCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListInstanceCustomPages(listInstanceCustomPagesOptions *ListInstanceCustomPagesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listInstanceCustomPagesPath := "/v1/testString/custom_pages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listInstanceCustomPagesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}], "result_info": {"page": 1, "per_page": 20, "total_pages": 1, "count": 10, "total_count": 10}}`)
				}))
			})
			It(`Invoke ListInstanceCustomPages successfully`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListInstanceCustomPages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceCustomPagesOptions model
				listInstanceCustomPagesOptionsModel := new(custompagesv1.ListInstanceCustomPagesOptions)
				listInstanceCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListInstanceCustomPages(listInstanceCustomPagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListInstanceCustomPages with error: Operation request error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListInstanceCustomPagesOptions model
				listInstanceCustomPagesOptionsModel := new(custompagesv1.ListInstanceCustomPagesOptions)
				listInstanceCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListInstanceCustomPages(listInstanceCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInstanceCustomPage(getInstanceCustomPageOptions *GetInstanceCustomPageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getInstanceCustomPagePath := "/v1/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getInstanceCustomPagePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInstanceCustomPage with error: Operation response processing error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetInstanceCustomPageOptions model
				getInstanceCustomPageOptionsModel := new(custompagesv1.GetInstanceCustomPageOptions)
				getInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetInstanceCustomPage(getInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetInstanceCustomPage(getInstanceCustomPageOptions *GetInstanceCustomPageOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getInstanceCustomPagePath := "/v1/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getInstanceCustomPagePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}}`)
				}))
			})
			It(`Invoke GetInstanceCustomPage successfully`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetInstanceCustomPage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInstanceCustomPageOptions model
				getInstanceCustomPageOptionsModel := new(custompagesv1.GetInstanceCustomPageOptions)
				getInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetInstanceCustomPage(getInstanceCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetInstanceCustomPage with error: Operation validation and request error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetInstanceCustomPageOptions model
				getInstanceCustomPageOptionsModel := new(custompagesv1.GetInstanceCustomPageOptions)
				getInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetInstanceCustomPage(getInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetInstanceCustomPageOptions model with no property values
				getInstanceCustomPageOptionsModelNew := new(custompagesv1.GetInstanceCustomPageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetInstanceCustomPage(getInstanceCustomPageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateInstanceCustomPage(updateInstanceCustomPageOptions *UpdateInstanceCustomPageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateInstanceCustomPagePath := "/v1/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateInstanceCustomPagePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateInstanceCustomPage with error: Operation response processing error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateInstanceCustomPageOptions model
				updateInstanceCustomPageOptionsModel := new(custompagesv1.UpdateInstanceCustomPageOptions)
				updateInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateInstanceCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateInstanceCustomPageOptionsModel.State = core.StringPtr("customized")
				updateInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateInstanceCustomPage(updateInstanceCustomPageOptions *UpdateInstanceCustomPageOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateInstanceCustomPagePath := "/v1/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateInstanceCustomPagePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}}`)
				}))
			})
			It(`Invoke UpdateInstanceCustomPage successfully`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateInstanceCustomPage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateInstanceCustomPageOptions model
				updateInstanceCustomPageOptionsModel := new(custompagesv1.UpdateInstanceCustomPageOptions)
				updateInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateInstanceCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateInstanceCustomPageOptionsModel.State = core.StringPtr("customized")
				updateInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateInstanceCustomPage with error: Operation validation and request error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateInstanceCustomPageOptions model
				updateInstanceCustomPageOptionsModel := new(custompagesv1.UpdateInstanceCustomPageOptions)
				updateInstanceCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateInstanceCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateInstanceCustomPageOptionsModel.State = core.StringPtr("customized")
				updateInstanceCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateInstanceCustomPageOptions model with no property values
				updateInstanceCustomPageOptionsModelNew := new(custompagesv1.UpdateInstanceCustomPageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateInstanceCustomPage(updateInstanceCustomPageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListZoneCustomPages(listZoneCustomPagesOptions *ListZoneCustomPagesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listZoneCustomPagesPath := "/v1/testString/zones/testString/custom_pages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listZoneCustomPagesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListZoneCustomPages with error: Operation response processing error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListZoneCustomPagesOptions model
				listZoneCustomPagesOptionsModel := new(custompagesv1.ListZoneCustomPagesOptions)
				listZoneCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListZoneCustomPages(listZoneCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListZoneCustomPages(listZoneCustomPagesOptions *ListZoneCustomPagesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listZoneCustomPagesPath := "/v1/testString/zones/testString/custom_pages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listZoneCustomPagesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}], "result_info": {"page": 1, "per_page": 20, "total_pages": 1, "count": 10, "total_count": 10}}`)
				}))
			})
			It(`Invoke ListZoneCustomPages successfully`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListZoneCustomPages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListZoneCustomPagesOptions model
				listZoneCustomPagesOptionsModel := new(custompagesv1.ListZoneCustomPagesOptions)
				listZoneCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListZoneCustomPages(listZoneCustomPagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListZoneCustomPages with error: Operation request error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListZoneCustomPagesOptions model
				listZoneCustomPagesOptionsModel := new(custompagesv1.ListZoneCustomPagesOptions)
				listZoneCustomPagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListZoneCustomPages(listZoneCustomPagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneCustomPage(getZoneCustomPageOptions *GetZoneCustomPageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneCustomPagePath := "/v1/testString/zones/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getZoneCustomPagePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneCustomPage with error: Operation response processing error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetZoneCustomPageOptions model
				getZoneCustomPageOptionsModel := new(custompagesv1.GetZoneCustomPageOptions)
				getZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetZoneCustomPage(getZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetZoneCustomPage(getZoneCustomPageOptions *GetZoneCustomPageOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneCustomPagePath := "/v1/testString/zones/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getZoneCustomPagePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}}`)
				}))
			})
			It(`Invoke GetZoneCustomPage successfully`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetZoneCustomPage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneCustomPageOptions model
				getZoneCustomPageOptionsModel := new(custompagesv1.GetZoneCustomPageOptions)
				getZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetZoneCustomPage(getZoneCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetZoneCustomPage with error: Operation validation and request error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetZoneCustomPageOptions model
				getZoneCustomPageOptionsModel := new(custompagesv1.GetZoneCustomPageOptions)
				getZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				getZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetZoneCustomPage(getZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneCustomPageOptions model with no property values
				getZoneCustomPageOptionsModelNew := new(custompagesv1.GetZoneCustomPageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetZoneCustomPage(getZoneCustomPageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZoneCustomPage(updateZoneCustomPageOptions *UpdateZoneCustomPageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneCustomPagePath := "/v1/testString/zones/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateZoneCustomPagePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZoneCustomPage with error: Operation response processing error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneCustomPageOptions model
				updateZoneCustomPageOptionsModel := new(custompagesv1.UpdateZoneCustomPageOptions)
				updateZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateZoneCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateZoneCustomPageOptionsModel.State = core.StringPtr("customized")
				updateZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateZoneCustomPage(updateZoneCustomPageOptions *UpdateZoneCustomPageOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneCustomPagePath := "/v1/testString/zones/testString/custom_pages/basic_challenge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateZoneCustomPagePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "basic_challenge", "description": "Basic Challenge", "required_tokens": ["::CAPTCHA_BOX::"], "preview_target": "block:basic-sec-captcha", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00", "url": "https://www.example.com/basic_challenge_error.html", "state": "customized"}}`)
				}))
			})
			It(`Invoke UpdateZoneCustomPage successfully`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateZoneCustomPage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateZoneCustomPageOptions model
				updateZoneCustomPageOptionsModel := new(custompagesv1.UpdateZoneCustomPageOptions)
				updateZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateZoneCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateZoneCustomPageOptionsModel.State = core.StringPtr("customized")
				updateZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateZoneCustomPage with error: Operation validation and request error`, func() {
				testService, testServiceErr := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneCustomPageOptions model
				updateZoneCustomPageOptionsModel := new(custompagesv1.UpdateZoneCustomPageOptions)
				updateZoneCustomPageOptionsModel.PageIdentifier = core.StringPtr("basic_challenge")
				updateZoneCustomPageOptionsModel.URL = core.StringPtr("https://www.example.com/basic_challenge_error.html")
				updateZoneCustomPageOptionsModel.State = core.StringPtr("customized")
				updateZoneCustomPageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateZoneCustomPageOptions model with no property values
				updateZoneCustomPageOptionsModelNew := new(custompagesv1.UpdateZoneCustomPageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateZoneCustomPage(updateZoneCustomPageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			crn := "testString"
			zoneIdentifier := "testString"
			testService, _ := custompagesv1.NewCustomPagesV1(&custompagesv1.CustomPagesV1Options{
				URL:            "http://custompagesv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewGetInstanceCustomPageOptions successfully`, func() {
				// Construct an instance of the GetInstanceCustomPageOptions model
				pageIdentifier := "basic_challenge"
				getInstanceCustomPageOptionsModel := testService.NewGetInstanceCustomPageOptions(pageIdentifier)
				getInstanceCustomPageOptionsModel.SetPageIdentifier("basic_challenge")
				getInstanceCustomPageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInstanceCustomPageOptionsModel).ToNot(BeNil())
				Expect(getInstanceCustomPageOptionsModel.PageIdentifier).To(Equal(core.StringPtr("basic_challenge")))
				Expect(getInstanceCustomPageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneCustomPageOptions successfully`, func() {
				// Construct an instance of the GetZoneCustomPageOptions model
				pageIdentifier := "basic_challenge"
				getZoneCustomPageOptionsModel := testService.NewGetZoneCustomPageOptions(pageIdentifier)
				getZoneCustomPageOptionsModel.SetPageIdentifier("basic_challenge")
				getZoneCustomPageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneCustomPageOptionsModel).ToNot(BeNil())
				Expect(getZoneCustomPageOptionsModel.PageIdentifier).To(Equal(core.StringPtr("basic_challenge")))
				Expect(getZoneCustomPageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListInstanceCustomPagesOptions successfully`, func() {
				// Construct an instance of the ListInstanceCustomPagesOptions model
				listInstanceCustomPagesOptionsModel := testService.NewListInstanceCustomPagesOptions()
				listInstanceCustomPagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listInstanceCustomPagesOptionsModel).ToNot(BeNil())
				Expect(listInstanceCustomPagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListZoneCustomPagesOptions successfully`, func() {
				// Construct an instance of the ListZoneCustomPagesOptions model
				listZoneCustomPagesOptionsModel := testService.NewListZoneCustomPagesOptions()
				listZoneCustomPagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listZoneCustomPagesOptionsModel).ToNot(BeNil())
				Expect(listZoneCustomPagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateInstanceCustomPageOptions successfully`, func() {
				// Construct an instance of the UpdateInstanceCustomPageOptions model
				pageIdentifier := "basic_challenge"
				updateInstanceCustomPageOptionsModel := testService.NewUpdateInstanceCustomPageOptions(pageIdentifier)
				updateInstanceCustomPageOptionsModel.SetPageIdentifier("basic_challenge")
				updateInstanceCustomPageOptionsModel.SetURL("https://www.example.com/basic_challenge_error.html")
				updateInstanceCustomPageOptionsModel.SetState("customized")
				updateInstanceCustomPageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateInstanceCustomPageOptionsModel).ToNot(BeNil())
				Expect(updateInstanceCustomPageOptionsModel.PageIdentifier).To(Equal(core.StringPtr("basic_challenge")))
				Expect(updateInstanceCustomPageOptionsModel.URL).To(Equal(core.StringPtr("https://www.example.com/basic_challenge_error.html")))
				Expect(updateInstanceCustomPageOptionsModel.State).To(Equal(core.StringPtr("customized")))
				Expect(updateInstanceCustomPageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneCustomPageOptions successfully`, func() {
				// Construct an instance of the UpdateZoneCustomPageOptions model
				pageIdentifier := "basic_challenge"
				updateZoneCustomPageOptionsModel := testService.NewUpdateZoneCustomPageOptions(pageIdentifier)
				updateZoneCustomPageOptionsModel.SetPageIdentifier("basic_challenge")
				updateZoneCustomPageOptionsModel.SetURL("https://www.example.com/basic_challenge_error.html")
				updateZoneCustomPageOptionsModel.SetState("customized")
				updateZoneCustomPageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneCustomPageOptionsModel).ToNot(BeNil())
				Expect(updateZoneCustomPageOptionsModel.PageIdentifier).To(Equal(core.StringPtr("basic_challenge")))
				Expect(updateZoneCustomPageOptionsModel.URL).To(Equal(core.StringPtr("https://www.example.com/basic_challenge_error.html")))
				Expect(updateZoneCustomPageOptionsModel.State).To(Equal(core.StringPtr("customized")))
				Expect(updateZoneCustomPageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
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
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
