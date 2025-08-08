// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.1-067d600b-20250616-154447
*/

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIamIdentityPreferenceDataSourceBasic(t *testing.T) {
	identityPreferenceResponseAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	identityPreferenceResponseIamID := fmt.Sprintf("tf_iam_id_%d", acctest.RandIntRange(10, 100))
	identityPreferenceResponseService := fmt.Sprintf("tf_service_%d", acctest.RandIntRange(10, 100))
	identityPreferenceResponsePreferenceID := fmt.Sprintf("tf_preference_id_%d", acctest.RandIntRange(10, 100))
	identityPreferenceResponseValueString := fmt.Sprintf("tf_value_string_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamIdentityPreferenceDataSourceConfigBasic(identityPreferenceResponseAccountID, identityPreferenceResponseIamID, identityPreferenceResponseService, identityPreferenceResponsePreferenceID, identityPreferenceResponseValueString),
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

func testAccCheckIBMIamIdentityPreferenceDataSourceConfigBasic(identityPreferenceResponseAccountID string, identityPreferenceResponseIamID string, identityPreferenceResponseService string, identityPreferenceResponsePreferenceID string, identityPreferenceResponseValueString string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_identity_preference" "iam_identity_preference_instance" {
			account_id = "%s"
			iam_id = "%s"
			service = "%s"
			preference_id = "%s"
			value_string = "%s"
		}

		data "ibm_iam_identity_preference" "iam_identity_preference_instance" {
			account_id = ibm_iam_identity_preference.iam_identity_preference_instance.account_id
			iam_id = ibm_iam_identity_preference.iam_identity_preference_instance.iam_id
			service = ibm_iam_identity_preference.iam_identity_preference_instance.service
			preference_id = ibm_iam_identity_preference.iam_identity_preference_instance.preference_id
		}
	`, identityPreferenceResponseAccountID, identityPreferenceResponseIamID, identityPreferenceResponseService, identityPreferenceResponsePreferenceID, identityPreferenceResponseValueString)
}

