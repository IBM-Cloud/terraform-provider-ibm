// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmArbitrarySecretDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmArbitrarySecretDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_arbitrary_secret.sm_arbitrary_secret", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_arbitrary_secret.sm_arbitrary_secret", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_arbitrary_secret.sm_arbitrary_secret", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_arbitrary_secret.sm_arbitrary_secret", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_arbitrary_secret.sm_arbitrary_secret", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_arbitrary_secret.sm_arbitrary_secret", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_arbitrary_secret.sm_arbitrary_secret", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_arbitrary_secret.sm_arbitrary_secret", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_arbitrary_secret.sm_arbitrary_secret", "versions_total"),
				),
			},
		},
	})
}

func testAccCheckIbmSmArbitrarySecretDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_arbitrary_secret" "sm_arbitrary_secret_instance" {
			name = "test_arbitrary_secret_terraform"
			instance_id   = "%s"
  			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			payload = "secret-credentials"
  			secret_group_id = "default"
		}

		data "ibm_sm_arbitrary_secret" "sm_arbitrary_secret" {
			instance_id   = "%s"
			region = "%s"
			secret_id = ibm_sm_arbitrary_secret.sm_arbitrary_secret_instance.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
