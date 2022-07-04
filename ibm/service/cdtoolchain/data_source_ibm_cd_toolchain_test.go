// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainDataSourceBasic(t *testing.T) {
	getToolchainByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	getToolchainByIDResponseResourceGroupID := acc.CdResourceGroupID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainDataSourceConfigBasic(getToolchainByIDResponseName, getToolchainByIDResponseResourceGroupID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "created_by"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainDataSourceAllArgs(t *testing.T) {
	getToolchainByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	getToolchainByIDResponseResourceGroupID := acc.CdResourceGroupID
	getToolchainByIDResponseDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainDataSourceConfig(getToolchainByIDResponseName, getToolchainByIDResponseResourceGroupID, getToolchainByIDResponseDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "created_by"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainDataSourceConfigBasic(getToolchainByIDResponseName string, getToolchainByIDResponseResourceGroupID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}

		data "ibm_cd_toolchain" "cd_toolchain" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
		}
	`, getToolchainByIDResponseName, getToolchainByIDResponseResourceGroupID)
}

func testAccCheckIBMCdToolchainDataSourceConfig(getToolchainByIDResponseName string, getToolchainByIDResponseResourceGroupID string, getToolchainByIDResponseDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
			description = "%s"
		}

		data "ibm_cd_toolchain" "cd_toolchain" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
		}
	`, getToolchainByIDResponseName, getToolchainByIDResponseResourceGroupID, getToolchainByIDResponseDescription)
}
