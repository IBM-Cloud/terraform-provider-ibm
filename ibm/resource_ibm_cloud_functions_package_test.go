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

func TestAccCloudFunctionsPackage_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageNameUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageWithAnnotations(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageWithAnnotationsUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageWithParameters(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageWithParametersUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "publish", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageUpdatePublish(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "publish", "true"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsPackage_Bind_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())
	bindName := "/whisk.system/alarms"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageBindCreate(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "bind_package_name", bindName),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageNameBindUpdate(updatedName, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageBindWithAnnotations(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageBindWithAnnotationsUpdate(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageBindWithParameters(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageBindWithParametersUpdate(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "publish", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageBindUpdatePublish(name, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.bindpackage", "publish", "true"),
				),
			},
		},
	})
}

func TestAccCloudFunctionsPackage_Import(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudFunctionsPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckCloudFunctionsPackageImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudFunctionsPackageExists("ibm_cloud_functions_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_cloud_functions_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_cloud_functions_package.package",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudFunctionsPackageExists(n string, obj *whisk.Package) resource.TestCheckFunc {

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

		pkg, _, err := client.Packages.Get(name)
		if err != nil {
			return err
		}

		*obj = *pkg
		return nil
	}
}

func testAccCheckCloudFunctionsPackageDestroy(s *terraform.State) error {
	client, err := testAccProvider.Meta().(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cloud_functions_package" {
			continue
		}

		name := rs.Primary.ID
		_, _, err := client.Packages.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for IBM Cloud Functions Package (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckCloudFunctionsPackageCreate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "package" {
	   name = "%s"
}`, name)

}

func testAccCheckCloudFunctionsPackageNameUpdate(updatedName string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "package" {
	   name = "%s"
}`, updatedName)
}

func testAccCheckCloudFunctionsPackageWithAnnotations(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "package" {
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

func testAccCheckCloudFunctionsPackageWithAnnotationsUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "package" {
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

func testAccCheckCloudFunctionsPackageWithParameters(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "package" {
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

func testAccCheckCloudFunctionsPackageWithParametersUpdate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "package" {
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

func testAccCheckCloudFunctionsPackageImport(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "package" {
   	name = "%s"
}`, name)

}

func testAccCheckCloudFunctionsPackageUpdatePublish(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "package" {
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

func testAccCheckCloudFunctionsPackageBindCreate(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
}`, name, bind)

}

func testAccCheckCloudFunctionsPackageNameBindUpdate(updatedName, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "bindpackage" {
	   name = "%s"
	   bind_package_name = "%s"
}`, updatedName, bind)
}

func testAccCheckCloudFunctionsPackageBindWithAnnotations(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "bindpackage" {
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
func testAccCheckCloudFunctionsPackageBindWithAnnotationsUpdate(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "bindpackage" {
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

func testAccCheckCloudFunctionsPackageBindWithParameters(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "bindpackage" {
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

func testAccCheckCloudFunctionsPackageBindWithParametersUpdate(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "bindpackage" {
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

func testAccCheckCloudFunctionsPackageBindUpdatePublish(name, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_cloud_functions_package" "bindpackage" {
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
