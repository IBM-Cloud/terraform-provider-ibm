package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMNetworkGatewayVlanAtachment_Basic(t *testing.T) {

	hostname1 := fmt.Sprintf("tfuat%s", acctest.RandString(11))
	gatewayName := fmt.Sprintf("tfuatgw%s", acctest.RandString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkGatewayVlanAttachment_basic(gatewayName, hostname1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_gateway_vlan_association.attachment", "bypass", "true"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMNetworkGatewayVlanAttachment_update(gatewayName, hostname1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_gateway_vlan_association.attachment", "bypass", "false"),
				),
			},
		},
	})
}

func TestAccIBMNetworkGatewayVlanAtachment_Import_Update(t *testing.T) {

	hostname1 := fmt.Sprintf("tfuat%s", acctest.RandString(11))
	gatewayName := fmt.Sprintf("tfuatgw%s", acctest.RandString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkGatewayVlanAttachment_import_update(gatewayName, hostname1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_gateway_vlan_association.attachment", "bypass", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkGatewayVlanAttachment_basic(gatewayName, hostName1 string) string {
	return fmt.Sprintf(`
			resource "ibm_network_vlan" "test_vlan" {
			  name            = "test_vlan"
			  datacenter      = "lon02"
			  type            = "PUBLIC"
			  router_hostname = "fcr01a.lon02"
			}
			resource "ibm_network_vlan" "test_vlan_assoc" {
			   name = "test_vlan_assoc"
			   datacenter = "lon02"
			   type = "PUBLIC"
			   router_hostname = "fcr01a.lon02"
			}
			
			resource "ibm_network_vlan" "test_vlan_p" {
			  name            = "tfuat_mult_subnet"
			  datacenter      = "lon02"
			  type            = "PRIVATE"
			  router_hostname = "bcr01a.lon02"
			}
			
			resource "ibm_network_gateway" "gw" {
			  name = "%s"
			
			  members = [{
			    hostname             = "%s"
			    domain               = "terraformuat.ibm.com"
			    datacenter           = "lon02"
			    network_speed        = 100
			    private_network_only = false
			    tcp_monitoring       = true
			    process_key_name     = "INTEL_SINGLE_XEON_1270_3_50"
			    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			    redundant_network    = false
			    disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
			    public_bandwidth     = 20000
			    memory               = 8
			    ipv6_enabled         = true
			    public_vlan_id       = "${ibm_network_vlan.test_vlan.id}"
			    private_vlan_id      = "${ibm_network_vlan.test_vlan_p.id}"
			    user_metadata        = "{\"value\":\"newvalue\"}"
			    notes                = "member 1"
			    tags                 = ["gateway tags 1", "terraform test tags 1"]
			  }]
			}
		  resource "ibm_network_gateway_vlan_association" "attachment"{
			  gateway_id = "${ibm_network_gateway.gw.id}"
			  network_vlan_id = "${ibm_network_vlan.test_vlan_assoc.id}"
		  }
		  `, gatewayName, hostName1)

}

func testAccCheckIBMNetworkGatewayVlanAttachment_update(gatewayName, hostName1 string) string {
	return fmt.Sprintf(`
					resource "ibm_network_vlan" "test_vlan" {
			  name            = "test_vlan"
			  datacenter      = "lon02"
			  type            = "PUBLIC"
			  router_hostname = "fcr01a.lon02"
			}
			
			resource "ibm_network_vlan" "test_vlan_assoc" {
			   name = "test_vlan_assoc"
			   datacenter = "lon02"
			   type = "PUBLIC"
			   router_hostname = "fcr01a.lon02"
			}
			
			resource "ibm_network_vlan" "test_vlan_p" {
			  name            = "tfuat_mult_subnet"
			  datacenter      = "lon02"
			  type            = "PRIVATE"
			  router_hostname = "bcr01a.lon02"
			}
			
			resource "ibm_network_gateway" "gw" {
			  name = "%s"
			
			  members = [{
			    hostname             = "%s"
			    domain               = "terraformuat.ibm.com"
			    datacenter           = "lon02"
			    network_speed        = 100
			    private_network_only = false
			    tcp_monitoring       = true
			    process_key_name     = "INTEL_SINGLE_XEON_1270_3_50"
			    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			    redundant_network    = false
			    disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
			    public_bandwidth     = 20000
			    memory               = 8
			    ipv6_enabled         = true
			    public_vlan_id       = "${ibm_network_vlan.test_vlan.id}"
			    private_vlan_id      = "${ibm_network_vlan.test_vlan_p.id}"
			    user_metadata        = "{\"value\":\"newvalue\"}"
			    notes                = "member 1"
			    tags                 = ["gateway tags 1", "terraform test tags 1"]
			  }]
			}
			  resource "ibm_network_gateway_vlan_association" "attachment"{
				  gateway_id = "${ibm_network_gateway.gw.id}"
				  network_vlan_id = "${ibm_network_vlan.test_vlan_assoc.id}"
				  bypass = false
			  }		  `, gatewayName, hostName1)

}

func testAccCheckIBMNetworkGatewayVlanAttachment_import_update(gatewayName, hostName1 string) string {
	return fmt.Sprintf(`
				resource "ibm_network_vlan" "test_vlan" {
			  name            = "test_vlan"
			  datacenter      = "lon02"
			  type            = "PUBLIC"
			  router_hostname = "fcr01a.lon02"
			}
			
			resource "ibm_network_vlan" "test_vlan_assoc" {
			   name = "test_vlan_assoc"
			   datacenter = "lon02"
			   type = "PUBLIC"
			   router_hostname = "fcr01a.lon02"
			}
			
			resource "ibm_network_vlan" "test_vlan_p" {
			  name            = "tfuat_mult_subnet"
			  datacenter      = "lon02"
			  type            = "PRIVATE"
			  router_hostname = "bcr01a.lon02"
			}
			
			resource "ibm_network_gateway" "gw" {
			  name = "%s"
			
			  members = [{
			    hostname             = "%s"
			    domain               = "terraformuat.ibm.com"
			    datacenter           = "lon02"
			    network_speed        = 100
			    private_network_only = false
			    tcp_monitoring       = true
			    process_key_name     = "INTEL_SINGLE_XEON_1270_3_50"
			    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			    redundant_network    = false
			    disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
			    public_bandwidth     = 20000
			    memory               = 8
			    ipv6_enabled         = true
			    public_vlan_id       = "${ibm_network_vlan.test_vlan.id}"
			    private_vlan_id      = "${ibm_network_vlan.test_vlan_p.id}"
			    user_metadata        = "{\"value\":\"newvalue\"}"
			    notes                = "member 1"
			    tags                 = ["gateway tags 1", "terraform test tags 1"]
			  }]
			}
		  resource "ibm_network_gateway_vlan_association" "attachment"{
			  gateway_id = "${ibm_network_gateway.gw.id}"
			  network_vlan_id = "${ibm_network_vlan.test_vlan_assoc.id}"
			  bypass = true
		  }
		  `, gatewayName, hostName1)

}
