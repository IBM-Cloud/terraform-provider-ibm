// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsPoliciesDataSourceBasic(t *testing.T) {
	policyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyPriority := "type_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPoliciesDataSourceConfigBasic(policyName, policyPriority),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "id"),
				),
			},
		},
	})
}

func TestAccIbmLogsPoliciesDataSourceAllArgs(t *testing.T) {
	policyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	policyPriority := "type_unspecified"
	policyEnabled := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPoliciesDataSourceConfig(policyName, policyDescription, policyPriority, policyEnabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "enabled_only"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "source_type"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.company_id"),
					resource.TestCheckResourceAttr("data.ibm_logs_policies.logs_policies_instance", "policies.0.name", policyName),
					resource.TestCheckResourceAttr("data.ibm_logs_policies.logs_policies_instance", "policies.0.description", policyDescription),
					resource.TestCheckResourceAttr("data.ibm_logs_policies.logs_policies_instance", "policies.0.priority", policyPriority),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.deleted"),
					resource.TestCheckResourceAttr("data.ibm_logs_policies.logs_policies_instance", "policies.0.enabled", policyEnabled),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.order"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.updated_at"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsPoliciesDataSourceConfigBasic(policyName string, policyPriority string) string {
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

		data "ibm_logs_policies" "logs_policies_instance" {
			instance_id  = ibm_logs_policy.logs_policy_instance.instance_id
			region       = ibm_logs_policy.logs_policy_instance.region
			enabled_only = true
			source_type  = "logs"
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, policyName, policyPriority)
}

func testAccCheckIbmLogsPoliciesDataSourceConfig(policyName string, policyDescription string, policyPriority string, policyEnabled string) string {
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

		data "ibm_logs_policies" "logs_policies_instance" {
			instance_id  = ibm_logs_policy.logs_policy_instance.instance_id
			region       = ibm_logs_policy.logs_policy_instance.region
			enabled_only = true
			source_type = "logs"
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, policyName, policyDescription, policyPriority, policyEnabled)
}
