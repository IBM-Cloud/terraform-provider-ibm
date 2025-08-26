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

var iamCredentialsSecretName = "terraform-test-iam-secret"
var modifiedIamCredentialsSecretName = "modified-terraform-test-iam-secret"
var iamCredentialsTtl = "259200"          // 3 days in seconds
var modifiedIamCredentialsTtl = "7776000" // 3 months in seconds

func TestAccIbmSmIamCredentialsSecretBasic(t *testing.T) {
	resourceName := "ibm_sm_iam_credentials_secret.sm_iam_credentials_secret_basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmIamCredentialsSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: iamCredentialsSecretConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "api_key"),
					resource.TestCheckResourceAttrSet(resourceName, "api_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "service_id_is_static"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "2"),
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

func TestAccIbmSmIamCredentialsSecretAllArgs(t *testing.T) {
	resourceName := "ibm_sm_iam_credentials_secret.sm_iam_credentials_secret"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmIamCredentialsSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: iamCredentialsSecretConfigAllArgs(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmIamCredentialsSecretCreated(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "api_key"),
					resource.TestCheckResourceAttrSet(resourceName, "api_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "service_id_is_static"),
					resource.TestCheckResourceAttrSet(resourceName, "next_rotation_date"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "2"),
				),
			},
			resource.TestStep{
				Config: iamCredentialsSecretConfigUpdated(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmIamCredentialsSecretUpdated(resourceName),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "2"),
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

var iamCredentialsSecretBasicConfigFormat = `
		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret_basic" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
   			ttl = "%s"
			%s
			depends_on = [
				ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration_instance
			]
		}`

var iamCredentialsSecretFullConfigFormat = `
		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
  			custom_metadata = %s
   			ttl = "%s"
			reuse_api_key = true
			rotation %s
			%s
			depends_on = [
				ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration_instance
			]
		}`

func iamCredentialsEngineConfig() string {
	if acc.SecretsManagerIamCredentialsConfigurationApiKey == "" {
		return ""
	} else {
		return fmt.Sprintf(`resource "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource"
			api_key = "%s"
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			acc.SecretsManagerIamCredentialsConfigurationApiKey)
	}
}

// Return either service_id field or access_groups field (name value pair)
func iamCredentialAccessField() string {
	if acc.SecretsManagerIamCredentialsSecretServiceId != "" {
		return fmt.Sprintf(`service_id = "%s"`, acc.SecretsManagerIamCredentialsSecretServiceId)
	} else {
		return fmt.Sprintf(`access_groups = ["%s"]`, acc.SecretsManagerIamCredentialsSecretServiceAccessGroup)
	}

}

func iamCredentialsSecretConfigBasic() string {
	return iamCredentialsEngineConfig() +
		fmt.Sprintf(iamCredentialsSecretBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			iamCredentialsSecretName, iamCredentialsTtl, iamCredentialAccessField())
}

func iamCredentialsSecretConfigAllArgs() string {
	return iamCredentialsEngineConfig() +
		fmt.Sprintf(iamCredentialsSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			iamCredentialsSecretName, description, label, customMetadata, iamCredentialsTtl, rotationPolicy, iamCredentialAccessField())
}

func iamCredentialsSecretConfigUpdated() string {
	return iamCredentialsEngineConfig() +
		fmt.Sprintf(iamCredentialsSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
			modifiedIamCredentialsSecretName, modifiedDescription, modifiedLabel,
			modifiedCustomMetadata, modifiedIamCredentialsTtl, modifiedRotationPolicy, iamCredentialAccessField())
}

func testAccCheckIbmSmIamCredentialsSecretCreated(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		iAMCredentialsSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := iAMCredentialsSecretIntf.(*secretsmanagerv2.IAMCredentialsSecret)

		if err := verifyAttr(*secret.Name, iamCredentialsSecretName, "secret name"); err != nil {
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
		return nil
	}
}

func testAccCheckIbmSmIamCredentialsSecretUpdated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		iamCredentialsSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := iamCredentialsSecretIntf.(*secretsmanagerv2.IAMCredentialsSecret)

		if err := verifyAttr(*secret.Name, modifiedIamCredentialsSecretName, "secret name"); err != nil {
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
