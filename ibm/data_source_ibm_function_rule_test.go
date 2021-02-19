/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccFunctionRuleDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	actionName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	triggerName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
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
		  kind = "nodejs:6"
		  code = file("test-fixtures/hellonode.js")
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
