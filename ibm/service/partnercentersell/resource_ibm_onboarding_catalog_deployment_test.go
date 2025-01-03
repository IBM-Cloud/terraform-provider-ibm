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
	productID := acc.PcsOnboardingProductWithCatalogProduct
	catalogProductID := acc.PcsOnboardingCatalogProductId
	catalogPlanID := acc.PcsOnboardingCatalogPlanId
	objectId := fmt.Sprintf("test-object-id-terraform-%d", acctest.RandIntRange(10, 100))
	name := "test-deployment-name-terraform"
	active := "true"
	disabled := "false"
	kind := "deployment"
	nameUpdate := "test-deployment-name-terraform"
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "deployment"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckPartnerCenterSell(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfigBasic(productID, catalogProductID, catalogPlanID, name, active, disabled, kind, objectId),
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
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfigBasic(productID, catalogProductID, catalogPlanID, nameUpdate, activeUpdate, disabledUpdate, kindUpdate, objectId),
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
	productID := acc.PcsOnboardingProductWithCatalogProduct
	catalogProductID := acc.PcsOnboardingCatalogProductId
	catalogPlanID := acc.PcsOnboardingCatalogPlanId
	objectId := fmt.Sprintf("test-object-id-terraform-%d", acctest.RandIntRange(10, 100))
	env := "current"
	name := "test-deployment-name-terraform"
	active := "true"
	disabled := "false"
	kind := "deployment"
	envUpdate := "current"
	nameUpdate := "test-deployment-name-terraform"
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "deployment"
	overviewUiEn := "display_name"
	overviewUiEnUpdate := "display_name_2"
	rcCompatible := "true"
	rcCompatibleUpdate := "false"
	iamCompatible := "true"
	iamCompatibleUpdate := "false"
	deploymentBrokerName := "broker-petra-1"
	deploymentBrokerNameUpdate := "broker-petra-2"
	bulletTitleName := "title"
	bulletTitleNameUpdate := "title-2"
	mediaCaption := "Some random minecraft Video"
	mediaCaptionUpdate := "Some random minecraft Video 2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckPartnerCenterSell(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfig(productID, catalogProductID, catalogPlanID, env, name, active, disabled, kind, objectId, overviewUiEn, rcCompatible, iamCompatible, deploymentBrokerName, bulletTitleName, mediaCaption),
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
				Config: testAccCheckIbmOnboardingCatalogDeploymentUpdateConfig(productID, catalogProductID, catalogPlanID, envUpdate, nameUpdate, activeUpdate, disabledUpdate, kindUpdate, objectId, overviewUiEnUpdate, rcCompatibleUpdate, iamCompatibleUpdate, deploymentBrokerNameUpdate, bulletTitleNameUpdate, mediaCaptionUpdate),
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
				ResourceName:      "ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance",
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateVerifyIgnore: []string{
					"env", "product_id", "catalog_product_id", "catalog_plan_id", "geo_tags",
				},
			},
		},
	})
}

func testAccCheckIbmOnboardingCatalogDeploymentConfigBasic(productID string, catalogProductID string, catalogPlanID string, name string, active string, disabled string, kind string, objectId string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_catalog_deployment" "onboarding_catalog_deployment_instance" {
			product_id = "%s"
			catalog_product_id = "%s"
			catalog_plan_id = "%s"
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			object_id = "%s"
			tags = ["sample"]
			object_provider {
				name = "name"
				email = "email@email.com"
			}
			metadata {
				service {
				  	rc_provisionable = true
  					iam_compatible = true
		}
                rc_compatible =	false
            }
		
		}
	`, productID, catalogProductID, catalogPlanID, name, active, disabled, kind, objectId)
}

func testAccCheckIbmOnboardingCatalogDeploymentConfig(productID string, catalogProductID string, catalogPlanID string, env string, name string, active string, disabled string, kind string, objectId string, overviewUiEn string, rcCompatible string, iamCompatible string, deploymentBrokerName string, bulletTitleName string, mediaCaption string) string {
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
			object_id = "%s"
			overview_ui {
				en {
					display_name = "%s"
					description = "description"
					long_description = "long_description"
				}
			}
			tags = ["sample"]
			object_provider {
				name = "name"
				email = "email@email.com"
			}
			metadata {
				rc_compatible = "%s"
				service {
				  	rc_provisionable = true
  					iam_compatible = "%s"
				}
				deployment {
					broker {
						name = "%s"
						guid = "guid"
					}
					location = "ams03"
					location_url = "https://globalcatalog.test.cloud.ibm.com/api/v1/ams03"
					target_crn = "crn:v1:staging:public::ams03:::environment:staging-ams03"
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
	`, productID, catalogProductID, catalogPlanID, env, name, active, disabled, kind, objectId, overviewUiEn, rcCompatible, iamCompatible, deploymentBrokerName, bulletTitleName, mediaCaption)
}

func testAccCheckIbmOnboardingCatalogDeploymentUpdateConfig(productID string, catalogProductID string, catalogPlanID string, env string, name string, active string, disabled string, kind string, objectId string, overviewUiEn string, rcCompatible string, iamCompatible string, deploymentBrokerName string, bulletTitleName string, mediaCaption string) string {
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
			object_id = "%s"
			overview_ui {
				en {
					display_name = "%s"
									description = "description"
					long_description = "long_description"
								}
			}
			tags = ["sample", "moresample"]
			object_provider {
				name = "name"
				email = "email@email.com"
								}
			metadata {
				rc_compatible = "%s"
				service {
				  	rc_provisionable = true
  					iam_compatible = "%s"
				}
				deployment {
					broker {
						name = "%s"
						guid = "guid"
								}
					location = "ams03"
					location_url = "https://globalcatalog.test.cloud.ibm.com/api/v1/ams03"
					target_crn = "crn:v1:staging:public::ams03:::environment:staging-ams03"
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
	`, productID, catalogProductID, catalogPlanID, env, name, active, disabled, kind, objectId, overviewUiEn, rcCompatible, iamCompatible, deploymentBrokerName, bulletTitleName, mediaCaption)
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
		globalCatalogMetadataUiStringsContentModel["embeddable_dashboard"] = "testString"

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

		globalCatalogMetadataUiModel := make(map[string]interface{})
		globalCatalogMetadataUiModel["strings"] = []map[string]interface{}{globalCatalogMetadataUiStringsModel}
		globalCatalogMetadataUiModel["urls"] = []map[string]interface{}{globalCatalogMetadataUiUrlsModel}
		globalCatalogMetadataUiModel["hidden"] = true
		globalCatalogMetadataUiModel["side_by_side_index"] = float64(72.5)

		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
		globalCatalogMetadataServiceCustomParametersModel["options"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel["value"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["associations"] = map[string]interface{}{"anyKey": "anyValue"}
		globalCatalogMetadataServiceCustomParametersModel["validation_url"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["options_url"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["invalidmessage"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["description"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["required"] = true
		globalCatalogMetadataServiceCustomParametersModel["pattern"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["placeholder"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["readonly"] = true
		globalCatalogMetadataServiceCustomParametersModel["hidden"] = true
		globalCatalogMetadataServiceCustomParametersModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		globalCatalogDeploymentMetadataServiceModel := make(map[string]interface{})
		globalCatalogDeploymentMetadataServiceModel["rc_provisionable"] = true
		globalCatalogDeploymentMetadataServiceModel["iam_compatible"] = true
		globalCatalogDeploymentMetadataServiceModel["bindable"] = true
		globalCatalogDeploymentMetadataServiceModel["plan_updateable"] = true
		globalCatalogDeploymentMetadataServiceModel["service_key_supported"] = true
		globalCatalogDeploymentMetadataServiceModel["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersModel}

		globalCatalogMetadataDeploymentBrokerModel := make(map[string]interface{})
		globalCatalogMetadataDeploymentBrokerModel["name"] = "testString"
		globalCatalogMetadataDeploymentBrokerModel["guid"] = "testString"

		globalCatalogMetadataDeploymentModel := make(map[string]interface{})
		globalCatalogMetadataDeploymentModel["broker"] = []map[string]interface{}{globalCatalogMetadataDeploymentBrokerModel}
		globalCatalogMetadataDeploymentModel["location"] = "testString"
		globalCatalogMetadataDeploymentModel["location_url"] = "testString"
		globalCatalogMetadataDeploymentModel["target_crn"] = "testString"

		model := make(map[string]interface{})
		model["rc_compatible"] = true
		model["ui"] = []map[string]interface{}{globalCatalogMetadataUiModel}
		model["service"] = []map[string]interface{}{globalCatalogDeploymentMetadataServiceModel}
		model["deployment"] = []map[string]interface{}{globalCatalogMetadataDeploymentModel}

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
	globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

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

	globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
	globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
	globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
	globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
	globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
	globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
	globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
	globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Associations = map[string]interface{}{"anyKey": "anyValue"}
	globalCatalogMetadataServiceCustomParametersModel.ValidationURL = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.OptionsURL = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Invalidmessage = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Description = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Required = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.Pattern = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Placeholder = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Readonly = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.Hidden = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	globalCatalogDeploymentMetadataServiceModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadataService)
	globalCatalogDeploymentMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
	globalCatalogDeploymentMetadataServiceModel.IamCompatible = core.BoolPtr(true)
	globalCatalogDeploymentMetadataServiceModel.Bindable = core.BoolPtr(true)
	globalCatalogDeploymentMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
	globalCatalogDeploymentMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)
	globalCatalogDeploymentMetadataServiceModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel}

	globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
	globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("testString")
	globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("testString")

	globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
	globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
	globalCatalogMetadataDeploymentModel.Location = core.StringPtr("testString")
	globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("testString")
	globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
	model.RcCompatible = core.BoolPtr(true)
	model.Ui = globalCatalogMetadataUiModel
	model.Service = globalCatalogDeploymentMetadataServiceModel
	model.Deployment = globalCatalogMetadataDeploymentModel

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
		globalCatalogMetadataUiStringsContentModel["embeddable_dashboard"] = "testString"

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
	globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

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
		globalCatalogMetadataUiStringsContentModel["embeddable_dashboard"] = "testString"

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
	globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

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
		model["embeddable_dashboard"] = "testString"

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
	model.EmbeddableDashboard = core.StringPtr("testString")

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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataUIUrlsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogDeploymentMetadataServiceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
		globalCatalogMetadataServiceCustomParametersModel["options"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel["value"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["associations"] = map[string]interface{}{"anyKey": "anyValue"}
		globalCatalogMetadataServiceCustomParametersModel["validation_url"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["options_url"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["invalidmessage"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["description"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["required"] = true
		globalCatalogMetadataServiceCustomParametersModel["pattern"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["placeholder"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["readonly"] = true
		globalCatalogMetadataServiceCustomParametersModel["hidden"] = true
		globalCatalogMetadataServiceCustomParametersModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		model := make(map[string]interface{})
		model["rc_provisionable"] = true
		model["iam_compatible"] = true
		model["bindable"] = true
		model["plan_updateable"] = true
		model["service_key_supported"] = true
		model["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
	globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
	globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
	globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Associations = map[string]interface{}{"anyKey": "anyValue"}
	globalCatalogMetadataServiceCustomParametersModel.ValidationURL = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.OptionsURL = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Invalidmessage = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Description = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Required = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.Pattern = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Placeholder = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Readonly = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.Hidden = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	model := new(partnercentersellv1.GlobalCatalogDeploymentMetadataService)
	model.RcProvisionable = core.BoolPtr(true)
	model.IamCompatible = core.BoolPtr(true)
	model.Bindable = core.BoolPtr(true)
	model.PlanUpdateable = core.BoolPtr(true)
	model.ServiceKeySupported = core.BoolPtr(true)
	model.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogDeploymentMetadataServiceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		model := make(map[string]interface{})
		model["displayname"] = "testString"
		model["name"] = "testString"
		model["type"] = "text"
		model["options"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
		model["value"] = []string{"testString"}
		model["layout"] = "testString"
		model["associations"] = map[string]interface{}{"anyKey": "anyValue"}
		model["validation_url"] = "testString"
		model["options_url"] = "testString"
		model["invalidmessage"] = "testString"
		model["description"] = "testString"
		model["required"] = true
		model["pattern"] = "testString"
		model["placeholder"] = "testString"
		model["readonly"] = true
		model["hidden"] = true
		model["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
	globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
	model.Displayname = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Type = core.StringPtr("text")
	model.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
	model.Value = []string{"testString"}
	model.Layout = core.StringPtr("testString")
	model.Associations = map[string]interface{}{"anyKey": "anyValue"}
	model.ValidationURL = core.StringPtr("testString")
	model.OptionsURL = core.StringPtr("testString")
	model.Invalidmessage = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.Required = core.BoolPtr(true)
	model.Pattern = core.StringPtr("testString")
	model.Placeholder = core.StringPtr("testString")
	model.Readonly = core.BoolPtr(true)
	model.Hidden = core.BoolPtr(true)
	model.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		model := make(map[string]interface{})
		model["displayname"] = "testString"
		model["value"] = "testString"
		model["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
	model.Displayname = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")
	model.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersI18nToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		model := make(map[string]interface{})
		model["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	model.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersI18nToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersI18nFieldsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["displayname"] = "testString"
		model["description"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	model.Displayname = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersI18nFieldsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataDeploymentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataDeploymentBrokerModel := make(map[string]interface{})
		globalCatalogMetadataDeploymentBrokerModel["name"] = "testString"
		globalCatalogMetadataDeploymentBrokerModel["guid"] = "testString"

		model := make(map[string]interface{})
		model["broker"] = []map[string]interface{}{globalCatalogMetadataDeploymentBrokerModel}
		model["location"] = "testString"
		model["location_url"] = "testString"
		model["target_crn"] = "testString"

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
	globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("testString")
	globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
	model.Broker = globalCatalogMetadataDeploymentBrokerModel
	model.Location = core.StringPtr("testString")
	model.LocationURL = core.StringPtr("testString")
	model.TargetCrn = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataDeploymentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataDeploymentBrokerToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["guid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
	model.Name = core.StringPtr("testString")
	model.Guid = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataDeploymentBrokerToMap(model)
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
		globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

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

		globalCatalogMetadataUiModel := new(partnercentersellv1.GlobalCatalogMetadataUI)
		globalCatalogMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
		globalCatalogMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
		globalCatalogMetadataUiModel.Hidden = core.BoolPtr(true)
		globalCatalogMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))

		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
		globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
		globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
		globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Associations = map[string]interface{}{"anyKey": "anyValue"}
		globalCatalogMetadataServiceCustomParametersModel.ValidationURL = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.OptionsURL = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Invalidmessage = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Description = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Required = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.Pattern = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Placeholder = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Readonly = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.Hidden = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		globalCatalogDeploymentMetadataServiceModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadataService)
		globalCatalogDeploymentMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
		globalCatalogDeploymentMetadataServiceModel.IamCompatible = core.BoolPtr(true)
		globalCatalogDeploymentMetadataServiceModel.Bindable = core.BoolPtr(true)
		globalCatalogDeploymentMetadataServiceModel.PlanUpdateable = core.BoolPtr(true)
		globalCatalogDeploymentMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)
		globalCatalogDeploymentMetadataServiceModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel}

		globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
		globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("testString")
		globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("testString")

		globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
		globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
		globalCatalogMetadataDeploymentModel.Location = core.StringPtr("testString")
		globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("testString")
		globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogDeploymentMetadata)
		model.RcCompatible = core.BoolPtr(true)
		model.Ui = globalCatalogMetadataUiModel
		model.Service = globalCatalogDeploymentMetadataServiceModel
		model.Deployment = globalCatalogMetadataDeploymentModel

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
	globalCatalogMetadataUiStringsContentModel["embeddable_dashboard"] = "testString"

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

	globalCatalogMetadataUiModel := make(map[string]interface{})
	globalCatalogMetadataUiModel["strings"] = []interface{}{globalCatalogMetadataUiStringsModel}
	globalCatalogMetadataUiModel["urls"] = []interface{}{globalCatalogMetadataUiUrlsModel}
	globalCatalogMetadataUiModel["hidden"] = true
	globalCatalogMetadataUiModel["side_by_side_index"] = float64(72.5)

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
	globalCatalogMetadataServiceCustomParametersModel["options"] = []interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel["value"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["associations"] = map[string]interface{}{"anyKey": "anyValue"}
	globalCatalogMetadataServiceCustomParametersModel["validation_url"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["options_url"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["invalidmessage"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["description"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["required"] = true
	globalCatalogMetadataServiceCustomParametersModel["pattern"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["placeholder"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["readonly"] = true
	globalCatalogMetadataServiceCustomParametersModel["hidden"] = true
	globalCatalogMetadataServiceCustomParametersModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	globalCatalogDeploymentMetadataServiceModel := make(map[string]interface{})
	globalCatalogDeploymentMetadataServiceModel["rc_provisionable"] = true
	globalCatalogDeploymentMetadataServiceModel["iam_compatible"] = true
	globalCatalogDeploymentMetadataServiceModel["bindable"] = true
	globalCatalogDeploymentMetadataServiceModel["plan_updateable"] = true
	globalCatalogDeploymentMetadataServiceModel["service_key_supported"] = true
	globalCatalogDeploymentMetadataServiceModel["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersModel}

	globalCatalogMetadataDeploymentBrokerModel := make(map[string]interface{})
	globalCatalogMetadataDeploymentBrokerModel["name"] = "testString"
	globalCatalogMetadataDeploymentBrokerModel["guid"] = "testString"

	globalCatalogMetadataDeploymentModel := make(map[string]interface{})
	globalCatalogMetadataDeploymentModel["broker"] = []interface{}{globalCatalogMetadataDeploymentBrokerModel}
	globalCatalogMetadataDeploymentModel["location"] = "testString"
	globalCatalogMetadataDeploymentModel["location_url"] = "testString"
	globalCatalogMetadataDeploymentModel["target_crn"] = "testString"

	model := make(map[string]interface{})
	model["rc_compatible"] = true
	model["ui"] = []interface{}{globalCatalogMetadataUiModel}
	model["service"] = []interface{}{globalCatalogDeploymentMetadataServiceModel}
	model["deployment"] = []interface{}{globalCatalogMetadataDeploymentModel}

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
		globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

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
	globalCatalogMetadataUiStringsContentModel["embeddable_dashboard"] = "testString"

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
		globalCatalogMetadataUiStringsContentModel.EmbeddableDashboard = core.StringPtr("testString")

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
	globalCatalogMetadataUiStringsContentModel["embeddable_dashboard"] = "testString"

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
		model.EmbeddableDashboard = core.StringPtr("testString")

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
	model["embeddable_dashboard"] = "testString"

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

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataUIUrls(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogDeploymentMetadataService(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogDeploymentMetadataService) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
		globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
		globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
		globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Associations = map[string]interface{}{"anyKey": "anyValue"}
		globalCatalogMetadataServiceCustomParametersModel.ValidationURL = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.OptionsURL = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Invalidmessage = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Description = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Required = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.Pattern = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Placeholder = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Readonly = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.Hidden = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		model := new(partnercentersellv1.GlobalCatalogDeploymentMetadataService)
		model.RcProvisionable = core.BoolPtr(true)
		model.IamCompatible = core.BoolPtr(true)
		model.Bindable = core.BoolPtr(true)
		model.PlanUpdateable = core.BoolPtr(true)
		model.ServiceKeySupported = core.BoolPtr(true)
		model.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
	globalCatalogMetadataServiceCustomParametersModel["options"] = []interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel["value"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["associations"] = map[string]interface{}{"anyKey": "anyValue"}
	globalCatalogMetadataServiceCustomParametersModel["validation_url"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["options_url"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["invalidmessage"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["description"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["required"] = true
	globalCatalogMetadataServiceCustomParametersModel["pattern"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["placeholder"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["readonly"] = true
	globalCatalogMetadataServiceCustomParametersModel["hidden"] = true
	globalCatalogMetadataServiceCustomParametersModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	model := make(map[string]interface{})
	model["rc_provisionable"] = true
	model["iam_compatible"] = true
	model["bindable"] = true
	model["plan_updateable"] = true
	model["service_key_supported"] = true
	model["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogDeploymentMetadataService(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParameters(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
		globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
		model.Displayname = core.StringPtr("testString")
		model.Name = core.StringPtr("testString")
		model.Type = core.StringPtr("text")
		model.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
		model.Value = []string{"testString"}
		model.Layout = core.StringPtr("testString")
		model.Associations = map[string]interface{}{"anyKey": "anyValue"}
		model.ValidationURL = core.StringPtr("testString")
		model.OptionsURL = core.StringPtr("testString")
		model.Invalidmessage = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")
		model.Required = core.BoolPtr(true)
		model.Pattern = core.StringPtr("testString")
		model.Placeholder = core.StringPtr("testString")
		model.Readonly = core.BoolPtr(true)
		model.Hidden = core.BoolPtr(true)
		model.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	model := make(map[string]interface{})
	model["displayname"] = "testString"
	model["name"] = "testString"
	model["type"] = "text"
	model["options"] = []interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
	model["value"] = []interface{}{"testString"}
	model["layout"] = "testString"
	model["associations"] = map[string]interface{}{"anyKey": "anyValue"}
	model["validation_url"] = "testString"
	model["options_url"] = "testString"
	model["invalidmessage"] = "testString"
	model["description"] = "testString"
	model["required"] = true
	model["pattern"] = "testString"
	model["placeholder"] = "testString"
	model["readonly"] = true
	model["hidden"] = true
	model["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersOptions(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
		model.Displayname = core.StringPtr("testString")
		model.Value = core.StringPtr("testString")
		model.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	model := make(map[string]interface{})
	model["displayname"] = "testString"
	model["value"] = "testString"
	model["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersOptions(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersI18n(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		model.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	model := make(map[string]interface{})
	model["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersI18n(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersI18nFields(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields) {
		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		model.Displayname = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["displayname"] = "testString"
	model["description"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersI18nFields(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataDeployment(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataDeployment) {
		globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
		globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("testString")
		globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
		model.Broker = globalCatalogMetadataDeploymentBrokerModel
		model.Location = core.StringPtr("testString")
		model.LocationURL = core.StringPtr("testString")
		model.TargetCrn = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataDeploymentBrokerModel := make(map[string]interface{})
	globalCatalogMetadataDeploymentBrokerModel["name"] = "testString"
	globalCatalogMetadataDeploymentBrokerModel["guid"] = "testString"

	model := make(map[string]interface{})
	model["broker"] = []interface{}{globalCatalogMetadataDeploymentBrokerModel}
	model["location"] = "testString"
	model["location_url"] = "testString"
	model["target_crn"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataDeployment(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataDeploymentBroker(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataDeploymentBroker) {
		model := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
		model.Name = core.StringPtr("testString")
		model.Guid = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["guid"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataDeploymentBroker(model)
	assert.Nil(t, err)
	checkResult(result)
}
