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

func TestAccIBMNetworkGateway_Basic(t *testing.T) {
	var networkGateway datatypes.Hardware

	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMNetworkGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMNetworkGatewayConfig_basic(hostname),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMNetworkGatewayExists("ibm_network_gateway.terraform-acceptance-test-1", &networkGateway),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "hostname", hostname),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "datacenter", "ams01"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "network_speed", "100"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "private_network_only", "false"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "tcp_monitoring", "true"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "process_key_name", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "os_key_name", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "redundant_network", "false"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "public_bandwidth", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "memory", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "ipv6_enabled", "true"),
					CheckStringSet(
						"ibm_network_gateway.terraform-acceptance-test-1",
						"tags", []string{"collectd"},
					),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkGatewayConfig_basic(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_network_gateway" "terraform-acceptance-test-1" {
			hostname               = "%s"
			domain                 = "terraformuat.ibm.com"
			datacenter             = "ams01"
			network_speed          = 100
			private_network_only   = false
			tcp_monitoring         = true
			process_key_name       = "INTEL_SINGLE_XEON_1270_3_40_2"
			os_key_name            = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			redundant_network      = false
			public_bandwidth       = 20000
			memory                 = 4
			ipv6_enabled           = true
			server_instances       = [{"networkVlanID" = 645086,"bypass" = true},
			                          {"networkVlanID" = 637374,"bypass" = true}]
		  }
`, hostname)
}

func testAccCheckIBMNetworkGatewayExists(n string, networkGateway *datatypes.Hardware) resource.TestCheckFunc {
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

		service := services.GetHardwareService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
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
	service := services.GetHardwareService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

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
