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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/mqcloud"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmMqcloudVirtualPrivateEndpointGatewaysDataSourceBasic(t *testing.T) {
	virtualPrivateEndpointGatewayDetailsServiceInstanceGuid := acc.MqcloudCapacityID
	virtualPrivateEndpointGatewayDetailsName := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	virtualPrivateEndpointGatewayDetailsTargetCrn := acc.MqCloudVirtualPrivateEndPointTargetCrn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudVirtualPrivateEndpointGatewaysDataSourceConfigBasic(virtualPrivateEndpointGatewayDetailsServiceInstanceGuid, virtualPrivateEndpointGatewayDetailsName, virtualPrivateEndpointGatewayDetailsTargetCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "virtual_private_endpoint_gateways.#"),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "virtual_private_endpoint_gateways.0.name", virtualPrivateEndpointGatewayDetailsName),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "virtual_private_endpoint_gateways.0.target_crn", virtualPrivateEndpointGatewayDetailsTargetCrn),
				),
			},
		},
	})
}

func TestAccIbmMqcloudVirtualPrivateEndpointGatewaysDataSourceAllArgs(t *testing.T) {
	virtualPrivateEndpointGatewayDetailsServiceInstanceGuid := acc.MqcloudCapacityID
	virtualPrivateEndpointGatewayDetailsTrustedProfile := acc.MqCloudVirtualPrivateEndPointTrustedProfile
	virtualPrivateEndpointGatewayDetailsName := fmt.Sprintf("tf-name-%d", acctest.RandIntRange(10, 100))
	virtualPrivateEndpointGatewayDetailsTargetCrn := acc.MqCloudVirtualPrivateEndPointTargetCrn

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudVirtualPrivateEndpointGatewaysDataSourceConfig(virtualPrivateEndpointGatewayDetailsServiceInstanceGuid, virtualPrivateEndpointGatewayDetailsTrustedProfile, virtualPrivateEndpointGatewayDetailsName, virtualPrivateEndpointGatewayDetailsTargetCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "trusted_profile"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "virtual_private_endpoint_gateways.#"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "virtual_private_endpoint_gateways.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "virtual_private_endpoint_gateways.0.id"),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "virtual_private_endpoint_gateways.0.name", virtualPrivateEndpointGatewayDetailsName),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "virtual_private_endpoint_gateways.0.target_crn", virtualPrivateEndpointGatewayDetailsTargetCrn),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_virtual_private_endpoint_gateways.mqcloud_virtual_private_endpoint_gateways_instance", "virtual_private_endpoint_gateways.0.status"),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudVirtualPrivateEndpointGatewaysDataSourceConfigBasic(virtualPrivateEndpointGatewayDetailsServiceInstanceGuid string, virtualPrivateEndpointGatewayDetailsName string, virtualPrivateEndpointGatewayDetailsTargetCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
			service_instance_guid = "%s"
			name = "%s"
			target_crn = "%s"
		}

		data "ibm_mqcloud_virtual_private_endpoint_gateways" "mqcloud_virtual_private_endpoint_gateways_instance" {
			service_instance_guid = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.service_instance_guid
			trusted_profile = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.trusted_profile
			name = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.name
		}
	`, virtualPrivateEndpointGatewayDetailsServiceInstanceGuid, virtualPrivateEndpointGatewayDetailsName, virtualPrivateEndpointGatewayDetailsTargetCrn)
}

func testAccCheckIbmMqcloudVirtualPrivateEndpointGatewaysDataSourceConfig(virtualPrivateEndpointGatewayDetailsServiceInstanceGuid string, virtualPrivateEndpointGatewayDetailsTrustedProfile string, virtualPrivateEndpointGatewayDetailsName string, virtualPrivateEndpointGatewayDetailsTargetCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
			service_instance_guid = "%s"
			trusted_profile = "%s"
			name = "%s"
			target_crn = "%s"
		}

		data "ibm_mqcloud_virtual_private_endpoint_gateways" "mqcloud_virtual_private_endpoint_gateways_instance" {
			service_instance_guid = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.service_instance_guid
			trusted_profile = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.trusted_profile
			name = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.name
		}
	`, virtualPrivateEndpointGatewayDetailsServiceInstanceGuid, virtualPrivateEndpointGatewayDetailsTrustedProfile, virtualPrivateEndpointGatewayDetailsName, virtualPrivateEndpointGatewayDetailsTargetCrn)
}

func TestDataSourceIbmMqcloudVirtualPrivateEndpointGatewaysVirtualPrivateEndpointGatewayDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "testString"
		model["id"] = "testString"
		model["name"] = "testString"
		model["target_crn"] = "testString"
		model["status"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(mqcloudv1.VirtualPrivateEndpointGatewayDetails)
	model.Href = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.TargetCrn = core.StringPtr("testString")
	model.Status = core.StringPtr("testString")

	result, err := mqcloud.DataSourceIbmMqcloudVirtualPrivateEndpointGatewaysVirtualPrivateEndpointGatewayDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
