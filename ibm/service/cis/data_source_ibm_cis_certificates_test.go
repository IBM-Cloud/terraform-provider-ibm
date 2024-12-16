// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisCertificatesDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_certificates.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
	return testAccCheckCisCertificateOrderConfigBasic() + `
	data "ibm_cis_certificates" "test" {
		cis_id    = ibm_cis_certificate_order.test.cis_id
		domain_id = ibm_cis_certificate_order.test.domain_id
	}`
}
