/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/namespace-go-sdk/ibmcloudfunctionsnamespaceapiv1"
)

func TestAccFunctionNamespace_Basic(t *testing.T) {
	var instance string
	name := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resourceGroupName := "Default"
	updateName := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionNamespaceDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionNamespaceCreate(name, resourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionNamespaceExists("ibm_function_namespace.namespace", instance),
					resource.TestCheckResourceAttr("ibm_function_namespace.namespace", "name", name),
					resource.TestCheckResourceAttr("ibm_function_namespace.namespace", "location", "us-south"),
				),
			},
			resource.TestStep{
				Config: testAccCheckFunctionNamespaceUpdate(updateName, resourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionNamespaceExists("ibm_function_namespace.namespace", instance),
					resource.TestCheckResourceAttr("ibm_function_namespace.namespace", "name", updateName),
					resource.TestCheckResourceAttr("ibm_function_namespace.namespace", "location", "us-south"),
				),
			},
		},
	})
}

func TestAccFunctionNamespace_Import(t *testing.T) {
	var instance string
	name := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resourceGroupName := "Default"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFunctionNamespaceDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckFunctionNamespaceImport(resourceGroupName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionNamespaceExists("ibm_function_namespace.namespace", instance),
					resource.TestCheckResourceAttr("ibm_function_namespace.namespace", "name", name),
					resource.TestCheckResourceAttr("ibm_function_namespace.namespace", "location", "us-south"),
				),
			},

			resource.TestStep{
				ResourceName:            "ibm_function_namespace.namespace",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"resource_group_id"},
			},
		},
	})
}

func testAccCheckFunctionNamespaceExists(n string, instance string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		ID := rs.Primary.ID

		nsClient, err := testAccProvider.Meta().(ClientSession).IAMNamespaceAPI()
		if err != nil {
			return err
		}

		getOptions := &ibmcloudfunctionsnamespaceapiv1.GetNamespaceOptions{
			ID: &ID,
			//Headers: headers,
		}
		instance1, _, err := nsClient.GetNamespace(getOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Namesapce (IAM): %s\n", err)
		}

		instance = *instance1.ID
		return nil
	}
}

func testAccCheckFunctionNamespaceDestroy(s *terraform.State) error {
	nsClient, err := testAccProvider.Meta().(ClientSession).IAMNamespaceAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_namespace" {
			continue
		}

		ID := rs.Primary.ID

		delOptions := &ibmcloudfunctionsnamespaceapiv1.DeleteNamespaceOptions{
			ID: &ID,
		}
		response, err := nsClient.DeleteNamespace(delOptions)
		if err != nil && response.StatusCode != 404 {
			log.Printf("Error deleting namespace (IAM): %s", response)
			return err
		}

		getOptions := &ibmcloudfunctionsnamespaceapiv1.GetNamespaceOptions{
			ID: &ID,
		}
		_, _, err = nsClient.GetNamespace(getOptions)
		if err == nil {
			return fmt.Errorf("Namespace still exists: %s", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckFunctionNamespaceCreate(name, resourceGroupName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
                name = "%s"
        }

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	  
	  
`, resourceGroupName, name)

}

func testAccCheckFunctionNamespaceUpdate(name, resourceGroupName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
                name = "%s"
        }

        resource "ibm_function_namespace" "namespace" {
                name                = "%s"
                resource_group_id   = data.ibm_resource_group.test_acc.id
        }

`, resourceGroupName, name)

}

func testAccCheckFunctionNamespaceImport(resourceGroupName, name string) string {
	return fmt.Sprintf(`

        data "ibm_resource_group" "test_acc" {
                name = "%s"
        }

        resource "ibm_function_namespace" "namespace" {
                name                = "%s"
                resource_group_id   = data.ibm_resource_group.test_acc.id
        }

`, resourceGroupName, name)

}
