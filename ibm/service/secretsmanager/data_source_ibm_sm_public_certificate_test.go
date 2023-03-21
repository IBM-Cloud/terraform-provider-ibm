// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIbmSmPublicCertificateDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPublicCertificateDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "versions_total"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "common_name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "key_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate.sm_public_certificate", "rotation.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSmPublicCertificateDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "public_cert_ca_lets_encrypt-terraform-test-datasource"
			lets_encrypt_environment = "%s"
			lets_encrypt_private_key = "%s"
		}

		resource "ibm_sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis_instance" {
			  instance_id   = "%s"
       		  region        = "%s"
			  cloud_internet_services_crn = "%s"
  			  name = "cloud-internet-services-config-terraform-test"
		}

		resource "ibm_sm_public_certificate" "sm_public_certificate_instance" {
			instance_id = "%s"
			region = "%s"
  			name = "public-certificate-terraform-tests"
  			secret_group_id = "default"
  			common_name = "%s"
  			ca = ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt_instance.name
  			dns = ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis_instance.name
		}

		data "ibm_sm_public_certificate" "sm_public_certificate" {
			instance_id   = "%s"
			region        = "%s"
			secret_id = ibm_sm_public_certificate.sm_public_certificate_instance.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateLetsEncryptEnvironment, acc.SecretsManagerPublicCertificateLetsEncryptPrivateKey,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateCisCrn,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateCommonName,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
