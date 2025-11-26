// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsDataAccessRuleBasic(t *testing.T) {
	var conf logsv0.DataAccessRule
	displayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	defaultExpression := "true"
	displayNameUpdate := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	defaultExpressionUpdate := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsDataAccessRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRuleConfigBasic(displayName, defaultExpression),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsDataAccessRuleExists("ibm_logs_data_access_rule.logs_data_access_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "display_name", displayName),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "default_expression", defaultExpression),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRuleConfigBasic(displayNameUpdate, defaultExpressionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "display_name", displayNameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "default_expression", defaultExpressionUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsDataAccessRuleAllArgs(t *testing.T) {
	var conf logsv0.DataAccessRule
	displayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	defaultExpression := "<v1>true"
	displayNameUpdate := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	defaultExpressionUpdate := "<v1>false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsDataAccessRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRuleConfig(displayName, description, defaultExpression),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsDataAccessRuleExists("ibm_logs_data_access_rule.logs_data_access_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "display_name", displayName),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "default_expression", defaultExpression),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRuleConfig(displayNameUpdate, descriptionUpdate, defaultExpressionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "display_name", displayNameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "default_expression", defaultExpressionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_data_access_rule.logs_data_access_rule_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsDataAccessRuleConfigBasic(displayName string, defaultExpression string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_data_access_rule" "logs_data_access_rule_instance" {
			instance_id  = "%s"
			region       = "%s"
			display_name = "%s"
			filters {
				entity_type = "logs"
				expression  = "<v1> foo == 'bar'"
			}
			default_expression = "%s"
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, displayName, defaultExpression)
}

func testAccCheckIbmLogsDataAccessRuleConfig(displayName string, description string, defaultExpression string) string {
	return fmt.Sprintf(`

		resource "ibm_logs_data_access_rule" "logs_data_access_rule_instance" {
			instance_id  = "%s"
			region       = "%s"
			display_name = "%s"
			description  = "%s"
			filters {
				entity_type = "logs"
				expression  = "<v1> foo == 'bar'"
			}
			default_expression = "%s"
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, displayName, description, defaultExpression)
}

func testAccCheckIbmLogsDataAccessRuleExists(n string, obj logsv0.DataAccessRule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		listDataAccessRulesOptions := &logsv0.ListDataAccessRulesOptions{}
		accessRuleID := core.UUIDPtr(strfmt.UUID(resourceID[2]))

		listDataAccessRulesOptions.ID = []strfmt.UUID{*accessRuleID}
		dataAccessRuleCollection, _, err := logsClient.ListDataAccessRules(listDataAccessRulesOptions)
		if err != nil {
			return err
		}
		if dataAccessRuleCollection != nil && len(dataAccessRuleCollection.DataAccessRules) > 0 && &dataAccessRuleCollection.DataAccessRules[0] != nil && dataAccessRuleCollection.DataAccessRules[0].ID == accessRuleID {
			dataAccessRule := dataAccessRuleCollection.DataAccessRules[0]
			obj = dataAccessRule
			return nil
		}

		return nil
	}
}

func testAccCheckIbmLogsDataAccessRuleDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_data_access_rule" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		listDataAccessRulesOptions := &logsv0.ListDataAccessRulesOptions{}
		accessRuleID := core.UUIDPtr(strfmt.UUID(resourceID[2]))

		listDataAccessRulesOptions.ID = []strfmt.UUID{*accessRuleID}

		dataAccessRuleCollection, _, err := logsClient.ListDataAccessRules(listDataAccessRulesOptions)

		if dataAccessRuleCollection != nil && len(dataAccessRuleCollection.DataAccessRules) > 0 && &dataAccessRuleCollection.DataAccessRules[0] != nil && dataAccessRuleCollection.DataAccessRules[0].ID == accessRuleID {
			return fmt.Errorf("logs_data_access_rule still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
