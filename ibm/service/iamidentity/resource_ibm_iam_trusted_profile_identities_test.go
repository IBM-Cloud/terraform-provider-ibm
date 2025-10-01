// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMIamTrustedProfileIdentitiesBasic(t *testing.T) {
	var conf iamidentityv1.ProfileIdentitiesResponse
	profileID := acc.IAMTrustedProfileID
	accountId := acc.IAMAccountId
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	ibmID1 := acc.Ibmid1
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileIdentitiesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileIdentitiesConfigBasic(name, profileID, ibmID1, accountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileIdentitiesExists("ibm_iam_trusted_profile_identities.iam_trusted_profile_identities", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identities.iam_trusted_profile_identities", "profile_id", profileID),
				),
			},
			{
				Config: testAccCheckIBMIamTrustedProfileIdentitiesConfigBasic(name, profileID, ibmID1, accountId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_identities.iam_trusted_profile_identities", "profile_id", profileID),
					testCheckIdentitiesCount(),
				),
			},
			{
				ResourceName:      "ibm_iam_trusted_profile_identities.iam_trusted_profile_identities",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileIdentitiesConfigBasic(name string, profileID string, ibmID1 string, accountId string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_service_id" "serviceID" {
			name        = "%s"
			description = "ServiceID for test"
		}

		resource "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities" {
			profile_id = "%s"
			identities {
				type       = "serviceid"
				iam_id     = ibm_iam_service_id.serviceID.iam_id
				identifier = ibm_iam_service_id.serviceID.id
			}
			identities {
				iam_id     = "%s"
				type       = "user"
				identifier = "%s"
				accounts = ["%s"]
				description = "tf_description_profile identity description"
            }
		}
	`, name, profileID, ibmID1, ibmID1, accountId)
}

func testCheckIdentitiesCount() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources["ibm_iam_trusted_profile_identities.iam_trusted_profile_identities"]
		if !ok {
			return fmt.Errorf("not found: %s", "ibm_iam_trusted_profile_identities.iam_trusted_profile_identities")
		}

		countStr := rs.Primary.Attributes["identities.#"]
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return err
		}

		if count < 2 {
			return fmt.Errorf("expected at least 2 identities, got %d", count)
		}

		return nil
	}
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
