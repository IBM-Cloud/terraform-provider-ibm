// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmConfigurationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmConfigurationsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_configurations.sm_configurations", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configurations.sm_configurations", "configurations.#"),
				),
			},
		},
	})
}

func TestAccIbmSmConfigurationsDataSourceCryptoKey(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmConfigurationsDataSourceConfigCryptoKey(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_configurations.sm_configurations", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configurations.sm_configurations", "configurations.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSmConfigurationsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource-iam-configuration"
			api_key = "%s"
		}
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
		resource "ibm_sm_private_certificate_configuration_template" "ibm_sm_private_certificate_configuration_template_instance" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_private_certificate_configuration_intermediate_ca.ibm_sm_private_certificate_configuration_intermediate_ca_instance.name
			allow_any_name = true
			name = "template-terraform-private-cert-datasource-test"
		}
		resource "ibm_sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis_instance" {
			  instance_id   = "%s"
       		  region        = "%s"
			  cloud_internet_services_crn = "%s"
  			  name = "cloud-internet-services-config-terraform-test"
		}
		resource "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" "sm_public_certificate_configuration_dns_classic_infrastructure_instance" {
			instance_id   = "%s"
			region        = "%s"
  			classic_infrastructure_username = "%s"
			classic_infrastructure_password = "%s"
  			name = "classic-infrastructure-config-terraform-test"
		}
		data "ibm_sm_configurations" "sm_configurations" {
			instance_id   = "%s"
			region        = "%s"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsConfigurationApiKey,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerPublicCertificateCisCrn, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerPublicCertificateClassicInfrastructureUsername, acc.SecretsManagerPublicCertificateClassicInfrastructurePassword,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmConfigurationsDataSourceConfigCryptoKey() string {
	return privateCertificateIntermediateCAConfigCryptoKey() + fmt.Sprintf(`
		data "ibm_sm_configurations" "sm_configurations" {
			instance_id   = "%s"
			region        = "%s"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
