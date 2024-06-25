// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVPNGatewayConnectionPeerCidrsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVPNGatewayConnectionPeerCidrsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection_peer_cidrs.is_vpn_gateway_connection_cidrs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection_peer_cidrs.is_vpn_gateway_connection_cidrs_instance", "vpn_gateway"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection_peer_cidrs.is_vpn_gateway_connection_cidrs_instance", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection_peer_cidrs.is_vpn_gateway_connection_cidrs_instance", "cidrs.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNGatewayConnectionPeerCidrsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_vpn_gateway_connection_peer_cidrs" "is_vpn_gateway_connection_cidrs_instance" {
			vpn_gateway = "vpn_gateway"
			id = "id"
		}
	`)
}
