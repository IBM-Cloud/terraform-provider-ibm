// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

func TestAccIbmSmPublicCertificateBasic(t *testing.T) {
	var conf secretsmanagerv2.PublicCertificate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPublicCertificateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPublicCertificateConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPublicCertificateExists("ibm_sm_public_certificate.sm_public_certificate", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_public_certificate.sm_public_certificate",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmPublicCertificateConfigBasic() string {
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

		resource "ibm_sm_public_certificate" "sm_public_certificate" {
			instance_id = "%s"
			region = "%s"
  			name = "public-certificate-terraform-tests"
  			secret_group_id = "default"
  			common_name = "%s"
  			ca = ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt_instance.name
  			dns = ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis_instance.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateLetsEncryptEnvironment, acc.SecretsManagerPublicCertificateLetsEncryptPrivateKey,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateCisCrn,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateCommonName)
}

func testAccCheckIbmSmPublicCertificateExists(n string, obj secretsmanagerv2.PublicCertificate) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
		if err != nil {
			return err
		}

		secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		publicCertificateIntf, _, err := secretsManagerClient.GetSecret(getSecretOptions)
		if err != nil {
			return err
		}

		publicCertificate := publicCertificateIntf.(*secretsmanagerv2.PublicCertificate)
		obj = *publicCertificate
		return nil
	}
}

func testAccCheckIbmSmPublicCertificateDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_public_certificate" {
			continue
		}

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("PublicCertificate still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PublicCertificate (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
