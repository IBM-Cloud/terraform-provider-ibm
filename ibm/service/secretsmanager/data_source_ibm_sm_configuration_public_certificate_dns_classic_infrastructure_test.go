// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"
)

func TestAccIbmSmConfigurationPublicCertificateDnsClassicInfrastructureDataSourceBasic(t *testing.T) {
	//resource.Test(t, resource.TestCase{
	//	PreCheck:  func() { acc.TestAccPreCheck(t) },
	//	Providers: acc.TestAccProviders,
	//	Steps: []resource.TestStep{
	//		resource.TestStep{
	//			Config: testAccCheckIbmSmConfigurationPublicCertificateDnsClassicInfrastructureDataSourceConfigBasic(),
	//			Check: resource.ComposeTestCheckFunc(
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_classic_infrastructure.sm_configuration_public_certificate_dns_classic_infrastructure", "id"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_classic_infrastructure.sm_configuration_public_certificate_dns_classic_infrastructure", "name"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_classic_infrastructure.sm_configuration_public_certificate_dns_classic_infrastructure", "config_type"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_classic_infrastructure.sm_configuration_public_certificate_dns_classic_infrastructure", "secret_type"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_classic_infrastructure.sm_configuration_public_certificate_dns_classic_infrastructure", "created_by"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_classic_infrastructure.sm_configuration_public_certificate_dns_classic_infrastructure", "created_at"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_classic_infrastructure.sm_configuration_public_certificate_dns_classic_infrastructure", "updated_at"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_classic_infrastructure.sm_configuration_public_certificate_dns_classic_infrastructure", "classic_infrastructure_username"),
	//				resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_public_certificate_dns_classic_infrastructure.sm_configuration_public_certificate_dns_classic_infrastructure", "classic_infrastructure_password"),
	//			),
	//		},
	//	},
	//})
}

func testAccCheckIbmSmConfigurationPublicCertificateDnsClassicInfrastructureDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_configuration_public_certificate_dns_classic_infrastructure" "sm_configuration_public_certificate_dns_classic_infrastructure_instance" {
			configuration_prototype {
				config_type = "public_cert_configuration_ca_lets_encrypt"
				name = "my-example-engine-config"
				lets_encrypt_environment = "production"
				lets_encrypt_private_key = "lets_encrypt_private_key"
				lets_encrypt_preferred_chain = "lets_encrypt_preferred_chain"
			}
		}

		data "ibm_sm_configuration_public_certificate_dns_classic_infrastructure" "sm_configuration_public_certificate_dns_classic_infrastructure_instance" {
			name = "configuration-name"
		}
	`)
}
