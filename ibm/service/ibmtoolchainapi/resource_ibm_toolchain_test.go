// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibmtoolchainapi_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/org-ids/toolchain-go-sdk/ibmtoolchainapiv2"
)

func TestAccIbmToolchainBasic(t *testing.T) {
	var conf ibmtoolchainapiv2.ToolchainResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	generator := "API"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainConfigBasic(name, generator),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainExists("ibm_toolchain.toolchain", conf),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", name),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "generator", generator),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmToolchainConfigBasic(nameUpdate, generator),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "generator", generator),
				),
			},
		},
	})
}

func TestAccIbmToolchainAllArgs(t *testing.T) {
	var conf ibmtoolchainapiv2.ToolchainResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	generator := "API"
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	key := fmt.Sprintf("tf_key_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainConfig(name, generator, description, key),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainExists("ibm_toolchain.toolchain", conf),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", name),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "generator", generator),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "description", description),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "key", key),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmToolchainConfig(nameUpdate, generator, descriptionUpdate, key),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "generator", generator),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_toolchain.toolchain", "key", key),
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

func testAccCheckIbmToolchainConfigBasic(name string, generator string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain" "toolchain" {
			name = "%s"
			generator = "%s"
		}
	`, name, generator)
}

func testAccCheckIbmToolchainConfig(name string, generator string, description string, key string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain" "toolchain" {
			name = "%s"
			generator = "%s"
			description = "%s"
			key = "%s"
			container {
				guid = "d02d29f1-e7bb-4977-8a6f-26d7b7bb893e"
				type = "organization_guid"
			}
		}
	`, name, generator, description, key)
}

func testAccCheckIbmToolchainExists(n string, obj ibmtoolchainapiv2.ToolchainResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ibmToolchainApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IbmToolchainApiV2()
		if err != nil {
			return err
		}

		getToolchainOptions := &ibmtoolchainapiv2.GetToolchainOptions{}

		getToolchainOptions.SetToolchainGuid(rs.Primary.ID)

		toolchainResponse, _, err := ibmToolchainApiClient.GetToolchain(getToolchainOptions)
		if err != nil {
			return err
		}

		obj = *toolchainResponse
		return nil
	}
}

func testAccCheckIbmToolchainDestroy(s *terraform.State) error {
	ibmToolchainApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain" {
			continue
		}

		getToolchainOptions := &ibmtoolchainapiv2.GetToolchainOptions{}

		getToolchainOptions.SetToolchainGuid(rs.Primary.ID)

		// Try to find the key
		_, response, err := ibmToolchainApiClient.GetToolchain(getToolchainOptions)

		if err == nil {
			return fmt.Errorf("toolchain still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
