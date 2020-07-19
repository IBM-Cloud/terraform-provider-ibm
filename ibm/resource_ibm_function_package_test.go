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

func TestAccFunctionPackage_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionPackageCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageNameUpdate(updatedName, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageWithAnnotations(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageWithAnnotationsUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageWithParameters(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageWithParametersUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckFunctionPackageUpdatePublish(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
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
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	bindName := "/whisk.system/alarms"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindCreate(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "bind_package_name", bindName),
				),
			},
			resource.TestStep{
				Config: testAccCheckFunctionPackageNameBindUpdate(updatedName, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindWithAnnotations(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindWithAnnotationsUpdate(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindWithParameters(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			resource.TestStep{
				Config: testAccCheckFunctionPackageBindWithParametersUpdate(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},
			resource.TestStep{
				Config: testAccCheckFunctionPackageBindUpdatePublish(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
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
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionPackageImport(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
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

	bxSession, err := testAccProvider.Meta().(ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_package" {
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

		_, _, err = client.Packages.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("Error waiting for IBM Cloud Function Package (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckFunctionPackageCreate(name string, namespace string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
	   name = "%s"
	   namespace = "%s"
}`, name, namespace)

}

func testAccCheckFunctionPackageNameUpdate(updatedName string, namespace string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "package" {
	   name = "%s"
	   namespace = "%s"
}`, updatedName, namespace)
}

func testAccCheckFunctionPackageWithAnnotations(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
		name                     = "%s"
		namespace                = "%s"
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
	  
	  }
`, name, namespace)

}

func testAccCheckFunctionPackageWithAnnotationsUpdate(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
		name                     = "%s"
		namespace                = "%s"
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF 
	  }
`, name, namespace)

}

func testAccCheckFunctionPackageWithParameters(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
		name                    = "%s"
		namespace               = "%s"
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
	  
	  }
	  
`, name, namespace)

}

func testAccCheckFunctionPackageWithParametersUpdate(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
		name                    = "%s"
		namespace               = "%s"
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
	  
	  }
`, name, namespace)

}

func testAccCheckFunctionPackageImport(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
   		name = "%s"
		namespace = "%s"	
	}
`, name, namespace)

}

func testAccCheckFunctionPackageUpdatePublish(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
		name                    = "%s"
		namespace               = "%s"
		publish                 = true
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
	  
	  }
`, name, namespace)
}

func testAccCheckFunctionPackageBindCreate(name, namespace, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "bindpackage" {
	   name = "%s"
	   namespace = "%s"
	   bind_package_name = "%s"
}`, name, namespace, bind)

}

func testAccCheckFunctionPackageNameBindUpdate(updatedName, namespace, bind string) string {
	return fmt.Sprintf(`
	
resource "ibm_function_package" "bindpackage" {
	   name = "%s"
	   namespace = "%s"
	   bind_package_name = "%s"
}`, updatedName, namespace, bind)
}

func testAccCheckFunctionPackageBindWithAnnotations(name, namespace, bind string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "bindpackage" {
		name                     = "%s"
		namespace                = "%s"
		bind_package_name        = "%s"
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
	  
	  }
`, name, namespace, bind)

}
func testAccCheckFunctionPackageBindWithAnnotationsUpdate(name, namespace, bind string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "bindpackage" {
		name                     = "%s"
		namespace                = "%s"
		bind_package_name        = "%s"
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"binded alaram package"
		  }
	  ]
	  EOF
	  
	  }
`, name, namespace, bind)

}

func testAccCheckFunctionPackageBindWithParameters(name, namespace, bind string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "bindpackage" {
		name                    = "%s"
		namespace               = "%s"
		bind_package_name       = "%s"
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
	  
	  }
`, name, namespace, bind)

}

func testAccCheckFunctionPackageBindWithParametersUpdate(name, namespace, bind string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "bindpackage" {
		name                    = "%s"
		namespace		= "%s"
		bind_package_name       = "%s"
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
	  
	  }
`, name, namespace, bind)

}

func testAccCheckFunctionPackageBindUpdatePublish(name, namespace, bind string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "bindpackage" {
		name                    = "%s"
		namespace		= "%s"
		bind_package_name       = "%s"
		publish                 = true
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
	  
	  }
	  
`, name, namespace, bind)
}
