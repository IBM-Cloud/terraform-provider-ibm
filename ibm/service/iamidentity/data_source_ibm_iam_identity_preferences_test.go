// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIamIdentityPreferencesDataSourceBasic(t *testing.T) {
	accountID := acc.IAMAccountId
	iamID := acc.IAMTrustedProfileID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamIdentityPreferencesDataSourceConfigBasic(accountID, iamID),
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

func testAccCheckIBMIamIdentityPreferencesDataSourceConfigBasic(accountID string, iamID string) string {
	return fmt.Sprintf(`
		data "ibm_iam_identity_preferences" "iam_identity_preferences_instance" {
			account_id = "%s"
			iam_id = "%s"
		}
	`, accountID, iamID)
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
