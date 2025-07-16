//go:build integration

/**
 * (C) Copyright IBM Corp. 2023.
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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the globalcatalogv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`GlobalCatalogV1 Integration Tests`, func() {
	const externalConfigFile = "../global_catalog.env"

	var (
		err                  error
		globalCatalogService *globalcatalogv1.GlobalCatalogV1
		serviceURL           string
		config               map[string]string
		fetchedEntry         *globalcatalogv1.CatalogEntry
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(globalcatalogv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			globalCatalogServiceOptions := &globalcatalogv1.GlobalCatalogV1Options{}

			globalCatalogService, err = globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(globalCatalogServiceOptions)
			Expect(err).To(BeNil())
			Expect(globalCatalogService).ToNot(BeNil())
			Expect(globalCatalogService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			globalCatalogService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListCatalogEntries - Returns parent catalog entries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCatalogEntries(listCatalogEntriesOptions *ListCatalogEntriesOptions)`, func() {
			listCatalogEntriesOptions := &globalcatalogv1.ListCatalogEntriesOptions{
				Include:    core.StringPtr("testString"),
				Q:          core.StringPtr("testString"),
				Descending: core.StringPtr("testString"),
				Languages:  core.StringPtr("testString"),
				Catalog:    core.BoolPtr(true),
				Complete:   core.BoolPtr(true),
				Offset:     core.Int64Ptr(int64(38)),
				Limit:      core.Int64Ptr(int64(200)),
			}

			entrySearchResult, response, err := globalCatalogService.ListCatalogEntries(listCatalogEntriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entrySearchResult).ToNot(BeNil())
		})
	})

	Describe(`CreateCatalogEntry - Create a catalog entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCatalogEntry(createCatalogEntryOptions *CreateCatalogEntryOptions)`, func() {
			overviewModel := &globalcatalogv1.Overview{
				DisplayName:         core.StringPtr("testString"),
				LongDescription:     core.StringPtr("testString"),
				Description:         core.StringPtr("testString"),
				FeaturedDescription: core.StringPtr("testString"),
			}

			imageModel := &globalcatalogv1.Image{
				Image:        core.StringPtr("testString"),
				SmallImage:   core.StringPtr("testString"),
				MediumImage:  core.StringPtr("testString"),
				FeatureImage: core.StringPtr("testString"),
			}

			providerModel := &globalcatalogv1.Provider{
				Email:        core.StringPtr("testString@ibm.com"),
				Name:         core.StringPtr("testString"),
				Contact:      core.StringPtr("testString"),
				SupportEmail: core.StringPtr("testString@ibm.com"),
				Phone:        core.StringPtr("testString"),
			}

			cfMetaDataModel := &globalcatalogv1.CfMetaData{
				Type:                         core.StringPtr("testString"),
				IamCompatible:                core.BoolPtr(true),
				UniqueAPIKey:                 core.BoolPtr(true),
				Provisionable:                core.BoolPtr(true),
				Bindable:                     core.BoolPtr(true),
				AsyncProvisioningSupported:   core.BoolPtr(true),
				AsyncUnprovisioningSupported: core.BoolPtr(true),
				Requires:                     []string{"testString"},
				PlanUpdateable:               core.BoolPtr(true),
				State:                        core.StringPtr("testString"),
				ServiceCheckEnabled:          core.BoolPtr(true),
				TestCheckInterval:            core.Int64Ptr(int64(38)),
				ServiceKeySupported:          core.BoolPtr(true),
				CfGUID:                       make(map[string]string),
			}

			planMetaDataModel := &globalcatalogv1.PlanMetaData{
				Bindable:                     core.BoolPtr(true),
				Reservable:                   core.BoolPtr(true),
				AllowInternalUsers:           core.BoolPtr(true),
				AsyncProvisioningSupported:   core.BoolPtr(true),
				AsyncUnprovisioningSupported: core.BoolPtr(true),
				TestCheckInterval:            core.Int64Ptr(int64(38)),
				SingleScopeInstance:          core.StringPtr("testString"),
				ServiceCheckEnabled:          core.BoolPtr(true),
				CfGUID:                       make(map[string]string),
			}

			aliasMetaDataModel := &globalcatalogv1.AliasMetaData{
				Type:   core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
			}

			sourceMetaDataModel := &globalcatalogv1.SourceMetaData{
				Path: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				URL:  core.StringPtr("testString"),
			}

			templateMetaDataModel := &globalcatalogv1.TemplateMetaData{
				Services:             []string{"testString"},
				DefaultMemory:        core.Int64Ptr(int64(38)),
				StartCmd:             core.StringPtr("testString"),
				Source:               sourceMetaDataModel,
				RuntimeCatalogID:     core.StringPtr("testString"),
				CfRuntimeID:          core.StringPtr("testString"),
				TemplateID:           core.StringPtr("testString"),
				ExecutableFile:       core.StringPtr("testString"),
				Buildpack:            core.StringPtr("testString"),
				EnvironmentVariables: make(map[string]string),
			}

			bulletsModel := &globalcatalogv1.Bullets{
				Title:       core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Icon:        core.StringPtr("testString"),
				Quantity:    core.Int64Ptr(int64(38)),
			}

			uiMediaSourceMetaDataModel := &globalcatalogv1.UIMediaSourceMetaData{
				Type: core.StringPtr("testString"),
				URL:  core.StringPtr("testString"),
			}

			uiMetaMediaModel := &globalcatalogv1.UIMetaMedia{
				Caption:      core.StringPtr("testString"),
				ThumbnailURL: core.StringPtr("testString"),
				Type:         core.StringPtr("testString"),
				URL:          core.StringPtr("testString"),
				Source:       []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel},
			}

			stringsModel := &globalcatalogv1.Strings{
				Bullets:              []globalcatalogv1.Bullets{*bulletsModel},
				Media:                []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel},
				NotCreatableMsg:      core.StringPtr("testString"),
				NotCreatableRobotMsg: core.StringPtr("testString"),
				DeprecationWarning:   core.StringPtr("testString"),
				PopupWarningMessage:  core.StringPtr("testString"),
				Instruction:          core.StringPtr("testString"),
			}

			urlsModel := &globalcatalogv1.Urls{
				DocURL:              core.StringPtr("testString"),
				InstructionsURL:     core.StringPtr("testString"),
				APIURL:              core.StringPtr("testString"),
				CreateURL:           core.StringPtr("testString"),
				SdkDownloadURL:      core.StringPtr("testString"),
				TermsURL:            core.StringPtr("testString"),
				CustomCreatePageURL: core.StringPtr("testString"),
				CatalogDetailsURL:   core.StringPtr("testString"),
				DeprecationDocURL:   core.StringPtr("testString"),
				DashboardURL:        core.StringPtr("testString"),
				RegistrationURL:     core.StringPtr("testString"),
				Apidocsurl:          core.StringPtr("testString"),
			}

			uiMetaDataModel := &globalcatalogv1.UIMetaData{
				Strings:                      make(map[string]globalcatalogv1.Strings),
				Urls:                         urlsModel,
				EmbeddableDashboard:          core.StringPtr("testString"),
				EmbeddableDashboardFullWidth: core.BoolPtr(true),
				NavigationOrder:              []string{"testString"},
				NotCreatable:                 core.BoolPtr(true),
				PrimaryOfferingID:            core.StringPtr("testString"),
				AccessibleDuringProvision:    core.BoolPtr(true),
				SideBySideIndex:              core.Int64Ptr(int64(38)),
				EndOfServiceTime:             CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Hidden:                       core.BoolPtr(true),
				HideLiteMetering:             core.BoolPtr(true),
				NoUpgradeNextStep:            core.BoolPtr(true),
			}
			uiMetaDataModel.Strings["en"] = *stringsModel

			drMetaDataModel := &globalcatalogv1.DrMetaData{
				Dr:          core.BoolPtr(true),
				Description: core.StringPtr("testString"),
			}

			slaMetaDataModel := &globalcatalogv1.SLAMetaData{
				Terms:          core.StringPtr("testString"),
				Tenancy:        core.StringPtr("testString"),
				Provisioning:   core.Float64Ptr(float64(72.5)),
				Responsiveness: core.Float64Ptr(float64(72.5)),
				Dr:             drMetaDataModel,
			}

			callbacksModel := &globalcatalogv1.Callbacks{
				ControllerURL:            core.StringPtr("testString"),
				BrokerURL:                core.StringPtr("testString"),
				BrokerProxyURL:           core.StringPtr("testString"),
				DashboardURL:             core.StringPtr("testString"),
				DashboardDataURL:         core.StringPtr("testString"),
				DashboardDetailTabURL:    core.StringPtr("testString"),
				DashboardDetailTabExtURL: core.StringPtr("testString"),
				ServiceMonitorAPI:        core.StringPtr("testString"),
				ServiceMonitorApp:        core.StringPtr("testString"),
				APIEndpoint:              make(map[string]string),
			}

			priceModel := &globalcatalogv1.Price{
				QuantityTier: core.Int64Ptr(int64(38)),
				Price:        core.Float64Ptr(float64(72.5)),
			}

			amountModel := &globalcatalogv1.Amount{
				Country:  core.StringPtr("testString"),
				Currency: core.StringPtr("testString"),
				Prices:   []globalcatalogv1.Price{*priceModel},
			}

			startingPriceModel := &globalcatalogv1.StartingPrice{
				PlanID:       core.StringPtr("testString"),
				DeploymentID: core.StringPtr("testString"),
				Unit:         core.StringPtr("testString"),
				Amount:       []globalcatalogv1.Amount{*amountModel},
			}

			pricingSetModel := &globalcatalogv1.PricingSet{
				Type:          core.StringPtr("testString"),
				Origin:        core.StringPtr("testString"),
				StartingPrice: startingPriceModel,
			}

			brokerModel := &globalcatalogv1.Broker{
				Name: core.StringPtr("testString"),
				GUID: core.StringPtr("testString"),
			}

			deploymentBaseModel := &globalcatalogv1.DeploymentBase{
				Location:            core.StringPtr("testString"),
				LocationURL:         core.StringPtr("testString"),
				OriginalLocation:    core.StringPtr("testString"),
				TargetCRN:           core.StringPtr("testString"),
				ServiceCRN:          core.StringPtr("testString"),
				MccpID:              core.StringPtr("testString"),
				Broker:              brokerModel,
				SupportsRcMigration: core.BoolPtr(true),
				TargetNetwork:       core.StringPtr("testString"),
			}

			objectMetadataSetModel := &globalcatalogv1.ObjectMetadataSet{
				RcCompatible: core.BoolPtr(true),
				Service:      cfMetaDataModel,
				Plan:         planMetaDataModel,
				Alias:        aliasMetaDataModel,
				Template:     templateMetaDataModel,
				UI:           uiMetaDataModel,
				Compliance:   []string{"testString"},
				SLA:          slaMetaDataModel,
				Callbacks:    callbacksModel,
				OriginalName: core.StringPtr("testString"),
				Version:      core.StringPtr("testString"),
				Other:        map[string]interface{}{"anyKey": "anyValue"},
				Pricing:      pricingSetModel,
				Deployment:   deploymentBaseModel,
			}

			createCatalogEntryOptions := &globalcatalogv1.CreateCatalogEntryOptions{
				Name:       core.StringPtr("testString"),
				Kind:       core.StringPtr("service"),
				OverviewUI: make(map[string]globalcatalogv1.Overview),
				Images:     imageModel,
				Disabled:   core.BoolPtr(false),
				Tags:       []string{"testString"},
				Provider:   providerModel,
				ID:         core.StringPtr("testString"),
				Group:      core.BoolPtr(true),
				Active:     core.BoolPtr(true),
				Metadata:   objectMetadataSetModel,
			}
			createCatalogEntryOptions.OverviewUI["en"] = *overviewModel

			catalogEntry, response, err := globalCatalogService.CreateCatalogEntry(createCatalogEntryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(catalogEntry).ToNot(BeNil())
		})
	})

	Describe(`GetCatalogEntry - Get a specific catalog object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalogEntry(getCatalogEntryOptions *GetCatalogEntryOptions)`, func() {
			getCatalogEntryOptions := &globalcatalogv1.GetCatalogEntryOptions{
				ID:        core.StringPtr("testString"),
				Include:   core.StringPtr("testString"),
				Languages: core.StringPtr("testString"),
				Complete:  core.BoolPtr(true),
				Depth:     core.Int64Ptr(int64(38)),
			}

			catalogEntry, response, err := globalCatalogService.GetCatalogEntry(getCatalogEntryOptions)
			fetchedEntry = catalogEntry
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogEntry).ToNot(BeNil())
		})
	})

	Describe(`UpdateCatalogEntry - Update a catalog entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateCatalogEntry(updateCatalogEntryOptions *UpdateCatalogEntryOptions)`, func() {
			overviewModel := &globalcatalogv1.Overview{
				DisplayName:         core.StringPtr("testString"),
				LongDescription:     core.StringPtr("testString"),
				Description:         core.StringPtr("testString"),
				FeaturedDescription: core.StringPtr("testString"),
			}

			imageModel := &globalcatalogv1.Image{
				Image:        core.StringPtr("testString"),
				SmallImage:   core.StringPtr("testString"),
				MediumImage:  core.StringPtr("testString"),
				FeatureImage: core.StringPtr("testString"),
			}

			providerModel := &globalcatalogv1.Provider{
				Email:        core.StringPtr("testString@ibm.com"),
				Name:         core.StringPtr("testString"),
				Contact:      core.StringPtr("testString"),
				SupportEmail: core.StringPtr("testString@ibm.com"),
				Phone:        core.StringPtr("testString"),
			}

			cfMetaDataModel := &globalcatalogv1.CfMetaData{
				Type:                         core.StringPtr("testString"),
				IamCompatible:                core.BoolPtr(true),
				UniqueAPIKey:                 core.BoolPtr(true),
				Provisionable:                core.BoolPtr(true),
				Bindable:                     core.BoolPtr(true),
				AsyncProvisioningSupported:   core.BoolPtr(true),
				AsyncUnprovisioningSupported: core.BoolPtr(true),
				Requires:                     []string{"testString"},
				PlanUpdateable:               core.BoolPtr(true),
				State:                        core.StringPtr("testString"),
				ServiceCheckEnabled:          core.BoolPtr(true),
				TestCheckInterval:            core.Int64Ptr(int64(38)),
				ServiceKeySupported:          core.BoolPtr(true),
				CfGUID:                       make(map[string]string),
			}

			planMetaDataModel := &globalcatalogv1.PlanMetaData{
				Bindable:                     core.BoolPtr(true),
				Reservable:                   core.BoolPtr(true),
				AllowInternalUsers:           core.BoolPtr(true),
				AsyncProvisioningSupported:   core.BoolPtr(true),
				AsyncUnprovisioningSupported: core.BoolPtr(true),
				TestCheckInterval:            core.Int64Ptr(int64(38)),
				SingleScopeInstance:          core.StringPtr("testString"),
				ServiceCheckEnabled:          core.BoolPtr(true),
				CfGUID:                       make(map[string]string),
			}

			aliasMetaDataModel := &globalcatalogv1.AliasMetaData{
				Type:   core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
			}

			sourceMetaDataModel := &globalcatalogv1.SourceMetaData{
				Path: core.StringPtr("testString"),
				Type: core.StringPtr("testString"),
				URL:  core.StringPtr("testString"),
			}

			templateMetaDataModel := &globalcatalogv1.TemplateMetaData{
				Services:             []string{"testString"},
				DefaultMemory:        core.Int64Ptr(int64(38)),
				StartCmd:             core.StringPtr("testString"),
				Source:               sourceMetaDataModel,
				RuntimeCatalogID:     core.StringPtr("testString"),
				CfRuntimeID:          core.StringPtr("testString"),
				TemplateID:           core.StringPtr("testString"),
				ExecutableFile:       core.StringPtr("testString"),
				Buildpack:            core.StringPtr("testString"),
				EnvironmentVariables: make(map[string]string),
			}

			bulletsModel := &globalcatalogv1.Bullets{
				Title:       core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Icon:        core.StringPtr("testString"),
				Quantity:    core.Int64Ptr(int64(38)),
			}

			uiMediaSourceMetaDataModel := &globalcatalogv1.UIMediaSourceMetaData{
				Type: core.StringPtr("testString"),
				URL:  core.StringPtr("testString"),
			}

			uiMetaMediaModel := &globalcatalogv1.UIMetaMedia{
				Caption:      core.StringPtr("testString"),
				ThumbnailURL: core.StringPtr("testString"),
				Type:         core.StringPtr("testString"),
				URL:          core.StringPtr("testString"),
				Source:       []globalcatalogv1.UIMediaSourceMetaData{*uiMediaSourceMetaDataModel},
			}

			stringsModel := &globalcatalogv1.Strings{
				Bullets:              []globalcatalogv1.Bullets{*bulletsModel},
				Media:                []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel},
				NotCreatableMsg:      core.StringPtr("testString"),
				NotCreatableRobotMsg: core.StringPtr("testString"),
				DeprecationWarning:   core.StringPtr("testString"),
				PopupWarningMessage:  core.StringPtr("testString"),
				Instruction:          core.StringPtr("testString"),
			}

			urlsModel := &globalcatalogv1.Urls{
				DocURL:              core.StringPtr("testString"),
				InstructionsURL:     core.StringPtr("testString"),
				APIURL:              core.StringPtr("testString"),
				CreateURL:           core.StringPtr("testString"),
				SdkDownloadURL:      core.StringPtr("testString"),
				TermsURL:            core.StringPtr("testString"),
				CustomCreatePageURL: core.StringPtr("testString"),
				CatalogDetailsURL:   core.StringPtr("testString"),
				DeprecationDocURL:   core.StringPtr("testString"),
				DashboardURL:        core.StringPtr("testString"),
				RegistrationURL:     core.StringPtr("testString"),
				Apidocsurl:          core.StringPtr("testString"),
			}

			uiMetaDataModel := &globalcatalogv1.UIMetaData{
				Strings:                      make(map[string]globalcatalogv1.Strings),
				Urls:                         urlsModel,
				EmbeddableDashboard:          core.StringPtr("testString"),
				EmbeddableDashboardFullWidth: core.BoolPtr(true),
				NavigationOrder:              []string{"testString"},
				NotCreatable:                 core.BoolPtr(true),
				PrimaryOfferingID:            core.StringPtr("testString"),
				AccessibleDuringProvision:    core.BoolPtr(true),
				SideBySideIndex:              core.Int64Ptr(int64(38)),
				EndOfServiceTime:             CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Hidden:                       core.BoolPtr(true),
				HideLiteMetering:             core.BoolPtr(true),
				NoUpgradeNextStep:            core.BoolPtr(true),
			}
			uiMetaDataModel.Strings["en"] = *stringsModel

			drMetaDataModel := &globalcatalogv1.DrMetaData{
				Dr:          core.BoolPtr(true),
				Description: core.StringPtr("testString"),
			}

			slaMetaDataModel := &globalcatalogv1.SLAMetaData{
				Terms:          core.StringPtr("testString"),
				Tenancy:        core.StringPtr("testString"),
				Provisioning:   core.Float64Ptr(float64(72.5)),
				Responsiveness: core.Float64Ptr(float64(72.5)),
				Dr:             drMetaDataModel,
			}

			callbacksModel := &globalcatalogv1.Callbacks{
				ControllerURL:            core.StringPtr("testString"),
				BrokerURL:                core.StringPtr("testString"),
				BrokerProxyURL:           core.StringPtr("testString"),
				DashboardURL:             core.StringPtr("testString"),
				DashboardDataURL:         core.StringPtr("testString"),
				DashboardDetailTabURL:    core.StringPtr("testString"),
				DashboardDetailTabExtURL: core.StringPtr("testString"),
				ServiceMonitorAPI:        core.StringPtr("testString"),
				ServiceMonitorApp:        core.StringPtr("testString"),
				APIEndpoint:              make(map[string]string),
			}

			priceModel := &globalcatalogv1.Price{
				QuantityTier: core.Int64Ptr(int64(38)),
				Price:        core.Float64Ptr(float64(72.5)),
			}

			amountModel := &globalcatalogv1.Amount{
				Country:  core.StringPtr("testString"),
				Currency: core.StringPtr("testString"),
				Prices:   []globalcatalogv1.Price{*priceModel},
			}

			startingPriceModel := &globalcatalogv1.StartingPrice{
				PlanID:       core.StringPtr("testString"),
				DeploymentID: core.StringPtr("testString"),
				Unit:         core.StringPtr("testString"),
				Amount:       []globalcatalogv1.Amount{*amountModel},
			}

			pricingSetModel := &globalcatalogv1.PricingSet{
				Type:          core.StringPtr("testString"),
				Origin:        core.StringPtr("testString"),
				StartingPrice: startingPriceModel,
			}

			brokerModel := &globalcatalogv1.Broker{
				Name: core.StringPtr("testString"),
				GUID: core.StringPtr("testString"),
			}

			deploymentBaseModel := &globalcatalogv1.DeploymentBase{
				Location:            core.StringPtr("testString"),
				LocationURL:         core.StringPtr("testString"),
				OriginalLocation:    core.StringPtr("testString"),
				TargetCRN:           core.StringPtr("testString"),
				ServiceCRN:          core.StringPtr("testString"),
				MccpID:              core.StringPtr("testString"),
				Broker:              brokerModel,
				SupportsRcMigration: core.BoolPtr(true),
				TargetNetwork:       core.StringPtr("testString"),
			}

			objectMetadataSetModel := &globalcatalogv1.ObjectMetadataSet{
				RcCompatible: core.BoolPtr(true),
				Service:      cfMetaDataModel,
				Plan:         planMetaDataModel,
				Alias:        aliasMetaDataModel,
				Template:     templateMetaDataModel,
				UI:           uiMetaDataModel,
				Compliance:   []string{"testString"},
				SLA:          slaMetaDataModel,
				Callbacks:    callbacksModel,
				OriginalName: core.StringPtr("testString"),
				Version:      core.StringPtr("testString"),
				Other:        map[string]interface{}{"anyKey": "anyValue"},
				Pricing:      pricingSetModel,
				Deployment:   deploymentBaseModel,
			}

			updateCatalogEntryOptions := &globalcatalogv1.UpdateCatalogEntryOptions{
				ID:         core.StringPtr("testString"),
				Name:       core.StringPtr("testString"),
				Kind:       core.StringPtr("service"),
				OverviewUI: make(map[string]globalcatalogv1.Overview),
				Images:     imageModel,
				Disabled:   core.BoolPtr(false),
				Tags:       []string{"testString"},
				Provider:   providerModel,
				Group:      core.BoolPtr(true),
				Active:     core.BoolPtr(true),
				Metadata:   objectMetadataSetModel,
				Move:       core.StringPtr("testString"),
				URL:        fetchedEntry.URL,
			}
			updateCatalogEntryOptions.OverviewUI["en"] = *overviewModel

			catalogEntry, response, err := globalCatalogService.UpdateCatalogEntry(updateCatalogEntryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogEntry).ToNot(BeNil())
		})
	})

	Describe(`GetChildObjects - Get child catalog entries of a specific kind`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetChildObjects(getChildObjectsOptions *GetChildObjectsOptions)`, func() {
			getChildObjectsOptions := &globalcatalogv1.GetChildObjectsOptions{
				ID:         core.StringPtr("testString"),
				Kind:       core.StringPtr("testString"),
				Include:    core.StringPtr("testString"),
				Q:          core.StringPtr("testString"),
				Descending: core.StringPtr("testString"),
				Languages:  core.StringPtr("testString"),
				Complete:   core.BoolPtr(true),
				Offset:     core.Int64Ptr(int64(38)),
				Limit:      core.Int64Ptr(int64(200)),
			}

			entrySearchResult, response, err := globalCatalogService.GetChildObjects(getChildObjectsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(entrySearchResult).ToNot(BeNil())
		})
	})

	Describe(`RestoreCatalogEntry - Restore archived catalog entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestoreCatalogEntry(restoreCatalogEntryOptions *RestoreCatalogEntryOptions)`, func() {
			restoreCatalogEntryOptions := &globalcatalogv1.RestoreCatalogEntryOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := globalCatalogService.RestoreCatalogEntry(restoreCatalogEntryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetVisibility - Get the visibility constraints for an object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVisibility(getVisibilityOptions *GetVisibilityOptions)`, func() {
			getVisibilityOptions := &globalcatalogv1.GetVisibilityOptions{
				ID: core.StringPtr("testString"),
			}

			visibility, response, err := globalCatalogService.GetVisibility(getVisibilityOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(visibility).ToNot(BeNil())
		})
	})

	Describe(`UpdateVisibility - Update visibility`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVisibility(updateVisibilityOptions *UpdateVisibilityOptions)`, func() {
			Skip("Not testing")
			visibilityDetailAccountsModel := &globalcatalogv1.VisibilityDetailAccounts{
				Accountid: core.StringPtr("testString"),
			}

			visibilityDetailModel := &globalcatalogv1.VisibilityDetail{
				Accounts: visibilityDetailAccountsModel,
			}

			updateVisibilityOptions := &globalcatalogv1.UpdateVisibilityOptions{
				ID:           core.StringPtr("testString"),
				Restrictions: core.StringPtr("private"),
				Extendable:   core.BoolPtr(true),
				Include:      visibilityDetailModel,
				Exclude:      visibilityDetailModel,
			}

			response, err := globalCatalogService.UpdateVisibility(updateVisibilityOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetPricing - Get the pricing for an object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPricing(getPricingOptions *GetPricingOptions)`, func() {
			getPricingOptions := &globalcatalogv1.GetPricingOptions{
				ID: core.StringPtr("testString"),
			}

			pricingGet, response, err := globalCatalogService.GetPricing(getPricingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pricingGet).ToNot(BeNil())
		})
	})

	Describe(`GetPricingDeployments - Get the pricing deployments for a plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPricingDeployments(getPricingDeploymentsOptions *GetPricingDeploymentsOptions)`, func() {
			getPricingDeploymentsOptions := &globalcatalogv1.GetPricingDeploymentsOptions{
				ID: core.StringPtr("testString"),
			}

			pricingSearchResult, response, err := globalCatalogService.GetPricingDeployments(getPricingDeploymentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pricingSearchResult).ToNot(BeNil())
		})
	})

	Describe(`GetAuditLogs - Get the audit logs for an object`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAuditLogs(getAuditLogsOptions *GetAuditLogsOptions)`, func() {
			Skip("Not testing")
			getAuditLogsOptions := &globalcatalogv1.GetAuditLogsOptions{
				ID:        core.StringPtr("testString"),
				Ascending: core.StringPtr("false"),
				Startat:   core.StringPtr("testString"),
				Offset:    core.Int64Ptr(int64(38)),
				Limit:     core.Int64Ptr(int64(200)),
			}

			auditSearchResult, response, err := globalCatalogService.GetAuditLogs(getAuditLogsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(auditSearchResult).ToNot(BeNil())
		})
	})

	Describe(`ListArtifacts - Get artifacts`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListArtifacts(listArtifactsOptions *ListArtifactsOptions)`, func() {
			Skip("Not testing")
			listArtifactsOptions := &globalcatalogv1.ListArtifactsOptions{
				ObjectID: core.StringPtr("testString"),
			}

			artifacts, response, err := globalCatalogService.ListArtifacts(listArtifactsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(artifacts).ToNot(BeNil())
		})
	})

	Describe(`GetArtifact - Get artifact`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetArtifact(getArtifactOptions *GetArtifactOptions)`, func() {
			Skip("Not testing")
			getArtifactOptions := &globalcatalogv1.GetArtifactOptions{
				ObjectID:   core.StringPtr("testString"),
				ArtifactID: core.StringPtr("testString"),
				Accept:     core.StringPtr("testString"),
			}

			result, response, err := globalCatalogService.GetArtifact(getArtifactOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`UploadArtifact - Upload artifact`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UploadArtifact(uploadArtifactOptions *UploadArtifactOptions)`, func() {
			Skip("Not testing")
			uploadArtifactOptions := &globalcatalogv1.UploadArtifactOptions{
				ObjectID:    core.StringPtr("testString"),
				ArtifactID:  core.StringPtr("testString"),
				Artifact:    CreateMockReader("This is a mock file."),
				ContentType: core.StringPtr("testString"),
			}

			response, err := globalCatalogService.UploadArtifact(uploadArtifactOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`DeleteCatalogEntry - Delete a catalog entry`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCatalogEntry(deleteCatalogEntryOptions *DeleteCatalogEntryOptions)`, func() {
			deleteCatalogEntryOptions := &globalcatalogv1.DeleteCatalogEntryOptions{
				ID:    core.StringPtr("testString"),
				Force: core.BoolPtr(true),
			}

			response, err := globalCatalogService.DeleteCatalogEntry(deleteCatalogEntryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`DeleteArtifact - Delete artifact`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteArtifact(deleteArtifactOptions *DeleteArtifactOptions)`, func() {
			Skip("Not testing")
			deleteArtifactOptions := &globalcatalogv1.DeleteArtifactOptions{
				ObjectID:   core.StringPtr("testString"),
				ArtifactID: core.StringPtr("testString"),
			}

			response, err := globalCatalogService.DeleteArtifact(deleteArtifactOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
