package ibm

import (
	"fmt"
	"testing"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
)

func TestAccFunctionAction_NodeJS(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionNodeJS(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.nodehello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "exec.0.kind", "nodejs:6"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_NodeJSWithParams(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionNodeJSWithParams(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.nodehellowithparameter", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "exec.0.kind", "nodejs:6"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_NodeJSZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionNodeJSZip(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.nodezip", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "exec.0.kind", "nodejs:6"),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_Python(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionPython(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.pythonhello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "exec.0.kind", "python"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_PythonZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionPythonZip(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.pythonzip", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "exec.0.kind", "python"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_PHP(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionPHP(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.phphello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "exec.0.kind", "php:7.1"),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_PHPZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionPHPZip(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.phpzip", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "exec.0.kind", "php:7.1"),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_Swift(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionSwift(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.swifthello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "exec.0.kind", "swift:3.1.1"),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_Sequence(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionSequence(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.sequence", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "exec.0.kind", "sequence"),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_Basic(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.action", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.memory", "256"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionActionUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.action", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "publish", "true"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.log_size", "5"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.timeout", "50000"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccFunctionAction_Import(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionActionImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.import", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.import", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_action.import", "publish", "false"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_function_action.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckFunctionActionExists(n string, obj *whisk.Action) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		client, err := testAccProvider.Meta().(ClientSession).FunctionClient()
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

func testAccCheckFunctionActionDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).FunctionClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_action" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Actions.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for IBM Cloud Function Action (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckFunctionActionNodeJS(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "nodehello" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonode.js")}"
			}
		  }
	
`, name)

}

func testAccCheckFunctionActionNodeJSWithParams(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "nodehellowithparameter" {
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

func testAccCheckFunctionActionNodeJSZip(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "nodezip" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${base64encode("${file("test-fixtures/nodeaction.zip")}")}"
			}
		  }
	
`, name)

}

func testAccCheckFunctionActionPython(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "pythonhello" {
			name = "%s"		  
			exec = {
			  kind = "python"
			  code = "${file("test-fixtures/helloPython.py")}"
			}
		  }
	
`, name)

}

func testAccCheckFunctionActionPythonZip(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "pythonzip" {
			name = "%s"		  
			exec = {
			  kind = "python"
			  code = "${base64encode("${file("test-fixtures/pythonaction.zip")}")}"
			}
		  }
	
`, name)

}

func testAccCheckFunctionActionPHP(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "phphello" {
			name = "%s"		  
			exec = {
			  kind = "php:7.1"
			  code = "${file("test-fixtures/hellophp.php")}"
			}
		  }
	
`, name)

}

func testAccCheckFunctionActionPHPZip(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "phpzip" {
			name = "%s"		  
			exec = {
			  kind = "php:7.1"
			  code = "${base64encode("${file("test-fixtures/phpaction.zip")}")}"
			}
		  }
	
`, name)

}

func testAccCheckFunctionActionSwift(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "swifthello" {
			name = "%s"		  
			exec = {
			  kind = "swift:3.1.1"
			  code = "${file("test-fixtures/helloSwift.swift")}"
			}
		  }
	
`, name)

}

func testAccCheckFunctionActionSequence(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "sequence" {
			name = "%s"		  
			exec = {
			  kind = "sequence"
			  components = ["/whisk.system/utils/split","/whisk.system/utils/sort"]
			}
		  }
	
`, name)

}

func testAccCheckFunctionActionCreate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "action" {
			name = "%s"		  
			exec = {
			  kind = "nodejs:6"
			  code = "${file("test-fixtures/hellonode.js")}"
			}
			limits = {

			}
			}
`, name)

}

func testAccCheckFunctionActionUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "action" {
			name = "%s"	
			publish = "true"
			limits = {
				log_size = 5
				timeout = 50000
				}	  
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

func testAccCheckFunctionActionImport(name string) string {
	return fmt.Sprintf(`
		resource "ibm_function_action" "import" {
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
