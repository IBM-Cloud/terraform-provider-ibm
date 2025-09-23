// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	// . "github.com/Mavrickk3/terraform-provider-ibm/ibm/unittest"
)

func TestAccIbmLogsRuleGroupsDataSourceBasic(t *testing.T) {
	ruleGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupsDataSourceConfigBasic(ruleGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_groups.logs_rule_groups_instance", "id"),
				),
			},
		},
	})
}

func TestAccIbmLogsRuleGroupsDataSourceAllArgs(t *testing.T) {
	ruleGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	ruleGroupDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	ruleGroupEnabled := "false"
	ruleGroupOrder := fmt.Sprintf("%d", acctest.RandIntRange(0, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupsDataSourceConfig(ruleGroupName, ruleGroupDescription, ruleGroupEnabled, ruleGroupOrder),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_groups.logs_rule_groups_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_groups.logs_rule_groups_instance", "rulegroups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_rule_groups.logs_rule_groups_instance", "rulegroups.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_rule_groups.logs_rule_groups_instance", "rulegroups.0.name", ruleGroupName),
					resource.TestCheckResourceAttr("data.ibm_logs_rule_groups.logs_rule_groups_instance", "rulegroups.0.description", ruleGroupDescription),
					resource.TestCheckResourceAttr("data.ibm_logs_rule_groups.logs_rule_groups_instance", "rulegroups.0.enabled", ruleGroupEnabled),
					resource.TestCheckResourceAttr("data.ibm_logs_rule_groups.logs_rule_groups_instance", "rulegroups.0.order", ruleGroupOrder),
				),
			},
		},
	})
}

func testAccCheckIbmLogsRuleGroupsDataSourceConfigBasic(ruleGroupName string) string {
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

		data "ibm_logs_rule_groups" "logs_rule_groups_instance" {
			depends_on = [
				ibm_logs_rule_group.logs_rule_group_instance
			]
			instance_id = ibm_logs_rule_group.logs_rule_group_instance.instance_id
			region 		= ibm_logs_rule_group.logs_rule_group_instance.region
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, ruleGroupName)
}

func testAccCheckIbmLogsRuleGroupsDataSourceConfig(ruleGroupName string, ruleGroupDescription string, ruleGroupEnabled string, ruleGroupOrder string) string {
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

		data "ibm_logs_rule_groups" "logs_rule_groups_instance" {
			depends_on  = [
				ibm_logs_rule_group.logs_rule_group_instance
			]
			instance_id = ibm_logs_rule_group.logs_rule_group_instance.instance_id
			region 		= ibm_logs_rule_group.logs_rule_group_instance.region
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, ruleGroupName, ruleGroupDescription, ruleGroupEnabled, ruleGroupOrder)
}
