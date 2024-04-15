// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsPolicyBasic(t *testing.T) {
	var conf logsv0.Policy
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	priority := "type_unspecified"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	priorityUpdate := "type_high"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyConfigBasic(name, priority),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsPolicyExists("ibm_logs_policy.logs_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "priority", priority),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyConfigBasic(nameUpdate, priorityUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "priority", priorityUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsPolicyAllArgs(t *testing.T) {
	var conf logsv0.Policy
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	priority := "type_unspecified"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	priorityUpdate := "type_high"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyConfig(name, description, priority),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsPolicyExists("ibm_logs_policy.logs_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "priority", priority),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyConfig(nameUpdate, descriptionUpdate, priorityUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "priority", priorityUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_policy.logs_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsPolicyConfigBasic(name string, priority string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_policy" "logs_policy_instance" {
			name = "%s"
			priority = "%s"
		}
	`, name, priority)
}

func testAccCheckIbmLogsPolicyConfig(name string, description string, priority string) string {
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
	`, name, description, priority)
}

func testAccCheckIbmLogsPolicyExists(n string, obj logsv0.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getPolicyOptions := &logsv0.GetPolicyOptions{}

		getPolicyOptions.SetID(resourceID[2])

		policyIntf, _, err := logsClient.GetPolicy(getPolicyOptions)
		if err != nil {
			return err
		}

		policy := policyIntf.(*logsv0.Policy)
		obj = *policy
		return nil
	}
}

func testAccCheckIbmLogsPolicyDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_policy" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getPolicyOptions := &logsv0.GetPolicyOptions{}

		getPolicyOptions.SetID(resourceID[2])

		// Try to find the key
		_, response, err := logsClient.GetPolicy(getPolicyOptions)

		if err == nil {
			return fmt.Errorf("logs_policy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_policy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmLogsPolicyQuotaV1RuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["rule_type_id"] = "unspecified"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.QuotaV1Rule)
	model.RuleTypeID = core.StringPtr("unspecified")
	model.Name = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsPolicyQuotaV1RuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsPolicyQuotaV1ArchiveRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.QuotaV1ArchiveRetention)
	model.ID = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsPolicyQuotaV1ArchiveRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsPolicyQuotaV1LogRulesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["severities"] = []string{"unspecified"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.QuotaV1LogRules)
	model.Severities = []string{"unspecified"}

	result, err := logs.ResourceIbmLogsPolicyQuotaV1LogRulesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsPolicyMapToQuotaV1Rule(t *testing.T) {
	checkResult := func(result *logsv0.QuotaV1Rule) {
		model := new(logsv0.QuotaV1Rule)
		model.RuleTypeID = core.StringPtr("unspecified")
		model.Name = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["rule_type_id"] = "unspecified"
	model["name"] = "testString"

	result, err := logs.ResourceIbmLogsPolicyMapToQuotaV1Rule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsPolicyMapToQuotaV1ArchiveRetention(t *testing.T) {
	checkResult := func(result *logsv0.QuotaV1ArchiveRetention) {
		model := new(logsv0.QuotaV1ArchiveRetention)
		model.ID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"

	result, err := logs.ResourceIbmLogsPolicyMapToQuotaV1ArchiveRetention(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsPolicyMapToQuotaV1LogRules(t *testing.T) {
	checkResult := func(result *logsv0.QuotaV1LogRules) {
		model := new(logsv0.QuotaV1LogRules)
		model.Severities = []string{"unspecified"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["severities"] = []interface{}{"unspecified"}

	result, err := logs.ResourceIbmLogsPolicyMapToQuotaV1LogRules(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsPolicyMapToPolicyPrototype(t *testing.T) {
	checkResult := func(result logsv0.PolicyPrototypeIntf) {
		quotaV1RuleModel := new(logsv0.QuotaV1Rule)
		quotaV1RuleModel.RuleTypeID = core.StringPtr("is")
		quotaV1RuleModel.Name = core.StringPtr("cs-rest-test")

		quotaV1ArchiveRetentionModel := new(logsv0.QuotaV1ArchiveRetention)
		quotaV1ArchiveRetentionModel.ID = core.StringPtr("testString")

		quotaV1LogRulesModel := new(logsv0.QuotaV1LogRules)
		quotaV1LogRulesModel.Severities = []string{"debug", "verbose", "info", "warning", "error"}

		model := new(logsv0.PolicyPrototype)
		model.Name = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")
		model.Priority = core.StringPtr("type_unspecified")
		model.ApplicationRule = quotaV1RuleModel
		model.SubsystemRule = quotaV1RuleModel
		model.ArchiveRetention = quotaV1ArchiveRetentionModel
		model.LogRules = quotaV1LogRulesModel

		assert.Equal(t, result, model)
	}

	quotaV1RuleModel := make(map[string]interface{})
	quotaV1RuleModel["rule_type_id"] = "is"
	quotaV1RuleModel["name"] = "cs-rest-test"

	quotaV1ArchiveRetentionModel := make(map[string]interface{})
	quotaV1ArchiveRetentionModel["id"] = "testString"

	quotaV1LogRulesModel := make(map[string]interface{})
	quotaV1LogRulesModel["severities"] = []interface{}{"debug", "verbose", "info", "warning", "error"}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["description"] = "testString"
	model["priority"] = "type_unspecified"
	model["application_rule"] = []interface{}{quotaV1RuleModel}
	model["subsystem_rule"] = []interface{}{quotaV1RuleModel}
	model["archive_retention"] = []interface{}{quotaV1ArchiveRetentionModel}
	model["log_rules"] = []interface{}{quotaV1LogRulesModel}

	result, err := logs.ResourceIbmLogsPolicyMapToPolicyPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsPolicyMapToPolicyPrototypeQuotaV1CreatePolicyRequestSourceTypeRulesLogRules(t *testing.T) {
	checkResult := func(result *logsv0.PolicyPrototypeQuotaV1CreatePolicyRequestSourceTypeRulesLogRules) {
		quotaV1RuleModel := new(logsv0.QuotaV1Rule)
		quotaV1RuleModel.RuleTypeID = core.StringPtr("unspecified")
		quotaV1RuleModel.Name = core.StringPtr("testString")

		quotaV1ArchiveRetentionModel := new(logsv0.QuotaV1ArchiveRetention)
		quotaV1ArchiveRetentionModel.ID = core.StringPtr("testString")

		quotaV1LogRulesModel := new(logsv0.QuotaV1LogRules)
		quotaV1LogRulesModel.Severities = []string{"unspecified"}

		model := new(logsv0.PolicyPrototypeQuotaV1CreatePolicyRequestSourceTypeRulesLogRules)
		model.Name = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")
		model.Priority = core.StringPtr("type_unspecified")
		model.ApplicationRule = quotaV1RuleModel
		model.SubsystemRule = quotaV1RuleModel
		model.ArchiveRetention = quotaV1ArchiveRetentionModel
		model.LogRules = quotaV1LogRulesModel

		assert.Equal(t, result, model)
	}

	quotaV1RuleModel := make(map[string]interface{})
	quotaV1RuleModel["rule_type_id"] = "unspecified"
	quotaV1RuleModel["name"] = "testString"

	quotaV1ArchiveRetentionModel := make(map[string]interface{})
	quotaV1ArchiveRetentionModel["id"] = "testString"

	quotaV1LogRulesModel := make(map[string]interface{})
	quotaV1LogRulesModel["severities"] = []interface{}{"unspecified"}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["description"] = "testString"
	model["priority"] = "type_unspecified"
	model["application_rule"] = []interface{}{quotaV1RuleModel}
	model["subsystem_rule"] = []interface{}{quotaV1RuleModel}
	model["archive_retention"] = []interface{}{quotaV1ArchiveRetentionModel}
	model["log_rules"] = []interface{}{quotaV1LogRulesModel}

	result, err := logs.ResourceIbmLogsPolicyMapToPolicyPrototypeQuotaV1CreatePolicyRequestSourceTypeRulesLogRules(model)
	assert.Nil(t, err)
	checkResult(result)
}
