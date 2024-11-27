// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmImportedCertificateMetadataDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmImportedCertificateMetadataDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "versions_total"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "signing_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "expiration_date"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "intermediate_included"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "issuer"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "private_key_included"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "serial_number"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "validity.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSmImportedCertificateMetadataDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_imported_certificate" "sm_imported_certificate_instance" {
			instance_id   = "%s"
			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			secret_group_id = "default"
			name = "imported_cert_terraform_test"	
			certificate = file("%s")
		}
		data "ibm_sm_imported_certificate_metadata" "sm_imported_certificate_metadata" {
			instance_id = "%s"
			region = "%s"
			secret_id = ibm_sm_imported_certificate.sm_imported_certificate_instance.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerImportedCertificatePathToCertificate, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
