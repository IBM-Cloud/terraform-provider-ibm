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

func TestAccIBMCdToolchainToolJenkinsDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJenkinsDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "state"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolJenkinsDataSourceAllArgs(t *testing.T) {
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	toolName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolJenkinsDataSourceConfig(tcName, rgName, toolName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "tool_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "toolchain_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "referent.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolJenkinsDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "jenkins-tool-01"
				dashboard_url = "https://jenkins.mycompany.example.com"
				api_user_name = "<api_user_name>"
				api_token = "<api_token>"
			}
		}

		data "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins.toolchain_id
			tool_id = ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins.tool_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdToolchainToolJenkinsDataSourceConfig(tcName string, rgName string, toolName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "jenkins-tool-01"
				dashboard_url = "https://jenkins.mycompany.example.com"
				api_user_name = "<api_user_name>"
				api_token = "<api_token>"
			}
			name = "%s"
		}

		data "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins" {
			toolchain_id = ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins.toolchain_id
			tool_id = ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins.tool_id
		}
	`, rgName, tcName, toolName)
}
