// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMUserSettings_Basic(t *testing.T) {
	t.Skip()
	var allowedIP string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserSettingsBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserSettingsExists("ibm_iam_user_settings.user_settings", allowedIP),
					resource.TestCheckResourceAttr("ibm_iam_user_settings.user_settings", "allowed_ip_addresses.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMUserSettingsUpdate(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_settings.user_settings", "allowed_ip_addresses.#", "4"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserSettingsDestroy(s *terraform.State) error {

	userManagement, err := acc.TestAccProvider.Meta().(conns.ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	client := userManagement.UserInvite()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_user_settings" {
			continue
		}

		usermail := rs.Primary.ID

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		accountID := userDetails.UserAccount

		iamID, err := flex.GetIBMUniqueId(accountID, usermail, acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		UserSetting, UserSettingError := client.GetUserSettings(accountID, iamID)
		if UserSettingError == nil && UserSetting.AllowedIPAddresses != "" {
			return fmt.Errorf("Allowed IP setting still exists: %s", usermail)
		}
	}

	return nil
}

func testAccCheckIBMIAMUserSettingsExists(n string, ip string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No ID is set")
		}

		userManagement, err := acc.TestAccProvider.Meta().(conns.ClientSession).UserManagementAPI()
		if err != nil {
			return err
		}

		client := userManagement.UserInvite()

		usermail := rs.Primary.ID

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		accountID := userDetails.UserAccount

		iamID, err := flex.GetIBMUniqueId(accountID, usermail, acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		UserSetting, UserSettingError := client.GetUserSettings(accountID, iamID)
		if UserSettingError != nil {
			return fmt.Errorf("ERROR in getting user settings: %s", rs.Primary.ID)
		}

		ip = UserSetting.AllowedIPAddresses
		return nil
	}
}

func testAccCheckIBMIAMUserSettingsBasic() string {
	return fmt.Sprintf(`

		  
	resource "ibm_iam_user_settings" "user_settings" {
		iam_id = "%s"
		allowed_ip_addresses = ["192.168.0.0","192.168.0.1"]
	  }

	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserSettingsUpdate() string {
	return fmt.Sprintf(`

		  
	resource "ibm_iam_user_settings" "user_settings" {
		iam_id = "%s"
		allowed_ip_addresses = ["192.168.0.2","192.168.0.3","192.168.0.4","192.168.0.5"]
	  }

	`, acc.IAMUser)
}
