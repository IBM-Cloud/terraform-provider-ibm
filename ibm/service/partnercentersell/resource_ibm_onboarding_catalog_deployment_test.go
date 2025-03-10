// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package partnercentersell_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

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
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfigBasic(name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogDeploymentExists("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfigBasic(nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
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
	env := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	objectID := fmt.Sprintf("tf_object_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	active := "true"
	disabled := "true"
	kind := "deployment"
	envUpdate := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	objectIDUpdate := fmt.Sprintf("tf_object_id_%d", acctest.RandIntRange(10, 100))
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
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfig(env, objectID, name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogDeploymentExists("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "env", env),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "object_id", objectID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogDeploymentConfig(envUpdate, objectIDUpdate, nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "env", envUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance", "object_id", objectIDUpdate),
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

func testAccCheckIbmOnboardingCatalogDeploymentConfigBasic(name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_catalog_deployment" "onboarding_catalog_deployment_instance" {
			product_id = ibm_onboarding_product.onboarding_product_instance.id
			catalog_product_id = ibm_onboarding_catalog_product.onboarding_catalog_product_instance.onboarding_catalog_product_id
			catalog_plan_id = ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance.onboarding_catalog_plan_id
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			object_provider {
				name = "name"
				email = "email"
			}
		}
	`, name, active, disabled, kind)
}

func testAccCheckIbmOnboardingCatalogDeploymentConfig(env string, objectID string, name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_catalog_deployment" "onboarding_catalog_deployment_instance" {
			product_id = ibm_onboarding_product.onboarding_product_instance.id
			catalog_product_id = ibm_onboarding_catalog_product.onboarding_catalog_product_instance.onboarding_catalog_product_id
			catalog_plan_id = ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance.onboarding_catalog_plan_id
			env = "%s"
			object_id = "%s"
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
				service {
					rc_provisionable = true
					iam_compatible = true
					service_key_supported = true
					parameters {
						displayname = "displayname"
						name = "name"
						type = "text"
						options {
							displayname = "displayname"
							value = "value"
							i18n {
								en {
									displayname = "displayname"
									description = "description"
								}
								de {
									displayname = "displayname"
									description = "description"
								}
								es {
									displayname = "displayname"
									description = "description"
								}
								fr {
									displayname = "displayname"
									description = "description"
								}
								it {
									displayname = "displayname"
									description = "description"
								}
								ja {
									displayname = "displayname"
									description = "description"
								}
								ko {
									displayname = "displayname"
									description = "description"
								}
								pt_br {
									displayname = "displayname"
									description = "description"
								}
								zh_tw {
									displayname = "displayname"
									description = "description"
								}
								zh_cn {
									displayname = "displayname"
									description = "description"
								}
							}
						}
						value = [ "value" ]
						layout = "layout"
						associations = { "key" = "anything as a string" }
						validation_url = "validation_url"
						options_url = "options_url"
						invalidmessage = "invalidmessage"
						description = "description"
						required = true
						pattern = "pattern"
						placeholder = "placeholder"
						readonly = true
						hidden = true
						i18n {
							en {
								displayname = "displayname"
								description = "description"
							}
							de {
								displayname = "displayname"
								description = "description"
							}
							es {
								displayname = "displayname"
								description = "description"
							}
							fr {
								displayname = "displayname"
								description = "description"
							}
							it {
								displayname = "displayname"
								description = "description"
							}
							ja {
								displayname = "displayname"
								description = "description"
							}
							ko {
								displayname = "displayname"
								description = "description"
							}
							pt_br {
								displayname = "displayname"
								description = "description"
							}
							zh_tw {
								displayname = "displayname"
								description = "description"
							}
							zh_cn {
								displayname = "displayname"
								description = "description"
							}
						}
					}
				}
				deployment {
					broker {
						name = "name"
						guid = "guid"
					}
					location = "location"
					location_url = "location_url"
					target_crn = "target_crn"
				}
			}
		}
	`, env, objectID, name, active, disabled, kind)
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
