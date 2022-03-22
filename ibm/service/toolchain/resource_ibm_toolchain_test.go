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

func TestAccIbmToolchainBasic(t *testing.T) {
	var conf toolchainv2.GetToolchainByIdResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroupID := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	generator := fmt.Sprintf("tf_generator_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	generatorUpdate := fmt.Sprintf("tf_generator_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainConfigBasic(name, resourceGroupID, generator),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainExists("ibm_toolchain.toolchain", conf),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", name),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "generator", generator),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmToolchainConfigBasic(nameUpdate, resourceGroupID, generatorUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "generator", generatorUpdate),
				),
			},
		},
	})
}

func TestAccIbmToolchainAllArgs(t *testing.T) {
	var conf toolchainv2.GetToolchainByIdResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroupID := fmt.Sprintf("tf_resource_group_id_%d", acctest.RandIntRange(10, 100))
	generator := fmt.Sprintf("tf_generator_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	generatorUpdate := fmt.Sprintf("tf_generator_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainConfig(name, resourceGroupID, generator, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainExists("ibm_toolchain.toolchain", conf),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", name),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "generator", generator),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmToolchainConfig(nameUpdate, resourceGroupID, generatorUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "generator", generatorUpdate),
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

func testAccCheckIbmToolchainConfigBasic(name string, resourceGroupID string, generator string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain" "toolchain" {
			name = "%s"
			resource_group_id = "%s"
			generator = "%s"
		}
	`, name, resourceGroupID, generator)
}

func testAccCheckIbmToolchainConfig(name string, resourceGroupID string, generator string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain" "toolchain" {
			name = "%s"
			resource_group_id = "%s"
			generator = "%s"
			description = "%s"
			template = "FIXME"
		}
	`, name, resourceGroupID, generator, description)
}

func testAccCheckIbmToolchainExists(n string, obj toolchainv2.GetToolchainByIdResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
		if err != nil {
			return err
		}

		getToolchainByIdOptions := &toolchainv2.GetToolchainByIdOptions{}

		getToolchainByIdOptions.SetToolchainID(rs.Primary.ID)

		getToolchainByIdResponse, _, err := toolchainClient.GetToolchainByID(getToolchainByIdOptions)
		if err != nil {
			return err
		}

		obj = *getToolchainByIdResponse
		return nil
	}
}

func testAccCheckIbmToolchainDestroy(s *terraform.State) error {
	toolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain" {
			continue
		}

		getToolchainByIdOptions := &toolchainv2.GetToolchainByIdOptions{}

		getToolchainByIdOptions.SetToolchainID(rs.Primary.ID)

		// Try to find the key
		_, response, err := toolchainClient.GetToolchainByID(getToolchainByIdOptions)

		if err == nil {
			return fmt.Errorf("toolchain still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
