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

func TestAccIBMCdToolchainToolSecuritycomplianceDataSourceBasic(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSecuritycomplianceDataSourceConfigBasic(getToolByIDResponseToolchainID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolSecuritycomplianceDataSourceAllArgs(t *testing.T) {
	getToolByIDResponseToolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	getToolByIDResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSecuritycomplianceDataSourceConfig(getToolByIDResponseToolchainID, getToolByIDResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "get_tool_by_id_response_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolSecuritycomplianceDataSourceConfigBasic(getToolByIDResponseToolchainID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
			toolchain_id = "%s"
		}

		data "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
			toolchain_id = ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID)
}

func testAccCheckIBMCdToolchainToolSecuritycomplianceDataSourceConfig(getToolByIDResponseToolchainID string, getToolByIDResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				evidence_repo_name = "evidence_repo_name"
				trigger_scan = "disabled"
				location = "location"
				evidence_namespace = "evidence_namespace"
				api-key = "api-key"
				scope = "scope"
				profile = "profile"
			}
		}

		data "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
			toolchain_id = ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance.toolchain_id
			tool_id = "tool_id"
		}
	`, getToolByIDResponseToolchainID, getToolByIDResponseName)
}
