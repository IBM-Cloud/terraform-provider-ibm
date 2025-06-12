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

//SecretsManagerIamCredentialsConfigurationApiKey = os.Getenv("SECRETS_MANAGER_IAM_CREDENTIALS_CONFIGURATION_API_KEY")

func TestAccIbmSmIamCredentialsConfigurationBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmIamCredentialsConfigurationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmIamCredentialsConfigurationConfigBasic(acc.SecretsManagerIamCredentialsConfigurationApiKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmIamCredentialsConfigurationExists("ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration", acc.SecretsManagerIamCredentialsConfigurationApiKey),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmIamCredentialsConfigurationConfigBasic(apiKey string) string {
	return fmt.Sprintf(`

		resource "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource-iam-configuration"
			api_key = "%s"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, apiKey)
}

func testAccCheckIbmSmIamCredentialsConfigurationExists(n string, expectedApiKey string) resource.TestCheckFunc {

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

		iAMCredentialsConfigurationIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		iAMCredentialsConfiguration := iAMCredentialsConfigurationIntf.(*secretsmanagerv2.IAMCredentialsConfiguration)
		if *iAMCredentialsConfiguration.ApiKey != expectedApiKey {
			return fmt.Errorf("Wrong API Key")
		}
		return nil
	}
}

func testAccCheckIbmSmIamCredentialsConfigurationDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_iam_credentials_configuration" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		// Try to find the key
		_, response, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)

		if err == nil {
			return fmt.Errorf("IAMCredentialsConfiguration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for IAMCredentialsConfiguration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
