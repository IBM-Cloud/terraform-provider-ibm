// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIbmSmPublicCertificateConfigurationDnsClassicInfrastructureDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPublicCertificateConfigurationDnsClassicInfrastructureDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", "config_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", "classic_infrastructure_username"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", "classic_infrastructure_password"),
				),
			},
		},
	})
}

func testAccCheckIbmSmPublicCertificateConfigurationDnsClassicInfrastructureDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" "sm_public_certificate_configuration_dns_classic_infrastructure_instance" {
			instance_id   = "%s"
			region        = "%s"
  			classic_infrastructure_username = "%s"
			classic_infrastructure_password = "%s"
  			name = "classic-infrastructure-config-terraform-test"
		}

		data "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" "sm_public_certificate_configuration_dns_classic_infrastructure" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure_instance.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateClassicInfrastructureUsername, acc.SecretsManagerPublicCertificateClassicInfrastructurePassword, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
