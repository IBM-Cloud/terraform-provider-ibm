// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
	return fmt.Sprintf(`
		data "ibm_iam_trusted_profile_claim_rules" "iam_trusted_profile_claim_rules" {
			profile_id = "%s"
		}
	`, acc.IAMTrustedProfileID)
}
