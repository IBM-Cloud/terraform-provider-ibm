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

func TestAccIbmToolchainIntegrationBasic(t *testing.T) {
	var conf ibmtoolchainapiv2.ServiceResponse
	serviceID := fmt.Sprintf("tf_service_id_%d", acctest.RandIntRange(10, 100))
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainIntegrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainIntegrationConfigBasic(serviceID, toolchainID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainIntegrationExists("ibm_toolchain_integration.toolchain_integration", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_integration.toolchain_integration", "service_id", serviceID),
					resource.TestCheckResourceAttr("ibm_toolchain_integration.toolchain_integration", "toolchain_id", toolchainID),
				),
			},
		},
	})
}

func TestAccIbmToolchainIntegrationAllArgs(t *testing.T) {
	var conf ibmtoolchainapiv2.ServiceResponse
	serviceID := fmt.Sprintf("tf_service_id_%d", acctest.RandIntRange(10, 100))
	toolchainID := fmt.Sprintf("tf_toolchain_id_%d", acctest.RandIntRange(10, 100))
	parameters := fmt.Sprintf("tf_parameters_%d", acctest.RandIntRange(10, 100))
	parametersUpdate := fmt.Sprintf("tf_parameters_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmToolchainIntegrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmToolchainIntegrationConfig(serviceID, toolchainID, parameters),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmToolchainIntegrationExists("ibm_toolchain_integration.toolchain_integration", conf),
					resource.TestCheckResourceAttr("ibm_toolchain_integration.toolchain_integration", "service_id", serviceID),
					resource.TestCheckResourceAttr("ibm_toolchain_integration.toolchain_integration", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_integration.toolchain_integration", "parameters", parameters),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmToolchainIntegrationConfig(serviceID, toolchainID, parametersUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_toolchain_integration.toolchain_integration", "service_id", serviceID),
					resource.TestCheckResourceAttr("ibm_toolchain_integration.toolchain_integration", "toolchain_id", toolchainID),
					resource.TestCheckResourceAttr("ibm_toolchain_integration.toolchain_integration", "parameters", parametersUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_toolchain_integration.toolchain_integration",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmToolchainIntegrationConfigBasic(serviceID string, toolchainID string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_integration" "toolchain_integration" {
			service_id = "%s"
			toolchain_id = "%s"
		}
	`, serviceID, toolchainID)
}

func testAccCheckIbmToolchainIntegrationConfig(serviceID string, toolchainID string, parameters string) string {
	return fmt.Sprintf(`

		resource "ibm_toolchain_integration" "toolchain_integration" {
			service_id = "%s"
			toolchain_id = "%s"
			parameters = "%s"
			parameters_references = "FIXME"
			container {
				guid = "d02d29f1-e7bb-4977-8a6f-26d7b7bb893e"
				type = "organization_guid"
			}
		}
	`, serviceID, toolchainID, parameters)
}

func testAccCheckIbmToolchainIntegrationExists(n string, obj ibmtoolchainapiv2.ServiceResponse) resource.TestCheckFunc {

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

func testAccCheckIbmToolchainIntegrationDestroy(s *terraform.State) error {
	ibmToolchainApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IbmToolchainApiV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_toolchain_integration" {
			continue
		}

		getServiceInstanceOptions := &ibmtoolchainapiv2.GetServiceInstanceOptions{}

		getServiceInstanceOptions.SetServiceInstanceID(rs.Primary.ID)

		// Try to find the key
		_, response, err := ibmToolchainApiClient.GetServiceInstance(getServiceInstanceOptions)

		if err == nil {
			return fmt.Errorf("toolchain_integration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for toolchain_integration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
