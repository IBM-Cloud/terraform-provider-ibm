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

package resourcecontrollerv2_test

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
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ResourceControllerV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(resourceControllerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(resourceControllerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "https://resourcecontrollerv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(resourceControllerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_CONTROLLER_URL":       "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{})
				Expect(resourceControllerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := resourceControllerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != resourceControllerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(resourceControllerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(resourceControllerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
					URL: "https://testService/api",
				})
				Expect(resourceControllerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := resourceControllerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != resourceControllerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(resourceControllerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(resourceControllerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{})
				err := resourceControllerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := resourceControllerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != resourceControllerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(resourceControllerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(resourceControllerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_CONTROLLER_URL":       "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{})

			It(`Instantiate service client with error`, func() {
				Expect(resourceControllerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_CONTROLLER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(resourceControllerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = resourcecontrollerv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions) - Operation response error`, func() {
		listResourceInstancesPath := "/v2/resource_instances"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceInstancesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sub_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceInstances with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := new(resourcecontrollerv2.ListResourceInstancesOptions)
				listResourceInstancesOptionsModel.GUID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Name = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourcePlanID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Type = core.StringPtr("testString")
				listResourceInstancesOptionsModel.SubType = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceInstancesOptionsModel.Start = core.StringPtr("testString")
				listResourceInstancesOptionsModel.State = core.StringPtr("active")
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.ListResourceInstances(listResourceInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.ListResourceInstances(listResourceInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions)`, func() {
		listResourceInstancesPath := "/v2/resource_instances"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sub_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}]}`)
				}))
			})
			It(`Invoke ListResourceInstances successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := new(resourcecontrollerv2.ListResourceInstancesOptions)
				listResourceInstancesOptionsModel.GUID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Name = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourcePlanID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Type = core.StringPtr("testString")
				listResourceInstancesOptionsModel.SubType = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceInstancesOptionsModel.Start = core.StringPtr("testString")
				listResourceInstancesOptionsModel.State = core.StringPtr("active")
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.ListResourceInstancesWithContext(ctx, listResourceInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.ListResourceInstances(listResourceInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.ListResourceInstancesWithContext(ctx, listResourceInstancesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listResourceInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sub_type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}]}`)
				}))
			})
			It(`Invoke ListResourceInstances successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.ListResourceInstances(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := new(resourcecontrollerv2.ListResourceInstancesOptions)
				listResourceInstancesOptionsModel.GUID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Name = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourcePlanID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Type = core.StringPtr("testString")
				listResourceInstancesOptionsModel.SubType = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceInstancesOptionsModel.Start = core.StringPtr("testString")
				listResourceInstancesOptionsModel.State = core.StringPtr("active")
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.ListResourceInstances(listResourceInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListResourceInstances with error: Operation request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := new(resourcecontrollerv2.ListResourceInstancesOptions)
				listResourceInstancesOptionsModel.GUID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Name = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourcePlanID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Type = core.StringPtr("testString")
				listResourceInstancesOptionsModel.SubType = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceInstancesOptionsModel.Start = core.StringPtr("testString")
				listResourceInstancesOptionsModel.State = core.StringPtr("active")
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.ListResourceInstances(listResourceInstancesOptionsModel)
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
			It(`Invoke ListResourceInstances successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := new(resourcecontrollerv2.ListResourceInstancesOptions)
				listResourceInstancesOptionsModel.GUID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Name = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourcePlanID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Type = core.StringPtr("testString")
				listResourceInstancesOptionsModel.SubType = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceInstancesOptionsModel.Start = core.StringPtr("testString")
				listResourceInstancesOptionsModel.State = core.StringPtr("active")
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.ListResourceInstances(listResourceInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resourcecontrollerv2.ResourceInstancesList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(resourcecontrollerv2.ResourceInstancesList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(resourcecontrollerv2.ResourceInstancesList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","scheduled_reclaim_at":"2019-01-01T12:00:00.000Z","restored_at":"2019-01-01T12:00:00.000Z","restored_by":"RestoredBy","scheduled_reclaim_by":"ScheduledReclaimBy","name":"Name","region_id":"RegionID","account_id":"AccountID","reseller_channel_id":"ResellerChannelID","resource_plan_id":"ResourcePlanID","resource_group_id":"ResourceGroupID","resource_group_crn":"ResourceGroupCRN","target_crn":"TargetCRN","onetime_credentials":true,"parameters":{"anyKey":"anyValue"},"allow_cleanup":true,"crn":"CRN","state":"active","type":"Type","sub_type":"SubType","resource_id":"ResourceID","dashboard_url":"DashboardURL","last_operation":{"type":"Type","state":"in progress","sub_type":"SubType","async":false,"description":"Description","reason_code":"ReasonCode","poll_after":9,"cancelable":true,"poll":true},"resource_aliases_url":"ResourceAliasesURL","resource_bindings_url":"ResourceBindingsURL","resource_keys_url":"ResourceKeysURL","plan_history":[{"resource_plan_id":"ResourcePlanID","start_date":"2019-01-01T12:00:00.000Z","requestor_id":"RequestorID"}],"migrated":true,"extensions":{"anyKey":"anyValue"},"controlled_by":"ControlledBy","locked":true}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","scheduled_reclaim_at":"2019-01-01T12:00:00.000Z","restored_at":"2019-01-01T12:00:00.000Z","restored_by":"RestoredBy","scheduled_reclaim_by":"ScheduledReclaimBy","name":"Name","region_id":"RegionID","account_id":"AccountID","reseller_channel_id":"ResellerChannelID","resource_plan_id":"ResourcePlanID","resource_group_id":"ResourceGroupID","resource_group_crn":"ResourceGroupCRN","target_crn":"TargetCRN","onetime_credentials":true,"parameters":{"anyKey":"anyValue"},"allow_cleanup":true,"crn":"CRN","state":"active","type":"Type","sub_type":"SubType","resource_id":"ResourceID","dashboard_url":"DashboardURL","last_operation":{"type":"Type","state":"in progress","sub_type":"SubType","async":false,"description":"Description","reason_code":"ReasonCode","poll_after":9,"cancelable":true,"poll":true},"resource_aliases_url":"ResourceAliasesURL","resource_bindings_url":"ResourceBindingsURL","resource_keys_url":"ResourceKeysURL","plan_history":[{"resource_plan_id":"ResourcePlanID","start_date":"2019-01-01T12:00:00.000Z","requestor_id":"RequestorID"}],"migrated":true,"extensions":{"anyKey":"anyValue"},"controlled_by":"ControlledBy","locked":true}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ResourceInstancesPager.GetNext successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceInstancesOptionsModel := &resourcecontrollerv2.ListResourceInstancesOptions{
					GUID:            core.StringPtr("testString"),
					Name:            core.StringPtr("testString"),
					ResourceGroupID: core.StringPtr("testString"),
					ResourceID:      core.StringPtr("testString"),
					ResourcePlanID:  core.StringPtr("testString"),
					Type:            core.StringPtr("testString"),
					SubType:         core.StringPtr("testString"),
					Limit:           core.Int64Ptr(int64(10)),
					State:           core.StringPtr("active"),
					UpdatedFrom:     core.StringPtr("2021-01-01"),
					UpdatedTo:       core.StringPtr("2021-01-01"),
				}

				pager, err := resourceControllerService.NewResourceInstancesPager(listResourceInstancesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceInstance
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ResourceInstancesPager.GetAll successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceInstancesOptionsModel := &resourcecontrollerv2.ListResourceInstancesOptions{
					GUID:            core.StringPtr("testString"),
					Name:            core.StringPtr("testString"),
					ResourceGroupID: core.StringPtr("testString"),
					ResourceID:      core.StringPtr("testString"),
					ResourcePlanID:  core.StringPtr("testString"),
					Type:            core.StringPtr("testString"),
					SubType:         core.StringPtr("testString"),
					Limit:           core.Int64Ptr(int64(10)),
					State:           core.StringPtr("active"),
					UpdatedFrom:     core.StringPtr("2021-01-01"),
					UpdatedTo:       core.StringPtr("2021-01-01"),
				}

				pager, err := resourceControllerService.NewResourceInstancesPager(listResourceInstancesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateResourceInstance(createResourceInstanceOptions *CreateResourceInstanceOptions) - Operation response error`, func() {
		createResourceInstancePath := "/v2/resource_instances"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceInstancePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", false)))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceInstance with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsModel := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				createResourceInstanceOptionsModel.Name = core.StringPtr("ExampleResourceInstance")
				createResourceInstanceOptionsModel.Target = core.StringPtr("global")
				createResourceInstanceOptionsModel.ResourceGroup = core.StringPtr("13aa3ee48c3b44ddb64c05c79f7ab8ef")
				createResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a10e4960-3685-11e9-b210-d663bd873d93")
				createResourceInstanceOptionsModel.Tags = []string{"testString"}
				createResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createResourceInstanceOptionsModel.EntityLock = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResourceInstance(createResourceInstanceOptions *CreateResourceInstanceOptions)`, func() {
		createResourceInstancePath := "/v2/resource_instances"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceInstancePath))
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

					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", false)))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke CreateResourceInstance successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsModel := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				createResourceInstanceOptionsModel.Name = core.StringPtr("ExampleResourceInstance")
				createResourceInstanceOptionsModel.Target = core.StringPtr("global")
				createResourceInstanceOptionsModel.ResourceGroup = core.StringPtr("13aa3ee48c3b44ddb64c05c79f7ab8ef")
				createResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a10e4960-3685-11e9-b210-d663bd873d93")
				createResourceInstanceOptionsModel.Tags = []string{"testString"}
				createResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createResourceInstanceOptionsModel.EntityLock = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.CreateResourceInstanceWithContext(ctx, createResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.CreateResourceInstanceWithContext(ctx, createResourceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createResourceInstancePath))
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

					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", false)))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke CreateResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.CreateResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsModel := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				createResourceInstanceOptionsModel.Name = core.StringPtr("ExampleResourceInstance")
				createResourceInstanceOptionsModel.Target = core.StringPtr("global")
				createResourceInstanceOptionsModel.ResourceGroup = core.StringPtr("13aa3ee48c3b44ddb64c05c79f7ab8ef")
				createResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a10e4960-3685-11e9-b210-d663bd873d93")
				createResourceInstanceOptionsModel.Tags = []string{"testString"}
				createResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createResourceInstanceOptionsModel.EntityLock = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateResourceInstance with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsModel := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				createResourceInstanceOptionsModel.Name = core.StringPtr("ExampleResourceInstance")
				createResourceInstanceOptionsModel.Target = core.StringPtr("global")
				createResourceInstanceOptionsModel.ResourceGroup = core.StringPtr("13aa3ee48c3b44ddb64c05c79f7ab8ef")
				createResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a10e4960-3685-11e9-b210-d663bd873d93")
				createResourceInstanceOptionsModel.Tags = []string{"testString"}
				createResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createResourceInstanceOptionsModel.EntityLock = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceInstanceOptions model with no property values
				createResourceInstanceOptionsModelNew := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.CreateResourceInstance(createResourceInstanceOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsModel := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				createResourceInstanceOptionsModel.Name = core.StringPtr("ExampleResourceInstance")
				createResourceInstanceOptionsModel.Target = core.StringPtr("global")
				createResourceInstanceOptionsModel.ResourceGroup = core.StringPtr("13aa3ee48c3b44ddb64c05c79f7ab8ef")
				createResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a10e4960-3685-11e9-b210-d663bd873d93")
				createResourceInstanceOptionsModel.Tags = []string{"testString"}
				createResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				createResourceInstanceOptionsModel.EntityLock = core.BoolPtr(false)
				createResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.CreateResourceInstance(createResourceInstanceOptionsModel)
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
	Describe(`GetResourceInstance(getResourceInstanceOptions *GetResourceInstanceOptions) - Operation response error`, func() {
		getResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceInstancePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceInstance with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceInstanceOptions model
				getResourceInstanceOptionsModel := new(resourcecontrollerv2.GetResourceInstanceOptions)
				getResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				getResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceInstance(getResourceInstanceOptions *GetResourceInstanceOptions)`, func() {
		getResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceInstancePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke GetResourceInstance successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceInstanceOptions model
				getResourceInstanceOptionsModel := new(resourcecontrollerv2.GetResourceInstanceOptions)
				getResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				getResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.GetResourceInstanceWithContext(ctx, getResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.GetResourceInstanceWithContext(ctx, getResourceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceInstancePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke GetResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.GetResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceInstanceOptions model
				getResourceInstanceOptionsModel := new(resourcecontrollerv2.GetResourceInstanceOptions)
				getResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				getResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceInstance with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceInstanceOptions model
				getResourceInstanceOptionsModel := new(resourcecontrollerv2.GetResourceInstanceOptions)
				getResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				getResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceInstanceOptions model with no property values
				getResourceInstanceOptionsModelNew := new(resourcecontrollerv2.GetResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.GetResourceInstance(getResourceInstanceOptionsModelNew)
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
			It(`Invoke GetResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceInstanceOptions model
				getResourceInstanceOptionsModel := new(resourcecontrollerv2.GetResourceInstanceOptions)
				getResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				getResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.GetResourceInstance(getResourceInstanceOptionsModel)
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
	Describe(`DeleteResourceInstance(deleteResourceInstanceOptions *DeleteResourceInstanceOptions)`, func() {
		deleteResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for recursive query parameter
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := resourceControllerService.DeleteResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceInstanceOptions model
				deleteResourceInstanceOptionsModel := new(resourcecontrollerv2.DeleteResourceInstanceOptions)
				deleteResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				deleteResourceInstanceOptionsModel.Recursive = core.BoolPtr(false)
				deleteResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = resourceControllerService.DeleteResourceInstance(deleteResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceInstance with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceInstanceOptions model
				deleteResourceInstanceOptionsModel := new(resourcecontrollerv2.DeleteResourceInstanceOptions)
				deleteResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				deleteResourceInstanceOptionsModel.Recursive = core.BoolPtr(false)
				deleteResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := resourceControllerService.DeleteResourceInstance(deleteResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceInstanceOptions model with no property values
				deleteResourceInstanceOptionsModelNew := new(resourcecontrollerv2.DeleteResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = resourceControllerService.DeleteResourceInstance(deleteResourceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceInstance(updateResourceInstanceOptions *UpdateResourceInstanceOptions) - Operation response error`, func() {
		updateResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceInstance with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceInstanceOptions model
				updateResourceInstanceOptionsModel := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				updateResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.Name = core.StringPtr("UpdatedExampleResourceInstance")
				updateResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				updateResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceInstance(updateResourceInstanceOptions *UpdateResourceInstanceOptions)`, func() {
		updateResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))

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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke UpdateResourceInstance successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the UpdateResourceInstanceOptions model
				updateResourceInstanceOptionsModel := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				updateResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.Name = core.StringPtr("UpdatedExampleResourceInstance")
				updateResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				updateResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.UpdateResourceInstanceWithContext(ctx, updateResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.UpdateResourceInstanceWithContext(ctx, updateResourceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke UpdateResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.UpdateResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceInstanceOptions model
				updateResourceInstanceOptionsModel := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				updateResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.Name = core.StringPtr("UpdatedExampleResourceInstance")
				updateResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				updateResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateResourceInstance with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceInstanceOptions model
				updateResourceInstanceOptionsModel := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				updateResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.Name = core.StringPtr("UpdatedExampleResourceInstance")
				updateResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				updateResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceInstanceOptions model with no property values
				updateResourceInstanceOptionsModelNew := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.UpdateResourceInstance(updateResourceInstanceOptionsModelNew)
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
			It(`Invoke UpdateResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceInstanceOptions model
				updateResourceInstanceOptionsModel := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				updateResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.Name = core.StringPtr("UpdatedExampleResourceInstance")
				updateResourceInstanceOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				updateResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
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
	Describe(`ListResourceAliasesForInstance(listResourceAliasesForInstanceOptions *ListResourceAliasesForInstanceOptions) - Operation response error`, func() {
		listResourceAliasesForInstancePath := "/v2/resource_instances/testString/resource_aliases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceAliasesForInstancePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceAliasesForInstance with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceAliasesForInstanceOptions model
				listResourceAliasesForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceAliasesForInstanceOptions)
				listResourceAliasesForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListResourceAliasesForInstance(listResourceAliasesForInstanceOptions *ListResourceAliasesForInstanceOptions)`, func() {
		listResourceAliasesForInstancePath := "/v2/resource_instances/testString/resource_aliases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceAliasesForInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}]}`)
				}))
			})
			It(`Invoke ListResourceAliasesForInstance successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ListResourceAliasesForInstanceOptions model
				listResourceAliasesForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceAliasesForInstanceOptions)
				listResourceAliasesForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.ListResourceAliasesForInstanceWithContext(ctx, listResourceAliasesForInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.ListResourceAliasesForInstanceWithContext(ctx, listResourceAliasesForInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listResourceAliasesForInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}]}`)
				}))
			})
			It(`Invoke ListResourceAliasesForInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.ListResourceAliasesForInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceAliasesForInstanceOptions model
				listResourceAliasesForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceAliasesForInstanceOptions)
				listResourceAliasesForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListResourceAliasesForInstance with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceAliasesForInstanceOptions model
				listResourceAliasesForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceAliasesForInstanceOptions)
				listResourceAliasesForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListResourceAliasesForInstanceOptions model with no property values
				listResourceAliasesForInstanceOptionsModelNew := new(resourcecontrollerv2.ListResourceAliasesForInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptionsModelNew)
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
			It(`Invoke ListResourceAliasesForInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceAliasesForInstanceOptions model
				listResourceAliasesForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceAliasesForInstanceOptions)
				listResourceAliasesForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resourcecontrollerv2.ResourceAliasesList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(resourcecontrollerv2.ResourceAliasesList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(resourcecontrollerv2.ResourceAliasesList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceAliasesForInstancePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","name":"Name","resource_instance_id":"ResourceInstanceID","target_crn":"TargetCRN","account_id":"AccountID","resource_id":"ResourceID","resource_group_id":"ResourceGroupID","crn":"CRN","region_instance_id":"RegionInstanceID","region_instance_crn":"RegionInstanceCRN","state":"State","migrated":true,"resource_instance_url":"ResourceInstanceURL","resource_bindings_url":"ResourceBindingsURL","resource_keys_url":"ResourceKeysURL"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","name":"Name","resource_instance_id":"ResourceInstanceID","target_crn":"TargetCRN","account_id":"AccountID","resource_id":"ResourceID","resource_group_id":"ResourceGroupID","crn":"CRN","region_instance_id":"RegionInstanceID","region_instance_crn":"RegionInstanceCRN","state":"State","migrated":true,"resource_instance_url":"ResourceInstanceURL","resource_bindings_url":"ResourceBindingsURL","resource_keys_url":"ResourceKeysURL"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ResourceAliasesForInstancePager.GetNext successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceAliasesForInstanceOptionsModel := &resourcecontrollerv2.ListResourceAliasesForInstanceOptions{
					ID:    core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := resourceControllerService.NewResourceAliasesForInstancePager(listResourceAliasesForInstanceOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceAlias
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ResourceAliasesForInstancePager.GetAll successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceAliasesForInstanceOptionsModel := &resourcecontrollerv2.ListResourceAliasesForInstanceOptions{
					ID:    core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := resourceControllerService.NewResourceAliasesForInstancePager(listResourceAliasesForInstanceOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ListResourceKeysForInstance(listResourceKeysForInstanceOptions *ListResourceKeysForInstanceOptions) - Operation response error`, func() {
		listResourceKeysForInstancePath := "/v2/resource_instances/testString/resource_keys"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceKeysForInstancePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceKeysForInstance with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceKeysForInstanceOptions model
				listResourceKeysForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceKeysForInstanceOptions)
				listResourceKeysForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListResourceKeysForInstance(listResourceKeysForInstanceOptions *ListResourceKeysForInstanceOptions)`, func() {
		listResourceKeysForInstancePath := "/v2/resource_instances/testString/resource_keys"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceKeysForInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}]}`)
				}))
			})
			It(`Invoke ListResourceKeysForInstance successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ListResourceKeysForInstanceOptions model
				listResourceKeysForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceKeysForInstanceOptions)
				listResourceKeysForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.ListResourceKeysForInstanceWithContext(ctx, listResourceKeysForInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.ListResourceKeysForInstanceWithContext(ctx, listResourceKeysForInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listResourceKeysForInstancePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}]}`)
				}))
			})
			It(`Invoke ListResourceKeysForInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.ListResourceKeysForInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceKeysForInstanceOptions model
				listResourceKeysForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceKeysForInstanceOptions)
				listResourceKeysForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListResourceKeysForInstance with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceKeysForInstanceOptions model
				listResourceKeysForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceKeysForInstanceOptions)
				listResourceKeysForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListResourceKeysForInstanceOptions model with no property values
				listResourceKeysForInstanceOptionsModelNew := new(resourcecontrollerv2.ListResourceKeysForInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptionsModelNew)
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
			It(`Invoke ListResourceKeysForInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceKeysForInstanceOptions model
				listResourceKeysForInstanceOptionsModel := new(resourcecontrollerv2.ListResourceKeysForInstanceOptions)
				listResourceKeysForInstanceOptionsModel.ID = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysForInstanceOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysForInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resourcecontrollerv2.ResourceKeysList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(resourcecontrollerv2.ResourceKeysList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(resourcecontrollerv2.ResourceKeysList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceKeysForInstancePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","source_crn":"SourceCRN","name":"Name","crn":"CRN","state":"State","account_id":"AccountID","resource_group_id":"ResourceGroupID","resource_id":"ResourceID","onetime_credentials":true,"credentials":{"REDACTED":"REDACTED","apikey":"Apikey","iam_apikey_description":"IamApikeyDescription","iam_apikey_name":"IamApikeyName","iam_role_crn":"IamRoleCRN","iam_serviceid_crn":"IamServiceidCRN"},"iam_compatible":false,"migrated":true,"resource_instance_url":"ResourceInstanceURL","resource_alias_url":"ResourceAliasURL"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","source_crn":"SourceCRN","name":"Name","crn":"CRN","state":"State","account_id":"AccountID","resource_group_id":"ResourceGroupID","resource_id":"ResourceID","onetime_credentials":true,"credentials":{"REDACTED":"REDACTED","apikey":"Apikey","iam_apikey_description":"IamApikeyDescription","iam_apikey_name":"IamApikeyName","iam_role_crn":"IamRoleCRN","iam_serviceid_crn":"IamServiceidCRN"},"iam_compatible":false,"migrated":true,"resource_instance_url":"ResourceInstanceURL","resource_alias_url":"ResourceAliasURL"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ResourceKeysForInstancePager.GetNext successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceKeysForInstanceOptionsModel := &resourcecontrollerv2.ListResourceKeysForInstanceOptions{
					ID:    core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := resourceControllerService.NewResourceKeysForInstancePager(listResourceKeysForInstanceOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceKey
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ResourceKeysForInstancePager.GetAll successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceKeysForInstanceOptionsModel := &resourcecontrollerv2.ListResourceKeysForInstanceOptions{
					ID:    core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := resourceControllerService.NewResourceKeysForInstancePager(listResourceKeysForInstanceOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`LockResourceInstance(lockResourceInstanceOptions *LockResourceInstanceOptions) - Operation response error`, func() {
		lockResourceInstancePath := "/v2/resource_instances/testString/lock"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(lockResourceInstancePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke LockResourceInstance with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the LockResourceInstanceOptions model
				lockResourceInstanceOptionsModel := new(resourcecontrollerv2.LockResourceInstanceOptions)
				lockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				lockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.LockResourceInstance(lockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.LockResourceInstance(lockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`LockResourceInstance(lockResourceInstanceOptions *LockResourceInstanceOptions)`, func() {
		lockResourceInstancePath := "/v2/resource_instances/testString/lock"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(lockResourceInstancePath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke LockResourceInstance successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the LockResourceInstanceOptions model
				lockResourceInstanceOptionsModel := new(resourcecontrollerv2.LockResourceInstanceOptions)
				lockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				lockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.LockResourceInstanceWithContext(ctx, lockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.LockResourceInstance(lockResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.LockResourceInstanceWithContext(ctx, lockResourceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(lockResourceInstancePath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke LockResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.LockResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LockResourceInstanceOptions model
				lockResourceInstanceOptionsModel := new(resourcecontrollerv2.LockResourceInstanceOptions)
				lockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				lockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.LockResourceInstance(lockResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke LockResourceInstance with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the LockResourceInstanceOptions model
				lockResourceInstanceOptionsModel := new(resourcecontrollerv2.LockResourceInstanceOptions)
				lockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				lockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.LockResourceInstance(lockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the LockResourceInstanceOptions model with no property values
				lockResourceInstanceOptionsModelNew := new(resourcecontrollerv2.LockResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.LockResourceInstance(lockResourceInstanceOptionsModelNew)
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
			It(`Invoke LockResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the LockResourceInstanceOptions model
				lockResourceInstanceOptionsModel := new(resourcecontrollerv2.LockResourceInstanceOptions)
				lockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				lockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.LockResourceInstance(lockResourceInstanceOptionsModel)
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
	Describe(`UnlockResourceInstance(unlockResourceInstanceOptions *UnlockResourceInstanceOptions) - Operation response error`, func() {
		unlockResourceInstancePath := "/v2/resource_instances/testString/lock"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unlockResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UnlockResourceInstance with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UnlockResourceInstanceOptions model
				unlockResourceInstanceOptionsModel := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				unlockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				unlockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.UnlockResourceInstance(unlockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.UnlockResourceInstance(unlockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UnlockResourceInstance(unlockResourceInstanceOptions *UnlockResourceInstanceOptions)`, func() {
		unlockResourceInstancePath := "/v2/resource_instances/testString/lock"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unlockResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke UnlockResourceInstance successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the UnlockResourceInstanceOptions model
				unlockResourceInstanceOptionsModel := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				unlockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				unlockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.UnlockResourceInstanceWithContext(ctx, unlockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.UnlockResourceInstance(unlockResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.UnlockResourceInstanceWithContext(ctx, unlockResourceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(unlockResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke UnlockResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.UnlockResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UnlockResourceInstanceOptions model
				unlockResourceInstanceOptionsModel := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				unlockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				unlockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.UnlockResourceInstance(unlockResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UnlockResourceInstance with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UnlockResourceInstanceOptions model
				unlockResourceInstanceOptionsModel := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				unlockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				unlockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.UnlockResourceInstance(unlockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UnlockResourceInstanceOptions model with no property values
				unlockResourceInstanceOptionsModelNew := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.UnlockResourceInstance(unlockResourceInstanceOptionsModelNew)
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
			It(`Invoke UnlockResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UnlockResourceInstanceOptions model
				unlockResourceInstanceOptionsModel := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				unlockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				unlockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.UnlockResourceInstance(unlockResourceInstanceOptionsModel)
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
	Describe(`CancelLastopResourceInstance(cancelLastopResourceInstanceOptions *CancelLastopResourceInstanceOptions) - Operation response error`, func() {
		cancelLastopResourceInstancePath := "/v2/resource_instances/testString/last_operation"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cancelLastopResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CancelLastopResourceInstance with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the CancelLastopResourceInstanceOptions model
				cancelLastopResourceInstanceOptionsModel := new(resourcecontrollerv2.CancelLastopResourceInstanceOptions)
				cancelLastopResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				cancelLastopResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CancelLastopResourceInstance(cancelLastopResourceInstanceOptions *CancelLastopResourceInstanceOptions)`, func() {
		cancelLastopResourceInstancePath := "/v2/resource_instances/testString/last_operation"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(cancelLastopResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke CancelLastopResourceInstance successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the CancelLastopResourceInstanceOptions model
				cancelLastopResourceInstanceOptionsModel := new(resourcecontrollerv2.CancelLastopResourceInstanceOptions)
				cancelLastopResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				cancelLastopResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.CancelLastopResourceInstanceWithContext(ctx, cancelLastopResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.CancelLastopResourceInstanceWithContext(ctx, cancelLastopResourceInstanceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(cancelLastopResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "scheduled_reclaim_at": "2019-01-01T12:00:00.000Z", "restored_at": "2019-01-01T12:00:00.000Z", "restored_by": "RestoredBy", "scheduled_reclaim_by": "ScheduledReclaimBy", "name": "Name", "region_id": "RegionID", "account_id": "AccountID", "reseller_channel_id": "ResellerChannelID", "resource_plan_id": "ResourcePlanID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCRN", "target_crn": "TargetCRN", "onetime_credentials": true, "parameters": {"anyKey": "anyValue"}, "allow_cleanup": true, "crn": "CRN", "state": "active", "type": "Type", "sub_type": "SubType", "resource_id": "ResourceID", "dashboard_url": "DashboardURL", "last_operation": {"type": "Type", "state": "in progress", "sub_type": "SubType", "async": false, "description": "Description", "reason_code": "ReasonCode", "poll_after": 9, "cancelable": true, "poll": true}, "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00.000Z", "requestor_id": "RequestorID"}], "migrated": true, "extensions": {"anyKey": "anyValue"}, "controlled_by": "ControlledBy", "locked": true}`)
				}))
			})
			It(`Invoke CancelLastopResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.CancelLastopResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CancelLastopResourceInstanceOptions model
				cancelLastopResourceInstanceOptionsModel := new(resourcecontrollerv2.CancelLastopResourceInstanceOptions)
				cancelLastopResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				cancelLastopResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CancelLastopResourceInstance with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the CancelLastopResourceInstanceOptions model
				cancelLastopResourceInstanceOptionsModel := new(resourcecontrollerv2.CancelLastopResourceInstanceOptions)
				cancelLastopResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				cancelLastopResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CancelLastopResourceInstanceOptions model with no property values
				cancelLastopResourceInstanceOptionsModelNew := new(resourcecontrollerv2.CancelLastopResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptionsModelNew)
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
			It(`Invoke CancelLastopResourceInstance successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the CancelLastopResourceInstanceOptions model
				cancelLastopResourceInstanceOptionsModel := new(resourcecontrollerv2.CancelLastopResourceInstanceOptions)
				cancelLastopResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				cancelLastopResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptionsModel)
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
	Describe(`ListResourceKeys(listResourceKeysOptions *ListResourceKeysOptions) - Operation response error`, func() {
		listResourceKeysPath := "/v2/resource_keys"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceKeysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceKeys with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := new(resourcecontrollerv2.ListResourceKeysOptions)
				listResourceKeysOptionsModel.GUID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Name = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.ListResourceKeys(listResourceKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.ListResourceKeys(listResourceKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListResourceKeys(listResourceKeysOptions *ListResourceKeysOptions)`, func() {
		listResourceKeysPath := "/v2/resource_keys"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceKeysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}]}`)
				}))
			})
			It(`Invoke ListResourceKeys successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := new(resourcecontrollerv2.ListResourceKeysOptions)
				listResourceKeysOptionsModel.GUID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Name = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.ListResourceKeysWithContext(ctx, listResourceKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.ListResourceKeys(listResourceKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.ListResourceKeysWithContext(ctx, listResourceKeysOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listResourceKeysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}]}`)
				}))
			})
			It(`Invoke ListResourceKeys successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.ListResourceKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := new(resourcecontrollerv2.ListResourceKeysOptions)
				listResourceKeysOptionsModel.GUID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Name = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.ListResourceKeys(listResourceKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListResourceKeys with error: Operation request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := new(resourcecontrollerv2.ListResourceKeysOptions)
				listResourceKeysOptionsModel.GUID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Name = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.ListResourceKeys(listResourceKeysOptionsModel)
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
			It(`Invoke ListResourceKeys successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := new(resourcecontrollerv2.ListResourceKeysOptions)
				listResourceKeysOptionsModel.GUID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Name = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceKeysOptionsModel.Start = core.StringPtr("testString")
				listResourceKeysOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.ListResourceKeys(listResourceKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resourcecontrollerv2.ResourceKeysList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(resourcecontrollerv2.ResourceKeysList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(resourcecontrollerv2.ResourceKeysList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceKeysPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","source_crn":"SourceCRN","name":"Name","crn":"CRN","state":"State","account_id":"AccountID","resource_group_id":"ResourceGroupID","resource_id":"ResourceID","onetime_credentials":true,"credentials":{"REDACTED":"REDACTED","apikey":"Apikey","iam_apikey_description":"IamApikeyDescription","iam_apikey_name":"IamApikeyName","iam_role_crn":"IamRoleCRN","iam_serviceid_crn":"IamServiceidCRN"},"iam_compatible":false,"migrated":true,"resource_instance_url":"ResourceInstanceURL","resource_alias_url":"ResourceAliasURL"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","source_crn":"SourceCRN","name":"Name","crn":"CRN","state":"State","account_id":"AccountID","resource_group_id":"ResourceGroupID","resource_id":"ResourceID","onetime_credentials":true,"credentials":{"REDACTED":"REDACTED","apikey":"Apikey","iam_apikey_description":"IamApikeyDescription","iam_apikey_name":"IamApikeyName","iam_role_crn":"IamRoleCRN","iam_serviceid_crn":"IamServiceidCRN"},"iam_compatible":false,"migrated":true,"resource_instance_url":"ResourceInstanceURL","resource_alias_url":"ResourceAliasURL"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ResourceKeysPager.GetNext successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceKeysOptionsModel := &resourcecontrollerv2.ListResourceKeysOptions{
					GUID:            core.StringPtr("testString"),
					Name:            core.StringPtr("testString"),
					ResourceGroupID: core.StringPtr("testString"),
					ResourceID:      core.StringPtr("testString"),
					Limit:           core.Int64Ptr(int64(10)),
					UpdatedFrom:     core.StringPtr("2021-01-01"),
					UpdatedTo:       core.StringPtr("2021-01-01"),
				}

				pager, err := resourceControllerService.NewResourceKeysPager(listResourceKeysOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceKey
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ResourceKeysPager.GetAll successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceKeysOptionsModel := &resourcecontrollerv2.ListResourceKeysOptions{
					GUID:            core.StringPtr("testString"),
					Name:            core.StringPtr("testString"),
					ResourceGroupID: core.StringPtr("testString"),
					ResourceID:      core.StringPtr("testString"),
					Limit:           core.Int64Ptr(int64(10)),
					UpdatedFrom:     core.StringPtr("2021-01-01"),
					UpdatedTo:       core.StringPtr("2021-01-01"),
				}

				pager, err := resourceControllerService.NewResourceKeysPager(listResourceKeysOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateResourceKey(createResourceKeyOptions *CreateResourceKeyOptions) - Operation response error`, func() {
		createResourceKeyPath := "/v2/resource_keys"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceKeyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceKey with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				resourceKeyPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceKeyPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsModel := new(resourcecontrollerv2.CreateResourceKeyOptions)
				createResourceKeyOptionsModel.Name = core.StringPtr("ExampleResourceKey")
				createResourceKeyOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceKeyOptionsModel.Parameters = resourceKeyPostParametersModel
				createResourceKeyOptionsModel.Role = core.StringPtr("Writer")
				createResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.CreateResourceKey(createResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.CreateResourceKey(createResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResourceKey(createResourceKeyOptions *CreateResourceKeyOptions)`, func() {
		createResourceKeyPath := "/v2/resource_keys"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceKeyPath))
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke CreateResourceKey successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				resourceKeyPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceKeyPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsModel := new(resourcecontrollerv2.CreateResourceKeyOptions)
				createResourceKeyOptionsModel.Name = core.StringPtr("ExampleResourceKey")
				createResourceKeyOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceKeyOptionsModel.Parameters = resourceKeyPostParametersModel
				createResourceKeyOptionsModel.Role = core.StringPtr("Writer")
				createResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.CreateResourceKeyWithContext(ctx, createResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.CreateResourceKey(createResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.CreateResourceKeyWithContext(ctx, createResourceKeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createResourceKeyPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke CreateResourceKey successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.CreateResourceKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				resourceKeyPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceKeyPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsModel := new(resourcecontrollerv2.CreateResourceKeyOptions)
				createResourceKeyOptionsModel.Name = core.StringPtr("ExampleResourceKey")
				createResourceKeyOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceKeyOptionsModel.Parameters = resourceKeyPostParametersModel
				createResourceKeyOptionsModel.Role = core.StringPtr("Writer")
				createResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.CreateResourceKey(createResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateResourceKey with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				resourceKeyPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceKeyPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsModel := new(resourcecontrollerv2.CreateResourceKeyOptions)
				createResourceKeyOptionsModel.Name = core.StringPtr("ExampleResourceKey")
				createResourceKeyOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceKeyOptionsModel.Parameters = resourceKeyPostParametersModel
				createResourceKeyOptionsModel.Role = core.StringPtr("Writer")
				createResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.CreateResourceKey(createResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceKeyOptions model with no property values
				createResourceKeyOptionsModelNew := new(resourcecontrollerv2.CreateResourceKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.CreateResourceKey(createResourceKeyOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateResourceKey successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				resourceKeyPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceKeyPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsModel := new(resourcecontrollerv2.CreateResourceKeyOptions)
				createResourceKeyOptionsModel.Name = core.StringPtr("ExampleResourceKey")
				createResourceKeyOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceKeyOptionsModel.Parameters = resourceKeyPostParametersModel
				createResourceKeyOptionsModel.Role = core.StringPtr("Writer")
				createResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.CreateResourceKey(createResourceKeyOptionsModel)
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
	Describe(`GetResourceKey(getResourceKeyOptions *GetResourceKeyOptions) - Operation response error`, func() {
		getResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceKeyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceKey with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceKeyOptions model
				getResourceKeyOptionsModel := new(resourcecontrollerv2.GetResourceKeyOptions)
				getResourceKeyOptionsModel.ID = core.StringPtr("testString")
				getResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.GetResourceKey(getResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.GetResourceKey(getResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceKey(getResourceKeyOptions *GetResourceKeyOptions)`, func() {
		getResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceKeyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke GetResourceKey successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceKeyOptions model
				getResourceKeyOptionsModel := new(resourcecontrollerv2.GetResourceKeyOptions)
				getResourceKeyOptionsModel.ID = core.StringPtr("testString")
				getResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.GetResourceKeyWithContext(ctx, getResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.GetResourceKey(getResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.GetResourceKeyWithContext(ctx, getResourceKeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceKeyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke GetResourceKey successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.GetResourceKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceKeyOptions model
				getResourceKeyOptionsModel := new(resourcecontrollerv2.GetResourceKeyOptions)
				getResourceKeyOptionsModel.ID = core.StringPtr("testString")
				getResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.GetResourceKey(getResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceKey with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceKeyOptions model
				getResourceKeyOptionsModel := new(resourcecontrollerv2.GetResourceKeyOptions)
				getResourceKeyOptionsModel.ID = core.StringPtr("testString")
				getResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.GetResourceKey(getResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceKeyOptions model with no property values
				getResourceKeyOptionsModelNew := new(resourcecontrollerv2.GetResourceKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.GetResourceKey(getResourceKeyOptionsModelNew)
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
			It(`Invoke GetResourceKey successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceKeyOptions model
				getResourceKeyOptionsModel := new(resourcecontrollerv2.GetResourceKeyOptions)
				getResourceKeyOptionsModel.ID = core.StringPtr("testString")
				getResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.GetResourceKey(getResourceKeyOptionsModel)
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
	Describe(`DeleteResourceKey(deleteResourceKeyOptions *DeleteResourceKeyOptions)`, func() {
		deleteResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteResourceKeyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceKey successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := resourceControllerService.DeleteResourceKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceKeyOptions model
				deleteResourceKeyOptionsModel := new(resourcecontrollerv2.DeleteResourceKeyOptions)
				deleteResourceKeyOptionsModel.ID = core.StringPtr("testString")
				deleteResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = resourceControllerService.DeleteResourceKey(deleteResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceKey with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceKeyOptions model
				deleteResourceKeyOptionsModel := new(resourcecontrollerv2.DeleteResourceKeyOptions)
				deleteResourceKeyOptionsModel.ID = core.StringPtr("testString")
				deleteResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := resourceControllerService.DeleteResourceKey(deleteResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceKeyOptions model with no property values
				deleteResourceKeyOptionsModelNew := new(resourcecontrollerv2.DeleteResourceKeyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = resourceControllerService.DeleteResourceKey(deleteResourceKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceKey(updateResourceKeyOptions *UpdateResourceKeyOptions) - Operation response error`, func() {
		updateResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceKeyPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceKey with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceKeyOptions model
				updateResourceKeyOptionsModel := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				updateResourceKeyOptionsModel.ID = core.StringPtr("testString")
				updateResourceKeyOptionsModel.Name = core.StringPtr("UpdatedExampleResourceKey")
				updateResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.UpdateResourceKey(updateResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.UpdateResourceKey(updateResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceKey(updateResourceKeyOptions *UpdateResourceKeyOptions)`, func() {
		updateResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceKeyPath))
					Expect(req.Method).To(Equal("PATCH"))

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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke UpdateResourceKey successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the UpdateResourceKeyOptions model
				updateResourceKeyOptionsModel := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				updateResourceKeyOptionsModel.ID = core.StringPtr("testString")
				updateResourceKeyOptionsModel.Name = core.StringPtr("UpdatedExampleResourceKey")
				updateResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.UpdateResourceKeyWithContext(ctx, updateResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.UpdateResourceKey(updateResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.UpdateResourceKeyWithContext(ctx, updateResourceKeyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceKeyPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "name": "Name", "crn": "CRN", "state": "State", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_id": "ResourceID", "onetime_credentials": true, "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke UpdateResourceKey successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.UpdateResourceKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceKeyOptions model
				updateResourceKeyOptionsModel := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				updateResourceKeyOptionsModel.ID = core.StringPtr("testString")
				updateResourceKeyOptionsModel.Name = core.StringPtr("UpdatedExampleResourceKey")
				updateResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.UpdateResourceKey(updateResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateResourceKey with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceKeyOptions model
				updateResourceKeyOptionsModel := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				updateResourceKeyOptionsModel.ID = core.StringPtr("testString")
				updateResourceKeyOptionsModel.Name = core.StringPtr("UpdatedExampleResourceKey")
				updateResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.UpdateResourceKey(updateResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceKeyOptions model with no property values
				updateResourceKeyOptionsModelNew := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.UpdateResourceKey(updateResourceKeyOptionsModelNew)
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
			It(`Invoke UpdateResourceKey successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceKeyOptions model
				updateResourceKeyOptionsModel := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				updateResourceKeyOptionsModel.ID = core.StringPtr("testString")
				updateResourceKeyOptionsModel.Name = core.StringPtr("UpdatedExampleResourceKey")
				updateResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.UpdateResourceKey(updateResourceKeyOptionsModel)
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
	Describe(`ListResourceBindings(listResourceBindingsOptions *ListResourceBindingsOptions) - Operation response error`, func() {
		listResourceBindingsPath := "/v2/resource_bindings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceBindingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region_binding_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceBindings with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := new(resourcecontrollerv2.ListResourceBindingsOptions)
				listResourceBindingsOptionsModel.GUID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Name = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.RegionBindingID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.ListResourceBindings(listResourceBindingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.ListResourceBindings(listResourceBindingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListResourceBindings(listResourceBindingsOptions *ListResourceBindingsOptions)`, func() {
		listResourceBindingsPath := "/v2/resource_bindings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceBindingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region_binding_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}]}`)
				}))
			})
			It(`Invoke ListResourceBindings successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := new(resourcecontrollerv2.ListResourceBindingsOptions)
				listResourceBindingsOptionsModel.GUID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Name = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.RegionBindingID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.ListResourceBindingsWithContext(ctx, listResourceBindingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.ListResourceBindings(listResourceBindingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.ListResourceBindingsWithContext(ctx, listResourceBindingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listResourceBindingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region_binding_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}]}`)
				}))
			})
			It(`Invoke ListResourceBindings successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.ListResourceBindings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := new(resourcecontrollerv2.ListResourceBindingsOptions)
				listResourceBindingsOptionsModel.GUID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Name = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.RegionBindingID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.ListResourceBindings(listResourceBindingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListResourceBindings with error: Operation request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := new(resourcecontrollerv2.ListResourceBindingsOptions)
				listResourceBindingsOptionsModel.GUID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Name = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.RegionBindingID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.ListResourceBindings(listResourceBindingsOptionsModel)
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
			It(`Invoke ListResourceBindings successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := new(resourcecontrollerv2.ListResourceBindingsOptions)
				listResourceBindingsOptionsModel.GUID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Name = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.RegionBindingID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceBindingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.ListResourceBindings(listResourceBindingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resourcecontrollerv2.ResourceBindingsList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(resourcecontrollerv2.ResourceBindingsList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(resourcecontrollerv2.ResourceBindingsList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceBindingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","source_crn":"SourceCRN","target_crn":"TargetCRN","crn":"CRN","region_binding_id":"RegionBindingID","region_binding_crn":"RegionBindingCRN","name":"Name","account_id":"AccountID","resource_group_id":"ResourceGroupID","state":"State","credentials":{"REDACTED":"REDACTED","apikey":"Apikey","iam_apikey_description":"IamApikeyDescription","iam_apikey_name":"IamApikeyName","iam_role_crn":"IamRoleCRN","iam_serviceid_crn":"IamServiceidCRN"},"iam_compatible":false,"resource_id":"ResourceID","migrated":true,"resource_alias_url":"ResourceAliasURL"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","source_crn":"SourceCRN","target_crn":"TargetCRN","crn":"CRN","region_binding_id":"RegionBindingID","region_binding_crn":"RegionBindingCRN","name":"Name","account_id":"AccountID","resource_group_id":"ResourceGroupID","state":"State","credentials":{"REDACTED":"REDACTED","apikey":"Apikey","iam_apikey_description":"IamApikeyDescription","iam_apikey_name":"IamApikeyName","iam_role_crn":"IamRoleCRN","iam_serviceid_crn":"IamServiceidCRN"},"iam_compatible":false,"resource_id":"ResourceID","migrated":true,"resource_alias_url":"ResourceAliasURL"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ResourceBindingsPager.GetNext successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceBindingsOptionsModel := &resourcecontrollerv2.ListResourceBindingsOptions{
					GUID:            core.StringPtr("testString"),
					Name:            core.StringPtr("testString"),
					ResourceGroupID: core.StringPtr("testString"),
					ResourceID:      core.StringPtr("testString"),
					RegionBindingID: core.StringPtr("testString"),
					Limit:           core.Int64Ptr(int64(10)),
					UpdatedFrom:     core.StringPtr("2021-01-01"),
					UpdatedTo:       core.StringPtr("2021-01-01"),
				}

				pager, err := resourceControllerService.NewResourceBindingsPager(listResourceBindingsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceBinding
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ResourceBindingsPager.GetAll successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceBindingsOptionsModel := &resourcecontrollerv2.ListResourceBindingsOptions{
					GUID:            core.StringPtr("testString"),
					Name:            core.StringPtr("testString"),
					ResourceGroupID: core.StringPtr("testString"),
					ResourceID:      core.StringPtr("testString"),
					RegionBindingID: core.StringPtr("testString"),
					Limit:           core.Int64Ptr(int64(10)),
					UpdatedFrom:     core.StringPtr("2021-01-01"),
					UpdatedTo:       core.StringPtr("2021-01-01"),
				}

				pager, err := resourceControllerService.NewResourceBindingsPager(listResourceBindingsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateResourceBinding(createResourceBindingOptions *CreateResourceBindingOptions) - Operation response error`, func() {
		createResourceBindingPath := "/v2/resource_bindings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceBindingPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceBinding with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				resourceBindingPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceBindingPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsModel := new(resourcecontrollerv2.CreateResourceBindingOptions)
				createResourceBindingOptionsModel.Source = core.StringPtr("faaec9d8-ec64-44d8-ab83-868632fac6a2")
				createResourceBindingOptionsModel.Target = core.StringPtr("crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205")
				createResourceBindingOptionsModel.Name = core.StringPtr("ExampleResourceBinding")
				createResourceBindingOptionsModel.Parameters = resourceBindingPostParametersModel
				createResourceBindingOptionsModel.Role = core.StringPtr("Writer")
				createResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.CreateResourceBinding(createResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.CreateResourceBinding(createResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResourceBinding(createResourceBindingOptions *CreateResourceBindingOptions)`, func() {
		createResourceBindingPath := "/v2/resource_bindings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceBindingPath))
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke CreateResourceBinding successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				resourceBindingPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceBindingPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsModel := new(resourcecontrollerv2.CreateResourceBindingOptions)
				createResourceBindingOptionsModel.Source = core.StringPtr("faaec9d8-ec64-44d8-ab83-868632fac6a2")
				createResourceBindingOptionsModel.Target = core.StringPtr("crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205")
				createResourceBindingOptionsModel.Name = core.StringPtr("ExampleResourceBinding")
				createResourceBindingOptionsModel.Parameters = resourceBindingPostParametersModel
				createResourceBindingOptionsModel.Role = core.StringPtr("Writer")
				createResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.CreateResourceBindingWithContext(ctx, createResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.CreateResourceBinding(createResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.CreateResourceBindingWithContext(ctx, createResourceBindingOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createResourceBindingPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke CreateResourceBinding successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.CreateResourceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				resourceBindingPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceBindingPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsModel := new(resourcecontrollerv2.CreateResourceBindingOptions)
				createResourceBindingOptionsModel.Source = core.StringPtr("faaec9d8-ec64-44d8-ab83-868632fac6a2")
				createResourceBindingOptionsModel.Target = core.StringPtr("crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205")
				createResourceBindingOptionsModel.Name = core.StringPtr("ExampleResourceBinding")
				createResourceBindingOptionsModel.Parameters = resourceBindingPostParametersModel
				createResourceBindingOptionsModel.Role = core.StringPtr("Writer")
				createResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.CreateResourceBinding(createResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateResourceBinding with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				resourceBindingPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceBindingPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsModel := new(resourcecontrollerv2.CreateResourceBindingOptions)
				createResourceBindingOptionsModel.Source = core.StringPtr("faaec9d8-ec64-44d8-ab83-868632fac6a2")
				createResourceBindingOptionsModel.Target = core.StringPtr("crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205")
				createResourceBindingOptionsModel.Name = core.StringPtr("ExampleResourceBinding")
				createResourceBindingOptionsModel.Parameters = resourceBindingPostParametersModel
				createResourceBindingOptionsModel.Role = core.StringPtr("Writer")
				createResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.CreateResourceBinding(createResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceBindingOptions model with no property values
				createResourceBindingOptionsModelNew := new(resourcecontrollerv2.CreateResourceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.CreateResourceBinding(createResourceBindingOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateResourceBinding successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				resourceBindingPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceBindingPostParametersModel.SetProperty("exampleParameter", "exampleValue")

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsModel := new(resourcecontrollerv2.CreateResourceBindingOptions)
				createResourceBindingOptionsModel.Source = core.StringPtr("faaec9d8-ec64-44d8-ab83-868632fac6a2")
				createResourceBindingOptionsModel.Target = core.StringPtr("crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205")
				createResourceBindingOptionsModel.Name = core.StringPtr("ExampleResourceBinding")
				createResourceBindingOptionsModel.Parameters = resourceBindingPostParametersModel
				createResourceBindingOptionsModel.Role = core.StringPtr("Writer")
				createResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.CreateResourceBinding(createResourceBindingOptionsModel)
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
	Describe(`GetResourceBinding(getResourceBindingOptions *GetResourceBindingOptions) - Operation response error`, func() {
		getResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceBindingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceBinding with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceBindingOptions model
				getResourceBindingOptionsModel := new(resourcecontrollerv2.GetResourceBindingOptions)
				getResourceBindingOptionsModel.ID = core.StringPtr("testString")
				getResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.GetResourceBinding(getResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.GetResourceBinding(getResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceBinding(getResourceBindingOptions *GetResourceBindingOptions)`, func() {
		getResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceBindingPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke GetResourceBinding successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceBindingOptions model
				getResourceBindingOptionsModel := new(resourcecontrollerv2.GetResourceBindingOptions)
				getResourceBindingOptionsModel.ID = core.StringPtr("testString")
				getResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.GetResourceBindingWithContext(ctx, getResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.GetResourceBinding(getResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.GetResourceBindingWithContext(ctx, getResourceBindingOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceBindingPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke GetResourceBinding successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.GetResourceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceBindingOptions model
				getResourceBindingOptionsModel := new(resourcecontrollerv2.GetResourceBindingOptions)
				getResourceBindingOptionsModel.ID = core.StringPtr("testString")
				getResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.GetResourceBinding(getResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceBinding with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceBindingOptions model
				getResourceBindingOptionsModel := new(resourcecontrollerv2.GetResourceBindingOptions)
				getResourceBindingOptionsModel.ID = core.StringPtr("testString")
				getResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.GetResourceBinding(getResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceBindingOptions model with no property values
				getResourceBindingOptionsModelNew := new(resourcecontrollerv2.GetResourceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.GetResourceBinding(getResourceBindingOptionsModelNew)
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
			It(`Invoke GetResourceBinding successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceBindingOptions model
				getResourceBindingOptionsModel := new(resourcecontrollerv2.GetResourceBindingOptions)
				getResourceBindingOptionsModel.ID = core.StringPtr("testString")
				getResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.GetResourceBinding(getResourceBindingOptionsModel)
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
	Describe(`DeleteResourceBinding(deleteResourceBindingOptions *DeleteResourceBindingOptions)`, func() {
		deleteResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteResourceBindingPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceBinding successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := resourceControllerService.DeleteResourceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceBindingOptions model
				deleteResourceBindingOptionsModel := new(resourcecontrollerv2.DeleteResourceBindingOptions)
				deleteResourceBindingOptionsModel.ID = core.StringPtr("testString")
				deleteResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = resourceControllerService.DeleteResourceBinding(deleteResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceBinding with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceBindingOptions model
				deleteResourceBindingOptionsModel := new(resourcecontrollerv2.DeleteResourceBindingOptions)
				deleteResourceBindingOptionsModel.ID = core.StringPtr("testString")
				deleteResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := resourceControllerService.DeleteResourceBinding(deleteResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceBindingOptions model with no property values
				deleteResourceBindingOptionsModelNew := new(resourcecontrollerv2.DeleteResourceBindingOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = resourceControllerService.DeleteResourceBinding(deleteResourceBindingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceBinding(updateResourceBindingOptions *UpdateResourceBindingOptions) - Operation response error`, func() {
		updateResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceBindingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceBinding with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceBindingOptions model
				updateResourceBindingOptionsModel := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				updateResourceBindingOptionsModel.ID = core.StringPtr("testString")
				updateResourceBindingOptionsModel.Name = core.StringPtr("UpdatedExampleResourceBinding")
				updateResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.UpdateResourceBinding(updateResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.UpdateResourceBinding(updateResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceBinding(updateResourceBindingOptions *UpdateResourceBindingOptions)`, func() {
		updateResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceBindingPath))
					Expect(req.Method).To(Equal("PATCH"))

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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke UpdateResourceBinding successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the UpdateResourceBindingOptions model
				updateResourceBindingOptionsModel := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				updateResourceBindingOptionsModel.ID = core.StringPtr("testString")
				updateResourceBindingOptionsModel.Name = core.StringPtr("UpdatedExampleResourceBinding")
				updateResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.UpdateResourceBindingWithContext(ctx, updateResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.UpdateResourceBinding(updateResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.UpdateResourceBindingWithContext(ctx, updateResourceBindingOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceBindingPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}`)
				}))
			})
			It(`Invoke UpdateResourceBinding successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.UpdateResourceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceBindingOptions model
				updateResourceBindingOptionsModel := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				updateResourceBindingOptionsModel.ID = core.StringPtr("testString")
				updateResourceBindingOptionsModel.Name = core.StringPtr("UpdatedExampleResourceBinding")
				updateResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.UpdateResourceBinding(updateResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateResourceBinding with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceBindingOptions model
				updateResourceBindingOptionsModel := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				updateResourceBindingOptionsModel.ID = core.StringPtr("testString")
				updateResourceBindingOptionsModel.Name = core.StringPtr("UpdatedExampleResourceBinding")
				updateResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.UpdateResourceBinding(updateResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceBindingOptions model with no property values
				updateResourceBindingOptionsModelNew := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.UpdateResourceBinding(updateResourceBindingOptionsModelNew)
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
			It(`Invoke UpdateResourceBinding successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceBindingOptions model
				updateResourceBindingOptionsModel := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				updateResourceBindingOptionsModel.ID = core.StringPtr("testString")
				updateResourceBindingOptionsModel.Name = core.StringPtr("UpdatedExampleResourceBinding")
				updateResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.UpdateResourceBinding(updateResourceBindingOptionsModel)
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
	Describe(`ListResourceAliases(listResourceAliasesOptions *ListResourceAliasesOptions) - Operation response error`, func() {
		listResourceAliasesPath := "/v2/resource_aliases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceAliasesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceAliases with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := new(resourcecontrollerv2.ListResourceAliasesOptions)
				listResourceAliasesOptionsModel.GUID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Name = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.RegionInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.ListResourceAliases(listResourceAliasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.ListResourceAliases(listResourceAliasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListResourceAliases(listResourceAliasesOptions *ListResourceAliasesOptions)`, func() {
		listResourceAliasesPath := "/v2/resource_aliases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceAliasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}]}`)
				}))
			})
			It(`Invoke ListResourceAliases successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := new(resourcecontrollerv2.ListResourceAliasesOptions)
				listResourceAliasesOptionsModel.GUID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Name = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.RegionInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.ListResourceAliasesWithContext(ctx, listResourceAliasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.ListResourceAliases(listResourceAliasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.ListResourceAliasesWithContext(ctx, listResourceAliasesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listResourceAliasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2021-01-01"}))
					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2021-01-01"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}]}`)
				}))
			})
			It(`Invoke ListResourceAliases successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.ListResourceAliases(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := new(resourcecontrollerv2.ListResourceAliasesOptions)
				listResourceAliasesOptionsModel.GUID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Name = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.RegionInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.ListResourceAliases(listResourceAliasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListResourceAliases with error: Operation request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := new(resourcecontrollerv2.ListResourceAliasesOptions)
				listResourceAliasesOptionsModel.GUID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Name = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.RegionInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.ListResourceAliases(listResourceAliasesOptionsModel)
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
			It(`Invoke ListResourceAliases successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := new(resourcecontrollerv2.ListResourceAliasesOptions)
				listResourceAliasesOptionsModel.GUID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Name = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.RegionInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceAliasesOptionsModel.Start = core.StringPtr("testString")
				listResourceAliasesOptionsModel.UpdatedFrom = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.UpdatedTo = core.StringPtr("2021-01-01")
				listResourceAliasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.ListResourceAliases(listResourceAliasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resourcecontrollerv2.ResourceAliasesList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(resourcecontrollerv2.ResourceAliasesList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(resourcecontrollerv2.ResourceAliasesList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceAliasesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","name":"Name","resource_instance_id":"ResourceInstanceID","target_crn":"TargetCRN","account_id":"AccountID","resource_id":"ResourceID","resource_group_id":"ResourceGroupID","crn":"CRN","region_instance_id":"RegionInstanceID","region_instance_crn":"RegionInstanceCRN","state":"State","migrated":true,"resource_instance_url":"ResourceInstanceURL","resource_bindings_url":"ResourceBindingsURL","resource_keys_url":"ResourceKeysURL"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","name":"Name","resource_instance_id":"ResourceInstanceID","target_crn":"TargetCRN","account_id":"AccountID","resource_id":"ResourceID","resource_group_id":"ResourceGroupID","crn":"CRN","region_instance_id":"RegionInstanceID","region_instance_crn":"RegionInstanceCRN","state":"State","migrated":true,"resource_instance_url":"ResourceInstanceURL","resource_bindings_url":"ResourceBindingsURL","resource_keys_url":"ResourceKeysURL"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ResourceAliasesPager.GetNext successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceAliasesOptionsModel := &resourcecontrollerv2.ListResourceAliasesOptions{
					GUID:               core.StringPtr("testString"),
					Name:               core.StringPtr("testString"),
					ResourceInstanceID: core.StringPtr("testString"),
					RegionInstanceID:   core.StringPtr("testString"),
					ResourceID:         core.StringPtr("testString"),
					ResourceGroupID:    core.StringPtr("testString"),
					Limit:              core.Int64Ptr(int64(10)),
					UpdatedFrom:        core.StringPtr("2021-01-01"),
					UpdatedTo:          core.StringPtr("2021-01-01"),
				}

				pager, err := resourceControllerService.NewResourceAliasesPager(listResourceAliasesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceAlias
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ResourceAliasesPager.GetAll successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceAliasesOptionsModel := &resourcecontrollerv2.ListResourceAliasesOptions{
					GUID:               core.StringPtr("testString"),
					Name:               core.StringPtr("testString"),
					ResourceInstanceID: core.StringPtr("testString"),
					RegionInstanceID:   core.StringPtr("testString"),
					ResourceID:         core.StringPtr("testString"),
					ResourceGroupID:    core.StringPtr("testString"),
					Limit:              core.Int64Ptr(int64(10)),
					UpdatedFrom:        core.StringPtr("2021-01-01"),
					UpdatedTo:          core.StringPtr("2021-01-01"),
				}

				pager, err := resourceControllerService.NewResourceAliasesPager(listResourceAliasesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateResourceAlias(createResourceAliasOptions *CreateResourceAliasOptions) - Operation response error`, func() {
		createResourceAliasPath := "/v2/resource_aliases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceAliasPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceAlias with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsModel := new(resourcecontrollerv2.CreateResourceAliasOptions)
				createResourceAliasOptionsModel.Name = core.StringPtr("ExampleResourceAlias")
				createResourceAliasOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceAliasOptionsModel.Target = core.StringPtr("crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7")
				createResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.CreateResourceAlias(createResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.CreateResourceAlias(createResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResourceAlias(createResourceAliasOptions *CreateResourceAliasOptions)`, func() {
		createResourceAliasPath := "/v2/resource_aliases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceAliasPath))
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}`)
				}))
			})
			It(`Invoke CreateResourceAlias successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsModel := new(resourcecontrollerv2.CreateResourceAliasOptions)
				createResourceAliasOptionsModel.Name = core.StringPtr("ExampleResourceAlias")
				createResourceAliasOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceAliasOptionsModel.Target = core.StringPtr("crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7")
				createResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.CreateResourceAliasWithContext(ctx, createResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.CreateResourceAlias(createResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.CreateResourceAliasWithContext(ctx, createResourceAliasOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createResourceAliasPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}`)
				}))
			})
			It(`Invoke CreateResourceAlias successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.CreateResourceAlias(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsModel := new(resourcecontrollerv2.CreateResourceAliasOptions)
				createResourceAliasOptionsModel.Name = core.StringPtr("ExampleResourceAlias")
				createResourceAliasOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceAliasOptionsModel.Target = core.StringPtr("crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7")
				createResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.CreateResourceAlias(createResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateResourceAlias with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsModel := new(resourcecontrollerv2.CreateResourceAliasOptions)
				createResourceAliasOptionsModel.Name = core.StringPtr("ExampleResourceAlias")
				createResourceAliasOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceAliasOptionsModel.Target = core.StringPtr("crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7")
				createResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.CreateResourceAlias(createResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceAliasOptions model with no property values
				createResourceAliasOptionsModelNew := new(resourcecontrollerv2.CreateResourceAliasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.CreateResourceAlias(createResourceAliasOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateResourceAlias successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsModel := new(resourcecontrollerv2.CreateResourceAliasOptions)
				createResourceAliasOptionsModel.Name = core.StringPtr("ExampleResourceAlias")
				createResourceAliasOptionsModel.Source = core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceAliasOptionsModel.Target = core.StringPtr("crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7")
				createResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.CreateResourceAlias(createResourceAliasOptionsModel)
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
	Describe(`GetResourceAlias(getResourceAliasOptions *GetResourceAliasOptions) - Operation response error`, func() {
		getResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceAliasPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceAlias with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceAliasOptions model
				getResourceAliasOptionsModel := new(resourcecontrollerv2.GetResourceAliasOptions)
				getResourceAliasOptionsModel.ID = core.StringPtr("testString")
				getResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.GetResourceAlias(getResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.GetResourceAlias(getResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceAlias(getResourceAliasOptions *GetResourceAliasOptions)`, func() {
		getResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceAliasPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}`)
				}))
			})
			It(`Invoke GetResourceAlias successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceAliasOptions model
				getResourceAliasOptionsModel := new(resourcecontrollerv2.GetResourceAliasOptions)
				getResourceAliasOptionsModel.ID = core.StringPtr("testString")
				getResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.GetResourceAliasWithContext(ctx, getResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.GetResourceAlias(getResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.GetResourceAliasWithContext(ctx, getResourceAliasOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceAliasPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}`)
				}))
			})
			It(`Invoke GetResourceAlias successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.GetResourceAlias(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceAliasOptions model
				getResourceAliasOptionsModel := new(resourcecontrollerv2.GetResourceAliasOptions)
				getResourceAliasOptionsModel.ID = core.StringPtr("testString")
				getResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.GetResourceAlias(getResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceAlias with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceAliasOptions model
				getResourceAliasOptionsModel := new(resourcecontrollerv2.GetResourceAliasOptions)
				getResourceAliasOptionsModel.ID = core.StringPtr("testString")
				getResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.GetResourceAlias(getResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceAliasOptions model with no property values
				getResourceAliasOptionsModelNew := new(resourcecontrollerv2.GetResourceAliasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.GetResourceAlias(getResourceAliasOptionsModelNew)
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
			It(`Invoke GetResourceAlias successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the GetResourceAliasOptions model
				getResourceAliasOptionsModel := new(resourcecontrollerv2.GetResourceAliasOptions)
				getResourceAliasOptionsModel.ID = core.StringPtr("testString")
				getResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.GetResourceAlias(getResourceAliasOptionsModel)
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
	Describe(`DeleteResourceAlias(deleteResourceAliasOptions *DeleteResourceAliasOptions)`, func() {
		deleteResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteResourceAliasPath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for recursive query parameter
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceAlias successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := resourceControllerService.DeleteResourceAlias(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceAliasOptions model
				deleteResourceAliasOptionsModel := new(resourcecontrollerv2.DeleteResourceAliasOptions)
				deleteResourceAliasOptionsModel.ID = core.StringPtr("testString")
				deleteResourceAliasOptionsModel.Recursive = core.BoolPtr(false)
				deleteResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = resourceControllerService.DeleteResourceAlias(deleteResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceAlias with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceAliasOptions model
				deleteResourceAliasOptionsModel := new(resourcecontrollerv2.DeleteResourceAliasOptions)
				deleteResourceAliasOptionsModel.ID = core.StringPtr("testString")
				deleteResourceAliasOptionsModel.Recursive = core.BoolPtr(false)
				deleteResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := resourceControllerService.DeleteResourceAlias(deleteResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceAliasOptions model with no property values
				deleteResourceAliasOptionsModelNew := new(resourcecontrollerv2.DeleteResourceAliasOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = resourceControllerService.DeleteResourceAlias(deleteResourceAliasOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceAlias(updateResourceAliasOptions *UpdateResourceAliasOptions) - Operation response error`, func() {
		updateResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceAliasPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceAlias with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceAliasOptions model
				updateResourceAliasOptionsModel := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				updateResourceAliasOptionsModel.ID = core.StringPtr("testString")
				updateResourceAliasOptionsModel.Name = core.StringPtr("UpdatedExampleResourceAlias")
				updateResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.UpdateResourceAlias(updateResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.UpdateResourceAlias(updateResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceAlias(updateResourceAliasOptions *UpdateResourceAliasOptions)`, func() {
		updateResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceAliasPath))
					Expect(req.Method).To(Equal("PATCH"))

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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}`)
				}))
			})
			It(`Invoke UpdateResourceAlias successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the UpdateResourceAliasOptions model
				updateResourceAliasOptionsModel := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				updateResourceAliasOptionsModel.ID = core.StringPtr("testString")
				updateResourceAliasOptionsModel.Name = core.StringPtr("UpdatedExampleResourceAlias")
				updateResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.UpdateResourceAliasWithContext(ctx, updateResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.UpdateResourceAlias(updateResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.UpdateResourceAliasWithContext(ctx, updateResourceAliasOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceAliasPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "name": "Name", "resource_instance_id": "ResourceInstanceID", "target_crn": "TargetCRN", "account_id": "AccountID", "resource_id": "ResourceID", "resource_group_id": "ResourceGroupID", "crn": "CRN", "region_instance_id": "RegionInstanceID", "region_instance_crn": "RegionInstanceCRN", "state": "State", "migrated": true, "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL"}`)
				}))
			})
			It(`Invoke UpdateResourceAlias successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.UpdateResourceAlias(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceAliasOptions model
				updateResourceAliasOptionsModel := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				updateResourceAliasOptionsModel.ID = core.StringPtr("testString")
				updateResourceAliasOptionsModel.Name = core.StringPtr("UpdatedExampleResourceAlias")
				updateResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.UpdateResourceAlias(updateResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateResourceAlias with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceAliasOptions model
				updateResourceAliasOptionsModel := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				updateResourceAliasOptionsModel.ID = core.StringPtr("testString")
				updateResourceAliasOptionsModel.Name = core.StringPtr("UpdatedExampleResourceAlias")
				updateResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.UpdateResourceAlias(updateResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceAliasOptions model with no property values
				updateResourceAliasOptionsModelNew := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.UpdateResourceAlias(updateResourceAliasOptionsModelNew)
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
			It(`Invoke UpdateResourceAlias successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceAliasOptions model
				updateResourceAliasOptionsModel := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				updateResourceAliasOptionsModel.ID = core.StringPtr("testString")
				updateResourceAliasOptionsModel.Name = core.StringPtr("UpdatedExampleResourceAlias")
				updateResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.UpdateResourceAlias(updateResourceAliasOptionsModel)
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
	Describe(`ListResourceBindingsForAlias(listResourceBindingsForAliasOptions *ListResourceBindingsForAliasOptions) - Operation response error`, func() {
		listResourceBindingsForAliasPath := "/v2/resource_aliases/testString/resource_bindings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceBindingsForAliasPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceBindingsForAlias with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceBindingsForAliasOptions model
				listResourceBindingsForAliasOptionsModel := new(resourcecontrollerv2.ListResourceBindingsForAliasOptions)
				listResourceBindingsForAliasOptionsModel.ID = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsForAliasOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListResourceBindingsForAlias(listResourceBindingsForAliasOptions *ListResourceBindingsForAliasOptions)`, func() {
		listResourceBindingsForAliasPath := "/v2/resource_aliases/testString/resource_bindings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceBindingsForAliasPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}]}`)
				}))
			})
			It(`Invoke ListResourceBindingsForAlias successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ListResourceBindingsForAliasOptions model
				listResourceBindingsForAliasOptionsModel := new(resourcecontrollerv2.ListResourceBindingsForAliasOptions)
				listResourceBindingsForAliasOptionsModel.ID = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsForAliasOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.ListResourceBindingsForAliasWithContext(ctx, listResourceBindingsForAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.ListResourceBindingsForAliasWithContext(ctx, listResourceBindingsForAliasOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listResourceBindingsForAliasPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "guid": "GUID", "url": "URL", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "deleted_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_by": "UpdatedBy", "deleted_by": "DeletedBy", "source_crn": "SourceCRN", "target_crn": "TargetCRN", "crn": "CRN", "region_binding_id": "RegionBindingID", "region_binding_crn": "RegionBindingCRN", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "state": "State", "credentials": {"REDACTED": "REDACTED", "apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCRN", "iam_serviceid_crn": "IamServiceidCRN"}, "iam_compatible": false, "resource_id": "ResourceID", "migrated": true, "resource_alias_url": "ResourceAliasURL"}]}`)
				}))
			})
			It(`Invoke ListResourceBindingsForAlias successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.ListResourceBindingsForAlias(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceBindingsForAliasOptions model
				listResourceBindingsForAliasOptionsModel := new(resourcecontrollerv2.ListResourceBindingsForAliasOptions)
				listResourceBindingsForAliasOptionsModel.ID = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsForAliasOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListResourceBindingsForAlias with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceBindingsForAliasOptions model
				listResourceBindingsForAliasOptionsModel := new(resourcecontrollerv2.ListResourceBindingsForAliasOptions)
				listResourceBindingsForAliasOptionsModel.ID = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsForAliasOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListResourceBindingsForAliasOptions model with no property values
				listResourceBindingsForAliasOptionsModelNew := new(resourcecontrollerv2.ListResourceBindingsForAliasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptionsModelNew)
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
			It(`Invoke ListResourceBindingsForAlias successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListResourceBindingsForAliasOptions model
				listResourceBindingsForAliasOptionsModel := new(resourcecontrollerv2.ListResourceBindingsForAliasOptions)
				listResourceBindingsForAliasOptionsModel.ID = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Limit = core.Int64Ptr(int64(10))
				listResourceBindingsForAliasOptionsModel.Start = core.StringPtr("testString")
				listResourceBindingsForAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(resourcecontrollerv2.ResourceBindingsList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(resourcecontrollerv2.ResourceBindingsList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(resourcecontrollerv2.ResourceBindingsList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceBindingsForAliasPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","source_crn":"SourceCRN","target_crn":"TargetCRN","crn":"CRN","region_binding_id":"RegionBindingID","region_binding_crn":"RegionBindingCRN","name":"Name","account_id":"AccountID","resource_group_id":"ResourceGroupID","state":"State","credentials":{"REDACTED":"REDACTED","apikey":"Apikey","iam_apikey_description":"IamApikeyDescription","iam_apikey_name":"IamApikeyName","iam_role_crn":"IamRoleCRN","iam_serviceid_crn":"IamServiceidCRN"},"iam_compatible":false,"resource_id":"ResourceID","migrated":true,"resource_alias_url":"ResourceAliasURL"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","guid":"GUID","url":"URL","created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z","deleted_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_by":"UpdatedBy","deleted_by":"DeletedBy","source_crn":"SourceCRN","target_crn":"TargetCRN","crn":"CRN","region_binding_id":"RegionBindingID","region_binding_crn":"RegionBindingCRN","name":"Name","account_id":"AccountID","resource_group_id":"ResourceGroupID","state":"State","credentials":{"REDACTED":"REDACTED","apikey":"Apikey","iam_apikey_description":"IamApikeyDescription","iam_apikey_name":"IamApikeyName","iam_role_crn":"IamRoleCRN","iam_serviceid_crn":"IamServiceidCRN"},"iam_compatible":false,"resource_id":"ResourceID","migrated":true,"resource_alias_url":"ResourceAliasURL"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ResourceBindingsForAliasPager.GetNext successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceBindingsForAliasOptionsModel := &resourcecontrollerv2.ListResourceBindingsForAliasOptions{
					ID:    core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := resourceControllerService.NewResourceBindingsForAliasPager(listResourceBindingsForAliasOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceBinding
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ResourceBindingsForAliasPager.GetAll successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				listResourceBindingsForAliasOptionsModel := &resourcecontrollerv2.ListResourceBindingsForAliasOptions{
					ID:    core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := resourceControllerService.NewResourceBindingsForAliasPager(listResourceBindingsForAliasOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ListReclamations(listReclamationsOptions *ListReclamationsOptions) - Operation response error`, func() {
		listReclamationsPath := "/v1/reclamations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReclamations with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(resourcecontrollerv2.ListReclamationsOptions)
				listReclamationsOptionsModel.AccountID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListReclamations(listReclamationsOptions *ListReclamationsOptions)`, func() {
		listReclamationsPath := "/v1/reclamations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resources": [{"id": "ID", "entity_id": "EntityID", "entity_type_id": "EntityTypeID", "entity_crn": "EntityCRN", "resource_instance_id": "ResourceInstanceID", "resource_group_id": "ResourceGroupID", "account_id": "AccountID", "policy_id": "PolicyID", "state": "State", "target_time": "TargetTime", "custom_properties": {"anyKey": "anyValue"}, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListReclamations successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(resourcecontrollerv2.ListReclamationsOptions)
				listReclamationsOptionsModel.AccountID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.ListReclamationsWithContext(ctx, listReclamationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.ListReclamationsWithContext(ctx, listReclamationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resources": [{"id": "ID", "entity_id": "EntityID", "entity_type_id": "EntityTypeID", "entity_crn": "EntityCRN", "resource_instance_id": "ResourceInstanceID", "resource_group_id": "ResourceGroupID", "account_id": "AccountID", "policy_id": "PolicyID", "state": "State", "target_time": "TargetTime", "custom_properties": {"anyKey": "anyValue"}, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListReclamations successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.ListReclamations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(resourcecontrollerv2.ListReclamationsOptions)
				listReclamationsOptionsModel.AccountID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListReclamations with error: Operation request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(resourcecontrollerv2.ListReclamationsOptions)
				listReclamationsOptionsModel.AccountID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.ListReclamations(listReclamationsOptionsModel)
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
			It(`Invoke ListReclamations successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(resourcecontrollerv2.ListReclamationsOptions)
				listReclamationsOptionsModel.AccountID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.ListReclamations(listReclamationsOptionsModel)
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
	Describe(`RunReclamationAction(runReclamationActionOptions *RunReclamationActionOptions) - Operation response error`, func() {
		runReclamationActionPath := "/v1/reclamations/testString/actions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runReclamationActionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RunReclamationAction with error: Operation response processing error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the RunReclamationActionOptions model
				runReclamationActionOptionsModel := new(resourcecontrollerv2.RunReclamationActionOptions)
				runReclamationActionOptionsModel.ID = core.StringPtr("testString")
				runReclamationActionOptionsModel.ActionName = core.StringPtr("testString")
				runReclamationActionOptionsModel.RequestBy = core.StringPtr("testString")
				runReclamationActionOptionsModel.Comment = core.StringPtr("testString")
				runReclamationActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceControllerService.RunReclamationAction(runReclamationActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceControllerService.EnableRetries(0, 0)
				result, response, operationErr = resourceControllerService.RunReclamationAction(runReclamationActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RunReclamationAction(runReclamationActionOptions *RunReclamationActionOptions)`, func() {
		runReclamationActionPath := "/v1/reclamations/testString/actions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runReclamationActionPath))
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_id": "EntityID", "entity_type_id": "EntityTypeID", "entity_crn": "EntityCRN", "resource_instance_id": "ResourceInstanceID", "resource_group_id": "ResourceGroupID", "account_id": "AccountID", "policy_id": "PolicyID", "state": "State", "target_time": "TargetTime", "custom_properties": {"anyKey": "anyValue"}, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke RunReclamationAction successfully with retries`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())
				resourceControllerService.EnableRetries(0, 0)

				// Construct an instance of the RunReclamationActionOptions model
				runReclamationActionOptionsModel := new(resourcecontrollerv2.RunReclamationActionOptions)
				runReclamationActionOptionsModel.ID = core.StringPtr("testString")
				runReclamationActionOptionsModel.ActionName = core.StringPtr("testString")
				runReclamationActionOptionsModel.RequestBy = core.StringPtr("testString")
				runReclamationActionOptionsModel.Comment = core.StringPtr("testString")
				runReclamationActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceControllerService.RunReclamationActionWithContext(ctx, runReclamationActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceControllerService.DisableRetries()
				result, response, operationErr := resourceControllerService.RunReclamationAction(runReclamationActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceControllerService.RunReclamationActionWithContext(ctx, runReclamationActionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(runReclamationActionPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "entity_id": "EntityID", "entity_type_id": "EntityTypeID", "entity_crn": "EntityCRN", "resource_instance_id": "ResourceInstanceID", "resource_group_id": "ResourceGroupID", "account_id": "AccountID", "policy_id": "PolicyID", "state": "State", "target_time": "TargetTime", "custom_properties": {"anyKey": "anyValue"}, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke RunReclamationAction successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceControllerService.RunReclamationAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RunReclamationActionOptions model
				runReclamationActionOptionsModel := new(resourcecontrollerv2.RunReclamationActionOptions)
				runReclamationActionOptionsModel.ID = core.StringPtr("testString")
				runReclamationActionOptionsModel.ActionName = core.StringPtr("testString")
				runReclamationActionOptionsModel.RequestBy = core.StringPtr("testString")
				runReclamationActionOptionsModel.Comment = core.StringPtr("testString")
				runReclamationActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceControllerService.RunReclamationAction(runReclamationActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RunReclamationAction with error: Operation validation and request error`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the RunReclamationActionOptions model
				runReclamationActionOptionsModel := new(resourcecontrollerv2.RunReclamationActionOptions)
				runReclamationActionOptionsModel.ID = core.StringPtr("testString")
				runReclamationActionOptionsModel.ActionName = core.StringPtr("testString")
				runReclamationActionOptionsModel.RequestBy = core.StringPtr("testString")
				runReclamationActionOptionsModel.Comment = core.StringPtr("testString")
				runReclamationActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceControllerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceControllerService.RunReclamationAction(runReclamationActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RunReclamationActionOptions model with no property values
				runReclamationActionOptionsModelNew := new(resourcecontrollerv2.RunReclamationActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceControllerService.RunReclamationAction(runReclamationActionOptionsModelNew)
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
			It(`Invoke RunReclamationAction successfully`, func() {
				resourceControllerService, serviceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceControllerService).ToNot(BeNil())

				// Construct an instance of the RunReclamationActionOptions model
				runReclamationActionOptionsModel := new(resourcecontrollerv2.RunReclamationActionOptions)
				runReclamationActionOptionsModel.ID = core.StringPtr("testString")
				runReclamationActionOptionsModel.ActionName = core.StringPtr("testString")
				runReclamationActionOptionsModel.RequestBy = core.StringPtr("testString")
				runReclamationActionOptionsModel.Comment = core.StringPtr("testString")
				runReclamationActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceControllerService.RunReclamationAction(runReclamationActionOptionsModel)
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
			resourceControllerService, _ := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL:           "http://resourcecontrollerv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCancelLastopResourceInstanceOptions successfully`, func() {
				// Construct an instance of the CancelLastopResourceInstanceOptions model
				id := "testString"
				cancelLastopResourceInstanceOptionsModel := resourceControllerService.NewCancelLastopResourceInstanceOptions(id)
				cancelLastopResourceInstanceOptionsModel.SetID("testString")
				cancelLastopResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cancelLastopResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(cancelLastopResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(cancelLastopResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateResourceAliasOptions successfully`, func() {
				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsName := "ExampleResourceAlias"
				createResourceAliasOptionsSource := "381fd51a-f251-4f95-aff4-2b03fa8caa63"
				createResourceAliasOptionsTarget := "crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7"
				createResourceAliasOptionsModel := resourceControllerService.NewCreateResourceAliasOptions(createResourceAliasOptionsName, createResourceAliasOptionsSource, createResourceAliasOptionsTarget)
				createResourceAliasOptionsModel.SetName("ExampleResourceAlias")
				createResourceAliasOptionsModel.SetSource("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceAliasOptionsModel.SetTarget("crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7")
				createResourceAliasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceAliasOptionsModel).ToNot(BeNil())
				Expect(createResourceAliasOptionsModel.Name).To(Equal(core.StringPtr("ExampleResourceAlias")))
				Expect(createResourceAliasOptionsModel.Source).To(Equal(core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")))
				Expect(createResourceAliasOptionsModel.Target).To(Equal(core.StringPtr("crn:v1:bluemix:public:bluemix:us-south:o/d35d4f0e-5076-4c89-9361-2522894b6548::cf-space:e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7")))
				Expect(createResourceAliasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateResourceBindingOptions successfully`, func() {
				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				Expect(resourceBindingPostParametersModel).ToNot(BeNil())
				resourceBindingPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceBindingPostParametersModel.SetProperty("exampleParameter", "exampleValue")
				Expect(resourceBindingPostParametersModel.ServiceidCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")))
				Expect(resourceBindingPostParametersModel.GetProperties()).ToNot(BeEmpty())
				Expect(resourceBindingPostParametersModel.GetProperty("exampleParameter")).To(Equal("exampleValue"))

				resourceBindingPostParametersModel.SetProperties(nil)
				Expect(resourceBindingPostParametersModel.GetProperties()).To(BeEmpty())

				resourceBindingPostParametersModelExpectedMap := make(map[string]interface{})
				resourceBindingPostParametersModelExpectedMap["exampleParameter"] = "exampleValue"
				resourceBindingPostParametersModel.SetProperties(resourceBindingPostParametersModelExpectedMap)
				resourceBindingPostParametersModelActualMap := resourceBindingPostParametersModel.GetProperties()
				Expect(resourceBindingPostParametersModelActualMap).To(Equal(resourceBindingPostParametersModelExpectedMap))

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsSource := "faaec9d8-ec64-44d8-ab83-868632fac6a2"
				createResourceBindingOptionsTarget := "crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205"
				createResourceBindingOptionsModel := resourceControllerService.NewCreateResourceBindingOptions(createResourceBindingOptionsSource, createResourceBindingOptionsTarget)
				createResourceBindingOptionsModel.SetSource("faaec9d8-ec64-44d8-ab83-868632fac6a2")
				createResourceBindingOptionsModel.SetTarget("crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205")
				createResourceBindingOptionsModel.SetName("ExampleResourceBinding")
				createResourceBindingOptionsModel.SetParameters(resourceBindingPostParametersModel)
				createResourceBindingOptionsModel.SetRole("Writer")
				createResourceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceBindingOptionsModel).ToNot(BeNil())
				Expect(createResourceBindingOptionsModel.Source).To(Equal(core.StringPtr("faaec9d8-ec64-44d8-ab83-868632fac6a2")))
				Expect(createResourceBindingOptionsModel.Target).To(Equal(core.StringPtr("crn:v1:staging:public:bluemix:us-south:s/e1773b6e-17b4-40c8-b5ed-d2a1c4b620d7::cf-application:8d9457e0-1303-4f32-b4b3-5525575f6205")))
				Expect(createResourceBindingOptionsModel.Name).To(Equal(core.StringPtr("ExampleResourceBinding")))
				Expect(createResourceBindingOptionsModel.Parameters).To(Equal(resourceBindingPostParametersModel))
				Expect(createResourceBindingOptionsModel.Role).To(Equal(core.StringPtr("Writer")))
				Expect(createResourceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateResourceInstanceOptions successfully`, func() {
				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsName := "ExampleResourceInstance"
				createResourceInstanceOptionsTarget := "global"
				createResourceInstanceOptionsResourceGroup := "13aa3ee48c3b44ddb64c05c79f7ab8ef"
				createResourceInstanceOptionsResourcePlanID := "a10e4960-3685-11e9-b210-d663bd873d93"
				createResourceInstanceOptionsModel := resourceControllerService.NewCreateResourceInstanceOptions(createResourceInstanceOptionsName, createResourceInstanceOptionsTarget, createResourceInstanceOptionsResourceGroup, createResourceInstanceOptionsResourcePlanID)
				createResourceInstanceOptionsModel.SetName("ExampleResourceInstance")
				createResourceInstanceOptionsModel.SetTarget("global")
				createResourceInstanceOptionsModel.SetResourceGroup("13aa3ee48c3b44ddb64c05c79f7ab8ef")
				createResourceInstanceOptionsModel.SetResourcePlanID("a10e4960-3685-11e9-b210-d663bd873d93")
				createResourceInstanceOptionsModel.SetTags([]string{"testString"})
				createResourceInstanceOptionsModel.SetAllowCleanup(false)
				createResourceInstanceOptionsModel.SetParameters(map[string]interface{}{"anyKey": "anyValue"})
				createResourceInstanceOptionsModel.SetEntityLock(false)
				createResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(createResourceInstanceOptionsModel.Name).To(Equal(core.StringPtr("ExampleResourceInstance")))
				Expect(createResourceInstanceOptionsModel.Target).To(Equal(core.StringPtr("global")))
				Expect(createResourceInstanceOptionsModel.ResourceGroup).To(Equal(core.StringPtr("13aa3ee48c3b44ddb64c05c79f7ab8ef")))
				Expect(createResourceInstanceOptionsModel.ResourcePlanID).To(Equal(core.StringPtr("a10e4960-3685-11e9-b210-d663bd873d93")))
				Expect(createResourceInstanceOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createResourceInstanceOptionsModel.AllowCleanup).To(Equal(core.BoolPtr(false)))
				Expect(createResourceInstanceOptionsModel.Parameters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createResourceInstanceOptionsModel.EntityLock).To(Equal(core.BoolPtr(false)))
				Expect(createResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateResourceKeyOptions successfully`, func() {
				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				Expect(resourceKeyPostParametersModel).ToNot(BeNil())
				resourceKeyPostParametersModel.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				resourceKeyPostParametersModel.SetProperty("exampleParameter", "exampleValue")
				Expect(resourceKeyPostParametersModel.ServiceidCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")))
				Expect(resourceKeyPostParametersModel.GetProperties()).ToNot(BeEmpty())
				Expect(resourceKeyPostParametersModel.GetProperty("exampleParameter")).To(Equal("exampleValue"))

				resourceKeyPostParametersModel.SetProperties(nil)
				Expect(resourceKeyPostParametersModel.GetProperties()).To(BeEmpty())

				resourceKeyPostParametersModelExpectedMap := make(map[string]interface{})
				resourceKeyPostParametersModelExpectedMap["exampleParameter"] = "exampleValue"
				resourceKeyPostParametersModel.SetProperties(resourceKeyPostParametersModelExpectedMap)
				resourceKeyPostParametersModelActualMap := resourceKeyPostParametersModel.GetProperties()
				Expect(resourceKeyPostParametersModelActualMap).To(Equal(resourceKeyPostParametersModelExpectedMap))

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsName := "ExampleResourceKey"
				createResourceKeyOptionsSource := "381fd51a-f251-4f95-aff4-2b03fa8caa63"
				createResourceKeyOptionsModel := resourceControllerService.NewCreateResourceKeyOptions(createResourceKeyOptionsName, createResourceKeyOptionsSource)
				createResourceKeyOptionsModel.SetName("ExampleResourceKey")
				createResourceKeyOptionsModel.SetSource("381fd51a-f251-4f95-aff4-2b03fa8caa63")
				createResourceKeyOptionsModel.SetParameters(resourceKeyPostParametersModel)
				createResourceKeyOptionsModel.SetRole("Writer")
				createResourceKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceKeyOptionsModel).ToNot(BeNil())
				Expect(createResourceKeyOptionsModel.Name).To(Equal(core.StringPtr("ExampleResourceKey")))
				Expect(createResourceKeyOptionsModel.Source).To(Equal(core.StringPtr("381fd51a-f251-4f95-aff4-2b03fa8caa63")))
				Expect(createResourceKeyOptionsModel.Parameters).To(Equal(resourceKeyPostParametersModel))
				Expect(createResourceKeyOptionsModel.Role).To(Equal(core.StringPtr("Writer")))
				Expect(createResourceKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceAliasOptions successfully`, func() {
				// Construct an instance of the DeleteResourceAliasOptions model
				id := "testString"
				deleteResourceAliasOptionsModel := resourceControllerService.NewDeleteResourceAliasOptions(id)
				deleteResourceAliasOptionsModel.SetID("testString")
				deleteResourceAliasOptionsModel.SetRecursive(false)
				deleteResourceAliasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceAliasOptionsModel).ToNot(BeNil())
				Expect(deleteResourceAliasOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceAliasOptionsModel.Recursive).To(Equal(core.BoolPtr(false)))
				Expect(deleteResourceAliasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceBindingOptions successfully`, func() {
				// Construct an instance of the DeleteResourceBindingOptions model
				id := "testString"
				deleteResourceBindingOptionsModel := resourceControllerService.NewDeleteResourceBindingOptions(id)
				deleteResourceBindingOptionsModel.SetID("testString")
				deleteResourceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceBindingOptionsModel).ToNot(BeNil())
				Expect(deleteResourceBindingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceInstanceOptions successfully`, func() {
				// Construct an instance of the DeleteResourceInstanceOptions model
				id := "testString"
				deleteResourceInstanceOptionsModel := resourceControllerService.NewDeleteResourceInstanceOptions(id)
				deleteResourceInstanceOptionsModel.SetID("testString")
				deleteResourceInstanceOptionsModel.SetRecursive(false)
				deleteResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(deleteResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceInstanceOptionsModel.Recursive).To(Equal(core.BoolPtr(false)))
				Expect(deleteResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceKeyOptions successfully`, func() {
				// Construct an instance of the DeleteResourceKeyOptions model
				id := "testString"
				deleteResourceKeyOptionsModel := resourceControllerService.NewDeleteResourceKeyOptions(id)
				deleteResourceKeyOptionsModel.SetID("testString")
				deleteResourceKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceKeyOptionsModel).ToNot(BeNil())
				Expect(deleteResourceKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceAliasOptions successfully`, func() {
				// Construct an instance of the GetResourceAliasOptions model
				id := "testString"
				getResourceAliasOptionsModel := resourceControllerService.NewGetResourceAliasOptions(id)
				getResourceAliasOptionsModel.SetID("testString")
				getResourceAliasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceAliasOptionsModel).ToNot(BeNil())
				Expect(getResourceAliasOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceAliasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceBindingOptions successfully`, func() {
				// Construct an instance of the GetResourceBindingOptions model
				id := "testString"
				getResourceBindingOptionsModel := resourceControllerService.NewGetResourceBindingOptions(id)
				getResourceBindingOptionsModel.SetID("testString")
				getResourceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceBindingOptionsModel).ToNot(BeNil())
				Expect(getResourceBindingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceInstanceOptions successfully`, func() {
				// Construct an instance of the GetResourceInstanceOptions model
				id := "testString"
				getResourceInstanceOptionsModel := resourceControllerService.NewGetResourceInstanceOptions(id)
				getResourceInstanceOptionsModel.SetID("testString")
				getResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(getResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceKeyOptions successfully`, func() {
				// Construct an instance of the GetResourceKeyOptions model
				id := "testString"
				getResourceKeyOptionsModel := resourceControllerService.NewGetResourceKeyOptions(id)
				getResourceKeyOptionsModel.SetID("testString")
				getResourceKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceKeyOptionsModel).ToNot(BeNil())
				Expect(getResourceKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListReclamationsOptions successfully`, func() {
				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := resourceControllerService.NewListReclamationsOptions()
				listReclamationsOptionsModel.SetAccountID("testString")
				listReclamationsOptionsModel.SetResourceInstanceID("testString")
				listReclamationsOptionsModel.SetResourceGroupID("testString")
				listReclamationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReclamationsOptionsModel).ToNot(BeNil())
				Expect(listReclamationsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listReclamationsOptionsModel.ResourceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listReclamationsOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listReclamationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceAliasesForInstanceOptions successfully`, func() {
				// Construct an instance of the ListResourceAliasesForInstanceOptions model
				id := "testString"
				listResourceAliasesForInstanceOptionsModel := resourceControllerService.NewListResourceAliasesForInstanceOptions(id)
				listResourceAliasesForInstanceOptionsModel.SetID("testString")
				listResourceAliasesForInstanceOptionsModel.SetLimit(int64(10))
				listResourceAliasesForInstanceOptionsModel.SetStart("testString")
				listResourceAliasesForInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceAliasesForInstanceOptionsModel).ToNot(BeNil())
				Expect(listResourceAliasesForInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesForInstanceOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listResourceAliasesForInstanceOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesForInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceAliasesOptions successfully`, func() {
				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := resourceControllerService.NewListResourceAliasesOptions()
				listResourceAliasesOptionsModel.SetGUID("testString")
				listResourceAliasesOptionsModel.SetName("testString")
				listResourceAliasesOptionsModel.SetResourceInstanceID("testString")
				listResourceAliasesOptionsModel.SetRegionInstanceID("testString")
				listResourceAliasesOptionsModel.SetResourceID("testString")
				listResourceAliasesOptionsModel.SetResourceGroupID("testString")
				listResourceAliasesOptionsModel.SetLimit(int64(10))
				listResourceAliasesOptionsModel.SetStart("testString")
				listResourceAliasesOptionsModel.SetUpdatedFrom("2021-01-01")
				listResourceAliasesOptionsModel.SetUpdatedTo("2021-01-01")
				listResourceAliasesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceAliasesOptionsModel).ToNot(BeNil())
				Expect(listResourceAliasesOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.ResourceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.RegionInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listResourceAliasesOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.UpdatedFrom).To(Equal(core.StringPtr("2021-01-01")))
				Expect(listResourceAliasesOptionsModel.UpdatedTo).To(Equal(core.StringPtr("2021-01-01")))
				Expect(listResourceAliasesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceBindingsForAliasOptions successfully`, func() {
				// Construct an instance of the ListResourceBindingsForAliasOptions model
				id := "testString"
				listResourceBindingsForAliasOptionsModel := resourceControllerService.NewListResourceBindingsForAliasOptions(id)
				listResourceBindingsForAliasOptionsModel.SetID("testString")
				listResourceBindingsForAliasOptionsModel.SetLimit(int64(10))
				listResourceBindingsForAliasOptionsModel.SetStart("testString")
				listResourceBindingsForAliasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceBindingsForAliasOptionsModel).ToNot(BeNil())
				Expect(listResourceBindingsForAliasOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsForAliasOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listResourceBindingsForAliasOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsForAliasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceBindingsOptions successfully`, func() {
				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := resourceControllerService.NewListResourceBindingsOptions()
				listResourceBindingsOptionsModel.SetGUID("testString")
				listResourceBindingsOptionsModel.SetName("testString")
				listResourceBindingsOptionsModel.SetResourceGroupID("testString")
				listResourceBindingsOptionsModel.SetResourceID("testString")
				listResourceBindingsOptionsModel.SetRegionBindingID("testString")
				listResourceBindingsOptionsModel.SetLimit(int64(10))
				listResourceBindingsOptionsModel.SetStart("testString")
				listResourceBindingsOptionsModel.SetUpdatedFrom("2021-01-01")
				listResourceBindingsOptionsModel.SetUpdatedTo("2021-01-01")
				listResourceBindingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceBindingsOptionsModel).ToNot(BeNil())
				Expect(listResourceBindingsOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.RegionBindingID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listResourceBindingsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.UpdatedFrom).To(Equal(core.StringPtr("2021-01-01")))
				Expect(listResourceBindingsOptionsModel.UpdatedTo).To(Equal(core.StringPtr("2021-01-01")))
				Expect(listResourceBindingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceInstancesOptions successfully`, func() {
				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := resourceControllerService.NewListResourceInstancesOptions()
				listResourceInstancesOptionsModel.SetGUID("testString")
				listResourceInstancesOptionsModel.SetName("testString")
				listResourceInstancesOptionsModel.SetResourceGroupID("testString")
				listResourceInstancesOptionsModel.SetResourceID("testString")
				listResourceInstancesOptionsModel.SetResourcePlanID("testString")
				listResourceInstancesOptionsModel.SetType("testString")
				listResourceInstancesOptionsModel.SetSubType("testString")
				listResourceInstancesOptionsModel.SetLimit(int64(10))
				listResourceInstancesOptionsModel.SetStart("testString")
				listResourceInstancesOptionsModel.SetState("active")
				listResourceInstancesOptionsModel.SetUpdatedFrom("2021-01-01")
				listResourceInstancesOptionsModel.SetUpdatedTo("2021-01-01")
				listResourceInstancesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceInstancesOptionsModel).ToNot(BeNil())
				Expect(listResourceInstancesOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.ResourcePlanID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.SubType).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listResourceInstancesOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.State).To(Equal(core.StringPtr("active")))
				Expect(listResourceInstancesOptionsModel.UpdatedFrom).To(Equal(core.StringPtr("2021-01-01")))
				Expect(listResourceInstancesOptionsModel.UpdatedTo).To(Equal(core.StringPtr("2021-01-01")))
				Expect(listResourceInstancesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceKeysForInstanceOptions successfully`, func() {
				// Construct an instance of the ListResourceKeysForInstanceOptions model
				id := "testString"
				listResourceKeysForInstanceOptionsModel := resourceControllerService.NewListResourceKeysForInstanceOptions(id)
				listResourceKeysForInstanceOptionsModel.SetID("testString")
				listResourceKeysForInstanceOptionsModel.SetLimit(int64(10))
				listResourceKeysForInstanceOptionsModel.SetStart("testString")
				listResourceKeysForInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceKeysForInstanceOptionsModel).ToNot(BeNil())
				Expect(listResourceKeysForInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysForInstanceOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listResourceKeysForInstanceOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysForInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceKeysOptions successfully`, func() {
				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := resourceControllerService.NewListResourceKeysOptions()
				listResourceKeysOptionsModel.SetGUID("testString")
				listResourceKeysOptionsModel.SetName("testString")
				listResourceKeysOptionsModel.SetResourceGroupID("testString")
				listResourceKeysOptionsModel.SetResourceID("testString")
				listResourceKeysOptionsModel.SetLimit(int64(10))
				listResourceKeysOptionsModel.SetStart("testString")
				listResourceKeysOptionsModel.SetUpdatedFrom("2021-01-01")
				listResourceKeysOptionsModel.SetUpdatedTo("2021-01-01")
				listResourceKeysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceKeysOptionsModel).ToNot(BeNil())
				Expect(listResourceKeysOptionsModel.GUID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listResourceKeysOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.UpdatedFrom).To(Equal(core.StringPtr("2021-01-01")))
				Expect(listResourceKeysOptionsModel.UpdatedTo).To(Equal(core.StringPtr("2021-01-01")))
				Expect(listResourceKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLockResourceInstanceOptions successfully`, func() {
				// Construct an instance of the LockResourceInstanceOptions model
				id := "testString"
				lockResourceInstanceOptionsModel := resourceControllerService.NewLockResourceInstanceOptions(id)
				lockResourceInstanceOptionsModel.SetID("testString")
				lockResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(lockResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(lockResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(lockResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRunReclamationActionOptions successfully`, func() {
				// Construct an instance of the RunReclamationActionOptions model
				id := "testString"
				actionName := "testString"
				runReclamationActionOptionsModel := resourceControllerService.NewRunReclamationActionOptions(id, actionName)
				runReclamationActionOptionsModel.SetID("testString")
				runReclamationActionOptionsModel.SetActionName("testString")
				runReclamationActionOptionsModel.SetRequestBy("testString")
				runReclamationActionOptionsModel.SetComment("testString")
				runReclamationActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(runReclamationActionOptionsModel).ToNot(BeNil())
				Expect(runReclamationActionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(runReclamationActionOptionsModel.ActionName).To(Equal(core.StringPtr("testString")))
				Expect(runReclamationActionOptionsModel.RequestBy).To(Equal(core.StringPtr("testString")))
				Expect(runReclamationActionOptionsModel.Comment).To(Equal(core.StringPtr("testString")))
				Expect(runReclamationActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnlockResourceInstanceOptions successfully`, func() {
				// Construct an instance of the UnlockResourceInstanceOptions model
				id := "testString"
				unlockResourceInstanceOptionsModel := resourceControllerService.NewUnlockResourceInstanceOptions(id)
				unlockResourceInstanceOptionsModel.SetID("testString")
				unlockResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unlockResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(unlockResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(unlockResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceAliasOptions successfully`, func() {
				// Construct an instance of the UpdateResourceAliasOptions model
				id := "testString"
				updateResourceAliasOptionsName := "UpdatedExampleResourceAlias"
				updateResourceAliasOptionsModel := resourceControllerService.NewUpdateResourceAliasOptions(id, updateResourceAliasOptionsName)
				updateResourceAliasOptionsModel.SetID("testString")
				updateResourceAliasOptionsModel.SetName("UpdatedExampleResourceAlias")
				updateResourceAliasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceAliasOptionsModel).ToNot(BeNil())
				Expect(updateResourceAliasOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceAliasOptionsModel.Name).To(Equal(core.StringPtr("UpdatedExampleResourceAlias")))
				Expect(updateResourceAliasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceBindingOptions successfully`, func() {
				// Construct an instance of the UpdateResourceBindingOptions model
				id := "testString"
				updateResourceBindingOptionsName := "UpdatedExampleResourceBinding"
				updateResourceBindingOptionsModel := resourceControllerService.NewUpdateResourceBindingOptions(id, updateResourceBindingOptionsName)
				updateResourceBindingOptionsModel.SetID("testString")
				updateResourceBindingOptionsModel.SetName("UpdatedExampleResourceBinding")
				updateResourceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceBindingOptionsModel).ToNot(BeNil())
				Expect(updateResourceBindingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceBindingOptionsModel.Name).To(Equal(core.StringPtr("UpdatedExampleResourceBinding")))
				Expect(updateResourceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceInstanceOptions successfully`, func() {
				// Construct an instance of the UpdateResourceInstanceOptions model
				id := "testString"
				updateResourceInstanceOptionsModel := resourceControllerService.NewUpdateResourceInstanceOptions(id)
				updateResourceInstanceOptionsModel.SetID("testString")
				updateResourceInstanceOptionsModel.SetName("UpdatedExampleResourceInstance")
				updateResourceInstanceOptionsModel.SetParameters(map[string]interface{}{"anyKey": "anyValue"})
				updateResourceInstanceOptionsModel.SetResourcePlanID("testString")
				updateResourceInstanceOptionsModel.SetAllowCleanup(true)
				updateResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(updateResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceInstanceOptionsModel.Name).To(Equal(core.StringPtr("UpdatedExampleResourceInstance")))
				Expect(updateResourceInstanceOptionsModel.Parameters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateResourceInstanceOptionsModel.ResourcePlanID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceInstanceOptionsModel.AllowCleanup).To(Equal(core.BoolPtr(true)))
				Expect(updateResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceKeyOptions successfully`, func() {
				// Construct an instance of the UpdateResourceKeyOptions model
				id := "testString"
				updateResourceKeyOptionsName := "UpdatedExampleResourceKey"
				updateResourceKeyOptionsModel := resourceControllerService.NewUpdateResourceKeyOptions(id, updateResourceKeyOptionsName)
				updateResourceKeyOptionsModel.SetID("testString")
				updateResourceKeyOptionsModel.SetName("UpdatedExampleResourceKey")
				updateResourceKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceKeyOptionsModel).ToNot(BeNil())
				Expect(updateResourceKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceKeyOptionsModel.Name).To(Equal(core.StringPtr("UpdatedExampleResourceKey")))
				Expect(updateResourceKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalResourceBindingPostParameters successfully`, func() {
			// Construct an instance of the model.
			model := new(resourcecontrollerv2.ResourceBindingPostParameters)
			model.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourcecontrollerv2.ResourceBindingPostParameters
			err = resourcecontrollerv2.UnmarshalResourceBindingPostParameters(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalResourceKeyPostParameters successfully`, func() {
			// Construct an instance of the model.
			model := new(resourcecontrollerv2.ResourceKeyPostParameters)
			model.ServiceidCRN = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourcecontrollerv2.ResourceKeyPostParameters
			err = resourcecontrollerv2.UnmarshalResourceKeyPostParameters(raw, &result)
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
