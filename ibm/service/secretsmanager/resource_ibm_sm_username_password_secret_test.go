// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

var username = "user"
var modifiedUsername = "modified_user"
var password = "password"
var modifiedPassword = "modified_password"
var usernamePasswordSecretName = "terraform-test-username-secret"
var modifiedUsernamePasswordSecretName = "modified-terraform-test-username-secret"
var passwordGenerationPolicy = `{
				length = 17
				include_digits = false
				include_symbols = false
				include_uppercase = true
			}`
var passwordGenerationPolicyJSON = "{\"length\":17,\"include_digits\":false,\"include_symbols\":false,\"include_uppercase\":true}"

var modifiedPasswordGenerationPolicy = `{
				length = 26
				include_digits = true
				include_symbols = true
				include_uppercase = false
			}`
var modifiedPasswordGenerationPolicyJSON = "{\"length\":26,\"include_digits\":true,\"include_symbols\":true,\"include_uppercase\":false}"

func TestAccIbmSmUsernamePasswordSecretBasic(t *testing.T) {
	resourceName := "ibm_sm_username_password_secret.sm_username_password_secret_basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmUsernamePasswordSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: usernamePasswordSecretConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "retrieved_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "retrieved_at"},
			},
		},
	})
}

func TestAccIbmSmUsernamePasswordSecretAllArgs(t *testing.T) {
	resourceName := "ibm_sm_username_password_secret.sm_username_password_secret"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmUsernamePasswordSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: usernamePasswordSecretConfigAllArgs(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmUsernamePasswordSecretCreated(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "next_rotation_date"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
				),
			},
			resource.TestStep{
				Config: usernamePasswordSecretConfigUpdated(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmUsernamePasswordSecretUpdated(resourceName),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "retrieved_at"},
			},
		},
	})
}

var usernamePasswordSecretBasicConfigFormat = `
		resource "ibm_sm_username_password_secret" "sm_username_password_secret_basic" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
			username = "%s"
			password = "%s"
		}`

var usernamePasswordSecretFullConfigFormat = `
		resource "ibm_sm_username_password_secret" "sm_username_password_secret" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
			username = "%s"
			password = "%s"
  			expiration_date = "%s"
  			custom_metadata = %s
			secret_group_id = "default"
			rotation %s
			password_generation_policy %s
		}`

func usernamePasswordSecretConfigBasic() string {
	return fmt.Sprintf(usernamePasswordSecretBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		usernamePasswordSecretName, username, password)
}

func usernamePasswordSecretConfigAllArgs() string {
	return fmt.Sprintf(usernamePasswordSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		usernamePasswordSecretName, description, label, username, password, expirationDate, customMetadata, rotationPolicy,
		passwordGenerationPolicy)
}

func usernamePasswordSecretConfigUpdated() string {
	return fmt.Sprintf(usernamePasswordSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		modifiedUsernamePasswordSecretName, modifiedDescription, modifiedLabel, modifiedUsername, modifiedPassword,
		modifiedExpirationDate, modifiedCustomMetadata, modifiedRotationPolicy, modifiedPasswordGenerationPolicy)
}

func testAccCheckIbmSmUsernamePasswordSecretCreated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		usernamePasswordSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := usernamePasswordSecretIntf.(*secretsmanagerv2.UsernamePasswordSecret)

		if err := verifyAttr(*secret.Name, usernamePasswordSecretName, "secret name"); err != nil {
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
		if err := verifyDateAttr(secret.ExpirationDate, expirationDate, "expiration date"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, customMetadata, "custom metadata"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Username, username, "username"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Password, password, "password"); err != nil {
			return err
		}
		if err := verifyAttr(getAutoRotate(secret.Rotation), "true", "auto_rotate"); err != nil {
			return err
		}
		pwdPolicyJson, _ := json.Marshal(secret.PasswordGenerationPolicy)
		if err := verifyAttr(string(pwdPolicyJson), passwordGenerationPolicyJSON, "password_generation_policy"); err != nil {
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

func testAccCheckIbmSmUsernamePasswordSecretUpdated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		usernamePasswordSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := usernamePasswordSecretIntf.(*secretsmanagerv2.UsernamePasswordSecret)

		if err := verifyAttr(*secret.Name, modifiedUsernamePasswordSecretName, "secret name"); err != nil {
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
		if err := verifyDateAttr(secret.ExpirationDate, modifiedExpirationDate, "expiration date after update"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, modifiedCustomMetadata, "custom metadata after update"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Username, modifiedUsername, "username after update"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Password, modifiedPassword, "password after update"); err != nil {
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
		pwdPolicyJson, _ := json.Marshal(secret.PasswordGenerationPolicy)
		if err := verifyAttr(string(pwdPolicyJson), modifiedPasswordGenerationPolicyJSON, "password_generation_policy"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmUsernamePasswordSecretDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_username_password_secret" {
			continue
		}

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("UsernamePasswordSecret still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for UsernamePasswordSecret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
