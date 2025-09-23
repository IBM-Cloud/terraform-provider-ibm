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

func TestAccIbmSmPublicCertificateConfigurationDNSClassicInfrastructureBasic(t *testing.T) {
	var conf secretsmanagerv2.PublicCertificateConfigurationDNSClassicInfrastructure

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmPublicCertificateConfigurationDNSClassicInfrastructureDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPublicCertificateConfigurationDNSClassicInfrastructureConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPublicCertificateConfigurationDNSClassicInfrastructureExists("ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmPublicCertificateConfigurationDNSClassicInfrastructureConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" "sm_public_certificate_configuration_dns_classic_infrastructure" {
			instance_id   = "%s"
			region        = "%s"
  			classic_infrastructure_username = "%s"
			classic_infrastructure_password = "%s"
  			name = "classic-infrastructure-config-terraform-test"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateClassicInfrastructureUsername, acc.SecretsManagerPublicCertificateClassicInfrastructurePassword)
}

func testAccCheckIbmSmPublicCertificateConfigurationDNSClassicInfrastructureExists(n string, obj secretsmanagerv2.PublicCertificateConfigurationDNSClassicInfrastructure) resource.TestCheckFunc {

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

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		publicCertificateConfigurationDNSClassicInfrastructureIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		publicCertificateConfigurationDNSClassicInfrastructure := publicCertificateConfigurationDNSClassicInfrastructureIntf.(*secretsmanagerv2.PublicCertificateConfigurationDNSClassicInfrastructure)
		obj = *publicCertificateConfigurationDNSClassicInfrastructure
		return nil
	}
}

func testAccCheckIbmSmPublicCertificateConfigurationDNSClassicInfrastructureDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		// Try to find the key
		_, response, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)

		if err == nil {
			return fmt.Errorf("PublicCertificateConfigurationDNSClassicInfrastructure still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PublicCertificateConfigurationDNSClassicInfrastructure (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
