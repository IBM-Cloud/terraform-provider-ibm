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

func TestAccIbmLogsPoliciesDataSourceBasic(t *testing.T) {
	policyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyPriority := "type_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
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
					resource.TestCheckResourceAttr("data.ibm_logs_policies.logs_policies_instance", "policies.0.name", policyName),
					resource.TestCheckResourceAttr("data.ibm_logs_policies.logs_policies_instance", "policies.0.description", policyDescription),
					resource.TestCheckResourceAttr("data.ibm_logs_policies.logs_policies_instance", "policies.0.priority", policyPriority),
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
			name = "%s"
			priority = "%s"
		}

		data "ibm_logs_policies" "logs_policies_instance" {
			enabled_only = true
			source_type = "unspecified"
		}
	`, policyName, policyPriority)
}

func testAccCheckIbmLogsPoliciesDataSourceConfig(policyName string, policyDescription string, policyPriority string) string {
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

		data "ibm_logs_policies" "logs_policies_instance" {
			enabled_only = true
			source_type = "unspecified"
		}
	`, policyName, policyDescription, policyPriority)
}

func TestDataSourceIbmLogsPoliciesPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		quotaV1RuleModel := make(map[string]interface{})
		quotaV1RuleModel["rule_type_id"] = "is"
		quotaV1RuleModel["name"] = "cs-rest-test"

		quotaV1ArchiveRetentionModel := make(map[string]interface{})
		quotaV1ArchiveRetentionModel["id"] = "testString"

		quotaV1LogRulesModel := make(map[string]interface{})
		quotaV1LogRulesModel["severities"] = []string{"debug", "verbose", "info", "warning", "error"}

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["company_id"] = int(38)
		model["name"] = "testString"
		model["description"] = "testString"
		model["priority"] = "type_unspecified"
		model["deleted"] = true
		model["enabled"] = true
		model["order"] = int(38)
		model["application_rule"] = []map[string]interface{}{quotaV1RuleModel}
		model["subsystem_rule"] = []map[string]interface{}{quotaV1RuleModel}
		model["created_at"] = "testString"
		model["updated_at"] = "testString"
		model["archive_retention"] = []map[string]interface{}{quotaV1ArchiveRetentionModel}
		model["log_rules"] = []map[string]interface{}{quotaV1LogRulesModel}

		assert.Equal(t, result, model)
	}

	quotaV1RuleModel := new(logsv0.QuotaV1Rule)
	quotaV1RuleModel.RuleTypeID = core.StringPtr("is")
	quotaV1RuleModel.Name = core.StringPtr("cs-rest-test")

	quotaV1ArchiveRetentionModel := new(logsv0.QuotaV1ArchiveRetention)
	quotaV1ArchiveRetentionModel.ID = core.StringPtr("testString")

	quotaV1LogRulesModel := new(logsv0.QuotaV1LogRules)
	quotaV1LogRulesModel.Severities = []string{"debug", "verbose", "info", "warning", "error"}

	model := new(logsv0.Policy)
	model.ID = core.StringPtr("testString")
	model.CompanyID = core.Int64Ptr(int64(38))
	model.Name = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.Priority = core.StringPtr("type_unspecified")
	model.Deleted = core.BoolPtr(true)
	model.Enabled = core.BoolPtr(true)
	model.Order = core.Int64Ptr(int64(38))
	model.ApplicationRule = quotaV1RuleModel
	model.SubsystemRule = quotaV1RuleModel
	model.CreatedAt = core.StringPtr("testString")
	model.UpdatedAt = core.StringPtr("testString")
	model.ArchiveRetention = quotaV1ArchiveRetentionModel
	model.LogRules = quotaV1LogRulesModel

	result, err := logs.DataSourceIbmLogsPoliciesPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
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

func TestDataSourceIbmLogsPoliciesQuotaV1ArchiveRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.QuotaV1ArchiveRetention)
	model.ID = core.StringPtr("testString")

	result, err := logs.DataSourceIbmLogsPoliciesQuotaV1ArchiveRetentionToMap(model)
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

func TestDataSourceIbmLogsPoliciesPolicyQuotaV1PolicySourceTypeRulesLogRulesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		quotaV1RuleModel := make(map[string]interface{})
		quotaV1RuleModel["rule_type_id"] = "unspecified"
		quotaV1RuleModel["name"] = "testString"

		quotaV1ArchiveRetentionModel := make(map[string]interface{})
		quotaV1ArchiveRetentionModel["id"] = "testString"

		quotaV1LogRulesModel := make(map[string]interface{})
		quotaV1LogRulesModel["severities"] = []string{"unspecified"}

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["company_id"] = int(38)
		model["name"] = "testString"
		model["description"] = "testString"
		model["priority"] = "type_unspecified"
		model["deleted"] = true
		model["enabled"] = true
		model["order"] = int(38)
		model["application_rule"] = []map[string]interface{}{quotaV1RuleModel}
		model["subsystem_rule"] = []map[string]interface{}{quotaV1RuleModel}
		model["created_at"] = "testString"
		model["updated_at"] = "testString"
		model["archive_retention"] = []map[string]interface{}{quotaV1ArchiveRetentionModel}
		model["log_rules"] = []map[string]interface{}{quotaV1LogRulesModel}

		assert.Equal(t, result, model)
	}

	quotaV1RuleModel := new(logsv0.QuotaV1Rule)
	quotaV1RuleModel.RuleTypeID = core.StringPtr("unspecified")
	quotaV1RuleModel.Name = core.StringPtr("testString")

	quotaV1ArchiveRetentionModel := new(logsv0.QuotaV1ArchiveRetention)
	quotaV1ArchiveRetentionModel.ID = core.StringPtr("testString")

	quotaV1LogRulesModel := new(logsv0.QuotaV1LogRules)
	quotaV1LogRulesModel.Severities = []string{"unspecified"}

	model := new(logsv0.PolicyQuotaV1PolicySourceTypeRulesLogRules)
	model.ID = core.StringPtr("testString")
	model.CompanyID = core.Int64Ptr(int64(38))
	model.Name = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.Priority = core.StringPtr("type_unspecified")
	model.Deleted = core.BoolPtr(true)
	model.Enabled = core.BoolPtr(true)
	model.Order = core.Int64Ptr(int64(38))
	model.ApplicationRule = quotaV1RuleModel
	model.SubsystemRule = quotaV1RuleModel
	model.CreatedAt = core.StringPtr("testString")
	model.UpdatedAt = core.StringPtr("testString")
	model.ArchiveRetention = quotaV1ArchiveRetentionModel
	model.LogRules = quotaV1LogRulesModel

	result, err := logs.DataSourceIbmLogsPoliciesPolicyQuotaV1PolicySourceTypeRulesLogRulesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
