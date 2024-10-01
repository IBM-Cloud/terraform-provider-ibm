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

func TestAccIbmSmPrivateCertificateConfigurationIntermediateCABasic(t *testing.T) {
	resourceName := "ibm_sm_private_certificate_configuration_intermediate_ca.private_cert_intermediate_ca_basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: privateCertificateIntermediateCAConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "max_ttl_seconds"),
					resource.TestCheckResourceAttrSet(resourceName, "crl_expiry_seconds"),
					resource.TestCheckResourceAttrSet(resourceName, "serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "expiration_date"),
					resource.TestCheckResourceAttrSet(resourceName, "data.0.certificate"),
					resource.TestCheckResourceAttr(resourceName, "status", "certificate_template_required"),
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

func TestAccIbmSmPrivateCertificateConfigurationIntermediateCAllArgs(t *testing.T) {
	resourceName := "ibm_sm_private_certificate_configuration_intermediate_ca.private_cert_intermediate_ca"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: privateCertificateIntermediateCAConfigAllArgs("9980000", "10h", "true", "false", "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCAExists(resourceName, 9980000, 36000, true, false, true),
				),
			},
			resource.TestStep{
				Config: privateCertificateIntermediateCAConfigAllArgs("8870001", "20h", "false", "true", "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCAExists(resourceName, 8870001, 72000, false, true, false),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"crl_expiry", "max_ttl", "max_path_length",
					"permitted_dns_domains", "ttl", "use_csr_values"},
			},
		},
	})
}

func TestAccIbmSmPrivateCertificateConfigurationIntermediateCACryptoKey(t *testing.T) {
	resourceName := "ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_cert_intermediate_ca_crypto_key"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: privateCertificateIntermediateCAConfigCryptoKey(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCAExists(resourceName, 94680000., 259200, false, true, true),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"crl_expiry", "max_ttl", "max_path_length",
					"permitted_dns_domains", "ttl", "use_csr_values"},
			},
		},
	})
}

func rootCaConfig() string {
	return fmt.Sprintf(`

		resource "ibm_sm_private_certificate_configuration_root_ca" "root_ca_terraform" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "9990000"
			common_name = "ibm.com"
			name = "root-ca-terraform-private-cert-datasource-test"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func privateCertificateIntermediateCAConfigBasic() string {
	return rootCaConfig() + fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_intermediate_ca" "private_cert_intermediate_ca_basic" {
  			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			issuer = ibm_sm_private_certificate_configuration_root_ca.root_ca_terraform.name
			signing_method = "internal"
			name = "intermediate-ca-terraform-private-cert-test-basic"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func privateCertificateIntermediateCAConfigAllArgs(maxTtl, crlExpiry, crlDisable,
	crlDistributionPointsEncoded, issuingCertificatesUrlsEncoded string) string {
	return rootCaConfig() + fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_intermediate_ca" "private_cert_intermediate_ca" {
  			instance_id   = "%s"
			region        = "%s"
			name = "intermediate-ca-terraform-private-cert-test"
			max_ttl = "%s"
			common_name = "ibm.com"
			signing_method = "internal"
			issuer = ibm_sm_private_certificate_configuration_root_ca.root_ca_terraform.name
			crl_expiry = "%s"
			crl_disable = %s
			crl_distribution_points_encoded = %s
			issuing_certificates_urls_encoded = %s
			alt_names = ["ibm.com","example.com"]
			ip_sans = "90.180.210.30, 80.111.222.33"
			uri_sans = "http://www.example.com, http://www.ibm.com"
			format = "pem_bundle"
			private_key_format = "pkcs8"
			key_type = "ec"
			key_bits = 521
			exclude_cn_from_sans = true
			ou = ["ou1", "ou2"]
			organization = ["org1", "org2"]
			country = ["us"]
			locality = ["San Francisco"]
			province = ["PV"]
			street_address = ["123 Main St."]
			postal_code = ["12345"]
			ttl = "5000000"
			permitted_dns_domains = ["example.com"]
			use_csr_values = true
			max_path_length = 80
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, maxTtl, crlExpiry, crlDisable,
		crlDistributionPointsEncoded, issuingCertificatesUrlsEncoded)
}

func privateCertificateIntermediateCAConfigCryptoKey() string {
	return privateCertificateRootCAConfigCryptoKey() + fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_intermediate_ca" "sm_private_cert_intermediate_ca_crypto_key" {
			depends_on     = [ibm_sm_private_certificate_configuration_root_ca.sm_private_cert_root_ca_crypto_key]
			instance_id   = "%s"
			region        = "%s"
			name = "intermediate-ca-terraform-private-cert-test"
			max_ttl = "26300h"
			ttl = "2190h"
			issuing_certificates_urls_encoded = true
			crl_distribution_points_encoded = true
			crl_disable = false
			key_type = "rsa"
			key_bits = 4096
			signing_method = "internal"
			issuer = ibm_sm_private_certificate_configuration_root_ca.sm_private_cert_root_ca_crypto_key.name
			common_name = "ibm.com"
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

func testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCAExists(resourceName string, maxTtl, crlExpiry int, crlDisable,
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

		privateCertificateConfigurationIntermediateCAIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		privateCertificateConfigurationIntermediateCA := privateCertificateConfigurationIntermediateCAIntf.(*secretsmanagerv2.PrivateCertificateConfigurationIntermediateCA)
		intermediateCA := *privateCertificateConfigurationIntermediateCA
		if err := verifyAttr(*intermediateCA.Name, "intermediate-ca-terraform-private-cert-test", "configuration name"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*intermediateCA.MaxTtlSeconds), maxTtl, "max_ttl_seconds"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*intermediateCA.CrlExpirySeconds), crlExpiry, "CRL expiry seconds"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*intermediateCA.CrlDisable, crlDisable, "CRL disable"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*intermediateCA.CrlDistributionPointsEncoded, crlDistributionPointsEncoded, "crlDistributionPointsEncoded"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*intermediateCA.IssuingCertificatesUrlsEncoded, issuingCertificatesUrlsEncoded, "issuingCertificatesUrlsEncoded"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmPrivateCertificateConfigurationIntermediateCADestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_private_certificate_configuration_intermediate_ca" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		// Try to find the key
		_, response, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)

		if err == nil {
			return fmt.Errorf("PrivateCertificateConfigurationIntermediateCA still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PrivateCertificateConfigurationIntermediateCA (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
