// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIamIdentityPreferenceBasic(t *testing.T) {
	var conf iamidentityv1.IdentityPreferenceResponse
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	iamID := fmt.Sprintf("tf_iam_id_%d", acctest.RandIntRange(10, 100))
	service := fmt.Sprintf("tf_service_%d", acctest.RandIntRange(10, 100))
	preferenceID := fmt.Sprintf("tf_preference_id_%d", acctest.RandIntRange(10, 100))
	valueString := fmt.Sprintf("tf_value_string_%d", acctest.RandIntRange(10, 100))
	valueStringUpdate := fmt.Sprintf("tf_value_string_%d", acctest.RandIntRange(10, 100))

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
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "iam_id", iamID),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "service", service),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "preference_id", preferenceID),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "value_string", valueString),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamIdentityPreferenceConfigBasic(accountID, iamID, service, preferenceID, valueStringUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "iam_id", iamID),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "service", service),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "preference_id", preferenceID),
					resource.TestCheckResourceAttr("ibm_iam_identity_preference.iam_identity_preference_instance", "value_string", valueStringUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_identity_preference.iam_identity_preference",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamIdentityPreferenceConfigBasic(accountID string, iamID string, service string, preferenceID string, valueString string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_identity_preference" "iam_identity_preference_instance" {
			account_id = "%s"
			iam_id = "%s"
			service = "%s"
			preference_id = "%s"
			value_string = "%s"
		}
	`, accountID, iamID, service, preferenceID, valueString)
}

func testAccCheckIBMIamIdentityPreferenceExists(n string, obj iamidentityv1.IdentityPreferenceResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IamIdentityV1()
		if err != nil {
			return err
		}

		getPreferenceOnScopeAccountOptions := &iamidentityv1.GetPreferenceOnScopeAccountOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPreferenceOnScopeAccountOptions.SetAccountID(parts[0])
		getPreferenceOnScopeAccountOptions.SetIamID(parts[1])
		getPreferenceOnScopeAccountOptions.SetService(parts[2])
		getPreferenceOnScopeAccountOptions.SetPreferenceID(parts[3])
		getPreferenceOnScopeAccountOptions.SetPreferenceID(parts[4])

		identityPreferenceResponse, _, err := iamIdentityClient.GetPreferenceOnScopeAccount(getPreferenceOnScopeAccountOptions)
		if err != nil {
			return err
		}

		obj = *identityPreferenceResponse
		return nil
	}
}

func testAccCheckIBMIamIdentityPreferenceDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IamIdentityV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_identity_preference" {
			continue
		}

		getPreferenceOnScopeAccountOptions := &iamidentityv1.GetPreferenceOnScopeAccountOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPreferenceOnScopeAccountOptions.SetAccountID(parts[0])
		getPreferenceOnScopeAccountOptions.SetIamID(parts[1])
		getPreferenceOnScopeAccountOptions.SetService(parts[2])
		getPreferenceOnScopeAccountOptions.SetPreferenceID(parts[3])
		getPreferenceOnScopeAccountOptions.SetPreferenceID(parts[4])

		// Try to find the key
		_, response, err := iamIdentityClient.GetPreferenceOnScopeAccount(getPreferenceOnScopeAccountOptions)

		if err == nil {
			return fmt.Errorf("iam_identity_preference still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_identity_preference (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
