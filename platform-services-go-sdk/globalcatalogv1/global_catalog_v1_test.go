/**
 * (C) Copyright IBM Corp. 2025.
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

package globalcatalogv1_test

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
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`GlobalCatalogV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(globalCatalogService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(globalCatalogService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "https://globalcatalogv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(globalCatalogService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL":       "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{})
				Expect(globalCatalogService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := globalCatalogService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalCatalogService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalCatalogService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalCatalogService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
					URL: "https://testService/api",
				})
				Expect(globalCatalogService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := globalCatalogService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalCatalogService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalCatalogService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalCatalogService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{})
				err := globalCatalogService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := globalCatalogService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalCatalogService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalCatalogService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalCatalogService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL":       "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(globalCatalogService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalCatalogService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = globalcatalogv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListCatalogEntries(listCatalogEntriesOptions *ListCatalogEntriesOptions) - Operation response error`, func() {
		listCatalogEntriesPath := "/"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogEntriesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))
					// TODO: Add check for catalog query parameter
					// TODO: Add check for complete query parameter
					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCatalogEntries with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := new(globalcatalogv1.ListCatalogEntriesOptions)
				listCatalogEntriesOptionsModel.Account = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Include = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Q = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.SortBy = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Descending = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Languages = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Catalog = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Complete = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCatalogEntriesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listCatalogEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.ListCatalogEntries(listCatalogEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.ListCatalogEntries(listCatalogEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCatalogEntries(listCatalogEntriesOptions *ListCatalogEntriesOptions)`, func() {
		listCatalogEntriesPath := "/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogEntriesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))
					// TODO: Add check for catalog query parameter
					// TODO: Add check for complete query parameter
					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListCatalogEntries successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := new(globalcatalogv1.ListCatalogEntriesOptions)
				listCatalogEntriesOptionsModel.Account = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Include = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Q = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.SortBy = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Descending = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Languages = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Catalog = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Complete = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCatalogEntriesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listCatalogEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.ListCatalogEntriesWithContext(ctx, listCatalogEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.ListCatalogEntries(listCatalogEntriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.ListCatalogEntriesWithContext(ctx, listCatalogEntriesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogEntriesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))
					// TODO: Add check for catalog query parameter
					// TODO: Add check for complete query parameter
					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListCatalogEntries successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.ListCatalogEntries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := new(globalcatalogv1.ListCatalogEntriesOptions)
				listCatalogEntriesOptionsModel.Account = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Include = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Q = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.SortBy = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Descending = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Languages = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Catalog = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Complete = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCatalogEntriesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listCatalogEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.ListCatalogEntries(listCatalogEntriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCatalogEntries with error: Operation request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := new(globalcatalogv1.ListCatalogEntriesOptions)
				listCatalogEntriesOptionsModel.Account = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Include = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Q = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.SortBy = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Descending = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Languages = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Catalog = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Complete = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCatalogEntriesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listCatalogEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.ListCatalogEntries(listCatalogEntriesOptionsModel)
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
			It(`Invoke ListCatalogEntries successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := new(globalcatalogv1.ListCatalogEntriesOptions)
				listCatalogEntriesOptionsModel.Account = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Include = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Q = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.SortBy = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Descending = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Languages = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Catalog = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Complete = core.BoolPtr(true)
				listCatalogEntriesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCatalogEntriesOptionsModel.Limit = core.Int64Ptr(int64(50))
				listCatalogEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalCatalogService.ListCatalogEntries(listCatalogEntriesOptionsModel)
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
	Describe(`CreateCatalogEntry(createCatalogEntryOptions *CreateCatalogEntryOptions) - Operation response error`, func() {
		createCatalogEntryPath := "/"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogEntryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCatalogEntry with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsModel := new(globalcatalogv1.CreateCatalogEntryOptions)
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				createCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				createCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				createCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				createCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.CreateCatalogEntry(createCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.CreateCatalogEntry(createCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogEntry(createCatalogEntryOptions *CreateCatalogEntryOptions)`, func() {
		createCatalogEntryPath := "/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogEntryPath))
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

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateCatalogEntry successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsModel := new(globalcatalogv1.CreateCatalogEntryOptions)
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				createCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				createCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				createCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				createCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.CreateCatalogEntryWithContext(ctx, createCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.CreateCatalogEntry(createCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.CreateCatalogEntryWithContext(ctx, createCatalogEntryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogEntryPath))
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

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateCatalogEntry successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.CreateCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsModel := new(globalcatalogv1.CreateCatalogEntryOptions)
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				createCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				createCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				createCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				createCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.CreateCatalogEntry(createCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCatalogEntry with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsModel := new(globalcatalogv1.CreateCatalogEntryOptions)
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				createCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				createCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				createCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				createCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.CreateCatalogEntry(createCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCatalogEntryOptions model with no property values
				createCatalogEntryOptionsModelNew := new(globalcatalogv1.CreateCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.CreateCatalogEntry(createCatalogEntryOptionsModelNew)
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
			It(`Invoke CreateCatalogEntry successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsModel := new(globalcatalogv1.CreateCatalogEntryOptions)
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				createCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				createCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				createCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				createCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel

				// Invoke operation
				result, response, operationErr := globalCatalogService.CreateCatalogEntry(createCatalogEntryOptionsModel)
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
	Describe(`GetCatalogEntry(getCatalogEntryOptions *GetCatalogEntryOptions) - Operation response error`, func() {
		getCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogEntryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))
					// TODO: Add check for complete query parameter
					Expect(req.URL.Query()["depth"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalogEntry with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetCatalogEntryOptions model
				getCatalogEntryOptionsModel := new(globalcatalogv1.GetCatalogEntryOptions)
				getCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Include = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Languages = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Complete = core.BoolPtr(true)
				getCatalogEntryOptionsModel.Depth = core.Int64Ptr(int64(38))
				getCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.GetCatalogEntry(getCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.GetCatalogEntry(getCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogEntry(getCatalogEntryOptions *GetCatalogEntryOptions)`, func() {
		getCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogEntryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))
					// TODO: Add check for complete query parameter
					Expect(req.URL.Query()["depth"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetCatalogEntry successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the GetCatalogEntryOptions model
				getCatalogEntryOptionsModel := new(globalcatalogv1.GetCatalogEntryOptions)
				getCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Include = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Languages = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Complete = core.BoolPtr(true)
				getCatalogEntryOptionsModel.Depth = core.Int64Ptr(int64(38))
				getCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.GetCatalogEntryWithContext(ctx, getCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.GetCatalogEntry(getCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.GetCatalogEntryWithContext(ctx, getCatalogEntryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogEntryPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))
					// TODO: Add check for complete query parameter
					Expect(req.URL.Query()["depth"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetCatalogEntry successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.GetCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogEntryOptions model
				getCatalogEntryOptionsModel := new(globalcatalogv1.GetCatalogEntryOptions)
				getCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Include = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Languages = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Complete = core.BoolPtr(true)
				getCatalogEntryOptionsModel.Depth = core.Int64Ptr(int64(38))
				getCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.GetCatalogEntry(getCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCatalogEntry with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetCatalogEntryOptions model
				getCatalogEntryOptionsModel := new(globalcatalogv1.GetCatalogEntryOptions)
				getCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Include = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Languages = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Complete = core.BoolPtr(true)
				getCatalogEntryOptionsModel.Depth = core.Int64Ptr(int64(38))
				getCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.GetCatalogEntry(getCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCatalogEntryOptions model with no property values
				getCatalogEntryOptionsModelNew := new(globalcatalogv1.GetCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.GetCatalogEntry(getCatalogEntryOptionsModelNew)
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
			It(`Invoke GetCatalogEntry successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetCatalogEntryOptions model
				getCatalogEntryOptionsModel := new(globalcatalogv1.GetCatalogEntryOptions)
				getCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Include = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Languages = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Complete = core.BoolPtr(true)
				getCatalogEntryOptionsModel.Depth = core.Int64Ptr(int64(38))
				getCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalCatalogService.GetCatalogEntry(getCatalogEntryOptionsModel)
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
	Describe(`UpdateCatalogEntry(updateCatalogEntryOptions *UpdateCatalogEntryOptions) - Operation response error`, func() {
		updateCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogEntryPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["move"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCatalogEntry with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				updateCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				updateCatalogEntryOptionsModel.Images = imageModel
				updateCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Tags = []string{"testString"}
				updateCatalogEntryOptionsModel.Provider = providerModel
				updateCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				updateCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Move = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				updateCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogEntry(updateCatalogEntryOptions *UpdateCatalogEntryOptions)`, func() {
		updateCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogEntryPath))
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

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["move"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateCatalogEntry successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				updateCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				updateCatalogEntryOptionsModel.Images = imageModel
				updateCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Tags = []string{"testString"}
				updateCatalogEntryOptionsModel.Provider = providerModel
				updateCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				updateCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Move = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				updateCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.UpdateCatalogEntryWithContext(ctx, updateCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.UpdateCatalogEntryWithContext(ctx, updateCatalogEntryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogEntryPath))
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

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["move"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateCatalogEntry successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.UpdateCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				updateCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				updateCatalogEntryOptionsModel.Images = imageModel
				updateCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Tags = []string{"testString"}
				updateCatalogEntryOptionsModel.Provider = providerModel
				updateCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				updateCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Move = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				updateCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCatalogEntry with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				updateCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				updateCatalogEntryOptionsModel.Images = imageModel
				updateCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Tags = []string{"testString"}
				updateCatalogEntryOptionsModel.Provider = providerModel
				updateCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				updateCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Move = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				updateCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCatalogEntryOptions model with no property values
				updateCatalogEntryOptionsModelNew := new(globalcatalogv1.UpdateCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptionsModelNew)
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
			It(`Invoke UpdateCatalogEntry successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				updateCatalogEntryOptionsModel.OverviewUI = map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				updateCatalogEntryOptionsModel.Images = imageModel
				updateCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Tags = []string{"testString"}
				updateCatalogEntryOptionsModel.Provider = providerModel
				updateCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.URL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				updateCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Move = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				updateCatalogEntryOptionsModel.OverviewUI["foo"] = *overviewModel

				// Invoke operation
				result, response, operationErr := globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptionsModel)
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
	Describe(`DeleteCatalogEntry(deleteCatalogEntryOptions *DeleteCatalogEntryOptions)`, func() {
		deleteCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCatalogEntryPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// TODO: Add check for force query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteCatalogEntry successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := globalCatalogService.DeleteCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCatalogEntryOptions model
				deleteCatalogEntryOptionsModel := new(globalcatalogv1.DeleteCatalogEntryOptions)
				deleteCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				deleteCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				deleteCatalogEntryOptionsModel.Force = core.BoolPtr(true)
				deleteCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = globalCatalogService.DeleteCatalogEntry(deleteCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCatalogEntry with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the DeleteCatalogEntryOptions model
				deleteCatalogEntryOptionsModel := new(globalcatalogv1.DeleteCatalogEntryOptions)
				deleteCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				deleteCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				deleteCatalogEntryOptionsModel.Force = core.BoolPtr(true)
				deleteCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := globalCatalogService.DeleteCatalogEntry(deleteCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCatalogEntryOptions model with no property values
				deleteCatalogEntryOptionsModelNew := new(globalcatalogv1.DeleteCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = globalCatalogService.DeleteCatalogEntry(deleteCatalogEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetChildObjects(getChildObjectsOptions *GetChildObjectsOptions) - Operation response error`, func() {
		getChildObjectsPath := "/testString/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChildObjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))
					// TODO: Add check for complete query parameter
					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetChildObjects with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetChildObjectsOptions model
				getChildObjectsOptionsModel := new(globalcatalogv1.GetChildObjectsOptions)
				getChildObjectsOptionsModel.ID = core.StringPtr("testString")
				getChildObjectsOptionsModel.Kind = core.StringPtr("testString")
				getChildObjectsOptionsModel.Account = core.StringPtr("testString")
				getChildObjectsOptionsModel.Include = core.StringPtr("testString")
				getChildObjectsOptionsModel.Q = core.StringPtr("testString")
				getChildObjectsOptionsModel.SortBy = core.StringPtr("testString")
				getChildObjectsOptionsModel.Descending = core.StringPtr("testString")
				getChildObjectsOptionsModel.Languages = core.StringPtr("testString")
				getChildObjectsOptionsModel.Complete = core.BoolPtr(true)
				getChildObjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getChildObjectsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getChildObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.GetChildObjects(getChildObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.GetChildObjects(getChildObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetChildObjects(getChildObjectsOptions *GetChildObjectsOptions)`, func() {
		getChildObjectsPath := "/testString/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChildObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))
					// TODO: Add check for complete query parameter
					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetChildObjects successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the GetChildObjectsOptions model
				getChildObjectsOptionsModel := new(globalcatalogv1.GetChildObjectsOptions)
				getChildObjectsOptionsModel.ID = core.StringPtr("testString")
				getChildObjectsOptionsModel.Kind = core.StringPtr("testString")
				getChildObjectsOptionsModel.Account = core.StringPtr("testString")
				getChildObjectsOptionsModel.Include = core.StringPtr("testString")
				getChildObjectsOptionsModel.Q = core.StringPtr("testString")
				getChildObjectsOptionsModel.SortBy = core.StringPtr("testString")
				getChildObjectsOptionsModel.Descending = core.StringPtr("testString")
				getChildObjectsOptionsModel.Languages = core.StringPtr("testString")
				getChildObjectsOptionsModel.Complete = core.BoolPtr(true)
				getChildObjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getChildObjectsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getChildObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.GetChildObjectsWithContext(ctx, getChildObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.GetChildObjects(getChildObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.GetChildObjectsWithContext(ctx, getChildObjectsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getChildObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))
					// TODO: Add check for complete query parameter
					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"name": "Name", "kind": "service", "overview_ui": {"mapKey": {"display_name": "DisplayName", "long_description": "LongDescription", "description": "Description", "featured_description": "FeaturedDescription"}}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "url": "URL", "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}, "crn_mask": "CRNMask", "parameters": {"anyKey": "anyValue"}, "user_defined_service": {"anyKey": "anyValue"}, "extension": {"anyKey": "anyValue"}, "paid_only": true, "custom_create_page_hybrid_enabled": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "provision_type": "ProvisionType", "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {"mapKey": {"bullets": [{"title": "Title", "description": "Description", "icon": "Icon", "quantity": 8}], "media": [{"caption": "Caption", "thumbnail_url": "ThumbnailURL", "type": "Type", "URL": "URL", "source": [{"type": "Type", "url": "URL"}]}], "not_creatable_msg": "NotCreatableMsg", "not_creatable__robot_msg": "NotCreatableRobotMsg", "deprecation_warning": "DeprecationWarning", "popup_warning_message": "PopupWarningMessage", "instruction": "Instruction"}}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "APIURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00.000Z", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": 12, "responsiveness": 14, "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorAPI", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"anyKey": "anyValue"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCRN", "service_crn": "ServiceCRN", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "GUID"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "CatalogCRN", "children_url": "ChildrenURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00.000Z", "updated": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetChildObjects successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.GetChildObjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetChildObjectsOptions model
				getChildObjectsOptionsModel := new(globalcatalogv1.GetChildObjectsOptions)
				getChildObjectsOptionsModel.ID = core.StringPtr("testString")
				getChildObjectsOptionsModel.Kind = core.StringPtr("testString")
				getChildObjectsOptionsModel.Account = core.StringPtr("testString")
				getChildObjectsOptionsModel.Include = core.StringPtr("testString")
				getChildObjectsOptionsModel.Q = core.StringPtr("testString")
				getChildObjectsOptionsModel.SortBy = core.StringPtr("testString")
				getChildObjectsOptionsModel.Descending = core.StringPtr("testString")
				getChildObjectsOptionsModel.Languages = core.StringPtr("testString")
				getChildObjectsOptionsModel.Complete = core.BoolPtr(true)
				getChildObjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getChildObjectsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getChildObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.GetChildObjects(getChildObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetChildObjects with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetChildObjectsOptions model
				getChildObjectsOptionsModel := new(globalcatalogv1.GetChildObjectsOptions)
				getChildObjectsOptionsModel.ID = core.StringPtr("testString")
				getChildObjectsOptionsModel.Kind = core.StringPtr("testString")
				getChildObjectsOptionsModel.Account = core.StringPtr("testString")
				getChildObjectsOptionsModel.Include = core.StringPtr("testString")
				getChildObjectsOptionsModel.Q = core.StringPtr("testString")
				getChildObjectsOptionsModel.SortBy = core.StringPtr("testString")
				getChildObjectsOptionsModel.Descending = core.StringPtr("testString")
				getChildObjectsOptionsModel.Languages = core.StringPtr("testString")
				getChildObjectsOptionsModel.Complete = core.BoolPtr(true)
				getChildObjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getChildObjectsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getChildObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.GetChildObjects(getChildObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetChildObjectsOptions model with no property values
				getChildObjectsOptionsModelNew := new(globalcatalogv1.GetChildObjectsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.GetChildObjects(getChildObjectsOptionsModelNew)
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
			It(`Invoke GetChildObjects successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetChildObjectsOptions model
				getChildObjectsOptionsModel := new(globalcatalogv1.GetChildObjectsOptions)
				getChildObjectsOptionsModel.ID = core.StringPtr("testString")
				getChildObjectsOptionsModel.Kind = core.StringPtr("testString")
				getChildObjectsOptionsModel.Account = core.StringPtr("testString")
				getChildObjectsOptionsModel.Include = core.StringPtr("testString")
				getChildObjectsOptionsModel.Q = core.StringPtr("testString")
				getChildObjectsOptionsModel.SortBy = core.StringPtr("testString")
				getChildObjectsOptionsModel.Descending = core.StringPtr("testString")
				getChildObjectsOptionsModel.Languages = core.StringPtr("testString")
				getChildObjectsOptionsModel.Complete = core.BoolPtr(true)
				getChildObjectsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getChildObjectsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getChildObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalCatalogService.GetChildObjects(getChildObjectsOptionsModel)
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
	Describe(`RestoreCatalogEntry(restoreCatalogEntryOptions *RestoreCatalogEntryOptions)`, func() {
		restoreCatalogEntryPath := "/testString/restore"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreCatalogEntryPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke RestoreCatalogEntry successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := globalCatalogService.RestoreCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RestoreCatalogEntryOptions model
				restoreCatalogEntryOptionsModel := new(globalcatalogv1.RestoreCatalogEntryOptions)
				restoreCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				restoreCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				restoreCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = globalCatalogService.RestoreCatalogEntry(restoreCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RestoreCatalogEntry with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the RestoreCatalogEntryOptions model
				restoreCatalogEntryOptionsModel := new(globalcatalogv1.RestoreCatalogEntryOptions)
				restoreCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				restoreCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				restoreCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := globalCatalogService.RestoreCatalogEntry(restoreCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RestoreCatalogEntryOptions model with no property values
				restoreCatalogEntryOptionsModelNew := new(globalcatalogv1.RestoreCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = globalCatalogService.RestoreCatalogEntry(restoreCatalogEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVisibility(getVisibilityOptions *GetVisibilityOptions) - Operation response error`, func() {
		getVisibilityPath := "/testString/visibility"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVisibilityPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVisibility with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetVisibilityOptions model
				getVisibilityOptionsModel := new(globalcatalogv1.GetVisibilityOptions)
				getVisibilityOptionsModel.ID = core.StringPtr("testString")
				getVisibilityOptionsModel.Account = core.StringPtr("testString")
				getVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.GetVisibility(getVisibilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.GetVisibility(getVisibilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVisibility(getVisibilityOptions *GetVisibilityOptions)`, func() {
		getVisibilityPath := "/testString/visibility"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVisibilityPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"restrictions": "Restrictions", "owner": "Owner", "extendable": true, "include": {"accounts": {"_accountid_": "Accountid"}}, "exclude": {"accounts": {"_accountid_": "Accountid"}}, "approved": true}`)
				}))
			})
			It(`Invoke GetVisibility successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the GetVisibilityOptions model
				getVisibilityOptionsModel := new(globalcatalogv1.GetVisibilityOptions)
				getVisibilityOptionsModel.ID = core.StringPtr("testString")
				getVisibilityOptionsModel.Account = core.StringPtr("testString")
				getVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.GetVisibilityWithContext(ctx, getVisibilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.GetVisibility(getVisibilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.GetVisibilityWithContext(ctx, getVisibilityOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getVisibilityPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"restrictions": "Restrictions", "owner": "Owner", "extendable": true, "include": {"accounts": {"_accountid_": "Accountid"}}, "exclude": {"accounts": {"_accountid_": "Accountid"}}, "approved": true}`)
				}))
			})
			It(`Invoke GetVisibility successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.GetVisibility(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVisibilityOptions model
				getVisibilityOptionsModel := new(globalcatalogv1.GetVisibilityOptions)
				getVisibilityOptionsModel.ID = core.StringPtr("testString")
				getVisibilityOptionsModel.Account = core.StringPtr("testString")
				getVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.GetVisibility(getVisibilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVisibility with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetVisibilityOptions model
				getVisibilityOptionsModel := new(globalcatalogv1.GetVisibilityOptions)
				getVisibilityOptionsModel.ID = core.StringPtr("testString")
				getVisibilityOptionsModel.Account = core.StringPtr("testString")
				getVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.GetVisibility(getVisibilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVisibilityOptions model with no property values
				getVisibilityOptionsModelNew := new(globalcatalogv1.GetVisibilityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.GetVisibility(getVisibilityOptionsModelNew)
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
			It(`Invoke GetVisibility successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetVisibilityOptions model
				getVisibilityOptionsModel := new(globalcatalogv1.GetVisibilityOptions)
				getVisibilityOptionsModel.ID = core.StringPtr("testString")
				getVisibilityOptionsModel.Account = core.StringPtr("testString")
				getVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalCatalogService.GetVisibility(getVisibilityOptionsModel)
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
	Describe(`UpdateVisibility(updateVisibilityOptions *UpdateVisibilityOptions)`, func() {
		updateVisibilityPath := "/testString/visibility"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVisibilityPath))
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

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateVisibility successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := globalCatalogService.UpdateVisibility(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the VisibilityDetailAccounts model
				visibilityDetailAccountsModel := new(globalcatalogv1.VisibilityDetailAccounts)
				visibilityDetailAccountsModel.Accountid = core.StringPtr("testString")

				// Construct an instance of the VisibilityDetail model
				visibilityDetailModel := new(globalcatalogv1.VisibilityDetail)
				visibilityDetailModel.Accounts = visibilityDetailAccountsModel

				// Construct an instance of the UpdateVisibilityOptions model
				updateVisibilityOptionsModel := new(globalcatalogv1.UpdateVisibilityOptions)
				updateVisibilityOptionsModel.ID = core.StringPtr("testString")
				updateVisibilityOptionsModel.Restrictions = core.StringPtr("testString")
				updateVisibilityOptionsModel.Extendable = core.BoolPtr(true)
				updateVisibilityOptionsModel.Include = visibilityDetailModel
				updateVisibilityOptionsModel.Exclude = visibilityDetailModel
				updateVisibilityOptionsModel.Account = core.StringPtr("testString")
				updateVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = globalCatalogService.UpdateVisibility(updateVisibilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateVisibility with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the VisibilityDetailAccounts model
				visibilityDetailAccountsModel := new(globalcatalogv1.VisibilityDetailAccounts)
				visibilityDetailAccountsModel.Accountid = core.StringPtr("testString")

				// Construct an instance of the VisibilityDetail model
				visibilityDetailModel := new(globalcatalogv1.VisibilityDetail)
				visibilityDetailModel.Accounts = visibilityDetailAccountsModel

				// Construct an instance of the UpdateVisibilityOptions model
				updateVisibilityOptionsModel := new(globalcatalogv1.UpdateVisibilityOptions)
				updateVisibilityOptionsModel.ID = core.StringPtr("testString")
				updateVisibilityOptionsModel.Restrictions = core.StringPtr("testString")
				updateVisibilityOptionsModel.Extendable = core.BoolPtr(true)
				updateVisibilityOptionsModel.Include = visibilityDetailModel
				updateVisibilityOptionsModel.Exclude = visibilityDetailModel
				updateVisibilityOptionsModel.Account = core.StringPtr("testString")
				updateVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := globalCatalogService.UpdateVisibility(updateVisibilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateVisibilityOptions model with no property values
				updateVisibilityOptionsModelNew := new(globalcatalogv1.UpdateVisibilityOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = globalCatalogService.UpdateVisibility(updateVisibilityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPricing(getPricingOptions *GetPricingOptions) - Operation response error`, func() {
		getPricingPath := "/testString/pricing"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPricingPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deployment_region"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPricing with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetPricingOptions model
				getPricingOptionsModel := new(globalcatalogv1.GetPricingOptions)
				getPricingOptionsModel.ID = core.StringPtr("testString")
				getPricingOptionsModel.Account = core.StringPtr("testString")
				getPricingOptionsModel.DeploymentRegion = core.StringPtr("testString")
				getPricingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.GetPricing(getPricingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.GetPricing(getPricingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPricing(getPricingOptions *GetPricingOptions)`, func() {
		getPricingPath := "/testString/pricing"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPricingPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deployment_region"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}`)
				}))
			})
			It(`Invoke GetPricing successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the GetPricingOptions model
				getPricingOptionsModel := new(globalcatalogv1.GetPricingOptions)
				getPricingOptionsModel.ID = core.StringPtr("testString")
				getPricingOptionsModel.Account = core.StringPtr("testString")
				getPricingOptionsModel.DeploymentRegion = core.StringPtr("testString")
				getPricingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.GetPricingWithContext(ctx, getPricingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.GetPricing(getPricingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.GetPricingWithContext(ctx, getPricingOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPricingPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deployment_region"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}`)
				}))
			})
			It(`Invoke GetPricing successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.GetPricing(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPricingOptions model
				getPricingOptionsModel := new(globalcatalogv1.GetPricingOptions)
				getPricingOptionsModel.ID = core.StringPtr("testString")
				getPricingOptionsModel.Account = core.StringPtr("testString")
				getPricingOptionsModel.DeploymentRegion = core.StringPtr("testString")
				getPricingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.GetPricing(getPricingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPricing with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetPricingOptions model
				getPricingOptionsModel := new(globalcatalogv1.GetPricingOptions)
				getPricingOptionsModel.ID = core.StringPtr("testString")
				getPricingOptionsModel.Account = core.StringPtr("testString")
				getPricingOptionsModel.DeploymentRegion = core.StringPtr("testString")
				getPricingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.GetPricing(getPricingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPricingOptions model with no property values
				getPricingOptionsModelNew := new(globalcatalogv1.GetPricingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.GetPricing(getPricingOptionsModelNew)
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
			It(`Invoke GetPricing successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetPricingOptions model
				getPricingOptionsModel := new(globalcatalogv1.GetPricingOptions)
				getPricingOptionsModel.ID = core.StringPtr("testString")
				getPricingOptionsModel.Account = core.StringPtr("testString")
				getPricingOptionsModel.DeploymentRegion = core.StringPtr("testString")
				getPricingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalCatalogService.GetPricing(getPricingOptionsModel)
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
	Describe(`GetPricingDeployments(getPricingDeploymentsOptions *GetPricingDeploymentsOptions) - Operation response error`, func() {
		getPricingDeploymentsPath := "/testString/pricing/deployment"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPricingDeploymentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPricingDeployments with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetPricingDeploymentsOptions model
				getPricingDeploymentsOptionsModel := new(globalcatalogv1.GetPricingDeploymentsOptions)
				getPricingDeploymentsOptionsModel.ID = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Account = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.GetPricingDeployments(getPricingDeploymentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.GetPricingDeployments(getPricingDeploymentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPricingDeployments(getPricingDeploymentsOptions *GetPricingDeploymentsOptions)`, func() {
		getPricingDeploymentsPath := "/testString/pricing/deployment"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPricingDeploymentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}]}`)
				}))
			})
			It(`Invoke GetPricingDeployments successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the GetPricingDeploymentsOptions model
				getPricingDeploymentsOptionsModel := new(globalcatalogv1.GetPricingDeploymentsOptions)
				getPricingDeploymentsOptionsModel.ID = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Account = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.GetPricingDeploymentsWithContext(ctx, getPricingDeploymentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.GetPricingDeployments(getPricingDeploymentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.GetPricingDeploymentsWithContext(ctx, getPricingDeploymentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPricingDeploymentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"deployment_id": "DeploymentID", "deployment_location": "DeploymentLocation", "deployment_region": "DeploymentRegion", "deployment_location_no_price_available": true, "type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": 18, "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "price": 5}]}], "additional_properties": {"anyKey": "anyValue"}}], "deployment_regions": ["DeploymentRegions"], "effective_from": "2019-01-01T12:00:00.000Z", "effective_until": "2019-01-01T12:00:00.000Z", "require_login": true, "pricing_catalog_url": "PricingCatalogURL", "sales_avenue": ["SalesAvenue"]}]}`)
				}))
			})
			It(`Invoke GetPricingDeployments successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.GetPricingDeployments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPricingDeploymentsOptions model
				getPricingDeploymentsOptionsModel := new(globalcatalogv1.GetPricingDeploymentsOptions)
				getPricingDeploymentsOptionsModel.ID = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Account = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.GetPricingDeployments(getPricingDeploymentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPricingDeployments with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetPricingDeploymentsOptions model
				getPricingDeploymentsOptionsModel := new(globalcatalogv1.GetPricingDeploymentsOptions)
				getPricingDeploymentsOptionsModel.ID = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Account = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.GetPricingDeployments(getPricingDeploymentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPricingDeploymentsOptions model with no property values
				getPricingDeploymentsOptionsModelNew := new(globalcatalogv1.GetPricingDeploymentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.GetPricingDeployments(getPricingDeploymentsOptionsModelNew)
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
			It(`Invoke GetPricingDeployments successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetPricingDeploymentsOptions model
				getPricingDeploymentsOptionsModel := new(globalcatalogv1.GetPricingDeploymentsOptions)
				getPricingDeploymentsOptionsModel.ID = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Account = core.StringPtr("testString")
				getPricingDeploymentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalCatalogService.GetPricingDeployments(getPricingDeploymentsOptionsModel)
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
	Describe(`GetAuditLogs(getAuditLogsOptions *GetAuditLogsOptions) - Operation response error`, func() {
		getAuditLogsPath := "/testString/logs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAuditLogsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["ascending"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["startat"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAuditLogs with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetAuditLogsOptions model
				getAuditLogsOptionsModel := new(globalcatalogv1.GetAuditLogsOptions)
				getAuditLogsOptionsModel.ID = core.StringPtr("testString")
				getAuditLogsOptionsModel.Account = core.StringPtr("testString")
				getAuditLogsOptionsModel.Ascending = core.StringPtr("false")
				getAuditLogsOptionsModel.Startat = core.StringPtr("testString")
				getAuditLogsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getAuditLogsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getAuditLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.GetAuditLogs(getAuditLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.GetAuditLogs(getAuditLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAuditLogs(getAuditLogsOptions *GetAuditLogsOptions)`, func() {
		getAuditLogsPath := "/testString/logs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAuditLogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["ascending"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["startat"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "effective": {"restrictions": "Restrictions", "owner": "Owner", "extendable": true, "include": {"accounts": {"_accountid_": "Accountid"}}, "exclude": {"accounts": {"_accountid_": "Accountid"}}, "approved": true}, "time": "2019-01-01T12:00:00.000Z", "who_id": "WhoID", "who_name": "WhoName", "who_email": "WhoEmail", "instance": "Instance", "gid": "Gid", "type": "Type", "message": "Message", "data": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke GetAuditLogs successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the GetAuditLogsOptions model
				getAuditLogsOptionsModel := new(globalcatalogv1.GetAuditLogsOptions)
				getAuditLogsOptionsModel.ID = core.StringPtr("testString")
				getAuditLogsOptionsModel.Account = core.StringPtr("testString")
				getAuditLogsOptionsModel.Ascending = core.StringPtr("false")
				getAuditLogsOptionsModel.Startat = core.StringPtr("testString")
				getAuditLogsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getAuditLogsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getAuditLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.GetAuditLogsWithContext(ctx, getAuditLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.GetAuditLogs(getAuditLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.GetAuditLogsWithContext(ctx, getAuditLogsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAuditLogsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["ascending"]).To(Equal([]string{"false"}))
					Expect(req.URL.Query()["startat"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "effective": {"restrictions": "Restrictions", "owner": "Owner", "extendable": true, "include": {"accounts": {"_accountid_": "Accountid"}}, "exclude": {"accounts": {"_accountid_": "Accountid"}}, "approved": true}, "time": "2019-01-01T12:00:00.000Z", "who_id": "WhoID", "who_name": "WhoName", "who_email": "WhoEmail", "instance": "Instance", "gid": "Gid", "type": "Type", "message": "Message", "data": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke GetAuditLogs successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.GetAuditLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAuditLogsOptions model
				getAuditLogsOptionsModel := new(globalcatalogv1.GetAuditLogsOptions)
				getAuditLogsOptionsModel.ID = core.StringPtr("testString")
				getAuditLogsOptionsModel.Account = core.StringPtr("testString")
				getAuditLogsOptionsModel.Ascending = core.StringPtr("false")
				getAuditLogsOptionsModel.Startat = core.StringPtr("testString")
				getAuditLogsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getAuditLogsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getAuditLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.GetAuditLogs(getAuditLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAuditLogs with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetAuditLogsOptions model
				getAuditLogsOptionsModel := new(globalcatalogv1.GetAuditLogsOptions)
				getAuditLogsOptionsModel.ID = core.StringPtr("testString")
				getAuditLogsOptionsModel.Account = core.StringPtr("testString")
				getAuditLogsOptionsModel.Ascending = core.StringPtr("false")
				getAuditLogsOptionsModel.Startat = core.StringPtr("testString")
				getAuditLogsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getAuditLogsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getAuditLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.GetAuditLogs(getAuditLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAuditLogsOptions model with no property values
				getAuditLogsOptionsModelNew := new(globalcatalogv1.GetAuditLogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.GetAuditLogs(getAuditLogsOptionsModelNew)
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
			It(`Invoke GetAuditLogs successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetAuditLogsOptions model
				getAuditLogsOptionsModel := new(globalcatalogv1.GetAuditLogsOptions)
				getAuditLogsOptionsModel.ID = core.StringPtr("testString")
				getAuditLogsOptionsModel.Account = core.StringPtr("testString")
				getAuditLogsOptionsModel.Ascending = core.StringPtr("false")
				getAuditLogsOptionsModel.Startat = core.StringPtr("testString")
				getAuditLogsOptionsModel.Offset = core.Int64Ptr(int64(0))
				getAuditLogsOptionsModel.Limit = core.Int64Ptr(int64(50))
				getAuditLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalCatalogService.GetAuditLogs(getAuditLogsOptionsModel)
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
	Describe(`ListArtifacts(listArtifactsOptions *ListArtifactsOptions) - Operation response error`, func() {
		listArtifactsPath := "/testString/artifacts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listArtifactsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListArtifacts with error: Operation response processing error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the ListArtifactsOptions model
				listArtifactsOptionsModel := new(globalcatalogv1.ListArtifactsOptions)
				listArtifactsOptionsModel.ObjectID = core.StringPtr("testString")
				listArtifactsOptionsModel.Account = core.StringPtr("testString")
				listArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalCatalogService.ListArtifacts(listArtifactsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalCatalogService.EnableRetries(0, 0)
				result, response, operationErr = globalCatalogService.ListArtifacts(listArtifactsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListArtifacts(listArtifactsOptions *ListArtifactsOptions)`, func() {
		listArtifactsPath := "/testString/artifacts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listArtifactsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "resources": [{"name": "Name", "updated": "2019-01-01T12:00:00.000Z", "url": "URL", "etag": "Etag", "size": 4}]}`)
				}))
			})
			It(`Invoke ListArtifacts successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the ListArtifactsOptions model
				listArtifactsOptionsModel := new(globalcatalogv1.ListArtifactsOptions)
				listArtifactsOptionsModel.ObjectID = core.StringPtr("testString")
				listArtifactsOptionsModel.Account = core.StringPtr("testString")
				listArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.ListArtifactsWithContext(ctx, listArtifactsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.ListArtifacts(listArtifactsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.ListArtifactsWithContext(ctx, listArtifactsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listArtifactsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "resources": [{"name": "Name", "updated": "2019-01-01T12:00:00.000Z", "url": "URL", "etag": "Etag", "size": 4}]}`)
				}))
			})
			It(`Invoke ListArtifacts successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.ListArtifacts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListArtifactsOptions model
				listArtifactsOptionsModel := new(globalcatalogv1.ListArtifactsOptions)
				listArtifactsOptionsModel.ObjectID = core.StringPtr("testString")
				listArtifactsOptionsModel.Account = core.StringPtr("testString")
				listArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.ListArtifacts(listArtifactsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListArtifacts with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the ListArtifactsOptions model
				listArtifactsOptionsModel := new(globalcatalogv1.ListArtifactsOptions)
				listArtifactsOptionsModel.ObjectID = core.StringPtr("testString")
				listArtifactsOptionsModel.Account = core.StringPtr("testString")
				listArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.ListArtifacts(listArtifactsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListArtifactsOptions model with no property values
				listArtifactsOptionsModelNew := new(globalcatalogv1.ListArtifactsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.ListArtifacts(listArtifactsOptionsModelNew)
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
			It(`Invoke ListArtifacts successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the ListArtifactsOptions model
				listArtifactsOptionsModel := new(globalcatalogv1.ListArtifactsOptions)
				listArtifactsOptionsModel.ObjectID = core.StringPtr("testString")
				listArtifactsOptionsModel.Account = core.StringPtr("testString")
				listArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalCatalogService.ListArtifacts(listArtifactsOptionsModel)
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
	Describe(`GetArtifact(getArtifactOptions *GetArtifactOptions)`, func() {
		getArtifactPath := "/testString/artifacts/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getArtifactPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetArtifact successfully with retries`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())
				globalCatalogService.EnableRetries(0, 0)

				// Construct an instance of the GetArtifactOptions model
				getArtifactOptionsModel := new(globalcatalogv1.GetArtifactOptions)
				getArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				getArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				getArtifactOptionsModel.Accept = core.StringPtr("testString")
				getArtifactOptionsModel.Account = core.StringPtr("testString")
				getArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := globalCatalogService.GetArtifactWithContext(ctx, getArtifactOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				globalCatalogService.DisableRetries()
				result, response, operationErr := globalCatalogService.GetArtifact(getArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = globalCatalogService.GetArtifactWithContext(ctx, getArtifactOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getArtifactPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetArtifact successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalCatalogService.GetArtifact(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetArtifactOptions model
				getArtifactOptionsModel := new(globalcatalogv1.GetArtifactOptions)
				getArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				getArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				getArtifactOptionsModel.Accept = core.StringPtr("testString")
				getArtifactOptionsModel.Account = core.StringPtr("testString")
				getArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalCatalogService.GetArtifact(getArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetArtifact with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetArtifactOptions model
				getArtifactOptionsModel := new(globalcatalogv1.GetArtifactOptions)
				getArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				getArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				getArtifactOptionsModel.Accept = core.StringPtr("testString")
				getArtifactOptionsModel.Account = core.StringPtr("testString")
				getArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalCatalogService.GetArtifact(getArtifactOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetArtifactOptions model with no property values
				getArtifactOptionsModelNew := new(globalcatalogv1.GetArtifactOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalCatalogService.GetArtifact(getArtifactOptionsModelNew)
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
			It(`Invoke GetArtifact successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the GetArtifactOptions model
				getArtifactOptionsModel := new(globalcatalogv1.GetArtifactOptions)
				getArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				getArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				getArtifactOptionsModel.Accept = core.StringPtr("testString")
				getArtifactOptionsModel.Account = core.StringPtr("testString")
				getArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := globalCatalogService.GetArtifact(getArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := io.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UploadArtifact(uploadArtifactOptions *UploadArtifactOptions)`, func() {
		uploadArtifactPath := "/testString/artifacts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadArtifactPath))
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

					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UploadArtifact successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := globalCatalogService.UploadArtifact(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UploadArtifactOptions model
				uploadArtifactOptionsModel := new(globalcatalogv1.UploadArtifactOptions)
				uploadArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				uploadArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				uploadArtifactOptionsModel.Artifact = CreateMockReader("This is a mock file.")
				uploadArtifactOptionsModel.ContentType = core.StringPtr("testString")
				uploadArtifactOptionsModel.Account = core.StringPtr("testString")
				uploadArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = globalCatalogService.UploadArtifact(uploadArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UploadArtifact with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the UploadArtifactOptions model
				uploadArtifactOptionsModel := new(globalcatalogv1.UploadArtifactOptions)
				uploadArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				uploadArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				uploadArtifactOptionsModel.Artifact = CreateMockReader("This is a mock file.")
				uploadArtifactOptionsModel.ContentType = core.StringPtr("testString")
				uploadArtifactOptionsModel.Account = core.StringPtr("testString")
				uploadArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := globalCatalogService.UploadArtifact(uploadArtifactOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UploadArtifactOptions model with no property values
				uploadArtifactOptionsModelNew := new(globalcatalogv1.UploadArtifactOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = globalCatalogService.UploadArtifact(uploadArtifactOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteArtifact(deleteArtifactOptions *DeleteArtifactOptions)`, func() {
		deleteArtifactPath := "/testString/artifacts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteArtifactPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteArtifact successfully`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := globalCatalogService.DeleteArtifact(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteArtifactOptions model
				deleteArtifactOptionsModel := new(globalcatalogv1.DeleteArtifactOptions)
				deleteArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				deleteArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				deleteArtifactOptionsModel.Account = core.StringPtr("testString")
				deleteArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = globalCatalogService.DeleteArtifact(deleteArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteArtifact with error: Operation validation and request error`, func() {
				globalCatalogService, serviceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalCatalogService).ToNot(BeNil())

				// Construct an instance of the DeleteArtifactOptions model
				deleteArtifactOptionsModel := new(globalcatalogv1.DeleteArtifactOptions)
				deleteArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				deleteArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				deleteArtifactOptionsModel.Account = core.StringPtr("testString")
				deleteArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalCatalogService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := globalCatalogService.DeleteArtifact(deleteArtifactOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteArtifactOptions model with no property values
				deleteArtifactOptionsModelNew := new(globalcatalogv1.DeleteArtifactOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = globalCatalogService.DeleteArtifact(deleteArtifactOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			globalCatalogService, _ := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL:           "http://globalcatalogv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateCatalogEntryOptions successfully`, func() {
				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				Expect(overviewModel).ToNot(BeNil())
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")
				Expect(overviewModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.FeaturedDescription).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				Expect(imageModel).ToNot(BeNil())
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")
				Expect(imageModel.Image).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.SmallImage).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.MediumImage).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.FeatureImage).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				Expect(providerModel).ToNot(BeNil())
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")
				Expect(providerModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.SupportEmail).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Phone).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				Expect(cfMetaDataModel).ToNot(BeNil())
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)
				Expect(cfMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.IamCompatible).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.UniqueAPIKey).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Provisionable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.AsyncProvisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.AsyncUnprovisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Requires).To(Equal([]string{"testString"}))
				Expect(cfMetaDataModel.PlanUpdateable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.State).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.ServiceCheckEnabled).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.TestCheckInterval).To(Equal(core.Int64Ptr(int64(38))))
				Expect(cfMetaDataModel.ServiceKeySupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.CfGUID).To(Equal(map[string]string{"key1": "testString"}))
				Expect(cfMetaDataModel.CRNMask).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.Parameters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(cfMetaDataModel.UserDefinedService).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(cfMetaDataModel.Extension).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(cfMetaDataModel.PaidOnly).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.CustomCreatePageHybridEnabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				Expect(planMetaDataModel).ToNot(BeNil())
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				Expect(planMetaDataModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.Reservable).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AllowInternalUsers).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AsyncProvisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AsyncUnprovisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.ProvisionType).To(Equal(core.StringPtr("testString")))
				Expect(planMetaDataModel.TestCheckInterval).To(Equal(core.Int64Ptr(int64(38))))
				Expect(planMetaDataModel.SingleScopeInstance).To(Equal(core.StringPtr("testString")))
				Expect(planMetaDataModel.ServiceCheckEnabled).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.CfGUID).To(Equal(map[string]string{"key1": "testString"}))

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				Expect(aliasMetaDataModel).ToNot(BeNil())
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")
				Expect(aliasMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(aliasMetaDataModel.PlanID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				Expect(sourceMetaDataModel).ToNot(BeNil())
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")
				Expect(sourceMetaDataModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(sourceMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(sourceMetaDataModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				Expect(templateMetaDataModel).ToNot(BeNil())
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}
				Expect(templateMetaDataModel.Services).To(Equal([]string{"testString"}))
				Expect(templateMetaDataModel.DefaultMemory).To(Equal(core.Int64Ptr(int64(38))))
				Expect(templateMetaDataModel.StartCmd).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.Source).To(Equal(sourceMetaDataModel))
				Expect(templateMetaDataModel.RuntimeCatalogID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.CfRuntimeID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.ExecutableFile).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.Buildpack).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.EnvironmentVariables).To(Equal(map[string]string{"key1": "testString"}))

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				Expect(bulletsModel).ToNot(BeNil())
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))
				Expect(bulletsModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Icon).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Quantity).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				Expect(uiMediaSourceMetaDataModel).ToNot(BeNil())
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")
				Expect(uiMediaSourceMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(uiMediaSourceMetaDataModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				Expect(uiMetaMediaModel).ToNot(BeNil())
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}
				Expect(uiMetaMediaModel.Caption).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.ThumbnailURL).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.Source).To(Equal([]globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}))

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				Expect(stringsModel).ToNot(BeNil())
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")
				Expect(stringsModel.Bullets).To(Equal([]globalcatalogv1.Bullets{*bulletsModel}))
				Expect(stringsModel.Media).To(Equal([]globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}))
				Expect(stringsModel.NotCreatableMsg).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.NotCreatableRobotMsg).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.DeprecationWarning).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.PopupWarningMessage).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.Instruction).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				Expect(urlsModel).ToNot(BeNil())
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")
				Expect(urlsModel.DocURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.InstructionsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.APIURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CreateURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.SdkDownloadURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.TermsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CustomCreatePageURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CatalogDetailsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.DeprecationDocURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.DashboardURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.RegistrationURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.Apidocsurl).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				Expect(uiMetaDataModel).ToNot(BeNil())
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel
				Expect(uiMetaDataModel.Urls).To(Equal(urlsModel))
				Expect(uiMetaDataModel.EmbeddableDashboard).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaDataModel.EmbeddableDashboardFullWidth).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.NavigationOrder).To(Equal([]string{"testString"}))
				Expect(uiMetaDataModel.NotCreatable).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.PrimaryOfferingID).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaDataModel.AccessibleDuringProvision).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.SideBySideIndex).To(Equal(core.Int64Ptr(int64(38))))
				Expect(uiMetaDataModel.EndOfServiceTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(uiMetaDataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.HideLiteMetering).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.NoUpgradeNextStep).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.Strings["foo"]).To(Equal(*stringsModel))

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				Expect(drMetaDataModel).ToNot(BeNil())
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")
				Expect(drMetaDataModel.Dr).To(Equal(core.BoolPtr(true)))
				Expect(drMetaDataModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				Expect(slaMetaDataModel).ToNot(BeNil())
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel
				Expect(slaMetaDataModel.Terms).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Tenancy).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Provisioning).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(slaMetaDataModel.Responsiveness).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(slaMetaDataModel.Dr).To(Equal(drMetaDataModel))

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				Expect(callbacksModel).ToNot(BeNil())
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}
				Expect(callbacksModel.ControllerURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.BrokerURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.BrokerProxyURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDataURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDetailTabURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDetailTabExtURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ServiceMonitorAPI).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ServiceMonitorApp).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.APIEndpoint).To(Equal(map[string]string{"key1": "testString"}))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				Expect(priceModel).ToNot(BeNil())
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))
				Expect(priceModel.QuantityTier).To(Equal(core.Int64Ptr(int64(38))))
				Expect(priceModel.Price).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				Expect(amountModel).ToNot(BeNil())
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}
				Expect(amountModel.Country).To(Equal(core.StringPtr("testString")))
				Expect(amountModel.Currency).To(Equal(core.StringPtr("testString")))
				Expect(amountModel.Prices).To(Equal([]globalcatalogv1.Price{*priceModel}))

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				Expect(startingPriceModel).ToNot(BeNil())
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}
				Expect(startingPriceModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.DeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.Unit).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.Amount).To(Equal([]globalcatalogv1.Amount{*amountModel}))

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				Expect(pricingSetModel).ToNot(BeNil())
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel
				Expect(pricingSetModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(pricingSetModel.Origin).To(Equal(core.StringPtr("testString")))
				Expect(pricingSetModel.StartingPrice).To(Equal(startingPriceModel))

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				Expect(brokerModel).ToNot(BeNil())
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")
				Expect(brokerModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(brokerModel.GUID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				Expect(deploymentBaseModel).ToNot(BeNil())
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")
				Expect(deploymentBaseModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.LocationURL).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.OriginalLocation).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.TargetCRN).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.ServiceCRN).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.MccpID).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.Broker).To(Equal(brokerModel))
				Expect(deploymentBaseModel.SupportsRcMigration).To(Equal(core.BoolPtr(true)))
				Expect(deploymentBaseModel.TargetNetwork).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				Expect(objectMetadataSetModel).ToNot(BeNil())
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel
				Expect(objectMetadataSetModel.RcCompatible).To(Equal(core.BoolPtr(true)))
				Expect(objectMetadataSetModel.Service).To(Equal(cfMetaDataModel))
				Expect(objectMetadataSetModel.Plan).To(Equal(planMetaDataModel))
				Expect(objectMetadataSetModel.Alias).To(Equal(aliasMetaDataModel))
				Expect(objectMetadataSetModel.Template).To(Equal(templateMetaDataModel))
				Expect(objectMetadataSetModel.UI).To(Equal(uiMetaDataModel))
				Expect(objectMetadataSetModel.Compliance).To(Equal([]string{"testString"}))
				Expect(objectMetadataSetModel.SLA).To(Equal(slaMetaDataModel))
				Expect(objectMetadataSetModel.Callbacks).To(Equal(callbacksModel))
				Expect(objectMetadataSetModel.OriginalName).To(Equal(core.StringPtr("testString")))
				Expect(objectMetadataSetModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(objectMetadataSetModel.Other).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(objectMetadataSetModel.Pricing).To(Equal(pricingSetModel))
				Expect(objectMetadataSetModel.Deployment).To(Equal(deploymentBaseModel))

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsName := "testString"
				createCatalogEntryOptionsKind := "service"
				createCatalogEntryOptionsOverviewUI := map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				var createCatalogEntryOptionsImages *globalcatalogv1.Image = nil
				createCatalogEntryOptionsDisabled := true
				createCatalogEntryOptionsTags := []string{"testString"}
				var createCatalogEntryOptionsProvider *globalcatalogv1.Provider = nil
				createCatalogEntryOptionsID := "testString"
				createCatalogEntryOptionsModel := globalCatalogService.NewCreateCatalogEntryOptions(createCatalogEntryOptionsName, createCatalogEntryOptionsKind, createCatalogEntryOptionsOverviewUI, createCatalogEntryOptionsImages, createCatalogEntryOptionsDisabled, createCatalogEntryOptionsTags, createCatalogEntryOptionsProvider, createCatalogEntryOptionsID)
				createCatalogEntryOptionsModel.SetName("testString")
				createCatalogEntryOptionsModel.SetKind("service")
				createCatalogEntryOptionsModel.SetOverviewUI(map[string]globalcatalogv1.Overview{"key1": *overviewModel})
				createCatalogEntryOptionsModel.SetImages(imageModel)
				createCatalogEntryOptionsModel.SetDisabled(true)
				createCatalogEntryOptionsModel.SetTags([]string{"testString"})
				createCatalogEntryOptionsModel.SetProvider(providerModel)
				createCatalogEntryOptionsModel.SetID("testString")
				createCatalogEntryOptionsModel.SetParentID("testString")
				createCatalogEntryOptionsModel.SetGroup(true)
				createCatalogEntryOptionsModel.SetActive(true)
				createCatalogEntryOptionsModel.SetURL("testString")
				createCatalogEntryOptionsModel.SetMetadata(objectMetadataSetModel)
				createCatalogEntryOptionsModel.SetAccount("testString")
				createCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(createCatalogEntryOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogEntryOptionsModel.Kind).To(Equal(core.StringPtr("service")))
				Expect(createCatalogEntryOptionsModel.OverviewUI).To(Equal(map[string]globalcatalogv1.Overview{"key1": *overviewModel}))
				Expect(createCatalogEntryOptionsModel.Images).To(Equal(imageModel))
				Expect(createCatalogEntryOptionsModel.Disabled).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogEntryOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createCatalogEntryOptionsModel.Provider).To(Equal(providerModel))
				Expect(createCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogEntryOptionsModel.ParentID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogEntryOptionsModel.Group).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogEntryOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogEntryOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogEntryOptionsModel.Metadata).To(Equal(objectMetadataSetModel))
				Expect(createCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteArtifactOptions successfully`, func() {
				// Construct an instance of the DeleteArtifactOptions model
				objectID := "testString"
				artifactID := "testString"
				deleteArtifactOptionsModel := globalCatalogService.NewDeleteArtifactOptions(objectID, artifactID)
				deleteArtifactOptionsModel.SetObjectID("testString")
				deleteArtifactOptionsModel.SetArtifactID("testString")
				deleteArtifactOptionsModel.SetAccount("testString")
				deleteArtifactOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteArtifactOptionsModel).ToNot(BeNil())
				Expect(deleteArtifactOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteArtifactOptionsModel.ArtifactID).To(Equal(core.StringPtr("testString")))
				Expect(deleteArtifactOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(deleteArtifactOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCatalogEntryOptions successfully`, func() {
				// Construct an instance of the DeleteCatalogEntryOptions model
				id := "testString"
				deleteCatalogEntryOptionsModel := globalCatalogService.NewDeleteCatalogEntryOptions(id)
				deleteCatalogEntryOptionsModel.SetID("testString")
				deleteCatalogEntryOptionsModel.SetAccount("testString")
				deleteCatalogEntryOptionsModel.SetForce(true)
				deleteCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(deleteCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogEntryOptionsModel.Force).To(Equal(core.BoolPtr(true)))
				Expect(deleteCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetArtifactOptions successfully`, func() {
				// Construct an instance of the GetArtifactOptions model
				objectID := "testString"
				artifactID := "testString"
				getArtifactOptionsModel := globalCatalogService.NewGetArtifactOptions(objectID, artifactID)
				getArtifactOptionsModel.SetObjectID("testString")
				getArtifactOptionsModel.SetArtifactID("testString")
				getArtifactOptionsModel.SetAccept("testString")
				getArtifactOptionsModel.SetAccount("testString")
				getArtifactOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getArtifactOptionsModel).ToNot(BeNil())
				Expect(getArtifactOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(getArtifactOptionsModel.ArtifactID).To(Equal(core.StringPtr("testString")))
				Expect(getArtifactOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(getArtifactOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getArtifactOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAuditLogsOptions successfully`, func() {
				// Construct an instance of the GetAuditLogsOptions model
				id := "testString"
				getAuditLogsOptionsModel := globalCatalogService.NewGetAuditLogsOptions(id)
				getAuditLogsOptionsModel.SetID("testString")
				getAuditLogsOptionsModel.SetAccount("testString")
				getAuditLogsOptionsModel.SetAscending("false")
				getAuditLogsOptionsModel.SetStartat("testString")
				getAuditLogsOptionsModel.SetOffset(int64(0))
				getAuditLogsOptionsModel.SetLimit(int64(50))
				getAuditLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAuditLogsOptionsModel).ToNot(BeNil())
				Expect(getAuditLogsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getAuditLogsOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getAuditLogsOptionsModel.Ascending).To(Equal(core.StringPtr("false")))
				Expect(getAuditLogsOptionsModel.Startat).To(Equal(core.StringPtr("testString")))
				Expect(getAuditLogsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getAuditLogsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(getAuditLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogEntryOptions successfully`, func() {
				// Construct an instance of the GetCatalogEntryOptions model
				id := "testString"
				getCatalogEntryOptionsModel := globalCatalogService.NewGetCatalogEntryOptions(id)
				getCatalogEntryOptionsModel.SetID("testString")
				getCatalogEntryOptionsModel.SetAccount("testString")
				getCatalogEntryOptionsModel.SetInclude("testString")
				getCatalogEntryOptionsModel.SetLanguages("testString")
				getCatalogEntryOptionsModel.SetComplete(true)
				getCatalogEntryOptionsModel.SetDepth(int64(38))
				getCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(getCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogEntryOptionsModel.Include).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogEntryOptionsModel.Languages).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogEntryOptionsModel.Complete).To(Equal(core.BoolPtr(true)))
				Expect(getCatalogEntryOptionsModel.Depth).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetChildObjectsOptions successfully`, func() {
				// Construct an instance of the GetChildObjectsOptions model
				id := "testString"
				kind := "testString"
				getChildObjectsOptionsModel := globalCatalogService.NewGetChildObjectsOptions(id, kind)
				getChildObjectsOptionsModel.SetID("testString")
				getChildObjectsOptionsModel.SetKind("testString")
				getChildObjectsOptionsModel.SetAccount("testString")
				getChildObjectsOptionsModel.SetInclude("testString")
				getChildObjectsOptionsModel.SetQ("testString")
				getChildObjectsOptionsModel.SetSortBy("testString")
				getChildObjectsOptionsModel.SetDescending("testString")
				getChildObjectsOptionsModel.SetLanguages("testString")
				getChildObjectsOptionsModel.SetComplete(true)
				getChildObjectsOptionsModel.SetOffset(int64(0))
				getChildObjectsOptionsModel.SetLimit(int64(50))
				getChildObjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getChildObjectsOptionsModel).ToNot(BeNil())
				Expect(getChildObjectsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Kind).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Include).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Q).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.SortBy).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Descending).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Languages).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Complete).To(Equal(core.BoolPtr(true)))
				Expect(getChildObjectsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getChildObjectsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(getChildObjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPricingDeploymentsOptions successfully`, func() {
				// Construct an instance of the GetPricingDeploymentsOptions model
				id := "testString"
				getPricingDeploymentsOptionsModel := globalCatalogService.NewGetPricingDeploymentsOptions(id)
				getPricingDeploymentsOptionsModel.SetID("testString")
				getPricingDeploymentsOptionsModel.SetAccount("testString")
				getPricingDeploymentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPricingDeploymentsOptionsModel).ToNot(BeNil())
				Expect(getPricingDeploymentsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getPricingDeploymentsOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getPricingDeploymentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPricingOptions successfully`, func() {
				// Construct an instance of the GetPricingOptions model
				id := "testString"
				getPricingOptionsModel := globalCatalogService.NewGetPricingOptions(id)
				getPricingOptionsModel.SetID("testString")
				getPricingOptionsModel.SetAccount("testString")
				getPricingOptionsModel.SetDeploymentRegion("testString")
				getPricingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPricingOptionsModel).ToNot(BeNil())
				Expect(getPricingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getPricingOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getPricingOptionsModel.DeploymentRegion).To(Equal(core.StringPtr("testString")))
				Expect(getPricingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVisibilityOptions successfully`, func() {
				// Construct an instance of the GetVisibilityOptions model
				id := "testString"
				getVisibilityOptionsModel := globalCatalogService.NewGetVisibilityOptions(id)
				getVisibilityOptionsModel.SetID("testString")
				getVisibilityOptionsModel.SetAccount("testString")
				getVisibilityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVisibilityOptionsModel).ToNot(BeNil())
				Expect(getVisibilityOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getVisibilityOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getVisibilityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImage successfully`, func() {
				image := "testString"
				_model, err := globalCatalogService.NewImage(image)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListArtifactsOptions successfully`, func() {
				// Construct an instance of the ListArtifactsOptions model
				objectID := "testString"
				listArtifactsOptionsModel := globalCatalogService.NewListArtifactsOptions(objectID)
				listArtifactsOptionsModel.SetObjectID("testString")
				listArtifactsOptionsModel.SetAccount("testString")
				listArtifactsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listArtifactsOptionsModel).ToNot(BeNil())
				Expect(listArtifactsOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(listArtifactsOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(listArtifactsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCatalogEntriesOptions successfully`, func() {
				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := globalCatalogService.NewListCatalogEntriesOptions()
				listCatalogEntriesOptionsModel.SetAccount("testString")
				listCatalogEntriesOptionsModel.SetInclude("testString")
				listCatalogEntriesOptionsModel.SetQ("testString")
				listCatalogEntriesOptionsModel.SetSortBy("testString")
				listCatalogEntriesOptionsModel.SetDescending("testString")
				listCatalogEntriesOptionsModel.SetLanguages("testString")
				listCatalogEntriesOptionsModel.SetCatalog(true)
				listCatalogEntriesOptionsModel.SetComplete(true)
				listCatalogEntriesOptionsModel.SetOffset(int64(0))
				listCatalogEntriesOptionsModel.SetLimit(int64(50))
				listCatalogEntriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCatalogEntriesOptionsModel).ToNot(BeNil())
				Expect(listCatalogEntriesOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Include).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Q).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.SortBy).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Descending).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Languages).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Catalog).To(Equal(core.BoolPtr(true)))
				Expect(listCatalogEntriesOptionsModel.Complete).To(Equal(core.BoolPtr(true)))
				Expect(listCatalogEntriesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listCatalogEntriesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listCatalogEntriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewOverview successfully`, func() {
				displayName := "testString"
				longDescription := "testString"
				description := "testString"
				_model, err := globalCatalogService.NewOverview(displayName, longDescription, description)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProvider successfully`, func() {
				email := "testString"
				name := "testString"
				_model, err := globalCatalogService.NewProvider(email, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRestoreCatalogEntryOptions successfully`, func() {
				// Construct an instance of the RestoreCatalogEntryOptions model
				id := "testString"
				restoreCatalogEntryOptionsModel := globalCatalogService.NewRestoreCatalogEntryOptions(id)
				restoreCatalogEntryOptionsModel.SetID("testString")
				restoreCatalogEntryOptionsModel.SetAccount("testString")
				restoreCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(restoreCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(restoreCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(restoreCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCatalogEntryOptions successfully`, func() {
				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				Expect(overviewModel).ToNot(BeNil())
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")
				Expect(overviewModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.FeaturedDescription).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				Expect(imageModel).ToNot(BeNil())
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")
				Expect(imageModel.Image).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.SmallImage).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.MediumImage).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.FeatureImage).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				Expect(providerModel).ToNot(BeNil())
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")
				Expect(providerModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.SupportEmail).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Phone).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CfMetaData model
				cfMetaDataModel := new(globalcatalogv1.CfMetaData)
				Expect(cfMetaDataModel).ToNot(BeNil())
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueAPIKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				cfMetaDataModel.CRNMask = core.StringPtr("testString")
				cfMetaDataModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.Extension = map[string]interface{}{"anyKey": "anyValue"}
				cfMetaDataModel.PaidOnly = core.BoolPtr(true)
				cfMetaDataModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)
				Expect(cfMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.IamCompatible).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.UniqueAPIKey).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Provisionable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.AsyncProvisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.AsyncUnprovisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Requires).To(Equal([]string{"testString"}))
				Expect(cfMetaDataModel.PlanUpdateable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.State).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.ServiceCheckEnabled).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.TestCheckInterval).To(Equal(core.Int64Ptr(int64(38))))
				Expect(cfMetaDataModel.ServiceKeySupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.CfGUID).To(Equal(map[string]string{"key1": "testString"}))
				Expect(cfMetaDataModel.CRNMask).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.Parameters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(cfMetaDataModel.UserDefinedService).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(cfMetaDataModel.Extension).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(cfMetaDataModel.PaidOnly).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.CustomCreatePageHybridEnabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				Expect(planMetaDataModel).ToNot(BeNil())
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.ProvisionType = core.StringPtr("testString")
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGUID = map[string]string{"key1": "testString"}
				Expect(planMetaDataModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.Reservable).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AllowInternalUsers).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AsyncProvisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AsyncUnprovisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.ProvisionType).To(Equal(core.StringPtr("testString")))
				Expect(planMetaDataModel.TestCheckInterval).To(Equal(core.Int64Ptr(int64(38))))
				Expect(planMetaDataModel.SingleScopeInstance).To(Equal(core.StringPtr("testString")))
				Expect(planMetaDataModel.ServiceCheckEnabled).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.CfGUID).To(Equal(map[string]string{"key1": "testString"}))

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				Expect(aliasMetaDataModel).ToNot(BeNil())
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")
				Expect(aliasMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(aliasMetaDataModel.PlanID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				Expect(sourceMetaDataModel).ToNot(BeNil())
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")
				Expect(sourceMetaDataModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(sourceMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(sourceMetaDataModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				Expect(templateMetaDataModel).ToNot(BeNil())
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = map[string]string{"key1": "testString"}
				Expect(templateMetaDataModel.Services).To(Equal([]string{"testString"}))
				Expect(templateMetaDataModel.DefaultMemory).To(Equal(core.Int64Ptr(int64(38))))
				Expect(templateMetaDataModel.StartCmd).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.Source).To(Equal(sourceMetaDataModel))
				Expect(templateMetaDataModel.RuntimeCatalogID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.CfRuntimeID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.ExecutableFile).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.Buildpack).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.EnvironmentVariables).To(Equal(map[string]string{"key1": "testString"}))

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				Expect(bulletsModel).ToNot(BeNil())
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))
				Expect(bulletsModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Icon).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Quantity).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the UIMediaSourceMetaData model
				uiMediaSourceMetaDataModel := new(globalcatalogv1.UIMediaSourceMetaData)
				Expect(uiMediaSourceMetaDataModel).ToNot(BeNil())
				uiMediaSourceMetaDataModel.Type = core.StringPtr("testString")
				uiMediaSourceMetaDataModel.URL = core.StringPtr("testString")
				Expect(uiMediaSourceMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(uiMediaSourceMetaDataModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				Expect(uiMetaMediaModel).ToNot(BeNil())
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}
				Expect(uiMetaMediaModel.Caption).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.ThumbnailURL).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.Source).To(Equal([]globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel}))

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				Expect(stringsModel).ToNot(BeNil())
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")
				Expect(stringsModel.Bullets).To(Equal([]globalcatalogv1.Bullets{*bulletsModel}))
				Expect(stringsModel.Media).To(Equal([]globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}))
				Expect(stringsModel.NotCreatableMsg).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.NotCreatableRobotMsg).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.DeprecationWarning).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.PopupWarningMessage).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.Instruction).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Urls model
				urlsModel := new(globalcatalogv1.Urls)
				Expect(urlsModel).ToNot(BeNil())
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.APIURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")
				Expect(urlsModel.DocURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.InstructionsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.APIURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CreateURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.SdkDownloadURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.TermsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CustomCreatePageURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CatalogDetailsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.DeprecationDocURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.DashboardURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.RegistrationURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.Apidocsurl).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				Expect(uiMetaDataModel).ToNot(BeNil())
				uiMetaDataModel.Strings = map[string]globalcatalogv1.Strings{"key1": *stringsModel}
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				uiMetaDataModel.Strings["foo"] = *stringsModel
				Expect(uiMetaDataModel.Urls).To(Equal(urlsModel))
				Expect(uiMetaDataModel.EmbeddableDashboard).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaDataModel.EmbeddableDashboardFullWidth).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.NavigationOrder).To(Equal([]string{"testString"}))
				Expect(uiMetaDataModel.NotCreatable).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.PrimaryOfferingID).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaDataModel.AccessibleDuringProvision).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.SideBySideIndex).To(Equal(core.Int64Ptr(int64(38))))
				Expect(uiMetaDataModel.EndOfServiceTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(uiMetaDataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.HideLiteMetering).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.NoUpgradeNextStep).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.Strings["foo"]).To(Equal(*stringsModel))

				// Construct an instance of the DrMetaData model
				drMetaDataModel := new(globalcatalogv1.DrMetaData)
				Expect(drMetaDataModel).ToNot(BeNil())
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")
				Expect(drMetaDataModel.Dr).To(Equal(core.BoolPtr(true)))
				Expect(drMetaDataModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				Expect(slaMetaDataModel).ToNot(BeNil())
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Responsiveness = core.Float64Ptr(float64(72.5))
				slaMetaDataModel.Dr = drMetaDataModel
				Expect(slaMetaDataModel.Terms).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Tenancy).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Provisioning).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(slaMetaDataModel.Responsiveness).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(slaMetaDataModel.Dr).To(Equal(drMetaDataModel))

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				Expect(callbacksModel).ToNot(BeNil())
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorAPI = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.APIEndpoint = map[string]string{"key1": "testString"}
				Expect(callbacksModel.ControllerURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.BrokerURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.BrokerProxyURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDataURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDetailTabURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDetailTabExtURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ServiceMonitorAPI).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ServiceMonitorApp).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.APIEndpoint).To(Equal(map[string]string{"key1": "testString"}))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				Expect(priceModel).ToNot(BeNil())
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(float64(72.5))
				Expect(priceModel.QuantityTier).To(Equal(core.Int64Ptr(int64(38))))
				Expect(priceModel.Price).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				Expect(amountModel).ToNot(BeNil())
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}
				Expect(amountModel.Country).To(Equal(core.StringPtr("testString")))
				Expect(amountModel.Currency).To(Equal(core.StringPtr("testString")))
				Expect(amountModel.Prices).To(Equal([]globalcatalogv1.Price{*priceModel}))

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				Expect(startingPriceModel).ToNot(BeNil())
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}
				Expect(startingPriceModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.DeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.Unit).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.Amount).To(Equal([]globalcatalogv1.Amount{*amountModel}))

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				Expect(pricingSetModel).ToNot(BeNil())
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel
				Expect(pricingSetModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(pricingSetModel.Origin).To(Equal(core.StringPtr("testString")))
				Expect(pricingSetModel.StartingPrice).To(Equal(startingPriceModel))

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				Expect(brokerModel).ToNot(BeNil())
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.GUID = core.StringPtr("testString")
				Expect(brokerModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(brokerModel.GUID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				Expect(deploymentBaseModel).ToNot(BeNil())
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCRN = core.StringPtr("testString")
				deploymentBaseModel.ServiceCRN = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")
				Expect(deploymentBaseModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.LocationURL).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.OriginalLocation).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.TargetCRN).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.ServiceCRN).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.MccpID).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.Broker).To(Equal(brokerModel))
				Expect(deploymentBaseModel.SupportsRcMigration).To(Equal(core.BoolPtr(true)))
				Expect(deploymentBaseModel.TargetNetwork).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				Expect(objectMetadataSetModel).ToNot(BeNil())
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.UI = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.SLA = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = map[string]interface{}{"anyKey": "anyValue"}
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel
				Expect(objectMetadataSetModel.RcCompatible).To(Equal(core.BoolPtr(true)))
				Expect(objectMetadataSetModel.Service).To(Equal(cfMetaDataModel))
				Expect(objectMetadataSetModel.Plan).To(Equal(planMetaDataModel))
				Expect(objectMetadataSetModel.Alias).To(Equal(aliasMetaDataModel))
				Expect(objectMetadataSetModel.Template).To(Equal(templateMetaDataModel))
				Expect(objectMetadataSetModel.UI).To(Equal(uiMetaDataModel))
				Expect(objectMetadataSetModel.Compliance).To(Equal([]string{"testString"}))
				Expect(objectMetadataSetModel.SLA).To(Equal(slaMetaDataModel))
				Expect(objectMetadataSetModel.Callbacks).To(Equal(callbacksModel))
				Expect(objectMetadataSetModel.OriginalName).To(Equal(core.StringPtr("testString")))
				Expect(objectMetadataSetModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(objectMetadataSetModel.Other).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(objectMetadataSetModel.Pricing).To(Equal(pricingSetModel))
				Expect(objectMetadataSetModel.Deployment).To(Equal(deploymentBaseModel))

				// Construct an instance of the UpdateCatalogEntryOptions model
				id := "testString"
				updateCatalogEntryOptionsName := "testString"
				updateCatalogEntryOptionsKind := "service"
				updateCatalogEntryOptionsOverviewUI := map[string]globalcatalogv1.Overview{"key1": *overviewModel}
				var updateCatalogEntryOptionsImages *globalcatalogv1.Image = nil
				updateCatalogEntryOptionsDisabled := true
				updateCatalogEntryOptionsTags := []string{"testString"}
				var updateCatalogEntryOptionsProvider *globalcatalogv1.Provider = nil
				updateCatalogEntryOptionsModel := globalCatalogService.NewUpdateCatalogEntryOptions(id, updateCatalogEntryOptionsName, updateCatalogEntryOptionsKind, updateCatalogEntryOptionsOverviewUI, updateCatalogEntryOptionsImages, updateCatalogEntryOptionsDisabled, updateCatalogEntryOptionsTags, updateCatalogEntryOptionsProvider)
				updateCatalogEntryOptionsModel.SetID("testString")
				updateCatalogEntryOptionsModel.SetName("testString")
				updateCatalogEntryOptionsModel.SetKind("service")
				updateCatalogEntryOptionsModel.SetOverviewUI(map[string]globalcatalogv1.Overview{"key1": *overviewModel})
				updateCatalogEntryOptionsModel.SetImages(imageModel)
				updateCatalogEntryOptionsModel.SetDisabled(true)
				updateCatalogEntryOptionsModel.SetTags([]string{"testString"})
				updateCatalogEntryOptionsModel.SetProvider(providerModel)
				updateCatalogEntryOptionsModel.SetParentID("testString")
				updateCatalogEntryOptionsModel.SetGroup(true)
				updateCatalogEntryOptionsModel.SetActive(true)
				updateCatalogEntryOptionsModel.SetURL("testString")
				updateCatalogEntryOptionsModel.SetMetadata(objectMetadataSetModel)
				updateCatalogEntryOptionsModel.SetAccount("testString")
				updateCatalogEntryOptionsModel.SetMove("testString")
				updateCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(updateCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Kind).To(Equal(core.StringPtr("service")))
				Expect(updateCatalogEntryOptionsModel.OverviewUI).To(Equal(map[string]globalcatalogv1.Overview{"key1": *overviewModel}))
				Expect(updateCatalogEntryOptionsModel.Images).To(Equal(imageModel))
				Expect(updateCatalogEntryOptionsModel.Disabled).To(Equal(core.BoolPtr(true)))
				Expect(updateCatalogEntryOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(updateCatalogEntryOptionsModel.Provider).To(Equal(providerModel))
				Expect(updateCatalogEntryOptionsModel.ParentID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Group).To(Equal(core.BoolPtr(true)))
				Expect(updateCatalogEntryOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(updateCatalogEntryOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Metadata).To(Equal(objectMetadataSetModel))
				Expect(updateCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Move).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateVisibilityOptions successfully`, func() {
				// Construct an instance of the VisibilityDetailAccounts model
				visibilityDetailAccountsModel := new(globalcatalogv1.VisibilityDetailAccounts)
				Expect(visibilityDetailAccountsModel).ToNot(BeNil())
				visibilityDetailAccountsModel.Accountid = core.StringPtr("testString")
				Expect(visibilityDetailAccountsModel.Accountid).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VisibilityDetail model
				visibilityDetailModel := new(globalcatalogv1.VisibilityDetail)
				Expect(visibilityDetailModel).ToNot(BeNil())
				visibilityDetailModel.Accounts = visibilityDetailAccountsModel
				Expect(visibilityDetailModel.Accounts).To(Equal(visibilityDetailAccountsModel))

				// Construct an instance of the UpdateVisibilityOptions model
				id := "testString"
				updateVisibilityOptionsModel := globalCatalogService.NewUpdateVisibilityOptions(id)
				updateVisibilityOptionsModel.SetID("testString")
				updateVisibilityOptionsModel.SetRestrictions("testString")
				updateVisibilityOptionsModel.SetExtendable(true)
				updateVisibilityOptionsModel.SetInclude(visibilityDetailModel)
				updateVisibilityOptionsModel.SetExclude(visibilityDetailModel)
				updateVisibilityOptionsModel.SetAccount("testString")
				updateVisibilityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateVisibilityOptionsModel).ToNot(BeNil())
				Expect(updateVisibilityOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateVisibilityOptionsModel.Restrictions).To(Equal(core.StringPtr("testString")))
				Expect(updateVisibilityOptionsModel.Extendable).To(Equal(core.BoolPtr(true)))
				Expect(updateVisibilityOptionsModel.Include).To(Equal(visibilityDetailModel))
				Expect(updateVisibilityOptionsModel.Exclude).To(Equal(visibilityDetailModel))
				Expect(updateVisibilityOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(updateVisibilityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadArtifactOptions successfully`, func() {
				// Construct an instance of the UploadArtifactOptions model
				objectID := "testString"
				artifactID := "testString"
				uploadArtifactOptionsModel := globalCatalogService.NewUploadArtifactOptions(objectID, artifactID)
				uploadArtifactOptionsModel.SetObjectID("testString")
				uploadArtifactOptionsModel.SetArtifactID("testString")
				uploadArtifactOptionsModel.SetArtifact(CreateMockReader("This is a mock file."))
				uploadArtifactOptionsModel.SetContentType("testString")
				uploadArtifactOptionsModel.SetAccount("testString")
				uploadArtifactOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadArtifactOptionsModel).ToNot(BeNil())
				Expect(uploadArtifactOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(uploadArtifactOptionsModel.ArtifactID).To(Equal(core.StringPtr("testString")))
				Expect(uploadArtifactOptionsModel.Artifact).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(uploadArtifactOptionsModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(uploadArtifactOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(uploadArtifactOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVisibilityDetail successfully`, func() {
				var accounts *globalcatalogv1.VisibilityDetailAccounts = nil
				_, err := globalCatalogService.NewVisibilityDetail(accounts)
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAliasMetaData successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.AliasMetaData)
			model.Type = core.StringPtr("testString")
			model.PlanID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.AliasMetaData
			err = globalcatalogv1.UnmarshalAliasMetaData(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAmount successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Amount)
			model.Country = core.StringPtr("testString")
			model.Currency = core.StringPtr("testString")
			model.Prices = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Amount
			err = globalcatalogv1.UnmarshalAmount(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalBroker successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Broker)
			model.Name = core.StringPtr("testString")
			model.GUID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Broker
			err = globalcatalogv1.UnmarshalBroker(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalBullets successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Bullets)
			model.Title = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Icon = core.StringPtr("testString")
			model.Quantity = core.Int64Ptr(int64(38))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Bullets
			err = globalcatalogv1.UnmarshalBullets(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCfMetaData successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.CfMetaData)
			model.Type = core.StringPtr("testString")
			model.IamCompatible = core.BoolPtr(true)
			model.UniqueAPIKey = core.BoolPtr(true)
			model.Provisionable = core.BoolPtr(true)
			model.Bindable = core.BoolPtr(true)
			model.AsyncProvisioningSupported = core.BoolPtr(true)
			model.AsyncUnprovisioningSupported = core.BoolPtr(true)
			model.Requires = []string{"testString"}
			model.PlanUpdateable = core.BoolPtr(true)
			model.State = core.StringPtr("testString")
			model.ServiceCheckEnabled = core.BoolPtr(true)
			model.TestCheckInterval = core.Int64Ptr(int64(38))
			model.ServiceKeySupported = core.BoolPtr(true)
			model.CfGUID = map[string]string{"key1": "testString"}
			model.CRNMask = core.StringPtr("testString")
			model.Parameters = map[string]interface{}{"anyKey": "anyValue"}
			model.UserDefinedService = map[string]interface{}{"anyKey": "anyValue"}
			model.Extension = map[string]interface{}{"anyKey": "anyValue"}
			model.PaidOnly = core.BoolPtr(true)
			model.CustomCreatePageHybridEnabled = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.CfMetaData
			err = globalcatalogv1.UnmarshalCfMetaData(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCallbacks successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Callbacks)
			model.ControllerURL = core.StringPtr("testString")
			model.BrokerURL = core.StringPtr("testString")
			model.BrokerProxyURL = core.StringPtr("testString")
			model.DashboardURL = core.StringPtr("testString")
			model.DashboardDataURL = core.StringPtr("testString")
			model.DashboardDetailTabURL = core.StringPtr("testString")
			model.DashboardDetailTabExtURL = core.StringPtr("testString")
			model.ServiceMonitorAPI = core.StringPtr("testString")
			model.ServiceMonitorApp = core.StringPtr("testString")
			model.APIEndpoint = map[string]string{"key1": "testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Callbacks
			err = globalcatalogv1.UnmarshalCallbacks(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDrMetaData successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.DrMetaData)
			model.Dr = core.BoolPtr(true)
			model.Description = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.DrMetaData
			err = globalcatalogv1.UnmarshalDrMetaData(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDeploymentBase successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.DeploymentBase)
			model.Location = core.StringPtr("testString")
			model.LocationURL = core.StringPtr("testString")
			model.OriginalLocation = core.StringPtr("testString")
			model.TargetCRN = core.StringPtr("testString")
			model.ServiceCRN = core.StringPtr("testString")
			model.MccpID = core.StringPtr("testString")
			model.Broker = nil
			model.SupportsRcMigration = core.BoolPtr(true)
			model.TargetNetwork = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.DeploymentBase
			err = globalcatalogv1.UnmarshalDeploymentBase(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalImage successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Image)
			model.Image = core.StringPtr("testString")
			model.SmallImage = core.StringPtr("testString")
			model.MediumImage = core.StringPtr("testString")
			model.FeatureImage = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Image
			err = globalcatalogv1.UnmarshalImage(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalObjectMetadataSet successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.ObjectMetadataSet)
			model.RcCompatible = core.BoolPtr(true)
			model.Service = nil
			model.Plan = nil
			model.Alias = nil
			model.Template = nil
			model.UI = nil
			model.Compliance = []string{"testString"}
			model.SLA = nil
			model.Callbacks = nil
			model.OriginalName = core.StringPtr("testString")
			model.Version = core.StringPtr("testString")
			model.Other = map[string]interface{}{"anyKey": "anyValue"}
			model.Pricing = nil
			model.Deployment = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.ObjectMetadataSet
			err = globalcatalogv1.UnmarshalObjectMetadataSet(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalOverview successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Overview)
			model.DisplayName = core.StringPtr("testString")
			model.LongDescription = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.FeaturedDescription = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Overview
			err = globalcatalogv1.UnmarshalOverview(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPlanMetaData successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.PlanMetaData)
			model.Bindable = core.BoolPtr(true)
			model.Reservable = core.BoolPtr(true)
			model.AllowInternalUsers = core.BoolPtr(true)
			model.AsyncProvisioningSupported = core.BoolPtr(true)
			model.AsyncUnprovisioningSupported = core.BoolPtr(true)
			model.ProvisionType = core.StringPtr("testString")
			model.TestCheckInterval = core.Int64Ptr(int64(38))
			model.SingleScopeInstance = core.StringPtr("testString")
			model.ServiceCheckEnabled = core.BoolPtr(true)
			model.CfGUID = map[string]string{"key1": "testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.PlanMetaData
			err = globalcatalogv1.UnmarshalPlanMetaData(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPrice successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Price)
			model.QuantityTier = core.Int64Ptr(int64(38))
			model.Price = core.Float64Ptr(float64(72.5))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Price
			err = globalcatalogv1.UnmarshalPrice(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPricingSet successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.PricingSet)
			model.Type = core.StringPtr("testString")
			model.Origin = core.StringPtr("testString")
			model.StartingPrice = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.PricingSet
			err = globalcatalogv1.UnmarshalPricingSet(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalProvider successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Provider)
			model.Email = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")
			model.Contact = core.StringPtr("testString")
			model.SupportEmail = core.StringPtr("testString")
			model.Phone = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Provider
			err = globalcatalogv1.UnmarshalProvider(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSLAMetaData successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.SLAMetaData)
			model.Terms = core.StringPtr("testString")
			model.Tenancy = core.StringPtr("testString")
			model.Provisioning = core.Float64Ptr(float64(72.5))
			model.Responsiveness = core.Float64Ptr(float64(72.5))
			model.Dr = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.SLAMetaData
			err = globalcatalogv1.UnmarshalSLAMetaData(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSourceMetaData successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.SourceMetaData)
			model.Path = core.StringPtr("testString")
			model.Type = core.StringPtr("testString")
			model.URL = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.SourceMetaData
			err = globalcatalogv1.UnmarshalSourceMetaData(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalStartingPrice successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.StartingPrice)
			model.PlanID = core.StringPtr("testString")
			model.DeploymentID = core.StringPtr("testString")
			model.Unit = core.StringPtr("testString")
			model.Amount = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.StartingPrice
			err = globalcatalogv1.UnmarshalStartingPrice(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalStrings successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Strings)
			model.Bullets = nil
			model.Media = nil
			model.NotCreatableMsg = core.StringPtr("testString")
			model.NotCreatableRobotMsg = core.StringPtr("testString")
			model.DeprecationWarning = core.StringPtr("testString")
			model.PopupWarningMessage = core.StringPtr("testString")
			model.Instruction = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Strings
			err = globalcatalogv1.UnmarshalStrings(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTemplateMetaData successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.TemplateMetaData)
			model.Services = []string{"testString"}
			model.DefaultMemory = core.Int64Ptr(int64(38))
			model.StartCmd = core.StringPtr("testString")
			model.Source = nil
			model.RuntimeCatalogID = core.StringPtr("testString")
			model.CfRuntimeID = core.StringPtr("testString")
			model.TemplateID = core.StringPtr("testString")
			model.ExecutableFile = core.StringPtr("testString")
			model.Buildpack = core.StringPtr("testString")
			model.EnvironmentVariables = map[string]string{"key1": "testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.TemplateMetaData
			err = globalcatalogv1.UnmarshalTemplateMetaData(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUIMediaSourceMetaData successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.UIMediaSourceMetaData)
			model.Type = core.StringPtr("testString")
			model.URL = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.UIMediaSourceMetaData
			err = globalcatalogv1.UnmarshalUIMediaSourceMetaData(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUIMetaData successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.UIMetaData)
			model.Strings = nil
			model.Urls = nil
			model.EmbeddableDashboard = core.StringPtr("testString")
			model.EmbeddableDashboardFullWidth = core.BoolPtr(true)
			model.NavigationOrder = []string{"testString"}
			model.NotCreatable = core.BoolPtr(true)
			model.PrimaryOfferingID = core.StringPtr("testString")
			model.AccessibleDuringProvision = core.BoolPtr(true)
			model.SideBySideIndex = core.Int64Ptr(int64(38))
			model.EndOfServiceTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
			model.Hidden = core.BoolPtr(true)
			model.HideLiteMetering = core.BoolPtr(true)
			model.NoUpgradeNextStep = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.UIMetaData
			err = globalcatalogv1.UnmarshalUIMetaData(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUIMetaMedia successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.UIMetaMedia)
			model.Caption = core.StringPtr("testString")
			model.ThumbnailURL = core.StringPtr("testString")
			model.Type = core.StringPtr("testString")
			model.URL = core.StringPtr("testString")
			model.Source = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.UIMetaMedia
			err = globalcatalogv1.UnmarshalUIMetaMedia(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUrls successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Urls)
			model.DocURL = core.StringPtr("testString")
			model.InstructionsURL = core.StringPtr("testString")
			model.APIURL = core.StringPtr("testString")
			model.CreateURL = core.StringPtr("testString")
			model.SdkDownloadURL = core.StringPtr("testString")
			model.TermsURL = core.StringPtr("testString")
			model.CustomCreatePageURL = core.StringPtr("testString")
			model.CatalogDetailsURL = core.StringPtr("testString")
			model.DeprecationDocURL = core.StringPtr("testString")
			model.DashboardURL = core.StringPtr("testString")
			model.RegistrationURL = core.StringPtr("testString")
			model.Apidocsurl = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Urls
			err = globalcatalogv1.UnmarshalUrls(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVisibility successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.Visibility)
			model.Restrictions = core.StringPtr("testString")
			model.Owner = core.StringPtr("testString")
			model.Extendable = core.BoolPtr(true)
			model.Include = nil
			model.Exclude = nil
			model.Approved = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.Visibility
			err = globalcatalogv1.UnmarshalVisibility(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVisibilityDetail successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.VisibilityDetail)
			model.Accounts = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.VisibilityDetail
			err = globalcatalogv1.UnmarshalVisibilityDetail(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalVisibilityDetailAccounts successfully`, func() {
			// Construct an instance of the model.
			model := new(globalcatalogv1.VisibilityDetailAccounts)
			model.Accountid = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *globalcatalogv1.VisibilityDetailAccounts
			err = globalcatalogv1.UnmarshalVisibilityDetailAccounts(raw, &result)
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
