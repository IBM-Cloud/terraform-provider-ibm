// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSecretsManagerSecretDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSecretsManagerSecretDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_secrets_manager_secret.secrets_manager_secret", "secret_type", secretsManagerSecretType),
					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secret.secrets_manager_secret", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secret.secrets_manager_secret", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secret.secrets_manager_secret", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secret.secrets_manager_secret", "metadata.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSecretsManagerSecretDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_secrets_manager_secret" "secrets_manager_secret" {
			instance_id = "%s"
			secret_type = "%s"
			secret_id = "%s"
		}
	`, secretsManagerInstanceID, secretsManagerSecretType, secretsManagerSecretID)
}
