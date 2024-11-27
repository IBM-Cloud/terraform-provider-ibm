// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVPNGatewayConnectionLocalCidrsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVPNGatewayConnectionLocalCidrsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection_local_cidrs.is_vpn_gateway_connection_cidrs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection_local_cidrs.is_vpn_gateway_connection_cidrs_instance", "vpn_gateway"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection_local_cidrs.is_vpn_gateway_connection_cidrs_instance", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection_local_cidrs.is_vpn_gateway_connection_cidrs_instance", "cidrs.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNGatewayConnectionLocalCidrsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_vpn_gateway_connection_local_cidrs" "is_vpn_gateway_connection_cidrs_instance" {
			vpn_gateway = "vpn_gateway"
			id = "id"
		}
	`)
}
