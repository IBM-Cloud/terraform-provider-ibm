/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMDNSDomainDataSource_Basic(t *testing.T) {

	var domainName = acctest.RandString(16) + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
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
