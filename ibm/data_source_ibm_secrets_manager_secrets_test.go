// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSecretsManagerSecretsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSecretsManagerSecretsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "secret_type", secretsManagerSecretType),
					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "metadata.#"),
					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "secrets.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSecretsManagerSecretsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_secrets_manager_secrets" "secrets_manager_secrets" {
			instance_id = "%s"
			secret_type = "%s"
		}

		output "WorkSpaceValues" {
			value = data.ibm_secrets_manager_secrets.secrets_manager_secrets.secret_type
		}
	`, secretsManagerInstanceID, secretsManagerSecretType)
}
