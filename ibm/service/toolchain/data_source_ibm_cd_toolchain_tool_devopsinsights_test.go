// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainToolDevopsinsightsDataSourceBasic(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolDevopsinsightsDataSourceConfigBasic(getIntegrationByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "get_integration_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolDevopsinsightsDataSourceAllArgs(t *testing.T) {
	getIntegrationByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getIntegrationByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolDevopsinsightsDataSourceConfig(getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "integration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "get_integration_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolDevopsinsightsDataSourceConfigBasic(getIntegrationByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_devopsinsights" "cd_toolchain_tool_devopsinsights" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_devopsinsights" "cd_toolchain_tool_devopsinsights" {
			toolchain_id = ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolDevopsinsightsDataSourceConfig(getIntegrationByIDResponseToolchainID string, getIntegrationByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_devopsinsights" "cd_toolchain_tool_devopsinsights" {
			toolchain_id = "%s"
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_devopsinsights" "cd_toolchain_tool_devopsinsights" {
			toolchain_id = ibm_cd_toolchain_tool_devopsinsights.cd_toolchain_tool_devopsinsights.toolchain_id
			integration_id = "integration_id"
		}
	`, getIntegrationByIDResponseToolchainID, getIntegrationByIDResponseName)
}
