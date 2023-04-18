// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIamTrustedProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles", "profiles.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`

		data "ibm_iam_trusted_profiles" "iam_trusted_profiles" {
			account_id = "%s"
			name = "name"
		}
	`, acc.IAMAccountId)
}
