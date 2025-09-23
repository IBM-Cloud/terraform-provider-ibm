// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package functions_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/apache/openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/Mavrickk3/bluemix-go/bmxerror"
)

func TestAccCFFunctionRule_Basic(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_rule_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	actionName := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	triggerName := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	updatedTriggerName := fmt.Sprintf("terraform_update_trigger_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionRuleDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionRuleCreate(actionName, triggerName, name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionRuleExists("ibm_function_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "trigger_name", triggerName),
				),
			},
			{
				Config: testAccCheckCFFunctionRuleUpdate(updatedTriggerName, name, namespace),
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

func TestAccIAMFunctionRule_Basic(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_rule_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	actionName := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	triggerName := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	updatedTriggerName := fmt.Sprintf("terraform_update_trigger_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionRuleDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionRuleCreate(actionName, triggerName, name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionRuleExists("ibm_function_rule.rule", &conf),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "name", name),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_rule.rule", "trigger_name", triggerName),
				),
			},
			{
				Config: testAccCheckIAMFunctionRuleUpdate(updatedTriggerName, name, namespace),
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

func TestAccCFFunctionRule_Import(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_rule_%d", acctest.RandIntRange(10, 100))
	triggeName := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionRuleDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionRuleImport(triggeName, name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionRuleExists("ibm_function_rule.import", &conf),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "name", name),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "publish", "false"),
				),
			},

			{
				ResourceName:      "ibm_function_rule.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIAMFunctionRule_Import(t *testing.T) {
	var conf whisk.Rule
	name := fmt.Sprintf("terraform_rule_%d", acctest.RandIntRange(10, 100))
	triggeName := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionRuleDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionRuleImport(triggeName, name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionRuleExists("ibm_function_rule.import", &conf),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "name", name),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_rule.import", "publish", "false"),
				),
			},

			{
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

		parts, err := flex.CfIdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		namespace := parts[0]
		name := parts[1]

		functionNamespaceAPI, err := acc.TestAccProvider.Meta().(conns.ClientSession).FunctionIAMNamespaceAPI()
		if err != nil {
			return err
		}

		bxSession, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
		if err != nil {
			return err
		}

		wskClient, err := conns.SetupOpenWhiskClientConfig(namespace, bxSession, functionNamespaceAPI)
		if err != nil {
			return err

		}

		rule, _, err := wskClient.Rules.Get(name)
		if err != nil {
			return err
		}

		*obj = *rule
		return nil
	}
}

func testAccCheckFunctionRuleDestroy(s *terraform.State) error {
	functionNamespaceAPI, err := acc.TestAccProvider.Meta().(conns.ClientSession).FunctionIAMNamespaceAPI()
	if err != nil {
		return err
	}

	bxSession, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_rule" {
			continue
		}

		parts, err := flex.CfIdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		namespace := parts[0]
		name := parts[1]

		client, err := conns.SetupOpenWhiskClientConfig(namespace, bxSession, functionNamespaceAPI)
		if err != nil && strings.Contains(err.Error(), "is not in the list of entitled namespaces") {
			return nil
		}
		if err != nil {
			return err
		}

		_, _, err = client.Rules.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("[ERROR] Error waiting for IBM Cloud Function Rule (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckIAMFunctionRuleCreate(actionName, triggerName, name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "action" {
		depends_on = [ibm_function_namespace.namespace]
		name      = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonode.js")
		}
	  }
	  
	  resource "ibm_function_trigger" "trigger" {
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
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
		namespace = ibm_function_namespace.namespace.name
		trigger_name = ibm_function_trigger.trigger.name
		action_name  = ibm_function_action.action.name
	  }
	  
`, namespace, actionName, triggerName, name)

}

func testAccCheckIAMFunctionRuleUpdate(updatedTriggerName, name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_trigger" "triggerUpdated" {
		name                     = "%s"
		namespace 				 = ibm_function_namespace.namespace.name
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
		namespace = ibm_function_namespace.namespace.name
		trigger_name = ibm_function_trigger.triggerUpdated.name
		action_name  = "/whisk.system/cloudant/delete-attachment"
	  }
`, namespace, updatedTriggerName, name)

}

func testAccCheckIAMFunctionRuleImport(triggerName, name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_trigger" "trigger" {
		name                     = "%s"
		namespace 				 = ibm_function_namespace.namespace.name
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
		namespace = ibm_function_namespace.namespace.name
		trigger_name = ibm_function_trigger.trigger.name
		action_name  = "/whisk.system/cloudant/delete-attachment"
	  }
`, namespace, triggerName, name)

}

func testAccCheckCFFunctionRuleImport(triggerName, name, namespace string) string {
	return fmt.Sprintf(`

	resource "ibm_function_trigger" "trigger" {
		name                     = "%s"
		namespace 				 = "%s"
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
		namespace = "%s"
		trigger_name = ibm_function_trigger.trigger.name
		action_name  = "/whisk.system/cloudant/delete-attachment"
	  }
`, triggerName, namespace, name, namespace)

}

func testAccCheckCFFunctionRuleCreate(actionName, triggerName, name, namespace string) string {
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

func testAccCheckCFFunctionRuleUpdate(updatedTriggerName, name, namespace string) string {
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
