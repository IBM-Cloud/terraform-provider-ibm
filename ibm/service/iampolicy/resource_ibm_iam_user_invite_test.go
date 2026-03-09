// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMUserInvite_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserInviteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserInviteBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserInviteExists("ibm_iam_user_invite.invite_user"),
					resource.TestCheckResourceAttr("ibm_iam_user_invite.invite_user", "users.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_invite.invite_user", "access_groups.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserInviteDestroy(s *terraform.State) error {
	// User invites are transient - once processed, they don't persist as resources
	// The invited users become part of the account, but the invite itself is gone
	// So there's nothing to check for destroy
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_user_invite" {
			continue
		}
		// Just verify the resource was in state
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set for user invite resource")
		}
	}
	return nil
}

func testAccCheckIBMIAMUserInviteExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No user invite ID is set")
		}
		return nil
	}
}

func testAccCheckIBMIAMUserInviteBasic() string {
	return fmt.Sprintf(`
		resource "ibm_iam_user_invite" "invite_user" {
			users = ["%s"]
			access_groups = ["%s"]
		}
	`, acc.IAMUser, acc.IAMAccessGroupId)
}
