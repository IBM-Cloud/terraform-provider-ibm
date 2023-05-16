// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

func TestAccIbmSmArbitrarySecretBasic(t *testing.T) {
	var conf secretsmanagerv2.ArbitrarySecret

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmArbitrarySecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmArbitrarySecretConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmArbitrarySecretExists("ibm_sm_arbitrary_secret.sm_arbitrary_secret", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_arbitrary_secret.sm_arbitrary_secret",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmArbitrarySecretConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_arbitrary_secret" "sm_arbitrary_secret" {
			name = "terraform-test-arbitrary-secret-resource"
			instance_id   = "%s"
  			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			payload = "secret-credentials"
  			secret_group_id = "default"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmArbitrarySecretExists(n string, obj secretsmanagerv2.ArbitrarySecret) resource.TestCheckFunc {

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

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		arbitrarySecretIntf, _, err := secretsManagerClient.GetSecret(getSecretOptions)
		if err != nil {
			return err
		}

		arbitrarySecret := arbitrarySecretIntf.(*secretsmanagerv2.ArbitrarySecret)
		obj = *arbitrarySecret
		return nil
	}
}

func testAccCheckIbmSmArbitrarySecretDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_arbitrary_secret" {
			continue
		}

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("ArbitrarySecret still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ArbitrarySecret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func getClientWithInstanceEndpointTest(originalClient *secretsmanagerv2.SecretsManagerV2) *secretsmanagerv2.SecretsManagerV2 {
	// build the api endpoint
	domain := "appdomain.cloud"
	if strings.Contains(os.Getenv("IBMCLOUD_IAM_API_ENDPOINT"), "test") {
		domain = "test.appdomain.cloud"
	}
	endpoint := fmt.Sprintf("https://%s.%s.secrets-manager.%s", acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, domain)
	newClient := &secretsmanagerv2.SecretsManagerV2{
		Service: originalClient.Service.Clone(),
	}
	newClient.Service.SetServiceURL(endpoint)

	return newClient
}
