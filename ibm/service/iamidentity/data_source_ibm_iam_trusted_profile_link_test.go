// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMTrustedProfileLinkDataSourceBasic(t *testing.T) {
	profileLinkProfileName := fmt.Sprintf("tf_profile_%d", acctest.RandIntRange(10, 100))
	profileLinkCrType := "IKS_SA"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileLinkDataSourceConfigBasic(profileLinkProfileName, profileLinkCrType),
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

func TestAccIBMIAMTrustedProfileLinkDataSourceAllArgs(t *testing.T) {
	profileLinkProfileName := fmt.Sprintf("tf_profile_%d", acctest.RandIntRange(10, 100))
	profileLinkCrType := "IKS_SA"
	profileLinkName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileLinkDataSourceConfig(profileLinkProfileName, profileLinkCrType, profileLinkName),
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

func testAccCheckIBMIamTrustedProfileLinkDataSourceConfigBasic(profileLinkProfileName string, profileLinkCrType string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
		}
		resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id
			cr_type = "%s"
			link {
				crn = ibm_iam_trusted_profile.iam_trusted_profile.crn # just need a valid CRN format for testing, does not need to be CR
				namespace = "namespace"
				name = "name"
			}
		}

		data "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile_id = ibm_iam_trusted_profile_link.iam_trusted_profile_link.profile_id
			link_id = ibm_iam_trusted_profile_link.iam_trusted_profile_link.link_id
		}
	`, profileLinkProfileName, profileLinkCrType)
}

func testAccCheckIBMIamTrustedProfileLinkDataSourceConfig(profileLinkProfileName string, profileLinkCrType string, profileLinkName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
		}
		resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id
			cr_type = "%s"
			link {
				crn = ibm_iam_trusted_profile.iam_trusted_profile.crn # just need a valid CRN format for testing, does not need to be CR
				namespace = "namespace"
				name = "name"
			}
			name = "%s"
		}

		data "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile_id = ibm_iam_trusted_profile_link.iam_trusted_profile_link.profile_id
			link_id = ibm_iam_trusted_profile_link.iam_trusted_profile_link.link_id
		}
	`, profileLinkProfileName, profileLinkCrType, profileLinkName)
}
