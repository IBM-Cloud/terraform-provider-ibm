// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	// . "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
)

func TestAccIbmLogsRuleGroupDataSourceBasic(t *testing.T) {
	ruleGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupDataSourceConfigBasic(ruleGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "rule_subgroups.#"),
				),
			},
		},
	})
}

func TestAccIbmLogsRuleGroupDataSourceAllArgs(t *testing.T) {
	ruleGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	ruleGroupDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	ruleGroupEnabled := "true"
	ruleGroupOrder := fmt.Sprintf("%d", acctest.RandIntRange(0, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupDataSourceConfig(ruleGroupName, ruleGroupDescription, ruleGroupEnabled, ruleGroupOrder),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "rule_matchers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "rule_subgroups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "rule_subgroups.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_rule_group.logs_rule_group_instance", "rule_subgroups.0.enabled", ruleGroupEnabled),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_group.logs_rule_group_instance", "rule_subgroups.0.order"),
					resource.TestCheckResourceAttr("data.ibm_logs_rule_group.logs_rule_group_instance", "order", ruleGroupOrder),
				),
			},
		},
	})
}

func testAccCheckIbmLogsRuleGroupDataSourceConfigBasic(ruleGroupName string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_rule_group" "logs_rule_group_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			description = "test description"
			enabled     = true
			rule_matchers {
			subsystem_name {
				value = "mysql"
			}
			}
			rule_subgroups {
			rules {
				name         = "mysql-parse"
				source_field = "text"
				parameters {
				parse_parameters {
					destination_field = "text"
					rule              = "(?P<timestamp>[^,]+),(?P<hostname>[^,]+),(?P<username>[^,]+),(?P<ip>[^,]+),(?P<connectionId>[0-9]+),(?P<queryId>[0-9]+),(?P<operation>[^,]+),(?P<database>[^,]+),'?(?P<object>.*)'?,(?P<returnCode>[0-9]+)"
				}
				}
				enabled = true
				order   = 1
			}
		
			enabled = true
			order   = 1
			}
			order = 1
		}

		data "ibm_logs_rule_group" "logs_rule_group_instance" {
			instance_id = ibm_logs_rule_group.logs_rule_group_instance.instance_id
			region 		= ibm_logs_rule_group.logs_rule_group_instance.region
			group_id 	= ibm_logs_rule_group.logs_rule_group_instance.rule_group_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, ruleGroupName)
}

func testAccCheckIbmLogsRuleGroupDataSourceConfig(ruleGroupName string, ruleGroupDescription string, ruleGroupEnabled string, ruleGroupOrder string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_rule_group" "logs_rule_group_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			description = "%s"
			enabled     = %s
			rule_matchers {
			subsystem_name {
				value = "mysql"
			}
			}
			rule_subgroups {
			rules {
				name         = "mysql-parse"
				source_field = "text"
				parameters {
				parse_parameters {
					destination_field = "text"
					rule              = "(?P<timestamp>[^,]+),(?P<hostname>[^,]+),(?P<username>[^,]+),(?P<ip>[^,]+),(?P<connectionId>[0-9]+),(?P<queryId>[0-9]+),(?P<operation>[^,]+),(?P<database>[^,]+),'?(?P<object>.*)'?,(?P<returnCode>[0-9]+)"
				}
				}
				enabled = true
				order   = 1
			}
		
			enabled = true
			order   = 1
			}
			order = %s
		}

		data "ibm_logs_rule_group" "logs_rule_group_instance" {
			instance_id = ibm_logs_rule_group.logs_rule_group_instance.instance_id
			region 		= ibm_logs_rule_group.logs_rule_group_instance.region
			group_id 	= ibm_logs_rule_group.logs_rule_group_instance.rule_group_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, ruleGroupName, ruleGroupDescription, ruleGroupEnabled, ruleGroupOrder)
}

// Todo @kavya498: Fix unit testcases
// func TestDataSourceIbmLogsRuleGroupRulesV1RuleMatcherToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleMatcherToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1ApplicationNameConstraintToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["value"] = "testString"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1ApplicationNameConstraint)
// 	model.Value = core.StringPtr("testString")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1ApplicationNameConstraintToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1SubsystemNameConstraintToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["value"] = "testString"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1SubsystemNameConstraint)
// 	model.Value = core.StringPtr("testString")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1SubsystemNameConstraintToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1SeverityConstraintToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["value"] = "debug_or_unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1SeverityConstraint)
// 	model.Value = core.StringPtr("debug_or_unspecified")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1SeverityConstraintToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintApplicationNameToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1ApplicationNameConstraintModel := make(map[string]interface{})
// 		rulesV1ApplicationNameConstraintModel["value"] = "testString"

// 		model := make(map[string]interface{})
// 		model["application_name"] = []map[string]interface{}{rulesV1ApplicationNameConstraintModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ApplicationNameConstraintModel := new(logsv0.RulesV1ApplicationNameConstraint)
// 	rulesV1ApplicationNameConstraintModel.Value = core.StringPtr("testString")

// 	model := new(logsv0.RulesV1RuleMatcherConstraintApplicationName)
// 	model.ApplicationName = rulesV1ApplicationNameConstraintModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintApplicationNameToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintSubsystemNameToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1SubsystemNameConstraintModel := make(map[string]interface{})
// 		rulesV1SubsystemNameConstraintModel["value"] = "testString"

// 		model := make(map[string]interface{})
// 		model["subsystem_name"] = []map[string]interface{}{rulesV1SubsystemNameConstraintModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1SubsystemNameConstraintModel := new(logsv0.RulesV1SubsystemNameConstraint)
// 	rulesV1SubsystemNameConstraintModel.Value = core.StringPtr("testString")

// 	model := new(logsv0.RulesV1RuleMatcherConstraintSubsystemName)
// 	model.SubsystemName = rulesV1SubsystemNameConstraintModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintSubsystemNameToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintSeverityToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1SeverityConstraintModel := make(map[string]interface{})
// 		rulesV1SeverityConstraintModel["value"] = "debug_or_unspecified"

// 		model := make(map[string]interface{})
// 		model["severity"] = []map[string]interface{}{rulesV1SeverityConstraintModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1SeverityConstraintModel := new(logsv0.RulesV1SeverityConstraint)
// 	rulesV1SeverityConstraintModel.Value = core.StringPtr("debug_or_unspecified")

// 	model := new(logsv0.RulesV1RuleMatcherConstraintSeverity)
// 	model.Severity = rulesV1SeverityConstraintModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleMatcherConstraintSeverityToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleSubgroupToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1ExtractParametersModel := make(map[string]interface{})
// 		rulesV1ExtractParametersModel["rule"] = "testString"

// 		rulesV1RuleParametersModel := make(map[string]interface{})
// 		rulesV1RuleParametersModel["extract_parameters"] = []map[string]interface{}{rulesV1ExtractParametersModel}

// 		rulesV1RuleModel := make(map[string]interface{})
// 		rulesV1RuleModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		rulesV1RuleModel["name"] = "testString"
// 		rulesV1RuleModel["description"] = "testString"
// 		rulesV1RuleModel["source_field"] = "logObj.source"
// 		rulesV1RuleModel["parameters"] = []map[string]interface{}{rulesV1RuleParametersModel}
// 		rulesV1RuleModel["enabled"] = true
// 		rulesV1RuleModel["order"] = int(0)

// 		model := make(map[string]interface{})
// 		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		model["rules"] = []map[string]interface{}{rulesV1RuleModel}
// 		model["enabled"] = true
// 		model["order"] = int(0)

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
// 	rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

// 	rulesV1RuleParametersModel := new(logsv0.RulesV1RuleParametersRuleParametersExtractParameters)
// 	rulesV1RuleParametersModel.ExtractParameters = rulesV1ExtractParametersModel

// 	rulesV1RuleModel := new(logsv0.RulesV1Rule)
// 	rulesV1RuleModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	rulesV1RuleModel.Name = core.StringPtr("testString")
// 	rulesV1RuleModel.Description = core.StringPtr("testString")
// 	rulesV1RuleModel.SourceField = core.StringPtr("logObj.source")
// 	rulesV1RuleModel.Parameters = rulesV1RuleParametersModel
// 	rulesV1RuleModel.Enabled = core.BoolPtr(true)
// 	rulesV1RuleModel.Order = core.Int64Ptr(int64(0))

// 	model := new(logsv0.RulesV1RuleSubgroup)
// 	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	model.Rules = []logsv0.RulesV1Rule{*rulesV1RuleModel}
// 	model.Enabled = core.BoolPtr(true)
// 	model.Order = core.Int64Ptr(int64(0))

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleSubgroupToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1ExtractParametersModel := make(map[string]interface{})
// 		rulesV1ExtractParametersModel["rule"] = "testString"

// 		rulesV1RuleParametersModel := make(map[string]interface{})
// 		rulesV1RuleParametersModel["extract_parameters"] = []map[string]interface{}{rulesV1ExtractParametersModel}

// 		model := make(map[string]interface{})
// 		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
// 		model["name"] = "testString"
// 		model["description"] = "testString"
// 		model["source_field"] = "logObj.source"
// 		model["parameters"] = []map[string]interface{}{rulesV1RuleParametersModel}
// 		model["enabled"] = true
// 		model["order"] = int(0)

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
// 	rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

// 	rulesV1RuleParametersModel := new(logsv0.RulesV1RuleParametersRuleParametersExtractParameters)
// 	rulesV1RuleParametersModel.ExtractParameters = rulesV1ExtractParametersModel

// 	model := new(logsv0.RulesV1Rule)
// 	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
// 	model.Name = core.StringPtr("testString")
// 	model.Description = core.StringPtr("testString")
// 	model.SourceField = core.StringPtr("logObj.source")
// 	model.Parameters = rulesV1RuleParametersModel
// 	model.Enabled = core.BoolPtr(true)
// 	model.Order = core.Int64Ptr(int64(0))

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1ExtractParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["rule"] = "testString"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1ExtractParameters)
// 	model.Rule = core.StringPtr("testString")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1ExtractParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1JSONExtractParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["destination_field"] = "category_or_unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1JSONExtractParameters)
// 	model.DestinationField = core.StringPtr("category_or_unspecified")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1JSONExtractParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1ReplaceParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["destination_field"] = "text.message"
// 		model["replace_new_val"] = "***"
// 		model["rule"] = "the password is (?P<password>[A-Za-z0-9!@#$].)"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1ReplaceParameters)
// 	model.DestinationField = core.StringPtr("text.message")
// 	model.ReplaceNewVal = core.StringPtr("***")
// 	model.Rule = core.StringPtr("the password is (?P<password>[A-Za-z0-9!@#$].)")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1ReplaceParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1ParseParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["destination_field"] = "text.message"
// 		model["rule"] = "testString"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1ParseParameters)
// 	model.DestinationField = core.StringPtr("text.message")
// 	model.Rule = core.StringPtr("testString")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1ParseParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1AllowParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["keep_blocked_logs"] = true
// 		model["rule"] = "^this log should be kept!!!.*$"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1AllowParameters)
// 	model.KeepBlockedLogs = core.BoolPtr(true)
// 	model.Rule = core.StringPtr("^this log should be kept!!!.*$")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1AllowParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1BlockParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["keep_blocked_logs"] = true
// 		model["rule"] = "^this log should be blocked!!!.*$"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1BlockParameters)
// 	model.KeepBlockedLogs = core.BoolPtr(true)
// 	model.Rule = core.StringPtr("^this log should be blocked!!!.*$")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1BlockParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1ExtractTimestampParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["standard"] = "strftime_or_unspecified"
// 		model["format"] = "%Y-%m-%ddT%H:%M:%S.%f%z"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1ExtractTimestampParameters)
// 	model.Standard = core.StringPtr("strftime_or_unspecified")
// 	model.Format = core.StringPtr("%Y-%m-%ddT%H:%M:%S.%f%z")

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1ExtractTimestampParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RemoveFieldsParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["fields"] = []string{"testString"}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1RemoveFieldsParameters)
// 	model.Fields = []string{"testString"}

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RemoveFieldsParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1JSONStringifyParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["destination_field"] = "json.stringified"
// 		model["delete_source"] = true

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1JSONStringifyParameters)
// 	model.DestinationField = core.StringPtr("json.stringified")
// 	model.DeleteSource = core.BoolPtr(true)

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1JSONStringifyParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1JSONParseParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["destination_field"] = "json.content"
// 		model["delete_source"] = true
// 		model["override_dest"] = true

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.RulesV1JSONParseParameters)
// 	model.DestinationField = core.StringPtr("json.content")
// 	model.DeleteSource = core.BoolPtr(true)
// 	model.OverrideDest = core.BoolPtr(true)

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1JSONParseParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersExtractParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1ExtractParametersModel := make(map[string]interface{})
// 		rulesV1ExtractParametersModel["rule"] = "testString"

// 		model := make(map[string]interface{})
// 		model["extract_parameters"] = []map[string]interface{}{rulesV1ExtractParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ExtractParametersModel := new(logsv0.RulesV1ExtractParameters)
// 	rulesV1ExtractParametersModel.Rule = core.StringPtr("testString")

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersExtractParameters)
// 	model.ExtractParameters = rulesV1ExtractParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersExtractParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONExtractParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1JSONExtractParametersModel := make(map[string]interface{})
// 		rulesV1JSONExtractParametersModel["destination_field"] = "category_or_unspecified"

// 		model := make(map[string]interface{})
// 		model["json_extract_parameters"] = []map[string]interface{}{rulesV1JSONExtractParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1JSONExtractParametersModel := new(logsv0.RulesV1JSONExtractParameters)
// 	rulesV1JSONExtractParametersModel.DestinationField = core.StringPtr("category_or_unspecified")

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersJSONExtractParameters)
// 	model.JSONExtractParameters = rulesV1JSONExtractParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONExtractParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersReplaceParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1ReplaceParametersModel := make(map[string]interface{})
// 		rulesV1ReplaceParametersModel["destination_field"] = "text.message"
// 		rulesV1ReplaceParametersModel["replace_new_val"] = "***"
// 		rulesV1ReplaceParametersModel["rule"] = "the password is (?P<password>[A-Za-z0-9!@#$].)"

// 		model := make(map[string]interface{})
// 		model["replace_parameters"] = []map[string]interface{}{rulesV1ReplaceParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ReplaceParametersModel := new(logsv0.RulesV1ReplaceParameters)
// 	rulesV1ReplaceParametersModel.DestinationField = core.StringPtr("text.message")
// 	rulesV1ReplaceParametersModel.ReplaceNewVal = core.StringPtr("***")
// 	rulesV1ReplaceParametersModel.Rule = core.StringPtr("the password is (?P<password>[A-Za-z0-9!@#$].)")

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersReplaceParameters)
// 	model.ReplaceParameters = rulesV1ReplaceParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersReplaceParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersParseParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1ParseParametersModel := make(map[string]interface{})
// 		rulesV1ParseParametersModel["destination_field"] = "text.message"
// 		rulesV1ParseParametersModel["rule"] = "testString"

// 		model := make(map[string]interface{})
// 		model["parse_parameters"] = []map[string]interface{}{rulesV1ParseParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ParseParametersModel := new(logsv0.RulesV1ParseParameters)
// 	rulesV1ParseParametersModel.DestinationField = core.StringPtr("text.message")
// 	rulesV1ParseParametersModel.Rule = core.StringPtr("testString")

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersParseParameters)
// 	model.ParseParameters = rulesV1ParseParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersParseParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersAllowParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1AllowParametersModel := make(map[string]interface{})
// 		rulesV1AllowParametersModel["keep_blocked_logs"] = true
// 		rulesV1AllowParametersModel["rule"] = "^this log should be kept!!!.*$"

// 		model := make(map[string]interface{})
// 		model["allow_parameters"] = []map[string]interface{}{rulesV1AllowParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1AllowParametersModel := new(logsv0.RulesV1AllowParameters)
// 	rulesV1AllowParametersModel.KeepBlockedLogs = core.BoolPtr(true)
// 	rulesV1AllowParametersModel.Rule = core.StringPtr("^this log should be kept!!!.*$")

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersAllowParameters)
// 	model.AllowParameters = rulesV1AllowParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersAllowParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersBlockParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1BlockParametersModel := make(map[string]interface{})
// 		rulesV1BlockParametersModel["keep_blocked_logs"] = true
// 		rulesV1BlockParametersModel["rule"] = "^this log should be blocked!!!.*$"

// 		model := make(map[string]interface{})
// 		model["block_parameters"] = []map[string]interface{}{rulesV1BlockParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1BlockParametersModel := new(logsv0.RulesV1BlockParameters)
// 	rulesV1BlockParametersModel.KeepBlockedLogs = core.BoolPtr(true)
// 	rulesV1BlockParametersModel.Rule = core.StringPtr("^this log should be blocked!!!.*$")

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersBlockParameters)
// 	model.BlockParameters = rulesV1BlockParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersBlockParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersExtractTimestampParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1ExtractTimestampParametersModel := make(map[string]interface{})
// 		rulesV1ExtractTimestampParametersModel["standard"] = "strftime_or_unspecified"
// 		rulesV1ExtractTimestampParametersModel["format"] = "%Y-%m-%ddT%H:%M:%S.%f%z"

// 		model := make(map[string]interface{})
// 		model["extract_timestamp_parameters"] = []map[string]interface{}{rulesV1ExtractTimestampParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1ExtractTimestampParametersModel := new(logsv0.RulesV1ExtractTimestampParameters)
// 	rulesV1ExtractTimestampParametersModel.Standard = core.StringPtr("strftime_or_unspecified")
// 	rulesV1ExtractTimestampParametersModel.Format = core.StringPtr("%Y-%m-%ddT%H:%M:%S.%f%z")

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersExtractTimestampParameters)
// 	model.ExtractTimestampParameters = rulesV1ExtractTimestampParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersExtractTimestampParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersRemoveFieldsParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1RemoveFieldsParametersModel := make(map[string]interface{})
// 		rulesV1RemoveFieldsParametersModel["fields"] = []string{"testString"}

// 		model := make(map[string]interface{})
// 		model["remove_fields_parameters"] = []map[string]interface{}{rulesV1RemoveFieldsParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1RemoveFieldsParametersModel := new(logsv0.RulesV1RemoveFieldsParameters)
// 	rulesV1RemoveFieldsParametersModel.Fields = []string{"testString"}

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersRemoveFieldsParameters)
// 	model.RemoveFieldsParameters = rulesV1RemoveFieldsParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersRemoveFieldsParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONStringifyParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1JSONStringifyParametersModel := make(map[string]interface{})
// 		rulesV1JSONStringifyParametersModel["destination_field"] = "json.stringified"
// 		rulesV1JSONStringifyParametersModel["delete_source"] = true

// 		model := make(map[string]interface{})
// 		model["json_stringify_parameters"] = []map[string]interface{}{rulesV1JSONStringifyParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1JSONStringifyParametersModel := new(logsv0.RulesV1JSONStringifyParameters)
// 	rulesV1JSONStringifyParametersModel.DestinationField = core.StringPtr("json.stringified")
// 	rulesV1JSONStringifyParametersModel.DeleteSource = core.BoolPtr(true)

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersJSONStringifyParameters)
// 	model.JSONStringifyParameters = rulesV1JSONStringifyParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONStringifyParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONParseParametersToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		rulesV1JSONParseParametersModel := make(map[string]interface{})
// 		rulesV1JSONParseParametersModel["destination_field"] = "json.content"
// 		rulesV1JSONParseParametersModel["delete_source"] = true
// 		rulesV1JSONParseParametersModel["override_dest"] = true

// 		model := make(map[string]interface{})
// 		model["json_parse_parameters"] = []map[string]interface{}{rulesV1JSONParseParametersModel}

// 		assert.Equal(t, result, model)
// 	}

// 	rulesV1JSONParseParametersModel := new(logsv0.RulesV1JSONParseParameters)
// 	rulesV1JSONParseParametersModel.DestinationField = core.StringPtr("json.content")
// 	rulesV1JSONParseParametersModel.DeleteSource = core.BoolPtr(true)
// 	rulesV1JSONParseParametersModel.OverrideDest = core.BoolPtr(true)

// 	model := new(logsv0.RulesV1RuleParametersRuleParametersJSONParseParameters)
// 	model.JSONParseParameters = rulesV1JSONParseParametersModel

// 	result, err := logs.DataSourceIbmLogsRuleGroupRulesV1RuleParametersRuleParametersJSONParseParametersToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }
