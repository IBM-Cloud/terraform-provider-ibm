// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.98.0-8be2046a-20241205-162752
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIamTrustedProfileIdentitiesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileIdentitiesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identities.iam_trusted_profile_identities", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identities.iam_trusted_profile_identities", "profile_id"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileIdentitiesDataSourceConfigBasic() string {
	profileID := acc.IAMTrustedProfileID
	return fmt.Sprintf(`
		data "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities" {
			profile_id = "%s"
		}
	`, profileID)
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
