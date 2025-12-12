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

func TestAccIBMIsVirtualEndpointGatewayResourceBindingsDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-eg-%d-test", acctest.RandIntRange(10, 100))
	endpointGatewayName := fmt.Sprintf("tf-eg-%d-test", acctest.RandIntRange(10, 100))
	endpointGatewayRBName := fmt.Sprintf("tf-rb-eg-%d-test", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualEndpointGatewayResourceBindingsDataSourceConfigBasic(vpcname, endpointGatewayName, endpointGatewayRBName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "endpoint_gateway_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.0.service_endpoint"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.0.target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_virtual_endpoint_gateway_resource_bindings.testacc_vpe_rb_ds", "resource_bindings.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualEndpointGatewayResourceBindingsDataSourceConfigBasic(vpcname, endpointGatewayName, endpointGatewayRBName string) string {
	return testAccCheckIBMIsVirtualEndpointGatewayResourceBindingConfigBasic(vpcname, endpointGatewayName, endpointGatewayRBName) + fmt.Sprintf(`
		data "ibm_is_virtual_endpoint_gateway_resource_bindings" "testacc_vpe_rb_ds" {
			depends_on = [ ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb ]
			endpoint_gateway_id = ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb.endpoint_gateway_id
		}
	`)
}
