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

var privateCertName = "terraform-test-private-cert"
var modifiedPrivateCertName = "modified-terraform-test-private-cert"

func TestAccIbmSmPrivateCertificateBasic(t *testing.T) {
	resourceName := "ibm_sm_private_certificate.sm_private_certificate_basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: privateCertificateConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_authority"),
					resource.TestCheckResourceAttrSet(resourceName, "expiration_date"),
					resource.TestCheckResourceAttrSet(resourceName, "key_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "signing_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate"),
					resource.TestCheckResourceAttrSet(resourceName, "private_key"),
					resource.TestCheckResourceAttrSet(resourceName, "issuing_ca"),
					resource.TestCheckResourceAttrSet(resourceName, "ca_chain.0"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_before"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_after"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key_format", "ttl", "updated_at"},
			},
		},
	})
}

func TestAccIbmSmPrivateCertificateAllArgs(t *testing.T) {
	resourceName := "ibm_sm_private_certificate.sm_private_certificate"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: privateCertificateConfigAllArgs(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateCreated(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_authority"),
					resource.TestCheckResourceAttrSet(resourceName, "expiration_date"),
					resource.TestCheckResourceAttrSet(resourceName, "key_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "signing_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate"),
					resource.TestCheckResourceAttrSet(resourceName, "private_key"),
					resource.TestCheckResourceAttrSet(resourceName, "issuing_ca"),
					resource.TestCheckResourceAttrSet(resourceName, "ca_chain.0"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_before"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_after"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
				),
			},
			{
				Config: privateCertificateConfigUpdated(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateUpdated(resourceName),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key_format", "ip_sans", "format", "exclude_cn_from_sans", "ttl", "updated_at"},
			},
		},
	})
}

func configRootCa() string {
	return fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_root_ca" "ibm_sm_private_certificate_configuration_root_ca_instance" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "9990000"
			common_name = "ibm.com"
			crl_expiry = "10000h"
			name = "root-ca-terraform-private-cert-test"
		}
		`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func configIntermediateCa() string {
	return fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_intermediate_ca" "ibm_sm_private_certificate_configuration_intermediate_ca_instance" {
  			instance_id   = "%s"
			region        = "%s"
			max_ttl = "9980000"
			common_name = "ibm.com"
			issuer = ibm_sm_private_certificate_configuration_root_ca.ibm_sm_private_certificate_configuration_root_ca_instance.name
			signing_method = "internal"
			name = "intermediate-ca-terraform-private-cert-test"
		}
		`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func configCertificateTemplate() string {
	return fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_template" "sm_private_certificate_configuration_template_instance" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_private_certificate_configuration_intermediate_ca.ibm_sm_private_certificate_configuration_intermediate_ca_instance.name
			allow_any_name = true
			allowed_domains = ["example.com"]
			name = "template-terraform-private-cert-test"
		}
		`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

var privateCertBasicConfigFormat = `
		resource "ibm_sm_private_certificate" "sm_private_certificate_basic" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			certificate_template = ibm_sm_private_certificate_configuration_template.sm_private_certificate_configuration_template_instance.name
  			common_name = "ibm.com"
		}`

var privateCertFullConfigFormat = `
		resource "ibm_sm_private_certificate" "sm_private_certificate" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
  			custom_metadata = %s
			secret_group_id = "default"
  			certificate_template = ibm_sm_private_certificate_configuration_template.sm_private_certificate_configuration_template_instance.name
  			common_name = "ibm.com"
            alt_names = ["example.com"]
            ip_sans = ""
            other_sans = []
            format = "pem"
            private_key_format = "pkcs8"
            exclude_cn_from_sans = true
  			ttl = "7890000"
			rotation %s
		}`

func privateCertificateConfigBasic() string {
	return configRootCa() + configIntermediateCa() + configCertificateTemplate() +
		fmt.Sprintf(privateCertBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			privateCertName)
}

func privateCertificateConfigAllArgs() string {
	return configRootCa() + configIntermediateCa() + configCertificateTemplate() +
		fmt.Sprintf(privateCertFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			privateCertName, description, label, customMetadata, rotationPolicy)
}

func privateCertificateConfigUpdated() string {
	return configRootCa() + configIntermediateCa() + configCertificateTemplate() +
		fmt.Sprintf(privateCertFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			modifiedPrivateCertName, modifiedDescription, modifiedLabel, modifiedCustomMetadata, modifiedRotationPolicy)
}

func testAccCheckIbmSmPrivateCertificateCreated(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		privateCertIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := privateCertIntf.(*secretsmanagerv2.PrivateCertificate)

		if err := verifyAttr(*secret.Name, privateCertName, "secret name"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Description, description, "secret description"); err != nil {
			return err
		}
		if len(secret.Labels) != 1 {
			return fmt.Errorf("Wrong number of labels: %d", len(secret.Labels))
		}
		if err := verifyAttr(secret.Labels[0], label, "label"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, customMetadata, "custom metadata"); err != nil {
			return err
		}
		if err := verifyAttr(getAutoRotate(secret.Rotation), "true", "auto_rotate"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationUnit(secret.Rotation), "day", "rotation unit"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationInterval(secret.Rotation), "1", "rotation interval"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmPrivateCertificateUpdated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		privateCertIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := privateCertIntf.(*secretsmanagerv2.PrivateCertificate)
		if err := verifyAttr(*secret.Name, modifiedPrivateCertName, "secret name after update"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Description, modifiedDescription, "secret description after update"); err != nil {
			return err
		}
		if len(secret.Labels) != 1 {
			return fmt.Errorf("Wrong number of labels after update: %d", len(secret.Labels))
		}
		if err := verifyAttr(secret.Labels[0], modifiedLabel, "label after update"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, modifiedCustomMetadata, "custom metadata after update"); err != nil {
			return err
		}
		if err := verifyAttr(getAutoRotate(secret.Rotation), "true", "auto_rotate"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationUnit(secret.Rotation), "month", "rotation unit"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationInterval(secret.Rotation), "2", "rotation interval"); err != nil {
			return err
		}
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
