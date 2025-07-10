/**
 * (C) Copyright IBM Corp. 2024.
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

package contextbasedrestrictionsv1_test

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
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ContextBasedRestrictionsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(contextBasedRestrictionsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(contextBasedRestrictionsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
				URL: "https://contextbasedrestrictionsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(contextBasedRestrictionsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONTEXT_BASED_RESTRICTIONS_URL":       "https://contextbasedrestrictionsv1/api",
				"CONTEXT_BASED_RESTRICTIONS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1UsingExternalConfig(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{})
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := contextBasedRestrictionsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != contextBasedRestrictionsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(contextBasedRestrictionsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(contextBasedRestrictionsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1UsingExternalConfig(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL: "https://testService/api",
				})
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := contextBasedRestrictionsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != contextBasedRestrictionsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(contextBasedRestrictionsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(contextBasedRestrictionsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1UsingExternalConfig(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{})
				err := contextBasedRestrictionsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := contextBasedRestrictionsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != contextBasedRestrictionsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(contextBasedRestrictionsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(contextBasedRestrictionsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONTEXT_BASED_RESTRICTIONS_URL":       "https://contextbasedrestrictionsv1/api",
				"CONTEXT_BASED_RESTRICTIONS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1UsingExternalConfig(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(contextBasedRestrictionsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CONTEXT_BASED_RESTRICTIONS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1UsingExternalConfig(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(contextBasedRestrictionsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = contextbasedrestrictionsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateZone(createZoneOptions *CreateZoneOptions) - Operation response error`, func() {
		createZonePath := "/v1/zones"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZonePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZone with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(contextbasedrestrictionsv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("an example of zone")
				createZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				createZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				createZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				createZoneOptionsModel.TransactionID = core.StringPtr("testString")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.CreateZone(createZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.CreateZone(createZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateZone(createZoneOptions *CreateZoneOptions)`, func() {
		createZonePath := "/v1/zones"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZonePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "address_count": 12, "excluded_count": 13, "name": "Name", "account_id": "AccountID", "description": "Description", "addresses": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "excluded": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateZone successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(contextbasedrestrictionsv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("an example of zone")
				createZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				createZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				createZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				createZoneOptionsModel.TransactionID = core.StringPtr("testString")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.CreateZoneWithContext(ctx, createZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.CreateZone(createZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.CreateZoneWithContext(ctx, createZoneOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createZonePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "address_count": 12, "excluded_count": 13, "name": "Name", "account_id": "AccountID", "description": "Description", "addresses": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "excluded": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateZone successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.CreateZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(contextbasedrestrictionsv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("an example of zone")
				createZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				createZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				createZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				createZoneOptionsModel.TransactionID = core.StringPtr("testString")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.CreateZone(createZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateZone with error: Operation request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(contextbasedrestrictionsv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("an example of zone")
				createZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				createZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				createZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				createZoneOptionsModel.TransactionID = core.StringPtr("testString")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.CreateZone(createZoneOptionsModel)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateZone successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(contextbasedrestrictionsv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("an example of zone")
				createZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				createZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				createZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				createZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				createZoneOptionsModel.TransactionID = core.StringPtr("testString")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.CreateZone(createZoneOptionsModel)
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
	Describe(`ListZones(listZonesOptions *ListZonesOptions) - Operation response error`, func() {
		listZonesPath := "/v1/zones"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZonesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListZones with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(contextbasedrestrictionsv1.ListZonesOptions)
				listZonesOptionsModel.AccountID = core.StringPtr("testString")
				listZonesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZonesOptionsModel.TransactionID = core.StringPtr("testString")
				listZonesOptionsModel.Name = core.StringPtr("testString")
				listZonesOptionsModel.Sort = core.StringPtr("testString")
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.ListZones(listZonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.ListZones(listZonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListZones(listZonesOptions *ListZonesOptions)`, func() {
		listZonesPath := "/v1/zones"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listZonesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "zones": [{"id": "ID", "crn": "CRN", "name": "Name", "description": "Description", "addresses_preview": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "address_count": 12, "excluded_count": 13, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListZones successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(contextbasedrestrictionsv1.ListZonesOptions)
				listZonesOptionsModel.AccountID = core.StringPtr("testString")
				listZonesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZonesOptionsModel.TransactionID = core.StringPtr("testString")
				listZonesOptionsModel.Name = core.StringPtr("testString")
				listZonesOptionsModel.Sort = core.StringPtr("testString")
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.ListZonesWithContext(ctx, listZonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.ListZones(listZonesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.ListZonesWithContext(ctx, listZonesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listZonesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "zones": [{"id": "ID", "crn": "CRN", "name": "Name", "description": "Description", "addresses_preview": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "address_count": 12, "excluded_count": 13, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListZones successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.ListZones(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(contextbasedrestrictionsv1.ListZonesOptions)
				listZonesOptionsModel.AccountID = core.StringPtr("testString")
				listZonesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZonesOptionsModel.TransactionID = core.StringPtr("testString")
				listZonesOptionsModel.Name = core.StringPtr("testString")
				listZonesOptionsModel.Sort = core.StringPtr("testString")
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.ListZones(listZonesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListZones with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(contextbasedrestrictionsv1.ListZonesOptions)
				listZonesOptionsModel.AccountID = core.StringPtr("testString")
				listZonesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZonesOptionsModel.TransactionID = core.StringPtr("testString")
				listZonesOptionsModel.Name = core.StringPtr("testString")
				listZonesOptionsModel.Sort = core.StringPtr("testString")
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.ListZones(listZonesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListZonesOptions model with no property values
				listZonesOptionsModelNew := new(contextbasedrestrictionsv1.ListZonesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = contextBasedRestrictionsService.ListZones(listZonesOptionsModelNew)
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
			It(`Invoke ListZones successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(contextbasedrestrictionsv1.ListZonesOptions)
				listZonesOptionsModel.AccountID = core.StringPtr("testString")
				listZonesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listZonesOptionsModel.TransactionID = core.StringPtr("testString")
				listZonesOptionsModel.Name = core.StringPtr("testString")
				listZonesOptionsModel.Sort = core.StringPtr("testString")
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.ListZones(listZonesOptionsModel)
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
	Describe(`GetZone(getZoneOptions *GetZoneOptions) - Operation response error`, func() {
		getZonePath := "/v1/zones/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZonePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZone with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(contextbasedrestrictionsv1.GetZoneOptions)
				getZoneOptionsModel.ZoneID = core.StringPtr("testString")
				getZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOptionsModel.TransactionID = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.GetZone(getZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.GetZone(getZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZone(getZoneOptions *GetZoneOptions)`, func() {
		getZonePath := "/v1/zones/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZonePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "address_count": 12, "excluded_count": 13, "name": "Name", "account_id": "AccountID", "description": "Description", "addresses": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "excluded": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetZone successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(contextbasedrestrictionsv1.GetZoneOptions)
				getZoneOptionsModel.ZoneID = core.StringPtr("testString")
				getZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOptionsModel.TransactionID = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.GetZoneWithContext(ctx, getZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.GetZone(getZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.GetZoneWithContext(ctx, getZoneOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getZonePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "address_count": 12, "excluded_count": 13, "name": "Name", "account_id": "AccountID", "description": "Description", "addresses": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "excluded": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetZone successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.GetZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(contextbasedrestrictionsv1.GetZoneOptions)
				getZoneOptionsModel.ZoneID = core.StringPtr("testString")
				getZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOptionsModel.TransactionID = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.GetZone(getZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetZone with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(contextbasedrestrictionsv1.GetZoneOptions)
				getZoneOptionsModel.ZoneID = core.StringPtr("testString")
				getZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOptionsModel.TransactionID = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.GetZone(getZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneOptions model with no property values
				getZoneOptionsModelNew := new(contextbasedrestrictionsv1.GetZoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = contextBasedRestrictionsService.GetZone(getZoneOptionsModelNew)
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
			It(`Invoke GetZone successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(contextbasedrestrictionsv1.GetZoneOptions)
				getZoneOptionsModel.ZoneID = core.StringPtr("testString")
				getZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				getZoneOptionsModel.TransactionID = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.GetZone(getZoneOptionsModel)
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
	Describe(`ReplaceZone(replaceZoneOptions *ReplaceZoneOptions) - Operation response error`, func() {
		replaceZonePath := "/v1/zones/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceZonePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceZone with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the ReplaceZoneOptions model
				replaceZoneOptionsModel := new(contextbasedrestrictionsv1.ReplaceZoneOptions)
				replaceZoneOptionsModel.ZoneID = core.StringPtr("testString")
				replaceZoneOptionsModel.IfMatch = core.StringPtr("testString")
				replaceZoneOptionsModel.Name = core.StringPtr("an example of zone")
				replaceZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				replaceZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				replaceZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceZoneOptionsModel.TransactionID = core.StringPtr("testString")
				replaceZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.ReplaceZone(replaceZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceZone(replaceZoneOptions *ReplaceZoneOptions)`, func() {
		replaceZonePath := "/v1/zones/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceZonePath))
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
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "address_count": 12, "excluded_count": 13, "name": "Name", "account_id": "AccountID", "description": "Description", "addresses": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "excluded": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplaceZone successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the ReplaceZoneOptions model
				replaceZoneOptionsModel := new(contextbasedrestrictionsv1.ReplaceZoneOptions)
				replaceZoneOptionsModel.ZoneID = core.StringPtr("testString")
				replaceZoneOptionsModel.IfMatch = core.StringPtr("testString")
				replaceZoneOptionsModel.Name = core.StringPtr("an example of zone")
				replaceZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				replaceZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				replaceZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceZoneOptionsModel.TransactionID = core.StringPtr("testString")
				replaceZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.ReplaceZoneWithContext(ctx, replaceZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.ReplaceZoneWithContext(ctx, replaceZoneOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceZonePath))
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
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "address_count": 12, "excluded_count": 13, "name": "Name", "account_id": "AccountID", "description": "Description", "addresses": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "excluded": [{"type": "ipAddress", "value": "Value", "id": "ID"}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplaceZone successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.ReplaceZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the ReplaceZoneOptions model
				replaceZoneOptionsModel := new(contextbasedrestrictionsv1.ReplaceZoneOptions)
				replaceZoneOptionsModel.ZoneID = core.StringPtr("testString")
				replaceZoneOptionsModel.IfMatch = core.StringPtr("testString")
				replaceZoneOptionsModel.Name = core.StringPtr("an example of zone")
				replaceZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				replaceZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				replaceZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceZoneOptionsModel.TransactionID = core.StringPtr("testString")
				replaceZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.ReplaceZone(replaceZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceZone with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the ReplaceZoneOptions model
				replaceZoneOptionsModel := new(contextbasedrestrictionsv1.ReplaceZoneOptions)
				replaceZoneOptionsModel.ZoneID = core.StringPtr("testString")
				replaceZoneOptionsModel.IfMatch = core.StringPtr("testString")
				replaceZoneOptionsModel.Name = core.StringPtr("an example of zone")
				replaceZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				replaceZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				replaceZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceZoneOptionsModel.TransactionID = core.StringPtr("testString")
				replaceZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceZoneOptions model with no property values
				replaceZoneOptionsModelNew := new(contextbasedrestrictionsv1.ReplaceZoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = contextBasedRestrictionsService.ReplaceZone(replaceZoneOptionsModelNew)
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
			It(`Invoke ReplaceZone successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")

				// Construct an instance of the ReplaceZoneOptions model
				replaceZoneOptionsModel := new(contextbasedrestrictionsv1.ReplaceZoneOptions)
				replaceZoneOptionsModel.ZoneID = core.StringPtr("testString")
				replaceZoneOptionsModel.IfMatch = core.StringPtr("testString")
				replaceZoneOptionsModel.Name = core.StringPtr("an example of zone")
				replaceZoneOptionsModel.AccountID = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				replaceZoneOptionsModel.Description = core.StringPtr("this is an example of zone")
				replaceZoneOptionsModel.Addresses = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.Excluded = []contextbasedrestrictionsv1.AddressIntf{addressModel}
				replaceZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceZoneOptionsModel.TransactionID = core.StringPtr("testString")
				replaceZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptionsModel)
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
	Describe(`DeleteZone(deleteZoneOptions *DeleteZoneOptions)`, func() {
		deleteZonePath := "/v1/zones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZonePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteZone successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := contextBasedRestrictionsService.DeleteZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(contextbasedrestrictionsv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneID = core.StringPtr("testString")
				deleteZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteZoneOptionsModel.TransactionID = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = contextBasedRestrictionsService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteZone with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(contextbasedrestrictionsv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneID = core.StringPtr("testString")
				deleteZoneOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteZoneOptionsModel.TransactionID = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := contextBasedRestrictionsService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteZoneOptions model with no property values
				deleteZoneOptionsModelNew := new(contextbasedrestrictionsv1.DeleteZoneOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = contextBasedRestrictionsService.DeleteZone(deleteZoneOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAvailableServicerefTargets(listAvailableServicerefTargetsOptions *ListAvailableServicerefTargetsOptions) - Operation response error`, func() {
		listAvailableServicerefTargetsPath := "/v1/zones/serviceref_targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAvailableServicerefTargetsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"all"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAvailableServicerefTargets with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListAvailableServicerefTargetsOptions model
				listAvailableServicerefTargetsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServicerefTargetsOptions)
				listAvailableServicerefTargetsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.Type = core.StringPtr("all")
				listAvailableServicerefTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServicerefTargets(listAvailableServicerefTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.ListAvailableServicerefTargets(listAvailableServicerefTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAvailableServicerefTargets(listAvailableServicerefTargetsOptions *ListAvailableServicerefTargetsOptions)`, func() {
		listAvailableServicerefTargetsPath := "/v1/zones/serviceref_targets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAvailableServicerefTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"all"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "targets": [{"service_name": "ServiceName", "service_type": "ServiceType", "locations": [{"display_name": "DisplayName", "kind": "Kind", "name": "Name"}]}]}`)
				}))
			})
			It(`Invoke ListAvailableServicerefTargets successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the ListAvailableServicerefTargetsOptions model
				listAvailableServicerefTargetsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServicerefTargetsOptions)
				listAvailableServicerefTargetsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.Type = core.StringPtr("all")
				listAvailableServicerefTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.ListAvailableServicerefTargetsWithContext(ctx, listAvailableServicerefTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServicerefTargets(listAvailableServicerefTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.ListAvailableServicerefTargetsWithContext(ctx, listAvailableServicerefTargetsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAvailableServicerefTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"all"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "targets": [{"service_name": "ServiceName", "service_type": "ServiceType", "locations": [{"display_name": "DisplayName", "kind": "Kind", "name": "Name"}]}]}`)
				}))
			})
			It(`Invoke ListAvailableServicerefTargets successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServicerefTargets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAvailableServicerefTargetsOptions model
				listAvailableServicerefTargetsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServicerefTargetsOptions)
				listAvailableServicerefTargetsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.Type = core.StringPtr("all")
				listAvailableServicerefTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.ListAvailableServicerefTargets(listAvailableServicerefTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAvailableServicerefTargets with error: Operation request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListAvailableServicerefTargetsOptions model
				listAvailableServicerefTargetsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServicerefTargetsOptions)
				listAvailableServicerefTargetsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.Type = core.StringPtr("all")
				listAvailableServicerefTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServicerefTargets(listAvailableServicerefTargetsOptionsModel)
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
			It(`Invoke ListAvailableServicerefTargets successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListAvailableServicerefTargetsOptions model
				listAvailableServicerefTargetsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServicerefTargetsOptions)
				listAvailableServicerefTargetsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServicerefTargetsOptionsModel.Type = core.StringPtr("all")
				listAvailableServicerefTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServicerefTargets(listAvailableServicerefTargetsOptionsModel)
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
	Describe(`GetServicerefTarget(getServicerefTargetOptions *GetServicerefTargetOptions) - Operation response error`, func() {
		getServicerefTargetPath := "/v1/zones/serviceref_targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServicerefTargetPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetServicerefTarget with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetServicerefTargetOptions model
				getServicerefTargetOptionsModel := new(contextbasedrestrictionsv1.GetServicerefTargetOptions)
				getServicerefTargetOptionsModel.ServiceName = core.StringPtr("testString")
				getServicerefTargetOptionsModel.XCorrelationID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.TransactionID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetServicerefTarget(getServicerefTargetOptions *GetServicerefTargetOptions)`, func() {
		getServicerefTargetPath := "/v1/zones/serviceref_targets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServicerefTargetPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"service_name": "ServiceName", "service_type": "ServiceType", "locations": [{"display_name": "DisplayName", "kind": "Kind", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetServicerefTarget successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the GetServicerefTargetOptions model
				getServicerefTargetOptionsModel := new(contextbasedrestrictionsv1.GetServicerefTargetOptions)
				getServicerefTargetOptionsModel.ServiceName = core.StringPtr("testString")
				getServicerefTargetOptionsModel.XCorrelationID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.TransactionID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.GetServicerefTargetWithContext(ctx, getServicerefTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.GetServicerefTargetWithContext(ctx, getServicerefTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getServicerefTargetPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"service_name": "ServiceName", "service_type": "ServiceType", "locations": [{"display_name": "DisplayName", "kind": "Kind", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetServicerefTarget successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.GetServicerefTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServicerefTargetOptions model
				getServicerefTargetOptionsModel := new(contextbasedrestrictionsv1.GetServicerefTargetOptions)
				getServicerefTargetOptionsModel.ServiceName = core.StringPtr("testString")
				getServicerefTargetOptionsModel.XCorrelationID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.TransactionID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetServicerefTarget with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetServicerefTargetOptions model
				getServicerefTargetOptionsModel := new(contextbasedrestrictionsv1.GetServicerefTargetOptions)
				getServicerefTargetOptionsModel.ServiceName = core.StringPtr("testString")
				getServicerefTargetOptionsModel.XCorrelationID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.TransactionID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetServicerefTargetOptions model with no property values
				getServicerefTargetOptionsModelNew := new(contextbasedrestrictionsv1.GetServicerefTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptionsModelNew)
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
			It(`Invoke GetServicerefTarget successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetServicerefTargetOptions model
				getServicerefTargetOptionsModel := new(contextbasedrestrictionsv1.GetServicerefTargetOptions)
				getServicerefTargetOptionsModel.ServiceName = core.StringPtr("testString")
				getServicerefTargetOptionsModel.XCorrelationID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.TransactionID = core.StringPtr("testString")
				getServicerefTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptionsModel)
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
	Describe(`CreateRule(createRuleOptions *CreateRuleOptions) - Operation response error`, func() {
		createRulePath := "/v1/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRule with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("65810ac762004f22ac19f8f8edf70a34")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(contextbasedrestrictionsv1.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				createRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				createRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				createRuleOptionsModel.Operations = newRuleOperationsModel
				createRuleOptionsModel.EnforcementMode = core.StringPtr("enabled")
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
		createRulePath := "/v1/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRulePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "description": "Description", "contexts": [{"attributes": [{"name": "Name", "value": "Value"}]}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "operations": {"api_types": [{"api_type_id": "APITypeID"}]}, "enforcement_mode": "enabled", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateRule successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("65810ac762004f22ac19f8f8edf70a34")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(contextbasedrestrictionsv1.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				createRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				createRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				createRuleOptionsModel.Operations = newRuleOperationsModel
				createRuleOptionsModel.EnforcementMode = core.StringPtr("enabled")
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.CreateRuleWithContext(ctx, createRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.CreateRuleWithContext(ctx, createRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createRulePath))
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "description": "Description", "contexts": [{"attributes": [{"name": "Name", "value": "Value"}]}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "operations": {"api_types": [{"api_type_id": "APITypeID"}]}, "enforcement_mode": "enabled", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreateRule successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.CreateRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("65810ac762004f22ac19f8f8edf70a34")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(contextbasedrestrictionsv1.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				createRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				createRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				createRuleOptionsModel.Operations = newRuleOperationsModel
				createRuleOptionsModel.EnforcementMode = core.StringPtr("enabled")
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.CreateRule(createRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRule with error: Operation request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("65810ac762004f22ac19f8f8edf70a34")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(contextbasedrestrictionsv1.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				createRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				createRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				createRuleOptionsModel.Operations = newRuleOperationsModel
				createRuleOptionsModel.EnforcementMode = core.StringPtr("enabled")
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.CreateRule(createRuleOptionsModel)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateRule successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("65810ac762004f22ac19f8f8edf70a34")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := new(contextbasedrestrictionsv1.CreateRuleOptions)
				createRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				createRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				createRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				createRuleOptionsModel.Operations = newRuleOperationsModel
				createRuleOptionsModel.EnforcementMode = core.StringPtr("enabled")
				createRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				createRuleOptionsModel.TransactionID = core.StringPtr("testString")
				createRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.CreateRule(createRuleOptionsModel)
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
	Describe(`ListRules(listRulesOptions *ListRulesOptions) - Operation response error`, func() {
		listRulesPath := "/v1/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_instance"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["zone_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["enforcement_mode"]).To(Equal([]string{"enabled"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRules with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(contextbasedrestrictionsv1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("testString")
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Region = core.StringPtr("testString")
				listRulesOptionsModel.Resource = core.StringPtr("testString")
				listRulesOptionsModel.ResourceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceInstance = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.ServiceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listRulesOptionsModel.ZoneID = core.StringPtr("testString")
				listRulesOptionsModel.Sort = core.StringPtr("testString")
				listRulesOptionsModel.EnforcementMode = core.StringPtr("enabled")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRules(listRulesOptions *ListRulesOptions)`, func() {
		listRulesPath := "/v1/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_instance"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["zone_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["enforcement_mode"]).To(Equal([]string{"enabled"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "rules": [{"id": "ID", "crn": "CRN", "description": "Description", "contexts": [{"attributes": [{"name": "Name", "value": "Value"}]}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "operations": {"api_types": [{"api_type_id": "APITypeID"}]}, "enforcement_mode": "enabled", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListRules successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(contextbasedrestrictionsv1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("testString")
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Region = core.StringPtr("testString")
				listRulesOptionsModel.Resource = core.StringPtr("testString")
				listRulesOptionsModel.ResourceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceInstance = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.ServiceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listRulesOptionsModel.ZoneID = core.StringPtr("testString")
				listRulesOptionsModel.Sort = core.StringPtr("testString")
				listRulesOptionsModel.EnforcementMode = core.StringPtr("enabled")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.ListRulesWithContext(ctx, listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.ListRulesWithContext(ctx, listRulesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_instance"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["zone_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["enforcement_mode"]).To(Equal([]string{"enabled"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "rules": [{"id": "ID", "crn": "CRN", "description": "Description", "contexts": [{"attributes": [{"name": "Name", "value": "Value"}]}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "operations": {"api_types": [{"api_type_id": "APITypeID"}]}, "enforcement_mode": "enabled", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListRules successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.ListRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(contextbasedrestrictionsv1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("testString")
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Region = core.StringPtr("testString")
				listRulesOptionsModel.Resource = core.StringPtr("testString")
				listRulesOptionsModel.ResourceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceInstance = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.ServiceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listRulesOptionsModel.ZoneID = core.StringPtr("testString")
				listRulesOptionsModel.Sort = core.StringPtr("testString")
				listRulesOptionsModel.EnforcementMode = core.StringPtr("enabled")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.ListRules(listRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRules with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(contextbasedrestrictionsv1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("testString")
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Region = core.StringPtr("testString")
				listRulesOptionsModel.Resource = core.StringPtr("testString")
				listRulesOptionsModel.ResourceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceInstance = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.ServiceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listRulesOptionsModel.ZoneID = core.StringPtr("testString")
				listRulesOptionsModel.Sort = core.StringPtr("testString")
				listRulesOptionsModel.EnforcementMode = core.StringPtr("enabled")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.ListRules(listRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListRulesOptions model with no property values
				listRulesOptionsModelNew := new(contextbasedrestrictionsv1.ListRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = contextBasedRestrictionsService.ListRules(listRulesOptionsModelNew)
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
			It(`Invoke ListRules successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListRulesOptions model
				listRulesOptionsModel := new(contextbasedrestrictionsv1.ListRulesOptions)
				listRulesOptionsModel.AccountID = core.StringPtr("testString")
				listRulesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listRulesOptionsModel.Region = core.StringPtr("testString")
				listRulesOptionsModel.Resource = core.StringPtr("testString")
				listRulesOptionsModel.ResourceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceInstance = core.StringPtr("testString")
				listRulesOptionsModel.ServiceName = core.StringPtr("testString")
				listRulesOptionsModel.ServiceType = core.StringPtr("testString")
				listRulesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listRulesOptionsModel.ZoneID = core.StringPtr("testString")
				listRulesOptionsModel.Sort = core.StringPtr("testString")
				listRulesOptionsModel.EnforcementMode = core.StringPtr("enabled")
				listRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.ListRules(listRulesOptionsModel)
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
	Describe(`GetRule(getRuleOptions *GetRuleOptions) - Operation response error`, func() {
		getRulePath := "/v1/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRule with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(contextbasedrestrictionsv1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
		getRulePath := "/v1/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "description": "Description", "contexts": [{"attributes": [{"name": "Name", "value": "Value"}]}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "operations": {"api_types": [{"api_type_id": "APITypeID"}]}, "enforcement_mode": "enabled", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetRule successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(contextbasedrestrictionsv1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.GetRuleWithContext(ctx, getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.GetRuleWithContext(ctx, getRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "description": "Description", "contexts": [{"attributes": [{"name": "Name", "value": "Value"}]}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "operations": {"api_types": [{"api_type_id": "APITypeID"}]}, "enforcement_mode": "enabled", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetRule successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.GetRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(contextbasedrestrictionsv1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.GetRule(getRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRule with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(contextbasedrestrictionsv1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.GetRule(getRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRuleOptions model with no property values
				getRuleOptionsModelNew := new(contextbasedrestrictionsv1.GetRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = contextBasedRestrictionsService.GetRule(getRuleOptionsModelNew)
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
			It(`Invoke GetRule successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetRuleOptions model
				getRuleOptionsModel := new(contextbasedrestrictionsv1.GetRuleOptions)
				getRuleOptionsModel.RuleID = core.StringPtr("testString")
				getRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				getRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.GetRule(getRuleOptionsModel)
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
	Describe(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions) - Operation response error`, func() {
		replaceRulePath := "/v1/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRulePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceRule with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("76921bd873115033bd2a0909fe081b45")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(contextbasedrestrictionsv1.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				replaceRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				replaceRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				replaceRuleOptionsModel.Operations = newRuleOperationsModel
				replaceRuleOptionsModel.EnforcementMode = core.StringPtr("disabled")
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions)`, func() {
		replaceRulePath := "/v1/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRulePath))
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
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "description": "Description", "contexts": [{"attributes": [{"name": "Name", "value": "Value"}]}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "operations": {"api_types": [{"api_type_id": "APITypeID"}]}, "enforcement_mode": "enabled", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplaceRule successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("76921bd873115033bd2a0909fe081b45")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(contextbasedrestrictionsv1.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				replaceRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				replaceRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				replaceRuleOptionsModel.Operations = newRuleOperationsModel
				replaceRuleOptionsModel.EnforcementMode = core.StringPtr("disabled")
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.ReplaceRuleWithContext(ctx, replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.ReplaceRuleWithContext(ctx, replaceRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceRulePath))
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
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "description": "Description", "contexts": [{"attributes": [{"name": "Name", "value": "Value"}]}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "operations": {"api_types": [{"api_type_id": "APITypeID"}]}, "enforcement_mode": "enabled", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplaceRule successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.ReplaceRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("76921bd873115033bd2a0909fe081b45")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(contextbasedrestrictionsv1.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				replaceRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				replaceRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				replaceRuleOptionsModel.Operations = newRuleOperationsModel
				replaceRuleOptionsModel.EnforcementMode = core.StringPtr("disabled")
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceRule with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("76921bd873115033bd2a0909fe081b45")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(contextbasedrestrictionsv1.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				replaceRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				replaceRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				replaceRuleOptionsModel.Operations = newRuleOperationsModel
				replaceRuleOptionsModel.EnforcementMode = core.StringPtr("disabled")
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceRuleOptions model with no property values
				replaceRuleOptionsModelNew := new(contextbasedrestrictionsv1.ReplaceRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = contextBasedRestrictionsService.ReplaceRule(replaceRuleOptionsModelNew)
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
			It(`Invoke ReplaceRule successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("76921bd873115033bd2a0909fe081b45")

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}

				// Construct an instance of the ReplaceRuleOptions model
				replaceRuleOptionsModel := new(contextbasedrestrictionsv1.ReplaceRuleOptions)
				replaceRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRuleOptionsModel.Description = core.StringPtr("this is an example of rule")
				replaceRuleOptionsModel.Contexts = []contextbasedrestrictionsv1.RuleContext{*ruleContextModel}
				replaceRuleOptionsModel.Resources = []contextbasedrestrictionsv1.Resource{*resourceModel}
				replaceRuleOptionsModel.Operations = newRuleOperationsModel
				replaceRuleOptionsModel.EnforcementMode = core.StringPtr("disabled")
				replaceRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				replaceRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptionsModel)
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
	Describe(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
		deleteRulePath := "/v1/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRule successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := contextBasedRestrictionsService.DeleteRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(contextbasedrestrictionsv1.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteRuleOptionsModel.TransactionID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = contextBasedRestrictionsService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRule with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the DeleteRuleOptions model
				deleteRuleOptionsModel := new(contextbasedrestrictionsv1.DeleteRuleOptions)
				deleteRuleOptionsModel.RuleID = core.StringPtr("testString")
				deleteRuleOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteRuleOptionsModel.TransactionID = core.StringPtr("testString")
				deleteRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := contextBasedRestrictionsService.DeleteRule(deleteRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRuleOptions model with no property values
				deleteRuleOptionsModelNew := new(contextbasedrestrictionsv1.DeleteRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = contextBasedRestrictionsService.DeleteRule(deleteRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) - Operation response error`, func() {
		getAccountSettingsPath := "/v1/account_settings/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountSettings with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(contextbasedrestrictionsv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptionsModel)
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
		getAccountSettingsPath := "/v1/account_settings/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "rule_count_limit": 14, "zone_count_limit": 14, "tags_rule_count_limit": 18, "current_rule_count": 16, "current_zone_count": 16, "current_tags_rule_count": 20, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(contextbasedrestrictionsv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
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

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "CRN", "rule_count_limit": 14, "zone_count_limit": 14, "tags_rule_count_limit": 18, "current_rule_count": 16, "current_zone_count": 16, "current_tags_rule_count": 20, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.GetAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(contextbasedrestrictionsv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountSettings with error: Operation validation and request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(contextbasedrestrictionsv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountSettingsOptions model with no property values
				getAccountSettingsOptionsModelNew := new(contextbasedrestrictionsv1.GetAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptionsModelNew)
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
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(contextbasedrestrictionsv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.XCorrelationID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptionsModel)
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
	Describe(`ListAvailableServiceOperations(listAvailableServiceOperationsOptions *ListAvailableServiceOperationsOptions) - Operation response error`, func() {
		listAvailableServiceOperationsPath := "/v1/operations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAvailableServiceOperationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_type"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAvailableServiceOperations with error: Operation response processing error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListAvailableServiceOperationsOptions model
				listAvailableServiceOperationsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServiceOperationsOptions)
				listAvailableServiceOperationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceName = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ResourceType = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				contextBasedRestrictionsService.EnableRetries(0, 0)
				result, response, operationErr = contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAvailableServiceOperations(listAvailableServiceOperationsOptions *ListAvailableServiceOperationsOptions)`, func() {
		listAvailableServiceOperationsPath := "/v1/operations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAvailableServiceOperationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_type"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"api_types": [{"api_type_id": "APITypeID", "display_name": "DisplayName", "description": "Description", "type": "Type", "actions": [{"action_id": "ActionID", "description": "Description"}], "enforcement_modes": ["EnforcementModes"]}]}`)
				}))
			})
			It(`Invoke ListAvailableServiceOperations successfully with retries`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())
				contextBasedRestrictionsService.EnableRetries(0, 0)

				// Construct an instance of the ListAvailableServiceOperationsOptions model
				listAvailableServiceOperationsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServiceOperationsOptions)
				listAvailableServiceOperationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceName = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ResourceType = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := contextBasedRestrictionsService.ListAvailableServiceOperationsWithContext(ctx, listAvailableServiceOperationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				contextBasedRestrictionsService.DisableRetries()
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = contextBasedRestrictionsService.ListAvailableServiceOperationsWithContext(ctx, listAvailableServiceOperationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAvailableServiceOperationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_type"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"api_types": [{"api_type_id": "APITypeID", "display_name": "DisplayName", "description": "Description", "type": "Type", "actions": [{"action_id": "ActionID", "description": "Description"}], "enforcement_modes": ["EnforcementModes"]}]}`)
				}))
			})
			It(`Invoke ListAvailableServiceOperations successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServiceOperations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAvailableServiceOperationsOptions model
				listAvailableServiceOperationsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServiceOperationsOptions)
				listAvailableServiceOperationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceName = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ResourceType = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAvailableServiceOperations with error: Operation request error`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListAvailableServiceOperationsOptions model
				listAvailableServiceOperationsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServiceOperationsOptions)
				listAvailableServiceOperationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceName = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ResourceType = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := contextBasedRestrictionsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptionsModel)
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
			It(`Invoke ListAvailableServiceOperations successfully`, func() {
				contextBasedRestrictionsService, serviceErr := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(contextBasedRestrictionsService).ToNot(BeNil())

				// Construct an instance of the ListAvailableServiceOperationsOptions model
				listAvailableServiceOperationsOptionsModel := new(contextbasedrestrictionsv1.ListAvailableServiceOperationsOptions)
				listAvailableServiceOperationsOptionsModel.XCorrelationID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.TransactionID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceName = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.ResourceType = core.StringPtr("testString")
				listAvailableServiceOperationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptionsModel)
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
			contextBasedRestrictionsService, _ := contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(&contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
				URL:           "http://contextbasedrestrictionsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateRuleOptions successfully`, func() {
				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				Expect(ruleContextAttributeModel).ToNot(BeNil())
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("65810ac762004f22ac19f8f8edf70a34")
				Expect(ruleContextAttributeModel.Name).To(Equal(core.StringPtr("networkZoneId")))
				Expect(ruleContextAttributeModel.Value).To(Equal(core.StringPtr("65810ac762004f22ac19f8f8edf70a34")))

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				Expect(ruleContextModel).ToNot(BeNil())
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}
				Expect(ruleContextModel.Attributes).To(Equal([]contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}))

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				Expect(resourceAttributeModel).ToNot(BeNil())
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")
				Expect(resourceAttributeModel.Name).To(Equal(core.StringPtr("accountId")))
				Expect(resourceAttributeModel.Value).To(Equal(core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")))
				Expect(resourceAttributeModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				Expect(resourceTagAttributeModel).ToNot(BeNil())
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")
				Expect(resourceTagAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(resourceTagAttributeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(resourceTagAttributeModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				Expect(resourceModel).ToNot(BeNil())
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}
				Expect(resourceModel.Attributes).To(Equal([]contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}))
				Expect(resourceModel.Tags).To(Equal([]contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}))

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				Expect(newRuleOperationsAPITypesItemModel).ToNot(BeNil())
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")
				Expect(newRuleOperationsAPITypesItemModel.APITypeID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				Expect(newRuleOperationsModel).ToNot(BeNil())
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}
				Expect(newRuleOperationsModel.APITypes).To(Equal([]contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}))

				// Construct an instance of the CreateRuleOptions model
				createRuleOptionsModel := contextBasedRestrictionsService.NewCreateRuleOptions()
				createRuleOptionsModel.SetDescription("this is an example of rule")
				createRuleOptionsModel.SetContexts([]contextbasedrestrictionsv1.RuleContext{*ruleContextModel})
				createRuleOptionsModel.SetResources([]contextbasedrestrictionsv1.Resource{*resourceModel})
				createRuleOptionsModel.SetOperations(newRuleOperationsModel)
				createRuleOptionsModel.SetEnforcementMode("enabled")
				createRuleOptionsModel.SetXCorrelationID("testString")
				createRuleOptionsModel.SetTransactionID("testString")
				createRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRuleOptionsModel).ToNot(BeNil())
				Expect(createRuleOptionsModel.Description).To(Equal(core.StringPtr("this is an example of rule")))
				Expect(createRuleOptionsModel.Contexts).To(Equal([]contextbasedrestrictionsv1.RuleContext{*ruleContextModel}))
				Expect(createRuleOptionsModel.Resources).To(Equal([]contextbasedrestrictionsv1.Resource{*resourceModel}))
				Expect(createRuleOptionsModel.Operations).To(Equal(newRuleOperationsModel))
				Expect(createRuleOptionsModel.EnforcementMode).To(Equal(core.StringPtr("enabled")))
				Expect(createRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateZoneOptions successfully`, func() {
				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				Expect(addressModel).ToNot(BeNil())
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")
				Expect(addressModel.Type).To(Equal(core.StringPtr("ipAddress")))
				Expect(addressModel.Value).To(Equal(core.StringPtr("169.23.56.234")))
				Expect(addressModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := contextBasedRestrictionsService.NewCreateZoneOptions()
				createZoneOptionsModel.SetName("an example of zone")
				createZoneOptionsModel.SetAccountID("12ab34cd56ef78ab90cd12ef34ab56cd")
				createZoneOptionsModel.SetDescription("this is an example of zone")
				createZoneOptionsModel.SetAddresses([]contextbasedrestrictionsv1.AddressIntf{addressModel})
				createZoneOptionsModel.SetExcluded([]contextbasedrestrictionsv1.AddressIntf{addressModel})
				createZoneOptionsModel.SetXCorrelationID("testString")
				createZoneOptionsModel.SetTransactionID("testString")
				createZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneOptionsModel).ToNot(BeNil())
				Expect(createZoneOptionsModel.Name).To(Equal(core.StringPtr("an example of zone")))
				Expect(createZoneOptionsModel.AccountID).To(Equal(core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")))
				Expect(createZoneOptionsModel.Description).To(Equal(core.StringPtr("this is an example of zone")))
				Expect(createZoneOptionsModel.Addresses).To(Equal([]contextbasedrestrictionsv1.AddressIntf{addressModel}))
				Expect(createZoneOptionsModel.Excluded).To(Equal([]contextbasedrestrictionsv1.AddressIntf{addressModel}))
				Expect(createZoneOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(createZoneOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRuleOptions successfully`, func() {
				// Construct an instance of the DeleteRuleOptions model
				ruleID := "testString"
				deleteRuleOptionsModel := contextBasedRestrictionsService.NewDeleteRuleOptions(ruleID)
				deleteRuleOptionsModel.SetRuleID("testString")
				deleteRuleOptionsModel.SetXCorrelationID("testString")
				deleteRuleOptionsModel.SetTransactionID("testString")
				deleteRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRuleOptionsModel).ToNot(BeNil())
				Expect(deleteRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneOptions successfully`, func() {
				// Construct an instance of the DeleteZoneOptions model
				zoneID := "testString"
				deleteZoneOptionsModel := contextBasedRestrictionsService.NewDeleteZoneOptions(zoneID)
				deleteZoneOptionsModel.SetZoneID("testString")
				deleteZoneOptionsModel.SetXCorrelationID("testString")
				deleteZoneOptionsModel.SetTransactionID("testString")
				deleteZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneOptionsModel).ToNot(BeNil())
				Expect(deleteZoneOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountSettingsOptions successfully`, func() {
				// Construct an instance of the GetAccountSettingsOptions model
				accountID := "testString"
				getAccountSettingsOptionsModel := contextBasedRestrictionsService.NewGetAccountSettingsOptions(accountID)
				getAccountSettingsOptionsModel.SetAccountID("testString")
				getAccountSettingsOptionsModel.SetXCorrelationID("testString")
				getAccountSettingsOptionsModel.SetTransactionID("testString")
				getAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(getAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRuleOptions successfully`, func() {
				// Construct an instance of the GetRuleOptions model
				ruleID := "testString"
				getRuleOptionsModel := contextBasedRestrictionsService.NewGetRuleOptions(ruleID)
				getRuleOptionsModel.SetRuleID("testString")
				getRuleOptionsModel.SetXCorrelationID("testString")
				getRuleOptionsModel.SetTransactionID("testString")
				getRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRuleOptionsModel).ToNot(BeNil())
				Expect(getRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServicerefTargetOptions successfully`, func() {
				// Construct an instance of the GetServicerefTargetOptions model
				serviceName := "testString"
				getServicerefTargetOptionsModel := contextBasedRestrictionsService.NewGetServicerefTargetOptions(serviceName)
				getServicerefTargetOptionsModel.SetServiceName("testString")
				getServicerefTargetOptionsModel.SetXCorrelationID("testString")
				getServicerefTargetOptionsModel.SetTransactionID("testString")
				getServicerefTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServicerefTargetOptionsModel).ToNot(BeNil())
				Expect(getServicerefTargetOptionsModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(getServicerefTargetOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getServicerefTargetOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getServicerefTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneOptions successfully`, func() {
				// Construct an instance of the GetZoneOptions model
				zoneID := "testString"
				getZoneOptionsModel := contextBasedRestrictionsService.NewGetZoneOptions(zoneID)
				getZoneOptionsModel.SetZoneID("testString")
				getZoneOptionsModel.SetXCorrelationID("testString")
				getZoneOptionsModel.SetTransactionID("testString")
				getZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneOptionsModel).ToNot(BeNil())
				Expect(getZoneOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(getZoneOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(getZoneOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAvailableServiceOperationsOptions successfully`, func() {
				// Construct an instance of the ListAvailableServiceOperationsOptions model
				listAvailableServiceOperationsOptionsModel := contextBasedRestrictionsService.NewListAvailableServiceOperationsOptions()
				listAvailableServiceOperationsOptionsModel.SetXCorrelationID("testString")
				listAvailableServiceOperationsOptionsModel.SetTransactionID("testString")
				listAvailableServiceOperationsOptionsModel.SetServiceName("testString")
				listAvailableServiceOperationsOptionsModel.SetServiceGroupID("testString")
				listAvailableServiceOperationsOptionsModel.SetResourceType("testString")
				listAvailableServiceOperationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAvailableServiceOperationsOptionsModel).ToNot(BeNil())
				Expect(listAvailableServiceOperationsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listAvailableServiceOperationsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAvailableServiceOperationsOptionsModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(listAvailableServiceOperationsOptionsModel.ServiceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAvailableServiceOperationsOptionsModel.ResourceType).To(Equal(core.StringPtr("testString")))
				Expect(listAvailableServiceOperationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAvailableServicerefTargetsOptions successfully`, func() {
				// Construct an instance of the ListAvailableServicerefTargetsOptions model
				listAvailableServicerefTargetsOptionsModel := contextBasedRestrictionsService.NewListAvailableServicerefTargetsOptions()
				listAvailableServicerefTargetsOptionsModel.SetXCorrelationID("testString")
				listAvailableServicerefTargetsOptionsModel.SetTransactionID("testString")
				listAvailableServicerefTargetsOptionsModel.SetType("all")
				listAvailableServicerefTargetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAvailableServicerefTargetsOptionsModel).ToNot(BeNil())
				Expect(listAvailableServicerefTargetsOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listAvailableServicerefTargetsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAvailableServicerefTargetsOptionsModel.Type).To(Equal(core.StringPtr("all")))
				Expect(listAvailableServicerefTargetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRulesOptions successfully`, func() {
				// Construct an instance of the ListRulesOptions model
				accountID := "testString"
				listRulesOptionsModel := contextBasedRestrictionsService.NewListRulesOptions(accountID)
				listRulesOptionsModel.SetAccountID("testString")
				listRulesOptionsModel.SetXCorrelationID("testString")
				listRulesOptionsModel.SetTransactionID("testString")
				listRulesOptionsModel.SetRegion("testString")
				listRulesOptionsModel.SetResource("testString")
				listRulesOptionsModel.SetResourceType("testString")
				listRulesOptionsModel.SetServiceInstance("testString")
				listRulesOptionsModel.SetServiceName("testString")
				listRulesOptionsModel.SetServiceType("testString")
				listRulesOptionsModel.SetServiceGroupID("testString")
				listRulesOptionsModel.SetZoneID("testString")
				listRulesOptionsModel.SetSort("testString")
				listRulesOptionsModel.SetEnforcementMode("enabled")
				listRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRulesOptionsModel).ToNot(BeNil())
				Expect(listRulesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.Resource).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.ResourceType).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.ServiceInstance).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.ServiceType).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.ServiceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listRulesOptionsModel.EnforcementMode).To(Equal(core.StringPtr("enabled")))
				Expect(listRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListZonesOptions successfully`, func() {
				// Construct an instance of the ListZonesOptions model
				accountID := "testString"
				listZonesOptionsModel := contextBasedRestrictionsService.NewListZonesOptions(accountID)
				listZonesOptionsModel.SetAccountID("testString")
				listZonesOptionsModel.SetXCorrelationID("testString")
				listZonesOptionsModel.SetTransactionID("testString")
				listZonesOptionsModel.SetName("testString")
				listZonesOptionsModel.SetSort("testString")
				listZonesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listZonesOptionsModel).ToNot(BeNil())
				Expect(listZonesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listZonesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listZonesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listZonesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listZonesOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listZonesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewNewRuleOperations successfully`, func() {
				apiTypes := []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{}
				_model, err := contextBasedRestrictionsService.NewNewRuleOperations(apiTypes)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNewRuleOperationsAPITypesItem successfully`, func() {
				apiTypeID := "testString"
				_model, err := contextBasedRestrictionsService.NewNewRuleOperationsAPITypesItem(apiTypeID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplaceRuleOptions successfully`, func() {
				// Construct an instance of the RuleContextAttribute model
				ruleContextAttributeModel := new(contextbasedrestrictionsv1.RuleContextAttribute)
				Expect(ruleContextAttributeModel).ToNot(BeNil())
				ruleContextAttributeModel.Name = core.StringPtr("networkZoneId")
				ruleContextAttributeModel.Value = core.StringPtr("76921bd873115033bd2a0909fe081b45")
				Expect(ruleContextAttributeModel.Name).To(Equal(core.StringPtr("networkZoneId")))
				Expect(ruleContextAttributeModel.Value).To(Equal(core.StringPtr("76921bd873115033bd2a0909fe081b45")))

				// Construct an instance of the RuleContext model
				ruleContextModel := new(contextbasedrestrictionsv1.RuleContext)
				Expect(ruleContextModel).ToNot(BeNil())
				ruleContextModel.Attributes = []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}
				Expect(ruleContextModel.Attributes).To(Equal([]contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel}))

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(contextbasedrestrictionsv1.ResourceAttribute)
				Expect(resourceAttributeModel).ToNot(BeNil())
				resourceAttributeModel.Name = core.StringPtr("accountId")
				resourceAttributeModel.Value = core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")
				resourceAttributeModel.Operator = core.StringPtr("testString")
				Expect(resourceAttributeModel.Name).To(Equal(core.StringPtr("accountId")))
				Expect(resourceAttributeModel.Value).To(Equal(core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")))
				Expect(resourceAttributeModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ResourceTagAttribute model
				resourceTagAttributeModel := new(contextbasedrestrictionsv1.ResourceTagAttribute)
				Expect(resourceTagAttributeModel).ToNot(BeNil())
				resourceTagAttributeModel.Name = core.StringPtr("testString")
				resourceTagAttributeModel.Value = core.StringPtr("testString")
				resourceTagAttributeModel.Operator = core.StringPtr("testString")
				Expect(resourceTagAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(resourceTagAttributeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(resourceTagAttributeModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Resource model
				resourceModel := new(contextbasedrestrictionsv1.Resource)
				Expect(resourceModel).ToNot(BeNil())
				resourceModel.Attributes = []contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}
				resourceModel.Tags = []contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}
				Expect(resourceModel.Attributes).To(Equal([]contextbasedrestrictionsv1.ResourceAttribute{*resourceAttributeModel}))
				Expect(resourceModel.Tags).To(Equal([]contextbasedrestrictionsv1.ResourceTagAttribute{*resourceTagAttributeModel}))

				// Construct an instance of the NewRuleOperationsAPITypesItem model
				newRuleOperationsAPITypesItemModel := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
				Expect(newRuleOperationsAPITypesItemModel).ToNot(BeNil())
				newRuleOperationsAPITypesItemModel.APITypeID = core.StringPtr("testString")
				Expect(newRuleOperationsAPITypesItemModel.APITypeID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the NewRuleOperations model
				newRuleOperationsModel := new(contextbasedrestrictionsv1.NewRuleOperations)
				Expect(newRuleOperationsModel).ToNot(BeNil())
				newRuleOperationsModel.APITypes = []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}
				Expect(newRuleOperationsModel.APITypes).To(Equal([]contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{*newRuleOperationsAPITypesItemModel}))

				// Construct an instance of the ReplaceRuleOptions model
				ruleID := "testString"
				ifMatch := "testString"
				replaceRuleOptionsModel := contextBasedRestrictionsService.NewReplaceRuleOptions(ruleID, ifMatch)
				replaceRuleOptionsModel.SetRuleID("testString")
				replaceRuleOptionsModel.SetIfMatch("testString")
				replaceRuleOptionsModel.SetDescription("this is an example of rule")
				replaceRuleOptionsModel.SetContexts([]contextbasedrestrictionsv1.RuleContext{*ruleContextModel})
				replaceRuleOptionsModel.SetResources([]contextbasedrestrictionsv1.Resource{*resourceModel})
				replaceRuleOptionsModel.SetOperations(newRuleOperationsModel)
				replaceRuleOptionsModel.SetEnforcementMode("disabled")
				replaceRuleOptionsModel.SetXCorrelationID("testString")
				replaceRuleOptionsModel.SetTransactionID("testString")
				replaceRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceRuleOptionsModel).ToNot(BeNil())
				Expect(replaceRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.Description).To(Equal(core.StringPtr("this is an example of rule")))
				Expect(replaceRuleOptionsModel.Contexts).To(Equal([]contextbasedrestrictionsv1.RuleContext{*ruleContextModel}))
				Expect(replaceRuleOptionsModel.Resources).To(Equal([]contextbasedrestrictionsv1.Resource{*resourceModel}))
				Expect(replaceRuleOptionsModel.Operations).To(Equal(newRuleOperationsModel))
				Expect(replaceRuleOptionsModel.EnforcementMode).To(Equal(core.StringPtr("disabled")))
				Expect(replaceRuleOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceZoneOptions successfully`, func() {
				// Construct an instance of the AddressIPAddress model
				addressModel := new(contextbasedrestrictionsv1.AddressIPAddress)
				Expect(addressModel).ToNot(BeNil())
				addressModel.Type = core.StringPtr("ipAddress")
				addressModel.Value = core.StringPtr("169.23.56.234")
				addressModel.ID = core.StringPtr("testString")
				Expect(addressModel.Type).To(Equal(core.StringPtr("ipAddress")))
				Expect(addressModel.Value).To(Equal(core.StringPtr("169.23.56.234")))
				Expect(addressModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceZoneOptions model
				zoneID := "testString"
				ifMatch := "testString"
				replaceZoneOptionsModel := contextBasedRestrictionsService.NewReplaceZoneOptions(zoneID, ifMatch)
				replaceZoneOptionsModel.SetZoneID("testString")
				replaceZoneOptionsModel.SetIfMatch("testString")
				replaceZoneOptionsModel.SetName("an example of zone")
				replaceZoneOptionsModel.SetAccountID("12ab34cd56ef78ab90cd12ef34ab56cd")
				replaceZoneOptionsModel.SetDescription("this is an example of zone")
				replaceZoneOptionsModel.SetAddresses([]contextbasedrestrictionsv1.AddressIntf{addressModel})
				replaceZoneOptionsModel.SetExcluded([]contextbasedrestrictionsv1.AddressIntf{addressModel})
				replaceZoneOptionsModel.SetXCorrelationID("testString")
				replaceZoneOptionsModel.SetTransactionID("testString")
				replaceZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceZoneOptionsModel).ToNot(BeNil())
				Expect(replaceZoneOptionsModel.ZoneID).To(Equal(core.StringPtr("testString")))
				Expect(replaceZoneOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceZoneOptionsModel.Name).To(Equal(core.StringPtr("an example of zone")))
				Expect(replaceZoneOptionsModel.AccountID).To(Equal(core.StringPtr("12ab34cd56ef78ab90cd12ef34ab56cd")))
				Expect(replaceZoneOptionsModel.Description).To(Equal(core.StringPtr("this is an example of zone")))
				Expect(replaceZoneOptionsModel.Addresses).To(Equal([]contextbasedrestrictionsv1.AddressIntf{addressModel}))
				Expect(replaceZoneOptionsModel.Excluded).To(Equal([]contextbasedrestrictionsv1.AddressIntf{addressModel}))
				Expect(replaceZoneOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(replaceZoneOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResource successfully`, func() {
				attributes := []contextbasedrestrictionsv1.ResourceAttribute{}
				_model, err := contextBasedRestrictionsService.NewResource(attributes)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceAttribute successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := contextBasedRestrictionsService.NewResourceAttribute(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceTagAttribute successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := contextBasedRestrictionsService.NewResourceTagAttribute(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleContext successfully`, func() {
				attributes := []contextbasedrestrictionsv1.RuleContextAttribute{}
				_model, err := contextBasedRestrictionsService.NewRuleContext(attributes)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleContextAttribute successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := contextBasedRestrictionsService.NewRuleContextAttribute(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewServiceRefValue successfully`, func() {
				accountID := "testString"
				_model, err := contextBasedRestrictionsService.NewServiceRefValue(accountID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAddressIPAddress successfully`, func() {
				typeVar := "ipAddress"
				value := "testString"
				_model, err := contextBasedRestrictionsService.NewAddressIPAddress(typeVar, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAddressIPAddressRange successfully`, func() {
				typeVar := "ipRange"
				value := "testString"
				_model, err := contextBasedRestrictionsService.NewAddressIPAddressRange(typeVar, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAddressServiceRef successfully`, func() {
				typeVar := "serviceRef"
				var ref *contextbasedrestrictionsv1.ServiceRefValue = nil
				_, err := contextBasedRestrictionsService.NewAddressServiceRef(typeVar, ref)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewAddressSubnet successfully`, func() {
				typeVar := "subnet"
				value := "testString"
				_model, err := contextBasedRestrictionsService.NewAddressSubnet(typeVar, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAddressVPC successfully`, func() {
				typeVar := "vpc"
				value := "testString"
				_model, err := contextBasedRestrictionsService.NewAddressVPC(typeVar, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAddress successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.Address)
			model.Type = core.StringPtr("ipAddress")
			model.Value = core.StringPtr("testString")
			model.ID = core.StringPtr("testString")
			model.Ref = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result interface{}
			err = contextbasedrestrictionsv1.UnmarshalAddress(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
		It(`Invoke UnmarshalNewRuleOperations successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.NewRuleOperations)
			model.APITypes = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.NewRuleOperations
			err = contextbasedrestrictionsv1.UnmarshalNewRuleOperations(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalNewRuleOperationsAPITypesItem successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem)
			model.APITypeID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem
			err = contextbasedrestrictionsv1.UnmarshalNewRuleOperationsAPITypesItem(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalResource successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.Resource)
			model.Attributes = nil
			model.Tags = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.Resource
			err = contextbasedrestrictionsv1.UnmarshalResource(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalResourceAttribute successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.ResourceAttribute)
			model.Name = core.StringPtr("testString")
			model.Value = core.StringPtr("testString")
			model.Operator = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.ResourceAttribute
			err = contextbasedrestrictionsv1.UnmarshalResourceAttribute(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalResourceTagAttribute successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.ResourceTagAttribute)
			model.Name = core.StringPtr("testString")
			model.Value = core.StringPtr("testString")
			model.Operator = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.ResourceTagAttribute
			err = contextbasedrestrictionsv1.UnmarshalResourceTagAttribute(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRuleContext successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.RuleContext)
			model.Attributes = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.RuleContext
			err = contextbasedrestrictionsv1.UnmarshalRuleContext(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRuleContextAttribute successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.RuleContextAttribute)
			model.Name = core.StringPtr("testString")
			model.Value = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.RuleContextAttribute
			err = contextbasedrestrictionsv1.UnmarshalRuleContextAttribute(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalServiceRefValue successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.ServiceRefValue)
			model.AccountID = core.StringPtr("testString")
			model.ServiceType = core.StringPtr("testString")
			model.ServiceName = core.StringPtr("testString")
			model.ServiceInstance = core.StringPtr("testString")
			model.Location = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.ServiceRefValue
			err = contextbasedrestrictionsv1.UnmarshalServiceRefValue(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAddressIPAddress successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.AddressIPAddress)
			model.Type = core.StringPtr("ipAddress")
			model.Value = core.StringPtr("testString")
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.AddressIPAddress
			err = contextbasedrestrictionsv1.UnmarshalAddressIPAddress(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAddressIPAddressRange successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.AddressIPAddressRange)
			model.Type = core.StringPtr("ipRange")
			model.Value = core.StringPtr("testString")
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.AddressIPAddressRange
			err = contextbasedrestrictionsv1.UnmarshalAddressIPAddressRange(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAddressServiceRef successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.AddressServiceRef)
			model.Type = core.StringPtr("serviceRef")
			model.Ref = nil
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.AddressServiceRef
			err = contextbasedrestrictionsv1.UnmarshalAddressServiceRef(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAddressSubnet successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.AddressSubnet)
			model.Type = core.StringPtr("subnet")
			model.Value = core.StringPtr("testString")
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.AddressSubnet
			err = contextbasedrestrictionsv1.UnmarshalAddressSubnet(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAddressVPC successfully`, func() {
			// Construct an instance of the model.
			model := new(contextbasedrestrictionsv1.AddressVPC)
			model.Type = core.StringPtr("vpc")
			model.Value = core.StringPtr("testString")
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *contextbasedrestrictionsv1.AddressVPC
			err = contextbasedrestrictionsv1.UnmarshalAddressVPC(raw, &result)
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
