// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.91.0-d9755c53-20240605-153412
 */

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsDataAccessRulesDataSourceBasic(t *testing.T) {
	dataAccessRuleDisplayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	dataAccessRuleDefaultExpression := "<v1>true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRulesDataSourceConfigBasic(dataAccessRuleDisplayName, dataAccessRuleDefaultExpression),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.#"),
				),
			},
		},
	})
}

func TestAccIbmLogsDataAccessRulesDataSourceAllArgs(t *testing.T) {
	dataAccessRuleDisplayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	dataAccessRuleDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	dataAccessRuleDefaultExpression := "<v1>true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRulesDataSourceConfig(dataAccessRuleDisplayName, dataAccessRuleDescription, dataAccessRuleDefaultExpression),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "logs_data_access_rules_id.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.display_name", dataAccessRuleDisplayName),
					resource.TestCheckResourceAttr("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.description", dataAccessRuleDescription),
					resource.TestCheckResourceAttr("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.default_expression", dataAccessRuleDefaultExpression),
				),
			},
		},
	})
}

func testAccCheckIbmLogsDataAccessRulesDataSourceConfigBasic(dataAccessRuleDisplayName string, dataAccessRuleDefaultExpression string) string {
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
 
		 data "ibm_logs_data_access_rules" "logs_data_access_rules_instance" {
			 instance_id 			  = ibm_logs_data_access_rule.logs_data_access_rule_instance.instance_id
			 region      			  = ibm_logs_data_access_rule.logs_data_access_rule_instance.region
		 }
	 `, acc.LogsInstanceId, acc.LogsInstanceRegion, dataAccessRuleDisplayName, dataAccessRuleDefaultExpression)
}

func testAccCheckIbmLogsDataAccessRulesDataSourceConfig(dataAccessRuleDisplayName string, dataAccessRuleDescription string, dataAccessRuleDefaultExpression string) string {
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
 
		 data "ibm_logs_data_access_rules" "logs_data_access_rules_instance" {
			 instance_id 			  = ibm_logs_data_access_rule.logs_data_access_rule_instance.instance_id
			 region      			  = ibm_logs_data_access_rule.logs_data_access_rule_instance.region
			 logs_data_access_rules_id = [ibm_logs_data_access_rule.logs_data_access_rule_instance.access_rule_id]
		 }
	 `, acc.LogsInstanceId, acc.LogsInstanceRegion, dataAccessRuleDisplayName, dataAccessRuleDescription, dataAccessRuleDefaultExpression)
}
