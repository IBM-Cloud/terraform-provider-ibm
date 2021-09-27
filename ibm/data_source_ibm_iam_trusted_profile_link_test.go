// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIamTrustedProfileLinkDataSourceBasic(t *testing.T) {
	profileLinkProfileID := fmt.Sprintf("tf_profile_id_%d", acctest.RandIntRange(10, 100))
	profileLinkCrType := fmt.Sprintf("tf_cr_type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileLinkDataSourceConfigBasic(profileLinkProfileID, profileLinkCrType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "link_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "cr_type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "link.#"),
				),
			},
		},
	})
}

func TestAccIBMIamTrustedProfileLinkDataSourceAllArgs(t *testing.T) {
	profileLinkProfileID := fmt.Sprintf("tf_profile_id_%d", acctest.RandIntRange(10, 100))
	profileLinkCrType := fmt.Sprintf("tf_cr_type_%d", acctest.RandIntRange(10, 100))
	profileLinkName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileLinkDataSourceConfig(profileLinkProfileID, profileLinkCrType, profileLinkName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "link_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "cr_type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_link.iam_trusted_profile_link", "link.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileLinkDataSourceConfigBasic(profileLinkProfileID string, profileLinkCrType string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile_id = "%s"
			cr_type = "%s"
			link {
				crn = "crn"
				namespace = "namespace"
				name = "name"
			}
		}

		data "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile-id = ibm_iam_trusted_profile_link.iam_trusted_profile_link.profile_id
			link-id = "link-id"
		}
	`, profileLinkProfileID, profileLinkCrType)
}

func testAccCheckIBMIamTrustedProfileLinkDataSourceConfig(profileLinkProfileID string, profileLinkCrType string, profileLinkName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile_id = "%s"
			cr_type = "%s"
			link {
				crn = "crn"
				namespace = "namespace"
				name = "name"
			}
			name = "%s"
		}

		data "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile-id = ibm_iam_trusted_profile_link.iam_trusted_profile_link.profile_id
			link-id = "link-id"
		}
	`, profileLinkProfileID, profileLinkCrType, profileLinkName)
}
