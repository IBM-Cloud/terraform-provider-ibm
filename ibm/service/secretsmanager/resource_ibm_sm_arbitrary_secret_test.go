// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

var arbitrarySecretName = "terraform-test-arbitrary-secret"
var modifiedArbitrarySecretName = "modified-terraform-test-arbitrary-secret"
var payload = "secret-credentials"
var modifiedPayload = "modified-credentials"

func TestAccIbmSmArbitrarySecretBasic(t *testing.T) {
	resourceName := "ibm_sm_arbitrary_secret.sm_arbitrary_secret_basic"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmArbitrarySecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: arbitrarySecretConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIbmSmArbitrarySecretAllArgs(t *testing.T) {
	resourceName := "ibm_sm_arbitrary_secret.sm_arbitrary_secret"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmArbitrarySecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: arbitrarySecretConfigAllArgs(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmArbitrarySecretCreated(resourceName),
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
				Config: testAccCheckIbmSmArbitrarySecretConfigUpdated(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmArbitrarySecretUpdated(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

var arbitrarySecretBasicConfigFormat = `
		resource "ibm_sm_arbitrary_secret" "sm_arbitrary_secret_basic" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			payload = "%s"
		}`

var arbitrarySecretFullConfigFormat = `
		resource "ibm_sm_arbitrary_secret" "sm_arbitrary_secret" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
  			payload = "%s"
  			expiration_date = "%s"
  			custom_metadata = %s
			secret_group_id = "default"
		}`

func arbitrarySecretConfigBasic() string {
	return fmt.Sprintf(arbitrarySecretBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		arbitrarySecretName, payload)
}

func arbitrarySecretConfigAllArgs() string {
	return fmt.Sprintf(arbitrarySecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		arbitrarySecretName, description, label, payload, expirationDate, customMetadata)
}

func testAccCheckIbmSmArbitrarySecretConfigUpdated() string {
	return fmt.Sprintf(arbitrarySecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		modifiedArbitrarySecretName, modifiedDescription, modifiedLabel, modifiedPayload, modifiedExpirationDate, modifiedCustomMetadata)
}

func testAccCheckIbmSmArbitrarySecretCreated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		arbitrarySecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := arbitrarySecretIntf.(*secretsmanagerv2.ArbitrarySecret)

		if err := verifyAttr(*secret.Name, arbitrarySecretName, "secret name"); err != nil {
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
		if err := verifyAttr(*secret.Payload, payload, "payload"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmArbitrarySecretUpdated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		arbitrarySecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := arbitrarySecretIntf.(*secretsmanagerv2.ArbitrarySecret)
		if err := verifyAttr(*secret.Name, modifiedArbitrarySecretName, "secret name after update"); err != nil {
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
		if err := verifyAttr(*secret.Payload, modifiedPayload, "payload after update"); err != nil {
			return err
		}
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
