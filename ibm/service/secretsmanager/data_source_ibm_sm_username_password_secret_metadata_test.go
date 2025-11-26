// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmUsernamePasswordSecretMetadataDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmUsernamePasswordSecretMetadataDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "retrieved_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "versions_total"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "rotation.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_username_password_secret_metadata.sm_username_password_secret_metadata", "password_generation_policy.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSmUsernamePasswordSecretMetadataDataSourceConfigBasic() string {
	return fmt.Sprintf(`
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

		data "ibm_sm_username_password_secret_metadata" "sm_username_password_secret_metadata" {
			instance_id   = "%s"
			region        = "%s"
			secret_id = ibm_sm_username_password_secret.sm_username_password_secret_instance.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
