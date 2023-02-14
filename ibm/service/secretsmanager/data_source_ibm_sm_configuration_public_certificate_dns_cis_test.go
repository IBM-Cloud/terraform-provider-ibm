// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"
)

func TestAccIbmSmConfigurationPublicCertificateDnsCisDataSourceBasic(t *testing.T) {
	//resource.Test(t, resource.TestCase{
	//	PreCheck:  func() { acc.TestAccPreCheck(t) },
	//	Providers: acc.TestAccProviders,
	//	Steps: []resource.TestStep{
	//		resource.TestStep{
	//			Config: testAccCheckIbmSmConfigurationPublicCertificateDnsCisDataSourceConfigBasic(),
	//			Check: resource.ComposeTestCheckFunc(
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_cis.sm_configuration_public_certificate_dns_cis", "id"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_cis.sm_configuration_public_certificate_dns_cis", "name"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_cis.sm_configuration_public_certificate_dns_cis", "config_type"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_cis.sm_configuration_public_certificate_dns_cis", "secret_type"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_cis.sm_configuration_public_certificate_dns_cis", "created_by"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_cis.sm_configuration_public_certificate_dns_cis", "created_at"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_cis.sm_configuration_public_certificate_dns_cis", "updated_at"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_cis.sm_configuration_public_certificate_dns_cis", "cloud_internet_services_crn"),
	//			),
	//		},
	//	},
	//})
}

func testAccCheckIbmSmConfigurationPublicCertificateDnsCisDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_configuration_public_certificate_dns_cis" "sm_configuration_public_certificate_dns_cis_instance" {
			configuration_prototype {
				config_type = "public_cert_configuration_ca_lets_encrypt"
				name = "my-example-engine-config"
				lets_encrypt_environment = "production"
				lets_encrypt_private_key = "lets_encrypt_private_key"
				lets_encrypt_preferred_chain = "lets_encrypt_preferred_chain"
			}
		}

		data "ibm_sm_configuration_public_certificate_dns_cis" "sm_configuration_public_certificate_dns_cis_instance" {
			name = "configuration-name"
		}
	`)
}
