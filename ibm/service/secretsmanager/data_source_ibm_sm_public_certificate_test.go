// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"
)

func TestAccIbmSmPublicCertificateDataSourceBasic(t *testing.T) {
	//resource.Test(t, resource.TestCase{
	//	PreCheck:  func() { acc.TestAccPreCheck(t) },
	//	Providers: acc.TestAccProviders,
	//	Steps: []resource.TestStep{
	//		resource.TestStep{
	//			Config: testAccCheckIbmSmPublicCertificateDataSourceConfigBasic(),
	//			Check: resource.ComposeTestCheckFunc(
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "id"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "id"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "created_by"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "created_at"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "crn"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "secret_group_id"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "secret_type"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "updated_at"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "versions_total"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "common_name"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "key_algorithm"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "rotation.#"),
	//			),
	//		},
	//	},
	//})
}

func testAccCheckIbmSmPublicCertificateDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_public_certificate" "sm_public_certificate_instance" {
			secret_prototype {
				custom_metadata = {"key":"value"}
				description = "Extended description for this secret."
				expiration_date = "2022-04-12T23:20:50.520Z"
				labels = [ "my-label" ]
				name = "my-secret-example"
				secret_group_id = "default"
				secret_type = "arbitrary"
				payload = "secret-credentials"
				version_custom_metadata = {"key":"value"}
			}
		}

		data "ibm_sm_public_certificate" "sm_public_certificate_instance" {
			id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
		}
	`)
}
