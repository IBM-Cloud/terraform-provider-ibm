// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVirtualEndpointGatewayResourceBindingBasic(t *testing.T) {
	var conf vpcv1.EndpointGatewayResourceBinding
	vpcname := fmt.Sprintf("tf-vpc-eg-%d-test", acctest.RandIntRange(10, 100))
	endpointGatewayName := fmt.Sprintf("tf-eg-%d-test", acctest.RandIntRange(10, 100))
	endpointGatewayRBName := fmt.Sprintf("tf-rb-eg-%d-test", acctest.RandIntRange(10, 100))
	endpointGatewayRBNameUpdated := fmt.Sprintf("tf-rb-eg-updated-%d-test", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualEndpointGatewayResourceBindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualEndpointGatewayResourceBindingConfigBasic(vpcname, endpointGatewayName, endpointGatewayRBName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualEndpointGatewayResourceBindingExists("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "endpoint_gateway_id"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "endpoint_gateway_resource_binding_id"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "service_endpoint"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "type"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "target.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "target.0.crn"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "name", endpointGatewayRBName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualEndpointGatewayResourceBindingConfigBasic(vpcname, endpointGatewayName, endpointGatewayRBNameUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualEndpointGatewayResourceBindingExists("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "endpoint_gateway_id"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "endpoint_gateway_resource_binding_id"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "service_endpoint"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "type"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "target.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "target.0.crn"),
					resource.TestCheckResourceAttr("ibm_is_virtual_endpoint_gateway_resource_binding.testacc_vpe_rb", "name", endpointGatewayRBNameUpdated),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualEndpointGatewayResourceBindingConfigBasic(vpcName, egname, egResourceBindingName string) string {
	return fmt.Sprintf(`

		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
		resource "ibm_is_virtual_endpoint_gateway" "testacc_vpe" {
			name = "%s"
			vpc  = ibm_is_vpc.testacc_vpc.id

			target {
				resource_type = "%s"
				crn           = "%s"
			}
			dns_resolution_binding_mode = "disabled"
		}

		resource "ibm_is_virtual_endpoint_gateway_resource_binding" "testacc_vpe_rb" {
			name = "%s"
			endpoint_gateway_id  = ibm_is_virtual_endpoint_gateway.testacc_vpe.id
			target {
				crn = "%s"
			}
		}
	`, vpcName, egname, acc.IsEndpointGatewayTargetType, acc.IsEndpointGatewayTargetCRN, egResourceBindingName, acc.IsResourceBindingCRN)
}

func testAccCheckIBMIsVirtualEndpointGatewayResourceBindingExists(n string, obj vpcv1.EndpointGatewayResourceBinding) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getEndpointGatewayResourceBindingOptions := &vpcv1.GetEndpointGatewayResourceBindingOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getEndpointGatewayResourceBindingOptions.SetEndpointGatewayID(parts[0])
		getEndpointGatewayResourceBindingOptions.SetID(parts[1])

		endpointGatewayResourceBinding, _, err := vpcClient.GetEndpointGatewayResourceBinding(getEndpointGatewayResourceBindingOptions)
		if err != nil {
			return err
		}

		obj = *endpointGatewayResourceBinding
		return nil
	}
}

func testAccCheckIBMIsVirtualEndpointGatewayResourceBindingDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_virtual_endpoint_gateway_resource_binding" {
			continue
		}

		getEndpointGatewayResourceBindingOptions := &vpcv1.GetEndpointGatewayResourceBindingOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getEndpointGatewayResourceBindingOptions.SetEndpointGatewayID(parts[0])
		getEndpointGatewayResourceBindingOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetEndpointGatewayResourceBinding(getEndpointGatewayResourceBindingOptions)

		if err == nil {
			return fmt.Errorf("EndpointGatewayResourceBinding still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for EndpointGatewayResourceBinding (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
