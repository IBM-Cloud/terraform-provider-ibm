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
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

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
					// Check that at least one policy exists (the data source returns all policies)
					// We cannot assume our created policy will be at index 0
					testAccCheckIbmLogsPoliciesDataSourcePolicyExists(policyName, policyDescription, policyPriority),
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

// testAccCheckIbmLogsPoliciesDataSourcePolicyExists checks if a policy with the given attributes exists in the policies list
func testAccCheckIbmLogsPoliciesDataSourcePolicyExists(name, description, priority string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources["data.ibm_logs_policies.logs_policies_instance"]
		if !ok {
			return fmt.Errorf("Not found: data.ibm_logs_policies.logs_policies_instance")
		}

		// Get the number of policies
		policiesCount := 0
		for k := range rs.Primary.Attributes {
			if k == "policies.#" {
				fmt.Sscanf(rs.Primary.Attributes[k], "%d", &policiesCount)
				break
			}
		}

		if policiesCount == 0 {
			return fmt.Errorf("No policies found in data source")
		}

		// Search for our policy in the list
		for i := 0; i < policiesCount; i++ {
			policyName := rs.Primary.Attributes[fmt.Sprintf("policies.%d.name", i)]
			if policyName == name {
				// Found our policy, verify other attributes
				policyDesc := rs.Primary.Attributes[fmt.Sprintf("policies.%d.description", i)]
				policyPriority := rs.Primary.Attributes[fmt.Sprintf("policies.%d.priority", i)]

				if policyDesc != description {
					return fmt.Errorf("Policy description mismatch: expected %s, got %s", description, policyDesc)
				}
				if policyPriority != priority {
					return fmt.Errorf("Policy priority mismatch: expected %s, got %s", priority, policyPriority)
				}

				// All checks passed
				return nil
			}
		}

		return fmt.Errorf("Policy with name %s not found in the policies list", name)
	}
}
