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

package globalloadbalancerpoolsv0_test

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
	"github.ibm.com/ibmcloud/networking-go-sdk/globalloadbalancerpoolsv0"
)

var _ = Describe(`GlobalLoadBalancerPoolsV0`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
				URL: "https://globalloadbalancerpoolsv0/api",
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
			testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCER_POOLS_URL":       "https://globalloadbalancerpoolsv0/api",
				"GLOBAL_LOAD_BALANCER_POOLS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0UsingExternalConfig(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					Crn: core.StringPtr(crn),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0UsingExternalConfig(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
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
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0UsingExternalConfig(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
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
				"GLOBAL_LOAD_BALANCER_POOLS_URL":       "https://globalloadbalancerpoolsv0/api",
				"GLOBAL_LOAD_BALANCER_POOLS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0UsingExternalConfig(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
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
				"GLOBAL_LOAD_BALANCER_POOLS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0UsingExternalConfig(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
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
	Describe(`ListAllLoadBalancerPools(listAllLoadBalancerPoolsOptions *ListAllLoadBalancerPoolsOptions) - Operation response error`, func() {
		crn := "testString"
		listAllLoadBalancerPoolsPath := "/v1/testString/load_balancers/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllLoadBalancerPoolsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllLoadBalancerPools with error: Operation response processing error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllLoadBalancerPoolsOptions model
				listAllLoadBalancerPoolsOptionsModel := new(globalloadbalancerpoolsv0.ListAllLoadBalancerPoolsOptions)
				listAllLoadBalancerPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListAllLoadBalancerPools(listAllLoadBalancerPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllLoadBalancerPools(listAllLoadBalancerPoolsOptions *ListAllLoadBalancerPoolsOptions)`, func() {
		crn := "testString"
		listAllLoadBalancerPoolsPath := "/v1/testString/load_balancers/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllLoadBalancerPoolsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "17b5962d775c646f3f9725cbc7a53df4", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "description": "Primary data center - Provider XYZ", "name": "primary-dc-1", "enabled": true, "healthy": true, "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc", "origins": [{"name": "app-server-1", "address": "0.0.0.0", "enabled": true, "healthy": true}], "notification_email": "someone@example.com"}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListAllLoadBalancerPools successfully`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAllLoadBalancerPools(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllLoadBalancerPoolsOptions model
				listAllLoadBalancerPoolsOptionsModel := new(globalloadbalancerpoolsv0.ListAllLoadBalancerPoolsOptions)
				listAllLoadBalancerPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAllLoadBalancerPools(listAllLoadBalancerPoolsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListAllLoadBalancerPools with error: Operation request error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllLoadBalancerPoolsOptions model
				listAllLoadBalancerPoolsOptionsModel := new(globalloadbalancerpoolsv0.ListAllLoadBalancerPoolsOptions)
				listAllLoadBalancerPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListAllLoadBalancerPools(listAllLoadBalancerPoolsOptionsModel)
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
	Describe(`CreateLoadBalancerPool(createLoadBalancerPoolOptions *CreateLoadBalancerPoolOptions) - Operation response error`, func() {
		crn := "testString"
		createLoadBalancerPoolPath := "/v1/testString/load_balancers/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createLoadBalancerPoolPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLoadBalancerPool with error: Operation response processing error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerPoolReqOriginsItem model
				loadBalancerPoolReqOriginsItemModel := new(globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem)
				loadBalancerPoolReqOriginsItemModel.Name = core.StringPtr("app-server-1")
				loadBalancerPoolReqOriginsItemModel.Address = core.StringPtr("0.0.0.0")
				loadBalancerPoolReqOriginsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateLoadBalancerPoolOptions model
				createLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.CreateLoadBalancerPoolOptions)
				createLoadBalancerPoolOptionsModel.Name = core.StringPtr("primary-dc-1")
				createLoadBalancerPoolOptionsModel.CheckRegions = []string{"WNAM"}
				createLoadBalancerPoolOptionsModel.Origins = []globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel}
				createLoadBalancerPoolOptionsModel.Description = core.StringPtr("Primary data center - Provider XYZ")
				createLoadBalancerPoolOptionsModel.MinimumOrigins = core.Int64Ptr(int64(2))
				createLoadBalancerPoolOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerPoolOptionsModel.Monitor = core.StringPtr("f1aba936b94213e5b8dca0c0dbf1f9cc")
				createLoadBalancerPoolOptionsModel.NotificationEmail = core.StringPtr("someone@example.com")
				createLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateLoadBalancerPool(createLoadBalancerPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateLoadBalancerPool(createLoadBalancerPoolOptions *CreateLoadBalancerPoolOptions)`, func() {
		crn := "testString"
		createLoadBalancerPoolPath := "/v1/testString/load_balancers/pools"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createLoadBalancerPoolPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "17b5962d775c646f3f9725cbc7a53df4", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "description": "Primary data center - Provider XYZ", "name": "primary-dc-1", "enabled": true, "healthy": true, "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc", "origins": [{"name": "app-server-1", "address": "0.0.0.0", "enabled": true, "healthy": true}], "notification_email": "someone@example.com"}, "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke CreateLoadBalancerPool successfully`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateLoadBalancerPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerPoolReqOriginsItem model
				loadBalancerPoolReqOriginsItemModel := new(globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem)
				loadBalancerPoolReqOriginsItemModel.Name = core.StringPtr("app-server-1")
				loadBalancerPoolReqOriginsItemModel.Address = core.StringPtr("0.0.0.0")
				loadBalancerPoolReqOriginsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateLoadBalancerPoolOptions model
				createLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.CreateLoadBalancerPoolOptions)
				createLoadBalancerPoolOptionsModel.Name = core.StringPtr("primary-dc-1")
				createLoadBalancerPoolOptionsModel.CheckRegions = []string{"WNAM"}
				createLoadBalancerPoolOptionsModel.Origins = []globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel}
				createLoadBalancerPoolOptionsModel.Description = core.StringPtr("Primary data center - Provider XYZ")
				createLoadBalancerPoolOptionsModel.MinimumOrigins = core.Int64Ptr(int64(2))
				createLoadBalancerPoolOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerPoolOptionsModel.Monitor = core.StringPtr("f1aba936b94213e5b8dca0c0dbf1f9cc")
				createLoadBalancerPoolOptionsModel.NotificationEmail = core.StringPtr("someone@example.com")
				createLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateLoadBalancerPool(createLoadBalancerPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateLoadBalancerPool with error: Operation request error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerPoolReqOriginsItem model
				loadBalancerPoolReqOriginsItemModel := new(globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem)
				loadBalancerPoolReqOriginsItemModel.Name = core.StringPtr("app-server-1")
				loadBalancerPoolReqOriginsItemModel.Address = core.StringPtr("0.0.0.0")
				loadBalancerPoolReqOriginsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateLoadBalancerPoolOptions model
				createLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.CreateLoadBalancerPoolOptions)
				createLoadBalancerPoolOptionsModel.Name = core.StringPtr("primary-dc-1")
				createLoadBalancerPoolOptionsModel.CheckRegions = []string{"WNAM"}
				createLoadBalancerPoolOptionsModel.Origins = []globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel}
				createLoadBalancerPoolOptionsModel.Description = core.StringPtr("Primary data center - Provider XYZ")
				createLoadBalancerPoolOptionsModel.MinimumOrigins = core.Int64Ptr(int64(2))
				createLoadBalancerPoolOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerPoolOptionsModel.Monitor = core.StringPtr("f1aba936b94213e5b8dca0c0dbf1f9cc")
				createLoadBalancerPoolOptionsModel.NotificationEmail = core.StringPtr("someone@example.com")
				createLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateLoadBalancerPool(createLoadBalancerPoolOptionsModel)
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
	Describe(`GetLoadBalancerPool(getLoadBalancerPoolOptions *GetLoadBalancerPoolOptions) - Operation response error`, func() {
		crn := "testString"
		getLoadBalancerPoolPath := "/v1/testString/load_balancers/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getLoadBalancerPoolPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLoadBalancerPool with error: Operation response processing error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerPoolOptions model
				getLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.GetLoadBalancerPoolOptions)
				getLoadBalancerPoolOptionsModel.PoolIdentifier = core.StringPtr("testString")
				getLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetLoadBalancerPool(getLoadBalancerPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLoadBalancerPool(getLoadBalancerPoolOptions *GetLoadBalancerPoolOptions)`, func() {
		crn := "testString"
		getLoadBalancerPoolPath := "/v1/testString/load_balancers/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getLoadBalancerPoolPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "17b5962d775c646f3f9725cbc7a53df4", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "description": "Primary data center - Provider XYZ", "name": "primary-dc-1", "enabled": true, "healthy": true, "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc", "origins": [{"name": "app-server-1", "address": "0.0.0.0", "enabled": true, "healthy": true}], "notification_email": "someone@example.com"}, "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke GetLoadBalancerPool successfully`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoadBalancerPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerPoolOptions model
				getLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.GetLoadBalancerPoolOptions)
				getLoadBalancerPoolOptionsModel.PoolIdentifier = core.StringPtr("testString")
				getLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoadBalancerPool(getLoadBalancerPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetLoadBalancerPool with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerPoolOptions model
				getLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.GetLoadBalancerPoolOptions)
				getLoadBalancerPoolOptionsModel.PoolIdentifier = core.StringPtr("testString")
				getLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetLoadBalancerPool(getLoadBalancerPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLoadBalancerPoolOptions model with no property values
				getLoadBalancerPoolOptionsModelNew := new(globalloadbalancerpoolsv0.GetLoadBalancerPoolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetLoadBalancerPool(getLoadBalancerPoolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLoadBalancerPool(deleteLoadBalancerPoolOptions *DeleteLoadBalancerPoolOptions) - Operation response error`, func() {
		crn := "testString"
		deleteLoadBalancerPoolPath := "/v1/testString/load_balancers/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteLoadBalancerPoolPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLoadBalancerPool with error: Operation response processing error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteLoadBalancerPoolOptions model
				deleteLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.DeleteLoadBalancerPoolOptions)
				deleteLoadBalancerPoolOptionsModel.PoolIdentifier = core.StringPtr("testString")
				deleteLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteLoadBalancerPool(deleteLoadBalancerPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLoadBalancerPool(deleteLoadBalancerPoolOptions *DeleteLoadBalancerPoolOptions)`, func() {
		crn := "testString"
		deleteLoadBalancerPoolPath := "/v1/testString/load_balancers/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteLoadBalancerPoolPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "17b5962d775c646f3f9725cbc7a53df4"}}`)
				}))
			})
			It(`Invoke DeleteLoadBalancerPool successfully`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteLoadBalancerPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerPoolOptions model
				deleteLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.DeleteLoadBalancerPoolOptions)
				deleteLoadBalancerPoolOptionsModel.PoolIdentifier = core.StringPtr("testString")
				deleteLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteLoadBalancerPool(deleteLoadBalancerPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteLoadBalancerPool with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteLoadBalancerPoolOptions model
				deleteLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.DeleteLoadBalancerPoolOptions)
				deleteLoadBalancerPoolOptionsModel.PoolIdentifier = core.StringPtr("testString")
				deleteLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteLoadBalancerPool(deleteLoadBalancerPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteLoadBalancerPoolOptions model with no property values
				deleteLoadBalancerPoolOptionsModelNew := new(globalloadbalancerpoolsv0.DeleteLoadBalancerPoolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteLoadBalancerPool(deleteLoadBalancerPoolOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditLoadBalancerPool(editLoadBalancerPoolOptions *EditLoadBalancerPoolOptions) - Operation response error`, func() {
		crn := "testString"
		editLoadBalancerPoolPath := "/v1/testString/load_balancers/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editLoadBalancerPoolPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditLoadBalancerPool with error: Operation response processing error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerPoolReqOriginsItem model
				loadBalancerPoolReqOriginsItemModel := new(globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem)
				loadBalancerPoolReqOriginsItemModel.Name = core.StringPtr("app-server-1")
				loadBalancerPoolReqOriginsItemModel.Address = core.StringPtr("0.0.0.0")
				loadBalancerPoolReqOriginsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the EditLoadBalancerPoolOptions model
				editLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.EditLoadBalancerPoolOptions)
				editLoadBalancerPoolOptionsModel.PoolIdentifier = core.StringPtr("testString")
				editLoadBalancerPoolOptionsModel.Name = core.StringPtr("primary-dc-1")
				editLoadBalancerPoolOptionsModel.CheckRegions = []string{"WNAM"}
				editLoadBalancerPoolOptionsModel.Origins = []globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel}
				editLoadBalancerPoolOptionsModel.Description = core.StringPtr("Primary data center - Provider XYZ")
				editLoadBalancerPoolOptionsModel.MinimumOrigins = core.Int64Ptr(int64(2))
				editLoadBalancerPoolOptionsModel.Enabled = core.BoolPtr(true)
				editLoadBalancerPoolOptionsModel.Monitor = core.StringPtr("f1aba936b94213e5b8dca0c0dbf1f9cc")
				editLoadBalancerPoolOptionsModel.NotificationEmail = core.StringPtr("someone@example.com")
				editLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.EditLoadBalancerPool(editLoadBalancerPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditLoadBalancerPool(editLoadBalancerPoolOptions *EditLoadBalancerPoolOptions)`, func() {
		crn := "testString"
		editLoadBalancerPoolPath := "/v1/testString/load_balancers/pools/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editLoadBalancerPoolPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "17b5962d775c646f3f9725cbc7a53df4", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "description": "Primary data center - Provider XYZ", "name": "primary-dc-1", "enabled": true, "healthy": true, "monitor": "f1aba936b94213e5b8dca0c0dbf1f9cc", "origins": [{"name": "app-server-1", "address": "0.0.0.0", "enabled": true, "healthy": true}], "notification_email": "someone@example.com"}, "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke EditLoadBalancerPool successfully`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.EditLoadBalancerPool(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoadBalancerPoolReqOriginsItem model
				loadBalancerPoolReqOriginsItemModel := new(globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem)
				loadBalancerPoolReqOriginsItemModel.Name = core.StringPtr("app-server-1")
				loadBalancerPoolReqOriginsItemModel.Address = core.StringPtr("0.0.0.0")
				loadBalancerPoolReqOriginsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the EditLoadBalancerPoolOptions model
				editLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.EditLoadBalancerPoolOptions)
				editLoadBalancerPoolOptionsModel.PoolIdentifier = core.StringPtr("testString")
				editLoadBalancerPoolOptionsModel.Name = core.StringPtr("primary-dc-1")
				editLoadBalancerPoolOptionsModel.CheckRegions = []string{"WNAM"}
				editLoadBalancerPoolOptionsModel.Origins = []globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel}
				editLoadBalancerPoolOptionsModel.Description = core.StringPtr("Primary data center - Provider XYZ")
				editLoadBalancerPoolOptionsModel.MinimumOrigins = core.Int64Ptr(int64(2))
				editLoadBalancerPoolOptionsModel.Enabled = core.BoolPtr(true)
				editLoadBalancerPoolOptionsModel.Monitor = core.StringPtr("f1aba936b94213e5b8dca0c0dbf1f9cc")
				editLoadBalancerPoolOptionsModel.NotificationEmail = core.StringPtr("someone@example.com")
				editLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.EditLoadBalancerPool(editLoadBalancerPoolOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke EditLoadBalancerPool with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the LoadBalancerPoolReqOriginsItem model
				loadBalancerPoolReqOriginsItemModel := new(globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem)
				loadBalancerPoolReqOriginsItemModel.Name = core.StringPtr("app-server-1")
				loadBalancerPoolReqOriginsItemModel.Address = core.StringPtr("0.0.0.0")
				loadBalancerPoolReqOriginsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the EditLoadBalancerPoolOptions model
				editLoadBalancerPoolOptionsModel := new(globalloadbalancerpoolsv0.EditLoadBalancerPoolOptions)
				editLoadBalancerPoolOptionsModel.PoolIdentifier = core.StringPtr("testString")
				editLoadBalancerPoolOptionsModel.Name = core.StringPtr("primary-dc-1")
				editLoadBalancerPoolOptionsModel.CheckRegions = []string{"WNAM"}
				editLoadBalancerPoolOptionsModel.Origins = []globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel}
				editLoadBalancerPoolOptionsModel.Description = core.StringPtr("Primary data center - Provider XYZ")
				editLoadBalancerPoolOptionsModel.MinimumOrigins = core.Int64Ptr(int64(2))
				editLoadBalancerPoolOptionsModel.Enabled = core.BoolPtr(true)
				editLoadBalancerPoolOptionsModel.Monitor = core.StringPtr("f1aba936b94213e5b8dca0c0dbf1f9cc")
				editLoadBalancerPoolOptionsModel.NotificationEmail = core.StringPtr("someone@example.com")
				editLoadBalancerPoolOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.EditLoadBalancerPool(editLoadBalancerPoolOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditLoadBalancerPoolOptions model with no property values
				editLoadBalancerPoolOptionsModelNew := new(globalloadbalancerpoolsv0.EditLoadBalancerPoolOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.EditLoadBalancerPool(editLoadBalancerPoolOptionsModelNew)
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
			testService, _ := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(&globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
				URL:           "http://globalloadbalancerpoolsv0modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			It(`Invoke NewCreateLoadBalancerPoolOptions successfully`, func() {
				// Construct an instance of the LoadBalancerPoolReqOriginsItem model
				loadBalancerPoolReqOriginsItemModel := new(globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem)
				Expect(loadBalancerPoolReqOriginsItemModel).ToNot(BeNil())
				loadBalancerPoolReqOriginsItemModel.Name = core.StringPtr("app-server-1")
				loadBalancerPoolReqOriginsItemModel.Address = core.StringPtr("0.0.0.0")
				loadBalancerPoolReqOriginsItemModel.Enabled = core.BoolPtr(true)
				Expect(loadBalancerPoolReqOriginsItemModel.Name).To(Equal(core.StringPtr("app-server-1")))
				Expect(loadBalancerPoolReqOriginsItemModel.Address).To(Equal(core.StringPtr("0.0.0.0")))
				Expect(loadBalancerPoolReqOriginsItemModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateLoadBalancerPoolOptions model
				createLoadBalancerPoolOptionsModel := testService.NewCreateLoadBalancerPoolOptions()
				createLoadBalancerPoolOptionsModel.SetName("primary-dc-1")
				createLoadBalancerPoolOptionsModel.SetCheckRegions([]string{"WNAM"})
				createLoadBalancerPoolOptionsModel.SetOrigins([]globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel})
				createLoadBalancerPoolOptionsModel.SetDescription("Primary data center - Provider XYZ")
				createLoadBalancerPoolOptionsModel.SetMinimumOrigins(int64(2))
				createLoadBalancerPoolOptionsModel.SetEnabled(true)
				createLoadBalancerPoolOptionsModel.SetMonitor("f1aba936b94213e5b8dca0c0dbf1f9cc")
				createLoadBalancerPoolOptionsModel.SetNotificationEmail("someone@example.com")
				createLoadBalancerPoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLoadBalancerPoolOptionsModel).ToNot(BeNil())
				Expect(createLoadBalancerPoolOptionsModel.Name).To(Equal(core.StringPtr("primary-dc-1")))
				Expect(createLoadBalancerPoolOptionsModel.CheckRegions).To(Equal([]string{"WNAM"}))
				Expect(createLoadBalancerPoolOptionsModel.Origins).To(Equal([]globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel}))
				Expect(createLoadBalancerPoolOptionsModel.Description).To(Equal(core.StringPtr("Primary data center - Provider XYZ")))
				Expect(createLoadBalancerPoolOptionsModel.MinimumOrigins).To(Equal(core.Int64Ptr(int64(2))))
				Expect(createLoadBalancerPoolOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createLoadBalancerPoolOptionsModel.Monitor).To(Equal(core.StringPtr("f1aba936b94213e5b8dca0c0dbf1f9cc")))
				Expect(createLoadBalancerPoolOptionsModel.NotificationEmail).To(Equal(core.StringPtr("someone@example.com")))
				Expect(createLoadBalancerPoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLoadBalancerPoolOptions successfully`, func() {
				// Construct an instance of the DeleteLoadBalancerPoolOptions model
				poolIdentifier := "testString"
				deleteLoadBalancerPoolOptionsModel := testService.NewDeleteLoadBalancerPoolOptions(poolIdentifier)
				deleteLoadBalancerPoolOptionsModel.SetPoolIdentifier("testString")
				deleteLoadBalancerPoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLoadBalancerPoolOptionsModel).ToNot(BeNil())
				Expect(deleteLoadBalancerPoolOptionsModel.PoolIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerPoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditLoadBalancerPoolOptions successfully`, func() {
				// Construct an instance of the LoadBalancerPoolReqOriginsItem model
				loadBalancerPoolReqOriginsItemModel := new(globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem)
				Expect(loadBalancerPoolReqOriginsItemModel).ToNot(BeNil())
				loadBalancerPoolReqOriginsItemModel.Name = core.StringPtr("app-server-1")
				loadBalancerPoolReqOriginsItemModel.Address = core.StringPtr("0.0.0.0")
				loadBalancerPoolReqOriginsItemModel.Enabled = core.BoolPtr(true)
				Expect(loadBalancerPoolReqOriginsItemModel.Name).To(Equal(core.StringPtr("app-server-1")))
				Expect(loadBalancerPoolReqOriginsItemModel.Address).To(Equal(core.StringPtr("0.0.0.0")))
				Expect(loadBalancerPoolReqOriginsItemModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the EditLoadBalancerPoolOptions model
				poolIdentifier := "testString"
				editLoadBalancerPoolOptionsModel := testService.NewEditLoadBalancerPoolOptions(poolIdentifier)
				editLoadBalancerPoolOptionsModel.SetPoolIdentifier("testString")
				editLoadBalancerPoolOptionsModel.SetName("primary-dc-1")
				editLoadBalancerPoolOptionsModel.SetCheckRegions([]string{"WNAM"})
				editLoadBalancerPoolOptionsModel.SetOrigins([]globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel})
				editLoadBalancerPoolOptionsModel.SetDescription("Primary data center - Provider XYZ")
				editLoadBalancerPoolOptionsModel.SetMinimumOrigins(int64(2))
				editLoadBalancerPoolOptionsModel.SetEnabled(true)
				editLoadBalancerPoolOptionsModel.SetMonitor("f1aba936b94213e5b8dca0c0dbf1f9cc")
				editLoadBalancerPoolOptionsModel.SetNotificationEmail("someone@example.com")
				editLoadBalancerPoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editLoadBalancerPoolOptionsModel).ToNot(BeNil())
				Expect(editLoadBalancerPoolOptionsModel.PoolIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(editLoadBalancerPoolOptionsModel.Name).To(Equal(core.StringPtr("primary-dc-1")))
				Expect(editLoadBalancerPoolOptionsModel.CheckRegions).To(Equal([]string{"WNAM"}))
				Expect(editLoadBalancerPoolOptionsModel.Origins).To(Equal([]globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*loadBalancerPoolReqOriginsItemModel}))
				Expect(editLoadBalancerPoolOptionsModel.Description).To(Equal(core.StringPtr("Primary data center - Provider XYZ")))
				Expect(editLoadBalancerPoolOptionsModel.MinimumOrigins).To(Equal(core.Int64Ptr(int64(2))))
				Expect(editLoadBalancerPoolOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(editLoadBalancerPoolOptionsModel.Monitor).To(Equal(core.StringPtr("f1aba936b94213e5b8dca0c0dbf1f9cc")))
				Expect(editLoadBalancerPoolOptionsModel.NotificationEmail).To(Equal(core.StringPtr("someone@example.com")))
				Expect(editLoadBalancerPoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLoadBalancerPoolOptions successfully`, func() {
				// Construct an instance of the GetLoadBalancerPoolOptions model
				poolIdentifier := "testString"
				getLoadBalancerPoolOptionsModel := testService.NewGetLoadBalancerPoolOptions(poolIdentifier)
				getLoadBalancerPoolOptionsModel.SetPoolIdentifier("testString")
				getLoadBalancerPoolOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLoadBalancerPoolOptionsModel).ToNot(BeNil())
				Expect(getLoadBalancerPoolOptionsModel.PoolIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerPoolOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllLoadBalancerPoolsOptions successfully`, func() {
				// Construct an instance of the ListAllLoadBalancerPoolsOptions model
				listAllLoadBalancerPoolsOptionsModel := testService.NewListAllLoadBalancerPoolsOptions()
				listAllLoadBalancerPoolsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllLoadBalancerPoolsOptionsModel).ToNot(BeNil())
				Expect(listAllLoadBalancerPoolsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
