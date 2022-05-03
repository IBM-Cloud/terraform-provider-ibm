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

func TestAccIBMToolchainToolJenkinsBasic(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolJenkinsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolJenkinsConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolJenkinsExists("ibm_toolchain_tool_jenkins.toolchain_tool_jenkins", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_jenkins.toolchain_tool_jenkins", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMToolchainToolJenkinsAllArgs(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolJenkinsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolJenkinsConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolJenkinsExists("ibm_toolchain_tool_jenkins.toolchain_tool_jenkins", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_jenkins.toolchain_tool_jenkins", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_jenkins.toolchain_tool_jenkins", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolJenkinsConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain_tool_jenkins.toolchain_tool_jenkins", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_jenkins.toolchain_tool_jenkins", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_jenkins.toolchain_tool_jenkins",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMToolchainToolJenkinsConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_jenkins" "toolchain_tool_jenkins" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIBMToolchainToolJenkinsConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_jenkins" "toolchain_tool_jenkins" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				dashboard_url = "dashboard_url"
				webhook_url = "webhook_url"
				api_user_name = "api_user_name"
				api_token = "api_token"
			}
		}
	`, toolchainID, name)
}

func testAccCheckIBMToolchainToolJenkinsExists(n string, obj toolchainv2.GetIntegrationByIDResponse) resource.TestCheckFunc {

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

func testAccCheckIBMToolchainToolJenkinsDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_jenkins" {
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
			return fmt.Errorf("toolchain_tool_jenkins still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_jenkins (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
