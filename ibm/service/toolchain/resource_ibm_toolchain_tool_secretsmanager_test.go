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

func TestAccIBMToolchainToolSecretsmanagerBasic(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolSecretsmanagerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolSecretsmanagerConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolSecretsmanagerExists("ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIBMToolchainToolSecretsmanagerAllArgs(t *testing.T) {
	var conf toolchainv2.GetIntegrationByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainToolSecretsmanagerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolSecretsmanagerConfig(toolchainID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainToolSecretsmanagerExists("ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMToolchainToolSecretsmanagerConfig(toolchainID, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMToolchainToolSecretsmanagerConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_secretsmanager" "toolchain_tool_secretsmanager" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIBMToolchainToolSecretsmanagerConfig(toolchainID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_secretsmanager" "toolchain_tool_secretsmanager" {
			toolchain_id = "%s"
			name = "%s"
			parameters {
				name = "name"
				region = "region"
				resource-group = "resource-group"
				instance-name = "instance-name"
				integration-status = "integration-status"
			}
		}
	`, toolchainID, name)
}

func testAccCheckIBMToolchainToolSecretsmanagerExists(n string, obj toolchainv2.GetIntegrationByIDResponse) resource.TestCheckFunc {

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

func testAccCheckIBMToolchainToolSecretsmanagerDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_secretsmanager" {
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
			return fmt.Errorf("toolchain_tool_secretsmanager still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_secretsmanager (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
