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

package ibmcloudshellv1_test

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
	"github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IBMCloudShellV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudShellService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudShellService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
				URL: "https://ibmcloudshellv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudShellService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_SHELL_URL":       "https://ibmcloudshellv1/api",
				"IBM_CLOUD_SHELL_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IBMCloudShellV1Options{})
				Expect(ibmCloudShellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudShellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudShellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudShellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudShellService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudShellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudShellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudShellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudShellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudShellService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IBMCloudShellV1Options{})
				err := ibmCloudShellService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudShellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudShellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudShellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudShellService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_SHELL_URL":       "https://ibmcloudshellv1/api",
				"IBM_CLOUD_SHELL_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IBMCloudShellV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudShellService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_SHELL_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1UsingExternalConfig(&ibmcloudshellv1.IBMCloudShellV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudShellService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudshellv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) - Operation response error`, func() {
		getAccountSettingsPath := "/api/v1/user/accounts/testString/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountSettings with error: Operation response processing error`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(ibmcloudshellv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudShellService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudShellService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudShellService.GetAccountSettings(getAccountSettingsOptionsModel)
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
		getAccountSettingsPath := "/api/v1/user/accounts/testString/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_id": "ID", "_rev": "Rev", "account_id": "AccountID", "created_at": 9, "created_by": "CreatedBy", "default_enable_new_features": true, "default_enable_new_regions": false, "enabled": false, "features": [{"enabled": false, "key": "Key"}], "regions": [{"enabled": false, "key": "Key"}], "type": "Type", "updated_at": 9, "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully with retries`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())
				ibmCloudShellService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(ibmcloudshellv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudShellService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudShellService.DisableRetries()
				result, response, operationErr := ibmCloudShellService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudShellService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_id": "ID", "_rev": "Rev", "account_id": "AccountID", "created_at": 9, "created_by": "CreatedBy", "default_enable_new_features": true, "default_enable_new_regions": false, "enabled": false, "features": [{"enabled": false, "key": "Key"}], "regions": [{"enabled": false, "key": "Key"}], "type": "Type", "updated_at": 9, "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudShellService.GetAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(ibmcloudshellv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudShellService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountSettings with error: Operation validation and request error`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(ibmcloudshellv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudShellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudShellService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountSettingsOptions model with no property values
				getAccountSettingsOptionsModelNew := new(ibmcloudshellv1.GetAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudShellService.GetAccountSettings(getAccountSettingsOptionsModelNew)
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
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(ibmcloudshellv1.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudShellService.GetAccountSettings(getAccountSettingsOptionsModel)
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
	Describe(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions) - Operation response error`, func() {
		updateAccountSettingsPath := "/api/v1/user/accounts/testString/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountSettings with error: Operation response processing error`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(ibmcloudshellv1.Feature)
				featureModel.Enabled = core.BoolPtr(true)
				featureModel.Key = core.StringPtr("testString")

				// Construct an instance of the RegionSetting model
				regionSettingModel := new(ibmcloudshellv1.RegionSetting)
				regionSettingModel.Enabled = core.BoolPtr(true)
				regionSettingModel.Key = core.StringPtr("testString")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Rev = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.DefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.DefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Features = []ibmcloudshellv1.Feature{*featureModel}
				updateAccountSettingsOptionsModel.Regions = []ibmcloudshellv1.RegionSetting{*regionSettingModel}
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudShellService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudShellService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions)`, func() {
		updateAccountSettingsPath := "/api/v1/user/accounts/testString/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
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
					fmt.Fprintf(res, "%s", `{"_id": "ID", "_rev": "Rev", "account_id": "AccountID", "created_at": 9, "created_by": "CreatedBy", "default_enable_new_features": true, "default_enable_new_regions": false, "enabled": false, "features": [{"enabled": false, "key": "Key"}], "regions": [{"enabled": false, "key": "Key"}], "type": "Type", "updated_at": 9, "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully with retries`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())
				ibmCloudShellService.EnableRetries(0, 0)

				// Construct an instance of the Feature model
				featureModel := new(ibmcloudshellv1.Feature)
				featureModel.Enabled = core.BoolPtr(true)
				featureModel.Key = core.StringPtr("testString")

				// Construct an instance of the RegionSetting model
				regionSettingModel := new(ibmcloudshellv1.RegionSetting)
				regionSettingModel.Enabled = core.BoolPtr(true)
				regionSettingModel.Key = core.StringPtr("testString")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Rev = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.DefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.DefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Features = []ibmcloudshellv1.Feature{*featureModel}
				updateAccountSettingsOptionsModel.Regions = []ibmcloudshellv1.RegionSetting{*regionSettingModel}
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudShellService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudShellService.DisableRetries()
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudShellService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
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
					fmt.Fprintf(res, "%s", `{"_id": "ID", "_rev": "Rev", "account_id": "AccountID", "created_at": 9, "created_by": "CreatedBy", "default_enable_new_features": true, "default_enable_new_regions": false, "enabled": false, "features": [{"enabled": false, "key": "Key"}], "regions": [{"enabled": false, "key": "Key"}], "type": "Type", "updated_at": 9, "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(ibmcloudshellv1.Feature)
				featureModel.Enabled = core.BoolPtr(true)
				featureModel.Key = core.StringPtr("testString")

				// Construct an instance of the RegionSetting model
				regionSettingModel := new(ibmcloudshellv1.RegionSetting)
				regionSettingModel.Enabled = core.BoolPtr(true)
				regionSettingModel.Key = core.StringPtr("testString")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Rev = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.DefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.DefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Features = []ibmcloudshellv1.Feature{*featureModel}
				updateAccountSettingsOptionsModel.Regions = []ibmcloudshellv1.RegionSetting{*regionSettingModel}
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudShellService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateAccountSettings with error: Operation validation and request error`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(ibmcloudshellv1.Feature)
				featureModel.Enabled = core.BoolPtr(true)
				featureModel.Key = core.StringPtr("testString")

				// Construct an instance of the RegionSetting model
				regionSettingModel := new(ibmcloudshellv1.RegionSetting)
				regionSettingModel.Enabled = core.BoolPtr(true)
				regionSettingModel.Key = core.StringPtr("testString")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Rev = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.DefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.DefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Features = []ibmcloudshellv1.Feature{*featureModel}
				updateAccountSettingsOptionsModel.Regions = []ibmcloudshellv1.RegionSetting{*regionSettingModel}
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudShellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountSettingsOptions model with no property values
				updateAccountSettingsOptionsModelNew := new(ibmcloudshellv1.UpdateAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudShellService.UpdateAccountSettings(updateAccountSettingsOptionsModelNew)
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
			It(`Invoke UpdateAccountSettings successfully`, func() {
				ibmCloudShellService, serviceErr := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudShellService).ToNot(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(ibmcloudshellv1.Feature)
				featureModel.Enabled = core.BoolPtr(true)
				featureModel.Key = core.StringPtr("testString")

				// Construct an instance of the RegionSetting model
				regionSettingModel := new(ibmcloudshellv1.RegionSetting)
				regionSettingModel.Enabled = core.BoolPtr(true)
				regionSettingModel.Key = core.StringPtr("testString")

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(ibmcloudshellv1.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Rev = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.DefaultEnableNewFeatures = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.DefaultEnableNewRegions = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Enabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.Features = []ibmcloudshellv1.Feature{*featureModel}
				updateAccountSettingsOptionsModel.Regions = []ibmcloudshellv1.RegionSetting{*regionSettingModel}
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudShellService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
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
			ibmCloudShellService, _ := ibmcloudshellv1.NewIBMCloudShellV1(&ibmcloudshellv1.IBMCloudShellV1Options{
				URL:           "http://ibmcloudshellv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetAccountSettingsOptions successfully`, func() {
				// Construct an instance of the GetAccountSettingsOptions model
				accountID := "testString"
				getAccountSettingsOptionsModel := ibmCloudShellService.NewGetAccountSettingsOptions(accountID)
				getAccountSettingsOptionsModel.SetAccountID("testString")
				getAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(getAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountSettingsOptions successfully`, func() {
				// Construct an instance of the Feature model
				featureModel := new(ibmcloudshellv1.Feature)
				Expect(featureModel).ToNot(BeNil())
				featureModel.Enabled = core.BoolPtr(true)
				featureModel.Key = core.StringPtr("testString")
				Expect(featureModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(featureModel.Key).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RegionSetting model
				regionSettingModel := new(ibmcloudshellv1.RegionSetting)
				Expect(regionSettingModel).ToNot(BeNil())
				regionSettingModel.Enabled = core.BoolPtr(true)
				regionSettingModel.Key = core.StringPtr("testString")
				Expect(regionSettingModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(regionSettingModel.Key).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateAccountSettingsOptions model
				accountID := "testString"
				updateAccountSettingsOptionsModel := ibmCloudShellService.NewUpdateAccountSettingsOptions(accountID)
				updateAccountSettingsOptionsModel.SetAccountID("testString")
				updateAccountSettingsOptionsModel.SetRev("testString")
				updateAccountSettingsOptionsModel.SetDefaultEnableNewFeatures(true)
				updateAccountSettingsOptionsModel.SetDefaultEnableNewRegions(true)
				updateAccountSettingsOptionsModel.SetEnabled(true)
				updateAccountSettingsOptionsModel.SetFeatures([]ibmcloudshellv1.Feature{*featureModel})
				updateAccountSettingsOptionsModel.SetRegions([]ibmcloudshellv1.RegionSetting{*regionSettingModel})
				updateAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(updateAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.DefaultEnableNewFeatures).To(Equal(core.BoolPtr(true)))
				Expect(updateAccountSettingsOptionsModel.DefaultEnableNewRegions).To(Equal(core.BoolPtr(true)))
				Expect(updateAccountSettingsOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateAccountSettingsOptionsModel.Features).To(Equal([]ibmcloudshellv1.Feature{*featureModel}))
				Expect(updateAccountSettingsOptionsModel.Regions).To(Equal([]ibmcloudshellv1.RegionSetting{*regionSettingModel}))
				Expect(updateAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAccountSettings successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudshellv1.AccountSettings)
			model.ID = core.StringPtr("testString")
			model.Rev = core.StringPtr("testString")
			model.AccountID = core.StringPtr("testString")
			model.CreatedAt = core.Int64Ptr(int64(38))
			model.CreatedBy = core.StringPtr("testString")
			model.DefaultEnableNewFeatures = core.BoolPtr(true)
			model.DefaultEnableNewRegions = core.BoolPtr(true)
			model.Enabled = core.BoolPtr(true)
			model.Features = nil
			model.Regions = nil
			model.Type = core.StringPtr("testString")
			model.UpdatedAt = core.Int64Ptr(int64(38))
			model.UpdatedBy = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudshellv1.AccountSettings
			err = ibmcloudshellv1.UnmarshalAccountSettings(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalFeature successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudshellv1.Feature)
			model.Enabled = core.BoolPtr(true)
			model.Key = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudshellv1.Feature
			err = ibmcloudshellv1.UnmarshalFeature(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRegionSetting successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudshellv1.RegionSetting)
			model.Enabled = core.BoolPtr(true)
			model.Key = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudshellv1.RegionSetting
			err = ibmcloudshellv1.UnmarshalRegionSetting(raw, &result)
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
