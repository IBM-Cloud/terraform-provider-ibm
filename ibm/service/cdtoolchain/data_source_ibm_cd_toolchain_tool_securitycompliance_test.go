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
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSecuritycomplianceDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "tool_id"),
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
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolSecuritycomplianceDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance", "tool_id"),
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

func testAccCheckIBMCdToolchainToolSecuritycomplianceDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "compliance"
				evidence_namespace = "cd"
				evidence_repo_url = "https://github.example.com/<username>/compliance-evidence-<datestamp>"
			}
		}

		data "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
			toolchain_id = ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance.toolchain_id
			tool_id = ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolSecuritycomplianceDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "compliance"
				evidence_namespace = "cd"
				evidence_repo_url = "https://github.example.com/<username>/compliance-evidence-<datestamp>"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
			toolchain_id = ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance.toolchain_id
			tool_id = ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance.tool_id
		}
	`, rgName, tcName, toolName)
}
