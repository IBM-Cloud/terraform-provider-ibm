// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIamTrustedProfilesDataSourceBasic(t *testing.T) {
	trustedProfileName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	trustedProfileAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesDataSourceConfigBasic(trustedProfileName, trustedProfileAccountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "account_id"),
				),
			},
		},
	})
}

func TestAccIBMIamTrustedProfilesDataSourceAllArgs(t *testing.T) {
	trustedProfileName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	trustedProfileAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	trustedProfileDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesDataSourceConfig(trustedProfileName, trustedProfileAccountID, trustedProfileDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "context.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "ims_account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "ims_user_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "history.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "history.0.timestamp"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "history.0.iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "history.0.iam_id_account"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "history.0.action"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "history.0.message"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfilesDataSourceConfigBasic(trustedProfileName string, trustedProfileAccountID string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profiles" "iam_trusted_profiles" {
			name = "%s"
			account_id = "%s"
		}

		data "ibm_iam_trusted_profiles" "iam_trusted_profiles" {
			profile-id = "profile-id"
		}
	`, trustedProfileName, trustedProfileAccountID)
}

func testAccCheckIBMIamTrustedProfilesDataSourceConfig(trustedProfileName string, trustedProfileAccountID string, trustedProfileDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profiles" "iam_trusted_profiles" {
			name = "%s"
			account_id = "%s"
			description = "%s"
		}

		data "ibm_iam_trusted_profiles" "iam_trusted_profiles" {
			profile-id = "profile-id"
		}
	`, trustedProfileName, trustedProfileAccountID, trustedProfileDescription)
}
