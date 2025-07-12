//go:build examples

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
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Partner Center Sell service.
//
// The following configuration properties are assumed to be defined:
// PARTNER_CENTER_SELL_URL=<service base url>
// PARTNER_CENTER_SELL_AUTH_TYPE=iam
// PARTNER_CENTER_SELL_APIKEY=<IAM apikey>
// PARTNER_CENTER_SELL_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// PRODUCT_ID_APPROVED=<product id>
// PARTNER_CENTER_SELL_BADGE_ID=<badge id>
// PARTNER_CENTER_SELL_IAM_REGISTRATION_ID=<iam registration id>

// PARTNER_CENTER_SELL_ALT_AUTH_TYPE=iam
// PARTNER_CENTER_SELL_ALT_APIKEY=<IAM apikey>
// PARTNER_CENTER_SELL_ALT_URL=<service base url>

// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`PartnerCenterSellV1 Examples Tests`, func() {

	const externalConfigFile = "../partner_center_sell_v1.env"

	var (
		partnerCenterSellService    *partnercentersellv1.PartnerCenterSellV1
		partnerCenterSellServiceAlt *partnercentersellv1.PartnerCenterSellV1
		config                      map[string]string

		// Variables to hold link values
		accountId                             string
		productIdWithApprovedProgrammaticName string
		badgeId                               string
		brokerIdLink                          string
		catalogDeploymentIdLink               string
		catalogPlanIdLink                     string
		catalogProductIdLink                  string
		productIdLink                         string
		programmaticNameLink                  string
		registrationIdLink                    string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(partnercentersellv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			accountId = config["ACCOUNT_ID"]
			Expect(accountId).ToNot(BeEmpty())

			productIdWithApprovedProgrammaticName = config["PRODUCT_ID_APPROVED"]
			Expect(productIdWithApprovedProgrammaticName).ToNot(BeEmpty())

			badgeId = config["BADGE_ID"]
			Expect(badgeId).ToNot(BeEmpty())

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			partnerCenterSellServiceOptions := &partnercentersellv1.PartnerCenterSellV1Options{}

			partnerCenterSellService, err = partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(partnerCenterSellServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(partnerCenterSellService).ToNot(BeNil())
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

	Describe(`PartnerCenterSellV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRegistration request example`, func() {
			fmt.Println("\nCreateRegistration() result:")
			// begin-create_registration

			primaryContactModel := &partnercentersellv1.PrimaryContact{
				Name:  core.StringPtr("Company Representative"),
				Email: core.StringPtr("companyrep@email.com"),
			}

			createRegistrationOptions := partnerCenterSellServiceAlt.NewCreateRegistrationOptions(
				accountId,
				"Beautiful Company",
				primaryContactModel,
			)

			registration, response, err := partnerCenterSellServiceAlt.CreateRegistration(createRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(registration, "", "  ")
			fmt.Println(string(b))

			// end-create_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registration).ToNot(BeNil())

			registrationIdLink = *registration.ID
			fmt.Fprintf(GinkgoWriter, "Saved registrationIdLink value: %v\n", registrationIdLink)
		})
		It(`CreateOnboardingProduct request example`, func() {
			fmt.Println("\nCreateOnboardingProduct() result:")
			// begin-create_onboarding_product

			primaryContactModel := &partnercentersellv1.PrimaryContact{
				Name:  core.StringPtr("name"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			createOnboardingProductOptions := partnerCenterSellServiceAlt.NewCreateOnboardingProductOptions(
				"service",
				primaryContactModel,
			)

			onboardingProduct, response, err := partnerCenterSellServiceAlt.CreateOnboardingProduct(createOnboardingProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(onboardingProduct, "", "  ")
			fmt.Println(string(b))

			// end-create_onboarding_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(onboardingProduct).ToNot(BeNil())

			productIdLink = *onboardingProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved productIdLink value: %v\n", productIdLink)
		})
		It(`UpdateOnboardingProduct request example`, func() {
			fmt.Println("\nUpdateOnboardingProduct() result:")
			// begin-update_onboarding_product

			onboardingProductPatchModel := &partnercentersellv1.OnboardingProductPatch{}
			onboardingProductPatchModelAsPatch, asPatchErr := onboardingProductPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateOnboardingProductOptions := partnerCenterSellServiceAlt.NewUpdateOnboardingProductOptions(
				productIdLink,
				onboardingProductPatchModelAsPatch,
			)

			onboardingProduct, response, err := partnerCenterSellServiceAlt.UpdateOnboardingProduct(updateOnboardingProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(onboardingProduct, "", "  ")
			fmt.Println(string(b))

			// end-update_onboarding_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(onboardingProduct).ToNot(BeNil())

			productIdLink = *onboardingProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved productIdLink value: %v\n", productIdLink)
		})
		It(`CreateCatalogProduct request example`, func() {
			fmt.Println("\nCreateCatalogProduct() result:")
			// begin-create_catalog_product

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

			globalCatalogProductMetadataServicePrototypePatchModel := &partnercentersellv1.GlobalCatalogProductMetadataServicePrototypePatch{
				RcProvisionable: core.BoolPtr(true),
				IamCompatible:   core.BoolPtr(true),
			}

			globalCatalogProductMetadataPrototypePatchModel := &partnercentersellv1.GlobalCatalogProductMetadataPrototypePatch{
				RcCompatible: core.BoolPtr(true),
				Service:      globalCatalogProductMetadataServicePrototypePatchModel,
			}

			var randomInteger = strconv.Itoa(rand.Intn(1000))
			catalogProductName := fmt.Sprintf("gc-product-example-%s", randomInteger)

			createCatalogProductOptions := partnerCenterSellService.NewCreateCatalogProductOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductName,
				true,
				false,
				"service",
				[]string{"keyword", "support_ibm"},
				catalogProductProviderModel,
			)
			createCatalogProductOptions.SetOverviewUi(globalCatalogOverviewUiModel)
			createCatalogProductOptions.SetMetadata(globalCatalogProductMetadataPrototypePatchModel)

			globalCatalogProduct, response, err := partnerCenterSellService.CreateCatalogProduct(createCatalogProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogProduct, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogProduct).ToNot(BeNil())

			catalogProductIdLink = *globalCatalogProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogProductIdLink value: %v\n", catalogProductIdLink)
		})
		It(`UpdateCatalogProduct request example`, func() {
			fmt.Println("\nUpdateCatalogProduct() result:")
			// begin-update_catalog_product

			globalCatalogOverviewUiTranslatedContentModel := &partnercentersellv1.GlobalCatalogOverviewUITranslatedContent{
				DisplayName: core.StringPtr("My updated display name."),
			}

			globalCatalogOverviewUiModel := &partnercentersellv1.GlobalCatalogOverviewUI{
				En: globalCatalogOverviewUiTranslatedContentModel,
			}

			globalCatalogProductPatchModel := &partnercentersellv1.GlobalCatalogProductPatch{
				OverviewUi: globalCatalogOverviewUiModel,
			}
			globalCatalogProductPatchModelAsPatch, asPatchErr := globalCatalogProductPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogProductOptions := partnerCenterSellService.NewUpdateCatalogProductOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
				globalCatalogProductPatchModelAsPatch,
			)

			globalCatalogProduct, response, err := partnerCenterSellService.UpdateCatalogProduct(updateCatalogProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogProduct, "", "  ")
			fmt.Println(string(b))

			// end-update_catalog_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogProduct).ToNot(BeNil())

			catalogProductIdLink = *globalCatalogProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogProductIdLink value: %v\n", catalogProductIdLink)
		})
		It(`CreateCatalogPlan request example`, func() {
			fmt.Println("\nCreateCatalogPlan() result:")
			// begin-create_catalog_plan

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

			globalCatalogPlanMetadataServicePrototypePatchModel := &partnercentersellv1.GlobalCatalogPlanMetadataServicePrototypePatch{
				RcProvisionable: core.BoolPtr(false),
				IamCompatible:   core.BoolPtr(true),
			}

			globalCatalogMetadataPricingModel := &partnercentersellv1.GlobalCatalogMetadataPricing{
				Type:   core.StringPtr("paid"),
				Origin: core.StringPtr("pricing_catalog"),
			}

			globalCatalogPlanMetadataPrototypePatchModel := &partnercentersellv1.GlobalCatalogPlanMetadataPrototypePatch{
				RcCompatible: core.BoolPtr(true),
				Service:      globalCatalogPlanMetadataServicePrototypePatchModel,
				Pricing:      globalCatalogMetadataPricingModel,
			}

			createCatalogPlanOptions := partnerCenterSellService.NewCreateCatalogPlanOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
				"free-plan2",
				true,
				false,
				"plan",
				catalogProductProviderModel,
			)
			createCatalogPlanOptions.SetOverviewUi(globalCatalogOverviewUiModel)
			createCatalogPlanOptions.SetMetadata(globalCatalogPlanMetadataPrototypePatchModel)

			globalCatalogPlan, response, err := partnerCenterSellService.CreateCatalogPlan(createCatalogPlanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogPlan, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogPlan).ToNot(BeNil())

			catalogPlanIdLink = *globalCatalogPlan.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogPlanIdLink value: %v\n", catalogPlanIdLink)
		})
		It(`UpdateCatalogPlan request example`, func() {
			fmt.Println("\nUpdateCatalogPlan() result:")
			// begin-update_catalog_plan

			globalCatalogMetadataPricingModel := &partnercentersellv1.GlobalCatalogMetadataPricing{
				Type:   core.StringPtr("free"),
				Origin: core.StringPtr("pricing_catalog"),
			}

			globalCatalogPlanMetadataPrototypePatchModel := &partnercentersellv1.GlobalCatalogPlanMetadataPrototypePatch{
				RcCompatible: core.BoolPtr(true),
				Pricing:      globalCatalogMetadataPricingModel,
			}

			globalCatalogPlanPatchModel := &partnercentersellv1.GlobalCatalogPlanPatch{
				Metadata: globalCatalogPlanMetadataPrototypePatchModel,
			}
			globalCatalogPlanPatchModelAsPatch, asPatchErr := globalCatalogPlanPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogPlanOptions := partnerCenterSellService.NewUpdateCatalogPlanOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
				catalogPlanIdLink,
				globalCatalogPlanPatchModelAsPatch,
			)

			globalCatalogPlan, response, err := partnerCenterSellService.UpdateCatalogPlan(updateCatalogPlanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogPlan, "", "  ")
			fmt.Println(string(b))

			// end-update_catalog_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogPlan).ToNot(BeNil())

			catalogPlanIdLink = *globalCatalogPlan.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogPlanIdLink value: %v\n", catalogPlanIdLink)
		})
		It(`CreateCatalogDeployment request example`, func() {
			fmt.Println("\nCreateCatalogDeployment() result:")
			// begin-create_catalog_deployment

			catalogProductProviderModel := &partnercentersellv1.CatalogProductProvider{
				Name:  core.StringPtr("IBM"),
				Email: core.StringPtr("name.name@ibm.com"),
			}

			globalCatalogDeploymentMetadataServicePrototypePatchModel := &partnercentersellv1.GlobalCatalogDeploymentMetadataServicePrototypePatch{
				RcProvisionable: core.BoolPtr(true),
				IamCompatible:   core.BoolPtr(true),
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

			createCatalogDeploymentOptions := partnerCenterSellService.NewCreateCatalogDeploymentOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
				catalogPlanIdLink,
				"deployment-eu-de",
				true,
				false,
				"deployment",
				catalogProductProviderModel,
			)
			createCatalogDeploymentOptions.SetMetadata(globalCatalogDeploymentMetadataPrototypePatchModel)

			globalCatalogDeployment, response, err := partnerCenterSellService.CreateCatalogDeployment(createCatalogDeploymentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogDeployment, "", "  ")
			fmt.Println(string(b))

			// end-create_catalog_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(globalCatalogDeployment).ToNot(BeNil())

			catalogDeploymentIdLink = *globalCatalogDeployment.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogDeploymentIdLink value: %v\n", catalogDeploymentIdLink)
		})
		It(`UpdateCatalogDeployment request example`, func() {
			fmt.Println("\nUpdateCatalogDeployment() result:")
			// begin-update_catalog_deployment

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
				Deployment: globalCatalogMetadataDeploymentModel,
			}

			globalCatalogDeploymentPatchModel := &partnercentersellv1.GlobalCatalogDeploymentPatch{
				Metadata: globalCatalogDeploymentMetadataPrototypePatchModel,
			}
			globalCatalogDeploymentPatchModelAsPatch, asPatchErr := globalCatalogDeploymentPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateCatalogDeploymentOptions := partnerCenterSellService.NewUpdateCatalogDeploymentOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
				catalogPlanIdLink,
				catalogDeploymentIdLink,
				globalCatalogDeploymentPatchModelAsPatch,
			)

			globalCatalogDeployment, response, err := partnerCenterSellService.UpdateCatalogDeployment(updateCatalogDeploymentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogDeployment, "", "  ")
			fmt.Println(string(b))

			// end-update_catalog_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogDeployment).ToNot(BeNil())

			catalogDeploymentIdLink = *globalCatalogDeployment.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogDeploymentIdLink value: %v\n", catalogDeploymentIdLink)
		})
		It(`CreateIamRegistration request example`, func() {
			fmt.Println("\nCreateIamRegistration() result:")
			// begin-create_iam_registration

			createIamRegistrationOptions := partnerCenterSellService.NewCreateIamRegistrationOptions(
				productIdWithApprovedProgrammaticName,
				"pet-store",
			)

			iamServiceRegistration, response, err := partnerCenterSellService.CreateIamRegistration(createIamRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iamServiceRegistration, "", "  ")
			fmt.Println(string(b))

			// end-create_iam_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(iamServiceRegistration).ToNot(BeNil())

			programmaticNameLink = *iamServiceRegistration.Name
			fmt.Fprintf(GinkgoWriter, "Saved programmaticNameLink value: %v\n", programmaticNameLink)
		})
		It(`UpdateIamRegistration request example`, func() {
			fmt.Println("\nUpdateIamRegistration() result:")
			// begin-update_iam_registration

			iamServiceRegistrationPatchModel := &partnercentersellv1.IamServiceRegistrationPatch{}
			iamServiceRegistrationPatchModelAsPatch, asPatchErr := iamServiceRegistrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateIamRegistrationOptions := partnerCenterSellService.NewUpdateIamRegistrationOptions(
				productIdWithApprovedProgrammaticName,
				programmaticNameLink,
				iamServiceRegistrationPatchModelAsPatch,
			)

			iamServiceRegistration, response, err := partnerCenterSellService.UpdateIamRegistration(updateIamRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iamServiceRegistration, "", "  ")
			fmt.Println(string(b))

			// end-update_iam_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iamServiceRegistration).ToNot(BeNil())

			programmaticNameLink = *iamServiceRegistration.Name
			fmt.Fprintf(GinkgoWriter, "Saved programmaticNameLink value: %v\n", programmaticNameLink)
		})
		It(`CreateResourceBroker request example`, func() {
			fmt.Println("\nCreateResourceBroker() result:")
			// begin-create_resource_broker

			var randomInteger = strconv.Itoa(rand.Intn(1000))
			brokerName := fmt.Sprintf("broker-example-%s", randomInteger)
			brokerLink := fmt.Sprintf("https://broker-url-for-my-service.com/%s", randomInteger)

			createResourceBrokerOptions := partnerCenterSellService.NewCreateResourceBrokerOptions(
				"bearer-crn",
				brokerName,
				brokerLink,
				"provision_through",
			)

			createResourceBrokerOptions.SetState("active")
			createResourceBrokerOptions.SetAllowContextUpdates(false)
			createResourceBrokerOptions.SetCatalogType("service")
			createResourceBrokerOptions.SetRegion("global")

			broker, response, err := partnerCenterSellService.CreateResourceBroker(createResourceBrokerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(broker, "", "  ")
			fmt.Println(string(b))

			// end-create_resource_broker

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(broker).ToNot(BeNil())

			brokerIdLink = *broker.ID
			fmt.Fprintf(GinkgoWriter, "Saved brokerIdLink value: %v\n", brokerIdLink)
		})
		It(`GetRegistration request example`, func() {
			fmt.Println("\nGetRegistration() result:")
			// begin-get_registration

			getRegistrationOptions := partnerCenterSellServiceAlt.NewGetRegistrationOptions(
				registrationIdLink,
			)

			registration, response, err := partnerCenterSellServiceAlt.GetRegistration(getRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(registration, "", "  ")
			fmt.Println(string(b))

			// end-get_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registration).ToNot(BeNil())
		})
		It(`UpdateRegistration request example`, func() {
			fmt.Println("\nUpdateRegistration() result:")
			// begin-update_registration

			registrationPatchModel := &partnercentersellv1.RegistrationPatch{}
			registrationPatchModelAsPatch, asPatchErr := registrationPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateRegistrationOptions := partnerCenterSellServiceAlt.NewUpdateRegistrationOptions(
				registrationIdLink,
				registrationPatchModelAsPatch,
			)

			registration, response, err := partnerCenterSellServiceAlt.UpdateRegistration(updateRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(registration, "", "  ")
			fmt.Println(string(b))

			// end-update_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(registration).ToNot(BeNil())
		})
		It(`GetOnboardingProduct request example`, func() {
			fmt.Println("\nGetOnboardingProduct() result:")
			// begin-get_onboarding_product

			getOnboardingProductOptions := partnerCenterSellServiceAlt.NewGetOnboardingProductOptions(
				productIdLink,
			)

			onboardingProduct, response, err := partnerCenterSellServiceAlt.GetOnboardingProduct(getOnboardingProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(onboardingProduct, "", "  ")
			fmt.Println(string(b))

			// end-get_onboarding_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(onboardingProduct).ToNot(BeNil())

			productIdLink = *onboardingProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved productIdLink value: %v\n", productIdLink)
		})
		It(`GetCatalogProduct request example`, func() {
			fmt.Println("\nGetCatalogProduct() result:")
			// begin-get_catalog_product

			getCatalogProductOptions := partnerCenterSellService.NewGetCatalogProductOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
			)

			globalCatalogProduct, response, err := partnerCenterSellService.GetCatalogProduct(getCatalogProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogProduct, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogProduct).ToNot(BeNil())

			catalogProductIdLink = *globalCatalogProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogProductIdLink value: %v\n", catalogProductIdLink)
		})
		It(`GetCatalogPlan request example`, func() {
			fmt.Println("\nGetCatalogPlan() result:")
			// begin-get_catalog_plan

			getCatalogPlanOptions := partnerCenterSellService.NewGetCatalogPlanOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
				catalogPlanIdLink,
			)

			globalCatalogPlan, response, err := partnerCenterSellService.GetCatalogPlan(getCatalogPlanOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogPlan, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogPlan).ToNot(BeNil())

			catalogPlanIdLink = *globalCatalogPlan.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogPlanIdLink value: %v\n", catalogPlanIdLink)
		})
		It(`GetCatalogDeployment request example`, func() {
			fmt.Println("\nGetCatalogDeployment() result:")
			// begin-get_catalog_deployment

			getCatalogDeploymentOptions := partnerCenterSellService.NewGetCatalogDeploymentOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
				catalogPlanIdLink,
				catalogDeploymentIdLink,
			)

			globalCatalogDeployment, response, err := partnerCenterSellService.GetCatalogDeployment(getCatalogDeploymentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(globalCatalogDeployment, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(globalCatalogDeployment).ToNot(BeNil())

			catalogDeploymentIdLink = *globalCatalogDeployment.ID
			fmt.Fprintf(GinkgoWriter, "Saved catalogDeploymentIdLink value: %v\n", catalogDeploymentIdLink)
		})
		It(`GetIamRegistration request example`, func() {
			fmt.Println("\nGetIamRegistration() result:")
			// begin-get_iam_registration

			getIamRegistrationOptions := partnerCenterSellService.NewGetIamRegistrationOptions(
				productIdWithApprovedProgrammaticName,
				programmaticNameLink,
			)

			iamServiceRegistration, response, err := partnerCenterSellService.GetIamRegistration(getIamRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iamServiceRegistration, "", "  ")
			fmt.Println(string(b))

			// end-get_iam_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iamServiceRegistration).ToNot(BeNil())

			programmaticNameLink = *iamServiceRegistration.Name
			fmt.Fprintf(GinkgoWriter, "Saved programmaticNameLink value: %v\n", programmaticNameLink)
		})
		It(`UpdateResourceBroker request example`, func() {
			fmt.Println("\nUpdateResourceBroker() result:")
			// begin-update_resource_broker

			var randomInteger = strconv.Itoa(rand.Intn(1000))
			brokerLink := fmt.Sprintf("https://broker-url-for-my-service.com/%s", randomInteger)

			brokerPatchModel := &partnercentersellv1.BrokerPatch{
				BrokerURL: core.StringPtr(brokerLink),
			}
			brokerPatchModelAsPatch, asPatchErr := brokerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateResourceBrokerOptions := partnerCenterSellService.NewUpdateResourceBrokerOptions(
				brokerIdLink,
				brokerPatchModelAsPatch,
			)

			broker, response, err := partnerCenterSellService.UpdateResourceBroker(updateResourceBrokerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(broker, "", "  ")
			fmt.Println(string(b))

			// end-update_resource_broker

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(broker).ToNot(BeNil())
		})
		It(`GetResourceBroker request example`, func() {
			fmt.Println("\nGetResourceBroker() result:")
			// begin-get_resource_broker

			getResourceBrokerOptions := partnerCenterSellService.NewGetResourceBrokerOptions(
				brokerIdLink,
			)

			broker, response, err := partnerCenterSellService.GetResourceBroker(getResourceBrokerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(broker, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_broker

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(broker).ToNot(BeNil())
		})
		It(`ListProductBadges request example`, func() {
			fmt.Println("\nListProductBadges() result:")
			// begin-list_product_badges

			listProductBadgesOptions := partnerCenterSellService.NewListProductBadgesOptions()

			productBadgeCollection, response, err := partnerCenterSellService.ListProductBadges(listProductBadgesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(productBadgeCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_product_badges

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productBadgeCollection).ToNot(BeNil())
		})
		It(`GetProductBadge request example`, func() {
			fmt.Println("\nGetProductBadge() result:")
			// begin-get_product_badge

			getProductBadgeOptions := partnerCenterSellService.NewGetProductBadgeOptions(
				CreateMockUUID(badgeId),
			)

			productBadge, response, err := partnerCenterSellService.GetProductBadge(getProductBadgeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(productBadge, "", "  ")
			fmt.Println(string(b))

			// end-get_product_badge

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(productBadge).ToNot(BeNil())
		})
		It(`DeleteCatalogDeployment request example`, func() {
			// begin-delete_catalog_deployment

			deleteCatalogDeploymentOptions := partnerCenterSellService.NewDeleteCatalogDeploymentOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
				catalogPlanIdLink,
				catalogDeploymentIdLink,
			)

			response, err := partnerCenterSellService.DeleteCatalogDeployment(deleteCatalogDeploymentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteCatalogDeployment(): %d\n", response.StatusCode)
			}

			// end-delete_catalog_deployment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteCatalogPlan request example`, func() {
			// begin-delete_catalog_plan

			deleteCatalogPlanOptions := partnerCenterSellService.NewDeleteCatalogPlanOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
				catalogPlanIdLink,
			)

			response, err := partnerCenterSellService.DeleteCatalogPlan(deleteCatalogPlanOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteCatalogPlan(): %d\n", response.StatusCode)
			}

			// end-delete_catalog_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteCatalogProduct request example`, func() {
			// begin-delete_catalog_product

			deleteCatalogProductOptions := partnerCenterSellService.NewDeleteCatalogProductOptions(
				productIdWithApprovedProgrammaticName,
				catalogProductIdLink,
			)

			response, err := partnerCenterSellService.DeleteCatalogProduct(deleteCatalogProductOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteCatalogProduct(): %d\n", response.StatusCode)
			}

			// end-delete_catalog_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteIamRegistration request example`, func() {
			// begin-delete_iam_registration

			deleteIamRegistrationOptions := partnerCenterSellService.NewDeleteIamRegistrationOptions(
				productIdWithApprovedProgrammaticName,
				programmaticNameLink,
			)

			response, err := partnerCenterSellService.DeleteIamRegistration(deleteIamRegistrationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteIamRegistration(): %d\n", response.StatusCode)
			}

			// end-delete_iam_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteOnboardingProduct request example`, func() {
			// begin-delete_onboarding_product

			deleteOnboardingProductOptions := partnerCenterSellServiceAlt.NewDeleteOnboardingProductOptions(
				productIdLink,
			)

			response, err := partnerCenterSellServiceAlt.DeleteOnboardingProduct(deleteOnboardingProductOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteOnboardingProduct(): %d\n", response.StatusCode)
			}

			// end-delete_onboarding_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteRegistration request example`, func() {
			// begin-delete_registration

			deleteRegistrationOptions := partnerCenterSellServiceAlt.NewDeleteRegistrationOptions(
				registrationIdLink,
			)

			response, err := partnerCenterSellServiceAlt.DeleteRegistration(deleteRegistrationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteRegistration(): %d\n", response.StatusCode)
			}

			// end-delete_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteResourceBroker request example`, func() {
			// begin-delete_resource_broker

			deleteResourceBrokerOptions := partnerCenterSellService.NewDeleteResourceBrokerOptions(
				brokerIdLink,
			)

			response, err := partnerCenterSellService.DeleteResourceBroker(deleteResourceBrokerOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteResourceBroker(): %d\n", response.StatusCode)
			}

			// end-delete_resource_broker

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
