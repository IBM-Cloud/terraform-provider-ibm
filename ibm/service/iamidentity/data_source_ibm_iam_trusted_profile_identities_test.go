// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.103.0-e8b84313-20250402-201816
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIamTrustedProfileIdentitiesDataSourceBasic(t *testing.T) {
	profileIdentitiesResponseProfileID := acc.IAMTrustedProfileID
	profileIdentitiesResponseIfMatch := fmt.Sprintf("tf_if_match_%d", acctest.RandIntRange(10, 100))
	ibmID1 := acc.Ibmid1

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: func() string {
					var _ string = profileIdentitiesResponseIfMatch
					return testAccCheckIBMIamTrustedProfileIdentitiesDataSourceConfigBasic(profileIdentitiesResponseProfileID, ibmID1)
				}(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identities.iam_trusted_profile_identities_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identities.iam_trusted_profile_identities_instance", "profile_id"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileIdentitiesDataSourceConfigBasic(profileIdentitiesResponseProfileID, ibmID1 string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities_instance" {
			profile_id = "%s"
		  identities {
		    iam_id     = "%s"
			type       = "user"
			identifier = "%s"
			accounts = ["86a1004d3f1848a291de32874cb48120"]
			description = "tf_description_profile identity description"
			}
		}

		data "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities_instance" {
			profile_id = ibm_iam_trusted_profile_identities.iam_trusted_profile_identities_instance.profile_id
		}
	`, profileIdentitiesResponseProfileID, ibmID1, ibmID1)
}

func TestDataSourceIBMIamTrustedProfileIdentitiesProfileIdentityResponseToMap(t *testing.T) {
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

	result, err := iamidentity.DataSourceIBMIamTrustedProfileIdentitiesProfileIdentityResponseToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
