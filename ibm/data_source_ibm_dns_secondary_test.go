package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMDNSSecondaryDataSource_Basic(t *testing.T) {

	var domainName = acctest.RandString(16) + ".com"
	var domainName1 = acctest.RandString(16) + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckIBMDNSSecondaryDataSourceConfig_basic, domainName, domainName1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_dns_secondary.secondary_domain_id", "zone_name", domainName),
					resource.TestMatchResourceAttr("data.ibm_dns_secondary.secondary_domain_id", "id", regexp.MustCompile("^[0-9]+$")),
					resource.TestCheckResourceAttr("data.ibm_dns_secondary.secondary_domain_id", "master_ip_address", "172.16.0.2"),
					resource.TestCheckResourceAttr("data.ibm_dns_secondary.secondary_domain_id", "transfer_frequency", "10"),
				),
			},
		},
	})
}

func TestAccIBMDNSSecondaryDataSource_InvalidZone(t *testing.T) {

	var domainName = acctest.RandString(16) + ".com"
	var domainName1 = acctest.RandString(16) + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      fmt.Sprintf(testAccCheckIBMDNSSecondaryDataSourceConfig_invlaid_zone, domainName, domainName1),
				ExpectError: regexp.MustCompile(fmt.Sprintf("No secondary zone found with name: %s", domainName1)),
			},
		},
	})
}

// The datasource to apply
const testAccCheckIBMDNSSecondaryDataSourceConfig_basic = `
resource "ibm_dns_secondary" "ds_secondary_domain_test" {
	zone_name = "%s"
	transfer_frequency = 10
	master_ip_address = "172.16.0.2"

}
resource "ibm_dns_secondary" "ds_secondary_domain_test1" {
	zone_name = "%s"
	transfer_frequency = 10
	master_ip_address = "172.16.0.2"

}
data "ibm_dns_secondary" "secondary_domain_id" {
    zone_name = "${ibm_dns_secondary.ds_secondary_domain_test.zone_name}"
}
`

const testAccCheckIBMDNSSecondaryDataSourceConfig_invlaid_zone = `

resource "ibm_dns_secondary" "ds_secondary_domain_test" {
	zone_name = "%s"
	transfer_frequency = 10
	master_ip_address = "172.16.0.2"

}

data "ibm_dns_secondary" "secondary_domain_id" {
    zone_name = "%s"
}
`
