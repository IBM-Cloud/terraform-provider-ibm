// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
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
