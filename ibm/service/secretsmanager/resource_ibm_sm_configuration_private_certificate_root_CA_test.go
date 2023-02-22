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

func TestAccIbmSmConfigurationPrivateCertificateRootCABasic(t *testing.T) {
	var conf secretsmanagerv2.PrivateCertificateConfigurationRootCA

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmConfigurationPrivateCertificateRootCADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmConfigurationPrivateCertificateRootCAConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmConfigurationPrivateCertificateRootCAExists("ibm_sm_configuration_private_certificate_root_ca.sm_configuration_private_certificate_root_ca", conf),
				),
			},
			//resource.TestStep{
			//	ResourceName:      "ibm_sm_configuration_private_certificate_root_CA.sm_configuration_private_certificate_root_ca",
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//},
		},
	})
}

func testAccCheckIbmSmConfigurationPrivateCertificateRootCAConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_configuration_private_certificate_root_ca" "sm_configuration_private_certificate_root_ca" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			crl_expiry = "10000h"
			name = "root-ca-terraform-private-cert-datasource-test"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmConfigurationPrivateCertificateRootCAExists(n string, obj secretsmanagerv2.PrivateCertificateConfigurationRootCA) resource.TestCheckFunc {

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

		privateCertificateConfigurationRootCAIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		privateCertificateConfigurationRootCA := privateCertificateConfigurationRootCAIntf.(*secretsmanagerv2.PrivateCertificateConfigurationRootCA)
		obj = *privateCertificateConfigurationRootCA
		return nil
	}
}

func testAccCheckIbmSmConfigurationPrivateCertificateRootCADestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_configuration_private_certificate_root_ca" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		getConfigurationOptions.SetName(rs.Primary.ID)

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
