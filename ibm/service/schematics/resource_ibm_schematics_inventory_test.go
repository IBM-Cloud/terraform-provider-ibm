// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIBMSchematicsInventoryBasic(t *testing.T) {
	var conf schematicsv1.InventoryResourceRecord

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsInventoryDestroy,
		Steps: []resource.TestStep{
			{
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
	resourceGroup := ""
	inventoriesIni := fmt.Sprintf("tf_inventories_ini_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	locationUpdate := location
	resourceGroupUpdate := resourceGroup
	inventoriesIniUpdate := fmt.Sprintf("tf_inventories_ini_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsInventoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsInventoryConfig(name, description, location, resourceGroup, inventoriesIni),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsInventoryExists("ibm_schematics_inventory.schematics_inventory", conf),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "description", description),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "location", location),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "resource_group", resourceGroup),
				),
			},
			{
				Config: testAccCheckIBMSchematicsInventoryConfig(nameUpdate, descriptionUpdate, locationUpdate, resourceGroupUpdate, inventoriesIniUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_inventory.schematics_inventory", "resource_group", resourceGroupUpdate),
				),
			},
			{
				ResourceName:      "ibm_schematics_inventory.schematics_inventory",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsInventoryConfigBasic() string {
	return `

		resource "ibm_schematics_inventory" "schematics_inventory" {
			name = "test_inventory"
			location = "us-south"
		}
	`
}

func testAccCheckIBMSchematicsInventoryConfig(name string, description string, location string, resourceGroup string, inventoriesIni string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_inventory" "schematics_inventory" {
			name = "%s"
			description = "%s"
			location = "us-south"
		}
	`, name, description)
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
			return fmt.Errorf("[ERROR] Error checking for schematics_inventory (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
