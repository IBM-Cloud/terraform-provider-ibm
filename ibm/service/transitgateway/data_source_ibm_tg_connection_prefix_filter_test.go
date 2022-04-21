// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIBMTransitGatewayConnectionPrefixFilterDataSource_basic(t *testing.T) {
	randNum := acctest.RandIntRange(10, 100)
	gatewayName := fmt.Sprintf("gateway-name-%d", randNum)
	location := fmt.Sprintf("us-south")
	connectionName := fmt.Sprintf("connection-name-%d", randNum)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTransitGatewayDataConnectionPrefixFilterSourceConfig(gatewayName, location, connectionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tg_connection_prefix_filter.ds_tg_prefix_filter", "filter_id"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewayDataConnectionPrefixFilterSourceConfig(gatewayName, location, connectionName string) string {
	return fmt.Sprintf(`

	resource "ibm_tg_gateway" "test_tg_gateway" {
		name="%s"
		location="%s"
		global=true
	}

	resource "ibm_tg_connection" "test_tg_connection"{
		gateway = ibm_tg_gateway.test_tg_gateway.id
		network_type = "classic"
		name = "%s"
	}

	resource "ibm_tg_connection_prefix_filter" "test_tg_prefix_filter" {
		gateway = ibm_tg_gateway.test_tg_gateway.id
		connection_id = ibm_tg_connection.test_tg_connection.connection_id
		action = "permit"
		prefix = "10.0.0.0/16"
	}
	
	data "ibm_tg_connection_prefix_filter" "ds_tg_prefix_filter" {
		gateway = ibm_tg_gateway.test_tg_gateway.id
		connection_id = ibm_tg_connection.test_tg_connection.connection_id
		filter_id = ibm_tg_connection_prefix_filter.test_tg_prefix_filter.filter_id
	}
	`, gatewayName, location, connectionName)
}
