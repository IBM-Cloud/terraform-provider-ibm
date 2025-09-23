// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func TestAccIbmSmPublicCertificateActionValidateManualDnsBasic(t *testing.T) {

	var conf secretsmanagerv2.PublicCertificate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSmPublicCertificateActionValidateManualDnsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSmPublicCertificateActionValidateManualDnsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSmPublicCertificateActionValidateManualDnsExists("ibm_sm_public_certificate.sm_public_certificate", conf),
				),
			},
		},
	})
}

func testAccCheckSmPublicCertificateActionValidateManualDnsConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "public_cert_ca_lets_encrypt-terraform-test-datasource"
			lets_encrypt_environment = "%s"
			lets_encrypt_private_key = "%s"
		}

		resource "ibm_sm_public_certificate" "sm_public_certificate" {
			instance_id = "%s"
			region = "%s"
  			name = "public-certificate-terraform-tests"
  			secret_group_id = "default"
  			common_name = "%s"
  			ca = ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt_instance.name
  			dns = "manual"
		}
		
		resource "ibm_cis_dns_record" "test_dns_txt_record" {
			cis_id  = "%s"
			domain_id = "%s"
  			name    = ibm_sm_public_certificate.sm_public_certificate.issuance_info[0].challenges[0].txt_record_name
  			type    = "TXT"
  			content = ibm_sm_public_certificate.sm_public_certificate.issuance_info[0].challenges[0].txt_record_value
		}

		resource "ibm_sm_public_certificate_action_validate_manual_dns" "sm_public_certificate_action_validate_manual_dns_instance" {
			instance_id   = "%s"
			region        = "%s"
			secret_id     = ibm_sm_public_certificate.sm_public_certificate.secret_id
			depends_on = [
				ibm_cis_dns_record.test_dns_txt_record
			]
		}


	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateLetsEncryptEnvironment, acc.SecretsManagerPublicCertificateLetsEncryptPrivateKey,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateCommonName,
		acc.SecretsManagerPublicCertificateCisCrn, acc.SecretsManagerValidateManualDnsCisZoneId,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckSmPublicCertificateActionValidateManualDnsExists(n string, obj secretsmanagerv2.PublicCertificate) resource.TestCheckFunc {

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

		if *publicCertificate.StateDescription != "active" {
			return fmt.Errorf("Error checking for PublicCertificateActionValidateManualDNS (%s): Secret's state is: %s, instead of 'active'", rs.Primary.ID, *publicCertificate.StateDescription)

		}
		obj = *publicCertificate
		return nil
	}
}

func testAccCheckSmPublicCertificateActionValidateManualDnsDestroy(s *terraform.State) error {
	return nil
}
