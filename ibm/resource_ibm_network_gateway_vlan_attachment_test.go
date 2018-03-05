package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMNetworkGatewayVlanAtachment_Basic(t *testing.T) {

	hostname1 := fmt.Sprintf("tfuat%s", acctest.RandString(11))
	hostname2 := fmt.Sprintf("tfuat%s", acctest.RandString(11))
	gatewayName := fmt.Sprintf("tfuatgw%s", acctest.RandString(12))

	config := "ibm_network_gateway.gw"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkGatewayVlanAttachment_basic(gatewayName, hostname1, hostname1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_gateway_vlan_association.attachment", "bypass", "true"),
					resource.TestCheckResourceAttr(
						config, "members.#", "2"),
					resource.TestCheckResourceAttr(
						config, "members.0.hostname", hostname1),
					resource.TestCheckResourceAttr(
						config, "members.0.domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						config, "members.0.datacenter", "ams01"),
					resource.TestCheckResourceAttr(
						config, "members.0.network_speed", "100"),
					resource.TestCheckResourceAttr(
						config, "members.0.private_network_only", "false"),
					resource.TestCheckResourceAttr(
						config, "members.0.ipv6_enabled", "true"),
					resource.TestCheckResourceAttr(
						config, "members.0.user_metadata", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						config, "members.0.notes", "member 1"),
					CheckStringSet(
						config,
						"members.0.tags", []string{"gateway tags 1", "terraform test tags 1"},
					),
					resource.TestCheckResourceAttr(
						config, "members.1.hostname", hostname2),
					resource.TestCheckResourceAttr(
						config, "members.1.domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						config, "members.1.datacenter", "ams01"),
					resource.TestCheckResourceAttr(
						config, "members.1.network_speed", "100"),
					resource.TestCheckResourceAttr(
						config, "members.1.private_network_only", "false"),
					resource.TestCheckResourceAttr(
						config, "members.1.user_metadata", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						config, "members.1.ipv6_enabled", "true"),
					resource.TestCheckResourceAttr(
						config, "members.1.notes", "member2"),
					CheckStringSet(
						config,
						"members.1.tags", []string{"gateway tags 2", "terraform test tags 2"},
					),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMNetworkGatewayVlanAttachment_update(gatewayName, hostname1, hostname2),
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
	hostname2 := fmt.Sprintf("tfuat%s", acctest.RandString(11))
	gatewayName := fmt.Sprintf("tfuatgw%s", acctest.RandString(12))

	config := "ibm_network_gateway.gw"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkGatewayVlanAttachment_import_update(gatewayName, hostname1, hostname1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_network_gateway_vlan_association.attachment", "bypass", "false"),
					resource.TestCheckResourceAttr(
						config, "members.#", "2"),
					resource.TestCheckResourceAttr(
						config, "members.0.hostname", hostname1),
					resource.TestCheckResourceAttr(
						config, "members.0.domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						config, "members.0.datacenter", "ams01"),
					resource.TestCheckResourceAttr(
						config, "members.0.network_speed", "100"),
					resource.TestCheckResourceAttr(
						config, "members.0.private_network_only", "false"),
					resource.TestCheckResourceAttr(
						config, "members.0.ipv6_enabled", "true"),
					resource.TestCheckResourceAttr(
						config, "members.0.user_metadata", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						config, "members.0.notes", "member 1"),
					CheckStringSet(
						config,
						"members.0.tags", []string{"gateway tags 1", "terraform test tags 1"},
					),
					resource.TestCheckResourceAttr(
						config, "members.1.hostname", hostname2),
					resource.TestCheckResourceAttr(
						config, "members.1.domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						config, "members.1.datacenter", "ams01"),
					resource.TestCheckResourceAttr(
						config, "members.1.network_speed", "100"),
					resource.TestCheckResourceAttr(
						config, "members.1.private_network_only", "false"),
					resource.TestCheckResourceAttr(
						config, "members.1.user_metadata", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						config, "members.1.ipv6_enabled", "true"),
					resource.TestCheckResourceAttr(
						config, "members.1.notes", "member2"),
					CheckStringSet(
						config,
						"members.1.tags", []string{"gateway tags 2", "terraform test tags 2"},
					),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkGatewayVlanAttachment_basic(gatewayName, hostName1, hostName2 string) string {
	return fmt.Sprintf(`
	resource "ibm_network_gateway" "gw" {
	       name   = "%s"
			members = [{
			    hostname             = "%s"
			    domain               = "terraformuat.ibm.com"
			    datacenter           = "ams01"
			    network_speed        = 100
			    private_network_only = false
			    tcp_monitoring       = true
			    process_key_name     = "INTEL_SINGLE_XEON_1270_3_40_2"
			    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			    redundant_network    = false
			    disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
			    public_bandwidth     = 20000
			    memory               = 8
			    ipv6_enabled         = true
			    public_vlan_id       = 2225905
				private_vlan_id      = 2225915
				user_metadata        = "{\"value\":\"newvalue\"}"
				notes                = "member 1"
				tags                 = ["gateway tags 1", "terraform test tags 1"]

			  },{
			    hostname             = "%s"
			    domain               = "terraformuat.ibm.com"
				datacenter           = "ams01"
			    public_vlan_id       = 2225905
			    private_vlan_id      = 2225915
			    network_speed        = 100
			    private_network_only = false
			    tcp_monitoring       = true
			    process_key_name     = "INTEL_SINGLE_XEON_1270_3_40_2"
			    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			    redundant_network    = false
			    disk_key_names       = ["HARD_DRIVE_1_00_TB_SATA_III"]
				public_bandwidth     = 20000
				user_metadata        = "{\"value\":\"newvalue\"}"
			    memory               = 16
				ipv6_enabled         = true
				notes                = "member 2"
				tags                 = ["gateway tags 2", "terraform test tags 2"]
			  }
			  ]
		  }
		  resource "ibm_network_gateway_vlan_association" "attachment"{
			  gateway_id = "${ibm_network_gateway.gw.id}"
			  network_vlan_id = 2225915
		  }
		  `, gatewayName, hostName1, hostName2)

}

func testAccCheckIBMNetworkGatewayVlanAttachment_update(gatewayName, hostName1, hostName2 string) string {
	return fmt.Sprintf(`
		resource "ibm_network_gateway" "gw" {
	       name   = "%s"
		   members = [{
			hostname             = "%s"
			domain               = "terraformuat.ibm.com"
			datacenter           = "ams01"
			network_speed        = 100
			private_network_only = false
			tcp_monitoring       = true
			process_key_name     = "INTEL_SINGLE_XEON_1270_3_40_2"
			os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			redundant_network    = false
			disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
			public_bandwidth     = 20000
			memory               = 8
			ipv6_enabled         = true
			public_vlan_id       = 2225905
			private_vlan_id      = 2225915
			notes                = "member 1"
			tags                 = ["gateway tags 1", "terraform test tags 1"]

		  },{
			hostname             = "%s"
			domain               = "terraformuat.ibm.com"
			datacenter           = "ams01"
			public_vlan_id       = 2225905
			private_vlan_id      = 2225915
			network_speed        = 100
			private_network_only = false
			tcp_monitoring       = true
			process_key_name     = "INTEL_SINGLE_XEON_1270_3_40_2"
			os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			redundant_network    = false
			disk_key_names       = ["HARD_DRIVE_1_00_TB_SATA_III"]
			public_bandwidth     = 20000
			memory               = 16
			ipv6_enabled         = true
			notes                = "member 2"
			tags                 = ["gateway tags 2", "terraform test tags 2"]
		  }
		  ]
		  }
		  resource "ibm_network_gateway_vlan_association" "attachment"{
			  gateway_id = "${ibm_network_gateway.gw.id}"
			  network_vlan_id = 2225915
			  bypass = false
		  }		  `, gatewayName, hostName1, hostName2)

}

func testAccCheckIBMNetworkGatewayVlanAttachment_import_update(gatewayName, hostName1, hostName2 string) string {
	return fmt.Sprintf(`
	resource "ibm_network_gateway" "gw" {
	       name   = "%s"
			members = [{
			    hostname             = "%s"
			    domain               = "terraformuat.ibm.com"
			    datacenter           = "ams01"
			    network_speed        = 100
			    private_network_only = false
			    tcp_monitoring       = true
			    process_key_name     = "INTEL_SINGLE_XEON_1270_3_40_2"
			    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			    redundant_network    = false
			    disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
			    public_bandwidth     = 20000
			    memory               = 8
			    ipv6_enabled         = true
			    public_vlan_id       = 2225905
				private_vlan_id      = 2225915
				user_metadata        = "{\"value\":\"newvalue\"}"
				notes                = "member 1"
				tags                 = ["gateway tags 1", "terraform test tags 1"]

			  },{
			    hostname             = "%s"
			    domain               = "terraformuat.ibm.com"
				datacenter           = "ams01"
			    public_vlan_id       = 2225905
			    private_vlan_id      = 2225915
			    network_speed        = 100
			    private_network_only = false
			    tcp_monitoring       = true
			    process_key_name     = "INTEL_SINGLE_XEON_1270_3_40_2"
			    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			    redundant_network    = false
			    disk_key_names       = ["HARD_DRIVE_1_00_TB_SATA_III"]
				public_bandwidth     = 20000
				user_metadata        = "{\"value\":\"newvalue\"}"
			    memory               = 16
				ipv6_enabled         = true
				notes                = "member 2"
				tags                 = ["gateway tags 2", "terraform test tags 2"]
			  }
			  ]
		  }
		  resource "ibm_network_gateway_vlan_association" "attachment"{
			  gateway_id = "${ibm_network_gateway.gw.id}"
			  network_vlan_id = 2225905
			  bypass = true
		  }
		  `, gatewayName, hostName1, hostName2)

}
