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

func TestAccIbmOnboardingCatalogPlanBasic(t *testing.T) {
	var conf partnercentersellv1.GlobalCatalogPlan
	productID := fmt.Sprintf("tf_product_id_%d", acctest.RandIntRange(10, 100))
	catalogProductID := fmt.Sprintf("tf_catalog_product_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	active := "true"
	disabled := "true"
	kind := "plan"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "plan"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogPlanDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogPlanConfigBasic(productID, catalogProductID, name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogPlanExists("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "catalog_product_id", catalogProductID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogPlanConfigBasic(productID, catalogProductID, nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "catalog_product_id", catalogProductID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "active", activeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "kind", kindUpdate),
				),
			},
		},
	})
}

func TestAccIbmOnboardingCatalogPlanAllArgs(t *testing.T) {
	var conf partnercentersellv1.GlobalCatalogPlan
	productID := fmt.Sprintf("tf_product_id_%d", acctest.RandIntRange(10, 100))
	catalogProductID := fmt.Sprintf("tf_catalog_product_id_%d", acctest.RandIntRange(10, 100))
	env := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	active := "true"
	disabled := "true"
	kind := "plan"
	envUpdate := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "plan"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogPlanDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogPlanConfig(productID, catalogProductID, env, name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogPlanExists("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "catalog_product_id", catalogProductID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "env", env),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogPlanConfig(productID, catalogProductID, envUpdate, nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "catalog_product_id", catalogProductID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "env", envUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "active", activeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance", "kind", kindUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_onboarding_catalog_plan.onboarding_catalog_plan",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmOnboardingCatalogPlanConfigBasic(productID string, catalogProductID string, name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_catalog_plan" "onboarding_catalog_plan_instance" {
			product_id = "%s"
			catalog_product_id = "%s"
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
	`, productID, catalogProductID, name, active, disabled, kind)
}

func testAccCheckIbmOnboardingCatalogPlanConfig(productID string, catalogProductID string, env string, name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_catalog_plan" "onboarding_catalog_plan_instance" {
			product_id = "%s"
			catalog_product_id = "%s"
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
				pricing {
					type = "free"
					origin = "global_catalog"
				}
			}
		}
	`, productID, catalogProductID, env, name, active, disabled, kind)
}

func testAccCheckIbmOnboardingCatalogPlanExists(n string, obj partnercentersellv1.GlobalCatalogPlan) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
		if err != nil {
			return err
		}

		getCatalogPlanOptions := &partnercentersellv1.GetCatalogPlanOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getCatalogPlanOptions.SetProductID(parts[0])
		getCatalogPlanOptions.SetCatalogProductID(parts[1])
		getCatalogPlanOptions.SetCatalogPlanID(parts[2])

		globalCatalogPlan, _, err := partnerCenterSellClient.GetCatalogPlan(getCatalogPlanOptions)
		if err != nil {
			return err
		}

		obj = *globalCatalogPlan
		return nil
	}
}

func testAccCheckIbmOnboardingCatalogPlanDestroy(s *terraform.State) error {
	partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_onboarding_catalog_plan" {
			continue
		}

		getCatalogPlanOptions := &partnercentersellv1.GetCatalogPlanOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getCatalogPlanOptions.SetProductID(parts[0])
		getCatalogPlanOptions.SetCatalogProductID(parts[1])
		getCatalogPlanOptions.SetCatalogPlanID(parts[2])

		// Try to find the key
		_, response, err := partnerCenterSellClient.GetCatalogPlan(getCatalogPlanOptions)

		if err == nil {
			return fmt.Errorf("onboarding_catalog_plan still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for onboarding_catalog_plan (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogOverviewUIToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogOverviewUIToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogOverviewUITranslatedContentToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogOverviewUITranslatedContentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanCatalogProductProviderToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["email"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogProductProvider)
	model.Name = core.StringPtr("testString")
	model.Email = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanCatalogProductProviderToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataToMap(t *testing.T) {
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

		globalCatalogMetadataPricingModel := make(map[string]interface{})
		globalCatalogMetadataPricingModel["type"] = "free"
		globalCatalogMetadataPricingModel["origin"] = "global_catalog"

		model := make(map[string]interface{})
		model["rc_compatible"] = true
		model["ui"] = []map[string]interface{}{globalCatalogMetadataUiModel}
		model["pricing"] = []map[string]interface{}{globalCatalogMetadataPricingModel}

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

	globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
	globalCatalogMetadataPricingModel.Type = core.StringPtr("free")
	globalCatalogMetadataPricingModel.Origin = core.StringPtr("global_catalog")

	model := new(partnercentersellv1.GlobalCatalogPlanMetadata)
	model.RcCompatible = core.BoolPtr(true)
	model.Ui = globalCatalogMetadataUiModel
	model.Pricing = globalCatalogMetadataPricingModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIStringsToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIStringsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIStringsContentToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIStringsContentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanCatalogHighlightItemToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanCatalogHighlightItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanCatalogProductMediaItemToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanCatalogProductMediaItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIUrlsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["doc_url"] = "testString"
		model["terms_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	model.DocURL = core.StringPtr("testString")
	model.TermsURL = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIUrlsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataPricingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["type"] = "free"
		model["origin"] = "global_catalog"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataPricing)
	model.Type = core.StringPtr("free")
	model.Origin = core.StringPtr("global_catalog")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataPricingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToCatalogProductProvider(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogProductProvider) {
		model := new(partnercentersellv1.CatalogProductProvider)
		model.Name = core.StringPtr("testString")
		model.Email = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["email"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToCatalogProductProvider(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogOverviewUI(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogOverviewUI(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogOverviewUITranslatedContent(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogOverviewUITranslatedContent(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadata(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogPlanMetadata) {
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

		globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
		globalCatalogMetadataPricingModel.Type = core.StringPtr("free")
		globalCatalogMetadataPricingModel.Origin = core.StringPtr("global_catalog")

		model := new(partnercentersellv1.GlobalCatalogPlanMetadata)
		model.RcCompatible = core.BoolPtr(true)
		model.Ui = globalCatalogMetadataUiModel
		model.Pricing = globalCatalogMetadataPricingModel

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

	globalCatalogMetadataPricingModel := make(map[string]interface{})
	globalCatalogMetadataPricingModel["type"] = "free"
	globalCatalogMetadataPricingModel["origin"] = "global_catalog"

	model := make(map[string]interface{})
	model["rc_compatible"] = true
	model["ui"] = []interface{}{globalCatalogMetadataUiModel}
	model["pricing"] = []interface{}{globalCatalogMetadataPricingModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadata(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUI(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUI(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIStrings(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIStrings(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIStringsContent(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIStringsContent(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToCatalogHighlightItem(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToCatalogHighlightItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToCatalogProductMediaItem(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToCatalogProductMediaItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIUrls(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIUrls) {
		model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		model.DocURL = core.StringPtr("testString")
		model.TermsURL = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["doc_url"] = "testString"
	model["terms_url"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIUrls(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataPricing(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataPricing) {
		model := new(partnercentersellv1.GlobalCatalogMetadataPricing)
		model.Type = core.StringPtr("free")
		model.Origin = core.StringPtr("global_catalog")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["type"] = "free"
	model["origin"] = "global_catalog"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataPricing(model)
	assert.Nil(t, err)
	checkResult(result)
}
