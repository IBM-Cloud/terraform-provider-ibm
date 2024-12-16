// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
)

func TestAccIbmMqcloudVirtualPrivateEndpointGatewayBasic(t *testing.T) {
	t.Parallel()
	var conf mqcloudv1.VirtualPrivateEndpointGatewayDetails
	serviceInstanceGuid := acc.MqcloudCapacityID
	name := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	targetCrn := acc.MqCloudVirtualPrivateEndPointTargetCrn

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckMqcloud(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayConfigBasic(serviceInstanceGuid, name, targetCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayExists("ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", conf),
					resource.TestCheckResourceAttr("ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "target_crn", targetCrn),
				),
			},
		},
	})
}

func TestAccIbmMqcloudVirtualPrivateEndpointGatewayAllArgs(t *testing.T) {
	t.Parallel()
	var conf mqcloudv1.VirtualPrivateEndpointGatewayDetails
	serviceInstanceGuid := acc.MqcloudCapacityID
	trustedProfile := acc.MqCloudVirtualPrivateEndPointTrustedProfile
	name := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	targetCrn := acc.MqCloudVirtualPrivateEndPointTargetCrn

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayConfig(serviceInstanceGuid, trustedProfile, name, targetCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayExists("ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", conf),
					resource.TestCheckResourceAttr("ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "trusted_profile", trustedProfile),
					resource.TestCheckResourceAttr("ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance", "target_crn", targetCrn),
				),
			},
			{
				ResourceName:            "ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"trusted_profile", "service_instance_guid"},
			},
		},
	})
}

func testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayConfigBasic(serviceInstanceGuid string, name string, targetCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
			service_instance_guid = "%s"
			name = "%s"
			target_crn = "%s"
		}
	`, serviceInstanceGuid, name, targetCrn)
}

func testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayConfig(serviceInstanceGuid string, trustedProfile string, name string, targetCrn string) string {
	return fmt.Sprintf(`

		resource "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
			service_instance_guid = "%s"
			trusted_profile = "%s"
			name = "%s"
			target_crn = "%s"
		}
	`, serviceInstanceGuid, trustedProfile, name, targetCrn)
}

func testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayExists(n string, obj mqcloudv1.VirtualPrivateEndpointGatewayDetails) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
		if err != nil {
			return err
		}

		getVirtualPrivateEndpointGatewayOptions := &mqcloudv1.GetVirtualPrivateEndpointGatewayOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVirtualPrivateEndpointGatewayOptions.SetServiceInstanceGuid(parts[0])
		getVirtualPrivateEndpointGatewayOptions.SetVirtualPrivateEndpointGatewayGuid(parts[1])

		virtualPrivateEndpointGatewayDetails, _, err := mqcloudClient.GetVirtualPrivateEndpointGateway(getVirtualPrivateEndpointGatewayOptions)
		if err != nil {
			return err
		}

		obj = *virtualPrivateEndpointGatewayDetails
		return nil
	}
}

func testAccCheckIbmMqcloudVirtualPrivateEndpointGatewayDestroy(s *terraform.State) error {
	mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_mqcloud_virtual_private_endpoint_gateway" {
			continue
		}

		getVirtualPrivateEndpointGatewayOptions := &mqcloudv1.GetVirtualPrivateEndpointGatewayOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVirtualPrivateEndpointGatewayOptions.SetServiceInstanceGuid(parts[0])
		getVirtualPrivateEndpointGatewayOptions.SetVirtualPrivateEndpointGatewayGuid(parts[1])

		// Try to find the key
		_, response, err := mqcloudClient.GetVirtualPrivateEndpointGateway(getVirtualPrivateEndpointGatewayOptions)

		if err == nil {
			return fmt.Errorf("mqcloud_virtual_private_endpoint_gateway still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for mqcloud_virtual_private_endpoint_gateway (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
