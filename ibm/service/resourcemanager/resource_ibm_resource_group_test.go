// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMResourceGroupBasic(t *testing.T) {
	var conf string
	resourceGroupName := fmt.Sprintf("tf-rg-%d", acctest.RandIntRange(10, 100))
	resourceGroupUpdateName := fmt.Sprintf("tf-rg-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupBasic(resourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceGroupExists("ibm_resource_group.resourceGroup", &conf),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "name", resourceGroupName),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "default", "false"),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "state", "ACTIVE"),
				),
			},
			{
				Config: testAccCheckIBMResourceGroupBasic(resourceGroupUpdateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceGroupExists("ibm_resource_group.resourceGroup", &conf),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "name", resourceGroupUpdateName),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "default", "false"),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "state", "ACTIVE"),
				),
			},
			{
				ResourceName:      "ibm_resource_group.resourceGroup",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMResourceGroupWithTags(t *testing.T) {
	var conf string
	resourceGroupName := fmt.Sprintf("tf-rg-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceGroupWithtags(resourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceGroupExists("ibm_resource_group.resourceGroup", &conf),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "name", resourceGroupName),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "default", "false"),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "state", "ACTIVE"),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMResourceGroupWithupdatedTags(resourceGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceGroupExists("ibm_resource_group.resourceGroup", &conf),
					resource.TestCheckResourceAttr("ibm_resource_group.resourceGroup", "tags.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMResourceGroupExists(n string, obj *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceManagerV2API()
		if err != nil {
			return err
		}
		resourceGroupID := rs.Primary.ID

		resourceGroupGet := rg.GetResourceGroupOptions{
			ID: &resourceGroupID,
		}

		resourceGroup, resp, err := rsContClient.GetResourceGroup(&resourceGroupGet)
		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving resource group: %s\n Response code is: %+v", err, resp)
		}

		obj = resourceGroup.ID
		return nil
	}
}

func testAccCheckIBMResourceGroupDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_resource_group" {
			continue
		}

		resourceGroupID := rs.Primary.ID
		resourceGroupGet := rg.GetResourceGroupOptions{
			ID: &resourceGroupID,
		}

		_, resp, err := rsContClient.GetResourceGroup(&resourceGroupGet)

		if err == nil {
			if resp != nil && resp.StatusCode == 404 {
				return nil
			}
			return fmt.Errorf("Resource group still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMResourceGroupBasic(resourceGroupName string) string {
	return fmt.Sprintf(`
		  
		  resource "ibm_resource_group" "resourceGroup" {
			name     = "%s"
		  }
	`, resourceGroupName)
}

func testAccCheckIBMResourceGroupWithtags(resourceGroupName string) string {
	return fmt.Sprintf(`
		  
		  resource "ibm_resource_group" "resourceGroup" {
			name     = "%s"
			tags     = ["one"]
		  }
	`, resourceGroupName)
}

func testAccCheckIBMResourceGroupWithupdatedTags(resourceGroupName string) string {
	return fmt.Sprintf(`
		  
		  resource "ibm_resource_group" "resourceGroup" {
			name     = "%s"
			tags     = ["one", "two"]
		  }
	`, resourceGroupName)
}
