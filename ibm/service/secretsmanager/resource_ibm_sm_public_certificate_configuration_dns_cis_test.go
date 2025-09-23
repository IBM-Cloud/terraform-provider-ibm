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

func TestAccIbmSmConfigurationPublicCertificateDnsCisBasic(t *testing.T) {
	var conf secretsmanagerv2.PublicCertificateConfigurationDNSCloudInternetServices

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmConfigurationPublicCertificateDnsCisDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmConfigurationPublicCertificateDnsCisConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmConfigurationPublicCertificateDnsCisExists("ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmConfigurationPublicCertificateDnsCisConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis" {
			  instance_id   = "%s"
       		  region        = "%s"
			  cloud_internet_services_crn = "%s"
  			  name = "cloud-internet-services-config-terraform-test"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerPublicCertificateCisCrn)
}

func testAccCheckIbmSmConfigurationPublicCertificateDnsCisExists(n string, obj secretsmanagerv2.PublicCertificateConfigurationDNSCloudInternetServices) resource.TestCheckFunc {

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

		publicCertificateConfigurationDNSCloudInternetServicesIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		publicCertificateConfigurationDNSCloudInternetServices := publicCertificateConfigurationDNSCloudInternetServicesIntf.(*secretsmanagerv2.PublicCertificateConfigurationDNSCloudInternetServices)
		obj = *publicCertificateConfigurationDNSCloudInternetServices
		return nil
	}
}

func testAccCheckIbmSmConfigurationPublicCertificateDnsCisDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_public_certificate_configuration_dns_cis" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		// Try to find the key
		_, response, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)

		if err == nil {
			return fmt.Errorf("PublicCertificateConfigurationDNSCloudInternetServices still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for PublicCertificateConfigurationDNSCloudInternetServices (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
