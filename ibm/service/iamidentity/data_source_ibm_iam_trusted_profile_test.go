// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIAMTrustedProfileDataSourceBasic(t *testing.T) {
	trustedProfileName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileDataSourceConfigBasic(trustedProfileName),
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

func TestAccIBMIAMTrustedProfileDataSourceAllArgs(t *testing.T) {
	trustedProfileName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	trustedProfileDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileDataSourceConfig(trustedProfileName, trustedProfileDescription),
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
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileDataSourceConfigBasic(trustedProfileName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
		}

		data "ibm_iam_trusted_profile" "iam_trusted_profile" {
			profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id
		}
	`, trustedProfileName)
}

func testAccCheckIBMIamTrustedProfileDataSourceConfig(trustedProfileName string, trustedProfileDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
			description = "%s"
		}

		data "ibm_iam_trusted_profile" "iam_trusted_profile" {
			profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id
		}
	`, trustedProfileName, trustedProfileDescription)
}
