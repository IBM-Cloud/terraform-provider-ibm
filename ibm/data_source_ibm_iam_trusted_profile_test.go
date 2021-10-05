// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIamTrustedProfileDataSourceBasic(t *testing.T) {
	trustedProfileName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	trustedProfileAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileDataSourceConfigBasic(trustedProfileName, trustedProfileAccountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "account_id"),
				),
			},
		},
	})
}

func TestAccIBMIamTrustedProfileDataSourceAllArgs(t *testing.T) {
	trustedProfileName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	trustedProfileAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	trustedProfileDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileDataSourceConfig(trustedProfileName, trustedProfileAccountID, trustedProfileDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "ims_account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "ims_user_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "history.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "history.0.timestamp"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "history.0.iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "history.0.iam_id_account"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "history.0.action"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile", "history.0.message"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileDataSourceConfigBasic(trustedProfileName string, trustedProfileAccountID string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
			account_id = "%s"
		}

		data "ibm_iam_trusted_profile" "iam_trusted_profile" {
			profile_id = "profile-id"
		}
	`, trustedProfileName, trustedProfileAccountID)
}

func testAccCheckIBMIamTrustedProfileDataSourceConfig(trustedProfileName string, trustedProfileAccountID string, trustedProfileDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
			account_id = "%s"
			description = "%s"
		}

		data "ibm_iam_trusted_profile" "iam_trusted_profile" {
			profile_id = "profile-id"
		}
	`, trustedProfileName, trustedProfileAccountID, trustedProfileDescription)
}
