// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsPrivatePathServiceGatewayEndpointGatewayBindingsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayEndpointGatewayBindingsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_bindings.is_private_path_service_gateway_endpoint_gateway_bindings", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_bindings.is_private_path_service_gateway_endpoint_gateway_bindings", "private_path_service_gateway_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_bindings.is_private_path_service_gateway_endpoint_gateway_bindings", "endpoint_gateway_bindings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_bindings.is_private_path_service_gateway_endpoint_gateway_bindings", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_bindings.is_private_path_service_gateway_endpoint_gateway_bindings", "limit"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_bindings.is_private_path_service_gateway_endpoint_gateway_bindings", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayEndpointGatewayBindingsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_private_path_service_gateway_endpoint_gateway_bindings" "is_private_path_service_gateway_endpoint_gateway_bindings_instance" {
			private_path_service_gateway_id = "private_path_service_gateway_id"
		}
	`)
}

