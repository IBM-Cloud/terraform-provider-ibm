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

package globalloadbalancereventsv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/globalloadbalancereventsv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`GlobalLoadBalancerEventsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
				URL: "https://globalloadbalancereventsv1/api",
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
			testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCER_EVENTS_URL":       "https://globalloadbalancereventsv1/api",
				"GLOBAL_LOAD_BALANCER_EVENTS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1UsingExternalConfig(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
					Crn: core.StringPtr(crn),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1UsingExternalConfig(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
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
				testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1UsingExternalConfig(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
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
				"GLOBAL_LOAD_BALANCER_EVENTS_URL":       "https://globalloadbalancereventsv1/api",
				"GLOBAL_LOAD_BALANCER_EVENTS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1UsingExternalConfig(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
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
				"GLOBAL_LOAD_BALANCER_EVENTS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1UsingExternalConfig(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
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
	Describe(`GetLoadBalancerEvents(getLoadBalancerEventsOptions *GetLoadBalancerEventsOptions) - Operation response error`, func() {
		crn := "testString"
		getLoadBalancerEventsPath := "/v1/testString/load_balancers/events"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getLoadBalancerEventsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLoadBalancerEvents with error: Operation response processing error`, func() {
				testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerEventsOptions model
				getLoadBalancerEventsOptionsModel := new(globalloadbalancereventsv1.GetLoadBalancerEventsOptions)
				getLoadBalancerEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetLoadBalancerEvents(getLoadBalancerEventsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLoadBalancerEvents(getLoadBalancerEventsOptions *GetLoadBalancerEventsOptions)`, func() {
		crn := "testString"
		getLoadBalancerEventsPath := "/v1/testString/load_balancers/events"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getLoadBalancerEventsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "result": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "timestamp": "2019-01-01T12:00:00", "pool": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "name": "some-pool", "healthy": true, "changed": true, "minimum_origins": 1}], "origins": [{"name": "f1aba936b94213e5b8dca0c0dbf1f9cc", "address": "1.2.3.4", "ip": "1.2.3.4", "enabled": true, "healthy": true, "failure_reason": "No failures", "changed": true}]}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetLoadBalancerEvents successfully`, func() {
				testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLoadBalancerEvents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerEventsOptions model
				getLoadBalancerEventsOptionsModel := new(globalloadbalancereventsv1.GetLoadBalancerEventsOptions)
				getLoadBalancerEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLoadBalancerEvents(getLoadBalancerEventsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetLoadBalancerEvents with error: Operation request error`, func() {
				testService, testServiceErr := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerEventsOptions model
				getLoadBalancerEventsOptionsModel := new(globalloadbalancereventsv1.GetLoadBalancerEventsOptions)
				getLoadBalancerEventsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetLoadBalancerEvents(getLoadBalancerEventsOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			crn := "testString"
			testService, _ := globalloadbalancereventsv1.NewGlobalLoadBalancerEventsV1(&globalloadbalancereventsv1.GlobalLoadBalancerEventsV1Options{
				URL:           "http://globalloadbalancereventsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
			})
			It(`Invoke NewGetLoadBalancerEventsOptions successfully`, func() {
				// Construct an instance of the GetLoadBalancerEventsOptions model
				getLoadBalancerEventsOptionsModel := testService.NewGetLoadBalancerEventsOptions()
				getLoadBalancerEventsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLoadBalancerEventsOptionsModel).ToNot(BeNil())
				Expect(getLoadBalancerEventsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
