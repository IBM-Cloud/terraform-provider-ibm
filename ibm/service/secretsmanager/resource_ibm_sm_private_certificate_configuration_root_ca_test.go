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

func TestAccIbmSmPrivateCertificateConfigurationRootCABasic(t *testing.T) {
	resourceName := "ibm_sm_private_certificate_configuration_root_ca.sm_private_cert_root_ca_basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateConfigurationRootCADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: privateCertificateRootCAConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "max_ttl_seconds"),
					resource.TestCheckResourceAttrSet(resourceName, "ttl_seconds"),
					resource.TestCheckResourceAttrSet(resourceName, "crl_expiry_seconds"),
					resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "expiration_date"),
					resource.TestCheckResourceAttrSet(resourceName, "data.0.certificate"),
					resource.TestCheckResourceAttr(resourceName, "status", "configured"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"max_ttl"},
			},
		},
	})
}

func TestAccIbmSmPrivateCertificateConfigurationRootCAllArgs(t *testing.T) {
	resourceName := "ibm_sm_private_certificate_configuration_root_ca.sm_private_cert_root_ca"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateConfigurationRootCADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: privateCertificateRootCAConfigAllArgs("250000", "10h", "true", "true", "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateConfigurationRootCAExists(resourceName, 250000, 36000, true, true, true),
				),
			},
			resource.TestStep{
				Config: privateCertificateRootCAConfigAllArgs("12345", "20h", "false", "false", "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateConfigurationRootCAExists(resourceName, 12345, 72000, false, false, false),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"crl_expiry", "max_ttl", "ttl"},
			},
		},
	})
}

func TestAccIbmSmPrivateCertificateConfigurationRootCACryptoKey(t *testing.T) {
	resourceName := "ibm_sm_private_certificate_configuration_root_ca.sm_private_cert_root_ca_crypto_key"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateConfigurationRootCADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: privateCertificateRootCAConfigCryptoKey(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateConfigurationRootCAExists(resourceName, 157788000, 259200, false, true, true),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"crl_expiry", "max_ttl", "ttl"},
			},
		},
	})
}

var rootCaBasicConfigFormat = `
		resource "ibm_sm_private_certificate_configuration_root_ca" "sm_private_cert_root_ca_basic" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			name = "root-ca-terraform-private-cert--test-basic"
		}`

var rootCaFullConfigFormat = `
		resource "ibm_sm_private_certificate_configuration_root_ca" "sm_private_cert_root_ca" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "%s"
			crl_expiry = "%s"
   			crl_disable = %s
			crl_distribution_points_encoded = %s
			issuing_certificates_urls_encoded = %s
			common_name = "ibm.com"
			name = "root-ca-terraform-private-cert-test"
			alt_names = ["ibm.com", "example.com"]
			ip_sans = "90.180.210.30, 80.111.222.33"
			uri_sans = "http://www.example.com, http://www.ibm.com"
			other_sans = ["1.3.6.1.4.1.1;UTF8:some_value"]
			ttl = "10h"
			format = "pem_bundle"
			private_key_format = "pkcs8"
			key_type = "ec"
			key_bits = 384
			max_path_length = 80
			exclude_cn_from_sans = true
			permitted_dns_domains  = ["example.com"]
			ou = ["ou1", "ou2"]
			organization = ["org1", "org2"]
			country = ["us"]
			locality = ["San Francisco"]
			province = ["PV"]
			street_address = ["123 Main St."]
			postal_code = ["12345"]
		}`

func iamCredentialsSecretConfigCryptoKey() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret_instance_crypto_key" {
			instance_id   = "%s"
			region        = "%s"
			name = "iam-credentials-for-crypto-key-terraform-tests"
            service_id = "%s"
  			reuse_api_key = true
  			ttl = "259200"
			rotation {
				auto_rotate = true
				interval = 1
				unit = "day"
			}
			depends_on = [
				ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration_instance
			]
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerPrivateCertificateConfigurationCryptoKeyIAMSecretServiceId)
}

func privateCertificateRootCAConfigCryptoKey() string {
	return iamCredentialsEngineConfig() + iamCredentialsSecretConfigCryptoKey() + fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_root_ca" "sm_private_cert_root_ca_crypto_key" {
			depends_on     = [ibm_sm_iam_credentials_secret.sm_iam_credentials_secret_instance_crypto_key]
			instance_id   = "%s"
			region        = "%s"
			name = "root-ca-terraform-private-cert-test"
			max_ttl       = "43830h"
			ttl = "2190h"
			crl_disable = false
			crl_expiry = "72h"
			crl_distribution_points_encoded = true
			issuing_certificates_urls_encoded = true
			key_type = "rsa"
			key_bits = 4096
			common_name   = "ibm.com"
			alt_names = ["ddd.com", "aaa.com"]
			crypto_key {
				allow_generate_key = true
				label = "e2e-tf-test"
				provider {
					type = "%s"
					instance_crn = "%s"
					pin_iam_credentials_secret_id = ibm_sm_iam_credentials_secret.sm_iam_credentials_secret_instance_crypto_key.secret_id
					private_keystore_id = "%s"
				}
			}
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderType,
		acc.SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderInstanceCrn,
		acc.SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderPrivateKeystoreId)
}

func privateCertificateRootCAConfigBasic() string {
	return fmt.Sprintf(rootCaBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func privateCertificateRootCAConfigAllArgs(maxTtl, crlExpiry, crlDisable,
	crlDistributionPointsEncoded, issuingCertificatesUrlsEncoded string) string {
	return fmt.Sprintf(rootCaFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		maxTtl, crlExpiry, crlDisable, crlDistributionPointsEncoded, issuingCertificatesUrlsEncoded)
}

func testAccCheckIbmSmPrivateCertificateConfigurationRootCAExists(resourceName string, maxTtl, crlExpiry int, crlDisable,
	crlDistributionPointsEncoded, issuingCertificatesUrlsEncoded bool) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
		if err != nil {
			return err
		}

		secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		privateCertificateConfigurationRootCAIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		rootCA := privateCertificateConfigurationRootCAIntf.(*secretsmanagerv2.PrivateCertificateConfigurationRootCA)
		if err := verifyAttr(*rootCA.Name, "root-ca-terraform-private-cert-test", "configuration name"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*rootCA.MaxTtlSeconds), maxTtl, "max_ttl_seconds"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*rootCA.CrlExpirySeconds), crlExpiry, "CRL expiry seconds"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*rootCA.CrlDisable, crlDisable, "CRL disable"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*rootCA.CrlDistributionPointsEncoded, crlDistributionPointsEncoded, "crlDistributionPointsEncoded"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*rootCA.IssuingCertificatesUrlsEncoded, issuingCertificatesUrlsEncoded, "issuingCertificatesUrlsEncoded"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmPrivateCertificateConfigurationRootCADestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_private_certificate_configuration_root_ca" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		// Try to find the key
		_, response, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)

		if err == nil {
			return fmt.Errorf("PrivateCertificateConfigurationRootCA still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PrivateCertificateConfigurationRootCA (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
