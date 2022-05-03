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

func TestAccIBMToolchainToolRationalteamconcertBasic(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolRationalteamconcertDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolRationalteamconcertConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolRationalteamconcertExists("ibm_toolchain_tool_rationalteamconcert.toolchain_tool_rationalteamconcert", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_rationalteamconcert.toolchain_tool_rationalteamconcert", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMToolchainToolRationalteamconcertAllArgs(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolRationalteamconcertDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolRationalteamconcertConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolRationalteamconcertExists("ibm_toolchain_tool_rationalteamconcert.toolchain_tool_rationalteamconcert", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_rationalteamconcert.toolchain_tool_rationalteamconcert", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_rationalteamconcert.toolchain_tool_rationalteamconcert", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolRationalteamconcertConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain_tool_rationalteamconcert.toolchain_tool_rationalteamconcert", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_rationalteamconcert.toolchain_tool_rationalteamconcert", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_rationalteamconcert.toolchain_tool_rationalteamconcert",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMToolchainToolRationalteamconcertConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_rationalteamconcert" "toolchain_tool_rationalteamconcert" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIBMToolchainToolRationalteamconcertConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_rationalteamconcert" "toolchain_tool_rationalteamconcert" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				server_url = "server_url"
				user_id = "user_id"
				password = "password"
				type = "new"
				project_area = "project_area"
				process_template = "process_template"
				enable_traceability = true
			}
		}
	`, toolchainID, name)
}

func testAccCheckIBMToolchainToolRationalteamconcertExists(n string, obj toolchainv2.GetIntegrationByIDResponse) resource.TestCheckFunc {

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

func testAccCheckIBMToolchainToolRationalteamconcertDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_rationalteamconcert" {
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
			return fmt.Errorf("toolchain_tool_rationalteamconcert still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_rationalteamconcert (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
