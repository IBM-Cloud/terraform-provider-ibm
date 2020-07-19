package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/apache/openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
)

func TestAccFunctionRule_Basic(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	actionName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	triggerName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updatedTriggerName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionRuleDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionRuleCreate(actionName, triggerName, name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionRuleExists("ibm_function_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "trigger_name", triggerName),
				),
			},
			resource.TestStep{
				Config: testAccCheckFunctionRuleUpdate(updatedTriggerName, name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionRuleExists("ibm_function_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "trigger_name", updatedTriggerName),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "action_name", "/whisk.system/cloudant/delete-attachment"),
				),
			},
		},
	})
}

func TestAccFunctionRule_Import(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	triggeName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionRuleDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionRuleImport(triggeName, name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionRuleExists("ibm_function_rule.import", &conf),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "name", name),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "publish", "false"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_function_rule.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckFunctionRuleExists(n string, obj *whisk.Rule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		parts, err := cfIdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		namespace := parts[0]
		name := parts[1]

		client, err := testAccProvider.Meta().(ClientSession).FunctionClient()
		if err != nil {
			return err
		}

		bxSession, err := testAccProvider.Meta().(ClientSession).BluemixSession()
		if err != nil {
			return err
		}
		client, err = setupOpenWhiskClientConfig(namespace, bxSession.Config, client)
		if err != nil {
			return err

		}

		rule, _, err := client.Rules.Get(name)
		if err != nil {
			return err
		}

		*obj = *rule
		return nil
	}
}

func testAccCheckFunctionRuleDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).FunctionClient()
	if err != nil {
		return err
	}

	bxSession, err := testAccProvider.Meta().(ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_rule" {
			continue
		}

		parts, err := cfIdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		namespace := parts[0]
		name := parts[1]

		client, err = setupOpenWhiskClientConfig(namespace, bxSession.Config, client)
		if err != nil {
			return err

		}

		_, _, err = client.Rules.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for IBM Cloud Function Rule (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckFunctionRuleCreate(actionName, triggerName, name, namespace string) string {
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
		name = "%s"
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
	  
`, actionName, namespace, triggerName, namespace, name, namespace)

}

func testAccCheckFunctionRuleUpdate(updatedTriggerName, name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_trigger" "triggerUpdated" {
		name                     = "%s"
		namespace		 = "%s"
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
		trigger_name = ibm_function_trigger.triggerUpdated.name
		action_name  = "/whisk.system/cloudant/delete-attachment"
	  }
`, updatedTriggerName, namespace, name, namespace)

}

func testAccCheckFunctionRuleImport(triggerName, name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_trigger" "trigger" {
		name                     = "%s"
		namespace		 = "%s"
		user_defined_annotations = <<EOF
					   [
			   {
					   "key":"sample trigger",
					   "value":"Trigger for hello action"
			   }
					   ]
	  
	  EOF
	  
	  }
	  
	  resource "ibm_function_rule" "import" {
		name         = "%s"
		namespace    = "%s"
		trigger_name = ibm_function_trigger.trigger.name
		action_name  = "/whisk.system/cloudant/delete-attachment"
	  }
`, triggerName, namespace, name, namespace)

}
