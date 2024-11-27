// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmImportedCertificateDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmImportedCertificateDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "versions_total"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "signing_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "expiration_date"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "intermediate_included"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "issuer"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "private_key_included"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "serial_number"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "validity.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate", "certificate"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate_by_name", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate.sm_imported_certificate_by_name", "secret_group_name"),
				),
			},
		},
	})
}

func testAccCheckIbmSmImportedCertificateDataSourceConfigBasic() string {
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

		data "ibm_sm_imported_certificate" "sm_imported_certificate" {
			instance_id = "%s"
			region = "%s"
			secret_id = ibm_sm_imported_certificate.sm_imported_certificate_instance.secret_id
		}

		data "ibm_sm_imported_certificate" "sm_imported_certificate_by_name" {
			instance_id   = "%s"
			region = "%s"
			name = ibm_sm_imported_certificate.sm_imported_certificate_instance.name
			secret_group_name = "default"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerImportedCertificatePathToCertificate, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
