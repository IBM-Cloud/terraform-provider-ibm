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

package zoneratelimitsv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/zoneratelimitsv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`ZoneRateLimitsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
				URL:            "https://zoneratelimitsv1/api",
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
			testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{})
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
				"ZONE_RATE_LIMITS_URL":       "https://zoneratelimitsv1/api",
				"ZONE_RATE_LIMITS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1UsingExternalConfig(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1UsingExternalConfig(&zoneratelimitsv1.ZoneRateLimitsV1Options{
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
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1UsingExternalConfig(&zoneratelimitsv1.ZoneRateLimitsV1Options{
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
				"ZONE_RATE_LIMITS_URL":       "https://zoneratelimitsv1/api",
				"ZONE_RATE_LIMITS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1UsingExternalConfig(&zoneratelimitsv1.ZoneRateLimitsV1Options{
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
				"ZONE_RATE_LIMITS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1UsingExternalConfig(&zoneratelimitsv1.ZoneRateLimitsV1Options{
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
	Describe(`ListAllZoneRateLimits(listAllZoneRateLimitsOptions *ListAllZoneRateLimitsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneRateLimitsPath := "/v1/testString/zones/testString/rate_limits"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllZoneRateLimitsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllZoneRateLimits with error: Operation response processing error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneRateLimitsOptions model
				listAllZoneRateLimitsOptionsModel := new(zoneratelimitsv1.ListAllZoneRateLimitsOptions)
				listAllZoneRateLimitsOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneRateLimitsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneRateLimitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListAllZoneRateLimits(listAllZoneRateLimitsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllZoneRateLimits(listAllZoneRateLimitsOptions *ListAllZoneRateLimitsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneRateLimitsPath := "/v1/testString/zones/testString/rate_limits"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllZoneRateLimitsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["[]"]], "messages": [["[]"]], "result": [{"id": "92f17202ed8bd63d69a66b86a49a8f6b", "disabled": false, "description": "Prevent multiple login failures to mitigate brute force attacks", "bypass": [{"name": "url", "value": "example.com/*"}], "threshold": 1000, "period": 60, "correlate": {"by": "nat"}, "action": {"mode": "simulate", "timeout": 60, "response": {"content_type": "text/plain", "body": "This request has been rate-limited."}}, "match": {"request": {"methods": ["_ALL_"], "schemes": ["_ALL_"], "url": "*.example.org/path*"}, "response": {"status": [403], "headers": [{"name": "Cf-Cache-Status", "op": "ne", "value": "HIT"}], "origin_traffic": false}}}], "result_info": {"page": 1, "per_page": 10, "count": 1, "total_count": 1}}`)
				}))
			})
			It(`Invoke ListAllZoneRateLimits successfully`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAllZoneRateLimits(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllZoneRateLimitsOptions model
				listAllZoneRateLimitsOptionsModel := new(zoneratelimitsv1.ListAllZoneRateLimitsOptions)
				listAllZoneRateLimitsOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneRateLimitsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneRateLimitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAllZoneRateLimits(listAllZoneRateLimitsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListAllZoneRateLimits with error: Operation request error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneRateLimitsOptions model
				listAllZoneRateLimitsOptionsModel := new(zoneratelimitsv1.ListAllZoneRateLimitsOptions)
				listAllZoneRateLimitsOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneRateLimitsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneRateLimitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListAllZoneRateLimits(listAllZoneRateLimitsOptionsModel)
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
	Describe(`CreateZoneRateLimits(createZoneRateLimitsOptions *CreateZoneRateLimitsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneRateLimitsPath := "/v1/testString/zones/testString/rate_limits"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createZoneRateLimitsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZoneRateLimits with error: Operation response processing error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RatelimitInputMatchResponseHeadersItem model
				ratelimitInputMatchResponseHeadersItemModel := new(zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem)
				ratelimitInputMatchResponseHeadersItemModel.Name = core.StringPtr("Cf-Cache-Status")
				ratelimitInputMatchResponseHeadersItemModel.Op = core.StringPtr("ne")
				ratelimitInputMatchResponseHeadersItemModel.Value = core.StringPtr("HIT")

				// Construct an instance of the RatelimitInputActionResponse model
				ratelimitInputActionResponseModel := new(zoneratelimitsv1.RatelimitInputActionResponse)
				ratelimitInputActionResponseModel.ContentType = core.StringPtr("text/plain")
				ratelimitInputActionResponseModel.Body = core.StringPtr("This request has been rate-limited.")

				// Construct an instance of the RatelimitInputMatchRequest model
				ratelimitInputMatchRequestModel := new(zoneratelimitsv1.RatelimitInputMatchRequest)
				ratelimitInputMatchRequestModel.Methods = []string{"GET"}
				ratelimitInputMatchRequestModel.Schemes = []string{"HTTP"}
				ratelimitInputMatchRequestModel.URL = core.StringPtr("*.example.org/path*")

				// Construct an instance of the RatelimitInputMatchResponse model
				ratelimitInputMatchResponseModel := new(zoneratelimitsv1.RatelimitInputMatchResponse)
				ratelimitInputMatchResponseModel.Status = []int64{int64(403)}
				ratelimitInputMatchResponseModel.HeadersVar = []zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}
				ratelimitInputMatchResponseModel.OriginTraffic = core.BoolPtr(false)

				// Construct an instance of the RatelimitInputAction model
				ratelimitInputActionModel := new(zoneratelimitsv1.RatelimitInputAction)
				ratelimitInputActionModel.Mode = core.StringPtr("simulate")
				ratelimitInputActionModel.Timeout = core.Int64Ptr(int64(60))
				ratelimitInputActionModel.Response = ratelimitInputActionResponseModel

				// Construct an instance of the RatelimitInputBypassItem model
				ratelimitInputBypassItemModel := new(zoneratelimitsv1.RatelimitInputBypassItem)
				ratelimitInputBypassItemModel.Name = core.StringPtr("url")
				ratelimitInputBypassItemModel.Value = core.StringPtr("api.example.com/*")

				// Construct an instance of the RatelimitInputCorrelate model
				ratelimitInputCorrelateModel := new(zoneratelimitsv1.RatelimitInputCorrelate)
				ratelimitInputCorrelateModel.By = core.StringPtr("nat")

				// Construct an instance of the RatelimitInputMatch model
				ratelimitInputMatchModel := new(zoneratelimitsv1.RatelimitInputMatch)
				ratelimitInputMatchModel.Request = ratelimitInputMatchRequestModel
				ratelimitInputMatchModel.Response = ratelimitInputMatchResponseModel

				// Construct an instance of the CreateZoneRateLimitsOptions model
				createZoneRateLimitsOptionsModel := new(zoneratelimitsv1.CreateZoneRateLimitsOptions)
				createZoneRateLimitsOptionsModel.Disabled = core.BoolPtr(false)
				createZoneRateLimitsOptionsModel.Description = core.StringPtr("Prevent multiple login failures to mitigate brute force attacks")
				createZoneRateLimitsOptionsModel.Bypass = []zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel}
				createZoneRateLimitsOptionsModel.Threshold = core.Int64Ptr(int64(1000))
				createZoneRateLimitsOptionsModel.Period = core.Int64Ptr(int64(60))
				createZoneRateLimitsOptionsModel.Action = ratelimitInputActionModel
				createZoneRateLimitsOptionsModel.Correlate = ratelimitInputCorrelateModel
				createZoneRateLimitsOptionsModel.Match = ratelimitInputMatchModel
				createZoneRateLimitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateZoneRateLimits(createZoneRateLimitsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateZoneRateLimits(createZoneRateLimitsOptions *CreateZoneRateLimitsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneRateLimitsPath := "/v1/testString/zones/testString/rate_limits"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createZoneRateLimitsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["[]"]], "messages": [["[]"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "disabled": false, "description": "Prevent multiple login failures to mitigate brute force attacks", "bypass": [{"name": "url", "value": "example.com/*"}], "threshold": 1000, "period": 60, "correlate": {"by": "nat"}, "action": {"mode": "simulate", "timeout": 60, "response": {"content_type": "text/plain", "body": "This request has been rate-limited."}}, "match": {"request": {"methods": ["_ALL_"], "schemes": ["_ALL_"], "url": "*.example.org/path*"}, "response": {"status": [403], "headers": [{"name": "Cf-Cache-Status", "op": "ne", "value": "HIT"}], "origin_traffic": false}}}}`)
				}))
			})
			It(`Invoke CreateZoneRateLimits successfully`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateZoneRateLimits(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RatelimitInputMatchResponseHeadersItem model
				ratelimitInputMatchResponseHeadersItemModel := new(zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem)
				ratelimitInputMatchResponseHeadersItemModel.Name = core.StringPtr("Cf-Cache-Status")
				ratelimitInputMatchResponseHeadersItemModel.Op = core.StringPtr("ne")
				ratelimitInputMatchResponseHeadersItemModel.Value = core.StringPtr("HIT")

				// Construct an instance of the RatelimitInputActionResponse model
				ratelimitInputActionResponseModel := new(zoneratelimitsv1.RatelimitInputActionResponse)
				ratelimitInputActionResponseModel.ContentType = core.StringPtr("text/plain")
				ratelimitInputActionResponseModel.Body = core.StringPtr("This request has been rate-limited.")

				// Construct an instance of the RatelimitInputMatchRequest model
				ratelimitInputMatchRequestModel := new(zoneratelimitsv1.RatelimitInputMatchRequest)
				ratelimitInputMatchRequestModel.Methods = []string{"GET"}
				ratelimitInputMatchRequestModel.Schemes = []string{"HTTP"}
				ratelimitInputMatchRequestModel.URL = core.StringPtr("*.example.org/path*")

				// Construct an instance of the RatelimitInputMatchResponse model
				ratelimitInputMatchResponseModel := new(zoneratelimitsv1.RatelimitInputMatchResponse)
				ratelimitInputMatchResponseModel.Status = []int64{int64(403)}
				ratelimitInputMatchResponseModel.HeadersVar = []zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}
				ratelimitInputMatchResponseModel.OriginTraffic = core.BoolPtr(false)

				// Construct an instance of the RatelimitInputAction model
				ratelimitInputActionModel := new(zoneratelimitsv1.RatelimitInputAction)
				ratelimitInputActionModel.Mode = core.StringPtr("simulate")
				ratelimitInputActionModel.Timeout = core.Int64Ptr(int64(60))
				ratelimitInputActionModel.Response = ratelimitInputActionResponseModel

				// Construct an instance of the RatelimitInputBypassItem model
				ratelimitInputBypassItemModel := new(zoneratelimitsv1.RatelimitInputBypassItem)
				ratelimitInputBypassItemModel.Name = core.StringPtr("url")
				ratelimitInputBypassItemModel.Value = core.StringPtr("api.example.com/*")

				// Construct an instance of the RatelimitInputCorrelate model
				ratelimitInputCorrelateModel := new(zoneratelimitsv1.RatelimitInputCorrelate)
				ratelimitInputCorrelateModel.By = core.StringPtr("nat")

				// Construct an instance of the RatelimitInputMatch model
				ratelimitInputMatchModel := new(zoneratelimitsv1.RatelimitInputMatch)
				ratelimitInputMatchModel.Request = ratelimitInputMatchRequestModel
				ratelimitInputMatchModel.Response = ratelimitInputMatchResponseModel

				// Construct an instance of the CreateZoneRateLimitsOptions model
				createZoneRateLimitsOptionsModel := new(zoneratelimitsv1.CreateZoneRateLimitsOptions)
				createZoneRateLimitsOptionsModel.Disabled = core.BoolPtr(false)
				createZoneRateLimitsOptionsModel.Description = core.StringPtr("Prevent multiple login failures to mitigate brute force attacks")
				createZoneRateLimitsOptionsModel.Bypass = []zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel}
				createZoneRateLimitsOptionsModel.Threshold = core.Int64Ptr(int64(1000))
				createZoneRateLimitsOptionsModel.Period = core.Int64Ptr(int64(60))
				createZoneRateLimitsOptionsModel.Action = ratelimitInputActionModel
				createZoneRateLimitsOptionsModel.Correlate = ratelimitInputCorrelateModel
				createZoneRateLimitsOptionsModel.Match = ratelimitInputMatchModel
				createZoneRateLimitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateZoneRateLimits(createZoneRateLimitsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateZoneRateLimits with error: Operation request error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RatelimitInputMatchResponseHeadersItem model
				ratelimitInputMatchResponseHeadersItemModel := new(zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem)
				ratelimitInputMatchResponseHeadersItemModel.Name = core.StringPtr("Cf-Cache-Status")
				ratelimitInputMatchResponseHeadersItemModel.Op = core.StringPtr("ne")
				ratelimitInputMatchResponseHeadersItemModel.Value = core.StringPtr("HIT")

				// Construct an instance of the RatelimitInputActionResponse model
				ratelimitInputActionResponseModel := new(zoneratelimitsv1.RatelimitInputActionResponse)
				ratelimitInputActionResponseModel.ContentType = core.StringPtr("text/plain")
				ratelimitInputActionResponseModel.Body = core.StringPtr("This request has been rate-limited.")

				// Construct an instance of the RatelimitInputMatchRequest model
				ratelimitInputMatchRequestModel := new(zoneratelimitsv1.RatelimitInputMatchRequest)
				ratelimitInputMatchRequestModel.Methods = []string{"GET"}
				ratelimitInputMatchRequestModel.Schemes = []string{"HTTP"}
				ratelimitInputMatchRequestModel.URL = core.StringPtr("*.example.org/path*")

				// Construct an instance of the RatelimitInputMatchResponse model
				ratelimitInputMatchResponseModel := new(zoneratelimitsv1.RatelimitInputMatchResponse)
				ratelimitInputMatchResponseModel.Status = []int64{int64(403)}
				ratelimitInputMatchResponseModel.HeadersVar = []zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}
				ratelimitInputMatchResponseModel.OriginTraffic = core.BoolPtr(false)

				// Construct an instance of the RatelimitInputAction model
				ratelimitInputActionModel := new(zoneratelimitsv1.RatelimitInputAction)
				ratelimitInputActionModel.Mode = core.StringPtr("simulate")
				ratelimitInputActionModel.Timeout = core.Int64Ptr(int64(60))
				ratelimitInputActionModel.Response = ratelimitInputActionResponseModel

				// Construct an instance of the RatelimitInputBypassItem model
				ratelimitInputBypassItemModel := new(zoneratelimitsv1.RatelimitInputBypassItem)
				ratelimitInputBypassItemModel.Name = core.StringPtr("url")
				ratelimitInputBypassItemModel.Value = core.StringPtr("api.example.com/*")

				// Construct an instance of the RatelimitInputCorrelate model
				ratelimitInputCorrelateModel := new(zoneratelimitsv1.RatelimitInputCorrelate)
				ratelimitInputCorrelateModel.By = core.StringPtr("nat")

				// Construct an instance of the RatelimitInputMatch model
				ratelimitInputMatchModel := new(zoneratelimitsv1.RatelimitInputMatch)
				ratelimitInputMatchModel.Request = ratelimitInputMatchRequestModel
				ratelimitInputMatchModel.Response = ratelimitInputMatchResponseModel

				// Construct an instance of the CreateZoneRateLimitsOptions model
				createZoneRateLimitsOptionsModel := new(zoneratelimitsv1.CreateZoneRateLimitsOptions)
				createZoneRateLimitsOptionsModel.Disabled = core.BoolPtr(false)
				createZoneRateLimitsOptionsModel.Description = core.StringPtr("Prevent multiple login failures to mitigate brute force attacks")
				createZoneRateLimitsOptionsModel.Bypass = []zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel}
				createZoneRateLimitsOptionsModel.Threshold = core.Int64Ptr(int64(1000))
				createZoneRateLimitsOptionsModel.Period = core.Int64Ptr(int64(60))
				createZoneRateLimitsOptionsModel.Action = ratelimitInputActionModel
				createZoneRateLimitsOptionsModel.Correlate = ratelimitInputCorrelateModel
				createZoneRateLimitsOptionsModel.Match = ratelimitInputMatchModel
				createZoneRateLimitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateZoneRateLimits(createZoneRateLimitsOptionsModel)
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
	Describe(`DeleteZoneRateLimit(deleteZoneRateLimitOptions *DeleteZoneRateLimitOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneRateLimitPath := "/v1/testString/zones/testString/rate_limits/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteZoneRateLimitPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZoneRateLimit with error: Operation response processing error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneRateLimitOptions model
				deleteZoneRateLimitOptionsModel := new(zoneratelimitsv1.DeleteZoneRateLimitOptions)
				deleteZoneRateLimitOptionsModel.RateLimitIdentifier = core.StringPtr("testString")
				deleteZoneRateLimitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteZoneRateLimit(deleteZoneRateLimitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteZoneRateLimit(deleteZoneRateLimitOptions *DeleteZoneRateLimitOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneRateLimitPath := "/v1/testString/zones/testString/rate_limits/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteZoneRateLimitPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["[]"]], "messages": [["[]"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteZoneRateLimit successfully`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteZoneRateLimit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneRateLimitOptions model
				deleteZoneRateLimitOptionsModel := new(zoneratelimitsv1.DeleteZoneRateLimitOptions)
				deleteZoneRateLimitOptionsModel.RateLimitIdentifier = core.StringPtr("testString")
				deleteZoneRateLimitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteZoneRateLimit(deleteZoneRateLimitOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteZoneRateLimit with error: Operation validation and request error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneRateLimitOptions model
				deleteZoneRateLimitOptionsModel := new(zoneratelimitsv1.DeleteZoneRateLimitOptions)
				deleteZoneRateLimitOptionsModel.RateLimitIdentifier = core.StringPtr("testString")
				deleteZoneRateLimitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteZoneRateLimit(deleteZoneRateLimitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneRateLimitOptions model with no property values
				deleteZoneRateLimitOptionsModelNew := new(zoneratelimitsv1.DeleteZoneRateLimitOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteZoneRateLimit(deleteZoneRateLimitOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRateLimit(getRateLimitOptions *GetRateLimitOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getRateLimitPath := "/v1/testString/zones/testString/rate_limits/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getRateLimitPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRateLimit with error: Operation response processing error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetRateLimitOptions model
				getRateLimitOptionsModel := new(zoneratelimitsv1.GetRateLimitOptions)
				getRateLimitOptionsModel.RateLimitIdentifier = core.StringPtr("testString")
				getRateLimitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetRateLimit(getRateLimitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRateLimit(getRateLimitOptions *GetRateLimitOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getRateLimitPath := "/v1/testString/zones/testString/rate_limits/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getRateLimitPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["[]"]], "messages": [["[]"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "disabled": false, "description": "Prevent multiple login failures to mitigate brute force attacks", "bypass": [{"name": "url", "value": "example.com/*"}], "threshold": 1000, "period": 60, "correlate": {"by": "nat"}, "action": {"mode": "simulate", "timeout": 60, "response": {"content_type": "text/plain", "body": "This request has been rate-limited."}}, "match": {"request": {"methods": ["_ALL_"], "schemes": ["_ALL_"], "url": "*.example.org/path*"}, "response": {"status": [403], "headers": [{"name": "Cf-Cache-Status", "op": "ne", "value": "HIT"}], "origin_traffic": false}}}}`)
				}))
			})
			It(`Invoke GetRateLimit successfully`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetRateLimit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRateLimitOptions model
				getRateLimitOptionsModel := new(zoneratelimitsv1.GetRateLimitOptions)
				getRateLimitOptionsModel.RateLimitIdentifier = core.StringPtr("testString")
				getRateLimitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetRateLimit(getRateLimitOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetRateLimit with error: Operation validation and request error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetRateLimitOptions model
				getRateLimitOptionsModel := new(zoneratelimitsv1.GetRateLimitOptions)
				getRateLimitOptionsModel.RateLimitIdentifier = core.StringPtr("testString")
				getRateLimitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetRateLimit(getRateLimitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRateLimitOptions model with no property values
				getRateLimitOptionsModelNew := new(zoneratelimitsv1.GetRateLimitOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetRateLimit(getRateLimitOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRateLimit(updateRateLimitOptions *UpdateRateLimitOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateRateLimitPath := "/v1/testString/zones/testString/rate_limits/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateRateLimitPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRateLimit with error: Operation response processing error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RatelimitInputMatchResponseHeadersItem model
				ratelimitInputMatchResponseHeadersItemModel := new(zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem)
				ratelimitInputMatchResponseHeadersItemModel.Name = core.StringPtr("Cf-Cache-Status")
				ratelimitInputMatchResponseHeadersItemModel.Op = core.StringPtr("ne")
				ratelimitInputMatchResponseHeadersItemModel.Value = core.StringPtr("HIT")

				// Construct an instance of the RatelimitInputActionResponse model
				ratelimitInputActionResponseModel := new(zoneratelimitsv1.RatelimitInputActionResponse)
				ratelimitInputActionResponseModel.ContentType = core.StringPtr("text/plain")
				ratelimitInputActionResponseModel.Body = core.StringPtr("This request has been rate-limited.")

				// Construct an instance of the RatelimitInputMatchRequest model
				ratelimitInputMatchRequestModel := new(zoneratelimitsv1.RatelimitInputMatchRequest)
				ratelimitInputMatchRequestModel.Methods = []string{"GET"}
				ratelimitInputMatchRequestModel.Schemes = []string{"HTTP"}
				ratelimitInputMatchRequestModel.URL = core.StringPtr("*.example.org/path*")

				// Construct an instance of the RatelimitInputMatchResponse model
				ratelimitInputMatchResponseModel := new(zoneratelimitsv1.RatelimitInputMatchResponse)
				ratelimitInputMatchResponseModel.Status = []int64{int64(403)}
				ratelimitInputMatchResponseModel.HeadersVar = []zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}
				ratelimitInputMatchResponseModel.OriginTraffic = core.BoolPtr(false)

				// Construct an instance of the RatelimitInputAction model
				ratelimitInputActionModel := new(zoneratelimitsv1.RatelimitInputAction)
				ratelimitInputActionModel.Mode = core.StringPtr("simulate")
				ratelimitInputActionModel.Timeout = core.Int64Ptr(int64(60))
				ratelimitInputActionModel.Response = ratelimitInputActionResponseModel

				// Construct an instance of the RatelimitInputBypassItem model
				ratelimitInputBypassItemModel := new(zoneratelimitsv1.RatelimitInputBypassItem)
				ratelimitInputBypassItemModel.Name = core.StringPtr("url")
				ratelimitInputBypassItemModel.Value = core.StringPtr("api.example.com/*")

				// Construct an instance of the RatelimitInputCorrelate model
				ratelimitInputCorrelateModel := new(zoneratelimitsv1.RatelimitInputCorrelate)
				ratelimitInputCorrelateModel.By = core.StringPtr("nat")

				// Construct an instance of the RatelimitInputMatch model
				ratelimitInputMatchModel := new(zoneratelimitsv1.RatelimitInputMatch)
				ratelimitInputMatchModel.Request = ratelimitInputMatchRequestModel
				ratelimitInputMatchModel.Response = ratelimitInputMatchResponseModel

				// Construct an instance of the UpdateRateLimitOptions model
				updateRateLimitOptionsModel := new(zoneratelimitsv1.UpdateRateLimitOptions)
				updateRateLimitOptionsModel.RateLimitIdentifier = core.StringPtr("testString")
				updateRateLimitOptionsModel.Disabled = core.BoolPtr(false)
				updateRateLimitOptionsModel.Description = core.StringPtr("Prevent multiple login failures to mitigate brute force attacks")
				updateRateLimitOptionsModel.Bypass = []zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel}
				updateRateLimitOptionsModel.Threshold = core.Int64Ptr(int64(1000))
				updateRateLimitOptionsModel.Period = core.Int64Ptr(int64(60))
				updateRateLimitOptionsModel.Action = ratelimitInputActionModel
				updateRateLimitOptionsModel.Correlate = ratelimitInputCorrelateModel
				updateRateLimitOptionsModel.Match = ratelimitInputMatchModel
				updateRateLimitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateRateLimit(updateRateLimitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateRateLimit(updateRateLimitOptions *UpdateRateLimitOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateRateLimitPath := "/v1/testString/zones/testString/rate_limits/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateRateLimitPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["[]"]], "messages": [["[]"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "disabled": false, "description": "Prevent multiple login failures to mitigate brute force attacks", "bypass": [{"name": "url", "value": "example.com/*"}], "threshold": 1000, "period": 60, "correlate": {"by": "nat"}, "action": {"mode": "simulate", "timeout": 60, "response": {"content_type": "text/plain", "body": "This request has been rate-limited."}}, "match": {"request": {"methods": ["_ALL_"], "schemes": ["_ALL_"], "url": "*.example.org/path*"}, "response": {"status": [403], "headers": [{"name": "Cf-Cache-Status", "op": "ne", "value": "HIT"}], "origin_traffic": false}}}}`)
				}))
			})
			It(`Invoke UpdateRateLimit successfully`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateRateLimit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RatelimitInputMatchResponseHeadersItem model
				ratelimitInputMatchResponseHeadersItemModel := new(zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem)
				ratelimitInputMatchResponseHeadersItemModel.Name = core.StringPtr("Cf-Cache-Status")
				ratelimitInputMatchResponseHeadersItemModel.Op = core.StringPtr("ne")
				ratelimitInputMatchResponseHeadersItemModel.Value = core.StringPtr("HIT")

				// Construct an instance of the RatelimitInputActionResponse model
				ratelimitInputActionResponseModel := new(zoneratelimitsv1.RatelimitInputActionResponse)
				ratelimitInputActionResponseModel.ContentType = core.StringPtr("text/plain")
				ratelimitInputActionResponseModel.Body = core.StringPtr("This request has been rate-limited.")

				// Construct an instance of the RatelimitInputMatchRequest model
				ratelimitInputMatchRequestModel := new(zoneratelimitsv1.RatelimitInputMatchRequest)
				ratelimitInputMatchRequestModel.Methods = []string{"GET"}
				ratelimitInputMatchRequestModel.Schemes = []string{"HTTP"}
				ratelimitInputMatchRequestModel.URL = core.StringPtr("*.example.org/path*")

				// Construct an instance of the RatelimitInputMatchResponse model
				ratelimitInputMatchResponseModel := new(zoneratelimitsv1.RatelimitInputMatchResponse)
				ratelimitInputMatchResponseModel.Status = []int64{int64(403)}
				ratelimitInputMatchResponseModel.HeadersVar = []zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}
				ratelimitInputMatchResponseModel.OriginTraffic = core.BoolPtr(false)

				// Construct an instance of the RatelimitInputAction model
				ratelimitInputActionModel := new(zoneratelimitsv1.RatelimitInputAction)
				ratelimitInputActionModel.Mode = core.StringPtr("simulate")
				ratelimitInputActionModel.Timeout = core.Int64Ptr(int64(60))
				ratelimitInputActionModel.Response = ratelimitInputActionResponseModel

				// Construct an instance of the RatelimitInputBypassItem model
				ratelimitInputBypassItemModel := new(zoneratelimitsv1.RatelimitInputBypassItem)
				ratelimitInputBypassItemModel.Name = core.StringPtr("url")
				ratelimitInputBypassItemModel.Value = core.StringPtr("api.example.com/*")

				// Construct an instance of the RatelimitInputCorrelate model
				ratelimitInputCorrelateModel := new(zoneratelimitsv1.RatelimitInputCorrelate)
				ratelimitInputCorrelateModel.By = core.StringPtr("nat")

				// Construct an instance of the RatelimitInputMatch model
				ratelimitInputMatchModel := new(zoneratelimitsv1.RatelimitInputMatch)
				ratelimitInputMatchModel.Request = ratelimitInputMatchRequestModel
				ratelimitInputMatchModel.Response = ratelimitInputMatchResponseModel

				// Construct an instance of the UpdateRateLimitOptions model
				updateRateLimitOptionsModel := new(zoneratelimitsv1.UpdateRateLimitOptions)
				updateRateLimitOptionsModel.RateLimitIdentifier = core.StringPtr("testString")
				updateRateLimitOptionsModel.Disabled = core.BoolPtr(false)
				updateRateLimitOptionsModel.Description = core.StringPtr("Prevent multiple login failures to mitigate brute force attacks")
				updateRateLimitOptionsModel.Bypass = []zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel}
				updateRateLimitOptionsModel.Threshold = core.Int64Ptr(int64(1000))
				updateRateLimitOptionsModel.Period = core.Int64Ptr(int64(60))
				updateRateLimitOptionsModel.Action = ratelimitInputActionModel
				updateRateLimitOptionsModel.Correlate = ratelimitInputCorrelateModel
				updateRateLimitOptionsModel.Match = ratelimitInputMatchModel
				updateRateLimitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateRateLimit(updateRateLimitOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateRateLimit with error: Operation validation and request error`, func() {
				testService, testServiceErr := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RatelimitInputMatchResponseHeadersItem model
				ratelimitInputMatchResponseHeadersItemModel := new(zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem)
				ratelimitInputMatchResponseHeadersItemModel.Name = core.StringPtr("Cf-Cache-Status")
				ratelimitInputMatchResponseHeadersItemModel.Op = core.StringPtr("ne")
				ratelimitInputMatchResponseHeadersItemModel.Value = core.StringPtr("HIT")

				// Construct an instance of the RatelimitInputActionResponse model
				ratelimitInputActionResponseModel := new(zoneratelimitsv1.RatelimitInputActionResponse)
				ratelimitInputActionResponseModel.ContentType = core.StringPtr("text/plain")
				ratelimitInputActionResponseModel.Body = core.StringPtr("This request has been rate-limited.")

				// Construct an instance of the RatelimitInputMatchRequest model
				ratelimitInputMatchRequestModel := new(zoneratelimitsv1.RatelimitInputMatchRequest)
				ratelimitInputMatchRequestModel.Methods = []string{"GET"}
				ratelimitInputMatchRequestModel.Schemes = []string{"HTTP"}
				ratelimitInputMatchRequestModel.URL = core.StringPtr("*.example.org/path*")

				// Construct an instance of the RatelimitInputMatchResponse model
				ratelimitInputMatchResponseModel := new(zoneratelimitsv1.RatelimitInputMatchResponse)
				ratelimitInputMatchResponseModel.Status = []int64{int64(403)}
				ratelimitInputMatchResponseModel.HeadersVar = []zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}
				ratelimitInputMatchResponseModel.OriginTraffic = core.BoolPtr(false)

				// Construct an instance of the RatelimitInputAction model
				ratelimitInputActionModel := new(zoneratelimitsv1.RatelimitInputAction)
				ratelimitInputActionModel.Mode = core.StringPtr("simulate")
				ratelimitInputActionModel.Timeout = core.Int64Ptr(int64(60))
				ratelimitInputActionModel.Response = ratelimitInputActionResponseModel

				// Construct an instance of the RatelimitInputBypassItem model
				ratelimitInputBypassItemModel := new(zoneratelimitsv1.RatelimitInputBypassItem)
				ratelimitInputBypassItemModel.Name = core.StringPtr("url")
				ratelimitInputBypassItemModel.Value = core.StringPtr("api.example.com/*")

				// Construct an instance of the RatelimitInputCorrelate model
				ratelimitInputCorrelateModel := new(zoneratelimitsv1.RatelimitInputCorrelate)
				ratelimitInputCorrelateModel.By = core.StringPtr("nat")

				// Construct an instance of the RatelimitInputMatch model
				ratelimitInputMatchModel := new(zoneratelimitsv1.RatelimitInputMatch)
				ratelimitInputMatchModel.Request = ratelimitInputMatchRequestModel
				ratelimitInputMatchModel.Response = ratelimitInputMatchResponseModel

				// Construct an instance of the UpdateRateLimitOptions model
				updateRateLimitOptionsModel := new(zoneratelimitsv1.UpdateRateLimitOptions)
				updateRateLimitOptionsModel.RateLimitIdentifier = core.StringPtr("testString")
				updateRateLimitOptionsModel.Disabled = core.BoolPtr(false)
				updateRateLimitOptionsModel.Description = core.StringPtr("Prevent multiple login failures to mitigate brute force attacks")
				updateRateLimitOptionsModel.Bypass = []zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel}
				updateRateLimitOptionsModel.Threshold = core.Int64Ptr(int64(1000))
				updateRateLimitOptionsModel.Period = core.Int64Ptr(int64(60))
				updateRateLimitOptionsModel.Action = ratelimitInputActionModel
				updateRateLimitOptionsModel.Correlate = ratelimitInputCorrelateModel
				updateRateLimitOptionsModel.Match = ratelimitInputMatchModel
				updateRateLimitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateRateLimit(updateRateLimitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRateLimitOptions model with no property values
				updateRateLimitOptionsModelNew := new(zoneratelimitsv1.UpdateRateLimitOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateRateLimit(updateRateLimitOptionsModelNew)
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
			testService, _ := zoneratelimitsv1.NewZoneRateLimitsV1(&zoneratelimitsv1.ZoneRateLimitsV1Options{
				URL:            "http://zoneratelimitsv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCreateZoneRateLimitsOptions successfully`, func() {
				// Construct an instance of the RatelimitInputMatchResponseHeadersItem model
				ratelimitInputMatchResponseHeadersItemModel := new(zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem)
				Expect(ratelimitInputMatchResponseHeadersItemModel).ToNot(BeNil())
				ratelimitInputMatchResponseHeadersItemModel.Name = core.StringPtr("Cf-Cache-Status")
				ratelimitInputMatchResponseHeadersItemModel.Op = core.StringPtr("ne")
				ratelimitInputMatchResponseHeadersItemModel.Value = core.StringPtr("HIT")
				Expect(ratelimitInputMatchResponseHeadersItemModel.Name).To(Equal(core.StringPtr("Cf-Cache-Status")))
				Expect(ratelimitInputMatchResponseHeadersItemModel.Op).To(Equal(core.StringPtr("ne")))
				Expect(ratelimitInputMatchResponseHeadersItemModel.Value).To(Equal(core.StringPtr("HIT")))

				// Construct an instance of the RatelimitInputActionResponse model
				ratelimitInputActionResponseModel := new(zoneratelimitsv1.RatelimitInputActionResponse)
				Expect(ratelimitInputActionResponseModel).ToNot(BeNil())
				ratelimitInputActionResponseModel.ContentType = core.StringPtr("text/plain")
				ratelimitInputActionResponseModel.Body = core.StringPtr("This request has been rate-limited.")
				Expect(ratelimitInputActionResponseModel.ContentType).To(Equal(core.StringPtr("text/plain")))
				Expect(ratelimitInputActionResponseModel.Body).To(Equal(core.StringPtr("This request has been rate-limited.")))

				// Construct an instance of the RatelimitInputMatchRequest model
				ratelimitInputMatchRequestModel := new(zoneratelimitsv1.RatelimitInputMatchRequest)
				Expect(ratelimitInputMatchRequestModel).ToNot(BeNil())
				ratelimitInputMatchRequestModel.Methods = []string{"GET"}
				ratelimitInputMatchRequestModel.Schemes = []string{"HTTP"}
				ratelimitInputMatchRequestModel.URL = core.StringPtr("*.example.org/path*")
				Expect(ratelimitInputMatchRequestModel.Methods).To(Equal([]string{"GET"}))
				Expect(ratelimitInputMatchRequestModel.Schemes).To(Equal([]string{"HTTP"}))
				Expect(ratelimitInputMatchRequestModel.URL).To(Equal(core.StringPtr("*.example.org/path*")))

				// Construct an instance of the RatelimitInputMatchResponse model
				ratelimitInputMatchResponseModel := new(zoneratelimitsv1.RatelimitInputMatchResponse)
				Expect(ratelimitInputMatchResponseModel).ToNot(BeNil())
				ratelimitInputMatchResponseModel.Status = []int64{int64(403)}
				ratelimitInputMatchResponseModel.HeadersVar = []zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}
				ratelimitInputMatchResponseModel.OriginTraffic = core.BoolPtr(false)
				Expect(ratelimitInputMatchResponseModel.Status).To(Equal([]int64{int64(403)}))
				Expect(ratelimitInputMatchResponseModel.HeadersVar).To(Equal([]zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}))
				Expect(ratelimitInputMatchResponseModel.OriginTraffic).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the RatelimitInputAction model
				ratelimitInputActionModel := new(zoneratelimitsv1.RatelimitInputAction)
				Expect(ratelimitInputActionModel).ToNot(BeNil())
				ratelimitInputActionModel.Mode = core.StringPtr("simulate")
				ratelimitInputActionModel.Timeout = core.Int64Ptr(int64(60))
				ratelimitInputActionModel.Response = ratelimitInputActionResponseModel
				Expect(ratelimitInputActionModel.Mode).To(Equal(core.StringPtr("simulate")))
				Expect(ratelimitInputActionModel.Timeout).To(Equal(core.Int64Ptr(int64(60))))
				Expect(ratelimitInputActionModel.Response).To(Equal(ratelimitInputActionResponseModel))

				// Construct an instance of the RatelimitInputBypassItem model
				ratelimitInputBypassItemModel := new(zoneratelimitsv1.RatelimitInputBypassItem)
				Expect(ratelimitInputBypassItemModel).ToNot(BeNil())
				ratelimitInputBypassItemModel.Name = core.StringPtr("url")
				ratelimitInputBypassItemModel.Value = core.StringPtr("api.example.com/*")
				Expect(ratelimitInputBypassItemModel.Name).To(Equal(core.StringPtr("url")))
				Expect(ratelimitInputBypassItemModel.Value).To(Equal(core.StringPtr("api.example.com/*")))

				// Construct an instance of the RatelimitInputCorrelate model
				ratelimitInputCorrelateModel := new(zoneratelimitsv1.RatelimitInputCorrelate)
				Expect(ratelimitInputCorrelateModel).ToNot(BeNil())
				ratelimitInputCorrelateModel.By = core.StringPtr("nat")
				Expect(ratelimitInputCorrelateModel.By).To(Equal(core.StringPtr("nat")))

				// Construct an instance of the RatelimitInputMatch model
				ratelimitInputMatchModel := new(zoneratelimitsv1.RatelimitInputMatch)
				Expect(ratelimitInputMatchModel).ToNot(BeNil())
				ratelimitInputMatchModel.Request = ratelimitInputMatchRequestModel
				ratelimitInputMatchModel.Response = ratelimitInputMatchResponseModel
				Expect(ratelimitInputMatchModel.Request).To(Equal(ratelimitInputMatchRequestModel))
				Expect(ratelimitInputMatchModel.Response).To(Equal(ratelimitInputMatchResponseModel))

				// Construct an instance of the CreateZoneRateLimitsOptions model
				createZoneRateLimitsOptionsModel := testService.NewCreateZoneRateLimitsOptions()
				createZoneRateLimitsOptionsModel.SetDisabled(false)
				createZoneRateLimitsOptionsModel.SetDescription("Prevent multiple login failures to mitigate brute force attacks")
				createZoneRateLimitsOptionsModel.SetBypass([]zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel})
				createZoneRateLimitsOptionsModel.SetThreshold(int64(1000))
				createZoneRateLimitsOptionsModel.SetPeriod(int64(60))
				createZoneRateLimitsOptionsModel.SetAction(ratelimitInputActionModel)
				createZoneRateLimitsOptionsModel.SetCorrelate(ratelimitInputCorrelateModel)
				createZoneRateLimitsOptionsModel.SetMatch(ratelimitInputMatchModel)
				createZoneRateLimitsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneRateLimitsOptionsModel).ToNot(BeNil())
				Expect(createZoneRateLimitsOptionsModel.Disabled).To(Equal(core.BoolPtr(false)))
				Expect(createZoneRateLimitsOptionsModel.Description).To(Equal(core.StringPtr("Prevent multiple login failures to mitigate brute force attacks")))
				Expect(createZoneRateLimitsOptionsModel.Bypass).To(Equal([]zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel}))
				Expect(createZoneRateLimitsOptionsModel.Threshold).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(createZoneRateLimitsOptionsModel.Period).To(Equal(core.Int64Ptr(int64(60))))
				Expect(createZoneRateLimitsOptionsModel.Action).To(Equal(ratelimitInputActionModel))
				Expect(createZoneRateLimitsOptionsModel.Correlate).To(Equal(ratelimitInputCorrelateModel))
				Expect(createZoneRateLimitsOptionsModel.Match).To(Equal(ratelimitInputMatchModel))
				Expect(createZoneRateLimitsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneRateLimitOptions successfully`, func() {
				// Construct an instance of the DeleteZoneRateLimitOptions model
				rateLimitIdentifier := "testString"
				deleteZoneRateLimitOptionsModel := testService.NewDeleteZoneRateLimitOptions(rateLimitIdentifier)
				deleteZoneRateLimitOptionsModel.SetRateLimitIdentifier("testString")
				deleteZoneRateLimitOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneRateLimitOptionsModel).ToNot(BeNil())
				Expect(deleteZoneRateLimitOptionsModel.RateLimitIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneRateLimitOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRateLimitOptions successfully`, func() {
				// Construct an instance of the GetRateLimitOptions model
				rateLimitIdentifier := "testString"
				getRateLimitOptionsModel := testService.NewGetRateLimitOptions(rateLimitIdentifier)
				getRateLimitOptionsModel.SetRateLimitIdentifier("testString")
				getRateLimitOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRateLimitOptionsModel).ToNot(BeNil())
				Expect(getRateLimitOptionsModel.RateLimitIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getRateLimitOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllZoneRateLimitsOptions successfully`, func() {
				// Construct an instance of the ListAllZoneRateLimitsOptions model
				listAllZoneRateLimitsOptionsModel := testService.NewListAllZoneRateLimitsOptions()
				listAllZoneRateLimitsOptionsModel.SetPage(int64(38))
				listAllZoneRateLimitsOptionsModel.SetPerPage(int64(5))
				listAllZoneRateLimitsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllZoneRateLimitsOptionsModel).ToNot(BeNil())
				Expect(listAllZoneRateLimitsOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAllZoneRateLimitsOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(5))))
				Expect(listAllZoneRateLimitsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRatelimitInputAction successfully`, func() {
				mode := "simulate"
				model, err := testService.NewRatelimitInputAction(mode)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRatelimitInputBypassItem successfully`, func() {
				name := "url"
				value := "api.example.com/*"
				model, err := testService.NewRatelimitInputBypassItem(name, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRatelimitInputCorrelate successfully`, func() {
				by := "nat"
				model, err := testService.NewRatelimitInputCorrelate(by)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRatelimitInputMatchRequest successfully`, func() {
				url := "*.example.org/path*"
				model, err := testService.NewRatelimitInputMatchRequest(url)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRatelimitInputMatchResponseHeadersItem successfully`, func() {
				name := "Cf-Cache-Status"
				op := "ne"
				value := "HIT"
				model, err := testService.NewRatelimitInputMatchResponseHeadersItem(name, op, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateRateLimitOptions successfully`, func() {
				// Construct an instance of the RatelimitInputMatchResponseHeadersItem model
				ratelimitInputMatchResponseHeadersItemModel := new(zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem)
				Expect(ratelimitInputMatchResponseHeadersItemModel).ToNot(BeNil())
				ratelimitInputMatchResponseHeadersItemModel.Name = core.StringPtr("Cf-Cache-Status")
				ratelimitInputMatchResponseHeadersItemModel.Op = core.StringPtr("ne")
				ratelimitInputMatchResponseHeadersItemModel.Value = core.StringPtr("HIT")
				Expect(ratelimitInputMatchResponseHeadersItemModel.Name).To(Equal(core.StringPtr("Cf-Cache-Status")))
				Expect(ratelimitInputMatchResponseHeadersItemModel.Op).To(Equal(core.StringPtr("ne")))
				Expect(ratelimitInputMatchResponseHeadersItemModel.Value).To(Equal(core.StringPtr("HIT")))

				// Construct an instance of the RatelimitInputActionResponse model
				ratelimitInputActionResponseModel := new(zoneratelimitsv1.RatelimitInputActionResponse)
				Expect(ratelimitInputActionResponseModel).ToNot(BeNil())
				ratelimitInputActionResponseModel.ContentType = core.StringPtr("text/plain")
				ratelimitInputActionResponseModel.Body = core.StringPtr("This request has been rate-limited.")
				Expect(ratelimitInputActionResponseModel.ContentType).To(Equal(core.StringPtr("text/plain")))
				Expect(ratelimitInputActionResponseModel.Body).To(Equal(core.StringPtr("This request has been rate-limited.")))

				// Construct an instance of the RatelimitInputMatchRequest model
				ratelimitInputMatchRequestModel := new(zoneratelimitsv1.RatelimitInputMatchRequest)
				Expect(ratelimitInputMatchRequestModel).ToNot(BeNil())
				ratelimitInputMatchRequestModel.Methods = []string{"GET"}
				ratelimitInputMatchRequestModel.Schemes = []string{"HTTP"}
				ratelimitInputMatchRequestModel.URL = core.StringPtr("*.example.org/path*")
				Expect(ratelimitInputMatchRequestModel.Methods).To(Equal([]string{"GET"}))
				Expect(ratelimitInputMatchRequestModel.Schemes).To(Equal([]string{"HTTP"}))
				Expect(ratelimitInputMatchRequestModel.URL).To(Equal(core.StringPtr("*.example.org/path*")))

				// Construct an instance of the RatelimitInputMatchResponse model
				ratelimitInputMatchResponseModel := new(zoneratelimitsv1.RatelimitInputMatchResponse)
				Expect(ratelimitInputMatchResponseModel).ToNot(BeNil())
				ratelimitInputMatchResponseModel.Status = []int64{int64(403)}
				ratelimitInputMatchResponseModel.HeadersVar = []zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}
				ratelimitInputMatchResponseModel.OriginTraffic = core.BoolPtr(false)
				Expect(ratelimitInputMatchResponseModel.Status).To(Equal([]int64{int64(403)}))
				Expect(ratelimitInputMatchResponseModel.HeadersVar).To(Equal([]zoneratelimitsv1.RatelimitInputMatchResponseHeadersItem{*ratelimitInputMatchResponseHeadersItemModel}))
				Expect(ratelimitInputMatchResponseModel.OriginTraffic).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the RatelimitInputAction model
				ratelimitInputActionModel := new(zoneratelimitsv1.RatelimitInputAction)
				Expect(ratelimitInputActionModel).ToNot(BeNil())
				ratelimitInputActionModel.Mode = core.StringPtr("simulate")
				ratelimitInputActionModel.Timeout = core.Int64Ptr(int64(60))
				ratelimitInputActionModel.Response = ratelimitInputActionResponseModel
				Expect(ratelimitInputActionModel.Mode).To(Equal(core.StringPtr("simulate")))
				Expect(ratelimitInputActionModel.Timeout).To(Equal(core.Int64Ptr(int64(60))))
				Expect(ratelimitInputActionModel.Response).To(Equal(ratelimitInputActionResponseModel))

				// Construct an instance of the RatelimitInputBypassItem model
				ratelimitInputBypassItemModel := new(zoneratelimitsv1.RatelimitInputBypassItem)
				Expect(ratelimitInputBypassItemModel).ToNot(BeNil())
				ratelimitInputBypassItemModel.Name = core.StringPtr("url")
				ratelimitInputBypassItemModel.Value = core.StringPtr("api.example.com/*")
				Expect(ratelimitInputBypassItemModel.Name).To(Equal(core.StringPtr("url")))
				Expect(ratelimitInputBypassItemModel.Value).To(Equal(core.StringPtr("api.example.com/*")))

				// Construct an instance of the RatelimitInputCorrelate model
				ratelimitInputCorrelateModel := new(zoneratelimitsv1.RatelimitInputCorrelate)
				Expect(ratelimitInputCorrelateModel).ToNot(BeNil())
				ratelimitInputCorrelateModel.By = core.StringPtr("nat")
				Expect(ratelimitInputCorrelateModel.By).To(Equal(core.StringPtr("nat")))

				// Construct an instance of the RatelimitInputMatch model
				ratelimitInputMatchModel := new(zoneratelimitsv1.RatelimitInputMatch)
				Expect(ratelimitInputMatchModel).ToNot(BeNil())
				ratelimitInputMatchModel.Request = ratelimitInputMatchRequestModel
				ratelimitInputMatchModel.Response = ratelimitInputMatchResponseModel
				Expect(ratelimitInputMatchModel.Request).To(Equal(ratelimitInputMatchRequestModel))
				Expect(ratelimitInputMatchModel.Response).To(Equal(ratelimitInputMatchResponseModel))

				// Construct an instance of the UpdateRateLimitOptions model
				rateLimitIdentifier := "testString"
				updateRateLimitOptionsModel := testService.NewUpdateRateLimitOptions(rateLimitIdentifier)
				updateRateLimitOptionsModel.SetRateLimitIdentifier("testString")
				updateRateLimitOptionsModel.SetDisabled(false)
				updateRateLimitOptionsModel.SetDescription("Prevent multiple login failures to mitigate brute force attacks")
				updateRateLimitOptionsModel.SetBypass([]zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel})
				updateRateLimitOptionsModel.SetThreshold(int64(1000))
				updateRateLimitOptionsModel.SetPeriod(int64(60))
				updateRateLimitOptionsModel.SetAction(ratelimitInputActionModel)
				updateRateLimitOptionsModel.SetCorrelate(ratelimitInputCorrelateModel)
				updateRateLimitOptionsModel.SetMatch(ratelimitInputMatchModel)
				updateRateLimitOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRateLimitOptionsModel).ToNot(BeNil())
				Expect(updateRateLimitOptionsModel.RateLimitIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateRateLimitOptionsModel.Disabled).To(Equal(core.BoolPtr(false)))
				Expect(updateRateLimitOptionsModel.Description).To(Equal(core.StringPtr("Prevent multiple login failures to mitigate brute force attacks")))
				Expect(updateRateLimitOptionsModel.Bypass).To(Equal([]zoneratelimitsv1.RatelimitInputBypassItem{*ratelimitInputBypassItemModel}))
				Expect(updateRateLimitOptionsModel.Threshold).To(Equal(core.Int64Ptr(int64(1000))))
				Expect(updateRateLimitOptionsModel.Period).To(Equal(core.Int64Ptr(int64(60))))
				Expect(updateRateLimitOptionsModel.Action).To(Equal(ratelimitInputActionModel))
				Expect(updateRateLimitOptionsModel.Correlate).To(Equal(ratelimitInputCorrelateModel))
				Expect(updateRateLimitOptionsModel.Match).To(Equal(ratelimitInputMatchModel))
				Expect(updateRateLimitOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
