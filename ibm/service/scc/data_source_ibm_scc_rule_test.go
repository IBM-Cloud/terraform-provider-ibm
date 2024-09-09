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
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleDataSourceConfigBasic(acc.SccInstanceID, ruleDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "required_config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "labels.#"),
				),
			},
		},
	})
}

func TestAccIbmSccRuleDataSourceAllArgs(t *testing.T) {
	ruleDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	ruleVersion := "0.0.1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleDataSourceConfig(acc.SccInstanceID, ruleDescription, ruleVersion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "import.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "required_config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "labels.#"),
				),
			},
		},
	})
}

func TestAccIbmSccRuleDataSourcePreexistingOrIps(t *testing.T) {
	// rule with a required_config using "or" and "ips_not_equal"
	ruleID := "rule-9407e5a8-ec51-4228-a01a-0f32364224a6"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleDataSourcePreExistingRuleID(acc.SccInstanceID, ruleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "import.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "required_config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "labels.#"),
					resource.TestCheckResourceAttr("data.ibm_scc_rule.scc_rule_instance", "required_config.0.or.1.value", "[0.0.0.0/0]"),
				),
			},
		},
	})
}
func TestAccIbmSccRuleDataSourcePreexistingAndNumerical(t *testing.T) {
	// rule with a required_config using "and" and numerical value
	ruleID := "rule-0e5151b1-9caf-433c-b4e5-be3d505e458e"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleDataSourcePreExistingRuleID(acc.SccInstanceID, ruleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "import.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "required_config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "labels.#"),
					resource.TestCheckResourceAttr("data.ibm_scc_rule.scc_rule_instance", "required_config.0.and.1.value", "0"),
				),
			},
		},
	})
}
func TestAccIbmSccRuleDataSourcePreexistingSubRuleAny(t *testing.T) {
	// rule with a required_config using a subRule
	ruleID := "rule-5910ed25-7ad7-42d0-8e42-905df0123346"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleDataSourcePreExistingRuleID(acc.SccInstanceID, ruleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_on"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "import.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "required_config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_rule.scc_rule_instance", "labels.#"),
					resource.TestCheckResourceAttr("data.ibm_scc_rule.scc_rule_instance", "required_config.0.and.0.any.0.required_config.0.value", "${this-logdnaat}.region_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccRuleDataSourceConfigBasic(instanceID string, ruleDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_rule" "scc_rule_instance" {
			instance_id = "%s"
			description = "%s"
			target {
				service_name = "cloud-object-storage"
				resource_kind = "bucket"
			}
			labels = ["FIX_ME"]
			required_config {
				description = "required_config_description"
				and {
					description = "description"
					property = "storage_class"
					operator = "string_equals"
					value = "smart"
				}
			}
			version = "0.0.1"
		}

		data "ibm_scc_rule" "scc_rule_instance" {
			instance_id = resource.ibm_scc_rule.scc_rule_instance.instance_id
			rule_id = resource.ibm_scc_rule.scc_rule_instance.rule_id
		}
	`, instanceID, ruleDescription)
}

func testAccCheckIbmSccRuleDataSourceConfig(instanceID string, ruleDescription string, ruleVersion string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_rule" "scc_rule_instance" {
			instance_id = "%s"
			description = "%s"
			target {
				service_name = "cloud-object-storage"
				resource_kind = "bucket"
			}
			labels = ["FIX_ME"]
			required_config {
				description = "required_config_description"
				and {
					description = "description"
					property = "storage_class"
					operator = "string_equals"
					value = "smart"
				}
			}
			version = "%s"
		}

		data "ibm_scc_rule" "scc_rule_instance" {
			instance_id = resource.ibm_scc_rule.scc_rule_instance.instance_id
			rule_id = resource.ibm_scc_rule.scc_rule_instance.rule_id
		}
	`, instanceID, ruleDescription, ruleVersion)
}

func testAccCheckIbmSccRuleDataSourcePreExistingRuleID(instanceID string, ruleID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_rule" "scc_rule_instance" {
			instance_id = "%s"
			rule_id = "%s"
		}
	`, instanceID, ruleID)
}
