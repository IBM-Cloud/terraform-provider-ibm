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

package zonesv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/zonesv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`ZonesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
				URL: "https://zonesv1/api",
				Crn: core.StringPtr(crn),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONES_URL":       "https://zonesv1/api",
				"ZONES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
					Crn: core.StringPtr(crn),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
					Crn: core.StringPtr(crn),
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
				"ZONES_URL":       "https://zonesv1/api",
				"ZONES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
				Crn: core.StringPtr(crn),
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
				"ZONES_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := zonesv1.NewZonesV1UsingExternalConfig(&zonesv1.ZonesV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListZones(listZonesOptions *ListZonesOptions) - Operation response error`, func() {
		crn := "testString"
		listZonesPath := "/v1/testString/zones"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listZonesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListZones with error: Operation response processing error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(zonesv1.ListZonesOptions)
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListZones(listZonesOptionsModel)
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
		crn := "testString"
		listZonesPath := "/v1/testString/zones"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listZonesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"]}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListZones successfully`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListZones(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(zonesv1.ListZonesOptions)
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListZones(listZonesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListZones with error: Operation request error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := new(zonesv1.ListZonesOptions)
				listZonesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListZones(listZonesOptionsModel)
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
	Describe(`CreateZone(createZoneOptions *CreateZoneOptions) - Operation response error`, func() {
		crn := "testString"
		createZonePath := "/v1/testString/zones"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createZonePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZone with error: Operation response processing error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(zonesv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("test-example.com")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateZone(createZoneOptionsModel)
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
		crn := "testString"
		createZonePath := "/v1/testString/zones"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createZonePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"]}}`)
				}))
			})
			It(`Invoke CreateZone successfully`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(zonesv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("test-example.com")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateZone(createZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateZone with error: Operation request error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := new(zonesv1.CreateZoneOptions)
				createZoneOptionsModel.Name = core.StringPtr("test-example.com")
				createZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateZone(createZoneOptionsModel)
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
	Describe(`DeleteZone(deleteZoneOptions *DeleteZoneOptions) - Operation response error`, func() {
		crn := "testString"
		deleteZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteZonePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZone with error: Operation response processing error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(zonesv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteZone(deleteZoneOptions *DeleteZoneOptions)`, func() {
		crn := "testString"
		deleteZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteZonePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteZone successfully`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(zonesv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteZone with error: Operation validation and request error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneOptions model
				deleteZoneOptionsModel := new(zonesv1.DeleteZoneOptions)
				deleteZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				deleteZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteZone(deleteZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneOptions model with no property values
				deleteZoneOptionsModelNew := new(zonesv1.DeleteZoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteZone(deleteZoneOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZone(getZoneOptions *GetZoneOptions) - Operation response error`, func() {
		crn := "testString"
		getZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getZonePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZone with error: Operation response processing error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(zonesv1.GetZoneOptions)
				getZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetZone(getZoneOptionsModel)
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
		crn := "testString"
		getZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getZonePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"]}}`)
				}))
			})
			It(`Invoke GetZone successfully`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(zonesv1.GetZoneOptions)
				getZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetZone(getZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetZone with error: Operation validation and request error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetZoneOptions model
				getZoneOptionsModel := new(zonesv1.GetZoneOptions)
				getZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				getZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetZone(getZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneOptions model with no property values
				getZoneOptionsModelNew := new(zonesv1.GetZoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetZone(getZoneOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZone(updateZoneOptions *UpdateZoneOptions) - Operation response error`, func() {
		crn := "testString"
		updateZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateZonePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZone with error: Operation response processing error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneOptions model
				updateZoneOptionsModel := new(zonesv1.UpdateZoneOptions)
				updateZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateZoneOptionsModel.Paused = core.BoolPtr(false)
				updateZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateZone(updateZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateZone(updateZoneOptions *UpdateZoneOptions)`, func() {
		crn := "testString"
		updateZonePath := "/v1/testString/zones/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateZonePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "test-example.com", "original_registrar": "GoDaddy", "original_dnshost": "NameCheap", "status": "active", "paused": false, "original_name_servers": ["ns1.originaldnshost.com"], "name_servers": ["ns001.name.cloud.ibm.com"]}}`)
				}))
			})
			It(`Invoke UpdateZone successfully`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateZone(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateZoneOptions model
				updateZoneOptionsModel := new(zonesv1.UpdateZoneOptions)
				updateZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateZoneOptionsModel.Paused = core.BoolPtr(false)
				updateZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateZone(updateZoneOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateZone with error: Operation validation and request error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneOptions model
				updateZoneOptionsModel := new(zonesv1.UpdateZoneOptions)
				updateZoneOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				updateZoneOptionsModel.Paused = core.BoolPtr(false)
				updateZoneOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateZone(updateZoneOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateZoneOptions model with no property values
				updateZoneOptionsModelNew := new(zonesv1.UpdateZoneOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateZone(updateZoneOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ZoneActivationCheck(zoneActivationCheckOptions *ZoneActivationCheckOptions) - Operation response error`, func() {
		crn := "testString"
		zoneActivationCheckPath := "/v1/testString/zones/testString/activation_check"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(zoneActivationCheckPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ZoneActivationCheck with error: Operation response processing error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ZoneActivationCheckOptions model
				zoneActivationCheckOptionsModel := new(zonesv1.ZoneActivationCheckOptions)
				zoneActivationCheckOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				zoneActivationCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ZoneActivationCheck(zoneActivationCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ZoneActivationCheck(zoneActivationCheckOptions *ZoneActivationCheckOptions)`, func() {
		crn := "testString"
		zoneActivationCheckPath := "/v1/testString/zones/testString/activation_check"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(zoneActivationCheckPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke ZoneActivationCheck successfully`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ZoneActivationCheck(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ZoneActivationCheckOptions model
				zoneActivationCheckOptionsModel := new(zonesv1.ZoneActivationCheckOptions)
				zoneActivationCheckOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				zoneActivationCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ZoneActivationCheck(zoneActivationCheckOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ZoneActivationCheck with error: Operation validation and request error`, func() {
				testService, testServiceErr := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ZoneActivationCheckOptions model
				zoneActivationCheckOptionsModel := new(zonesv1.ZoneActivationCheckOptions)
				zoneActivationCheckOptionsModel.ZoneIdentifier = core.StringPtr("testString")
				zoneActivationCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ZoneActivationCheck(zoneActivationCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ZoneActivationCheckOptions model with no property values
				zoneActivationCheckOptionsModelNew := new(zonesv1.ZoneActivationCheckOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ZoneActivationCheck(zoneActivationCheckOptionsModelNew)
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
			testService, _ := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
				URL:           "http://zonesv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			It(`Invoke NewCreateZoneOptions successfully`, func() {
				// Construct an instance of the CreateZoneOptions model
				createZoneOptionsModel := testService.NewCreateZoneOptions()
				createZoneOptionsModel.SetName("test-example.com")
				createZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneOptionsModel).ToNot(BeNil())
				Expect(createZoneOptionsModel.Name).To(Equal(core.StringPtr("test-example.com")))
				Expect(createZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneOptions successfully`, func() {
				// Construct an instance of the DeleteZoneOptions model
				zoneIdentifier := "testString"
				deleteZoneOptionsModel := testService.NewDeleteZoneOptions(zoneIdentifier)
				deleteZoneOptionsModel.SetZoneIdentifier("testString")
				deleteZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneOptionsModel).ToNot(BeNil())
				Expect(deleteZoneOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneOptions successfully`, func() {
				// Construct an instance of the GetZoneOptions model
				zoneIdentifier := "testString"
				getZoneOptionsModel := testService.NewGetZoneOptions(zoneIdentifier)
				getZoneOptionsModel.SetZoneIdentifier("testString")
				getZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneOptionsModel).ToNot(BeNil())
				Expect(getZoneOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListZonesOptions successfully`, func() {
				// Construct an instance of the ListZonesOptions model
				listZonesOptionsModel := testService.NewListZonesOptions()
				listZonesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listZonesOptionsModel).ToNot(BeNil())
				Expect(listZonesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneOptions successfully`, func() {
				// Construct an instance of the UpdateZoneOptions model
				zoneIdentifier := "testString"
				updateZoneOptionsModel := testService.NewUpdateZoneOptions(zoneIdentifier)
				updateZoneOptionsModel.SetZoneIdentifier("testString")
				updateZoneOptionsModel.SetPaused(false)
				updateZoneOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneOptionsModel).ToNot(BeNil())
				Expect(updateZoneOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneOptionsModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(updateZoneOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewZoneActivationCheckOptions successfully`, func() {
				// Construct an instance of the ZoneActivationCheckOptions model
				zoneIdentifier := "testString"
				zoneActivationCheckOptionsModel := testService.NewZoneActivationCheckOptions(zoneIdentifier)
				zoneActivationCheckOptionsModel.SetZoneIdentifier("testString")
				zoneActivationCheckOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(zoneActivationCheckOptionsModel).ToNot(BeNil())
				Expect(zoneActivationCheckOptionsModel.ZoneIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(zoneActivationCheckOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
