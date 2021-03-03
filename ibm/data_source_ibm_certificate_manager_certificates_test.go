// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCertificateManagerCertificatesDataSource_Basic(t *testing.T) {
	cmsName := fmt.Sprintf("tf-acc-test1-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCertificateManagerCertificatesDataSourceConfig_basic(cmsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_certificate_manager_certificates.certs", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMCertificateManagerCertificatesDataSourceConfig_basic(cmsName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "cm" {
		name     = "%s"
		location = "us-south"
		service  = "cloudcerts"
		plan     = "free"
	}
	data "ibm_certificate_manager_certificates" "certs"{
		certificate_manager_instance_id=ibm_resource_instance.cm.id
	}
	  
	`, cmsName)
}
