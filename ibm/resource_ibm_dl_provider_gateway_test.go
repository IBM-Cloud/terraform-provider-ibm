// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccIBMDLProviderGateway_basic(t *testing.T) {
	var instance string
	gatewayname := fmt.Sprintf("tf-gateway-name-%d", acctest.RandIntRange(10, 100))
	//	newgatewayname := fmt.Sprintf("newgateway-name-%d", acctest.RandIntRange(10, 100))
	custAccID := "3f455c4c574447adbc14bda52f80e62f"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDLProviderGatewayDestroy, // Delete test case
		Steps: []resource.TestStep{
			{
				//Create test case
				Config: testAccCheckIBMDLProviderGatewayConfig(gatewayname, custAccID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLProviderGatewayExists("ibm_dl_provider_gateway.test_dl_gateway", instance),
					resource.TestCheckResourceAttr("ibm_dl_provider_gateway.test_dl_gateway", "name", gatewayname),
				),
			},
		},
	})
}

func testAccCheckIBMDLProviderGatewayConfig(gatewayname, custAccID string) string {
	return fmt.Sprintf(`
	data "ibm_dl_provider_ports" "test_ds_dl_ports" {
	}
	  resource "ibm_dl_provider_gateway" "test_dl_gateway" {
		bgp_asn =  64999
		bgp_ibm_cidr =  "169.254.10.29/30"
		bgp_cer_cidr = "169.254.10.30/30"
		name = "%s"
		customer_account_id = "%s"
		speed_mbps = 1000
		port =  data.ibm_dl_provider_ports.test_ds_dl_ports.ports[0].port_id
	  }
	  
	  `, gatewayname, custAccID)
}

func testAccCheckIBMDLProviderGatewayExists(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		directLink, err := directlinkProviderClient(testAccProvider.Meta())
		if err != nil {
			return err
		}

		getOptions := directLink.NewGetProviderGatewayOptions(rs.Primary.ID)

		instance1, response, err := directLink.GetProviderGateway(getOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Direct Link Provider Gateway: %s\n%s", err, response)
		}
		instance = *instance1.ID
		return nil
	}
}

func testAccCheckIBMDLProviderGatewayDestroy(s *terraform.State) error {
	directLink, err := directlinkProviderClient(testAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dl_provider_gateway" {
			log.Printf("Destroy called ...%s", rs.Primary.ID)
			getOptions := directLink.NewGetProviderGatewayOptions(rs.Primary.ID)

			_, _, err = directLink.GetProviderGateway(getOptions)

			if err == nil {
				return fmt.Errorf("gateway still exists: %s", rs.Primary.ID)
			}
		}
	}
	return nil
}
