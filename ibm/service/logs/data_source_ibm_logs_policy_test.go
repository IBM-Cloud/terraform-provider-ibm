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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
					testAccCheckIbmLogsPolicyDataSourceBeforeExists("data.ibm_logs_policy.logs_policy_instance"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "company_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "priority"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "deleted"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "order"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "application_rule.#"),
					testAccCheckIbmLogsPolicyDataSourceSubsystemRuleExists("data.ibm_logs_policy.logs_policy_instance"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "updated_at"),
					testAccCheckIbmLogsPolicyDataSourceArchiveRetentionTagExists("data.ibm_logs_policy.logs_policy_instance"),
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

// testAccCheckIbmLogsPolicyDataSourceBeforeExists checks if the 'before' attribute exists
// This is a conditional check since 'before' is only set when there's a policy that comes before this one
func testAccCheckIbmLogsPolicyDataSourceBeforeExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		// Check if 'before' attribute exists in the state
		// If it exists, verify it's set; if it doesn't exist, that's also valid
		if _, ok := rs.Primary.Attributes["before.#"]; ok {
			// 'before' exists, so it should be set
			if rs.Primary.Attributes["before.#"] == "" {
				return fmt.Errorf("Attribute 'before.#' is empty")
			}
		}
		// If 'before' doesn't exist, that's fine - it's optional
		return nil
	}
}

// testAccCheckIbmLogsPolicyDataSourceSubsystemRuleExists checks if the 'subsystem_rule' attribute exists
// This is a conditional check since 'subsystem_rule' is optional and may not always be present
func testAccCheckIbmLogsPolicyDataSourceSubsystemRuleExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		// Check if 'subsystem_rule' attribute exists in the state
		// If it exists, verify it's set; if it doesn't exist, that's also valid
		if _, ok := rs.Primary.Attributes["subsystem_rule.#"]; ok {
			// 'subsystem_rule' exists, so it should be set
			if rs.Primary.Attributes["subsystem_rule.#"] == "" {
				return fmt.Errorf("Attribute 'subsystem_rule.#' is empty")
			}
		}
		// If 'subsystem_rule' doesn't exist, that's fine - it's optional
		return nil
	}
}

// testAccCheckIbmLogsPolicyDataSourceArchiveRetentionTagExists checks if the 'archive_retention_tag' attribute exists
// This is a conditional check since 'archive_retention_tag' is only present when retention tags are active
func testAccCheckIbmLogsPolicyDataSourceArchiveRetentionTagExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		// Check if 'archive_retention_tag' attribute exists in the state
		// If it exists, verify it's set; if it doesn't exist, that's also valid (retention tags not active)
		if _, ok := rs.Primary.Attributes["archive_retention_tag"]; ok {
			// 'archive_retention_tag' exists, so it should be set
			if rs.Primary.Attributes["archive_retention_tag"] == "" {
				return fmt.Errorf("Attribute 'archive_retention_tag' is empty")
			}
		}
		// If 'archive_retention_tag' doesn't exist, that's fine - retention tags may not be active
		return nil
	}
}
