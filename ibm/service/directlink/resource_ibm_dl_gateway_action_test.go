// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

/*
import (
	"errors"
	"fmt"
	"log"
	"testing"

	//dlProviderV2 "github.com/IBM/networking-go-sdk/directlinkproviderv2"
	//"github.com/IBM/networking-go-sdk/directlinkv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	//"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func TestAccIBMDLGatewayAction_basic(t *testing.T) {
	//var instance string
	gatewayname := fmt.Sprintf("provider-gateway-name12-%d", acctest.RandIntRange(10, 100))
	//custAccID := "3f455c4c574447adbc14bda52f80e62f" // bbsdldv1 account
	//node := "data.ibm_dl_gateway.test_ibm_dl_gateway1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDLProviderGatewayActionDestroy, // Delete test case
		Steps: []resource.TestStep{
			{
				//Create test case
				Config: testAccCheckIBMDLGatewayActionConfig(gatewayname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dl_gateway_action.test_dl_gateway_action", "operational_status", "create_reject"),
				),
			},
		},
	})
}

func testAccCheckIBMDLGatewayActionConfig(gatewayname string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		alias = "packet_fabric"
		ibmcloud_api_key      = "__BDfQHiKvN5Dsnsni4YerMqoPUgOIbso0f645JXn4R1"
		region                = "us-south"
		ibmcloud_timeout      = 300
	  }
	resource "ibm_dl_provider_gateway" "test_dl_gateway" {
		provider = ibm.packet_fabric
        bgp_asn =  64999
        name = "%s"
        customer_account_id = "3f455c4c574447adbc14bda52f80e62f"
        speed_mbps = 1000
        port = "3aa86cea-454d-4586-8247-222a36f7d1fe"
        vlan = 25
    }
	resource "ibm_dl_gateway_action" "test_dl_gateway_action" {
        action = "create_gateway_reject"
        global = true
        metered = true
		gateway = ibm_dl_provider_gateway.test_dl_gateway.id
		depends_on = [ ibm_dl_provider_gateway.test_dl_gateway ]
    }
	  `, gatewayname)
}

/*
	func directlinkClient(meta interface{}) (*directlinkv1.DirectLinkV1, error) {
		sess, err := meta.(conns.ClientSession).DirectlinkV1API()
		return sess, err
	}

	func testAccCheckIBMDLGatewayDestroy(s *terraform.State) error {
		directLink, err := directlinkClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_dl_gateway" {
				log.Printf("Destroy called ...%s", rs.Primary.ID)
				getOptions := &directlinkv1.GetGatewayOptions{
					ID: &rs.Primary.ID,
				}
				_, _, err = directLink.GetGateway(getOptions)

				if err == nil {
					return fmt.Errorf("gateway still exists: %s", rs.Primary.ID)
				}
			}
		}
		return nil
	}

func testAccCheckIBMDLProviderGatewayActionExists(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		directLink, err := directlinkProviderClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		getOptions := directLink.NewGetProviderGatewayOptions(rs.Primary.ID)

		instance1, response, err := directLink.GetProviderGateway(getOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting Direct Link Provider Gateway: %s\n%s", err, response)
		}
		instance = *instance1.ID
		return nil
	}
}

func testAccCheckIBMDLProviderGatewayActionDestroy(s *terraform.State) error {

	directLink, err := directlinkProviderClient(acc.TestAccProvider.Meta())
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
				//return nil
			}
		}
	}
	return nil
}
*/
