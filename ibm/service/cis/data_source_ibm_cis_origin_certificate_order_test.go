// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisOriginCertificateDataSource_Basic(t *testing.T) {
	name := "data.ibm_cis_certificate_order.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisOriginCertificateDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hosts.#", "1"),
				),
			},
		},
	})
}

func testAccCheckCisOriginCertificateDataSourceConfigBasic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprint(`
	data "ibm_cis_origin_certificates" "test" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
	  }
	`, acc.CisDomainStatic)
}
