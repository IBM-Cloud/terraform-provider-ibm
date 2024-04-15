// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsRuleGroupBasic(t *testing.T) {
	var conf logsv0.RuleGroup
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsRuleGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsRuleGroupExists("ibm_logs_rule_group.logs_rule_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsRuleGroupAllArgs(t *testing.T) {
	var conf logsv0.RuleGroup
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	creator := fmt.Sprintf("tf_creator_%d", acctest.RandIntRange(10, 100))
	enabled := "false"
	order := fmt.Sprintf("%d", acctest.RandIntRange(0, 4294967295))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	creatorUpdate := fmt.Sprintf("tf_creator_%d", acctest.RandIntRange(10, 100))
	enabledUpdate := "true"
	orderUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 4294967295))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsRuleGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupConfig(name, description, creator, enabled, order),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsRuleGroupExists("ibm_logs_rule_group.logs_rule_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "creator", creator),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "order", order),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupConfig(nameUpdate, descriptionUpdate, creatorUpdate, enabledUpdate, orderUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "creator", creatorUpdate),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "order", orderUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_rule_group.logs_rule_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsRuleGroupConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_rule_group" "logs_rule_group_instance" {
			name = "%s"
			rule_subgroups {
				id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
				rules {
					id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
					name = "name"
					description = "description"
					source_field = "logObj.source"
					parameters {
						extract_parameters {
							rule = "rule"
						}
					}
					enabled = true
					order = 0
				}
				enabled = true
				order = 0
			}
		}
	`, name)
}

func testAccCheckIbmLogsRuleGroupConfig(name string, description string, creator string, enabled string, order string) string {
	return fmt.Sprintf(`

		resource "ibm_logs_rule_group" "logs_rule_group_instance" {
			name = "%s"
			description = "%s"
			creator = "%s"
			enabled = %s
			rule_matchers {
				application_name {
					value = "value"
				}
			}
			rule_subgroups {
				id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
				rules {
					id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
					name = "name"
					description = "description"
					source_field = "logObj.source"
					parameters {
						extract_parameters {
							rule = "rule"
						}
					}
					enabled = true
					order = 0
				}
				enabled = true
				order = 0
			}
			order = %s
		}
	`, name, description, creator, enabled, order)
}

func testAccCheckIbmLogsRuleGroupExists(n string, obj logsv0.RuleGroup) resource.TestCheckFunc {

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

		getRuleGroupOptions := &logsv0.GetRuleGroupOptions{}

		getRuleGroupOptions.SetGroupID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		ruleGroup, _, err := logsClient.GetRuleGroup(getRuleGroupOptions)
		if err != nil {
			return err
		}

		obj = *ruleGroup
		return nil
	}
}

func testAccCheckIbmLogsRuleGroupDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_rule_group" {
			continue
		}

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getRuleGroupOptions := &logsv0.GetRuleGroupOptions{}

		getRuleGroupOptions.SetGroupID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		// Try to find the key
		_, response, err := logsClient.GetRuleGroup(getRuleGroupOptions)

		if err == nil {
			return fmt.Errorf("logs_rule_group still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_rule_group (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsRuleGroupRulesV1RuleMatcherToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1ApplicationNameConstraintModel := make(map[string]interface{})
// 		rulesV1ApplicationNameConstraintModel["value"] = "testString"

// 		model := make(map[string]interface{})
// 		model["application_name"] = []map[string]interface{}{rulesV1ApplicationNameConstraintModel}
// 		model["subsystem_name"] = []map[string]interface{}{rulesV1SubsystemNameConstraintModel}
// 		model["severity"] = []map[string]interface{}{rulesV1SeverityConstraintModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ApplicationNameConstraintModel := new(logsv0.RulesV1ApplicationNameConstraint)
// 	rulesV1ApplicationNameConstraintModel.Value = core.StringPtr("testString")

// 	model := new(logsv0.RulesV1RuleMatcher)
// 	model.ApplicationName = rulesV1ApplicationNameConstraintModel
// 	model.SubsystemName = rulesV1SubsystemNameConstraintModel
// 	model.Severity = rulesV1SeverityConstraintModel

// 	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleMatcherToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsRuleGroupRulesV1ApplicationNameConstraintToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1ApplicationNameConstraint)
	model.Value = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1ApplicationNameConstraintToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1SubsystemNameConstraintToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1SubsystemNameConstraint)
	model.Value = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1SubsystemNameConstraintToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1SeverityConstraintToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["value"] = "debug_or_unspecified"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1SeverityConstraint)
	model.Value = core.StringPtr("debug_or_unspecified")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1SeverityConstraintToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintApplicationNameToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1ApplicationNameConstraintModel := make(map[string]interface{})
		rulesV1ApplicationNameConstraintModel["value"] = "testString"

		model := make(map[string]interface{})
		model["application_name"] = []map[string]interface{}{rulesV1ApplicationNameConstraintModel}

		assert.Equal(t, result, model)
	}

	rulesV1ApplicationNameConstraintModel := new(logsv0.RulesV1ApplicationNameConstraint)
	rulesV1ApplicationNameConstraintModel.Value = core.StringPtr("testString")

	model := new(logsv0.RulesV1RuleMatcherConstraintApplicationName)
	model.ApplicationName = rulesV1ApplicationNameConstraintModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintApplicationNameToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintSubsystemNameToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1SubsystemNameConstraintModel := make(map[string]interface{})
		rulesV1SubsystemNameConstraintModel["value"] = "testString"

		model := make(map[string]interface{})
		model["subsystem_name"] = []map[string]interface{}{rulesV1SubsystemNameConstraintModel}

		assert.Equal(t, result, model)
	}

	rulesV1SubsystemNameConstraintModel := new(logsv0.RulesV1SubsystemNameConstraint)
	rulesV1SubsystemNameConstraintModel.Value = core.StringPtr("testString")

	model := new(logsv0.RulesV1RuleMatcherConstraintSubsystemName)
	model.SubsystemName = rulesV1SubsystemNameConstraintModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintSubsystemNameToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintSeverityToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1SeverityConstraintModel := make(map[string]interface{})
		rulesV1SeverityConstraintModel["value"] = "debug_or_unspecified"

		model := make(map[string]interface{})
		model["severity"] = []map[string]interface{}{rulesV1SeverityConstraintModel}

		assert.Equal(t, result, model)
	}

	rulesV1SeverityConstraintModel := new(logsv0.RulesV1SeverityConstraint)
	rulesV1SeverityConstraintModel.Value = core.StringPtr("debug_or_unspecified")

	model := new(logsv0.RulesV1RuleMatcherConstraintSeverity)
	model.Severity = rulesV1SeverityConstraintModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintSeverityToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleSubgroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1ExtractParametersModel := make(map[string]interface{})
		rulesV1ExtractParametersModel["rule"] = "testString"

		rulesV1RuleParametersModel := make(map[string]interface{})
		rulesV1RuleParametersModel["extract_parameters"] = []map[string]interface{}{rulesV1ExtractParametersModel}

		rulesV1RuleModel := make(map[string]interface{})
		rulesV1RuleModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		rulesV1RuleModel["name"] = "testString"
		rulesV1RuleModel["description"] = "testString"
		rulesV1RuleModel["source_field"] = "logObj.source"
		rulesV1RuleModel["parameters"] = []map[string]interface{}{rulesV1RuleParametersModel}
		rulesV1RuleModel["enabled"] = true
		rulesV1RuleModel["order"] = int(0)

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["rules"] = []map[string]interface{}{rulesV1RuleModel}
		model["enabled"] = true
		model["order"] = int(0)

		assert.Equal(t, result, model)
	}

	rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
	rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

	rulesV1RuleParametersModel := new(logsv0.RulesV1RuleParametersRuleParametersExtractParameters)
	rulesV1RuleParametersModel.ExtractParameters = rulesV1ExtractParametersModel

	rulesV1RuleModel := new(logsv0.RulesV1Rule)
	rulesV1RuleModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	rulesV1RuleModel.Name = core.StringPtr("testString")
	rulesV1RuleModel.Description = core.StringPtr("testString")
	rulesV1RuleModel.SourceField = core.StringPtr("logObj.source")
	rulesV1RuleModel.Parameters = rulesV1RuleParametersModel
	rulesV1RuleModel.Enabled = core.BoolPtr(true)
	rulesV1RuleModel.Order = core.Int64Ptr(int64(0))

	model := new(logsv0.RulesV1RuleSubgroup)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Rules = []logsv0.RulesV1Rule{*rulesV1RuleModel}
	model.Enabled = core.BoolPtr(true)
	model.Order = core.Int64Ptr(int64(0))

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleSubgroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1ExtractParametersModel := make(map[string]interface{})
		rulesV1ExtractParametersModel["rule"] = "testString"

		rulesV1RuleParametersModel := make(map[string]interface{})
		rulesV1RuleParametersModel["extract_parameters"] = []map[string]interface{}{rulesV1ExtractParametersModel}

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "testString"
		model["description"] = "testString"
		model["source_field"] = "logObj.source"
		model["parameters"] = []map[string]interface{}{rulesV1RuleParametersModel}
		model["enabled"] = true
		model["order"] = int(0)

		assert.Equal(t, result, model)
	}

	rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
	rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

	rulesV1RuleParametersModel := new(logsv0.RulesV1RuleParametersRuleParametersExtractParameters)
	rulesV1RuleParametersModel.ExtractParameters = rulesV1ExtractParametersModel

	model := new(logsv0.RulesV1Rule)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.SourceField = core.StringPtr("logObj.source")
	model.Parameters = rulesV1RuleParametersModel
	model.Enabled = core.BoolPtr(true)
	model.Order = core.Int64Ptr(int64(0))

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsRuleGroupRulesV1RuleParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1ExtractParametersModel := make(map[string]interface{})
// 		rulesV1ExtractParametersModel["rule"] = "testString"

// 		model := make(map[string]interface{})
// 		model["extract_parameters"] = []map[string]interface{}{rulesV1ExtractParametersModel}
// 		model["json_extract_parameters"] = []map[string]interface{}{rulesV1JSONExtractParametersModel}
// 		model["replace_parameters"] = []map[string]interface{}{rulesV1ReplaceParametersModel}
// 		model["parse_parameters"] = []map[string]interface{}{rulesV1ParseParametersModel}
// 		model["allow_parameters"] = []map[string]interface{}{rulesV1AllowParametersModel}
// 		model["block_parameters"] = []map[string]interface{}{rulesV1BlockParametersModel}
// 		model["extract_timestamp_parameters"] = []map[string]interface{}{rulesV1ExtractTimestampParametersModel}
// 		model["remove_fields_parameters"] = []map[string]interface{}{rulesV1RemoveFieldsParametersModel}
// 		model["json_stringify_parameters"] = []map[string]interface{}{rulesV1JSONStringifyParametersModel}
// 		model["json_parse_parameters"] = []map[string]interface{}{rulesV1JSONParseParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
// 	rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

// 	model := new(logsv0.RulesV1RuleParameters)
// 	model.ExtractParameters = rulesV1ExtractParametersModel
// 	model.JSONExtractParameters = rulesV1JSONExtractParametersModel
// 	model.ReplaceParameters = rulesV1ReplaceParametersModel
// 	model.ParseParameters = rulesV1ParseParametersModel
// 	model.AllowParameters = rulesV1AllowParametersModel
// 	model.BlockParameters = rulesV1BlockParametersModel
// 	model.ExtractTimestampParameters = rulesV1ExtractTimestampParametersModel
// 	model.RemoveFieldsParameters = rulesV1RemoveFieldsParametersModel
// 	model.JSONStringifyParameters = rulesV1JSONStringifyParametersModel
// 	model.JSONParseParameters = rulesV1JSONParseParametersModel

// 	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsRuleGroupRulesV1ExtractParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["rule"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1ExtractParameters)
	model.Rule = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1ExtractParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1JSONExtractParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["destination_field"] = "category_or_unspecified"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1JSONExtractParameters)
	model.DestinationField = core.StringPtr("category_or_unspecified")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1JSONExtractParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1ReplaceParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["destination_field"] = "text.message"
		model["replace_new_val"] = "***"
		model["rule"] = "the password is (?P<password>[A-Za-z0-9!@#$].)"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1ReplaceParameters)
	model.DestinationField = core.StringPtr("text.message")
	model.ReplaceNewVal = core.StringPtr("***")
	model.Rule = core.StringPtr("the password is (?P<password>[A-Za-z0-9!@#$].)")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1ReplaceParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1ParseParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["destination_field"] = "text.message"
		model["rule"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1ParseParameters)
	model.DestinationField = core.StringPtr("text.message")
	model.Rule = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1ParseParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1AllowParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["keep_blocked_logs"] = true
		model["rule"] = "^this log should be kept!!!.*$"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1AllowParameters)
	model.KeepBlockedLogs = core.BoolPtr(true)
	model.Rule = core.StringPtr("^this log should be kept!!!.*$")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1AllowParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1BlockParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["keep_blocked_logs"] = true
		model["rule"] = "^this log should be blocked!!!.*$"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1BlockParameters)
	model.KeepBlockedLogs = core.BoolPtr(true)
	model.Rule = core.StringPtr("^this log should be blocked!!!.*$")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1BlockParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1ExtractTimestampParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["standard"] = "strftime_or_unspecified"
		model["format"] = "%Y-%m-%ddT%H:%M:%S.%f%z"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1ExtractTimestampParameters)
	model.Standard = core.StringPtr("strftime_or_unspecified")
	model.Format = core.StringPtr("%Y-%m-%ddT%H:%M:%S.%f%z")

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1ExtractTimestampParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RemoveFieldsParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["fields"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1RemoveFieldsParameters)
	model.Fields = []string{"testString"}

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RemoveFieldsParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1JSONStringifyParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["destination_field"] = "json.stringified"
		model["delete_source"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1JSONStringifyParameters)
	model.DestinationField = core.StringPtr("json.stringified")
	model.DeleteSource = core.BoolPtr(true)

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1JSONStringifyParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1JSONParseParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["destination_field"] = "json.content"
		model["delete_source"] = true
		model["override_dest"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.RulesV1JSONParseParameters)
	model.DestinationField = core.StringPtr("json.content")
	model.DeleteSource = core.BoolPtr(true)
	model.OverrideDest = core.BoolPtr(true)

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1JSONParseParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersExtractParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1ExtractParametersModel := make(map[string]interface{})
		rulesV1ExtractParametersModel["rule"] = "testString"

		model := make(map[string]interface{})
		model["extract_parameters"] = []map[string]interface{}{rulesV1ExtractParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
	rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

	model := new(logsv0.RulesV1RuleParametersRuleParametersExtractParameters)
	model.ExtractParameters = rulesV1ExtractParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersExtractParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONExtractParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1JSONExtractParametersModel := make(map[string]interface{})
		rulesV1JSONExtractParametersModel["destination_field"] = "category_or_unspecified"

		model := make(map[string]interface{})
		model["json_extract_parameters"] = []map[string]interface{}{rulesV1JSONExtractParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1JSONExtractParametersModel := new(logsv0.RulesV1JSONExtractParameters)
	rulesV1JSONExtractParametersModel.DestinationField = core.StringPtr("category_or_unspecified")

	model := new(logsv0.RulesV1RuleParametersRuleParametersJSONExtractParameters)
	model.JSONExtractParameters = rulesV1JSONExtractParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONExtractParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersReplaceParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1ReplaceParametersModel := make(map[string]interface{})
		rulesV1ReplaceParametersModel["destination_field"] = "text.message"
		rulesV1ReplaceParametersModel["replace_new_val"] = "***"
		rulesV1ReplaceParametersModel["rule"] = "the password is (?P<password>[A-Za-z0-9!@#$].)"

		model := make(map[string]interface{})
		model["replace_parameters"] = []map[string]interface{}{rulesV1ReplaceParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1ReplaceParametersModel := new(logsv0.RulesV1ReplaceParameters)
	rulesV1ReplaceParametersModel.DestinationField = core.StringPtr("text.message")
	rulesV1ReplaceParametersModel.ReplaceNewVal = core.StringPtr("***")
	rulesV1ReplaceParametersModel.Rule = core.StringPtr("the password is (?P<password>[A-Za-z0-9!@#$].)")

	model := new(logsv0.RulesV1RuleParametersRuleParametersReplaceParameters)
	model.ReplaceParameters = rulesV1ReplaceParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersReplaceParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersParseParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1ParseParametersModel := make(map[string]interface{})
		rulesV1ParseParametersModel["destination_field"] = "text.message"
		rulesV1ParseParametersModel["rule"] = "testString"

		model := make(map[string]interface{})
		model["parse_parameters"] = []map[string]interface{}{rulesV1ParseParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1ParseParametersModel := new(logsv0.RulesV1ParseParameters)
	rulesV1ParseParametersModel.DestinationField = core.StringPtr("text.message")
	rulesV1ParseParametersModel.Rule = core.StringPtr("testString")

	model := new(logsv0.RulesV1RuleParametersRuleParametersParseParameters)
	model.ParseParameters = rulesV1ParseParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersParseParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersAllowParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1AllowParametersModel := make(map[string]interface{})
		rulesV1AllowParametersModel["keep_blocked_logs"] = true
		rulesV1AllowParametersModel["rule"] = "^this log should be kept!!!.*$"

		model := make(map[string]interface{})
		model["allow_parameters"] = []map[string]interface{}{rulesV1AllowParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1AllowParametersModel := new(logsv0.RulesV1AllowParameters)
	rulesV1AllowParametersModel.KeepBlockedLogs = core.BoolPtr(true)
	rulesV1AllowParametersModel.Rule = core.StringPtr("^this log should be kept!!!.*$")

	model := new(logsv0.RulesV1RuleParametersRuleParametersAllowParameters)
	model.AllowParameters = rulesV1AllowParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersAllowParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersBlockParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1BlockParametersModel := make(map[string]interface{})
		rulesV1BlockParametersModel["keep_blocked_logs"] = true
		rulesV1BlockParametersModel["rule"] = "^this log should be blocked!!!.*$"

		model := make(map[string]interface{})
		model["block_parameters"] = []map[string]interface{}{rulesV1BlockParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1BlockParametersModel := new(logsv0.RulesV1BlockParameters)
	rulesV1BlockParametersModel.KeepBlockedLogs = core.BoolPtr(true)
	rulesV1BlockParametersModel.Rule = core.StringPtr("^this log should be blocked!!!.*$")

	model := new(logsv0.RulesV1RuleParametersRuleParametersBlockParameters)
	model.BlockParameters = rulesV1BlockParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersBlockParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersExtractTimestampParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1ExtractTimestampParametersModel := make(map[string]interface{})
		rulesV1ExtractTimestampParametersModel["standard"] = "strftime_or_unspecified"
		rulesV1ExtractTimestampParametersModel["format"] = "%Y-%m-%ddT%H:%M:%S.%f%z"

		model := make(map[string]interface{})
		model["extract_timestamp_parameters"] = []map[string]interface{}{rulesV1ExtractTimestampParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1ExtractTimestampParametersModel := new(logsv0.RulesV1ExtractTimestampParameters)
	rulesV1ExtractTimestampParametersModel.Standard = core.StringPtr("strftime_or_unspecified")
	rulesV1ExtractTimestampParametersModel.Format = core.StringPtr("%Y-%m-%ddT%H:%M:%S.%f%z")

	model := new(logsv0.RulesV1RuleParametersRuleParametersExtractTimestampParameters)
	model.ExtractTimestampParameters = rulesV1ExtractTimestampParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersExtractTimestampParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersRemoveFieldsParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1RemoveFieldsParametersModel := make(map[string]interface{})
		rulesV1RemoveFieldsParametersModel["fields"] = []string{"testString"}

		model := make(map[string]interface{})
		model["remove_fields_parameters"] = []map[string]interface{}{rulesV1RemoveFieldsParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1RemoveFieldsParametersModel := new(logsv0.RulesV1RemoveFieldsParameters)
	rulesV1RemoveFieldsParametersModel.Fields = []string{"testString"}

	model := new(logsv0.RulesV1RuleParametersRuleParametersRemoveFieldsParameters)
	model.RemoveFieldsParameters = rulesV1RemoveFieldsParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersRemoveFieldsParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONStringifyParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1JSONStringifyParametersModel := make(map[string]interface{})
		rulesV1JSONStringifyParametersModel["destination_field"] = "json.stringified"
		rulesV1JSONStringifyParametersModel["delete_source"] = true

		model := make(map[string]interface{})
		model["json_stringify_parameters"] = []map[string]interface{}{rulesV1JSONStringifyParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1JSONStringifyParametersModel := new(logsv0.RulesV1JSONStringifyParameters)
	rulesV1JSONStringifyParametersModel.DestinationField = core.StringPtr("json.stringified")
	rulesV1JSONStringifyParametersModel.DeleteSource = core.BoolPtr(true)

	model := new(logsv0.RulesV1RuleParametersRuleParametersJSONStringifyParameters)
	model.JSONStringifyParameters = rulesV1JSONStringifyParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONStringifyParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONParseParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		rulesV1JSONParseParametersModel := make(map[string]interface{})
		rulesV1JSONParseParametersModel["destination_field"] = "json.content"
		rulesV1JSONParseParametersModel["delete_source"] = true
		rulesV1JSONParseParametersModel["override_dest"] = true

		model := make(map[string]interface{})
		model["json_parse_parameters"] = []map[string]interface{}{rulesV1JSONParseParametersModel}

		assert.Equal(t, result, model)
	}

	rulesV1JSONParseParametersModel := new(logsv0.RulesV1JSONParseParameters)
	rulesV1JSONParseParametersModel.DestinationField = core.StringPtr("json.content")
	rulesV1JSONParseParametersModel.DeleteSource = core.BoolPtr(true)
	rulesV1JSONParseParametersModel.OverrideDest = core.BoolPtr(true)

	model := new(logsv0.RulesV1RuleParametersRuleParametersJSONParseParameters)
	model.JSONParseParameters = rulesV1JSONParseParametersModel

	result, err := logs.ResourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONParseParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1CreateRuleGroupRequestCreateRuleSubgroup(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1CreateRuleGroupRequestCreateRuleSubgroup) {
		rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
		rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

		rulesV1RuleParametersModel := new(logsv0.RulesV1RuleParametersRuleParametersExtractParameters)
		rulesV1RuleParametersModel.ExtractParameters = rulesV1ExtractParametersModel

		rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel := new(logsv0.RulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRule)
		rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel.Name = core.StringPtr("Extract Timestamp")
		rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel.SourceField = core.StringPtr("log_obj.date_time")
		rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel.Parameters = rulesV1RuleParametersModel
		rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel.Enabled = core.BoolPtr(true)
		rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel.Order = core.Int64Ptr(int64(0))
		rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel.Description = core.StringPtr("Extract timestamp with ISO format for Mysql logs")

		model := new(logsv0.RulesV1CreateRuleGroupRequestCreateRuleSubgroup)
		model.Rules = []logsv0.RulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRule{*rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel}
		model.Enabled = core.BoolPtr(true)
		model.Order = core.Int64Ptr(int64(0))

		assert.Equal(t, result, model)
	}

	rulesV1ExtractParametersModel := make(map[string]interface{})
	rulesV1ExtractParametersModel["rule"] = "testString"

	rulesV1RuleParametersModel := make(map[string]interface{})
	rulesV1RuleParametersModel["extract_parameters"] = []interface{}{rulesV1ExtractParametersModel}

	rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel := make(map[string]interface{})
	rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel["name"] = "Extract Timestamp"
	rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel["source_field"] = "log_obj.date_time"
	rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel["parameters"] = []interface{}{rulesV1RuleParametersModel}
	rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel["enabled"] = true
	rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel["order"] = int(0)
	rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel["description"] = "Extract timestamp with ISO format for Mysql logs"

	model := make(map[string]interface{})
	model["rules"] = []interface{}{rulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRuleModel}
	model["enabled"] = true
	model["order"] = int(0)

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1CreateRuleGroupRequestCreateRuleSubgroup(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRule(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRule) {
		rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
		rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

		rulesV1RuleParametersModel := new(logsv0.RulesV1RuleParametersRuleParametersExtractParameters)
		rulesV1RuleParametersModel.ExtractParameters = rulesV1ExtractParametersModel

		model := new(logsv0.RulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRule)
		model.Name = core.StringPtr("Extract Timestamp")
		model.SourceField = core.StringPtr("log_obj.date_time")
		model.Parameters = rulesV1RuleParametersModel
		model.Enabled = core.BoolPtr(true)
		model.Order = core.Int64Ptr(int64(0))
		model.Description = core.StringPtr("Extract timestamp with ISO format for Mysql logs")

		assert.Equal(t, result, model)
	}

	rulesV1ExtractParametersModel := make(map[string]interface{})
	rulesV1ExtractParametersModel["rule"] = "testString"

	rulesV1RuleParametersModel := make(map[string]interface{})
	rulesV1RuleParametersModel["extract_parameters"] = []interface{}{rulesV1ExtractParametersModel}

	model := make(map[string]interface{})
	model["name"] = "Extract Timestamp"
	model["source_field"] = "log_obj.date_time"
	model["parameters"] = []interface{}{rulesV1RuleParametersModel}
	model["enabled"] = true
	model["order"] = int(0)
	model["description"] = "Extract timestamp with ISO format for Mysql logs"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1CreateRuleGroupRequestCreateRuleSubgroupCreateRule(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParameters(t *testing.T) {
// 	checkResult := func(result logsv0.RulesV1RuleParametersIntf) {
// 		rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
// 		rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

// 		model := new(logsv0.RulesV1RuleParameters)
// 		model.ExtractParameters = rulesV1ExtractParametersModel
// 		model.JSONExtractParameters = rulesV1JSONExtractParametersModel
// 		model.ReplaceParameters = rulesV1ReplaceParametersModel
// 		model.ParseParameters = rulesV1ParseParametersModel
// 		model.AllowParameters = rulesV1AllowParametersModel
// 		model.BlockParameters = rulesV1BlockParametersModel
// 		model.ExtractTimestampParameters = rulesV1ExtractTimestampParametersModel
// 		model.RemoveFieldsParameters = rulesV1RemoveFieldsParametersModel
// 		model.JSONStringifyParameters = rulesV1JSONStringifyParametersModel
// 		model.JSONParseParameters = rulesV1JSONParseParametersModel

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ExtractParametersModel := make(map[string]interface{})
// 	rulesV1ExtractParametersModel["rule"] = "testString"

// 	model := make(map[string]interface{})
// 	model["extract_parameters"] = []interface{}{rulesV1ExtractParametersModel}
// 	model["json_extract_parameters"] = []interface{}{rulesV1JSONExtractParametersModel}
// 	model["replace_parameters"] = []interface{}{rulesV1ReplaceParametersModel}
// 	model["parse_parameters"] = []interface{}{rulesV1ParseParametersModel}
// 	model["allow_parameters"] = []interface{}{rulesV1AllowParametersModel}
// 	model["block_parameters"] = []interface{}{rulesV1BlockParametersModel}
// 	model["extract_timestamp_parameters"] = []interface{}{rulesV1ExtractTimestampParametersModel}
// 	model["remove_fields_parameters"] = []interface{}{rulesV1RemoveFieldsParametersModel}
// 	model["json_stringify_parameters"] = []interface{}{rulesV1JSONStringifyParametersModel}
// 	model["json_parse_parameters"] = []interface{}{rulesV1JSONParseParametersModel}

// 	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParameters(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsRuleGroupMapToRulesV1ExtractParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1ExtractParameters) {
		model := new(logsv0.RulesV1ExtractParameters)
		model.Rule = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["rule"] = "testString"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1ExtractParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1JSONExtractParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1JSONExtractParameters) {
		model := new(logsv0.RulesV1JSONExtractParameters)
		model.DestinationField = core.StringPtr("category_or_unspecified")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["destination_field"] = "category_or_unspecified"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1JSONExtractParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1ReplaceParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1ReplaceParameters) {
		model := new(logsv0.RulesV1ReplaceParameters)
		model.DestinationField = core.StringPtr("text.message")
		model.ReplaceNewVal = core.StringPtr("***")
		model.Rule = core.StringPtr("the password is (?P<password>[A-Za-z0-9!@#$].)")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["destination_field"] = "text.message"
	model["replace_new_val"] = "***"
	model["rule"] = "the password is (?P<password>[A-Za-z0-9!@#$].)"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1ReplaceParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1ParseParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1ParseParameters) {
		model := new(logsv0.RulesV1ParseParameters)
		model.DestinationField = core.StringPtr("text.message")
		model.Rule = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["destination_field"] = "text.message"
	model["rule"] = "testString"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1ParseParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1AllowParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1AllowParameters) {
		model := new(logsv0.RulesV1AllowParameters)
		model.KeepBlockedLogs = core.BoolPtr(true)
		model.Rule = core.StringPtr("^this log should be kept!!!.*$")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["keep_blocked_logs"] = true
	model["rule"] = "^this log should be kept!!!.*$"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1AllowParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1BlockParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1BlockParameters) {
		model := new(logsv0.RulesV1BlockParameters)
		model.KeepBlockedLogs = core.BoolPtr(true)
		model.Rule = core.StringPtr("^this log should be blocked!!!.*$")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["keep_blocked_logs"] = true
	model["rule"] = "^this log should be blocked!!!.*$"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1BlockParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1ExtractTimestampParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1ExtractTimestampParameters) {
		model := new(logsv0.RulesV1ExtractTimestampParameters)
		model.Standard = core.StringPtr("strftime_or_unspecified")
		model.Format = core.StringPtr("%Y-%m-%ddT%H:%M:%S.%f%z")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["standard"] = "strftime_or_unspecified"
	model["format"] = "%Y-%m-%ddT%H:%M:%S.%f%z"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1ExtractTimestampParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RemoveFieldsParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RemoveFieldsParameters) {
		model := new(logsv0.RulesV1RemoveFieldsParameters)
		model.Fields = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["fields"] = []interface{}{"testString"}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RemoveFieldsParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1JSONStringifyParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1JSONStringifyParameters) {
		model := new(logsv0.RulesV1JSONStringifyParameters)
		model.DestinationField = core.StringPtr("json.stringified")
		model.DeleteSource = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["destination_field"] = "json.stringified"
	model["delete_source"] = true

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1JSONStringifyParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1JSONParseParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1JSONParseParameters) {
		model := new(logsv0.RulesV1JSONParseParameters)
		model.DestinationField = core.StringPtr("json.content")
		model.DeleteSource = core.BoolPtr(true)
		model.OverrideDest = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["destination_field"] = "json.content"
	model["delete_source"] = true
	model["override_dest"] = true

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1JSONParseParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersExtractParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersExtractParameters) {
		rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
		rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

		model := new(logsv0.RulesV1RuleParametersRuleParametersExtractParameters)
		model.ExtractParameters = rulesV1ExtractParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1ExtractParametersModel := make(map[string]interface{})
	rulesV1ExtractParametersModel["rule"] = "testString"

	model := make(map[string]interface{})
	model["extract_parameters"] = []interface{}{rulesV1ExtractParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersExtractParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersJSONExtractParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersJSONExtractParameters) {
		rulesV1JSONExtractParametersModel := new(logsv0.RulesV1JSONExtractParameters)
		rulesV1JSONExtractParametersModel.DestinationField = core.StringPtr("category_or_unspecified")

		model := new(logsv0.RulesV1RuleParametersRuleParametersJSONExtractParameters)
		model.JSONExtractParameters = rulesV1JSONExtractParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1JSONExtractParametersModel := make(map[string]interface{})
	rulesV1JSONExtractParametersModel["destination_field"] = "category_or_unspecified"

	model := make(map[string]interface{})
	model["json_extract_parameters"] = []interface{}{rulesV1JSONExtractParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersJSONExtractParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersReplaceParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersReplaceParameters) {
		rulesV1ReplaceParametersModel := new(logsv0.RulesV1ReplaceParameters)
		rulesV1ReplaceParametersModel.DestinationField = core.StringPtr("text.message")
		rulesV1ReplaceParametersModel.ReplaceNewVal = core.StringPtr("***")
		rulesV1ReplaceParametersModel.Rule = core.StringPtr("the password is (?P<password>[A-Za-z0-9!@#$].)")

		model := new(logsv0.RulesV1RuleParametersRuleParametersReplaceParameters)
		model.ReplaceParameters = rulesV1ReplaceParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1ReplaceParametersModel := make(map[string]interface{})
	rulesV1ReplaceParametersModel["destination_field"] = "text.message"
	rulesV1ReplaceParametersModel["replace_new_val"] = "***"
	rulesV1ReplaceParametersModel["rule"] = "the password is (?P<password>[A-Za-z0-9!@#$].)"

	model := make(map[string]interface{})
	model["replace_parameters"] = []interface{}{rulesV1ReplaceParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersReplaceParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersParseParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersParseParameters) {
		rulesV1ParseParametersModel := new(logsv0.RulesV1ParseParameters)
		rulesV1ParseParametersModel.DestinationField = core.StringPtr("text.message")
		rulesV1ParseParametersModel.Rule = core.StringPtr("testString")

		model := new(logsv0.RulesV1RuleParametersRuleParametersParseParameters)
		model.ParseParameters = rulesV1ParseParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1ParseParametersModel := make(map[string]interface{})
	rulesV1ParseParametersModel["destination_field"] = "text.message"
	rulesV1ParseParametersModel["rule"] = "testString"

	model := make(map[string]interface{})
	model["parse_parameters"] = []interface{}{rulesV1ParseParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersParseParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersAllowParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersAllowParameters) {
		rulesV1AllowParametersModel := new(logsv0.RulesV1AllowParameters)
		rulesV1AllowParametersModel.KeepBlockedLogs = core.BoolPtr(true)
		rulesV1AllowParametersModel.Rule = core.StringPtr("^this log should be kept!!!.*$")

		model := new(logsv0.RulesV1RuleParametersRuleParametersAllowParameters)
		model.AllowParameters = rulesV1AllowParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1AllowParametersModel := make(map[string]interface{})
	rulesV1AllowParametersModel["keep_blocked_logs"] = true
	rulesV1AllowParametersModel["rule"] = "^this log should be kept!!!.*$"

	model := make(map[string]interface{})
	model["allow_parameters"] = []interface{}{rulesV1AllowParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersAllowParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersBlockParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersBlockParameters) {
		rulesV1BlockParametersModel := new(logsv0.RulesV1BlockParameters)
		rulesV1BlockParametersModel.KeepBlockedLogs = core.BoolPtr(true)
		rulesV1BlockParametersModel.Rule = core.StringPtr("^this log should be blocked!!!.*$")

		model := new(logsv0.RulesV1RuleParametersRuleParametersBlockParameters)
		model.BlockParameters = rulesV1BlockParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1BlockParametersModel := make(map[string]interface{})
	rulesV1BlockParametersModel["keep_blocked_logs"] = true
	rulesV1BlockParametersModel["rule"] = "^this log should be blocked!!!.*$"

	model := make(map[string]interface{})
	model["block_parameters"] = []interface{}{rulesV1BlockParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersBlockParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersExtractTimestampParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersExtractTimestampParameters) {
		rulesV1ExtractTimestampParametersModel := new(logsv0.RulesV1ExtractTimestampParameters)
		rulesV1ExtractTimestampParametersModel.Standard = core.StringPtr("strftime_or_unspecified")
		rulesV1ExtractTimestampParametersModel.Format = core.StringPtr("%Y-%m-%ddT%H:%M:%S.%f%z")

		model := new(logsv0.RulesV1RuleParametersRuleParametersExtractTimestampParameters)
		model.ExtractTimestampParameters = rulesV1ExtractTimestampParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1ExtractTimestampParametersModel := make(map[string]interface{})
	rulesV1ExtractTimestampParametersModel["standard"] = "strftime_or_unspecified"
	rulesV1ExtractTimestampParametersModel["format"] = "%Y-%m-%ddT%H:%M:%S.%f%z"

	model := make(map[string]interface{})
	model["extract_timestamp_parameters"] = []interface{}{rulesV1ExtractTimestampParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersExtractTimestampParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersRemoveFieldsParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersRemoveFieldsParameters) {
		rulesV1RemoveFieldsParametersModel := new(logsv0.RulesV1RemoveFieldsParameters)
		rulesV1RemoveFieldsParametersModel.Fields = []string{"testString"}

		model := new(logsv0.RulesV1RuleParametersRuleParametersRemoveFieldsParameters)
		model.RemoveFieldsParameters = rulesV1RemoveFieldsParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1RemoveFieldsParametersModel := make(map[string]interface{})
	rulesV1RemoveFieldsParametersModel["fields"] = []interface{}{"testString"}

	model := make(map[string]interface{})
	model["remove_fields_parameters"] = []interface{}{rulesV1RemoveFieldsParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersRemoveFieldsParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersJSONStringifyParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersJSONStringifyParameters) {
		rulesV1JSONStringifyParametersModel := new(logsv0.RulesV1JSONStringifyParameters)
		rulesV1JSONStringifyParametersModel.DestinationField = core.StringPtr("json.stringified")
		rulesV1JSONStringifyParametersModel.DeleteSource = core.BoolPtr(true)

		model := new(logsv0.RulesV1RuleParametersRuleParametersJSONStringifyParameters)
		model.JSONStringifyParameters = rulesV1JSONStringifyParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1JSONStringifyParametersModel := make(map[string]interface{})
	rulesV1JSONStringifyParametersModel["destination_field"] = "json.stringified"
	rulesV1JSONStringifyParametersModel["delete_source"] = true

	model := make(map[string]interface{})
	model["json_stringify_parameters"] = []interface{}{rulesV1JSONStringifyParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersJSONStringifyParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersJSONParseParameters(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleParametersRuleParametersJSONParseParameters) {
		rulesV1JSONParseParametersModel := new(logsv0.RulesV1JSONParseParameters)
		rulesV1JSONParseParametersModel.DestinationField = core.StringPtr("json.content")
		rulesV1JSONParseParametersModel.DeleteSource = core.BoolPtr(true)
		rulesV1JSONParseParametersModel.OverrideDest = core.BoolPtr(true)

		model := new(logsv0.RulesV1RuleParametersRuleParametersJSONParseParameters)
		model.JSONParseParameters = rulesV1JSONParseParametersModel

		assert.Equal(t, result, model)
	}

	rulesV1JSONParseParametersModel := make(map[string]interface{})
	rulesV1JSONParseParametersModel["destination_field"] = "json.content"
	rulesV1JSONParseParametersModel["delete_source"] = true
	rulesV1JSONParseParametersModel["override_dest"] = true

	model := make(map[string]interface{})
	model["json_parse_parameters"] = []interface{}{rulesV1JSONParseParametersModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleParametersRuleParametersJSONParseParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsRuleGroupMapToRulesV1RuleMatcher(t *testing.T) {
// 	checkResult := func(result logsv0.RulesV1RuleMatcherIntf) {
// 		rulesV1ApplicationNameConstraintModel := new(logsv0.RulesV1ApplicationNameConstraint)
// 		rulesV1ApplicationNameConstraintModel.Value = core.StringPtr("testString")

// 		model := new(logsv0.RulesV1RuleMatcher)
// 		model.ApplicationName = rulesV1ApplicationNameConstraintModel
// 		model.SubsystemName = rulesV1SubsystemNameConstraintModel
// 		model.Severity = rulesV1SeverityConstraintModel

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ApplicationNameConstraintModel := make(map[string]interface{})
// 	rulesV1ApplicationNameConstraintModel["value"] = "testString"

// 	model := make(map[string]interface{})
// 	model["application_name"] = []interface{}{rulesV1ApplicationNameConstraintModel}
// 	model["subsystem_name"] = []interface{}{rulesV1SubsystemNameConstraintModel}
// 	model["severity"] = []interface{}{rulesV1SeverityConstraintModel}

// 	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleMatcher(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsRuleGroupMapToRulesV1ApplicationNameConstraint(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1ApplicationNameConstraint) {
		model := new(logsv0.RulesV1ApplicationNameConstraint)
		model.Value = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["value"] = "testString"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1ApplicationNameConstraint(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1SubsystemNameConstraint(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1SubsystemNameConstraint) {
		model := new(logsv0.RulesV1SubsystemNameConstraint)
		model.Value = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["value"] = "testString"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1SubsystemNameConstraint(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1SeverityConstraint(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1SeverityConstraint) {
		model := new(logsv0.RulesV1SeverityConstraint)
		model.Value = core.StringPtr("debug_or_unspecified")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["value"] = "debug_or_unspecified"

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1SeverityConstraint(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleMatcherConstraintApplicationName(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleMatcherConstraintApplicationName) {
		rulesV1ApplicationNameConstraintModel := new(logsv0.RulesV1ApplicationNameConstraint)
		rulesV1ApplicationNameConstraintModel.Value = core.StringPtr("testString")

		model := new(logsv0.RulesV1RuleMatcherConstraintApplicationName)
		model.ApplicationName = rulesV1ApplicationNameConstraintModel

		assert.Equal(t, result, model)
	}

	rulesV1ApplicationNameConstraintModel := make(map[string]interface{})
	rulesV1ApplicationNameConstraintModel["value"] = "testString"

	model := make(map[string]interface{})
	model["application_name"] = []interface{}{rulesV1ApplicationNameConstraintModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleMatcherConstraintApplicationName(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleMatcherConstraintSubsystemName(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleMatcherConstraintSubsystemName) {
		rulesV1SubsystemNameConstraintModel := new(logsv0.RulesV1SubsystemNameConstraint)
		rulesV1SubsystemNameConstraintModel.Value = core.StringPtr("testString")

		model := new(logsv0.RulesV1RuleMatcherConstraintSubsystemName)
		model.SubsystemName = rulesV1SubsystemNameConstraintModel

		assert.Equal(t, result, model)
	}

	rulesV1SubsystemNameConstraintModel := make(map[string]interface{})
	rulesV1SubsystemNameConstraintModel["value"] = "testString"

	model := make(map[string]interface{})
	model["subsystem_name"] = []interface{}{rulesV1SubsystemNameConstraintModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleMatcherConstraintSubsystemName(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsRuleGroupMapToRulesV1RuleMatcherConstraintSeverity(t *testing.T) {
	checkResult := func(result *logsv0.RulesV1RuleMatcherConstraintSeverity) {
		rulesV1SeverityConstraintModel := new(logsv0.RulesV1SeverityConstraint)
		rulesV1SeverityConstraintModel.Value = core.StringPtr("debug_or_unspecified")

		model := new(logsv0.RulesV1RuleMatcherConstraintSeverity)
		model.Severity = rulesV1SeverityConstraintModel

		assert.Equal(t, result, model)
	}

	rulesV1SeverityConstraintModel := make(map[string]interface{})
	rulesV1SeverityConstraintModel["value"] = "debug_or_unspecified"

	model := make(map[string]interface{})
	model["severity"] = []interface{}{rulesV1SeverityConstraintModel}

	result, err := logs.ResourceIbmLogsRuleGroupMapToRulesV1RuleMatcherConstraintSeverity(model)
	assert.Nil(t, err)
	checkResult(result)
}
