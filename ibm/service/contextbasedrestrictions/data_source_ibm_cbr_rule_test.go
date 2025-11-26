// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.95.2-120e65bc-20240924-152329
 */

package contextbasedrestrictions_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCbrRuleDataSourceBasic(t *testing.T) {
	accountID, _ := getTestAccountAndZoneID()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleDataSourceConfigBasic(accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "contexts.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "last_modified_by_id"),
				),
			},
		},
	})
}

func TestAccIBMCbrRuleDataSourceAllArgs(t *testing.T) {
	ruleDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	ruleEnforcementMode := "enabled"
	accountID, _ := getTestAccountAndZoneID()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleDataSourceConfig(ruleDescription, ruleEnforcementMode, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "contexts.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "operations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "enforcement_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule_instance", "last_modified_by_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCbrRuleDataSourceConfigBasic(accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone_instance" {
			name = "Test Zone Data Source Config Basic"
			description = "Test Zone Data Source Config Basic"
			account_id = "%s"
			addresses {
				type = "ipRange"
				value = "169.23.22.0-169.23.22.255"
			}
		}

		resource "ibm_cbr_rule" "cbr_rule_instance" {
 			description = "Test Rule Data Source Config Basic"
  			contexts {
    			attributes {
      				name = "networkZoneId"
      				value = ibm_cbr_zone.cbr_zone_instance.id
    			}
  			}
  			resources {
    			attributes {
      				name = "accountId"
      				value = "%s"
    			}
    			attributes {
      				name = "serviceName"
      				value = "iam-groups"
    			}
  			}
		}

		data "ibm_cbr_rule" "cbr_rule_instance" {
			rule_id = ibm_cbr_rule.cbr_rule_instance.id
		}
	`, accountID, accountID)
}

func testAccCheckIBMCbrRuleDataSourceConfig(ruleDescription, ruleEnforcementMode, accountID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_zone" "cbr_zone_instance" {
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
      				value = resource.ibm_cbr_zone.cbr_zone_instance.id
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
					value = "tag_name"
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

		data "ibm_cbr_rule" "cbr_rule_instance" {
			rule_id = ibm_cbr_rule.cbr_rule_instance.id
		}
	`, accountID, ruleDescription, accountID, ruleEnforcementMode)
}
