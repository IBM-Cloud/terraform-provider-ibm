// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmCustomCredentialsConfigurationDataSource(t *testing.T) {
	dataSourceName := "data.ibm_sm_custom_credentials_configuration.my_config"
	taskTimeout := "3h"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmCustomCredentialsConfigurationDataSourceConfig(taskTimeout),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "created_by"),
					resource.TestCheckResourceAttrSet(dataSourceName, "created_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "updated_at"),
					resource.TestCheckResourceAttr(dataSourceName, "task_timeout", taskTimeout),
					resource.TestCheckResourceAttr(dataSourceName, "code_engine.0.job_name", acc.SecretsManagerCodeEngineJobName),
				),
			},
		},
	})
}

func testAccCheckIbmSmCustomCredentialsConfigurationDataSourceConfig(taskTimeout string) string {
	return iamCredentialSecretConfigForCustomCredentials() + fmt.Sprintf(`
		resource "ibm_sm_custom_credentials_configuration" "my_config" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource-custom-configuration"
			api_key_ref = ibm_sm_iam_credentials_secret.iam_credentials_for_custom_credentials.secret_id
			task_timeout  = "%s"
            code_engine {
				job_name   = "%s"
				project_id = "%s"
				region     = "%s"
			}
		}

		data "ibm_sm_custom_credentials_configuration" "my_config" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_custom_credentials_configuration.my_config.name
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, taskTimeout,
		acc.SecretsManagerCodeEngineJobName, acc.SecretsManagerCodeEngineProjectId, acc.SecretsManagerCodeEngineRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func iamCredentialSecretConfigForCustomCredentials() string {
	return iamCredentialsEngineConfig() + fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_secret" "iam_credentials_for_custom_credentials" {
			instance_id   = "%s"
			region        = "%s"
			name = "iam-credentials-for-custom-credentials-terraform-tests"
            service_id = "%s"
  			reuse_api_key = true
  			ttl = "259200"
			rotation {
				auto_rotate = true
				interval = 1
				unit = "day"
			}
			depends_on = [
				ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration_instance
			]
		}`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerServiceIdForCustomCredentials)
}
