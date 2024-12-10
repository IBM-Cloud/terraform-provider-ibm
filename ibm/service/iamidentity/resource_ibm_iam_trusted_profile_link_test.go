// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMIAMTrustedProfileLinkBasic(t *testing.T) {
	var conf iamidentityv1.ProfileLink
	profileName := fmt.Sprintf("tf_profile_%d", acctest.RandIntRange(10, 100))
	crType := "IKS_SA"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileLinkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileLinkConfigBasic(profileName, crType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileLinkExists("ibm_iam_trusted_profile_link.iam_trusted_profile_link", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_link.iam_trusted_profile_link", "cr_type", crType),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfileLinkAllArgs(t *testing.T) {
	var conf iamidentityv1.ProfileLink
	profileName := fmt.Sprintf("tf_profile_%d", acctest.RandIntRange(10, 100))
	crType := "IKS_SA"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileLinkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileLinkConfig(profileName, crType, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileLinkExists("ibm_iam_trusted_profile_link.iam_trusted_profile_link", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_link.iam_trusted_profile_link", "cr_type", crType),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_link.iam_trusted_profile_link", "name", name),
				),
			},
			{
				ResourceName:      "ibm_iam_trusted_profile_link.iam_trusted_profile_link",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileLinkConfigBasic(profileName string, crType string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
		}
		resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id
			cr_type = "%s"
			link {
				crn = "%s"
				namespace = "namespace"
				name = "name"
			}
		}
	`, profileName, crType, acc.IksSa)
}

func testAccCheckIBMIamTrustedProfileLinkConfig(profileName string, crType string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
		}
		resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
			profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id
			cr_type = "%s"
			link {
				crn = "%s"
				namespace = "namespace"
				name = "name"
			}
			name = "%s"
		}
	`, profileName, crType, acc.IksSa, name)
}

func testAccCheckIBMIamTrustedProfileLinkExists(n string, obj iamidentityv1.ProfileLink) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getLinkOptions := &iamidentityv1.GetLinkOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getLinkOptions.SetProfileID(parts[0])
		getLinkOptions.SetLinkID(parts[1])

		profileLink, _, err := iamIdentityClient.GetLink(getLinkOptions)
		if err != nil {
			return err
		}

		obj = *profileLink
		return nil
	}
}

func testAccCheckIBMIamTrustedProfileLinkDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile_link" {
			continue
		}

		getLinkOptions := &iamidentityv1.GetLinkOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getLinkOptions.SetProfileID(parts[0])
		getLinkOptions.SetLinkID(parts[1])

		// Try to find the key
		_, response, err := iamIdentityClient.GetLink(getLinkOptions)

		if err == nil {
			return fmt.Errorf("iam_trusted_profile_link still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for iam_trusted_profile_link (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
