// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIamUserMfaEnrollmentsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamUserMfaEnrollmentsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_user_mfa_enrollments.iam_user_mfa_enrollments", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_user_mfa_enrollments.iam_user_mfa_enrollments", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_user_mfa_enrollments.iam_user_mfa_enrollments", "iam_id"),
				),
			},
		},
	})
}

func testAccCheckIBMIamUserMfaEnrollmentsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_user_mfa_enrollments" "iam_user_mfa_enrollments" {
			account_id = "%s"
			iam_id = "%s"
		}
	`, acc.IAMAccountId, acc.Ibmid1)
}
