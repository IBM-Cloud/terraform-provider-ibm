// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMDNSDomainRegistrationDataSource_Basic(t *testing.T) {

	var domainName = "wcpclouduk.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
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
