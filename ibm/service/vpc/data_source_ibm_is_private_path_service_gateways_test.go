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
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "limit"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "private_path_service_gateways.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_private_path_service_gateways.is_private_path_service_gateways", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPrivatePathServiceGatewaysDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_private_path_service_gateways" "is_private_path_service_gateways_instance" {
		}
	`)
}

