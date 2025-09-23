// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamaccessgroup_test

import (
	"fmt"
	iamaccessgroups "github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMAccessGroupAccountSettingsBasic(t *testing.T) {
	var conf iamaccessgroups.AccountSettings

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIamAccessGroupAccountSettingsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_access_group_account_settings.iam_access_group_account_settings", conf),
				),
			},
			{
				Config: testAccCheckIbmIamAccessGroupAccountSettingsConfigBasic(),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupAccountSettingsUpdates(t *testing.T) {
	var conf iamaccessgroups.AccountSettings
	publicAccessEnabled := "false"
	publicAccessEnabledUpdated := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIamAccessGroupAccountSettingsConfig(publicAccessEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamAccountSettingsExists("ibm_iam_access_group_account_settings.iam_access_group_account_settings", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_account_settings.iam_access_group_account_settings", "public_access_enabled", publicAccessEnabled),
				),
			},
			{
				Config: testAccCheckIbmIamAccessGroupAccountSettingsConfig(publicAccessEnabledUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group_account_settings.iam_access_group_account_settings", "public_access_enabled", publicAccessEnabledUpdated),
				),
			},
			{
				ResourceName:      "ibm_iam_access_group_account_settings.iam_access_group_account_settings",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func testAccCheckIbmIamAccessGroupAccountSettingsConfigBasic() string {
	return `

		resource "ibm_iam_access_group_account_settings" "iam_access_group_account_settings" {
			public_access_enabled = true
		}
	`
}

func testAccCheckIbmIamAccessGroupAccountSettingsConfig(publicAccessEnabled string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group_account_settings" "iam_access_group_account_settings" {
			public_access_enabled = %s
		}
	`, publicAccessEnabled)
}

func testAccCheckIbmIamAccountSettingsExists(n string, obj iamaccessgroups.AccountSettings) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamAccessGroupsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
		if err != nil {
			return err
		}

		getAccountSettingsOptions := iamAccessGroupsClient.NewGetAccountSettingsOptions(rs.Primary.ID)

		accountSettingsResponse, _, err := iamAccessGroupsClient.GetAccountSettings(getAccountSettingsOptions)
		if err != nil {
			return err
		}

		obj = *accountSettingsResponse

		return nil
	}
}
