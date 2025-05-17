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
	objectId := fmt.Sprintf("test-object-id-terraform-2-%d", acctest.RandIntRange(10, 100))
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

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckPartnerCenterSell(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfig(productID, catalogProductID, catalogPlanID, env, name, active, disabled, kind, objectId, overviewUiEn, rcCompatible, iamCompatible, deploymentBrokerName),
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
				Config: testAccCheckIbmOnboardingCatalogDeploymentUpdateConfig(productID, catalogProductID, catalogPlanID, envUpdate, nameUpdate, activeUpdate, disabledUpdate, kindUpdate, objectId, overviewUiEnUpdate, rcCompatibleUpdate, iamCompatibleUpdate, deploymentBrokerNameUpdate),
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
			object_provider {
				name = "name"
				email = "email@email.com"
			}
			metadata {
                rc_compatible =	false
				service {
				  	rc_provisionable = true
  					iam_compatible = true
		}
            }
		}
	`, productID, catalogProductID, catalogPlanID, name, active, disabled, kind, objectId)
}

func testAccCheckIbmOnboardingCatalogDeploymentConfig(productID string, catalogProductID string, catalogPlanID string, env string, name string, active string, disabled string, kind string, objectId string, overviewUiEn string, rcCompatible string, iamCompatible string, deploymentBrokerName string) string {
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
                rc_compatible =	"%s"
				service {
					rc_provisionable = true
  					iam_compatible = "%s"
					service_key_supported = true
					parameters {
                		displayname = "test"
                		name = "test"
						type = "text"
                		value = ["test"]
                		description = "test"
						associations {
							plan {
								show_for = [ "plan-id" ]
								options_refresh = true
							}
							parameters {
								name = "name"
								show_for = [ "parameter" ]
								options_refresh = true
							}
							location {
								show_for = [ "eu-gb" ]
							}
						}
            		}
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
            }
		}
	`, productID, catalogProductID, catalogPlanID, env, name, active, disabled, kind, objectId, overviewUiEn, rcCompatible, iamCompatible, deploymentBrokerName)
}

func testAccCheckIbmOnboardingCatalogDeploymentUpdateConfig(productID string, catalogProductID string, catalogPlanID string, env string, name string, active string, disabled string, kind string, objectId string, overviewUiEn string, rcCompatible string, iamCompatible string, deploymentBrokerName string) string {
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
                rc_compatible =	"%s"
				service {
				  	rc_provisionable = true
  					iam_compatible = "%s"
					service_key_supported = false
					parameters {
                		displayname = "test"
                		name = "test"
			    		type = "text"
                		value = ["test"]
                		description = "test"
						associations {
							plan {
								show_for = [ "plan-id", "plan-id-2" ]
								options_refresh = false
							}
							parameters {
								name = "name"
								show_for = [ "parameter", "parameter2" ]
								options_refresh = false
							}
							parameters {
								name = "name2"
								show_for = [ "parameter2" ]
								options_refresh = true
							}
							location {
								show_for = [ "us-east" ]
							}
						}
					}
					parameters {
                		displayname = "test2"
                		name = "test2"
			    		type = "text"
                		value = ["test2"]
                		description = "test2"
					}
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
			}
		}
	`, productID, catalogProductID, catalogPlanID, env, name, active, disabled, kind, objectId, overviewUiEn, rcCompatible, iamCompatible, deploymentBrokerName)
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

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

		globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
		globalCatalogMetadataServiceCustomParametersModel["options"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel["value"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["associations"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
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
		model["service"] = []map[string]interface{}{globalCatalogDeploymentMetadataServiceModel}
		model["deployment"] = []map[string]interface{}{globalCatalogMetadataDeploymentModel}

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

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

	globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
	globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
	globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
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
	model.Service = globalCatalogDeploymentMetadataServiceModel
	model.Deployment = globalCatalogMetadataDeploymentModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogDeploymentMetadataToMap(model)
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

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

		globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
		globalCatalogMetadataServiceCustomParametersModel["options"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel["value"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["associations"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
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
		model["unique_api_key"] = true
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

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

	globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
	globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
	globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
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
	model.UniqueApiKey = core.BoolPtr(true)
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

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

		model := make(map[string]interface{})
		model["displayname"] = "testString"
		model["name"] = "testString"
		model["type"] = "text"
		model["options"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
		model["value"] = []string{"testString"}
		model["layout"] = "testString"
		model["associations"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
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

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
	model.Displayname = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Type = core.StringPtr("text")
	model.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
	model.Value = []string{"testString"}
	model.Layout = core.StringPtr("testString")
	model.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
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

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersAssociationsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []string{"testString"}

		model := make(map[string]interface{})
		model["plan"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
		model["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		model["location"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
	model.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
	model.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	model.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersAssociationsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersAssociationsPlanToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["show_for"] = []string{"testString"}
		model["options_refresh"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	model.ShowFor = []string{"testString"}
	model.OptionsRefresh = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersAssociationsPlanToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersAssociationsParametersItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["show_for"] = []string{"testString"}
		model["options_refresh"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	model.Name = core.StringPtr("testString")
	model.ShowFor = []string{"testString"}
	model.OptionsRefresh = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersAssociationsParametersItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersAssociationsLocationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["show_for"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	model.ShowFor = []string{"testString"}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentGlobalCatalogMetadataServiceCustomParametersAssociationsLocationToMap(model)
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

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogDeploymentMetadataPrototypePatch(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogDeploymentMetadataPrototypePatch) {
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

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

		globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
		globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
		globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
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

		globalCatalogDeploymentMetadataServicePrototypePatchModel := new(partnercentersellv1.GlobalCatalogDeploymentMetadataServicePrototypePatch)
		globalCatalogDeploymentMetadataServicePrototypePatchModel.RcProvisionable = core.BoolPtr(true)
		globalCatalogDeploymentMetadataServicePrototypePatchModel.IamCompatible = core.BoolPtr(true)
		globalCatalogDeploymentMetadataServicePrototypePatchModel.ServiceKeySupported = core.BoolPtr(true)
		globalCatalogDeploymentMetadataServicePrototypePatchModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel}

		globalCatalogMetadataDeploymentBrokerModel := new(partnercentersellv1.GlobalCatalogMetadataDeploymentBroker)
		globalCatalogMetadataDeploymentBrokerModel.Name = core.StringPtr("testString")
		globalCatalogMetadataDeploymentBrokerModel.Guid = core.StringPtr("testString")

		globalCatalogMetadataDeploymentModel := new(partnercentersellv1.GlobalCatalogMetadataDeployment)
		globalCatalogMetadataDeploymentModel.Broker = globalCatalogMetadataDeploymentBrokerModel
		globalCatalogMetadataDeploymentModel.Location = core.StringPtr("testString")
		globalCatalogMetadataDeploymentModel.LocationURL = core.StringPtr("testString")
		globalCatalogMetadataDeploymentModel.TargetCrn = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogDeploymentMetadataPrototypePatch)
		model.RcCompatible = core.BoolPtr(true)
		model.Service = globalCatalogDeploymentMetadataServicePrototypePatchModel
		model.Deployment = globalCatalogMetadataDeploymentModel

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

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []interface{}{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

	globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
	globalCatalogMetadataServiceCustomParametersModel["options"] = []interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel["value"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["associations"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
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

	globalCatalogDeploymentMetadataServicePrototypePatchModel := make(map[string]interface{})
	globalCatalogDeploymentMetadataServicePrototypePatchModel["rc_provisionable"] = true
	globalCatalogDeploymentMetadataServicePrototypePatchModel["iam_compatible"] = true
	globalCatalogDeploymentMetadataServicePrototypePatchModel["service_key_supported"] = true
	globalCatalogDeploymentMetadataServicePrototypePatchModel["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersModel}

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
	model["service"] = []interface{}{globalCatalogDeploymentMetadataServicePrototypePatchModel}
	model["deployment"] = []interface{}{globalCatalogMetadataDeploymentModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogDeploymentMetadataPrototypePatch(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogDeploymentMetadataServicePrototypePatch(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogDeploymentMetadataServicePrototypePatch) {
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

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

		globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
		globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
		globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
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

		model := new(partnercentersellv1.GlobalCatalogDeploymentMetadataServicePrototypePatch)
		model.RcProvisionable = core.BoolPtr(true)
		model.IamCompatible = core.BoolPtr(true)
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

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []interface{}{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

	globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
	globalCatalogMetadataServiceCustomParametersModel["options"] = []interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel["value"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["associations"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
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
	model["service_key_supported"] = true
	model["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogDeploymentMetadataServicePrototypePatch(model)
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

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
		model.Displayname = core.StringPtr("testString")
		model.Name = core.StringPtr("testString")
		model.Type = core.StringPtr("text")
		model.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
		model.Value = []string{"testString"}
		model.Layout = core.StringPtr("testString")
		model.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
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

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []interface{}{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

	model := make(map[string]interface{})
	model["displayname"] = "testString"
	model["name"] = "testString"
	model["type"] = "text"
	model["options"] = []interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
	model["value"] = []interface{}{"testString"}
	model["layout"] = "testString"
	model["associations"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
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

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersAssociations(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations) {
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
		model.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
		model.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		model.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []interface{}{"testString"}

	model := make(map[string]interface{})
	model["plan"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
	model["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	model["location"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersAssociations(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersAssociationsPlan(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan) {
		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		model.ShowFor = []string{"testString"}
		model.OptionsRefresh = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["show_for"] = []interface{}{"testString"}
	model["options_refresh"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersAssociationsPlan(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem) {
		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		model.Name = core.StringPtr("testString")
		model.ShowFor = []string{"testString"}
		model.OptionsRefresh = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["show_for"] = []interface{}{"testString"}
	model["options_refresh"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersAssociationsLocation(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation) {
		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		model.ShowFor = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["show_for"] = []interface{}{"testString"}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogDeploymentMapToGlobalCatalogMetadataServiceCustomParametersAssociationsLocation(model)
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
