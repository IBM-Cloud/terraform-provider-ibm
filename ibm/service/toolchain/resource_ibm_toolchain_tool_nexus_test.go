// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package toolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/org-ids/toolchain-go-sdk/toolchainv2"
)

func TestAccIBMToolchainToolNexusBasic(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolNexusDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolNexusConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolNexusExists("ibm_toolchain_tool_nexus.toolchain_tool_nexus", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_nexus.toolchain_tool_nexus", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMToolchainToolNexusAllArgs(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolNexusDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolNexusConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolNexusExists("ibm_toolchain_tool_nexus.toolchain_tool_nexus", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_nexus.toolchain_tool_nexus", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_nexus.toolchain_tool_nexus", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolNexusConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain_tool_nexus.toolchain_tool_nexus", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_nexus.toolchain_tool_nexus", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_nexus.toolchain_tool_nexus",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMToolchainToolNexusConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_nexus" "toolchain_tool_nexus" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIBMToolchainToolNexusConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_nexus" "toolchain_tool_nexus" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				dashboard_url = "dashboard_url"
				type = "npm"
				user_id = "user_id"
				token = "token"
				release_url = "release_url"
				mirror_url = "mirror_url"
				snapshot_url = "snapshot_url"
			}
		}
	`, toolchainID, name)
}

func testAccCheckIBMToolchainToolNexusExists(n string, obj toolchainv2.GetIntegrationByIDResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
		if err != nil {
			return err
		}

		getIntegrationByIDOptions := &toolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		getIntegrationByIDResponse, _, err := toolchainClient.GetIntegrationByID(getIntegrationByIDOptions)
		if err != nil {
			return err
		}

		obj = *getIntegrationByIDResponse
		return nil
	}
}

func testAccCheckIBMToolchainToolNexusDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_nexus" {
			continue
		}

		getIntegrationByIDOptions := &toolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		// Try to find the key
		_, response, err := toolchainClient.GetIntegrationByID(getIntegrationByIDOptions)

		if err == nil {
			return fmt.Errorf("toolchain_tool_nexus still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_nexus (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
