// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisAdvancedCertificatepackOrder_Basic(t *testing.T) {
	name := "ibm_cis_advanced_certificate_pack_order.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisAdvancedCertificatePackOrderConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hosts.#", "1"),
				),
			},
		},
	})
}

func testAccCheckCisAdvancedCertificatePackOrderConfigBasic() string {
	return fmt.Sprintf(`
	resource "ibm_cis_advanced_certificate_pack_order" "test" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
		hosts     = ["%[1]s"]
		certificate_authority = "lets_encrypt"
		validation_method = "txt"
    	validity = 90
	  }
	`, acc.CisDomainStatic)
}
