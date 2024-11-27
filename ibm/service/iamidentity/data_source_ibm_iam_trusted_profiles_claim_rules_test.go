// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIamTrustedProfilesClaimRulesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfilesClaimRulesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rules.iam_trusted_profile_claim_rules", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rules.iam_trusted_profile_claim_rules", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rules.iam_trusted_profile_claim_rules", "rules.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfilesClaimRulesDataSourceConfigBasic() string {
	return `
		data "ibm_iam_trusted_profile_claim_rules" "iam_trusted_profile_claim_rules" {
			profile_id = "profile_id"
		}
	`
}
