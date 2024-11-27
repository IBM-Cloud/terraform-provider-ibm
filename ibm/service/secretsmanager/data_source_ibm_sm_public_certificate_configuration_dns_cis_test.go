// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmSmPublicCertificateConfigurationDnsCisDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPublicCertificateConfigurationDnsCisDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis", "config_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis", "cloud_internet_services_crn"),
				),
			},
		},
	})
}

func testAccCheckIbmSmPublicCertificateConfigurationDnsCisDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis_instance" {
			  instance_id   = "%s"
       		  region        = "%s"
			  cloud_internet_services_crn = "%s"
  			  name = "cloud-internet-services-config-terraform-test"
		}

		data "ibm_sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis" {
			  instance_id   = "%s"
       		  region        = "%s"
			  name = ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis_instance.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateCisCrn, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
