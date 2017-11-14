package ibm

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/helpers/location"
	"github.com/softlayer/softlayer-go/helpers/product"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)
func TestAccIBMNetworkGateway_Basic(t *testing.T) {
	var networkGateway datatypes.Network_Gateway

	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
						"ibm_network_gateway.terraform-acceptance-test-1", "os_reference_code", "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "datacenter", "ams01"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "network_speed", "100"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "hourly_billing", "false"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "private_network_only", "false"),
					resource.TestCheckResourceAttr(
						"ibm_network_gateway.terraform-acceptance-test-1", "user_metadata", "{\"value\":\"newvalue\"}"),
					CheckStringSet(
						"ibm_network_gateway.terraform-acceptance-test-1",
						"tags", []string{"collectd"},
					),
				),
			},
}
func testAccCheckIBMNetworkGatewayConfig_basic(hostname string) string {
	return fmt.Sprintf(`
		resource "ibm_compute_bare_metal" "terraform-acceptance-test-1" {
			hostname               = "%s"
			domain                 = "terraformuat.ibm.com"
			os_reference_code      = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
			datacenter             = "ams01"
			network_speed          = 100
			hourly_billing         = true
			private_network_only   = false
			tags                   = ["collectd"]
			notes                  = "baremetal notes"
		}
		`, hostname)
}

func testAccCheckIBMNetworkGatewayExists(n string, bareMetal *datatypes.Hardware) resource.TestCheckFunc {
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
		bm, err := service.Id(id).GetObject()
		if err != nil {
			return err
		}

		fmt.Printf("The ID is %d", *bm.Id)

		if *bm.Id != id {
			return errors.New("Network Gateway not found")
		}

		*bareMetal = bm

		return nil
	}
}
