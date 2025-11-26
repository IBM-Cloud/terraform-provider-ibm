// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

var kvSecretName = "terraform-test-kv-secret"
var modifiedKvSecretName = "modified-terraform-test-kv-secret"
var kvSecretData = `{"secret_key":"secret_value"}`
var modifiedKvSecretData = `{"modified_key":"modified_value"}`

func TestAccIbmSmKvSecretBasic(t *testing.T) {
	resourceName := "ibm_sm_kv_secret.sm_kv_secret_basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmKvSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: kvSecretConfigBasic(),
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

func TestAccIbmSmKvSecretAllArgs(t *testing.T) {
	resourceName := "ibm_sm_kv_secret.sm_kv_secret"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmKvSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: kvSecretConfigAllArgs(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmKvSecretCreated(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
				),
			},
			{
				Config: kvSecretConfigUpdated(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmKvSecretUpdated(resourceName),
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

var kvSecretBasicConfigFormat = `
		resource "ibm_sm_kv_secret" "sm_kv_secret_basic" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			data = %s
		}`

var kvSecretFullConfigFormat = `
		resource "ibm_sm_kv_secret" "sm_kv_secret" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
  			data = %s
			custom_metadata = %s
			secret_group_id = "default"
		}`

func kvSecretConfigBasic() string {
	return fmt.Sprintf(kvSecretBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		kvSecretName, kvSecretData)
}

func kvSecretConfigAllArgs() string {
	return fmt.Sprintf(kvSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		kvSecretName, description, label, kvSecretData, customMetadata)
}

func kvSecretConfigUpdated() string {
	return fmt.Sprintf(kvSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		modifiedKvSecretName, modifiedDescription, modifiedLabel, modifiedKvSecretData, modifiedCustomMetadata)
}

func testAccCheckIbmSmKvSecretCreated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		kvSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := kvSecretIntf.(*secretsmanagerv2.KVSecret)

		if err := verifyAttr(*secret.Name, kvSecretName, "secret name"); err != nil {
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
		if err := verifyJsonAttr(secret.Data, kvSecretData, "secret data"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmKvSecretUpdated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		kvSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := kvSecretIntf.(*secretsmanagerv2.KVSecret)
		if err := verifyAttr(*secret.Name, modifiedKvSecretName, "secret name after update"); err != nil {
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
		if err := verifyJsonAttr(secret.Data, modifiedKvSecretData, "payload after update"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmKvSecretDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_kv_secret" {
			continue
		}

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("KVSecret still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for KVSecret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
