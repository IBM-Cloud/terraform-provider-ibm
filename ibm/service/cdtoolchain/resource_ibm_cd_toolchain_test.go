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
	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
)

func TestAccIBMCdToolchainBasic(t *testing.T) {
	var conf cdtoolchainv2.Toolchain
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroupID := acc.CdResourceGroupID
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainConfigBasic(name, resourceGroupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainExists("ibm_cd_toolchain.cd_toolchain", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "resource_group_id", resourceGroupID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainConfigBasic(nameUpdate, resourceGroupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "resource_group_id", resourceGroupID),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainAllArgs(t *testing.T) {
	var conf cdtoolchainv2.Toolchain
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroupID := acc.CdResourceGroupID
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainConfig(name, resourceGroupID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainExists("ibm_cd_toolchain.cd_toolchain", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainConfig(nameUpdate, resourceGroupID, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_toolchain.cd_toolchain",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdToolchainConfigBasic(name string, resourceGroupID string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}
	`, name, resourceGroupID)
}

func testAccCheckIBMCdToolchainConfig(name string, resourceGroupID string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
			description = "%s"
		}
	`, name, resourceGroupID, description)
}

func testAccCheckIBMCdToolchainExists(n string, obj cdtoolchainv2.Toolchain) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
		if err != nil {
			return err
		}

		getToolchainByIDOptions := &cdtoolchainv2.GetToolchainByIDOptions{}

		getToolchainByIDOptions.SetToolchainID(rs.Primary.ID)

		getToolchainByIDResponse, _, err := cdToolchainClient.GetToolchainByID(getToolchainByIDOptions)
		if err != nil {
			return err
		}

		obj = *getToolchainByIDResponse
		return nil
	}
}

func testAccCheckIBMCdToolchainDestroy(s *terraform.State) error {
	cdToolchainClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain" {
			continue
		}

		getToolchainByIDOptions := &cdtoolchainv2.GetToolchainByIDOptions{}

		getToolchainByIDOptions.SetToolchainID(rs.Primary.ID)

		// Try to find the key
		_, response, err := cdToolchainClient.GetToolchainByID(getToolchainByIDOptions)

		if err == nil {
			return fmt.Errorf("cd_toolchain still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_toolchain (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
