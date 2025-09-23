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

func TestAccIBMIsPrivatePathServiceGatewayDataSourceBasic(t *testing.T) {
	accessPolicy := "deny"
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tf-test-lb%dd", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-test-ppsg%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "default_access_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "endpoint_gateways_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "published"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "load_balancer.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "load_balancer.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "load_balancer.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "load_balancer.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "load_balancer.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "resource_group.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "service_endpoints"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "vpc.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "vpc.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "vpc.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "vpc.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "zonal_affinity"),

					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "default_access_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "endpoint_gateways_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "published"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "load_balancer.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "load_balancer.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "load_balancer.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "load_balancer.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "load_balancer.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "resource_group.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "service_endpoints"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "vpc.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "vpc.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "vpc.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "vpc.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway_by_name", "zonal_affinity"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayDataSourceConfigBasic(vpcname, subnetname, zone, cidr, lbname, accessPolicy, name string) string {
	return testAccCheckIBMIsPrivatePathServiceGatewayConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, accessPolicy, name) + fmt.Sprintf(`
		data "ibm_is_private_path_service_gateway" "is_private_path_service_gateway" {
			private_path_service_gateway = ibm_is_private_path_service_gateway.is_private_path_service_gateway.id
		}
		data "ibm_is_private_path_service_gateway" "is_private_path_service_gateway_by_name" {
			private_path_service_gateway_name = ibm_is_private_path_service_gateway.is_private_path_service_gateway.name
		}
	`)
}
