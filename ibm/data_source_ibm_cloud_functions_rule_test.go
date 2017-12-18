package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudFunctionsRuleDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	actionName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	triggerName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsRuleDataSource(actionName, triggerName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_rule.rule", "trigger_name", triggerName),
					resource.TestCheckResourceAttr("data.ibm_cloud_functions_rule.datarule", "name", name),
				),
			},
		},
	})
}

func testAccCheckCloudFunctionsRuleDataSource(actionName, triggerName, name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_action" "action" {
	name = "%s"		  
	exec = {
	  kind = "nodejs:6"
	  code = "${file("test-fixtures/hellonode.js")}"
	}
  }
  resource "ibm_cloud_functions_trigger" "trigger" {
	name = "%s"
	feed = [
		{
			  name = "/whisk.system/alarms/alarm"
			  parameters = <<EOF
			[
				{
					"key":"cron",
					"value":"0 */2 * * *"
				}
			]
		EOF
	 },
 ]

 user_defined_annotations = <<EOF
 [
{
 "key":"sample trigger",
 "value":"Trigger for hello action"
}
 ]
 EOF
}
resource "ibm_cloud_functions_rule" "rule" {
name = "%s"
trigger_name = "${ibm_cloud_functions_trigger.trigger.name}"
action_name = "${ibm_cloud_functions_action.action.name}"

}
data "ibm_cloud_functions_rule" "datarule" {
	name = "${ibm_cloud_functions_rule.rule.name}"

}
`, actionName, triggerName, name)

}
