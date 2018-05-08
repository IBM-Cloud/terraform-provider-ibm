package ibm

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func TestAccIBMNetworkGateway_standalone(t *testing.T) {
	var networkGateway datatypes.Network_Gateway

	hostname := fmt.Sprintf("tfuat%s", acctest.RandString(11))
	gatewayName := fmt.Sprintf("tfuat-gw-%s", acctest.RandString(7))
	config := "ibm_network_gateway.standalone"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMNetworkGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMNetworkGatewayStandaloneConfig(gatewayName, hostname),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMNetworkGatewayExists(config, &networkGateway),
					resource.TestCheckResourceAttr(
						config, "members.#", "1"),
					resource.TestCheckResourceAttr(
						config, "name", gatewayName),
					resource.TestCheckResourceAttr(
						config, "associated_vlans.#", "0"),
					resource.TestCheckResourceAttr(
						config, "status", "Active"),
				),
			},
		},
	})
}

func TestAccIBMNetworkGateway_ha_similar_members(t *testing.T) {
	var networkGateway datatypes.Network_Gateway
	hostname1 := fmt.Sprintf("tfuat%s", acctest.RandString(11))
	hostname2 := fmt.Sprintf("tfuat%s", acctest.RandString(11))
	gatewayName := fmt.Sprintf("tfuat-gw-%s", acctest.RandString(7))
	config := "ibm_network_gateway.ha_same_conf"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMNetworkGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMNetworkGatewaySameHardwareConfig(gatewayName, hostname1, hostname2),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMNetworkGatewayExists(config, &networkGateway),
					resource.TestCheckResourceAttr(
						config, "name", gatewayName),
					resource.TestCheckResourceAttr(
						config, "members.#", "2"),
					resource.TestCheckResourceAttr(
						config, "associated_vlans.#", "0"),
					resource.TestCheckResourceAttr(
						config, "status", "Active"),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkGatewayExists(n string, networkGateway *datatypes.Network_Gateway) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Network Gateway ID is set")
		}

		id, err := strconv.Atoi(rs.Primary.ID)

		if err != nil {
			return err
		}

		service := services.GetNetworkGatewayService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		ng, err := service.Id(id).GetObject()
		if err != nil {
			return err
		}

		fmt.Printf("The ID is %d", *ng.Id)

		if *ng.Id != id {
			return errors.New("Network Gateway not found")
		}

		*networkGateway = ng

		return nil
	}
}
func testAccCheckIBMNetworkGatewayDestroy(s *terraform.State) error {
	service := services.GetNetworkGatewayService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_network_gateway" {
			continue
		}

		id, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the Network Gateway
		_, err := service.Id(id).GetObject()

		// Wait
		if err != nil {
			if apiErr, ok := err.(sl.Error); !ok || apiErr.StatusCode != 404 {
				return fmt.Errorf(
					"Error waiting for Network Gateway (%d) to be destroyed: %s",
					id, err)
			}
		}
	}

	return nil
}

func testAccCheckIBMNetworkGatewayStandaloneConfig(gwName, hostname string) string {
	return fmt.Sprintf(`
resource "ibm_network_gateway" "standalone" {
	       name   = "%s"
	       members {
				hostname               = "%s"
				domain                 = "terraformuat.ibm.com"
				datacenter             = "ams01"
				network_speed          = 100
				private_network_only   = false
				tcp_monitoring         = true
				process_key_name       = "INTEL_SINGLE_XEON_1270_3_50"
				os_key_name            = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
				redundant_network      = false
				disk_key_names         = [ "HARD_DRIVE_2_00TB_SATA_II" ]
<<<<<<< HEAD
				public_bandwidth       = "BANDWIDTH_20000_GB"
				memory                 = 4
=======
				public_bandwidth       = 20000
				memory                 = 8
>>>>>>> master
				ipv6_enabled           = true
				tags                   = ["gateway_test", "terraform_test"]
		   }
		  }
`, gwName, hostname)
}

func testAccCheckIBMNetworkGatewaySameHardwareConfig(gatewayName, hostname1, hostname2 string) string {
	return fmt.Sprintf(`
		data "ibm_compute_ssh_key" "key" {
			label       = "test-lbaas"
			most_recent = true
		  }
		  resource "ibm_network_gateway" "ha_same_conf" {
			name = "%s"
			ssh_key_ids = ["${data.ibm_compute_ssh_key.key.id}"]
			members {
			  hostname             = "%s"
			  domain               = "terraformuat.ibm.com"
			  datacenter           = "ams01"
			  network_speed        = 100
			  private_network_only = false
			  ssh_key_ids          = ["${data.ibm_compute_ssh_key.key.id}"]
			  tcp_monitoring       = true
			  process_key_name     = "INTEL_SINGLE_XEON_1270_3_50"
			  os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			  redundant_network    = false
			  disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
			  public_bandwidth     = "BANDWIDTH_20000_GB"
			  memory               = 8
			  tags                 = ["gateway tags 1", "terraform test tags 1"]
			  notes                = "gateway notes 1"
			  ipv6_enabled         = true
			}
			members {
			  hostname             = "%s"
			  domain               = "terraformuat.ibm.com"
			  datacenter           = "ams01"
			  network_speed        = 100
			  private_network_only = false
			  ssh_key_ids          = ["${data.ibm_compute_ssh_key.key.id}"]
			  tcp_monitoring       = true
			  process_key_name     = "INTEL_SINGLE_XEON_1270_3_50"
			  os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			  redundant_network    = false
			  disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
			  public_bandwidth     = "BANDWIDTH_20000_GB"
			  memory               = 8
			  tags                 = ["gateway tags 2", "terraform test tags 2"]
			  notes                = "gateway notes 2"
			  ipv6_enabled         = true
			}
		  }
`, gatewayName, hostname1, hostname2)
}
