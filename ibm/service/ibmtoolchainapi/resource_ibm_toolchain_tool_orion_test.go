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

func TestAccIbmToolchainToolOrionBasic(t *testing.T) {
	var conf ibmtoolchainapiv2.ServiceResponse
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainToolOrionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainToolOrionConfigBasic(toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainToolOrionExists("ibm_toolchain_tool_orion.toolchain_tool_orion", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_tool_orion.toolchain_tool_orion", "toolchain_id", toolchainID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_tool_orion.toolchain_tool_orion",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmToolchainToolOrionConfigBasic(toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_tool_orion" "toolchain_tool_orion" {
			toolchain_id = "%s"
		}
	`, toolchainID)
}

func testAccCheckIbmToolchainToolOrionExists(n string, obj ibmtoolchainapiv2.ServiceResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ibmToolchainApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IbmToolchainApiV2()
		if err != nil {
			return err
		}

		getServiceInstanceOptions := &ibmtoolchainapiv2.GetServiceInstanceOptions{}

		getServiceInstanceOptions.SetServiceInstanceID(rs.Primary.ID)

		serviceResponse, _, err := ibmToolchainApiClient.GetServiceInstance(getServiceInstanceOptions)
		if err != nil {
			return err
		}

		obj = *serviceResponse
		return nil
	}
}

func testAccCheckIbmToolchainToolOrionDestroy(s *terraform.State) error {
	ibmToolchainApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_tool_orion" {
			continue
		}

		getServiceInstanceOptions := &ibmtoolchainapiv2.GetServiceInstanceOptions{}

		getServiceInstanceOptions.SetServiceInstanceID(rs.Primary.ID)

		// Try to find the key
		_, response, err := ibmToolchainApiClient.GetServiceInstance(getServiceInstanceOptions)

		if err == nil {
			return fmt.Errorf("toolchain_tool_orion still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_tool_orion (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
