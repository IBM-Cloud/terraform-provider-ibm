// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package logs_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmLogsPolicyDataSourceBasic(t *testing.T) {
	policyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyPriority := "type_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyDataSourceConfigBasic(policyName, policyPriority),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "logs_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "company_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "order"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "updated_at"),
				),
			},
		},
	})
}

func TestAccIbmLogsPolicyDataSourceAllArgs(t *testing.T) {
	policyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	policyPriority := "type_unspecified"
	policyEnabled := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyDataSourceConfig(policyName, policyDescription, policyPriority, policyEnabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "logs_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "before.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "company_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "priority"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "deleted"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "order"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "application_rule.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "subsystem_rule.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "archive_retention.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "log_rules.#"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsPolicyDataSourceConfigBasic(policyName string, policyPriority string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_policy" "logs_policy_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			description = "Test description"
			priority    = "%s"
			application_rule {
				name         = "otel-links-test"
				rule_type_id = "start_with"
			}
			log_rules {
				severities = ["info"]
			}
		}

		data "ibm_logs_policy" "logs_policy_instance" {
			instance_id    = ibm_logs_policy.logs_policy_instance.instance_id
			region         = ibm_logs_policy.logs_policy_instance.region
			logs_policy_id = ibm_logs_policy.logs_policy_instance.policy_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, policyName, policyPriority)
}

func testAccCheckIbmLogsPolicyDataSourceConfig(policyName string, policyDescription string, policyPriority string, policyEnabled string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_policy" "logs_policy_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			description = "%s"
			priority    = "%s"
			enabled = %s
			application_rule {
				name         = "otel-links-test"
				rule_type_id = "start_with"
			}
			log_rules {
				severities = ["info"]
			}
		}

		data "ibm_logs_policy" "logs_policy_instance" {
			instance_id    = ibm_logs_policy.logs_policy_instance.instance_id
			region         = ibm_logs_policy.logs_policy_instance.region
			logs_policy_id = ibm_logs_policy.logs_policy_instance.policy_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, policyName, policyDescription, policyPriority, policyEnabled)
}
