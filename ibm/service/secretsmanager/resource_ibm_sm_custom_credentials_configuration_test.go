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

//SecretsManagerIamCredentialsConfigurationApiKey = os.Getenv("SECRETS_MANAGER_IAM_CREDENTIALS_CONFIGURATION_API_KEY")

func TestAccIbmSmCustomCredentialsConfiguration(t *testing.T) {
	resourceName := "ibm_sm_custom_credentials_configuration.my_config"
	taskTimeout := "3h"
	modifiedTaskTimeout := "14h"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmCustomCredentialsConfigurationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmCustomCredentialsConfigurationConfig(taskTimeout),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmCustomCredentialsConfigurationExists(resourceName, taskTimeout),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "schema.0.credentials.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "schema.0.credentials.1.name"),
					resource.TestCheckResourceAttrSet(resourceName, "schema.0.credentials.2.name"),
					resource.TestCheckResourceAttrSet(resourceName, "schema.0.credentials.0.format"),
					resource.TestCheckResourceAttrSet(resourceName, "schema.0.parameters.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "schema.0.parameters.1.name"),
					resource.TestCheckResourceAttrSet(resourceName, "schema.0.parameters.2.name"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSmCustomCredentialsConfigurationConfig(modifiedTaskTimeout),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmCustomCredentialsConfigurationExists(resourceName, modifiedTaskTimeout),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_custom_credentials_configuration.my_config",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmCustomCredentialsConfigurationConfig(taskTimeout string) string {
	return iamCredentialSecretConfigForCustomCredentials() + fmt.Sprintf(`
		resource "ibm_sm_custom_credentials_configuration" "my_config" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource-custom-configuration"
			api_key_ref = ibm_sm_iam_credentials_secret.iam_credentials_for_custom_credentials.secret_id
			task_timeout  = "%s"
            code_engine {
				job_name   = "%s"
				project_id = "%s"
				region     = "%s"
			}
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, taskTimeout,
		acc.SecretsManagerCodeEngineJobName, acc.SecretsManagerCodeEngineProjectId, acc.SecretsManagerCodeEngineRegion)
}

func testAccCheckIbmSmCustomCredentialsConfigurationExists(n string, expectedTaskTimeout string) resource.TestCheckFunc {

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

		customCredentialsConfigurationIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		customCredentialsConfiguration := customCredentialsConfigurationIntf.(*secretsmanagerv2.CustomCredentialsConfiguration)
		if *customCredentialsConfiguration.TaskTimeout != expectedTaskTimeout {
			return fmt.Errorf("Wrong task timeout")
		}
		if *customCredentialsConfiguration.CodeEngine.ProjectID != acc.SecretsManagerCodeEngineProjectId {
			return fmt.Errorf("Wrong code engine project ID")
		}
		if *customCredentialsConfiguration.CodeEngine.JobName != acc.SecretsManagerCodeEngineJobName {
			return fmt.Errorf("Wrong code engine job name")
		}
		if *customCredentialsConfiguration.CodeEngine.Region != acc.SecretsManagerCodeEngineRegion {
			return fmt.Errorf("Wrong code engine region")
		}
		return nil
	}
}

func testAccCheckIbmSmCustomCredentialsConfigurationDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_custom_credentials_configuration" {
			continue
		}

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		// Try to find the key
		_, response, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)

		if err == nil {
			return fmt.Errorf("CustomCredentialsConfiguration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for CustomCredentialsConfiguration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
