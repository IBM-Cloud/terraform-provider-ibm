// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "private_path_service_gateway_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "account.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "expiration_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateway_endpoint_gateway_binding.is_private_path_service_gateway_endpoint_gateway_binding", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_private_path_service_gateway_endpoint_gateway_binding" "is_private_path_service_gateway_endpoint_gateway_binding_instance" {
			private_path_service_gateway = "private_path_service_gateway_id"
			id = "id"
		}
	`)
}
