// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolPipelineBasic(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgID := acc.CdResourceGroupID
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolPipelineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPipelineConfigBasic(tcName, rgID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolPipelineExists("ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolPipelineAllArgs(t *testing.T) {
	var conf cdtoolchainv2.ToolchainTool
	rgID := acc.CdResourceGroupID
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolPipelineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPipelineConfig(tcName, rgID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolPipelineExists("ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPipelineConfig(tcName, rgID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "toolchain_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolPipelineConfigBasic(tcName string, rgID string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}

		resource "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "name"
				type = "tekton"
				ui_pipeline = true
			}
		}
	`, tcName, rgID)
}

func testAccCheckIBMCdToolchainToolPipelineConfig(tcName string, rgID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}

		resource "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
				toolchain_id = ibm_cd_toolchain.cd_toolchain.id
				parameters {
					name = "name"
					type = "tekton"
					ui_pipeline = true
				}
				name = "%s"
		}
	`, tcName, rgID, name)
}

func testAccCheckIBMCdToolchainToolPipelineExists(n string, obj cdtoolchainv2.ToolchainTool) resource.TestCheckFunc {

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

		getToolByIDResponse, _, err := cdToolchainClient.GetToolByID(getToolByIDOptions)
		if err != nil {
			return err
		}

		obj = *getToolByIDResponse
		return nil
	}
}

func testAccCheckIBMCdToolchainToolPipelineDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_pipeline" {
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
			return fmt.Errorf("cd_toolchain_tool_pipeline still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_pipeline (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
