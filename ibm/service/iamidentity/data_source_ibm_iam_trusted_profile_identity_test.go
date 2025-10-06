// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIamTrustedProfileIdentityDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileIdentityDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "identity_type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "identifier_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_identity.iam_trusted_profile_identity", "type"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileIdentityDataSourceConfigBasic() string {
	profileID := acc.IAMTrustedProfileID
	identifier := acc.Ibmid1
	accountId := acc.IAMAccountId

	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_identity" "profile_identity" {
			profile_id = "%s"
			identity_type = "user"
			identifier = "%s"
			type = "user"
			accounts = [
				"%s"
			]
		}
		
		data "ibm_iam_trusted_profile_identity" "iam_trusted_profile_identity" {
			profile_id = ibm_iam_trusted_profile_identity.profile_identity.profile_id
			identity_type = ibm_iam_trusted_profile_identity.profile_identity.identity_type
			identifier_id = ibm_iam_trusted_profile_identity.profile_identity.identifier
		}
	`, profileID, identifier, accountId)
}
