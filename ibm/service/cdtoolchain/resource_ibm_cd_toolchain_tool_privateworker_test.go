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
	"github.ibm.com/org-ids/toolchain-go-sdk/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolPrivateworkerBasic(t *testing.T) {
	var conf cdtoolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolPrivateworkerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPrivateworkerConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolPrivateworkerExists("ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainToolPrivateworkerAllArgs(t *testing.T) {
	var conf cdtoolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolPrivateworkerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPrivateworkerConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolPrivateworkerExists("ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolPrivateworkerConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain_tool_privateworker.cd_toolchain_tool_privateworker",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolPrivateworkerConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIBMCdToolchainToolPrivateworkerConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain_tool_privateworker" "cd_toolchain_tool_privateworker" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				workerQueueCredentials = "workerQueueCredentials"
				workerQueueIdentifier = "workerQueueIdentifier"
			}
		}
	`, toolchainID, name)
}

func testAccCheckIBMCdToolchainToolPrivateworkerExists(n string, obj cdtoolchainv2.GetIntegrationByIDResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
		if err != nil {
			return err
		}

		getIntegrationByIDOptions := &cdtoolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		getIntegrationByIDResponse, _, err := cdToolchainClient.GetIntegrationByID(getIntegrationByIDOptions)
		if err != nil {
			return err
		}

		obj = *getIntegrationByIDResponse
		return nil
	}
}

func testAccCheckIBMCdToolchainToolPrivateworkerDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_privateworker" {
			continue
		}

		getIntegrationByIDOptions := &cdtoolchainv2.GetIntegrationByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIntegrationByIDOptions.SetToolchainID(parts[0])
		getIntegrationByIDOptions.SetIntegrationID(parts[1])

		// Try to find the key
		_, response, err := cdToolchainClient.GetIntegrationByID(getIntegrationByIDOptions)

		if err == nil {
			return fmt.Errorf("cd_toolchain_tool_privateworker still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain_tool_privateworker (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
