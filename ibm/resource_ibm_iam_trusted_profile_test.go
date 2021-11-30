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

func TestAccIBMIAMTrustedProfileBasic(t *testing.T) {
	var conf iamidentityv1.TrustedProfile
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileExists("ibm_iam_trusted_profile.iam_trusted_profile", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfileAllArgs(t *testing.T) {
	var conf iamidentityv1.TrustedProfile
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileConfig(name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileExists("ibm_iam_trusted_profile.iam_trusted_profile", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileConfig(nameUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.iam_trusted_profile", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_trusted_profile.iam_trusted_profile",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileConfigBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
		}
	`, name)
}

func testAccCheckIBMIamTrustedProfileConfig(name string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
			description = "%s"
		}
	`, name, description)
}

func testAccCheckIBMIamTrustedProfileExists(n string, obj iamidentityv1.TrustedProfile) resource.TestCheckFunc {

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

func testAccCheckIBMIamTrustedProfileDestroy(s *terraform.State) error {
	iamIdentityClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile" {
			continue
		}

		getProfileOptions := &iamidentityv1.GetProfileOptions{}

		getProfileOptions.SetProfileID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamIdentityClient.GetProfile(getProfileOptions)

		if err == nil {
			return fmt.Errorf("iam_trusted_profile still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_trusted_profile (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
