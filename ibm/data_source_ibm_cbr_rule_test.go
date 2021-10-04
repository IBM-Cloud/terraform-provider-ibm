// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCbrRuleDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
	ruleTransactionID := fmt.Sprintf("tf_transaction_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCbrRuleDataSourceConfig(ruleDescription, ruleTransactionID),
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

func testAccCheckIBMCbrRuleDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cbr_rule" "cbr_rule" {
		}

		data "ibm_cbr_rule" "cbr_rule" {
			rule_id = "rule_id"
		}
	`)
}

func testAccCheckIBMCbrRuleDataSourceConfig(ruleDescription string, ruleTransactionID string) string {
	return fmt.Sprintf(`
		resource "ibm_cbr_rule" "cbr_rule" {
			description = "%s"
			contexts {
				attributes {
					name = "name"
					value = "value"
				}
			}
			resources {
				attributes {
					name = "name"
					value = "value"
					operator = "operator"
				}
				tags {
					name = "name"
					value = "value"
					operator = "operator"
				}
			}
			transaction_id = "%s"
		}

		data "ibm_cbr_rule" "cbr_rule" {
			rule_id = "rule_id"
		}
	`, ruleDescription, ruleTransactionID)
}
