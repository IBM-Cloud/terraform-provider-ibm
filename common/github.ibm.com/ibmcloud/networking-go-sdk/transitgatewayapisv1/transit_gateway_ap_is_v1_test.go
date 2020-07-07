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
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/transitgatewayapisv1"
)

var _ = Describe(`TransitGatewayApIsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := CreateMockDate()
		It(`Instantiate service client`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       version,
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
				URL:     "{BAD_URL_STRING",
				Version: version,
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := CreateMockDate()
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_AP_IS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_AP_IS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					Version: version,
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				"TRANSIT_GATEWAY_AP_IS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_AP_IS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				"TRANSIT_GATEWAY_AP_IS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
					fmt.Fprintf(res, `{"transit_gateways": [{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "global": true, "location": "us-south", "created_at": "2019-01-01T12:00:00", "resource_group": {"id": "56969d60-43e9-465c-883c-b9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d60-43e9-465c-883c-b9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListTransitGateways successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d60-43e9-465c-883c-b9f7363e78e8")

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
					fmt.Fprintf(res, `{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "global": true, "location": "us-south", "created_at": "2019-01-01T12:00:00", "resource_group": {"id": "56969d60-43e9-465c-883c-b9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d60-43e9-465c-883c-b9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateTransitGateway successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				resourceGroupIdentityModel.ID = core.StringPtr("56969d60-43e9-465c-883c-b9f7363e78e8")

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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("56969d60-43e9-465c-883c-b9f7363e78e8")

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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
	Describe(`DetailTransitGateway(detailTransitGatewayOptions *DetailTransitGatewayOptions) - Operation response error`, func() {
		version := CreateMockDate()
		detailTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(detailTransitGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DetailTransitGateway with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DetailTransitGatewayOptions model
				detailTransitGatewayOptionsModel := new(transitgatewayapisv1.DetailTransitGatewayOptions)
				detailTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				detailTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DetailTransitGateway(detailTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DetailTransitGateway(detailTransitGatewayOptions *DetailTransitGatewayOptions)`, func() {
		version := CreateMockDate()
		detailTransitGatewayPath := "/transit_gateways/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(detailTransitGatewayPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "global": true, "location": "us-south", "created_at": "2019-01-01T12:00:00", "resource_group": {"id": "56969d60-43e9-465c-883c-b9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d60-43e9-465c-883c-b9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke DetailTransitGateway successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DetailTransitGateway(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DetailTransitGatewayOptions model
				detailTransitGatewayOptionsModel := new(transitgatewayapisv1.DetailTransitGatewayOptions)
				detailTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				detailTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DetailTransitGateway(detailTransitGatewayOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DetailTransitGateway with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DetailTransitGatewayOptions model
				detailTransitGatewayOptionsModel := new(transitgatewayapisv1.DetailTransitGatewayOptions)
				detailTransitGatewayOptionsModel.ID = core.StringPtr("testString")
				detailTransitGatewayOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DetailTransitGateway(detailTransitGatewayOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DetailTransitGatewayOptions model with no property values
				detailTransitGatewayOptionsModelNew := new(transitgatewayapisv1.DetailTransitGatewayOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DetailTransitGateway(detailTransitGatewayOptionsModelNew)
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
					fmt.Fprintf(res, `{"id": "ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "crn": "crn:v1:bluemix:public:transit:dal03:a/57a7d05f36894e3cb9b46a43556d903e::gateway:ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4", "name": "my-transit-gateway-in-TransitGateway", "global": true, "location": "us-south", "created_at": "2019-01-01T12:00:00", "resource_group": {"id": "56969d60-43e9-465c-883c-b9f7363e78e8", "href": "https://resource-manager.bluemix.net/v1/resource_groups/56969d60-43e9-465c-883c-b9f7363e78e8"}, "status": "available", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateTransitGateway successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       version,
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
				URL:     "{BAD_URL_STRING",
				Version: version,
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := CreateMockDate()
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_AP_IS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_AP_IS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					Version: version,
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				"TRANSIT_GATEWAY_AP_IS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_AP_IS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				"TRANSIT_GATEWAY_AP_IS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
					fmt.Fprintf(res, `{"connections": [{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531-connection_id", "created_at": "2019-01-01T12:00:00", "status": "attached", "updated_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListTransitGatewayConnections successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
					fmt.Fprintf(res, `{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531-connection_id", "created_at": "2019-01-01T12:00:00", "status": "attached", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateTransitGatewayConnection successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				createTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateTransitGatewayConnection(createTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateTransitGatewayConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
	Describe(`DetailTransitGatewayConnection(detailTransitGatewayConnectionOptions *DetailTransitGatewayConnectionOptions) - Operation response error`, func() {
		version := CreateMockDate()
		detailTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(detailTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DetailTransitGatewayConnection with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DetailTransitGatewayConnectionOptions model
				detailTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.DetailTransitGatewayConnectionOptions)
				detailTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				detailTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				detailTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DetailTransitGatewayConnection(detailTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DetailTransitGatewayConnection(detailTransitGatewayConnectionOptions *DetailTransitGatewayConnectionOptions)`, func() {
		version := CreateMockDate()
		detailTransitGatewayConnectionPath := "/transit_gateways/testString/connections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(detailTransitGatewayConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531-connection_id", "created_at": "2019-01-01T12:00:00", "status": "attached", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke DetailTransitGatewayConnection successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DetailTransitGatewayConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DetailTransitGatewayConnectionOptions model
				detailTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.DetailTransitGatewayConnectionOptions)
				detailTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				detailTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				detailTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DetailTransitGatewayConnection(detailTransitGatewayConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DetailTransitGatewayConnection with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DetailTransitGatewayConnectionOptions model
				detailTransitGatewayConnectionOptionsModel := new(transitgatewayapisv1.DetailTransitGatewayConnectionOptions)
				detailTransitGatewayConnectionOptionsModel.TransitGatewayID = core.StringPtr("testString")
				detailTransitGatewayConnectionOptionsModel.ID = core.StringPtr("testString")
				detailTransitGatewayConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DetailTransitGatewayConnection(detailTransitGatewayConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DetailTransitGatewayConnectionOptions model with no property values
				detailTransitGatewayConnectionOptionsModelNew := new(transitgatewayapisv1.DetailTransitGatewayConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DetailTransitGatewayConnection(detailTransitGatewayConnectionOptionsModelNew)
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
					fmt.Fprintf(res, `{"name": "Transit_Service_BWTN_SJ_DL", "network_id": "crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b", "network_type": "vpc", "id": "1a15dca5-7e33-45e1-b7c5-bc690e569531-connection_id", "created_at": "2019-01-01T12:00:00", "status": "attached", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateTransitGatewayConnection successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
	Describe(`Service constructor tests`, func() {
		version := CreateMockDate()
		It(`Instantiate service client`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       version,
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
				URL:     "{BAD_URL_STRING",
				Version: version,
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := CreateMockDate()
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TRANSIT_GATEWAY_AP_IS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_AP_IS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					Version: version,
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				"TRANSIT_GATEWAY_AP_IS_URL":       "https://transitgatewayapisv1/api",
				"TRANSIT_GATEWAY_AP_IS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				"TRANSIT_GATEWAY_AP_IS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1UsingExternalConfig(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
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
	Describe(`DetailGatewayLocation(detailGatewayLocationOptions *DetailGatewayLocationOptions) - Operation response error`, func() {
		version := CreateMockDate()
		detailGatewayLocationPath := "/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(detailGatewayLocationPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DetailGatewayLocation with error: Operation response processing error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DetailGatewayLocationOptions model
				detailGatewayLocationOptionsModel := new(transitgatewayapisv1.DetailGatewayLocationOptions)
				detailGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				detailGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DetailGatewayLocation(detailGatewayLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DetailGatewayLocation(detailGatewayLocationOptions *DetailGatewayLocationOptions)`, func() {
		version := CreateMockDate()
		detailGatewayLocationPath := "/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(detailGatewayLocationPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for version query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"billing_location": "us", "name": "us-south", "type": "region", "local_connection_locations": [{"display_name": "Dallas", "name": "us-south", "type": "region"}]}`)
				}))
			})
			It(`Invoke DetailGatewayLocation successfully`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DetailGatewayLocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DetailGatewayLocationOptions model
				detailGatewayLocationOptionsModel := new(transitgatewayapisv1.DetailGatewayLocationOptions)
				detailGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				detailGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DetailGatewayLocation(detailGatewayLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DetailGatewayLocation with error: Operation validation and request error`, func() {
				testService, testServiceErr := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DetailGatewayLocationOptions model
				detailGatewayLocationOptionsModel := new(transitgatewayapisv1.DetailGatewayLocationOptions)
				detailGatewayLocationOptionsModel.Name = core.StringPtr("testString")
				detailGatewayLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DetailGatewayLocation(detailGatewayLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DetailGatewayLocationOptions model with no property values
				detailGatewayLocationOptionsModelNew := new(transitgatewayapisv1.DetailGatewayLocationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DetailGatewayLocation(detailGatewayLocationOptionsModelNew)
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
			testService, _ := transitgatewayapisv1.NewTransitGatewayApIsV1(&transitgatewayapisv1.TransitGatewayApIsV1Options{
				URL:           "http://transitgatewayapisv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       version,
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
				createTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(createTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkType).To(Equal(core.StringPtr("vpc")))
				Expect(createTransitGatewayConnectionOptionsModel.Name).To(Equal(core.StringPtr("Transit_Service_BWTN_SJ_DL")))
				Expect(createTransitGatewayConnectionOptionsModel.NetworkID).To(Equal(core.StringPtr("crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b")))
				Expect(createTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTransitGatewayOptions successfully`, func() {
				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(transitgatewayapisv1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("56969d60-43e9-465c-883c-b9f7363e78e8")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("56969d60-43e9-465c-883c-b9f7363e78e8")))

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
			It(`Invoke NewDetailGatewayLocationOptions successfully`, func() {
				// Construct an instance of the DetailGatewayLocationOptions model
				name := "testString"
				detailGatewayLocationOptionsModel := testService.NewDetailGatewayLocationOptions(name)
				detailGatewayLocationOptionsModel.SetName("testString")
				detailGatewayLocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(detailGatewayLocationOptionsModel).ToNot(BeNil())
				Expect(detailGatewayLocationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(detailGatewayLocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDetailTransitGatewayConnectionOptions successfully`, func() {
				// Construct an instance of the DetailTransitGatewayConnectionOptions model
				transitGatewayID := "testString"
				id := "testString"
				detailTransitGatewayConnectionOptionsModel := testService.NewDetailTransitGatewayConnectionOptions(transitGatewayID, id)
				detailTransitGatewayConnectionOptionsModel.SetTransitGatewayID("testString")
				detailTransitGatewayConnectionOptionsModel.SetID("testString")
				detailTransitGatewayConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(detailTransitGatewayConnectionOptionsModel).ToNot(BeNil())
				Expect(detailTransitGatewayConnectionOptionsModel.TransitGatewayID).To(Equal(core.StringPtr("testString")))
				Expect(detailTransitGatewayConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(detailTransitGatewayConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDetailTransitGatewayOptions successfully`, func() {
				// Construct an instance of the DetailTransitGatewayOptions model
				id := "testString"
				detailTransitGatewayOptionsModel := testService.NewDetailTransitGatewayOptions(id)
				detailTransitGatewayOptionsModel.SetID("testString")
				detailTransitGatewayOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(detailTransitGatewayOptionsModel).ToNot(BeNil())
				Expect(detailTransitGatewayOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(detailTransitGatewayOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				id := "56969d60-43e9-465c-883c-b9f7363e78e8"
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
