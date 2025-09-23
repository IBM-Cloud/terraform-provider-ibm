// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmPrivateCertificateConfigurationIntermediateCADataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCADataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", "config_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", "secret_type"),
					//resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", "created_by"),
					//resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", "created_at"),
					//resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", "signing_method"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", "common_name"),
				),
			},
		},
	})
}

func TestAccIbmSmPrivateCertificateConfigurationIntermediateCADataSourceCryptoKey(t *testing.T) {
	if acc.SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderInstanceCrn != "" {
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { acc.TestAccPreCheck(t) },
			Providers: acc.TestAccProviders,
			Steps: []resource.TestStep{
				resource.TestStep{
					Config: testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCADataSourceConfigCryptoKey(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "id"),
						resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "name"),
						resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "config_type"),
						resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "secret_type"),
						//resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "created_by"),
						//resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "created_at"),
						//resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "updated_at"),
						resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "signing_method"),
						resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "common_name"),
						resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_crypto_key", "crypto_key.#"),
					),
				},
			},
		})
	}
}

func testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCADataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_root_ca" "ibm_sm_private_certificate_configuration_root_ca_instance" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			crl_expiry = "10000h"
			name = "root-ca-terraform-private-cert-datasource-test"
		}
		resource "ibm_sm_private_certificate_configuration_intermediate_ca" "ibm_sm_private_certificate_configuration_intermediate_ca_instance" {
  			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			issuer = ibm_sm_private_certificate_configuration_root_ca.ibm_sm_private_certificate_configuration_root_ca_instance.name
			signing_method = "internal"
			name = "intermediate-ca-terraform-private-cert-datasource-test"
		}

		data "ibm_sm_private_certificate_configuration_intermediate_ca" "sm_private_certificate_configuration_intermediate_ca" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_private_certificate_configuration_intermediate_ca.ibm_sm_private_certificate_configuration_intermediate_ca_instance.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCADataSourceConfigCryptoKey() string {
	return privateCertificateIntermediateCAConfigCryptoKey() + fmt.Sprintf(`
		data "ibm_sm_private_certificate_configuration_intermediate_ca" "sm_private_certificate_configuration_intermediate_ca_crypto_key" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_cert_intermediate_ca_crypto_key.name
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
