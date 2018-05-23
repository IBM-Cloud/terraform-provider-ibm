package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMDnsSecondary_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDNSSecondaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckIBMDnsSecondaryConfig, zoneName, transferFrequency1, masterIPAddress1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDnsSecondaryZoneExists("ibm_dns_secondary.dns-secondary-zone-1"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "zone_name", zoneName),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "transfer_frequency", "10"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "master_ip_address", masterIPAddress1),
				),
			},
			{
				Config: fmt.Sprintf(testAccCheckIBMDnsSecondaryConfig, zoneName, transferFrequency2, masterIPAddress2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "transfer_frequency", "15"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "master_ip_address", masterIPAddress2),
				),
			},
		},
	})
}

func TestAccIBMDnsSecondary_Basic_Tags(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDNSSecondaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckIBMDnsSecondaryConfig_basic_tags, zoneName, transferFrequency1, masterIPAddress1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDnsSecondaryZoneExists("ibm_dns_secondary.dns-secondary-zone-1"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "zone_name", zoneName),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "transfer_frequency", "10"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "master_ip_address", masterIPAddress1),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "tags.#", "2"),
				),
			},
			{
				Config: fmt.Sprintf(testAccCheckIBMDnsSecondaryConfig_updated_tags, zoneName, transferFrequency2, masterIPAddress2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "transfer_frequency", "15"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "master_ip_address", masterIPAddress2),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMDnsSecondary_Import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDNSSecondaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckIBMDnsSecondaryConfig, zoneName, transferFrequency1, masterIPAddress1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDnsSecondaryZoneExists("ibm_dns_secondary.dns-secondary-zone-1"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "zone_name", zoneName),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "transfer_frequency", "10"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "master_ip_address", masterIPAddress1),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_dns_secondary.dns-secondary-zone-1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMDnsSecondaryZoneExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		dnsId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetDnsSecondaryService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		foundSecondaryZone, err := service.Id(dnsId).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundSecondaryZone.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

func testAccCheckIBMDNSSecondaryDestroy(s *terraform.State) error {
	service := services.GetDnsSecondaryService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_secondary" {
			continue
		}

		dnsId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the domain
		_, err := service.Id(dnsId).GetObject()

		if err == nil {
			return fmt.Errorf("Dns secondary zone with id %d still exists", dnsId)
		}
	}

	return nil
}

const testAccCheckIBMDnsSecondaryConfig = `
resource "ibm_dns_secondary" "dns-secondary-zone-1" {
    zone_name = "%s"
    transfer_frequency = "%d"
    master_ip_address = "%s"
}
`

const testAccCheckIBMDnsSecondaryConfig_basic_tags = `
resource "ibm_dns_secondary" "dns-secondary-zone-1" {
    zone_name = "%s"
    transfer_frequency = "%d"
	master_ip_address = "%s"
	tags = ["one", "two"]
}
`

const testAccCheckIBMDnsSecondaryConfig_updated_tags = `
resource "ibm_dns_secondary" "dns-secondary-zone-1" {
    zone_name = "%s"
    transfer_frequency = "%d"
	master_ip_address = "%s"
	tags = ["one", "two", "three"]
}
`

var zoneName = fmt.Sprintf("tfuatdomain%s.com", acctest.RandString(10))
var masterIPAddress1 = "172.16.0.1"
var masterIPAddress2 = "172.16.0.2"
var transferFrequency1 = 10
var transferFrequency2 = 15
