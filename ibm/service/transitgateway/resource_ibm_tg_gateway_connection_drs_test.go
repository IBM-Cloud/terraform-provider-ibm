// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMTransitGatewayConnectionVpnGateway_basic(t *testing.T) {
	if acc.Tg_cross_network_id == "" {
		t.Skip("Skipping TestAccIBMTransitGatewayConnectionVpnGateway_basic because IBM_TG_CROSS_NETWORK_ID is not set")
	}
	var tgConnection string
	var randNum = acctest.RandIntRange(10, 100)
	connectionName := fmt.Sprintf("tg-connection-vpngw-%d", randNum)
	gatewayName := fmt.Sprintf("tg-gateway-name-%d", randNum)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTransitGatewayVpnGatewayConnectionConfig(gatewayName, connectionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_tg_vpngw_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_tg_vpngw_connection", "name", connectionName),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_tg_vpngw_connection", "network_type", "vpn_gateway"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewayVpnGatewayConnectionConfig(gatewayName, connectionName string) string {
	return fmt.Sprintf(`
resource "ibm_tg_gateway" "test_tg_gateway" {
	name     = "%s"
	location = "us-south"
	global   = true
}

resource "ibm_tg_connection" "test_tg_vpngw_connection" {
	gateway      = ibm_tg_gateway.test_tg_gateway.id
	network_type = "vpn_gateway"
	name         = "%s"
	network_id   = "%s"
}
`, gatewayName, connectionName, acc.Tg_cross_network_id)
}
