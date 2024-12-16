// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisMtlsApp_Basic(t *testing.T) {
	name := "ibm_cis_mtls_app." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisMtlsAppBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "domain", "darunya.austest-10.cistest-load.com"),
					resource.TestCheckResourceAttr(name, "name", "MTLS-APP"),
					resource.TestCheckResourceAttr(name, "policy_name", "MTLS-Policy"),
				),
			},
		},
	})
}

func testAccCheckCisMtlsAppBasic1(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_mtls_app" "%[1]s" {
		cis_id                         = data.ibm_cis.cis.id
		domain_id                      = data.ibm_cis_domain.cis_domain.domain_id
		domain                         = "darunya.austest-10.cistest-load.com"
		name                           = "MTLS-APP"
		policy_name                    = "Default Policy"
	  }
`, id)
}
