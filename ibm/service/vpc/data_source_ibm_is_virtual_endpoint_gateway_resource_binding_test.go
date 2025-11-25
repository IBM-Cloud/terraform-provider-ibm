// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVirtualEndpointGatewayResourceBindingDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-eg-%d-test", acctest.RandIntRange(10, 100))
	endpointGatewayName := fmt.Sprintf("tf-eg-%d-test", acctest.RandIntRange(10, 100))
	endpointGatewayRBName := fmt.Sprintf("tf-rb-eg-%d-test", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualEndpointGatewayResourceBindingDataSourceConfigBasic(vpcname, endpointGatewayName, endpointGatewayRBName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "endpoint_gateway_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "endpoint_gateway_resource_binding_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "lifecycle_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "service_endpoint"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb_ds", "type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualEndpointGatewayResourceBindingDataSourceConfigBasic(vpcname, endpointGatewayName, endpointGatewayRBName string) string {
	return testAccCheckIBMIsVirtualEndpointGatewayResourceBindingConfigBasic(vpcname, endpointGatewayName, endpointGatewayRBName) + fmt.Sprintf(`
		data "ibm_is_virtual_endpoint_gateway_resource_binding" "testacc_vpe_rb_ds" {
			endpoint_gateway_id = ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb.endpoint_gateway_id
			endpoint_gateway_resource_binding_id = ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb.endpoint_gateway_resource_binding_id
		}
	`)
}
