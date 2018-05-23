package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMLbVpx_Basic(t *testing.T) {
	var nadc datatypes.Network_Application_Delivery_Controller

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbVpxConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMLbVpxExists("ibm_lb_vpx.testacc_foobar_vpx", &nadc),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "type", "NetScaler VPX"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "speed", "10"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "plan", "Standard"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "ip_count", "2"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "version", "10.1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "vip_pool.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMLbVpxWithIPCount1(t *testing.T) {
	var nadc datatypes.Network_Application_Delivery_Controller

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbVpxWithIPCount1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMLbVpxExists("ibm_lb_vpx.testacc_foobar_vpx", &nadc),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "type", "NetScaler VPX"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "speed", "10"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "plan", "Standard"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "ip_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "version", "11.0"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "vip_pool.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMLbVpxWithTag(t *testing.T) {
	var nadc datatypes.Network_Application_Delivery_Controller

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbVpxWithTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMLbVpxExists("ibm_lb_vpx.testacc_foobar_vpx", &nadc),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "type", "NetScaler VPX"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "speed", "10"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "plan", "Standard"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "ip_count", "2"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "version", "10.1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "vip_pool.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMLbVpxWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMLbVpxExists("ibm_lb_vpx.testacc_foobar_vpx", &nadc),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "type", "NetScaler VPX"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "datacenter", "dal09"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "speed", "10"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "plan", "Standard"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "ip_count", "2"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "version", "10.1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "vip_pool.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx.testacc_foobar_vpx", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMLbVpxExists(n string, nadc *datatypes.Network_Application_Delivery_Controller) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		nadcId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetNetworkApplicationDeliveryControllerService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		found, err := service.Id(nadcId).GetObject()
		if err != nil {
			return err
		}

		if strconv.Itoa(int(*found.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*nadc = found

		return nil
	}
}

const testAccCheckIBMLbVpxConfig_basic = `
resource "ibm_lb_vpx" "testacc_foobar_vpx" {
    datacenter = "dal09"
    speed = 10
    version = "10.1"
    plan = "Standard"
    ip_count = 2
}`

const testAccCheckIBMLbVpxWithIPCount1 = `
resource "ibm_lb_vpx" "testacc_foobar_vpx" {
    datacenter = "dal09"
    speed = 10
    version = "11.0"
    plan = "Standard"
    ip_count = 1
}`

const testAccCheckIBMLbVpxWithTag = `
resource "ibm_lb_vpx" "testacc_foobar_vpx" {
    datacenter = "dal09"
    speed = 10
    version = "10.1"
    plan = "Standard"
	ip_count = 2
	tags = ["one", "two"]
}`

const testAccCheckIBMLbVpxWithUpdatedTag = `
resource "ibm_lb_vpx" "testacc_foobar_vpx" {
    datacenter = "dal09"
    speed = 10
    version = "10.1"
    plan = "Standard"
	ip_count = 2
	tags = ["one", "two", "three"]
}`
