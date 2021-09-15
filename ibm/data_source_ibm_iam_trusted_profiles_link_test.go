// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIamTrustedProfilesLinkDataSourceBasic(t *testing.T) {
	profileLinkProfileID := fmt.Sprintf("tf_profile_id_%d", acctest.RandIntRange(10, 100))
	profileLinkCrType := fmt.Sprintf("tf_cr_type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesLinkDataSourceConfigBasic(profileLinkProfileID, profileLinkCrType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "link_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "cr_type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "link.#"),
				),
			},
		},
	})
}

func TestAccIBMIamTrustedProfilesLinkDataSourceAllArgs(t *testing.T) {
	profileLinkProfileID := fmt.Sprintf("tf_profile_id_%d", acctest.RandIntRange(10, 100))
	profileLinkCrType := fmt.Sprintf("tf_cr_type_%d", acctest.RandIntRange(10, 100))
	profileLinkName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesLinkDataSourceConfig(profileLinkProfileID, profileLinkCrType, profileLinkName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "link_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "cr_type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles_link.iam_trusted_profiles_link", "link.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfilesLinkDataSourceConfigBasic(profileLinkProfileID string, profileLinkCrType string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profiles_link" "iam_trusted_profiles_link" {
			profile_id = "%s"
			cr_type = "%s"
			link {
				crn = "crn"
				namespace = "namespace"
				name = "name"
			}
		}

		data "ibm_iam_trusted_profiles_link" "iam_trusted_profiles_link" {
			profile-id = ibm_iam_trusted_profiles_link.iam_trusted_profiles_link.profile_id
			link-id = "link-id"
		}
	`, profileLinkProfileID, profileLinkCrType)
}

func testAccCheckIBMIamTrustedProfilesLinkDataSourceConfig(profileLinkProfileID string, profileLinkCrType string, profileLinkName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profiles_link" "iam_trusted_profiles_link" {
			profile_id = "%s"
			cr_type = "%s"
			link {
				crn = "crn"
				namespace = "namespace"
				name = "name"
			}
			name = "%s"
		}

		data "ibm_iam_trusted_profiles_link" "iam_trusted_profiles_link" {
			profile-id = ibm_iam_trusted_profiles_link.iam_trusted_profiles_link.profile_id
			link-id = "link-id"
		}
	`, profileLinkProfileID, profileLinkCrType, profileLinkName)
}
