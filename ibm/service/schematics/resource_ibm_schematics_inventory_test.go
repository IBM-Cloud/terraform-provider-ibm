// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIBMSchematicsInventoryBasic(t *testing.T) {
	var conf schematicsv1.InventoryResourceRecord

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsInventoryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsInventoryConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsInventoryExists("ibm_schematics_inventory.schematics_inventory", conf),
				),
			},
		},
	})
}

func TestAccIBMSchematicsInventoryAllArgs(t *testing.T) {
	var conf schematicsv1.InventoryResourceRecord
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	location := "us-south"
	resourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	inventoriesIni := fmt.Sprintf("tf_inventories_ini_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	locationUpdate := "eu-de"
	resourceGroupUpdate := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	inventoriesIniUpdate := fmt.Sprintf("tf_inventories_ini_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsInventoryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsInventoryConfig(name, description, location, resourceGroup, inventoriesIni),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsInventoryExists("ibm_schematics_inventory.schematics_inventory", conf),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "description", description),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "location", location),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "inventories_ini", inventoriesIni),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsInventoryConfig(nameUpdate, descriptionUpdate, locationUpdate, resourceGroupUpdate, inventoriesIniUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "inventories_ini", inventoriesIniUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_schematics_inventory.schematics_inventory",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsInventoryConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_schematics_inventory" "schematics_inventory" {
		}
	`)
}

func testAccCheckIBMSchematicsInventoryConfig(name string, description string, location string, resourceGroup string, inventoriesIni string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_inventory" "schematics_inventory" {
			name = "%s"
			description = "%s"
			location = "%s"
			resource_group = "%s"
			inventories_ini = "%s"
			// resource_queries = ["FIXME"]
		}
	`, name, description, location, resourceGroup, inventoriesIni)
}

func testAccCheckIBMSchematicsInventoryExists(n string, obj schematicsv1.InventoryResourceRecord) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getInventoryOptions := &schematicsv1.GetInventoryOptions{}

		getInventoryOptions.SetInventoryID(rs.Primary.ID)

		inventoryResourceRecord, _, err := schematicsClient.GetInventory(getInventoryOptions)
		if err != nil {
			return err
		}

		obj = *inventoryResourceRecord
		return nil
	}
}

func testAccCheckIBMSchematicsInventoryDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_inventory" {
			continue
		}

		getInventoryOptions := &schematicsv1.GetInventoryOptions{}

		getInventoryOptions.SetInventoryID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetInventory(getInventoryOptions)

		if err == nil {
			return fmt.Errorf("schematics_inventory still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_inventory (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
