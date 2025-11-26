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
)

func TestAccIBMIAMTrustedProfileDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "account_id"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_trusted_profile" "iam_trusted_profile_instance" {
			profile_id = "%s"
		}
	`, acc.IAMTrustedProfileID)
}
