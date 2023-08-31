// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccRuleDataSourceBasic(t *testing.T) {
	ruleDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccRuleDataSourceConfigBasic(ruleDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "required_config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "labels.#"),
				),
			},
		},
	})
}

func TestAccIbmSccRuleDataSourceAllArgs(t *testing.T) {
	ruleDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	ruleVersion := fmt.Sprintf("tf_version_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccRuleDataSourceConfig(ruleDescription, ruleVersion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "import.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "required_config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule", "labels.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSccRuleDataSourceConfigBasic(ruleDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_rule" "scc_rule_instance" {
			description = "%s"
			target {
				service_name = "service_name"
				service_display_name = "service_display_name"
				resource_kind = "resource_kind"
				additional_target_attributes {
					name = "name"
					operator = "string_equals"
					value = "value"
				}
			}
			required_config {
				description = "description"
				and {
					description = "description"
					or {
						description = "description"
						property = "property"
						operator = "string_equals"
						value = "anything as a string"
					}
				}
			}
		}

		data "ibm_scc_rule" "scc_rule_instance" {
			rule_id = ibm_scc_rule.scc_rule_instance.rule_id
		}
	`, ruleDescription)
}

func testAccCheckIbmSccRuleDataSourceConfig(ruleDescription string, ruleVersion string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_rule" "scc_rule_instance" {
			description = "%s"
			version = "%s"
			import {
				parameters {
					name = "name"
					display_name = "display_name"
					description = "description"
					type = "string"
				}
			}
			target {
				service_name = "service_name"
				service_display_name = "service_display_name"
				resource_kind = "resource_kind"
				additional_target_attributes {
					name = "name"
					operator = "string_equals"
					value = "value"
				}
			}
			required_config {
				description = "description"
				and {
					description = "description"
					or {
						description = "description"
						property = "property"
						operator = "string_equals"
						value = "anything as a string"
					}
				}
			}
			labels = "FIXME"
		}

		data "ibm_scc_rule" "scc_rule_instance" {
			rule_id = ibm_scc_rule.scc_rule_instance.rule_id
		}
	`, ruleDescription, ruleVersion)
}
