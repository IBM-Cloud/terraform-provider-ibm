// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmPublicCertificateConfigurationCALetsEncryptDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPublicCertificateConfigurationCALetsEncryptDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt", "lets_encrypt_environment"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt", "lets_encrypt_private_key"),
				),
			},
		},
	})
}

func testAccCheckIbmSmPublicCertificateConfigurationCALetsEncryptDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "public_cert_ca_lets_encrypt-terraform-test-datasource"
			lets_encrypt_environment = "%s"
			lets_encrypt_private_key = "%s"
		}

		data "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt_instance.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateLetsEncryptEnvironment, acc.SecretsManagerPublicCertificateLetsEncryptPrivateKey, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
