// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmMqcloudVirtualPrivateEndpointGatewayDataSourceBasic(t *testing.T) {
	virtualPrivateEndpointGatewayDetailsServiceInstanceGuid := acc.MqcloudCapacityID
	virtualPrivateEndpointGatewayDetailsName := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	virtualPrivateEndpointGatewayDetailsTargetCrn := acc.MqCloudVirtualPrivateEndPointTargetCrn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayDataSourceConfigBasic(virtualPrivateEndpointGatewayDetailsServiceInstanceGuid, virtualPrivateEndpointGatewayDetailsName, virtualPrivateEndpointGatewayDetailsTargetCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "virtual_private_endpoint_gateway_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "target_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "status"),
				),
			},
		},
	})
}

func TestAccIbmMqcloudVirtualPrivateEndpointGatewayDataSourceAllArgs(t *testing.T) {
	virtualPrivateEndpointGatewayDetailsServiceInstanceGuid := acc.MqcloudCapacityID
	virtualPrivateEndpointGatewayDetailsTrustedProfile := acc.MqCloudVirtualPrivateEndPointTrustedProfile
	virtualPrivateEndpointGatewayDetailsName := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	virtualPrivateEndpointGatewayDetailsTargetCrn := acc.MqCloudVirtualPrivateEndPointTargetCrn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayDataSourceConfig(virtualPrivateEndpointGatewayDetailsServiceInstanceGuid, virtualPrivateEndpointGatewayDetailsTrustedProfile, virtualPrivateEndpointGatewayDetailsName, virtualPrivateEndpointGatewayDetailsTargetCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "virtual_private_endpoint_gateway_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "trusted_profile"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "target_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "status"),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayDataSourceConfigBasic(virtualPrivateEndpointGatewayDetailsServiceInstanceGuid string, virtualPrivateEndpointGatewayDetailsName string, virtualPrivateEndpointGatewayDetailsTargetCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
			service_instance_guid = "%s"
			name = "%s"
			target_crn = "%s"
		}

		data "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
			service_instance_guid = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.service_instance_guid
			virtual_private_endpoint_gateway_guid = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.virtual_private_endpoint_gateway_guid
			trusted_profile = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.trusted_profile
		}
	`, virtualPrivateEndpointGatewayDetailsServiceInstanceGuid, virtualPrivateEndpointGatewayDetailsName, virtualPrivateEndpointGatewayDetailsTargetCrn)
}

func testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayDataSourceConfig(virtualPrivateEndpointGatewayDetailsServiceInstanceGuid string, virtualPrivateEndpointGatewayDetailsTrustedProfile string, virtualPrivateEndpointGatewayDetailsName string, virtualPrivateEndpointGatewayDetailsTargetCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
			service_instance_guid = "%s"
			trusted_profile = "%s"
			name = "%s"
			target_crn = "%s"
		}

		data "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
			service_instance_guid = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.service_instance_guid
			virtual_private_endpoint_gateway_guid = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.virtual_private_endpoint_gateway_guid
			trusted_profile = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.trusted_profile
		}
	`, virtualPrivateEndpointGatewayDetailsServiceInstanceGuid, virtualPrivateEndpointGatewayDetailsTrustedProfile, virtualPrivateEndpointGatewayDetailsName, virtualPrivateEndpointGatewayDetailsTargetCrn)
}
