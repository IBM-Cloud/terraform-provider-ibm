// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"testing"

	"github.com/IBM-Cloud/bluemix-go/api/functions"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFunctionNamespace_Basic(t *testing.T) {
	var instance string
	name := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resourceGroupName := "default"
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
	resourceGroupName := "default"

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

		nsClient, err := testAccProvider.Meta().(ClientSession).FunctionIAMNamespaceAPI()
		if err != nil {
			return err
		}

		getOptions := functions.GetNamespaceOptions{
			ID: &ID,
		}
		instance1, err := nsClient.Namespaces().GetNamespace(getOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Namesapce (IAM): %s", err)
		}

		instance = *instance1.ID
		return nil
	}
}

func testAccCheckFunctionNamespaceDestroy(s *terraform.State) error {
	nsClient, err := testAccProvider.Meta().(ClientSession).FunctionIAMNamespaceAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_namespace" {
			continue
		}

		ID := rs.Primary.ID
		_, err := nsClient.Namespaces().DeleteNamespace(ID)
		if err != nil {
			log.Printf("Error deleting namespace (IAM): %s", err)
			return err
		}

		getOptions := functions.GetNamespaceOptions{
			ID: &ID,
		}
		_, err = nsClient.Namespaces().GetNamespace(getOptions)
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
