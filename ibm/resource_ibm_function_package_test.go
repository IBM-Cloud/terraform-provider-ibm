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

func TestAccFunctionPackage_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionPackageCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageNameUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageWithAnnotations(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageWithAnnotationsUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageWithParameters(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageWithParametersUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckFunctionPackageUpdatePublish(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "true"),
				),
			},
		},
	})
}

func TestAccFunctionPackage_Bind_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())
	bindName := "/whisk.system/alarms"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindCreate(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "bind_package_name", bindName),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageNameBindUpdate(updatedName, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindWithAnnotations(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindWithAnnotationsUpdate(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindWithParameters(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindWithParametersUpdate(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckFunctionPackageBindUpdatePublish(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "true"),
				),
			},
		},
	})
}

func TestAccFunctionPackage_Import(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionPackageImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_function_package.package",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckFunctionPackageExists(n string, obj *whisk.Package) resource.TestCheckFunc {

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

		pkg, _, err := client.Packages.Get(name)
		if err != nil {
			return err
		}

		*obj = *pkg
		return nil
	}
}

func testAccCheckFunctionPackageDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).FunctionClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_package" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Packages.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for IBM Cloud Function Package (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckFunctionPackageCreate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
	   name = "%s"
}`, name)

}

func testAccCheckFunctionPackageNameUpdate(updatedName string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
	   name = "%s"
}`, updatedName)
}

func testAccCheckFunctionPackageWithAnnotations(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
   	name = "%s"
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    },
    {
        "key":"sampleOutput",
        "value": {
			"count": 3
		}
    },
    {
        "key":"final",
        "value": [
			{
				"description": "A string",
				"name": "payload",
				"required": true
			}
		]
    }
]
EOF

}`, name)

}

func testAccCheckFunctionPackageWithAnnotationsUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
   	name = "%s"
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name)

}

func testAccCheckFunctionPackageWithParameters(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
   	name = "%s"
	user_defined_parameters = <<EOF
	[
    {
        "key":"place",
        "value":"city"
    },
    {
        "key":"parameter",
        "value": {
			"count": 3
		}
    },
    {
        "key":"final",
        "value": [
			{
				"description": "Set of Values",
				"name": "payload",
				"required": true
			}
		]
    }
]
EOF
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name)

}

func testAccCheckFunctionPackageWithParametersUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
   	name = "%s"
	user_defined_parameters = <<EOF
	[
    {
        "key":"name",
        "value":"utils"
    }
]
EOF
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name)

}

func testAccCheckFunctionPackageImport(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
   	name = "%s"
}`, name)

}

func testAccCheckFunctionPackageUpdatePublish(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
	   name = "%s"
	   publish = true
	   user_defined_parameters = <<EOF
	[
    {
        "key":"name",
        "value":"utils"
    }
]
EOF
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF
}`, name)
}

func testAccCheckFunctionPackageBindCreate(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
}`, name, bind)

}

func testAccCheckFunctionPackageNameBindUpdate(updatedName, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
}`, updatedName, bind)
}

func testAccCheckFunctionPackageBindWithAnnotations(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"binded alaram package"
    },
    {
        "key":"sampleOutput",
        "value": {
			"count": 3
		}
    },
    {
        "key":"final",
        "value": [
			{
				"description": "A string",
				"name": "payload",
				"required": true
			}
		]
    }
]
EOF

}`, name, bind)

}
func testAccCheckFunctionPackageBindWithAnnotationsUpdate(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "bindpackage" {
	name = "%s"
	bind_package_name = "%s"
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"binded alaram package"
    }
]
EOF

}`, name, bind)

}

func testAccCheckFunctionPackageBindWithParameters(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "bindpackage" {
	name = "%s"
	bind_package_name = "%s"
	user_defined_parameters = <<EOF
        [
    {
        "key":"cron",
        "value":"0 0 1 0 *"
    },
    {
        "key":"trigger_payload ",
        "value":"{'message':'bye old Year!'}"
    },
    {
        "key":"maxTriggers",
        "value":1
    },
    {
        "key":"userdefined",
        "value":"test"
    }
]
EOF

	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name, bind)

}

func testAccCheckFunctionPackageBindWithParametersUpdate(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "bindpackage" {
	name = "%s"
	bind_package_name = "%s"
	user_defined_parameters = <<EOF
	   [
   {
	   "key":"cron",
	   "value":"0 0 1 0 *"
   }
]
EOF
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF

}`, name, bind)

}

func testAccCheckFunctionPackageBindUpdatePublish(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
	   publish = true
	   user_defined_parameters = <<EOF
	   [
   {
	   "key":"cron",
	   "value":"0 0 1 0 *"
   }
]
EOF
	user_defined_annotations = <<EOF
	[
    {
        "key":"description",
        "value":"Count words in a string"
    }
]
EOF
}`, name, bind)
}
