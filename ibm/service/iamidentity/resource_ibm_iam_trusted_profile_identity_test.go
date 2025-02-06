// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMIamTrustedProfileIdentityBasic(t *testing.T) {
	var conf iamidentityv1.ProfileIdentityResponse
	profileID := acc.IAMTrustedProfileID
	identityType := "serviceid"
	identifier := acc.IAMServiceId
	typeVar := "serviceid"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileIdentityDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileIdentityConfigBasic(profileID, identityType, identifier, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileIdentityExists("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_instance", "profile_id", profileID),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_instance", "identity_type", identityType),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_instance", "identifier", identifier),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_instance", "type", typeVar),
				),
			},
		},
	})
}

func TestAccIBMIamTrustedProfileIdentityAllArgs(t *testing.T) {
	var conf iamidentityv1.ProfileIdentityResponse
	profileID := acc.IAMTrustedProfileID
	identityType := "user"
	identifier := acc.Ibmid2
	typeVar := "user"
	description := fmt.Sprintf("tf_description_%s", "profile identity description")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileIdentityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileIdentityConfig(profileID, identityType, identifier, typeVar, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileIdentityExists("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "profile_id", profileID),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "identity_type", identityType),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "identifier", identifier),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "description", description),
				),
			},
			{
				ResourceName:      "ibm_iam_trusted_profile_identity.iam_trusted_profile_identity",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileIdentityConfigBasic(profileID string, identityType string, identifier string, typeVar string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile_identity" "iam_trusted_profile_identity" {
			profile_id = "%s"
			identity_type = "%s"
			identifier = "%s"
			type = "%s"
		}
	`, profileID, identityType, identifier, typeVar)
}

func testAccCheckIBMIamTrustedProfileIdentityConfig(profileID string, identityType string, identifier string, typeVar string, description string) string {
	acountId := acc.IAMAccountId
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile_identity" "iam_trusted_profile_identity" {
			profile_id = "%s"
			identity_type = "%s"
			identifier = "%s"
			type = "%s"
			accounts = [
				"%s"
			]
			description = "%s"
		}
	`, profileID, identityType, identifier, typeVar, acountId, description)
}

func testAccCheckIBMIamTrustedProfileIdentityExists(n string, obj iamidentityv1.ProfileIdentityResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getProfileIdentityOptions := &iamidentityv1.GetProfileIdentityOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "|")
		if err != nil {
			return err
		}

		getProfileIdentityOptions.SetProfileID(parts[0])
		getProfileIdentityOptions.SetIdentityType(parts[1])
		getProfileIdentityOptions.SetIdentifierID(parts[2])

		profileIdentityResponse, _, err := iamIdentityClient.GetProfileIdentity(getProfileIdentityOptions)
		if err != nil {
			return err
		}

		obj = *profileIdentityResponse
		return nil
	}
}

func testAccCheckIBMIamTrustedProfileIdentityDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile_identity" {
			continue
		}

		getProfileIdentityOptions := &iamidentityv1.GetProfileIdentityOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "|")
		if err != nil {
			return err
		}

		getProfileIdentityOptions.SetProfileID(parts[0])
		getProfileIdentityOptions.SetIdentityType(parts[1])
		getProfileIdentityOptions.SetIdentifierID(parts[2])

		// Try to find the key
		_, response, err := iamIdentityClient.GetProfileIdentity(getProfileIdentityOptions)

		if err == nil {
			return fmt.Errorf("iam_trusted_profile_identity still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 400 {
			return fmt.Errorf("Error checking for iam_trusted_profile_identity (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
