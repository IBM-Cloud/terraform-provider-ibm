// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

var customCredentialsSecretName = "terraform-test-custom-secret"
var modifiedCustomCredentialsSecretName = "modified-terraform-test-custom-secret"
var customCredentialsTtl = "259200"          // 3 days in seconds
var modifiedCustomCredentialsTtl = "7776000" // 3 months in seconds
var customCredentialsIntParam = 8
var modifiedCustomCredentialsIntParam = 67
var customCredentialsStrParam = "first"
var modifiedCustomCredentialsStrParam = "second"
var customCredentialsBoolParam = true
var modifiedCustomCredentialsBoolParam = false

func TestAccIbmSmCustomCredentialsSecretBasic(t *testing.T) {
	resourceName := "ibm_sm_custom_credentials_secret.sm_custom_credentials_secret_basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmCustomCredentialsSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: customCredentialsSecretConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at"},
			},
		},
	})
}

func TestAccIbmSmCustomCredentialsSecretAllArgs(t *testing.T) {
	resourceName := "ibm_sm_custom_credentials_secret.sm_custom_credentials_secret"
	// The expected credential values for the e2e code engine job we're using:
	expectedStrCredential := customCredentialsStrParam + "_output"
	expectedBoolCredential := !customCredentialsBoolParam
	expectedIntCredential := customCredentialsIntParam + 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmCustomCredentialsSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: customCredentialsSecretConfigAllArgs(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmCustomCredentialsSecretCreated(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "next_rotation_date"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "ttl", customCredentialsTtl),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.string_values.str_param", customCredentialsStrParam),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.integer_values.int_param", strconv.Itoa(customCredentialsIntParam)),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.boolean_values.bool_param", strconv.FormatBool(customCredentialsBoolParam)),
					resource.TestCheckResourceAttr(resourceName, "credentials_content.0.string_values.str_credential", expectedStrCredential),
					resource.TestCheckResourceAttr(resourceName, "credentials_content.0.integer_values.int_credential", strconv.Itoa(expectedIntCredential)),
					resource.TestCheckResourceAttr(resourceName, "credentials_content.0.boolean_values.bool_credential", strconv.FormatBool(expectedBoolCredential)),
				),
			},
			resource.TestStep{
				Config: customCredentialsSecretConfigUpdated(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmCustomCredentialsSecretUpdated(resourceName),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at"},
			},
		},
	})
}

var customCredentialsSecretBasicConfigFormat = `
		resource "ibm_sm_custom_credentials_secret" "sm_custom_credentials_secret_basic" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
   			ttl = "%s"
			configuration = ibm_sm_custom_credentials_configuration.test_config.name
			parameters {
				integer_values = {
					int_param = 6
				}
				string_values = {
					str_param = "prefix"
				}
				boolean_values = {
					bool_param = "true"
				}
			}
		}
`

var customCredentialsSecretFullConfigFormat = `
		resource "ibm_sm_custom_credentials_secret" "sm_custom_credentials_secret" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
  			custom_metadata = %s
   			ttl = "%s"
			rotation %s
			configuration = ibm_sm_custom_credentials_configuration.test_config.name
			parameters {
				integer_values = {
					int_param = %d
				}
				string_values = {
					str_param = "%s"
				}
				boolean_values = {
					bool_param = %t
				}
			}
		}
`

func customCredentialsEngineConfig() string {
	return iamCredentialSecretConfigForCustomCredentials() + fmt.Sprintf(`
		resource "ibm_sm_custom_credentials_configuration" "test_config" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource-custom-configuration"
			api_key_ref = ibm_sm_iam_credentials_secret.iam_credentials_for_custom_credentials.secret_id
			task_timeout  = "1h"
            code_engine {
				job_name   = "%s"
				project_id = "%s"
				region     = "%s"
			}
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerCodeEngineJobName, acc.SecretsManagerCodeEngineProjectId, acc.SecretsManagerCodeEngineRegion)
}

func customCredentialsSecretConfigBasic() string {
	return customCredentialsEngineConfig() +
		fmt.Sprintf(customCredentialsSecretBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			customCredentialsSecretName, customCredentialsTtl)
}

func customCredentialsSecretConfigAllArgs() string {
	return customCredentialsEngineConfig() +
		fmt.Sprintf(customCredentialsSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			customCredentialsSecretName, description, label, customMetadata, customCredentialsTtl, rotationPolicy, customCredentialsIntParam, customCredentialsStrParam, customCredentialsBoolParam)
}

func customCredentialsSecretConfigUpdated() string {
	return customCredentialsEngineConfig() +
		fmt.Sprintf(customCredentialsSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			modifiedCustomCredentialsSecretName, modifiedDescription, modifiedLabel,
			modifiedCustomMetadata, modifiedCustomCredentialsTtl, modifiedRotationPolicy, modifiedCustomCredentialsIntParam, modifiedCustomCredentialsStrParam, modifiedCustomCredentialsBoolParam)
}

func testAccCheckIbmSmCustomCredentialsSecretCreated(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		customCredentialsSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := customCredentialsSecretIntf.(*secretsmanagerv2.CustomCredentialsSecret)

		if err := verifyAttr(*secret.Name, customCredentialsSecretName, "secret name"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Description, description, "secret description"); err != nil {
			return err
		}
		if len(secret.Labels) != 1 {
			return fmt.Errorf("Wrong number of labels: %d", len(secret.Labels))
		}
		if err := verifyAttr(secret.Labels[0], label, "label"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, customMetadata, "custom metadata"); err != nil {
			return err
		}
		if err := verifyAttr(getAutoRotate(secret.Rotation), "true", "auto_rotate"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationUnit(secret.Rotation), "day", "rotation unit"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationInterval(secret.Rotation), "1", "rotation interval"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.TTL, customCredentialsTtl, "TTL"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmCustomCredentialsSecretUpdated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		customCredentialsSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := customCredentialsSecretIntf.(*secretsmanagerv2.CustomCredentialsSecret)

		if err := verifyAttr(*secret.Name, modifiedCustomCredentialsSecretName, "secret name"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Description, modifiedDescription, "secret description after update"); err != nil {
			return err
		}
		if len(secret.Labels) != 1 {
			return fmt.Errorf("Wrong number of labels after update: %d", len(secret.Labels))
		}
		if err := verifyAttr(secret.Labels[0], modifiedLabel, "label after update"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, modifiedCustomMetadata, "custom metadata after update"); err != nil {
			return err
		}
		if err := verifyAttr(getAutoRotate(secret.Rotation), "true", "auto_rotate after update"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationUnit(secret.Rotation), "month", "rotation unit after update"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationInterval(secret.Rotation), "2", "rotation interval after update"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.TTL, modifiedCustomCredentialsTtl, "TTL after update"); err != nil {
			return err
		}
		if len(secret.Parameters) != 3 {
			return fmt.Errorf("Wrong number of labels after update: %d", len(secret.Parameters))
		}
		if err := verifyAttr(secret.Parameters["str_param"].(string), modifiedCustomCredentialsStrParam, "str_param after update"); err != nil {
			return err
		}
		if err := verifyIntAttr(int(secret.Parameters["int_param"].(float64)), modifiedCustomCredentialsIntParam, "int_param after update"); err != nil {
			return err
		}
		if err := verifyBoolAttr(secret.Parameters["bool_param"].(bool), modifiedCustomCredentialsBoolParam, "bool_param after update"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmCustomCredentialsSecretDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_custom_credentials_secret" {
			continue
		}

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("CustomCredentialsSecret still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for CustomCredentialsSecret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
