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

func TestAccIAMFunctionTrigger_Basic(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionTriggerDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionTriggerCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "publish", "false"),
				),
			},

			{
				Config: testAccCheckIAMFunctionTriggerUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "version", "0.0.2"),
				),
			},
		},
	})
}

func TestAccIAMFunctionTrigger_Feed_Basic(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionTriggerDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionTriggerFeedCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.feedtrigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "feed.0.name", "/whisk.system/alarms/alarm"),
				),
			},

			{
				Config: testAccCheckIAMFunctionTriggerFeedUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.feedtrigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "feed.0.name", "/whisk.system/alarms/alarm"),
				),
			},
		},
	})
}

func TestAccIAMFunctionTrigger_Import(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionTriggerDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionTriggerImport(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.import", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "publish", "false"),
				),
			},

			{
				ResourceName:      "ibm_function_trigger.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCFFunctionTrigger_Basic(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionTriggerDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionTriggerCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "publish", "false"),
				),
			},

			{
				Config: testAccCheckCFFunctionTriggerUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "version", "0.0.2"),
				),
			},
		},
	})
}

func TestAccCFFunctionTrigger_Feed_Basic(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionTriggerDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionTriggerFeedCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.feedtrigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "feed.0.name", "/whisk.system/alarms/alarm"),
				),
			},

			{
				Config: testAccCheckCFFunctionTriggerFeedUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.feedtrigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "feed.0.name", "/whisk.system/alarms/alarm"),
				),
			},
		},
	})
}

func TestAccCFFunctionTrigger_Import(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_trigger_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionTriggerDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionTriggerImport(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.import", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "publish", "false"),
				),
			},

			{
				ResourceName:      "ibm_function_trigger.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckFunctionTriggerExists(n string, obj *whisk.Trigger) resource.TestCheckFunc {

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

		client, err := conns.SetupOpenWhiskClientConfig(namespace, bxSession, functionNamespaceAPI)
		if err != nil {
			return err

		}

		trigger, _, err := client.Triggers.Get(name)
		if err != nil {
			return err
		}

		*obj = *trigger
		return nil
	}
}

func testAccCheckFunctionTriggerDestroy(s *terraform.State) error {
	functionNamespaceAPI, err := acc.TestAccProvider.Meta().(conns.ClientSession).FunctionIAMNamespaceAPI()
	if err != nil {
		return err
	}

	bxSession, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_trigger" {
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
		_, _, err = client.Triggers.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("[ERROR] Error waiting for IBM Cloud Function Trigger (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckIAMFunctionTriggerCreate(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}


	resource "ibm_function_trigger" "trigger" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
	}
`, namespace, name)

}

func testAccCheckIAMFunctionTriggerUpdate(name, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_trigger" "trigger" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		user_defined_parameters = <<EOF
							  [
									  {
										 "key":"place",
											  "value":"India"
								 }
						 ]
	  
	  EOF
	  
	  
		user_defined_annotations = <<EOF
				 [
						 {
								"key":"Description",
								 "value":"Sample code to display hello"
						}
				]
	  
	  EOF
	  
	  }
`, namespace, name)

}

func testAccCheckIAMFunctionTriggerFeedCreate(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_trigger" "feedtrigger" {
		depends_on = [ibm_function_namespace.namespace]
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
`, namespace, name)

}

func testAccCheckIAMFunctionTriggerFeedUpdate(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_trigger" "feedtrigger" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		feed {
		  name = "/whisk.system/alarms/alarm"
	  
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
	  
	  
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"place",
			  "value":"India"
		  }
			  ]
	  
	  EOF
	  
	  }
	  
`, namespace, name)

}

func testAccCheckIAMFunctionTriggerImport(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}


	resource "ibm_function_trigger" "import" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		user_defined_parameters = <<EOF
							  [
									  {
										 "key":"place",
											  "value":"India"
								 }
						 ]
	  
	  EOF
	  
	  
		user_defined_annotations = <<EOF
				 [
						 {
								"key":"Description",
								 "value":"Sample code to display hello"
						}
				]
	  
	  EOF
	  
	  }
`, namespace, name)

}

func testAccCheckCFFunctionTriggerCreate(name, namespace string) string {
	return fmt.Sprintf(`
		resource "ibm_function_trigger" "trigger" {
			name = "%s"		  
			namespace = "%s"
			}
`, name, namespace)

}

func testAccCheckCFFunctionTriggerUpdate(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_trigger" "trigger" {
		name                    = "%s"
		namespace               = "%s"
		user_defined_parameters = <<EOF
							  [
									  {
										 "key":"place",
											  "value":"India"
								 }
						 ]
	  
	  EOF
	  
	  
		user_defined_annotations = <<EOF
				 [
						 {
								"key":"Description",
								 "value":"Sample code to display hello"
						}
				]
	  
	  EOF
	  
	  }
`, name, namespace)

}

func testAccCheckCFFunctionTriggerFeedCreate(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_trigger" "feedtrigger" {
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
`, name, namespace)

}

func testAccCheckCFFunctionTriggerFeedUpdate(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_trigger" "feedtrigger" {
		name = "%s"
		namespace = "%s"
		feed {
		  name = "/whisk.system/alarms/alarm"
	  
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
	  
	  
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"place",
			  "value":"India"
		  }
			  ]
	  
	  EOF
	  
	  }
	  
`, name, namespace)

}

func testAccCheckCFFunctionTriggerImport(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_trigger" "import" {
		name                    = "%s"
		namespace		= "%s"
		user_defined_parameters = <<EOF
							  [
									  {
										 "key":"place",
											  "value":"India"
								 }
						 ]
	  
	  EOF
	  
	  
		user_defined_annotations = <<EOF
				 [
						 {
								"key":"Description",
								 "value":"Sample code to display hello"
						}
				]
	  
	  EOF
	  
	  }
`, name, namespace)

}
