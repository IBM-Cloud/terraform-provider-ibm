// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

// ---------------------------------------------------------------------------
// Bulk / load tests for ibm_iam_user_invite
//
// These tests are opt-in: they only run when IBM_IAM_INVITE_USERS_LIST is set
// to a comma-separated list of IBM Cloud email addresses that are NOT yet
// members of the account under test.
//
// Example:
//   IBM_IAM_INVITE_USERS_LIST="u1@corp.com,u2@corp.com,...,u10@corp.com" \
//   IBM_IAM_ACCESS_GROUP_ID="abc-123" \
//   go test ./ibm/service/iampolicy/ -run TestAccIBMIAMUserInvite_Bulk -v
// ---------------------------------------------------------------------------

// bulkUserList parses IBM_IAM_INVITE_USERS_LIST and returns the emails as a
// slice, trimming whitespace from individual entries.
func bulkUserList() []string {
	raw := os.Getenv("IBM_IAM_INVITE_USERS_LIST")
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	emails := make([]string, 0, len(parts))
	for _, p := range parts {
		if e := strings.TrimSpace(p); e != "" {
			emails = append(emails, e)
		}
	}
	return emails
}

// quotedList renders a Go string slice as a Terraform list literal,
// e.g. ["a@x.com", "b@x.com"].
func quotedList(emails []string) string {
	quoted := make([]string, len(emails))
	for i, e := range emails {
		quoted[i] = fmt.Sprintf("%q", e)
	}
	return "[" + strings.Join(quoted, ", ") + "]"
}

// TestAccIBMIAMUserInvite_BulkSmall invites 5 users in a single call.
// This is the minimum bulk load test: it validates that the async polling loop
// handles multiple concurrent PROCESSING → PENDING transitions without
// returning prematurely or timing out.
func TestAccIBMIAMUserInvite_BulkSmall(t *testing.T) {
	emails := bulkUserList()
	if len(emails) < 5 {
		t.Skip("IBM_IAM_INVITE_USERS_LIST must contain at least 5 email addresses to run BulkSmall")
	}
	batch := emails[:5]

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIAMUserInviteBulk(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserInviteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserInviteBulk(batch),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserInviteExists("ibm_iam_user_invite.bulk_invite"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_invite.bulk_invite", "users.#",
						strconv.Itoa(len(batch)),
					),
				),
			},
		},
	})
}

// TestAccIBMIAMUserInvite_BulkLarge invites 10 users in a single call.
// This is the primary load test that exercises the async polling logic under
// higher concurrency, verifying all users eventually reach PENDING within
// the 5-minute retry window.
func TestAccIBMIAMUserInvite_BulkLarge(t *testing.T) {
	emails := bulkUserList()
	if len(emails) < 10 {
		t.Skip("IBM_IAM_INVITE_USERS_LIST must contain at least 10 email addresses to run BulkLarge")
	}
	batch := emails[:10]

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIAMUserInviteBulk(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserInviteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserInviteBulk(batch),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserInviteExists("ibm_iam_user_invite.bulk_invite"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_invite.bulk_invite", "users.#",
						strconv.Itoa(len(batch)),
					),
				),
			},
		},
	})
}

// TestAccIBMIAMUserInvite_BulkIncrementalGrowth validates that expanding the
// users list across two apply steps does not regress:
//   - Step 1: invites the first 3 users and waits for them to settle.
//   - Step 2: expands to 6 users, exercising the update path of the resource
//     and the polling loop's handling of a mixed-state list (some users
//     already PENDING/ACTIVE from step 1, others freshly PROCESSING).
func TestAccIBMIAMUserInvite_BulkIncrementalGrowth(t *testing.T) {
	emails := bulkUserList()
	if len(emails) < 6 {
		t.Skip("IBM_IAM_INVITE_USERS_LIST must contain at least 6 email addresses to run BulkIncrementalGrowth")
	}
	initialBatch := emails[:3]
	expandedBatch := emails[:6]

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckIAMUserInviteBulk(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserInviteDestroy,
		Steps: []resource.TestStep{
			// Step 1: invite 3 users; confirm all settle before expanding.
			{
				Config: testAccCheckIBMIAMUserInviteBulk(initialBatch),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserInviteExists("ibm_iam_user_invite.bulk_invite"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_invite.bulk_invite", "users.#",
						strconv.Itoa(len(initialBatch)),
					),
				),
			},
			// Step 2: expand to 6 users. The polling loop must correctly
			// categorise the original 3 (now PENDING or ACTIVE) alongside the
			// new 3 (PROCESSING or missing) and wait only for the new ones.
			{
				Config: testAccCheckIBMIAMUserInviteBulk(expandedBatch),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserInviteExists("ibm_iam_user_invite.bulk_invite"),
					resource.TestCheckResourceAttr(
						"ibm_iam_user_invite.bulk_invite", "users.#",
						strconv.Itoa(len(expandedBatch)),
					),
				),
			},
		},
	})
}

// testAccCheckIBMIAMUserInviteBulk renders a Terraform config that invites
// the given batch of users into the configured access group.
func testAccCheckIBMIAMUserInviteBulk(emails []string) string {
	return fmt.Sprintf(`
resource "ibm_iam_user_invite" "bulk_invite" {
  users         = %s
  access_groups = ["%s"]
}
`, quotedList(emails), acc.IAMAccessGroupId)
}
