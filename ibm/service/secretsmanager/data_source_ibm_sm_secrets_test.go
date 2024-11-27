// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmSecretsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmSecretsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_secrets.sm_secrets", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secrets.sm_secrets", "secrets.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSmSecretsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_arbitrary_secret" "sm_arbitrary_secret_instance" {
			name = "test_arbitrary_secret"
			instance_id   = "%s"
  			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			payload = "secret-credentials"
  			secret_group_id = "default"
		}

		resource "ibm_sm_username_password_secret" "sm_username_password_secret_instance" {
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

		resource "ibm_sm_kv_secret" "sm_kv_secret_instance" {
			  instance_id   = "%s"
       		  region        = "%s"
  			  custom_metadata = {"key":"value"}
  			  data = {"key":"value"}
			  description = "Extended description for this secret."
  			  labels = ["my-label"]
  			  secret_group_id = "default"
			  name = "kv-secret-terraform-test"
		}

		data "ibm_sm_secrets" "sm_secrets" {
			instance_id = "%s"
			region = "%s"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
