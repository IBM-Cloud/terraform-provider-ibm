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

func TestAccIbmOnboardingIamRegistrationBasic(t *testing.T) {
	var conf partnercentersellv1.IamServiceRegistration
	productID := acc.PcsOnboardingProductWithCatalogProduct
	name := acc.PcsIamServiceRegistrationId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckPartnerCenterSell(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingIamRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingIamRegistrationConfigBasic(productID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingIamRegistrationExists("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "product_id", productID),
				),
			},
		},
	})
}

func TestAccIbmOnboardingIamRegistrationAllArgs(t *testing.T) {
	var conf partnercentersellv1.IamServiceRegistration
	productID := acc.PcsOnboardingProductWithCatalogProduct
	env := "current"
	name := acc.PcsIamServiceRegistrationId
	roleDisplayName := fmt.Sprintf("random-%d", acctest.RandIntRange(10, 100))
	iamRegistrationRole := fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", acc.PcsIamServiceRegistrationId, roleDisplayName)
	enabled := "true"
	serviceType := "platform_service"
	envUpdate := "current"
	roleDisplayNameUpdate := fmt.Sprintf("random-%d", acctest.RandIntRange(10, 100))
	iamRegistrationRoleUpdate := fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", acc.PcsIamServiceRegistrationId, roleDisplayNameUpdate)
	nameUpdate := acc.PcsIamServiceRegistrationId
	enabledUpdate := "true"
	serviceTypeUpdate := "service"
	actionDescription := "default"
	actionDescriptionUpdate := "default_2"
	supportedAttributeDisplayName := "default"
	supportedAttributeDisplayNameUpdate := "default_2"
	supportedAttributeInputDetailsDisplayName := "default"
	supportedAttributeInputDetailsDisplayNameUpdate := "default_2"
	supportedAuthorizationSubjectsService := "serviceName"
	supportedAuthorizationSubjectsServiceUpdate := "serviceName2"
	environmentAttributesValues := "public"
	environmentAttributesValuesUpdate := "private"
	supportedAnonymousAccessesAdditionalPropValue := "additional"
	supportedAnonymousAccessesAdditionalPropValueUpdate := "additionals"
	// supportedAnonymousAccessesAccId := "account_id_2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckPartnerCenterSell(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingIamRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingIamRegistrationConfig(
					productID,
					env,
					name,
					enabled,
					serviceType,
					iamRegistrationRole,
					roleDisplayName,
					acc.PcsIamServiceRegistrationId,
					actionDescription,
					supportedAttributeDisplayName,
					supportedAttributeInputDetailsDisplayName,
					supportedAuthorizationSubjectsService,
					environmentAttributesValues,
					supportedAnonymousAccessesAdditionalPropValue,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingIamRegistrationExists("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "env", env),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "service_type", serviceType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingIamRegistrationUpdateConfig(
					productID,
					envUpdate,
					nameUpdate,
					enabledUpdate,
					serviceTypeUpdate,
					iamRegistrationRoleUpdate,
					roleDisplayNameUpdate,
					acc.PcsIamServiceRegistrationId,
					actionDescriptionUpdate,
					supportedAttributeDisplayNameUpdate,
					supportedAttributeInputDetailsDisplayNameUpdate,
					supportedAuthorizationSubjectsServiceUpdate,
					environmentAttributesValuesUpdate,
					supportedAnonymousAccessesAdditionalPropValueUpdate,
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "product_id", productID),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "env", envUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_iam_registration.onboarding_iam_registration_instance", "service_type", serviceTypeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_onboarding_iam_registration.onboarding_iam_registration_instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"env", "product_id", "service_type"},
			},
		},
	})
}

func testAccCheckIbmOnboardingIamRegistrationConfigBasic(productID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_iam_registration" "onboarding_iam_registration_instance" {
			product_id = "%s"
			name = "%s"
			enabled = true
			display_name {
				default = "%s"
			}
		}
	`, productID, name, name)
}

func testAccCheckIbmOnboardingIamRegistrationConfig(
	productID string,
	env string,
	name string,
	enabled string,
	serviceType string,
	iamRegistrationRole string,
	roleDisplayName string,
	iamRegistrationID string,
	actionDescription string,
	supportedAttributeDisplayName string,
	supportedAttributeInputDetailsDisplayName string,
	supportedAuthorizationSubjectsService string,
	environmentAttributesValues string,
	supportedAnonymousAccessesAdditionalPropValue string,
) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_iam_registration" "onboarding_iam_registration_instance" {
			product_id = "%s"
			env = "%s"
			name = "%s"
			enabled = %s
			service_type = "%s"
			actions {
				id = "id"
				roles = [ "%s" ]
				description {
					default = "%s"
					en = "en"
					de = "de"
					es = "es"
					fr = "fr"
					it = "it"
					ja = "ja"
					ko = "ko"
					pt_br = "pt_br"
					zh_tw = "zh_tw"
					zh_cn = "zh_cn"
				}
				display_name {
					default = "default"
					en = "en"
					de = "de"
					es = "es"
					fr = "fr"
					it = "it"
					ja = "ja"
					ko = "ko"
					pt_br = "pt_br"
					zh_tw = "zh_tw"
					zh_cn = "zh_cn"
				}
				options {
					hidden = true
				}
			}
			additional_policy_scopes = ["%s"]
			display_name {
				default = "%s"
				en = "en"
				de = "de"
				es = "es"
				fr = "fr"
				it = "it"
				ja = "ja"
				ko = "ko"
				pt_br = "pt_br"
				zh_tw = "zh_tw"
				zh_cn = "zh_cn"
			}
			parent_ids = ["05ca8653-de25-49fa-a14d-aaa5d373bc21"]	
			supported_attributes {
				key = "testString"
				options {
					operators = [ "stringEquals" ]
					hidden = false
					policy_types = [ "access" ]
					is_empty_value_supported = true
					is_string_exists_false_value_supported = true
				}
				display_name {
					default = "%s"
					en = "en"
					de = "de"
					es = "es"
					fr = "fr"
					it = "it"
					ja = "ja"
					ko = "ko"
					pt_br = "pt_br"
					zh_tw = "zh_tw"
					zh_cn = "zh_cn"
				}
				description {
					default = "default"
					en = "en"
					de = "de"
					es = "es"
					fr = "fr"
					it = "it"
					ja = "ja"
					ko = "ko"
					pt_br = "pt_br"
					zh_tw = "zh_tw"
					zh_cn = "zh_cn"
				}
				ui {
					input_type = "selector"
					input_details {
						type = "gst"
						values {
							value = "testString"
							display_name {
								default = "%s"
								en = "testString"
								de = "testString"
								es = "testString"
								fr = "testString"
								it = "testString"
								ja = "testString"
								ko = "testString"
								pt_br = "testString"
								zh_tw = "testString"
								zh_cn = "testString"
							}
						}
						gst {
							query = "query"
							value_property_name = "teststring"
							input_option_label = "{name} - {instance_id}"
						}
					}
				}
			}
			supported_authorization_subjects {
				attributes {
					service_name = "%s"
					resource_type = "testString"
				}
				roles = [ "%s" ]
			}
			supported_roles {
				id = "%s"
				description {
					default = "desc"
				}
				display_name {
					default = "%s"
				}
				options {
					access_policy = true
					policy_type = [ "access" ]
				}
			}
			supported_network {
				environment_attributes {
					key = "networkType"
					values = [ "%s" ]
					options {
						hidden = false
					}
				}
			}
			supported_anonymous_accesses {
				attributes {
					account_id = "account_id"
					service_name = "%s"
					additional_properties = { "testString" = "%s" }
				}
				roles = [ "%s" ]
			}
		}
	`, productID, env, name, enabled, serviceType, iamRegistrationRole, actionDescription, name, name, supportedAttributeDisplayName, supportedAttributeInputDetailsDisplayName, supportedAuthorizationSubjectsService, iamRegistrationRole, iamRegistrationRole, roleDisplayName, environmentAttributesValues, iamRegistrationID, supportedAnonymousAccessesAdditionalPropValue, iamRegistrationRole)
}

func testAccCheckIbmOnboardingIamRegistrationUpdateConfig(
	productID string,
	env string,
	name string,
	enabled string,
	serviceType string,
	iamRegistrationRole string,
	roleDisplayName string,
	iamRegistrationID string,
	actionDescription string,
	supportedAttributeDisplayName string,
	supportedAttributeInputDetailsDisplayName string,
	supportedAuthorizationSubjectsService string,
	environmentAttributesValues string,
	supportedAnonymousAccessesAdditionalPropValue string,
) string {
	roleDisplayName2 := fmt.Sprintf("random-2-%d", acctest.RandIntRange(10, 100))
	iamRegistrationRole2 := fmt.Sprintf("crn:v1:bluemix:public:%s::::serviceRole:%s", iamRegistrationID, roleDisplayName2)

	return fmt.Sprintf(`
		resource "ibm_onboarding_iam_registration" "onboarding_iam_registration_instance" {
			product_id = "%s"
			env = "%s"
			name = "%s"
			enabled = %s
			service_type = "%s"
			actions {
				id = "id"
				roles = [ "%s", "%s" ]
				description {
					default = "%s"
					en = "en"
					de = "de"
					es = "es"
					fr = "fr"
					it = "it"
					ja = "ja"
					ko = "ko"
					pt_br = "pt_br"
					zh_tw = "zh_tw"
					zh_cn = "zh_cn"
				}
							display_name {
								default = "default"
								en = "en"
								de = "de"
								es = "es"
								fr = "fr"
								it = "it"
								ja = "ja"
								ko = "ko"
								pt_br = "pt_br"
								zh_tw = "zh_tw"
								zh_cn = "zh_cn"
							}
				options {
					hidden = true
				}
			}
			actions {
				id = "idtwo"
				roles = [ "%s" ]
				description {
					default = "default"
					en = "en"
					de = "de"
					es = "es"
					fr = "fr"
					it = "it"
					ja = "ja"
					ko = "ko"
					pt_br = "pt_br"
					zh_tw = "zh_tw"
					zh_cn = "zh_cn"
				}
				display_name {
					default = "default"
					en = "en"
					de = "de"
					es = "es"
					fr = "fr"
					it = "it"
					ja = "ja"
					ko = "ko"
					pt_br = "pt_br"
					zh_tw = "zh_tw"
					zh_cn = "zh_cn"
				}
				options {
					hidden = true
				}
			}
			additional_policy_scopes = ["%s", "%s.some"]
			display_name {
				default = "%s"
				en = "en"
				de = "de"
				es = "es"
				fr = "fr"
				it = "it"
				ja = "ja"
				ko = "ko"
				pt_br = "pt_br"
				zh_tw = "zh_tw"
				zh_cn = "zh_cn"
						}
			parent_ids = ["05ca8653-de25-49fa-a14d-aaa5d373bc22"]	
			supported_attributes {
				key = "testString"
				options {
					operators = [ "stringEquals" ]
					hidden = false
					policy_types = [ "access" ]
					is_empty_value_supported = true
					is_string_exists_false_value_supported = true
				}
				display_name {
					default = "%s"
					en = "en"
					de = "de"
					es = "es"
					fr = "fr"
					it = "it"
					ja = "ja"
					ko = "ko"
					pt_br = "pt_br"
					zh_tw = "zh_tw"
					zh_cn = "zh_cn"
				}
				description {
					default = "default"
					en = "en"
					de = "de"
					es = "es"
					fr = "fr"
					it = "it"
					ja = "ja"
					ko = "ko"
					pt_br = "pt_br"
					zh_tw = "zh_tw"
					zh_cn = "zh_cn"
				}
				ui {
					input_type = "selector"
					input_details {
						type = "gst"
						values {
							value = "testString"
							display_name {
								default = "%s"
								en = "testString"
								de = "testString"
								es = "testString"
								fr = "testString"
								it = "testString"
								ja = "testString"
								ko = "testString"
								pt_br = "testString"
								zh_tw = "testString"
								zh_cn = "testString"
							}
						}
						gst {
							query = "query"
							value_property_name = "teststring"
							input_option_label = "{name} - {instance_id}"
						}
					}
				}
			}
			supported_attributes {
        		key = "some-attribute"
        		display_name {
            		default = "some-attribute"
        		}
        		description {
            		default = "some-attribute"
        		}
       			 ui {
            		input_type = "string"
       			}
    		}
			supported_authorization_subjects {
				attributes {
					service_name = "%s"
					resource_type = "testString"
				}
				roles = [ "%s" ]
			}
			supported_roles {
				id = "%s"
				description {
					default = "desc"
				}
				display_name {
					default = "%s"
				}
				options {
					access_policy = true
					policy_type = [ "access" ]
				}
			}
			supported_roles {
				id = "%s"
				description {
					default = "default"
				}
				display_name {
					default = "%s"
				}
				options {
					access_policy = true
					policy_type = [ "access" ]
				}
			}
			supported_network {
				environment_attributes {
					key = "networkType"
					values = [ "%s" ]
					options {
						hidden = true
					}
				}
			}
			supported_anonymous_accesses {
				attributes {
					account_id = "account_id"
					service_name = "%s"
					additional_properties = { "testString" = "%s" }
				}
				roles = [ "%s" ]
			}
			supported_anonymous_accesses {
				attributes {
					account_id = "account_id"
					service_name = "%s"
					additional_properties = { "testString" = "something" }
				}
				roles = [ "%s" ]
			}
		}
	`, productID, env, name, enabled, serviceType, iamRegistrationRole, iamRegistrationRole2, iamRegistrationRole, iamRegistrationRole, name, name, name, supportedAttributeDisplayName, supportedAttributeInputDetailsDisplayName, supportedAuthorizationSubjectsService, iamRegistrationRole, iamRegistrationRole, roleDisplayName, iamRegistrationRole2, roleDisplayName2, environmentAttributesValues, iamRegistrationID, supportedAnonymousAccessesAdditionalPropValue, iamRegistrationRole, iamRegistrationID, iamRegistrationRole)
}

func testAccCheckIbmOnboardingIamRegistrationExists(n string, obj partnercentersellv1.IamServiceRegistration) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
		if err != nil {
			return err
		}

		getIamRegistrationOptions := &partnercentersellv1.GetIamRegistrationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIamRegistrationOptions.SetProductID(parts[0])
		getIamRegistrationOptions.SetProgrammaticName(parts[1])

		iamServiceRegistration, _, err := partnerCenterSellClient.GetIamRegistration(getIamRegistrationOptions)
		if err != nil {
			return err
		}

		obj = *iamServiceRegistration
		return nil
	}
}

func testAccCheckIbmOnboardingIamRegistrationDestroy(s *terraform.State) error {
	partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_onboarding_iam_registration" {
			continue
		}

		getIamRegistrationOptions := &partnercentersellv1.GetIamRegistrationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIamRegistrationOptions.SetProductID(parts[0])
		getIamRegistrationOptions.SetProgrammaticName(parts[1])

		// Try to find the key
		_, response, err := partnerCenterSellClient.GetIamRegistration(getIamRegistrationOptions)

		if err == nil {
			return fmt.Errorf("onboarding_iam_registration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for onboarding_iam_registration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationActionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		iamServiceRegistrationDescriptionObjectModel := make(map[string]interface{})
		iamServiceRegistrationDescriptionObjectModel["default"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["en"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["de"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["es"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["fr"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["it"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["ja"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["ko"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["pt_br"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["zh_tw"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["zh_cn"] = "testString"

		iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
		iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

		iamServiceRegistrationActionOptionsModel := make(map[string]interface{})
		iamServiceRegistrationActionOptionsModel["hidden"] = true

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["roles"] = []string{"testString"}
		model["description"] = []map[string]interface{}{iamServiceRegistrationDescriptionObjectModel}
		model["display_name"] = []map[string]interface{}{iamServiceRegistrationDisplayNameObjectModel}
		model["options"] = []map[string]interface{}{iamServiceRegistrationActionOptionsModel}

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
	iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("testString")

	iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
	iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

	iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
	iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

	model := new(partnercentersellv1.IamServiceRegistrationAction)
	model.ID = core.StringPtr("testString")
	model.Roles = []string{"testString"}
	model.Description = iamServiceRegistrationDescriptionObjectModel
	model.DisplayName = iamServiceRegistrationDisplayNameObjectModel
	model.Options = iamServiceRegistrationActionOptionsModel

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationActionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationDescriptionObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["default"] = "testString"
		model["en"] = "testString"
		model["de"] = "testString"
		model["es"] = "testString"
		model["fr"] = "testString"
		model["it"] = "testString"
		model["ja"] = "testString"
		model["ko"] = "testString"
		model["pt_br"] = "testString"
		model["zh_tw"] = "testString"
		model["zh_cn"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
	model.Default = core.StringPtr("testString")
	model.En = core.StringPtr("testString")
	model.De = core.StringPtr("testString")
	model.Es = core.StringPtr("testString")
	model.Fr = core.StringPtr("testString")
	model.It = core.StringPtr("testString")
	model.Ja = core.StringPtr("testString")
	model.Ko = core.StringPtr("testString")
	model.PtBr = core.StringPtr("testString")
	model.ZhTw = core.StringPtr("testString")
	model.ZhCn = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationDescriptionObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationDisplayNameObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["default"] = "testString"
		model["en"] = "testString"
		model["de"] = "testString"
		model["es"] = "testString"
		model["fr"] = "testString"
		model["it"] = "testString"
		model["ja"] = "testString"
		model["ko"] = "testString"
		model["pt_br"] = "testString"
		model["zh_tw"] = "testString"
		model["zh_cn"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
	model.Default = core.StringPtr("testString")
	model.En = core.StringPtr("testString")
	model.De = core.StringPtr("testString")
	model.Es = core.StringPtr("testString")
	model.Fr = core.StringPtr("testString")
	model.It = core.StringPtr("testString")
	model.Ja = core.StringPtr("testString")
	model.Ko = core.StringPtr("testString")
	model.PtBr = core.StringPtr("testString")
	model.ZhTw = core.StringPtr("testString")
	model.ZhCn = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationDisplayNameObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationActionOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hidden"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.IamServiceRegistrationActionOptions)
	model.Hidden = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationActionOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationResourceHierarchyAttributeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationResourceHierarchyAttributeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedAnonymousAccessToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		iamServiceRegistrationSupportedAnonymousAccessAttributesModel := make(map[string]interface{})
		iamServiceRegistrationSupportedAnonymousAccessAttributesModel["account_id"] = "testString"
		iamServiceRegistrationSupportedAnonymousAccessAttributesModel["service_name"] = "testString"
		iamServiceRegistrationSupportedAnonymousAccessAttributesModel["additional_properties"] = map[string]interface{}{"key1": "testString"}

		model := make(map[string]interface{})
		model["attributes"] = []map[string]interface{}{iamServiceRegistrationSupportedAnonymousAccessAttributesModel}
		model["roles"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
	iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
	iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
	iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}

	model := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
	model.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
	model.Roles = []string{"testString"}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedAnonymousAccessToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedAnonymousAccessAttributesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["account_id"] = "testString"
		model["service_name"] = "testString"
		model["additional_properties"] = map[string]interface{}{"key1": "testString"}

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
	model.AccountID = core.StringPtr("testString")
	model.ServiceName = core.StringPtr("testString")
	model.AdditionalProperties = map[string]string{"key1": "testString"}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedAnonymousAccessAttributesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedAttributeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportedAttributesOptionsResourceHierarchyKeyModel := make(map[string]interface{})
		supportedAttributesOptionsResourceHierarchyKeyModel["key"] = "testString"
		supportedAttributesOptionsResourceHierarchyKeyModel["value"] = "testString"

		supportedAttributesOptionsResourceHierarchyValueModel := make(map[string]interface{})
		supportedAttributesOptionsResourceHierarchyValueModel["key"] = "testString"

		supportedAttributesOptionsResourceHierarchyModel := make(map[string]interface{})
		supportedAttributesOptionsResourceHierarchyModel["key"] = []map[string]interface{}{supportedAttributesOptionsResourceHierarchyKeyModel}
		supportedAttributesOptionsResourceHierarchyModel["value"] = []map[string]interface{}{supportedAttributesOptionsResourceHierarchyValueModel}

		supportedAttributesOptionsModel := make(map[string]interface{})
		supportedAttributesOptionsModel["operators"] = []string{"stringEquals"}
		supportedAttributesOptionsModel["hidden"] = true
		supportedAttributesOptionsModel["supported_patterns"] = []string{"testString"}
		supportedAttributesOptionsModel["policy_types"] = []string{"access"}
		supportedAttributesOptionsModel["is_empty_value_supported"] = true
		supportedAttributesOptionsModel["is_string_exists_false_value_supported"] = true
		supportedAttributesOptionsModel["key"] = "testString"
		supportedAttributesOptionsModel["resource_hierarchy"] = []map[string]interface{}{supportedAttributesOptionsResourceHierarchyModel}

		iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
		iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

		iamServiceRegistrationDescriptionObjectModel := make(map[string]interface{})
		iamServiceRegistrationDescriptionObjectModel["default"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["en"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["de"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["es"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["fr"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["it"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["ja"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["ko"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["pt_br"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["zh_tw"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["zh_cn"] = "testString"

		supportedAttributeUiInputValueModel := make(map[string]interface{})
		supportedAttributeUiInputValueModel["value"] = "testString"
		supportedAttributeUiInputValueModel["display_name"] = []map[string]interface{}{iamServiceRegistrationDisplayNameObjectModel}

		supportedAttributeUiInputGstModel := make(map[string]interface{})
		supportedAttributeUiInputGstModel["query"] = "testString"
		supportedAttributeUiInputGstModel["value_property_name"] = "testString"
		supportedAttributeUiInputGstModel["label_property_name"] = "testString"
		supportedAttributeUiInputGstModel["input_option_label"] = "testString"

		supportedAttributeUiInputUrlModel := make(map[string]interface{})
		supportedAttributeUiInputUrlModel["url_endpoint"] = "testString"
		supportedAttributeUiInputUrlModel["input_option_label"] = "testString"

		supportedAttributeUiInputDetailsModel := make(map[string]interface{})
		supportedAttributeUiInputDetailsModel["type"] = "testString"
		supportedAttributeUiInputDetailsModel["values"] = []map[string]interface{}{supportedAttributeUiInputValueModel}
		supportedAttributeUiInputDetailsModel["gst"] = []map[string]interface{}{supportedAttributeUiInputGstModel}
		supportedAttributeUiInputDetailsModel["url"] = []map[string]interface{}{supportedAttributeUiInputUrlModel}

		supportedAttributeUiModel := make(map[string]interface{})
		supportedAttributeUiModel["input_type"] = "testString"
		supportedAttributeUiModel["input_details"] = []map[string]interface{}{supportedAttributeUiInputDetailsModel}

		model := make(map[string]interface{})
		model["key"] = "testString"
		model["options"] = []map[string]interface{}{supportedAttributesOptionsModel}
		model["display_name"] = []map[string]interface{}{iamServiceRegistrationDisplayNameObjectModel}
		model["description"] = []map[string]interface{}{iamServiceRegistrationDescriptionObjectModel}
		model["ui"] = []map[string]interface{}{supportedAttributeUiModel}

		assert.Equal(t, result, model)
	}

	supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
	supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
	supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

	supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
	supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

	supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
	supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
	supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

	supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
	supportedAttributesOptionsModel.Operators = []string{"stringEquals"}
	supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
	supportedAttributesOptionsModel.SupportedPatterns = []string{"testString"}
	supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
	supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
	supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
	supportedAttributesOptionsModel.Key = core.StringPtr("testString")
	supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

	iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
	iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

	iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
	iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("testString")

	supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
	supportedAttributeUiInputValueModel.Value = core.StringPtr("testString")
	supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

	supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
	supportedAttributeUiInputGstModel.Query = core.StringPtr("testString")
	supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("testString")
	supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
	supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("testString")

	supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
	supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
	supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

	supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
	supportedAttributeUiInputDetailsModel.Type = core.StringPtr("testString")
	supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
	supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
	supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

	supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
	supportedAttributeUiModel.InputType = core.StringPtr("testString")
	supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

	model := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
	model.Key = core.StringPtr("testString")
	model.Options = supportedAttributesOptionsModel
	model.DisplayName = iamServiceRegistrationDisplayNameObjectModel
	model.Description = iamServiceRegistrationDescriptionObjectModel
	model.Ui = supportedAttributeUiModel

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedAttributeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedAttributesOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportedAttributesOptionsResourceHierarchyKeyModel := make(map[string]interface{})
		supportedAttributesOptionsResourceHierarchyKeyModel["key"] = "testString"
		supportedAttributesOptionsResourceHierarchyKeyModel["value"] = "testString"

		supportedAttributesOptionsResourceHierarchyValueModel := make(map[string]interface{})
		supportedAttributesOptionsResourceHierarchyValueModel["key"] = "testString"

		supportedAttributesOptionsResourceHierarchyModel := make(map[string]interface{})
		supportedAttributesOptionsResourceHierarchyModel["key"] = []map[string]interface{}{supportedAttributesOptionsResourceHierarchyKeyModel}
		supportedAttributesOptionsResourceHierarchyModel["value"] = []map[string]interface{}{supportedAttributesOptionsResourceHierarchyValueModel}

		model := make(map[string]interface{})
		model["operators"] = []string{"stringEquals"}
		model["hidden"] = true
		model["supported_patterns"] = []string{"testString"}
		model["policy_types"] = []string{"access"}
		model["is_empty_value_supported"] = true
		model["is_string_exists_false_value_supported"] = true
		model["key"] = "testString"
		model["resource_hierarchy"] = []map[string]interface{}{supportedAttributesOptionsResourceHierarchyModel}

		assert.Equal(t, result, model)
	}

	supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
	supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
	supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

	supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
	supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

	supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
	supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
	supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

	model := new(partnercentersellv1.SupportedAttributesOptions)
	model.Operators = []string{"stringEquals"}
	model.Hidden = core.BoolPtr(true)
	model.SupportedPatterns = []string{"testString"}
	model.PolicyTypes = []string{"access"}
	model.IsEmptyValueSupported = core.BoolPtr(true)
	model.IsStringExistsFalseValueSupported = core.BoolPtr(true)
	model.Key = core.StringPtr("testString")
	model.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedAttributesOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedAttributesOptionsResourceHierarchyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportedAttributesOptionsResourceHierarchyKeyModel := make(map[string]interface{})
		supportedAttributesOptionsResourceHierarchyKeyModel["key"] = "testString"
		supportedAttributesOptionsResourceHierarchyKeyModel["value"] = "testString"

		supportedAttributesOptionsResourceHierarchyValueModel := make(map[string]interface{})
		supportedAttributesOptionsResourceHierarchyValueModel["key"] = "testString"

		model := make(map[string]interface{})
		model["key"] = []map[string]interface{}{supportedAttributesOptionsResourceHierarchyKeyModel}
		model["value"] = []map[string]interface{}{supportedAttributesOptionsResourceHierarchyValueModel}

		assert.Equal(t, result, model)
	}

	supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
	supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
	supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

	supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
	supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

	model := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
	model.Key = supportedAttributesOptionsResourceHierarchyKeyModel
	model.Value = supportedAttributesOptionsResourceHierarchyValueModel

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedAttributesOptionsResourceHierarchyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedAttributesOptionsResourceHierarchyKeyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedAttributesOptionsResourceHierarchyKeyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedAttributesOptionsResourceHierarchyValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
	model.Key = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedAttributesOptionsResourceHierarchyValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedAttributeUiToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
		iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

		supportedAttributeUiInputValueModel := make(map[string]interface{})
		supportedAttributeUiInputValueModel["value"] = "testString"
		supportedAttributeUiInputValueModel["display_name"] = []map[string]interface{}{iamServiceRegistrationDisplayNameObjectModel}

		supportedAttributeUiInputGstModel := make(map[string]interface{})
		supportedAttributeUiInputGstModel["query"] = "testString"
		supportedAttributeUiInputGstModel["value_property_name"] = "testString"
		supportedAttributeUiInputGstModel["label_property_name"] = "testString"
		supportedAttributeUiInputGstModel["input_option_label"] = "testString"

		supportedAttributeUiInputUrlModel := make(map[string]interface{})
		supportedAttributeUiInputUrlModel["url_endpoint"] = "testString"
		supportedAttributeUiInputUrlModel["input_option_label"] = "testString"

		supportedAttributeUiInputDetailsModel := make(map[string]interface{})
		supportedAttributeUiInputDetailsModel["type"] = "testString"
		supportedAttributeUiInputDetailsModel["values"] = []map[string]interface{}{supportedAttributeUiInputValueModel}
		supportedAttributeUiInputDetailsModel["gst"] = []map[string]interface{}{supportedAttributeUiInputGstModel}
		supportedAttributeUiInputDetailsModel["url"] = []map[string]interface{}{supportedAttributeUiInputUrlModel}

		model := make(map[string]interface{})
		model["input_type"] = "testString"
		model["input_details"] = []map[string]interface{}{supportedAttributeUiInputDetailsModel}

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
	iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

	supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
	supportedAttributeUiInputValueModel.Value = core.StringPtr("testString")
	supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

	supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
	supportedAttributeUiInputGstModel.Query = core.StringPtr("testString")
	supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("testString")
	supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
	supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("testString")

	supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
	supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
	supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

	supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
	supportedAttributeUiInputDetailsModel.Type = core.StringPtr("testString")
	supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
	supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
	supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

	model := new(partnercentersellv1.SupportedAttributeUi)
	model.InputType = core.StringPtr("testString")
	model.InputDetails = supportedAttributeUiInputDetailsModel

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedAttributeUiToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedAttributeUiInputDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
		iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

		supportedAttributeUiInputValueModel := make(map[string]interface{})
		supportedAttributeUiInputValueModel["value"] = "testString"
		supportedAttributeUiInputValueModel["display_name"] = []map[string]interface{}{iamServiceRegistrationDisplayNameObjectModel}

		supportedAttributeUiInputGstModel := make(map[string]interface{})
		supportedAttributeUiInputGstModel["query"] = "testString"
		supportedAttributeUiInputGstModel["value_property_name"] = "testString"
		supportedAttributeUiInputGstModel["label_property_name"] = "testString"
		supportedAttributeUiInputGstModel["input_option_label"] = "testString"

		supportedAttributeUiInputUrlModel := make(map[string]interface{})
		supportedAttributeUiInputUrlModel["url_endpoint"] = "testString"
		supportedAttributeUiInputUrlModel["input_option_label"] = "testString"

		model := make(map[string]interface{})
		model["type"] = "testString"
		model["values"] = []map[string]interface{}{supportedAttributeUiInputValueModel}
		model["gst"] = []map[string]interface{}{supportedAttributeUiInputGstModel}
		model["url"] = []map[string]interface{}{supportedAttributeUiInputUrlModel}

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
	iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

	supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
	supportedAttributeUiInputValueModel.Value = core.StringPtr("testString")
	supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

	supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
	supportedAttributeUiInputGstModel.Query = core.StringPtr("testString")
	supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("testString")
	supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
	supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("testString")

	supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
	supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
	supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

	model := new(partnercentersellv1.SupportedAttributeUiInputDetails)
	model.Type = core.StringPtr("testString")
	model.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
	model.Gst = supportedAttributeUiInputGstModel
	model.URL = supportedAttributeUiInputUrlModel

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedAttributeUiInputDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedAttributeUiInputValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
		iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

		model := make(map[string]interface{})
		model["value"] = "testString"
		model["display_name"] = []map[string]interface{}{iamServiceRegistrationDisplayNameObjectModel}

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
	iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

	model := new(partnercentersellv1.SupportedAttributeUiInputValue)
	model.Value = core.StringPtr("testString")
	model.DisplayName = iamServiceRegistrationDisplayNameObjectModel

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedAttributeUiInputValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedAttributeUiInputGstToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["query"] = "testString"
		model["value_property_name"] = "testString"
		model["label_property_name"] = "testString"
		model["input_option_label"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportedAttributeUiInputGst)
	model.Query = core.StringPtr("testString")
	model.ValuePropertyName = core.StringPtr("testString")
	model.LabelPropertyName = core.StringPtr("testString")
	model.InputOptionLabel = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedAttributeUiInputGstToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedAttributeUiInputURLToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["url_endpoint"] = "testString"
		model["input_option_label"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportedAttributeUiInputURL)
	model.UrlEndpoint = core.StringPtr("testString")
	model.InputOptionLabel = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedAttributeUiInputURLToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedAuthorizationSubjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportAuthorizationSubjectAttributeModel := make(map[string]interface{})
		supportAuthorizationSubjectAttributeModel["service_name"] = "testString"
		supportAuthorizationSubjectAttributeModel["resource_type"] = "testString"

		model := make(map[string]interface{})
		model["attributes"] = []map[string]interface{}{supportAuthorizationSubjectAttributeModel}
		model["roles"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
	supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
	supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

	model := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
	model.Attributes = supportAuthorizationSubjectAttributeModel
	model.Roles = []string{"testString"}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedAuthorizationSubjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportAuthorizationSubjectAttributeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["service_name"] = "testString"
		model["resource_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
	model.ServiceName = core.StringPtr("testString")
	model.ResourceType = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportAuthorizationSubjectAttributeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedRoleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		iamServiceRegistrationDescriptionObjectModel := make(map[string]interface{})
		iamServiceRegistrationDescriptionObjectModel["default"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["en"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["de"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["es"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["fr"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["it"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["ja"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["ko"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["pt_br"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["zh_tw"] = "testString"
		iamServiceRegistrationDescriptionObjectModel["zh_cn"] = "testString"

		iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
		iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
		iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

		supportedRoleOptionsModel := make(map[string]interface{})
		supportedRoleOptionsModel["access_policy"] = true
		supportedRoleOptionsModel["policy_type"] = []string{"access"}
		supportedRoleOptionsModel["account_type"] = "enterprise"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["description"] = []map[string]interface{}{iamServiceRegistrationDescriptionObjectModel}
		model["display_name"] = []map[string]interface{}{iamServiceRegistrationDisplayNameObjectModel}
		model["options"] = []map[string]interface{}{supportedRoleOptionsModel}

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
	iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("testString")
	iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("testString")

	iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
	iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
	iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

	supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
	supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
	supportedRoleOptionsModel.PolicyType = []string{"access"}
	supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

	model := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
	model.ID = core.StringPtr("testString")
	model.Description = iamServiceRegistrationDescriptionObjectModel
	model.DisplayName = iamServiceRegistrationDisplayNameObjectModel
	model.Options = supportedRoleOptionsModel

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedRoleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationSupportedRoleOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["access_policy"] = true
		model["policy_type"] = []string{"access"}
		model["account_type"] = "enterprise"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportedRoleOptions)
	model.AccessPolicy = core.BoolPtr(true)
	model.PolicyType = []string{"access"}
	model.AccountType = core.StringPtr("enterprise")

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationSupportedRoleOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedNetworkToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		environmentAttributeOptionsModel := make(map[string]interface{})
		environmentAttributeOptionsModel["hidden"] = true

		environmentAttributeModel := make(map[string]interface{})
		environmentAttributeModel["key"] = "testString"
		environmentAttributeModel["values"] = []string{"testString"}
		environmentAttributeModel["options"] = []map[string]interface{}{environmentAttributeOptionsModel}

		model := make(map[string]interface{})
		model["environment_attributes"] = []map[string]interface{}{environmentAttributeModel}

		assert.Equal(t, result, model)
	}

	environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
	environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

	environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
	environmentAttributeModel.Key = core.StringPtr("testString")
	environmentAttributeModel.Values = []string{"testString"}
	environmentAttributeModel.Options = environmentAttributeOptionsModel

	model := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
	model.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationIamServiceRegistrationSupportedNetworkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationEnvironmentAttributeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		environmentAttributeOptionsModel := make(map[string]interface{})
		environmentAttributeOptionsModel["hidden"] = true

		model := make(map[string]interface{})
		model["key"] = "testString"
		model["values"] = []string{"testString"}
		model["options"] = []map[string]interface{}{environmentAttributeOptionsModel}

		assert.Equal(t, result, model)
	}

	environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
	environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

	model := new(partnercentersellv1.EnvironmentAttribute)
	model.Key = core.StringPtr("testString")
	model.Values = []string{"testString"}
	model.Options = environmentAttributeOptionsModel

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationEnvironmentAttributeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationEnvironmentAttributeOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hidden"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.EnvironmentAttributeOptions)
	model.Hidden = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationEnvironmentAttributeOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationAction(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationAction) {
		iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
		iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("testString")

		iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
		iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

		iamServiceRegistrationActionOptionsModel := new(partnercentersellv1.IamServiceRegistrationActionOptions)
		iamServiceRegistrationActionOptionsModel.Hidden = core.BoolPtr(true)

		model := new(partnercentersellv1.IamServiceRegistrationAction)
		model.ID = core.StringPtr("testString")
		model.Roles = []string{"testString"}
		model.Description = iamServiceRegistrationDescriptionObjectModel
		model.DisplayName = iamServiceRegistrationDisplayNameObjectModel
		model.Options = iamServiceRegistrationActionOptionsModel

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDescriptionObjectModel := make(map[string]interface{})
	iamServiceRegistrationDescriptionObjectModel["default"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["en"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["de"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["es"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["fr"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["it"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["ja"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["ko"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["pt_br"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["zh_tw"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["zh_cn"] = "testString"

	iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
	iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

	iamServiceRegistrationActionOptionsModel := make(map[string]interface{})
	iamServiceRegistrationActionOptionsModel["hidden"] = true

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["roles"] = []interface{}{"testString"}
	model["description"] = []interface{}{iamServiceRegistrationDescriptionObjectModel}
	model["display_name"] = []interface{}{iamServiceRegistrationDisplayNameObjectModel}
	model["options"] = []interface{}{iamServiceRegistrationActionOptionsModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationAction(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationDescriptionObject(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationDescriptionObject) {
		model := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
		model.Default = core.StringPtr("testString")
		model.En = core.StringPtr("testString")
		model.De = core.StringPtr("testString")
		model.Es = core.StringPtr("testString")
		model.Fr = core.StringPtr("testString")
		model.It = core.StringPtr("testString")
		model.Ja = core.StringPtr("testString")
		model.Ko = core.StringPtr("testString")
		model.PtBr = core.StringPtr("testString")
		model.ZhTw = core.StringPtr("testString")
		model.ZhCn = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["default"] = "testString"
	model["en"] = "testString"
	model["de"] = "testString"
	model["es"] = "testString"
	model["fr"] = "testString"
	model["it"] = "testString"
	model["ja"] = "testString"
	model["ko"] = "testString"
	model["pt_br"] = "testString"
	model["zh_tw"] = "testString"
	model["zh_cn"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationDescriptionObject(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationDisplayNameObject(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationDisplayNameObject) {
		model := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
		model.Default = core.StringPtr("testString")
		model.En = core.StringPtr("testString")
		model.De = core.StringPtr("testString")
		model.Es = core.StringPtr("testString")
		model.Fr = core.StringPtr("testString")
		model.It = core.StringPtr("testString")
		model.Ja = core.StringPtr("testString")
		model.Ko = core.StringPtr("testString")
		model.PtBr = core.StringPtr("testString")
		model.ZhTw = core.StringPtr("testString")
		model.ZhCn = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["default"] = "testString"
	model["en"] = "testString"
	model["de"] = "testString"
	model["es"] = "testString"
	model["fr"] = "testString"
	model["it"] = "testString"
	model["ja"] = "testString"
	model["ko"] = "testString"
	model["pt_br"] = "testString"
	model["zh_tw"] = "testString"
	model["zh_cn"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationDisplayNameObject(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationActionOptions(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationActionOptions) {
		model := new(partnercentersellv1.IamServiceRegistrationActionOptions)
		model.Hidden = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["hidden"] = true

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationActionOptions(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationResourceHierarchyAttribute(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute) {
		model := new(partnercentersellv1.IamServiceRegistrationResourceHierarchyAttribute)
		model.Key = core.StringPtr("testString")
		model.Value = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["key"] = "testString"
	model["value"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationResourceHierarchyAttribute(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedAnonymousAccess(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess) {
		iamServiceRegistrationSupportedAnonymousAccessAttributesModel := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
		iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AccountID = core.StringPtr("testString")
		iamServiceRegistrationSupportedAnonymousAccessAttributesModel.ServiceName = core.StringPtr("testString")
		iamServiceRegistrationSupportedAnonymousAccessAttributesModel.AdditionalProperties = map[string]string{"key1": "testString"}

		model := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccess)
		model.Attributes = iamServiceRegistrationSupportedAnonymousAccessAttributesModel
		model.Roles = []string{"testString"}

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationSupportedAnonymousAccessAttributesModel := make(map[string]interface{})
	iamServiceRegistrationSupportedAnonymousAccessAttributesModel["account_id"] = "testString"
	iamServiceRegistrationSupportedAnonymousAccessAttributesModel["service_name"] = "testString"
	iamServiceRegistrationSupportedAnonymousAccessAttributesModel["additional_properties"] = map[string]interface{}{"key1": "testString"}

	model := make(map[string]interface{})
	model["attributes"] = []interface{}{iamServiceRegistrationSupportedAnonymousAccessAttributesModel}
	model["roles"] = []interface{}{"testString"}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedAnonymousAccess(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedAnonymousAccessAttributes(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes) {
		model := new(partnercentersellv1.IamServiceRegistrationSupportedAnonymousAccessAttributes)
		model.AccountID = core.StringPtr("testString")
		model.ServiceName = core.StringPtr("testString")
		model.AdditionalProperties = map[string]string{"key1": "testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["account_id"] = "testString"
	model["service_name"] = "testString"
	model["additional_properties"] = map[string]interface{}{"key1": "testString"}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedAnonymousAccessAttributes(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedAttribute(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationSupportedAttribute) {
		supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
		supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
		supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

		supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
		supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

		supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
		supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
		supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

		supportedAttributesOptionsModel := new(partnercentersellv1.SupportedAttributesOptions)
		supportedAttributesOptionsModel.Operators = []string{"stringEquals"}
		supportedAttributesOptionsModel.Hidden = core.BoolPtr(true)
		supportedAttributesOptionsModel.SupportedPatterns = []string{"testString"}
		supportedAttributesOptionsModel.PolicyTypes = []string{"access"}
		supportedAttributesOptionsModel.IsEmptyValueSupported = core.BoolPtr(true)
		supportedAttributesOptionsModel.IsStringExistsFalseValueSupported = core.BoolPtr(true)
		supportedAttributesOptionsModel.Key = core.StringPtr("testString")
		supportedAttributesOptionsModel.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

		iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
		iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

		iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
		iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("testString")

		supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
		supportedAttributeUiInputValueModel.Value = core.StringPtr("testString")
		supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

		supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
		supportedAttributeUiInputGstModel.Query = core.StringPtr("testString")
		supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("testString")
		supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
		supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("testString")

		supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
		supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
		supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

		supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
		supportedAttributeUiInputDetailsModel.Type = core.StringPtr("testString")
		supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
		supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
		supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

		supportedAttributeUiModel := new(partnercentersellv1.SupportedAttributeUi)
		supportedAttributeUiModel.InputType = core.StringPtr("testString")
		supportedAttributeUiModel.InputDetails = supportedAttributeUiInputDetailsModel

		model := new(partnercentersellv1.IamServiceRegistrationSupportedAttribute)
		model.Key = core.StringPtr("testString")
		model.Options = supportedAttributesOptionsModel
		model.DisplayName = iamServiceRegistrationDisplayNameObjectModel
		model.Description = iamServiceRegistrationDescriptionObjectModel
		model.Ui = supportedAttributeUiModel

		assert.Equal(t, result, model)
	}

	supportedAttributesOptionsResourceHierarchyKeyModel := make(map[string]interface{})
	supportedAttributesOptionsResourceHierarchyKeyModel["key"] = "testString"
	supportedAttributesOptionsResourceHierarchyKeyModel["value"] = "testString"

	supportedAttributesOptionsResourceHierarchyValueModel := make(map[string]interface{})
	supportedAttributesOptionsResourceHierarchyValueModel["key"] = "testString"

	supportedAttributesOptionsResourceHierarchyModel := make(map[string]interface{})
	supportedAttributesOptionsResourceHierarchyModel["key"] = []interface{}{supportedAttributesOptionsResourceHierarchyKeyModel}
	supportedAttributesOptionsResourceHierarchyModel["value"] = []interface{}{supportedAttributesOptionsResourceHierarchyValueModel}

	supportedAttributesOptionsModel := make(map[string]interface{})
	supportedAttributesOptionsModel["operators"] = []interface{}{"stringEquals"}
	supportedAttributesOptionsModel["hidden"] = true
	supportedAttributesOptionsModel["supported_patterns"] = []interface{}{"testString"}
	supportedAttributesOptionsModel["policy_types"] = []interface{}{"access"}
	supportedAttributesOptionsModel["is_empty_value_supported"] = true
	supportedAttributesOptionsModel["is_string_exists_false_value_supported"] = true
	supportedAttributesOptionsModel["key"] = "testString"
	supportedAttributesOptionsModel["resource_hierarchy"] = []interface{}{supportedAttributesOptionsResourceHierarchyModel}

	iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
	iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

	iamServiceRegistrationDescriptionObjectModel := make(map[string]interface{})
	iamServiceRegistrationDescriptionObjectModel["default"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["en"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["de"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["es"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["fr"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["it"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["ja"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["ko"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["pt_br"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["zh_tw"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["zh_cn"] = "testString"

	supportedAttributeUiInputValueModel := make(map[string]interface{})
	supportedAttributeUiInputValueModel["value"] = "testString"
	supportedAttributeUiInputValueModel["display_name"] = []interface{}{iamServiceRegistrationDisplayNameObjectModel}

	supportedAttributeUiInputGstModel := make(map[string]interface{})
	supportedAttributeUiInputGstModel["query"] = "testString"
	supportedAttributeUiInputGstModel["value_property_name"] = "testString"
	supportedAttributeUiInputGstModel["label_property_name"] = "testString"
	supportedAttributeUiInputGstModel["input_option_label"] = "testString"

	supportedAttributeUiInputUrlModel := make(map[string]interface{})
	supportedAttributeUiInputUrlModel["url_endpoint"] = "testString"
	supportedAttributeUiInputUrlModel["input_option_label"] = "testString"

	supportedAttributeUiInputDetailsModel := make(map[string]interface{})
	supportedAttributeUiInputDetailsModel["type"] = "testString"
	supportedAttributeUiInputDetailsModel["values"] = []interface{}{supportedAttributeUiInputValueModel}
	supportedAttributeUiInputDetailsModel["gst"] = []interface{}{supportedAttributeUiInputGstModel}
	supportedAttributeUiInputDetailsModel["url"] = []interface{}{supportedAttributeUiInputUrlModel}

	supportedAttributeUiModel := make(map[string]interface{})
	supportedAttributeUiModel["input_type"] = "testString"
	supportedAttributeUiModel["input_details"] = []interface{}{supportedAttributeUiInputDetailsModel}

	model := make(map[string]interface{})
	model["key"] = "testString"
	model["options"] = []interface{}{supportedAttributesOptionsModel}
	model["display_name"] = []interface{}{iamServiceRegistrationDisplayNameObjectModel}
	model["description"] = []interface{}{iamServiceRegistrationDescriptionObjectModel}
	model["ui"] = []interface{}{supportedAttributeUiModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedAttribute(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedAttributesOptions(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedAttributesOptions) {
		supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
		supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
		supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

		supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
		supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

		supportedAttributesOptionsResourceHierarchyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
		supportedAttributesOptionsResourceHierarchyModel.Key = supportedAttributesOptionsResourceHierarchyKeyModel
		supportedAttributesOptionsResourceHierarchyModel.Value = supportedAttributesOptionsResourceHierarchyValueModel

		model := new(partnercentersellv1.SupportedAttributesOptions)
		model.Operators = []string{"stringEquals"}
		model.Hidden = core.BoolPtr(true)
		model.SupportedPatterns = []string{"testString"}
		model.PolicyTypes = []string{"access"}
		model.IsEmptyValueSupported = core.BoolPtr(true)
		model.IsStringExistsFalseValueSupported = core.BoolPtr(true)
		model.Key = core.StringPtr("testString")
		model.ResourceHierarchy = supportedAttributesOptionsResourceHierarchyModel

		assert.Equal(t, result, model)
	}

	supportedAttributesOptionsResourceHierarchyKeyModel := make(map[string]interface{})
	supportedAttributesOptionsResourceHierarchyKeyModel["key"] = "testString"
	supportedAttributesOptionsResourceHierarchyKeyModel["value"] = "testString"

	supportedAttributesOptionsResourceHierarchyValueModel := make(map[string]interface{})
	supportedAttributesOptionsResourceHierarchyValueModel["key"] = "testString"

	supportedAttributesOptionsResourceHierarchyModel := make(map[string]interface{})
	supportedAttributesOptionsResourceHierarchyModel["key"] = []interface{}{supportedAttributesOptionsResourceHierarchyKeyModel}
	supportedAttributesOptionsResourceHierarchyModel["value"] = []interface{}{supportedAttributesOptionsResourceHierarchyValueModel}

	model := make(map[string]interface{})
	model["operators"] = []interface{}{"stringEquals"}
	model["hidden"] = true
	model["supported_patterns"] = []interface{}{"testString"}
	model["policy_types"] = []interface{}{"access"}
	model["is_empty_value_supported"] = true
	model["is_string_exists_false_value_supported"] = true
	model["key"] = "testString"
	model["resource_hierarchy"] = []interface{}{supportedAttributesOptionsResourceHierarchyModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedAttributesOptions(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedAttributesOptionsResourceHierarchy(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedAttributesOptionsResourceHierarchy) {
		supportedAttributesOptionsResourceHierarchyKeyModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
		supportedAttributesOptionsResourceHierarchyKeyModel.Key = core.StringPtr("testString")
		supportedAttributesOptionsResourceHierarchyKeyModel.Value = core.StringPtr("testString")

		supportedAttributesOptionsResourceHierarchyValueModel := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
		supportedAttributesOptionsResourceHierarchyValueModel.Key = core.StringPtr("testString")

		model := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchy)
		model.Key = supportedAttributesOptionsResourceHierarchyKeyModel
		model.Value = supportedAttributesOptionsResourceHierarchyValueModel

		assert.Equal(t, result, model)
	}

	supportedAttributesOptionsResourceHierarchyKeyModel := make(map[string]interface{})
	supportedAttributesOptionsResourceHierarchyKeyModel["key"] = "testString"
	supportedAttributesOptionsResourceHierarchyKeyModel["value"] = "testString"

	supportedAttributesOptionsResourceHierarchyValueModel := make(map[string]interface{})
	supportedAttributesOptionsResourceHierarchyValueModel["key"] = "testString"

	model := make(map[string]interface{})
	model["key"] = []interface{}{supportedAttributesOptionsResourceHierarchyKeyModel}
	model["value"] = []interface{}{supportedAttributesOptionsResourceHierarchyValueModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedAttributesOptionsResourceHierarchy(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedAttributesOptionsResourceHierarchyKey(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey) {
		model := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyKey)
		model.Key = core.StringPtr("testString")
		model.Value = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["key"] = "testString"
	model["value"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedAttributesOptionsResourceHierarchyKey(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedAttributesOptionsResourceHierarchyValue(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue) {
		model := new(partnercentersellv1.SupportedAttributesOptionsResourceHierarchyValue)
		model.Key = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["key"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedAttributesOptionsResourceHierarchyValue(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUi(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedAttributeUi) {
		iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
		iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

		supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
		supportedAttributeUiInputValueModel.Value = core.StringPtr("testString")
		supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

		supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
		supportedAttributeUiInputGstModel.Query = core.StringPtr("testString")
		supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("testString")
		supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
		supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("testString")

		supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
		supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
		supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

		supportedAttributeUiInputDetailsModel := new(partnercentersellv1.SupportedAttributeUiInputDetails)
		supportedAttributeUiInputDetailsModel.Type = core.StringPtr("testString")
		supportedAttributeUiInputDetailsModel.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
		supportedAttributeUiInputDetailsModel.Gst = supportedAttributeUiInputGstModel
		supportedAttributeUiInputDetailsModel.URL = supportedAttributeUiInputUrlModel

		model := new(partnercentersellv1.SupportedAttributeUi)
		model.InputType = core.StringPtr("testString")
		model.InputDetails = supportedAttributeUiInputDetailsModel

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
	iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

	supportedAttributeUiInputValueModel := make(map[string]interface{})
	supportedAttributeUiInputValueModel["value"] = "testString"
	supportedAttributeUiInputValueModel["display_name"] = []interface{}{iamServiceRegistrationDisplayNameObjectModel}

	supportedAttributeUiInputGstModel := make(map[string]interface{})
	supportedAttributeUiInputGstModel["query"] = "testString"
	supportedAttributeUiInputGstModel["value_property_name"] = "testString"
	supportedAttributeUiInputGstModel["label_property_name"] = "testString"
	supportedAttributeUiInputGstModel["input_option_label"] = "testString"

	supportedAttributeUiInputUrlModel := make(map[string]interface{})
	supportedAttributeUiInputUrlModel["url_endpoint"] = "testString"
	supportedAttributeUiInputUrlModel["input_option_label"] = "testString"

	supportedAttributeUiInputDetailsModel := make(map[string]interface{})
	supportedAttributeUiInputDetailsModel["type"] = "testString"
	supportedAttributeUiInputDetailsModel["values"] = []interface{}{supportedAttributeUiInputValueModel}
	supportedAttributeUiInputDetailsModel["gst"] = []interface{}{supportedAttributeUiInputGstModel}
	supportedAttributeUiInputDetailsModel["url"] = []interface{}{supportedAttributeUiInputUrlModel}

	model := make(map[string]interface{})
	model["input_type"] = "testString"
	model["input_details"] = []interface{}{supportedAttributeUiInputDetailsModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUi(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUiInputDetails(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedAttributeUiInputDetails) {
		iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
		iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

		supportedAttributeUiInputValueModel := new(partnercentersellv1.SupportedAttributeUiInputValue)
		supportedAttributeUiInputValueModel.Value = core.StringPtr("testString")
		supportedAttributeUiInputValueModel.DisplayName = iamServiceRegistrationDisplayNameObjectModel

		supportedAttributeUiInputGstModel := new(partnercentersellv1.SupportedAttributeUiInputGst)
		supportedAttributeUiInputGstModel.Query = core.StringPtr("testString")
		supportedAttributeUiInputGstModel.ValuePropertyName = core.StringPtr("testString")
		supportedAttributeUiInputGstModel.LabelPropertyName = core.StringPtr("testString")
		supportedAttributeUiInputGstModel.InputOptionLabel = core.StringPtr("testString")

		supportedAttributeUiInputUrlModel := new(partnercentersellv1.SupportedAttributeUiInputURL)
		supportedAttributeUiInputUrlModel.UrlEndpoint = core.StringPtr("testString")
		supportedAttributeUiInputUrlModel.InputOptionLabel = core.StringPtr("testString")

		model := new(partnercentersellv1.SupportedAttributeUiInputDetails)
		model.Type = core.StringPtr("testString")
		model.Values = []partnercentersellv1.SupportedAttributeUiInputValue{*supportedAttributeUiInputValueModel}
		model.Gst = supportedAttributeUiInputGstModel
		model.URL = supportedAttributeUiInputUrlModel

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
	iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

	supportedAttributeUiInputValueModel := make(map[string]interface{})
	supportedAttributeUiInputValueModel["value"] = "testString"
	supportedAttributeUiInputValueModel["display_name"] = []interface{}{iamServiceRegistrationDisplayNameObjectModel}

	supportedAttributeUiInputGstModel := make(map[string]interface{})
	supportedAttributeUiInputGstModel["query"] = "testString"
	supportedAttributeUiInputGstModel["value_property_name"] = "testString"
	supportedAttributeUiInputGstModel["label_property_name"] = "testString"
	supportedAttributeUiInputGstModel["input_option_label"] = "testString"

	supportedAttributeUiInputUrlModel := make(map[string]interface{})
	supportedAttributeUiInputUrlModel["url_endpoint"] = "testString"
	supportedAttributeUiInputUrlModel["input_option_label"] = "testString"

	model := make(map[string]interface{})
	model["type"] = "testString"
	model["values"] = []interface{}{supportedAttributeUiInputValueModel}
	model["gst"] = []interface{}{supportedAttributeUiInputGstModel}
	model["url"] = []interface{}{supportedAttributeUiInputUrlModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUiInputDetails(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUiInputValue(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedAttributeUiInputValue) {
		iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
		iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

		model := new(partnercentersellv1.SupportedAttributeUiInputValue)
		model.Value = core.StringPtr("testString")
		model.DisplayName = iamServiceRegistrationDisplayNameObjectModel

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
	iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

	model := make(map[string]interface{})
	model["value"] = "testString"
	model["display_name"] = []interface{}{iamServiceRegistrationDisplayNameObjectModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUiInputValue(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUiInputGst(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedAttributeUiInputGst) {
		model := new(partnercentersellv1.SupportedAttributeUiInputGst)
		model.Query = core.StringPtr("testString")
		model.ValuePropertyName = core.StringPtr("testString")
		model.LabelPropertyName = core.StringPtr("testString")
		model.InputOptionLabel = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["query"] = "testString"
	model["value_property_name"] = "testString"
	model["label_property_name"] = "testString"
	model["input_option_label"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUiInputGst(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUiInputURL(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedAttributeUiInputURL) {
		model := new(partnercentersellv1.SupportedAttributeUiInputURL)
		model.UrlEndpoint = core.StringPtr("testString")
		model.InputOptionLabel = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["url_endpoint"] = "testString"
	model["input_option_label"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedAttributeUiInputURL(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedAuthorizationSubject(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject) {
		supportAuthorizationSubjectAttributeModel := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
		supportAuthorizationSubjectAttributeModel.ServiceName = core.StringPtr("testString")
		supportAuthorizationSubjectAttributeModel.ResourceType = core.StringPtr("testString")

		model := new(partnercentersellv1.IamServiceRegistrationSupportedAuthorizationSubject)
		model.Attributes = supportAuthorizationSubjectAttributeModel
		model.Roles = []string{"testString"}

		assert.Equal(t, result, model)
	}

	supportAuthorizationSubjectAttributeModel := make(map[string]interface{})
	supportAuthorizationSubjectAttributeModel["service_name"] = "testString"
	supportAuthorizationSubjectAttributeModel["resource_type"] = "testString"

	model := make(map[string]interface{})
	model["attributes"] = []interface{}{supportAuthorizationSubjectAttributeModel}
	model["roles"] = []interface{}{"testString"}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedAuthorizationSubject(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportAuthorizationSubjectAttribute(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportAuthorizationSubjectAttribute) {
		model := new(partnercentersellv1.SupportAuthorizationSubjectAttribute)
		model.ServiceName = core.StringPtr("testString")
		model.ResourceType = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["service_name"] = "testString"
	model["resource_type"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportAuthorizationSubjectAttribute(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedRole(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationSupportedRole) {
		iamServiceRegistrationDescriptionObjectModel := new(partnercentersellv1.IamServiceRegistrationDescriptionObject)
		iamServiceRegistrationDescriptionObjectModel.Default = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.En = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.De = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Es = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Fr = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.It = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Ja = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.Ko = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.PtBr = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.ZhTw = core.StringPtr("testString")
		iamServiceRegistrationDescriptionObjectModel.ZhCn = core.StringPtr("testString")

		iamServiceRegistrationDisplayNameObjectModel := new(partnercentersellv1.IamServiceRegistrationDisplayNameObject)
		iamServiceRegistrationDisplayNameObjectModel.Default = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.En = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.De = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Es = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Fr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.It = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ja = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.Ko = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.PtBr = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhTw = core.StringPtr("testString")
		iamServiceRegistrationDisplayNameObjectModel.ZhCn = core.StringPtr("testString")

		supportedRoleOptionsModel := new(partnercentersellv1.SupportedRoleOptions)
		supportedRoleOptionsModel.AccessPolicy = core.BoolPtr(true)
		supportedRoleOptionsModel.PolicyType = []string{"access"}
		supportedRoleOptionsModel.AccountType = core.StringPtr("enterprise")

		model := new(partnercentersellv1.IamServiceRegistrationSupportedRole)
		model.ID = core.StringPtr("testString")
		model.Description = iamServiceRegistrationDescriptionObjectModel
		model.DisplayName = iamServiceRegistrationDisplayNameObjectModel
		model.Options = supportedRoleOptionsModel

		assert.Equal(t, result, model)
	}

	iamServiceRegistrationDescriptionObjectModel := make(map[string]interface{})
	iamServiceRegistrationDescriptionObjectModel["default"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["en"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["de"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["es"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["fr"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["it"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["ja"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["ko"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["pt_br"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["zh_tw"] = "testString"
	iamServiceRegistrationDescriptionObjectModel["zh_cn"] = "testString"

	iamServiceRegistrationDisplayNameObjectModel := make(map[string]interface{})
	iamServiceRegistrationDisplayNameObjectModel["default"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["en"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["de"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["es"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["fr"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["it"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ja"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["ko"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["pt_br"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_tw"] = "testString"
	iamServiceRegistrationDisplayNameObjectModel["zh_cn"] = "testString"

	supportedRoleOptionsModel := make(map[string]interface{})
	supportedRoleOptionsModel["access_policy"] = true
	supportedRoleOptionsModel["policy_type"] = []interface{}{"access"}
	supportedRoleOptionsModel["account_type"] = "enterprise"

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["description"] = []interface{}{iamServiceRegistrationDescriptionObjectModel}
	model["display_name"] = []interface{}{iamServiceRegistrationDisplayNameObjectModel}
	model["options"] = []interface{}{supportedRoleOptionsModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedRole(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToSupportedRoleOptions(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportedRoleOptions) {
		model := new(partnercentersellv1.SupportedRoleOptions)
		model.AccessPolicy = core.BoolPtr(true)
		model.PolicyType = []string{"access"}
		model.AccountType = core.StringPtr("enterprise")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["access_policy"] = true
	model["policy_type"] = []interface{}{"access"}
	model["account_type"] = "enterprise"

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToSupportedRoleOptions(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedNetwork(t *testing.T) {
	checkResult := func(result *partnercentersellv1.IamServiceRegistrationSupportedNetwork) {
		environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
		environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

		environmentAttributeModel := new(partnercentersellv1.EnvironmentAttribute)
		environmentAttributeModel.Key = core.StringPtr("testString")
		environmentAttributeModel.Values = []string{"testString"}
		environmentAttributeModel.Options = environmentAttributeOptionsModel

		model := new(partnercentersellv1.IamServiceRegistrationSupportedNetwork)
		model.EnvironmentAttributes = []partnercentersellv1.EnvironmentAttribute{*environmentAttributeModel}

		assert.Equal(t, result, model)
	}

	environmentAttributeOptionsModel := make(map[string]interface{})
	environmentAttributeOptionsModel["hidden"] = true

	environmentAttributeModel := make(map[string]interface{})
	environmentAttributeModel["key"] = "testString"
	environmentAttributeModel["values"] = []interface{}{"testString"}
	environmentAttributeModel["options"] = []interface{}{environmentAttributeOptionsModel}

	model := make(map[string]interface{})
	model["environment_attributes"] = []interface{}{environmentAttributeModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToIamServiceRegistrationSupportedNetwork(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToEnvironmentAttribute(t *testing.T) {
	checkResult := func(result *partnercentersellv1.EnvironmentAttribute) {
		environmentAttributeOptionsModel := new(partnercentersellv1.EnvironmentAttributeOptions)
		environmentAttributeOptionsModel.Hidden = core.BoolPtr(true)

		model := new(partnercentersellv1.EnvironmentAttribute)
		model.Key = core.StringPtr("testString")
		model.Values = []string{"testString"}
		model.Options = environmentAttributeOptionsModel

		assert.Equal(t, result, model)
	}

	environmentAttributeOptionsModel := make(map[string]interface{})
	environmentAttributeOptionsModel["hidden"] = true

	model := make(map[string]interface{})
	model["key"] = "testString"
	model["values"] = []interface{}{"testString"}
	model["options"] = []interface{}{environmentAttributeOptionsModel}

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToEnvironmentAttribute(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingIamRegistrationMapToEnvironmentAttributeOptions(t *testing.T) {
	checkResult := func(result *partnercentersellv1.EnvironmentAttributeOptions) {
		model := new(partnercentersellv1.EnvironmentAttributeOptions)
		model.Hidden = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["hidden"] = true

	result, err := partnercentersell.ResourceIbmOnboardingIamRegistrationMapToEnvironmentAttributeOptions(model)
	assert.Nil(t, err)
	checkResult(result)
}
