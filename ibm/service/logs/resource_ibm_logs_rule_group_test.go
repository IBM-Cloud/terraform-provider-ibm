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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsRuleGroupBasic(t *testing.T) {
	var conf logsv0.RuleGroup
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
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
	enabled := "false"
	order := fmt.Sprintf("%d", acctest.RandIntRange(0, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enabledUpdate := "true"
	orderUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsRuleGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupConfig(name, description, enabled, order),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsRuleGroupExists("ibm_logs_rule_group.logs_rule_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "order", order),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsRuleGroupConfig(nameUpdate, descriptionUpdate, enabledUpdate, orderUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_logs_rule_group.logs_rule_group_instance", "order", orderUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_rule_group.logs_rule_group_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsRuleGroupConfigBasic(name string) string {
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name)
}

func testAccCheckIbmLogsRuleGroupConfig(name string, description string, enabled string, order string) string {
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, description, enabled, order)
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
