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

func TestAccIbmSmPrivateCertificateBasic(t *testing.T) {
	//var conf secretsmanagerv2.PrivateCertificate
	//
	//resource.Test(t, resource.TestCase{
	//	PreCheck:     func() { acc.TestAccPreCheck(t) },
	//	Providers:    acc.TestAccProviders,
	//	CheckDestroy: testAccCheckIbmSmPrivateCertificateDestroy,
	//	Steps: []resource.TestStep{
	//		resource.TestStep{
	//			Config: testAccCheckIbmSmPrivateCertificateConfigBasic(),
	//			Check: resource.ComposeAggregateTestCheckFunc(
	//				testAccCheckIbmSmPrivateCertificateExists("ibm_sm_private_certificate.sm_private_certificate", conf),
	//			),
	//		},
	//		resource.TestStep{
	//			ResourceName:      "ibm_sm_private_certificate.sm_private_certificate",
	//			ImportState:       true,
	//			ImportStateVerify: true,
	//		},
	//	},
	//})
}

func testAccCheckIbmSmPrivateCertificateConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_private_certificate" "sm_private_certificate_instance" {
			secret_prototype {
				custom_metadata = {"key":"value"}
				description = "Extended description for this secret."
				expiration_date = "2022-04-12T23:20:50.520Z"
				labels = [ "my-label" ]
				name = "my-secret-example"
				secret_group_id = "default"
				secret_type = "arbitrary"
				payload = "secret-credentials"
				version_custom_metadata = {"key":"value"}
			}
		}
	`)
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

		getSecretOptions.SetID(rs.Primary.ID)

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

		getSecretOptions.SetID(rs.Primary.ID)

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
