// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package functions_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFunctionRuleDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_rule_%d", acctest.RandIntRange(10, 100))
	actionName := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	triggerName := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckFunctionRuleDataSource(actionName, triggerName, name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "trigger_name", triggerName),
					resource.TestCheckResourceAttr("data.ibm_function_rule.datarule", "name", name),
				),
			},
		},
	})
}

func testAccCheckFunctionRuleDataSource(actionName, triggerName, name, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_action" "action" {
		name      = "%s"
		namespace = "%s"
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonode.js")
		}
	  }
	  
	  resource "ibm_function_trigger" "trigger" {
		name      = "%s"
		namespace = "%s"
		feed {
		  name       = "/whisk.system/alarms/alarm"
		  parameters = <<EOF
							  [
									  {
											  "key":"cron",
											  "value":"0 */2 * * *"
									  }
							  ]
	  
	  EOF
	  
		}
	  
		user_defined_annotations = <<EOF
	   [
	  {
	   "key":"sample trigger",
	   "value":"Trigger for hello action"
	  }
	   ]
	  
	  EOF
	  
	  }
	  
	  resource "ibm_function_rule" "rule" {
		name         = "%s"
		namespace    = "%s"
		trigger_name = ibm_function_trigger.trigger.name
		action_name  = ibm_function_action.action.name
	  }
	  
	  data "ibm_function_rule" "datarule" {
		name      = ibm_function_rule.rule.name
		namespace = ibm_function_rule.rule.namespace
	  }
	  
`, actionName, namespace, triggerName, namespace, name, namespace)

}
