// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSchematicsInventoryDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsInventoryDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "inventory_id"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsInventoryDataSourceAllArgs(t *testing.T) {
	inventoryResourceRecordName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	inventoryResourceRecordDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	inventoryResourceRecordLocation := "us-south"
	inventoryResourceRecordResourceGroup := "Default"
	inventoryResourceRecordInventoriesIni := fmt.Sprintf("tf_inventories_ini_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsInventoryDataSourceConfig(inventoryResourceRecordName, inventoryResourceRecordDescription, inventoryResourceRecordLocation, inventoryResourceRecordResourceGroup, inventoryResourceRecordInventoriesIni),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "inventory_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_inventory.schematics_inventory", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsInventoryDataSourceConfigBasic() string {
	return `
		resource "ibm_schematics_inventory" "schematics_inventory" {
			name = "test_inventory"
			location = "us-south"
		}

		data "ibm_schematics_inventory" "schematics_inventory" {
			inventory_id = ibm_schematics_inventory.schematics_inventory.id
		}
	`
}

func testAccCheckIBMSchematicsInventoryDataSourceConfig(inventoryResourceRecordName string, inventoryResourceRecordDescription string, inventoryResourceRecordLocation string, inventoryResourceRecordResourceGroup string, inventoryResourceRecordInventoriesIni string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_inventory" "schematics_inventory" {
			name = "%s"
			description = "%s"
			location = "%s"
			resource_group = "%s"
		}

		data "ibm_schematics_inventory" "schematics_inventory" {
			inventory_id = ibm_schematics_inventory.schematics_inventory.id
		}
	`, inventoryResourceRecordName, inventoryResourceRecordDescription, inventoryResourceRecordLocation, inventoryResourceRecordResourceGroup)
}
