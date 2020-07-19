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

func TestAccFunctionTrigger_Basic(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionTriggerCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.trigger", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionTriggerUpdate(name, namespace),
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

func TestAccFunctionTrigger_Feed_Basic(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionTriggerFeedCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.feedtrigger", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_trigger.feedtrigger", "feed.0.name", "/whisk.system/alarms/alarm"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionTriggerFeedUpdate(name, namespace),
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

func TestAccFunctionTrigger_Import(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionTriggerImport(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionTriggerExists("ibm_function_trigger.import", &conf),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "name", name),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_trigger.import", "publish", "false"),
				),
			},

			resource.TestStep{
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

		trigger, _, err := client.Triggers.Get(name)
		if err != nil {
			return err
		}

		*obj = *trigger
		return nil
	}
}

func testAccCheckFunctionTriggerDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).FunctionClient()
	if err != nil {
		return err
	}

	bxSession, err := testAccProvider.Meta().(ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_trigger" {
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
		_, _, err = client.Triggers.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for IBM Cloud Function Trigger (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckFunctionTriggerCreate(name, namespace string) string {
	return fmt.Sprintf(`
		resource "ibm_function_trigger" "trigger" {
			name = "%s"		  
			namespace = "%s"
			}
`, name, namespace)

}

func testAccCheckFunctionTriggerUpdate(name, namespace string) string {
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

func testAccCheckFunctionTriggerFeedCreate(name, namespace string) string {
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

func testAccCheckFunctionTriggerFeedUpdate(name, namespace string) string {
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

func testAccCheckFunctionTriggerImport(name, namespace string) string {
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
