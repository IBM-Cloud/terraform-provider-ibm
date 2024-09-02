// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsPrivatePathServiceGatewaysDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewaysDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.default_access_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.endpoint_gateway_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.published"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.load_balancer.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.load_balancer.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.load_balancer.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.load_balancer.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.load_balancer.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.resource_group.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.service_endpoints"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.vpc.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.vpc.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.vpc.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.vpc.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.0.zonal_affinity"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewaysDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_private_path_service_gateways" "is_private_path_service_gateways" {
		}
	`)
}
