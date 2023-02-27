// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/secretsmanagerv2"
)

func TestAccIbmSmConfigurationPrivateCertificateTemplateBasic(t *testing.T) {
	var conf secretsmanagerv2.PrivateCertificateConfigurationTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmConfigurationPrivateCertificateTemplateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmConfigurationPrivateCertificateTemplateConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmConfigurationPrivateCertificateTemplateExists("ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template", conf),
				),
			},
			//resource.TestStep{
			//	ResourceName:      "ibm_sm_configuration_private_certificate_template.sm_configuration_private_certificate_template",
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//},
		},
	})
}

func testAccCheckIbmSmConfigurationPrivateCertificateTemplateConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_configuration_private_certificate_root_ca" "ibm_sm_configuration_private_certificate_root_ca_instance" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			crl_expiry = "10000h"
			name = "root-ca-terraform-private-cert-datasource-test"
		}
		resource "ibm_sm_configuration_private_certificate_intermediate_ca" "ibm_sm_configuration_private_certificate_intermediate_ca_instance" {
  			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			issuer = ibm_sm_configuration_private_certificate_root_ca.ibm_sm_configuration_private_certificate_root_ca_instance.name
			signing_method = "internal"
			name = "intermediate-ca-terraform-private-cert-datasource-test"
		}
		resource "ibm_sm_configuration_private_certificate_template" "sm_configuration_private_certificate_template" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_configuration_private_certificate_intermediate_ca.ibm_sm_configuration_private_certificate_intermediate_ca_instance.name
			allow_any_name = true
			name = "template-terraform-private-cert-datasource-test"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmConfigurationPrivateCertificateTemplateExists(n string, obj secretsmanagerv2.PrivateCertificateConfigurationTemplate) resource.TestCheckFunc {

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

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		getConfigurationOptions.SetName(rs.Primary.ID)

		privateCertificateConfigurationTemplateIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		privateCertificateConfigurationTemplate := privateCertificateConfigurationTemplateIntf.(*secretsmanagerv2.PrivateCertificateConfigurationTemplate)
		obj = *privateCertificateConfigurationTemplate
		return nil
	}
}

func testAccCheckIbmSmConfigurationPrivateCertificateTemplateDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_configuration_private_certificate_template" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		getConfigurationOptions.SetName(rs.Primary.ID)

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
