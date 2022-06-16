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

func TestAccIBMCdToolchainToolAppconfigDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgID := acc.CdResourceGroupID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolAppconfigDataSourceConfigBasic(tcName, rgID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolAppconfigDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgID := acc.CdResourceGroupID
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolAppconfigDataSourceConfig(tcName, rgID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolAppconfigDataSourceConfigBasic(tcName string, rgID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}

		resource "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "name"
				region = "region"
				resource_group = "resource-group"
				instance_name = "instance-name"
				environment_name = "environment-name"
				collection_name = "collection-name"
			}
		}

		data "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.toolchain_id
			tool_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.tool_id
		}
	`, tcName, rgID)
}

func testAccCheckIBMCdToolchainToolAppconfigDataSourceConfig(tcName string, rgID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}

		resource "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "name"
				region = "region"
				resource_group = "resource-group"
				instance_name = "instance-name"
				environment_name = "environment-name"
				collection_name = "collection-name"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig" {
			toolchain_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.toolchain_id
			tool_id = ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig.tool_id
		}
	`, tcName, rgID, getToolByIDResponseName)
}
