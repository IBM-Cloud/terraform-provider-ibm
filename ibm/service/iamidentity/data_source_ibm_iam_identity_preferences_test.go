// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIamIdentityPreferencesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamIdentityPreferencesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_identity_preferences.iam_identity_preferences_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_identity_preferences.iam_identity_preferences_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_identity_preferences.iam_identity_preferences_instance", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_identity_preferences.iam_identity_preferences_instance", "preferences.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamIdentityPreferencesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_identity_preferences" "iam_identity_preferences_instance" {
			account_id = "account_id"
			iam_id = "iam_id"
		}
	`)
}

func TestDataSourceIBMIamIdentityPreferencesIdentityPreferenceResponseToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["service"] = "testString"
		model["id"] = "testString"
		model["account_id"] = "testString"
		model["scope"] = "testString"
		model["value_string"] = "testString"
		model["value_list_of_strings"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(iamidentityv1.IdentityPreferenceResponse)
	model.Service = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")
	model.AccountID = core.StringPtr("testString")
	model.Scope = core.StringPtr("testString")
	model.ValueString = core.StringPtr("testString")
	model.ValueListOfStrings = []string{"testString"}

	result, err := iamidentity.DataSourceIBMIamIdentityPreferencesIdentityPreferenceResponseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
