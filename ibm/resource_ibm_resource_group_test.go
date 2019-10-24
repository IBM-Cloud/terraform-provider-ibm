package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/bluemix-go/models"

	"strings"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMResourceGroup_Basic(t *testing.T) {
	var conf models.ResourceGroup
	resourceGroupName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceGroupUpdateName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMResourceGroup_basic(resourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceGroupExists("ibm_resource_group.resourceGroup", &conf),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "name", resourceGroupName),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "default", "false"),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "state", "ACTIVE"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMResourceGroup_basic(resourceGroupUpdateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceGroupExists("ibm_resource_group.resourceGroup", &conf),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "name", resourceGroupUpdateName),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "default", "false"),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "state", "ACTIVE"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_resource_group.resourceGroup",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMResourceGroup_With_Tags(t *testing.T) {
	var conf models.ResourceGroup
	resourceGroupName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMResourceGroup_with_tags(resourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceGroupExists("ibm_resource_group.resourceGroup", &conf),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "name", resourceGroupName),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "default", "false"),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "state", "ACTIVE"),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "tags.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMResourceGroup_with_updated_tags(resourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceGroupExists("ibm_resource_group.resourceGroup", &conf),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "tags.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMResourceGroupExists(n string, obj *models.ResourceGroup) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceManagementAPI()
		if err != nil {
			return err
		}
		resourceGroupID := rs.Primary.ID

		resourceGroup, err := rsContClient.ResourceGroup().Get(resourceGroupID)
		if err != nil {
			return err
		}

		obj = resourceGroup
		return nil
	}
}

func testAccCheckIBMResourceGroupDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).ResourceManagementAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_resource_group" {
			continue
		}

		resourceGroupID := rs.Primary.ID

		// Try to find the key
		_, err := rsContClient.ResourceGroup().Get(resourceGroupID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for resource group (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMResourceGroup_basic(resourceGroupName string) string {
	return fmt.Sprintf(`
		  
		  resource "ibm_resource_group" "resourceGroup" {
			name     = "%s"
		  }
	`, resourceGroupName)
}

func testAccCheckIBMResourceGroup_with_tags(resourceGroupName string) string {
	return fmt.Sprintf(`
		  
		  resource "ibm_resource_group" "resourceGroup" {
			name     = "%s"
			tags     = ["one"]
		  }
	`, resourceGroupName)
}

func testAccCheckIBMResourceGroup_with_updated_tags(resourceGroupName string) string {
	return fmt.Sprintf(`
		  
		  resource "ibm_resource_group" "resourceGroup" {
			name     = "%s"
			tags     = ["one", "two"]
		  }
	`, resourceGroupName)
}
