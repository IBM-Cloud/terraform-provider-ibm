/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"regexp"
	"testing"
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

const testAccCheckIBMDNSDomainRegistrationDataSourceConfig_basic = `
data "ibm_dns_domain_registration" "wcpclouduk" {
    name = "%s"
}
`
