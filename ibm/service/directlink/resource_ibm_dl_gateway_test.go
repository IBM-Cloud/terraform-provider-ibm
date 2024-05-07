// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"errors"
	"fmt"
	"log"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/networking-go-sdk/directlinkv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMDLGateway_basic(t *testing.T) {
	var instance string
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	newgatewayname := fmt.Sprintf("newgateway-name-%d", acctest.RandIntRange(10, 100))
	custname := fmt.Sprintf("customer-name-%d", acctest.RandIntRange(10, 100))
	carriername := fmt.Sprintf("carrier-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayDestroy, // Delete test case
		Steps: []resource.TestStep{
			{
				//Create test case
				Config: testAccCheckIBMDLGatewayConfig(gatewayname, custname, carriername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayExists("ibm_dl_gateway.test_dl_gateway", instance),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_gateway", "name", gatewayname),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_gateway", "customer_name", custname),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_gateway", "carrier_name", carriername),
				),
			},
			{
				//Update test case
				Config: testAccCheckIBMDLGatewayConfig(newgatewayname, custname, carriername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayExists("ibm_dl_gateway.test_dl_gateway", instance),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_gateway", "name", newgatewayname),
				),
			},
		},
	})
}
func TestAccIBMDLGatewayConnect_basic(t *testing.T) {
	var instance string
	connectgatewayname := fmt.Sprintf("gateway-connect-%d", acctest.RandIntRange(10, 100))
	exprefix := "10.0.0.0/16"
	exupdatedPrefix := "10.0.0.0/17"
	imprefix := "10.0.0.0/16"
	imupdatedPrefix := "10.0.0.0/17"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayDestroy, // Delete test case
		Steps: []resource.TestStep{

			{
				//dl connect  test case
				Config: testAccCheckIBMDLConnectGatewayConfig(connectgatewayname, exprefix, imprefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayExists("ibm_dl_gateway.test_dl_connect", instance),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_connect", "name", connectgatewayname),
					//resource.TestCheckResourceAttr("data.ibm_dl_export_route_filter.test_dl_export_route_filter", "prefix", exprefix),
					//resource.TestCheckResourceAttr("data.ibm_dl_import_route_filter.test_dl_import_route_filter", "prefix", imprefix),
					//resource.TestCheckResourceAttrSet("ibm_dl_gateway.test_dl_connect", "as_prepends.#"),
				),
			},
			{
				//Update test case
				Config: testAccCheckIBMDLConnectGatewayConfig(connectgatewayname, exupdatedPrefix, imupdatedPrefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayExists("ibm_dl_gateway.test_dl_connect", instance),
					//resource.TestCheckResourceAttr("data.ibm_dl_export_route_filter.test_dl_export_route_filter", "prefix", exupdatedPrefix),
					//resource.TestCheckResourceAttr("data.ibm_dl_import_route_filter.test_dl_import_route_filter", "prefix", imupdatedPrefix),
				),
			},
		},
	})
}

func testAccCheckIBMDLGatewayConfig(gatewayname, custname, carriername string) string {
	return fmt.Sprintf(`
	data "ibm_dl_routers" "test1" {
		offering_type = "dedicated"
		location_name = "dal10"
	}
	  resource "ibm_dl_gateway" "test_dl_gateway" {
		bgp_asn =  64999
        global = true
        metered = false
        name = "%s"
        speed_mbps = 1000
        type =  "dedicated"
		cross_connect_router = data.ibm_dl_routers.test1.cross_connect_routers[0].router_name
        location_name = data.ibm_dl_routers.test1.location_name
		customer_name = "%s"
        carrier_name = "%s"
	  }
	  
	  `, gatewayname, custname, carriername)
}

func testAccCheckIBMDLConnectGatewayConfig(gatewayname string, exprefix string, imprefix string) string {
	return fmt.Sprintf(`
	data "ibm_dl_ports" "ds_dlports" {
	}
	  resource "ibm_dl_gateway" "test_dl_connect" {
		bgp_asn =  64999
        global = true
        metered = false
        name = "%s"
        speed_mbps = 1000
		type =  "connect"
		port =  data.ibm_dl_ports.ds_dlports.ports[0].port_id
		export_route_filters {
			action = "deny"
			prefix = "%s"
			ge =17
			le = 28
		}
		import_route_filters {
			action = "deny"
			prefix = "%s"
			ge =17
			le = 28
		}
	}
	/*
	data "ibm_dl_export_route_filters" "test_dl_export_route_filters" {
		gateway = ibm_dl_gateway.test_dl_connect.id
    }
	data "ibm_dl_export_route_filter" "test_dl_export_route_filter" {
		gateway = ibm_dl_gateway.test_dl_connect.id
		id = data.ibm_dl_export_route_filters.test_dl_export_route_filters.export_route_filters[0].export_route_filter_id
    }
	
	data "ibm_dl_import_route_filters" "test_dl_import_route_filters" {
		gateway = ibm_dl_gateway.test_dl_connect.id
    }
	data "ibm_dl_import_route_filter" "test_dl_import_route_filter" {
		gateway = ibm_dl_gateway.test_dl_connect.id
		id = data.ibm_dl_import_route_filters.test_dl_import_route_filters.import_route_filters[0].import_route_filter_id
    }
	*/
	  `, gatewayname, exprefix, imprefix)
}

func directlinkClient(meta interface{}) (*directlinkv1.DirectLinkV1, error) {
	ibmsess, err := meta.(conns.ClientSession).DirectlinkV1API()
	return ibmsess, err
}

func testAccCheckIBMDLGatewayExists(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		directLink, err := directlinkClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		getOptions := &directlinkv1.GetGatewayOptions{
			ID: &rs.Primary.ID,
		}
		instanceIntf, response, err := directLink.GetGateway(getOptions)

		if (err != nil) || (instanceIntf == nil) {
			return fmt.Errorf("[ERROR] Error Getting Direct Link Gateway (Dedicated Template): %s\n%s", err, response)
		}
		instance1 := instanceIntf.(*directlinkv1.GetGatewayResponse)
		instance = *instance1.ID
		return nil
	}
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
