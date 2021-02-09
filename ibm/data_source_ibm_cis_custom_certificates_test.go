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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisCustomCertificatesDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_custom_certificates.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisCustomCertificatesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "custom_certificates.0.id"),
					resource.TestCheckResourceAttrSet(node, "custom_certificates.0.bundle_method"),
				),
			},
		},
	})
}

func testAccCheckIBMCisCustomCertificatesDataSourceConfig() string {

	certMgrInstanceName := fmt.Sprintf("testacc-cert-manager-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	domainName := fmt.Sprintf("%s.%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum), cisDomainStatic)

	return testAccCheckCisCertificateUploadConfigBasic(certMgrInstanceName, domainName) +
		fmt.Sprintf(`
	data "ibm_cis_custom_certificates" "test" {
		cis_id    = ibm_cis_certificate_upload.test.cis_id
		domain_id = ibm_cis_certificate_upload.test.domain_id
	  }`)
}
