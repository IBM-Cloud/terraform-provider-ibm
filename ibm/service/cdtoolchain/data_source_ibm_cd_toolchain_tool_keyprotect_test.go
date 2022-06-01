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

func TestAccIBMCdToolchainToolKeyprotectDataSourceBasic(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfigBasic(getIntegrationByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "get_integration_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolKeyprotectDataSourceAllArgs(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getIntegrationByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfig(getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "get_integration_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfigBasic(getIntegrationByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolKeyprotectDataSourceConfig(getIntegrationByIDResponseToolchainID string, getIntegrationByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				region = "region"
				resource-group = "resource-group"
				instance-name = "instance-name"
				integration-status = "integration-status"
			}
		}

		data "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
			toolchain_id = ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName)
}
