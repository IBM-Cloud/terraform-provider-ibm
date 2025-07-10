//go:build integration

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

package partnercentersellv1_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the partnercentersellv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PartnerCenterSellV1 Integration Tests`, func() {
	const externalConfigFile = "../partner_center_sell_v1.env"

	var (
		err                         error
		partnerCenterSellService    *partnercentersellv1.PartnerCenterSellV1
		partnerCenterSellServiceAlt *partnercentersellv1.PartnerCenterSellV1
		serviceURL                  string
		config                      map[string]string

		// Variables to hold link values
		brokerIdLink                          string
		catalogDeploymentIdLink               string
		catalogPlanIdLink                     string
		catalogProductIdLink                  string
		productIdLink                         string
		programmaticNameLink                  string
		registrationIdLink                    string
		accountId                             string
		badgeId                               string
		iamServiceRegistrationId              string
		productIdWithApprovedProgrammaticName string
		env                                   string
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
			config, err = core.GetServiceProperties(partnercentersellv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			accountId = config["ACCOUNT_ID"]
			Expect(accountId).ToNot(BeEmpty())

			badgeId = config["BADGE_ID"]
			Expect(badgeId).ToNot(BeEmpty())

			iamServiceRegistrationId = config["IAM_REGISTRATION_ID"]
			Expect(iamServiceRegistrationId).ToNot(BeEmpty())

			productIdWithApprovedProgrammaticName = config["PRODUCT_ID_APPROVED"]
			Expect(productIdWithApprovedProgrammaticName).ToNot(BeEmpty())

			env = "current"

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			partnerCenterSellServiceOptions := &partnercentersellv1.PartnerCenterSellV1Options{}

			partnerCenterSellService, err = partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(partnerCenterSellServiceOptions)
			Expect(err).To(BeNil())
			Expect(partnerCenterSellService).ToNot(BeNil())
			Expect(partnerCenterSellService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			partnerCenterSellService.EnableRetries(4, 30*time.Second)
		})
		It("Successfully construct the service client instance with alternative credentials", func() {
			var err error

			// begin-common

			partnerCenterSellServiceOptions := &partnercentersellv1.PartnerCenterSellV1Options{
				ServiceName: "partner_center_sell_alt",
			}

			partnerCenterSellServiceAlt, err = partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(partnerCenterSellServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(partnerCenterSellServiceAlt).ToNot(BeNil())
		})
	})

	Describe(`CreateRegistration - Register your account in Partner Center - Sell`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRegistration(createRegistrationOptions *CreateRegistrationOptions)`, func() {
			primaryContactModel := &partnercentersellv1.PrimaryContact{
				Name:  core.StringPtr("Company Representative"),
				Email: core.StringPtr("companyrep@email.com"),
			}

			createRegistrationOptions := &partnercentersellv1.CreateRegistrationOptions{
				AccountID:      core.StringPtr(accountId),
				CompanyName:    core.StringPtr("company_sdk"),
				PrimaryContact: primaryContactModel,
			}

			registration, response, err := partnerCenterSellServiceAlt.CreateRegistration(createRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registration).ToNot(BeNil())

			registrationIdLink = *registration.ID
			fmt.Fprintf(GinkgoWriter, "Saved registrationIdLink value: %v\n", registrationIdLink)
		})
	})

	Describe(`CreateOnboardingProduct - Create a product to onboard`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateOnboardingProduct(createOnboardingProductOptions *CreateOnboardingProductOptions)`, func() {
			primaryContactModel := &partnercentersellv1.PrimaryContact{
				Name:  core.StringPtr("name"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			onboardingProductSupportEscalationContactItemsModel := &partnercentersellv1.OnboardingProductSupportEscalationContactItems{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
				Role:  core.StringPtr("admin"),
			}

			onboardingProductSupportModel := &partnercentersellv1.OnboardingProductSupport{
				EscalationContacts: []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel},
			}

			createOnboardingProductOptions := &partnercentersellv1.CreateOnboardingProductOptions{
				Type:           core.StringPtr("service"),
				PrimaryContact: primaryContactModel,
				EccnNumber:     core.StringPtr("5D002.C.1"),
				EroClass:       core.StringPtr("A6VR"),
				Unspsc:         core.Float64Ptr(25191503),
				TaxAssessment:  core.StringPtr("PAAS"),
				Support:        onboardingProductSupportModel,
			}

			onboardingProduct, response, err := partnerCenterSellServiceAlt.CreateOnboardingProduct(createOnboardingProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(onboardingProduct).ToNot(BeNil())

			productIdLink = *onboardingProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved productIdLink value: %v\n", productIdLink)
		})
	})

	Describe(`UpdateOnboardingProduct - Update an onboarding product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateOnboardingProduct(updateOnboardingProductOptions *UpdateOnboardingProductOptions)`, func() {
			primaryContactModel := &partnercentersellv1.PrimaryContact{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
			}

			onboardingProductSupportEscalationContactItemsModel := &partnercentersellv1.OnboardingProductSupportEscalationContactItems{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
				Role:  core.StringPtr("admin"),
			}

			onboardingProductSupportModel := &partnercentersellv1.OnboardingProductSupport{
				EscalationContacts: []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel},
			}

			onboardingProductPatchModel := &partnercentersellv1.OnboardingProductPatch{
				PrimaryContact: primaryContactModel,
				EccnNumber:     core.StringPtr("5D002.C.1"),
				EroClass:       core.StringPtr("A6VR"),
				Unspsc:         core.Float64Ptr(25191503),
				TaxAssessment:  core.StringPtr("PAAS"),
				Support:        onboardingProductSupportModel,
			}
			onboardingProductPatchModelAsPatch, asPatchErr := onboardingProductPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateOnboardingProductOptions := &partnercentersellv1.UpdateOnboardingProductOptions{
				ProductID:              &productIdLink,
				OnboardingProductPatch: onboardingProductPatchModelAsPatch,
			}

			onboardingProduct, response, err := partnerCenterSellServiceAlt.UpdateOnboardingProduct(updateOnboardingProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(onboardingProduct).ToNot(BeNil())

			productIdLink = *onboardingProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved productIdLink value: %v\n", productIdLink)
		})
	})

	Describe(`CreateCatalogProduct - Create a global catalog product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCatalogProduct(createCatalogProductOptions *CreateCatalogProductOptions)`, func() {
			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("IBM"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName:     core.StringPtr("My product display name."),
				Description:     core.StringPtr("My product description."),
				LongDescription: core.StringPtr("My product description long description."),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			globalCatalogProductImagesModel := &partnercentersellv1.GlobalCatalogProductImages{
				Image: core.StringPtr("https://http.cat/images/100.jpg"),
			}

			catalogHighlightItemModel := &partnercentersellv1.CatalogHighlightItem{
				Description: core.StringPtr("highlight desc"),
				Title:       core.StringPtr("Title"),
			}

			catalogProductMediaItemModel := &partnercentersellv1.CatalogProductMediaItem{
				Caption:   core.StringPtr("testString"),
				Thumbnail: core.StringPtr("testString"),
				Type:      core.StringPtr("image"),
				URL:       core.StringPtr("https://http.cat/images/100.jpg"),
			}

			globalCatalogMetadataUiNavigationItemModel := &partnercentersellv1.GlobalCatalogMetadataUINavigationItem{
				ID:    core.StringPtr("testString"),
				URL:   core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
			}

			globalCatalogMetadataUiStringsContentModel := &partnercentersellv1.GlobalCatalogMetadataUIStringsContent{
				Bullets:         []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel},
				Media:           []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel},
				NavigationItems: []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel},
			}

			globalCatalogMetadataUiStringsModel := &partnercentersellv1.GlobalCatalogMetadataUIStrings{
				En: globalCatalogMetadataUiStringsContentModel,
			}

			globalCatalogMetadataUiUrlsModel := &partnercentersellv1.GlobalCatalogMetadataUIUrls{
				DocURL:              core.StringPtr("https://http.cat/doc"),
				ApidocsURL:          core.StringPtr("https://http.cat/doc"),
				TermsURL:            core.StringPtr("https://http.cat/doc"),
				InstructionsURL:     core.StringPtr("https://http.cat/doc"),
				CatalogDetailsURL:   core.StringPtr("https://http.cat/doc"),
				CustomCreatePageURL: core.StringPtr("https://http.cat/doc"),
				Dashboard:           core.StringPtr("https://http.cat/doc"),
			}

			globalCatalogProductMetadataUiModel := &partnercentersellv1.GlobalCatalogProductMetadataUI{
				Strings:                   globalCatalogMetadataUiStringsModel,
				Urls:                      globalCatalogMetadataUiUrlsModel,
				Hidden:                    core.BoolPtr(true),
				SideBySideIndex:           core.Float64Ptr(float64(72)),
				EmbeddableDashboard:       core.StringPtr("https://http.cat/emb"),
				AccessibleDuringProvision: core.BoolPtr(true),
				PrimaryOfferingID:         core.StringPtr("offeringId"),
			}

			globalCatalogMetadataServiceCustomParametersI18nFieldsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields{
				Displayname: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			globalCatalogMetadataServiceCustomParametersI18nModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n{
				En:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				De:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Es:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Fr:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				It:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Ja:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Ko:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				PtBr: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				ZhTw: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				ZhCn: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
			}

			globalCatalogMetadataServiceCustomParametersOptionsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{
				Displayname: core.StringPtr("testString"),
				Value:       core.StringPtr("testString"),
				I18n:        globalCatalogMetadataServiceCustomParametersI18nModel,
			}

			globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan{
				ShowFor:        []string{"fake-plan-id"},
				OptionsRefresh: core.BoolPtr(true),
			}

			globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{
				Name:           core.StringPtr("fake-parameter"),
				ShowFor:        []string{"1"},
				OptionsRefresh: core.BoolPtr(true),
			}

			globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation{
				ShowFor: []string{"eu-gb"},
			}

			globalCatalogMetadataServiceCustomParametersAssociationsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations{
				Plan:       globalCatalogMetadataServiceCustomParametersAssociationsPlanModel,
				Parameters: []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel},
				Location:   globalCatalogMetadataServiceCustomParametersAssociationsLocationModel,
			}

			globalCatalogMetadataServiceCustomParametersModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{
				Displayname:    core.StringPtr("Display_name"),
				Name:           core.StringPtr("Sample_name"),
				Type:           core.StringPtr("text"),
				Options:        []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel},
				Value:          []string{"sample_value"},
				Layout:         core.StringPtr("sample_layout"),
				Associations:   globalCatalogMetadataServiceCustomParametersAssociationsModel,
				ValidationURL:  core.StringPtr("https://http.cat/valid"),
				OptionsURL:     core.StringPtr("https://http.cat/option"),
				Invalidmessage: core.StringPtr("Invalid message"),
				Description:    core.StringPtr("Sample description"),
				Required:       core.BoolPtr(true),
				Pattern:        core.StringPtr("."),
				Placeholder:    core.StringPtr("Placeholder"),
				Readonly:       core.BoolPtr(false),
				Hidden:         core.BoolPtr(false),
				I18n:           globalCatalogMetadataServiceCustomParametersI18nModel,
			}

			globalCatalogProductMetadataServicePrototypePatchModel := &partnercentersellv1.GlobalCatalogProductMetadataServicePrototypePatch{
				RcProvisionable:               core.BoolPtr(true),
				IamCompatible:                 core.BoolPtr(true),
				ServiceKeySupported:           core.BoolPtr(true),
				UniqueApiKey:                  core.BoolPtr(true),
				AsyncProvisioningSupported:    core.BoolPtr(true),
				AsyncUnprovisioningSupported:  core.BoolPtr(true),
				CustomCreatePageHybridEnabled: core.BoolPtr(true),
				Parameters:                    []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel},
			}

			supportTimeIntervalModel := &partnercentersellv1.SupportTimeInterval{
				Value: core.Float64Ptr(float64(72)),
				Type:  core.StringPtr("testString"),
			}

			supportEscalationModel := &partnercentersellv1.SupportEscalation{
				Contact:            core.StringPtr("testString"),
				EscalationWaitTime: supportTimeIntervalModel,
				ResponseWaitTime:   supportTimeIntervalModel,
			}

			supportDetailsItemAvailabilityTimeModel := &partnercentersellv1.SupportDetailsItemAvailabilityTime{
				Day:       core.Float64Ptr(float64(72)),
				StartTime: core.StringPtr("10:00"),
				EndTime:   core.StringPtr("18:00"),
			}

			supportDetailsItemAvailabilityModel := &partnercentersellv1.SupportDetailsItemAvailability{
				Times:           []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel},
				Timezone:        core.StringPtr("testString"),
				AlwaysAvailable: core.BoolPtr(true),
			}

			supportDetailsItemModel := &partnercentersellv1.SupportDetailsItem{
				Type:             core.StringPtr("support_site"),
				Contact:          core.StringPtr("testString"),
				ResponseWaitTime: supportTimeIntervalModel,
				Availability:     supportDetailsItemAvailabilityModel,
			}

			globalCatalogProductMetadataOtherPcSupportModel := &partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport{
				URL:               core.StringPtr("https://http.cat/"),
				StatusURL:         core.StringPtr("https://http.cat/status"),
				Locations:         []string{"hu"},
				Languages:         []string{"hu"},
				Process:           core.StringPtr("testString"),
				ProcessI18n:       map[string]string{"anyKey": "anyValue"},
				SupportType:       core.StringPtr("community"),
				SupportEscalation: supportEscalationModel,
				SupportDetails:    []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel},
			}

			globalCatalogProductMetadataOtherPcModel := &partnercentersellv1.GlobalCatalogProductMetadataOtherPC{
				Support: globalCatalogProductMetadataOtherPcSupportModel,
			}

			globalCatalogProductMetadataOtherCompositeChildModel := &partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{
				Kind: core.StringPtr("service"),
				Name: core.StringPtr("test.string"),
			}

			globalCatalogProductMetadataOtherCompositeModel := &partnercentersellv1.GlobalCatalogProductMetadataOtherComposite{
				CompositeKind: core.StringPtr("service"),
				CompositeTag:  core.StringPtr("test.string"),
				Children:      []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel},
			}

			globalCatalogProductMetadataOtherModel := &partnercentersellv1.GlobalCatalogProductMetadataOther{
				PC:        globalCatalogProductMetadataOtherPcModel,
				Composite: globalCatalogProductMetadataOtherCompositeModel,
			}

			globalCatalogProductMetadataPrototypePatchModel := &partnercentersellv1.GlobalCatalogProductMetadataPrototypePatch{
				RcCompatible: core.BoolPtr(true),
				Ui:           globalCatalogProductMetadataUiModel,
				Service:      globalCatalogProductMetadataServicePrototypePatchModel,
				Other:        globalCatalogProductMetadataOtherModel,
			}

			var randomInteger = strconv.Itoa(rand.Intn(1000))
			objectId := fmt.Sprintf("random-id-%s", randomInteger)
			randomName := fmt.Sprintf("random-name-%s", randomInteger)

			createCatalogProductOptions := &partnercentersellv1.CreateCatalogProductOptions{
				ProductID:      core.StringPtr(productIdWithApprovedProgrammaticName),
				Name:           core.StringPtr(randomName),
				Active:         core.BoolPtr(true),
				Disabled:       core.BoolPtr(false),
				Kind:           core.StringPtr("service"),
				Tags:           []string{"keyword", "support_ibm"},
				ObjectProvider: catalogProductProviderModel,
				ObjectID:       core.StringPtr(objectId),
				OverviewUi:     globalCatalogOverviewUiModel,
				Images:         globalCatalogProductImagesModel,
				Metadata:       globalCatalogProductMetadataPrototypePatchModel,
				Env:            core.StringPtr(env),
			}

			globalCatalogProduct, response, err := partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogProduct).ToNot(BeNil())

			catalogProductIdLink = *globalCatalogProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogProductIdLink value: %v\n", catalogProductIdLink)
		})
	})

	Describe(`UpdateCatalogProduct - Update a global catalog product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateCatalogProduct(updateCatalogProductOptions *UpdateCatalogProductOptions)`, func() {
			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName:     core.StringPtr("test product up"),
				Description:     core.StringPtr("test product desc up"),
				LongDescription: core.StringPtr("test product desc long up"),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			globalCatalogProductImagesModel := &partnercentersellv1.GlobalCatalogProductImages{
				Image: core.StringPtr("testString"),
			}

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
			}

			catalogHighlightItemModel := &partnercentersellv1.CatalogHighlightItem{
				Description: core.StringPtr("testString"),
				Title:       core.StringPtr("testString"),
			}

			catalogProductMediaItemModel := &partnercentersellv1.CatalogProductMediaItem{
				Caption:   core.StringPtr("testString"),
				Thumbnail: core.StringPtr("testString"),
				Type:      core.StringPtr("image"),
				URL:       core.StringPtr("https://http.cat/images/200.jpg"),
			}

			globalCatalogMetadataUiNavigationItemModel := &partnercentersellv1.GlobalCatalogMetadataUINavigationItem{
				ID:    core.StringPtr("testString"),
				URL:   core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
			}

			globalCatalogMetadataUiStringsContentModel := &partnercentersellv1.GlobalCatalogMetadataUIStringsContent{
				Bullets:         []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel},
				Media:           []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel},
				NavigationItems: []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel},
			}

			globalCatalogMetadataUiStringsModel := &partnercentersellv1.GlobalCatalogMetadataUIStrings{
				En: globalCatalogMetadataUiStringsContentModel,
			}

			globalCatalogMetadataUiUrlsModel := &partnercentersellv1.GlobalCatalogMetadataUIUrls{
				DocURL:              core.StringPtr("https://http.cat/elua"),
				ApidocsURL:          core.StringPtr("https://http.cat/elua"),
				TermsURL:            core.StringPtr("https://http.cat/elua"),
				InstructionsURL:     core.StringPtr("https://http.cat/elua"),
				CatalogDetailsURL:   core.StringPtr("https://http.cat/elua"),
				CustomCreatePageURL: core.StringPtr("https://http.cat/elua"),
				Dashboard:           core.StringPtr("https://http.cat/elua"),
			}

			globalCatalogProductMetadataUiModel := &partnercentersellv1.GlobalCatalogProductMetadataUI{
				Strings:                   globalCatalogMetadataUiStringsModel,
				Urls:                      globalCatalogMetadataUiUrlsModel,
				Hidden:                    core.BoolPtr(false),
				SideBySideIndex:           core.Float64Ptr(float64(72)),
				EmbeddableDashboard:       core.StringPtr("https://http.cat/emb"),
				AccessibleDuringProvision: core.BoolPtr(true),
				PrimaryOfferingID:         core.StringPtr("offeringId"),
			}

			globalCatalogMetadataServiceCustomParametersI18nFieldsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields{
				Displayname: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			globalCatalogMetadataServiceCustomParametersI18nModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n{
				En:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				De:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Es:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Fr:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				It:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Ja:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Ko:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				PtBr: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				ZhTw: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				ZhCn: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
			}

			globalCatalogMetadataServiceCustomParametersOptionsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{
				Displayname: core.StringPtr("testString"),
				Value:       core.StringPtr("testString"),
				I18n:        globalCatalogMetadataServiceCustomParametersI18nModel,
			}

			globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan{
				ShowFor:        []string{"fake-plan-i-2"},
				OptionsRefresh: core.BoolPtr(true),
			}

			globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{
				Name:           core.StringPtr("fake-parameter-2"),
				ShowFor:        []string{"2"},
				OptionsRefresh: core.BoolPtr(true),
			}

			globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation{
				ShowFor: []string{"us-east"},
			}

			globalCatalogMetadataServiceCustomParametersAssociationsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations{
				Plan:       globalCatalogMetadataServiceCustomParametersAssociationsPlanModel,
				Parameters: []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel},
				Location:   globalCatalogMetadataServiceCustomParametersAssociationsLocationModel,
			}

			globalCatalogMetadataServiceCustomParametersModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{
				Displayname:    core.StringPtr("Display_name"),
				Name:           core.StringPtr("Sample_name"),
				Type:           core.StringPtr("text"),
				Options:        []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel},
				Value:          []string{"sample_value"},
				Layout:         core.StringPtr("sample_layout"),
				Associations:   globalCatalogMetadataServiceCustomParametersAssociationsModel,
				ValidationURL:  core.StringPtr("https://http.cat/valid"),
				OptionsURL:     core.StringPtr("https://http.cat/option"),
				Invalidmessage: core.StringPtr("Invalid message"),
				Description:    core.StringPtr("Sample description"),
				Required:       core.BoolPtr(true),
				Pattern:        core.StringPtr("."),
				Placeholder:    core.StringPtr("Placeholder"),
				Readonly:       core.BoolPtr(false),
				Hidden:         core.BoolPtr(false),
				I18n:           globalCatalogMetadataServiceCustomParametersI18nModel,
			}

			globalCatalogProductMetadataServicePrototypePatchModel := &partnercentersellv1.GlobalCatalogProductMetadataServicePrototypePatch{
				RcProvisionable:               core.BoolPtr(true),
				IamCompatible:                 core.BoolPtr(true),
				ServiceKeySupported:           core.BoolPtr(true),
				UniqueApiKey:                  core.BoolPtr(true),
				AsyncProvisioningSupported:    core.BoolPtr(true),
				AsyncUnprovisioningSupported:  core.BoolPtr(true),
				CustomCreatePageHybridEnabled: core.BoolPtr(true),
				Parameters:                    []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel},
			}

			supportTimeIntervalModel := &partnercentersellv1.SupportTimeInterval{
				Value: core.Float64Ptr(float64(72)),
				Type:  core.StringPtr("testString"),
			}

			supportEscalationModel := &partnercentersellv1.SupportEscalation{
				Contact:            core.StringPtr("testString"),
				EscalationWaitTime: supportTimeIntervalModel,
				ResponseWaitTime:   supportTimeIntervalModel,
			}

			supportDetailsItemAvailabilityTimeModel := &partnercentersellv1.SupportDetailsItemAvailabilityTime{
				Day:       core.Float64Ptr(float64(72)),
				StartTime: core.StringPtr("10:00"),
				EndTime:   core.StringPtr("10:00"),
			}

			supportDetailsItemAvailabilityModel := &partnercentersellv1.SupportDetailsItemAvailability{
				Times:           []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel},
				Timezone:        core.StringPtr("testString"),
				AlwaysAvailable: core.BoolPtr(true),
			}

			supportDetailsItemModel := &partnercentersellv1.SupportDetailsItem{
				Type:             core.StringPtr("support_site"),
				Contact:          core.StringPtr("testString"),
				ResponseWaitTime: supportTimeIntervalModel,
				Availability:     supportDetailsItemAvailabilityModel,
			}

			globalCatalogProductMetadataOtherPcSupportModel := &partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport{
				URL:               core.StringPtr("https://http.cat/"),
				StatusURL:         core.StringPtr("https://http.cat/status"),
				Locations:         []string{"hu"},
				Languages:         []string{"hu"},
				Process:           core.StringPtr("testString"),
				ProcessI18n:       map[string]string{"anyKey": "anyValue"},
				SupportType:       core.StringPtr("community"),
				SupportEscalation: supportEscalationModel,
				SupportDetails:    []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel},
			}

			globalCatalogProductMetadataOtherPcModel := &partnercentersellv1.GlobalCatalogProductMetadataOtherPC{
				Support: globalCatalogProductMetadataOtherPcSupportModel,
			}

			globalCatalogProductMetadataOtherCompositeChildModel := &partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{
				Kind: core.StringPtr("service"),
				Name: core.StringPtr("testString"),
			}

			globalCatalogProductMetadataOtherCompositeModel := &partnercentersellv1.GlobalCatalogProductMetadataOtherComposite{
				CompositeKind: core.StringPtr("service"),
				CompositeTag:  core.StringPtr("testString"),
				Children:      []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel},
			}

			globalCatalogProductMetadataOtherModel := &partnercentersellv1.GlobalCatalogProductMetadataOther{
				PC:        globalCatalogProductMetadataOtherPcModel,
				Composite: globalCatalogProductMetadataOtherCompositeModel,
			}

			globalCatalogProductMetadataPrototypePatchModel := &partnercentersellv1.GlobalCatalogProductMetadataPrototypePatch{
				RcCompatible: core.BoolPtr(true),
				Ui:           globalCatalogProductMetadataUiModel,
				Service:      globalCatalogProductMetadataServicePrototypePatchModel,
				Other:        globalCatalogProductMetadataOtherModel,
			}

			globalCatalogProductPatchModel := &partnercentersellv1.GlobalCatalogProductPatch{
				Active:         core.BoolPtr(true),
				Disabled:       core.BoolPtr(false),
				OverviewUi:     globalCatalogOverviewUiModel,
				Tags:           []string{"tag"},
				Images:         globalCatalogProductImagesModel,
				ObjectProvider: catalogProductProviderModel,
				Metadata:       globalCatalogProductMetadataPrototypePatchModel,
			}
			globalCatalogProductPatchModelAsPatch, asPatchErr := globalCatalogProductPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogProductOptions := &partnercentersellv1.UpdateCatalogProductOptions{
				ProductID:                 core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID:          &catalogProductIdLink,
				GlobalCatalogProductPatch: globalCatalogProductPatchModelAsPatch,
				Env:                       core.StringPtr(env),
			}

			globalCatalogProduct, response, err := partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogProduct).ToNot(BeNil())

			catalogProductIdLink = *globalCatalogProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogProductIdLink value: %v\n", catalogProductIdLink)
		})
	})

	Describe(`CreateCatalogPlan - Create a pricing plan in global catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCatalogPlan(createCatalogPlanOptions *CreateCatalogPlanOptions)`, func() {
			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("IBM"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName:     core.StringPtr("My plan"),
				Description:     core.StringPtr("My plan description."),
				LongDescription: core.StringPtr("My plan long description."),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			catalogHighlightItemModel := &partnercentersellv1.CatalogHighlightItem{
				Description: core.StringPtr("testString"),
				Title:       core.StringPtr("testString"),
			}

			catalogProductMediaItemModel := &partnercentersellv1.CatalogProductMediaItem{
				Caption:   core.StringPtr("testString"),
				Thumbnail: core.StringPtr("testString"),
				Type:      core.StringPtr("image"),
				URL:       core.StringPtr("testString"),
			}

			globalCatalogMetadataUiNavigationItemModel := &partnercentersellv1.GlobalCatalogMetadataUINavigationItem{
				ID:    core.StringPtr("testString"),
				URL:   core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
			}

			globalCatalogMetadataUiStringsContentModel := &partnercentersellv1.GlobalCatalogMetadataUIStringsContent{
				Bullets:         []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel},
				Media:           []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel},
				NavigationItems: []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel},
			}

			globalCatalogMetadataUiStringsModel := &partnercentersellv1.GlobalCatalogMetadataUIStrings{
				En: globalCatalogMetadataUiStringsContentModel,
			}

			globalCatalogMetadataUiUrlsModel := &partnercentersellv1.GlobalCatalogMetadataUIUrls{
				DocURL:              core.StringPtr("https://http.cat/do"),
				ApidocsURL:          core.StringPtr("https://http.cat/do"),
				TermsURL:            core.StringPtr("https://http.cat/do"),
				InstructionsURL:     core.StringPtr("https://http.cat/do"),
				CatalogDetailsURL:   core.StringPtr("https://http.cat/do"),
				CustomCreatePageURL: core.StringPtr("https://http.cat/do"),
				Dashboard:           core.StringPtr("https://http.cat/do"),
			}

			globalCatalogPlanMetadataUiModel := &partnercentersellv1.GlobalCatalogPlanMetadataUI{
				Strings:         globalCatalogMetadataUiStringsModel,
				Urls:            globalCatalogMetadataUiUrlsModel,
				Hidden:          core.BoolPtr(true),
				SideBySideIndex: core.Float64Ptr(float64(72)),
			}

			globalCatalogPlanMetadataServicePrototypePatchModel := &partnercentersellv1.GlobalCatalogPlanMetadataServicePrototypePatch{
				RcProvisionable:     core.BoolPtr(false),
				IamCompatible:       core.BoolPtr(true),
				Bindable:            core.BoolPtr(true),
				PlanUpdateable:      core.BoolPtr(true),
				ServiceKeySupported: core.BoolPtr(true),
			}

			globalCatalogMetadataPricingModel := &partnercentersellv1.GlobalCatalogMetadataPricing{
				Type:        core.StringPtr("paid"),
				Origin:      core.StringPtr("pricing_catalog"),
				SalesAvenue: []string{"seller"},
			}

			globalCatalogPlanMetadataPlanModel := &partnercentersellv1.GlobalCatalogPlanMetadataPlan{
				AllowInternalUsers: core.BoolPtr(true),
				ProvisionType:      core.StringPtr("ibm_cloud"),
				Reservable:         core.BoolPtr(true),
			}

			globalCatalogPlanMetadataOtherResourceControllerModel := &partnercentersellv1.GlobalCatalogPlanMetadataOtherResourceController{
				SubscriptionProviderID: core.StringPtr("crn:v1:staging:public:resource-controller::a/280d69caa3744c7b8e09878d4009c07a::resource-broker:17061cd2-911a-4c37-b8aa-991b99493d32"),
			}

			globalCatalogPlanMetadataOtherModel := &partnercentersellv1.GlobalCatalogPlanMetadataOther{
				ResourceController: globalCatalogPlanMetadataOtherResourceControllerModel,
			}

			globalCatalogPlanMetadataPrototypePatchModel := &partnercentersellv1.GlobalCatalogPlanMetadataPrototypePatch{
				RcCompatible: core.BoolPtr(true),
				Ui:           globalCatalogPlanMetadataUiModel,
				Service:      globalCatalogPlanMetadataServicePrototypePatchModel,
				Pricing:      globalCatalogMetadataPricingModel,
				Plan:         globalCatalogPlanMetadataPlanModel,
				Other:        globalCatalogPlanMetadataOtherModel,
			}

			var randomInteger = strconv.Itoa(rand.Intn(1000))
			objectId := fmt.Sprintf("random-id-%s", randomInteger)

			createCatalogPlanOptions := &partnercentersellv1.CreateCatalogPlanOptions{
				ProductID:        core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID: &catalogProductIdLink,
				Name:             core.StringPtr("free-plan2"),
				Active:           core.BoolPtr(true),
				Disabled:         core.BoolPtr(false),
				Kind:             core.StringPtr("plan"),
				Tags:             []string{"ibm_created"},
				ObjectProvider:   catalogProductProviderModel,
				ObjectID:         core.StringPtr(objectId),
				OverviewUi:       globalCatalogOverviewUiModel,
				PricingTags:      []string{"free"},
				Metadata:         globalCatalogPlanMetadataPrototypePatchModel,
				Env:              core.StringPtr(env),
			}

			globalCatalogPlan, response, err := partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogPlan).ToNot(BeNil())

			catalogPlanIdLink = *globalCatalogPlan.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogPlanIdLink value: %v\n", catalogPlanIdLink)
		})
	})

	Describe(`UpdateCatalogPlan - Update a global catalog plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateCatalogPlan(updateCatalogPlanOptions *UpdateCatalogPlanOptions)`, func() {
			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName:     core.StringPtr("test plan up"),
				Description:     core.StringPtr("test plan desc up"),
				LongDescription: core.StringPtr("test plan desc long up"),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
			}

			catalogHighlightItemModel := &partnercentersellv1.CatalogHighlightItem{
				Description: core.StringPtr("testString"),
				Title:       core.StringPtr("testString"),
			}

			catalogProductMediaItemModel := &partnercentersellv1.CatalogProductMediaItem{
				Caption:   core.StringPtr("testString"),
				Thumbnail: core.StringPtr("testString"),
				Type:      core.StringPtr("image"),
				URL:       core.StringPtr("https://http.cat/images/200.jpg"),
			}

			globalCatalogMetadataUiNavigationItemModel := &partnercentersellv1.GlobalCatalogMetadataUINavigationItem{
				ID:    core.StringPtr("testString"),
				URL:   core.StringPtr("testString"),
				Label: core.StringPtr("testString"),
			}

			globalCatalogMetadataUiStringsContentModel := &partnercentersellv1.GlobalCatalogMetadataUIStringsContent{
				Bullets:         []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel},
				Media:           []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel},
				NavigationItems: []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel},
			}

			globalCatalogMetadataUiStringsModel := &partnercentersellv1.GlobalCatalogMetadataUIStrings{
				En: globalCatalogMetadataUiStringsContentModel,
			}

			globalCatalogMetadataUiUrlsModel := &partnercentersellv1.GlobalCatalogMetadataUIUrls{
				DocURL:              core.StringPtr("https://http.cat/elua"),
				ApidocsURL:          core.StringPtr("https://http.cat/elua"),
				TermsURL:            core.StringPtr("https://http.cat/elua"),
				InstructionsURL:     core.StringPtr("https://http.cat/elua"),
				CatalogDetailsURL:   core.StringPtr("https://http.cat/elua"),
				CustomCreatePageURL: core.StringPtr("https://http.cat/elua"),
				Dashboard:           core.StringPtr("https://http.cat/elua"),
			}

			globalCatalogPlanMetadataUiModel := &partnercentersellv1.GlobalCatalogPlanMetadataUI{
				Strings:         globalCatalogMetadataUiStringsModel,
				Urls:            globalCatalogMetadataUiUrlsModel,
				Hidden:          core.BoolPtr(true),
				SideBySideIndex: core.Float64Ptr(float64(72)),
			}

			globalCatalogPlanMetadataServicePrototypePatchModel := &partnercentersellv1.GlobalCatalogPlanMetadataServicePrototypePatch{
				RcProvisionable:     core.BoolPtr(true),
				IamCompatible:       core.BoolPtr(true),
				Bindable:            core.BoolPtr(true),
				PlanUpdateable:      core.BoolPtr(true),
				ServiceKeySupported: core.BoolPtr(true),
			}

			globalCatalogMetadataPricingModel := &partnercentersellv1.GlobalCatalogMetadataPricing{
				Type:        core.StringPtr("free"),
				Origin:      core.StringPtr("pricing_catalog"),
				SalesAvenue: []string{"seller"},
			}

			globalCatalogPlanMetadataPlanModel := &partnercentersellv1.GlobalCatalogPlanMetadataPlan{
				AllowInternalUsers: core.BoolPtr(true),
				ProvisionType:      core.StringPtr("ibm_cloud"),
				Reservable:         core.BoolPtr(true),
			}

			globalCatalogPlanMetadataOtherResourceControllerModel := &partnercentersellv1.GlobalCatalogPlanMetadataOtherResourceController{
				SubscriptionProviderID: core.StringPtr("crn:v1:staging:public:resource-controller::a/280d69caa3744c7b8e09878d4009c07a::resource-broker:17061cd2-911a-4c37-b8aa-991b99493d32"),
			}

			globalCatalogPlanMetadataOtherModel := &partnercentersellv1.GlobalCatalogPlanMetadataOther{
				ResourceController: globalCatalogPlanMetadataOtherResourceControllerModel,
			}

			globalCatalogPlanMetadataPrototypePatchModel := &partnercentersellv1.GlobalCatalogPlanMetadataPrototypePatch{
				RcCompatible: core.BoolPtr(true),
				Ui:           globalCatalogPlanMetadataUiModel,
				Service:      globalCatalogPlanMetadataServicePrototypePatchModel,
				Pricing:      globalCatalogMetadataPricingModel,
				Plan:         globalCatalogPlanMetadataPlanModel,
				Other:        globalCatalogPlanMetadataOtherModel,
			}

			globalCatalogPlanPatchModel := &partnercentersellv1.GlobalCatalogPlanPatch{
				Active:         core.BoolPtr(true),
				Disabled:       core.BoolPtr(false),
				OverviewUi:     globalCatalogOverviewUiModel,
				Tags:           []string{"testString"},
				PricingTags:    []string{"free"},
				ObjectProvider: catalogProductProviderModel,
				Metadata:       globalCatalogPlanMetadataPrototypePatchModel,
			}
			globalCatalogPlanPatchModelAsPatch, asPatchErr := globalCatalogPlanPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogPlanOptions := &partnercentersellv1.UpdateCatalogPlanOptions{
				ProductID:              core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID:       &catalogProductIdLink,
				CatalogPlanID:          &catalogPlanIdLink,
				GlobalCatalogPlanPatch: globalCatalogPlanPatchModelAsPatch,
				Env:                    core.StringPtr(env),
			}

			globalCatalogPlan, response, err := partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogPlan).ToNot(BeNil())

			catalogPlanIdLink = *globalCatalogPlan.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogPlanIdLink value: %v\n", catalogPlanIdLink)
		})
	})

	Describe(`CreateCatalogDeployment - Create a global catalog deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCatalogDeployment(createCatalogDeploymentOptions *CreateCatalogDeploymentOptions)`, func() {
			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("IBM"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName:     core.StringPtr("test plan"),
				Description:     core.StringPtr("test plan desc"),
				LongDescription: core.StringPtr("test plan desc long"),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			globalCatalogMetadataServiceCustomParametersI18nFieldsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields{
				Displayname: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			globalCatalogMetadataServiceCustomParametersI18nModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n{
				En:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				De:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Es:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Fr:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				It:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Ja:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Ko:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				PtBr: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				ZhTw: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				ZhCn: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
			}

			globalCatalogMetadataServiceCustomParametersOptionsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{
				Displayname: core.StringPtr("testString"),
				Value:       core.StringPtr("testString"),
				I18n:        globalCatalogMetadataServiceCustomParametersI18nModel,
			}

			globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan{
				ShowFor:        []string{"fake-id"},
				OptionsRefresh: core.BoolPtr(true),
			}

			globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{
				Name:           core.StringPtr("fake-parameter"),
				ShowFor:        []string{"1"},
				OptionsRefresh: core.BoolPtr(true),
			}

			globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation{
				ShowFor: []string{"eu-gb"},
			}

			globalCatalogMetadataServiceCustomParametersAssociationsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations{
				Plan:       globalCatalogMetadataServiceCustomParametersAssociationsPlanModel,
				Parameters: []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel},
				Location:   globalCatalogMetadataServiceCustomParametersAssociationsLocationModel,
			}

			globalCatalogMetadataServiceCustomParametersModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{
				Displayname:    core.StringPtr("Display_name"),
				Name:           core.StringPtr("Sample_name"),
				Type:           core.StringPtr("text"),
				Options:        []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel},
				Value:          []string{"sample_value"},
				Layout:         core.StringPtr("sample_layout"),
				Associations:   globalCatalogMetadataServiceCustomParametersAssociationsModel,
				ValidationURL:  core.StringPtr("https://http.cat/valid"),
				OptionsURL:     core.StringPtr("https://http.cat/option"),
				Invalidmessage: core.StringPtr("Invalid message"),
				Description:    core.StringPtr("Sample description"),
				Required:       core.BoolPtr(true),
				Pattern:        core.StringPtr("."),
				Placeholder:    core.StringPtr("Placeholder"),
				Readonly:       core.BoolPtr(false),
				Hidden:         core.BoolPtr(false),
				I18n:           globalCatalogMetadataServiceCustomParametersI18nModel,
			}

			globalCatalogDeploymentMetadataServicePrototypePatchModel := &partnercentersellv1.GlobalCatalogDeploymentMetadataServicePrototypePatch{
				RcProvisionable:     core.BoolPtr(true),
				IamCompatible:       core.BoolPtr(true),
				ServiceKeySupported: core.BoolPtr(true),
				Parameters:          []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel},
			}

			globalCatalogMetadataDeploymentBrokerModel := &partnercentersellv1.GlobalCatalogMetadataDeploymentBroker{
				Name: core.StringPtr("brokerunique1234"),
				Guid: core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3c2"),
			}

			globalCatalogMetadataDeploymentModel := &partnercentersellv1.GlobalCatalogMetadataDeployment{
				Broker:      globalCatalogMetadataDeploymentBrokerModel,
				Location:    core.StringPtr("eu-gb"),
				LocationURL: core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb"),
				TargetCrn:   core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb"),
			}

			globalCatalogDeploymentMetadataPrototypePatchModel := &partnercentersellv1.GlobalCatalogDeploymentMetadataPrototypePatch{
				RcCompatible: core.BoolPtr(true),
				Service:      globalCatalogDeploymentMetadataServicePrototypePatchModel,
				Deployment:   globalCatalogMetadataDeploymentModel,
			}

			var randomInteger = strconv.Itoa(rand.Intn(1000))
			objectId := fmt.Sprintf("random-id-%s", randomInteger)

			createCatalogDeploymentOptions := &partnercentersellv1.CreateCatalogDeploymentOptions{
				ProductID:        core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID: &catalogProductIdLink,
				CatalogPlanID:    &catalogPlanIdLink,
				Name:             core.StringPtr("deployment-eu-de"),
				Active:           core.BoolPtr(true),
				Disabled:         core.BoolPtr(false),
				Kind:             core.StringPtr("deployment"),
				Tags:             []string{"eu-gb"},
				ObjectProvider:   catalogProductProviderModel,
				ObjectID:         core.StringPtr(objectId),
				OverviewUi:       globalCatalogOverviewUiModel,
				Metadata:         globalCatalogDeploymentMetadataPrototypePatchModel,
				Env:              core.StringPtr(env),
			}

			globalCatalogDeployment, response, err := partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogDeployment).ToNot(BeNil())

			catalogDeploymentIdLink = *globalCatalogDeployment.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogDeploymentIdLink value: %v\n", catalogDeploymentIdLink)
		})
	})

	Describe(`UpdateCatalogDeployment - Update a global catalog deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateCatalogDeployment(updateCatalogDeploymentOptions *UpdateCatalogDeploymentOptions)`, func() {
			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName:     core.StringPtr("test plan up"),
				Description:     core.StringPtr("test plan desc up"),
				LongDescription: core.StringPtr("test plan desc long up"),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("IBM"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			globalCatalogMetadataServiceCustomParametersI18nFieldsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields{
				Displayname: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			globalCatalogMetadataServiceCustomParametersI18nModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n{
				En:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				De:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Es:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Fr:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				It:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Ja:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				Ko:   globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				PtBr: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				ZhTw: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
				ZhCn: globalCatalogMetadataServiceCustomParametersI18nFieldsModel,
			}

			globalCatalogMetadataServiceCustomParametersOptionsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{
				Displayname: core.StringPtr("testString"),
				Value:       core.StringPtr("testString"),
				I18n:        globalCatalogMetadataServiceCustomParametersI18nModel,
			}

			globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan{
				ShowFor:        []string{"fake-id-2"},
				OptionsRefresh: core.BoolPtr(true),
			}

			globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{
				Name:           core.StringPtr("parameter-2"),
				ShowFor:        []string{"2"},
				OptionsRefresh: core.BoolPtr(true),
			}

			globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation{
				ShowFor: []string{"us-east"},
			}

			globalCatalogMetadataServiceCustomParametersAssociationsModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations{
				Plan:       globalCatalogMetadataServiceCustomParametersAssociationsPlanModel,
				Parameters: []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel},
				Location:   globalCatalogMetadataServiceCustomParametersAssociationsLocationModel,
			}

			globalCatalogMetadataServiceCustomParametersModel := &partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{
				Displayname:    core.StringPtr("Display_name"),
				Name:           core.StringPtr("Sample_name"),
				Type:           core.StringPtr("text"),
				Options:        []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel},
				Value:          []string{"sample_value"},
				Layout:         core.StringPtr("sample_layout"),
				Associations:   globalCatalogMetadataServiceCustomParametersAssociationsModel,
				ValidationURL:  core.StringPtr("https://http.cat/valid"),
				OptionsURL:     core.StringPtr("https://http.cat/option"),
				Invalidmessage: core.StringPtr("Invalid message"),
				Description:    core.StringPtr("Sample description"),
				Required:       core.BoolPtr(true),
				Pattern:        core.StringPtr("."),
				Placeholder:    core.StringPtr("Placeholder"),
				Readonly:       core.BoolPtr(false),
				Hidden:         core.BoolPtr(false),
				I18n:           globalCatalogMetadataServiceCustomParametersI18nModel,
			}

			globalCatalogDeploymentMetadataServicePrototypePatchModel := &partnercentersellv1.GlobalCatalogDeploymentMetadataServicePrototypePatch{
				RcProvisionable:     core.BoolPtr(true),
				IamCompatible:       core.BoolPtr(true),
				ServiceKeySupported: core.BoolPtr(true),
				Parameters:          []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel},
			}

			globalCatalogMetadataDeploymentBrokerModel := &partnercentersellv1.GlobalCatalogMetadataDeploymentBroker{
				Name: core.StringPtr("another-broker"),
				Guid: core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Aresource-controller%3A%3Aa%2F4a5c3c51b97a446fbb1d0e1ef089823b%3A%3Aresource-broker%3A5fb34e97-74f6-47a6-900c-07eed308d3cf"),
			}

			globalCatalogMetadataDeploymentModel := &partnercentersellv1.GlobalCatalogMetadataDeployment{
				Broker:      globalCatalogMetadataDeploymentBrokerModel,
				Location:    core.StringPtr("eu-gb"),
				LocationURL: core.StringPtr("https://globalcatalog.test.cloud.ibm.com/api/v1/eu-gb"),
				TargetCrn:   core.StringPtr("crn:v1:staging:public::eu-gb:::environment:staging-eu-gb"),
			}

			globalCatalogDeploymentMetadataPrototypePatchModel := &partnercentersellv1.GlobalCatalogDeploymentMetadataPrototypePatch{
				RcCompatible: core.BoolPtr(true),
				Service:      globalCatalogDeploymentMetadataServicePrototypePatchModel,
				Deployment:   globalCatalogMetadataDeploymentModel,
			}

			globalCatalogDeploymentPatchModel := &partnercentersellv1.GlobalCatalogDeploymentPatch{
				Active:         core.BoolPtr(true),
				Disabled:       core.BoolPtr(false),
				OverviewUi:     globalCatalogOverviewUiModel,
				Tags:           []string{"testString"},
				ObjectProvider: catalogProductProviderModel,
				Metadata:       globalCatalogDeploymentMetadataPrototypePatchModel,
			}
			globalCatalogDeploymentPatchModelAsPatch, asPatchErr := globalCatalogDeploymentPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogDeploymentOptions := &partnercentersellv1.UpdateCatalogDeploymentOptions{
				ProductID:                    core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID:             &catalogProductIdLink,
				CatalogPlanID:                &catalogPlanIdLink,
				CatalogDeploymentID:          &catalogDeploymentIdLink,
				GlobalCatalogDeploymentPatch: globalCatalogDeploymentPatchModelAsPatch,
				Env:                          core.StringPtr(env),
			}

			globalCatalogDeployment, response, err := partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogDeployment).ToNot(BeNil())

			catalogDeploymentIdLink = *globalCatalogDeployment.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogDeploymentIdLink value: %v\n", catalogDeploymentIdLink)
		})
	})

	Describe(`CreateIamRegistration - Create IAM registration for your service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		var randomInteger = strconv.Itoa(rand.Intn(1000))
		roleDisplayName := fmt.Sprintf("random-%s-2", randomInteger)

		It(`CreateIamRegistration(createIamRegistrationOptions *CreateIamRegistrationOptions)`, func() {
			iamServiceRegistrationDescriptionObjectModel := &partnercentersellv1.IamServiceRegistrationDescriptionObject{
				Default: core.StringPtr("View dashboard"),
				En:      core.StringPtr("View dashboard"),
				De:      core.StringPtr("View dashboard"),
				Es:      core.StringPtr("View dashboard"),
				Fr:      core.StringPtr("View dashboard"),
				It:      core.StringPtr("View dashboard"),
				Ja:      core.StringPtr("View dashboard"),
				Ko:      core.StringPtr("View dashboard"),
				PtBr:    core.StringPtr("View dashboard"),
				ZhTw:    core.StringPtr("View dashboard"),
				ZhCn:    core.StringPtr("View dashboard"),
			}

			iamServiceRegistrationDisplayNameObjectModel := &partnercentersellv1.IamServiceRegistrationDisplayNameObject{
				Default: core.StringPtr(roleDisplayName),
				En:      core.StringPtr(roleDisplayName),
				De:      core.StringPtr(roleDisplayName),
				Es:      core.StringPtr(roleDisplayName),
				Fr:      core.StringPtr(roleDisplayName),
				It:      core.StringPtr(roleDisplayName),
				Ja:      core.StringPtr(roleDisplayName),
				Ko:      core.StringPtr(roleDisplayName),
				PtBr:    core.StringPtr(roleDisplayName),
				ZhTw:    core.StringPtr(roleDisplayName),
				ZhCn:    core.StringPtr(roleDisplayName),
			}

			iamServiceRegistrationActionOptionsModel := &partnercentersellv1.IamServiceRegistrationActionOptions{
				Hidden: core.BoolPtr(true),
			}

			iamServiceRegistrationActionModel := &partnercentersellv1.IamServiceRegistrationAction{
				ID:          core.StringPtr(fmt.Sprintf("%s.dashboard.view", iamServiceRegistrationId)),
				Roles:       []string{fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", iamServiceRegistrationId, roleDisplayName)},
				Description: iamServiceRegistrationDescriptionObjectModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
				Options:     iamServiceRegistrationActionOptionsModel,
			}

			iamServiceRegistrationSupportedAnonymousAccessAttributesModel := &partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes{
				AccountID:            core.StringPtr("testString"),
				ServiceName:          core.StringPtr(iamServiceRegistrationId),
				AdditionalProperties: map[string]string{"testString": "dsgdsfgsd576456"},
			}

			iamServiceRegistrationSupportedAnonymousAccessModel := &partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{
				Attributes: iamServiceRegistrationSupportedAnonymousAccessAttributesModel,
				Roles:      []string{fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", iamServiceRegistrationId, roleDisplayName)},
			}

			supportedAttributesOptionsResourceHierarchyKeyModel := &partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey{
				Key:   core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			supportedAttributesOptionsResourceHierarchyValueModel := &partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue{
				Key: core.StringPtr("testString"),
			}

			supportedAttributesOptionsResourceHierarchyModel := &partnercentersellv1.SupportedAttributesOptionsResourceHierarchy{
				Key:   supportedAttributesOptionsResourceHierarchyKeyModel,
				Value: supportedAttributesOptionsResourceHierarchyValueModel,
			}

			supportedAttributesOptionsModel := &partnercentersellv1.SupportedAttributesOptions{
				Operators:                         []string{"stringEquals"},
				Hidden:                            core.BoolPtr(false),
				SupportedPatterns:                 []string{"attribute-based-condition:resource:literal-and-wildcard"},
				PolicyTypes:                       []string{"access"},
				IsEmptyValueSupported:             core.BoolPtr(true),
				IsStringExistsFalseValueSupported: core.BoolPtr(true),
				ResourceHierarchy:                 supportedAttributesOptionsResourceHierarchyModel,
			}

			supportedAttributeUiInputValueModel := &partnercentersellv1.SupportedAttributeUiInputValue{
				Value:       core.StringPtr("testString"),
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
			}

			supportedAttributeUiInputGstModel := &partnercentersellv1.SupportedAttributeUiInputGst{
				Query:             core.StringPtr("ghost query"),
				ValuePropertyName: core.StringPtr("instance"),
				InputOptionLabel:  core.StringPtr("{name} - {instance_id}"),
			}

			supportedAttributeUiInputDetailsModel := &partnercentersellv1.SupportedAttributeUiInputDetails{
				Type:   core.StringPtr("gst"),
				Values: []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel},
				Gst:    supportedAttributeUiInputGstModel,
			}

			supportedAttributeUiModel := &partnercentersellv1.SupportedAttributeUi{
				InputType:    core.StringPtr("selector"),
				InputDetails: supportedAttributeUiInputDetailsModel,
			}

			iamServiceRegistrationSupportedAttributeModel := &partnercentersellv1.IamServiceRegistrationSupportedAttribute{
				Key:         core.StringPtr("testString"),
				Options:     supportedAttributesOptionsModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
				Description: iamServiceRegistrationDescriptionObjectModel,
				Ui:          supportedAttributeUiModel,
			}

			supportAuthorizationSubjectAttributeModel := &partnercentersellv1.SupportAuthorizationSubjectAttribute{
				ServiceName:  core.StringPtr("testString"),
				ResourceType: core.StringPtr("testString"),
			}

			iamServiceRegistrationSupportedAuthorizationSubjectModel := &partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{
				Attributes: supportAuthorizationSubjectAttributeModel,
				Roles:      []string{fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", iamServiceRegistrationId, roleDisplayName)},
			}

			supportedRoleOptionsModel := &partnercentersellv1.SupportedRoleOptions{
				AccessPolicy: core.BoolPtr(true),
				PolicyType:   []string{"access"},
			}

			iamServiceRegistrationSupportedRoleModel := &partnercentersellv1.IamServiceRegistrationSupportedRole{
				ID:          core.StringPtr(fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", iamServiceRegistrationId, roleDisplayName)),
				Description: iamServiceRegistrationDescriptionObjectModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
				Options:     supportedRoleOptionsModel,
			}

			environmentAttributeOptionsModel := &partnercentersellv1.EnvironmentAttributeOptions{
				Hidden: core.BoolPtr(false),
			}

			environmentAttributeModel := &partnercentersellv1.EnvironmentAttribute{
				Key:     core.StringPtr("networkType"),
				Values:  []string{"public"},
				Options: environmentAttributeOptionsModel,
			}

			iamServiceRegistrationSupportedNetworkModel := &partnercentersellv1.IamServiceRegistrationSupportedNetwork{
				EnvironmentAttributes: []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel},
			}

			createIamRegistrationOptions := &partnercentersellv1.CreateIamRegistrationOptions{
				ProductID:                      core.StringPtr(productIdWithApprovedProgrammaticName),
				Name:                           core.StringPtr(iamServiceRegistrationId),
				Enabled:                        core.BoolPtr(true),
				ServiceType:                    core.StringPtr("service"),
				Actions:                        []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel},
				ParentIds:                      []string{"3bee3c3c-998c-432a-adff-b387750ceb49"},
				DisplayName:                    iamServiceRegistrationDisplayNameObjectModel,
				SupportedAnonymousAccesses:     []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel},
				SupportedAttributes:            []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel},
				SupportedAuthorizationSubjects: []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel},
				SupportedRoles:                 []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel},
				SupportedNetwork:               iamServiceRegistrationSupportedNetworkModel,
				Env:                            core.StringPtr(env),
			}

			iamServiceRegistration, response, err := partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(iamServiceRegistration).ToNot(BeNil())

			programmaticNameLink = *iamServiceRegistration.Name
			fmt.Fprintf(GinkgoWriter, "Saved programmaticNameLink value: %v\n", programmaticNameLink)
		})
	})

	Describe(`UpdateIamRegistration - Update IAM registration for your service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateIamRegistration(updateIamRegistrationOptions *UpdateIamRegistrationOptions)`, func() {
			var randomInteger = strconv.Itoa(rand.Intn(10000))
			roleDisplayName := fmt.Sprintf("random-%s-2", randomInteger)

			iamServiceRegistrationDescriptionObjectModel := &partnercentersellv1.IamServiceRegistrationDescriptionObject{
				Default: core.StringPtr("View dashboard"),
				En:      core.StringPtr("View dashboard"),
				De:      core.StringPtr("View dashboard"),
				Es:      core.StringPtr("View dashboard"),
				Fr:      core.StringPtr("View dashboard"),
				It:      core.StringPtr("View dashboard"),
				Ja:      core.StringPtr("View dashboard"),
				Ko:      core.StringPtr("View dashboard"),
				PtBr:    core.StringPtr("View dashboard"),
				ZhTw:    core.StringPtr("View dashboard"),
				ZhCn:    core.StringPtr("View dashboard"),
			}

			iamServiceRegistrationDisplayNameObjectModel := &partnercentersellv1.IamServiceRegistrationDisplayNameObject{
				Default: core.StringPtr(roleDisplayName),
				En:      core.StringPtr(roleDisplayName),
				De:      core.StringPtr(roleDisplayName),
				Es:      core.StringPtr(roleDisplayName),
				Fr:      core.StringPtr(roleDisplayName),
				It:      core.StringPtr(roleDisplayName),
				Ja:      core.StringPtr(roleDisplayName),
				Ko:      core.StringPtr(roleDisplayName),
				PtBr:    core.StringPtr(roleDisplayName),
				ZhTw:    core.StringPtr(roleDisplayName),
				ZhCn:    core.StringPtr(roleDisplayName),
			}

			iamServiceRegistrationActionOptionsModel := &partnercentersellv1.IamServiceRegistrationActionOptions{
				Hidden: core.BoolPtr(true),
			}

			iamServiceRegistrationActionModel := &partnercentersellv1.IamServiceRegistrationAction{
				ID:          core.StringPtr(fmt.Sprintf("%s.dashboard.view", iamServiceRegistrationId)),
				Roles:       []string{fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", iamServiceRegistrationId, roleDisplayName)},
				Description: iamServiceRegistrationDescriptionObjectModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
				Options:     iamServiceRegistrationActionOptionsModel,
			}

			iamServiceRegistrationSupportedAnonymousAccessAttributesModel := &partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes{
				AccountID:            core.StringPtr("testString"),
				ServiceName:          core.StringPtr(iamServiceRegistrationId),
				AdditionalProperties: map[string]string{"testString": "fakeAccount"},
			}

			iamServiceRegistrationSupportedAnonymousAccessModel := &partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{
				Attributes: iamServiceRegistrationSupportedAnonymousAccessAttributesModel,
				Roles:      []string{fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", iamServiceRegistrationId, roleDisplayName)},
			}

			supportedAttributesOptionsResourceHierarchyKeyModel := &partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey{
				Key:   core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			supportedAttributesOptionsResourceHierarchyValueModel := &partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue{
				Key: core.StringPtr("testString"),
			}

			supportedAttributesOptionsResourceHierarchyModel := &partnercentersellv1.SupportedAttributesOptionsResourceHierarchy{
				Key:   supportedAttributesOptionsResourceHierarchyKeyModel,
				Value: supportedAttributesOptionsResourceHierarchyValueModel,
			}

			supportedAttributesOptionsModel := &partnercentersellv1.SupportedAttributesOptions{
				Operators:                         []string{"stringEquals"},
				Hidden:                            core.BoolPtr(false),
				SupportedPatterns:                 []string{"attribute-based-condition:resource:literal-and-wildcard"},
				PolicyTypes:                       []string{"access"},
				IsEmptyValueSupported:             core.BoolPtr(true),
				IsStringExistsFalseValueSupported: core.BoolPtr(true),
				ResourceHierarchy:                 supportedAttributesOptionsResourceHierarchyModel,
			}

			supportedAttributeUiInputValueModel := &partnercentersellv1.SupportedAttributeUiInputValue{
				Value:       core.StringPtr("testString"),
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
			}

			supportedAttributeUiInputGstModel := &partnercentersellv1.SupportedAttributeUiInputGst{
				Query:             core.StringPtr("ghost query"),
				ValuePropertyName: core.StringPtr("instance"),
				InputOptionLabel:  core.StringPtr("{name} - {instance_id}"),
			}

			supportedAttributeUiInputDetailsModel := &partnercentersellv1.SupportedAttributeUiInputDetails{
				Type:   core.StringPtr("gst"),
				Values: []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel},
				Gst:    supportedAttributeUiInputGstModel,
			}

			supportedAttributeUiModel := &partnercentersellv1.SupportedAttributeUi{
				InputType:    core.StringPtr("selector"),
				InputDetails: supportedAttributeUiInputDetailsModel,
			}

			iamServiceRegistrationSupportedAttributeModel := &partnercentersellv1.IamServiceRegistrationSupportedAttribute{
				Key:         core.StringPtr("testString"),
				Options:     supportedAttributesOptionsModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
				Description: iamServiceRegistrationDescriptionObjectModel,
				Ui:          supportedAttributeUiModel,
			}

			supportAuthorizationSubjectAttributeModel := &partnercentersellv1.SupportAuthorizationSubjectAttribute{
				ServiceName:  core.StringPtr("testString"),
				ResourceType: core.StringPtr("testString"),
			}

			iamServiceRegistrationSupportedAuthorizationSubjectModel := &partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{
				Attributes: supportAuthorizationSubjectAttributeModel,
				Roles:      []string{fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", iamServiceRegistrationId, roleDisplayName)},
			}

			supportedRoleOptionsModel := &partnercentersellv1.SupportedRoleOptions{
				AccessPolicy: core.BoolPtr(true),
				PolicyType:   []string{"access"},
			}

			iamServiceRegistrationSupportedRoleModel := &partnercentersellv1.IamServiceRegistrationSupportedRole{
				ID:          core.StringPtr(fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", iamServiceRegistrationId, roleDisplayName)),
				Description: iamServiceRegistrationDescriptionObjectModel,
				DisplayName: iamServiceRegistrationDisplayNameObjectModel,
				Options:     supportedRoleOptionsModel,
			}

			environmentAttributeOptionsModel := &partnercentersellv1.EnvironmentAttributeOptions{
				Hidden: core.BoolPtr(true),
			}

			environmentAttributeModel := &partnercentersellv1.EnvironmentAttribute{
				Key:     core.StringPtr("networkType"),
				Values:  []string{"public"},
				Options: environmentAttributeOptionsModel,
			}

			iamServiceRegistrationSupportedNetworkModel := &partnercentersellv1.IamServiceRegistrationSupportedNetwork{
				EnvironmentAttributes: []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel},
			}

			iamServiceRegistrationPatchModel := &partnercentersellv1.IamServiceRegistrationPatch{
				Enabled:                        core.BoolPtr(true),
				ServiceType:                    core.StringPtr("service"),
				Actions:                        []partnercentersellv1.IamServiceRegistrationAction{*iamServiceRegistrationActionModel},
				AdditionalPolicyScopes:         []string{iamServiceRegistrationId},
				DisplayName:                    iamServiceRegistrationDisplayNameObjectModel,
				ParentIds:                      []string{"3bee3c3c-998c-432a-adff-b387750ceb49"},
				SupportedAnonymousAccesses:     []partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess{*iamServiceRegistrationSupportedAnonymousAccessModel},
				SupportedAttributes:            []partnercentersellv1.IamServiceRegistrationSupportedAttribute{*iamServiceRegistrationSupportedAttributeModel},
				SupportedAuthorizationSubjects: []partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject{*iamServiceRegistrationSupportedAuthorizationSubjectModel},
				SupportedRoles:                 []partnercentersellv1.IamServiceRegistrationSupportedRole{*iamServiceRegistrationSupportedRoleModel},
				SupportedNetwork:               iamServiceRegistrationSupportedNetworkModel,
			}
			iamServiceRegistrationPatchModelAsPatch, asPatchErr := iamServiceRegistrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateIamRegistrationOptions := &partnercentersellv1.UpdateIamRegistrationOptions{
				ProductID:            core.StringPtr(productIdWithApprovedProgrammaticName),
				ProgrammaticName:     &programmaticNameLink,
				IamRegistrationPatch: iamServiceRegistrationPatchModelAsPatch,
				Env:                  core.StringPtr(env),
			}

			iamServiceRegistration, response, err := partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iamServiceRegistration).ToNot(BeNil())

			programmaticNameLink = *iamServiceRegistration.Name
			fmt.Fprintf(GinkgoWriter, "Saved programmaticNameLink value: %v\n", programmaticNameLink)
		})
	})

	Describe(`CreateResourceBroker - Create a broker`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResourceBroker(createResourceBrokerOptions *CreateResourceBrokerOptions)`, func() {
			var randomInteger = strconv.Itoa(rand.Intn(1000))
			brokerUrl := fmt.Sprintf("https://broker-url-for-service.com/%s", randomInteger)
			brokerName := fmt.Sprintf("petya_test_2_%s", randomInteger)

			createResourceBrokerOptions := &partnercentersellv1.CreateResourceBrokerOptions{
				AuthUsername:        core.StringPtr("apikey"),
				AuthPassword:        core.StringPtr("0GANZzXiTurnXTF_000-FVk500800sdkrTHAt000y00y"),
				AuthScheme:          core.StringPtr("bearer"),
				Name:                core.StringPtr(brokerName),
				BrokerURL:           core.StringPtr(brokerUrl),
				Type:                core.StringPtr("provision_through"),
				State:               core.StringPtr("active"),
				AllowContextUpdates: core.BoolPtr(true),
				CatalogType:         core.StringPtr("service"),
				Region:              core.StringPtr("global"),
				Env:                 core.StringPtr(env),
			}

			broker, response, err := partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(broker).ToNot(BeNil())

			brokerIdLink = *broker.ID
			fmt.Fprintf(GinkgoWriter, "Saved brokerIdLink value: %v\n", brokerIdLink)
		})
	})

	Describe(`GetRegistration - Retrieve a Partner Center - Sell registration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRegistration(getRegistrationOptions *GetRegistrationOptions)`, func() {
			getRegistrationOptions := &partnercentersellv1.GetRegistrationOptions{
				RegistrationID: &registrationIdLink,
			}

			registration, response, err := partnerCenterSellServiceAlt.GetRegistration(getRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registration).ToNot(BeNil())
		})
	})

	Describe(`UpdateRegistration - Update a Partner Center - Sell registration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateRegistration(updateRegistrationOptions *UpdateRegistrationOptions)`, func() {
			primaryContactModel := &partnercentersellv1.PrimaryContact{
				Name:  core.StringPtr("Petra"),
				Email: core.StringPtr("petra@ibm.com"),
			}

			registrationPatchModel := &partnercentersellv1.RegistrationPatch{
				CompanyName:    core.StringPtr("company_sdk_new"),
				PrimaryContact: primaryContactModel,
			}
			registrationPatchModelAsPatch, asPatchErr := registrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateRegistrationOptions := &partnercentersellv1.UpdateRegistrationOptions{
				RegistrationID:    &registrationIdLink,
				RegistrationPatch: registrationPatchModelAsPatch,
			}

			registration, response, err := partnerCenterSellServiceAlt.UpdateRegistration(updateRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registration).ToNot(BeNil())
		})
	})

	Describe(`GetOnboardingProduct - Get an onboarding product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOnboardingProduct(getOnboardingProductOptions *GetOnboardingProductOptions)`, func() {
			getOnboardingProductOptions := &partnercentersellv1.GetOnboardingProductOptions{
				ProductID: &productIdLink,
			}

			onboardingProduct, response, err := partnerCenterSellServiceAlt.GetOnboardingProduct(getOnboardingProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(onboardingProduct).ToNot(BeNil())

			productIdLink = *onboardingProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved productIdLink value: %v\n", productIdLink)
		})
	})

	Describe(`GetCatalogProduct - Get a global catalog product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalogProduct(getCatalogProductOptions *GetCatalogProductOptions)`, func() {
			getCatalogProductOptions := &partnercentersellv1.GetCatalogProductOptions{
				ProductID:        core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID: &catalogProductIdLink,
				Env:              core.StringPtr(env),
			}

			globalCatalogProduct, response, err := partnerCenterSellService.GetCatalogProduct(getCatalogProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogProduct).ToNot(BeNil())

			catalogProductIdLink = *globalCatalogProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogProductIdLink value: %v\n", catalogProductIdLink)
		})
	})

	Describe(`GetCatalogPlan - Get a global catalog pricing plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalogPlan(getCatalogPlanOptions *GetCatalogPlanOptions)`, func() {
			getCatalogPlanOptions := &partnercentersellv1.GetCatalogPlanOptions{
				ProductID:        core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID: &catalogProductIdLink,
				CatalogPlanID:    &catalogPlanIdLink,
				Env:              core.StringPtr(env),
			}

			globalCatalogPlan, response, err := partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogPlan).ToNot(BeNil())

			catalogPlanIdLink = *globalCatalogPlan.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogPlanIdLink value: %v\n", catalogPlanIdLink)
		})
	})

	Describe(`GetCatalogDeployment - Get a global catalog deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCatalogDeployment(getCatalogDeploymentOptions *GetCatalogDeploymentOptions)`, func() {
			getCatalogDeploymentOptions := &partnercentersellv1.GetCatalogDeploymentOptions{
				ProductID:           core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID:    &catalogProductIdLink,
				CatalogPlanID:       &catalogPlanIdLink,
				CatalogDeploymentID: &catalogDeploymentIdLink,
				Env:                 core.StringPtr(env),
			}

			globalCatalogDeployment, response, err := partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogDeployment).ToNot(BeNil())

			catalogDeploymentIdLink = *globalCatalogDeployment.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogDeploymentIdLink value: %v\n", catalogDeploymentIdLink)
		})
	})

	Describe(`GetIamRegistration - Get IAM registration for your service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetIamRegistration(getIamRegistrationOptions *GetIamRegistrationOptions)`, func() {
			getIamRegistrationOptions := &partnercentersellv1.GetIamRegistrationOptions{
				ProductID:        core.StringPtr(productIdWithApprovedProgrammaticName),
				ProgrammaticName: core.StringPtr(iamServiceRegistrationId),
				Env:              core.StringPtr(env),
			}

			iamServiceRegistration, response, err := partnerCenterSellService.GetIamRegistration(getIamRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iamServiceRegistration).ToNot(BeNil())

			programmaticNameLink = *iamServiceRegistration.Name
			fmt.Fprintf(GinkgoWriter, "Saved programmaticNameLink value: %v\n", programmaticNameLink)
		})
	})

	Describe(`UpdateResourceBroker - Update broker details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateResourceBroker(updateResourceBrokerOptions *UpdateResourceBrokerOptions)`, func() {
			var randomInteger = strconv.Itoa(rand.Intn(1000))
			brokerUrl := fmt.Sprintf("https://broker-url-for-my-service.com/%s", randomInteger)

			brokerPatchModel := &partnercentersellv1.BrokerPatch{
				AuthUsername:        core.StringPtr("apikey"),
				AuthPassword:        core.StringPtr("0GANZzXiTurnXTF_000-FVk500800sdkrTHAt000y00y"),
				AuthScheme:          core.StringPtr("bearer"),
				BrokerURL:           core.StringPtr(brokerUrl),
				State:               core.StringPtr("active"),
				AllowContextUpdates: core.BoolPtr(true),
				CatalogType:         core.StringPtr("service"),
				Region:              core.StringPtr("global"),
			}
			brokerPatchModelAsPatch, asPatchErr := brokerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateResourceBrokerOptions := &partnercentersellv1.UpdateResourceBrokerOptions{
				BrokerID:    &brokerIdLink,
				BrokerPatch: brokerPatchModelAsPatch,
				Env:         core.StringPtr(env),
			}

			broker, response, err := partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(broker).ToNot(BeNil())
		})
	})

	Describe(`GetResourceBroker - Get a broker`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceBroker(getResourceBrokerOptions *GetResourceBrokerOptions)`, func() {
			getResourceBrokerOptions := &partnercentersellv1.GetResourceBrokerOptions{
				BrokerID: &brokerIdLink,
				Env:      core.StringPtr(env),
			}

			broker, response, err := partnerCenterSellService.GetResourceBroker(getResourceBrokerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(broker).ToNot(BeNil())
		})
	})

	Describe(`ListProductBadges - List badges`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProductBadges(listProductBadgesOptions *ListProductBadgesOptions)`, func() {
			listProductBadgesOptions := &partnercentersellv1.ListProductBadgesOptions{
				Limit: core.Int64Ptr(int64(100)),
				Start: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			productBadgeCollection, response, err := partnerCenterSellService.ListProductBadges(listProductBadgesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productBadgeCollection).ToNot(BeNil())
		})
	})

	Describe(`GetProductBadge - Get badge`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProductBadge(getProductBadgeOptions *GetProductBadgeOptions)`, func() {
			getProductBadgeOptions := &partnercentersellv1.GetProductBadgeOptions{
				BadgeID: CreateMockUUID(badgeId),
			}

			productBadge, response, err := partnerCenterSellService.GetProductBadge(getProductBadgeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productBadge).ToNot(BeNil())
		})
	})
	Describe(`DeleteCatalogDeployment - Delete a global catalog deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCatalogDeployment(deleteCatalogDeploymentOptions *DeleteCatalogDeploymentOptions)`, func() {
			deleteCatalogDeploymentOptions := &partnercentersellv1.DeleteCatalogDeploymentOptions{
				ProductID:           core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID:    &catalogProductIdLink,
				CatalogPlanID:       &catalogPlanIdLink,
				CatalogDeploymentID: &catalogDeploymentIdLink,
				Env:                 core.StringPtr(env),
			}
			response, err := partnerCenterSellService.DeleteCatalogDeployment(deleteCatalogDeploymentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteCatalogPlan - Delete a global catalog pricing plan`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCatalogPlan(deleteCatalogPlanOptions *DeleteCatalogPlanOptions)`, func() {
			deleteCatalogPlanOptions := &partnercentersellv1.DeleteCatalogPlanOptions{
				ProductID:        core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID: &catalogProductIdLink,
				CatalogPlanID:    &catalogPlanIdLink,
				Env:              core.StringPtr(env),
			}

			response, err := partnerCenterSellService.DeleteCatalogPlan(deleteCatalogPlanOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteCatalogProduct - Delete a global catalog product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCatalogProduct(deleteCatalogProductOptions *DeleteCatalogProductOptions)`, func() {
			deleteCatalogProductOptions := &partnercentersellv1.DeleteCatalogProductOptions{
				ProductID:        core.StringPtr(productIdWithApprovedProgrammaticName),
				CatalogProductID: &catalogProductIdLink,
				Env:              core.StringPtr(env),
			}

			response, err := partnerCenterSellService.DeleteCatalogProduct(deleteCatalogProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteIamRegistration - Delete IAM registration for your service`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteIamRegistration(deleteIamRegistrationOptions *DeleteIamRegistrationOptions)`, func() {
			deleteIamRegistrationOptions := &partnercentersellv1.DeleteIamRegistrationOptions{
				ProductID:        core.StringPtr(productIdWithApprovedProgrammaticName),
				ProgrammaticName: core.StringPtr(iamServiceRegistrationId),
				Env:              core.StringPtr(env),
			}

			response, err := partnerCenterSellService.DeleteIamRegistration(deleteIamRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteOnboardingProduct - Delete an onboarding product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteOnboardingProduct(deleteOnboardingProductOptions *DeleteOnboardingProductOptions)`, func() {
			deleteOnboardingProductOptions := &partnercentersellv1.DeleteOnboardingProductOptions{
				ProductID: &productIdLink,
			}

			response, err := partnerCenterSellServiceAlt.DeleteOnboardingProduct(deleteOnboardingProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
	Describe(`DeleteRegistration - Delete your registration in Partner - Center Sell`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRegistration(deleteRegistrationOptions *DeleteRegistrationOptions)`, func() {
			deleteRegistrationOptions := &partnercentersellv1.DeleteRegistrationOptions{
				RegistrationID: &registrationIdLink,
			}

			response, err := partnerCenterSellServiceAlt.DeleteRegistration(deleteRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
	Describe(`DeleteResourceBroker - Remove a broker`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteResourceBroker(deleteResourceBrokerOptions *DeleteResourceBrokerOptions)`, func() {
			deleteResourceBrokerOptions := &partnercentersellv1.DeleteResourceBrokerOptions{
				BrokerID:          &brokerIdLink,
				Env:               core.StringPtr(env),
				RemoveFromAccount: core.BoolPtr(true),
			}

			response, err := partnerCenterSellService.DeleteResourceBroker(deleteResourceBrokerOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
