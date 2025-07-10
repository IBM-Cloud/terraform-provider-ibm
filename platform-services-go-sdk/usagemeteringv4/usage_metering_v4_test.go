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

package usagemeteringv4_test

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
	"github.com/IBM/platform-services-go-sdk/usagemeteringv4"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`UsageMeteringV4`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(usageMeteringService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(usageMeteringService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
				URL: "https://usagemeteringv4/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(usageMeteringService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USAGE_METERING_URL":       "https://usagemeteringv4/api",
				"USAGE_METERING_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4UsingExternalConfig(&usagemeteringv4.UsageMeteringV4Options{})
				Expect(usageMeteringService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := usageMeteringService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != usageMeteringService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(usageMeteringService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(usageMeteringService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4UsingExternalConfig(&usagemeteringv4.UsageMeteringV4Options{
					URL: "https://testService/api",
				})
				Expect(usageMeteringService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(usageMeteringService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := usageMeteringService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != usageMeteringService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(usageMeteringService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(usageMeteringService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4UsingExternalConfig(&usagemeteringv4.UsageMeteringV4Options{})
				err := usageMeteringService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(usageMeteringService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(usageMeteringService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := usageMeteringService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != usageMeteringService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(usageMeteringService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(usageMeteringService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USAGE_METERING_URL":       "https://usagemeteringv4/api",
				"USAGE_METERING_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4UsingExternalConfig(&usagemeteringv4.UsageMeteringV4Options{})

			It(`Instantiate service client with error`, func() {
				Expect(usageMeteringService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USAGE_METERING_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4UsingExternalConfig(&usagemeteringv4.UsageMeteringV4Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(usageMeteringService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = usagemeteringv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ReportResourceUsage(reportResourceUsageOptions *ReportResourceUsageOptions) - Operation response error`, func() {
		reportResourceUsagePath := "/v4/metering/resources/testString/usage"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(reportResourceUsagePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReportResourceUsage with error: Operation response processing error`, func() {
				usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageMeteringService).ToNot(BeNil())

				// Construct an instance of the MeasureAndQuantity model
				measureAndQuantityModel := new(usagemeteringv4.MeasureAndQuantity)
				measureAndQuantityModel.Measure = core.StringPtr("STORAGE")
				measureAndQuantityModel.Quantity = core.StringPtr("1")

				// Construct an instance of the ResourceInstanceUsage model
				resourceInstanceUsageModel := new(usagemeteringv4.ResourceInstanceUsage)
				resourceInstanceUsageModel.ResourceInstanceID = core.StringPtr("crn:v1:bluemix:staging:database-service:us-south:a/1c8ae972c35e470d994b6faff9494ce1:793ff3d3-9fe3-4329-9ea0-404703a3c371::")
				resourceInstanceUsageModel.PlanID = core.StringPtr("database-lite")
				resourceInstanceUsageModel.Region = core.StringPtr("us-south")
				resourceInstanceUsageModel.Start = core.Int64Ptr(int64(1485907200000))
				resourceInstanceUsageModel.End = core.Int64Ptr(int64(1485907200000))
				resourceInstanceUsageModel.MeasuredUsage = []usagemeteringv4.MeasureAndQuantity{*measureAndQuantityModel}
				resourceInstanceUsageModel.ConsumerID = core.StringPtr("cf-application:ed20abbe-8870-44e6-90f7-56d764c21127")

				// Construct an instance of the ReportResourceUsageOptions model
				reportResourceUsageOptionsModel := new(usagemeteringv4.ReportResourceUsageOptions)
				reportResourceUsageOptionsModel.ResourceID = core.StringPtr("testString")
				reportResourceUsageOptionsModel.ResourceUsage = []usagemeteringv4.ResourceInstanceUsage{*resourceInstanceUsageModel}
				reportResourceUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := usageMeteringService.ReportResourceUsage(reportResourceUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				usageMeteringService.EnableRetries(0, 0)
				result, response, operationErr = usageMeteringService.ReportResourceUsage(reportResourceUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReportResourceUsage(reportResourceUsageOptions *ReportResourceUsageOptions)`, func() {
		reportResourceUsagePath := "/v4/metering/resources/testString/usage"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(reportResourceUsagePath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"resources": [{"status": 6, "location": "Location", "code": "Code", "message": "Message"}]}`)
				}))
			})
			It(`Invoke ReportResourceUsage successfully`, func() {
				usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageMeteringService).ToNot(BeNil())
				usageMeteringService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := usageMeteringService.ReportResourceUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MeasureAndQuantity model
				measureAndQuantityModel := new(usagemeteringv4.MeasureAndQuantity)
				measureAndQuantityModel.Measure = core.StringPtr("STORAGE")
				measureAndQuantityModel.Quantity = core.StringPtr("1")

				// Construct an instance of the ResourceInstanceUsage model
				resourceInstanceUsageModel := new(usagemeteringv4.ResourceInstanceUsage)
				resourceInstanceUsageModel.ResourceInstanceID = core.StringPtr("crn:v1:bluemix:staging:database-service:us-south:a/1c8ae972c35e470d994b6faff9494ce1:793ff3d3-9fe3-4329-9ea0-404703a3c371::")
				resourceInstanceUsageModel.PlanID = core.StringPtr("database-lite")
				resourceInstanceUsageModel.Region = core.StringPtr("us-south")
				resourceInstanceUsageModel.Start = core.Int64Ptr(int64(1485907200000))
				resourceInstanceUsageModel.End = core.Int64Ptr(int64(1485907200000))
				resourceInstanceUsageModel.MeasuredUsage = []usagemeteringv4.MeasureAndQuantity{*measureAndQuantityModel}
				resourceInstanceUsageModel.ConsumerID = core.StringPtr("cf-application:ed20abbe-8870-44e6-90f7-56d764c21127")

				// Construct an instance of the ReportResourceUsageOptions model
				reportResourceUsageOptionsModel := new(usagemeteringv4.ReportResourceUsageOptions)
				reportResourceUsageOptionsModel.ResourceID = core.StringPtr("testString")
				reportResourceUsageOptionsModel.ResourceUsage = []usagemeteringv4.ResourceInstanceUsage{*resourceInstanceUsageModel}
				reportResourceUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = usageMeteringService.ReportResourceUsage(reportResourceUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = usageMeteringService.ReportResourceUsageWithContext(ctx, reportResourceUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				usageMeteringService.DisableRetries()
				result, response, operationErr = usageMeteringService.ReportResourceUsage(reportResourceUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = usageMeteringService.ReportResourceUsageWithContext(ctx, reportResourceUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReportResourceUsage with error: Operation validation and request error`, func() {
				usageMeteringService, serviceErr := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageMeteringService).ToNot(BeNil())

				// Construct an instance of the MeasureAndQuantity model
				measureAndQuantityModel := new(usagemeteringv4.MeasureAndQuantity)
				measureAndQuantityModel.Measure = core.StringPtr("STORAGE")
				measureAndQuantityModel.Quantity = core.StringPtr("1")

				// Construct an instance of the ResourceInstanceUsage model
				resourceInstanceUsageModel := new(usagemeteringv4.ResourceInstanceUsage)
				resourceInstanceUsageModel.ResourceInstanceID = core.StringPtr("crn:v1:bluemix:staging:database-service:us-south:a/1c8ae972c35e470d994b6faff9494ce1:793ff3d3-9fe3-4329-9ea0-404703a3c371::")
				resourceInstanceUsageModel.PlanID = core.StringPtr("database-lite")
				resourceInstanceUsageModel.Region = core.StringPtr("us-south")
				resourceInstanceUsageModel.Start = core.Int64Ptr(int64(1485907200000))
				resourceInstanceUsageModel.End = core.Int64Ptr(int64(1485907200000))
				resourceInstanceUsageModel.MeasuredUsage = []usagemeteringv4.MeasureAndQuantity{*measureAndQuantityModel}
				resourceInstanceUsageModel.ConsumerID = core.StringPtr("cf-application:ed20abbe-8870-44e6-90f7-56d764c21127")

				// Construct an instance of the ReportResourceUsageOptions model
				reportResourceUsageOptionsModel := new(usagemeteringv4.ReportResourceUsageOptions)
				reportResourceUsageOptionsModel.ResourceID = core.StringPtr("testString")
				reportResourceUsageOptionsModel.ResourceUsage = []usagemeteringv4.ResourceInstanceUsage{*resourceInstanceUsageModel}
				reportResourceUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := usageMeteringService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := usageMeteringService.ReportResourceUsage(reportResourceUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReportResourceUsageOptions model with no property values
				reportResourceUsageOptionsModelNew := new(usagemeteringv4.ReportResourceUsageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = usageMeteringService.ReportResourceUsage(reportResourceUsageOptionsModelNew)
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
			usageMeteringService, _ := usagemeteringv4.NewUsageMeteringV4(&usagemeteringv4.UsageMeteringV4Options{
				URL:           "http://usagemeteringv4modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewReportResourceUsageOptions successfully`, func() {
				// Construct an instance of the MeasureAndQuantity model
				measureAndQuantityModel := new(usagemeteringv4.MeasureAndQuantity)
				Expect(measureAndQuantityModel).ToNot(BeNil())
				measureAndQuantityModel.Measure = core.StringPtr("QUERIES")
				measureAndQuantityModel.Quantity = core.StringPtr("100")
				Expect(measureAndQuantityModel.Measure).To(Equal(core.StringPtr("QUERIES")))
				Expect(measureAndQuantityModel.Quantity).To(Equal(core.StringPtr("100")))

				// Construct an instance of the ResourceInstanceUsage model
				resourceInstanceUsageModel := new(usagemeteringv4.ResourceInstanceUsage)
				Expect(resourceInstanceUsageModel).ToNot(BeNil())
				resourceInstanceUsageModel.ResourceInstanceID = core.StringPtr("crn:v1:bluemix:staging:database-service:us-south:a/1c8ae972c35e470d994b6faff9494ce1:793ff3d3-9fe3-4329-9ea0-404703a3c371::")
				resourceInstanceUsageModel.PlanID = core.StringPtr("database-lite")
				resourceInstanceUsageModel.Region = core.StringPtr("us-south")
				resourceInstanceUsageModel.Start = core.Int64Ptr(int64(1485907200001))
				resourceInstanceUsageModel.End = core.Int64Ptr(int64(1485910800000))
				resourceInstanceUsageModel.MeasuredUsage = []usagemeteringv4.MeasureAndQuantity{*measureAndQuantityModel}
				resourceInstanceUsageModel.ConsumerID = core.StringPtr("cf-application:ed20abbe-8870-44e6-90f7-56d764c21127")
				Expect(resourceInstanceUsageModel.ResourceInstanceID).To(Equal(core.StringPtr("crn:v1:bluemix:staging:database-service:us-south:a/1c8ae972c35e470d994b6faff9494ce1:793ff3d3-9fe3-4329-9ea0-404703a3c371::")))
				Expect(resourceInstanceUsageModel.PlanID).To(Equal(core.StringPtr("database-lite")))
				Expect(resourceInstanceUsageModel.Region).To(Equal(core.StringPtr("us-south")))
				Expect(resourceInstanceUsageModel.Start).To(Equal(core.Int64Ptr(int64(1485907200001))))
				Expect(resourceInstanceUsageModel.End).To(Equal(core.Int64Ptr(int64(1485910800000))))
				Expect(resourceInstanceUsageModel.MeasuredUsage).To(Equal([]usagemeteringv4.MeasureAndQuantity{*measureAndQuantityModel}))
				Expect(resourceInstanceUsageModel.ConsumerID).To(Equal(core.StringPtr("cf-application:ed20abbe-8870-44e6-90f7-56d764c21127")))

				// Construct an instance of the ReportResourceUsageOptions model
				resourceID := "testString"
				resourceUsage := []usagemeteringv4.ResourceInstanceUsage{}
				reportResourceUsageOptionsModel := usageMeteringService.NewReportResourceUsageOptions(resourceID, resourceUsage)
				reportResourceUsageOptionsModel.SetResourceID("testString")
				reportResourceUsageOptionsModel.SetResourceUsage([]usagemeteringv4.ResourceInstanceUsage{*resourceInstanceUsageModel})
				reportResourceUsageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(reportResourceUsageOptionsModel).ToNot(BeNil())
				Expect(reportResourceUsageOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(reportResourceUsageOptionsModel.ResourceUsage).To(Equal([]usagemeteringv4.ResourceInstanceUsage{*resourceInstanceUsageModel}))
				Expect(reportResourceUsageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMeasureAndQuantity successfully`, func() {
				measure := "STORAGE"
				quantity := core.StringPtr("1")
				model, err := usageMeteringService.NewMeasureAndQuantity(measure, quantity)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceInstanceUsage successfully`, func() {
				resourceInstanceID := "crn:v1:bluemix:staging:database-service:us-south:a/1c8ae972c35e470d994b6faff9494ce1:793ff3d3-9fe3-4329-9ea0-404703a3c371::"
				planID := "database-lite"
				start := int64(1485907200000)
				end := int64(1485907200000)
				measuredUsage := []usagemeteringv4.MeasureAndQuantity{}
				model, err := usageMeteringService.NewResourceInstanceUsage(resourceInstanceID, planID, start, end, measuredUsage)
				Expect(model).ToNot(BeNil())
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
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
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
