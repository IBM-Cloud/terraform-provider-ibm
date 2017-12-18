package ibm

import (
	"fmt"
	"testing"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/IBM-Bluemix/bluemix-go/bmxerror"
)

func TestAccCloudFunctionsTrigger_Basic(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsTriggerCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsTriggerExists("ibm_cloud_functions_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.trigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.trigger", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsTriggerUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsTriggerExists("ibm_cloud_functions_trigger.trigger", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.trigger", "version", "0.0.2"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsTrigger_Feed_Basic(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsTriggerFeedCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsTriggerExists("ibm_cloud_functions_trigger.feedtrigger", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.feedtrigger", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.feedtrigger", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.feedtrigger", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsTriggerFeedUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsTriggerExists("ibm_cloud_functions_trigger.feedtrigger", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.feedtrigger", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.feedtrigger", "version", "0.0.2"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsTrigger_Import(t *testing.T) {
	var conf whisk.Trigger
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsTriggerDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsTriggerImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsTriggerExists("ibm_cloud_functions_trigger.import", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.import", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_trigger.import", "publish", "false"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_cloud_functions_trigger.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudFunctionsTriggerExists(n string, obj *whisk.Trigger) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		client, err := testAccProvider.Meta().(ClientSession).CloudFunctionsClient()
		if err != nil {
			return err
		}
		name := rs.Primary.ID

		trigger, _, err := client.Triggers.Get(name)
		if err != nil {
			return err
		}

		*obj = *trigger
		return nil
	}
}

func testAccCheckCloudFunctionsTriggerDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloud_functions_trigger" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Triggers.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for IBM Cloud Functions Trigger (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckCloudFunctionsTriggerCreate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_trigger" "trigger" {
			name = "%s"		  
			}
`, name)

}

func testAccCheckCloudFunctionsTriggerUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_trigger" "trigger" {
			name = "%s"		  
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
`, name)

}

func testAccCheckCloudFunctionsTriggerFeedCreate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_trigger" "feedtrigger" {
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
	
`, name)

}

func testAccCheckCloudFunctionsTriggerFeedUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_trigger" "feedtrigger" {
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

  user_defined_parameters = <<EOF
        [
    {
        "key":"place",
        "value":"India"
    }
        ]
        EOF
}
`, name)

}

func testAccCheckCloudFunctionsTriggerImport(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_trigger" "import" {
			name = "%s"	
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
`, name)

}
