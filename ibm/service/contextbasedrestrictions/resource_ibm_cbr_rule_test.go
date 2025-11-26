// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package contextbasedrestrictions_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMCbrRuleBasic(t *testing.T) {
	var conf contextbasedrestrictionsv1.Rule

	accountID, _ := getTestAccountAndZoneID()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCbr(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfigBasic(accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrRuleExists("ibm_cbr_rule.cbr_rule_instance", conf),
				),
			},
		},
	})
}

func TestAccIBMCbrRuleAllArgs(t *testing.T) {
	var conf contextbasedrestrictionsv1.Rule
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enforcementMode := "enabled"
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enforcementModeUpdate := "report"

	accountID, _ := getTestAccountAndZoneID()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCbr(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCbrRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfig(description, enforcementMode, accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCbrRuleExists("ibm_cbr_rule.cbr_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "enforcement_mode", enforcementMode),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleConfigUpdate(descriptionUpdate, enforcementModeUpdate, accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cbr_rule.cbr_rule_instance", "enforcement_mode", enforcementModeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cbr_rule.cbr_rule_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCbrRuleConfigBasic(accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		resource "ibm_cbr_rule" "cbr_rule_instance" {
  			description = "test rule config basic"
  			contexts {
    			attributes {
      				name = "networkZoneId"
      				value = ibm_cbr_zone.cbr_zone.id
    			}
  			}
			resources {
    			attributes {
      				name = "accountId"
      				value = "%s"
    			}
    			attributes {
      				name = "serviceName"
      				value = "user-management"
    			}
    			tags {
      				name     = "tag_name"
      				value    = "tag_value"
    			}
  			}
			enforcement_mode = "disabled"
		}
	`, accountID, accountID)
}

func testAccCheckIBMCbrRuleConfig(description string, enforcementMode string, accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		resource "ibm_cbr_rule" "cbr_rule_instance" {
			description = "%s"
			contexts {
    			attributes {
      				name = "networkZoneId"
      				value = ibm_cbr_zone.cbr_zone.id
    			}
			}
			resources {
    			attributes {
      				name = "accountId"
      				value = "%s"
    			}
    			attributes {
      				name = "serviceName"
      				value = "containers-kubernetes"
    			}
				tags {
					name = "name"
					value = "value"
					operator = "stringEquals"
				}
			}
			operations {
				api_types {
					api_type_id = "crn:v1:bluemix:public:containers-kubernetes::::api-type:management"
				}
			}
			enforcement_mode = "%s"
		}
	`, accountID, description, accountID, enforcementMode)
}

func testAccCheckIBMCbrRuleConfigUpdate(description string, enforcementMode string, accountID string) string {
	os.Setenv("IBMCLOUD_CONTEXT_BASED_RESTRICTIONS_ENDPOINT", "https://testing-2-eu-gb.network-policy.test.cloud.ibm.com")
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		resource "ibm_cbr_rule" "cbr_rule_instance" {
			description = "%s"
			contexts {
				attributes {
					name = "networkZoneId"
					value = ibm_cbr_zone.cbr_zone.id
				}
			}
			resources {
				attributes {
					name = "serviceName"
					value = "containers-kubernetes"
				}
				attributes {
					name = "accountId"
					value = "%s"
				}
				tags {
					name = "name"
					value = "value"
					operator = "stringEquals"
				}
			}
			operations {
				api_types {
					api_type_id = "crn:v1:bluemix:public:containers-kubernetes::::api-type:management"
				}
			}
			enforcement_mode = "%s"
		}
	`, accountID, description, accountID, enforcementMode)
}

func testAccCheckIBMCbrRuleExists(n string, obj contextbasedrestrictionsv1.Rule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
		if err != nil {
			return err
		}

		getRuleOptions := &contextbasedrestrictionsv1.GetRuleOptions{}

		getRuleOptions.SetRuleID(rs.Primary.ID)

		rule, _, err := contextBasedRestrictionsClient.GetRule(getRuleOptions)
		if err != nil {
			return err
		}

		obj = *rule
		return nil
	}
}

func testAccCheckIBMCbrRuleDestroy(s *terraform.State) error {
	contextBasedRestrictionsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContextBasedRestrictionsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cbr_rule" {
			continue
		}

		getRuleOptions := &contextbasedrestrictionsv1.GetRuleOptions{}

		getRuleOptions.SetRuleID(rs.Primary.ID)

		// Try to find the key
		_, response, err := contextBasedRestrictionsClient.GetRule(getRuleOptions)

		if err == nil {
			return fmt.Errorf("cbr_rule still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cbr_rule (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
