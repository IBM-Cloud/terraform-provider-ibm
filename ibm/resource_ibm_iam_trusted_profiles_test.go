// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMIamTrustedProfilesBasic(t *testing.T) {
	var conf iamidentityv1.TrustedProfile
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfilesDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesConfigBasic(name, accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfilesExists("ibm_iam_trusted_profiles.iam_trusted_profiles", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "account_id", accountID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesConfigBasic(nameUpdate, accountID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "account_id", accountID),
				),
			},
		},
	})
}

func TestAccIBMIamTrustedProfilesAllArgs(t *testing.T) {
	var conf iamidentityv1.TrustedProfile
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfilesDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesConfig(name, accountID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfilesExists("ibm_iam_trusted_profiles.iam_trusted_profiles", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesConfig(nameUpdate, accountID, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profiles.iam_trusted_profiles", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_trusted_profiles.iam_trusted_profiles",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfilesConfigBasic(name string, accountID string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profiles" "iam_trusted_profiles" {
			name = "%s"
			account_id = "%s"
		}
	`, name, accountID)
}

func testAccCheckIBMIamTrustedProfilesConfig(name string, accountID string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profiles" "iam_trusted_profiles" {
			name = "%s"
			account_id = "%s"
			description = "%s"
		}
	`, name, accountID, description)
}

func testAccCheckIBMIamTrustedProfilesExists(n string, obj iamidentityv1.TrustedProfile) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getProfileOptions := &iamidentityv1.GetProfileOptions{}

		getProfileOptions.SetProfileID(rs.Primary.ID)

		trustedProfile, _, err := iamIdentityClient.GetProfile(getProfileOptions)
		if err != nil {
			return err
		}

		obj = *trustedProfile
		return nil
	}
}

func testAccCheckIBMIamTrustedProfilesDestroy(s *terraform.State) error {
	iamIdentityClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profiles" {
			continue
		}

		getProfileOptions := &iamidentityv1.GetProfileOptions{}

		getProfileOptions.SetProfileID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamIdentityClient.GetProfile(getProfileOptions)

		if err == nil {
			return fmt.Errorf("iam_trusted_profiles still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_trusted_profiles (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
