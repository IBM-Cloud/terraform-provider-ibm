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

func TestAccIbmSmIamCredentialsSecretBasic(t *testing.T) {
	var conf secretsmanagerv2.IAMCredentialsSecret

	if acc.SecretsManagerIamCredentialsConfigurationApiKey != "" {
		if acc.SecretsManagerIamCredentialsSecretServiceId != "" {
			resource.Test(t, resource.TestCase{
				PreCheck:     func() { acc.TestAccPreCheck(t) },
				Providers:    acc.TestAccProviders,
				CheckDestroy: testAccCheckIbmSmIamCredentialsSecretDestroy,
				Steps: []resource.TestStep{
					resource.TestStep{
						Config: testAccCheckIbmSmIamCredentialsSecretConfigBasic_withApikey_serviceid(),
						Check: resource.ComposeAggregateTestCheckFunc(
							testAccCheckIbmSmIamCredentialsSecretExists("ibm_sm_iam_credentials_secret.sm_iam_credentials_secret", conf),
						),
					},
					resource.TestStep{
						ResourceName:      "ibm_sm_iam_credentials_secret.sm_iam_credentials_secret",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		} else {
			resource.Test(t, resource.TestCase{
				PreCheck:     func() { acc.TestAccPreCheck(t) },
				Providers:    acc.TestAccProviders,
				CheckDestroy: testAccCheckIbmSmIamCredentialsSecretDestroy,
				Steps: []resource.TestStep{
					resource.TestStep{
						Config: testAccCheckIbmSmIamCredentialsSecretConfigBasic_withApikey_accessGroup(),
						Check: resource.ComposeAggregateTestCheckFunc(
							testAccCheckIbmSmIamCredentialsSecretExists("ibm_sm_iam_credentials_secret.sm_iam_credentials_secret", conf),
						),
					},
					resource.TestStep{
						ResourceName:      "ibm_sm_iam_credentials_secret.sm_iam_credentials_secret",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}
	} else {
		if acc.SecretsManagerIamCredentialsSecretServiceId != "" {
			resource.Test(t, resource.TestCase{
				PreCheck:     func() { acc.TestAccPreCheck(t) },
				Providers:    acc.TestAccProviders,
				CheckDestroy: testAccCheckIbmSmIamCredentialsSecretDestroy,
				Steps: []resource.TestStep{
					resource.TestStep{
						Config: testAccCheckIbmSmIamCredentialsSecretConfigBasic_noApikey_serviceid(),
						Check: resource.ComposeAggregateTestCheckFunc(
							testAccCheckIbmSmIamCredentialsSecretExists("ibm_sm_iam_credentials_secret.sm_iam_credentials_secret", conf),
						),
					},
					resource.TestStep{
						ResourceName:      "ibm_sm_iam_credentials_secret.sm_iam_credentials_secret",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		} else {
			resource.Test(t, resource.TestCase{
				PreCheck:     func() { acc.TestAccPreCheck(t) },
				Providers:    acc.TestAccProviders,
				CheckDestroy: testAccCheckIbmSmIamCredentialsSecretDestroy,
				Steps: []resource.TestStep{
					resource.TestStep{
						Config: testAccCheckIbmSmIamCredentialsSecretConfigBasic_noApikey_accessGroup(),
						Check: resource.ComposeAggregateTestCheckFunc(
							testAccCheckIbmSmIamCredentialsSecretExists("ibm_sm_iam_credentials_secret.sm_iam_credentials_secret", conf),
						),
					},
					resource.TestStep{
						ResourceName:      "ibm_sm_iam_credentials_secret.sm_iam_credentials_secret",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}
	}
}

func testAccCheckIbmSmIamCredentialsSecretConfigBasic_withApikey_serviceid() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource"
			api_key = "%s"
		}

		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret" {
			instance_id   = "%s"
			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			service_id = "%s"
  			ttl = "1800"
			name = "terraform-test"
			reuse_api_key = true
			depends_on = [
				ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration_instance
			]
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsConfigurationApiKey, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsSecretServiceId)
}

func testAccCheckIbmSmIamCredentialsSecretConfigBasic_withApikey_accessGroup() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource"
			api_key = "%s"
		}

		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret" {
			instance_id   = "%s"
			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			access_groups = ["%s"]
  			ttl = "1800"
			name = "iam-credentials-test-terraform"
			reuse_api_key = true
			depends_on = [
				ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration_instance
			]
		}

	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsConfigurationApiKey, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsSecretServiceAccessGroup)
}

func testAccCheckIbmSmIamCredentialsSecretConfigBasic_noApikey_serviceid() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret" {
			instance_id   = "%s"
			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			service_id = "%s"
  			ttl = "1800"
			name = "iam-credentials-test-terraform"
			reuse_api_key = true
		}

	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsSecretServiceId)
}

func testAccCheckIbmSmIamCredentialsSecretConfigBasic_noApikey_accessGroup() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret" {
			instance_id   = "%s"
			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			access_groups = ["%s"]
  			ttl = "1800"
			name = "iam-credentials-test-terraform"
			reuse_api_key = true
		}

	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsSecretServiceAccessGroup)
}

func testAccCheckIbmSmIamCredentialsSecretExists(n string, obj secretsmanagerv2.IAMCredentialsSecret) resource.TestCheckFunc {

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

		iAMCredentialsSecretIntf, _, err := secretsManagerClient.GetSecret(getSecretOptions)
		if err != nil {
			return err
		}

		iAMCredentialsSecret := iAMCredentialsSecretIntf.(*secretsmanagerv2.IAMCredentialsSecret)
		obj = *iAMCredentialsSecret
		return nil
	}
}

func testAccCheckIbmSmIamCredentialsSecretDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_iam_credentials_secret" {
			continue
		}

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("IAMCredentialsSecret still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for IAMCredentialsSecret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
