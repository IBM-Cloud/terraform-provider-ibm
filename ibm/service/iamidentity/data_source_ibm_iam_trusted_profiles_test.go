// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.98.0-8be2046a-20241205-162752
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIamTrustedProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles_instance", "profiles.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_trusted_profiles" "iam_trusted_profiles_instance" {
			account_id = "%s"
			include_history = true
		}
	`, acc.AccountId)
}
