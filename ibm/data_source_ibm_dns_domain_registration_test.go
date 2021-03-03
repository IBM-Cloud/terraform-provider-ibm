// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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
