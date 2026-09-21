// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMIamIdentityPreferenceBasic(t *testing.T) {
	var conf iamidentityv1.IdentityPreferenceResponse
	accountID := acc.IAMAccountId
	iamID := acc.IAMTrustedProfileID
	service := "console"
	preferenceID := "landing_page"
	valueString := "/iam"
	valueStringUpdate := "/devops"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamIdentityPreferenceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamIdentityPreferenceConfigBasic(accountID, iamID, service, preferenceID, valueString),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamIdentityPreferenceExists("ibm_iam_identity_preference.iam_identity_preference_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "service", service),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "preference_id", preferenceID),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "value_string", valueString),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamIdentityPreferenceConfigBasic(accountID, iamID, service, preferenceID, valueStringUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "service", service),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "preference_id", preferenceID),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "value_string", valueStringUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_identity_preference.iam_identity_preference_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamIdentityPreferenceConfigBasic(accountID string, iamID string, service string, preferenceID string, valueString string) string {
	// Creation of preference not supported, so import existing preference to a resource
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
	`, accountID, iamID, service, preferenceID, valueString)
}

func testAccCheckIBMIamIdentityPreferenceExists(n string, obj iamidentityv1.IdentityPreferenceResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getPreferencesOnScopeAccountOptions := &iamidentityv1.GetPreferencesOnScopeAccountOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPreferencesOnScopeAccountOptions.SetAccountID(parts[0])
		getPreferencesOnScopeAccountOptions.SetIamID(parts[1])
		getPreferencesOnScopeAccountOptions.SetService(parts[2])
		getPreferencesOnScopeAccountOptions.SetPreferenceID(parts[3])

		identityPreferenceResponse, _, err := iamIdentityClient.GetPreferencesOnScopeAccount(getPreferencesOnScopeAccountOptions)
		if err != nil {
			return err
		}

		obj = *identityPreferenceResponse
		return nil
	}
}

func testAccCheckIBMIamIdentityPreferenceDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_identity_preference" {
			continue
		}

		deletePreferenceOnScopeAccountOptions := &iamidentityv1.DeletePreferencesOnScopeAccountOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		deletePreferenceOnScopeAccountOptions.SetAccountID(parts[0])
		deletePreferenceOnScopeAccountOptions.SetIamID(parts[1])
		deletePreferenceOnScopeAccountOptions.SetService(parts[2])
		deletePreferenceOnScopeAccountOptions.SetPreferenceID(parts[3])

		// Verify the preference was deleted (reset to defaults) by calling delete again
		// A successful delete returns 204, calling delete on an already deleted preference should also return 204
		response, err := iamIdentityClient.DeletePreferencesOnScopeAccount(deletePreferenceOnScopeAccountOptions)

		if err != nil && response != nil && response.StatusCode != 204 {
			return fmt.Errorf("Error verifying iam_identity_preference (%s) was destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
