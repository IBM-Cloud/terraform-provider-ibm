// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccRuleDataSourceBasic(t *testing.T) {
	ruleDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckSccInstanceID(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleDataSourceConfigBasic(instanceID, ruleDescription),
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
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckSccInstanceID(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccRuleDataSourceConfig(instanceID, ruleDescription, ruleVersion),
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
