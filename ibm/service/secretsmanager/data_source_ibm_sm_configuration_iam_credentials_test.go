// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmConfigurationIamCredentialsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmConfigurationIamCredentialsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_iam_credentials.sm_configuration_iam_credentials", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_iam_credentials.sm_configuration_iam_credentials", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_iam_credentials.sm_configuration_iam_credentials", "config_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_iam_credentials.sm_configuration_iam_credentials", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_iam_credentials.sm_configuration_iam_credentials", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_iam_credentials.sm_configuration_iam_credentials", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_configuration_iam_credentials.sm_configuration_iam_credentials", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIbmSmConfigurationIamCredentialsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_configuration_iam_credentials" "sm_configuration_iam_credentials_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource-iam-configuration"
			api_key = "%s"
		}

		data "ibm_sm_configuration_iam_credentials" "sm_configuration_iam_credentials" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_configuration_iam_credentials.sm_configuration_iam_credentials_instance.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsConfigurationApiKey, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
