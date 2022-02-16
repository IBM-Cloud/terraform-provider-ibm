// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/scc-go-sdk/v3/adminserviceapiv1"
)

func TestAccIbmSccAccountSettingsResourceBasic(t *testing.T) {
	var conf adminserviceapiv1.LocationID
	locationID := "us"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccAccountSettingsResourceConfigBasic(locationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccLocationIdExists("ibm_scc_account_settings.ibm_scc_account_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_account_settings.ibm_scc_account_settings_instance", "location_id", locationID),
				),
			},
		},
	})
}

func testAccCheckIbmSccAccountSettingsResourceConfigBasic(locationID string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_account_settings" "ibm_scc_account_settings_instance" {
			location_id = "%s"
		}
	`, locationID)
}

func testAccCheckIBMSccLocationIdExists(n string, obj adminserviceapiv1.LocationID) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		adminServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AdminServiceApiV1()
		if err != nil {
			return err
		}

		getAccountSettingsOption := &adminserviceapiv1.GetSettingsOptions{}

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}
		accountID := userDetails.UserAccount

		getAccountSettingsOption.SetAccountID(accountID)

		accountSettings, _, err := adminServiceClient.GetSettings(getAccountSettingsOption)
		if err != nil {
			return err
		}

		obj = *accountSettings.Location
		return nil
	}
}
