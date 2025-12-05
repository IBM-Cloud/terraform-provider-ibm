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

package drautomationservicev1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

var _ = Describe(`DrAutomationServiceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(drAutomationServiceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(drAutomationServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
				URL: "https://drautomationservicev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(drAutomationServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DR_AUTOMATION_SERVICE_URL": "https://drautomationservicev1/api",
				"DR_AUTOMATION_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1UsingExternalConfig(&drautomationservicev1.DrAutomationServiceV1Options{
				})
				Expect(drAutomationServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := drAutomationServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != drAutomationServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(drAutomationServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(drAutomationServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1UsingExternalConfig(&drautomationservicev1.DrAutomationServiceV1Options{
					URL: "https://testService/api",
				})
				Expect(drAutomationServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := drAutomationServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != drAutomationServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(drAutomationServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(drAutomationServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1UsingExternalConfig(&drautomationservicev1.DrAutomationServiceV1Options{
				})
				err := drAutomationServiceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := drAutomationServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != drAutomationServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(drAutomationServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(drAutomationServiceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DR_AUTOMATION_SERVICE_URL": "https://drautomationservicev1/api",
				"DR_AUTOMATION_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1UsingExternalConfig(&drautomationservicev1.DrAutomationServiceV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(drAutomationServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DR_AUTOMATION_SERVICE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1UsingExternalConfig(&drautomationservicev1.DrAutomationServiceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(drAutomationServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = drautomationservicev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetApikey(getApikeyOptions *GetApikeyOptions) - Operation response error`, func() {
		getApikeyPath := "/drautomation/v1/apikey/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApikeyPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApikey with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetApikeyOptions model
				getApikeyOptionsModel := new(drautomationservicev1.GetApikeyOptions)
				getApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetApikey(getApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetApikey(getApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApikey(getApikeyOptions *GetApikeyOptions)`, func() {
		getApikeyPath := "/drautomation/v1/apikey/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApikeyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Key is valid.", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::", "status": "Active"}`)
				}))
			})
			It(`Invoke GetApikey successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetApikeyOptions model
				getApikeyOptionsModel := new(drautomationservicev1.GetApikeyOptions)
				getApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetApikeyWithContext(ctx, getApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetApikey(getApikeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetApikeyWithContext(ctx, getApikeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getApikeyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Key is valid.", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::", "status": "Active"}`)
				}))
			})
			It(`Invoke GetApikey successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetApikey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApikeyOptions model
				getApikeyOptionsModel := new(drautomationservicev1.GetApikeyOptions)
				getApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetApikey(getApikeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApikey with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetApikeyOptions model
				getApikeyOptionsModel := new(drautomationservicev1.GetApikeyOptions)
				getApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetApikey(getApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApikeyOptions model with no property values
				getApikeyOptionsModelNew := new(drautomationservicev1.GetApikeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetApikey(getApikeyOptionsModelNew)
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
			It(`Invoke GetApikey successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetApikeyOptions model
				getApikeyOptionsModel := new(drautomationservicev1.GetApikeyOptions)
				getApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetApikey(getApikeyOptionsModel)
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
	Describe(`CreateApikey(createApikeyOptions *CreateApikeyOptions) - Operation response error`, func() {
		createApikeyPath := "/drautomation/v1/apikey/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApikeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateApikey with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateApikeyOptions model
				createApikeyOptionsModel := new(drautomationservicev1.CreateApikeyOptions)
				createApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createApikeyOptionsModel.APIKey = core.StringPtr("abcdefrg_izklmnop_fxbEED")
				createApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.CreateApikey(createApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.CreateApikey(createApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateApikey(createApikeyOptions *CreateApikeyOptions)`, func() {
		createApikeyPath := "/drautomation/v1/apikey/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createApikeyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Key is valid.", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::", "status": "Active"}`)
				}))
			})
			It(`Invoke CreateApikey successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the CreateApikeyOptions model
				createApikeyOptionsModel := new(drautomationservicev1.CreateApikeyOptions)
				createApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createApikeyOptionsModel.APIKey = core.StringPtr("abcdefrg_izklmnop_fxbEED")
				createApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.CreateApikeyWithContext(ctx, createApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.CreateApikey(createApikeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.CreateApikeyWithContext(ctx, createApikeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createApikeyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Key is valid.", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::", "status": "Active"}`)
				}))
			})
			It(`Invoke CreateApikey successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.CreateApikey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateApikeyOptions model
				createApikeyOptionsModel := new(drautomationservicev1.CreateApikeyOptions)
				createApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createApikeyOptionsModel.APIKey = core.StringPtr("abcdefrg_izklmnop_fxbEED")
				createApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.CreateApikey(createApikeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateApikey with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateApikeyOptions model
				createApikeyOptionsModel := new(drautomationservicev1.CreateApikeyOptions)
				createApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createApikeyOptionsModel.APIKey = core.StringPtr("abcdefrg_izklmnop_fxbEED")
				createApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.CreateApikey(createApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateApikeyOptions model with no property values
				createApikeyOptionsModelNew := new(drautomationservicev1.CreateApikeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.CreateApikey(createApikeyOptionsModelNew)
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
			It(`Invoke CreateApikey successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateApikeyOptions model
				createApikeyOptionsModel := new(drautomationservicev1.CreateApikeyOptions)
				createApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createApikeyOptionsModel.APIKey = core.StringPtr("abcdefrg_izklmnop_fxbEED")
				createApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.CreateApikey(createApikeyOptionsModel)
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
	Describe(`UpdateApikey(updateApikeyOptions *UpdateApikeyOptions) - Operation response error`, func() {
		updateApikeyPath := "/drautomation/v1/apikey/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateApikeyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateApikey with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the UpdateApikeyOptions model
				updateApikeyOptionsModel := new(drautomationservicev1.UpdateApikeyOptions)
				updateApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				updateApikeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				updateApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.UpdateApikey(updateApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.UpdateApikey(updateApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateApikey(updateApikeyOptions *UpdateApikeyOptions)`, func() {
		updateApikeyPath := "/drautomation/v1/apikey/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateApikeyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Key is valid.", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::", "status": "Active"}`)
				}))
			})
			It(`Invoke UpdateApikey successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the UpdateApikeyOptions model
				updateApikeyOptionsModel := new(drautomationservicev1.UpdateApikeyOptions)
				updateApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				updateApikeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				updateApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.UpdateApikeyWithContext(ctx, updateApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.UpdateApikey(updateApikeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.UpdateApikeyWithContext(ctx, updateApikeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateApikeyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Key is valid.", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::", "status": "Active"}`)
				}))
			})
			It(`Invoke UpdateApikey successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.UpdateApikey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateApikeyOptions model
				updateApikeyOptionsModel := new(drautomationservicev1.UpdateApikeyOptions)
				updateApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				updateApikeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				updateApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.UpdateApikey(updateApikeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateApikey with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the UpdateApikeyOptions model
				updateApikeyOptionsModel := new(drautomationservicev1.UpdateApikeyOptions)
				updateApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				updateApikeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				updateApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.UpdateApikey(updateApikeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateApikeyOptions model with no property values
				updateApikeyOptionsModelNew := new(drautomationservicev1.UpdateApikeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.UpdateApikey(updateApikeyOptionsModelNew)
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
			It(`Invoke UpdateApikey successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the UpdateApikeyOptions model
				updateApikeyOptionsModel := new(drautomationservicev1.UpdateApikeyOptions)
				updateApikeyOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				updateApikeyOptionsModel.APIKey = core.StringPtr("adfadfdsafsdfdsf")
				updateApikeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateApikeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.UpdateApikey(updateApikeyOptionsModel)
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
	Describe(`GetDrGrsLocationPair(getDrGrsLocationPairOptions *GetDrGrsLocationPairOptions) - Operation response error`, func() {
		getDrGrsLocationPairPath := "/drautomation/v1/dr_grs_location_pairs/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDrGrsLocationPairPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDrGrsLocationPair with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrGrsLocationPairOptions model
				getDrGrsLocationPairOptionsModel := new(drautomationservicev1.GetDrGrsLocationPairOptions)
				getDrGrsLocationPairOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrGrsLocationPairOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrGrsLocationPairOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetDrGrsLocationPair(getDrGrsLocationPairOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetDrGrsLocationPair(getDrGrsLocationPairOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDrGrsLocationPair(getDrGrsLocationPairOptions *GetDrGrsLocationPairOptions)`, func() {
		getDrGrsLocationPairPath := "/drautomation/v1/dr_grs_location_pairs/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDrGrsLocationPairPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location_pairs": {"mapKey": "Inner"}}`)
				}))
			})
			It(`Invoke GetDrGrsLocationPair successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetDrGrsLocationPairOptions model
				getDrGrsLocationPairOptionsModel := new(drautomationservicev1.GetDrGrsLocationPairOptions)
				getDrGrsLocationPairOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrGrsLocationPairOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrGrsLocationPairOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetDrGrsLocationPairWithContext(ctx, getDrGrsLocationPairOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetDrGrsLocationPair(getDrGrsLocationPairOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetDrGrsLocationPairWithContext(ctx, getDrGrsLocationPairOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDrGrsLocationPairPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location_pairs": {"mapKey": "Inner"}}`)
				}))
			})
			It(`Invoke GetDrGrsLocationPair successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetDrGrsLocationPair(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDrGrsLocationPairOptions model
				getDrGrsLocationPairOptionsModel := new(drautomationservicev1.GetDrGrsLocationPairOptions)
				getDrGrsLocationPairOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrGrsLocationPairOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrGrsLocationPairOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetDrGrsLocationPair(getDrGrsLocationPairOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDrGrsLocationPair with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrGrsLocationPairOptions model
				getDrGrsLocationPairOptionsModel := new(drautomationservicev1.GetDrGrsLocationPairOptions)
				getDrGrsLocationPairOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrGrsLocationPairOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrGrsLocationPairOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetDrGrsLocationPair(getDrGrsLocationPairOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDrGrsLocationPairOptions model with no property values
				getDrGrsLocationPairOptionsModelNew := new(drautomationservicev1.GetDrGrsLocationPairOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetDrGrsLocationPair(getDrGrsLocationPairOptionsModelNew)
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
			It(`Invoke GetDrGrsLocationPair successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrGrsLocationPairOptions model
				getDrGrsLocationPairOptionsModel := new(drautomationservicev1.GetDrGrsLocationPairOptions)
				getDrGrsLocationPairOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrGrsLocationPairOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrGrsLocationPairOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetDrGrsLocationPair(getDrGrsLocationPairOptionsModel)
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
	Describe(`GetDrLocations(getDrLocationsOptions *GetDrLocationsOptions) - Operation response error`, func() {
		getDrLocationsPath := "/drautomation/v1/dr_locations/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDrLocationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDrLocations with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrLocationsOptions model
				getDrLocationsOptionsModel := new(drautomationservicev1.GetDrLocationsOptions)
				getDrLocationsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrLocationsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetDrLocations(getDrLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetDrLocations(getDrLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDrLocations(getDrLocationsOptions *GetDrLocationsOptions)`, func() {
		getDrLocationsPath := "/drautomation/v1/dr_locations/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDrLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dr_locations": [{"id": "loc123", "name": "US-East-1"}]}`)
				}))
			})
			It(`Invoke GetDrLocations successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetDrLocationsOptions model
				getDrLocationsOptionsModel := new(drautomationservicev1.GetDrLocationsOptions)
				getDrLocationsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrLocationsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetDrLocationsWithContext(ctx, getDrLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetDrLocations(getDrLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetDrLocationsWithContext(ctx, getDrLocationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDrLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dr_locations": [{"id": "loc123", "name": "US-East-1"}]}`)
				}))
			})
			It(`Invoke GetDrLocations successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetDrLocations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDrLocationsOptions model
				getDrLocationsOptionsModel := new(drautomationservicev1.GetDrLocationsOptions)
				getDrLocationsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrLocationsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetDrLocations(getDrLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDrLocations with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrLocationsOptions model
				getDrLocationsOptionsModel := new(drautomationservicev1.GetDrLocationsOptions)
				getDrLocationsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrLocationsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetDrLocations(getDrLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDrLocationsOptions model with no property values
				getDrLocationsOptionsModelNew := new(drautomationservicev1.GetDrLocationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetDrLocations(getDrLocationsOptionsModelNew)
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
			It(`Invoke GetDrLocations successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrLocationsOptions model
				getDrLocationsOptionsModel := new(drautomationservicev1.GetDrLocationsOptions)
				getDrLocationsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrLocationsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetDrLocations(getDrLocationsOptionsModel)
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
	Describe(`GetDrManagedVM(getDrManagedVMOptions *GetDrManagedVMOptions) - Operation response error`, func() {
		getDrManagedVMPath := "/drautomation/v1/dr_managed_vms/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDrManagedVMPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDrManagedVM with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrManagedVMOptions model
				getDrManagedVMOptionsModel := new(drautomationservicev1.GetDrManagedVMOptions)
				getDrManagedVMOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrManagedVMOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrManagedVMOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetDrManagedVM(getDrManagedVMOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetDrManagedVM(getDrManagedVMOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDrManagedVM(getDrManagedVMOptions *GetDrManagedVMOptions)`, func() {
		getDrManagedVMPath := "/drautomation/v1/dr_managed_vms/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDrManagedVMPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"managed_vm_list": {"mapKey": {"core": "0.50", "dr_average_time": "10", "dr_region": "nyc02", "memory": "4", "region": "lon04", "vm_name": "example_vm", "workgroup_name": "Workgroup1", "workspace_name": "Workspace_dallas01"}}}`)
				}))
			})
			It(`Invoke GetDrManagedVM successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetDrManagedVMOptions model
				getDrManagedVMOptionsModel := new(drautomationservicev1.GetDrManagedVMOptions)
				getDrManagedVMOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrManagedVMOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrManagedVMOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetDrManagedVMWithContext(ctx, getDrManagedVMOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetDrManagedVM(getDrManagedVMOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetDrManagedVMWithContext(ctx, getDrManagedVMOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDrManagedVMPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"managed_vm_list": {"mapKey": {"core": "0.50", "dr_average_time": "10", "dr_region": "nyc02", "memory": "4", "region": "lon04", "vm_name": "example_vm", "workgroup_name": "Workgroup1", "workspace_name": "Workspace_dallas01"}}}`)
				}))
			})
			It(`Invoke GetDrManagedVM successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetDrManagedVM(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDrManagedVMOptions model
				getDrManagedVMOptionsModel := new(drautomationservicev1.GetDrManagedVMOptions)
				getDrManagedVMOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrManagedVMOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrManagedVMOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetDrManagedVM(getDrManagedVMOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDrManagedVM with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrManagedVMOptions model
				getDrManagedVMOptionsModel := new(drautomationservicev1.GetDrManagedVMOptions)
				getDrManagedVMOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrManagedVMOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrManagedVMOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetDrManagedVM(getDrManagedVMOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDrManagedVMOptions model with no property values
				getDrManagedVMOptionsModelNew := new(drautomationservicev1.GetDrManagedVMOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetDrManagedVM(getDrManagedVMOptionsModelNew)
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
			It(`Invoke GetDrManagedVM successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrManagedVMOptions model
				getDrManagedVMOptionsModel := new(drautomationservicev1.GetDrManagedVMOptions)
				getDrManagedVMOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrManagedVMOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrManagedVMOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetDrManagedVM(getDrManagedVMOptionsModel)
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
	Describe(`GetDrSummary(getDrSummaryOptions *GetDrSummaryOptions) - Operation response error`, func() {
		getDrSummaryPath := "/drautomation/v1/dr_summary/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDrSummaryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDrSummary with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrSummaryOptions model
				getDrSummaryOptionsModel := new(drautomationservicev1.GetDrSummaryOptions)
				getDrSummaryOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrSummaryOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetDrSummary(getDrSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetDrSummary(getDrSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDrSummary(getDrSummaryOptions *GetDrSummaryOptions)`, func() {
		getDrSummaryPath := "/drautomation/v1/dr_summary/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDrSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"managed_vm_list": {"anyKey": "anyValue"}, "orchestrator_details": {"last_updated_orchestrator_deployment_time": "2025-10-16T09:28:13.696Z", "last_updated_standby_orchestrator_deployment_time": "2025-10-16T09:28:13.696Z", "latest_orchestrator_time": "2025-10-16T09:28:13.696Z", "location_id": "LocationID", "mfa_enabled": "MfaEnabled", "orch_ext_connectivity_status": "OrchExtConnectivityStatus", "orch_standby_node_addition_status": "OrchStandbyNodeAdditionStatus", "orchestrator_cluster_message": "OrchestratorClusterMessage", "orchestrator_config_status": "OrchestratorConfigStatus", "orchestrator_group_leader": "OrchestratorGroupLeader", "orchestrator_location_type": "OrchestratorLocationType", "orchestrator_name": "OrchestratorName", "orchestrator_status": "OrchestratorStatus", "orchestrator_workspace_name": "OrchestratorWorkspaceName", "proxy_ip": "ProxyIP", "schematic_workspace_name": "SchematicWorkspaceName", "schematic_workspace_status": "SchematicWorkspaceStatus", "ssh_key_name": "SSHKeyName", "standby_orchestrator_name": "StandbyOrchestratorName", "standby_orchestrator_status": "StandbyOrchestratorStatus", "standby_orchestrator_workspace_name": "StandbyOrchestratorWorkspaceName", "transit_gateway_name": "TransitGatewayName", "vpc_name": "VPCName"}, "service_details": {"crn": "CRN", "deployment_name": "DeploymentName", "description": "Description", "orchestrator_ha": true, "plan_name": "PlanName", "primary_ip_address": "PrimaryIPAddress", "primary_orchestrator_dashboard_url": "PrimaryOrchestratorDashboardURL", "recovery_location": "RecoveryLocation", "resource_group": "ResourceGroup", "standby_description": "StandbyDescription", "standby_ip_address": "StandbyIPAddress", "standby_orchestrator_dashboard_url": "StandbyOrchestratorDashboardURL", "standby_status": "StandbyStatus", "status": "Status"}}`)
				}))
			})
			It(`Invoke GetDrSummary successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetDrSummaryOptions model
				getDrSummaryOptionsModel := new(drautomationservicev1.GetDrSummaryOptions)
				getDrSummaryOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrSummaryOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetDrSummaryWithContext(ctx, getDrSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetDrSummary(getDrSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetDrSummaryWithContext(ctx, getDrSummaryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDrSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"managed_vm_list": {"anyKey": "anyValue"}, "orchestrator_details": {"last_updated_orchestrator_deployment_time": "2025-10-16T09:28:13.696Z", "last_updated_standby_orchestrator_deployment_time": "2025-10-16T09:28:13.696Z", "latest_orchestrator_time": "2025-10-16T09:28:13.696Z", "location_id": "LocationID", "mfa_enabled": "MfaEnabled", "orch_ext_connectivity_status": "OrchExtConnectivityStatus", "orch_standby_node_addition_status": "OrchStandbyNodeAdditionStatus", "orchestrator_cluster_message": "OrchestratorClusterMessage", "orchestrator_config_status": "OrchestratorConfigStatus", "orchestrator_group_leader": "OrchestratorGroupLeader", "orchestrator_location_type": "OrchestratorLocationType", "orchestrator_name": "OrchestratorName", "orchestrator_status": "OrchestratorStatus", "orchestrator_workspace_name": "OrchestratorWorkspaceName", "proxy_ip": "ProxyIP", "schematic_workspace_name": "SchematicWorkspaceName", "schematic_workspace_status": "SchematicWorkspaceStatus", "ssh_key_name": "SSHKeyName", "standby_orchestrator_name": "StandbyOrchestratorName", "standby_orchestrator_status": "StandbyOrchestratorStatus", "standby_orchestrator_workspace_name": "StandbyOrchestratorWorkspaceName", "transit_gateway_name": "TransitGatewayName", "vpc_name": "VPCName"}, "service_details": {"crn": "CRN", "deployment_name": "DeploymentName", "description": "Description", "orchestrator_ha": true, "plan_name": "PlanName", "primary_ip_address": "PrimaryIPAddress", "primary_orchestrator_dashboard_url": "PrimaryOrchestratorDashboardURL", "recovery_location": "RecoveryLocation", "resource_group": "ResourceGroup", "standby_description": "StandbyDescription", "standby_ip_address": "StandbyIPAddress", "standby_orchestrator_dashboard_url": "StandbyOrchestratorDashboardURL", "standby_status": "StandbyStatus", "status": "Status"}}`)
				}))
			})
			It(`Invoke GetDrSummary successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetDrSummary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDrSummaryOptions model
				getDrSummaryOptionsModel := new(drautomationservicev1.GetDrSummaryOptions)
				getDrSummaryOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrSummaryOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetDrSummary(getDrSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDrSummary with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrSummaryOptions model
				getDrSummaryOptionsModel := new(drautomationservicev1.GetDrSummaryOptions)
				getDrSummaryOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrSummaryOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetDrSummary(getDrSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDrSummaryOptions model with no property values
				getDrSummaryOptionsModelNew := new(drautomationservicev1.GetDrSummaryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetDrSummary(getDrSummaryOptionsModelNew)
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
			It(`Invoke GetDrSummary successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetDrSummaryOptions model
				getDrSummaryOptionsModel := new(drautomationservicev1.GetDrSummaryOptions)
				getDrSummaryOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrSummaryOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDrSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetDrSummary(getDrSummaryOptionsModel)
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
	Describe(`GetMachineType(getMachineTypeOptions *GetMachineTypeOptions) - Operation response error`, func() {
		getMachineTypePath := "/drautomation/v1/machinetypes/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMachineTypePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["primary_workspace_name"]).To(Equal([]string{"Test-workspace-wdc06"}))
					Expect(req.URL.Query()["standby_workspace_name"]).To(Equal([]string{"Test-workspace-wdc07"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMachineType with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetMachineTypeOptions model
				getMachineTypeOptionsModel := new(drautomationservicev1.GetMachineTypeOptions)
				getMachineTypeOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getMachineTypeOptionsModel.PrimaryWorkspaceName = core.StringPtr("Test-workspace-wdc06")
				getMachineTypeOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getMachineTypeOptionsModel.StandbyWorkspaceName = core.StringPtr("Test-workspace-wdc07")
				getMachineTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetMachineType(getMachineTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetMachineType(getMachineTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMachineType(getMachineTypeOptions *GetMachineTypeOptions)`, func() {
		getMachineTypePath := "/drautomation/v1/machinetypes/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMachineTypePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["primary_workspace_name"]).To(Equal([]string{"Test-workspace-wdc06"}))
					Expect(req.URL.Query()["standby_workspace_name"]).To(Equal([]string{"Test-workspace-wdc07"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"workspaces": {"mapKey": ["Inner"]}}`)
				}))
			})
			It(`Invoke GetMachineType successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetMachineTypeOptions model
				getMachineTypeOptionsModel := new(drautomationservicev1.GetMachineTypeOptions)
				getMachineTypeOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getMachineTypeOptionsModel.PrimaryWorkspaceName = core.StringPtr("Test-workspace-wdc06")
				getMachineTypeOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getMachineTypeOptionsModel.StandbyWorkspaceName = core.StringPtr("Test-workspace-wdc07")
				getMachineTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetMachineTypeWithContext(ctx, getMachineTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetMachineType(getMachineTypeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetMachineTypeWithContext(ctx, getMachineTypeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getMachineTypePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["primary_workspace_name"]).To(Equal([]string{"Test-workspace-wdc06"}))
					Expect(req.URL.Query()["standby_workspace_name"]).To(Equal([]string{"Test-workspace-wdc07"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"workspaces": {"mapKey": ["Inner"]}}`)
				}))
			})
			It(`Invoke GetMachineType successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetMachineType(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMachineTypeOptions model
				getMachineTypeOptionsModel := new(drautomationservicev1.GetMachineTypeOptions)
				getMachineTypeOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getMachineTypeOptionsModel.PrimaryWorkspaceName = core.StringPtr("Test-workspace-wdc06")
				getMachineTypeOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getMachineTypeOptionsModel.StandbyWorkspaceName = core.StringPtr("Test-workspace-wdc07")
				getMachineTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetMachineType(getMachineTypeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMachineType with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetMachineTypeOptions model
				getMachineTypeOptionsModel := new(drautomationservicev1.GetMachineTypeOptions)
				getMachineTypeOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getMachineTypeOptionsModel.PrimaryWorkspaceName = core.StringPtr("Test-workspace-wdc06")
				getMachineTypeOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getMachineTypeOptionsModel.StandbyWorkspaceName = core.StringPtr("Test-workspace-wdc07")
				getMachineTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetMachineType(getMachineTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetMachineTypeOptions model with no property values
				getMachineTypeOptionsModelNew := new(drautomationservicev1.GetMachineTypeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetMachineType(getMachineTypeOptionsModelNew)
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
			It(`Invoke GetMachineType successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetMachineTypeOptions model
				getMachineTypeOptionsModel := new(drautomationservicev1.GetMachineTypeOptions)
				getMachineTypeOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getMachineTypeOptionsModel.PrimaryWorkspaceName = core.StringPtr("Test-workspace-wdc06")
				getMachineTypeOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getMachineTypeOptionsModel.StandbyWorkspaceName = core.StringPtr("Test-workspace-wdc07")
				getMachineTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetMachineType(getMachineTypeOptionsModel)
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
	Describe(`GetPowervsWorkspaces(getPowervsWorkspacesOptions *GetPowervsWorkspacesOptions) - Operation response error`, func() {
		getPowervsWorkspacesPath := "/drautomation/v1/powervs_workspaces/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPowervsWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["location_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPowervsWorkspaces with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPowervsWorkspacesOptions model
				getPowervsWorkspacesOptionsModel := new(drautomationservicev1.GetPowervsWorkspacesOptions)
				getPowervsWorkspacesOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getPowervsWorkspacesOptionsModel.LocationID = core.StringPtr("testString")
				getPowervsWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetPowervsWorkspaces(getPowervsWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetPowervsWorkspaces(getPowervsWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPowervsWorkspaces(getPowervsWorkspacesOptions *GetPowervsWorkspacesOptions)`, func() {
		getPowervsWorkspacesPath := "/drautomation/v1/powervs_workspaces/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPowervsWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["location_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dr_standby_workspace_description": "anyValue", "dr_standby_workspaces": [{"details": {"crn": "crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::"}, "id": "ID", "location": {"region": "lon06", "type": "data-center", "url": "https://lon.power-iaas.cloud.ibm.com"}, "name": "Name", "status": "Status"}], "dr_workspace_description": "anyValue", "dr_workspaces": [{"default": true, "details": {"crn": "crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::"}, "id": "ID", "location": {"region": "lon06", "type": "data-center", "url": "https://lon.power-iaas.cloud.ibm.com"}, "name": "Name", "status": "active"}]}`)
				}))
			})
			It(`Invoke GetPowervsWorkspaces successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetPowervsWorkspacesOptions model
				getPowervsWorkspacesOptionsModel := new(drautomationservicev1.GetPowervsWorkspacesOptions)
				getPowervsWorkspacesOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getPowervsWorkspacesOptionsModel.LocationID = core.StringPtr("testString")
				getPowervsWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetPowervsWorkspacesWithContext(ctx, getPowervsWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetPowervsWorkspaces(getPowervsWorkspacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetPowervsWorkspacesWithContext(ctx, getPowervsWorkspacesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPowervsWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["location_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dr_standby_workspace_description": "anyValue", "dr_standby_workspaces": [{"details": {"crn": "crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::"}, "id": "ID", "location": {"region": "lon06", "type": "data-center", "url": "https://lon.power-iaas.cloud.ibm.com"}, "name": "Name", "status": "Status"}], "dr_workspace_description": "anyValue", "dr_workspaces": [{"default": true, "details": {"crn": "crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::"}, "id": "ID", "location": {"region": "lon06", "type": "data-center", "url": "https://lon.power-iaas.cloud.ibm.com"}, "name": "Name", "status": "active"}]}`)
				}))
			})
			It(`Invoke GetPowervsWorkspaces successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetPowervsWorkspaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPowervsWorkspacesOptions model
				getPowervsWorkspacesOptionsModel := new(drautomationservicev1.GetPowervsWorkspacesOptions)
				getPowervsWorkspacesOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getPowervsWorkspacesOptionsModel.LocationID = core.StringPtr("testString")
				getPowervsWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetPowervsWorkspaces(getPowervsWorkspacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPowervsWorkspaces with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPowervsWorkspacesOptions model
				getPowervsWorkspacesOptionsModel := new(drautomationservicev1.GetPowervsWorkspacesOptions)
				getPowervsWorkspacesOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getPowervsWorkspacesOptionsModel.LocationID = core.StringPtr("testString")
				getPowervsWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetPowervsWorkspaces(getPowervsWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPowervsWorkspacesOptions model with no property values
				getPowervsWorkspacesOptionsModelNew := new(drautomationservicev1.GetPowervsWorkspacesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetPowervsWorkspaces(getPowervsWorkspacesOptionsModelNew)
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
			It(`Invoke GetPowervsWorkspaces successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetPowervsWorkspacesOptions model
				getPowervsWorkspacesOptionsModel := new(drautomationservicev1.GetPowervsWorkspacesOptions)
				getPowervsWorkspacesOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getPowervsWorkspacesOptionsModel.LocationID = core.StringPtr("testString")
				getPowervsWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetPowervsWorkspaces(getPowervsWorkspacesOptionsModel)
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
	Describe(`GetManageDr(getManageDrOptions *GetManageDrOptions) - Operation response error`, func() {
		getManageDrPath := "/drautomation/v1/manage_dr/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getManageDrPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetManageDr with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetManageDrOptions model
				getManageDrOptionsModel := new(drautomationservicev1.GetManageDrOptions)
				getManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetManageDr(getManageDrOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetManageDr(getManageDrOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetManageDr(getManageDrOptions *GetManageDrOptions)`, func() {
		getManageDrPath := "/drautomation/v1/manage_dr/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getManageDrPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dashboard_url": "https://power-dra.test.cloud.ibm.com/power-dra-ui?instance_id=crn:v1:bluemix:public:power-dr-automation:us-south:a/fe3c2ccd058e407c81e1dba2b5c0e0d6:e3d09875-bbf8-4d8a-b52c-abefb67a53c5::", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"}`)
				}))
			})
			It(`Invoke GetManageDr successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetManageDrOptions model
				getManageDrOptionsModel := new(drautomationservicev1.GetManageDrOptions)
				getManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetManageDrWithContext(ctx, getManageDrOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetManageDr(getManageDrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetManageDrWithContext(ctx, getManageDrOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getManageDrPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dashboard_url": "https://power-dra.test.cloud.ibm.com/power-dra-ui?instance_id=crn:v1:bluemix:public:power-dr-automation:us-south:a/fe3c2ccd058e407c81e1dba2b5c0e0d6:e3d09875-bbf8-4d8a-b52c-abefb67a53c5::", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"}`)
				}))
			})
			It(`Invoke GetManageDr successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetManageDr(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetManageDrOptions model
				getManageDrOptionsModel := new(drautomationservicev1.GetManageDrOptions)
				getManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetManageDr(getManageDrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetManageDr with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetManageDrOptions model
				getManageDrOptionsModel := new(drautomationservicev1.GetManageDrOptions)
				getManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetManageDr(getManageDrOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetManageDrOptions model with no property values
				getManageDrOptionsModelNew := new(drautomationservicev1.GetManageDrOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetManageDr(getManageDrOptionsModelNew)
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
			It(`Invoke GetManageDr successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetManageDrOptions model
				getManageDrOptionsModel := new(drautomationservicev1.GetManageDrOptions)
				getManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetManageDr(getManageDrOptionsModel)
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
	Describe(`CreateManageDr(createManageDrOptions *CreateManageDrOptions) - Operation response error`, func() {
		createManageDrPath := "/drautomation/v1/manage_dr/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createManageDrPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["stand_by_redeploy"]).To(Equal([]string{"testString"}))
					// TODO: Add check for accepts_incomplete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateManageDr with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateManageDrOptions model
				createManageDrOptionsModel := new(drautomationservicev1.CreateManageDrOptions)
				createManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createManageDrOptionsModel.LocationID = core.StringPtr("dal10")
				createManageDrOptionsModel.MachineType = core.StringPtr("bx2-4x16")
				createManageDrOptionsModel.OrchestratorLocationType = core.StringPtr("off-premises")
				createManageDrOptionsModel.OrchestratorName = core.StringPtr("adminUser")
				createManageDrOptionsModel.OrchestratorPassword = core.StringPtr("testString")
				createManageDrOptionsModel.OrchestratorWorkspaceID = core.StringPtr("orch-workspace-01")
				createManageDrOptionsModel.APIKey = core.StringPtr("testString")
				createManageDrOptionsModel.ClientID = core.StringPtr("abcd-97d2-1234-bf62-8eaecc67a1234")
				createManageDrOptionsModel.ClientSecret = core.StringPtr("abcd1234xM1y123wK6qR9123456789bE2jG0pabcdefgh")
				createManageDrOptionsModel.GUID = core.StringPtr("123e4567-e89b-12d3-a456-426614174000")
				createManageDrOptionsModel.OrchestratorHa = core.BoolPtr(true)
				createManageDrOptionsModel.ProxyIP = core.StringPtr("10.40.30.10:8888")
				createManageDrOptionsModel.RegionID = core.StringPtr("us-south")
				createManageDrOptionsModel.ResourceInstance = core.StringPtr("crn:v1:bluemix:public:resource-controller::res123")
				createManageDrOptionsModel.Secret = core.StringPtr("testString")
				createManageDrOptionsModel.SecretGroup = core.StringPtr("default-secret-group")
				createManageDrOptionsModel.SSHKeyName = core.StringPtr("my-ssh-key")
				createManageDrOptionsModel.StandbyMachineType = core.StringPtr("bx2-8x32")
				createManageDrOptionsModel.StandbyOrchestratorName = core.StringPtr("standbyAdmin")
				createManageDrOptionsModel.StandbyOrchestratorWorkspaceID = core.StringPtr("orch-standby-02")
				createManageDrOptionsModel.StandbyTier = core.StringPtr("Premium")
				createManageDrOptionsModel.TenantName = core.StringPtr("xxx.ibm.com")
				createManageDrOptionsModel.Tier = core.StringPtr("Standard")
				createManageDrOptionsModel.StandByRedeploy = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				createManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.CreateManageDr(createManageDrOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.CreateManageDr(createManageDrOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateManageDr(createManageDrOptions *CreateManageDrOptions)`, func() {
		createManageDrPath := "/drautomation/v1/manage_dr/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createManageDrPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["stand_by_redeploy"]).To(Equal([]string{"testString"}))
					// TODO: Add check for accepts_incomplete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dashboard_url": "https://power-dra.test.cloud.ibm.com/power-dra-ui?instance_id=crn:v1:bluemix:public:power-dr-automation:us-south:a/fe3c2ccd058e407c81e1dba2b5c0e0d6:e3d09875-bbf8-4d8a-b52c-abefb67a53c5::", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"}`)
				}))
			})
			It(`Invoke CreateManageDr successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the CreateManageDrOptions model
				createManageDrOptionsModel := new(drautomationservicev1.CreateManageDrOptions)
				createManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createManageDrOptionsModel.LocationID = core.StringPtr("dal10")
				createManageDrOptionsModel.MachineType = core.StringPtr("bx2-4x16")
				createManageDrOptionsModel.OrchestratorLocationType = core.StringPtr("off-premises")
				createManageDrOptionsModel.OrchestratorName = core.StringPtr("adminUser")
				createManageDrOptionsModel.OrchestratorPassword = core.StringPtr("testString")
				createManageDrOptionsModel.OrchestratorWorkspaceID = core.StringPtr("orch-workspace-01")
				createManageDrOptionsModel.APIKey = core.StringPtr("testString")
				createManageDrOptionsModel.ClientID = core.StringPtr("abcd-97d2-1234-bf62-8eaecc67a1234")
				createManageDrOptionsModel.ClientSecret = core.StringPtr("abcd1234xM1y123wK6qR9123456789bE2jG0pabcdefgh")
				createManageDrOptionsModel.GUID = core.StringPtr("123e4567-e89b-12d3-a456-426614174000")
				createManageDrOptionsModel.OrchestratorHa = core.BoolPtr(true)
				createManageDrOptionsModel.ProxyIP = core.StringPtr("10.40.30.10:8888")
				createManageDrOptionsModel.RegionID = core.StringPtr("us-south")
				createManageDrOptionsModel.ResourceInstance = core.StringPtr("crn:v1:bluemix:public:resource-controller::res123")
				createManageDrOptionsModel.Secret = core.StringPtr("testString")
				createManageDrOptionsModel.SecretGroup = core.StringPtr("default-secret-group")
				createManageDrOptionsModel.SSHKeyName = core.StringPtr("my-ssh-key")
				createManageDrOptionsModel.StandbyMachineType = core.StringPtr("bx2-8x32")
				createManageDrOptionsModel.StandbyOrchestratorName = core.StringPtr("standbyAdmin")
				createManageDrOptionsModel.StandbyOrchestratorWorkspaceID = core.StringPtr("orch-standby-02")
				createManageDrOptionsModel.StandbyTier = core.StringPtr("Premium")
				createManageDrOptionsModel.TenantName = core.StringPtr("xxx.ibm.com")
				createManageDrOptionsModel.Tier = core.StringPtr("Standard")
				createManageDrOptionsModel.StandByRedeploy = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				createManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.CreateManageDrWithContext(ctx, createManageDrOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.CreateManageDr(createManageDrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.CreateManageDrWithContext(ctx, createManageDrOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createManageDrPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["stand_by_redeploy"]).To(Equal([]string{"testString"}))
					// TODO: Add check for accepts_incomplete query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dashboard_url": "https://power-dra.test.cloud.ibm.com/power-dra-ui?instance_id=crn:v1:bluemix:public:power-dr-automation:us-south:a/fe3c2ccd058e407c81e1dba2b5c0e0d6:e3d09875-bbf8-4d8a-b52c-abefb67a53c5::", "id": "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"}`)
				}))
			})
			It(`Invoke CreateManageDr successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.CreateManageDr(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateManageDrOptions model
				createManageDrOptionsModel := new(drautomationservicev1.CreateManageDrOptions)
				createManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createManageDrOptionsModel.LocationID = core.StringPtr("dal10")
				createManageDrOptionsModel.MachineType = core.StringPtr("bx2-4x16")
				createManageDrOptionsModel.OrchestratorLocationType = core.StringPtr("off-premises")
				createManageDrOptionsModel.OrchestratorName = core.StringPtr("adminUser")
				createManageDrOptionsModel.OrchestratorPassword = core.StringPtr("testString")
				createManageDrOptionsModel.OrchestratorWorkspaceID = core.StringPtr("orch-workspace-01")
				createManageDrOptionsModel.APIKey = core.StringPtr("testString")
				createManageDrOptionsModel.ClientID = core.StringPtr("abcd-97d2-1234-bf62-8eaecc67a1234")
				createManageDrOptionsModel.ClientSecret = core.StringPtr("abcd1234xM1y123wK6qR9123456789bE2jG0pabcdefgh")
				createManageDrOptionsModel.GUID = core.StringPtr("123e4567-e89b-12d3-a456-426614174000")
				createManageDrOptionsModel.OrchestratorHa = core.BoolPtr(true)
				createManageDrOptionsModel.ProxyIP = core.StringPtr("10.40.30.10:8888")
				createManageDrOptionsModel.RegionID = core.StringPtr("us-south")
				createManageDrOptionsModel.ResourceInstance = core.StringPtr("crn:v1:bluemix:public:resource-controller::res123")
				createManageDrOptionsModel.Secret = core.StringPtr("testString")
				createManageDrOptionsModel.SecretGroup = core.StringPtr("default-secret-group")
				createManageDrOptionsModel.SSHKeyName = core.StringPtr("my-ssh-key")
				createManageDrOptionsModel.StandbyMachineType = core.StringPtr("bx2-8x32")
				createManageDrOptionsModel.StandbyOrchestratorName = core.StringPtr("standbyAdmin")
				createManageDrOptionsModel.StandbyOrchestratorWorkspaceID = core.StringPtr("orch-standby-02")
				createManageDrOptionsModel.StandbyTier = core.StringPtr("Premium")
				createManageDrOptionsModel.TenantName = core.StringPtr("xxx.ibm.com")
				createManageDrOptionsModel.Tier = core.StringPtr("Standard")
				createManageDrOptionsModel.StandByRedeploy = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				createManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.CreateManageDr(createManageDrOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateManageDr with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateManageDrOptions model
				createManageDrOptionsModel := new(drautomationservicev1.CreateManageDrOptions)
				createManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createManageDrOptionsModel.LocationID = core.StringPtr("dal10")
				createManageDrOptionsModel.MachineType = core.StringPtr("bx2-4x16")
				createManageDrOptionsModel.OrchestratorLocationType = core.StringPtr("off-premises")
				createManageDrOptionsModel.OrchestratorName = core.StringPtr("adminUser")
				createManageDrOptionsModel.OrchestratorPassword = core.StringPtr("testString")
				createManageDrOptionsModel.OrchestratorWorkspaceID = core.StringPtr("orch-workspace-01")
				createManageDrOptionsModel.APIKey = core.StringPtr("testString")
				createManageDrOptionsModel.ClientID = core.StringPtr("abcd-97d2-1234-bf62-8eaecc67a1234")
				createManageDrOptionsModel.ClientSecret = core.StringPtr("abcd1234xM1y123wK6qR9123456789bE2jG0pabcdefgh")
				createManageDrOptionsModel.GUID = core.StringPtr("123e4567-e89b-12d3-a456-426614174000")
				createManageDrOptionsModel.OrchestratorHa = core.BoolPtr(true)
				createManageDrOptionsModel.ProxyIP = core.StringPtr("10.40.30.10:8888")
				createManageDrOptionsModel.RegionID = core.StringPtr("us-south")
				createManageDrOptionsModel.ResourceInstance = core.StringPtr("crn:v1:bluemix:public:resource-controller::res123")
				createManageDrOptionsModel.Secret = core.StringPtr("testString")
				createManageDrOptionsModel.SecretGroup = core.StringPtr("default-secret-group")
				createManageDrOptionsModel.SSHKeyName = core.StringPtr("my-ssh-key")
				createManageDrOptionsModel.StandbyMachineType = core.StringPtr("bx2-8x32")
				createManageDrOptionsModel.StandbyOrchestratorName = core.StringPtr("standbyAdmin")
				createManageDrOptionsModel.StandbyOrchestratorWorkspaceID = core.StringPtr("orch-standby-02")
				createManageDrOptionsModel.StandbyTier = core.StringPtr("Premium")
				createManageDrOptionsModel.TenantName = core.StringPtr("xxx.ibm.com")
				createManageDrOptionsModel.Tier = core.StringPtr("Standard")
				createManageDrOptionsModel.StandByRedeploy = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				createManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.CreateManageDr(createManageDrOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateManageDrOptions model with no property values
				createManageDrOptionsModelNew := new(drautomationservicev1.CreateManageDrOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.CreateManageDr(createManageDrOptionsModelNew)
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
			It(`Invoke CreateManageDr successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the CreateManageDrOptions model
				createManageDrOptionsModel := new(drautomationservicev1.CreateManageDrOptions)
				createManageDrOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				createManageDrOptionsModel.LocationID = core.StringPtr("dal10")
				createManageDrOptionsModel.MachineType = core.StringPtr("bx2-4x16")
				createManageDrOptionsModel.OrchestratorLocationType = core.StringPtr("off-premises")
				createManageDrOptionsModel.OrchestratorName = core.StringPtr("adminUser")
				createManageDrOptionsModel.OrchestratorPassword = core.StringPtr("testString")
				createManageDrOptionsModel.OrchestratorWorkspaceID = core.StringPtr("orch-workspace-01")
				createManageDrOptionsModel.APIKey = core.StringPtr("testString")
				createManageDrOptionsModel.ClientID = core.StringPtr("abcd-97d2-1234-bf62-8eaecc67a1234")
				createManageDrOptionsModel.ClientSecret = core.StringPtr("abcd1234xM1y123wK6qR9123456789bE2jG0pabcdefgh")
				createManageDrOptionsModel.GUID = core.StringPtr("123e4567-e89b-12d3-a456-426614174000")
				createManageDrOptionsModel.OrchestratorHa = core.BoolPtr(true)
				createManageDrOptionsModel.ProxyIP = core.StringPtr("10.40.30.10:8888")
				createManageDrOptionsModel.RegionID = core.StringPtr("us-south")
				createManageDrOptionsModel.ResourceInstance = core.StringPtr("crn:v1:bluemix:public:resource-controller::res123")
				createManageDrOptionsModel.Secret = core.StringPtr("testString")
				createManageDrOptionsModel.SecretGroup = core.StringPtr("default-secret-group")
				createManageDrOptionsModel.SSHKeyName = core.StringPtr("my-ssh-key")
				createManageDrOptionsModel.StandbyMachineType = core.StringPtr("bx2-8x32")
				createManageDrOptionsModel.StandbyOrchestratorName = core.StringPtr("standbyAdmin")
				createManageDrOptionsModel.StandbyOrchestratorWorkspaceID = core.StringPtr("orch-standby-02")
				createManageDrOptionsModel.StandbyTier = core.StringPtr("Premium")
				createManageDrOptionsModel.TenantName = core.StringPtr("xxx.ibm.com")
				createManageDrOptionsModel.Tier = core.StringPtr("Standard")
				createManageDrOptionsModel.StandByRedeploy = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createManageDrOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				createManageDrOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.CreateManageDr(createManageDrOptionsModel)
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
	Describe(`GetLastOperation(getLastOperationOptions *GetLastOperationOptions) - Operation response error`, func() {
		getLastOperationPath := "/drautomation/v1/last_operation/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLastOperation with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(drautomationservicev1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getLastOperationOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLastOperation(getLastOperationOptions *GetLastOperationOptions)`, func() {
		getLastOperationPath := "/drautomation/v1/last_operation/123456d3-1122-3344-b67d-4389b44b7bf9"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "crn:v1:staging:public:power-dr-automation:global:a/2c5d7270091f495795350e9adfa8399c:86e0c9a9-80f4-4fcf-88a0-07643de01bb8::", "deployment_name": "dr-deployment-instance-1", "last_updated_orchestrator_deployment_time": "2025-10-16T09:28:13.696Z", "last_updated_standby_orchestrator_deployment_time": "2025-10-16T09:28:13.696Z", "mfa_enabled": "true", "orch_ext_connectivity_status": "Connected", "orch_standby_node_addtion_status": "Completed", "orchestrator_cluster_message": "Cluster healthy", "orchestrator_config_status": "Configured", "orchestrator_ha": true, "plan_name": "DR Automation Private Plan", "primary_description": "2/5: Creating primary orchestrator VM.", "primary_ip_address": "192.168.1.10", "primary_orchestrator_status": "orchestrator-VM-creation-in-progress", "recovery_location": "us-east", "resource_group": "Default", "standby_description": "1/4: Service instance is downloading orchestrator image for standby VM creation.", "standby_ip_address": "192.168.1.11", "standby_status": "downloading-orchestrator-image", "status": "Running"}`)
				}))
			})
			It(`Invoke GetLastOperation successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(drautomationservicev1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getLastOperationOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetLastOperationWithContext(ctx, getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetLastOperationWithContext(ctx, getLastOperationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "crn:v1:staging:public:power-dr-automation:global:a/2c5d7270091f495795350e9adfa8399c:86e0c9a9-80f4-4fcf-88a0-07643de01bb8::", "deployment_name": "dr-deployment-instance-1", "last_updated_orchestrator_deployment_time": "2025-10-16T09:28:13.696Z", "last_updated_standby_orchestrator_deployment_time": "2025-10-16T09:28:13.696Z", "mfa_enabled": "true", "orch_ext_connectivity_status": "Connected", "orch_standby_node_addtion_status": "Completed", "orchestrator_cluster_message": "Cluster healthy", "orchestrator_config_status": "Configured", "orchestrator_ha": true, "plan_name": "DR Automation Private Plan", "primary_description": "2/5: Creating primary orchestrator VM.", "primary_ip_address": "192.168.1.10", "primary_orchestrator_status": "orchestrator-VM-creation-in-progress", "recovery_location": "us-east", "resource_group": "Default", "standby_description": "1/4: Service instance is downloading orchestrator image for standby VM creation.", "standby_ip_address": "192.168.1.11", "standby_status": "downloading-orchestrator-image", "status": "Running"}`)
				}))
			})
			It(`Invoke GetLastOperation successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetLastOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(drautomationservicev1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getLastOperationOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLastOperation with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(drautomationservicev1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getLastOperationOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLastOperationOptions model with no property values
				getLastOperationOptionsModelNew := new(drautomationservicev1.GetLastOperationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetLastOperation(getLastOperationOptionsModelNew)
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
			It(`Invoke GetLastOperation successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(drautomationservicev1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getLastOperationOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetLastOperation(getLastOperationOptionsModel)
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
	Describe(`ListEvents(listEventsOptions *ListEventsOptions) - Operation response error`, func() {
		listEventsPath := "/drautomation/v1/service_instances/123456d3-1122-3344-b67d-4389b44b7bf9/events"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEventsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["time"]).To(Equal([]string{"2025-06-19T23:59:59Z"}))
					Expect(req.URL.Query()["from_time"]).To(Equal([]string{"2025-06-19T00:00:00Z"}))
					Expect(req.URL.Query()["to_time"]).To(Equal([]string{"2025-06-19T23:59:59Z"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEvents with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the ListEventsOptions model
				listEventsOptionsModel := new(drautomationservicev1.ListEventsOptions)
				listEventsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				listEventsOptionsModel.Time = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.FromTime = core.StringPtr("2025-06-19T00:00:00Z")
				listEventsOptionsModel.ToTime = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.ListEvents(listEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.ListEvents(listEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEvents(listEventsOptions *ListEventsOptions)`, func() {
		listEventsPath := "/drautomation/v1/service_instances/123456d3-1122-3344-b67d-4389b44b7bf9/events"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEventsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["time"]).To(Equal([]string{"2025-06-19T23:59:59Z"}))
					Expect(req.URL.Query()["from_time"]).To(Equal([]string{"2025-06-19T00:00:00Z"}))
					Expect(req.URL.Query()["to_time"]).To(Equal([]string{"2025-06-19T23:59:59Z"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event": [{"action": "create", "api_source": "dr-automation-api", "event_id": "1cecfe43-43cd-4b1b-86be-30c2d3d2a25f", "level": "info", "message": "Service Instance created successfully", "message_data": {"anyKey": "anyValue"}, "metadata": {"anyKey": "anyValue"}, "resource": "ProvisionID", "time": "2025-06-23T07:12:49.840Z", "timestamp": "1750662769", "user": {"email": "abcuser@ibm.com", "name": "abcuser", "user_id": "IBMid-695000abc7E"}}]}`)
				}))
			})
			It(`Invoke ListEvents successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the ListEventsOptions model
				listEventsOptionsModel := new(drautomationservicev1.ListEventsOptions)
				listEventsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				listEventsOptionsModel.Time = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.FromTime = core.StringPtr("2025-06-19T00:00:00Z")
				listEventsOptionsModel.ToTime = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.ListEventsWithContext(ctx, listEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.ListEvents(listEventsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.ListEventsWithContext(ctx, listEventsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listEventsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["time"]).To(Equal([]string{"2025-06-19T23:59:59Z"}))
					Expect(req.URL.Query()["from_time"]).To(Equal([]string{"2025-06-19T00:00:00Z"}))
					Expect(req.URL.Query()["to_time"]).To(Equal([]string{"2025-06-19T23:59:59Z"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"event": [{"action": "create", "api_source": "dr-automation-api", "event_id": "1cecfe43-43cd-4b1b-86be-30c2d3d2a25f", "level": "info", "message": "Service Instance created successfully", "message_data": {"anyKey": "anyValue"}, "metadata": {"anyKey": "anyValue"}, "resource": "ProvisionID", "time": "2025-06-23T07:12:49.840Z", "timestamp": "1750662769", "user": {"email": "abcuser@ibm.com", "name": "abcuser", "user_id": "IBMid-695000abc7E"}}]}`)
				}))
			})
			It(`Invoke ListEvents successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.ListEvents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEventsOptions model
				listEventsOptionsModel := new(drautomationservicev1.ListEventsOptions)
				listEventsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				listEventsOptionsModel.Time = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.FromTime = core.StringPtr("2025-06-19T00:00:00Z")
				listEventsOptionsModel.ToTime = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.ListEvents(listEventsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEvents with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the ListEventsOptions model
				listEventsOptionsModel := new(drautomationservicev1.ListEventsOptions)
				listEventsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				listEventsOptionsModel.Time = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.FromTime = core.StringPtr("2025-06-19T00:00:00Z")
				listEventsOptionsModel.ToTime = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.ListEvents(listEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListEventsOptions model with no property values
				listEventsOptionsModelNew := new(drautomationservicev1.ListEventsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.ListEvents(listEventsOptionsModelNew)
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
			It(`Invoke ListEvents successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the ListEventsOptions model
				listEventsOptionsModel := new(drautomationservicev1.ListEventsOptions)
				listEventsOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				listEventsOptionsModel.Time = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.FromTime = core.StringPtr("2025-06-19T00:00:00Z")
				listEventsOptionsModel.ToTime = core.StringPtr("2025-06-19T23:59:59Z")
				listEventsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.ListEvents(listEventsOptionsModel)
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
	Describe(`GetEvent(getEventOptions *GetEventOptions) - Operation response error`, func() {
		getEventPath := "/drautomation/v1/service_instances/123456d3-1122-3344-b67d-4389b44b7bf9/events/00116b2a-9326-4024-839e-fb5364b76898"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEventPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEvent with error: Operation response processing error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetEventOptions model
				getEventOptionsModel := new(drautomationservicev1.GetEventOptions)
				getEventOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getEventOptionsModel.EventID = core.StringPtr("00116b2a-9326-4024-839e-fb5364b76898")
				getEventOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := drAutomationServiceService.GetEvent(getEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				drAutomationServiceService.EnableRetries(0, 0)
				result, response, operationErr = drAutomationServiceService.GetEvent(getEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEvent(getEventOptions *GetEventOptions)`, func() {
		getEventPath := "/drautomation/v1/service_instances/123456d3-1122-3344-b67d-4389b44b7bf9/events/00116b2a-9326-4024-839e-fb5364b76898"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEventPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "create", "api_source": "dr-automation-api", "event_id": "1cecfe43-43cd-4b1b-86be-30c2d3d2a25f", "level": "info", "message": "Service Instance created successfully", "message_data": {"anyKey": "anyValue"}, "metadata": {"anyKey": "anyValue"}, "resource": "ProvisionID", "time": "2025-06-23T07:12:49.840Z", "timestamp": "1750662769", "user": {"email": "abcuser@ibm.com", "name": "abcuser", "user_id": "IBMid-695000abc7E"}}`)
				}))
			})
			It(`Invoke GetEvent successfully with retries`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())
				drAutomationServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetEventOptions model
				getEventOptionsModel := new(drautomationservicev1.GetEventOptions)
				getEventOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getEventOptionsModel.EventID = core.StringPtr("00116b2a-9326-4024-839e-fb5364b76898")
				getEventOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := drAutomationServiceService.GetEventWithContext(ctx, getEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				drAutomationServiceService.DisableRetries()
				result, response, operationErr := drAutomationServiceService.GetEvent(getEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = drAutomationServiceService.GetEventWithContext(ctx, getEventOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getEventPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action": "create", "api_source": "dr-automation-api", "event_id": "1cecfe43-43cd-4b1b-86be-30c2d3d2a25f", "level": "info", "message": "Service Instance created successfully", "message_data": {"anyKey": "anyValue"}, "metadata": {"anyKey": "anyValue"}, "resource": "ProvisionID", "time": "2025-06-23T07:12:49.840Z", "timestamp": "1750662769", "user": {"email": "abcuser@ibm.com", "name": "abcuser", "user_id": "IBMid-695000abc7E"}}`)
				}))
			})
			It(`Invoke GetEvent successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := drAutomationServiceService.GetEvent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEventOptions model
				getEventOptionsModel := new(drautomationservicev1.GetEventOptions)
				getEventOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getEventOptionsModel.EventID = core.StringPtr("00116b2a-9326-4024-839e-fb5364b76898")
				getEventOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = drAutomationServiceService.GetEvent(getEventOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEvent with error: Operation validation and request error`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetEventOptions model
				getEventOptionsModel := new(drautomationservicev1.GetEventOptions)
				getEventOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getEventOptionsModel.EventID = core.StringPtr("00116b2a-9326-4024-839e-fb5364b76898")
				getEventOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := drAutomationServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := drAutomationServiceService.GetEvent(getEventOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEventOptions model with no property values
				getEventOptionsModelNew := new(drautomationservicev1.GetEventOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = drAutomationServiceService.GetEvent(getEventOptionsModelNew)
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
			It(`Invoke GetEvent successfully`, func() {
				drAutomationServiceService, serviceErr := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(drAutomationServiceService).ToNot(BeNil())

				// Construct an instance of the GetEventOptions model
				getEventOptionsModel := new(drautomationservicev1.GetEventOptions)
				getEventOptionsModel.InstanceID = core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")
				getEventOptionsModel.EventID = core.StringPtr("00116b2a-9326-4024-839e-fb5364b76898")
				getEventOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getEventOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := drAutomationServiceService.GetEvent(getEventOptionsModel)
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
			drAutomationServiceService, _ := drautomationservicev1.NewDrAutomationServiceV1(&drautomationservicev1.DrAutomationServiceV1Options{
				URL:           "http://drautomationservicev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateApikeyOptions successfully`, func() {
				// Construct an instance of the CreateApikeyOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				createApikeyOptionsAPIKey := "abcdefrg_izklmnop_fxbEED"
				createApikeyOptionsModel := drAutomationServiceService.NewCreateApikeyOptions(instanceID, createApikeyOptionsAPIKey)
				createApikeyOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				createApikeyOptionsModel.SetAPIKey("abcdefrg_izklmnop_fxbEED")
				createApikeyOptionsModel.SetAcceptLanguage("testString")
				createApikeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createApikeyOptionsModel).ToNot(BeNil())
				Expect(createApikeyOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(createApikeyOptionsModel.APIKey).To(Equal(core.StringPtr("abcdefrg_izklmnop_fxbEED")))
				Expect(createApikeyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createApikeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateManageDrOptions successfully`, func() {
				// Construct an instance of the CreateManageDrOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				createManageDrOptionsLocationID := "dal10"
				createManageDrOptionsMachineType := "bx2-4x16"
				createManageDrOptionsOrchestratorLocationType := "off-premises"
				createManageDrOptionsOrchestratorName := "adminUser"
				createManageDrOptionsOrchestratorPassword := "testString"
				createManageDrOptionsOrchestratorWorkspaceID := "orch-workspace-01"
				createManageDrOptionsModel := drAutomationServiceService.NewCreateManageDrOptions(instanceID, createManageDrOptionsLocationID, createManageDrOptionsMachineType, createManageDrOptionsOrchestratorLocationType, createManageDrOptionsOrchestratorName, createManageDrOptionsOrchestratorPassword, createManageDrOptionsOrchestratorWorkspaceID)
				createManageDrOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				createManageDrOptionsModel.SetLocationID("dal10")
				createManageDrOptionsModel.SetMachineType("bx2-4x16")
				createManageDrOptionsModel.SetOrchestratorLocationType("off-premises")
				createManageDrOptionsModel.SetOrchestratorName("adminUser")
				createManageDrOptionsModel.SetOrchestratorPassword("testString")
				createManageDrOptionsModel.SetOrchestratorWorkspaceID("orch-workspace-01")
				createManageDrOptionsModel.SetAPIKey("testString")
				createManageDrOptionsModel.SetClientID("abcd-97d2-1234-bf62-8eaecc67a1234")
				createManageDrOptionsModel.SetClientSecret("abcd1234xM1y123wK6qR9123456789bE2jG0pabcdefgh")
				createManageDrOptionsModel.SetGUID("123e4567-e89b-12d3-a456-426614174000")
				createManageDrOptionsModel.SetOrchestratorHa(true)
				createManageDrOptionsModel.SetProxyIP("10.40.30.10:8888")
				createManageDrOptionsModel.SetRegionID("us-south")
				createManageDrOptionsModel.SetResourceInstance("crn:v1:bluemix:public:resource-controller::res123")
				createManageDrOptionsModel.SetSecret("testString")
				createManageDrOptionsModel.SetSecretGroup("default-secret-group")
				createManageDrOptionsModel.SetSSHKeyName("my-ssh-key")
				createManageDrOptionsModel.SetStandbyMachineType("bx2-8x32")
				createManageDrOptionsModel.SetStandbyOrchestratorName("standbyAdmin")
				createManageDrOptionsModel.SetStandbyOrchestratorWorkspaceID("orch-standby-02")
				createManageDrOptionsModel.SetStandbyTier("Premium")
				createManageDrOptionsModel.SetTenantName("xxx.ibm.com")
				createManageDrOptionsModel.SetTier("Standard")
				createManageDrOptionsModel.SetStandByRedeploy("testString")
				createManageDrOptionsModel.SetAcceptLanguage("testString")
				createManageDrOptionsModel.SetAcceptsIncomplete(true)
				createManageDrOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createManageDrOptionsModel).ToNot(BeNil())
				Expect(createManageDrOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(createManageDrOptionsModel.LocationID).To(Equal(core.StringPtr("dal10")))
				Expect(createManageDrOptionsModel.MachineType).To(Equal(core.StringPtr("bx2-4x16")))
				Expect(createManageDrOptionsModel.OrchestratorLocationType).To(Equal(core.StringPtr("off-premises")))
				Expect(createManageDrOptionsModel.OrchestratorName).To(Equal(core.StringPtr("adminUser")))
				Expect(createManageDrOptionsModel.OrchestratorPassword).To(Equal(core.StringPtr("testString")))
				Expect(createManageDrOptionsModel.OrchestratorWorkspaceID).To(Equal(core.StringPtr("orch-workspace-01")))
				Expect(createManageDrOptionsModel.APIKey).To(Equal(core.StringPtr("testString")))
				Expect(createManageDrOptionsModel.ClientID).To(Equal(core.StringPtr("abcd-97d2-1234-bf62-8eaecc67a1234")))
				Expect(createManageDrOptionsModel.ClientSecret).To(Equal(core.StringPtr("abcd1234xM1y123wK6qR9123456789bE2jG0pabcdefgh")))
				Expect(createManageDrOptionsModel.GUID).To(Equal(core.StringPtr("123e4567-e89b-12d3-a456-426614174000")))
				Expect(createManageDrOptionsModel.OrchestratorHa).To(Equal(core.BoolPtr(true)))
				Expect(createManageDrOptionsModel.ProxyIP).To(Equal(core.StringPtr("10.40.30.10:8888")))
				Expect(createManageDrOptionsModel.RegionID).To(Equal(core.StringPtr("us-south")))
				Expect(createManageDrOptionsModel.ResourceInstance).To(Equal(core.StringPtr("crn:v1:bluemix:public:resource-controller::res123")))
				Expect(createManageDrOptionsModel.Secret).To(Equal(core.StringPtr("testString")))
				Expect(createManageDrOptionsModel.SecretGroup).To(Equal(core.StringPtr("default-secret-group")))
				Expect(createManageDrOptionsModel.SSHKeyName).To(Equal(core.StringPtr("my-ssh-key")))
				Expect(createManageDrOptionsModel.StandbyMachineType).To(Equal(core.StringPtr("bx2-8x32")))
				Expect(createManageDrOptionsModel.StandbyOrchestratorName).To(Equal(core.StringPtr("standbyAdmin")))
				Expect(createManageDrOptionsModel.StandbyOrchestratorWorkspaceID).To(Equal(core.StringPtr("orch-standby-02")))
				Expect(createManageDrOptionsModel.StandbyTier).To(Equal(core.StringPtr("Premium")))
				Expect(createManageDrOptionsModel.TenantName).To(Equal(core.StringPtr("xxx.ibm.com")))
				Expect(createManageDrOptionsModel.Tier).To(Equal(core.StringPtr("Standard")))
				Expect(createManageDrOptionsModel.StandByRedeploy).To(Equal(core.StringPtr("testString")))
				Expect(createManageDrOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createManageDrOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(true)))
				Expect(createManageDrOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetApikeyOptions successfully`, func() {
				// Construct an instance of the GetApikeyOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				getApikeyOptionsModel := drAutomationServiceService.NewGetApikeyOptions(instanceID)
				getApikeyOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getApikeyOptionsModel.SetAcceptLanguage("testString")
				getApikeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApikeyOptionsModel).ToNot(BeNil())
				Expect(getApikeyOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getApikeyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getApikeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDrGrsLocationPairOptions successfully`, func() {
				// Construct an instance of the GetDrGrsLocationPairOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				getDrGrsLocationPairOptionsModel := drAutomationServiceService.NewGetDrGrsLocationPairOptions(instanceID)
				getDrGrsLocationPairOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrGrsLocationPairOptionsModel.SetAcceptLanguage("testString")
				getDrGrsLocationPairOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDrGrsLocationPairOptionsModel).ToNot(BeNil())
				Expect(getDrGrsLocationPairOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getDrGrsLocationPairOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDrGrsLocationPairOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDrLocationsOptions successfully`, func() {
				// Construct an instance of the GetDrLocationsOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				getDrLocationsOptionsModel := drAutomationServiceService.NewGetDrLocationsOptions(instanceID)
				getDrLocationsOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrLocationsOptionsModel.SetAcceptLanguage("testString")
				getDrLocationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDrLocationsOptionsModel).ToNot(BeNil())
				Expect(getDrLocationsOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getDrLocationsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDrLocationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDrManagedVMOptions successfully`, func() {
				// Construct an instance of the GetDrManagedVMOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				getDrManagedVMOptionsModel := drAutomationServiceService.NewGetDrManagedVMOptions(instanceID)
				getDrManagedVMOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrManagedVMOptionsModel.SetAcceptLanguage("testString")
				getDrManagedVMOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDrManagedVMOptionsModel).ToNot(BeNil())
				Expect(getDrManagedVMOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getDrManagedVMOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDrManagedVMOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDrSummaryOptions successfully`, func() {
				// Construct an instance of the GetDrSummaryOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				getDrSummaryOptionsModel := drAutomationServiceService.NewGetDrSummaryOptions(instanceID)
				getDrSummaryOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getDrSummaryOptionsModel.SetAcceptLanguage("testString")
				getDrSummaryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDrSummaryOptionsModel).ToNot(BeNil())
				Expect(getDrSummaryOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getDrSummaryOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDrSummaryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEventOptions successfully`, func() {
				// Construct an instance of the GetEventOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				eventID := "00116b2a-9326-4024-839e-fb5364b76898"
				getEventOptionsModel := drAutomationServiceService.NewGetEventOptions(instanceID, eventID)
				getEventOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getEventOptionsModel.SetEventID("00116b2a-9326-4024-839e-fb5364b76898")
				getEventOptionsModel.SetAcceptLanguage("testString")
				getEventOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEventOptionsModel).ToNot(BeNil())
				Expect(getEventOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getEventOptionsModel.EventID).To(Equal(core.StringPtr("00116b2a-9326-4024-839e-fb5364b76898")))
				Expect(getEventOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getEventOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLastOperationOptions successfully`, func() {
				// Construct an instance of the GetLastOperationOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				getLastOperationOptionsModel := drAutomationServiceService.NewGetLastOperationOptions(instanceID)
				getLastOperationOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getLastOperationOptionsModel.SetAcceptLanguage("testString")
				getLastOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLastOperationOptionsModel).ToNot(BeNil())
				Expect(getLastOperationOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getLastOperationOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getLastOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMachineTypeOptions successfully`, func() {
				// Construct an instance of the GetMachineTypeOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				primaryWorkspaceName := "Test-workspace-wdc06"
				getMachineTypeOptionsModel := drAutomationServiceService.NewGetMachineTypeOptions(instanceID, primaryWorkspaceName)
				getMachineTypeOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getMachineTypeOptionsModel.SetPrimaryWorkspaceName("Test-workspace-wdc06")
				getMachineTypeOptionsModel.SetAcceptLanguage("testString")
				getMachineTypeOptionsModel.SetStandbyWorkspaceName("Test-workspace-wdc07")
				getMachineTypeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMachineTypeOptionsModel).ToNot(BeNil())
				Expect(getMachineTypeOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getMachineTypeOptionsModel.PrimaryWorkspaceName).To(Equal(core.StringPtr("Test-workspace-wdc06")))
				Expect(getMachineTypeOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getMachineTypeOptionsModel.StandbyWorkspaceName).To(Equal(core.StringPtr("Test-workspace-wdc07")))
				Expect(getMachineTypeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetManageDrOptions successfully`, func() {
				// Construct an instance of the GetManageDrOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				getManageDrOptionsModel := drAutomationServiceService.NewGetManageDrOptions(instanceID)
				getManageDrOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getManageDrOptionsModel.SetAcceptLanguage("testString")
				getManageDrOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getManageDrOptionsModel).ToNot(BeNil())
				Expect(getManageDrOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getManageDrOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getManageDrOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPowervsWorkspacesOptions successfully`, func() {
				// Construct an instance of the GetPowervsWorkspacesOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				locationID := "testString"
				getPowervsWorkspacesOptionsModel := drAutomationServiceService.NewGetPowervsWorkspacesOptions(instanceID, locationID)
				getPowervsWorkspacesOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				getPowervsWorkspacesOptionsModel.SetLocationID("testString")
				getPowervsWorkspacesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPowervsWorkspacesOptionsModel).ToNot(BeNil())
				Expect(getPowervsWorkspacesOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(getPowervsWorkspacesOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(getPowervsWorkspacesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEventsOptions successfully`, func() {
				// Construct an instance of the ListEventsOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				listEventsOptionsModel := drAutomationServiceService.NewListEventsOptions(instanceID)
				listEventsOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				listEventsOptionsModel.SetTime("2025-06-19T23:59:59Z")
				listEventsOptionsModel.SetFromTime("2025-06-19T00:00:00Z")
				listEventsOptionsModel.SetToTime("2025-06-19T23:59:59Z")
				listEventsOptionsModel.SetAcceptLanguage("testString")
				listEventsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEventsOptionsModel).ToNot(BeNil())
				Expect(listEventsOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(listEventsOptionsModel.Time).To(Equal(core.StringPtr("2025-06-19T23:59:59Z")))
				Expect(listEventsOptionsModel.FromTime).To(Equal(core.StringPtr("2025-06-19T00:00:00Z")))
				Expect(listEventsOptionsModel.ToTime).To(Equal(core.StringPtr("2025-06-19T23:59:59Z")))
				Expect(listEventsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listEventsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateApikeyOptions successfully`, func() {
				// Construct an instance of the UpdateApikeyOptions model
				instanceID := "123456d3-1122-3344-b67d-4389b44b7bf9"
				updateApikeyOptionsAPIKey := "adfadfdsafsdfdsf"
				updateApikeyOptionsModel := drAutomationServiceService.NewUpdateApikeyOptions(instanceID, updateApikeyOptionsAPIKey)
				updateApikeyOptionsModel.SetInstanceID("123456d3-1122-3344-b67d-4389b44b7bf9")
				updateApikeyOptionsModel.SetAPIKey("adfadfdsafsdfdsf")
				updateApikeyOptionsModel.SetAcceptLanguage("testString")
				updateApikeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateApikeyOptionsModel).ToNot(BeNil())
				Expect(updateApikeyOptionsModel.InstanceID).To(Equal(core.StringPtr("123456d3-1122-3344-b67d-4389b44b7bf9")))
				Expect(updateApikeyOptionsModel.APIKey).To(Equal(core.StringPtr("adfadfdsafsdfdsf")))
				Expect(updateApikeyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(updateApikeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
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
