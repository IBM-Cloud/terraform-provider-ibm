// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmPrivateCertificateConfigurationRootCADataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPrivateCertificateConfigurationRootCADataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca", "config_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca", "common_name"),
				),
			},
		},
	})
}

func TestAccIbmSmPrivateCertificateConfigurationRootCADataSourceCryptoKey(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPrivateCertificateConfigurationRootCADataSourceConfigCryptoKey(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_crypto_key", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_crypto_key", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_crypto_key", "config_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_crypto_key", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_crypto_key", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_crypto_key", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_crypto_key", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_crypto_key", "common_name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_crypto_key", "crypto_key.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSmPrivateCertificateConfigurationRootCADataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_root_ca" "ibm_sm_private_certificate_configuration_root_ca_instance" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			crl_expiry = "10000h"
			name = "root-ca-terraform-private-cert-datasource-test"
		}

		data "ibm_sm_private_certificate_configuration_root_ca" "sm_private_certificate_configuration_root_ca" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_private_certificate_configuration_root_ca.ibm_sm_private_certificate_configuration_root_ca_instance.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmPrivateCertificateConfigurationRootCADataSourceConfigCryptoKey() string {
	return privateCertificateRootCAConfigCryptoKey() + fmt.Sprintf(`
		data "ibm_sm_private_certificate_configuration_root_ca" "sm_private_certificate_configuration_root_ca_crypto_key" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_private_certificate_configuration_root_ca.sm_private_cert_root_ca_crypto_key.name
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)

}
