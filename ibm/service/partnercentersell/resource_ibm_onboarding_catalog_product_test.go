// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package partnercentersell_test

import (
	"fmt"
	"testing"

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

func TestAccIbmOnboardingCatalogProductBasic(t *testing.T) {
	var conf partnercentersellv1.GlobalCatalogProduct
	productID := acc.PcsOnboardingProductWithApprovedProgrammaticName
	name := "test-name-terraform-1"
	active := "true"
	disabled := "false"
	kind := "service"
	nameUpdate := "test-name-terraform-1"
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "service"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckPartnerCenterSell(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogProductDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogProductConfigBasic(productID, name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogProductExists("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogProductConfigBasic(productID, nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "active", activeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "kind", kindUpdate),
				),
			},
		},
	})
}

func TestAccIbmOnboardingCatalogProductAllArgs(t *testing.T) {
	var conf partnercentersellv1.GlobalCatalogProduct
	productID := acc.PcsOnboardingProductWithApprovedProgrammaticName
	env := "current"
	name := "test-name-terraform-2"
	active := "true"
	disabled := "false"
	kind := "service"
	envUpdate := "current"
	nameUpdate := "test-name-terraform-2"
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "service"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckPartnerCenterSell(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogProductDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogProductConfig(productID, env, name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogProductExists("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "env", env),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogProductConfig(productID, envUpdate, nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "env", envUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "active", activeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "kind", kindUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_onboarding_catalog_product.onboarding_catalog_product_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmOnboardingCatalogProductConfigBasic(productID string, name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_catalog_product" "onboarding_catalog_product_instance" {
			product_id = "%s"
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			tags = ["tag", "support_ibm"]
			object_provider {
				name = "name"
				email = "email@emai.com"
			}
			metadata {
				rc_compatible = false
			}
		}
	`, productID, name, active, disabled, kind)
}

func testAccCheckIbmOnboardingCatalogProductConfig(productID string, env string, name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_catalog_product" "onboarding_catalog_product_instance" {
			product_id = "%s"
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
			tags = ["tag"]
			images {
				image = "image"
			}
			object_provider {
				name = "name"
				email = "email@email.com"
			}
			metadata {
				rc_compatible = false
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
				other {
					pc {
						support {
							url = "url"
							status_url = "status_url"
							locations = [ "locations" ]
							languages = [ "languages" ]
							process = "process"
							support_type = "community"
							support_escalation {
								contact = "contact"
								escalation_wait_time {
									value = 1.0
									type = "type"
								}
								response_wait_time {
									value = 1.0
									type = "type"
								}
							}
							support_details {
								type = "support_site"
								contact = "contact"
								response_wait_time {
									value = 1.0
									type = "type"
								}
								availability {
									times {
										day = 1.0
										start_time = "start_time"
										end_time = "end_time"
									}
									timezone = "timezone"
									always_available = true
								}
							}
						}
					}
				}
			}
		}
	`, productID, env, name, active, disabled, kind)
}

func testAccCheckIbmOnboardingCatalogProductExists(n string, obj partnercentersellv1.GlobalCatalogProduct) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
		if err != nil {
			return err
		}

		getCatalogProductOptions := &partnercentersellv1.GetCatalogProductOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getCatalogProductOptions.SetProductID(parts[0])
		getCatalogProductOptions.SetCatalogProductID(parts[1])

		globalCatalogProduct, _, err := partnerCenterSellClient.GetCatalogProduct(getCatalogProductOptions)
		if err != nil {
			return err
		}

		obj = *globalCatalogProduct
		return nil
	}
}

func testAccCheckIbmOnboardingCatalogProductDestroy(s *terraform.State) error {
	partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_onboarding_catalog_product" {
			continue
		}

		getCatalogProductOptions := &partnercentersellv1.GetCatalogProductOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getCatalogProductOptions.SetProductID(parts[0])
		getCatalogProductOptions.SetCatalogProductID(parts[1])

		// Try to find the key
		_, response, err := partnerCenterSellClient.GetCatalogProduct(getCatalogProductOptions)

		if err == nil {
			return fmt.Errorf("onboarding_catalog_product still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for onboarding_catalog_product (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogOverviewUIToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogOverviewUIToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogOverviewUITranslatedContentToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogOverviewUITranslatedContentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductImagesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["image"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogProductImages)
	model.Image = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductImagesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductCatalogProductProviderToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["email"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogProductProvider)
	model.Name = core.StringPtr("testString")
	model.Email = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductCatalogProductProviderToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataToMap(t *testing.T) {
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

		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportEscalationModel := make(map[string]interface{})
		supportEscalationModel["contact"] = "testString"
		supportEscalationModel["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportEscalationModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		supportDetailsItemModel := make(map[string]interface{})
		supportDetailsItemModel["type"] = "support_site"
		supportDetailsItemModel["contact"] = "testString"
		supportDetailsItemModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportDetailsItemModel["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["locations"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["languages"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
		globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []map[string]interface{}{supportEscalationModel}
		globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []map[string]interface{}{supportDetailsItemModel}

		globalCatalogProductMetadataOtherPcModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcModel["support"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcSupportModel}

		globalCatalogProductMetadataOtherModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherModel["pc"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcModel}

		model := make(map[string]interface{})
		model["rc_compatible"] = true
		model["ui"] = []map[string]interface{}{globalCatalogMetadataUiModel}
		model["service"] = []map[string]interface{}{globalCatalogMetadataServiceModel}
		model["other"] = []map[string]interface{}{globalCatalogProductMetadataOtherModel}

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

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportEscalationModel := new(partnercentersellv1.SupportEscalation)
	supportEscalationModel.Contact = core.StringPtr("testString")
	supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
	supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
	supportDetailsItemModel.Type = core.StringPtr("support_site")
	supportDetailsItemModel.Contact = core.StringPtr("testString")
	supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
	supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

	globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
	globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
	globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
	globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

	globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
	globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

	globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
	globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel

	model := new(partnercentersellv1.GlobalCatalogProductMetadata)
	model.RcCompatible = core.BoolPtr(true)
	model.Ui = globalCatalogMetadataUiModel
	model.Service = globalCatalogMetadataServiceModel
	model.Other = globalCatalogProductMetadataOtherModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIStringsToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIStringsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIStringsContentToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIStringsContentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductCatalogHighlightItemToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductCatalogHighlightItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductCatalogProductMediaItemToMap(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductCatalogProductMediaItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIUrlsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["doc_url"] = "testString"
		model["terms_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	model.DocURL = core.StringPtr("testString")
	model.TermsURL = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIUrlsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["rc_provisionable"] = true
		model["iam_compatible"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataService)
	model.RcProvisionable = core.BoolPtr(true)
	model.IamCompatible = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportEscalationModel := make(map[string]interface{})
		supportEscalationModel["contact"] = "testString"
		supportEscalationModel["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportEscalationModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		supportDetailsItemModel := make(map[string]interface{})
		supportDetailsItemModel["type"] = "support_site"
		supportDetailsItemModel["contact"] = "testString"
		supportDetailsItemModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportDetailsItemModel["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["locations"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["languages"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
		globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []map[string]interface{}{supportEscalationModel}
		globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []map[string]interface{}{supportDetailsItemModel}

		globalCatalogProductMetadataOtherPcModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcModel["support"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcSupportModel}

		model := make(map[string]interface{})
		model["pc"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportEscalationModel := new(partnercentersellv1.SupportEscalation)
	supportEscalationModel.Contact = core.StringPtr("testString")
	supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
	supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
	supportDetailsItemModel.Type = core.StringPtr("support_site")
	supportDetailsItemModel.Contact = core.StringPtr("testString")
	supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
	supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

	globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
	globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
	globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
	globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

	globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
	globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

	model := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
	model.PC = globalCatalogProductMetadataOtherPcModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherPCToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportEscalationModel := make(map[string]interface{})
		supportEscalationModel["contact"] = "testString"
		supportEscalationModel["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportEscalationModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		supportDetailsItemModel := make(map[string]interface{})
		supportDetailsItemModel["type"] = "support_site"
		supportDetailsItemModel["contact"] = "testString"
		supportDetailsItemModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportDetailsItemModel["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["locations"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["languages"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
		globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []map[string]interface{}{supportEscalationModel}
		globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []map[string]interface{}{supportDetailsItemModel}

		model := make(map[string]interface{})
		model["support"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcSupportModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportEscalationModel := new(partnercentersellv1.SupportEscalation)
	supportEscalationModel.Contact = core.StringPtr("testString")
	supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
	supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
	supportDetailsItemModel.Type = core.StringPtr("support_site")
	supportDetailsItemModel.Contact = core.StringPtr("testString")
	supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
	supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

	globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
	globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
	globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
	globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

	model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
	model.Support = globalCatalogProductMetadataOtherPcSupportModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherPCToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherPCSupportToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportEscalationModel := make(map[string]interface{})
		supportEscalationModel["contact"] = "testString"
		supportEscalationModel["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportEscalationModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		supportDetailsItemModel := make(map[string]interface{})
		supportDetailsItemModel["type"] = "support_site"
		supportDetailsItemModel["contact"] = "testString"
		supportDetailsItemModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportDetailsItemModel["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		model := make(map[string]interface{})
		model["url"] = "testString"
		model["status_url"] = "testString"
		model["locations"] = []string{"testString"}
		model["languages"] = []string{"testString"}
		model["process"] = "testString"
		model["process_i18n"] = map[string]interface{}{"key1": "testString"}
		model["support_type"] = "community"
		model["support_escalation"] = []map[string]interface{}{supportEscalationModel}
		model["support_details"] = []map[string]interface{}{supportDetailsItemModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportEscalationModel := new(partnercentersellv1.SupportEscalation)
	supportEscalationModel.Contact = core.StringPtr("testString")
	supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
	supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
	supportDetailsItemModel.Type = core.StringPtr("support_site")
	supportDetailsItemModel.Contact = core.StringPtr("testString")
	supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
	supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

	model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
	model.URL = core.StringPtr("testString")
	model.StatusURL = core.StringPtr("testString")
	model.Locations = []string{"testString"}
	model.Languages = []string{"testString"}
	model.Process = core.StringPtr("testString")
	model.ProcessI18n = map[string]string{"key1": "testString"}
	model.SupportType = core.StringPtr("community")
	model.SupportEscalation = supportEscalationModel
	model.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherPCSupportToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportEscalationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		model := make(map[string]interface{})
		model["contact"] = "testString"
		model["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		model["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	model := new(partnercentersellv1.SupportEscalation)
	model.Contact = core.StringPtr("testString")
	model.EscalationWaitTime = supportTimeIntervalModel
	model.ResponseWaitTime = supportTimeIntervalModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportEscalationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportTimeIntervalToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["value"] = float64(72.5)
		model["type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportTimeInterval)
	model.Value = core.Float64Ptr(float64(72.5))
	model.Type = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportTimeIntervalToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportDetailsItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		model := make(map[string]interface{})
		model["type"] = "support_site"
		model["contact"] = "testString"
		model["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		model["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	model := new(partnercentersellv1.SupportDetailsItem)
	model.Type = core.StringPtr("support_site")
	model.Contact = core.StringPtr("testString")
	model.ResponseWaitTime = supportTimeIntervalModel
	model.Availability = supportDetailsItemAvailabilityModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportDetailsItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportDetailsItemAvailabilityToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		model := make(map[string]interface{})
		model["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		model["timezone"] = "testString"
		model["always_available"] = true

		assert.Equal(t, result, model)
	}

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	model := new(partnercentersellv1.SupportDetailsItemAvailability)
	model.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	model.Timezone = core.StringPtr("testString")
	model.AlwaysAvailable = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportDetailsItemAvailabilityToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportDetailsItemAvailabilityTimeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day"] = float64(72.5)
		model["start_time"] = "testString"
		model["end_time"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	model.Day = core.Float64Ptr(float64(72.5))
	model.StartTime = core.StringPtr("testString")
	model.EndTime = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportDetailsItemAvailabilityTimeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToCatalogProductProvider(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogProductProvider) {
		model := new(partnercentersellv1.CatalogProductProvider)
		model.Name = core.StringPtr("testString")
		model.Email = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["email"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToCatalogProductProvider(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogOverviewUI(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogOverviewUI(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogOverviewUITranslatedContent(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogOverviewUITranslatedContent(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductImages(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductImages) {
		model := new(partnercentersellv1.GlobalCatalogProductImages)
		model.Image = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["image"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductImages(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadata(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadata) {
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

		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportEscalationModel := new(partnercentersellv1.SupportEscalation)
		supportEscalationModel.Contact = core.StringPtr("testString")
		supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
		supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
		supportDetailsItemModel.Type = core.StringPtr("support_site")
		supportDetailsItemModel.Contact = core.StringPtr("testString")
		supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
		supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

		globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
		globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
		globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
		globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

		globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
		globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

		globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
		globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel

		model := new(partnercentersellv1.GlobalCatalogProductMetadata)
		model.RcCompatible = core.BoolPtr(true)
		model.Ui = globalCatalogMetadataUiModel
		model.Service = globalCatalogMetadataServiceModel
		model.Other = globalCatalogProductMetadataOtherModel

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

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportEscalationModel := make(map[string]interface{})
	supportEscalationModel["contact"] = "testString"
	supportEscalationModel["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportEscalationModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	supportDetailsItemModel := make(map[string]interface{})
	supportDetailsItemModel["type"] = "support_site"
	supportDetailsItemModel["contact"] = "testString"
	supportDetailsItemModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportDetailsItemModel["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["locations"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["languages"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
	globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []interface{}{supportEscalationModel}
	globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []interface{}{supportDetailsItemModel}

	globalCatalogProductMetadataOtherPcModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcModel["support"] = []interface{}{globalCatalogProductMetadataOtherPcSupportModel}

	globalCatalogProductMetadataOtherModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherModel["pc"] = []interface{}{globalCatalogProductMetadataOtherPcModel}

	model := make(map[string]interface{})
	model["rc_compatible"] = true
	model["ui"] = []interface{}{globalCatalogMetadataUiModel}
	model["service"] = []interface{}{globalCatalogMetadataServiceModel}
	model["other"] = []interface{}{globalCatalogProductMetadataOtherModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadata(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUI(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUI(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIStrings(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIStrings(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIStringsContent(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIStringsContent(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToCatalogHighlightItem(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToCatalogHighlightItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToCatalogProductMediaItem(t *testing.T) {
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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToCatalogProductMediaItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIUrls(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIUrls) {
		model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		model.DocURL = core.StringPtr("testString")
		model.TermsURL = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["doc_url"] = "testString"
	model["terms_url"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIUrls(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataService(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataService) {
		model := new(partnercentersellv1.GlobalCatalogMetadataService)
		model.RcProvisionable = core.BoolPtr(true)
		model.IamCompatible = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["rc_provisionable"] = true
	model["iam_compatible"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataService(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOther(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataOther) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportEscalationModel := new(partnercentersellv1.SupportEscalation)
		supportEscalationModel.Contact = core.StringPtr("testString")
		supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
		supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
		supportDetailsItemModel.Type = core.StringPtr("support_site")
		supportDetailsItemModel.Contact = core.StringPtr("testString")
		supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
		supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

		globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
		globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
		globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
		globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

		globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
		globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

		model := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
		model.PC = globalCatalogProductMetadataOtherPcModel

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportEscalationModel := make(map[string]interface{})
	supportEscalationModel["contact"] = "testString"
	supportEscalationModel["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportEscalationModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	supportDetailsItemModel := make(map[string]interface{})
	supportDetailsItemModel["type"] = "support_site"
	supportDetailsItemModel["contact"] = "testString"
	supportDetailsItemModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportDetailsItemModel["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["locations"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["languages"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
	globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []interface{}{supportEscalationModel}
	globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []interface{}{supportDetailsItemModel}

	globalCatalogProductMetadataOtherPcModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcModel["support"] = []interface{}{globalCatalogProductMetadataOtherPcSupportModel}

	model := make(map[string]interface{})
	model["pc"] = []interface{}{globalCatalogProductMetadataOtherPcModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOther(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherPC(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataOtherPC) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportEscalationModel := new(partnercentersellv1.SupportEscalation)
		supportEscalationModel.Contact = core.StringPtr("testString")
		supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
		supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
		supportDetailsItemModel.Type = core.StringPtr("support_site")
		supportDetailsItemModel.Contact = core.StringPtr("testString")
		supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
		supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

		globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
		globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
		globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
		globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

		model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
		model.Support = globalCatalogProductMetadataOtherPcSupportModel

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportEscalationModel := make(map[string]interface{})
	supportEscalationModel["contact"] = "testString"
	supportEscalationModel["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportEscalationModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	supportDetailsItemModel := make(map[string]interface{})
	supportDetailsItemModel["type"] = "support_site"
	supportDetailsItemModel["contact"] = "testString"
	supportDetailsItemModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportDetailsItemModel["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["locations"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["languages"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
	globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []interface{}{supportEscalationModel}
	globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []interface{}{supportDetailsItemModel}

	model := make(map[string]interface{})
	model["support"] = []interface{}{globalCatalogProductMetadataOtherPcSupportModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherPC(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherPCSupport(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportEscalationModel := new(partnercentersellv1.SupportEscalation)
		supportEscalationModel.Contact = core.StringPtr("testString")
		supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
		supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
		supportDetailsItemModel.Type = core.StringPtr("support_site")
		supportDetailsItemModel.Contact = core.StringPtr("testString")
		supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
		supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

		model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
		model.URL = core.StringPtr("testString")
		model.StatusURL = core.StringPtr("testString")
		model.Locations = []string{"testString"}
		model.Languages = []string{"testString"}
		model.Process = core.StringPtr("testString")
		model.ProcessI18n = map[string]string{"key1": "testString"}
		model.SupportType = core.StringPtr("community")
		model.SupportEscalation = supportEscalationModel
		model.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportEscalationModel := make(map[string]interface{})
	supportEscalationModel["contact"] = "testString"
	supportEscalationModel["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportEscalationModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	supportDetailsItemModel := make(map[string]interface{})
	supportDetailsItemModel["type"] = "support_site"
	supportDetailsItemModel["contact"] = "testString"
	supportDetailsItemModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportDetailsItemModel["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	model := make(map[string]interface{})
	model["url"] = "testString"
	model["status_url"] = "testString"
	model["locations"] = []interface{}{"testString"}
	model["languages"] = []interface{}{"testString"}
	model["process"] = "testString"
	model["process_i18n"] = map[string]interface{}{"key1": "testString"}
	model["support_type"] = "community"
	model["support_escalation"] = []interface{}{supportEscalationModel}
	model["support_details"] = []interface{}{supportDetailsItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherPCSupport(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportEscalation(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportEscalation) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		model := new(partnercentersellv1.SupportEscalation)
		model.Contact = core.StringPtr("testString")
		model.EscalationWaitTime = supportTimeIntervalModel
		model.ResponseWaitTime = supportTimeIntervalModel

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	model := make(map[string]interface{})
	model["contact"] = "testString"
	model["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	model["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportEscalation(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportTimeInterval(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportTimeInterval) {
		model := new(partnercentersellv1.SupportTimeInterval)
		model.Value = core.Float64Ptr(float64(72.5))
		model.Type = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["value"] = float64(72.5)
	model["type"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportTimeInterval(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportDetailsItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportDetailsItem) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		model := new(partnercentersellv1.SupportDetailsItem)
		model.Type = core.StringPtr("support_site")
		model.Contact = core.StringPtr("testString")
		model.ResponseWaitTime = supportTimeIntervalModel
		model.Availability = supportDetailsItemAvailabilityModel

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	model := make(map[string]interface{})
	model["type"] = "support_site"
	model["contact"] = "testString"
	model["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	model["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportDetailsItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportDetailsItemAvailability(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportDetailsItemAvailability) {
		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		model := new(partnercentersellv1.SupportDetailsItemAvailability)
		model.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		model.Timezone = core.StringPtr("testString")
		model.AlwaysAvailable = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	model := make(map[string]interface{})
	model["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	model["timezone"] = "testString"
	model["always_available"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportDetailsItemAvailability(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportDetailsItemAvailabilityTime(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportDetailsItemAvailabilityTime) {
		model := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		model.Day = core.Float64Ptr(float64(72.5))
		model.StartTime = core.StringPtr("testString")
		model.EndTime = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["day"] = float64(72.5)
	model["start_time"] = "testString"
	model["end_time"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportDetailsItemAvailabilityTime(model)
	assert.Nil(t, err)
	checkResult(result)
}
