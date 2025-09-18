// Copyright IBM Corp. 2025 All Rights Reserved.
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
	productID := acc.PcsOnboardingProductWithCatalogProduct
	catalogProductID := acc.PcsOnboardingCatalogProductId
	objectId := fmt.Sprintf("test-object-id-terraform-3-%d", acctest.RandIntRange(10, 100))
	name := "test-plan-name-terraform"
	active := "true"
	disabled := "false"
	kind := "plan"
	nameUpdate := "test-plan-name-terraform"
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "plan"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckPartnerCenterSell(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogPlanDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogPlanConfigBasic(productID, catalogProductID, name, active, disabled, kind, objectId),
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
				Config: testAccCheckIbmOnboardingCatalogPlanConfigBasic(productID, catalogProductID, nameUpdate, activeUpdate, disabledUpdate, kindUpdate, objectId),
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
	productID := acc.PcsOnboardingProductWithCatalogProduct
	catalogProductID := acc.PcsOnboardingCatalogProductId
	objectId := fmt.Sprintf("test-object-id-terraform-3-%d", acctest.RandIntRange(10, 100))
	env := "current"
	name := "test-plan-name-terraform"
	active := "true"
	disabled := "false"
	kind := "plan"
	envUpdate := "current"
	nameUpdate := "test-plan-name-terraform"
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "plan"
	overviewUiEn := "display_name"
	overviewUiEnUpdate := "display_name_2"
	rcCompatible := "false"
	rcCompatibleUpdate := "false"
	allowInternalUsers := "true"
	allowInternalUsersUpdate := "false"
	pricingType := "paid"
	pricingTypeUpdate := "free"
	bulletTitleName := "title"
	bulletTitleNameUpdate := "title-2"
	mediaCaption := "Some random minecraft Video"
	mediaCaptionUpdate := "Some random minecraft Video 2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckPartnerCenterSell(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogPlanDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogPlanConfig(productID, catalogProductID, env, name, active, disabled, kind, objectId, overviewUiEn, rcCompatible, pricingType, allowInternalUsers, bulletTitleName, mediaCaption),
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
				Config: testAccCheckIbmOnboardingCatalogPlanUpdateConfig(productID, catalogProductID, envUpdate, nameUpdate, activeUpdate, disabledUpdate, kindUpdate, objectId, overviewUiEnUpdate, rcCompatibleUpdate, pricingTypeUpdate, allowInternalUsersUpdate, bulletTitleNameUpdate, mediaCaptionUpdate),
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
				ResourceName:      "ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"env", "product_id", "catalog_product_id", "geo_tags", "object_id",
				},
			},
		},
	})
}

func testAccCheckIbmOnboardingCatalogPlanConfigBasic(productID string, catalogProductID string, name string, active string, disabled string, kind string, objectId string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_catalog_plan" "onboarding_catalog_plan_instance" {
			product_id = "%s"
			catalog_product_id = "%s"
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			object_id = "%s"
			object_provider {
				name = "name"
			}
			metadata {
                rc_compatible =	 false
                pricing {
                    type = "paid"
                    origin = "pricing_catalog"
					sales_avenue = ["seller"]
                }
				plan {
					allow_internal_users = true
					provision_type = "ibm_cloud"
					reservable = true
				}
			}
		}
	`, productID, catalogProductID, name, active, disabled, kind, objectId)
}

func testAccCheckIbmOnboardingCatalogPlanConfig(productID string, catalogProductID string, env string, name string, active string, disabled string, kind string, objectId string, overviewUiEn string, rcCompatible string, pricingType string, allowInternalUsers string, bulletTitleName string, mediaCaption string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_catalog_plan" "onboarding_catalog_plan_instance" {
			product_id = "%s"
			catalog_product_id = "%s"
			env = "%s"
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			object_id = "%s"
			overview_ui {
				en {
					display_name = "%s"
					description = "description"
					long_description = "long_description"
				}
			}
			tags = ["tag"]
			pricing_tags = ["free"]
			object_provider {
				name = "name"
			}
			metadata {
				rc_compatible = "%s"
    			other {
      				resource_controller {
        				subscription_provider_id = "crn:v1:staging:public:resource-controller::a/280d69caa3744c7b8e09878d4009c07a::resource-broker:17061cd2-911a-4c37-b8aa-991b99493d32"
      				}
    			}
				pricing {
					type = "%s"
					origin = "global_catalog"
					sales_avenue = [ "seller" ]
				}
				plan {
					allow_internal_users = "%s"
					reservable = true
					provision_type = "ibm_cloud"
				}
				service {
				    rc_provisionable = true
      				iam_compatible = false
				    bindable = true
      				plan_updateable = true
      				service_key_supported = true
				}
				ui {
					strings {
						en {
							bullets {
                        		title = "%s"
                        		description = "some1"
							}
							media {
                        		type = "youtube"
                        		url = "https://www.youtube.com/embed/HtkpMgNFYtE"
                        		caption = "%s"
                    		}
                		}
            		}
        		}
			}
		}
	`, productID, catalogProductID, env, name, active, disabled, kind, objectId, overviewUiEn, rcCompatible, pricingType, allowInternalUsers, bulletTitleName, mediaCaption)
}

func testAccCheckIbmOnboardingCatalogPlanUpdateConfig(productID string, catalogProductID string, env string, name string, active string, disabled string, kind string, objectId string, overviewUiEn string, rcCompatible string, pricingType string, allowInternalUsers string, bulletTitleName string, mediaCaption string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_catalog_plan" "onboarding_catalog_plan_instance" {
			product_id = "%s"
			catalog_product_id = "%s"
			env = "%s"
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			object_id = "%s"
			pricing_tags = ["free"]
			overview_ui {
				en {
					display_name = "%s"
					description = "description"
					long_description = "long_description"
				}
			}
			tags = ["tag"]
			object_provider {
				name = "name"
			}
			metadata {
			    other {
      				resource_controller {
        				subscription_provider_id = "crn:v1:staging:public:resource-controller::a/280d69caa3744c7b8e09878d4009c07a::resource-broker:17061cd2-911a-4c37-b8aa-991b99493d32"
      				}
    			}
				rc_compatible = "%s"
				pricing {
					type = "%s"
					origin = "global_catalog"
					sales_avenue = [ "catalog" ]
				}
				plan {
					allow_internal_users = "%s"
					reservable = true
					provision_type = "ibm_cloud"
				}
				service {
					rc_provisionable = true
      				iam_compatible = false
					bindable = true
					plan_updateable = true
					service_key_supported = true
				}
				ui {
            		strings {
                		en {
                    		bullets {
                        		title = "%s"
                        		description = "some1"
                    		}
							bullets {
                        		title = "newBullet"
                        		description = "some1"
                    		}
							media {
                        		type = "youtube"
                        		url = "https://www.youtube.com/embed/HtkpMgNFYtE"
                        		caption = "%s"
                    		}
							media {
                        		type = "youtube"
                        		url = "https://www.youtube.com/embed/HtkpMgNFYtE"
                        		caption = "newMedia"
                    		}
                		}
					}
				}
			}
		}
	`, productID, catalogProductID, env, name, active, disabled, kind, objectId, overviewUiEn, rcCompatible, pricingType, allowInternalUsers, bulletTitleName, mediaCaption)
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
		catalogHighlightItemModel["title"] = "testString"

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
		globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

		globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsContentModel["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel["media"] = []map[string]interface{}{catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel["navigation_items"] = []map[string]interface{}{globalCatalogMetadataUiNavigationItemModel}

		globalCatalogMetadataUiStringsModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsModel["en"] = []map[string]interface{}{globalCatalogMetadataUiStringsContentModel}

		globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
		globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["apidocs_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["instructions_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["catalog_details_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["custom_create_page_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["dashboard"] = "testString"

		globalCatalogPlanMetadataUiModel := make(map[string]interface{})
		globalCatalogPlanMetadataUiModel["strings"] = []map[string]interface{}{globalCatalogMetadataUiStringsModel}
		globalCatalogPlanMetadataUiModel["urls"] = []map[string]interface{}{globalCatalogMetadataUiUrlsModel}
		globalCatalogPlanMetadataUiModel["hidden"] = true
		globalCatalogPlanMetadataUiModel["side_by_side_index"] = float64(72.5)

		globalCatalogPlanMetadataServiceModel := make(map[string]interface{})
		globalCatalogPlanMetadataServiceModel["rc_provisionable"] = true
		globalCatalogPlanMetadataServiceModel["iam_compatible"] = true
		globalCatalogPlanMetadataServiceModel["bindable"] = true
		globalCatalogPlanMetadataServiceModel["plan_updateable"] = true
		globalCatalogPlanMetadataServiceModel["service_key_supported"] = true

		globalCatalogMetadataPricingModel := make(map[string]interface{})
		globalCatalogMetadataPricingModel["type"] = "free"
		globalCatalogMetadataPricingModel["origin"] = "global_catalog"
		globalCatalogMetadataPricingModel["sales_avenue"] = []string{"seller"}

		globalCatalogPlanMetadataPlanModel := make(map[string]interface{})
		globalCatalogPlanMetadataPlanModel["allow_internal_users"] = true
		globalCatalogPlanMetadataPlanModel["provision_type"] = "ibm_cloud"
		globalCatalogPlanMetadataPlanModel["reservable"] = true

		globalCatalogPlanMetadataOtherResourceControllerModel := make(map[string]interface{})
		globalCatalogPlanMetadataOtherResourceControllerModel["subscription_provider_id"] = "testString"

		globalCatalogPlanMetadataOtherModel := make(map[string]interface{})
		globalCatalogPlanMetadataOtherModel["resource_controller"] = []map[string]interface{}{globalCatalogPlanMetadataOtherResourceControllerModel}

		model := make(map[string]interface{})
		model["rc_compatible"] = true
		model["ui"] = []map[string]interface{}{globalCatalogPlanMetadataUiModel}
		model["service"] = []map[string]interface{}{globalCatalogPlanMetadataServiceModel}
		model["pricing"] = []map[string]interface{}{globalCatalogMetadataPricingModel}
		model["plan"] = []map[string]interface{}{globalCatalogPlanMetadataPlanModel}
		model["other"] = []map[string]interface{}{globalCatalogPlanMetadataOtherModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.Title = core.StringPtr("testString")

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

	globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

	globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
	globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

	globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

	globalCatalogPlanMetadataUiModel := new(partnercentersellv1.GlobalCatalogPlanMetadataUI)
	globalCatalogPlanMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
	globalCatalogPlanMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
	globalCatalogPlanMetadataUiModel.Hidden = core.BoolPtr(true)
	globalCatalogPlanMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

	globalCatalogPlanMetadataServiceModel := new(partnercentersellv1.GlobalCatalogPlanMetadataService)
	globalCatalogPlanMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
	globalCatalogPlanMetadataServiceModel.IamCompatible = core.BoolPtr(true)
	globalCatalogPlanMetadataServiceModel.Bindable = core.BoolPtr(true)
	globalCatalogPlanMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
	globalCatalogPlanMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)

	globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
	globalCatalogMetadataPricingModel.Type = core.StringPtr("free")
	globalCatalogMetadataPricingModel.Origin = core.StringPtr("global_catalog")
	globalCatalogMetadataPricingModel.SalesAvenue = []string{"seller"}

	globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
	globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
	globalCatalogPlanMetadataPlanModel.ProvisionType = core.StringPtr("ibm_cloud")
	globalCatalogPlanMetadataPlanModel.Reservable = core.BoolPtr(true)

	globalCatalogPlanMetadataOtherResourceControllerModel := new(partnercentersellv1.GlobalCatalogPlanMetadataOtherResourceController)
	globalCatalogPlanMetadataOtherResourceControllerModel.SubscriptionProviderID = core.StringPtr("testString")

	globalCatalogPlanMetadataOtherModel := new(partnercentersellv1.GlobalCatalogPlanMetadataOther)
	globalCatalogPlanMetadataOtherModel.ResourceController = globalCatalogPlanMetadataOtherResourceControllerModel

	model := new(partnercentersellv1.GlobalCatalogPlanMetadata)
	model.RcCompatible = core.BoolPtr(true)
	model.Ui = globalCatalogPlanMetadataUiModel
	model.Service = globalCatalogPlanMetadataServiceModel
	model.Pricing = globalCatalogMetadataPricingModel
	model.Plan = globalCatalogPlanMetadataPlanModel
	model.Other = globalCatalogPlanMetadataOtherModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataUIToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["title"] = "testString"

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
		globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

		globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsContentModel["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel["media"] = []map[string]interface{}{catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel["navigation_items"] = []map[string]interface{}{globalCatalogMetadataUiNavigationItemModel}

		globalCatalogMetadataUiStringsModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsModel["en"] = []map[string]interface{}{globalCatalogMetadataUiStringsContentModel}

		globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
		globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["apidocs_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["instructions_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["catalog_details_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["custom_create_page_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["dashboard"] = "testString"

		model := make(map[string]interface{})
		model["strings"] = []map[string]interface{}{globalCatalogMetadataUiStringsModel}
		model["urls"] = []map[string]interface{}{globalCatalogMetadataUiUrlsModel}
		model["hidden"] = true
		model["side_by_side_index"] = float64(72.5)

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.Title = core.StringPtr("testString")

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

	globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

	globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
	globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

	globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogPlanMetadataUI)
	model.Strings = globalCatalogMetadataUiStringsModel
	model.Urls = globalCatalogMetadataUiUrlsModel
	model.Hidden = core.BoolPtr(true)
	model.SideBySideIndex = core.Float64Ptr(float64(72.5))

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataUIToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIStringsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["title"] = "testString"

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
		globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

		globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsContentModel["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel["media"] = []map[string]interface{}{catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel["navigation_items"] = []map[string]interface{}{globalCatalogMetadataUiNavigationItemModel}

		model := make(map[string]interface{})
		model["en"] = []map[string]interface{}{globalCatalogMetadataUiStringsContentModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.Title = core.StringPtr("testString")

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

	globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

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
		catalogHighlightItemModel["title"] = "testString"

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
		globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

		model := make(map[string]interface{})
		model["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		model["media"] = []map[string]interface{}{catalogProductMediaItemModel}
		model["navigation_items"] = []map[string]interface{}{globalCatalogMetadataUiNavigationItemModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.Title = core.StringPtr("testString")

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	model.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	model.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
	model.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIStringsContentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanCatalogHighlightItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["description"] = "testString"
		model["title"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogHighlightItem)
	model.Description = core.StringPtr("testString")
	model.Title = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanCatalogHighlightItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanCatalogProductMediaItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["caption"] = "testString"
		model["thumbnail"] = "testString"
		model["type"] = "image"
		model["url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogProductMediaItem)
	model.Caption = core.StringPtr("testString")
	model.Thumbnail = core.StringPtr("testString")
	model.Type = core.StringPtr("image")
	model.URL = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanCatalogProductMediaItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUINavigationItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["url"] = "testString"
		model["label"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	model.ID = core.StringPtr("testString")
	model.URL = core.StringPtr("testString")
	model.Label = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUINavigationItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIUrlsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["doc_url"] = "testString"
		model["apidocs_url"] = "testString"
		model["terms_url"] = "testString"
		model["instructions_url"] = "testString"
		model["catalog_details_url"] = "testString"
		model["custom_create_page_url"] = "testString"
		model["dashboard"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	model.DocURL = core.StringPtr("testString")
	model.ApidocsURL = core.StringPtr("testString")
	model.TermsURL = core.StringPtr("testString")
	model.InstructionsURL = core.StringPtr("testString")
	model.CatalogDetailsURL = core.StringPtr("testString")
	model.CustomCreatePageURL = core.StringPtr("testString")
	model.Dashboard = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataUIUrlsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataServiceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["rc_provisionable"] = true
		model["iam_compatible"] = true
		model["bindable"] = true
		model["plan_updateable"] = true
		model["service_key_supported"] = true
		model["unique_api_key"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogPlanMetadataService)
	model.RcProvisionable = core.BoolPtr(true)
	model.IamCompatible = core.BoolPtr(true)
	model.Bindable = core.BoolPtr(true)
	model.PlanUpdateable = core.BoolPtr(true)
	model.ServiceKeySupported = core.BoolPtr(true)
	model.UniqueApiKey = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataServiceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataPricingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["type"] = "free"
		model["origin"] = "global_catalog"
		model["sales_avenue"] = []string{"seller"}

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataPricing)
	model.Type = core.StringPtr("free")
	model.Origin = core.StringPtr("global_catalog")
	model.SalesAvenue = []string{"seller"}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogMetadataPricingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataPlanToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["allow_internal_users"] = true
		model["bindable"] = true
		model["provision_type"] = "ibm_cloud"
		model["reservable"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
	model.AllowInternalUsers = core.BoolPtr(true)
	model.Bindable = core.BoolPtr(true)
	model.ProvisionType = core.StringPtr("ibm_cloud")
	model.Reservable = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataPlanToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataOtherToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogPlanMetadataOtherResourceControllerModel := make(map[string]interface{})
		globalCatalogPlanMetadataOtherResourceControllerModel["subscription_provider_id"] = "testString"

		model := make(map[string]interface{})
		model["resource_controller"] = []map[string]interface{}{globalCatalogPlanMetadataOtherResourceControllerModel}

		assert.Equal(t, result, model)
	}

	globalCatalogPlanMetadataOtherResourceControllerModel := new(partnercentersellv1.GlobalCatalogPlanMetadataOtherResourceController)
	globalCatalogPlanMetadataOtherResourceControllerModel.SubscriptionProviderID = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogPlanMetadataOther)
	model.ResourceController = globalCatalogPlanMetadataOtherResourceControllerModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataOtherToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataOtherResourceControllerToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["subscription_provider_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogPlanMetadataOtherResourceController)
	model.SubscriptionProviderID = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataOtherResourceControllerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataOtherTargetPlansItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogPlanMetadataOtherTargetPlansItem)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanGlobalCatalogPlanMetadataOtherTargetPlansItemToMap(model)
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

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataPrototypePatch(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogPlanMetadataPrototypePatch) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.Title = core.StringPtr("testString")

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

		globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

		globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
		globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

		globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

		globalCatalogPlanMetadataUiModel := new(partnercentersellv1.GlobalCatalogPlanMetadataUI)
		globalCatalogPlanMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
		globalCatalogPlanMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
		globalCatalogPlanMetadataUiModel.Hidden = core.BoolPtr(true)
		globalCatalogPlanMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

		globalCatalogPlanMetadataServicePrototypePatchModel := new(partnercentersellv1.GlobalCatalogPlanMetadataServicePrototypePatch)
		globalCatalogPlanMetadataServicePrototypePatchModel.RcProvisionable = core.BoolPtr(true)
		globalCatalogPlanMetadataServicePrototypePatchModel.IamCompatible = core.BoolPtr(true)
		globalCatalogPlanMetadataServicePrototypePatchModel.Bindable = core.BoolPtr(true)
		globalCatalogPlanMetadataServicePrototypePatchModel.PlanUpdateable = core.BoolPtr(true)
		globalCatalogPlanMetadataServicePrototypePatchModel.ServiceKeySupported = core.BoolPtr(true)

		globalCatalogMetadataPricingModel := new(partnercentersellv1.GlobalCatalogMetadataPricing)
		globalCatalogMetadataPricingModel.Type = core.StringPtr("free")
		globalCatalogMetadataPricingModel.Origin = core.StringPtr("global_catalog")
		globalCatalogMetadataPricingModel.SalesAvenue = []string{"seller"}

		globalCatalogPlanMetadataPlanModel := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
		globalCatalogPlanMetadataPlanModel.AllowInternalUsers = core.BoolPtr(true)
		globalCatalogPlanMetadataPlanModel.ProvisionType = core.StringPtr("ibm_cloud")
		globalCatalogPlanMetadataPlanModel.Reservable = core.BoolPtr(true)

		globalCatalogPlanMetadataOtherResourceControllerModel := new(partnercentersellv1.GlobalCatalogPlanMetadataOtherResourceController)
		globalCatalogPlanMetadataOtherResourceControllerModel.SubscriptionProviderID = core.StringPtr("testString")

		globalCatalogPlanMetadataOtherModel := new(partnercentersellv1.GlobalCatalogPlanMetadataOther)
		globalCatalogPlanMetadataOtherModel.ResourceController = globalCatalogPlanMetadataOtherResourceControllerModel

		model := new(partnercentersellv1.GlobalCatalogPlanMetadataPrototypePatch)
		model.RcCompatible = core.BoolPtr(true)
		model.Ui = globalCatalogPlanMetadataUiModel
		model.Service = globalCatalogPlanMetadataServicePrototypePatchModel
		model.Pricing = globalCatalogMetadataPricingModel
		model.Plan = globalCatalogPlanMetadataPlanModel
		model.Other = globalCatalogPlanMetadataOtherModel

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["title"] = "testString"

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
	globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

	globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsContentModel["bullets"] = []interface{}{catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel["media"] = []interface{}{catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel["navigation_items"] = []interface{}{globalCatalogMetadataUiNavigationItemModel}

	globalCatalogMetadataUiStringsModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsModel["en"] = []interface{}{globalCatalogMetadataUiStringsContentModel}

	globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
	globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["apidocs_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["instructions_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["catalog_details_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["custom_create_page_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["dashboard"] = "testString"

	globalCatalogPlanMetadataUiModel := make(map[string]interface{})
	globalCatalogPlanMetadataUiModel["strings"] = []interface{}{globalCatalogMetadataUiStringsModel}
	globalCatalogPlanMetadataUiModel["urls"] = []interface{}{globalCatalogMetadataUiUrlsModel}
	globalCatalogPlanMetadataUiModel["hidden"] = true
	globalCatalogPlanMetadataUiModel["side_by_side_index"] = float64(72.5)

	globalCatalogPlanMetadataServicePrototypePatchModel := make(map[string]interface{})
	globalCatalogPlanMetadataServicePrototypePatchModel["rc_provisionable"] = true
	globalCatalogPlanMetadataServicePrototypePatchModel["iam_compatible"] = true
	globalCatalogPlanMetadataServicePrototypePatchModel["bindable"] = true
	globalCatalogPlanMetadataServicePrototypePatchModel["plan_updateable"] = true
	globalCatalogPlanMetadataServicePrototypePatchModel["service_key_supported"] = true

	globalCatalogMetadataPricingModel := make(map[string]interface{})
	globalCatalogMetadataPricingModel["type"] = "free"
	globalCatalogMetadataPricingModel["origin"] = "global_catalog"
	globalCatalogMetadataPricingModel["sales_avenue"] = []interface{}{"seller"}

	globalCatalogPlanMetadataPlanModel := make(map[string]interface{})
	globalCatalogPlanMetadataPlanModel["allow_internal_users"] = true
	globalCatalogPlanMetadataPlanModel["provision_type"] = "ibm_cloud"
	globalCatalogPlanMetadataPlanModel["reservable"] = true

	globalCatalogPlanMetadataOtherResourceControllerModel := make(map[string]interface{})
	globalCatalogPlanMetadataOtherResourceControllerModel["subscription_provider_id"] = "testString"

	globalCatalogPlanMetadataOtherModel := make(map[string]interface{})
	globalCatalogPlanMetadataOtherModel["resource_controller"] = []interface{}{globalCatalogPlanMetadataOtherResourceControllerModel}

	model := make(map[string]interface{})
	model["rc_compatible"] = true
	model["ui"] = []interface{}{globalCatalogPlanMetadataUiModel}
	model["service"] = []interface{}{globalCatalogPlanMetadataServicePrototypePatchModel}
	model["pricing"] = []interface{}{globalCatalogMetadataPricingModel}
	model["plan"] = []interface{}{globalCatalogPlanMetadataPlanModel}
	model["other"] = []interface{}{globalCatalogPlanMetadataOtherModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataPrototypePatch(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataUI(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogPlanMetadataUI) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.Title = core.StringPtr("testString")

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

		globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

		globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
		globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

		globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogPlanMetadataUI)
		model.Strings = globalCatalogMetadataUiStringsModel
		model.Urls = globalCatalogMetadataUiUrlsModel
		model.Hidden = core.BoolPtr(true)
		model.SideBySideIndex = core.Float64Ptr(float64(72.5))

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["title"] = "testString"

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
	globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

	globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsContentModel["bullets"] = []interface{}{catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel["media"] = []interface{}{catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel["navigation_items"] = []interface{}{globalCatalogMetadataUiNavigationItemModel}

	globalCatalogMetadataUiStringsModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsModel["en"] = []interface{}{globalCatalogMetadataUiStringsContentModel}

	globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
	globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["apidocs_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["instructions_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["catalog_details_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["custom_create_page_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["dashboard"] = "testString"

	model := make(map[string]interface{})
	model["strings"] = []interface{}{globalCatalogMetadataUiStringsModel}
	model["urls"] = []interface{}{globalCatalogMetadataUiUrlsModel}
	model["hidden"] = true
	model["side_by_side_index"] = float64(72.5)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataUI(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIStrings(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIStrings) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.Title = core.StringPtr("testString")

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

		globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

		model := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
		model.En = globalCatalogMetadataUiStringsContentModel

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["title"] = "testString"

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
	globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

	globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsContentModel["bullets"] = []interface{}{catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel["media"] = []interface{}{catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel["navigation_items"] = []interface{}{globalCatalogMetadataUiNavigationItemModel}

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
		catalogHighlightItemModel.Title = core.StringPtr("testString")

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		model.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		model.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
		model.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["title"] = "testString"

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
	globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

	model := make(map[string]interface{})
	model["bullets"] = []interface{}{catalogHighlightItemModel}
	model["media"] = []interface{}{catalogProductMediaItemModel}
	model["navigation_items"] = []interface{}{globalCatalogMetadataUiNavigationItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIStringsContent(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToCatalogHighlightItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogHighlightItem) {
		model := new(partnercentersellv1.CatalogHighlightItem)
		model.Description = core.StringPtr("testString")
		model.Title = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["description"] = "testString"
	model["title"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToCatalogHighlightItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToCatalogProductMediaItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogProductMediaItem) {
		model := new(partnercentersellv1.CatalogProductMediaItem)
		model.Caption = core.StringPtr("testString")
		model.Thumbnail = core.StringPtr("testString")
		model.Type = core.StringPtr("image")
		model.URL = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["caption"] = "testString"
	model["thumbnail"] = "testString"
	model["type"] = "image"
	model["url"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToCatalogProductMediaItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUINavigationItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUINavigationItem) {
		model := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		model.ID = core.StringPtr("testString")
		model.URL = core.StringPtr("testString")
		model.Label = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["url"] = "testString"
	model["label"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUINavigationItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIUrls(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIUrls) {
		model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		model.DocURL = core.StringPtr("testString")
		model.ApidocsURL = core.StringPtr("testString")
		model.TermsURL = core.StringPtr("testString")
		model.InstructionsURL = core.StringPtr("testString")
		model.CatalogDetailsURL = core.StringPtr("testString")
		model.CustomCreatePageURL = core.StringPtr("testString")
		model.Dashboard = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["doc_url"] = "testString"
	model["apidocs_url"] = "testString"
	model["terms_url"] = "testString"
	model["instructions_url"] = "testString"
	model["catalog_details_url"] = "testString"
	model["custom_create_page_url"] = "testString"
	model["dashboard"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataUIUrls(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataServicePrototypePatch(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogPlanMetadataServicePrototypePatch) {
		model := new(partnercentersellv1.GlobalCatalogPlanMetadataServicePrototypePatch)
		model.RcProvisionable = core.BoolPtr(true)
		model.IamCompatible = core.BoolPtr(true)
		model.Bindable = core.BoolPtr(true)
		model.PlanUpdateable = core.BoolPtr(true)
		model.ServiceKeySupported = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["rc_provisionable"] = true
	model["iam_compatible"] = true
	model["bindable"] = true
	model["plan_updateable"] = true
	model["service_key_supported"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataServicePrototypePatch(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataPricing(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataPricing) {
		model := new(partnercentersellv1.GlobalCatalogMetadataPricing)
		model.Type = core.StringPtr("free")
		model.Origin = core.StringPtr("global_catalog")
		model.SalesAvenue = []string{"seller"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["type"] = "free"
	model["origin"] = "global_catalog"
	model["sales_avenue"] = []interface{}{"seller"}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogMetadataPricing(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataPlan(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogPlanMetadataPlan) {
		model := new(partnercentersellv1.GlobalCatalogPlanMetadataPlan)
		model.AllowInternalUsers = core.BoolPtr(true)
		model.Bindable = core.BoolPtr(true)
		model.ProvisionType = core.StringPtr("ibm_cloud")
		model.Reservable = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["allow_internal_users"] = true
	model["bindable"] = true
	model["provision_type"] = "ibm_cloud"
	model["reservable"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataPlan(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataOther(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogPlanMetadataOther) {
		globalCatalogPlanMetadataOtherResourceControllerModel := new(partnercentersellv1.GlobalCatalogPlanMetadataOtherResourceController)
		globalCatalogPlanMetadataOtherResourceControllerModel.SubscriptionProviderID = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogPlanMetadataOther)
		model.ResourceController = globalCatalogPlanMetadataOtherResourceControllerModel

		assert.Equal(t, result, model)
	}

	globalCatalogPlanMetadataOtherResourceControllerModel := make(map[string]interface{})
	globalCatalogPlanMetadataOtherResourceControllerModel["subscription_provider_id"] = "testString"

	model := make(map[string]interface{})
	model["resource_controller"] = []interface{}{globalCatalogPlanMetadataOtherResourceControllerModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataOther(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataOtherResourceController(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogPlanMetadataOtherResourceController) {
		model := new(partnercentersellv1.GlobalCatalogPlanMetadataOtherResourceController)
		model.SubscriptionProviderID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["subscription_provider_id"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataOtherResourceController(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataOtherTargetPlansItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogPlanMetadataOtherTargetPlansItem) {
		model := new(partnercentersellv1.GlobalCatalogPlanMetadataOtherTargetPlansItem)
		model.ID = core.StringPtr("testString")
		model.Name = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["name"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogPlanMapToGlobalCatalogPlanMetadataOtherTargetPlansItem(model)
	assert.Nil(t, err)
	checkResult(result)
}
