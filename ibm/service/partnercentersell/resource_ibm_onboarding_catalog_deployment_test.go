// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package partnercentersell_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/partnercentersell"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmOnboardingCatalogDeploymentBasic(t *testing.T) {
	var conf partnercentersellv1.GlobalCatalogDeployment
	productID := fmt.Sprintf("tf_product_id_%d", acctest.RandIntRange(10, 100))
	catalogProductID := fmt.Sprintf("tf_catalog_product_id_%d", acctest.RandIntRange(10, 100))
	catalogPlanID := fmt.Sprintf("tf_catalog_plan_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	active := "true"
	disabled := "true"
	kind := "deployment"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "deployment"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfigBasic(productID, catalogProductID, catalogPlanID, name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogDeploymentExists("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "catalog_product_id", catalogProductID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "catalog_plan_id", catalogPlanID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfigBasic(productID, catalogProductID, catalogPlanID, nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "catalog_product_id", catalogProductID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "catalog_plan_id", catalogPlanID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "active", activeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "kind", kindUpdate),
				),
			},
		},
	})
}

func TestAccIbmOnboardingCatalogDeploymentAllArgs(t *testing.T) {
	var conf partnercentersellv1.GlobalCatalogDeployment
	productID := fmt.Sprintf("tf_product_id_%d", acctest.RandIntRange(10, 100))
	catalogProductID := fmt.Sprintf("tf_catalog_product_id_%d", acctest.RandIntRange(10, 100))
	catalogPlanID := fmt.Sprintf("tf_catalog_plan_id_%d", acctest.RandIntRange(10, 100))
	env := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	active := "true"
	disabled := "true"
	kind := "deployment"
	envUpdate := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "deployment"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfig(productID, catalogProductID, catalogPlanID, env, name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogDeploymentExists("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "catalog_product_id", catalogProductID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "catalog_plan_id", catalogPlanID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "env", env),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfig(productID, catalogProductID, catalogPlanID, envUpdate, nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "catalog_product_id", catalogProductID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "catalog_plan_id", catalogPlanID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "env", envUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "active", activeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "kind", kindUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_onboarding_catalog_deployment.onboarding_catalog_deployment",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmOnboardingCatalogDeploymentConfigBasic(productID string, catalogProductID string, catalogPlanID string, name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_catalog_deployment" "onboarding_catalog_deployment_instance" {
			product_id = "%s"
			catalog_product_id = "%s"
			catalog_plan_id = "%s"
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			tags = "FIXME"
			object_provider {
				name = "name"
				email = "email"
			}
		}
	`, productID, catalogProductID, catalogPlanID, name, active, disabled, kind)
}

func testAccCheckIbmOnboardingCatalogDeploymentConfig(productID string, catalogProductID string, catalogPlanID string, env string, name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_catalog_deployment" "onboarding_catalog_deployment_instance" {
			product_id = "%s"
			catalog_product_id = "%s"
			catalog_plan_id = "%s"
			env = "%s"
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			overview_ui {
				en {
					display_name = "display_name"
					description = "description"
					long_description = "long_description"
				}
			}
			tags = "FIXME"
			object_provider {
				name = "name"
				email = "email"
			}
			metadata {
				rc_compatible = true
				ui {
					strings {
						en {
							bullets {
								description = "description"
								description_i18n = { "key" = "inner" }
								title = "title"
								title_i18n = { "key" = "inner" }
							}
							media {
								caption = "caption"
								caption_i18n = { "key" = "inner" }
								thumbnail = "thumbnail"
								type = "image"
								url = "url"
							}
						}
					}
					urls {
						doc_url = "doc_url"
						terms_url = "terms_url"
					}
					hidden = true
					side_by_side_index = 1.0
				}
				service {
					rc_provisionable = true
					iam_compatible = true
				}
			}
		}
	`, productID, catalogProductID, catalogPlanID, env, name, active, disabled, kind)
}

func testAccCheckIbmOnboardingCatalogDeploymentExists(n string, obj partnercentersellv1.GlobalCatalogDeployment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
		if err != nil {
			return err
		}

		getCatalogDeploymentOptions := &partnercentersellv1.GetCatalogDeploymentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getCatalogDeploymentOptions.SetProductID(parts[0])
		getCatalogDeploymentOptions.SetCatalogProductID(parts[1])
		getCatalogDeploymentOptions.SetCatalogPlanID(parts[2])
		getCatalogDeploymentOptions.SetCatalogDeploymentID(parts[3])

		globalCatalogDeployment, _, err := partnerCenterSellClient.GetCatalogDeployment(getCatalogDeploymentOptions)
		if err != nil {
			return err
		}

		obj = *globalCatalogDeployment
		return nil
	}
}

func testAccCheckIbmOnboardingCatalogDeploymentDestroy(s *terraform.State) error {
	partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_onboarding_catalog_deployment" {
			continue
		}

		getCatalogDeploymentOptions := &partnercentersellv1.GetCatalogDeploymentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getCatalogDeploymentOptions.SetProductID(parts[0])
		getCatalogDeploymentOptions.SetCatalogProductID(parts[1])
		getCatalogDeploymentOptions.SetCatalogPlanID(parts[2])
		getCatalogDeploymentOptions.SetCatalogDeploymentID(parts[3])

		// Try to find the key
		_, response, err := partnerCenterSellClient.GetCatalogDeployment(getCatalogDeploymentOptions)

		if err == nil {
			return fmt.Errorf("onboarding_catalog_deployment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for onboarding_catalog_deployment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogOverviewUIToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogOverviewUiTranslatedContentModel := make(map[string]interface{})
		globalCatalogOverviewUiTranslatedContentModel["display_name"] = "testString"
		globalCatalogOverviewUiTranslatedContentModel["description"] = "testString"
		globalCatalogOverviewUiTranslatedContentModel["long_description"] = "testString"

		model := make(map[string]interface{})
		model["en"] = []map[string]interface{}{globalCatalogOverviewUiTranslatedContentModel}

		assert.Equal(t, result, model)
	}

	globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
	globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
	globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
	globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogOverviewUI)
	model.En = globalCatalogOverviewUiTranslatedContentModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogOverviewUIToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogOverviewUITranslatedContentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["display_name"] = "testString"
		model["description"] = "testString"
		model["long_description"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
	model.DisplayName = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.LongDescription = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogOverviewUITranslatedContentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentCatalogProductProviderToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["email"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogProductProvider)
	model.Name = core.StringPtr("testString")
	model.Email = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentCatalogProductProviderToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogDeploymentMetadataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["description_i18n"] = map[string]interface{}{"key1": "testString"}
		catalogHighlightItemModel["title"] = "testString"
		catalogHighlightItemModel["title_i18n"] = map[string]interface{}{"key1": "testString"}

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["caption_i18n"] = map[string]interface{}{"key1": "testString"}
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsContentModel["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel["media"] = []map[string]interface{}{catalogProductMediaItemModel}

		globalCatalogMetadataUiStringsModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsModel["en"] = []map[string]interface{}{globalCatalogMetadataUiStringsContentModel}

		globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
		globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"

		globalCatalogMetadataUiModel := make(map[string]interface{})
		globalCatalogMetadataUiModel["strings"] = []map[string]interface{}{globalCatalogMetadataUiStringsModel}
		globalCatalogMetadataUiModel["urls"] = []map[string]interface{}{globalCatalogMetadataUiUrlsModel}
		globalCatalogMetadataUiModel["hidden"] = true
		globalCatalogMetadataUiModel["side_by_side_index"] = float64(72.5)

		globalCatalogMetadataServiceModel := make(map[string]interface{})
		globalCatalogMetadataServiceModel["rc_provisionable"] = true
		globalCatalogMetadataServiceModel["iam_compatible"] = true

		model := make(map[string]interface{})
		model["rc_compatible"] = true
		model["ui"] = []map[string]interface{}{globalCatalogMetadataUiModel}
		model["service"] = []map[string]interface{}{globalCatalogMetadataServiceModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
	catalogHighlightItemModel.Title = core.StringPtr("testString")
	catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}

	globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
	globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

	globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")

	globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
	globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
	globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
	globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
	globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

	globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
	globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
	globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)

	model := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
	model.RcCompatible = core.BoolPtr(true)
	model.Ui = globalCatalogMetadataUiModel
	model.Service = globalCatalogMetadataServiceModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogDeploymentMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataUIToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["description_i18n"] = map[string]interface{}{"key1": "testString"}
		catalogHighlightItemModel["title"] = "testString"
		catalogHighlightItemModel["title_i18n"] = map[string]interface{}{"key1": "testString"}

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["caption_i18n"] = map[string]interface{}{"key1": "testString"}
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsContentModel["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel["media"] = []map[string]interface{}{catalogProductMediaItemModel}

		globalCatalogMetadataUiStringsModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsModel["en"] = []map[string]interface{}{globalCatalogMetadataUiStringsContentModel}

		globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
		globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"

		model := make(map[string]interface{})
		model["strings"] = []map[string]interface{}{globalCatalogMetadataUiStringsModel}
		model["urls"] = []map[string]interface{}{globalCatalogMetadataUiUrlsModel}
		model["hidden"] = true
		model["side_by_side_index"] = float64(72.5)

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
	catalogHighlightItemModel.Title = core.StringPtr("testString")
	catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}

	globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
	globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

	globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogMetadataUI)
	model.Strings = globalCatalogMetadataUiStringsModel
	model.Urls = globalCatalogMetadataUiUrlsModel
	model.Hidden = core.BoolPtr(true)
	model.SideBySideIndex = core.Float64Ptr(float64(72.5))

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataUIToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataUIStringsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["description_i18n"] = map[string]interface{}{"key1": "testString"}
		catalogHighlightItemModel["title"] = "testString"
		catalogHighlightItemModel["title_i18n"] = map[string]interface{}{"key1": "testString"}

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["caption_i18n"] = map[string]interface{}{"key1": "testString"}
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsContentModel["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel["media"] = []map[string]interface{}{catalogProductMediaItemModel}

		model := make(map[string]interface{})
		model["en"] = []map[string]interface{}{globalCatalogMetadataUiStringsContentModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
	catalogHighlightItemModel.Title = core.StringPtr("testString")
	catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}

	model := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
	model.En = globalCatalogMetadataUiStringsContentModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataUIStringsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataUIStringsContentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["description_i18n"] = map[string]interface{}{"key1": "testString"}
		catalogHighlightItemModel["title"] = "testString"
		catalogHighlightItemModel["title_i18n"] = map[string]interface{}{"key1": "testString"}

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["caption_i18n"] = map[string]interface{}{"key1": "testString"}
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		model := make(map[string]interface{})
		model["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		model["media"] = []map[string]interface{}{catalogProductMediaItemModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
	catalogHighlightItemModel.Title = core.StringPtr("testString")
	catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	model.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	model.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataUIStringsContentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentCatalogHighlightItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["description"] = "testString"
		model["description_i18n"] = map[string]interface{}{"key1": "testString"}
		model["title"] = "testString"
		model["title_i18n"] = map[string]interface{}{"key1": "testString"}

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogHighlightItem)
	model.Description = core.StringPtr("testString")
	model.DescriptionI18n = map[string]string{"key1": "testString"}
	model.Title = core.StringPtr("testString")
	model.TitleI18n = map[string]string{"key1": "testString"}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentCatalogHighlightItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentCatalogProductMediaItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["caption"] = "testString"
		model["caption_i18n"] = map[string]interface{}{"key1": "testString"}
		model["thumbnail"] = "testString"
		model["type"] = "image"
		model["url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogProductMediaItem)
	model.Caption = core.StringPtr("testString")
	model.CaptionI18n = map[string]string{"key1": "testString"}
	model.Thumbnail = core.StringPtr("testString")
	model.Type = core.StringPtr("image")
	model.URL = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentCatalogProductMediaItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataUIUrlsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["doc_url"] = "testString"
		model["terms_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	model.DocURL = core.StringPtr("testString")
	model.TermsURL = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataUIUrlsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["rc_provisionable"] = true
		model["iam_compatible"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataService)
	model.RcProvisionable = core.BoolPtr(true)
	model.IamCompatible = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToCatalogProductProvider(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogProductProvider) {
		model := new(partnercentersellv1.CatalogProductProvider)
		model.Name = core.StringPtr("testString")
		model.Email = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["email"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToCatalogProductProvider(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogOverviewUI(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogOverviewUI) {
		globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
		globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
		globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
		globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogOverviewUI)
		model.En = globalCatalogOverviewUiTranslatedContentModel

		assert.Equal(t, result, model)
	}

	globalCatalogOverviewUiTranslatedContentModel := make(map[string]interface{})
	globalCatalogOverviewUiTranslatedContentModel["display_name"] = "testString"
	globalCatalogOverviewUiTranslatedContentModel["description"] = "testString"
	globalCatalogOverviewUiTranslatedContentModel["long_description"] = "testString"

	model := make(map[string]interface{})
	model["en"] = []interface{}{globalCatalogOverviewUiTranslatedContentModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogOverviewUI(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogOverviewUITranslatedContent(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogOverviewUITranslatedContent) {
		model := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
		model.DisplayName = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")
		model.LongDescription = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["display_name"] = "testString"
	model["description"] = "testString"
	model["long_description"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogOverviewUITranslatedContent(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogDeploymentMetadata(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogDeploymentMetadata) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
		catalogHighlightItemModel.Title = core.StringPtr("testString")
		catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}

		globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
		globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

		globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")

		globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
		globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
		globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
		globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
		globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

		globalCatalogMetadataServiceModel := new(partnercentersellv1.GlobalCatalogMetadataService)
		globalCatalogMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
		globalCatalogMetadataServiceModel.IamCompatible = core.BoolPtr(true)

		model := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
		model.RcCompatible = core.BoolPtr(true)
		model.Ui = globalCatalogMetadataUiModel
		model.Service = globalCatalogMetadataServiceModel

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["description_i18n"] = map[string]interface{}{"key1": "testString"}
	catalogHighlightItemModel["title"] = "testString"
	catalogHighlightItemModel["title_i18n"] = map[string]interface{}{"key1": "testString"}

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["caption_i18n"] = map[string]interface{}{"key1": "testString"}
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsContentModel["bullets"] = []interface{}{catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel["media"] = []interface{}{catalogProductMediaItemModel}

	globalCatalogMetadataUiStringsModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsModel["en"] = []interface{}{globalCatalogMetadataUiStringsContentModel}

	globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
	globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"

	globalCatalogMetadataUiModel := make(map[string]interface{})
	globalCatalogMetadataUiModel["strings"] = []interface{}{globalCatalogMetadataUiStringsModel}
	globalCatalogMetadataUiModel["urls"] = []interface{}{globalCatalogMetadataUiUrlsModel}
	globalCatalogMetadataUiModel["hidden"] = true
	globalCatalogMetadataUiModel["side_by_side_index"] = float64(72.5)

	globalCatalogMetadataServiceModel := make(map[string]interface{})
	globalCatalogMetadataServiceModel["rc_provisionable"] = true
	globalCatalogMetadataServiceModel["iam_compatible"] = true

	model := make(map[string]interface{})
	model["rc_compatible"] = true
	model["ui"] = []interface{}{globalCatalogMetadataUiModel}
	model["service"] = []interface{}{globalCatalogMetadataServiceModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogDeploymentMetadata(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataUI(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUI) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
		catalogHighlightItemModel.Title = core.StringPtr("testString")
		catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}

		globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
		globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

		globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogMetadataUI)
		model.Strings = globalCatalogMetadataUiStringsModel
		model.Urls = globalCatalogMetadataUiUrlsModel
		model.Hidden = core.BoolPtr(true)
		model.SideBySideIndex = core.Float64Ptr(float64(72.5))

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["description_i18n"] = map[string]interface{}{"key1": "testString"}
	catalogHighlightItemModel["title"] = "testString"
	catalogHighlightItemModel["title_i18n"] = map[string]interface{}{"key1": "testString"}

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["caption_i18n"] = map[string]interface{}{"key1": "testString"}
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsContentModel["bullets"] = []interface{}{catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel["media"] = []interface{}{catalogProductMediaItemModel}

	globalCatalogMetadataUiStringsModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsModel["en"] = []interface{}{globalCatalogMetadataUiStringsContentModel}

	globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
	globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"

	model := make(map[string]interface{})
	model["strings"] = []interface{}{globalCatalogMetadataUiStringsModel}
	model["urls"] = []interface{}{globalCatalogMetadataUiUrlsModel}
	model["hidden"] = true
	model["side_by_side_index"] = float64(72.5)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataUI(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataUIStrings(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIStrings) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
		catalogHighlightItemModel.Title = core.StringPtr("testString")
		catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}

		model := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
		model.En = globalCatalogMetadataUiStringsContentModel

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["description_i18n"] = map[string]interface{}{"key1": "testString"}
	catalogHighlightItemModel["title"] = "testString"
	catalogHighlightItemModel["title_i18n"] = map[string]interface{}{"key1": "testString"}

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["caption_i18n"] = map[string]interface{}{"key1": "testString"}
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsContentModel["bullets"] = []interface{}{catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel["media"] = []interface{}{catalogProductMediaItemModel}

	model := make(map[string]interface{})
	model["en"] = []interface{}{globalCatalogMetadataUiStringsContentModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataUIStrings(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataUIStringsContent(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIStringsContent) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.DescriptionI18n = map[string]string{"key1": "testString"}
		catalogHighlightItemModel.Title = core.StringPtr("testString")
		catalogHighlightItemModel.TitleI18n = map[string]string{"key1": "testString"}

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.CaptionI18n = map[string]string{"key1": "testString"}
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		model.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		model.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["description_i18n"] = map[string]interface{}{"key1": "testString"}
	catalogHighlightItemModel["title"] = "testString"
	catalogHighlightItemModel["title_i18n"] = map[string]interface{}{"key1": "testString"}

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["caption_i18n"] = map[string]interface{}{"key1": "testString"}
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	model := make(map[string]interface{})
	model["bullets"] = []interface{}{catalogHighlightItemModel}
	model["media"] = []interface{}{catalogProductMediaItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataUIStringsContent(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToCatalogHighlightItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogHighlightItem) {
		model := new(partnercentersellv1.CatalogHighlightItem)
		model.Description = core.StringPtr("testString")
		model.DescriptionI18n = map[string]string{"key1": "testString"}
		model.Title = core.StringPtr("testString")
		model.TitleI18n = map[string]string{"key1": "testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["description"] = "testString"
	model["description_i18n"] = map[string]interface{}{"key1": "testString"}
	model["title"] = "testString"
	model["title_i18n"] = map[string]interface{}{"key1": "testString"}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToCatalogHighlightItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToCatalogProductMediaItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogProductMediaItem) {
		model := new(partnercentersellv1.CatalogProductMediaItem)
		model.Caption = core.StringPtr("testString")
		model.CaptionI18n = map[string]string{"key1": "testString"}
		model.Thumbnail = core.StringPtr("testString")
		model.Type = core.StringPtr("image")
		model.URL = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["caption"] = "testString"
	model["caption_i18n"] = map[string]interface{}{"key1": "testString"}
	model["thumbnail"] = "testString"
	model["type"] = "image"
	model["url"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToCatalogProductMediaItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataUIUrls(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIUrls) {
		model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		model.DocURL = core.StringPtr("testString")
		model.TermsURL = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["doc_url"] = "testString"
	model["terms_url"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataUIUrls(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataService(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataService) {
		model := new(partnercentersellv1.GlobalCatalogMetadataService)
		model.RcProvisionable = core.BoolPtr(true)
		model.IamCompatible = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["rc_provisionable"] = true
	model["iam_compatible"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataService(model)
	assert.Nil(t, err)
	checkResult(result)
}
