// Copyright IBM Corp. 2022 All Rights Reserved.
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

func TestAccIbmSccAccountSettingsBasic(t *testing.T) {
	var conf adminserviceapiv1.AccountSettings

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccAccountSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccAccountSettingsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccAccountSettingsExists("ibm_scc_account_settings.scc_account_settings", conf),
					resource.TestCheckResourceAttr(
						"ibm_scc_account_settings.scc_account_settings",
						"location.0.location_id",
						"us",
					),
					resource.TestCheckResourceAttr(
						"ibm_scc_account_settings.scc_account_settings",
						"event_notifications.#",
						"1",
					),
				),
				// ExpectNonEmptyPlan: true,
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_account_settings.scc_account_settings",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccAccountSettingsConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_scc_account_settings" "scc_account_settings" {
            location {
                location_id = "us"
            }
            event_notifications {
            }
        }
	`)
}

func testAccCheckIbmSccAccountSettingsExists(n string, obj adminserviceapiv1.AccountSettings) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		// rs, ok := s.RootModule().Resources[n]
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		adminServiceApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AdminServiceApiV1()
		if err != nil {
			return err
		}

		getSettingsOptions := &adminserviceapiv1.GetSettingsOptions{}

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		getSettingsOptions.SetAccountID(userDetails.UserAccount)

		accountSettings, _, err := adminServiceApiClient.GetSettings(getSettingsOptions)
		if err != nil {
			return err
		}

		obj = *accountSettings
		return nil
	}
}

func testAccCheckIbmSccAccountSettingsDestroy(s *terraform.State) error {
	adminServiceApiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AdminServiceApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_account_settings" {
			continue
		}

		getSettingsOptions := &adminserviceapiv1.GetSettingsOptions{}

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}
		getSettingsOptions.SetAccountID(userDetails.UserAccount)

		// Try to find the key
		_, response, err := adminServiceApiClient.GetSettings(getSettingsOptions)
		if response.StatusCode == 404 {
			return fmt.Errorf("Error checking for scc_account_settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

	}

	return nil
}
