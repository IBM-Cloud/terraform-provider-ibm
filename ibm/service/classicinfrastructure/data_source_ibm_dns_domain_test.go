// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDNSDomainDataSource_Basic(t *testing.T) {

	var domainName = acctest.RandString(16) + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckIBMDNSDomainDataSourceConfig_basic, domainName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_dns_domain.domain_id", "name", domainName),
					resource.TestMatchResourceAttr("data.ibm_dns_domain.domain_id", "id", regexp.MustCompile("^[0-9]+$")),
				),
			},
		},
	})
}

// The datasource to apply
const testAccCheckIBMDNSDomainDataSourceConfig_basic = `
resource "ibm_dns_domain" "ds_domain_test" {
	name = "%s"
}
data "ibm_dns_domain" "domain_id" {
    name = "${ibm_dns_domain.ds_domain_test.name}"
}
`
