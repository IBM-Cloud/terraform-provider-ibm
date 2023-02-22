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

func TestAccIbmSmPublicCertificateBasic(t *testing.T) {
	//var conf secretsmanagerv2.PublicCertificate
	//
	//resource.Test(t, resource.TestCase{
	//	PreCheck:     func() { acc.TestAccPreCheck(t) },
	//	Providers:    acc.TestAccProviders,
	//	CheckDestroy: testAccCheckIbmSmPublicCertificateDestroy,
	//	Steps: []resource.TestStep{
	//		resource.TestStep{
	//			Config: testAccCheckIbmSmPublicCertificateConfigBasic(),
	//			Check: resource.ComposeAggregateTestCheckFunc(
	//				testAccCheckIbmSmPublicCertificateExists("ibm_sm_public_certificate.sm_public_certificate", conf),
	//			),
	//		},
	//		resource.TestStep{
	//			ResourceName:      "ibm_sm_public_certificate.sm_public_certificate",
	//			ImportState:       true,
	//			ImportStateVerify: true,
	//		},
	//	},
	//})
}

func testAccCheckIbmSmPublicCertificateConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_public_certificate" "sm_public_certificate_instance" {
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

func testAccCheckIbmSmPublicCertificateExists(n string, obj secretsmanagerv2.PublicCertificate) resource.TestCheckFunc {

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

		publicCertificateIntf, _, err := secretsManagerClient.GetSecret(getSecretOptions)
		if err != nil {
			return err
		}

		publicCertificate := publicCertificateIntf.(*secretsmanagerv2.PublicCertificate)
		obj = *publicCertificate
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

		getSecretOptions.SetID(rs.Primary.ID)

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
