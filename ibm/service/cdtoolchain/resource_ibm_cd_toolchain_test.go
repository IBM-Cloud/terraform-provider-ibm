// Copyright IBM Corp. 2022, 2023 All Rights Reserved.
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
)

func TestAccIBMCdToolchainBasic(t *testing.T) {
	var conf cdtoolchainv2.Toolchain
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainConfigBasic(name, rgName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainExists("ibm_cd_toolchain.cd_toolchain", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "name", name),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain.cd_toolchain", "resource_group_id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainConfigBasic(nameUpdate, rgName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain.cd_toolchain", "resource_group_id"),
				),
			},
		},
	})
}

func TestAccIBMCdToolchainAllArgs(t *testing.T) {
	var conf cdtoolchainv2.Toolchain
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainConfig(name, rgName, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainExists("ibm_cd_toolchain.cd_toolchain", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "name", name),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain.cd_toolchain", "resource_group_id"),
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainConfig(nameUpdate, rgName, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_toolchain.cd_toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttrSet("ibm_cd_toolchain.cd_toolchain", "resource_group_id"),
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

func testAccCheckIBMCdToolchainConfigBasic(name string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}
	  
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}
	`, rgName, name)
}

func testAccCheckIBMCdToolchainConfig(name string, rgName string, description string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}

		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
			description = "%s"
		}
	`, rgName, name, description)
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

		toolchain, _, err := cdToolchainClient.GetToolchainByID(getToolchainByIDOptions)
		if err != nil {
			return err
		}

		obj = *toolchain
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
