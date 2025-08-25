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

var publicCertName = "terraform-test-public-cert"
var modifiedPublicCertName = "modified-terraform-test-public-cert"

func TestAccIbmSmPublicCertificateBasic(t *testing.T) {
	resourceName := "ibm_sm_public_certificate.sm_public_certificate_basic"
	commonName := generatePublicCertCommonName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPublicCertificateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: publicCertificateConfigBasic(commonName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer"),
					resource.TestCheckResourceAttrSet(resourceName, "common_name"),
					resource.TestCheckResourceAttrSet(resourceName, "bundle_certs"),
					resource.TestCheckResourceAttrSet(resourceName, "ca"),
					resource.TestCheckResourceAttrSet(resourceName, "dns"),
					resource.TestCheckResourceAttrSet(resourceName, "key_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "signing_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_before"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_after"),
					resource.TestCheckResourceAttrSet(resourceName, "issuance_info.0.%"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate"),
					resource.TestCheckResourceAttrSet(resourceName, "private_key"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at"},
			},
		},
	})
}

func TestAccIbmSmPublicCertificateAllArgs(t *testing.T) {
	resourceName := "ibm_sm_public_certificate.sm_public_certificate"
	commonName := generatePublicCertCommonName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPublicCertificateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: publicCertificateConfigAllArgs(commonName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPublicCertificateCreated(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer"),
					resource.TestCheckResourceAttrSet(resourceName, "common_name"),
					resource.TestCheckResourceAttrSet(resourceName, "bundle_certs"),
					resource.TestCheckResourceAttrSet(resourceName, "ca"),
					resource.TestCheckResourceAttrSet(resourceName, "dns"),
					resource.TestCheckResourceAttrSet(resourceName, "key_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "signing_algorithm"),
					resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_before"),
					resource.TestCheckResourceAttrSet(resourceName, "validity.0.not_after"),
					resource.TestCheckResourceAttrSet(resourceName, "issuance_info.0.%"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate"),
					resource.TestCheckResourceAttrSet(resourceName, "private_key"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
				),
			},
			{
				Config: publicCertificateConfigUpdated(commonName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPublicCertificateUpdated(resourceName),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at"},
			},
		},
	})
}

var publicCertBasicConfigFormat = `
		resource "ibm_sm_public_certificate" "sm_public_certificate_basic" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			common_name = "%s"
  			ca = ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt_instance.name
  			dns = ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis_instance.name
		}`

var publicCertFullConfigFormat = `
		resource "ibm_sm_public_certificate" "sm_public_certificate" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
  			custom_metadata = %s
			secret_group_id = "default"
  			common_name = "%s"
  			ca = ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt_instance.name
  			dns = ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis_instance.name
			rotation {
				auto_rotate = true
				rotate_keys = %s
			}
		}`

func letsEncryptCaConfig() string {
	return fmt.Sprintf(`
		resource "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "public_cert_ca_lets_encrypt-terraform-test-datasource"
			lets_encrypt_environment = "%s"
			lets_encrypt_private_key = "%s"
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerPublicCertificateLetsEncryptEnvironment,
		acc.SecretsManagerPublicCertificateLetsEncryptPrivateKey)
}

func dnsCisConfig() string {
	return fmt.Sprintf(`
		resource "ibm_sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis_instance" {
			  instance_id   = "%s"
       		  region        = "%s"
			  cloud_internet_services_crn = "%s"
  			  name = "cloud-internet-services-config-terraform-test"
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateCisCrn)
}

func publicCertificateConfigBasic(commonName string) string {
	return letsEncryptCaConfig() + dnsCisConfig() +
		fmt.Sprintf(publicCertBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			publicCertName, commonName)
}

func publicCertificateConfigAllArgs(commonName string) string {
	rotateKeys := "false"
	return letsEncryptCaConfig() + dnsCisConfig() +
		fmt.Sprintf(publicCertFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			publicCertName, description, label, customMetadata, commonName,
			rotateKeys)
}

func publicCertificateConfigUpdated(commonName string) string {
	rotateKeys := "true"
	return letsEncryptCaConfig() + dnsCisConfig() +
		fmt.Sprintf(publicCertFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			modifiedPublicCertName, modifiedDescription, modifiedLabel, modifiedCustomMetadata, commonName,
			rotateKeys)
}

func testAccCheckIbmSmPublicCertificateCreated(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		publicCertIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := publicCertIntf.(*secretsmanagerv2.PublicCertificate)

		if err := verifyAttr(*secret.Name, publicCertName, "secret name"); err != nil {
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
		if err := verifyAttr(getRotateKeys(secret.Rotation), "false", "rotate_keys"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmPublicCertificateUpdated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		publicCertIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := publicCertIntf.(*secretsmanagerv2.PublicCertificate)
		if err := verifyAttr(*secret.Name, modifiedPublicCertName, "secret name after update"); err != nil {
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
		if err := verifyAttr(getAutoRotate(secret.Rotation), "true", "auto_rotate after update"); err != nil {
			return err
		}
		if err := verifyAttr(getRotateKeys(secret.Rotation), "true", "rotate_keys after update"); err != nil {
			return err
		}
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
