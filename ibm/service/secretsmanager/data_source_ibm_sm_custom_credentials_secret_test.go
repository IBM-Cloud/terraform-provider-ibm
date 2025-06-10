// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmCustomCredentialsSecretDataSource(t *testing.T) {
	dataSourceName := "ibm_sm_custom_credentials_secret.sm_custom_credentials_secret"
	expectedStrCredential := customCredentialsStrParam + "_output"
	expectedBoolCredential := !customCredentialsBoolParam
	expectedIntCredential := customCredentialsIntParam + 1
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: customCredentialsSecretDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIbmSmCustomCredentialsSecretCreated(dataSourceName),
					resource.TestCheckResourceAttrSet(dataSourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "created_by"),
					resource.TestCheckResourceAttrSet(dataSourceName, "created_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(dataSourceName, "crn"),
					resource.TestCheckResourceAttrSet(dataSourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(dataSourceName, "next_rotation_date"),
					resource.TestCheckResourceAttr(dataSourceName, "state", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "ttl", customCredentialsTtl),
					resource.TestCheckResourceAttr(dataSourceName, "parameters.0.string_values.str_param", customCredentialsStrParam),
					resource.TestCheckResourceAttr(dataSourceName, "parameters.0.integer_values.int_param", strconv.Itoa(customCredentialsIntParam)),
					resource.TestCheckResourceAttr(dataSourceName, "parameters.0.boolean_values.bool_param", strconv.FormatBool(customCredentialsBoolParam)),
					resource.TestCheckResourceAttr(dataSourceName, "credentials_content.0.string_values.str_credential", expectedStrCredential),
					resource.TestCheckResourceAttr(dataSourceName, "credentials_content.0.integer_values.int_credential", strconv.Itoa(expectedIntCredential)),
					resource.TestCheckResourceAttr(dataSourceName, "credentials_content.0.boolean_values.bool_credential", strconv.FormatBool(expectedBoolCredential)),
				),
			},
		},
	})
}

func customCredentialsSecretDataSourceConfig() string {
	return customCredentialsSecretConfigAllArgs() +
		fmt.Sprintf(`
		data "ibm_sm_custom_credentials_secret" "sm_custom_credenbtials_secret" {
			instance_id = "%s"
			region = "%s"
			secret_id = ibm_sm_custom_credentials_secret.sm_custom_credentials_secret.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
