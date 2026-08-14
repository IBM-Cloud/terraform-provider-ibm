// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.106.0-09823488-20250707-071701
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIamIdentityPreferenceDataSourceBasic(t *testing.T) {
	accountID := acc.IAMAccountId
	iamID := acc.IAMTrustedProfileID
	service := "console"
	preferenceID := "landing_page"
	valueString := "/iam"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamIdentityPreferenceDataSourceConfigBasic(accountID, iamID, service, preferenceID, valueString),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_identity_preference.iam_identity_preference_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_identity_preference.iam_identity_preference_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_identity_preference.iam_identity_preference_instance", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_identity_preference.iam_identity_preference_instance", "service"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_identity_preference.iam_identity_preference_instance", "preference_id"),
				),
			},
		},
	})
}

func testAccCheckIBMIamIdentityPreferenceDataSourceConfigBasic(accountID string, iamID string, service string, preferenceID string, valueString string) string {
	return fmt.Sprintf(`
		import {
			to = ibm_iam_identity_preference.iam_identity_preference_instance
			id = "%[1]s/%[2]s/%[3]s/%[4]s"
		}

		resource "ibm_iam_identity_preference" "iam_identity_preference_instance" {
			account_id = "%[1]s"
			service = "%[3]s"
			preference_id = "%[4]s"
			value_string = "%[5]s"
		}

		data "ibm_iam_identity_preference" "iam_identity_preference_instance" {
			account_id = "%[1]s"
			iam_id = "%[2]s"
			service = "%[3]s"
			preference_id = "%[4]s"

			depends_on = [ibm_iam_identity_preference.iam_identity_preference_instance]
		}
	`, accountID, iamID, service, preferenceID, valueString)
}
