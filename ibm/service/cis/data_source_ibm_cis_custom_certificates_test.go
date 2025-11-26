// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisCustomCertificatesDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_custom_certificates.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum), acc.CisDomainStatic)

	return testAccCheckCisCertificateUploadConfigBasic(certMgrInstanceName, domainName) +
		`
	data "ibm_cis_custom_certificates" "test" {
		cis_id    = ibm_cis_certificate_upload.test.cis_id
		domain_id = ibm_cis_certificate_upload.test.domain_id
	  }`
}
