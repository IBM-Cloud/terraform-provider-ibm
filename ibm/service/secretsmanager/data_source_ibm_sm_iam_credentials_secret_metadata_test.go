// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmIamCredentialsSecretMetadataDataSourceBasic(t *testing.T) {
	if acc.SecretsManagerIamCredentialsConfigurationApiKey != "" {
		if acc.SecretsManagerIamCredentialsSecretServiceId != "" {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { acc.TestAccPreCheck(t) },
				Providers: acc.TestAccProviders,
				Steps: []resource.TestStep{
					resource.TestStep{
						Config: testAccCheckIbmSmIamCredentialsSecretMetadataDataSourceConfigBasic_withApikey_serviceid(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "instance_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "created_by"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "created_at"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "crn"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_group_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_type"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "updated_at"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "versions_total"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "ttl"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "reuse_api_key"),
						),
					},
				},
			})
		} else {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { acc.TestAccPreCheck(t) },
				Providers: acc.TestAccProviders,
				Steps: []resource.TestStep{
					resource.TestStep{
						Config: testAccCheckIbmSmIamCredentialsSecretMetadataDataSourceConfigBasic_withApikey_accessGroup(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "instance_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "created_by"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "created_at"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "crn"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_group_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_type"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "updated_at"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "versions_total"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "ttl"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "reuse_api_key"),
						),
					},
				},
			})
		}
	} else {
		if acc.SecretsManagerIamCredentialsSecretServiceId != "" {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { acc.TestAccPreCheck(t) },
				Providers: acc.TestAccProviders,
				Steps: []resource.TestStep{
					resource.TestStep{
						Config: testAccCheckIbmSmIamCredentialsSecretMetadataDataSourceConfigBasic_noApikey_serviceid(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "instance_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "created_by"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "created_at"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "crn"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_group_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_type"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "updated_at"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "versions_total"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "ttl"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "reuse_api_key"),
						),
					},
				},
			})
		} else {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { acc.TestAccPreCheck(t) },
				Providers: acc.TestAccProviders,
				Steps: []resource.TestStep{
					resource.TestStep{
						Config: testAccCheckIbmSmIamCredentialsSecretMetadataDataSourceConfigBasic_noApikey_accessGroup(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "instance_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "created_by"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "created_at"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "crn"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_group_id"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "secret_type"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "updated_at"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "versions_total"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "ttl"),
							resource.TestCheckResourceAttrSet("data.ibm_sm_iam_credentials_secret_metadata.sm_iam_credentials_secret_metadata", "reuse_api_key"),
						),
					},
				},
			})
		}
	}

}

func testAccCheckIbmSmIamCredentialsSecretMetadataDataSourceConfigBasic_withApikey_serviceid() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource"
			api_key = "%s"
		}

		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret_instance" {
			instance_id   = "%s"
			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			service_id = "%s"
  			ttl = "1800"
			name = "iam-credentials-test-terraform"
			reuse_api_key = true
			depends_on = [
				ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration_instance
  			]
		}

		data "ibm_sm_iam_credentials_secret_metadata" "sm_iam_credentials_secret_metadata" {
			instance_id   = "%s"
			region        = "%s"
			secret_id = ibm_sm_iam_credentials_secret.sm_iam_credentials_secret_instance.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsConfigurationApiKey, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsSecretServiceId, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmIamCredentialsSecretMetadataDataSourceConfigBasic_withApikey_accessGroup() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = "terraform-test-datasource"
			api_key = "%s"
		}

		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret_instance" {
			instance_id   = "%s"
			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			access_groups = ["%s"]
  			ttl = "1800"
			name = "iam-credentials-test-terraform"
			reuse_api_key = true
			depends_on = [
				ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration_instance
  			]
		}

		data "ibm_sm_iam_credentials_secret_metadata" "sm_iam_credentials_secret_metadata" {
			instance_id   = "%s"
			region        = "%s"
			secret_id = ibm_sm_iam_credentials_secret.sm_iam_credentials_secret_instance.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsConfigurationApiKey, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsSecretServiceAccessGroup, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmIamCredentialsSecretMetadataDataSourceConfigBasic_noApikey_serviceid() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret_instance" {
			instance_id   = "%s"
			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			service_id = "%s"
  			ttl = "1800"
			name = "iam-credentials-test-terraform"
			reuse_api_key = true
		}

		data "ibm_sm_iam_credentials_secret_metadata" "sm_iam_credentials_secret_metadata" {
			instance_id   = "%s"
			region        = "%s"
			secret_id = ibm_sm_iam_credentials_secret.sm_iam_credentials_secret_instance.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsSecretServiceId, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmIamCredentialsSecretMetadataDataSourceConfigBasic_noApikey_accessGroup() string {
	return fmt.Sprintf(`
		resource "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret_instance" {
			instance_id   = "%s"
			region        = "%s"
  			custom_metadata = {"key":"value"}
  			description = "Extended description for this secret."
  			labels = ["my-label"]
  			access_groups = ["%s"]
  			ttl = "1800"
			name = "iam-credentials-test-terraform"
			reuse_api_key = true
		}

		data "ibm_sm_iam_credentials_secret_metadata" "sm_iam_credentials_secret_metadata" {
			instance_id   = "%s"
			region        = "%s"
			secret_id = ibm_sm_iam_credentials_secret.sm_iam_credentials_secret_instance.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerIamCredentialsSecretServiceAccessGroup, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
