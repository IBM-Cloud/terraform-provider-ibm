// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtoolchainv2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMCdToolchainToolBitbucketgitBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdBitbucketRepoUrl

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolBitbucketgitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolBitbucketgitConfigBasic(tcName, rgName, repoUrl),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolBitbucketgitExists("ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolBitbucketgitAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	repoUrl := acc.CdBitbucketRepoUrl
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolBitbucketgitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolBitbucketgitConfig(tcName, rgName, repoUrl, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolBitbucketgitExists("ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolBitbucketgitConfig(tcName, rgName, repoUrl, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolBitbucketgitConfigBasic(tcName string, rgName string, repoUrl string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_bitbucketgit" "cd_toolchain_tool_bitbucketgit" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				toolchain_issues_enabled = true
				enable_traceability = true
			}
			initialization {
				repo_url = "%s"
				type = "link"
			}
		}
	`, rgName, tcName, repoUrl)
}

func testAccCheckIBMCdToolchainToolBitbucketgitConfig(tcName string, rgName string, repoUrl string, name string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		resource "ibm_cd_toolchain_tool_bitbucketgit" "cd_toolchain_tool_bitbucketgit" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				toolchain_issues_enabled = true
				enable_traceability = true
			}
			initialization {
				repo_url = "%s"
				type = "link"
			}
			name = "%s"
		}
	`, rgName, tcName, repoUrl, name)
}

func testAccCheckIBMCdToolchainToolBitbucketgitExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
		if err != nil {
			return err
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		toolchainTool, _, err := cdToolchainClient.GetToolByID(getToolByIDOptions)
		if err != nil {
			return err
		}

		obj = *toolchainTool
		return nil
	}
}

func testAccCheckIBMCdToolchainToolBitbucketgitDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_bitbucketgit" {
			continue
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		// Try to find the key
		_, response, err := cdToolchainClient.GetToolByID(getToolByIDOptions)

		if err == nil {
			return fmt.Errorf("cd_toolchain_tool_bitbucketgit still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_bitbucketgit (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
