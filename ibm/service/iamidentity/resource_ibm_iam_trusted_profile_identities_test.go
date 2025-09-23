// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/iamidentity"
)

func TestAccIBMIamTrustedProfileIdentitiesBasic(t *testing.T) {
	var conf iamidentityv1.ProfileIdentitiesResponse
	profileID := acc.IAMTrustedProfileID
	ibmID1 := acc.Ibmid1
	ibmID2 := acc.Ibmid2
	identities := []map[string]interface{}{
		{
			"identity_type": "user",
			"identifier":    acc.Ibmid2,
			"type":          "user",
			"description":   fmt.Sprintf("tf_description_%s", "profile identity description"),
		},
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileIdentitiesDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileIdentitiesConfigBasic(profileID, ibmID1, ibmID2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileIdentitiesExists("ibm_iam_trusted_profile_identities.iam_trusted_profile_identities", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identities.iam_trusted_profile_identities", "profile_id", profileID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileIdentitiesConfigBasic(profileID, ibmID1, ibmID2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identities.iam_trusted_profile_identities", "profile_id", profileID),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identities.iam_trusted_profile_identities", "identities.0.type", identities[0]["type"].(string)),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_trusted_profile_identities.iam_trusted_profile_identities",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileIdentitiesConfigBasic(profileID string, ibmID1 string, ibmID2 string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities" {
          profile_id = "%s"
		  identities {
		    iam_id     = "%s"
			type       = "user"
			identifier = "%s"
			accounts = ["86a1004d3f1848a291de32874cb48120"]
			description = "tf_description_profile identity description"
			}
			identities {
		    iam_id     = "%s"
			type       = "user"
			identifier = "%s"
			accounts = ["86a1004d3f1848a291de32874cb48120"]
			description = "tf_description_profile identity description"
            }
		}
	`, profileID, ibmID1, ibmID1, ibmID2, ibmID2)
}

func testAccCheckIBMIamTrustedProfileIdentitiesExists(n string, obj iamidentityv1.ProfileIdentitiesResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getProfileIdentitiesOptions := &iamidentityv1.GetProfileIdentitiesOptions{}

		getProfileIdentitiesOptions.SetProfileID(rs.Primary.ID)

		profileIdentitiesResponse, _, err := iamIdentityClient.GetProfileIdentities(getProfileIdentitiesOptions)
		if err != nil {
			return err
		}

		obj = *profileIdentitiesResponse
		return nil
	}
}

func testAccCheckIBMIamTrustedProfileIdentitiesDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile_identities" {
			continue
		}

		getProfileIdentitiesOptions := &iamidentityv1.GetProfileIdentitiesOptions{}
		getProfileIdentitiesOptions.SetProfileID(rs.Primary.ID)

		resp, _, err := iamIdentityClient.GetProfileIdentities(getProfileIdentitiesOptions)
		if err != nil {
			return err
		}
		if len(resp.Identities) > 0 {
			return fmt.Errorf("iam_trusted_profile_identities still exist for profile: %s", rs.Primary.ID)
		}
	}
	return nil
}

func TestResourceIBMIamTrustedProfileIdentitiesProfileIdentityResponseToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["iam_id"] = "testString"
		model["identifier"] = "testString"
		model["type"] = "user"
		model["accounts"] = []string{"testString"}
		model["description"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(iamidentityv1.ProfileIdentityResponse)
	model.IamID = core.StringPtr("testString")
	model.Identifier = core.StringPtr("testString")
	model.Type = core.StringPtr("user")
	model.Accounts = []string{"testString"}
	model.Description = core.StringPtr("testString")

	result, err := iamidentity.ResourceIBMIamTrustedProfileIdentitiesProfileIdentityResponseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIamTrustedProfileIdentitiesMapToProfileIdentityRequest(t *testing.T) {
	checkResult := func(result *iamidentityv1.ProfileIdentityRequest) {
		model := new(iamidentityv1.ProfileIdentityRequest)
		model.Identifier = core.StringPtr("testString")
		model.Type = core.StringPtr("user")
		model.Accounts = []string{"testString"}
		model.Description = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["identifier"] = "testString"
	model["type"] = "user"
	model["accounts"] = []interface{}{"testString"}
	model["description"] = "testString"

	result, err := iamidentity.ResourceIBMIamTrustedProfileIdentitiesMapToProfileIdentityRequest(model)
	assert.Nil(t, err)
	checkResult(result)
}
