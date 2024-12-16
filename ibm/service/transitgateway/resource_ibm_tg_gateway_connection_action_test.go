// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMTransitGatewayConnectionAction_basic(t *testing.T) {
	var tgConnection string
	var randNum = acctest.RandIntRange(10, 100)
	connectionName := fmt.Sprintf("tg-connection-name-%d", randNum)
	gatewayName := fmt.Sprintf("tg-gateway-name-%d", randNum)
	var xacConnectionAction = "approve"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// tg cross account test
			{
				//Create test case
				Config: testAccCheckIBMTransitGatewayCrossAccConnectionWithActionConfig(gatewayName, connectionName, xacConnectionAction),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_tg_xac_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_tg_xac_connection", "request_status", "Attached"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewayCrossAccConnectionWithActionConfig(gatewayName, connectionName, connectionAction string) string {
	return fmt.Sprintf(`	
resource "ibm_tg_gateway" "test_tg_gateway"{
	provider = ibm
	name="%s"
	location="us-south"
	global=true
}

resource "ibm_tg_connection" "test_tg_xac_connection" {
	provider = ibm
	gateway = ibm_tg_gateway.test_tg_gateway.id
	network_type = "classic"
	name = "%s"
	network_account_id = "%s"
}

resource "ibm_tg_connection_action" "test_tg_xac_approve" {
	provider = ibm.account2
    gateway = ibm_tg_gateway.test_tg_gateway.id
    connection_id = ibm_tg_connection.test_tg_xac_connection.connection_id
    action = "%s"
}
	  `, gatewayName, connectionName, acc.Tg_cross_network_account_id, connectionAction)
}
