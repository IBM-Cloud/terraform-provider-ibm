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

func TestAccIbmSmPrivateCertificateConfigurationTemplateBasic(t *testing.T) {
	var resourceName = "ibm_sm_private_certificate_configuration_template.sm_private_cert_template_basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateConfigurationTemplateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: configTemplateBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testTemplateExistsBasic(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "ttl_seconds"),
					resource.TestCheckResourceAttrSet(resourceName, "max_ttl_seconds"),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIbmSmPrivateCertificateConfigurationTemplateAllArgs(t *testing.T) {
	var resourceName = "ibm_sm_private_certificate_configuration_template.sm_private_cert_template"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPrivateCertificateConfigurationTemplateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: configTemplateAllArgs(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testTemplateExistsAllArgs(resourceName),
				),
			},
			resource.TestStep{
				Config: configTemplateUpdate(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testTemplateUpdated(resourceName),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "max_ttl", "not_before_duration", "allow_wildcard_certificates"},
			},
		},
	})
}

func rootCaConfigForTemplate() string {
	return fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_root_ca" "root_ca_instance" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			name = "root-ca-terraform-private-cert-test"
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func intermediateCaConfig() string {
	return fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_intermediate_ca" "intermediate_ca_instance" {
  			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			issuer = ibm_sm_private_certificate_configuration_root_ca.root_ca_instance.name
			signing_method = "internal"
			name = "intermediate-ca-terraform-private-cert-test"
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func configTemplateBasic() string {
	return rootCaConfigForTemplate() + intermediateCaConfig() + fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_template" "sm_private_cert_template_basic" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_private_certificate_configuration_intermediate_ca.intermediate_ca_instance.name
			name = "template-terraform-test-basic"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func configTemplateAllArgs() string {
	return rootCaConfigForTemplate() + intermediateCaConfig() + fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_template" "sm_private_cert_template" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_private_certificate_configuration_intermediate_ca.intermediate_ca_instance.name
			name = "template-terraform-test"
			allowed_other_sans = ["1.3.6.1.4.1.1;UTF8:*"]
			max_ttl = "24h"
			allow_localhost = false
			allowed_domains = ["example.com"]
			allow_subdomains = false
			allow_bare_domains = false
			allow_glob_domains = false
			allow_any_name = false
			enforce_hostnames = false
			allow_ip_sans = false
			allowed_uri_sans = ["example.com"]
			server_flag = false
			client_flag = false
			code_signing_flag = false
			email_protection_flag = false
			key_type = "ec"
			key_bits = 384
			key_usage = ["DigitalSignature", "KeyAgreement", "KeyEncipherment"]
			ext_key_usage = ["anyExtendedKeyUsage"]
			ext_key_usage_oids = []
			use_csr_common_name = false
			use_csr_sans = false
			ttl = "10h"
			ou = ["ou1", "ou2"]
			organization = ["org1", "org2"]
			country = ["us"]
			locality = ["San Francisco"]
			province = ["PV"]
			street_address = ["123 Main St."]
			postal_code = ["12345"]
			require_cn = false
			policy_identifiers = []
			basic_constraints_valid_for_non_ca = false
			not_before_duration = "55s"
			allow_wildcard_certificates = false
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func configTemplateUpdate() string {
	return rootCaConfigForTemplate() + intermediateCaConfig() + fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_template" "sm_private_cert_template" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_private_certificate_configuration_intermediate_ca.intermediate_ca_instance.name
			name = "template-terraform-test"
			allowed_other_sans = ["1.3.6.1.4.1.1;UTF8:*"]
			max_ttl = "48h"
			allowed_domains_template = true
			allow_localhost = true
			allowed_domains = ["example.com"]
			allow_subdomains = true
			allow_bare_domains = true
			allow_glob_domains = true
			allow_any_name = true
			enforce_hostnames = true
			allow_ip_sans = true
			allowed_uri_sans = ["example.com"]
			server_flag = true
			client_flag = true
			code_signing_flag = true
			email_protection_flag = true
			key_type = "ec"
			key_bits = 384
			key_usage = ["DigitalSignature", "KeyAgreement", "KeyEncipherment"]
			ext_key_usage = ["anyExtendedKeyUsage"]
			ext_key_usage_oids = []
			use_csr_common_name = true
			use_csr_sans = true
			ttl = "20h"
			ou = ["ou1", "ou2"]
			organization = ["org1", "org2"]
			country = ["us"]
			locality = ["San Francisco"]
			province = ["PV"]
			street_address = ["123 Main St."]
			postal_code = ["12345"]
			require_cn = true
			policy_identifiers = []
			basic_constraints_valid_for_non_ca = true
			not_before_duration = "66s"
			allow_wildcard_certificates = false
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func getTemplate(s *terraform.State, resourceName string) (*secretsmanagerv2.PrivateCertificateConfigurationTemplate, error) {
	rs, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return nil, fmt.Errorf("Not found: %s", resourceName)
	}

	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return nil, err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

	id := strings.Split(rs.Primary.ID, "/")
	configName := id[2]
	getConfigurationOptions.SetName(configName)

	privateCertificateConfigurationTemplateIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
	if err != nil {
		return nil, err
	}

	template := privateCertificateConfigurationTemplateIntf.(*secretsmanagerv2.PrivateCertificateConfigurationTemplate)
	return template, nil
}

func testTemplateExistsBasic(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		template, err := getTemplate(s, resourceName)
		if err != nil {
			return err
		}
		if err := verifyAttr(*template.Name, "template-terraform-test-basic", "configuration name"); err != nil {
			return err
		}
		return nil
	}
}

func testTemplateExistsAllArgs(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		template, err := getTemplate(s, resourceName)
		if err != nil {
			return err
		}
		if err := verifyAttr(*template.Name, "template-terraform-test", "configuration name"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowLocalhost, false, "Allow Localhost"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowAnyName, false, "Allow Any Name"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowBareDomains, false, "Allow Bare Domains"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowSubdomains, false, "Allow Subdomains"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowedDomainsTemplate, false, "Allowed Domains Template"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowGlobDomains, false, "Allow Glob Domains"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowIpSans, false, "Allow Ip Sans"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.ServerFlag, false, "Server Flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.ClientFlag, false, "Client Flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.CodeSigningFlag, false, "Code signing Flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.EmailProtectionFlag, false, "Email Protection Flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.UseCsrSans, false, "Use Csr Sans"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.UseCsrCommonName, false, "Use Csr Common Name"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.RequireCn, false, "Require Cn"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.BasicConstraintsValidForNonCa, false, "BasicConstraintsValidForNonCa"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*template.TtlSeconds), 36000, "TTL"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*template.MaxTtlSeconds), 86400, "MaxTTL"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*template.NotBeforeDurationSeconds), 55, "NotBeforeDuration"); err != nil {
			return err
		}
		if err := verifyIntAttr(len(template.AllowedDomains), 1, "Num allowed domains"); err != nil {
			return err
		}
		if err := verifyAttr(template.AllowedDomains[0], "example.com", "allowed domain"); err != nil {
			return err
		}
		if err := verifyIntAttr(len(template.AllowedUriSans), 1, "Num allowed URI sans"); err != nil {
			return err
		}
		if err := verifyAttr(template.AllowedUriSans[0], "example.com", "allowed URI sans"); err != nil {
			return err
		}
		return nil
	}
}

func testTemplateUpdated(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		template, err := getTemplate(s, resourceName)
		if err != nil {
			return err
		}
		if err := verifyAttr(*template.Name, "template-terraform-test", "configuration name"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowLocalhost, true, "Allow Localhost"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowAnyName, true, "Allow Any Name"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowBareDomains, true, "Allow Bare Domains"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowSubdomains, true, "Allow Subdomains"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowedDomainsTemplate, true, "Allowed Domains Template"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowGlobDomains, true, "Allow Glob Domains"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.AllowIpSans, true, "Allow Ip Sans"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.ServerFlag, true, "Server Flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.ClientFlag, true, "Client Flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.CodeSigningFlag, true, "Code signing Flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.EmailProtectionFlag, true, "Email Protection Flag"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.UseCsrSans, true, "Use Csr Sans"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.UseCsrCommonName, true, "Use Csr Common Name"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.RequireCn, true, "Require Cn"); err != nil {
			return err
		}
		if err := verifyBoolAttr(*template.BasicConstraintsValidForNonCa, true, "BasicConstraintsValidForNonCa"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*template.TtlSeconds), 72000, "TTL"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*template.MaxTtlSeconds), 172800, "MaxTTL"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(*template.NotBeforeDurationSeconds), 66, "NotBeforeDuration"); err != nil {
			return err
		}
		if err := verifyIntAttr(len(template.AllowedDomains), 1, "Num allowed domains"); err != nil {
			return err
		}
		if err := verifyAttr(template.AllowedDomains[0], "example.com", "allowed domain"); err != nil {
			return err
		}
		if err := verifyIntAttr(len(template.AllowedUriSans), 1, "Num allowed URI sans"); err != nil {
			return err
		}
		if err := verifyAttr(template.AllowedUriSans[0], "example.com", "allowed URI sans"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmPrivateCertificateConfigurationTemplateDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_private_certificate_configuration_template" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		// Try to find the key
		_, response, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)

		if err == nil {
			return fmt.Errorf("PrivateCertificateConfigurationTemplate still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PrivateCertificateConfigurationTemplate (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
