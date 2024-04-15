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
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsPolicyDataSourceBasic(t *testing.T) {
	policyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyPriority := "type_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyDataSourceConfigBasic(policyName, policyPriority),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "logs_policy_id"),
				),
			},
		},
	})
}

func TestAccIbmLogsPolicyDataSourceAllArgs(t *testing.T) {
	policyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	policyPriority := "type_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyDataSourceConfig(policyName, policyDescription, policyPriority),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_policy.logs_policy_instance", "logs_policy_id"),
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
			name = "%s"
			priority = "%s"
		}

		data "ibm_logs_policy" "logs_policy_instance" {
			logs_policy_id = "logs_policy_id"
		}
	`, policyName, policyPriority)
}

func testAccCheckIbmLogsPolicyDataSourceConfig(policyName string, policyDescription string, policyPriority string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_policy" "logs_policy_instance" {
			name = "%s"
			description = "%s"
			priority = "%s"
			application_rule {
				rule_type_id = "unspecified"
				name = "name"
			}
			subsystem_rule {
				rule_type_id = "unspecified"
				name = "name"
			}
			archive_retention {
				id = "id"
			}
			log_rules {
				severities = [ "unspecified" ]
			}
		}

		data "ibm_logs_policy" "logs_policy_instance" {
			logs_policy_id = "logs_policy_id"
		}
	`, policyName, policyDescription, policyPriority)
}

func TestDataSourceIbmLogsPolicyQuotaV1RuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["rule_type_id"] = "unspecified"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.QuotaV1Rule)
	model.RuleTypeID = core.StringPtr("unspecified")
	model.Name = core.StringPtr("testString")

	result, err := logs.DataSourceIbmLogsPolicyQuotaV1RuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsPolicyQuotaV1ArchiveRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.QuotaV1ArchiveRetention)
	model.ID = core.StringPtr("testString")

	result, err := logs.DataSourceIbmLogsPolicyQuotaV1ArchiveRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsPolicyQuotaV1LogRulesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["severities"] = []string{"unspecified"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.QuotaV1LogRules)
	model.Severities = []string{"unspecified"}

	result, err := logs.DataSourceIbmLogsPolicyQuotaV1LogRulesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
