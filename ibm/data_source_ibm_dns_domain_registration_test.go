package ibm

import (
	"fmt"
	"regexp"
	"testing"

	//"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMDNSDomainRegistrationDataSource_Basic(t *testing.T) {

	var domainName = "wcpclouduk.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckIBMDNSDomainRegistrationDataSourceConfig_basic, domainName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_dns_domain_registration.wcpclouduk", "name", domainName),
					resource.TestMatchResourceAttr("data.ibm_dns_domain_registration.wcpclouduk", "id", regexp.MustCompile("^[0-9]+$")),
				),
			},
		},
	})
}

// The datasource to apply
// const testAccCheckIBMDNSDomainREgistrationDataSourceConfig_basic = `
// resource "ibm_dns_domain_registration" "ds_domain_test" {
// 	name = "%s"
// }
// data "ibm_dns_domain_registration" "domain_id" {
//     name = "${ibm_dns_domain_registration.ds_domain_test.name}"
// }
// `

// resource "ibm_dns_domain_registration" "wcpclouduk" {
// 	name = "%s"
// }
const testAccCheckIBMDNSDomainRegistrationDataSourceConfig_basic = `
data "ibm_dns_domain_registration" "wcpclouduk" {
    name = "%s"
}
`
