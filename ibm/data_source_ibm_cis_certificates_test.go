// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisCertificatesDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_certificates.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisCertificatesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "certificates.0.status"),
				),
			},
		},
	})
}

func testAccCheckIBMCisCertificatesDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckCisCertificateOrderConfigBasic() + fmt.Sprintf(`
	data "ibm_cis_certificates" "test" {
		cis_id    = ibm_cis_certificate_order.test.cis_id
		domain_id = ibm_cis_certificate_order.test.domain_id
	}`)
}
