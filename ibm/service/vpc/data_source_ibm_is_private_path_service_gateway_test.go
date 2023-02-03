// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsPrivatePathServiceGatewayDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "default_access_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "endpoint_gateways_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "load_balancer.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "published"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "region.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "service_endpoints.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway.is_private_path_service_gateway", "zonal_affinity"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_private_path_service_gateway" "is_private_path_service_gateway_instance" {
			id = "id"
		}
	`)
}
