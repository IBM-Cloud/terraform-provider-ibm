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

func TestAccIbmSmUsernamePasswordSecretBasic(t *testing.T) {
	var conf secretsmanagerv2.UsernamePasswordSecret

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmUsernamePasswordSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmUsernamePasswordSecretConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmUsernamePasswordSecretExists("ibm_sm_username_password_secret.sm_username_password_secret", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_username_password_secret.sm_username_password_secret",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmUsernamePasswordSecretConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_username_password_secret" "sm_username_password_secret" {
			  instance_id   = "%s"
              region        = "%s"
              custom_metadata = {"key":"value"}
              description = "Extended description for this secret."
              labels = ["my-label"]
              rotation {
                auto_rotate = true
                interval = 1
                unit = "day"
              }
              secret_group_id = "default"
              username = "username"
    		  password = "password"
			  name = "username_password-datasource-terraform-test"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmUsernamePasswordSecretExists(n string, obj secretsmanagerv2.UsernamePasswordSecret) resource.TestCheckFunc {

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

		usernamePasswordSecretIntf, _, err := secretsManagerClient.GetSecret(getSecretOptions)
		if err != nil {
			return err
		}

		usernamePasswordSecret := usernamePasswordSecretIntf.(*secretsmanagerv2.UsernamePasswordSecret)
		obj = *usernamePasswordSecret
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
