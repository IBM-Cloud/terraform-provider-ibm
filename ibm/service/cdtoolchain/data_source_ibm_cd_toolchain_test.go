// Copyright IBM Corp. 2022, 2023 All Rights Reserved.
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
	toolchainName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	toolchainResourceGroupName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainDataSourceConfigBasic(toolchainName, toolchainResourceGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "ui_href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "created_by"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainDataSourceAllArgs(t *testing.T) {
	toolchainName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	toolchainResourceGroupName := acc.CdResourceGroupName
	toolchainDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: ">=0.9.1",
			},
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainDataSourceConfig(toolchainName, toolchainResourceGroupName, toolchainDescription),
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
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "ui_href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain.cd_toolchain", "created_by"),
					resource.TestCheckResourceAttr("data.ibm_cd_toolchain.cd_toolchain", "tags.#", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainDataSourceConfigBasic(toolchainName string, toolchainResourceGroupName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		data "ibm_cd_toolchain" "cd_toolchain" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
		}
	`, toolchainResourceGroupName, toolchainName)
}

func testAccCheckIBMCdToolchainDataSourceConfig(toolchainName string, toolchainResourceGroupName string, toolchainDescription string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
			description = "%s"
		}

		resource "time_sleep" "wait_time" {
			create_duration = "10s"
			depends_on = [ibm_cd_toolchain.cd_toolchain]
		}
		  
		data "ibm_cd_toolchain" "cd_toolchain" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			depends_on = [time_sleep.wait_time]
		}
	`, toolchainResourceGroupName, toolchainName, toolchainDescription)
}
