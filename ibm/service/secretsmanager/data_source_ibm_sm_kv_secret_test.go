// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmKvSecretDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmKvSecretDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "versions_total"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret", "data.%"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret_by_name", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret_by_name", "secret_group_name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_kv_secret.sm_kv_secret_by_name", "secret_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSmKvSecretDataSourceConfigBasic() string {
	return fmt.Sprintf(`
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

		data "ibm_sm_kv_secret" "sm_kv_secret" {
			instance_id = "%s"
			region = "%s"
			secret_id = ibm_sm_kv_secret.sm_kv_secret_instance.secret_id
		}

		data "ibm_sm_kv_secret" "sm_kv_secret_by_name" {
			instance_id   = "%s"
			region = "%s"
			name = ibm_sm_kv_secret.sm_kv_secret_instance.name
			secret_group_name = "default"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
