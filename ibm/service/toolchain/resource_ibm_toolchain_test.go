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
	"github.ibm.com/org-ids/toolchain-go-sdk/toolchainv2"
)

func TestAccIBMToolchainBasic(t *testing.T) {
	var conf toolchainv2.GetToolchainByIDResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroupID := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainConfigBasic(name, resourceGroupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainExists("ibm_toolchain.toolchain", conf),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", name),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "resource_group_id", resourceGroupID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMToolchainConfigBasic(nameUpdate, resourceGroupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "resource_group_id", resourceGroupID),
				),
			},
		},
	})
}

func TestAccIBMToolchainAllArgs(t *testing.T) {
	var conf toolchainv2.GetToolchainByIDResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroupID := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMToolchainConfig(name, resourceGroupID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMToolchainExists("ibm_toolchain.toolchain", conf),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", name),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMToolchainConfig(nameUpdate, resourceGroupID, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain.toolchain",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMToolchainConfigBasic(name string, resourceGroupID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain" "toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}
	`, name, resourceGroupID)
}

func testAccCheckIBMToolchainConfig(name string, resourceGroupID string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain" "toolchain" {
			name = "%s"
			resource_group_id = "%s"
			description = "%s"
		}
	`, name, resourceGroupID, description)
}

func testAccCheckIBMToolchainExists(n string, obj toolchainv2.GetToolchainByIDResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
		if err != nil {
			return err
		}

		getToolchainByIDOptions := &toolchainv2.GetToolchainByIDOptions{}

		getToolchainByIDOptions.SetToolchainID(rs.Primary.ID)

		getToolchainByIDResponse, _, err := toolchainClient.GetToolchainByID(getToolchainByIDOptions)
		if err != nil {
			return err
		}

		obj = *getToolchainByIDResponse
		return nil
	}
}

func testAccCheckIBMToolchainDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain" {
			continue
		}

		getToolchainByIDOptions := &toolchainv2.GetToolchainByIDOptions{}

		getToolchainByIDOptions.SetToolchainID(rs.Primary.ID)

		// Try to find the key
		_, response, err := toolchainClient.GetToolchainByID(getToolchainByIDOptions)

		if err == nil {
			return fmt.Errorf("toolchain still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
