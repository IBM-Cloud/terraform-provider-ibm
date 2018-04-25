package softlayer

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMDnsSecondary_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDnsSecondaryConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDnsSecondaryZoneExists("ibm_dns_secondary.dns-secondary-zone-1"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "zoneName", "new-secondary-zone1.com"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "transferFrequency", "10"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "masterIpAddress", "172.16.0.1"),
				),
			},
			{
				Config: testAccCheckIBMDnsSecondaryConfig_updated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "transferFrequency", "15"),
					resource.TestCheckResourceAttr(
						"ibm_dns_secondary.dns-secondary-zone-1", "masterIpAddress", "172.16.0.2"),
				),
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

const testAccCheckIBMDnsSecondaryConfig_basic = `
resource "ibm_dns_secondary" "dns-secondary-zone-1" {
    zoneName = "new-secondary-zone1.com"
    transferFrequency = 10
    masterIpAddress = "172.16.0.1"
}
`

const testAccCheckIBMDnsSecondaryConfig_updated = `
resource "ibm_dns_secondary" "dns-secondary-zone-1" {
    zoneName = "new-secondary-zone1.com"
    transferFrequency = 15
    masterIpAddress = "172.16.0.2"
}
`
