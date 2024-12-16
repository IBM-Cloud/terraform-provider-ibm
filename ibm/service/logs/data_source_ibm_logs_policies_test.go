// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPoliciesDataSourceConfig(policyName, policyDescription, policyPriority),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "enabled_only"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "source_type"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.company_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.priority"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.deleted"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policies.logs_policies_instance", "policies.0.enabled"),
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

func testAccCheckIbmLogsPoliciesDataSourceConfig(policyName string, policyDescription string, policyPriority string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_policy" "logs_policy_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			description = "%s"
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, policyName, policyDescription, policyPriority)
}

func TestDataSourceIbmLogsPoliciesQuotaV1RuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["rule_type_id"] = "unspecified"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.QuotaV1Rule)
	model.RuleTypeID = core.StringPtr("unspecified")
	model.Name = core.StringPtr("testString")

	result, err := logs.DataSourceIbmLogsPoliciesQuotaV1RuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsPoliciesQuotaV1LogRulesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["severities"] = []string{"unspecified"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.QuotaV1LogRules)
	model.Severities = []string{"unspecified"}

	result, err := logs.DataSourceIbmLogsPoliciesQuotaV1LogRulesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
