// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"
)

func TestAccIbmSmPrivateCertificateMetadataDataSourceBasic(t *testing.T) {
	//resource.Test(t, resource.TestCase{
	//	PreCheck:  func() { acc.TestAccPreCheck(t) },
	//	Providers: acc.TestAccProviders,
	//	Steps: []resource.TestStep{
	//		resource.TestStep{
	//			Config: testAccCheckIbmSmPrivateCertificateMetadataDataSourceConfigBasic(),
	//			Check: resource.ComposeTestCheckFunc(
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "id"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "id"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "created_by"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "created_at"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "crn"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "secret_group_id"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "secret_type"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "updated_at"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "versions_total"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "signing_algorithm"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "certificate_template"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "common_name"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "expiration_date"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "issuer"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "serial_number"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_metadata.sm_private_certificate_metadata", "validity.#"),
	//			),
	//		},
	//	},
	//})
}

func testAccCheckIbmSmPrivateCertificateMetadataDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_sm_private_certificate_metadata" "sm_private_certificate_metadata_instance" {
			id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
		}
	`)
}
