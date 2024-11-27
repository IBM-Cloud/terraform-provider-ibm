// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmIamCredentialsConfigurationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmIamCredentialsConfigurationDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration", "config_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIbmSmIamCredentialsConfigurationDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource-iam-configuration"
			api_key = "%s"
		}

		data "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration_instance.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsConfigurationApiKey, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
