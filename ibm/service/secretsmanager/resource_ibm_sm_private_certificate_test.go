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

func TestAccIbmSmPrivateCertificateBasic(t *testing.T) {
	var conf secretsmanagerv2.PrivateCertificate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPrivateCertificateConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateExists("ibm_sm_private_certificate.sm_private_certificate", conf),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_sm_private_certificate.sm_private_certificate",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key_format", "ttl"},
			},
		},
	})
}

func testAccCheckIbmSmPrivateCertificateConfigBasic() string {
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
		resource "ibm_sm_private_certificate_configuration_template" "sm_private_certificate_configuration_template_instance" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_private_certificate_configuration_intermediate_ca.ibm_sm_private_certificate_configuration_intermediate_ca_instance.name
			allow_any_name = true
			name = "template-terraform-private-cert-test"
		}

		resource "ibm_sm_private_certificate" "sm_private_certificate" {
			instance_id = "%s"
			region = "%s"
		  	name = "private_cert_terraform-test"
  			description = "Extended description for this secret."
  			labels = [ "my-label", "another"]
  			custom_metadata = {"key":"value"}
  			certificate_template = ibm_sm_private_certificate_configuration_template.sm_private_certificate_configuration_template_instance.name
  			common_name = "ibm.com"
  			ttl = "1800"
  			secret_group_id = "default"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID,
		acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmPrivateCertificateExists(n string, obj secretsmanagerv2.PrivateCertificate) resource.TestCheckFunc {

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

		privateCertificateIntf, _, err := secretsManagerClient.GetSecret(getSecretOptions)
		if err != nil {
			return err
		}

		privateCertificate := privateCertificateIntf.(*secretsmanagerv2.PrivateCertificate)
		obj = *privateCertificate
		return nil
	}
}

func testAccCheckIbmSmPrivateCertificateDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_private_certificate" {
			continue
		}

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("PrivateCertificate still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PrivateCertificate (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
