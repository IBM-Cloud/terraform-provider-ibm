// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package accountmanagement_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Changes account-wide settings, so it only runs when IBM_ACCOUNT_TRAITS_TEST is
// set to "true" on an account where that is acceptable.
func TestAccIBMAccountTraitsBasic(t *testing.T) {
	if os.Getenv("IBM_ACCOUNT_TRAITS_TEST") != "true" {
		t.Skip("set IBM_ACCOUNT_TRAITS_TEST=true to run; it changes account-wide settings")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccAccountTraitsConfig(false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_account_traits.test", "account_id"),
					resource.TestCheckResourceAttr("ibm_account_traits.test", "team_directory_enabled", "false"),
					resource.TestCheckResourceAttrPair("data.ibm_account_traits.test", "team_directory_enabled", "ibm_account_traits.test", "team_directory_enabled"),
				),
			},
			{
				Config: testAccAccountTraitsConfig(true),
				Check:  resource.TestCheckResourceAttr("ibm_account_traits.test", "team_directory_enabled", "true"),
			},
			{
				ResourceName:            "ibm_account_traits.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"fs_ready", "team_directory_enabled"},
			},
		},
	})
}

func TestAccIBMAccountTraitsRequiresOneTrait(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      `resource "ibm_account_traits" "test" {}`,
				ExpectError: regexp.MustCompile("Set at least one of fs_ready or team_directory_enabled"),
			},
		},
	})
}

func testAccAccountTraitsConfig(teamDirectory bool) string {
	return fmt.Sprintf(`
		resource "ibm_account_traits" "test" {
			team_directory_enabled = %t
		}
		data "ibm_account_traits" "test" {
			account_id = ibm_account_traits.test.account_id
			depends_on = [ibm_account_traits.test]
		}
	`, teamDirectory)
}
