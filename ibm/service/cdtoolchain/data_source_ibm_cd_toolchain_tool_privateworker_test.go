// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolPrivateworkerDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPrivateworkerDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolPrivateworkerDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPrivateworkerDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolPrivateworkerDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "private-worker-tool-01"
				worker_queue_credentials = "<worker_queue_credentials>"
			}
		}

		data "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker.toolchain_id
			tool_id = ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolPrivateworkerDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "private-worker-tool-01"
				worker_queue_credentials = "<worker_queue_credentials>"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker.toolchain_id
			tool_id = ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker.tool_id
		}
	`, rgName, tcName, toolName)
}
