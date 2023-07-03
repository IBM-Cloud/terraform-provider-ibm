// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package contextbasedrestrictions_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCbrRuleDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "contexts.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "last_modified_by_id"),
				),
			},
		},
	})
}

func TestAccIBMCbrRuleDataSourceAllArgs(t *testing.T) {
	ruleDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	ruleEnforcementMode := "enabled"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleDataSourceConfig(ruleDescription, ruleEnforcementMode),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "contexts.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "resources.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "operations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "enforcement_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "created_by_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "last_modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cbr_rule.cbr_rule", "last_modified_by_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCbrRuleDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cbr_rule" "cbr_rule" {
 			description = "Test Rule Data Source Config Basic"
  			contexts {
    			attributes {
      				name = "networkZoneId"
      				value = "559052eb8f43302824e7ae490c0281eb"
    			}
  			}
  			resources {
    			attributes {
      				name = "accountId"
      				value = "12ab34cd56ef78ab90cd12ef34ab56cd"
    			}
    			attributes {
      				name = "serviceName"
      				value = "iam-groups"
    			}
  			}
		}
		data "ibm_cbr_rule" "cbr_rule" {
			rule_id = ibm_cbr_rule.cbr_rule.id
		}
	`)
}

func testAccCheckIBMCbrRuleDataSourceConfig(ruleDescription string, ruleEnforcementMode string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_rule" "cbr_rule" {
			description = "%s"
			contexts {
    			attributes {
      				name = "networkZoneId"
      				value = "559052eb8f43302824e7ae490c0281eb"
    			}
			}
			resources {
				attributes {
      				name = "accountId"
      				value = "12ab34cd56ef78ab90cd12ef34ab56cd"
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

		data "ibm_cbr_rule" "cbr_rule" {
			rule_id = ibm_cbr_rule.cbr_rule.id
		}
	`, ruleDescription, ruleEnforcementMode)
}
