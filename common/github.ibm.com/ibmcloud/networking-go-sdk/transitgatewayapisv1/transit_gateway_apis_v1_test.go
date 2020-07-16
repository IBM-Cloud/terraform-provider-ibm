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

package transitgatewayapisv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/transitgatewayapisv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`TransitGatewayApisV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := CreateMockDate()
		It(`Instantiate service client`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       version,
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:     "{BAD_URL_STRING",
				Version: version,
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:     "https://transitgatewayapisv1/api",
				Version: version,
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := CreateMockDate()
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: version,
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:     "https://testService/api",
					Version: version,
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: version,
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
				"TRANSIT_GATEWAY_APIS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Version: version,
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
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:     "{BAD_URL_STRING",
				Version: version,
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListTransitGateways(listTransitGatewaysOptions *ListTransitGatewaysOptions) - Operation response error`, func() {
		version := CreateMockDate()
		listTransitGatewaysPath := "/transit_gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listTransitGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTransitGateways with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := new(transitgatewayapisv1.ListTransitGatewaysOptions)
				listTransitGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListTransitGateways(listTransitGatewaysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListTransitGateways(listTransitGatewaysOptions *ListTransitGatewaysOptions)`, func() {
		version := CreateMockDate()
		listTransitGatewaysPath := "/transit_gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listTransitGatewaysPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"transit_gateways": [{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "location": "us-south", "created_at": "2019-01-01T12:00:00", "global": true, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListTransitGateways successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListTransitGateways(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := new(transitgatewayapisv1.ListTransitGatewaysOptions)
				listTransitGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListTransitGateways(listTransitGatewaysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListTransitGateways with error: Operation request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := new(transitgatewayapisv1.ListTransitGatewaysOptions)
				listTransitGatewaysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListTransitGateways(listTransitGatewaysOptionsModel)
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
	Describe(`CreateTransitGateway(createTransitGatewayOptions *CreateTransitGatewayOptions) - Operation response error`, func() {
		version := CreateMockDate()
		createTransitGatewayPath := "/transit_gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTransitGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				createTransitGatewayOptionsModel.Location = core.StringPtr("us-south")
				createTransitGatewayOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateTransitGateway(createTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateTransitGateway(createTransitGatewayOptions *CreateTransitGatewayOptions)`, func() {
		version := CreateMockDate()
		createTransitGatewayPath := "/transit_gateways"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "location": "us-south", "created_at": "2019-01-01T12:00:00", "global": true, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateTransitGateway successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateTransitGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				createTransitGatewayOptionsModel.Location = core.StringPtr("us-south")
				createTransitGatewayOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateTransitGateway(createTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateTransitGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				createTransitGatewayOptionsModel.Location = core.StringPtr("us-south")
				createTransitGatewayOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				createTransitGatewayOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateTransitGateway(createTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTransitGatewayOptions model with no property values
				createTransitGatewayOptionsModelNew := new(transitgatewayapisv1.CreateTransitGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateTransitGateway(createTransitGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteTransitGateway(deleteTransitGatewayOptions *DeleteTransitGatewayOptions)`, func() {
		version := CreateMockDate()
		deleteTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteTransitGatewayPath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for version query parameter

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTransitGateway successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteTransitGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTransitGatewayOptions model
				deleteTransitGatewayOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayOptions)
				deleteTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteTransitGateway(deleteTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTransitGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteTransitGatewayOptions model
				deleteTransitGatewayOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayOptions)
				deleteTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteTransitGateway(deleteTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTransitGatewayOptions model with no property values
				deleteTransitGatewayOptionsModelNew := new(transitgatewayapisv1.DeleteTransitGatewayOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteTransitGateway(deleteTransitGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGateway(getTransitGatewayOptions *GetTransitGatewayOptions) - Operation response error`, func() {
		version := CreateMockDate()
		getTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTransitGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTransitGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayOptions model
				getTransitGatewayOptionsModel := new(transitgatewayapisv1.GetTransitGatewayOptions)
				getTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetTransitGateway(getTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTransitGateway(getTransitGatewayOptions *GetTransitGatewayOptions)`, func() {
		version := CreateMockDate()
		getTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTransitGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "location": "us-south", "created_at": "2019-01-01T12:00:00", "global": true, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetTransitGateway successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetTransitGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTransitGatewayOptions model
				getTransitGatewayOptionsModel := new(transitgatewayapisv1.GetTransitGatewayOptions)
				getTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetTransitGateway(getTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetTransitGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayOptions model
				getTransitGatewayOptionsModel := new(transitgatewayapisv1.GetTransitGatewayOptions)
				getTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetTransitGateway(getTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTransitGatewayOptions model with no property values
				getTransitGatewayOptionsModelNew := new(transitgatewayapisv1.GetTransitGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetTransitGateway(getTransitGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTransitGateway(updateTransitGatewayOptions *UpdateTransitGatewayOptions) - Operation response error`, func() {
		version := CreateMockDate()
		updateTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateTransitGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTransitGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayOptions model
				updateTransitGatewayOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				updateTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway")
				updateTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateTransitGateway(updateTransitGatewayOptions *UpdateTransitGatewayOptions)`, func() {
		version := CreateMockDate()
		updateTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateTransitGatewayPath))
					Expect(req.Method).To(Equal("PATCH"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "location": "us-south", "created_at": "2019-01-01T12:00:00", "global": true, "resource_group": {"id": "56969d6043e9465c883cb9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d6043e9465c883cb9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateTransitGateway successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateTransitGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayOptions model
				updateTransitGatewayOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				updateTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway")
				updateTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateTransitGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayOptions model
				updateTransitGatewayOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				updateTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayOptionsModel.Global = core.BoolPtr(true)
				updateTransitGatewayOptionsModel.Name = core.StringPtr("my-transit-gateway")
				updateTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateTransitGateway(updateTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTransitGatewayOptions model with no property values
				updateTransitGatewayOptionsModelNew := new(transitgatewayapisv1.UpdateTransitGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateTransitGateway(updateTransitGatewayOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := CreateMockDate()
		It(`Instantiate service client`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       version,
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:     "{BAD_URL_STRING",
				Version: version,
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:     "https://transitgatewayapisv1/api",
				Version: version,
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := CreateMockDate()
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: version,
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:     "https://testService/api",
					Version: version,
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: version,
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
				"TRANSIT_GATEWAY_APIS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Version: version,
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
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:     "{BAD_URL_STRING",
				Version: version,
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListTransitGatewayConnections(listTransitGatewayConnectionsOptions *ListTransitGatewayConnectionsOptions) - Operation response error`, func() {
		version := CreateMockDate()
		listTransitGatewayConnectionsPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listTransitGatewayConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTransitGatewayConnections with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				listTransitGatewayConnectionsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				listTransitGatewayConnectionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListTransitGatewayConnections(listTransitGatewayConnectionsOptions *ListTransitGatewayConnectionsOptions)`, func() {
		version := CreateMockDate()
		listTransitGatewayConnectionsPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listTransitGatewayConnectionsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"connections": [{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "network_account_id": "28e4d90ac7504be694471ee66e70d0d5", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "created_at": "2019-01-01T12:00:00", "request_status": "pending", "status": "attached", "updated_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListTransitGatewayConnections successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListTransitGatewayConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				listTransitGatewayConnectionsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				listTransitGatewayConnectionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListTransitGatewayConnections with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				listTransitGatewayConnectionsOptionsModel := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				listTransitGatewayConnectionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				listTransitGatewayConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTransitGatewayConnectionsOptions model with no property values
				listTransitGatewayConnectionsOptionsModelNew := new(transitgatewayapisv1.ListTransitGatewayConnectionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListTransitGatewayConnections(listTransitGatewayConnectionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTransitGatewayConnection(createTransitGatewayConnectionOptions *CreateTransitGatewayConnectionOptions) - Operation response error`, func() {
		version := CreateMockDate()
		createTransitGatewayConnectionPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTransitGatewayConnection with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("28e4d90ac7504be694471ee66e70d0d5")
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateTransitGatewayConnection(createTransitGatewayConnectionOptions *CreateTransitGatewayConnectionOptions)`, func() {
		version := CreateMockDate()
		createTransitGatewayConnectionPath := "/transit_gateways/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "network_account_id": "28e4d90ac7504be694471ee66e70d0d5", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "created_at": "2019-01-01T12:00:00", "request_status": "pending", "status": "attached", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateTransitGatewayConnection successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateTransitGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("28e4d90ac7504be694471ee66e70d0d5")
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateTransitGatewayConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				createTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				createTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionOptionsModel.NetworkType = core.StringPtr("vpc")
				createTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.NetworkID = core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.NetworkAccountID = core.StringPtr("28e4d90ac7504be694471ee66e70d0d5")
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTransitGatewayConnectionOptions model with no property values
				createTransitGatewayConnectionOptionsModelNew := new(transitgatewayapisv1.CreateTransitGatewayConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptions *DeleteTransitGatewayConnectionOptions)`, func() {
		version := CreateMockDate()
		deleteTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for version query parameter

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTransitGatewayConnection successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteTransitGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTransitGatewayConnectionOptions model
				deleteTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayConnectionOptions)
				deleteTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTransitGatewayConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteTransitGatewayConnectionOptions model
				deleteTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.DeleteTransitGatewayConnectionOptions)
				deleteTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				deleteTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				deleteTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTransitGatewayConnectionOptions model with no property values
				deleteTransitGatewayConnectionOptionsModelNew := new(transitgatewayapisv1.DeleteTransitGatewayConnectionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteTransitGatewayConnection(deleteTransitGatewayConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTransitGatewayConnection(getTransitGatewayConnectionOptions *GetTransitGatewayConnectionOptions) - Operation response error`, func() {
		version := CreateMockDate()
		getTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTransitGatewayConnection with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionOptions model
				getTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				getTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTransitGatewayConnection(getTransitGatewayConnectionOptions *GetTransitGatewayConnectionOptions)`, func() {
		version := CreateMockDate()
		getTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "network_account_id": "28e4d90ac7504be694471ee66e70d0d5", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "created_at": "2019-01-01T12:00:00", "request_status": "pending", "status": "attached", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetTransitGatewayConnection successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetTransitGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionOptions model
				getTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				getTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetTransitGatewayConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetTransitGatewayConnectionOptions model
				getTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				getTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				getTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTransitGatewayConnectionOptions model with no property values
				getTransitGatewayConnectionOptionsModelNew := new(transitgatewayapisv1.GetTransitGatewayConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetTransitGatewayConnection(getTransitGatewayConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions *UpdateTransitGatewayConnectionOptions) - Operation response error`, func() {
		version := CreateMockDate()
		updateTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnection with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				updateTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				updateTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptions *UpdateTransitGatewayConnectionOptions)`, func() {
		version := CreateMockDate()
		updateTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("PATCH"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "network_account_id": "28e4d90ac7504be694471ee66e70d0d5", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531", "created_at": "2019-01-01T12:00:00", "request_status": "pending", "status": "attached", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnection successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateTransitGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				updateTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				updateTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateTransitGatewayConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				updateTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				updateTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				updateTransitGatewayConnectionOptionsModel.Name = core.StringPtr("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTransitGatewayConnectionOptions model with no property values
				updateTransitGatewayConnectionOptionsModelNew := new(transitgatewayapisv1.UpdateTransitGatewayConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateTransitGatewayConnection(updateTransitGatewayConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateTransitGatewayConnectionActions(createTransitGatewayConnectionActionsOptions *CreateTransitGatewayConnectionActionsOptions)`, func() {
		version := CreateMockDate()
		createTransitGatewayConnectionActionsPath := "/transit_gateways/testString/connections/testString/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createTransitGatewayConnectionActionsPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for version query parameter

					res.WriteHeader(204)
				}))
			})
			It(`Invoke CreateTransitGatewayConnectionActions successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.CreateTransitGatewayConnectionActions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionActionsOptions model
				createTransitGatewayConnectionActionsOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionActionsOptions)
				createTransitGatewayConnectionActionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionActionsOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayConnectionActionsOptionsModel.Action = core.StringPtr("approve")
				createTransitGatewayConnectionActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.CreateTransitGatewayConnectionActions(createTransitGatewayConnectionActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CreateTransitGatewayConnectionActions with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateTransitGatewayConnectionActionsOptions model
				createTransitGatewayConnectionActionsOptionsModel := new(transitgatewayapisv1.CreateTransitGatewayConnectionActionsOptions)
				createTransitGatewayConnectionActionsOptionsModel.TransitGatewayID = core.StringPtr("testString")
				createTransitGatewayConnectionActionsOptionsModel.ID = core.StringPtr("testString")
				createTransitGatewayConnectionActionsOptionsModel.Action = core.StringPtr("approve")
				createTransitGatewayConnectionActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.CreateTransitGatewayConnectionActions(createTransitGatewayConnectionActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CreateTransitGatewayConnectionActionsOptions model with no property values
				createTransitGatewayConnectionActionsOptionsModelNew := new(transitgatewayapisv1.CreateTransitGatewayConnectionActionsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.CreateTransitGatewayConnectionActions(createTransitGatewayConnectionActionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := CreateMockDate()
		It(`Instantiate service client`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       version,
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:     "{BAD_URL_STRING",
				Version: version,
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:     "https://transitgatewayapisv1/api",
				Version: version,
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := CreateMockDate()
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_APIS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: version,
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:     "https://testService/api",
					Version: version,
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
					Version: version,
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
				"TRANSIT_GATEWAY_APIS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				Version: version,
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
				"TRANSIT_GATEWAY_APIS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:     "{BAD_URL_STRING",
				Version: version,
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListGatewayLocations(listGatewayLocationsOptions *ListGatewayLocationsOptions) - Operation response error`, func() {
		version := CreateMockDate()
		listGatewayLocationsPath := "/locations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewayLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListGatewayLocations with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := new(transitgatewayapisv1.ListGatewayLocationsOptions)
				listGatewayLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListGatewayLocations(listGatewayLocationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListGatewayLocations(listGatewayLocationsOptions *ListGatewayLocationsOptions)`, func() {
		version := CreateMockDate()
		listGatewayLocationsPath := "/locations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listGatewayLocationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"locations": [{"billing_location": "us", "name": "us-south", "type": "region"}]}`)
				}))
			})
			It(`Invoke ListGatewayLocations successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListGatewayLocations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := new(transitgatewayapisv1.ListGatewayLocationsOptions)
				listGatewayLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListGatewayLocations(listGatewayLocationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListGatewayLocations with error: Operation request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := new(transitgatewayapisv1.ListGatewayLocationsOptions)
				listGatewayLocationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListGatewayLocations(listGatewayLocationsOptionsModel)
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
	Describe(`GetGatewayLocation(getGatewayLocationOptions *GetGatewayLocationOptions) - Operation response error`, func() {
		version := CreateMockDate()
		getGatewayLocationPath := "/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayLocationPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGatewayLocation with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetGatewayLocationOptions model
				getGatewayLocationOptionsModel := new(transitgatewayapisv1.GetGatewayLocationOptions)
				getGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				getGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetGatewayLocation(getGatewayLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetGatewayLocation(getGatewayLocationOptions *GetGatewayLocationOptions)`, func() {
		version := CreateMockDate()
		getGatewayLocationPath := "/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getGatewayLocationPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"billing_location": "us", "name": "us-south", "type": "region", "local_connection_locations": [{"display_name": "Dallas", "name": "us-south", "type": "region"}]}`)
				}))
			})
			It(`Invoke GetGatewayLocation successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetGatewayLocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGatewayLocationOptions model
				getGatewayLocationOptionsModel := new(transitgatewayapisv1.GetGatewayLocationOptions)
				getGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				getGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetGatewayLocation(getGatewayLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetGatewayLocation with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetGatewayLocationOptions model
				getGatewayLocationOptionsModel := new(transitgatewayapisv1.GetGatewayLocationOptions)
				getGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				getGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetGatewayLocation(getGatewayLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGatewayLocationOptions model with no property values
				getGatewayLocationOptionsModelNew := new(transitgatewayapisv1.GetGatewayLocationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetGatewayLocation(getGatewayLocationOptionsModelNew)
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
			version := CreateMockDate()
			testService, _ := transitgatewayapisv1.NewTransitGatewayApisV1(&transitgatewayapisv1.TransitGatewayApisV1Options{
				URL:           "http://transitgatewayapisv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       version,
			})
			It(`Invoke NewCreateTransitGatewayConnectionActionsOptions successfully`, func() {
				// Construct an instance of the CreateTransitGatewayConnectionActionsOptions model
				transitGatewayID := "testString"
				id := "testString"
				createTransitGatewayConnectionActionsOptionsAction := "approve"
				createTransitGatewayConnectionActionsOptionsModel := testService.NewCreateTransitGatewayConnectionActionsOptions(transitGatewayID, id, createTransitGatewayConnectionActionsOptionsAction)
				createTransitGatewayConnectionActionsOptionsModel.SetTransitGatewayID("testString")
				createTransitGatewayConnectionActionsOptionsModel.SetID("testString")
				createTransitGatewayConnectionActionsOptionsModel.SetAction("approve")
				createTransitGatewayConnectionActionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayConnectionActionsOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayConnectionActionsOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionActionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionActionsOptionsModel.Action).To(Equal(core.StringPtr("approve")))
				Expect(createTransitGatewayConnectionActionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the CreateTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				createTransitGatewayConnectionOptionsNetworkType := "vpc"
				createTransitGatewayConnectionOptionsModel := testService.NewCreateTransitGatewayConnectionOptions(transitGatewayID, createTransitGatewayConnectionOptionsNetworkType)
				createTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				createTransitGatewayConnectionOptionsModel.SetNetworkType("vpc")
				createTransitGatewayConnectionOptionsModel.SetName("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayConnectionOptionsModel.SetNetworkID("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")
				createTransitGatewayConnectionOptionsModel.SetNetworkAccountID("28e4d90ac7504be694471ee66e70d0d5")
				createTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkType).To(Equal(core.StringPtr("vpc")))
				Expect(createTransitGatewayConnectionOptionsModel.Name).To(Equal(core.StringPtr("Transit_Service_BWTN_SJ_DL")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkID).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkAccountID).To(Equal(core.StringPtr("28e4d90ac7504be694471ee66e70d0d5")))
				Expect(createTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTransitGatewayOptions successfully`, func() {
				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("56969d6043e9465c883cb9f7363e78e8")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("56969d6043e9465c883cb9f7363e78e8")))

				// Construct an instance of the CreateTransitGatewayOptions model
				createTransitGatewayOptionsLocation := "us-south"
				createTransitGatewayOptionsName := "Transit_Service_BWTN_SJ_DL"
				createTransitGatewayOptionsModel := testService.NewCreateTransitGatewayOptions(createTransitGatewayOptionsLocation, createTransitGatewayOptionsName)
				createTransitGatewayOptionsModel.SetLocation("us-south")
				createTransitGatewayOptionsModel.SetName("Transit_Service_BWTN_SJ_DL")
				createTransitGatewayOptionsModel.SetGlobal(true)
				createTransitGatewayOptionsModel.SetResourceGroup(resourceGroupIdentityModel)
				createTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayOptionsModel.Location).To(Equal(core.StringPtr("us-south")))
				Expect(createTransitGatewayOptionsModel.Name).To(Equal(core.StringPtr("Transit_Service_BWTN_SJ_DL")))
				Expect(createTransitGatewayOptionsModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(createTransitGatewayOptionsModel.ResourceGroup).To(Equal(resourceGroupIdentityModel))
				Expect(createTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the DeleteTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				id := "testString"
				deleteTransitGatewayConnectionOptionsModel := testService.NewDeleteTransitGatewayConnectionOptions(transitGatewayID, id)
				deleteTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				deleteTransitGatewayConnectionOptionsModel.SetID("testString")
				deleteTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(deleteTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTransitGatewayOptions successfully`, func() {
				// Construct an instance of the DeleteTransitGatewayOptions model
				id := "testString"
				deleteTransitGatewayOptionsModel := testService.NewDeleteTransitGatewayOptions(id)
				deleteTransitGatewayOptionsModel.SetID("testString")
				deleteTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(deleteTransitGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGatewayLocationOptions successfully`, func() {
				// Construct an instance of the GetGatewayLocationOptions model
				name := "testString"
				getGatewayLocationOptionsModel := testService.NewGetGatewayLocationOptions(name)
				getGatewayLocationOptionsModel.SetName("testString")
				getGatewayLocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGatewayLocationOptionsModel).ToNot(BeNil())
				Expect(getGatewayLocationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(getGatewayLocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the GetTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				id := "testString"
				getTransitGatewayConnectionOptionsModel := testService.NewGetTransitGatewayConnectionOptions(transitGatewayID, id)
				getTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				getTransitGatewayConnectionOptionsModel.SetID("testString")
				getTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(getTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTransitGatewayOptions successfully`, func() {
				// Construct an instance of the GetTransitGatewayOptions model
				id := "testString"
				getTransitGatewayOptionsModel := testService.NewGetTransitGatewayOptions(id)
				getTransitGatewayOptionsModel.SetID("testString")
				getTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(getTransitGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListGatewayLocationsOptions successfully`, func() {
				// Construct an instance of the ListGatewayLocationsOptions model
				listGatewayLocationsOptionsModel := testService.NewListGatewayLocationsOptions()
				listGatewayLocationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listGatewayLocationsOptionsModel).ToNot(BeNil())
				Expect(listGatewayLocationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTransitGatewayConnectionsOptions successfully`, func() {
				// Construct an instance of the ListTransitGatewayConnectionsOptions model
				transitGatewayID := "testString"
				listTransitGatewayConnectionsOptionsModel := testService.NewListTransitGatewayConnectionsOptions(transitGatewayID)
				listTransitGatewayConnectionsOptionsModel.SetTransitGatewayID("testString")
				listTransitGatewayConnectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTransitGatewayConnectionsOptionsModel).ToNot(BeNil())
				Expect(listTransitGatewayConnectionsOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(listTransitGatewayConnectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTransitGatewaysOptions successfully`, func() {
				// Construct an instance of the ListTransitGatewaysOptions model
				listTransitGatewaysOptionsModel := testService.NewListTransitGatewaysOptions()
				listTransitGatewaysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTransitGatewaysOptionsModel).ToNot(BeNil())
				Expect(listTransitGatewaysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResourceGroupIdentity successfully`, func() {
				id := "56969d6043e9465c883cb9f7363e78e8"
				model, err := testService.NewResourceGroupIdentity(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the UpdateTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				id := "testString"
				updateTransitGatewayConnectionOptionsModel := testService.NewUpdateTransitGatewayConnectionOptions(transitGatewayID, id)
				updateTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				updateTransitGatewayConnectionOptionsModel.SetID("testString")
				updateTransitGatewayConnectionOptionsModel.SetName("Transit_Service_BWTN_SJ_DL")
				updateTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(updateTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayConnectionOptionsModel.Name).To(Equal(core.StringPtr("Transit_Service_BWTN_SJ_DL")))
				Expect(updateTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTransitGatewayOptions successfully`, func() {
				// Construct an instance of the UpdateTransitGatewayOptions model
				id := "testString"
				updateTransitGatewayOptionsModel := testService.NewUpdateTransitGatewayOptions(id)
				updateTransitGatewayOptionsModel.SetID("testString")
				updateTransitGatewayOptionsModel.SetGlobal(true)
				updateTransitGatewayOptionsModel.SetName("my-transit-gateway")
				updateTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(updateTransitGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTransitGatewayOptionsModel.Global).To(Equal(core.BoolPtr(true)))
				Expect(updateTransitGatewayOptionsModel.Name).To(Equal(core.StringPtr("my-transit-gateway")))
				Expect(updateTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
