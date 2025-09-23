// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDataSourceBasic(t *testing.T) {
	accessPolicy := "deny"
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tf-test-lb%dd", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-test-ppsg%d", acctest.RandIntRange(10, 100))
	targetName := fmt.Sprintf("tf-egw-target%d", acctest.RandIntRange(10, 100))
	egwName := fmt.Sprintf("tf-egw%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name, egwName, targetName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "private_path_service_gateway"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.account.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.account.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.expiration_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "endpoint_gateway_bindings.0.updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDataSourceConfigBasic(vpcname, subnetname, zone, cidr, lbname, accessPolicy, name, egwName, targetName string) string {
	return testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name) + fmt.Sprintf(`
		resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
			name = "%s"
			target {
			name          = "%s"
			resource_type = "private_path_service_gateway"
			}
			vpc = ibm_is_vpc.testacc_vpc.id
			resource_group = data.ibm_resource_group.test_acc.id
		}
		data "ibm_is_private_path_service_gateway_endpoint_gateway_bindings" "is_private_path_service_gateway_endpoint_gateway_bindings" {
			private_path_service_gateway = ibm_is_private_path_service_gateway.is_private_path_service_gateway.id
		}
		data "ibm_is_private_path_service_gateway_endpoint_gateway_binding" "is_private_path_service_gateway_endpoint_gateway_binding" {
			private_path_service_gateway = ibm_is_private_path_service_gateway.is_private_path_service_gateway.id
			endpoint_gateway_binding = data.ibm_is_private_path_service_gateway_endpoint_gateway_bindings.is_private_path_service_gateway_endpoint_gateway_bindings.endpoint_gateway_bindings.0.id
		}
	`, egwName, targetName)
}
