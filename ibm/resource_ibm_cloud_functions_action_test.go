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

func TestAccCloudFunctionsAction_NodeJS(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionNodeJS(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.nodehello", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.nodehello", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.nodehello", "exec.1729670059.kind", "nodejs:6"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_NodeJSWithParams(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionNodeJSWithParams(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.nodehellowithparameter", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.nodehellowithparameter", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.nodehellowithparameter", "exec.2223737754.kind", "nodejs:6"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_NodeJSZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionNodeJSZip(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.nodezip", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.nodezip", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.nodezip", "exec.915916944.kind", "nodejs:6"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_Python(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionPython(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.pythonhello", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.pythonhello", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.pythonhello", "exec.2177277460.kind", "python"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_PythonZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionPythonZip(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.pythonzip", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.pythonzip", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.pythonzip", "exec.799663035.kind", "python"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_PHP(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionPHP(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.phphello", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.phphello", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.phphello", "exec.2196642062.kind", "php:7.1"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_PHPZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionPHPZip(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.phpzip", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.phpzip", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.phpzip", "exec.36392495.kind", "php:7.1"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_Swift(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionSwift(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.swifthello", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.swifthello", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.swifthello", "exec.455016215.kind", "swift:3.1.1"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_Sequence(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionSequence(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.sequence", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.sequence", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.sequence", "exec.3443818261.kind", "sequence"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_Basic(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.action", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.action", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.action", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.action", "publish", "true"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsAction_Import(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsActionImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsActionExists("ibm_cloud_functions_action.import", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.import", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_action.import", "publish", "false"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_cloud_functions_action.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudFunctionsActionExists(n string, obj *whisk.Action) resource.TestCheckFunc {

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

		action, _, err := client.Actions.Get(name)
		if err != nil {
			return err
		}

		*obj = *action
		return nil
	}
}

func testAccCheckCloudFunctionsActionDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloud_functions_action" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Actions.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for CloudFunctions Action (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckCloudFunctionsActionNodeJS(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "nodehello" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonode.js")}"
			}
		  }
	
`, name)

}

func testAccCheckCloudFunctionsActionNodeJSWithParams(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "nodehellowithparameter" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonodewithparameter.js")}"
			}
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

func testAccCheckCloudFunctionsActionNodeJSZip(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "nodezip" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${base64encode("${file("test-fixtures/nodeaction.zip")}")}"
			}
		  }
	
`, name)

}

func testAccCheckCloudFunctionsActionPython(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "pythonhello" {
			name = "%s"		  
			exec = {
			  kind = "python"
			  code = "${file("test-fixtures/hellopython.py")}"
			}
		  }
	
`, name)

}

func testAccCheckCloudFunctionsActionPythonZip(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "pythonzip" {
			name = "%s"		  
			exec = {
			  kind = "python"
			  code = "${base64encode("${file("test-fixtures/pythonaction.zip")}")}"
			}
		  }
	
`, name)

}

func testAccCheckCloudFunctionsActionPHP(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "phphello" {
			name = "%s"		  
			exec = {
			  kind = "php:7.1"
			  code = "${file("test-fixtures/hellophp.php")}"
			}
		  }
	
`, name)

}

func testAccCheckCloudFunctionsActionPHPZip(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "phpzip" {
			name = "%s"		  
			exec = {
			  kind = "php:7.1"
			  code = "${base64encode("${file("test-fixtures/phpaction.zip")}")}"
			}
		  }
	
`, name)

}

func testAccCheckCloudFunctionsActionSwift(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "swifthello" {
			name = "%s"		  
			exec = {
			  kind = "swift:3.1.1"
			  code = "${file("test-fixtures/helloSwift.swift")}"
			}
		  }
	
`, name)

}

func testAccCheckCloudFunctionsActionSequence(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "sequence" {
			name = "%s"		  
			exec = {
			  kind = "sequence"
			  components = ["/whisk.system/utils/split","/whisk.system/utils/sort"]
			}
		  }
	
`, name)

}

func testAccCheckCloudFunctionsActionCreate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "action" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonode.js")}"
			}
			}
`, name)

}

func testAccCheckCloudFunctionsActionUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "action" {
			name = "%s"	
			publish = "true"	  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonodewithparameter.js")}"
			}
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

func testAccCheckCloudFunctionsActionImport(name string) string {
	return fmt.Sprintf(`
		resource "ibm_cloud_functions_action" "import" {
			name = "%s"	
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonodewithparameter.js")}"
			}
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
