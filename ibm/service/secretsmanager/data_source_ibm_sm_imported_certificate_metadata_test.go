// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

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
					resource.TestCheckResourceAttrSet("data.ibm_sm_imported_certificate_metadata.sm_imported_certificate_metadata", "retrieved_at"),
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

func TestAccIbmSmImportedCertificateMetadataDataSourceManagedCSR(t *testing.T) {
	dataSourceName := "data.ibm_sm_imported_certificate_metadata.managed_csr"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmImportedCertificateMetadataDataSourceManagedCSR(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "crn"),
					resource.TestCheckResourceAttrSet(dataSourceName, "secret_type"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.common_name", "example.com"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.alt_names", "alt1"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.client_flag", "true"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.code_signing_flag", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.organization.0", "IBM"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.country.0", "IL"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.ou.0", "ILSL"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.postal_code.0", "5320047"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.province.0", "DAN"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.locality.0", "Givatayim"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.email_protection_flag", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.exclude_cn_from_sans", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.ext_key_usage", "timestamping"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.ip_sans", "127.0.0.1"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.key_bits", "2048"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.key_type", "rsa"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.key_usage", "DigitalSignature"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.other_sans", "1.3.6.1.4.1.311.21.2.3;utf8:*.example.com"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.policy_identifiers", ""),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.uri_sans", "https://www.example.com/test"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.user_ids", "user"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.require_cn", "true"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.rotate_keys", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "managed_csr.0.server_flag", "true"),
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

func testAccCheckIbmSmImportedCertificateMetadataDataSourceManagedCSR() string {
	return fmt.Sprintf(`
		resource "ibm_sm_imported_certificate" "managed_csr" {
			instance_id   = "%s"
			region        = "%s"
			name = "imported_cert_terraform_test_managed_csr"
			managed_csr {
				alt_names = "alt1"
				client_flag = true
				code_signing_flag = false
				common_name = "example.com"
				country = ["IL"]
				email_protection_flag = false
				exclude_cn_from_sans = false
				ext_key_usage = "timestamping"
				ext_key_usage_oids = "1.3.6.1.5.5.7.3.67"
				ip_sans = "127.0.0.1"
				key_bits = 2048
				key_type = "rsa"
				key_usage = "DigitalSignature"
				locality = ["Givatayim"]
				policy_identifiers = ""
				organization = ["IBM"]
				other_sans = "1.3.6.1.4.1.311.21.2.3;utf8:*.example.com"
				ou = ["ILSL"]
				postal_code = ["5320047"]
				province = ["DAN"]
				require_cn = true
				rotate_keys = false
				server_flag = true
				street_address = ["Ariel Sharon 4"]
				uri_sans = "https://www.example.com/test"
				user_ids = "user"
			}
		}

		data "ibm_sm_imported_certificate_metadata" "managed_csr" {
			instance_id = "%s"
			region = "%s"
			secret_id = ibm_sm_imported_certificate.managed_csr.secret_id
		}

	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
