// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMTransitGatewayConnectionPrefixFilter_basic(t *testing.T) {
	var tgPrefixFilter string
	randNum := acctest.RandIntRange(10, 100)
	gatewayName := fmt.Sprintf("gateway-name-%d", randNum)
	location := fmt.Sprintf("us-south")
	connectionName := fmt.Sprintf("connection-name-%d", randNum)
	prefix := "10.0.0.0/16"
	updatedPrefix := "10.0.0.0/17"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayConnectionPrefixFilterDestroy,
		Steps: []resource.TestStep{
			// Create test case
			{
				Config: testAccCheckIBMTransitGatewayConnectionPrefixFiltersConfig(gatewayName, location, connectionName, prefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionPrefixFilterExists("ibm_tg_connection_prefix_filter.test_tg_prefix_filter", tgPrefixFilter),
					resource.TestCheckResourceAttr("ibm_tg_connection_prefix_filter.test_tg_prefix_filter", "prefix", prefix),
				),
			},
			// Update test case
			{
				Config: testAccCheckIBMTransitGatewayConnectionPrefixFiltersConfig(gatewayName, location, connectionName, updatedPrefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionPrefixFilterExists("ibm_tg_connection_prefix_filter.test_tg_prefix_filter", tgPrefixFilter),
					resource.TestCheckResourceAttr("ibm_tg_connection_prefix_filter.test_tg_prefix_filter", "prefix", updatedPrefix),
				),
			},
		},
	},
	)
}

func testAccCheckIBMTransitGatewayConnectionPrefixFiltersConfig(gatewayName, location, connectionName, prefix string) string {
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
		prefix = "%s"
	}
	`, gatewayName, location, connectionName, prefix)
}

func testAccCheckIBMTransitGatewayConnectionPrefixFilterExists(n string, vc string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := transitgatewayClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayId := parts[0]
		connectionId := parts[1]
		filterId := parts[2]

		detailPrefixFilterOptions := &transitgatewayapisv1.GetTransitGatewayConnectionPrefixFilterOptions{}
		detailPrefixFilterOptions.SetTransitGatewayID(gatewayId)
		detailPrefixFilterOptions.SetID(connectionId)
		detailPrefixFilterOptions.SetFilterID(filterId)

		r, response, err := client.GetTransitGatewayConnectionPrefixFilter(detailPrefixFilterOptions)
		if err != nil {
			return fmt.Errorf("testAccCheckIBMTransitGatewayConnectionPrefixFilterExists: Error Getting Transit Gateway Connection Prefix Filter: %s\n%s", err, response)
		}

		vc = *r.ID
		return nil
	}
}

func testAccCheckIBMTransitGatewayConnectionPrefixFilterDestroy(s *terraform.State) error {
	client, err := transitgatewayClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tg_connection_prefix_filter" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayId := parts[0]
		connectionId := parts[1]
		filterId := parts[2]

		detailPrefixFilterOptions := &transitgatewayapisv1.GetTransitGatewayConnectionPrefixFilterOptions{}
		detailPrefixFilterOptions.SetTransitGatewayID(gatewayId)
		detailPrefixFilterOptions.SetID(connectionId)
		detailPrefixFilterOptions.SetFilterID(filterId)

		_, _, err = client.GetTransitGatewayConnectionPrefixFilter(detailPrefixFilterOptions)
		if err == nil {
			return fmt.Errorf(" transit gateway connection prefix filter still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}
